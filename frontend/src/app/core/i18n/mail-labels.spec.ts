import { EN } from './en';
import { ES } from './es';

/**
 * El editor de correos tiene catorce plantillas desde `v1.0.1`.
 *
 * **Regresión de QA-03:** la etiqueta de `auth.accountCreated` faltaba y se veía la clave cruda; los
 * textos seguían diciendo diez/trece y doce.
 */
describe('textos de los correos', () => {
  it('etiqueta la plantilla nueva en los dos idiomas', () => {
    expect(ES.plantillas['auth.accountCreated']).toBeTruthy();
    expect(EN.plantillas['auth.accountCreated']).toBeTruthy();
  });

  it('dice catorce/catorce y fourteen en los textos de la pantalla', () => {
    expect(ES.correos.descripcion).toContain('catorce');
    expect(EN.correos.descripcion).toContain('fourteen');
    expect(ES.configuracion.correosAyuda).toContain('catorce');
    expect(EN.configuracion.correosAyuda).toContain('fourteen');
  });

  it('actualiza los avisos de revisión de idioma y la clave desconocida', () => {
    expect(ES.errores['mail.translation.incomplete']).toContain('catorce');
    expect(EN.errores['mail.translation.incomplete']).toContain('fourteen');
    expect(ES.errores['settings.language.reviewRequired']).toContain('catorce');
    expect(EN.errores['settings.language.reviewRequired']).toContain('fourteen');
    expect(ES.errores['mail.template.unknownKey']).toContain('catorce');
    expect(EN.errores['mail.template.unknownKey']).toContain('fourteen');
  });
});
