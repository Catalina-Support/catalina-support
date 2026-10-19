import {
  AfterViewInit,
  Component,
  ElementRef,
  type OnDestroy,
  OnInit,
  computed,
  effect,
  inject,
  input,
  signal,
  viewChild,
} from '@angular/core';
import { ActivatedRoute, Router, RouterLink } from '@angular/router';

import { TranslationService } from '../../core/i18n/translation.service';
import { BrandService } from '../../core/services/brand.service';
import { SessionService } from '../../core/services/session.service';
import { Aviso, type MensajeDePantalla } from '../../shared/components/aviso';
import { Boton } from '../../shared/components/boton';
import { Campo } from '../../shared/components/campo';
import { Conmutador, type OpcionConmutador } from '../../shared/components/conmutador';
import { Dialogo } from '../../shared/components/dialogo';
import { Estado } from '../../shared/components/estado';
import { claveDelError } from '../../shared/errores';
import { consultaDeAncho, esPantallaAncha } from '../../shared/pantalla';
import {
  etiquetaDeEstado,
  fechaHora,
  interpolar,
  nombreDe,
  normalizarEtiqueta,
  opcionesDeEstado,
  textoDelResumen,
} from './etiquetas';
import {
  FILTROS_VACIOS,
  TIPO_INTERNO,
  TIPO_PRINCIPAL,
  TicketsService,
  type Categoria,
  type FiltrosDeBandeja,
  type PaginaDeTickets,
  type Prioridad,
  type Ticket,
} from './tickets.service';

/** Anchos del contenido, no de la ventana: el menú lateral también consume espacio. */
const ANCHO_DE_TABLA = 768;
const ANCHO_DE_TABLA_COMPLETA = 1240;

/**
 * Cuánto se espera entre relecturas y cuántas se hacen como mucho mientras el motor siga pendiente.
 *
 * El motor tarda, medido, **entre 12 y 24 segundos por campo**, y más si tiene cola (`docs/modules/ai.md`):
 * con ocho intentos de ocho segundos se le da algo más de un minuto, el mismo margen que usa la ficha.
 * **No es un sondeo eterno**: al llegar al tope se para (`docs/interfaz-y-experiencia.md`, enmienda del
 * 2026-09-29).
 */
const ESPERA_DEL_MOTOR = 8000;
const INTENTOS_DEL_MOTOR = 8;

interface EstadoDeNavegacion {
  readonly filtros: FiltrosDeBandeja;
  readonly pagina: PaginaDeTickets;
  readonly scrollY: number;
}

/** Memoria efímera: desaparece al recargar la pestaña porque nunca sale del módulo. */
const ESTADOS_DE_NAVEGACION = new Map<string, EstadoDeNavegacion>();

/** Lo que se espera antes de buscar: escribir un número entero es una búsqueda, no seis. */

/**
 * Las listas de tickets: **una sola pantalla con cuatro usos**.
 *
 * Es la misma tabla con distinto contenido, como manda `docs/modules/tickets.md`, sección 8:
 *
 * - **`mios`** (`/tickets`): **lo mío** —lo asignado o lo observado actualmente—. Soporte reúne
 *   principales y Desarrollo internos; la capacidad Admin no cambia el nivel operativo.
 * - **`principal`** (`/tickets/main`) y **`interno`** (`/tickets/internal`): **todo lo que hay**, para
 *   Soporte y Desarrollo, con el tipo fijo por la pantalla misma.
 * - **El Administrador** usa `/tickets` como siempre: los dos tipos, con su chip y **de sólo
 *   lectura** (`docs/interfaz-y-experiencia.md`, sección 4.4).
 *
 * Lo que lleva más tiempo esperando va primero, y lo cerrado no estorba: el orden lo decide el
 * backend, que es quien tiene los datos (sección 3.3).
 */
@Component({
  selector: 'app-bandeja-page',
  imports: [RouterLink, Aviso, Boton, Campo, Conmutador, Dialogo, Estado],
  templateUrl: './bandeja-page.html',
})
export class BandejaPage implements OnInit, AfterViewInit, OnDestroy {
  private readonly tickets = inject(TicketsService);
  private readonly textos = inject(TranslationService);
  private readonly marca = inject(BrandService);
  private readonly sesion = inject(SessionService);
  private readonly router = inject(Router);
  private readonly route = inject(ActivatedRoute);

  protected readonly pagina = signal<PaginaDeTickets>({
    tickets: [],
    total: 0,
    page: 1,
    perPage: 0,
  });

  protected readonly filtros = signal<FiltrosDeBandeja>(FILTROS_VACIOS);

  /**
   * Si el modal de búsqueda está abierto.
   *
   * **Los criterios viven en un modal** (decisión del responsable, 2026-09-27): eran tres filas de
   * filtros para cuatro criterios —y la de categorías se salía de la pantalla cuando el catálogo
   * crecía—, y así la lista se ve limpia y los criterios caben todos, los de hoy y los de mañana.
   */
  protected readonly buscando = signal(false);

  /**
   * Lo que se está escribiendo en el modal, **antes de aplicarlo**.
   *
   * Se edita aquí y no en los filtros de verdad: dentro de un modal, filtrar en vivo es confuso —la
   * lista de detrás iría cambiando mientras se rellenan cinco campos—, así que lo aplica el botón
   * «Buscar». Al abrir el modal, el borrador arranca con lo que ya está filtrado.
   */
  protected readonly borrador = signal<FiltrosDeBandeja>(FILTROS_VACIOS);

  /** El catálogo, para el chip de categoría. Lo ven todos los que han entrado. */
  protected readonly categorias = signal<readonly Categoria[]>([]);
  protected readonly cargando = signal(true);
  protected readonly cargandoMas = signal(false);
  protected readonly errorCargaMas = signal<string | null>(null);
  protected readonly anuncioCarga = signal('');
  protected readonly mensaje = signal<MensajeDePantalla | null>(null);
  protected readonly tooltipAbierto = signal<string | null>(null);

  /** Las relecturas que lleva hechas la vigilancia del motor y la que tiene programada, si hay alguna. */
  private intentosDelMotor = 0;
  private temporizadorDelMotor: ReturnType<typeof setTimeout> | null = null;

  /**
   * Cuál de las cuatro listas es. Lo trae la ruta (`data.listado`), así que la pantalla es la misma
   * para todas y lo que cambia es esto.
   */
  readonly listado = input<'mios' | 'principal' | 'interno'>('mios');

  private readonly contenedor = viewChild<ElementRef<HTMLElement>>('contenedor');
  private readonly contenedorLista = viewChild<ElementRef<HTMLElement>>('contenedorLista');
  private readonly marcadorDeCarga = viewChild<ElementRef<HTMLElement>>('marcadorDeCarga');
  private observadorDeAncho: ResizeObserver | null = null;
  private observadorDeCarga: IntersectionObserver | null = null;
  /** Invalida respuestas tardías cuando la persona cambia los filtros mientras llega una página. */
  private versionDeConsulta = 0;
  private readonly consultaDeTabla = consultaDeAncho(ANCHO_DE_TABLA);
  private readonly alCambiarAnchoDeTabla = (evento: MediaQueryListEvent) => {
    this.esAncha.set(evento.matches);
    if (evento.matches) setTimeout(() => this.llevarPaginaAlInicio(), 50);
  };

  /** Desde 768 px existe la tabla; por debajo sólo existen las tarjetas. */
  protected readonly esAncha = signal(esPantallaAncha(ANCHO_DE_TABLA));
  protected readonly muestraPersonas = signal(true);

  protected readonly papel = computed(() => this.sesion.usuario()?.role ?? '');
  protected readonly esUsuario = computed(() => this.papel() === 'usuario');
  protected readonly esSoporte = computed(() => this.papel() === 'soporte');
  protected readonly esDesarrollo = computed(() => this.papel() === 'desarrollo');
  protected readonly esAdministradorPuro = computed(() => this.papel() === 'administrador');

  /**
   * Quién puede crear un ticket desde aquí: el usuario, Soporte y Desarrollo, **y sólo en la
   * bandeja** —«Mis tickets»—, que es donde vive el botón desde que salió del menú
   * (decisión 53). Las dos listas del «todo» son para mirar y trabajar lo que hay, no para crear.
   */
  protected readonly puedeCrear = computed(
    () =>
      !this.esListaDelTodo() &&
      !this.esAdministradorPuro() &&
      (this.esUsuario() || this.esSoporte() || this.esDesarrollo()),
  );

  /** Si la bandeja enseña internos, que es lo que decide los chips de estado y el nombre de las cosas. */
  protected readonly mirandoInternos = computed(() => this.filtros().type === TIPO_INTERNO);

  protected readonly hayMas = computed(() => this.pagina().tickets.length < this.pagina().total);

  /** Si esta lista lleva el chip de tipo: «Mis tickets» de Soporte y Desarrollo, y la del Administrador. */
  /**
   * El chip de vista lo lleva **«Mis tickets» para Soporte y Desarrollo**: es donde se mezclan los
   * asignados, los que uno abrió, donde ha comentado y los que observa (decisión 62). El usuario no lo
   * necesita —todo lo suyo es suyo— y el Administrador mira la bandeja entera.
   */
  protected readonly conChipDeVista = computed(
    () => this.listado() === 'mios' && (this.esSoporte() || this.esDesarrollo()),
  );

  /**
   * «Mis tickets» abre combinando asignados y observados; los otros listados no tienen esta vista.
   */
  private vistaInicial(): { view?: string } {
    return this.conChipDeVista() ? { view: '' } : {};
  }

  /** Las tres vistas visibles en la cabecera, con la combinación como punto de entrada. */
  protected opcionesDeVista(): readonly OpcionConmutador[] {
    return [
      { valor: '', etiqueta: this.t().tickets.vistaTodos },
      { valor: 'assigned', etiqueta: this.t().tickets.vistaAsignados },
      { valor: 'watching', etiqueta: this.t().tickets.vistaObservo },
    ];
  }

  protected readonly conChipDeTipo = computed(
    () =>
      this.esAdministradorPuro(),
  );

  /** Si la lista enseña **todo** (las dos del «todo»), que es lo que decide el botón de crear. */
  protected readonly esListaDelTodo = computed(
    () => this.listado() !== 'mios' && !this.esAdministradorPuro(),
  );

  constructor() {
    effect(() => {
      const marcador = this.marcadorDeCarga()?.nativeElement;
      const hayMas = this.hayMas();
      const raiz = this.esAncha() ? (this.contenedorLista()?.nativeElement ?? null) : null;
      this.observadorDeCarga?.disconnect();
      this.observadorDeCarga = null;

      if (!marcador || !hayMas || typeof IntersectionObserver === 'undefined') return;
      this.observadorDeCarga = new IntersectionObserver(
        ([entrada]) => {
          if (entrada.isIntersecting) void this.cargarMasTickets();
        },
        { root: raiz, rootMargin: '400px 0px' },
      );
      this.observadorDeCarga.observe(marcador);
    });
  }

  ngAfterViewInit(): void {
    this.consultaDeTabla?.addEventListener('change', this.alCambiarAnchoDeTabla);
    if (this.esAncha()) setTimeout(() => this.llevarPaginaAlInicio(), 50);
    const elemento = this.contenedor()?.nativeElement;
    if (!elemento || typeof ResizeObserver === 'undefined') {
      return;
    }

    const actualizar = (ancho: number) => {
      // JSDOM no calcula geometría: en las pruebas se conserva la tabla completa por defecto.
      if (ancho <= 0) return;
      this.muestraPersonas.set(ancho >= ANCHO_DE_TABLA_COMPLETA);
    };

    actualizar(elemento.getBoundingClientRect().width);
    this.observadorDeAncho = new ResizeObserver(([entrada]) =>
      actualizar(entrada.contentRect.width),
    );
    this.observadorDeAncho.observe(elemento);
  }

  /**
   * La lista arranca aquí y no en el constructor: **cuál de las cuatro es lo trae la ruta**, y el
   * enrutador pone ese dato justo antes de `ngOnInit`. En el constructor la lista sería siempre «lo
   * mío», y una de las dos del «todo» pediría el tipo equivocado.
   */
  ngOnInit(): void {
    // El tipo de ticket: el que fija la lista cuando es una del «todo», y si no el de siempre —el
    // usuario y Soporte, los principales; Desarrollo, los internos; el Administrador, los principales,
    // que puede cambiar con su chip—.
    const consulta = this.route.snapshot.queryParamMap;
    const priorities = consulta.getAll('priority').filter(esPrioridad);
    const view =
      consulta.get('view') === 'all' ? '' : (consulta.get('view') ?? this.vistaInicial().view);
    this.filtros.set({
      ...FILTROS_VACIOS,
      ...this.vistaInicial(),
      view,
      type: consulta.get('type') === 'all' ? '' : (consulta.get('type') ?? this.tipoInicial(view)),
      state: consulta.get('state') ?? '',
      q: consulta.get('q') ?? '',
      category: consulta.get('category') ?? '',
      tag: consulta.get('tag') ?? '',
      priorities,
      order: consulta.get('order') === 'updated' ? 'updated' : 'priority',
      page: 1,
    });
    const guardado = ESTADOS_DE_NAVEGACION.get(this.claveDeNavegacion());
    if (guardado) {
      ESTADOS_DE_NAVEGACION.delete(this.claveDeNavegacion());
      this.filtros.set({ ...guardado.filtros });
      this.pagina.set(guardado.pagina);
      this.cargando.set(false);
      this.vigilarElMotor();
      // El enrutador restaura primero su propia posición (arriba); se aplica la memoria de esta
      // bandeja justo después para que no la sobrescriba al terminar la navegación.
      setTimeout(() => {
        if (this.esAncha()) {
          const lista = this.contenedorLista()?.nativeElement;
          if (lista) lista.scrollTop = guardado.scrollY;
        } else {
          document.documentElement.scrollTop = guardado.scrollY;
          document.body.scrollTop = guardado.scrollY;
        }
      }, 50);
    } else {
      void this.cargar();
    }
    void this.cargarCategorias();
  }

  protected t() {
    return this.textos.textos();
  }

  protected evaluarTooltip(clave: string, contenido: HTMLElement): void {
    const truncado =
      contenido.scrollHeight > contenido.clientHeight + 1 ||
      contenido.scrollWidth > contenido.clientWidth + 1;
    this.tooltipAbierto.set(truncado ? clave : null);
  }

  protected cerrarTooltip(clave: string): void {
    if (this.tooltipAbierto() === clave) this.tooltipAbierto.set(null);
  }

  protected tooltipVisible(clave: string): boolean {
    return this.tooltipAbierto() === clave;
  }

  private llevarPaginaAlInicio(): void {
    document.documentElement.scrollTop = 0;
    document.body.scrollTop = 0;
  }

  /** El título de la pantalla, que es como se llama esta lista. */
  protected titulo(): string {
    if (this.esAdministradorPuro()) {
      return this.t().menu.bandeja;
    }
    if (this.listado() === 'principal') {
      return this.t().menu.ticketsPrincipales;
    }
    if (this.listado() === 'interno') {
      return this.t().menu.ticketsInternos;
    }

    return this.t().menu.misTickets;
  }

  /** Los chips del estado: los del tipo de ticket que se está mirando. */
  protected opcionesDeEstado(): readonly OpcionConmutador[] {
    return opcionesDeEstado(this.t(), this.mirandoInternos(), this.esUsuario());
  }

  /** El chip de tipo: lo lleva quien puede tener de los dos —Soporte, Desarrollo y el Administrador—. */
  protected opcionesDeTipo(): readonly OpcionConmutador[] {
    return [
      { valor: TIPO_PRINCIPAL, etiqueta: this.t().tickets.tipoPrincipal },
      { valor: TIPO_INTERNO, etiqueta: this.t().tickets.tipoInterno },
      { valor: '', etiqueta: this.t().tickets.tipoTodos },
    ];
  }

  /**
   * El chip de categoría: todas y las del catálogo.
   *
   * Se enseña **tal y como lo devuelve el backend** —activas para todos, y las retiradas para quien
   * mantiene el catálogo—: una categoría retirada sigue clasificando tickets, así que poder filtrar por
   * ella es lo que hace falta para encontrarlos.
   */
  protected opcionesDeCategoria(): readonly OpcionConmutador[] {
    return [
      { valor: '', etiqueta: this.t().tickets.todasLasCategorias },
      ...this.categorias().map((categoria) => ({
        valor: String(categoria.id),
        etiqueta: categoria.name,
      })),
    ];
  }

  /** El estado de un ticket, como lo lee quien mira. */
  protected estado(ticket: Ticket): string {
    return etiquetaDeEstado(this.t(), ticket.state, this.esUsuario());
  }

  protected solicitante(ticket: Ticket): string {
    return nombreDe(ticket.requester) || '—';
  }

  protected responsable(ticket: Ticket): string {
    return nombreDe(ticket.assignee) || '—';
  }

  /** Summary estable del planteamiento original. */
  protected summary(ticket: Ticket): string {
    return textoDelResumen(ticket.insights?.summary, this.textos.idioma(), this.t());
  }

  /** Y la última acción. */
  protected ultimaAccion(ticket: Ticket): string {
    return textoDelResumen(ticket.insights?.ultimaAccion, this.textos.idioma(), this.t());
  }

  protected actualizado(ticket: Ticket): string {
    return fechaHora(ticket.updatedAt, this.textos.idioma(), this.marca.zonaHoraria());
  }

  protected fechaActualizada(ticket: Ticket): string {
    return this.actualizado(ticket).split(',')[0]?.trim() || '—';
  }

  protected horaActualizada(ticket: Ticket): string {
    return this.actualizado(ticket).split(',').slice(1).join(',').trim();
  }

  protected etiquetasVisibles(ticket: Ticket): readonly string[] {
    return (ticket.tags ?? []).slice(0, 2);
  }

  protected etiquetasOcultas(ticket: Ticket): readonly string[] {
    return (ticket.tags ?? []).slice(2);
  }

  protected prioridad(ticket: Ticket): string {
    const p = ticket.priority ?? 'medium';
    return this.etiquetaDePrioridad(p);
  }

  protected iconoDePrioridad(prioridad: Prioridad): string {
    return { low: '↓', medium: '•', high: '↑', critical: '!' }[prioridad];
  }

  protected etiquetaDePrioridad(prioridad: Prioridad): string {
    const t = this.t().tickets;
    return {
      low: t.prioridadBaja,
      medium: t.prioridadMedia,
      high: t.prioridadAlta,
      critical: t.prioridadCritica,
    }[prioridad];
  }

  protected alternarPrioridad(prioridad: Prioridad, activa: boolean): void {
    const actuales = this.borrador().priorities ?? [];
    this.borrador.update((filtros) => ({
      ...filtros,
      priorities: activa
        ? [...actuales, prioridad]
        : actuales.filter((valor) => valor !== prioridad),
    }));
  }

  protected hayFiltros(): boolean {
    const filtros = this.filtros();
    return Boolean(
      filtros.state ||
      filtros.q ||
      filtros.category ||
      filtros.tag ||
      filtros.priorities?.length ||
      (this.conChipDeTipo() && filtros.type !== this.tipoInicial()),
    );
  }

  protected filtrar(clave: 'type' | 'state' | 'view' | 'category', valor: string): void {
    this.filtros.update((filtros) => ({
      ...filtros,
      [clave]: valor,
      // Al cambiar de vista, el tipo se reajusta: «Observo» mira los dos.
      ...(clave === 'view' ? { type: this.tipoInicial(valor) } : {}),
      page: 1,
    }));
    void this.reiniciarYCargar();
  }

  /** Abre el modal con lo que ya está filtrado, para corregirlo en vez de empezar de cero. */
  protected abrirBusqueda(): void {
    this.borrador.set({ ...this.filtros() });
    this.buscando.set(true);
  }

  /** Cierra el modal **sin aplicar**: lo que se haya escrito en él se descarta. */
  protected cerrarBusqueda(): void {
    this.buscando.set(false);
  }

  /** Aplica lo del modal: es el botón «Buscar». */
  protected aplicarBusqueda(): void {
    this.filtros.set({ ...this.borrador(), page: 1 });
    this.buscando.set(false);
    void this.reiniciarYCargar();
  }

  /** Lo que se escribe en el modal: **no filtra todavía**, sólo rellena el borrador. */
  protected escribirEnElBorrador(
    clave: 'q' | 'tag' | 'state' | 'type' | 'view' | 'category',
    valor: string,
  ): void {
    // La etiqueta se normaliza **al aplicarla**, que es como está guardada: minúsculas y guiones.
    const limpio = clave === 'tag' ? normalizarEtiqueta(valor) : valor;

    this.borrador.update((borrador) => ({
      ...borrador,
      [clave]: limpio,
      // **Al cambiar de vista, el tipo se reajusta**: mirar lo que uno observa cruza los dos tipos, y el
      // «Mis tickets» de Desarrollo arranca en internos —que es su trabajo—, así que «Observo» tiene que
      // soltar ese filtro o no vería lo que sigue del otro lado. Es la misma regla que tenía el chip
      // cuando estaba a la vista: al moverlo al modal se perdió, y lo cazó la prueba de las menciones.
      ...(clave === 'view' ? { type: this.tipoInicial(limpio) } : {}),
    }));
  }

  /** Cuántos criterios están puestos: es lo que dice el resumen de la cabecera. */
  protected cuantosFiltros(): number {
    return this.clavesDeLosFiltros().length;
  }

  /** El resumen de la cabecera: «3 filtros · Buscar, Estado, Categoría». */
  protected resumenDeFiltros(): string {
    const cuantos = this.cuantosFiltros();
    const etiqueta =
      cuantos === 1
        ? this.t().tickets.resumenUnFiltro
        : interpolar(this.t().tickets.resumenFiltros, { n: String(cuantos) });

    return `${etiqueta} · ${this.clavesDeLosFiltros().join(', ')}`;
  }

  /**
   * Los criterios que están puestos, **por su nombre**: es lo que enseña el resumen de la lista, y lo
   * que hace que se entienda por qué salen pocos tickets sin tener que abrir el modal.
   */
  protected clavesDeLosFiltros(): readonly string[] {
    const filtros = this.filtros();
    const puestos: string[] = [];

    if (filtros.q) {
      puestos.push(this.t().tickets.buscarEtiqueta);
    }
    if (this.conChipDeTipo() && filtros.type !== this.tipoInicial()) {
      puestos.push(this.t().tickets.filtroTipo);
    }
    if (filtros.state) {
      puestos.push(this.t().tickets.filtroEstado);
    }
    if (filtros.category) {
      puestos.push(this.t().tickets.filtroCategoria);
    }
    if (filtros.tag) {
      puestos.push(this.t().tickets.filtroEtiqueta);
    }
    if (filtros.priorities?.length) {
      puestos.push(this.t().tickets.filtroPrioridad);
    }

    return puestos;
  }

  /** Quita todos los filtros y recarga: es el botón del resumen, y el del modal. */
  protected limpiarFiltros(): void {
    this.borrador.set({ ...FILTROS_VACIOS, type: this.tipoInicial() });
    this.filtros.set({ ...FILTROS_VACIOS, type: this.tipoInicial() });
    this.buscando.set(false);
    void this.reiniciarYCargar();
  }

  protected reintentarCarga(): void {
    void this.cargarMasTickets();
  }

  protected conservarAlAbrir(): void {
    this.guardarEstadoDeNavegacion();
  }

  protected irANuevo(): void {
    void this.router.navigate(['/tickets/new']);
  }

  /**
   * El tipo queda fijado por la lista y por el papel operativo. En «Mis tickets», Soporte consulta
   * principales y Desarrollo internos; una capacidad Admin acumulativa no amplía ese nivel.
   */
  private tipoInicial(_vista = this.filtros().view): string {
    if (this.listado() === 'principal') {
      return TIPO_PRINCIPAL;
    }
    if (this.listado() === 'interno') {
      return TIPO_INTERNO;
    }
    if (this.esDesarrollo()) {
      return TIPO_INTERNO;
    }

    return TIPO_PRINCIPAL;
  }

  /** Los filtros de la lista, con «lo mío» cuando la lista es la bandeja. */
  private filtrosDeLaLista(): FiltrosDeBandeja {
    const filtros = this.filtros();

    // Sólo la identidad de fábrica carece de papel operativo. Admin acumulativo conserva la bandeja
    // personal de Soporte o Desarrollo (docs/usuarios-y-permisos.md, sección 16).
    return this.listado() === 'mios' && !this.esAdministradorPuro()
      ? { ...filtros, mine: true }
      : filtros;
  }

  /**
   * Trae el catálogo para el chip de categoría.
   *
   * Si no llega, la lista se sigue viendo y sólo falta un filtro: no es un fallo de la pantalla, así
   * que no se saca un aviso por esto.
   */
  private async cargarCategorias(): Promise<void> {
    try {
      this.categorias.set((await this.tickets.categorias()).categories);
    } catch {
      this.categorias.set([]);
    }
  }

  private async cargar(): Promise<void> {
    void this.sincronizarDireccion();
    await this.leerLaLista();
    this.vigilarElMotor();
  }

  private async reiniciarYCargar(): Promise<void> {
    this.versionDeConsulta += 1;
    this.filtros.update((filtros) => ({ ...filtros, page: 1 }));
    this.pagina.set({ tickets: [], total: 0, page: 1, perPage: 20 });
    this.cargandoMas.set(false);
    this.errorCargaMas.set(null);
    this.anuncioCarga.set('');
    if (this.esAncha()) {
      const lista = this.contenedorLista()?.nativeElement;
      if (lista) lista.scrollTop = 0;
    } else {
      this.llevarPaginaAlInicio();
    }
    await this.cargar();
  }

  private sincronizarDireccion(): Promise<boolean> {
    const f = this.filtros();
    return this.router.navigate([], {
      relativeTo: this.route,
      replaceUrl: true,
      queryParams: {
        type: this.conChipDeTipo() && f.type !== this.tipoInicial(f.view) ? f.type || 'all' : null,
        state: f.state || null,
        q: f.q || null,
        category: f.category || null,
        tag: f.tag || null,
        view: this.conChipDeVista() && f.view !== this.vistaInicial().view ? f.view || 'all' : null,
        priority: f.priorities?.length ? f.priorities : null,
        order: f.order === 'updated' ? 'updated' : null,
        page: null,
      },
    });
  }

  /** La petición de la lista, sin tocar la vigilancia del motor: la comparten la carga y la relectura. */
  private async leerLaLista(): Promise<void> {
    const version = this.versionDeConsulta;
    this.cargando.set(true);

    try {
      const pagina = await this.tickets.listar(this.filtrosDeLaLista());
      if (version !== this.versionDeConsulta) return;
      this.pagina.set(pagina);
    } catch (error) {
      if (version !== this.versionDeConsulta) return;
      this.pagina.set({ tickets: [], total: 0, page: 1, perPage: 0 });
      this.mensaje.set({ forma: 'error', texto: this.textos.error(claveDelError(error)) });
    } finally {
      if (version === this.versionDeConsulta) this.cargando.set(false);
    }
  }

  private async cargarMasTickets(): Promise<void> {
    if (this.cargando() || this.cargandoMas() || !this.hayMas()) return;

    this.cargandoMas.set(true);
    this.errorCargaMas.set(null);
    const version = this.versionDeConsulta;
    const paginaSiguiente = this.pagina().page + 1;

    try {
      const recibida = await this.tickets.listar({
        ...this.filtrosDeLaLista(),
        page: paginaSiguiente,
      });
      if (version !== this.versionDeConsulta) return;
      const conocidos = new Set(this.pagina().tickets.map((ticket) => ticket.number));
      const nuevos = recibida.tickets.filter((ticket) => !conocidos.has(ticket.number));
      this.pagina.update((actual) => ({
        ...recibida,
        tickets: [...actual.tickets, ...nuevos],
      }));
      this.anuncioCarga.set(
        interpolar(this.t().tickets.cargaMasAnuncio, { n: String(nuevos.length) }),
      );
      this.programarRelectura();
    } catch {
      if (version === this.versionDeConsulta) {
        this.errorCargaMas.set(this.t().tickets.cargaMasError);
      }
    } finally {
      if (version === this.versionDeConsulta) this.cargandoMas.set(false);
    }
  }

  /**
   * **El motor se refresca solo** (enmienda del 2026-09-29): después de cualquier acción que recargue
   * la lista, si algún ticket se queda con el motivo o la última acción en `pendiente`, la pantalla
   * **vuelve a leer sola** hasta que dejen de estarlo.
   *
   * Es el mismo mecanismo que la ficha (`ticket-vista.ts`, `esperarAlMotor`) adaptado a una lista con
   * varios tickets: **se repite cada `ESPERA_DEL_MOTOR`** y **se para al llegar a `INTENTOS_DEL_MOTOR`**.
   * No se comparte con la ficha porque aquel archivo es de otro compañero y no se puede tocar.
   */
  private vigilarElMotor(): void {
    // Cada carga de verdad —filtro, página o búsqueda— estrena el presupuesto entero.
    this.intentosDelMotor = 0;
    this.programarRelectura();
  }

  /** Programa la siguiente relectura, si aún queda presupuesto y algo sigue pendiente. */
  private programarRelectura(): void {
    if (this.temporizadorDelMotor !== null) {
      clearTimeout(this.temporizadorDelMotor);
      this.temporizadorDelMotor = null;
    }

    if (!this.hayResumenesPendientes() || this.intentosDelMotor >= INTENTOS_DEL_MOTOR) {
      return;
    }

    this.temporizadorDelMotor = setTimeout(() => {
      this.temporizadorDelMotor = null;
      this.intentosDelMotor += 1;
      void this.refrescarListaCargada().then(() => this.programarRelectura());
    }, ESPERA_DEL_MOTOR);
  }

  /** Si algún ticket de la página sigue con el motor pendiente en cualquiera de los dos campos. */
  private hayResumenesPendientes(): boolean {
    return this.pagina().tickets.some(
      (ticket) =>
        ticket.insights?.summary.state === 'pendiente' ||
        ticket.insights?.ultimaAccion.state === 'pendiente',
    );
  }

  private async refrescarListaCargada(): Promise<void> {
    const ultimaPagina = this.pagina().page;
    const paginas: PaginaDeTickets[] = [];
    try {
      for (let page = 1; page <= ultimaPagina; page += 1) {
        paginas.push(await this.tickets.listar({ ...this.filtrosDeLaLista(), page }));
      }
      const tickets = paginas.flatMap((pagina) => pagina.tickets);
      const ultima = paginas.at(-1);
      if (ultima) this.pagina.set({ ...ultima, tickets });
    } catch {
      // El refresco de IA no sustituye ni vacía lo que la persona ya está leyendo.
    }
  }

  private claveDeNavegacion(): string {
    const { page: _page, ...filtros } = this.filtros();
    return `${this.listado()}:${JSON.stringify(filtros)}`;
  }

  private guardarEstadoDeNavegacion(): void {
    const scrollY = this.esAncha()
      ? (this.contenedorLista()?.nativeElement.scrollTop ?? 0)
      : window.scrollY || document.documentElement.scrollTop;
    ESTADOS_DE_NAVEGACION.set(this.claveDeNavegacion(), {
      filtros: { ...this.filtros() },
      pagina: this.pagina(),
      scrollY,
    });
  }

  ngOnDestroy(): void {
    this.observadorDeCarga?.disconnect();
    this.consultaDeTabla?.removeEventListener('change', this.alCambiarAnchoDeTabla);
    this.observadorDeAncho?.disconnect();
    this.observadorDeAncho = null;
    if (this.temporizadorDelMotor !== null) {
      clearTimeout(this.temporizadorDelMotor);
      this.temporizadorDelMotor = null;
    }
  }
}

function esPrioridad(valor: string): valor is Prioridad {
  return ['low', 'medium', 'high', 'critical'].includes(valor);
}
