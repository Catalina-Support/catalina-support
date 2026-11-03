import { rutaSinConsulta } from './sesion.guard';

/**
 * La guarda compara la ruta, no la URL completa.
 *
 * **Regresión de QA-01:** con `?ai=required`, comparar `state.url` contra `/settings` hacía que la
 * redirección se repitiera sin fin y la aplicación no llegara a cargar.
 */
describe('rutaSinConsulta', () => {
  it('quita la query y el fragmento', () => {
    expect(rutaSinConsulta('/settings?ai=required')).toBe('/settings');
    expect(rutaSinConsulta('/forbidden?ai=required')).toBe('/forbidden');
    expect(rutaSinConsulta('/tickets/CS-2026-0001?tab=history#top')).toBe(
      '/tickets/CS-2026-0001',
    );
  });

  it('deja intacta una ruta sin query', () => {
    expect(rutaSinConsulta('/settings')).toBe('/settings');
    expect(rutaSinConsulta('/')).toBe('/');
  });

  it('cae en la raíz si la URL viene vacía', () => {
    expect(rutaSinConsulta('')).toBe('/');
  });
});
