# Interfaz y experiencia

> **Estado:** as-built
> **Última actualización:** 2026-10-22
>
> **Propuesta coordinada del 2026-10-22.** La sección 51 reduce a 16 px el margen horizontal de las
> cuatro bandejas, elimina su máximo de ancho y fija Ticket y Action en los extremos con aire
> interior. El responsable cerró el repaso con 1A–4A. No se modifica código antes de aprobar esta
> propuesta junto con `modules/tickets.md` y `modules/ai.md`.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-20.** La sección 50 añade la tarjeta **Icono del sitio** y define
> la actualización inmediata del favicon en todas las pestañas. El responsable cerró 1A–14A y
> corrigió el fallback: sin icono propio se usa el logo claro efectivo. Pasaron 253 pruebas frontend,
> las fronteras y la compilación de producción.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-20.** La sección 49 sustituye las dos
> entradas generales por **Tickets**, reserva Usuarios y Clasificación a Admin y repasa la
> terminología visible completa del área de tickets.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-20.** El diálogo de alta de `/users`
> incorpora búsqueda de Active Directory por usuario o correo cuando AD es el método vigente,
> resultados seleccionables y confirmación explícita al vincular una cuenta local. Keycloak se
> muestra deshabilitado explicando que la cuenta nace en el primer acceso. Decisiones 4A–15A, con
> 6B, cerradas por el responsable.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-19.** La sección 47 añade el selector **Public / Private** al
> editor del principal, distingue visualmente las notas privadas y garantiza su ausencia total en la
> vista del Usuario. El responsable cerró y aprobó 1A–26A; se verificó el flujo real con ambos papeles.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-19.** La sección 46 lleva `All · Assigned · Watching` a la
> primera fila de **My tickets** y abre en `All`. La lista representa exclusivamente asignación u
> observación actuales, separada por nivel y sin ampliarse por la capacidad Admin. El responsable
> cerró 4A y 6A–9A y aprobó el documento.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-19.** La sección 45 corrige el estado activo del menú lateral.
> El hallazgo es que `/tickets` se evalúa hoy como prefijo y queda seleccionado a la vez que
> **Tickets principales**, **Tickets internos** o **Clasificación**. El responsable eligió 1A–4A:
> una sola entrada activa, fondo suave sin sombra, pertenencia de alta y detalle a **Mis tickets**,
> y pertenencia de las fichas de usuario y correo a **Usuarios** y **Configuración**. La función
> de resolución se cubre con pruebas unitarias; la compilación y las 246 pruebas del frontend pasan.
> El navegador confirmó selección única en todas las entradas y en alta y detalle de ticket, ficha
> de usuario y editor de correos.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-19.** La sección 43 corrige la
> composición del papel operativo y la insignia administrativa en la lista de usuarios. El
> responsable eligió 1A–4A: el papel ocupa
> la primera fila y la insignia **Admin / Administrador** una segunda, alineada a la izquierda y con
> una separación vertical pequeña y constante, tanto en la tabla de escritorio como en las tarjetas
> de pantallas estrechas. Pasaron las 243 pruebas del frontend, sus fronteras y la compilación. Se
> añadió una regresión Playwright común a PC y móvil; el entorno aislado no llegó a ejecutarla por un
> fallo de montaje de su volumen `node_modules`, ajeno a esta interfaz.
>
> **Corrección aprobada, implementada y verificada el 2026-10-19.** La sección 44 retira la tabla y las tarjetas
> del Dashboard, reduce sus filtros, corrige los tokens de color transparentes en los ocho temas,
> evita el solapamiento de fechas y sustituye CSV por un XLSX agregado de seis hojas. El responsable
> cerró 1A–4A, propuso XLSX en 5, eligió 6A–14A, corrigió el nombre en 15A y eligió 16A. No quedan
> decisiones abiertas. Pasaron las 241 pruebas del frontend, fronteras y compilación; el navegador
> confirmó las barras en los ocho temas, los cuatro filtros y la ausencia de detalle individual.
>
> **Implementado y verificado el 2026-10-18.** La sección 44 define la navegación y presentación
> accesible del nuevo Dashboard.
>
> **Implementado el 2026-10-18.** La sección 43 adapta cuentas, navegación y avisos al papel
> operativo más capacidad administrativa.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-17.** La sección 42 aprovecha el alto
> disponible en las tablas operativas de tickets y usuarios: 16 px exteriores arriba y abajo, 12 px
> entre cabecera y tabla, fila de acciones de 42 px y contenedor estable hasta el borde inferior.
> El hallazgo es que hoy el armazón reserva 32 px arriba y 40 px abajo, la bandeja añade 24 px entre
> bloques y sólo tickets calcula el alto restante; Usuarios conserva el flujo normal de página. El
> responsable cerró el repaso con 1A–15A; no quedan decisiones abiertas. Pasaron las 238 pruebas
> del frontend, sus fronteras y la compilación. El navegador confirmó en las bandejas de tickets y
> Usuarios la cabecera compacta, el uso del alto disponible y la conservación de las tarjetas móviles.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-17.** La sección 41 reduce y
> recompone el menú lateral: 208 px abierto y 56 px plegado, marca vertical centrada, logo mayor y
> un único control de plegado en la parte superior. También adopta **Classification / Clasificación**
> como nombre breve de la pantalla conjunta de categorías y etiquetas. El hallazgo de partida es que
> esta sección todavía documentaba y la implementación todavía usa 240/64 px y una cabecera
> horizontal que desplaza el logo al plegarse. El responsable cerró el repaso con 1A, 2A, 3B, 4C,
> 5A–14A; no quedan decisiones abiertas. Pasaron las 238 pruebas del frontend, sus fronteras y la
> compilación. El navegador confirmó escritorio abierto y plegado, cajón móvil, el control único,
> la marca centrada y Classification con sus dos tarjetas.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-16.** La sección 40 distingue el
> rótulo visible de `en espera` según quien consulta el ticket: el Usuario leerá `Waiting for you` /
> `Esperándote`, mientras Soporte, Desarrollo y el Administrador leerán `Waiting` / `Esperando`.
> La diferencia se aplicará en toda la interfaz, sin cambiar el valor técnico, la API, la base de
> datos ni los correos. El responsable eligió 1A–4A en el repaso previo; no quedan decisiones
> abiertas. Pasaron 76 pruebas focalizadas, las fronteras del frontend y la compilación.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-16.** La sección 39 convierte la tabla
> de las bandejas en el área desplazable que ocupa el alto restante, fija verticalmente sus
> encabezados, repara el separador de Ticket y elimina los tooltips duplicados. El responsable eligió
> 1A–13A. Como hallazgos, la sección 37 prohibía fijar verticalmente el encabezado y la sección 36.2
> combinaba el tooltip propio con `title`; ambos contratos quedaron corregidos. Pasaron nueve pruebas
> focalizadas, las fronteras y la compilación. El navegador verificó PC, tableta y móvil, scroll en
> ambos ejes, carga 20→27, restauración interna, ayuda única sólo al truncar y los ocho temas.
> **Corrección de compatibilidad del 2026-10-16:** Action/Acción declara `text-left` directamente en
> su `th`; depender de la herencia de la fila permitía que Firefox aplicara su alineación nativa.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-16.** La sección 38 sustituye la
> paginación visible de las bandejas por carga progresiva con scroll, acorta el estado inglés del
> usuario a `Waiting` y alinea a la izquierda el encabezado Action/Acción. El responsable eligió
> 1A–10A y no quedaron decisiones abiertas. Pasaron siete pruebas focalizadas y la compilación del
> frontend. El navegador confirmó 20 elementos iniciales y 27 tras llegar al final, sin paginación,
> en PC, tableta y móvil; también confirmó `Waiting`, las dos alineaciones y la restauración de los
> 27 elementos y de la posición al volver desde el detalle.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-15.** La sección 37 mantiene
> visible la columna Ticket durante el desplazamiento horizontal de las tablas amplia e intermedia.
> El responsable eligió 1A–7A: sólo se fija Ticket, únicamente en el eje horizontal, con su ancho
> actual, fondo opaco del tema y borde derecho; se comprobarán los ocho temas. Durante la preparación
> se detectó que la sección 36 impedía el desplazamiento que esta decisión necesita. El responsable
> corrigió el contrato con 1A–3A: desde 768 px se conserva una tabla con desplazamiento nativo, sin
> aviso adicional, y sólo por debajo se cambia a tarjetas. No quedan decisiones abiertas. La prueba
> específica y la compilación pasaron. En navegador, la
> tabla aparece a 768 px y las tarjetas a 767 px; al desplazar 165.5 px, la segunda columna se movió
> pero Ticket conservó exactamente su posición. Los ocho temas mostraron fondo opaco y borde de 1 px.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-15.** La sección 36 hace
> visible el encabezado de la última columna de las bandejas de PC. La comprobación real confirmó
> que «Abrir» sí vive dentro de un `td`, pero su `th` estaba oculto visualmente y por eso parecía
> quedar fuera de la tabla. El responsable eligió 1B, 2A, 3A, 4A, 5A, 6A, 7A–10A y, en el repaso
> integral de esta tabla, 1A–12A, 13B+13C, 14A–24A. Corrigió expresamente el orden de las columnas.
> No quedaron decisiones abiertas. Pasaron la prueba específica y la compilación del frontend; el
> navegador confirmó once columnas en ancho amplio, nueve en intermedio y tarjetas antes de que
> aparezca desplazamiento horizontal.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-15.** La sección 35 cambia
> Motivo/Reason por Resumen/Summary en todas las vistas, elimina los dos botones de reintento y deja
> estados discretos sólo en la ficha. El responsable eligió 1A–20B y no quedan decisiones abiertas.
> La interfaz quedó sin acciones manuales, presenta `—` sin resultado vigente y reserva el estado
> discreto para la ficha. Pasaron las 231 pruebas del frontend, sus fronteras y la compilación.
>
> **Corrección aprobada, implementada y verificada el 2026-10-15.** La sección 34
> iguala todos los desplegables a los 42 px reales del conmutador Main/Both/Internal. El responsable
> eligió 1A–5A: altura fija y global, sin modificar el conmutador ni la configuración de recarga del
> entorno. Pasaron las 232 pruebas del frontend y la compilación; un navegador limpio midió 42 px,
> fuente de 14 px y padding de 12 × 8 px. No quedan decisiones abiertas.
>
> **Ajuste implementado y verificado el 2026-10-15, con aprobación explícita del responsable.** La
> sección 33 registra que los 52 px se comprobaron correctamente en un navegador limpio, pero
> resultaron demasiado grandes. El responsable eligió 1A, 2A y 3A: todos los desplegables usan ahora
> 48 px, con el padding y la tipografía actuales. Pasaron las 232 pruebas del frontend, la
> compilación y la medición en una pestaña nueva; no quedan decisiones abiertas.
>
> **Corrección implementada y verificada el 2026-10-15, con aprobación explícita del responsable.**
> La sección 32 registra que los 44 px de la sección 31 sí llegaron al navegador, pero aún se
> percibían demasiado delgados. El responsable eligió 1A, 2A y 3A: todos los desplegables usan
> ahora 52 px y conservan la tipografía actual. Pasaron las 232 pruebas del frontend, la compilación
> y la medición real; no quedan decisiones abiertas.
>
> **Implementado y verificado el 2026-10-15, con aprobación explícita del responsable.** La sección
> 31 alinea la altura y el espaciado de todos los desplegables con los campos de texto y los botones.
> El responsable eligió 1A–5A; pasaron las 232 pruebas del frontend, la compilación y la revisión
> real en PC y móvil. No quedan decisiones abiertas.
>
> **Implementado y verificado el 2026-10-15, con aprobación explícita del responsable.** La sección 30 hace visibles la calidad de
> las propuestas y los errores de cada resumen sin perder borradores durante el refresco. El
> responsable cerró 1A–15A, con 8B, y corrigió la propuesta retirando dos ayudas redundantes del
> editor y del panel de revisión; no quedan decisiones abiertas.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección 29 sustituye el modal y sus cinco
> tonalidades por generación profesional inmediata y revisión inline independiente en cada editor.
> El responsable cerró 1A–15A con 3B, 6B, 7A, 10B y 11B; no quedan decisiones abiertas.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección
> 28 corrige el espacio interno
> que la vista doble todavía reserva para una ficha oculta y hace que cada conversación use por
> completo su mitad y habilita la comparación desde 1152 px. El responsable eligió 1A–6A, 7B y
> 8B; el repaso está cerrado sin decisiones abiertas.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección
> 27 da desplazamiento y ancho independientes a los dos tickets de la vista doble, sin estrecharlos
> en pantallas pequeñas.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección
> 26 hace visible el catálogo local desde que se abre la pantalla y convierte su botón en una
> actualización sin salto visual.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección
> 25 elimina las confirmaciones nativas del navegador y añade al paso 1 el selector completo de tema
> y el control institucional de color, reutilizando los componentes y reglas existentes.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección 24 convierte cada ficha de
> modelo en el único punto de activación o desactivación, lleva el saludo a un toast que identifica
> el modelo y convierte la propia vista previa del logo en el control de carga. También enmienda la
> política general de toast: todos admiten cierre manual y duran como máximo diez segundos.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección 23 corrige el catálogo local que
> hoy sólo muestra «Descargando», bloquea el avance prematuro del asistente, añade la prueba manual
> fechada y lleva los dos logos al primer paso. El comportamiento se cerró con el responsable antes
> de escribir esta propuesta.
>
> **Enmienda implementada para 1.0.0, 2026-10-14.** La sección 22 define selección, chip accesible,
> edición en el lugar, filtro múltiple y orden de prioridades en PC y móvil. No se modifica código
> antes de implementarla. El responsable la aprobó explícitamente.
>
> **Corrección aprobada y verificada el 2026-10-10:** el repaso de la sección 21 encontró que el
> documento prometía conservar el enlace de fuente en el menú plegado, mientras la implementación
> y su prueba lo ocultan hasta volver a desplegarlo. El responsable eligió 39A: documentar el
> comportamiento actual, sin añadir otro icono a la barra plegada. El responsable aprobó
> explícitamente 39A.
>
> **Enmienda implementada y verificada el 2026-10-10: fuente y licencia visibles.** La sección 21
> incorpora un enlace discreto en la entrada y el menú, alimentado por la marca pública. Pasó las
> pruebas de componente y el recorrido aislado en PC y móvil.
>
> **Enmienda implementada y verificada el 2026-10-08: ocultar la ayuda sin IA configurada.** La
> sección 20 hace que el botón y su modal dependan de la capacidad recibida con el detalle del
> ticket. El responsable aprobó 1A–3A y confirmó mantener la IA obligatoria.
>
> **Corrección aprobada e implementada el 2026-10-08.** No existe un recorrido real de ticket sin IA configurada:
> la obligatoriedad vigente lleva antes a Configuración. El responsable decidió conservarla. El
> estado falso se probará en el componente y el recorrido de navegador verificará el estado
> configurado. Las pruebas finales respetan esta frontera.
>
> **Enmienda implementada y verificada el 2026-10-08: modal «Mejorar con IA».** La sección 19
> describe el botón con icono, el editor sencillo, las cinco tonalidades, la revisión y sus estados.
> Pasó pruebas unitarias y el recorrido real en PC y móvil, en el entorno desechable.
>
> **Enmienda implementada y verificada el 2026-10-08 (bloque 6).** El único fleco visual registrado en
> `docs/arquitectura.md` era un repaso de formato sin alcance ni criterio de cierre. La sección 18
> define la revisión completa elegida por el responsable en 1A–14C. Lo existente continúa
> as-built; el responsable aprobó explícitamente esta enmienda el 2026-10-08.
>
> **Hallazgo corregido y verificado el 2026-10-08.** El bloque 5 encontró cifras de RAM distintas entre
> los selectores y el administrador, y un error genérico cuando un modelo no cabe. La sección 17
> alinea las cifras y define el toast elegido en 17A y 22A. La corrección fue aprobada explícitamente
> el 2026-10-08.
>
> **Enmienda propuesta el 2026-10-06.** Lo existente continúa as-built. La sección 16 define el
> idioma global como primer control y la experiencia para configurar, descargar y activar la IA
> obligatoria. El repaso quedó cerrado, la propuesta fue aprobada y la implementación fue verificada.
>
> **Enmendado el 2026-10-06**, con aprobación explícita del responsable. Los resultados de guardar y
> probar conexiones en `/settings` y `/setup` se mostraban en un aviso al principio del documento. En una pantalla
> larga, el resultado queda fuera de la parte visible y parece que la acción no respondió. El
> responsable confirmó un toast compartido, fijo en la ventana, para esos resultados. Los avisos que
> explican una condición de la pantalla permanecen dentro de su sección. Se implementó y verificó en
> PC y móvil.
>
> **Enmendado el 2026-10-05**, con aprobación explícita del responsable. Configuración tenía el mismo
> problema visual que tuvo el paso 3 de `/setup`: una lista larga puede ocultar la opción resaltada.
> El responsable confirmó aplicar en `/settings` el mismo campo de sólo lectura antes del buscador,
> con el identificador IANA exacto, actualización inmediata y la hora y el desfase debajo. Se
> implementó y se verificó en PC y móvil.
>
> **Enmendado el 2026-10-05**, con aprobación explícita del responsable. Al revisar todas las vistas se
> encontró que los controles de navegación de cabecera ya están a la izquierda en el alta y el
> detalle de tickets y en la ficha de usuario, pero `/mail` coloca a la derecha su vuelta a
> Configuración aunque este mismo documento dice desde el 2026-09-29 que debe ir a la izquierda.
> El responsable confirmó que se unifiquen los controles de cabecera como icono antes del título y
> que los enlaces de flujo dentro del contenido conserven su presentación. Se corrigió `/mail` y se
> verificó la posición y el regreso a Configuración en PC y móvil.
>
> **Enmendado el 2026-10-05**, conforme a `docs/primer-arranque.md` aprobado: el paso 3 de `/setup`
> muestra antes del buscador un campo de sólo lectura con la zona seleccionada. Enseña el
> identificador IANA exacto, se actualiza en el acto al elegir en la lista y permite seleccionar y
> copiar el texto; la hora y el desfase continúan debajo. Se verificó en PC y móvil.
>
> **Enmendado el 2026-10-04**, conforme a `docs/prueba-local.md` aprobado: `/setup` es la excepción
> inicial al idioma del navegador. Una instalación nueva se presenta entera en inglés; su selector
> cambia todos los textos, `html[lang]` y el valor recordado al instante. Al reanudar manda el idioma
> ya guardado en la instalación.
>
> **Enmendado el 2026-10-03**, conforme a `docs/prueba-local.md` aprobado: Los formularios de
> Keycloak en Configuración y primer arranque incorporan la dirección interna opcional, separada del
> emisor público, con ayuda ES/EN. El selector de Configuración permite elegir métodos todavía
> vacíos para mostrar sus campos; la selección es un borrador hasta guardar con éxito (enmienda
> §5.1.b aprobada).
>
> **Enmendado el 2026-09-30 (novena vez)**: **las tarjetas de Configuración se ordenaron y cada una
> lleva lo que su título dice**. Estaban descolocadas desde el cambio anterior: «La numeración y el
> reparto» tenía dentro sólo el idioma, y el prefijo y el reparto vivían en «Región horaria y dirección
> pública». Ahora son seis, en este orden: **La instalación** (nombre e idioma) · **La marca** (logo y
> color) · **Método de autenticación** · **La numeración y el reparto** · **Región horaria y dirección
> pública** · **Los correos** (`docs/modules/settings.md`, sección 5.11).
>
> **Enmendado el 2026-09-29 (octava vez)**: en **Configuración** entran **la región de la instalación** y
> **la dirección pública** (decisiones 13 a 15 de `docs/modules/settings.md`). La región se elige de una
> **lista de zonas con buscador** y la pantalla enseña **la hora que es en ella y su desfase**; la
> dirección se escribe con su esquema, su host y —si hace falta— su puerto, con **`localhost` permitido**
> y **un aviso, sin bloquear nada, cuando no es https**.
>
> **Enmendado el 2026-09-29 (séptima vez)**, a petición del responsable, al usar el producto: (1) el
> **número del ticket manda** —más grande, y crece en pantallas anchas— y **se retira el cartel «El ticket
> interno»**; (2) los **botones de sólo icono se achican** (de 44 px a 36, que sigue por encima de lo que
> se puede pulsar); (3) en **Configuración**, «Cómo se entra» pasa a llamarse **«Método de
> autenticación»**; (4) el botón **«Configuración» de la pantalla de correos** se cambia por **el botón de
> atrás** —la flecha, a la izquierda—, igual que en las listas de tickets; y (5) en **Usuarios**, el botón
> de mandar el enlace pasa a llamarse **«Cambiar contraseña»** y **sólo se enseña en las cuentas locales**
> (decisión 90 y `docs/usuarios-y-permisos.md`).
>
> **Enmendado el 2026-09-29 (sexta vez)**, a petición del responsable, al usar el producto: (1) en la
> cabecera del ticket, **el número de tu equipo va primero y con más peso** —Desarrollo ve primero el
> **interno**; Soporte y el Administrador, el **principal**— y **se retiran los enlaces «Ver el ticket
> interno» y «Ver el principal»**, porque para eso está el conmutador (decisión 87); (2) **«En espera»
> sólo se confirma**, sin pedir texto, y **cerrar sigue exigiendo el comentario** (decisión 88); y (3) la
> **lista de a quién etiquetar sale junto al cursor**, donde se está escribiendo, en vez de encima del
> bloque del comentario (decisión 89).
>
> **Enmendado el 2026-09-29 (quinta vez)**, a petición del responsable, al usar el producto: (1) en las
> tablas de resumen, **la columna «Ticket» enseña sólo el número** y **la categoría y las etiquetas se
> agrupan en una columna nueva, «Clasificación»** (decisión 86); (2) **el responsable se reasigna con un
> botón junto a su nombre**, que abre un modal —el campo «Asignar a» suelto desaparece—; (3) **los
> botones de volver van a la izquierda y son sólo un icono** (una flecha del proyecto, cuadrada); (4) el
> **área de trabajo es más ancha** en pantallas grandes —`px-5`, o `px-10` a partir de tablet—, sin llegar
> a ocupar todo el ancho; y (5) **los dos resúmenes del motor se refrescan solos** mientras están
> pendientes, después de **cualquier** cambio (antes sólo tras pulsar «Volver a resumir», y el campo se
> quedaba en «Generando…» hasta recargar: lo reportó el responsable como un cuelgue). **La mención, de
> paso, se ve como una mención**: con su `@` delante y con un color del tema que la distingue.
>
> **Enmendado el 2026-09-29 (cuarta vez)**, al aclararse quién etiqueta (decisión 85): el **modal de
> etiquetas** lo tienen **los dos equipos**, cada uno desde el ticket en el que trabaja —Soporte en el
> principal, Desarrollo en el interno—, y **las etiquetas son del principal** (el interno las hereda):
> cambiarlas desde cualquiera de los dos cambia las mismas. Es lo que quiere decir «clasificar es de los
> dos equipos», y **no** toca la regla de que Desarrollo no escribe el texto del principal.
>
> **Enmendado el 2026-09-29 (tercera vez)**, al comprobarse que dos decisiones se contradecían: **el
> catálogo se cura** (decisión 84). Escribir una etiqueta nueva en un ticket **ya no la crea**: quien
> etiqueta **elige de las que hay** —el modal de la ficha sólo ofrece las del catálogo— y **los nombres
> nuevos los estrena el Administrador** en la pantalla del catálogo. A quien no lo es, poner una
> etiqueta que no existe le responde **422** con su aviso, en vez de estrenarla sin querer.
>
> **Enmendado el 2026-09-29 (segunda vez)**, a petición del responsable: **el usuario no pone etiquetas**
> (decisión 82) —**el alta no lleva ese campo** y **la línea de etiquetas de la ficha no se le enseña**—,
> porque etiquetar es de **Soporte y Desarrollo**; y **se etiqueta en un modal**: **todas las etiquetas
> del catálogo** con cuántos tickets lleva cada una, un **buscador**, **casillas** para marcar y
> desmarcar varias a la vez —las del ticket salen marcadas— y **«Guardar» al final** (decisión 82). El
> **catálogo lo mantiene sólo el Administrador** (decisión 83), así que en la pantalla del catálogo los
> botones de crear, renombrar y retirar **no se le enseñan a Soporte**. Y **salir de la aplicación se
> confirma**: el botón de salir del menú abre su ventana antes de cerrar la sesión.
>
> **Enmendado el 2026-09-29**, **corrección del responsable** sobre la decisión 78: el desplegable del
> ticket **no es una lista de acciones, es para cambiar de estado**, y enseña **la lista de estados** —la
> misma en los principales y en los internos—, con **el estado actual como primera opción**, marcado y
> sin poder elegirlo, y **sólo los estados a los que el ticket puede pasar ahora** (decisión 80). Elegir
> uno abre **una sola ventana** que dice a cuál se pasa y, **si esa transición necesita un texto, lo pide
> ahí**: el motivo al escalar, qué se ha hecho al resolver, la pregunta al ponerse en espera y **el
> comentario al cerrar, que es obligatorio** (decisión 81). Y las acciones que tenían nombre **son esos
> mismos estados**: en un interno, «En espera» devuelve el caso a Soporte sin cerrarlo y «Cerrado» sin
> resolver lo devuelve y lo cierra; en un principal, «En espera» es preguntarle al usuario y «Cerrado»
> desde cerrado es reabrirlo.
>
> **Enmendado el 2026-09-28 (sexta vez)**, a petición del responsable: **la entrada, cuadrada**
> (decisión 79). El recuadro del logo no medía lo mismo que el formulario, y el idioma y el tema medían
> cada uno lo suyo —**medido: logo 136 px, formulario 329, controles 256, idioma 104 y tema 140**—. La
> causa era doble: el logo llevaba **tope de ancho propio** (288 y 320 px) y, sobre todo, **la columna
> de la entrada centra a sus hijos**, así que un componente **encoge hasta lo que mida su contenido**
> aunque por dentro ocupe el 100 %. El ancho se pide ahora **en el elemento** —el logo, la tarjeta y los
> controles—, y los dos desplegables **se reparten a medias**. Medido después: **380 = 380 = 380**, y
> 184 cada desplegable. Hay una prueba que lo mide.
>
> **Enmendado el 2026-09-28 (quinta vez)**, a petición del responsable: **las acciones del ticket, en un
> solo desplegable** (decisión 78). Eran cuatro botones sueltos en la ficha —«Preguntar al usuario»,
> «Escalar», «Resolver», «Cerrar» en el principal; «Pedir algo a Soporte», «No es un cambio de código»,
> «Resolver» y «Cerrar» en el interno— y el responsable no les veía el sentido sueltos. Ahora es **un
> desplegable de acciones**, nombradas **por lo que hacen** —no por el estado al que llevan, que sigue
> en su chip—, y **las que piden algo (el motivo, la pregunta) siguen abriendo su ventana**. Vive **en
> la fila de arriba cuando se ve un solo ticket** y **en la cabecera de cada columna cuando se ven los
> dos**. Y **«Escalar a Desarrollo» pasa a «Escalar»**.
>
> **Enmendado el 2026-09-28 (cuarta vez)**, a petición del responsable: **las ventanas salen centradas**
> (decisión 77). No salían: el `<dialog>` nativo se centra con el `margin: auto` de la hoja del
> navegador, y **la preparación de Tailwind pone `margin: 0` a todo**, así que aparecían en la esquina
> —medido: `x: 0, y: 0`—. Se les devuelve el `auto`, con altura máxima y desplazamiento propio para que
> una ventana alta no se salga. Y **el ✕ de la ventana se llama «Cerrar»**, no «Cancelar»: se llamaba
> igual que el botón de cancelar de dentro, y eran dos botones con el mismo nombre.
>
> **Enmendado el 2026-09-28 (tercera vez)**, a petición del responsable: **los iconos del ticket son los
> del proyecto** (decisión 76). El de copiar el número era el carácter `⧉` —un cuadrado con otro detrás,
> que a 16 px se lee como un cuadrado— y el lápiz de editar, `✏`, que **según la fuente del equipo sale
> como emoji de colores**: los dos se parecían. Ahora son **svg propios** —como los del menú y sin
> bibliotecas—: **dos hojas superpuestas** para copiar, un **visto** para confirmarlo y un **lápiz
> inclinado** para editar. El botón los recibe **por su nombre**, no por su carácter.
>
> **Enmendado el 2026-09-28 (segunda vez)**, a petición del responsable: en las listas, **el nombre del
> módulo, los filtros y «Nuevo ticket» van en una sola fila** (decisión 75), que era lo que faltaba para
> no gastar un renglón de cabecera: la tabla gana ese alto. En un ancho corto se reparten solos —el
> nombre, los filtros y el botón a la derecha si caben, y lo que no quepa, debajo—.
>
> **Enmendado el 2026-09-28**, a petición del responsable, al usar el ticket: **todo se edita donde se
> muestra** (decisión 74). En la ficha, la **categoría** es un desplegable en su propia línea, las
> **etiquetas** son fichas con su aspa y un campo pequeño para añadir, los **observadores** llevan su
> **«Añadir»** con el buscador de personas además del aspa de cada uno, y el **asunto** y la
> **descripción** llevan su lápiz que los convierte **en el sitio** en un campo con «Guardar» y
> «Cancelar». **Se van los formularios aparte** —el de «cambiar la clasificación» y el diálogo de
> editar el ticket—, que era lo que hacía que no se encontrara. Además: el número y el estado del
> ticket **se enseñaban dos veces** —los pintaba la cabecera de la página y la del propio ticket— y
> ahora los pinta sólo la de la página; y en el editor se puede **etiquetar escribiendo `@`**, con la
> lista filtrada **por nombre o por correo** mientras se escribe.
>
> **Enmendado el 2026-09-27 (quinta vez)**, a petición del responsable, al usarlo: **la lista lleva un
> solo botón «Buscar»**, que abre un modal con todos los criterios —texto, vista, tipo, estado,
> categoría y etiqueta—, y en su cabecera queda **el resumen de lo que está filtrado** con su botón de
> quitar; las tres filas de filtros se comían la pantalla y **la de categorías se salía** (sección 3.3).
> El título **«Qué se puede hacer»** de la tarjeta de acciones de la ficha **se retira**, porque se
> malinterpretaba. La **vista del ticket** pasa a ser **la del papel** —principal para Soporte, interno
> para Desarrollo—, «Los dos» enseña **sólo las dos conversaciones** sin las fichas, y el botón
> **«Retirar» de una categoría sólo se le enseña al Administrador**. La pantalla del catálogo pasa a
> tener **dos mitades con el mismo peso**, con las etiquetas pudiéndose **crear, renombrar y retirar**.
>
> **Enmendado el 2026-09-27 (cuarta vez)**, a petición del responsable: en Configuración, **cada
> método de entrada enseña sólo su configuración** —«Cuentas de la aplicación» no enseña ninguna, el
> directorio sale con la organización y el reino con Keycloak—. Antes se veían las dos a la vez
> (sección 3.8).
>
> **Enmendado el 2026-09-27 (tercera vez)**, a petición del responsable: **copiar el número** de un
> ticket es un **botón de sólo icono**, y al copiar **se cambia por un visto verde unos segundos**. El
> inventario de componentes (sección 6.3) recoge las dos cosas que hicieron falta: una **cuarta forma**
> de botón —la verde de la confirmación— y que un botón pueda ser **sólo un icono**, con su nombre
> accesible obligatorio.
>
> **Enmendado el 2026-09-27 (segunda vez)**, al entrar **las categorías y las etiquetas**
> (`docs/modules/tickets.md`, decisiones 64 a 68): el **alta pide la categoría** —es obligatoria: sin
> ella no se crea el ticket— y acepta **etiquetas** en minúsculas con guion medio y sugerencias; la
> **lista** las enseña como **chips en el renglón, debajo del número** (y en las tarjetas del móvil),
> con **chip de categoría y de etiqueta para filtrar**; la **ficha** deja cambiarlas a quien puede
> editar el ticket; y el menú del Administrador y de Soporte estrena la pantalla **«Categorías y
> etiquetas»**, que es donde se mantiene el catálogo.
>
> **Enmendado el 2026-09-27**, al entrar **las menciones y los observadores** (`docs/modules/tickets.md`,
> decisiones 58 a 63): el editor estrena un botón **«Etiquetar»** —y basta escribir `@`— que mete el
> nombre de un técnico o de un desarrollador **dentro del comentario, resaltado**; eso lo convierte en
> **observador** del ticket, que **no es lo mismo que el asignado**. «Mis tickets» lleva un chip nuevo
> —actualmente **Todos · Asignados · Observo**, según la sección 46—, la ficha tiene su lista de **«Observadores»** con su aspa para
> quitarlos (Soporte y Desarrollo), y quitar a alguien **se cuenta en la línea de tiempo**.
>
> **Enmendado el 2026-09-27**, al pedir el responsable **dos columnas nuevas en la lista de tickets**,
> que redacta un motor de inteligencia artificial (`docs/modules/ai.md`): **«Motivo»** —de qué va el
> ticket— y **«Última acción»** —qué fue lo último que pasó—, las dos **en el idioma de quien mira**.
> Y con ellas, **dos nombres que cambian**: la columna «Número» pasa a llamarse **«Ticket»** y «Última
> actualización» pasa a **«Fecha actualización»** (sección 3.3).
>
> **Enmendado el 2026-09-26**, al pedir el responsable que **la imagen y el vídeo no ocupen todo el
> ancho**: en la descripción y en los comentarios se ven **como mucho a 480 × 360** —una captura a
> pantalla completa se comía la conversación— y **al pulsarlos se abren en el visor grande**, que es
> donde se ven enteros y se descargan. **El vídeo pasa a ser una miniatura con su botón de
> reproducir**, y se reproduce en el visor (sección 6.5). Es la decisión 55 de `tickets.md`.
>
> **Enmendado el 2026-09-26**, al pedir el responsable que **la bandeja sea «lo mío»** y que existan
> **dos módulos nuevos para ver todo**: la Bandeja pasa a ser **Mis tickets** —su semántica vigente
> está en la sección 46—, y nacen **Tickets principales** y **Tickets internos**, sólo para
> Soporte y Desarrollo (secciones 3.2, 3.3 y 3.7). **«Nuevo ticket» sale del menú** —se confundía con
> un módulo— y se queda donde se crea, en el botón de la bandeja. **«Asignármelo» desaparece de la
> bandeja**: la reasignación se hace dentro del ticket, que es donde se ve a quién se le pasa. Las
> filas de la bandeja dicen **«Abrir»** y no «Abrir el ticket», sobra el texto «Lo que lleva más tiempo
> esperando va primero…» —la regla sigue, callada—, y **el idioma pasa a ser un desplegable como el del
> tema**, con su ancho y sin etiqueta a la vista (secciones 3.1 y 6.3).
>
> **Enmendado el 2026-09-26**, al implementar el editor con adjuntos: **pegar texto pega
> sin formato** —con los saltos de línea—, y **un archivo pegado sí entra como adjunto**, porque pegar
> el HTML de un correo traería etiquetas que el backend rechaza y se perdería el comentario entero por
> algo que la persona no ha escrito (sección 6.5). Y **el hueco de un archivo que no subió se ve
> marcado**, en el sitio donde iba.
>
> **Enmendado el 2026-09-26**, al pedir el responsable que **los adjuntos se vean dentro del texto** y
> que **la vista previa se vea también al escribir**: el cuadro de escribir pasa a ser un **editor con
> formato** (sección 6.5), el **visor** se agranda y gana **Descargar** y **Abrir en una pestaña**, y
> **se retira** la pieza de subida de archivos, porque adjuntar ya no es un recuadro aparte.
>
> **Enmendado el 2026-09-25**, al pedir el responsable que **la versión del sistema se
> vea**: en la fila de salir del menú lateral, alineada a la derecha, y **también en el pie de la
> pantalla de entrada** (secciones 3.1 y 6.4). **Es elección suya y no mi recomendación**: yo proponía
> sólo el menú, y decidió que también se vea antes de entrar, que es donde sirve para decir qué versión
> se está mirando al reportar un fallo.
>
> **Pasa a as-built el 2026-09-25**, y **enmendado el mismo día**: la **tarjeta de cómo se entra** en
> la pantalla de Configuración —el método, las dos configuraciones y sus dos pruebas de conexión— y lo
> que la pantalla de entrada ofrece según el método (sección 4.4). La forma de las pantallas que este
> documento fija está toda construida.
>
> **Enmendado el 2026-09-25**: la instalación tiene **nombre propio**, configurable, y
> **es él quien lo lee todo el mundo** —la pantalla de entrada, el menú lateral y **la pestaña del
> navegador**— en lugar del texto de fábrica (sección 6.4 y `docs/modules/settings.md`, sección 5.7).
> **El responsable corrigió dos cosas al pedirlo**: los correos y su remitente no cambian.
>
> **Enmendado el 2026-09-25**, al implementar las tres acciones que dependían del
> directorio: en las pantallas de usuarios, **reactivar sólo se ofrece donde puede funcionar** —en una
> cuenta de AD sí, y en una de Keycloak no, con lo que pasa contado en su lugar— y **la nota del
> origen** distingue lo que ocurre: AD puede consultarse desde el alta cuando es el método vigente;
> Keycloak continúa creando la cuenta al entrar (`docs/modules/users.md`, sección 15).
>
> Aprobado por el responsable el 2026-09-22, tras repasarlo en forma de preguntas mientras se
> escribía: **no hay barra superior** —el menú lateral lleva el producto, las opciones del papel y,
> abajo y separados, los controles (perfil, tema y salir)— y el papel `administrador` **ve los
> tickets en sólo lectura**. Fija los principios (intuitivo antes que bonito), la forma de la
> aplicación, la experiencia de los cuatro papeles, la vista doble principal/interno, el lenguaje
> visual con **Tailwind v4 y componentes propios**, los ocho temas y lo que se espera en cada
> dispositivo.
>
> Con este documento, la **documentación está completa**: siete documentos que cubren qué se
> construye, cómo se comporta, cómo se ve y cómo se despliega.
> **Enmendado el 2026-09-23**, al escribir el armazón de la sesión: se añade **de dónde sale el
> idioma de la interfaz** (principio 6: del navegador, y si no pide ninguno de los dos, inglés).
> Es una **corrección del responsable**: yo había propuesto español por defecto
> (`docs/modules/auth.md`, decisión 25).
>
> **Enmendado el 2026-09-24**, al construir el armazón: la sección 3.1 dice que **el
> estado plegado del menú se recuerda** y que **el menú enseña sólo lo que existe** —las entradas de
> la tabla 3.2 llegan con sus pantallas—, y la 6.4 recoge el logo en el menú.
>
> **Enmendado el 2026-09-24**, al entrar el logo de la institución: la sección 6.4 dice **dónde se
> enseña la marca** (entrada, menú lateral y Configuración) y **cómo se enseña cualquier logo**
> —entero, sin recortar y en un recuadro con el borde del tema—, para que sirva igual uno con fondo
> opaco que uno con transparencia. El logo y su reemplazo son del módulo `settings`
> (`docs/modules/settings.md`).
>
> **Enmendado el 2026-09-24**, al construir el editor de los correos: la sección 3.8 fija
> **su pantalla**, con los dos idiomas al lado, los marcadores que se insertan al pulsarlos y una vista
> previa que **renderiza el backend** con los datos de ejemplo de la prueba, para que lo que se ve sea
> lo que va a salir.
>
> **Enmendado el 2026-09-24**, al construir las pantallas de tickets: la sección 3.7 fija
> **las cuatro rutas** y quién ve cada una, y recoge que **el detalle es la misma pantalla para los
> cuatro papeles** y que **el prefijo y el reparto se configuran en Configuración**, no en una pantalla
> de `tickets`. La tabla de la sección 3.2 estrena **«Nuevo ticket» para Soporte** —tres entradas—:
> puede crear un ticket en nombre de otra persona y ese permiso necesitaba su puerta (**corrección del
> responsable**). El escalado se escribe **en un diálogo**, porque el motivo es obligatorio y la acción
> es deliberada, y los adjuntos se ven **con vista previa** cuando son imágenes o PDF.
>
> **Enmendado el 2026-09-24**, al empezar los tickets: la sección 3.3 añade el **chip de
> tipo** (Principales · Internos · Todos) en la bandeja del Administrador, que es el único papel que
> ve los dos tipos de ticket.
>
> **Enmendado el 2026-09-24**, al construir las pantallas de usuarios: la sección 3.6
> fija **las tres pantallas de `users`** —la lista, la ficha y el perfil propio—, sus direcciones y
> quién ve cada una. Las cinco decisiones son del responsable, del 2026-09-24, y salen de un hallazgo:
> **Soporte puede cambiar el nombre y los apellidos de otra cuenta, pero no ve la ficha de nadie**, así
> que ese permiso se ejerce **en línea, desde la propia lista**. La sección 3.1 recoge además que **el
> menú mide la zona visible y se desplaza por dentro** —medía el alto del documento, y en móvil dejaba
> el botón de salir fuera de la pantalla—, y la 6.3 que la tarjeta **elige el nivel de su título**,
> porque una pantalla con tres `h1` no se navega con un lector de pantalla. Los dos los encontró la capa
> de Playwright, no las pruebas de unidad.
>
> **Enmendado el 2026-09-23**, al montar el tema: la sección 6.2 dice **cómo se implementa**
> (un atributo en `html`, las paletas escritas una vez con `light-dark()`, la elección recordada en
> el navegador y el tema aplicado por código nuestro **porque la CSP de producción no admite scripts
> incrustados**). El conmutador tiene **dos opciones, claro y oscuro**: **«automático» no se muestra
> ni se elige**, porque seguir al sistema es no haber elegido. Las seis paletas fijas quedan para el
> armazón completo, cuando se aprueben sus nombres y sus colores.

## 1. Alcance de este documento

Fija **cómo se ve, cómo se navega y qué siente cada papel**: la estructura de las pantallas, el
lenguaje visual, el comportamiento en cada dispositivo y la experiencia de los cuatro papeles.

**No repite** lo que cada papel puede hacer (`docs/usuarios-y-permisos.md`), ni los recorridos
(`docs/flujos.md`), ni el modelo de datos (`docs/modules/tickets.md`). Cuando este documento y aquellos no
coincidan, mandan ellos: aquí sólo se decide cómo se presenta.

Está **aprobado**, así que habilita escribir código (Regla 0). La sección 11
recoge lo que he propuesto yo.

## 2. Los principios, en orden

El orden importa: cuando dos choquen, gana el de más arriba.

1. **Intuitivo antes que bonito.** Si hay que elegir entre algo que se entiende sin explicación y
   algo que se ve mejor, gana lo primero. Un producto que hay que explicar es un producto que se
   usa mal.
2. **Sin ruido** (el criterio del proyecto, `docs/propósito-y-alcance.md`): lo que no aporta, no
   está. Ni animaciones decorativas, ni contadores que nadie mira, ni avisos que se cierran sin
   leer.
3. **Cada uno ve su trabajo, no la aplicación entera.** El usuario no necesita saber que existe un
   ticket interno, ni un desarrollo necesita ver la administración de usuarios.
4. **El estado se lee de un vistazo.** Mirar la lista y saber qué está esperando a quién.
5. **Nada importante se hace por accidente.** Cerrar, escalar o borrar se confirman; lo demás, no.
6. **El idioma es global.** Una instalación nueva abre `/setup` en inglés y su primer control puede
   cambiar toda la pantalla. Después manda siempre el idioma guardado por la instalación, también
   antes de entrar y en sesiones abiertas; no hay selector ni preferencia por cuenta.
7. **Lo que falla, se explica.** Un error dice qué pasó y qué hacer, en el idioma de quien lo lee.

## 3. La forma de la aplicación

### 3.1 El armazón

```text
┌───────────┬───────────────────────────────────────────────┐
│  producto │                                               │
│           │                                               │
│  opciones │  contenido                                    │
│  del papel│                                               │
│           │                                               │
│  ───────  │                                               │
│  controles│                                               │
└───────────┴───────────────────────────────────────────────┘
```

- **No hay barra superior.** El menú lateral lo lleva todo, y el contenido gana todo el ancho: en
  una mesa de ayuda la pantalla se llena de listas, conversaciones y adjuntos, y una franja
  superior es espacio que se le quita a lo que importa.
- **El menú lateral tiene tres zonas**, separadas visualmente:
  1. **Arriba**: el producto —**el logo de la instalación y el nombre**—. Es lo único que no se mueve.
  2. **En medio**: las **opciones del papel** (sección 3.2). Es lo que cambia de una persona a otra.
  3. **Abajo, separado por una línea**: los **controles**, que son los mismos para todos —**tu
     nombre (y tu perfil)**, **el idioma**, **el tema** y **salir**—. Es el mismo reparto que usa
     Calibyou, y evita mezclar «lo que hago aquí» con «quién soy».
- **El idioma y el tema son el mismo control, uno debajo del otro** (decisión del responsable,
  2026-09-26): los dos son **desplegables**, ocupan **el mismo ancho** y **ninguno lleva la etiqueta a
  la vista** —se lee «Español» o «Claro»—, porque el valor ya dice qué es. El idioma se cambia
  **eligiendo cualquier valor de la lista**, y el nombre accesible sigue ahí, oculto, para quien navega
  con lector de pantalla. Antes el idioma eran dos botones segmentados y el tema un desplegable: dos
  formas para lo mismo, y dos anchos distintos en la misma columna.
- **La versión del sistema se lee en la fila de salir, alineada a la derecha** (2026-09-25): es un
  dato de la instalación, no una opción, así que **no es un botón ni ocupa una fila propia**: va con el
  control que ya está en esa fila. En el menú plegado se pierde con el texto, que es lo correcto: un
  número suelto al lado de un icono no dice nada. **En móvil se lee dentro del cajón**, que es donde
  vive esa fila.
- En PC el menú se puede **plegar** y queda una tira de iconos; en tablet se plegará solo. **El
  estado se recuerda** en el navegador, como el tema: quien trabaja con el menú plegado no tiene que
  plegarlo cada mañana.
- **Mide 13 rem (208 px)** y **3,5 rem (56 px) plegado**. La reducción entrega más ancho al
  contenido sin estrechar los blancos táctiles. El nombre de la instalación se muestra completo,
  centrado y en las líneas que necesite debajo del logo; no ensancha el menú y la cabecera puede
  crecer verticalmente.
- **El menú enseña lo que existe**: cada entrada llega **con su pantalla**, no antes —un menú que
  lleva a ninguna parte es peor que un menú corto—, y no se enseña a quien no le toca. El usuario ve
  una sola opción, **Mis tickets**, y eso es todo lo que necesita: pedir y seguir lo suyo.
- **El menú mide la zona visible, no el documento, y se desplaza por dentro si su contenido no cabe**
  (decisión del responsable, 2026-09-24, al construir las pantallas de usuarios). El cajón medía el
  alto del documento: en un móvil con barra de navegador lo visible son 839 px de 890, así que **el
  botón de salir quedaba por debajo de lo que se ve**, imposible de pulsar; y en PC, con una lista
  larga, la barra se estiraba a todo el contenido y sus controles acababan al final de la página. En
  PC la barra va **pegada arriba** mientras el contenido se desplaza, que es lo que la sección 7 llama
  «fijo».
- **En móvil sí hay una tira superior mínima**, con el botón de menú y el nombre del producto: sin
  ella no habría dónde poner el disparador. Todo lo demás vive en el menú, que se abre como cajón.

### 3.2 Las entradas del menú, por papel

| Papel | Opciones del menú |
| --- | --- |
| **Usuario** | Mis tickets |
| **Soporte** | Mis tickets · Tickets principales · Tickets internos · **Clasificación** · Usuarios |
| **Desarrollo** | Mis tickets · Tickets principales · Tickets internos |
| **Administrador** | Bandeja (sólo lectura) · **Clasificación** · Usuarios · Configuración |

**Cinco entradas como máximo** en el caso más cargado —Soporte— y cuatro el Administrador, y **los
controles no cuentan**: el perfil, **el idioma**, el tema y salir van abajo y son iguales para todos.
**Eran cuatro hasta el 2026-09-27**, que entró la pantalla hoy llamada **«Clasificación»** con las categorías: es una
pantalla de tickets —el catálogo son datos de los tickets, no configuración de la instalación— y por eso
vive en el menú y no dentro de Configuración (`docs/modules/tickets.md`, decisión 64). Si algún papel
llegara a necesitar más opciones, es señal de que algo se ha complicado en otro sitio.

**«Nuevo ticket» ya no está en el menú** (decisión del responsable, 2026-09-26): se confundía con un
módulo, y crear un ticket es una acción de la bandeja, no un sitio al que ir. El botón vive **dentro
de la bandeja** —arriba y también cuando está vacía—, así que **no se pierde nada**: la ruta
`/tickets/new` sigue existiendo y sigue siendo la misma pantalla.

**Soporte tiene cinco** porque, además de lo suyo, es el único papel —con Desarrollo— que ve **todo
lo que hay**, y ahora eso se ve en dos sitios de una vez: los principales y los internos. Antes estaba
todo junto en la bandeja, que es lo que dejaba de funcionar cuando la bandeja pasó a ser «lo mío».

### 3.3 Las bandejas

**Tres bandejas y una pantalla.** La misma pantalla de lista sirve para las tres —cambia el contenido
y el título—, porque separarlas en tres pantallas sería triplicar la tabla, los filtros y la paginación
para cambiar una condición (`docs/modules/tickets.md`, sección 8).

| Bandeja | Quién | Qué enseña |
| --- | --- | --- |
| **Mis tickets** (`/tickets`) | Usuario, Soporte y Desarrollo | Usuario: sus principales solicitados. Equipos: lo asignado o actualmente observado, limitado a su nivel operativo |
| **Tickets principales** (`/tickets/main`) | Soporte y Desarrollo | **Todos** los principales |
| **Tickets internos** (`/tickets/internal`) | Soporte y Desarrollo | **Todos** los internos |
| **Bandeja del Administrador** (`/tickets`) | Administrador | Los dos tipos, con su chip y de sólo lectura |

- **«Lo mío» es una relación vigente** (actualizado el 2026-10-19): para Soporte y Desarrollo es lo
  asignado o actualmente observado. Abrir o comentar no añade el ticket, y dejar de observarlo lo
  elimina salvo que permanezca asignado. Soporte ve aquí sólo principales y Desarrollo sólo internos.
- **La lista es una tabla** en PC y **tarjetas** en móvil, donde una tabla de ocho columnas no se lee.
  Las columnas son: **Ticket** · **Asunto** · **Motivo** · **Última acción** · **Clasificación** ·
  **Estado** · **Solicitante** · **Responsable** · **Fecha actualización** (los dos nombres de los
  extremos y las dos columnas del medio son del 2026-09-27: «Número» era lo que se lee como un ticket, y
  «Última actualización» es una fecha, no un suceso; **«Clasificación» es del 2026-09-29** y lleva la
  categoría y las etiquetas, decisión 86 — y **«Ticket» se queda con el número solo**—).
- **Todo se edita donde se muestra** (decisión del responsable, 2026-09-28): la categoría, las
  **etiquetas**, los observadores, el asunto y la descripción **se cambian en el sitio donde se leen**,
  no en un formulario aparte. El reparto, para **el texto y la categoría**, es el de siempre: **el
  solicitante en su ticket y Soporte**; Desarrollo no escribe el principal y el Administrador no toca
  nada, así que **a ellos no se les enseña el lápiz**. **Lo único que sí hace Desarrollo en un principal es seguir a quien lo sigue**: la
  lista de observadores —añadir y quitar— es de **los dos equipos** (decisión 63), porque seguir un
  ticket no es escribirlo. **Y las etiquetas también son de los dos** (decisión 85): se ponen, se cambian
  y se quitan **desde el ticket en el que trabaja cada uno** —Soporte en el principal, Desarrollo en el
  interno—, y como **son del principal** (el interno las hereda, decisión 67), cambiarlas desde cualquiera
  de los dos cambia las mismas. Lo encontró la prueba de interfaz al comprobar que no le dejábamos nada —la interfaz no ofrece lo que el backend rechazaría—. Cada control dice lo que
  hace y **el cambio se guarda al confirmarlo**, nunca al escribir.
- **Etiquetar escribiendo `@`** (decisión del responsable, 2026-09-28): en el editor, escribir `@` abre
  la lista **filtrada por lo que se sigue escribiendo**, y vale **el nombre o el correo** —el buscador
  enseña **los dos**, que es lo que permite distinguir a dos personas que se llaman igual—. La mención
  **se lee con el nombre**, como hasta ahora, y el botón «Etiquetar» sigue estando para quien no se
  acuerde del atajo.
- **Mandar el enlace pregunta antes** (decisión del responsable, 2026-09-28): el enlace es el de alta
  para quien nunca ha entrado y el de **recuperación** para quien ya tiene contraseña —la que tuviera
  deja de valer—, así que la pantalla **pregunta antes con una ventana**, y el aviso dice **cuál de los
  dos es** y cuánto caduca: 24 horas el de alta, 1 hora el de recuperación. No se usa un `confirm()` del
  navegador, como en el resto de las acciones que preguntan.
- **En la cabecera del ticket, el número del interno** (decisión del responsable, 2026-09-28): abriendo
  un interno, esa fila enseñaba sólo el número del principal y no se sabía cuál de los dos se estaba
  mirando. Va en apagado, porque es la referencia y no el título, y **el botón de volver es un icono con su nombre accesible «Volver»**
  —decía «Volver a la bandeja», y a la bandeja se vuelve desde muchos sitios—.
- **Los observadores se añaden también a mano** (decisión del responsable, 2026-09-28): además de
  llegar por una mención, se pueden **añadir desde su lista** con el buscador de personas y **quitar**
  con su aspa. Quien se añade así **queda igual** que quien llegó por una mención: todos son
  observadores, y la lista no distingue de dónde vino cada uno.
- **La tarjeta de acciones de la ficha no lleva título** (decisión del responsable, 2026-09-27): decía
  «Qué se puede hacer», que **se malinterpreta** —parece una promesa de lo que el sistema hará, y es
  sólo dónde están los botones—. Los botones siguen diciendo lo que hacen, que es lo que hacía falta.
- **La vista del ticket, por papel** (decisión del responsable, 2026-09-27): al abrir un ticket,
  **Soporte ve el principal** y **Desarrollo el interno**; los dos pueden cambiar a **Principal · Los
  dos · Interno** cuando quieran, y **abrir el número de un interno enseña el interno** aunque seas
  Soporte. **«Los dos» enseña sólo las dos conversaciones** —la columna de cada ticket—: con las cuatro
  columnas (dos conversaciones y dos fichas) «se ve muy saturado». Para actuar sobre uno, se elige ese.
- **Copiar el número es un botón de sólo icono** (decisión del responsable, 2026-09-27): el
  portapapeles, y **al copiar se cambia por un visto verde durante dos segundos** y vuelve solo. **El
  dibujo es el del proyecto** —dos hojas superpuestas—, y no se confunde con el lápiz de editar
  (2026-09-28). El
  nombre accesible **cambia con el estado** («Copiar el número» → «Copiado»), así que quien no ve el
  color sí oye la confirmación; el verde sale del tema (`--exito`), no escrito a mano.
- **El renglón lleva sus chips de clasificación** (2026-09-27, **movidos el 2026-09-29**, decisión 86):
  **un chip de categoría** —el *qué es* del ticket— y los de sus etiquetas. En la **tabla** viven en su
  propia columna, **«Clasificación»**; en las **tarjetas del móvil**, donde no hay columnas, siguen
  debajo del número. Un chip ocupa lo que mide su nombre y se lee de un vistazo. **La categoría no se
  recorta**: es una palabra o dos.
- **Un solo botón de búsqueda, y todo dentro** (decisión del responsable, 2026-09-27): la lista lleva
  **«Buscar»**, y el modal que abre tiene **todos los criterios** —el texto, la vista, el tipo, el
  estado, la categoría y la etiqueta— con sus botones de **buscar** y **limpiar** dentro. En su lugar
  no queda nada más que **el resumen de lo que está filtrado** —«3 filtros»— con su botón de quitar:
  la lista se ve limpia y **se entiende por qué salen pocos tickets**. Antes eran **tres filas de
  filtros** para cuatro criterios, y la de categorías **se salía de la pantalla** cuando el catálogo
  crecía; un modal no crece, se llena.
- **Y las dos (categoría y etiqueta) entran en la búsqueda**: buscar «red» encuentra la categoría «Red»
  y la etiqueta `red`, que es para lo que se pidieron.
- **«Motivo» y «Última acción» son redacciones del motor de IA** (`docs/modules/ai.md`), y se enseñan
  **recortadas a dos líneas**, porque son un resumen para barrer la lista y no el texto del ticket: el
  completo se lee en la ficha. Mientras el motor no ha contestado, dicen **«Generando…»**; si no hay
  motor o falló, lo dicen y **no se inventan nada** (decisiones 2 y 8 de `docs/modules/ai.md`).
- **En las tarjetas del móvil van las dos**, como dos líneas más: en móvil es donde más se agradece
  saber de qué va un ticket sin abrirlo.
- **La fila lleva «Abrir»**, no «Abrir el ticket» (2026-09-26): en una columna de acciones la palabra
  «ticket» sobra, y la fila entera ya habla de un ticket.
- **No hay «Asignármelo»** (2026-09-26). Se quitó de la tabla y de las tarjetas: **la reasignación se
  hace dentro del ticket**, en su ficha, donde se ve a quién se le pasa, y con las reglas de siempre
  (`docs/modules/tickets.md`, decisión 50). El atajo de la lista asignaba sin más, y eso es lo que
  dejaba sin hacer la reasignación a otro compañero.
- Los filtros son **chips visibles** encima de la lista, no un formulario escondido: estado, texto y
  nada más.
- **«Mis tickets» lleva en la primera fila el selector —Todos · Asignados · Observo—** para Soporte
  y Desarrollo (actualizado el 2026-10-19), y **arranca en «Todos»**. «Todos» es la unión sin
  duplicados de asignación y observación vigentes; el selector visible no se cuenta como búsqueda.
- **El nivel no se elige en «Mis tickets»**: Soporte ve sólo principales y Desarrollo sólo internos,
  incluso cuando la cuenta acumula capacidad Admin. **El usuario no lleva selector** y el
  Administrador puro mantiene su bandeja general de sólo lectura con el selector de tipo.
- **Las dos bandejas del «todo» no llevan chip de tipo** por lo de arriba: «Tickets principales»
  enseña principales y «Tickets internos», internos. Sí llevan el chip de **estado** y la búsqueda.
- El **estado va con color y con texto**: el color ayuda a barrer la lista, el texto dice qué es.
  Nunca sólo color (sección 8).
- Orden por defecto: **lo que lleva más tiempo esperando, primero**, dentro de lo que no está
  cerrado. Lo cerrado no estorba. **La regla sigue escrita aquí y ya no se escribe en la pantalla**
  (2026-09-26): el orden se ve en la lista, y una frase que explica lo que ya se ve es ruido.

### 3.4 El detalle del ticket: la pieza central

Es la pantalla donde se pasa el tiempo, así que se lleva la mayor parte del diseño.

```text
┌──────────────────────────────────────────────────────────────┐
│  ACME-2026-0042 · asunto                    [estado]         │
├──────────────────────────────┬───────────────────────────────┤
│  conversación                │  ficha: solicitante,          │
│  (comentarios, adjuntos,     │  responsable, fechas,         │
│   historial intercalado)     │  adjuntos                     │
│                              │  acciones                     │
├──────────────────────────────┴───────────────────────────────┤
│  escribir un comentario…                       [Enviar]      │
└──────────────────────────────────────────────────────────────┘
```

- **Una sola línea de tiempo**: los comentarios de las personas y lo que hizo el sistema, en orden.
  El historial no es una pestaña aparte: si cambió el estado, se ve ahí, entre los mensajes.
- **La ficha y las acciones van a la derecha** en PC: quién lo pidió, quién lo lleva, desde cuándo, y
  los botones de estado, asignación y escalado. En móvil bajan debajo, plegadas.
- **El número del ticket siempre visible y con un botón de copiar**: es lo que se dice en voz alta y
  lo que se pega en un correo.
- La caja de comentario está **siempre a la vista** al final de la conversación. Escribir es la
  acción más frecuente del sistema; esconderla detrás de un botón es un error de principiante.
- Un ticket **cerrado no admite comentarios** (`docs/flujos.md`): en su lugar aparece la explicación de
  por qué, y **reabrir se elige en el desplegable de acciones** (decisión 78), que es donde viven todas
  las acciones del ticket; no hay dos controles para lo mismo.

### 3.5 El usuario no lee jerga interna

El usuario **nunca** ve el ticket interno ni la palabra `escalado`. Y sus estados se le cuentan en
lenguaje llano:

| Estado interno | Lo que ve el usuario |
| --- | --- |
| `nuevo` | Recibido |
| `en progreso` | En curso |
| `en espera` | Esperándote |
| `escalado` | En curso *(el usuario no tiene por qué saber que ha pasado a Desarrollo)* |
| `resuelto` | Resuelto |
| `cerrado` | Cerrado |

### 3.6 Las pantallas de usuarios: la lista, la ficha y el perfil

Son **las primeras pantallas de producto**, y son tres, cada una con su dirección:

| Pantalla | Ruta | Quién la ve |
| --- | --- | --- |
| La lista de cuentas | `/users` | Soporte y Administrador |
| La ficha de una cuenta | `/users/{id}` | Administrador |
| El perfil propio | `/profile` | Cualquiera que haya entrado, menos la cuenta de fábrica |

- **La lista es una tabla en PC** —nombre, correo, papel, origen, estado y última entrada, más las
  acciones de la fila— y **tarjetas en móvil**: una tabla de seis columnas no se lee en un teléfono, y
  no se arrastra de lado (sección 7).
- **Los filtros son chips visibles encima de la lista**, como los de las bandejas (sección 3.3): papel,
  origen y estado, más la búsqueda por nombre o correo. Nada escondido detrás de un botón.
- **El orden es fijo —apellidos y nombre— y la lista se pagina**, con el total a la vista. **Ordenar
  por columna no se ofrece**: eso es de la tabla de las bandejas, que es donde hace falta.
- **El alta es un diálogo sobre la lista**, no una pantalla: seis campos y un botón. Al guardar, la
  lista se refresca con la cuenta nueva en su sitio, sin salir de donde se estaba.
- **Soporte edita el nombre y los apellidos en línea, desde la propia lista** (decisión del
  responsable, 2026-09-24). Soporte **no ve la ficha de nadie** (`docs/modules/users.md`, sección 4,
  decisión 2), y ese permiso tiene que poder ejercerse en algún sitio: el hallazgo fue justamente que,
  sin ficha, no había dónde. El Administrador cambia lo mismo desde la ficha.
- **La ficha** enseña la cuenta —nombre, correo, papel, origen, estado y cuándo entró por última
  vez— y **sus acciones**: cambiar el nombre, el correo, el papel y el idioma, desactivar o reactivar,
  cambiar el origen y **lanzar otra vez el enlace de contraseña**. Toda acción sensible dice lo que va
  a pasar antes de hacerlo (sección 4.4).
- **Reactivar se ofrece sólo donde puede funcionar**, y no es un detalle de forma: en una cuenta de
  **AD** el botón pregunta al directorio antes de devolver el acceso, y en una de **Keycloak** no se
  ofrece —se cuenta que **vuelve sola al entrar**, que es lo que pasa de verdad— porque su camino no
  tiene forma de comprobar quién sigue allí sin una cuenta de servicio en el reino
  (`docs/modules/users.md`, sección 5, punto 4). **La interfaz no ofrece lo que no se puede hacer.**
- **El botón de desactivar no se ofrece sobre uno mismo** (`docs/modules/users.md`, sección 7):
  desactivarse es quedarse fuera al instante y sin poder volver.
- **El origen se ofrece con sus tres valores**: `ad` se habilita sólo cuando es el método vigente y
  abre la búsqueda del directorio; `keycloak` permanece desactivado y explica que la cuenta se crea
  en el primer acceso (`docs/modules/users.md`, sección 15).
- **El perfil propio se abre desde tu nombre**, en la zona de controles del menú (sección 3.1), y
  lleva **tu nombre, tus apellidos y tu idioma**, más el enlace a cambiar la contraseña. El idioma se
  cambia aquí y en los controles: es el mismo ajuste, y es el que decide en qué idioma se te escriben
  los correos (`docs/modules/users.md`, sección 8).
- **La cuenta de fábrica no tiene perfil**: no está en la tabla de cuentas, así que **su nombre en el
  menú no lleva a ninguna parte** y la pantalla no se le enseña.

### 3.7 Las pantallas de tickets: la bandeja, el alta y el detalle

Cuatro rutas para las listas, y **el detalle es la misma pantalla para los cuatro papeles**: lo que
cambia es lo que cada uno puede hacer dentro, no la pantalla.

| Pantalla | Ruta | Quién la ve | Cómo se llama en el menú |
| --- | --- | --- | --- |
| **Mis tickets** | `/tickets` | Usuario, Soporte y Desarrollo | **Mis tickets** |
| **Tickets principales** | `/tickets/main` | Soporte y Desarrollo | **Tickets principales** |
| **Tickets internos** | `/tickets/internal` | Soporte y Desarrollo | **Tickets internos** |
| **La bandeja del Administrador** | `/tickets` | Administrador | **Bandeja** |
| El alta | `/tickets/new` | Usuario, Soporte y Desarrollo | — (se entra por el botón de la bandeja) |
| El detalle | `/tickets/{numero}` | Quien pueda verlo | — |
| **Clasificación** | `/tickets/categories` | Soporte (consulta el catálogo) y Administrador (lo mantiene) | **Clasificación** |
| El prefijo y el reparto | `/settings` | Administrador | **Configuración** |

- **Una sola pantalla de lista con cuatro nombres**: es el papel el que decide qué se ve dentro, y el
  documento ya lo llama distinto por papel. El nombre es lo único que cambia.
- **Las cuatro direcciones son del módulo `tickets`** y ninguna pide nada nuevo al backend: las listas
  ya se distinguen por su tipo y por de quién es el ticket (`docs/modules/tickets.md`, sección 5).
- **La pantalla del catálogo tiene dos mitades con el mismo peso** (decisión del responsable,
  2026-09-27): **categorías** y **etiquetas**, cada una con su lista, su cuenta de tickets y sus
  botones. Las etiquetas se pueden **crear**, **renombrar** —lo que vale para todos los tickets que la
  llevan— y **retirar** —lo que las quita de ellos, avisando antes—; el botón de retirar **sólo se le
  enseña al Administrador** y el de crear y renombrar, a Soporte y al Administrador.
- **La pantalla del catálogo es de `tickets` y no de Configuración** (2026-09-27): lista las
  categorías con su nombre y **cuántos tickets tiene cada una**, deja **crear y renombrar** a Soporte y
  al Administrador, y **retirar o volver a poner** sólo al Administrador —retirar una categoría en uso
  se permite y **no toca los tickets que la tienen**: deja de ofrecerse al crear—. Debajo, **la lista de
  etiquetas que se están usando**, con su cuenta, que es lo que hace que la gente vea que `red-wifi` ya
  existe y no escriba `wifi-red`.
- **El Administrador se queda con la bandeja de siempre**, con su chip de tipo y de sólo lectura
  (decisión del responsable, 2026-09-26): **no tiene tickets propios**, así que «Mis tickets» no
  significa nada para él, y las dos bandejas del «todo» le darían lo mismo que ya tiene, repartido en
  dos.
- **El alta es una pantalla y no un diálogo**, al revés que el alta de una cuenta: aquí se escribe un
  texto largo y se adjuntan archivos, y eso necesita sitio. En móvil es una pantalla completa.
- **Las direcciones las fijó el backend antes que la pantalla**: los correos que ya salen llevan
  `/tickets/ACME-2026-0042`, así que esa ruta no se elige aquí, se respeta.
- **El detalle de un ticket interno es el mismo detalle**, con su número (`INT-…`) y su conversación:
  no hay una pantalla aparte para los internos.
- **El prefijo de la numeración y el reparto se configuran en Configuración**, no en una pantalla de
  `tickets`: viven en la instalación y sus endpoints son de `settings`. Una pantalla del módulo
  `tickets` que llamara a `/api/settings` rompería la regla dura de modularidad. `tickets.md`, sección
  8, los listaba como pantalla suya y queda corregido (2026-09-24).

### 3.8 El editor de los correos

Vive en el módulo `mail`, en `/mail`, y **Configuración lleva hasta él**: es una pantalla de otro
módulo a propósito, porque llamar a `/api/mail/**` desde `settings` rompería la regla dura de
modularidad (sección 4.4).

| Pantalla | Ruta | Quién la ve |
| --- | --- | --- |
| El editor de los correos | `/mail` | Administrador |

- **A la izquierda, los trece correos**; a la derecha, el que se está editando: su **asunto**, su
  **cuerpo** y sus **marcadores**, con los imprescindibles marcados.
- **Los dos idiomas van al lado** —el mismo correo en español y en inglés, para ver que dicen lo
  mismo— **y en móvil se apilan**: en un teléfono no caben dos columnas.
- **El cuerpo se edita con formato, no escribiendo etiquetas**: negrita, cursiva, lista y enlace
  **sobre el texto seleccionado**, y el HTML lo escribe la aplicación. También **acepta HTML escrito a
  mano**, porque quien administra la instalación sabe lo que hace.
- **Los marcadores se insertan al pulsarlos**, en el punto donde esté el cursor.
- **La vista previa va al lado, y la renderiza el backend** con datos de ejemplo: no es una
  aproximación, es el correo que va a salir. Debajo, el botón de **«enviarme una prueba»**, que manda
  el correo de verdad a la dirección de quien lo pide.
- **Restaurar el texto de fábrica** está en cada plantilla, y el aviso de que falta un marcador
  imprescindible **no impide guardar**: hay avisos legítimos sin enlace.

## 4. Los cuatro enfoques

### 4.1 Usuario: pedir y seguir

Lo que hace un usuario es **crear un ticket y responder**. Todo lo demás sobra.

- **Al entrar**: sus tickets, los abiertos primero. Un botón grande y visible de **Nuevo ticket**.
- **Crear**: asunto, descripción y adjuntos. Tres campos, ninguno con jerga. Si se puede arrastrar
  una captura sobre la ventana, mejor que buscarla.
- **Seguir**: la conversación, con sus comentarios y los de Soporte. Cuando Soporte le pide algo, el
  estado lo dice —«Esperándote»— y el correo se lo ha avisado.
- **Cerrar y reabrir**: las dos, en el **desplegable de acciones** del ticket, y la de cerrar pregunta
  antes (`docs/usuarios-y-permisos.md`: cierra también el usuario).
- **Lo que no ve**: comentarios privados del principal, el ticket interno, la lista de usuarios, nada de
  administración. Su mundo son sus tickets.

### 4.2 Soporte: atender y escalar

Su problema no es entender la aplicación, es **no perder de vista qué le toca**.

- **Al entrar**: la **bandeja**, y dentro de ella lo que reclama atención primero — `nuevo`,
  `en espera` y `escalado`. Un vistazo y se sabe por dónde empezar.
- **Asignarse**: un clic desde la lista, sin abrir el ticket.
- **Atender**: comentar, mover el estado, pedir información al usuario (un botón que dice
  **«Preguntar al usuario»**, no «pasar a `en espera`»), resolver o **escalar**.
- **Escalar**: un formulario con **el motivo como campo principal** y obligatorio. Es lo que
  convierte una escalación en trabajo y no en un «pásalo tú»; la interfaz tiene que dejar claro que
  sin ese texto no se escala.
- **La vista doble** (sección 5) es suya: es quien traduce entre el usuario y Desarrollo.
- **Usuarios**: alta, desactivación y reseteo de contraseña, sin salir de la aplicación.

### 4.3 Desarrollo: entender y resolver

Desarrollo llega a un ticket que ya viene filtrado, y **no habla con el usuario**. Su pantalla tiene
que dejar eso claro.

- **Al entrar**: su bandeja de **tickets internos**.
- **Contexto**: lee el **principal** (la conversación con el usuario) y el **interno** donde trabaja.
  La vista doble es exactamente esto.
- **Acciones**: asignarse, **pedir algo a Soporte** («Necesito algo de Soporte»), y **resolver** con
  la explicación de qué se hizo.
- **Devolver**: un botón que dice **«No es un cambio de código»** y devuelve el caso a Soporte con un
  comentario. La interfaz no lo llama «cerrar sin resolver»: dice lo que significa.
- **Lo que no hace**: escribir en el ticket principal, comentar como si fuera Soporte, ni escribirle
  al usuario. La interfaz **no le ofrece** esas acciones, en vez de ofrecérselas y negárselas
  después.

### 4.4 Administrador: dejar el sistema en marcha

Su trabajo es de una vez al principio, no de todos los días.

- **La primera vez, aterriza en la puesta en marcha**: una pantalla que agrupa el **idioma de la
  instalación**, la **apariencia** (tema y color primario) y el **prefijo de numeración**, para
  dejarlo listo de una vez en lugar de ir a buscarlo por tres sitios. Después entra donde quiere, y
  esa misma pantalla sigue en el menú como **Configuración**.
- **Al entrar**: la lista de usuarios, que es lo que va a tocar.
- **Usuarios**: alta (con el rol y el origen), desactivar, reactivar, cambiar el origen y **lanzar un
  reseteo de contraseña**. Toda acción sensible dice lo que va a pasar antes de hacerlo.
- **Numeración**: el prefijo, con un aviso de que **no cambia los números ya emitidos** — es la
  duda que cualquiera tiene antes de tocarlo.
- **Y ahora también cómo se entra** (2026-09-25): la tarjeta del **método de entrada** —los tres en un
  desplegable, con la explicación del que está elegido— y **la configuración que toca según el método
  elegido**, con su botón de **«Probar la conexión»**. **Cada panel sale con su método y ninguno más**
  (decisión del responsable, 2026-09-27): con **«Cuentas de la aplicación»** no se enseña ninguna de
  las dos —no hay nada que configurar—, con la **organización** sale el del directorio, y con
  **Keycloak** el del reino. **Lo que no se está usando no es trabajo que toque hacer**, y las dos
  configuraciones a la vez invitaban a tocar la que no era. **Las opciones de un
  método sin configurar se enseñan pero no se pueden elegir**, y eso vale también para lo que se está
  escribiendo: en cuanto se rellenan el servidor y la base de búsqueda, la opción se enciende sin
  guardar antes. Los campos de los dos secretos **llegan siempre vacíos** y su ayuda dice si hay uno
  puesto: la pantalla nunca enseña un secreto, porque no lo recibe (`docs/modules/settings.md`,
  sección 5.8).
- **Configuración tiene ya todo lo suyo** (2026-09-24 y 2026-09-25): **el nombre de la instalación**
  —el campo, con su ayuda y su botón de guardar; se ve en toda la aplicación en cuanto se guarda—,
  **la marca** —los dos huecos del logo, con su vista previa, subir y volver al de fábrica—, **el color institucional** con **vista previa de cómo
  queda** antes de guardar, el **idioma global de la instalación** —primer control de la pantalla— y el
  **prefijo de numeración** con el aviso de que **no cambia los números ya emitidos**, más el
  **reparto y los avisos** de cada tipo de ticket. Todo lo de esta pantalla viene de un solo endpoint,
  el del módulo `settings`, más los del logo.
- **Los correos se editan en una pantalla del módulo `mail`**, a la que Configuración enlaza: cada
  plantilla con su asunto y su cuerpo, el idioma al lado —se editan los dos—, **los marcadores
  disponibles a la vista** (`{{nombre}}`, `{{enlace}}`, `{{numero}}`…), una **vista previa** y un
  botón de **«enviarme una prueba»**. Un marcador que no exista se avisa **al guardar**, no cuando el
  correo ha salido. `docs/modules/mail.md`, decisión 15.
- **Es una pantalla de otro módulo a propósito**: llamar a `/api/mail/**` desde `settings` rompería
  la regla dura de modularidad (`docs/arquitectura.md`, sección 4), así que la pantalla vive donde
  viven sus datos y Configuración sólo lleva hasta ella.
- **El cuerpo se edita con formato**, no escribiendo etiquetas: negrita, cursiva, lista y enlace
  sobre el texto seleccionado, y el HTML lo genera la aplicación. Ver la decisión 1 de `mail.md`.
- **Ve todo, sin tocar nada**: puede mirar cualquier ticket —también los internos y la vista
  doble— pero **no comenta, no mueve estados, no asigna y no cierra**. Es el precio de ser el papel
  que lo puede todo en configuración: donde no trabaja, mira.

## 5. La vista doble: principal e interno

Pedida a propósito, y es la pieza que hace usable el modelo de dos tickets. **Sólo existe para
Soporte y Desarrollo**, porque el usuario no ve el interno.

- **Cuándo aparece**: cuando el ticket tiene interno y el papel puede verlo. Si no hay interno, no
  hay nada que comparar y no aparece nada.
- **Cómo se controla**: un conmutador con tres posiciones —**Principal**, **Los dos**, **Interno**—
  en la cabecera del detalle. «Los dos» es el valor por defecto cuando existe el interno, porque
  quien está ahí quiere las dos cosas.
- **En PC**: dos columnas, la conversación de cada ticket y su ficha. Es comparar, y comparar es leer
  en paralelo.
- **En móvil y tablet vertical**: no caben dos columnas, así que el mismo conmutador cambia entre
  uno y otro. **Nunca dos columnas estrechas**: es peor que verlas por separado.
- **Cómo no confundirlos**: el interno lleva fondo distinto y su número `INT-…` bien visible. Si los
  dos se ven iguales, alguien va a escribir en el ticket equivocado.
- **Las acciones son del ticket que se está tocando**: en «Los dos» cada columna tiene su ficha y sus
  botones, y la caja de comentario es la de la columna donde estás. Escribir en el principal desde
  Desarrollo no es una opción que se ofrezca (sólo lectura).

## 6. El lenguaje visual

### 6.1 La base

**Tailwind CSS v4 con componentes propios.** Se escribe el aspecto con utilidades y **los componentes
se construyen aquí**, no se traen de una biblioteca.

Lo que eso significa, dicho claro:

- **El aspecto es exactamente el que se decida**, sin pelearse con los estilos de una biblioteca
  ajena ni con su sistema de temas.
- **Los componentes complejos son nuestros**: tabla, diálogo, menú, pestañas, subida de archivos,
  avisos. Cada uno hay que construirlo **y hacerlo accesible** (sección 8). Es el precio, y es un
  precio real: no es lo mismo escribir un botón que un diálogo que atrapa el foco y lo devuelve al
  cerrarse.
- **Inventario cerrado de componentes** (sección 6.3): lo que no esté en esa lista, se discute antes
  de construirlo, en lugar de aparecer por sorpresa.

Calibyou usa Tailwind 3; aquí se usa la **4** (la actual), que define los colores como variables CSS
y encaja mejor con el sistema de temas de la sección 6.2.

**Moderno, sobrio y sin adornos.** La tendencia de estos años —interfaces planas, tipografía grande,
mucho espacio en blanco, bordes suaves, sombras casi inexistentes— coincide con lo que hace falta
aquí: nada que compita con los datos.

| | |
| --- | --- |
| **Tipografía** | Una sola familia de sistema (la que ya usa la máquina), con escala de tamaños fija. Nada de fuentes descargadas: cargan tarde y en una mesa de ayuda no aportan |
| **Espaciado** | Escala de 4 y 8 píxeles, sin valores sueltos |
| **Bordes** | Radio pequeño y bordes finos; el borde separa mejor que la sombra |
| **Sombras** | Sólo en lo que flota de verdad: diálogos y menús |
| **Color** | Neutro para casi todo y **un color de acento** para lo accionable. El color se usa para señalar, no para decorar. Los colores no se escriben a mano en los componentes: salen de los temas (sección 6.2) |
| **Iconos** | Una sola librería, siempre con texto al lado en las acciones importantes |
| **Movimiento** | Casi ninguno: transiciones cortas al abrir un diálogo o plegar el menú, y nada más |

**El color de los estados** (siempre acompañado de su texto, nunca solo):

| Estado | Intención |
| --- | --- |
| `nuevo` | Llama la atención: acaba de llegar |
| `en progreso` | Neutro: alguien lo tiene |
| `en espera` | Ámbar: bloqueado, esperando a alguien |
| `escalado` | Distinto de todo lo demás: está en otro equipo |
| `resuelto` | Verde suave: hecho, pendiente de cerrar |
| `cerrado` | Gris apagado: fuera de la vista |

### 6.2 Los temas: ocho, y sólo cambian los colores

La forma de los controles **no cambia nunca**: un botón es el mismo botón en todos los temas. Lo que
cambia es la paleta. Eso es lo que hace que ocho temas no sean ocho interfaces distintas.

| | Cuántos | Qué son |
| --- | --- | --- |
| **Claros** | 3 | Tres paletas claras, fijas |
| **Oscuros** | 3 | Tres paletas oscuras, fijas |
| **Personalizables** | 1 claro y 1 oscuro | Se les cambia **el color primario**: el institucional, al gusto de quien administra |

- **Cómo se implementa**: los colores son **variables CSS** (`--color-primario`, `--color-fondo`,
  `--color-texto`, los colores de estado…). Un tema es un conjunto de valores de esas variables, y
  ningún componente sabe qué tema está activo. Añadir un tema noveno es escribir nueve líneas.
- **Los dos personalizables son los de por defecto.** Una instalación nueva entra con el tema
  personalizable **claro u oscuro según lo que pida el sistema operativo**, siempre con el color
  institucional: así lo primero que ve todo el mundo es la cara de la casa, sin que nadie tenga que
  configurar nada.
- **El color primario lo fija el administrador**, y sólo él: es el color de la institución, y si cada
  uno lo cambiara dejaría de serlo. Lo que cada persona puede cambiar es **el tema**, eligiendo
  entre los ocho desde la cabecera.
- **Elegir es opcional**: quien no toque nada se queda con el de la instalación. La elección de cada
  persona se recuerda **en su navegador**, que es lo simple: no hace falta tocar la base de datos ni
  añadir una pantalla de perfil.
- La paleta de cada tema tiene que **pasar el contraste** de la sección 8, incluidos los colores de
  estado. Un tema bonito que no se lee no entra.

**Cómo se implementa hoy** (2026-09-23, con lo que hay construido):

- **Un atributo en el elemento `html`**: `data-theme="claro"`, `data-theme="oscuro"` o **ninguno**,
  que es «seguir al sistema». Las paletas se escriben **una sola vez** con la función `light-dark()`
  de CSS y el atributo sólo cambia `color-scheme`: así no hay dos listas de colores que se puedan
  desincronizar, y añadir una paleta nueva son unas líneas, como dice la tabla de arriba.
- **La elección se recuerda en el navegador** (`localStorage`), que es lo simple y lo que ya está
  decidido: no toca la base de datos ni hace falta una pantalla de perfil.
- **El tema lo aplica código nuestro, nunca un `script` escrito en el `index.html`**: la CSP de
  producción es `script-src 'self'`, sin `'unsafe-inline'`, así que un script incrustado **no
  funcionaría en producción** y ensanchar la CSP por esto sería cambiar seguridad por comodidad. Se
  aplica antes de arrancar la aplicación, desde `main.ts`.
- **El conmutador enseña dos opciones: claro y oscuro. Nada más.** **«Automático» no se muestra ni se
  elige** (decisión del responsable, 2026-09-23): seguir al sistema es **no haber elegido**, que es el
  estado de fábrica, y por eso no hace falta una tercera opción que además sería la única forma de
  volver atrás. Mientras nadie toque nada, el conmutador **enseña marcado lo que se está viendo**
  —claro u oscuro según el sistema—, no una opción que no existe; en cuanto alguien elige, se queda
  con lo que eligió y el sistema deja de mandar. Si el sistema cambia con la página abierta y nadie
  ha elegido, la aplicación cambia con él, sin recargar.
- **Las seis paletas fijas** llegan con el armazón completo, cuando se aprueben sus nombres y sus
  colores. Mientras no exista la cabecera del armazón, el conmutador vive **al lado del de idioma**.
- **La paleta de fábrica sigue al sistema operativo** hasta que alguien elija otra cosa, que es la
  decisión 5 del registro.

### 6.4 La marca: el logo y el nombre de la instalación

La marca **es de la instalación**, no de cada persona, y la cambia un Administrador desde
Configuración (`docs/modules/settings.md`, secciones 5 y 5.7). Son **el logo y el nombre**, y se
enseñan juntos y siempre igual:

| Dónde | El logo | El nombre |
| --- | --- | --- |
| **Pantalla de entrada** | Arriba del formulario, centrado, **hasta 128 px en móvil y 160 px de tablet para arriba**: es lo primero que se ve | Debajo del título de la tarjeta, donde está el formulario: es lo primero que se lee |
| **Primer arranque** | Junto a la cabecera, **80 px de alto**: ocupa el mismo bloque visual que el título y su explicación sin competir con el contenido del paso | El nombre todavía no se enseña aquí: se está configurando |
| **Menú lateral** | Arriba, junto al nombre, **hasta 32 px de alto** | Arriba, junto al logo |
| **Pestaña del navegador** | — | Es el título de la pestaña, en todas las pantallas |
| **Configuración** | Los dos huecos con su vista previa y los botones de subir y volver al de fábrica | El campo, con su ayuda: si se deja vacío vuelve el nombre de fábrica |

**Y con la marca va la versión del sistema**, que no se configura: la escribe el código
(`docs/modules/settings.md`, sección 5.9).

| Dónde | La versión |
| --- | --- |
| **Menú lateral** | En la fila de salir, **alineada a la derecha**, en texto pequeño y apagado |
| **Pantalla de entrada** | **Debajo de los controles**, centrada y en texto pequeño y apagado: es el pie de la pantalla |
| **En ningún otro sitio** | No va en los correos, ni en el título de la pestaña, ni en las pantallas de producto |

- **Se escribe `v1.0.0`, con la `v`**: es como se lee una versión, y sin ella un `1.0.0` suelto
  parece cualquier otra cosa.
- **No es un enlace ni un botón**: no lleva a ningún sitio. Se lee y ya.
- **En la pantalla de entrada sí tiene sentido** (decisión del responsable, 2026-09-25): es donde
  alguien que no puede entrar mira para decir qué versión está usando, y la pantalla de entrada es lo
  único que ve.
- **No se enseña la fecha de compilación**: cambiaría en cada compilación sin decir nada nuevo, y dos
  instalaciones con el mismo código dirían cosas distintas.

- **El nombre no lleva el de la pantalla**: la pestaña se llama como la instalación y nada más (todas
  las rutas traían el mismo texto, así que no decían dónde estabas).
- **Se cuenta en caracteres y no pasa de 60**, que es lo que cabe en el menú y en la tarjeta de la
  entrada; el campo lo corta al escribir y el backend lo comprueba igual.

**El sistema tiene que servir para cualquier logo**, y eso se resuelve en cómo se enseña, no pidiendo
un archivo concreto:

- **No se recorta y no se estira**: se enseña entero, con su proporción, dentro del alto máximo.
- **Va en un recuadro de esquinas redondeadas y con un borde del tema.** Así un logo con **fondo
  opaco** —el de fábrica, por ejemplo, es un cuadrado negro— se ve como una pieza deliberada, y uno
  **con transparencia** se integra en el fondo. Ninguno de los dos se da por supuesto.
- **En la barra de controles, el selector de tema va sin etiqueta a la vista** (decisión del
  responsable, 2026-09-24): el desplegable ya enseña el tema que está puesto, así que la palabra
  «Tema» sólo ocupaba sitio. **El nombre accesible sigue ahí**, oculto a la vista: un lector de
  pantalla tiene que poder decir de qué es el control (sección 8).
- **Hay dos huecos, claro y oscuro, y los dos son opcionales** (sección 6.2 y
  `docs/modules/settings.md`, sección 5.1): si sólo hay uno, se usa en los ocho temas. **Y los dos
  logos de fábrica también son dos, uno por tema** (2026-09-26), así que una instalación sin logo
  propio se lee igual de bien en claro que en oscuro.
- **En los correos no va**: los correos son HTML sin imágenes (`docs/modules/mail.md`).

### 6.3 Inventario de componentes propios

Lo que hay que construir, y nada más:

`botón` (**cuatro formas** —principal, secundaria, peligrosa y la verde de la confirmación— **y puede
ser sólo un icono**, que entonces necesita su nombre accesible; **el icono es uno de los del proyecto**
—`app-icono`, `frontend/src/app/shared/components/icono.ts`, svg propios, nunca un carácter ni una
biblioteca—) · `campo de texto` · `área de texto` · `selector` · `casilla` · `etiqueta de estado` ·
`tabla` (con orden y paginación) · `tarjeta` (la lista en móvil) · `diálogo` · `menú` ·
`conmutador segmentado` (la vista doble y los filtros) · **`editor de texto con formato`, con los
adjuntos dentro** · `adjunto` (con vista previa) · `comentario` (con marcas de editado y eliminado) ·
`aviso` (éxito, error, información) · **`toast`** (resultado fijo en la ventana) · `estado vacío` ·
`indicador de carga` · `menú lateral` (con sus tres zonas) · `barra de móvil`.

Veinte piezas. Eran veinte antes de que **`subida de archivos con arrastrar y soltar` se retirara**
(2026-09-26); al integrarla en el editor quedaron diecinueve, y el toast aprobado el 2026-10-06 lleva
el inventario de nuevo a veinte. Adjuntar dejó de ser un recuadro aparte —ahora **se adjunta dentro
del texto**, con el editor— y tener dos sitios donde se hace lo mismo era el ruido que este producto
evita. Lo que hacía —arrastrar, elegir y filtrar por extensión con su aviso— **lo hace el editor**,
con las tres puertas de la sección 6.5.

Cada una se construye **una vez**, con sus estados (normal, hover, foco, deshabilitado, cargando) y su
comportamiento de teclado.

**Hechas hoy**: el botón, el campo de texto, el **área de texto**, el selector, el conmutador
segmentado, la tarjeta, el **aviso**, el **toast**, el **diálogo**, la **etiqueta de estado**, el **adjunto con su
vista previa**, el **editor con formato —con los adjuntos dentro, sección 6.5—**, el menú lateral con
sus tres zonas y la barra de móvil. Lo que falta —el comentario como pieza suelta y el indicador de
carga— llega con las pantallas que lo usan.

- **El diálogo es un `<dialog>` nativo** (2026-09-24): se cierra con **Escape**, atrapa el foco y lo
  **devuelve a donde estaba** sin que haya que programarlo. Es lo que hace el alta de una cuenta.
- **La tarjeta elige el nivel de su título** (2026-09-24): `h1` cuando la tarjeta *es* la pantalla —las
  de la sesión— y `h2` cuando va debajo del título de una pantalla de producto. Sin eso, la pantalla
  de usuarios tenía tres `h1`, y una página así no se navega con un lector de pantalla (sección 8).
- **Adjuntar tiene tres puertas y una sola regla** (2026-09-25, y desde el 2026-09-26 con el editor):
  **pegar**, **arrastrar** y **el botón**, y las tres **meten el archivo donde está el cursor**. Se
  comprueba lo mismo —la lista cerrada de extensiones— y se avisa igual de lo que no vale
  (`docs/modules/tickets.md`, secciones 2.3 y 4).
- **El selector puede enseñar una opción que no se puede elegir** (2026-09-24), con su texto de ayuda
  debajo: es como se enseñan hoy los orígenes de directorio, que están previstos y todavía no
  funcionan (sección 3.6).
- **Todos los desplegables conservan un `<select>` nativo**, pero dibujan su indicador con el color
  del texto y lo separan del borde derecho. El texto reserva ese espacio, por lo que nunca queda
  debajo de la flecha.

### 6.5 El cuadro de escribir: un editor con formato y los adjuntos dentro

**Escribir un ticket o un comentario no es rellenar un campo de texto**: lo que se escribe lleva
formato y **los archivos que se adjuntan se ven dentro del texto**, en el sitio donde se ponen
(`docs/modules/tickets.md`, sección 2.3). Un cuadro de texto normal no puede llevar una imagen dentro,
así que **el cuadro de escribir es un editor con formato**.

- **La barra del editor**: negrita, cursiva, subrayado, tachado, lista con viñetas, lista numerada,
  enlace y **Adjuntar**. Los botones actúan **sobre lo que está seleccionado** —es la misma regla del
  editor de los correos, sección 3.8— y **Adjuntar mete el archivo donde está el cursor**.
- **Tres puertas para adjuntar, y las tres meten el archivo en su sitio**: **pegar** (`Ctrl+V`),
  **arrastrarlo encima del editor** y **el botón**. Se acepta lo que esté en la lista cerrada de
  extensiones, y lo que no está se rechaza con el mismo aviso de siempre.
- **Al pegar texto, el texto entra sin formato** (2026-09-26): lo que se pega se queda como texto plano
  y **se respetan los saltos de línea**. Pegar el HTML de un correo o de Word traería etiquetas que no
  están en la lista blanca, y el comentario se rechazaría entero por algo que la persona no ha escrito
  —un `style`, un `<font>`—, que es la peor forma de perder un texto. Un **archivo** pegado (una
  captura, lo que se copia del explorador) sí entra, y entra como adjunto. Es la misma regla de
  `docs/modules/tickets.md`, sección 2.3.
- **Un enlace se pide con su dirección y se comprueba antes de ponerlo**: tiene que empezar por `http`,
  `https` o `mailto`. Lo que no, no se pone y se dice por qué, en vez de rechazar el comentario al
  guardarlo (es la misma regla que aplica el backend).
- **Se ve lo que va a salir**: la imagen y el vídeo **se ven de verdad** mientras se escribe, no como
  un texto que los representa. Un PDF o un Word se ven como **su enlace**, con su nombre.
- **Llamar a alguien se escribe en el comentario**: el editor tiene un botón **«Etiquetar»** —y
  escribir `@` abre lo mismo—, que enseña **la lista de técnicos y desarrolladores activos** y mete el
  nombre elegido **donde está el cursor**, resaltado como una **mención**. Quien lo lee ve a quién se
  está llamando, y esa persona pasa a **observar** el ticket: aparece en su «Observo» y en la lista de
  observadores de la ficha. **Sólo Soporte y Desarrollo pueden etiquetar** (el usuario no ve la lista de
  técnicos), y a quien se etiqueta **se le avisa por correo**. En la ficha, cada observador lleva su
  aspa para quitarlo, y quitarlo **queda en la línea de tiempo**.
- **Con tope, y el clic los abre grandes**: en el texto —escribiendo y leyendo— la imagen y el vídeo
  se ven **como mucho a 480 × 360** (decisión 55, del 2026-09-26), sin deformarse: lo que no puede es
  que una captura a pantalla completa se coma la conversación. **El tope es `min(100%, 480px)` y no
  `480px`**: un `max-width` sólo limita, así que con el valor fijo una imagen de 480 **se salía de la
  pantalla de un móvil de 412 px** y ensanchaba el documento entero —lo encontró la capa de
  Playwright, porque el clic en un botón de la misma tarjeta empezó a fallar con el scroll
  horizontal—. Ahora se encoge con la pantalla en vez de desbordarla. **Al pulsarlos se abren en el visor**,
  más grande, con **Descargar** y **Abrir en una pestaña**. Para eso el marco de lo que se ve es **un
  `button` de verdad** —se llega con el tabulador, se pulsa con Enter y dice qué es—, y el vídeo lleva
  encima **su botón de reproducir**: una imagen que se pulsa y no dice que se pulsa es una imagen que
  nadie descubre con un lector de pantalla (sección 8).
- **El vídeo se reproduce en el visor**, no en el texto: en la conversación se ve **una miniatura**
  —sin controles— y al pulsarla se abre el visor, que es donde se ve grande, se reproduce y se
  descarga. **Mientras se escribe sí lleva controles**, porque ahí no hay visor y lo que se está
  haciendo es comprobar lo que se acaba de adjuntar.
- **Los controles del reproductor los pone la pantalla**: lo que se guarda es
  `<video data-adjunto="grabacion.mp4">`, sin `controls`, porque **ese atributo no está en la lista
  blanca del backend** (`backend/modules/tickets/services/body.go`): con él, el texto entero se
  rechazaría. **El tamaño tampoco se guarda**: es cómo se pinta. Queda dicho aquí porque **el hallazgo
  fue justo el contrario**: el ejemplo de `docs/modules/tickets.md`, sección 2.3, escribía
  `<video data-adjunto="x.mp4" controls>`, que es un texto que el backend habría rechazado; el
  documento se corrigió en el mismo trabajo.
- **Un archivo que se nombró y no subió deja su hueco marcado**: si una subida falla después de
  publicar el comentario (decisión 44 de `tickets.md`), el sitio donde iba el archivo **se ve**, con su
  nombre y un aviso, en vez de quedar un hueco mudo que parezca un fallo de la pantalla. Cuando el
  archivo sube, el hueco se convierte en la imagen.
- **Quitar un adjunto es quitarlo del texto**: se borra como se borra cualquier cosa que se ha
  pegado. Y **un archivo que no está en el texto no se sube**, así que no hay que acordarse de
  quitarlo en dos sitios: hay uno.
- **El recuadro de archivos aparte desaparece** de estos formularios (`app-archivos` se retira, y el
  inventario de la sección 6.3 lo refleja): tener dos sitios donde se adjunta lo mismo era justo el
  ruido que este producto evita.
- **Lo que se escribe se cuenta en texto, no en HTML**: las etiquetas y las direcciones de los
  adjuntos son cosa del editor y de la pantalla, y **nadie las ve**. Eso incluye la búsqueda de la
  bandeja, que **quita las etiquetas antes de buscar** para que buscar «p» no encuentre todos los
  tickets con una imagen.
- **Y lo que se manda lleva sólo lo que el backend admite**: un adjunto sale como su `data-adjunto` y
  **nada más** —ni la dirección `blob:` de la vista previa, ni su `class`, ni un `style`—, porque el
  saneador del backend **rechaza el texto entero** con `422 tickets.body.notAllowed` en cuanto aparece
  algo que no está en su lista. Cuando eso pasa, la pantalla lo dice en castellano («El texto lleva algo
  que no se admite…»), no enseña la clave. **Los controles del vídeo tampoco se guardan**: los pone la
  pantalla al pintarlo.
- **Al leer, lo mismo**: la conversación pinta el texto con formato y, en su sitio, la imagen o el
  reproductor del vídeo; el PDF y los demás archivos, como enlaces. **Pulsar una imagen o un PDF abre
  el visor** —el modal, más grande, con **Descargar** y **Abrir en una pestaña**, decisión 48 de
  `tickets.md`—, y pulsar un Word o un Excel **lo descarga**.

## 7. Multi-dispositivo

**La prioridad es el PC**: es donde se trabaja, donde hay dos columnas y donde están las bandejas
completas. Tablet y móvil tienen que **funcionar de verdad**, no ser una promesa.

| | Móvil (< 768) | Tablet (768–1200) | PC (> 1200) |
| --- | --- | --- | --- |
| **Menú** | Botón en la cabecera | Plegado, se abre si se quiere | Fijo, plegable |
| **Listas** | Tarjetas | Tabla reducida | Tabla completa |
| **Detalle** | Una columna; la ficha baja debajo | Una columna más ancha | Conversación + ficha a la derecha |
| **Vista doble** | Conmutador, de uno en uno | Conmutador, de uno en uno | Las dos columnas |
| **Administración** | Consultar y acciones simples | Todo | Todo |

- **Objetivos táctiles de 44 px** como mínimo en móvil: un botón que no se acierta con el dedo es un
  botón que no existe.
- **Nada de desplazamiento horizontal**: si una tabla no cabe, se convierte en tarjetas, no se
  arrastra de lado.
- **Lo que se hace en móvil** (decidido, sección 11): **todo menos la administración y la vista
  doble**, que son lo que se deja para el PC.

## 8. Accesibilidad y confianza

- **Contraste suficiente** en texto y en los colores de estado, en claro y en oscuro.
- **Todo se puede usar con teclado**: foco visible siempre, orden de tabulación coherente y nada que
  sólo funcione con el ratón.
- **Nunca sólo color** para decir algo: el estado lleva su texto y los avisos llevan su icono.
- **Etiquetas de verdad** en los campos, no textos de ejemplo que desaparecen al escribir.
- **Los errores dicen qué pasó y qué hacer**, en el idioma de quien los lee. Un «Error 422» no es un
  mensaje. El backend manda una **clave** (`auth.invalidCredentials`) y la interfaz la traduce: por eso
  no basta con tener un diccionario de pantallas —**cada clave de error necesita su texto en los dos
  idiomas**, y una clave sin texto es un error sin mensaje.
- **Los diálogos cierran con Escape** y no roban el foco sin devolverlo.
- **Los adjuntos se anuncian** con su nombre y su tamaño antes de descargarlos.

**Como los componentes son nuestros, esto no viene de serie.** Lo que hay que hacer bien una vez y
reutilizar:

| Componente | Lo que no se puede olvidar |
| --- | --- |
| **Diálogo** | Foco dentro al abrirse, devuelto a donde estaba al cerrarse, Escape para cerrar y el fondo sin poder tocarse |
| **Menú** | Flechas para moverse, Enter para elegir, Escape para cerrar |
| **Conmutador segmentado** | Se anuncia como grupo de opciones, no como botones sueltos |
| **Tabla** | Encabezados de verdad, orden anunciado y navegación con teclado por las filas |
| **Subida de archivos** | Se puede hacer sin ratón, y cada archivo dice qué pasó con él |
| **Avisos** | Se anuncian solos a quien usa lector de pantalla, y no roban el foco |

## 9. Los estados que no son pantallas

Casi nadie diseña esto y es lo que más se sufre:

- **Lista vacía**: dice por qué está vacía y qué hacer. «No tienes tickets todavía» con el botón de
  crear al lado, no una tabla en blanco.
- **Bandeja vacía**: «No hay nada pendiente» — y que sea verdad, sin inventarse un mensaje motivador.
- **Cargando**: un indicador discreto y sólo si tarda; nada de pantallas de carga para medio segundo.
- **Sin permiso**: si alguien llega a una pantalla que no le toca, se le dice, no se le enseña en
  blanco.
- **Algo falló**: qué se intentaba, qué pasó y qué puede hacer. Y el camino de vuelta siempre visible.
- **Sin conexión o servidor caído**: un aviso claro, sin tecnicismos.

## 10. Detalles pequeños que se notan mucho

- **El número del ticket, siempre visible y copiable.**
- **La fecha de última actualización** en la lista y en el detalle: lo que lleva días parado se ve.
- **Sin refrescos automáticos ni contadores en vivo** (no hay tiempo real, `docs/modules/tickets.md`): la
  lista se actualiza al entrar y con un botón de refrescar. Un contador que no se actualiza solo es
  peor que no tenerlo.
- **Los adjuntos de imagen y PDF se previsualizan** dentro del ticket; el resto se descarga
  (`docs/modules/tickets.md`, sección 4).
- **Los correos son HTML, con una versión de texto automática** (`docs/modules/mail.md`): mismo tono
  que la interfaz, sin jerga y con estilos en línea, sin imágenes externas ni fuentes de terceros.

## 11. Lo que se decidió al repasar este documento

| # | Decisión | Quedó así |
| --- | --- | --- |
| 1 | **La base visual** | **Tailwind CSS v4 con componentes propios** (20 piezas, sección 6.3), no una biblioteca |
| 2 | **La vista doble** | Dos columnas en PC con **«Los dos» por defecto** cuando existe el interno; conmutador en móvil |
| 3 | **Qué se hace desde el móvil** | **Todo menos administración y vista doble** |
| 4 | **Los temas** | **Ocho: 3 claros + 3 oscuros fijos, más 1 claro y 1 oscuro personalizables**, que son los de por defecto. El **color primario lo fija el administrador**; cada persona elige tema. Sólo cambian colores, nunca la forma |
| 5 | **Modo claro y oscuro** | Sí desde la 1.0.0, siguiendo al sistema hasta que alguien elija |
| 6 | **Accesibilidad** | **Requisito**: contraste, teclado y foco en todas las pantallas, con AA como referencia, y comprobado al construir cada componente |
| 7 | **El tema, hoy** | Claro y oscuro (los dos personalizables de fábrica), con el mecanismo listo para las seis paletas fijas. De fábrica sigue al sistema, que es lo mismo que **no haber elegido** |
| 10 | **«Automático» no se elige ni se muestra** | Decisión del responsable (2026-09-23): el conmutador tiene dos opciones y mientras nadie elija **enseña marcado lo que se está viendo**. Seguir al sistema es el estado de fábrica, no una tercera opción |
| 11 | **El sistema manda en vivo** | Si cambia de claro a oscuro con la página abierta y nadie ha elegido, la aplicación cambia con él sin recargar |
| 8 | **Dónde vive el conmutador mientras no hay cabecera** | Al lado del de idioma, en la entrada y en el inicio: se muda a los controles de la cabecera cuando exista el armazón |
| 9 | **El tema no se aplica con un script incrustado** | La CSP de producción es `script-src 'self'`: se aplica desde `main.ts`, antes de arrancar |
| 13 | **Un comentario sin texto se lee por sus adjuntos** | **Elección del responsable, 2026-09-25**, al decidir que un comentario pueda ser sólo un archivo: en la conversación aparece con su autor, su fecha y sus archivos, **sin un hueco en blanco** que parezca un fallo de la pantalla. Y **si un archivo no sube después de publicar el comentario**, se avisa de cuál falta y se puede reintentar sin volver a escribirlo (`docs/modules/tickets.md`, decisiones 41 y 44) |
| 12 | **La versión del sistema** | **Elección del responsable, 2026-09-25**, y **no la que yo recomendaba**: se lee en la fila de salir del menú lateral, **alineada a la derecha**, y **también en el pie de la pantalla de entrada**. Es un dato, no un botón ni una fila propia, y **plegado no se enseña**. Yo proponía sólo el menú; él decidió que también se vea antes de entrar, que es donde sirve para decir qué versión se está mirando |

## 12. Qué habilita este documento

Con `docs/interfaz-y-experiencia.md` aprobado, la documentación cubre **qué se construye, cómo se
comporta y cómo se ve**. El código puede empezar por `auth` (`docs/modules/tickets.md`), y sus pantallas
—entrar, recuperar contraseña— se diseñan ya con estas reglas, en lugar de retocarlas después.

## 13. Navegación de regreso a la izquierda

### 13.1 Hallazgo e inventario

La regla aprobada el 2026-09-29 ya dice que los botones de volver van a la izquierda y son una
flecha de sólo icono con nombre accesible. El inventario de las vistas con ese control de cabecera
encontró:

| Vista | Destino | Posición encontrada | Cumplía la regla |
| --- | --- | --- | --- |
| `/mail` | Configuración | A la derecha del título | No |
| `/tickets/new` | Lista de tickets | A la izquierda del título | Sí |
| `/tickets/:número` | Lista de tickets | A la izquierda del número | Sí |
| `/users/:id` | Lista de usuarios | A la izquierda del nombre | Sí |

El comentario del propio `correos-page.html` también afirmaba que la flecha estaba a la izquierda
del título, pero la estructura `justify-between` la colocaba después del título y al extremo
derecho. Era una incoherencia entre el documento, el comentario y lo que se pintaba.

### 13.2 Regla

Todo control que **regresa desde una vista secundaria y vive en su cabecera** aparece como primer
elemento de esa cabecera, inmediatamente antes del título o identificador principal. Es una flecha
de sólo icono, conserva un nombre accesible y mantiene su destino actual.

Esta regla se aplica a los cuatro controles inventariados. Tres ya cumplían; la implementación
reordena la cabecera de `/mail` para colocar su enlace a Configuración antes del bloque de título y
descripción. No cambia el destino, el texto accesible, el icono ni el tamaño del control.

Quedan fuera:

- los enlaces que forman parte del flujo de una pantalla, como **Volver a iniciar sesión**;
- el botón **Volver a mi trabajo** de la vista sin permiso;
- acciones con otra finalidad que usan la palabra «volver», como restaurar un logo o un correo;
- añadir botones nuevos a vistas que hoy no tienen navegación de regreso;
- el backend, las rutas y los permisos.

### 13.3 Documentos y código

- Este documento define la regla general y registra el hallazgo.
- `docs/README.md` y `AGENTS.md` reflejan el estado final de la enmienda.
- `frontend/src/app/modules/mail/correos-page.html` coloca la flecha antes del bloque de título y
  descripción.
- La prueba E2E del editor de correos comprueba que la flecha es el primer elemento de la cabecera,
  antes del título, y que sigue regresando a Configuración. La suite la ejecuta en PC y móvil.

`docs/modules/mail.md` no cambia: describe las funciones del editor y no fija la composición de su
cabecera. Además, su comportamiento funcional no cambia.

### 13.4 Criterios de aceptación

1. En `/mail`, la flecha aparece inmediatamente a la izquierda del título **Los correos**.
2. En anchos de PC y móvil, el control permanece antes del título cuando la cabecera se reparte en
   más de una línea.
3. Su nombre accesible sigue siendo **Volver / Back** y su destino sigue siendo `/settings`.
4. Los tres controles que ya cumplen conservan su posición y su destino.
5. Los enlaces y acciones excluidos no cambian.

### 13.5 Decisión registrada

El responsable confirmó el 2026-10-05 la recomendación: unificar sólo los controles de navegación
de cabecera como icono a la izquierda del título. Los enlaces de flujo dentro del contenido quedan
fuera.

### 13.6 Registro del repaso

El responsable confirmó el 2026-10-05 que:

1. la flecha quedaría como primer elemento de la cabecera;
2. el título y su descripción permanecerían juntos, en un bloque a la derecha de la flecha;
3. no se añade ninguna excepción adicional a la regla ni se amplía el alcance.

El repaso queda cerrado sin decisiones abiertas. El responsable aprobó explícitamente la propuesta
con «Sí». La implementación respeta el alcance aprobado y el documento vuelve a `as-built`.

### 13.7 Verificación

- Las 216 pruebas unitarias del frontend pasan.
- Los cuatro casos E2E del editor de correos pasan: dos en PC y dos en móvil.
- La prueba comprueba que la flecha es el primer elemento de la cabecera, que su posición horizontal
  queda antes del título y que al pulsarla se regresa a `/settings`.

## 14. Región seleccionada visible en Configuración

### 14.1 Hallazgo

La tarjeta **Región horaria y dirección pública** de `/settings` ofrece el mismo buscador y la misma
lista desplazable que `/setup`. La opción elegida sólo se reconoce por el fondo y el peso de texto
dentro de esa lista; si queda fuera de la parte visible, la pantalla no nombra de forma estable la
zona que se guardará. La hora y el desfase permiten comprobarla, pero no dicen cuál es.

### 14.2 Comportamiento

Se reutiliza exactamente el patrón aprobado para `/setup`:

- campo de sólo lectura antes del buscador;
- etiqueta **Selected time zone / Región horaria seleccionada** según el idioma de la interfaz;
- identificador IANA exacto como valor;
- zona guardada al cargar la pantalla y actualización inmediata al elegir otra;
- texto seleccionable y copiable, sin edición directa;
- hora y desfase conservados debajo de la lista.

El valor del campo es el borrador actual del formulario. Elegir una zona lo cambia y habilita el
guardado existente; guardar, recibir un error o volver a cargar siguen las reglas actuales de
Configuración. El buscador, la lista, el resaltado, la dirección pública y el botón de guardar no
cambian.

### 14.3 Alcance

- `docs/modules/settings.md` define el comportamiento dentro de la tarjeta de región.
- Este documento define su presentación coherente con `/setup`.
- `docs/README.md` y `AGENTS.md` reflejan el estado final de la enmienda.
- `frontend/src/app/core/pages/settings-page.html` muestra el campo y las pruebas existentes de
  Configuración comprueban su comportamiento.
- Se reutilizan los textos ES/EN actuales; no se añaden claves de traducción.

Quedan fuera la API, la base de datos, la validación de zonas, `/setup`, la dirección pública y las
demás tarjetas de Configuración.

### 14.4 Criterios de aceptación

1. Al cargar `/settings`, el campo muestra la zona guardada.
2. Elegir `America/Guayaquil` actualiza el campo inmediatamente antes de guardar.
3. El campo es de sólo lectura y permite seleccionar y copiar el identificador.
4. La búsqueda, la lista, el resaltado, la hora y el desfase conservan su comportamiento.
5. Las pruebas unitarias comprueban el valor cargado, el cambio inmediato y el atributo de sólo
   lectura; Playwright comprueba la presencia del campo en PC y móvil.

### 14.5 Decisión registrada

El responsable confirmó el 2026-10-05 reutilizar exactamente el patrón de `/setup`, sin cambios en
el guardado, la API ni la base de datos.

### 14.6 Registro del repaso

El responsable confirmó el 2026-10-05 que `/settings` usará el mismo identificador técnico, la
misma apariencia y el mismo comportamiento de `/setup`, sin diferencias adicionales.

El repaso queda cerrado sin decisiones abiertas. El responsable aprobó explícitamente las dos
propuestas con «Sí». La implementación respeta el alcance aprobado y ambos documentos vuelven a
`as-built`.

### 14.7 Verificación

- Las 216 pruebas unitarias del frontend pasan; las de Configuración comprueban el valor inicial,
  el cambio inmediato y el atributo de sólo lectura.
- Los cuatro recorridos E2E afectados pasan: dos en PC y dos en móvil.
- Playwright comprueba el valor guardado, `America/Guayaquil` después de elegir, el atributo de sólo
  lectura y que el campo pertenece a la tarjeta de región y dirección.

## 15. Resultados de acciones en mensajes toast

### 15.1 Hallazgo

`/settings` y `/setup` tienen un único mensaje de pantalla colocado antes de su contenido. Guardar
una tarjeta o probar una conexión desde una sección inferior actualiza ese mensaje, pero la ventana
no vuelve al principio. La acción sí responde y el texto sí existe; queda fuera de la parte visible.
Esto explica por qué una prueba de conexión fallida parecía no mostrar ningún resultado.

### 15.2 Componente y comportamiento

Se añade un componente compartido de **toast** para los resultados transitorios de una acción:

- queda fijo en la ventana, arriba a la derecha en PC y con el ancho disponible, dejando margen, en
  móvil;
- sólo se muestra uno a la vez y el resultado más reciente sustituye al anterior;
- todas las formas desaparecen a los **10 segundos como máximo**;
- todas pueden cerrarse antes o ser sustituidas por una acción posterior;
- todos incluyen un botón de cierre con nombre accesible traducido;
- éxito e información se anuncian como estado; error y atención se anuncian como alerta, sin mover
  el foco ni desplazar el contenido;
- mantiene las cuatro formas visuales del aviso existente —éxito, error, información y atención— y
  usa los colores, iconos y contraste del tema activo.

El temporizador empieza de nuevo cuando llega un resultado distinto. Destruir la pantalla cancela
el temporizador pendiente, para que un mensaje antiguo no actúe después de navegar.

### 15.3 Qué es toast y qué sigue dentro de la página

Usan toast los resultados de acciones asíncronas que hoy alimentan el mensaje general de
`/settings` o `/setup`: guardar, subir o restaurar, avanzar en el asistente y probar una conexión.
También lo usan sus fallos de carga o de validación devueltos por la API cuando no pertenecen a un
campo concreto.

Permanecen dentro de su sección los avisos que describen un estado o una consecuencia mientras esa
condición exista, entre ellos la dirección pública sin HTTPS, la explicación de la cuenta de fábrica
y los avisos informativos propios de un paso. Los errores asociados a un campo conservan su lugar
junto al campo si una pantalla los presenta así.

### 15.4 Alcance

- Este documento define el componente compartido y su accesibilidad.
- `docs/modules/settings.md` define su uso en `/settings`.
- `docs/primer-arranque.md` define su uso en `/setup`.
- `docs/README.md` y `AGENTS.md` reflejan el estado de la enmienda.
- La implementación vive en `frontend/src/app/shared/components/**` y en las dos páginas afectadas.

Quedan fuera el backend, sus respuestas y códigos, la persistencia, las reglas de validación, los
mensajes de las demás pantallas y un servicio global de notificaciones para toda la aplicación.

### 15.5 Criterios de aceptación

1. Probar una conexión desde una tarjeta o paso inferior muestra su éxito o error dentro de la parte
   visible de la ventana, sin desplazarla al principio.
2. Guardar, subir, restaurar y avanzar muestran sus resultados por el mismo componente.
3. Todas las formas desaparecen después de 10 segundos como máximo y pueden cerrarse antes.
4. Un segundo resultado sustituye al primero y reinicia el temporizador que corresponda.
5. El botón de cierre funciona con teclado, tiene nombre accesible en español e inglés y el mensaje
   conserva el anuncio semántico adecuado.
6. Los avisos persistentes de contexto siguen dentro de su sección.
7. Las pruebas unitarias cubren sustitución, cierre y duración; Playwright comprueba la posición
   visible de los resultados de conexión en `/settings` y `/setup`, en PC y móvil.

### 15.6 Decisión registrada

El responsable confirmó el 2026-10-06 proponer el componente compartido para `/settings` y
`/setup`: fijo arriba a la derecha en PC y adaptado al ancho móvil, un mensaje a la vez, éxito con
cierre automático a los 10 segundos para todas las formas, cierre manual y avisos de contexto todavía
integrados en la página.

### 15.7 Registro del repaso

El responsable confirmó el 2026-10-06 que:

1. todos los resultados transitorios descritos de `/settings` y `/setup` usan el toast;
2. no hay cola: el resultado nuevo sustituye al anterior;
3. todas las formas desaparecen automáticamente después de 10 segundos como máximo y también
   pueden cerrarse o ser sustituidas antes.

El repaso queda cerrado sin decisiones abiertas. El responsable aprobó explícitamente los tres
documentos con «Apruebo» el 2026-10-06. La implementación respeta el alcance aprobado.

### 15.8 Implementación y verificación

- `app-toast` envuelve el aviso compartido, permanece fijo en la ventana y emite el cierre manual o
  automático; al destruirse cancela su temporizador.
- `app-aviso` admite un botón de cierre opcional. Los avisos integrados no lo reciben y conservan su
  presentación; error y atención usan `alert`, éxito e información usan `status`.
- Las 220 pruebas unitarias del frontend pasan; cuatro casos nuevos comprueban cierre, semántica,
  duración, persistencia y reinicio del temporizador.
- La compilación de producción termina correctamente.
- Los 20 casos Playwright seleccionados pasan en PC y móvil; 4 se omiten porque los servicios
  opcionales de directorio e IA no estaban levantados. La preparación de la instalación vacía
  comprobó también el toast de `/setup` en ambos anchos.

## 16. Implementación: idioma global y configuración guiada de IA

### 16.1 El idioma es el primer control

En `/setup` y `/settings`, Español/English aparece antes que cualquier otro dato. En el asistente
cambia todos los textos al elegirlo. En Configuración previsualiza la pantalla, pero una banda
explica que el cambio global todavía no se ha aplicado mientras se traducen y revisan correos.

Desaparecen todos los selectores personales de idioma: entrada, controles del armazón, perfil, alta
y ficha de usuario. El espacio liberado en el menú no se rellena con otro control. El tema continúa
siendo personal.

### 16.2 Selector de modalidad de IA

La tarjeta y el paso usan tres opciones grandes con nombre y explicación: **En este servidor**,
**En otro servidor** y **Proveedor externo**. Sólo se muestran los campos de la opción elegida. La
modalidad externa incluye un aviso persistente con los datos que saldrán y una casilla de aceptación
antes del botón **Probar y activar**.

Para proveedor se muestran OpenAI, Claude, DeepSeek y Compatible con OpenAI. El modelo es un
desplegable de sugerencias con **Otro modelo…**; al elegirlo aparece el identificador libre. Los
secretos usan campo de contraseña, indican si existe uno guardado y nunca muestran puntos que puedan
confundirse con su valor real.

### 16.3 Catálogo local y progreso

Cada modelo muestra nombre, tamaño de descarga, RAM recomendada y estado: no instalado, descargando,
instalado, activo o error. El actual 1.5B aparece primero y marcado como recomendado. Antes de una
descarga se muestran disco y RAM disponibles y se bloquea con una explicación si no alcanzan.

La ficha enlaza la fuente oficial y muestra licencia y checksum. Antes de la primera descarga de una
versión se abre la licencia y se pide una aceptación explícita; la pantalla dice quién la aceptó y
cuándo. Un cambio de versión, licencia o checksum vuelve a pedirla.

La descarga tiene barra de progreso, bytes descargados/total y acción de cancelar. Recargar conserva
el avance real que devuelve el servidor. Activar enseña las etapas «cargando», «probando» y «activo».
Los modelos inactivos tienen **Eliminar** con confirmación; el activo nunca ofrece esa acción.

En móvil las fichas se apilan, la barra ocupa el ancho y las acciones quedan debajo. Ningún estado
depende sólo del color; progreso y resultado se anuncian de forma accesible. Los resultados breves
usan el toast existente, mientras descarga, requisitos y estado activo permanecen dentro de la
tarjeta.

### 16.4 Revisión del cambio de idioma

La revisión presenta las trece plantillas como lista de pendientes. Al abrir una se ven origen y
borrador lado a lado en PC y apilados en móvil, con asunto, cuerpo y marcadores. Si ya existe un
destino personalizado, aparece como tercera versión y el Administrador elige conservarlo, usar el
borrador o combinarlos en el editor. No se reemplaza en silencio. **Aplicar idioma** sólo se habilita
cuando las doce están válidas; cancelar descarta los borradores y vuelve al idioma persistido.

Tras confirmar, la interfaz completa adopta el nuevo idioma. Los resúmenes anteriores se conservan y
siguen visibles en el idioma en que fueron redactados. Los nuevos y los que se vuelvan a pedir se
generan en el nuevo idioma global; no existe regeneración masiva al cambiarlo.

La advertencia económica aparece únicamente antes de traducir las plantillas con un proveedor
externo y muestra solicitudes, tokens y costo estimados. Rechazarla abre la misma revisión con los
destinos editables manualmente. La edición manual permanece disponible siempre desde Configuración.

Las sesiones que ya estaban abiertas reciben idioma y versión de configuración en la siguiente
respuesta autenticada. El interceptor actualiza la señal global y repinta textos y fechas sin
cerrar sesión ni consultar periódicamente. Hasta que esa pestaña haga otra petición conserva la
vista que ya tenía.

### 16.5 Criterios de aceptación propuestos

Playwright cubrirá los tres modos, secreto oculto, aviso de privacidad, prueba fallida que conserva
la configuración anterior, descarga/progreso/reanudación, licencia, recursos insuficientes,
activación y borrado, revisión con destino personalizado, confirmación comercial, cambio global,
actualización de una sesión abierta y ausencia de selectores personales, en PC y móvil.

### 16.6 Registro del repaso

El responsable cerró el repaso el 2026-10-06. Eligió una clave de cifrado obligatoria por entorno,
comparación de tres versiones para no pisar plantillas personalizadas, recuperación comercial bajo
decisión del Administrador, licencia e integridad visibles y actualización del idioma en sesiones
abiertas sin salir ni hacer sondeo. No quedan decisiones abiertas; propuesta aprobada.

## 17. Memoria coherente al activar modelos

Los selectores y fichas muestran mínimos de **2 GiB, 4 GiB y 6 GiB** para 1.5B, 3B y 7B. No mezclan
GB decimales con GiB ni presentan una estimación inferior a la puerta que aplica el administrador.

Cuando no cabe un modelo, el toast permanece hasta cerrarse y dice cuál se intentó activar, cuánta
RAM requiere y cuánta estaba disponible, con valores formateados en GiB. El modelo anterior continúa
marcado «Activo · Motor disponible». El mensaje funciona igual en `/setup` y `/settings`, en PC y
móvil, y no recomienda forzar swap.

La comprobación real en PC y Pixel 7 mostró 2,00/4,00/6,00 GiB. Al intentar el 7B, el toast permaneció
visible con 6,00 GiB requeridos y entre 5,06 y 5,08 GiB disponibles, mientras el 3B continuó marcado
«Activo · Motor disponible».

## 18. Repaso integral de formato

### 18.1 Alcance

El bloque 6 revisa todas las pantallas existentes: sesión, `/setup`, armazón, configuración, correo,
usuarios y tickets. Incluye estados vacíos, listas con contenido, formularios, diálogos, toasts,
adjuntos, vistas dobles y errores visibles. Se comprueban español e inglés en Chromium a
**1440×900**, **768×1024** y **412×915**.

La revisión cubre espaciado, alineación, jerarquía, legibilidad, cortes, desbordamientos, adaptación
responsive, foco visible, navegación por teclado, nombres accesibles, semántica, contraste y áreas
táctiles conforme a los criterios aplicables de **WCAG 2.2 AA**. Los ocho temas mantienen su prueba
automatizada de contraste; los temas claro y oscuro de fábrica reciben además el recorrido visual
detallado. Los seis temas fijos se comprueban mediante sus variables y pruebas comunes.

### 18.2 Inventario y prioridad

Cada hallazgo se registra antes de corregirse con pantalla, idioma, tema, resolución, forma de
reproducirlo y evidencia medible. La prioridad significa:

- **crítico**: bloquea una tarea;
- **alto**: impide usarla en un dispositivo o sin ratón;
- **medio**: dificulta lectura, comprensión o consistencia;
- **menor**: defecto cosmético perceptible.

El criterio de cierre elegido es **cero hallazgos abiertos de cualquier nivel**. Si aparece un cambio
que afecta comportamiento, alcance, datos, permisos o flujo, se detiene esa corrección y se abre otro
ciclo documental; no se reclasifica como formato para introducirlo en este bloque.

### 18.3 Regla de corrección y evidencia

Cuando el defecto se repite, se corrige el componente o estilo compartido y se verifican todos sus
usos. Un ajuste local se reserva para una composición propia de una sola pantalla. Se prefieren
mediciones y aserciones reproducibles sobre capturas de referencia: ancho, posición, desbordamiento,
orden de foco, nombre accesible, contraste y tamaño de objetivo se automatizan cuando aportan una
señal estable. El recorrido manual queda documentado con su resultado.

No se introduce una biblioteca de capturas visuales, no se añaden Firefox o WebKit, no se rediseñan
recorridos y no cambian backend, API, persistencia, permisos ni reglas del producto. Producción sigue
aparcada hasta el cierre de la versión 1.0.0.

### 18.4 Decisiones y repaso

El responsable eligió 1A–4A: todas las pantallas; formato, accesibilidad y responsive; PC, tableta y
móvil; inventario completo y corrección por prioridad. Eligió 5A–9A: ocho temas con recorrido
detallado en los dos de fábrica; ambos idiomas; Chromium; evidencia medible; y corrección común para
defectos repetidos. Cerró el repaso con 10A–13A y 14C: WCAG 2.2 AA aplicable, cuatro prioridades,
datos desechables con seeders, sin cambios funcionales y cero hallazgos abiertos. No quedan
decisiones abiertas.

### 18.5 Resultado as-built

El inventario recorrió las pantallas autenticadas de los cuatro papeles y las pantallas públicas en
español e inglés, con los temas claro y oscuro de fábrica y las tres resoluciones aprobadas. `/setup`
se comprobó además vacío en sus cuatro estados navegables: fueron **48 combinaciones** del asistente,
**150 vistas autenticadas y 18 públicas por idioma**. Las excepciones legítimas de WCAG para enlaces
en texto, controles con separación suficiente y contenido del correo no se registraron como defectos.

Se encontró un hallazgo medio: el conmutador de papel de `/users` medía 424 px en una ventana de
412 px cuando estaba en inglés. El componente compartido `Conmutador` ahora limita su ancho y parte
sus opciones en varias líneas; la corrección se verificó en los tres tamaños, ambos temas y con los
papeles Administrador y Soporte. No cambió ningún flujo ni contrato.

El cierre dejó **cero hallazgos abiertos**. Pasaron las 220 pruebas unitarias del frontend, 18 casos
afectados de Configuración con 4 omisiones previstas y la suite completa de Playwright con 180 casos
aprobados y 38 omisiones previstas. Los ocho temas conservaron su contraste medido.

## 19. Implementación: modal «Mejorar con IA»

### 19.1 Entrada y contenido

Los editores de descripción y comentarios permitidos a Soporte y Desarrollo incorporan, junto a sus
controles, un botón con icono y texto **«Mejorar con IA»**. En móvil puede ocupar su propia fila para
no comprimir el editor. No aparece para Usuario ni Administrador.

El botón exige un borrador no vacío y abre un modal con una copia en un `textarea` sencillo: sin
negrita, cursiva, menciones ni adjuntos. Editar esa copia no toca el editor principal. Debajo se elige
una de cinco tonalidades: profesional, cordial, breve, empática o técnica. Una nota indica que el
motor configurado procesará el texto y, cuando no sea local, que saldrá de la instalación conforme a
la aceptación administrativa existente.

### 19.2 Generación y revisión

«Mejorar» inicia una sola solicitud y queda deshabilitado mientras responde. El resultado sustituye
la copia dentro del modal para poder revisarlo y editarlo. Desde allí se puede **«Usar este texto»**,
**«Volver a generar»** o **«Cancelar»**. Regenerar usa el contenido visible y permite cambiar la
tonalidad. Usar el texto reemplaza el borrador del editor principal como texto plano con saltos de
línea; todavía no guarda ni publica nada.

Cancelar o cerrar conserva intacto el editor principal. Si se cierra durante una solicitud, una
respuesta tardía se descarta. Un fallo permanece dentro del modal, conserva su contenido y ofrece
reintentar o cancelar; no se muestra lejos del contexto mediante un toast.

### 19.3 Accesibilidad, adaptación y pruebas

El modal usa las piezas compartidas existentes, tiene título y descripción accesibles, atrapa y
restaura el foco, se cierra con Escape cuando no existe otra confirmación pendiente y anuncia carga,
error y resultado. En 412 px no crea desplazamiento horizontal y sus acciones pueden partirse en
varias filas.

Las pruebas unitarias cubren el contrato de la ruta, permisos por papel, texto vacío, tonalidad y
editor inválidos, contexto mínimo del prompt, respuesta, error conservando el borrador y aplicación
sin guardado automático. El recorrido Playwright comprueba en PC y móvil el modal, la tonalidad, la
revisión, la aplicación local y la ausencia del control para el Usuario. El cierre general previo de
la interfaz mantiene cubiertos ambos idiomas y los temas claro y oscuro de fábrica.

### 19.4 Decisiones y alcance

El responsable eligió 1A–19A y añadió que la IA reciba el nombre de la persona que requiere soporte
en el principal o de la persona de Soporte en el interno, para usarlo únicamente cuando la redacción
lo requiera. No quedan decisiones abiertas. La propuesta no introduce nuevos componentes de diseño,
formatos, adjuntos, persistencia, cuotas ni configuración administrativa.

## 20. Implementación: visibilidad de «Mejorar con IA»

La ficha sólo monta los botones «Mejorar con IA» y el modal compartido cuando el detalle devuelve
`capabilities.aiWriting: true`. Con `false` o con el campo ausente, no deja un control deshabilitado ni
un aviso ocupando espacio: la ayuda sencillamente no aparece. Esta regla se suma al borrador no vacío,
que continúa deshabilitando el botón una vez que la capacidad existe.

La pantalla no consulta configuración, salud ni proveedores por su cuenta. Al recargar el ticket
recibe otra vez la capacidad y refleja cualquier cambio administrativo. Una IA configurada pero
temporalmente caída mantiene el botón; el error sólo aparece si la persona intenta generar, dentro
del modal y conservando el texto.

Las pruebas de componente cubren capacidad verdadera, falsa y ausente. Playwright verifica en PC y
móvil que una instalación configurada muestre el botón y mantenga el recorrido de revisión. El caso
sin configuración queda en backend y componente: un navegador real sería dirigido a Configuración
antes de alcanzar esta vista, conforme a la IA obligatoria. No se modifican el modal, sus
tonalidades, los editores, la configuración administrativa ni el bloqueo global. El responsable
eligió 1A, 2A y 3A y confirmó mantener la obligatoriedad; no quedan decisiones abiertas.

## 21. Implementación: «Código fuente y licencia»

La pantalla de entrada muestra en su pie un enlace **«Código fuente y licencia»**, junto a la
versión sin competir con el formulario. El menú lateral no lo repite: su zona inferior conserva
sólo la cuenta, el tema, el idioma, salir y la versión.

La entrada lee `license` y `sourceUrl` de la marca pública. El enlace abre la fuente
correspondiente en otra pestaña con las protecciones habituales. Una compilación de desarrollo lo
indicará en el texto asociado a la versión. Si un backend anterior no entrega una URL válida, el
control no aparecerá. No se codifica una URL alternativa en el frontend.

Se comprobaron nombre accesible, foco, contraste, ausencia de desbordamiento y destino en PC y
móvil. No añade una pantalla legal, una ventana modal ni un control administrativo. El alcance y
la política de marca viven en `docs/proyecto-abierto.md`; el contrato, en
`docs/modules/settings.md`, sección 14.

La verificación original del 2026-10-10 cubrió las dos ubicaciones. El ajuste textual del
2026-10-15 retiró la repetición del menú por petición del responsable y conservó el enlace en la
entrada; `marca.spec.ts` comprueba ahora expresamente ambas condiciones.

## 22. Prioridad visible y operable

El alta ofrece las cuatro prioridades con media seleccionada. En listas y detalle aparece un chip
con texto traducido e icono; color e icono ayudan, pero el significado nunca depende de ellos. El
detalle permite a Soporte y, desde la vista interna autorizada, a Desarrollo cambiar el valor en el
lugar con Guardar y Cancelar. El Usuario y el Administrador lo ven sin editor. Cerrado no ofrece el
control.

El modal de búsqueda admite seleccionar varias prioridades. Las bandejas ordenan de fábrica crítica,
alta, media y baja, y dentro de cada grupo muestran primero el ticket más antiguo. La persona puede
cambiar el orden. Filtro y orden se reflejan en parámetros de URL y sobreviven a recarga, regreso y
enlaces compartidos. Tabla y tarjetas móviles muestran el mismo dato sin crear desplazamiento
horizontal.

El historial escribe las etiquetas traducidas anterior y nueva, el autor y la fecha. Si falla el
correo crítico, un toast persistente de atención confirma que la prioridad quedó guardada y que el
aviso no salió; no ofrece guardar de nuevo ni afirma que la operación falló completa.

## 23. Progreso, prueba local y logos en `/setup`

La ficha del modelo muestra una barra nativa accesible, porcentaje y `descargado / total`. Se
actualiza sin pulsar «Ver y actualizar el catálogo» mientras exista una fase transitoria y deja de
consultar al terminar o abandonar la pantalla. El texto nombra descarga, validación, activación y
prueba; ningún estado depende sólo del color. En móvil la barra ocupa el ancho y las acciones se
apilan.

Sólo se inicia una descarga. Durante ella aparecen **Cancelar** y la explicación de que el parcial se
conserva. Después aparecen **Reanudar** y **Eliminar descarga parcial**. El modelo instalado inactivo
ofrece **Eliminar modelo** tras confirmación; el activo no lo ofrece. Un fallo de activación conserva
la ficha y añade **Reintentar activación**.

En `/setup`, **Siguiente** está deshabilitado hasta que el modelo seleccionado figure activo y
saludable. Junto al botón se dice si falta descargar, validar, activar o probar. No se usa el error
genérico del motor como validación de una acción que la propia pantalla dejó ejecutar antes de
tiempo.

El activo saludable ofrece **¡Pruébame!** o **Try me!**. Mientras trabaja conserva el botón con su
estado accesible. Después enseña dentro de la ficha el saludo del modelo y, por separado, la fecha,
hora y zona IANA que el servidor le suministró. El resultado no usa toast porque debe poder leerse y
compararse; desaparece al repetir, cambiar de activo o salir. El error ocupa el mismo lugar y permite
reintentar.

El paso 1 coloca, después de idioma y nombre, las dos tarjetas de logo claro y oscuro que ya existen
en Configuración. Cada una enseña vista previa, estado de fábrica o propio, subir y restaurar. La
vista previa y el logo del encabezado cambian al guardar. Ninguna variante es obligatoria.

Las pruebas de componente y Playwright cubren textos completos en ambos idiomas, teclado, anuncios
accesibles, PC, móvil, recarga, cancelación, reanudación, fondo, error, borrado, saludo fechado y los
dos logos. Los toasts generales continúan reservados para resultados breves; progreso, saludo y
errores operativos permanecen en su ficha.

## 24. Acciones directas en fichas, saludo visible y carga desde el logo

### 24.1 Hallazgos

La implementación actual obliga a volver al selector situado antes del catálogo para escoger el
modelo, aunque cada ficha ya contiene sus demás acciones y estados. El saludo de prueba queda dentro
de una ficha larga y puede pasar inadvertido. En los logos, un control nativo separado enseña el
nombre del archivo y no deja claro que la vista previa es el lugar que se reemplaza.

La política vigente de la sección 15 tampoco coincide con la nueva decisión: éxito e información
duran cinco segundos, mientras error y atención pueden permanecer indefinidamente.

### 24.2 Modelos

Desaparece el selector superior del modelo local en `/setup` y `/settings`. Cada ficha instalada
incluye su propio botón **Activar** o **Desactivar** y un indicador textual de si está activa. La ficha
es la única fuente de selección y mantiene visibles descarga, progreso, salud, prueba y eliminación.

Se permite desactivar el único modelo activo. El motor queda sin modelo; `/setup` deshabilita
**Siguiente** y explica que falta activar uno, mientras `/settings` muestra esa condición sin impedir
administrar el resto. Desactivar no borra el archivo, por lo que puede volver a activarse desde la
misma ficha. El modelo activo sigue sin ofrecer eliminación.

### 24.3 Saludo y toast

**¡Pruébame!** / **Try me!** muestra el resultado mediante el toast compartido. El mensaje identifica
el nombre del modelo que respondió y contiene sólo su saludo. La fecha y hora suministradas ya forman
parte del saludo generado, por lo que no se repiten, y la zona IANA no se presenta. El resultado
anterior deja de ocupar espacio dentro de la ficha.

La política común de todo `app-toast` pasa a ser: cierre manual siempre disponible y desaparición
automática a los **10 segundos como máximo**, incluidas las formas de error y atención. Un resultado
nuevo continúa sustituyendo al anterior y reinicia el temporizador. Los avisos persistentes que
describen un estado siguen dentro de su sección y no se convierten en toast.

### 24.4 Logos

La vista previa de cada logo claro u oscuro es un botón accesible. Pulsarla abre el selector de
archivo; al elegir uno válido, se guarda y la imagen se actualiza. El control de archivo queda
oculto visualmente y no se enseña ruta ni nombre del archivo. Se conserva una acción separada para
volver al logo de fábrica y se indica con texto accesible qué variante se reemplazará.

### 24.5 Pruebas y alcance

Las pruebas de componente y navegador cubren teclado y lector de pantalla, activación y
desactivación desde cada ficha, ausencia del selector superior, bloqueo de **Siguiente**, toast del
saludo con nombre del modelo, cierre y límite de diez segundos para las cuatro formas, y carga de
ambos logos pulsando su vista previa en PC y móvil.

No cambia el catálogo 1.5B/3B/7B, la exclusión de descargas, las licencias, los límites de archivo,
las modalidades externas ni producción. El responsable eligió 1A, 2A y corrigió 3A a un máximo de
diez segundos cerrable para todos los toast. No quedan decisiones abiertas.

La implementación pasó las 227 pruebas del frontend y la compilación de producción. En el navegador
real se verificaron las fichas sin selector superior, desactivación y reactivación, bloqueo de
**Siguiente**, atribución del saludo sin repetir el reloj ni mostrar la zona, desaparición del toast a los diez
segundos y controles accesibles de logo sin nombre ni ruta de archivo.

## 25. Confirmaciones temáticas y apariencia en el paso 1

### 25.1 Hallazgos

Las cuatro confirmaciones de eliminación del administrador local de modelos usan `confirm()` del
navegador. Esa ventana pertenece al sistema, no recibe las variables del tema y en un tema oscuro
aparece blanca. Es la única confirmación de producto que todavía evita el componente propio.

El primer paso permite decidir idioma, nombre y logos, pero no presenta el selector de los ocho
temas ni el color institucional que ya existen en el armazón y en Configuración. La instalación se
puede personalizar, pero no ver completa mientras se configura.

### 25.2 Un solo diálogo para todas las confirmaciones

Toda confirmación de la aplicación usa `app-dialogo`, nunca `confirm()` del navegador. La regla
incluye las cuatro eliminaciones actuales de modelo instalado o descarga parcial en `/setup` y
`/settings`, y cualquier confirmación posterior. El diálogo hereda las variables del tema activo,
usa el botón peligroso para confirmar y el secundario para cancelar, atrapa el foco, devuelve el foco
al disparador y cierra con Escape. El texto continúa traducido y especifica qué recurso se elimina.

### 25.3 Tema y color en el paso 1

Junto al selector de idioma aparece el mismo selector completo del armazón: ocho temas agrupados en
claros y oscuros. Cambia la interfaz inmediatamente y guarda la preferencia en el navegador mediante
`ThemeService`; no se añade una preferencia a la instalación ni una opción «Automático».

El mismo bloque incorpora el control de color institucional de Configuración: selector libre, valor
hexadecimal y vista previa resuelta por el backend. El cambio se ve antes de guardar. Nombre, idioma
y color se guardan juntos al pulsar **Siguiente**. Cambiar de tema conserva el color elegido; sólo los
dos temas de fábrica usan ese acento, mientras los seis temas fijos mantienen su paleta.

### 25.4 Pruebas y alcance

Las pruebas cubren ausencia de `confirm()`, apertura, cancelación y confirmación del diálogo en los
ocho temas; foco, Escape y textos ES/EN; selección y memoria de los ocho temas; vista previa y
persistencia del color al avanzar; conservación del color al cambiar de tema; y presentación en PC
y móvil.

No cambian las ocho paletas, la derivación accesible del color, los límites de logo, el catálogo de
modelos, los permisos ni producción. El responsable eligió 1A, 2A y 3A para confirmaciones y tema, y
1A, 2A y 3A para el color. No quedan decisiones abiertas.

La implementación sustituye las cuatro llamadas nativas existentes y deja cero `confirm()` de
producto. Las 227 pruebas del frontend y la compilación de producción pasan. En navegador real se
verificaron memoria del tema tras recargar, presentación a 390 px, diálogo en tema oscuro, cierre
con Escape y devolución del foco al disparador, sin ejecutar el borrado.

## 26. Catálogo local siempre visible

### 26.1 Hallazgo

En `/setup` y `/settings`, el encabezado **Modelos locales** aparece antes de consultar el catálogo,
pero sus fichas permanecen ocultas hasta pulsar **Ver y refrescar el catálogo**. La acción mezcla dos
responsabilidades: revelar contenido que debe estar disponible desde el inicio y actualizarlo.

### 26.2 Carga y refresco

Al abrir el paso de IA de `/setup`, la tarjeta de IA de `/settings` o al cambiar a modalidad local,
la pantalla consulta automáticamente el catálogo. Las fichas, el espacio disponible y sus estados
se presentan sin una acción previa de la persona.

El botón pasa a llamarse **Refrescar catálogo** / **Refresh catalog**. Al pulsarlo vuelve a consultar
el estado real, muestra **Refrescando…** / **Refreshing…**, queda temporalmente deshabilitado y
conserva visibles las fichas actuales hasta recibir la respuesta. Una respuesta correcta sustituye
los datos en el sitio; un error conserva el catálogo anterior y se comunica mediante el toast común.

### 26.3 Pruebas y alcance

Las pruebas cubren carga automática, cambio a modalidad local, refresco manual sin ocultar fichas,
bloqueo de doble pulsación, error conservando datos y presentación en PC y móvil, en `/setup` y
`/settings`.

No cambia el catálogo 1.5B/3B/7B, las descargas, activación, licencias, API, permisos ni producción.
El responsable eligió 1A, 2A y 3A. No quedan decisiones abiertas.

La implementación pasó las 228 pruebas del frontend y la compilación de producción. El navegador
real confirmó en `/setup` que las tres fichas ya están presentes al entrar al paso 5 y continúan
visibles después de pulsar **Refresh catalog**.

## 27. Dos paneles de ticket independientes

### 27.1 Hallazgos

La vista **Los dos** coloca principal e interno dentro del mismo flujo vertical. El desplazamiento
pertenece a la página, por lo que mover una conversación mueve también la otra y se pierde la
referencia entre eventos. Además, el detalle conserva el máximo global de 88 rem aunque haya ancho
disponible alrededor, dejando una franja sin utilizar en pantallas grandes.

En anchos inferiores al punto de dos columnas, la implementación actual puede apilar ambos tickets,
aunque la regla documentada ya pedía ver uno por vez. Se corrige el desajuste de forma explícita.

### 27.2 PC: altura y desplazamiento propios

En PC, **Los dos** ocupa todo el ancho útil del área de contenido y lo divide en dos columnas iguales.
La cabecera general —volver, números, estado y conmutador— permanece visible. Debajo, cada ticket
ocupa la altura restante de la ventana y crea su propia región de desplazamiento vertical. La rueda,
el gesto, Page Up/Page Down y las barras actúan únicamente sobre la columna que tiene el puntero o
el foco.

Cada panel conserva su cabecera y su formulario de comentario dentro del contenido desplazable. El
editor aparece al final de su conversación; no queda fijo ni reduce permanentemente el área de
lectura. Las columnas no son redimensionables y parten el espacio 50/50.

### 27.3 Anchos estrechos

En móvil y tableta vertical no se apilan ambos paneles. El conmutador muestra Principal o Interno,
uno por vez, usando el ancho completo. La opción **Los dos** sólo produce las dos columnas cuando el
ancho permite compararlas; al entrar o estrechar la ventana se conserva el ticket del papel o el
último lado individual elegido.

Las vistas individuales mantienen su ancho y desplazamiento actuales. El cambio de ancho completo y
scroll interno pertenece únicamente a la comparación en PC.

### 27.4 Pruebas y alcance

Las pruebas de componente y navegador cubren scroll independiente sin mover la columna vecina,
cabecera general visible, columnas iguales sin franja desaprovechada, editores al final de cada
flujo, teclado, PC ancho y una sola vista en móvil y tableta vertical.

No cambian contenido, permisos, estados, acciones, API, modelo de datos ni producción. El responsable
eligió 1A, 2A, 3A y 4A. No quedan decisiones abiertas.

La implementación pasó las 228 pruebas del frontend y la compilación de producción. En navegador
real, a 1800 px, ambos paneles midieron 728 px, el desplazamiento del principal no movió el interno
y la cabecera común conservó su posición. A 900 px desapareció la opción **Los dos** y quedó un
único ticket a ancho completo.

## 28. Contenido a todo el ancho de cada panel

### 28.1 Hallazgo

La sección 27 ensanchó correctamente el contenedor exterior y repartió el área útil en dos mitades,
pero cada `app-ticket-vista` conserva por dentro la cuadrícula `2fr 1fr` de la vista individual. En
**Los dos**, la ficha de la columna derecha se oculta por decisión anterior, aunque su fracción sigue
reservándose. Por eso la conversación ocupa sólo dos tercios de cada mitad y quedan vacíos amplios
entre ambos tickets y a la derecha. La documentación decía que los paneles aprovechaban el ancho,
pero no registraba esta reserva interna; se reporta el desajuste antes de corregirlo.

### 28.2 Disposición aprobable

Cuando `soloLaConversacion` está activo, la cuadrícula interna tendrá una sola columna que ocupa el
100% del panel. Cabecera, asunto, motivo, descripción, comentarios, historial y editor de comentario
participan del mismo flujo ancho; no se añade un máximo intermedio ni una columna fantasma.

La separación entre principal e interno continúa en 24 px. El cambio se limita a **Los dos**: las
formas individuales mantienen su cuadrícula actual de conversación y ficha, sus anchos y su
comportamiento. Tampoco cambia el scroll independiente incorporado en la sección 27.

Las previsualizaciones de imágenes y videos conservan el límite de 480 × 360: el espacio nuevo no
las agranda artificialmente. El editor ocupa el nuevo ancho, pero mantiene su altura actual.

La comparación pasa de 1280 a 1152 px. Desde ese ancho aparecen **Los dos** y sus dos mitades; por
debajo sigue mostrándose un ticket por vez. El cambio permite usar la comparación en portátiles
pequeños sin llevarla hasta el ancho típico de una tableta vertical.

### 28.3 Pruebas y alcance

Las pruebas de componente comprobarán que `soloLaConversacion` selecciona una columna completa y
que una vista individual conserva `2fr 1fr`. El navegador verificará en PC que mensajes y editores
llegan al ancho útil de cada mitad, que sólo queda el espacio de 24 px entre paneles y que no aparece
scroll horizontal. Se repetirá la comprobación de scroll independiente y se probarán los límites de
1151 y 1152 px, además de una vista a 900 px.

No cambian textos, contenido, permisos, estados, acciones, API, modelo de datos ni producción. El
responsable eligió 1A, 2A, 3A, 4A, 5A, 6A, 7B y 8B. No quedan decisiones abiertas.

La implementación pasó las 229 pruebas del frontend y la compilación de producción. En navegador
real, a 1800 px, cada panel midió 728 px y su contenido 720 px, sin desbordamiento horizontal. A
1152 px se mostraron dos paneles de 404 px con 396 px de contenido; a 1151 px desapareció **Los
dos** y quedó una sola vista. El scroll independiente y las vistas individuales se conservaron.

## 29. Implementación: «Mejorar con IA» sin modal

### 29.1 Generación inmediata

El botón **Mejorar con IA** continúa junto a cada editor autorizado y exige texto. Al pulsarlo inicia
la solicitud en ese lugar, muestra actividad y usa siempre tono profesional. No abre modal, no pide
confirmación previa y no muestra selector de tonalidad. El aviso sobre tratamiento por motores
externos permanece únicamente en Configuración, conforme a la aceptación administrativa vigente.

Descripción y comentario administran estados independientes: cada uno puede conservar su propia
propuesta sin expulsar la del otro. Durante la espera el editor continúa disponible.

### 29.2 Panel de revisión

La respuesta aparece bajo el editor que la solicitó, en una región con título accesible. Es texto
plano de sólo lectura, visible y seleccionable; no se añade botón **Copiar**. Sus acciones son
**Usar propuesta** y **Conservar original**. La primera reemplaza el borrador local y cierra el
panel; la segunda descarta la salida, oculta el panel y deja el original intacto.

Con una propuesta vigente, el botón pasa a **Generar otra propuesta**. La nueva solicitud parte del
borrador actual y pide una paráfrasis profesional diferente de la última propuesta. Sólo se conserva
la salida más reciente.

Si la persona modifica el borrador después de iniciar la solicitud, ambos textos permanecen, pero
la propuesta se marca como desactualizada y **Usar propuesta** queda deshabilitado. Puede copiarse
por selección, conservarse el original o generarse otra propuesta desde el borrador nuevo. Así una
respuesta tardía nunca sobrescribe escritura reciente.

### 29.3 Errores, accesibilidad y pruebas

Un error no abre panel ni altera el borrador: usa el toast común, cerrable y con máximo de diez
segundos, y devuelve el botón a su estado inicial. La carga y el resultado se anuncian; el foco no
salta al aparecer el panel y las acciones pueden partirse en móvil sin desbordamiento horizontal.

Las pruebas cubren descripción y comentario independientes, carga, resultado seleccionable,
aplicación, descarte, borrador obsoleto, regeneración distinta, error por toast, PC y móvil, ambos
idiomas y ausencia del control por permisos o configuración. Se retiran las pruebas del modal y las
cinco tonalidades. No se añade persistencia, guardado automático ni confirmación adicional.

La implementación pasó las 230 pruebas del frontend y la compilación de producción. El recorrido
real con Soporte confirmó generación inline sin modal ni tonos, región accesible, descarte sin tocar
el borrador y bloqueo de **Usar propuesta** cuando el texto cambia. No se publicó ningún comentario
durante la verificación.

## 30. Implementación: calidad visible y refresco sin pérdida

### 30.1 Motivo y última acción

Durante una regeneración, cada bloque conserva el resumen anterior con una indicación accesible de
que está desactualizado y muestra actividad sin reemplazarlo por un marcador. Si falla, aparece un
mensaje específico de **Motivo** o **Última acción** y un botón **Reintentar** en ese mismo bloque.
Cuando nunca hubo texto válido, el error ocupa su lugar; no se muestra `...`, `—` ni una respuesta
inválida como contenido exitoso.

El sondeo y los reintentos actualizan esos dos bloques en sitio. No vacían campos, no cierran
editores ni retiran propuestas de IA. El foco y la selección permanecen donde estaban.

### 30.2 Advertencia de propuesta imperfecta

Una propuesta que sigue sin alcanzar la calidad mínima después del reintento aparece en el panel
inline con una advertencia persistente. Continúa visible, seleccionable y puede aplicarse por
decisión expresa; **Usar propuesta** no queda bloqueado por esta advertencia. El panel distingue esa
calidad insuficiente del estado obsoleto por cambio del borrador, que sí impide aplicarla hasta
regenerar.

La advertencia no usa toast porque debe permanecer junto al texto que califica. Los errores de red o
del motor continúan usando el toast común. Las pruebas cubren lectura accesible, teclado, PC, móvil,
ambos idiomas, textos largos, borrador intacto y diferencia visual entre advertencia, obsolescencia
y error.

### 30.3 Textos redundantes que desaparecen

El editor de comentarios no muestra la ayuda permanente que enumera formato, arrastre, pegado,
selección de archivos y comentario compuesto sólo por adjuntos; los controles y las tres puertas de
adjuntos continúan iguales y resultan suficientes por sí mismos. Se retiran el texto inglés
**“With formatting, and with the files inside: drag them here, paste them with Ctrl+V or look for
them. A comment can carry only files.”** y su equivalente en español.

Una propuesta válida tampoco muestra el estado neutro **“The proposal is ready to review.”** ni su
equivalente en español: el título, el texto y las acciones ya explican el estado. Se conservan los
mensajes que sí cambian una decisión —carga, borrador obsoleto, calidad insuficiente y error— y sus
anuncios accesibles no dependen de esa frase visible.

La verificación real del 2026-10-15 usó `CS-2026-0027`: mostró los dos resúmenes completos, mantuvo
literalmente el borrador de comentario y su propuesta durante un reintento de **Motivo**, enseñó
la advertencia inline cuando la segunda propuesta siguió siendo demasiado parecida y conservó
**Usar propuesta** habilitado. El menú lateral dejó de repetir el enlace «Código fuente y licencia»;
la entrada continúa mostrándolo.

## 31. Desplegables con la misma presencia que los demás controles

### 31.1 Hallazgo

El componente compartido de selector tiene hoy una altura mínima de 36 px y menos espaciado
interior que los campos de texto y los botones, que usan 44 px. Aunque varios selectores nativos ya
alcanzan 44 px, la diferencia del componente común hace que las listas se vean más delgadas y sin
armonía con el resto del formulario. El documento no fijaba una altura común para este control; se
registra el desajuste antes de cambiarlo.

### 31.2 Regla visual aplicada

Todos los desplegables de la aplicación, tanto el componente compartido como los `select` nativos,
tienen una altura mínima de **44 px** y el mismo espaciado interior horizontal y vertical de los
campos de texto: 12 px a los lados y 8 px arriba y abajo. La regla se aplica en formularios,
filtros, tablas, diálogos, configuración, asistente y menú lateral.

No hay una variante compacta: PC, tableta y móvil usan la misma altura. Se conservan el ancho,
tipografía, colores, borde, foco, icono, opciones y comportamiento actuales; este cambio no altera
el contenido ni el número de elementos que muestra cada lista.

### 31.3 Pruebas y alcance

Las pruebas de componente comprueban la altura y el espaciado del selector compartido. La revisión
del frontend confirmó que no quedan selectores nativos por debajo de 44 px y el navegador verificó una
muestra de formularios, filtros y menú lateral en PC y móvil, con sus temas claro y oscuro.

No cambian textos, datos, permisos, navegación, API, modelo de datos ni producción. El responsable
eligió 1A, 2A, 3A, 4A y 5A. No quedan decisiones abiertas.

La implementación pasó las 232 pruebas del frontend y la compilación de producción. En navegador
real, los ocho desplegables visibles de Configuración midieron 44 px y mostraron 8 px de espaciado
vertical y 12 px horizontal, tanto a 1440 × 900 como a 390 × 844. La comprobación incluyó tema,
idioma, método de entrada, reparto, avisos e IA, y se repitió con los temas claro y oscuro.

## 32. Corrección: desplegables de 52 px

### 32.1 Hallazgo posterior

La implementación de la sección 31 está presente: los controles se midieron en 44 px después de
recargar y ya no usan los 36 px anteriores. Sin embargo, la revisión visual del responsable en el
detalle del ticket mostró que **Asignar a**, el estado y **Tema** siguen percibiéndose demasiado
delgados respecto de los demás controles. La reducción de la captura acentuaba esa impresión, pero
la medida aplicada tampoco resuelve la presencia visual buscada.

### 32.2 Corrección aplicada

Todos los desplegables pasan a una altura mínima de **52 px**. La medida se aplica al componente
compartido y a cualquier `select` nativo, sin excepciones para formularios, diálogos, filtros, tablas,
asistente, Configuración o menú lateral. PC, tableta y móvil usan la misma altura.

Se conserva la tipografía `text-sm`, junto con los colores, bordes, foco, icono, ancho, opciones y
comportamiento actuales. El espaciado puede crecer sólo lo necesario para centrar el contenido dentro
de los 52 px; no se agranda el texto ni se introduce una variante compacta.

### 32.3 Pruebas y alcance

La prueba del componente exige 52 px y comprueba que la tipografía continúa en `text-sm`. La
revisión de código cubre los selectores nativos. El navegador mide los controles servidos realmente,
sin depender de una captura redimensionada.

No cambian textos, datos, permisos, navegación, API, modelo de datos ni producción. El responsable
eligió 1A, 2A y 3A. No quedan decisiones abiertas.

La implementación pasó las 232 pruebas del frontend y la compilación de producción. Una pestaña
nueva contra el servidor de desarrollo midió el selector de tema en **52 px**, con fuente de 14 px
y 8 px de espaciado vertical. El servidor ya había reconstruido automáticamente; no fue necesario
reiniciar los contenedores.

## 33. Ajuste final de desplegables a 48 px

### 33.1 Hallazgo visual

Los 52 px de la sección 32 se cargaron y midieron correctamente en un navegador distinto. La
primera pestaña conservaba el estilo anterior en caché, por lo que no era un fallo del contenedor ni
de la reconstrucción. Una vez visto el cambio real, el responsable determinó que 52 px ocupan
demasiado espacio.

### 33.2 Ajuste aplicado

Todos los desplegables tienen una altura mínima de **48 px**, como punto medio entre los 44 px que
se percibían delgados y los 52 px sobredimensionados. La medida continúa siendo global: componente
compartido y `select` nativos, formularios, diálogos, filtros, tablas, cabeceras, asistente,
Configuración y menú lateral, en PC, tableta y móvil.

Se conservan `text-sm`, los 12 px de espaciado horizontal y los 8 px verticales, además de colores,
bordes, foco, icono, ancho, opciones y comportamiento. No se crea una variante compacta.

### 33.3 Pruebas y alcance

La prueba del componente exige 48 px y conserva las aserciones de tipografía y padding. La
revisión cubre también el selector nativo, y un navegador limpio mide el resultado servido para
evitar confundir caché con la versión vigente.

No cambian textos, datos, permisos, navegación, API, modelo de datos ni producción. El responsable
eligió 1A, 2A y 3A. No quedan decisiones abiertas.

La implementación pasó las 232 pruebas del frontend y la compilación de producción. Una pestaña
nueva midió el selector servido en **48 px**, con fuente de 14 px, 12 px de espaciado horizontal y
8 px vertical. El frontend se reconstruyó automáticamente y no fue necesario reiniciar contenedores.

## 34. Desplegables iguales al conmutador segmentado

### 34.1 Medida de referencia

El conmutador Main/Both/Internal combina botones de 36 px, 2 px de padding exterior por lado y un
borde de 1 px por lado: su altura exterior es **42 px**. Los desplegables de 48 px aprobados en la
sección 33 quedan 6 px por encima y se ven desalineados cuando ambos controles comparten cabecera.

### 34.2 Normalización aplicada

Todos los desplegables usan una altura **fija de 42 px**, igual a la altura exterior del conmutador.
La regla es global y alcanza el componente compartido y los `select` nativos en formularios,
diálogos, filtros, tablas, cabeceras, asistente, Configuración y menú lateral, en PC, tableta y
móvil. El conmutador no cambia.

Se conservan `text-sm`, los 12 px de espaciado horizontal y los 8 px verticales, además de colores,
bordes, foco, icono, ancho, opciones y comportamiento. No existe variante compacta ni crecimiento
automático por contexto.

### 34.3 Pruebas y alcance

La prueba del componente exige una altura fija de 42 px y conserva las aserciones de tipografía y
padding. La revisión cubrió también el único selector nativo. Un navegador limpio midió el
desplegable compartido en **42 px**, con fuente de 14 px, padding horizontal de 12 px y vertical de
8 px; coincide con los 42 px exteriores derivados del conmutador Main/Both/Internal.

La implementación pasó las 232 pruebas del frontend y la compilación de producción.

No cambian textos, datos, permisos, navegación, API, modelo de datos, conmutadores ni producción. El
responsable eligió 1A, 2A y 3A. No quedan decisiones abiertas.

## 35. Implementación: Summary automático y sin botones

Las listas de PC, las tarjetas móviles y la ficha sustituyen **Reason/Motivo** por
**Summary/Resumen**. Last action/Última acción conserva su nombre. La tarjeta de la ficha mantiene
los dos textos, pero elimina **Try Reason again**, **Try Last action again** y sus equivalentes en
español. No aparece ninguna acción sustitutiva: ambos campos son resultados automáticos del ticket.

Sólo la ficha comunica `Generating…/Generando…` o un error breve junto al campo afectado. Conserva
la disposición y no usa toast para un proceso de fondo. En listas y tarjetas móviles, un resultado
no vigente se representa únicamente con `—`; no se añaden chips, iconos ni mensajes repetidos. El
lector de pantalla recibe el nombre nuevo y el estado cuando éste existe.

Las pruebas recorren ambos idiomas, PC y móvil, estado listo, pendiente y error, ausencia de todos
los botones y consistencia del rótulo en ficha, tabla y tarjeta. No cambian permisos, navegación,
tema, layout ni controles ajenos a insights.

## 36. Tabla resumen de tickets ordenada y legible

### 36.1 Hallazgo reproducido

En la bandeja de escritorio, el enlace **Open/Abrir** está técnicamente dentro de la tabla y ocupa
la undécima celda de cada fila. Sin embargo, el encabezado de esa columna existe sólo para lectores
de pantalla mediante `sr-only`. La captura y la medición en 1440 × 900 confirmaron que la tabla
termina exactamente después de esa celda, pero la ausencia de un rótulo visible hace que el botón
parezca colocado fuera de la tabla resumen.

### 36.2 Columnas y orden acordados

La tabla amplia conserva las once columnas, pero las ordena por la secuencia de lectura operativa:

1. **Ticket**;
2. **Priority/Prioridad**;
3. **State/Estado**;
4. **Subject/Asunto**;
5. **Classification/Clasificación**;
6. **Summary/Resumen**;
7. **Last action/Última acción**;
8. **Requested by/Solicitante**;
9. **Assignee/Responsable**;
10. **Update date/Fecha actualización**;
11. **Action/Acción**.

Summary recibe el mayor espacio flexible; después, Subject y Last action. Subject, Summary y Last
action ocupan como máximo dos líneas con puntos suspensivos y exponen el texto completo mediante la
ayuda nativa al pasar el cursor. Las filas crecen sólo hasta lo que requieren esos dos renglones y
todo el contenido se alinea arriba.

**Ticket** permanece siempre en una sola línea, sin truncado y con ancho suficiente para el número.
Priority y State ajustan su ancho al chip más largo y centran tanto encabezado como contenido.
Requested by y Assignee admiten hasta dos líneas y muestran `—` cuando no existe valor. Update date
usa una línea para la fecha y otra para la hora dentro de una columna compacta.

Classification conserva categoría y etiquetas juntas. Muestra como máximo dos líneas de chips; si
hay más etiquetas de las que caben, un indicador enfocable `+N` informa cuántas quedan ocultas. Al
pasar el cursor o enfocarlo enumera sus nombres; su nombre accesible también los enumera. No deja
crecer la fila sin límite.

La ayuda de Subject, Summary y Last action con el texto completo se abre tanto al pasar el cursor
como al enfocar la celda mediante teclado; el truncado no convierte esa información en exclusiva
del ratón.

### 36.3 Action visible y compacta

La última columna muestra el encabezado visible **Action** en inglés y **Acción** en español. El
encabezado se alinea a la izquierda, como los demás títulos; los botones permanecen centrados dentro
de sus celdas. La columna usa únicamente el ancho mínimo necesario
para el botón y su espaciado; no recibe un ancho fijo ni roba espacio adicional a asunto, Summary o
Last action.

El botón continúa diciendo **Open/Abrir** y conserva su enlace, foco, tamaño y permisos. El número
del ticket sigue siendo un enlace independiente. La fila completa no se vuelve pulsable, para no
abrir un ticket al seleccionar texto o interactuar con otro contenido.

El encabezado se mantiene en una sola línea y no ofrece ordenación: describe una acción, no un dato
ordenable.

### 36.4 Ancho intermedio y móvil

En ancho intermedio se ocultan primero Requested by y Assignee. Permanecen visibles:

`Ticket · Priority · State · Subject · Classification · Summary · Last action · Update date · Action`.

Cuando esas nueve columnas dejan de caber con legibilidad, la tabla permite desplazamiento horizontal
nativo y conserva sus medidas; no comprime los controles. Sólo por debajo de 768 px pasa a tarjetas.
El ancho real disponible sigue decidiendo si aparecen las once columnas o la variante intermedia,
por lo que responde también a abrir o plegar el menú lateral y no sólo a la ventana.

La tarjeta móvil adopta el mismo orden conceptual: Ticket, Priority, State, Subject,
Classification, Summary, Last action, Requested by, Assignee, Update date y Open. El botón
**Open/Abrir** permanece al final. Las personas ausentes muestran `—` también en la tarjeta.

### 36.5 Pruebas y alcance

Las pruebas comprueban ambos idiomas; orden y cantidad de columnas; Ticket sin salto; prioridad y
estado centrados, encabezado Action a la izquierda y su botón centrado; textos de dos líneas con contenido completo accesible; chips limitados y
contador `+N`; guiones para personas ausentes; fecha y hora en dos líneas; encabezado Action
visible, no ordenable y compacto; las nueve columnas intermedias; salto oportuno a tarjetas; y el
orden móvil. El navegador recorrerá escritorio amplio, ancho intermedio y móvil en inglés y español.
Priority y Update date conservan únicamente la ordenación que ya ofrece la bandeja; cambiar su
posición no añade encabezados interactivos ni duplica los controles de orden existentes.

No cambian datos, API, permisos, navegación, contenido de los resúmenes, filtros, orden de los
tickets, temas ni producción. Sí cambia el orden visual de las columnas y el punto responsive en el
que la tabla cede su lugar a las tarjetas, conforme a esta propuesta.

## 37. Ticket fijo durante el desplazamiento horizontal

En cualquier presentación que todavía use tabla —amplia o intermedia—, la primera columna,
**Ticket**, permanece visible al desplazar el contenido horizontalmente. Así, cada dato conserva la
referencia inequívoca al ticket de su fila. No se fija Priority ni Action y la fila completa sigue
sin ser un enlace.

Las celdas del cuerpo se fijan exclusivamente en horizontal. Desde la sección 39, el encabezado
Ticket es la esquina fija de los dos ejes y permanece también en el borde superior al bajar por una
lista larga. La columna conserva exactamente el ancho acordado en la sección 36; no gana espacio ni
obliga a truncar el número, que continúa en una sola línea.

Tanto el `th` como los `td` de Ticket usan el fondo completamente opaco de la superficie del tema,
un nivel de apilamiento suficiente y un borde derecho. El contenido desplazado nunca se transparenta
ni se dibuja por encima del número. El borde señala dónde termina la zona fija sin introducir una
sombra nueva.

Desde 768 px la tabla permanece disponible aunque su contenido sea más ancho que el contenedor y
ofrece desplazamiento horizontal nativo. No se añade degradado, texto ni icono para anunciarlo: el
propio contenido parcialmente visible indica que continúa. Por debajo de 768 px se usan tarjetas,
que no necesitan fijación ni desplazamiento lateral.

Las tarjetas no cambian: no tienen desplazamiento horizontal ni una columna que fijar. Tampoco
cambian datos, rutas, permisos, orden, filtros, paginación o comportamiento del enlace del número.

Las pruebas cubren tabla amplia e intermedia, desplazamiento horizontal real, número siempre visible,
encabezado fijo verticalmente, fondo opaco, separador derecho, foco y activación del enlace. La
comprobación visual recorre los ocho temas para confirmar que el fondo y el borde usan sus variables
y mantienen contraste.

## 38. Carga progresiva de las bandejas

Las cuatro bandejas dejan de mostrar el total, `Page/Página`, `Previous/Anterior` y `Next/Siguiente`.
La primera respuesta trae 20 tickets. Cuando la persona se acerca al final de la tabla o de las
tarjetas, un marcador no visible solicita automáticamente el bloque siguiente antes de que alcance
el último ticket.

Mientras llega ese bloque, los tickets ya visibles permanecen disponibles y un indicador discreto
al pie informa que se está cargando. Cada bloque recibido se añade en el orden entregado por el
backend, sin sustituir ni duplicar los anteriores. Una región viva accesible anuncia cuántos tickets
nuevos se incorporaron. Cuando no queda otra página, la lista termina sin mensaje adicional.

Si la carga progresiva falla, todo lo ya cargado se conserva. Al pie aparece el error con una acción
**Retry/Reintentar**; no hay reintento automático ni bloqueo de la lista. Un único bloque puede estar
en curso: varios cruces del marcador no producen solicitudes paralelas.

Cambiar búsqueda, filtros, vista u orden cancela la continuidad anterior, vacía sus resultados,
vuelve al inicio y solicita el primer bloque del nuevo criterio. El mismo comportamiento se usa en
tabla y tarjetas, en PC, tablet y móvil.

Al abrir un ticket y volver con el historial del navegador durante la misma sesión de navegación,
la bandeja recupera sus criterios, los bloques acumulados y la posición vertical. Esa memoria es
efímera: recargar o cerrar la pestaña vuelve a empezar por el primer bloque. Ir a otra bandeja no
mezcla sus resultados con la actual.

La etiqueta inglesa del estado `en espera` se presenta según el papel conforme a la sección 40. En
la tabla, **Action/Acción** se alinea a la izquierda como los demás encabezados; el botón Open/Abrir
continúa centrado en su celda.

Las pruebas cubren anexado sin duplicados, exclusión de solicitudes concurrentes, fin silencioso,
fallo y reintento, reinicio al cambiar criterios, anuncio accesible, tabla y tarjetas, restauración
al volver, descarte tras recarga, ambos idiomas, el texto `Waiting` y las dos alineaciones distintas
de encabezado y botón. La comprobación de navegador recorre PC, tablet y móvil.

## 39. Tabla de altura disponible y encabezado fijo

Desde 768 px, la cabecera de la bandeja —título, búsqueda, resumen de filtros y **New ticket/Nuevo
ticket**— permanece fuera del desplazamiento. La tabla ocupa todo el alto restante de la vista, con
un mínimo de 240 px, y su propio contenedor administra los desplazamientos vertical y horizontal.
Después de la última fila conserva 16 px de espacio interior para que el contenido no termine pegado
al borde. Por debajo de 768 px las tarjetas mantienen el desplazamiento normal de la página.

La fila de títulos permanece fija en el borde superior del contenedor mientras se recorren los
tickets. Conserva su fondo, tipografía y borde inferior actuales, con superficie opaca y el nivel de
apilamiento necesario; no añade sombra. La celda de encabezado **Ticket** queda fija en ambos ejes y
se dibuja por encima tanto de las demás cabeceras como de las celdas desplazadas.

El separador derecho de la columna Ticket deja de depender del borde de la celda fija: una línea de
1 px, independiente y superpuesta en su extremo derecho, conserva el color del tema y permanece
visible durante todo el desplazamiento horizontal. No se aumenta su grosor ni se introduce sombra.

Subject, Summary, Last action y el contador `+N` conservan una única ayuda propia, accesible con
cursor y teclado. Se elimina el atributo `title` que duplicaba esa ayuda. El tooltip se muestra sólo
cuando el contenido está realmente truncado; al enfocarlo con teclado permanece visible mientras
conserva el foco y se cierra al moverlo. El texto completo sigue disponible sin depender del ratón.

La carga progresiva observa el final del área desplazable interna y solicita el siguiente bloque al
acercarse, con la misma anticipación, exclusión de solicitudes y recuperación ante errores de la
sección 38. La barra horizontal pertenece al contenedor y permanece en su borde inferior, sin exigir
llegar a la última fila para desplazar las columnas.

Las pruebas cubren cálculo del alto restante y mínimo de 240 px, 16 px inferiores, tabla frente a
tarjetas, scroll vertical y horizontal combinados, encabezados fijos, esquina Ticket en ambos ejes,
separador persistente, ausencia de `title`, tooltip sólo con desbordamiento, foco y cierre, y carga
progresiva dentro del nuevo contenedor. La comprobación real recorre PC, tableta y móvil, y confirma
los ocho temas para el fondo, los bordes y el separador.

## 40. Estado `en espera` según quien lo consulta

El valor técnico continúa siendo `en espera`. Su rótulo visible depende del papel de la persona que
mira el ticket:

| Papel | Inglés | Español |
| --- | --- | --- |
| Usuario | `Waiting for you` | `Esperándote` |
| Soporte | `Waiting` | `Esperando` |
| Desarrollo | `Waiting` | `Esperando` |
| Administrador | `Waiting` | `Esperando` |

La distinción se aplica de forma coherente en toda la interfaz: tabla de bandeja, tarjetas móviles,
cabecera y ficha del ticket, historial y filtros que se muestran a cada papel. El Usuario sólo ve el
ticket principal; no se introduce una variante para el interno.

Los selectores con los que Soporte o Desarrollo cambian el estado conservan el rótulo operativo
`Waiting` / `Esperando`. No intentan nombrar allí a la persona esperada: el contexto y el texto que
se solicita al ejecutar la transición ya explican si se espera al Usuario o a Soporte.

Quedan fuera de alcance el valor técnico `en espera`, los contratos de API, la base de datos, las
reglas de transición, los permisos y los textos de correo. No se añade una preferencia personal ni
se persiste una nueva traducción.

Las pruebas deben cubrir los cuatro papeles en ambos idiomas y comprobar la etiqueta en bandeja,
tarjeta, ficha, historial y filtros. También deben confirmar que el selector técnico usa
`Waiting` / `Esperando` y que el backend continúa recibiendo `en espera`.

### 40.1 Registro del repaso

El responsable decidió el 2026-10-16:

1. aplicar la diferencia en toda la interfaz;
2. tratar al Administrador como papel técnico;
3. conservar el rótulo operativo en los selectores de estado;
4. limitar el cambio a la interfaz, sin modificar API, base de datos ni correos.

El repaso queda cerrado sin decisiones abiertas. El responsable aprobó explícitamente la propuesta
el 2026-10-16. La implementación respeta el alcance aprobado; pasaron 76 pruebas focalizadas, las
fronteras del frontend y la compilación.

## 41. Menú lateral más estrecho y marca vertical

El menú lateral conserva sus tres zonas y sus mismos permisos, rutas y controles, pero aprovecha
mejor el ancho de la vista. En escritorio y en el cajón móvil mide **208 px abierto**. En escritorio
plegado mide **56 px**; el cajón móvil no ofrece el estado plegado. Los enlaces conservan su área
táctil, sus iconos, el estado activo y el foco visible.

La cabecera abierta deja de ser una fila. El logo ocupa un marco de **80 × 80 px**, centrado, y el
nombre completo de la instalación aparece debajo, también centrado, en tantas líneas como necesite.
El nombre no ensancha el menú ni se trunca. El botón de plegar o cerrar ocupa la esquina superior
derecha en un espacio reservado: no empuja ni descentra el logo.

En el estado plegado, el único botón para desplegar se coloca arriba. El logo de **40 × 40 px** queda
centrado debajo, y un único separador cierra todo el bloque de marca. Se elimina el segundo botón de
desplegar que hoy se repite en la zona inferior. La marca no comparte fila con el control y, por
tanto, no se desplaza hacia la izquierda. El nombre queda oculto visualmente en este estado.

Los enlaces plegados mantienen el nombre mediante tooltip al pasar el puntero y mediante texto
accesible para tecnologías de asistencia. La reducción no elimina ni sustituye ninguna entrada.

La pantalla conjunta `/tickets/categories` y su entrada pasan a llamarse **Classification** en
inglés y **Clasificación** en español. Dentro conserva dos tarjetas claramente tituladas
**Categories / Categorías** y **Tags / Etiquetas**. No se separa en dos rutas ni en dos módulos, y no
cambian sus datos, API, permisos ni operaciones.

### 41.1 Hallazgo y alcance

El contrato anterior fija 240 px abierto y 64 px plegado, y la implementación vigente coloca logo,
nombre y control en una fila. En el estado plegado, los 64 px no alcanzan para logo y botón sin
desplazar la marca. Además, `Categories and tags` no cabe con comodidad en el ancho propuesto aunque
las dos materias ya forman una sola clasificación en listas y fichas.

Quedan fuera de alcance las demás etiquetas del menú, el orden de las entradas, la navegación, los
papeles, la persistencia del estado plegado, el comportamiento del cajón, el formato de los logos y
la administración de categorías o etiquetas.

### 41.2 Pruebas

Las pruebas deben cubrir 208 px abierto y 56 px plegado, logo de 80 y 40 px, nombre institucional
de varias líneas, un único botón de desplegar, ausencia de desplazamiento lateral de la marca,
tooltip y nombre accesible en los enlaces plegados, y la nueva denominación en ambos idiomas. La
comprobación en navegador recorrerá escritorio y móvil, menú abierto y plegado, nombre corto y largo,
los ocho temas y los papeles que ven Clasificación.

### 41.3 Registro del repaso

El responsable decidió el 2026-10-17:

1. conservar una sola pantalla y llamarla Classification / Clasificación;
2. usar 208 px abierta y 56 px plegada, también 208 px en el cajón móvil;
3. mostrar el logo a 80 × 80 px abierto y 40 × 40 px plegado;
4. mostrar el nombre completo, centrado y sin límite de líneas;
5. reservar la esquina superior derecha para plegar o cerrar sin mover la marca;
6. poner el único control de despliegue arriba del logo cuando el menú está plegado;
7. conservar un único separador debajo de todo el bloque de marca;
8. eliminar el control inferior duplicado y conservar tooltip y texto accesible en las entradas;
9. renombrar también el título de la pantalla y mantener dentro Categories y Tags.

El repaso queda cerrado sin decisiones abiertas. El responsable aprobó explícitamente la propuesta
el 2026-10-17. La implementación respeta el alcance aprobado; pasaron las 238 pruebas del frontend,
las fronteras y la compilación. La comprobación real cubrió escritorio abierto y plegado y cajón
móvil, y confirmó la marca centrada, el control único y Classification con Categories y Tags.

## 42. Altura disponible para las tablas operativas

Desde **768 px**, las bandejas de tickets y la lista de usuarios ocupan el alto disponible de la
ventana. El armazón deja **16 px arriba y 16 px abajo** en estas pantallas. La cabecera y la tabla se
distribuyen en una columna de alto limitado por la vista; la tabla crece con `flex` para llegar hasta
el margen inferior incluso cuando tiene pocos resultados.

La cabecera de cada listado conserva su jerarquía y contenido, pero reduce el espacio desperdiciado:
la separación entre la cabecera y la tabla es de **12 px**. En tickets, título, Search/Buscar,
resumen de filtros, limpieza y New ticket/Nuevo ticket comparten una sola fila de **42 px**, centrada
verticalmente cuando el ancho lo permite. Los botones de esa fila usan 42 px, igual que los
desplegables normalizados; el título conserva tamaño y peso.

En Usuarios se mantienen el título, su descripción y el panel completo de filtros. La descripción
continúa debajo del título con una separación menor, y sólo se reducen los márgenes externos del
panel: campos, conmutadores y blancos internos conservan sus medidas actuales. New account/Nueva
cuenta usa los mismos 42 px de la cabecera. La tabla pasa a administrar el alto restante como la de
tickets.

Los dos contenedores tienen un mínimo de **240 px** para seguir siendo utilizables en ventanas
bajas. Después de la última fila conservan **16 px de espacio interior**, aunque haya pocos
resultados. El borde exterior termina a 16 px del borde inferior de la ventana; no se ajusta a la
última fila ni cambia de altura con el número de resultados.

Por debajo de 768 px se conserva el flujo actual de tarjetas y de página. No se compactan controles
táctiles ni se fuerza una región desplazable de alto fijo.

### 42.1 Hallazgo y alcance

El armazón vigente aplica 32 px superiores y 40 px inferiores en escritorio. La bandeja de tickets
añade 24 px entre su cabecera y la tabla y calcula un alto de vista que descuenta 72 px. Usuarios no
calcula el alto restante: sus filtros y tabla crecen con el documento. Esto produce márgenes
desiguales y deja área vertical sin aprovechar.

El cambio se limita a las bandejas de tickets y a la lista de usuarios. Quedan fuera los catálogos,
el editor de correos, las fichas, formularios, configuración, tablas incrustadas, filas, columnas,
orden, filtros, carga progresiva, paginación y contratos de API.

### 42.2 Pruebas

Las pruebas deben comprobar desde 768 px: 16 px exteriores arriba y abajo, 12 px entre cabecera y
tabla, controles de cabecera de 42 px, tabla con mínimo de 240 px y 16 px interiores tras la última
fila. También deben cubrir tickets con y sin acción de alta, Usuarios con descripción y filtros,
pocos y muchos resultados, ventanas bajas y altas, y ausencia de cambio por debajo de 768 px. La
comprobación real recorrerá escritorio, tableta y móvil en ambos idiomas y con menú abierto y plegado.

### 42.3 Registro del repaso

El responsable decidió el 2026-10-17:

1. aplicar el patrón a las bandejas de tickets y a la tabla de usuarios;
2. dejar 16 px arriba y abajo y 12 px entre cabecera y tabla;
3. usar 42 px en los botones de las cabeceras, sin reducir el título;
4. conservar una sola fila compacta en tickets cuando el ancho lo permita;
5. mantener el contenedor hasta el borde inferior aunque haya pocos registros;
6. conservar en Usuarios la descripción y el panel de filtros, reduciendo sólo márgenes externos;
7. mantener un mínimo de tabla de 240 px y 16 px interiores después de la última fila;
8. aplicar el alto disponible desde 768 px y conservar el flujo móvil actual.

El repaso queda cerrado sin decisiones abiertas. El responsable aprobó explícitamente la propuesta
el 2026-10-17. La implementación respeta el alcance acordado: aplica el alto disponible sólo a las
bandejas de tickets y a Usuarios desde 768 px, compacta sus cabeceras y conserva el flujo móvil. Las
238 pruebas del frontend, las fronteras y la compilación pasaron; la comprobación real confirmó las
dos tablas en escritorio y las tarjetas de tickets en móvil.

## 43. Interfaz para Administrador acumulativo

Una cuenta combinada ve un único menú con la unión de las entradas de su papel operativo y las de
Administrador, sin duplicados ni selector de personalidad. Al entrar aterriza en la bandeja de su
papel operativo; la cuenta de fábrica conserva el inicio administrativo.

La insignia **Admin / Administrador** aparece con el papel en lista de usuarios, ficha, perfil propio
y bloque de cuenta al pie del menú. En la lista de usuarios, tanto la tabla de escritorio como las
tarjetas de pantallas estrechas presentan el papel operativo en la primera fila y la insignia en una
segunda fila, alineada a la izquierda y separada por un espacio pequeño y constante; nunca compiten
por el ancho de una misma línea. No se repite en tickets ni comentarios y no depende sólo del color.

Alta y ficha presentan selector de papel operativo y casilla independiente de Administrador. La
casilla sólo se habilita para Soporte o Desarrollo y sólo la modifica otro Administrador. Cambiar
entre ambos papeles no la desmarca.

Al retirar o desactivar al último Administrador ordinario, una ventana acorde al tema informa que la
cuenta de fábrica será la única vía administrativa y exige confirmación, sin impedir la acción.

La próxima carga de identidad muestra a la persona afectada un aviso bilingüe por concesión o
retirada. Puede cerrarse y, como excepción deliberada a los toast de diez segundos, no desaparece
automáticamente. La ficha muestra el historial con actor, acción y fecha; no nace una pantalla global.

Selector, casilla, insignia, confirmación, aviso e historial se verifican con teclado, lector de
pantalla, PC y móvil. Las 240 pruebas del frontend y sus fronteras pasaron el 2026-10-18.

## 44. Dashboard y reportería

Dashboard/Panel es la primera entrada del menú de Soporte, Desarrollo y Administrador y su destino
inicial. Usuario no la ve. Arriba presenta seis indicadores comparados con el periodo anterior; en
el centro, una línea de entradas/cierres y barras de estado, prioridad, categoría y responsable;
no presenta debajo tablas, tarjetas ni enlaces a tickets individuales.

Los filtros de periodo, tipo, estado y prioridad viven en la URL y se adaptan al ancho sin ocultar
su valor. Actualizar conserva el último
resultado visible. Cada gráfico tiene título, leyenda, valores accesibles y alternativa textual;
color nunca es la única diferencia. Las barras usan `bg-primario`, `text-primario`,
`text-primario-texto` y `bg-exito`, nunca los alias ingleses inexistentes. En móvil los bloques se
apilan. La serie es diaria hasta 31 días y semanal desde 32; sólo rotula fechas que caben y conserva
la fecha exacta de cada punto de forma accesible.

La única exportación visible se presenta como un botón con un icono de descarga a la izquierda de
«XLSX» y genera `report_AAAA-MM-DD_AAAA-MM-DD.xlsx`, localizado y dividido en seis hojas agregadas,
sin gráficas incrustadas ni datos individuales. La verificación recorrerá
escritorio y móvil, los ocho temas y ambos idiomas, y comprobará dimensiones y colores computados,
etiquetas temporales, ausencia de tickets y contenido del libro.

## 45. Estado activo único del menú lateral

El menú lateral muestra **una sola entrada activa**. Las rutas de sus listados coinciden de forma
exacta: `/tickets`, `/tickets/main`, `/tickets/internal` y `/tickets/categories` no se activan por
prefijo ni pueden iluminarse simultáneamente. La entrada activa conserva el fondo
`bg-superficie-suave`, el texto principal y el peso medio actuales; no añade sombra, borde ni una
segunda señal decorativa. El enlace expone además `aria-current="page"`.

Las pantallas que no tienen entrada propia conservan la pertenencia funcional acordada:

- `/tickets/new` y `/tickets/:numero` activan **Mis tickets / Bandeja**;
- `/users/:id` activa **Usuarios**;
- `/mail` activa **Configuración**;
- las demás rutas sin entrada propia no seleccionan ningún módulo.

La regla es idéntica con menú abierto, plegado o en cajón móvil y no depende del papel ni de que la
cuenta acumule Administración. Una función pura resuelve la pertenencia de la URL y sus pruebas
recorren las rutas visibles y auxiliares. La prueba E2E comprueba que existe exactamente un enlace
con `aria-current="page"` y que conserva el fondo suave. El recorrido manual con una cuenta
combinada confirmó todas las entradas y las cuatro clases de ruta auxiliar.

### 45.1 Registro de decisiones

El responsable eligió el 2026-10-19:

1. coincidencia exacta y una sola entrada activa (1A);
2. fondo suave actual, sin sombra (2A);
3. alta y detalle dentro de **Mis tickets / Bandeja** (3A);
4. ficha de usuario dentro de **Usuarios** y editor de correos dentro de **Configuración** (4A).

Quedan fuera del alcance el orden, las etiquetas, los iconos, los permisos, las rutas, los destinos
iniciales, la paleta de los temas y cualquier cambio de API o datos. No quedan decisiones abiertas;
el responsable aprobó explícitamente la propuesta y la implementación respeta el alcance acordado.

## 46. Cabecera y vistas de «Mis tickets»

La primera fila de **Mis tickets** contiene, junto al título y las acciones existentes, un selector
siempre visible con tres opciones: **All · Assigned · Watching** / **Todos · Asignados · Observo**.
La vista inicial es **All**, que combina tickets asignados y observados sin duplicados. Cambiar la
opción actualiza la lista y su dirección; no abre un modal ni se cuenta como filtro de búsqueda.

Soporte ve únicamente principales y Desarrollo únicamente internos. Una cuenta operativa con
capacidad Admin conserva exactamente esa presentación. Al dejar de observar un ticket no asignado,
éste desaparece al actualizar la bandeja; si estaba abierto en detalle, la ficha no se cierra ni se
interrumpe y al volver se refleja la lista vigente.

No se ofrece vista de no asignados ni de participación histórica. **Tickets principales** y
**Tickets internos** conservan sus filtros y su contenido general. Las pruebas cubren los tres
selectores, la combinación sin duplicados, ambos papeles operativos, una cuenta combinada y la
retirada de observador.

## 47. Comentarios públicos y privados

En el editor de comentarios del principal, Soporte ve un selector **Public · Private** / **Público ·
Privado**, con `Public` seleccionado al abrir. Desarrollo ve el mismo control limitado a `Private`,
pues no puede responder públicamente al Usuario. El selector no aparece en el interno ni al Usuario.
Los editores que acompañan movimientos de estado no lo muestran y publican siempre en público.

El equipo distingue una nota privada mediante el chip **Private / Privado** junto al autor y un
fondo sutil derivado del tema; no depende sólo del color. El Usuario no ve hueco, candado, contador,
entrada de historial ni señal de que existe. Los adjuntos privados se renderizan únicamente dentro
de esa nota para la audiencia autorizada.

El autor puede editar el comentario y alternar su visibilidad en ambos sentidos desde la acción de
edición existente. La pantalla actualiza inmediatamente conversación, adjuntos, búsqueda y
resúmenes; no abre confirmación al volver privado un comentario público. El historial técnico muestra
quién cambió **Public → Private** o **Private → Public** y cuándo, sin repetir el cuerpo.

Al escalar, la nota no se duplica en el hilo interno. En vista doble permanece en la columna del
principal; si se mira sólo el interno se consulta cambiando de vista. Las pruebas recorren los ocho
temas, ambos idiomas, PC y móvil, teclado y lector de pantalla, además de comprobar que el DOM del
Usuario no contiene el comentario ni sus adjuntos.

## 48. Buscar y crear una cuenta de Active Directory

El diálogo de alta de `/users` conserva **Cuentas de la aplicación** para el alta local. Cuando AD
sea el método vigente añade **Directorio de la organización**. Keycloak aparece deshabilitado con
la explicación «La cuenta se crea en su primer acceso»; si el método vigente no es AD, su búsqueda
no se ofrece.

La vista de directorio empieza con un solo campo **Usuario o correo**. Antes de tres caracteres no
consulta y muestra una ayuda breve. Después espera la pausa de escritura, cancela la petición
anterior y presenta hasta veinte resultados con nombre completo, usuario y correo. Tiene estados
distintos para buscando, sin coincidencias y directorio no disponible; teclado y lector de pantalla
pueden recorrer y seleccionar cada resultado.

Después de seleccionar se elige el papel con los límites actuales: Soporte sólo Usuario y
Administrador los papeles/capacidad que ya administra. No se editan nombre, apellidos ni correo,
porque vienen de AD. Crear vuelve a la lista, conserva filtros y muestra el resultado en el toast
común. No manda correo de alta.

En la interfaz en español, los rótulos visibles de cuenta dicen **Rol** —en el alta, los filtros,
las columnas y el perfil— por ser el término más directo para quien administra usuarios. El modelo
interno y los permisos no cambian.

Si el correo ya corresponde a una cuenta local, el diálogo temático explica antes de confirmar que
se conservará esa cuenta, pasará a AD y su contraseña local dejará de funcionar. Cancelar no cambia
nada. La confirmación no se reutiliza para otra persona ni sobrevive al cierre del diálogo.

Las pruebas automatizadas verifican las opciones según el método, que Keycloak no parezca
consultable y el contrato de búsqueda y alta. La suite completa del frontend (249 casos), sus
fronteras y la compilación de producción pasaron el 2026-10-20.

## 49. Navegación unificada y módulos administrativos

El menú de Soporte y Desarrollo sustituye **Tickets principales** y **Tickets internos** por una
sola entrada **Tickets**. **Mis tickets** permanece independiente. Al entrar, la pantalla abre
**Principales** y presenta un selector segmentado **Principales / Internos**; en móvil ese control se
convierte en un desplegable para no comprimir las dos etiquetas.

El selector escribe `type=main|internal` en la URL y responde al historial del navegador. Cada lado
conserva sus propios filtros mientras la pantalla sigue montada. No muestra contadores. Las antiguas
rutas desaparecen sin redirección, y el menú siempre enlaza a `/tickets?type=main`.

En la vista principal, Desarrollo ve un editor rotulado **Comentario interno**. Todo lo que publique
allí es privado automáticamente: no aparece selector de visibilidad ni un control bloqueado. Soporte
y Desarrollo pueden leerlo; el Usuario no puede inferirlo. Soporte conserva en sus propios
comentarios el selector público/privado vigente.

**Clasificación** y **Usuarios** sólo aparecen cuando la identidad tiene capacidad Administrador.
Una cuenta combinada ordena su menú como Panel, Mis tickets, Tickets, Clasificación, Usuarios y
Configuración. El Administrador puro conserva su navegación administrativa y no recibe Tickets.

La terminología visible del área se repasa completa en español e inglés: nombres de módulos,
encabezados, filtros, botones, ayudas y estados. Se conservan los conceptos **Mis tickets**,
**ticket principal** y **ticket interno**; el repaso elimina mezclas y nombres antiguos, no cambia el
modelo. Las 249 pruebas del frontend, sus fronteras y la compilación de producción pasaron el
2026-10-20; la navegación cubre roles puros y combinados y ausencia de entradas no autorizadas.

## 50. Icono del sitio y favicon

El bloque de marca de `/setup` y Configuración añade una tarjeta cuadrada **Icono del sitio**, aparte
de los logos claro y oscuro. La vista previa encaja la imagen completa y centrada, sin recortarla ni
deformarla. Pulsarla abre el selector de archivo; la ayuda admite PNG, JPEG, WebP y SVG, hasta 1 MB y
2000 píxeles por lado, con un mínimo de 32 × 32. La acción **Restaurar** sólo aparece cuando existe un
icono propio y explica que se volverá al logo claro.

El favicon de la pestaña se actualiza al terminar una subida o restauración, sin recargar. Si no hay
icono propio, también se actualiza cuando cambia o se restaura el logo claro. La misma actualización
se propaga a las demás pestañas abiertas de Catalina Support; una pestaña que se abra después obtiene
el valor efectivo al leer la marca pública. El enlace lleva el sello de versión para no reutilizar
una imagen anterior de la caché.

El icono es decorativo porque el nombre de la pestaña ya identifica la aplicación. La carga conserva
el foco, muestra el resultado en el toast común y no bloquea otras partes de Configuración. Se prueba
en los ocho temas, ambos idiomas, teclado, PC y móvil, incluidos los tres niveles de fallback y la
sincronización entre pestañas. No se cambia el icono de accesos directos instalados.

## 51. Ancho útil de las bandejas de tickets

Las cuatro bandejas de tickets ocupan el ancho disponible dentro del armazón y dejan 16 px a cada
lado en escritorio. Esta excepción sustituye para esas pantallas los 40 px generales y su límite de
88 rem; no ensancha formularios ni páginas de lectura. El menú lateral conserva sus anchos, control
de plegado y estado actuales.

La tabla reparte el espacio liberado entre sus columnas descriptivas. Ticket queda fija a la
izquierda y Action a la derecha cuando existe desplazamiento. La primera columna suma aire junto al
borde izquierdo y la última junto al derecho para que enlaces y botones no parezcan pegados al
marco. Separadores, fondos, foco y contraste usan los tokens de cada tema. En móvil permanecen las
tarjetas.

El navegador se comprueba con menú abierto y plegado, 1440 y 1920 px, un ancho intermedio con
scroll, zoom de usuario y móvil. La revisión visual exige que el botón final sea visible o quede
fijo durante el desplazamiento, que no cubra texto y que Summary y Last action mantengan su
legibilidad.

### 51.1 Registro de decisiones

El responsable aclaró el 2026-10-22 que el espacio desperdiciado está a izquierda y derecha de la
tabla, no en el menú. Eligió las cuatro bandejas (1A), 16 px (2A), Action fija (3A) y añadió padding
interior en ambos extremos. No quedan decisiones abiertas.

### 51.2 Implementación y verificación

La implementación deja 16 px efectivos mediante compensación localizada en la bandeja, sin cambiar
el armazón ni otras páginas. Action recibió 6 rem después de que la primera revisión visual detectó
que el ancho mínimo anterior recortaba encabezado y botón. La segunda pasada de Chromium confirmó
la tabla completa y sin desbordamiento en 1920 y 1440 px, y las tarjetas a 390 px. Pasaron las 252
pruebas Angular, la regresión focal, las fronteras y la compilación hacia un directorio limpio.
