# Dashboard y reportería

> **Estado:** as-built
> **Última actualización:** 2026-10-19
>
> **Enmienda aprobada, implementada y verificada el 2026-10-19: lectura directa de entradas y
> cierres.** La revisión del
> gráfico confirmó que no se están perdiendo cierres: los diez cierres de los datos locales
> ocurrieron entre el 3 de agosto y el 11 de septiembre de 2026 y, por tanto, quedan fuera del
> periodo consultado. El responsable eligió 1A–4A: conservar la fecha real del evento de cierre,
> mostrar siempre el valor de cada columna —incluido cero—, usar dos colores de serie propios y
> claramente distintos en los ocho temas, y explicar cuando el periodo no contiene ningún cierre.
> La interfaz y los textos bilingües quedaron implementados; dos pruebas nuevas cubren periodos con
> y sin cierres, las 243 pruebas del frontend y la compilación pasaron, y el navegador confirmó los
> valores permanentes y los colores calculados en los ocho temas. Un segundo rango, del 1 de agosto
> al 15 de septiembre, confirmó cierres reales agrupados por semana y la ausencia correcta del aviso.
>
> **Corrección aprobada, implementada y verificada el 2026-10-19.** La revisión real
> en los ocho temas encontró que los datos y tamaños de las barras existen, pero el color principal
> es transparente porque esta pantalla usa tokens ingleses que el tema no define. También encontró
> una tabla y tarjetas de tickets que contradicen la nueva decisión de que Dashboard muestre sólo
> estadísticas. El responsable cerró el repaso con 1A–4A, propuso XLSX en 5, eligió 6A–8A, 9A–14A,
> corrigió 15A al nombre `report_AAAA-MM-DD_AAAA-MM-DD.xlsx` y eligió 16A. No quedan decisiones
> abiertas. Se retiraron el detalle individual y su endpoint, se corrigieron los tokens, se redujeron
> los filtros y la exportación es ahora el XLSX agregado de seis hojas acordado.
>
> **Implementado y verificado el 2026-10-18, con aprobación explícita del responsable.** Este
> documento describe el primer módulo de Dashboard y reportería. El responsable cerró las
> decisiones 1A–45A, con la precisión de 24A de exportar ambos formatos de duración y la de 26A de
> excluir por completo al Usuario.

## 1. Propósito

El Dashboard ofrece una lectura operativa, no una evaluación de personas ni un sistema contractual
de SLA. Permite detectar carga, urgencia, antigüedad y tendencia únicamente mediante estadísticas
agregadas. No lista ni abre tickets individuales y no envía avisos ni correos.

## 2. Acceso y alcance

- Usuario no tiene ruta, menú, API ni acceso al Dashboard.
- Soporte ve principales e internos.
- Desarrollo ve únicamente internos.
- Administrador puro ve todo.
- Soporte o Desarrollo con capacidad Admin acumula el alcance administrativo y ve todo.

Dashboard/Panel aparece antes de las bandejas y es el destino inicial de Soporte, Desarrollo y
Administrador. Usuario continúa en Mis tickets.

## 3. Periodo, filtros y actualización

Abre con los últimos treinta días en la zona global de la instalación. Admite únicamente rango de
fechas, tipo, estado y prioridad. Los filtros viven en la URL y gobiernan indicadores, gráficos y
XLSX. Los datos se consultan al abrir, al cambiar filtros y al pulsar Actualizar; no hay sondeo ni
tiempo real. Categoría y responsable permanecen como desgloses, no como filtros.

## 4. Indicadores

La cabecera muestra abiertos, críticos abiertos, atrasados, tiempo medio de primera respuesta,
tiempo medio de resolución y reaperturas. Cada uno compara con el periodo inmediatamente anterior
de igual duración.

Abierto significa cualquier estado distinto de cerrado. Un valor inexistente se presenta con guion;
un ticket pendiente de respuesta dice Sin respuesta. Las medias nunca convierten ausencias en cero.

## 5. Definiciones temporales

- Primera respuesta del principal: primer comentario humano de Soporte.
- Primera respuesta del interno: primer comentario humano de Desarrollo.
- La media usa sólo tickets que ya recibieron respuesta y muestra aparte los pendientes.
- Resolución: tiempo desde la creación hasta la primera resolución; la media incluye las primeras
  resoluciones ocurridas en el periodo.
- Reaperturas se cuentan aparte.
- Entradas usan la fecha de creación; cierres usan la fecha del evento de cierre.
- La antigüedad acumula mientras el ticket está pendiente, se detiene al resolver y vuelve a
  acumular desde una reapertura.

## 6. Atraso configurable

Un ticket abierto está atrasado al superar, en horas corridas, el umbral de su prioridad. Los valores
iniciales son crítica 4 h, alta 24 h, media 72 h y baja 120 h. Sólo Administrador los modifica en la
tarjeta Reportería de Configuración; los equipos sólo los consultan. Son criterios operativos, no un
compromiso formal de SLA ni un porcentaje de cumplimiento.

## 7. Gráficos

Se usan barras y líneas accesibles, nunca gráficos circulares:

1. entradas frente a cierres;
2. distribución por estado;
3. distribución por prioridad;
4. diez categorías principales y Otras;
5. carga abierta por responsable, incluido Sin responsable.

La tendencia agrupa por día hasta 31 días y por semana desde 32 días. La interfaz muestra sólo las
etiquetas de fecha que caben sin solaparse. En entradas frente a cierres, cada columna muestra
siempre su valor numérico —también cuando es cero— sin exigir puntero, foco ni interacción; la ayuda
al puntero y el texto accesible conservan además la fecha, la serie y el valor exactos. Las series
usan dos tokens propios del Dashboard: entradas conserva la relación con el color primario y
cerradas recibe un color secundario claramente distinguible. Los dos tokens se definen y se
comprueban en los ocho temas; no dependen sólo de la leyenda para diferenciarse. Si todos los cierres
del periodo son cero, el gráfico permanece visible con sus ceros y añade «No tickets closed in this
period» / «No se cerraron tickets en este periodo». El gráfico de responsables mide carga, no
productividad.

## 8. Sin detalle de tickets

Dashboard no contiene tabla, tarjetas, carga progresiva, números, asuntos ni enlaces a tickets. El
detalle operativo continúa exclusivamente en las bandejas y fichas del módulo de tickets.

## 9. Libro XLSX

La única exportación es un libro XLSX de datos agregados con seis hojas: Resumen, Serie temporal,
Estado, Prioridad, Categoría y Responsable. Nombres de hojas, encabezados y valores traducibles usan
el idioma global. Resumen incluye el periodo, los cuatro filtros aplicados y los seis indicadores.
Las demás hojas contienen exactamente los agregados visibles; no añaden comparación oculta ni filas
de tickets.

Las hojas son tablas formateadas con encabezado fijo y filtros, sin gráficas incrustadas. Las fechas
son celdas reales con formato local y añaden valor ISO cuando sea útil para integración. Las
duraciones conservan valor humano y minutos enteros. El archivo se llama
`report_AAAA-MM-DD_AAAA-MM-DD.xlsx`, con los extremos del periodo filtrado. El servidor neutraliza
texto que pudiera interpretarse como fórmula.

## 10. Estados de interfaz

Indicadores y gráficos comparten carga, vacío y error sin borrar el último resultado válido durante
una actualización manual. En móvil se apilan sin añadir una lista de tickets. Las distribuciones
usan los tokens reales `primario` y `primario-texto`; la serie temporal usa sus tokens específicos
de entradas y cierres. Los valores, los dos colores entre sí y su contraste con el fondo se
comprueban en los ocho temas. Todo funciona con teclado, lector de pantalla y ambos idiomas.

## 11. Fuera de alcance

No incluye productividad individual, metas, SLA contractual, alertas, correos, tiempo real,
programación de reportes, CSV, PDF, preferencias personales, acceso del Usuario ni ningún detalle
o acceso a tickets individuales.

## 12. Verificación requerida

Las pruebas cubrirán cálculos con PostgreSQL real, permisos y alcance acumulativo, zona horaria,
periodos anteriores, los cuatro filtros URL, agrupación diaria/semanal, umbrales, reaperturas,
valores ausentes, XLSX de seis hojas localizado y neutralizado, ausencia de tickets, responsive en
PC y móvil, barras visibles en los ocho temas, valores permanentes incluido cero, diferenciación de
las dos series y el aviso bilingüe cuando el periodo no contiene cierres.

La implementación conserva `/api/dashboard/summary` y adapta
`/api/dashboard/export` al libro XLSX; la UI no consume detalle y se retiró el contrato
que sólo sostenía la lista individual. La ruta `/dashboard`, su entrada de menú y la tarjeta
Reportería de Configuración permanecen. PostgreSQL reconstruye estados por corte, intervalos activos
y hechos temporales; el servicio Dashboard agrega únicamente estadísticas autorizadas.

La comprobación recorrió Soporte, Desarrollo, Usuario rechazado y Administrador. La revisión
del 2026-10-18 añadió los hallazgos de color transparente en los ocho temas, fechas solapadas y
detalle de tickets fuera del nuevo alcance. Tras implementarlo pasaron los servicios Go, el repositorio
de reportería, las 241 pruebas del frontend, las fronteras y la compilación. El navegador confirmó
ausencia de tickets individuales, cuatro filtros, etiquetas temporales adaptativas y color efectivo
en los ocho temas; el XLSX descargado contiene seis hojas y XML válido.
