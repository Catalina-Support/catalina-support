import { ENTRADAS, claveDeEntrada, entradasPara, rutaDeMenuActiva } from './navigation';

describe('entradasPara', () => {
  it('el menú enseña lo que existe, y sólo lo que existe', () => {
    // La lista ya es la del documento de interfaz (sección 3.2). **«Nuevo ticket» salió el
    // 2026-09-26** —se confundía con un módulo— y en su lugar están las dos listas del «todo».
    const rutas = ENTRADAS.map((entrada) => entrada.ruta);

    expect(rutas).toEqual([
      '/dashboard',
      '/my-tickets',
      '/tickets',
      '/tickets/categories',
      '/users',
      '/settings',
    ]);
  });

  it('cada papel tiene sus entradas, y ninguna de más', () => {
    expect(entradasPara('usuario').map((entrada) => entrada.ruta)).toEqual(['/my-tickets']);

    expect(entradasPara('soporte').map((entrada) => entrada.ruta)).toEqual([
      '/dashboard',
      '/my-tickets',
      '/tickets',
    ]);

    expect(entradasPara('desarrollo').map((entrada) => entrada.ruta)).toEqual([
      '/dashboard',
      '/my-tickets',
      '/tickets',
    ]);

    expect(entradasPara('administrador').map((entrada) => entrada.ruta)).toEqual([
      '/dashboard',
      '/my-tickets',
      '/tickets/categories',
      '/users',
      '/settings',
    ]);
  });

  it('Clasificación la ve sólo quien tiene capacidad Administrador', () => {
    const entrada = ENTRADAS.find((candidata) => candidata.ruta === '/tickets/categories');

    expect(entrada?.clave).toBe('categorias');
    expect(entrada?.papeles).toEqual(['administrador']);
    for (const papel of ['usuario', 'soporte', 'desarrollo']) {
      expect(entradasPara(papel).map((candidata) => candidata.ruta)).not.toContain(
        '/tickets/categories',
      );
    }
  });

  it('la bandeja se llama distinto según quien mire', () => {
    // Una sola pantalla con dos nombres —«Mis tickets» para los tres que tienen tickets propios, y
    // «Bandeja» para el Administrador, que no tiene—, y el menú no se inventa dos entradas que llevan
    // al mismo sitio.
    const bandeja = ENTRADAS.find((entrada) => entrada.ruta === '/my-tickets')!;

    expect(claveDeEntrada(bandeja, 'usuario')).toBe('misTickets');
    expect(claveDeEntrada(bandeja, 'desarrollo')).toBe('misTickets');
    expect(claveDeEntrada(bandeja, 'soporte')).toBe('misTickets');
    expect(claveDeEntrada(bandeja, 'administrador')).toBe('bandeja');
  });

  it('la lista general unificada es de Soporte y Desarrollo', () => {
    const entrada = ENTRADAS.find((candidata) => candidata.ruta === '/tickets');
    expect(entrada?.papeles).toEqual(['soporte', 'desarrollo']);
    expect(entradasPara('usuario').map((candidata) => candidata.ruta)).not.toContain('/tickets');
    expect(entradasPara('administrador').map((candidata) => candidata.ruta)).not.toContain(
      '/tickets',
    );
  });

  it('Configuración sólo la ve un administrador', () => {
    expect(entradasPara('administrador').map((entrada) => entrada.ruta)).toContain('/settings');

    for (const papel of ['usuario', 'soporte', 'desarrollo']) {
      expect(entradasPara(papel).map((entrada) => entrada.ruta)).not.toContain('/settings');
    }
  });

  it('acumula las entradas administrativas sin perder las del papel operativo', () => {
    expect(entradasPara('soporte', true).map((entrada) => entrada.ruta)).toEqual([
      '/dashboard',
      '/my-tickets',
      '/tickets',
      '/tickets/categories',
      '/users',
      '/settings',
    ]);

    expect(entradasPara('desarrollo', true).map((entrada) => entrada.ruta)).toEqual([
      '/dashboard',
      '/my-tickets',
      '/tickets',
      '/tickets/categories',
      '/users',
      '/settings',
    ]);
  });

  it('no convierte a un usuario en administrador desde la navegación', () => {
    expect(entradasPara('usuario', true).map((entrada) => entrada.ruta)).toEqual(['/my-tickets']);
  });

  it('crear un ticket no está en el menú, para nadie', () => {
    // Es una acción de la bandeja, no un módulo (decisión 53): la ruta sigue existiendo y se entra
    // por el botón que hay dentro.
    expect(ENTRADAS.map((entrada) => entrada.ruta)).not.toContain('/tickets/new');
  });

  it('un papel que no conocemos no ve nada de nadie', () => {
    // Si el backend mandara un papel nuevo, el menú no se le adelanta enseñando entradas de más: el
    // backend lo rechazaría igual, pero el menú no promete lo que no hay.
    expect(entradasPara('superusuario')).toEqual([]);
    expect(entradasPara('')).toEqual([]);
  });

  it('cada entrada tiene su ruta y su clave de texto', () => {
    for (const entrada of ENTRADAS) {
      expect(entrada.ruta.startsWith('/')).toBe(true);
      expect(entrada.clave.length).toBeGreaterThan(0);
    }
  });
});

describe('rutaDeMenuActiva', () => {
  it('activa exactamente la bandeja personal o la lista general', () => {
    expect(rutaDeMenuActiva('/my-tickets')).toBe('/my-tickets');
    expect(rutaDeMenuActiva('/tickets')).toBe('/tickets');
    expect(rutaDeMenuActiva('/tickets/categories')).toBe('/tickets/categories');
  });

  it('mantiene las pantallas auxiliares dentro de su módulo acordado', () => {
    expect(rutaDeMenuActiva('/tickets/new')).toBe('/my-tickets');
    expect(rutaDeMenuActiva('/tickets/CS-2026-0001')).toBe('/my-tickets');
    expect(rutaDeMenuActiva('/tickets/INT-CS-2026-0001?view=both')).toBe('/my-tickets');
    expect(rutaDeMenuActiva('/users/42')).toBe('/users');
    expect(rutaDeMenuActiva('/mail')).toBe('/settings');
  });

  it('tolera query, fragmento y barra final, y no inventa selección para otras pantallas', () => {
    expect(rutaDeMenuActiva('/dashboard?from=2026-09-19#charts')).toBe('/dashboard');
    expect(rutaDeMenuActiva('/users/')).toBe('/users');
    expect(rutaDeMenuActiva('/profile')).toBeNull();
    expect(rutaDeMenuActiva('/change-password')).toBeNull();
  });
});
