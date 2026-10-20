# Documentación de Catalina-Support

> **Estado:** as-built
> **Última actualización:** 2026-10-20
>
> **Cambio coordinado aprobado, implementado y verificado el 2026-10-20:** `ambientes.md`,
> `arquitectura.md` y `prueba-local.md` sustituyen la migración manual de producción por un ejecutor
> Go automático, versionado, transaccional, con historial, checksum, bloqueo y respaldo previo. El
> responsable cerró y aprobó las decisiones 1A–25A. No se desplegó producción.
>
> **Cambio coordinado aprobado, implementado y verificado el 2026-10-20:** `modules/tickets.md`,
> `usuarios-y-permisos.md`, `modules/users.md` e `interfaz-y-experiencia.md` unifican las listas
> generales en **Tickets**, reservan Usuarios y Clasificación a la capacidad Admin y fijan los
> comentarios privados de Desarrollo en principales. Las decisiones 1A–25A quedaron cerradas con
> 6B, 12B, 14B, 20B y 25B. Pasaron las suites focales del backend, las 249 pruebas del frontend,
> las fronteras y la compilación; la suite Go global conserva el fallo ambiental conocido de
> reportes porque la base local no contiene sus tickets sembrados.
>
> **Cambio coordinado aprobado, implementado y verificado el 2026-10-20:** `modules/auth.md`,
> `modules/users.md`, `modules/settings.md` e `interfaz-y-experiencia.md` corrigen `objectGUID`
> binario, admiten usuario o correo al entrar por AD y permiten buscar y crear desde `/users` una
> cuenta AD activa. La búsqueda es parcial desde tres caracteres, devuelve hasta veinte personas y
> conserva el filtro base que excluye cuentas deshabilitadas. Keycloak mantiene el alta automática
> hasta disponer de una cuenta de servicio de consulta. El responsable cerró y aprobó 1A–15A, con
> 6B. Se verificó el acceso real por usuario y correo, pasaron las suites focales de backend y las
> pruebas y compilación del frontend; la suite Go global conserva un fallo ambiental de reportes
> porque la base local no contiene sus tickets sembrados.
>
> **Cambio coordinado aprobado, implementado y verificado el 2026-10-20:** `modules/mail.md`, `modules/settings.md`,
> `primer-arranque.md`, `arquitectura.md` y `prueba-local.md` convierten la prueba SMTP en un envío
> real opcional a una dirección efímera, con `From:` real, límite público, texto fijo y confirmación
> de aceptación sin prometer entrega. El responsable cerró y aprobó 1A–14A; pasaron las pruebas
> focalizadas, las 248 del frontend, fronteras y compilación.
>
> **Cambio coordinado aprobado, implementado y verificado el 2026-10-19:** `modules/mail.md`,
> `modules/settings.md`, `primer-arranque.md`, `arquitectura.md` y `prueba-local.md` distinguen
> STARTTLS, TLS directo y SMTP sin cifrado. El último exige aceptación administrativa, sólo admite
> IPv4 privada o loopback y no aparece por degradación automática; Mailpit es la única excepción
> explícita de desarrollo. El responsable cerró y aprobó 1A–10A; el relay real autenticó sin enviar
> correo y pasaron las pruebas focalizadas, las 247 del frontend, fronteras y compilación.
>
> **Cambio coordinado aprobado, implementado y verificado el 2026-10-19:** `modules/tickets.md`, `usuarios-y-permisos.md`,
> `interfaz-y-experiencia.md`, `modules/ai.md` y `modules/mail.md` incorporan comentarios privados
> en el principal, adjuntos protegidos, auditoría, correo interno y resúmenes separados por audiencia.
> La regla anterior que remitía todo contenido interno al ticket escalado queda identificada como
> hallazgo. El responsable cerró y aprobó 1A–26A; pasaron backend, 247 pruebas Angular, compilación
> de producción y el recorrido real Soporte/Usuario.
>
> **Cambio coordinado aprobado, implementado y verificado el 2026-10-19:** `modules/tickets.md`, `usuarios-y-permisos.md` e
> `interfaz-y-experiencia.md` redefinen **Mis tickets** como asignados u observados actuales,
> separados por nivel, con `All · Assigned · Watching` en la primera fila. Excluyen apertura,
> comentarios y no asignados; retirar la observación elimina la pertenencia salvo asignación. La
> capacidad Admin conserva el papel operativo. El responsable aprobó el documento; pasaron las
> pruebas de frontend y backend y se verificó el flujo con una cuenta Desarrollo + Admin.
>
> **Corrección coordinada aprobada, implementada y verificada el 2026-10-19:** `interfaz-y-experiencia.md` corrige la selección doble
> del menú lateral mediante una única entrada activa, coincidencias exactas y pertenencia explícita
> para las pantallas auxiliares. El responsable eligió 1A–4A. No cambia rutas, permisos, etiquetas,
> iconos ni temas. Pasaron la compilación y las 246 pruebas del frontend; el navegador confirmó
> selección única en todas las entradas y la pertenencia de las rutas auxiliares.
>
> **Corrección coordinada aprobada, implementada y verificada el 2026-10-19:** `modules/tickets.md` corrige la
> contaminación de cláusulas GORM al cambiar prioridad desde un interno. La reparación conserva la
> transacción y el comportamiento ya aprobado; no cambia permisos, API, interfaz ni datos. Pasaron
> la regresión aislada y toda la suite backend.
>
> **Corrección coordinada aprobada, implementada y verificada el 2026-10-19:** `propósito-y-alcance.md`,
> `interfaz-y-experiencia.md` y `modules/dashboard.md` convierten Dashboard en reportería sólo
> estadística, corrigen barras transparentes en los ocho temas, agrupan la tendencia sin solapar
> fechas, reducen los filtros y reemplazan CSV por un XLSX localizado de seis hojas. La revisión real
> encontró además que la pantalla actual lista tickets contra el nuevo alcance. El responsable cerró
> 1A–4A, propuso XLSX en 5, eligió 6A–14A, corrigió 15A y eligió 16A. Pasaron las pruebas específicas
> de Dashboard y reportería en backend, 241 pruebas
> frontend, fronteras, compilación, los ocho temas y la validación estructural del XLSX.
>
> **Propuesta coordinada del 2026-10-18:** `modules/ai.md`, `modules/settings.md`,
> `primer-arranque.md`, `arquitectura.md`, `ambientes.md` y `prueba-local.md` añaden aceleración
> local Auto/CPU/NVIDIA, una variante Compose opcional y cambio transaccional con respaldo seguro.
> El responsable cerró las decisiones 1A–25A; no quedan decisiones abiertas.
>
> **Dashboard implementado y verificado el 2026-10-18:** nace `modules/dashboard.md` y se actualizan
> alcance, permisos, tickets, settings, arquitectura, interfaz y prueba local. Incluye reportería
> operativa, filtros, tendencias, tiempos, umbrales y CSV. El responsable cerró 1A–45A, con 24A
> ampliada a los dos formatos de duración y 26A excluyendo al Usuario. Pasaron PostgreSQL real,
> pruebas de servicio, 241 pruebas frontend, fronteras, compilación y recorrido visual PC/móvil.
>
> **Cambio de interfaz del 2026-10-17, aprobado, implementado y verificado:**
> `interfaz-y-experiencia.md` amplía verticalmente las tablas operativas de tickets y usuarios,
> compacta sus cabeceras y conserva 16 px de aire superior, inferior e interior. El repaso 1A–15A
> quedó cerrado. Pasaron las 238 pruebas del frontend, las fronteras y la compilación; el navegador
> confirmó ambas tablas en escritorio y que las tarjetas de tickets permanecen intactas en móvil.
>
> **Cambio de interfaz del 2026-10-17, aprobado, implementado y verificado:**
> `interfaz-y-experiencia.md` reduce el menú lateral a 208/56 px, coloca el logo y el nombre en una
> cabecera vertical centrada, separa el único control de plegado por encima del logo cuando está
> plegado y renombra la pantalla conjunta como Classification/Clasificación. El repaso 1A–14A quedó
> cerrado. Pasaron las 238 pruebas del frontend, las fronteras y la compilación; el navegador
> confirmó escritorio abierto y plegado, cajón móvil, la marca centrada y Classification.
>
> **Portada pública cerrada el 2026-10-17:** `proyecto-abierto.md`, `ambientes.md` y
> `prueba-local.md` documentan el README bilingüe, el sitio oficial, la demo compartida, Ko-fi y
> dos capturas reales complementarias de la vista doble y la IA. Los enlaces relativos, la marca
> local restaurada y los Compose de desarrollo e IA quedaron verificados; no se tocó producción ni
> ningún entorno público.
>
> **Corrección coordinada del 2026-10-16, aprobada, implementada y verificada:**
> `modules/tickets.md` devuelve a Soporte la transición visible Esperando → En progreso que el
> backend ya admite. `modules/ai.md` impide que una generación antigua sobrescriba a la vigente,
> añade una recuperación automática diferida y finita y recupera una vez los errores existentes.
> Parte de tres fallos reproducidos durante el recorrido real de `CS-2026-0028` a `CS-2026-0030`.
> El responsable eligió 1A–4A. Pasaron 238 pruebas frontend, compilación, fronteras y las pruebas de
> IA contra repositorio real; el navegador ejecutó el retorno en `CS-2026-0011`. La suite backend
> conserva el fallo previo y ajeno de prioridad desde un interno. No vuelven las acciones manuales
> de IA.
>
> **Cambio del 2026-10-16, aprobado, implementado y verificado:** `interfaz-y-experiencia.md` distingue el
> estado `en espera` por papel en toda la UI: `Waiting for you` / `Esperándote` para el Usuario y
> `Waiting` / `Esperando` para Soporte, Desarrollo y Administrador. No cambia API, base de datos ni
> correos. Pasaron 76 pruebas focalizadas, las fronteras y la compilación del frontend.
>
> **Propuesta coordinada del 2026-10-15, pendiente de aprobación documental:**
> `interfaz-y-experiencia.md` y `modules/tickets.md` hacen visible **Action/Acción** sobre la columna
> de **Open/Abrir**, centrada y con ancho mínimo. La UI confirmó que la celda estaba dentro de la
> tabla, pero su encabezado `sr-only` provocaba la separación visual. El repaso posterior define
> orden, jerarquía de anchos, truncado, clasificación limitada, variantes intermedia y móvil y
> pruebas en ambos idiomas. El repaso final añade ayudas accesibles con teclado, etiquetas ocultas
> enumeradas, respuesta al ancho real del contenedor y conservación de la ordenación existente. No
> quedan decisiones abiertas y todavía no se modifica código.
>
> **Cambio coordinado del 2026-10-15, aprobado, implementado y verificado:** `modules/ai.md`,
> `modules/tickets.md`, `interfaz-y-experiencia.md` y `arquitectura.md` convierten Motivo/Reason en
> Summary/Resumen estable, retiran botones y endpoint de reintento, y migran el contrato afectado a
> `summary` y `escalationReason`. El responsable eligió 1A–20B, aceptó la regeneración total incluso
> con proveedor y no dejó decisiones abiertas. La migración se aplicó sólo en desarrollo; pasaron
> las pruebas específicas, las 231 pruebas del frontend, sus fronteras y la compilación. La suite
> completa del backend conserva el fallo previo de prioridad desde un interno, ajeno a este cambio.
>
> **Corrección coordinada del 2026-10-15, aprobada, implementada y verificada:** `modules/ai.md`,
> `modules/tickets.md` y `arquitectura.md` corrigen la última acción que copiaba la razón, entregan
> contexto funcional completo de principal e interno, vinculan evento y comentario, validan actor y
> acción y regeneran en segundo plano las filas existentes. El responsable eligió 1A–14A y aprobó
> explícitamente los documentos; no quedan decisiones abiertas.
>
> **Corrección coordinada del 2026-10-15, aprobada, implementada y verificada:**
> `interfaz-y-experiencia.md` fija todos los desplegables en 42 px para igualarlos al conmutador
> Main/Both/Internal. `ambientes.md` confirma que `ng serve` ya reconstruye y recarga con respuestas
> `no-cache`, por lo que no se añaden banderas redundantes ni polling. El responsable eligió 1A–5A.
> Pasaron las 232 pruebas del frontend, la compilación y la medición real de 42 px; no quedan
> decisiones abiertas.
>
> **Ajuste del 2026-10-15, aprobado, implementado y verificado:**
> `interfaz-y-experiencia.md` registra que los 52 px se sirvieron correctamente pero resultaron
> demasiado grandes. Todos los desplegables usan ahora 48 px, conservando tipografía y padding.
> Pasaron las 232 pruebas del frontend, la compilación y la medición en una pestaña nueva; no quedan
> decisiones abiertas.
>
> **Corrección del 2026-10-15, aprobada, implementada y verificada:**
> `interfaz-y-experiencia.md` registra que los 44 px sí se aplicaron, pero todavía se percibían
> demasiado delgados. Todos los desplegables usan ahora 52 px sin agrandar la tipografía. Pasaron
> las 232 pruebas del frontend, la compilación y la medición real; no fue necesario reiniciar los
> contenedores y no quedan decisiones abiertas.
>
> **Implementación del 2026-10-15, aprobada y verificada:** `interfaz-y-experiencia.md` alinea todos
> los desplegables con la altura mínima de 44 px y el espaciado de los campos de texto, sin variantes
> compactas y en todos los tamaños de pantalla. Pasaron las 232 pruebas del frontend, la compilación
> y la revisión real en PC y móvil; no quedan decisiones abiertas.
>
> **Implementación coordinada cerrada el 2026-10-15, con aprobación explícita:** `modules/ai.md`,
> `modules/tickets.md` e `interfaz-y-experiencia.md` rechazan resúmenes sin significado, conservan
> el texto anterior ante fallos, miden paráfrasis reales y evitan que el refresco de la IA destruya
> borradores o propuestas locales. Parte de los hallazgos reproducidos en `CS-2026-0026`; el repaso
> 1A–15A quedó cerrado, con 8B. La corrección posterior del responsable retiró la ayuda permanente
> del editor de comentarios y el estado neutro de propuesta lista. Pasaron 231 pruebas del frontend,
> la compilación, las pruebas afectadas del backend y el recorrido real `CS-2026-0027`; no quedan
> decisiones abiertas.
>
> **Implementación coordinada del 2026-10-14, aprobada y verificada:** `modules/ai.md`,
> `modules/tickets.md` e `interfaz-y-experiencia.md` eliminan el modal y las cinco tonalidades de
> «Mejorar con IA». La generación será profesional e inmediata, con propuesta inline por editor,
> aplicación o descarte y paráfrasis distinta al regenerar. Pasaron las 230 pruebas del frontend,
> la compilación de producción, las pruebas afectadas del backend y el recorrido real sin publicar;
> el único fallo de la suite completa del backend es el caso ajeno de prioridad registrado en esos
> documentos. No quedan decisiones abiertas.
>
> **Implementación coordinada del 2026-10-14, aprobada y verificada:** `modules/tickets.md` e
> `interfaz-y-experiencia.md` retiran la columna interna que permanece vacía en **Los dos** y hacen
> que cada conversación ocupe por completo su mitad, conservando 24 px entre paneles y sin tocar las
> vistas individuales. La comparación aparecerá desde 1152 px. El responsable eligió 1A–6A, 7B y
> 8B. Pasaron las 229 pruebas del frontend, la compilación y el recorrido real en 1800, 1152 y 1151
> px; no quedan decisiones abiertas.
>
> **Implementación coordinada del 2026-10-14, aprobada y verificada:** `modules/tickets.md` e
> `interfaz-y-experiencia.md` convierten la vista simultánea en dos paneles 50/50 con scroll propio,
> cabecera común visible y ancho útil completo; móvil y tableta vertical mantienen un ticket por vez.
> El responsable eligió 1A–4A. Pasaron las 228 pruebas del frontend, la compilación y el recorrido
> real en PC y tableta vertical; no quedan decisiones abiertas.
>
> **Implementación coordinada del 2026-10-14, aprobada y verificada:** `primer-arranque.md`,
> `modules/settings.md` e `interfaz-y-experiencia.md` muestran automáticamente el catálogo local y
> reservan el botón **Refrescar catálogo** para actualizar sin ocultar las fichas. Pasaron las 228
> pruebas del frontend, la compilación y el recorrido real de `/setup`.
>
> **Implementación coordinada del 2026-10-14, aprobada y verificada:** `primer-arranque.md`,
> `modules/settings.md` e `interfaz-y-experiencia.md` sustituyen todas las confirmaciones nativas por
> el diálogo temático y añaden al paso 1 los ocho temas y el color institucional con vista previa.
> Pasaron backend, 227 pruebas del frontend, compilación y recorrido real en PC y móvil.
>
> **Implementación coordinada del 2026-10-14, aprobada y verificada:** `primer-arranque.md`,
> `modules/ai.md`, `modules/settings.md` e `interfaz-y-experiencia.md` definen logos pulsables sin
> nombre de archivo, control de activación y desactivación desde cada ficha, saludo atribuido por
> toast y una política común de cierre manual con duración máxima de diez segundos. El responsable
> cerró 1A, 2A y corrigió 3A al límite común de diez segundos; no quedan decisiones abiertas.
>
> **Implementación coordinada del 2026-10-14, aprobada y verificada:** `primer-arranque.md`, `modules/ai.md`,
> `modules/settings.md` e `interfaz-y-experiencia.md` vuelven a `as-built` tras mostrar y controlar la
> descarga de modelos, activar y probar el elegido, administrar modelos inactivos y cargar los dos
> logos desde el paso 1. El hallazgo —la interfaz anterior no mostraba el progreso ni impedía
> avanzar durante la descarga— quedó corregido tras la aprobación explícita. Pasaron el administrador
> aislado, todo el backend, las 228 pruebas del frontend y el recorrido real del modelo 3B.
>
> **Prioridades implementadas para 1.0.0, 2026-10-14:** nueve documentos vuelven a `as-built` tras
> incorporar baja, media, alta y crítica, con permisos, historial, orden, filtro, correo único y
> pruebas. El responsable cerró 1A–25A con correcciones propias a 3A, 13A y 18A y aprobó
> explícitamente la propuesta coordinada.
>
> **Bloque 8 cerrado, 2026-10-13:** `prueba-local.md` y `ambientes.md` registran la cuarta pasada
> íntegra desde el repositorio público. Instalación vacía, IA, correo, seeders, AD, Keycloak y vuelta
> a local funcionaron; los logs terminaron limpios y el entorno desechable fue retirado.
>
> **Hallazgo corregido y verificado del tercer recorrido público, 2026-10-13:** `arquitectura.md`,
> `modules/settings.md`, `modules/auth.md`, `modules/mail.md`, `modules/users.md` y
> `modules/tickets.md` retiran los falsos errores de GORM para ausencias
> esperadas. `ambientes.md` y `prueba-local.md` conservan la cuarta repetición íntegra aprobada. El
> responsable aprobó 1A–4A; backend y `go vet` pasaron contra PostgreSQL real.
>
> **Propuesta coordinada del 2026-10-10:** `proyecto-abierto.md`, `arquitectura.md`,
> `modules/settings.md` e `interfaz-y-experiencia.md` preparan la publicación bajo AGPL-3.0-only,
> CLA sin cesión, política de marca, perfil comunitario, auditoría y enlace a la fuente. El
> responsable cerró 1A–25A. El repositorio sigue privado y no se modifica licencia, historial,
> configuración de GitHub ni código hasta aprobar los documentos.
>
> **Primer hallazgo cerrado el 2026-10-10:** la auditoría encontró dos vulnerabilidades críticas y
> cuatro altas en el lockfile de Angular. El responsable aprobó 26A; Angular quedó en 22.2.2, la
> auditoría terminó limpia y pasaron las pruebas, la compilación y el recorrido afectado.
>
> **Segundo hallazgo cerrado el 2026-10-10:** `go-logs v1.0.0`, dependencia directa del backend, no
> publicaba licencia. El responsable aprobó 27A y autorizó el cambio concreto: `go-logs v1.0.1` se
> publicó bajo MIT y Catalina Support ya resuelve y verificó ese artefacto licenciado.
>
> **Tercer hallazgo cerrado el 2026-10-10:** `proyecto-abierto.md` incorporó un inventario
> versionado de 29 dependencias de producción, sus textos completos, generación reproducible,
> comprobación de CI y copia junto a los artefactos, conforme a 28A–37A.
>
> **Cuarta enmienda cerrada el 2026-10-10:** la prueba real del CLA confirmó el flujo completo y
> reveló los identificadores públicos adicionales del esquema fijo de la acción. El responsable
> eligió y aprobó 38A; `proyecto-abierto.md` quedó `as-built` con el formato y la evidencia.
>
> **Metadatos y enlaces cerrados el 2026-10-10:** `arquitectura.md`, `modules/settings.md` e
> `interfaz-y-experiencia.md` vuelven a `as-built`. Backend, 227 pruebas del frontend y 12 casos
> aislados en PC y móvil pasaron; 39A alinea el menú plegado con su comportamiento real. El entorno
> desechable quedó limpio.
>
> **Configuración de GitHub aprobada el 2026-10-11:** `proyecto-abierto.md` concreta presentación,
> canales comunitarios, squash y protección de `dev` y `main` conforme a 40A–54A. La visibilidad no
> forma parte de esta enmienda; la aplicación y verificación de los ajustes sigue pendiente.
>
> **Corrección de publicación aprobada el 2026-10-11:** GitHub ya es público y la rama
> predeterminada `main`, atrasada, expone la licencia MIT anterior. `proyecto-abierto.md` registra
> 56A–65A: portada temporal en `dev`, segundo bundle, reescritura acotada de las tres ramas,
> conservación del CLA y sincronización posterior mediante PR. La implementación sigue pendiente.
>
> **Secuencia corregida y aprobada el 2026-10-11:** 66A mantiene `main` como rama principal
> y `dev` como desarrollo. La reescritura publicará primero `main` para corregir MIT y después las
> otras ramas y el PR. La implementación sigue pendiente.
>
> **Bloque 7 implementado y verificado el 2026-10-08.** `arquitectura.md` cierra la coherencia
> documental, automatiza las fronteras del frontend y deja decidido que no habrá un sistema de
> logging en el navegador. Las suites de frontend, build y backend pasaron; producción no se tocó.
>
> **Actualizado el 2026-10-08 (sexta vez):** la visibilidad de «Mejorar con IA» quedó implementada.
> El detalle devuelve `capabilities.aiWriting`; la interfaz sólo monta el botón y el modal cuando es
> verdadera. Configuración se distingue de salud y la IA continúa obligatoria. Pasaron backend, 225
> pruebas del frontend y Playwright en PC y móvil; los tres documentos vuelven a **as-built**, el
> entorno desechable quedó limpio y producción no se tocó.
>
> **Actualizado el 2026-10-08 (quinta vez):** `modules/ai.md`, `modules/tickets.md` e
> `interfaz-y-experiencia.md` pasan a **propuesta** para ocultar «Mejorar con IA» cuando no existe una
> configuración activa. El responsable eligió 1A–3A: configuración y no salud, capacidad dentro del
> detalle y actualización al recargar. No quedaron decisiones abiertas y la propuesta fue aprobada.
> Durante la implementación se confirmó que la IA obligatoria impide abrir tickets en una instalación
> sellada sin configurar. El responsable decidió mantener esa regla: el estado falso se probará en
> backend y componente, y Playwright conservará el recorrido configurado. La corrección fue aprobada
> y quedó implementada.
>
> **Actualizado el 2026-10-08 (cuarta vez):** el asistente de redacción aprobado quedó implementado.
> Soporte y Desarrollo pueden mejorar descripciones y comentarios con cinco tonos, revisar la
> propuesta y aplicarla al borrador sin publicar. El backend deriva el destinatario y aplica los
> permisos del ticket antes de delegar en IA. Pasaron backend, 224 pruebas del frontend y Playwright
> en PC y móvil; los cuatro documentos vuelven a **as-built** y producción no se tocó.
>
> **Actualizado el 2026-10-08 (tercera vez):** `modules/ai.md`, `modules/tickets.md`,
> `usuarios-y-permisos.md` e `interfaz-y-experiencia.md` pasan a **propuesta** para añadir el
> asistente de redacción con IA elegido en 1A–19A. Mejora borradores de Soporte y Desarrollo con
> contexto mínimo, revisión obligatoria y sin persistir prompts ni resultados. No quedan decisiones
> abiertas; el responsable aprobó explícitamente los cuatro documentos el 2026-10-08.
> Al preparar el código apareció una contradicción con la frontera modular. El responsable eligió
> 20A: la pantalla llamará a `/api/tickets/{number}/writing/improve` y `tickets` delegará en `ai`.
> El responsable aprobó explícitamente esta corrección; `modules/ai.md` y `modules/tickets.md`
> vuelven a **aprobado** para implementar.
> Durante la implementación se encontró que el plazo HTTP de 60 segundos cortaría un motor autorizado
> a tardar 240. El responsable eligió 21A: sólo la ruta de redacción dispondrá de 270 segundos. Los
> dos documentos vuelven a **aprobado** tras su aprobación explícita.
>
> **Actualizado el 2026-10-08 (segunda vez):** el bloque 6 cerró el repaso de formato de todas las
> pantallas en ambos idiomas, ocho temas y tres resoluciones. Corrigió en el componente compartido el
> único hallazgo real, un desbordamiento de `/users` en inglés y móvil. Pasaron 220 pruebas unitarias
> y Playwright completo (180 aprobadas y 38 omisiones previstas); quedaron cero hallazgos abiertos.
> Los cuatro documentos vuelven a **as-built** y producción permaneció fuera del recorrido.
>
> **Actualizado el 2026-10-08:** el bloque 5 verificó fuente y checksum de los tres modelos, corrigió
> el cambio 1.5B→3B y completó la validación real. La activación usa RSS recuperable, vuelve a medir,
> restaura el origen ante fallos y rechaza 7B sin interrumpir 3B; la interfaz muestra 2/4/6 GiB y el
> error con memoria requerida/disponible. Pasaron `ai-manager`, backend, 220 pruebas del frontend y
> el recorrido afectado en PC y móvil. La limpieza no dejó recursos `block5`. Los seis documentos
> del bloque vuelven a **as-built** y producción continúa aparcada hasta cerrar la versión 1.0.0.
>
> **Actualizado el 2026-10-07 (séptima vez):** se abre el bloque 5. `modules/ai.md`, `ambientes.md` y
> `prueba-local.md` pasan a **propuesta** para validar los GGUF reales 1.5B, 3B y 7B en un entorno
> desechable. El responsable eligió 1A–9A y cerró el repaso con 10A–14A: se miden recursos, tiempos,
> contrato y recuperación; se prueban `/setup` y `/settings` en PC y móvil; el 7B no se activa por
> falta de RAM; 3B y 7B se retiran primero, y la limpieza final destruye también el 1.5B junto con
> todos los recursos temporales. No quedan decisiones abiertas; los tres documentos fueron
> aprobados explícitamente el 2026-10-08 y la validación puede comenzar.
>
> **Actualizado el 2026-10-07 (sexta vez):** revisión de coherencia documental, sin cambios de
> comportamiento. Se alinean los resúmenes vigentes con los cinco pasos de `/setup`, el idioma
> global, las cinco tablas de configuración de una fila y las dos auxiliares de IA, y el motor local
> que sólo se levanta para esa modalidad. También se cierran referencias que aún describían como
> pendientes los cuatro bloques ya implementados y verificados. Todos sus documentos permanecen
> **as-built** y producción continúa aparcada hasta cerrar la versión 1.0.0.
>
> **Actualizado el 2026-10-07 (quinta vez):** el bloque 4 detecta diferencias entre la arquitectura
> aprobada y los clientes/administrador actuales de IA. `arquitectura.md`, `ambientes.md` y
> `modules/ai.md` pasaron a **propuesta** para cerrar redirecciones con credenciales, supervisión de
> `llama-server`, recursos efectivos y pruebas aisladas de los cuatro adaptadores. El responsable
> eligió 1A, 2A, 3A y 4A; el repaso cerró 5A, 6A, 7A y 8A sin asuntos abiertos. Falta únicamente la
> aprobación documental, concedida explícitamente el 2026-10-07. El bloque vuelve a **as-built** tras
> pasar backend, administrador local, 218 pruebas del frontend y 180 recorridos E2E ejecutados.
>
> **Actualizado el 2026-10-07 (cuarta vez):** al abrir el bloque 3 se encontró una contradicción en
> `primer-arranque.md`: la sección aprobada intercambiaba Correo e IA frente al resto del documento,
> el código y las pruebas. El responsable corrigió el orden: Correo permanece en el paso 4 e IA en
> el paso 5. El repaso final decidió confiar en la activación guardada al sellar, reunir privacidad
> y costo en una sola confirmación y reanudar descargas locales. El documento vuelve a
> **as-built** tras verificar servicios, 218 pruebas unitarias y 180 recorridos E2E ejecutados.
>
> **Actualizado el 2026-10-07 (tercera vez):** `flujos.md` vuelve a **as-built**. Los ocho avisos de
> tickets y los tres correos de cuenta consultan el idioma global vigente al enviarse, sin separar
> destinatarios por preferencias personales. Producción continúa aparcada hasta cerrar la 1.0.0.
>
> **Actualizado el 2026-10-07 (segunda vez):** `usuarios-y-permisos.md`, `modules/auth.md` y
> `modules/users.md` vuelven a **as-built**. La columna `users.language`, los campos de API y el
> estado de sesión desaparecen; los correos de cuenta consultan el idioma global. Producción sigue
> aparcada hasta el cierre de la versión 1.0.0.
>
> **Actualizado el 2026-10-07:** se corrige la propuesta coordinada del idioma. Sólo las once
> plantillas de correo se traducen; siempre pueden editarse manualmente y sólo un proveedor externo
> exige confirmar la estimación económica. Los resúmenes existentes se conservan visibles en su
> idioma original; los nuevos o recalculados usan el idioma global vigente. La corrección fue
> aprobada explícitamente y habilita implementar.
>
> **Actualizado el 2026-10-06 (tercera vez):** pasan a **propuesta** las enmiendas coordinadas de
> `modules/ai.md`, `modules/settings.md`, `primer-arranque.md`, `usuarios-y-permisos.md`,
> `modules/auth.md`, `modules/users.md`, `modules/mail.md`, `modules/tickets.md`, `flujos.md`,
> `interfaz-y-experiencia.md`, `arquitectura.md` y `ambientes.md`. Definen IA obligatoria —local,
> servidor propio o proveedor— administrada desde
> la interfaz, e idioma único para interfaz, correos e IA. El repaso quedó cerrado sin decisiones
> abiertas y el responsable aprobó explícitamente la propuesta. La implementación puede comenzar.
>
> **Actualizado el 2026-10-06 (segunda vez):** `prueba-local.md`, `ambientes.md` y
> `arquitectura.md` vuelven a **as-built**. La corrección conserva el código de pruebas de sólo
> lectura, arranca su backend una vez con `go run .` y retira `backend_tmp`. La ejecución canónica
> sin perfiles opcionales terminó con 177 casos aprobados y 39 omitidos, salida 0; `down -v` eliminó
> todos los recursos.
>
> **Actualizado el 2026-10-06:** `interfaz-y-experiencia.md`, `modules/settings.md` y
> `primer-arranque.md` vuelven a **as-built**. La enmienda sustituye los avisos generales situados al
> principio de `/settings` y `/setup` por un toast compartido, visible desde la posición actual. Los
> avisos permanentes permanecen en su sección. El repaso quedó cerrado, el responsable aprobó los
> tres documentos y la implementación pasó 220 pruebas unitarias, la compilación y los recorridos
> afectados de Playwright en PC y móvil.
>
> **Actualizado el 2026-10-05 (tercera vez)**: `modules/settings.md` e
> `interfaz-y-experiencia.md` vuelven a `as-built` tras mostrar explícitamente en `/settings` la zona
> seleccionada, con el mismo campo de sólo lectura ya usado en `/setup`. No cambia la persistencia,
> la API ni la base de datos. El repaso quedó cerrado, el responsable aprobó explícitamente ambas
> propuestas y las pruebas pasaron en PC y móvil.
>
> **Actualizado el 2026-10-05 (segunda vez)**: `interfaz-y-experiencia.md` vuelve a `as-built` tras
> encontrar que `/mail` coloca a la derecha su flecha de regreso aunque la regla aprobada y el propio
> comentario de la vista dicen que va a la izquierda. El inventario confirma que las otras tres
> cabeceras con este control ya cumplen. El alcance no incluye enlaces de flujo dentro del contenido;
> el repaso se cerró, el responsable aprobó explícitamente la enmienda y el cambio pasó las pruebas
> en PC y móvil.
>
> **Actualizado el 2026-10-05**: `primer-arranque.md` vuelve a `as-built`: el paso 3 de `/setup`
> muestra la región seleccionada en un campo de sólo lectura antes del buscador. La enmienda fue
> repasada, aprobada, implementada y verificada en PC y móvil. `interfaz-y-experiencia.md` recoge la
> misma regla visual.
>
> **Actualizado el 2026-10-04**: `prueba-local.md` vuelve a `as-built` tras implementar y verificar
> la migración automática antes del backend, `/setup` inicialmente en inglés y los valores
> editables de Mailpit en desarrollo.
>
> **Actualizado el 2026-10-03**: se añade `prueba-local.md`, implementado y verificado en Linux, para el
> recorrido local del arranque vacío, el esquema con un comando, los ejemplos que conservan
> configuración, Keycloak local y la suite aislada. Su enmienda §5.1.b también fue aprobada explícitamente.
>
> **Actualizado el 2026-10-02**: revisión del README para quien prueba el producto: instalación
> vacía primero y ejemplos opcionales, por corrección del responsable; `admin`/`admin` en local y
> `ADMIN_PASSWORD` en producción. Se corrige el pendiente desactualizado del asistente en
> `arquitectura.md` (7 enmiendas) y se reportan los límites reales del despliegue actual en
> `ambientes.md` (19 enmiendas). No se modifica código ni se abre producción.
>
> **Actualizado el 2026-10-01 (séptima vez)**: **los datos de ejemplo se siembran con un contenedor de
> un solo uso**, `docker compose -f dev.yml run --rm seed`, **igual en Linux, macOS y Windows** —deja
> fuera el problema del guion de bash, que en Windows no se ejecuta, y con él la copia de los
> adjuntos—. El servicio vive en `dev.yml` **con perfil propio**, así que **`up -d` no lo arranca**, y
> **los ejemplos son opcionales** (el esquema sigue siendo obligatorio). El contenedor corre **el guion
> de siempre, `scripts/dev-seed.sh`**, que también sigue valiendo en Linux y macOS fuera del
> contenedor. Quedan enmendados `ambientes.md` (**18 enmiendas**: secciones 3.1, 3.3, 5 y 9.3) y
> `arquitectura.md` (**6 enmiendas**: sección 10.3), y el `README.md` quita el aviso de Windows y el de
> los adjuntos. Comprobado el 2026-10-01: **dos pasadas seguidas** del comando («Listo: 11 cuentas, 25
> tickets y 5 adjuntos» las dos veces), 11 cuentas y 25 tickets en la base, **5 adjuntos en `_files/`**
> y `up -d` sin arrancar el servicio.
>
> **Actualizado el 2026-10-01 (sexta vez)**: **el servidor de desarrollo de Angular reenvía `/api` al
> backend**, así que **se entra en `http://127.0.0.1:11001` con `admin`/`admin` sin nginx** —el
> recorrido de quien se descarga el proyecto y sólo tiene Docker—. Se añade `frontend/proxy.conf.json`
> y se declara en `serve.options.proxyConfig` de `frontend/angular.json`, con destino
> **`http://backend:11002`** (el servicio dentro de la red del entorno) y **sin reescribir la ruta**.
> Es **sólo de desarrollo**: el `build` no la ve y **en producción sigue reenviando nginx**. El
> `README.md` quita el aviso de «pendiente de decisión». Enmendados: `arquitectura.md` (**5
> enmiendas**: la sección 9, la 10, la 11 y la fila de la 13, y nginx deja de ser necesario para
> entrar) y `ambientes.md` (**17 enmiendas**: sección 3.1 y 9.3). Comprobado el 2026-10-01 por la
> 11001: `GET /api/auth/methods` y `GET /api/health` → **200 JSON**, `POST /api/auth/login` con
> `admin`/`admin` → **200** con token, y Playwright contra `http://127.0.0.1:11001` → **14 casos en
> verde** (entrar y salir con la cuenta de fábrica).
>
> **Actualizado el 2026-10-01 (quinta vez)**: **el esquema y los ejemplos se pueden aplicar sin `bash`**.
> Los pasos de local del `README.md` quedan como un recorrido para quien llega de fuera —clonar,
> levantar, aplicar el esquema, entrar con `admin`/`admin`—, con el porqué de cada paso, y **los dos
> comandos de Docker que valen igual en PowerShell, CMD y bash**. `./scripts/dev-seed.sh` queda dicho
> como **lo que es, de Linux y macOS** (y como lo que es: **un extra de contenido, no un requisito para
> entrar**). `ambientes.md` pasa a **16 enmiendas** (su sección 3.2 estrena los dos comandos y las
> secciones 3.1, 3.3 y 5 separan las dos vías) y `arquitectura.md` a **4** (su sección 10.3, que
> empezaba por `up -d` sin el esquema, ya lo trae). Comprobado el 2026-10-01 sobre una base nueva
> (`catalina_support_prueba`, borrada después): el esquema deja 18 tablas, `GET /api/auth/methods` y
> `POST /api/auth/login` con `admin`/`admin` responden **200 sin datos de ejemplo**, y el mismo archivo
> aplicado dos veces seguidas da **0 errores**.
>
> **Actualizado el 2026-10-01 (cuarta vez)**: el `README.md` deja claro, en los pasos de local, que
> **en una base nueva `./scripts/dev-seed.sh` es lo que crea el esquema** —no sólo los ejemplos—, que
> **hasta que no se pasa la aplicación no tiene tablas** y que los síntomas son
> `relación "installation_settings" does not exist` y `relación "ai_insights" does not exist`; y
> **reordena los pasos** (entorno, después esquema y datos, y sólo entonces abrir). `ambientes.md`
> pasa a **15 enmiendas**: sus secciones 3.1, 3.3 y 5 dicen qué aplica el esquema en desarrollo y
> **cuándo** —antes de abrir la aplicación— y apuntan al `README.md` en vez de repetirlo.
>
> **Actualizado el 2026-10-01 (tercera vez)**: **el directorio de pruebas y Keycloak dejan `dev.yml`**
> y pasan a **`active-directory.yml`** y **`keycloak.yml`**, cada uno con **su propio comando**, y **el
> perfil `auth` desaparece**. La suite de los caminos de directorio pasa a **dos pasos** —levantar cada
> servicio con su archivo y correr la suite normal—, y los dos entran en la red
> **`catalina-support-dev`**, que **posee `dev.yml`** con un nombre fijo (sin `external`). Quedan
> enmendados `arquitectura.md` (**3 enmiendas**), `ambientes.md` (**14**) y `modules/auth.md` (**7**).
> Y el `README.md` dice ya **la contraseña de fábrica de desarrollo** —la de `config/env/dev.env`, que
> se versiona—, dejando claro que **la de producción es `ADMIN_PASSWORD` y no se escribe**.
>
> **Actualizado el 2026-10-01 (segunda vez)**: el `README.md` estrena **la tabla de las once cuentas
> de ejemplo de desarrollo** —nombre, correo, rol y contraseña—, y `ambientes.md` pasa a **13
> enmiendas** al **dejar de repetir esa lista** y apuntar a la tabla del `README.md` (sección 3.3),
> para que el dato viva en un solo sitio.
>
> **Actualizado el 2026-10-01**: los **tres elementos opcionales** —el motor de IA, el directorio de
> pruebas y Keycloak— quedan con **un comando de Docker cada uno** y dicho que son opcionales:
> `ambientes.md` pasa a **12 enmiendas** (sección 9.3: el directorio y Keycloak comparten el perfil
> `auth` pero **pueden levantarse por separado** nombrando el servicio, sin activar el perfil
> —comprobado—, y el comando de siempre `--profile auth up -d` sigue levantando los dos),
> `modules/ai.md` a **2 enmiendas** (se corrigen los recursos del motor: **1500m de memoria**, **sin
> tope de CPU** y **~1,44 GiB medidos**, donde decía «950 MB / 1,09 GB» y «2 CPU»; y se explica el
> **aviso esperado del volumen `ai_modelos`**) y `arquitectura.md` a **2 enmiendas** (la fila del motor
> medido, en la sección 13, con las mismas cifras: **1500 MiB de tope**, **~1,44 GiB —el 98% del
> tope—** y **sin tope de CPU**).
>
> **Actualizado el 2026-09-30 (tercera vez)**: los dos documentos de módulo afectados por las pruebas
> del asistente quedan al día: `modules/mail.md` pasa a **6 enmiendas** —estrena **`Probar`**, que
> conecta y autentica sin mandar ningún correo— y `modules/settings.md` a **8 enmiendas** —declara
> **`ProberDeCorreo`**, que el asistente reutiliza para el paso 4—.
>
> **Actualizado el 2026-09-30 (segunda vez)**: **el asistente de primer arranque prueba lo que pide**
> —sus dos endpoints públicos `POST /api/setup/entry/test` y `POST /api/setup/mail/test`, con el mismo
> candado del sello, y los botones de los pasos 2 y 4—, y la prueba del correo **conecta y autentica
> sin mandar ningún correo**. `primer-arranque.md` pasa a **1 enmienda**
> (`docs/primer-arranque.md`, sección 3.1).
>
> **Actualizado el 2026-09-30**: **el motor de IA se configura desde la pantalla de Configuración**
> —su dirección y su modelo, con su botón de «Probar la conexión»—, con el entorno como respaldo
> (`docs/modules/settings.md`, decisión 16, y `docs/modules/ai.md`, decisión 21). `modules/ai.md`
> pasa a **as-built**.
>
> **Actualizado el 2026-09-27**: nace `modules/ai.md`, el sexto módulo, con los dos campos que
> redacta el motor de IA en la lista de tickets.
>
> **Actualizado el 2026-09-26 (tercera vez)**: los adjuntos de texto y código —un `.sql` no se podía
> mandar— y **la imagen y el vídeo con tope de 480 × 360** que se abren en el visor al pulsarlos
> (`docs/modules/tickets.md`, decisiones 54 y 55).
>
> **Actualizado el 2026-09-26 (segunda vez)**: las listas de tickets pasan a ser **Mis tickets** (lo
> mío), **Tickets principales** y **Tickets internos**, y se quitan del menú «Nuevo ticket» y de la
> bandeja «Asignármelo» (decisiones del responsable; `docs/modules/tickets.md`, decisiones 50 a 53).
>
> **Actualizado el 2026-09-26**: se fija **cómo se cuenta una enmienda** (sección «El registro de
> enmiendas y cómo se cuenta») y se ponen los números de la tabla con esa regla, porque hasta hoy
> unos cuadraban y otros no; y `modules/mail.md` y `modules/users.md` pasan a **as-built**, que era
> lo que les tocaba desde que su código existe.

Índice de la documentación del proyecto. La **Regla 0** de `AGENTS.md` es obligatoria: no se
escribe código sin un documento aprobado antes.

## Documentos

| Documento | Estado | Cubre |
| --- | --- | --- |
| `README.md` (este archivo) | as-built | Índice, convenciones y estados |
| `proyecto-abierto.md` | **as-built** | Portada bilingüe, sitio oficial, demo pública y apoyo mediante Ko-fi |
| `prueba-local.md` | **as-built** | Verificación local de aceleración, respaldo y rollback |
| `arquitectura.md` | **as-built** | Frontera CPU/NVIDIA y reinicio transaccional |
| `propósito-y-alcance.md` | **as-built** (3 enmiendas previas) | Dashboard exclusivamente estadístico, sin evaluación ni SLA |
| `usuarios-y-permisos.md` | **as-built** | Permisos acumulativos y privacidad de comentarios del principal |
| `flujos.md` | **as-built** | Aviso único a todo Soporte al llegar por primera vez a crítica |
| `modules/settings.md` | **as-built** | Selección y estado efectivo de aceleración local |
| `modules/mail.md` | **as-built** | Aviso interno de comentario privado sin contenido sensible |
| `modules/auth.md` | **as-built** | Identidad con papel operativo y capacidad administrativa |
| `modules/users.md` | **as-built** | Persistencia, API, auditoría y avisos de administración |
| `interfaz-y-experiencia.md` | **as-built** | Selector Public/Private y representación accesible de notas privadas |
| `primer-arranque.md` | **as-built** (8 enmiendas previas) | Aceleración local en el paso 5 |
| `ambientes.md` | **as-built** | Variante opcional NVIDIA en Linux |
| `modules/ai.md` | **as-built** | Summary y Last action separados por audiencia |
| `modules/tickets.md` | **as-built** | Comentarios privados, adjuntos, búsqueda y auditoría del principal |
| `modules/dashboard.md` | **as-built** | Dashboard agregado, métricas, gráficos y XLSX localizado |

La cadena de producto **está completa**: siete documentos que cubren qué se construye, cómo se
comporta, cómo se ve y cómo se despliega. Además, **cada módulo tiene su documento**, que se escribe
justo antes de implementarlo, en este orden:

1. `modules/mail.md` — **entero**: backend y el editor.
2. `modules/settings.md` — **la pantalla de Configuración está hecha entera**: la marca, el color, el
   prefijo, el reparto y el idioma de la instalación.
3. `modules/auth.md` — **entero**: los tres caminos de entrada, hechos y verificados.
4. `modules/users.md` — **entero**: backend, las tres pantallas y las tres acciones del directorio.
5. `modules/tickets.md` — **entero**: backend y pantallas, con los adjuntos de los comentarios.

## Convenciones

### Cabecera obligatoria

Todo documento de `docs/` empieza así:

```markdown
# Título

> **Estado:** propuesta | aprobado | as-built
> **Última actualización:** AAAA-MM-DD
```

- `propuesta`: escrito, pendiente de aprobación. No se toca código de esa área.
- `aprobado`: aprobado por el responsable del proyecto. Habilita implementar.
- `as-built`: el código existe y el documento describe lo que hace hoy.

### El registro de enmiendas y cómo se cuenta

Debajo de la cabecera, cada documento lleva **su registro de enmiendas**, de la más reciente a la más
antigua: **una entrada por cambio**, con su fecha y qué cambió y por qué.

**Una enmienda es una entrada del registro que empieza por `Enmendado el …`.** Eso es lo que dice el
número que aparece en la tabla de arriba y en `AGENTS.md`, y se puede comprobar contando esas
entradas en la cabecera del documento. **`Pasa a as-built el …` no cuenta**: no cambia nada del
documento, cambia su estado.

Dos cosas se escriben aparte, y no son enmiendas:

- **La corrección del responsable**, cuando cambió lo propuesto, se anota **dentro de la entrada**
  («es una corrección del responsable», «yo proponía… y decidió…»).
- **Lo que se descubrió al implementar** también va dentro de su entrada, diciendo qué se encontró.

El registro es memoria de **por qué** un documento está como está. No se tachan ni se borran
entradas: si algo se corrige, se añade la entrada nueva y se dice qué queda corregido.

### Nombres y carpetas

- **Los documentos de producto** viven en `docs/`, y se llaman en español, en minúsculas, sin acentos
  ni espacios: `usuarios-y-permisos.md`.
- **Los documentos de módulo** viven en `docs/modules/` y se llaman **como el módulo**:
  `modules/auth.md`, `modules/users.md`. El nombre del módulo manda ahí, aunque sea en inglés, porque
  es el nombre que tiene su carpeta en el código.
- Un documento de módulo cuenta **cómo se construye** ese módulo. Lo que es decisión de producto —qué
  puede cada papel, cómo se comportan los tickets— vive en los documentos de `docs/`, y el de módulo
  no lo repite: lo da por escrito y apunta a él.

### Correspondencia con el código

La tabla documento ↔ código vive en `AGENTS.md`. Cuando un documento pase a `as-built`, hay que
añadir o actualizar su fila allí en el mismo cambio.

### Reglas de escritura

- Español, frases cortas, sin relleno y sin adornos de marketing.
- Cada documento dice **qué está decidido**, **qué está pendiente** y **qué queda fuera de alcance**.
- Lo que no esté decidido se marca como pendiente; no se rellena con suposiciones.
- Cuando un documento describa algo ya implementado, describe lo que el código hace de verdad,
  no lo que se pretendía.
