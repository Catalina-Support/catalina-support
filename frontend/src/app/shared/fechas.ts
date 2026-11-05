/**
 * Las fechas de la instalación.
 *
 * La instalación tiene **una zona horaria configurable** —un nombre IANA, como `America/Guayaquil`— y
 * es la que decide cómo se leen todas las fechas, en la interfaz y en los correos. **Las fechas
 * guardadas siguen en UTC**: cambiar la zona no mueve ningún ticket, sólo cambia la hora a la que se
 * lee. Sin esto, la misma actuación se vería a horas distintas según dónde estuviera quien mira, que
 * es justo lo que se arregla.
 *
 * Vive en `shared` porque lo usan los dos módulos que formatean fechas —`tickets` y `users`—, y la
 * regla de qué hacer cuando la zona no vale tiene que ser la misma en los dos.
 */

/**
 * La zona que se le puede pasar a `Intl`, o `undefined` para que use **la del navegador**.
 *
 * **La marca es pública y no puede romper la pantalla**: si el campo no llegara —`undefined`, o una
 * cadena vacía— o si `Intl` no conociera el nombre, se devuelve `undefined` y se lee con la zona del
 * navegador, que es lo que se hacía antes de que existiera este campo. Lanzar aquí dejaría la lista
 * de tickets sin pintar, y por un dato que sólo decide a qué hora se lee algo.
 */
export function zonaSegura(zona: string | undefined): string | undefined {
  const limpia = zona?.trim();
  if (!limpia) {
    return undefined;
  }

  try {
    // Es la única forma fiable de saber si `Intl` conoce el nombre: con una zona que no reconoce,
    // `Intl.DateTimeFormat` lanza `RangeError`.
    new Intl.DateTimeFormat('es', { timeZone: limpia });
    return limpia;
  } catch {
    return undefined;
  }
}

/**
 * Las opciones de `Intl.DateTimeFormat` con la zona de la instalación puesta, o **sin `timeZone`** si
 * la zona no vale —y entonces se lee con la del navegador—.
 *
 * El formato no cambia: sólo se le añade la zona.
 */
export function opcionesConZona(
  opciones: Intl.DateTimeFormatOptions,
  zona: string | undefined,
): Intl.DateTimeFormatOptions {
  const valida = zonaSegura(zona);
  return valida ? { ...opciones, timeZone: valida } : opciones;
}

/**
 * Una fecha en formato ISO legible: `AAAA-MM-DD hh:mm`, siempre en 24 horas.
 *
 * **El formato es el mismo en español y en inglés** (decisión del responsable, 2026-11-05): una
 * fecha no cambia de orden al cambiar de idioma, y `2026-09-29 21:00` no se confunde con el 9 de
 * septiembre. La hora sigue siendo la de la instalación; las fechas guardadas siguen en UTC.
 *
 * Se arma con las partes de `Intl` en vez de confiar en el orden de un locale: así el resultado no
 * depende de la versión del navegador ni de que el locale use 12 o 24 horas.
 */
export function fechaIso(iso: string | undefined, zona?: string): string {
  if (!iso) {
    return '';
  }

  const momento = new Date(iso);
  if (Number.isNaN(momento.getTime())) {
    return '';
  }

  const partes = new Intl.DateTimeFormat(
    'en-GB',
    opcionesConZona(
      {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        hourCycle: 'h23',
      },
      zona,
    ),
  ).formatToParts(momento);

  const valor = (tipo: Intl.DateTimeFormatPartTypes): string =>
    partes.find((parte) => parte.type === tipo)?.value ?? '';

  return `${valor('year')}-${valor('month')}-${valor('day')} ${valor('hour')}:${valor('minute')}`;
}
