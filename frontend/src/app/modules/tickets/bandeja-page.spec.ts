import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';
import { type ComponentFixture, TestBed } from '@angular/core/testing';
import { type WritableSignal, signal } from '@angular/core';
import { provideRouter } from '@angular/router';

import { TranslationService } from '../../core/i18n/translation.service';
import { SessionService, type Usuario } from '../../core/services/session.service';
import { BandejaPage } from './bandeja-page';
import { type PaginaDeTickets, type Ticket } from './tickets.service';

/**
 * **La tabla de resumen de las bandejas** (decisión 86) y **el refresco de los dos resúmenes del motor**
 * (enmienda del 2026-09-29): el número va solo, la categoría y las etiquetas viven juntas en
 * «Clasificación», el estado sigue en la suya, y la lista **se vuelve a leer sola** mientras algo siga
 * pendiente, con un tope de intentos.
 *
 * El aspecto en un navegador lo cubre la capa de Playwright; aquí se prueba lo que la pantalla enseña y
 * lo que pide.
 */

const ANA = {
  id: 9,
  name: 'Ana',
  lastName: 'Ruiz',
  email: 'ana@demo.com',
  role: 'usuario',
};

/** Un ticket ya resumido: el motor no está pendiente, así que no se programa ninguna relectura. */
const CON_CLASIFICACION: Ticket = {
  number: 'CS-2026-0001',
  internal: false,
  subject: 'No va la red',
  description: '',
  state: 'nuevo',
  category: { id: 1, name: 'Red', active: true },
  tags: ['red-wifi', 'urgente', 'oficina-norte'],
  requester: ANA,
  insights: {
    summary: { state: 'listo', text: 'No va la red', language: 'es' },
    ultimaAccion: { state: 'listo', text: 'Se avisó', language: 'es' },
  },
  createdAt: '2026-09-28T08:00:00Z',
  updatedAt: '2026-09-28T08:00:00Z',
};

const PAGINA: PaginaDeTickets = { tickets: [CON_CLASIFICACION], total: 1, page: 1, perPage: 20 };

/** El mismo ticket, pero con el motor todavía trabajando: es lo que dispara la relectura. */
const PAGINA_PENDIENTE: PaginaDeTickets = {
  ...PAGINA,
  tickets: [
    {
      ...CON_CLASIFICACION,
      insights: {
        summary: { state: 'pendiente' },
        ultimaAccion: { state: 'pendiente' },
      },
    },
  ],
};

const SIN_CATEGORIAS = { categories: [] };

describe('BandejaPage', () => {
  let http: HttpTestingController;
  let usuario: WritableSignal<Usuario | null>;

  beforeEach(async () => {
    localStorage.clear();
    usuario = signal<Usuario | null>(null);

    await TestBed.configureTestingModule({
      imports: [BandejaPage],
      providers: [
        provideHttpClient(),
        provideHttpClientTesting(),
        provideRouter([]),
        // La sesión se sustituye por lo único que mira la pantalla: quién ha entrado y su papel.
        { provide: SessionService, useValue: { usuario, esAdministrador: signal(false) } },
      ],
    }).compileComponents();

    TestBed.inject(TranslationService).cambiar('es');
    http = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    http.verify();
    localStorage.clear();
  });

  /** Deja correr las promesas del servicio: el backend de prueba contesta en el acto. */
  function esperar(): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve));
  }

  /** Monta la pantalla con un usuario normal y responde la lista y el catálogo. */
  async function montar(pagina: PaginaDeTickets = PAGINA): Promise<ComponentFixture<BandejaPage>> {
    usuario.set({
      id: 1,
      name: 'María',
      lastName: 'Pérez',
      email: 'maria@demo.com',
      role: 'usuario',
      origin: 'local',
      factory: false,
    });

    const fixture = TestBed.createComponent(BandejaPage);
    fixture.detectChanges();

    http.expectOne((peticion) => peticion.url === '/api/tickets').flush(pagina);
    http.expectOne('/api/tickets/categories').flush(SIN_CATEGORIAS);
    await esperar();
    fixture.detectChanges();

    return fixture;
  }

  /** Las cabeceras de la tabla, en orden. */
  function encabezados(host: HTMLElement): (string | undefined)[] {
    return [...host.querySelectorAll('thead th')].map((celda) => celda.textContent?.trim());
  }

  /** La fila de la tabla. */
  function celdasDeLaFila(host: HTMLElement): HTMLElement[] {
    return [...host.querySelectorAll('tbody tr:first-child td')] as HTMLElement[];
  }

  it('trata Desarrollo con Admin como papel operativo y abre Mis tickets en Todos', async () => {
    usuario.set({
      id: 20,
      name: 'Desarrollador',
      lastName: 'Uno',
      email: 'dev1@demo.com',
      role: 'desarrollo',
      isAdmin: true,
      origin: 'local',
      factory: false,
    });

    const fixture = TestBed.createComponent(BandejaPage);
    fixture.detectChanges();

    const peticion = http.expectOne((solicitud) => solicitud.url === '/api/tickets');
    expect(peticion.request.params.get('type')).toBe('interno');
    expect(peticion.request.params.get('mine')).toBe('1');
    expect(peticion.request.params.has('view')).toBe(false);
    peticion.flush(PAGINA);
    http.expectOne('/api/tickets/categories').flush(SIN_CATEGORIAS);
    await esperar();
    fixture.detectChanges();

    const cabecera = (fixture.nativeElement as HTMLElement).querySelector('main > div');
    expect(cabecera?.textContent).toContain('Todos');
    expect(cabecera?.textContent).toContain('Asignados');
    expect(cabecera?.textContent).toContain('Observo');
    expect(cabecera?.textContent).not.toContain('1 filtro');
  });

  it('ordena toda la tabla según el contrato y deja la acción como columna visible', async () => {
    const fixture = await montar();
    const host = fixture.nativeElement as HTMLElement;

    expect(encabezados(host)).toEqual([
      'Ticket',
      'Prioridad',
      'Estado',
      'Asunto',
      'Clasificación',
      'Resumen',
      'Última acción',
      'Solicitante',
      'Responsable',
      'Fecha actualización',
      'Acción',
    ]);

    const celdas = celdasDeLaFila(host);

    // **El número solo**: ni la categoría ni las etiquetas van debajo.
    expect(celdas[0].textContent?.trim()).toBe('CS-2026-0001');
    expect(celdas[0].classList).toContain('sticky');
    expect(celdas[0].classList).toContain('left-0');
    expect(celdas[0].classList).toContain('pl-4');
    expect(celdas.at(-1)?.classList).toContain('sticky');
    expect(celdas.at(-1)?.classList).toContain('right-0');
    expect(celdas.at(-1)?.classList).toContain('pr-4');
    expect(host.querySelector('thead th:last-child')?.classList).toContain('w-[6rem]');
    expect(host.querySelector('main')?.classList).toContain('lg:-mx-6');
    expect(host.querySelector('thead th')?.classList).toContain('sticky');
    expect(host.querySelector('section')?.classList).toContain('md:overflow-auto');
    expect(host.querySelector('section')?.classList).toContain('md:min-h-60');
    expect(host.querySelector('section')?.classList).toContain('pb-4');
    expect(host.querySelector('section')?.getAttribute('tabindex')).toBe('0');
    expect(celdas[0].querySelector('span[aria-hidden="true"]')).not.toBeNull();
    expect(celdas[0].textContent).not.toContain('Red');
    expect(celdas[0].textContent).not.toContain('red-wifi');

    // **La clasificación los agrupa**: el chip de la categoría y el de cada etiqueta.
    const clasificacion = celdas[4];
    expect(clasificacion.textContent).toContain('Red');
    expect(clasificacion.textContent).toContain('red-wifi');
    expect(clasificacion.textContent).toContain('urgente');
    expect(clasificacion.textContent).toContain('+1');
    expect(clasificacion.querySelector('[aria-label="oficina-norte"]')).not.toBeNull();

    // **El estado sigue en su columna**, con su chip.
    const estado = celdas[2];
    expect(estado.querySelector('app-estado')).not.toBeNull();
    // Y no se ha colado en la clasificación.
    expect(clasificacion.querySelector('app-estado')).toBeNull();

    expect(celdas[10].textContent?.trim()).toBe('Abrir');
  });

  it('fija las cabeceras, mantiene Ticket en ambos ejes y elimina los title duplicados', async () => {
    const fixture = await montar();
    const host = fixture.nativeElement as HTMLElement;
    const cabeceras = [...host.querySelectorAll('thead th')] as HTMLElement[];

    expect(cabeceras.every((cabecera) => cabecera.classList.contains('top-0'))).toBe(true);
    expect(cabeceras[0].classList).toContain('left-0');
    expect(cabeceras[0].classList).toContain('z-30');
    expect(cabeceras[0].querySelector('.bg-borde')).not.toBeNull();
    expect(celdasDeLaFila(host)[0].querySelector('.bg-borde')).not.toBeNull();

    const celdasConAyuda = [3, 5, 6].map((indice) =>
      celdasDeLaFila(host)[indice].querySelector('[tabindex="0"]'),
    );
    expect(celdasConAyuda.every((celda) => !celda?.hasAttribute('title'))).toBe(true);
  });

  it('muestra una sola ayuda propia únicamente cuando el texto está truncado', async () => {
    const fixture = await montar();
    const componente = fixture.componentInstance as unknown as {
      evaluarTooltip(clave: string, contenido: HTMLElement): void;
      cerrarTooltip(clave: string): void;
    };
    const contenido = document.createElement('span');
    Object.defineProperties(contenido, {
      scrollHeight: { configurable: true, value: 50 },
      clientHeight: { configurable: true, value: 20 },
      scrollWidth: { configurable: true, value: 100 },
      clientWidth: { configurable: true, value: 100 },
    });

    componente.evaluarTooltip('CS-2026-0001:summary', contenido);
    fixture.detectChanges();
    expect(
      (fixture.nativeElement as HTMLElement).querySelectorAll('[role="tooltip"]'),
    ).toHaveLength(2);

    componente.cerrarTooltip('CS-2026-0001:summary');
    Object.defineProperty(contenido, 'scrollHeight', { configurable: true, value: 20 });
    componente.evaluarTooltip('CS-2026-0001:summary', contenido);
    fixture.detectChanges();
    // Sólo queda la ayuda permanente de las etiquetas reunidas en +N.
    expect(
      (fixture.nativeElement as HTMLElement).querySelectorAll('[role="tooltip"]'),
    ).toHaveLength(1);
  });

  it('presenta también el encabezado singular de acción en inglés', async () => {
    TestBed.inject(TranslationService).cambiar('en');
    const fixture = await montar();
    const host = fixture.nativeElement as HTMLElement;
    const encabezado = [...host.querySelectorAll('thead th')].at(-1) as HTMLElement;
    expect(encabezado.textContent?.trim()).toBe('Action');
    expect(encabezado.classList).toContain('text-left');
    expect(encabezado.classList).not.toContain('text-center');
    expect(celdasDeLaFila(host).at(-1)?.classList).toContain('text-center');
  });

  it('presenta Waiting for you al usuario que debe responder', async () => {
    TestBed.inject(TranslationService).cambiar('en');
    const fixture = await montar({
      ...PAGINA,
      tickets: [{ ...CON_CLASIFICACION, state: 'en espera' }],
    });

    expect((fixture.nativeElement as HTMLElement).textContent).toContain('Waiting for you');
  });

  it('carga el siguiente bloque al llegar al final, sin paginación visible ni duplicados', async () => {
    const fixture = await montar({ ...PAGINA, total: 21, perPage: 20 });
    const host = fixture.nativeElement as HTMLElement;

    expect(host.textContent).not.toContain('Página 1 /');
    expect(host.textContent).not.toContain('Anterior');
    expect(host.textContent).not.toContain('Siguiente');

    const carga = (
      fixture.componentInstance as unknown as { cargarMasTickets(): Promise<void> }
    ).cargarMasTickets();
    const peticion = http.expectOne(
      (solicitud) =>
        solicitud.url === '/api/tickets' &&
        solicitud.params.get('page') === '2' &&
        solicitud.params.get('perPage') === '20',
    );
    peticion.flush({
      tickets: [
        CON_CLASIFICACION,
        { ...CON_CLASIFICACION, number: 'CS-2026-0002', subject: 'Segundo ticket' },
      ],
      total: 21,
      page: 2,
      perPage: 20,
    } satisfies PaginaDeTickets);
    await carga;
    fixture.detectChanges();

    expect(host.querySelectorAll('tbody tr')).toHaveLength(2);
    expect(host.textContent).toContain('Se cargaron 1 tickets más.');
  });

  it('conserva lo cargado si falla el siguiente bloque y permite reintentarlo', async () => {
    const fixture = await montar({ ...PAGINA, total: 21, perPage: 20 });
    const componente = fixture.componentInstance as unknown as {
      cargarMasTickets(): Promise<void>;
      reintentarCarga(): void;
    };

    const cargaFallida = componente.cargarMasTickets();
    http
      .expectOne((solicitud) => solicitud.params.get('page') === '2')
      .flush({ error: 'temporary' }, { status: 503, statusText: 'Unavailable' });
    await cargaFallida;
    fixture.detectChanges();

    const host = fixture.nativeElement as HTMLElement;
    expect(host.querySelectorAll('tbody tr')).toHaveLength(1);
    expect(host.textContent).toContain('No se pudieron cargar más tickets.');
    expect(host.textContent).toContain('Reintentar');

    componente.reintentarCarga();
    http
      .expectOne((solicitud) => solicitud.params.get('page') === '2')
      .flush({
        tickets: [{ ...CON_CLASIFICACION, number: 'CS-2026-0002' }],
        total: 21,
        page: 2,
        perPage: 20,
      } satisfies PaginaDeTickets);
    await esperar();
    fixture.detectChanges();
    expect(host.querySelectorAll('tbody tr')).toHaveLength(2);
  });

  it('vuelve a leer sola mientras algo esté pendiente, y para al llegar al tope', async () => {
    vi.useFakeTimers();
    try {
      usuario.set({
        id: 1,
        name: 'María',
        lastName: 'Pérez',
        email: 'maria@demo.com',
        role: 'usuario',
        origin: 'local',
        factory: false,
      });

      const fixture = TestBed.createComponent(BandejaPage);
      fixture.detectChanges();

      http.expectOne((peticion) => peticion.url === '/api/tickets').flush(PAGINA_PENDIENTE);
      http.expectOne('/api/tickets/categories').flush(SIN_CATEGORIAS);
      await vi.advanceTimersByTimeAsync(0);

      // Recién cargada, todavía no ha releído: la espera del motor aún no ha pasado.
      http.expectNone((peticion) => peticion.url === '/api/tickets');

      // **Ocho relecturas** (el tope), cada una con el motor todavía pendiente.
      for (let intento = 0; intento < 8; intento += 1) {
        await vi.advanceTimersByTimeAsync(8000);
        http.expectOne((peticion) => peticion.url === '/api/tickets').flush(PAGINA_PENDIENTE);
        await vi.advanceTimersByTimeAsync(0);
      }

      // Al llegar al tope **se para**: por mucho que se espere, no vuelve a pedir la lista.
      await vi.advanceTimersByTimeAsync(80_000);
      http.expectNone((peticion) => peticion.url === '/api/tickets');

      fixture.destroy();
    } finally {
      vi.useRealTimers();
    }
  });

  it('deja de releer en cuanto el motor termina', async () => {
    vi.useFakeTimers();
    try {
      usuario.set({
        id: 1,
        name: 'María',
        lastName: 'Pérez',
        email: 'maria@demo.com',
        role: 'usuario',
        origin: 'local',
        factory: false,
      });

      const fixture = TestBed.createComponent(BandejaPage);
      fixture.detectChanges();

      http.expectOne((peticion) => peticion.url === '/api/tickets').flush(PAGINA_PENDIENTE);
      http.expectOne('/api/tickets/categories').flush(SIN_CATEGORIAS);
      await vi.advanceTimersByTimeAsync(0);

      // La primera relectura ya encuentra los resúmenes listos.
      await vi.advanceTimersByTimeAsync(8000);
      http.expectOne((peticion) => peticion.url === '/api/tickets').flush(PAGINA);
      await vi.advanceTimersByTimeAsync(0);

      // Y no hay una segunda.
      await vi.advanceTimersByTimeAsync(80_000);
      http.expectNone((peticion) => peticion.url === '/api/tickets');

      fixture.destroy();
    } finally {
      vi.useRealTimers();
    }
  });

  it('descarta la respuesta anterior al cambiar Principales por Internos', async () => {
    usuario.set({
      id: 2,
      name: 'Sofía',
      lastName: 'Soporte',
      email: 'sofia@demo.com',
      role: 'soporte',
      origin: 'local',
      factory: false,
    });
    const fixture = TestBed.createComponent(BandejaPage);
    fixture.componentRef.setInput('listado', 'general');
    fixture.detectChanges();

    const principal = http.expectOne(
      (peticion) => peticion.url === '/api/tickets' && peticion.params.get('type') === 'principal',
    );
    http.expectOne('/api/tickets/categories').flush(SIN_CATEGORIAS);

    const componente = fixture.componentInstance as unknown as {
      aplicarTipoGeneral(tipo: string): Promise<void>;
    };
    const cambio = componente.aplicarTipoGeneral('interno');
    const interna = http.expectOne(
      (peticion) => peticion.url === '/api/tickets' && peticion.params.get('type') === 'interno',
    );
    interna.flush({
      tickets: [{ ...CON_CLASIFICACION, internal: true, number: 'INT-CS-2026-0001' }],
      total: 1,
      page: 1,
      perPage: 20,
    } satisfies PaginaDeTickets);
    await cambio;

    principal.flush(PAGINA);
    await esperar();
    fixture.detectChanges();
    expect(fixture.nativeElement.textContent as string).toContain('INT-CS-2026-0001');
    expect(fixture.nativeElement.querySelector('a[href="/tickets/CS-2026-0001"]')).toBeNull();
  });
});
