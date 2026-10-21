# Prueba local: instalación vacía, ejemplos y verificación aislada

> **Estado:** as-built
> **Última actualización:** 2026-10-21
>
> **Enmienda coordinada implementada y verificada el 2026-10-21: demo pública alojada.** La demo
> usa capacidades genéricas opt-in del producto y operación concreta en el repositorio privado;
> quedó publicada por HTTPS, con IA local, antivirus, cuotas y restauración A/B diaria. Producción
> no se modificó.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-20.** La sección 27 verifica la
> actualización autosuficiente y las migraciones versionadas. El responsable cerró y aprobó las
> decisiones 1A–25A. No se tocó producción.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-20: validar el envío SMTP de prueba.** La sección 26
> cubre destinatario efímero, `From:` real, datos sin guardar, secreto persistido, idioma, límite y
> exactitud del resultado. El responsable cerró y aprobó 1A–14A; las pruebas focalizadas y las 248
> del frontend pasaron, igual que fronteras y compilación.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-19: matriz de transporte SMTP.** La
> sección 25 verifica los tres modos, la ausencia de degradación y el límite privado del modo sin
> cifrar. El responsable cerró y aprobó 1A–10A; Mailpit, regresiones y relay real quedaron
> comprobados.
>
> **Aprobado e implementado el 2026-10-18.** La sección 24 registra la verificación portable de CPU,
> detección NVIDIA, respaldo Auto y restauración transaccional sin tocar producción.
>
> **Implementado y verificado el 2026-10-18.** La sección 23 añade la verificación del Dashboard,
> sus cálculos, permisos, CSV y responsive.
>
> **Enmienda implementada y verificada el 2026-10-17: README bilingüe y capturas públicas.** La
> sección 22 registra la comprobación del recorrido reordenado y de su equivalencia en español. Se
> restauraron los datos locales, la marca de fábrica y se verificaron visualmente la vista doble y
> el panel de IA. No se tocó ningún entorno público.
>
> **Enmienda implementada para 1.0.0, 2026-10-14.** La sección 21 añade la comprobación local de las
> cuatro prioridades, orden, filtros, historial y aviso crítico no bloqueante. No se modifica código
> antes de implementarla. El responsable la aprobó explícitamente.
>
> **Cuarta pasada pública completada, 2026-10-13.** La sección 20 registra el recorrido íntegro
> desde `main`: instalación vacía, modelo local real, correo, cuentas, ticket, IA, datos de ejemplo,
> AD, Keycloak y regreso a local. La inspección terminó sin falsos errores ni hallazgos abiertos y
> todos los recursos desechables fueron retirados; producción permaneció intacta.
>
> **Tercera reanudación aprobada, 2026-10-13.** El recorrido completó instalación, IA, correo,
> ejemplos, AD y Keycloak, pero la revisión final encontró falsos errores `record not found` en
> búsquedas cuya ausencia es normal. La sección 19 conserva la evidencia y exige una cuarta pasada
> completa después de la corrección. El responsable eligió y aprobó 1A–4A; la corrección ya pasó
> contra PostgreSQL real y falta integrarla en `main`.
>
> **Segunda reanudación aprobada el 2026-10-13.** El bloque 8 se detuvo antes de `/setup` porque `migrate` recibió `connection refused` después de un healthcheck por socket prematuro. La sección 18 registra tiempos, limpieza y repetición completa desde otro `main` corregido. El responsable aprobó la corrección y el plan; falta publicar y ejecutar la repetición completa.
>
>
> **Hallazgo del bloque 8 y reanudación propuesta, 2026-10-13.** La primera ejecución se detuvo al
> obtener un «Motivo» en español en un despliegue inglés. La sección 17 conserva la evidencia,
> remite la corrección a los documentos de IA y tickets y exige repetir el recorrido completo desde
> `main`. El responsable eligió 1A–12A, con 10B, y aprobó explícitamente esta enmienda el 2026-10-13.
>
> **Propuesta del bloque 8, 2026-10-12: recorrido público desde cero.** La sección 16 valida la
> experiencia de una persona externa desde un clon de `main`, usando exclusivamente el README, una
> instalación vacía, el modelo local real, los seeders opcionales y los tres métodos de entrada. El
> responsable cerró 1A–15B y aprobó explícitamente la propuesta el 2026-10-12. Producción, dominio
> y la etiqueta `v1.0.0` siguen fuera.
>
> **Enmienda implementada y verificada el 2026-10-08 (bloque 6).** La sección 15 añade el recorrido reproducible del
> repaso visual completo. Lo existente continúa as-built y la propuesta requiere aprobación antes
> de modificar frontend o pruebas; el responsable la aprobó explícitamente el 2026-10-08.
>
> **Enmienda implementada y verificada el 2026-10-08 (bloque 5).** La sección 14 contiene el recorrido manual aislado
> para validar los tres modelos reales, `/setup` y `/settings`, sin tocar desarrollo. El responsable
> cerró las decisiones 1A–14A. Lo existente continúa as-built; no se descargan modelos ni se cambia
> código hasta recibir aprobación explícita, concedida por el responsable el 2026-10-08.
>
> **Enmendado el 2026-10-06**, aprobado explícitamente por el responsable, implementado y verificado
> en el mismo trabajo. Al repetir la suite aislada se encontró que Docker no podía
> montar `backend_tmp:/app/tmp` dentro del bind `./backend:/app:ro`: intenta crear el punto de montaje
> bajo un padre de sólo lectura y el backend no llegaba a arrancar. El backend de `tests.yml` ejecuta
> ahora `go run .` una sola vez, conserva `/app:ro` y elimina ese volumen. La ejecución canónica sin
> servicios opcionales terminó con 177 casos aprobados y 39 omitidos, y la limpieza retiró todos los
> recursos de la instancia.
>
> **Enmendado el 2026-10-04**, aprobado explícitamente por el responsable e implementado en el mismo
> trabajo. Corrige el ruido de una base nueva —el backend consulta tablas antes de aplicar el
> esquema—, hace que el primer arranque
> empiece en inglés y que su selector traduzca toda la interfaz al instante, y precarga en desarrollo
> los valores editables de Mailpit. Las tres decisiones fueron confirmadas por el responsable antes
> de escribir esta propuesta. El repaso quedó cerrado y la aprobación «Lo apruebo» habilita su
> implementación.
>
> **Propuesta del 2026-10-03, aprobada explícitamente por el responsable.** Reúne las decisiones del
> responsable de las dos tandas de preguntas. La primera eligió local primero, esquema con un
> comando de Docker, pruebas locales y Keycloak opcional en localhost. La segunda eligió separar
> direcciones pública e interna de Keycloak, aislar las pruebas y conservar la configuración al
> cargar ejemplos. **Corrección previa del responsable:** la instalación vacía es el recorrido
> principal; el agente había sugerido comenzar con ejemplos. Esta propuesta mantiene esa corrección.
> Aprobación recibida después del repaso: «Lo apruebo». Implementado y verificado;
> la enmienda §5.1.b también recibió aprobación explícita.

## 1. Problema y resultado esperado

Quien descarga el proyecto debe poder seguir el README sin el dominio, nginx, certificados ni datos
previos de esta máquina. Primero levanta una instalación vacía, aplica el esquema con un comando,
completa el asistente y entra con la cuenta de fábrica. Puede crear cuentas y tickets propios,
cargar ejemplos si los quiere y probar AD o Keycloak como servicios opcionales.

Las pruebas automáticas se ejecutan contra una instalación separada. No cambian las cuentas,
tickets, configuración, correo ni adjuntos de la instalación que está usando esa persona.

Esta propuesta cubre **desarrollo y evaluación local**. No habilita desplegar producción.

## 2. Hallazgos que motivan la propuesta

| Hallazgo | Evidencia actual | Cambio propuesto |
| --- | --- | --- |
| Aplicar sólo el esquema exige copiar SQL y llamar a `psql` manualmente | README y `dev.yml` | Servicio de un solo uso para el esquema |
| Playwright apunta al dominio del proyecto | `dev.yml`, `playwright.config.ts` y preparación | Compose propio y direcciones de su red |
| Angular rechaza el nombre `frontend` | `frontend/angular.json` y comprobación del 2026-10-02 | Admitir explícitamente ese host de desarrollo |
| Los ejemplos sustituyen los ajustes del usuario | Últimos bloques de `v1.0.0_dev.sql` | Separar datos de ejemplo y configuración |
| Keycloak anuncia el dominio original | `keycloak.yml` y reino importado | Dirección pública local y conexión interna independiente |
| El cliente OIDC no lee ni comprueba `issuer` en descubrimiento | `backend/modules/auth/services/keycloak.go` | Comprobar el emisor público antes de utilizar los endpoints |

Los límites del script de producción ya reportados permanecen fuera de alcance. No se ejecuta
`prod-build.sh`, no se abre el dominio ni se crea la etiqueta de versión.

## 3. Recorrido principal: sin ejemplos

Comandos vigentes, desde la raíz del repositorio, iguales en PowerShell, CMD y bash:

```text
docker compose -f dev.yml up -d --build
docker compose -f dev.yml run --rm migrate
```

`migrate` es el nombre aprobado para el servicio que sólo aplica el esquema vigente:

- Tiene perfil propio y no arranca con `up -d`.
- Espera a PostgreSQL saludable, monta la migración en sólo lectura y usa el cliente de PostgreSQL
  dentro del contenedor. No necesita Docker, Go, Node ni `psql` en la máquina.
- Usa `ON_ERROR_STOP=1`, conserva la transacción e idempotencia del SQL y falla con código distinto
  de cero si la migración falla. El mensaje de éxito sólo aparece tras completarla.
- No carga ejemplos, no borra datos, no configura SMTP y no sella una base nueva.
- Repetirlo no reabre una instalación sellada ni altera sus ajustes.
- Mientras la 1.0.0 siga abierta aplica `v1.0.0.sql`; no introduce un sistema general de migraciones.

Al abrir `http://127.0.0.1:11001`, una base nueva enseña `/setup`. El README indica los cinco pasos,
el método local, los valores de Mailpit (`mail`, `1025`, sin TLS ni credenciales) y la IA, con sus pruebas.
Después de terminar se entra con `admin` / `admin`, se crean cuentas y se abren sus enlaces en
`http://127.0.0.1:11004`. La contraseña de producción sigue siendo `ADMIN_PASSWORD` del entorno.

No se cambia el flujo del asistente ni se añade un acceso que lo salte.

## 4. Ejemplos opcionales que conservan la configuración

El comando sigue siendo `docker compose -f dev.yml run --rm seed`.

En una instalación **sellada**, no modifica `installation_settings`, `directory_settings`,
`keycloak_settings` ni `ticket_settings`: conserva nombre, idioma, marca, región, dirección,
SMTP, método de entrada, directorios, IA, prefijo y reparto. Conserva el sello.

En una base **sin sellar**, completa la preparación local de los ejemplos con valores de demo:
entrada local, dirección `http://127.0.0.1:11001`, SMTP de Mailpit y ajustes necesarios para las
cuentas y tickets de ejemplo; deja la demo sellada. El README advierte que, en este caso, el seeder
sustituye el recorrido del asistente y puede reemplazar sus pasos todavía incompletos.

Se mantiene la semántica destructiva de los datos de ejemplo: se recrean tickets, categorías,
etiquetas, contadores, cuentas `@demo.com` y archivos adjuntos de desarrollo; se preserva la marca.
Conservar configuración **no** significa conservar esos datos. El aviso debe aparecer antes del
comando, tanto en el README como en su salida.

Las once cuentas mantienen sus credenciales de demo. Si la instalación conservada usa AD o
Keycloak, las cuentas locales de ejemplo no entran hasta volver a Local con la cuenta de fábrica.
No se cambia el método para hacerlas entrar automáticamente.

## 5. Keycloak: identidad pública y conexión interna

### 5.1 Configuración

Se añade una dirección interna **opcional** al bloque de Keycloak, tanto en Configuración como en
el paso de entrada del asistente. Vive en `keycloak_settings`, con migración idempotente dentro de
la versión abierta, y viaja en los DTO de settings y en `config.OIDC`.

Nombre de contrato aprobado: `internalIssuer` en JSON, `internal_issuer` en PostgreSQL. Es una URL
completa del mismo reino, no otro emisor ni un secreto. Vacía significa usar las direcciones
públicas, conservando las instalaciones que ya funcionan. No se añaden variables de entorno para
sustituir una configuración que vive en la base.

| Dato | Local en el navegador de la máquina |
| --- | --- |
| Aplicación y dirección pública | `http://127.0.0.1:11001` |
| Emisor público | `http://127.0.0.1:11006/sso/realms/catalina-support` |
| Dirección interna del mismo reino | `http://keycloak:8080/sso/realms/catalina-support` |
| Vuelta | `http://127.0.0.1:11001/api/auth/keycloak/callback` |
| Cliente y secreto de pruebas | Los del reino versionado, documentados como credenciales locales |

Keycloak sigue en su compose propio y en la red del entorno. El navegador llega directamente al
puerto 11006; no se necesita añadir `/sso` al proxy de Angular ni instalar nginx.
El reino permite la vuelta local exacta. No se añaden comodines de redirección.
La dirección anunciada por el contenedor debe poder configurarse para el entorno aislado de pruebas;
no se obliga a editar el reino ni a tener un archivo de hosts en la máquina.

### 5.1.b Enmienda aprobada: configurar un método todavía vacío

**Estado de esta enmienda: aprobada explícitamente por el responsable el 2026-10-03 («Lo apruebo»).** Hallazgo al verificar
Keycloak desde una instalación vacía: Configuración deshabilita AD y Keycloak si sus campos están
vacíos, pero los campos sólo se muestran al elegir ese método. No hay forma de rellenarlos por la
interfaz. El asistente ya permite esa elección provisional.

Se permite elegir AD o Keycloak en el formulario de Configuración aunque sus datos estén
vacíos. Elegir sólo muestra los campos; **el método en funcionamiento no cambia hasta guardar**.
El backend conserva su validación: guardar un método incompleto devuelve 422. Se conserva la cuenta
de fábrica. No se añaden pantallas, endpoints ni permisos. Se verificó que una instalación local
vacía permite rellenar y probar los dos métodos sin cargar ejemplos y que un guardado incompleto
no cambia el método persistido. La aprobación habilita el cambio del selector y su prueba de regresión.

### 5.2 Reglas de conexión y validación

- El descubrimiento se obtiene desde `internalIssuer`, si está configurado; en caso contrario,
  desde el emisor público.
- El `issuer` anunciado debe coincidir con el emisor público configurado. Ausente o distinto:
  la prueba y el acceso fallan con una clave traducida, antes de canjear credenciales.
- La URL de autorización enviada al navegador permanece pública.
- Sólo para peticiones del backend se sustituye el prefijo público del reino por el interno:
  descubrimiento, canje del código y lectura de UserInfo. Se conserva el resto de la ruta y consulta.
- La sustitución debe comparar URLs analizadas y límites de ruta, no reemplazar texto arbitrario.
  Los endpoints deben pertenecer al mismo origen y reino esperado; no se transforma un host o
  prefijo parecido ni se sigue una redirección a otro origen con secretos o tokens.
- Se admiten HTTP y HTTPS, puertos y rutas de reino; se rechazan credenciales incrustadas,
  fragmentos y consultas en las dos direcciones de configuración. No se desactiva la validación TLS.
- Vaciar la dirección interna restaura la conexión pública. Cambiar los ajustes invalida el
  descubrimiento recordado, como corresponde a una configuración leída en cada intento.
- Se mantienen el cliente confidencial, el `state` firmado y la sesión en el fragmento de vuelta.
  No se añade PKCE ni una nueva validación de ID tokens en este trabajo.

Ambos campos se explican con texto sencillo y ejemplos, en español e inglés. No se muestra la
URL interna al usuario que entra por Keycloak; el administrador es quien la configura.

## 6. Pruebas en una instalación aislada

Se implementa `tests.yml` en la raíz, con proyecto distinto del de desarrollo y producción:

- Contenedores, redes y volúmenes propios, sin nombres globales que coincidan con los actuales.
- Fuente montada sólo para construir o ejecutar; adjuntos, logs, dependencias y base en volúmenes
  propios. No monta `_files/`, `_logs/` ni los volúmenes de la instalación local.
- Frontend, backend, PostgreSQL y Mailpit sin puertos publicados en la máquina.
- LDAP y Keycloak opcionales mediante un perfil de pruebas de directorio, dentro de esta red.
  No reutilizan los servicios ni el reino de la instalación local.
- No entra en la red del motor local ni descarga modelos. Usa `ai-test`, un servidor compatible falso
  dentro de su propia red, para completar y probar la IA obligatoria sin proveedores reales.
- El navegador usa `http://frontend.localhost:11001`, permitido explícitamente en Angular. Los enlaces de
  correo usan esa misma base en la configuración exclusiva de pruebas.
- El Keycloak aislado anuncia `http://keycloak:8080/sso`; cliente, emisor, URL interna y vuelta
  coinciden con las direcciones accesibles al navegador del contenedor.
- Chromium resuelve `frontend.localhost` al servicio `frontend` sólo dentro del navegador de pruebas
  aisladas; el sufijo localhost proporciona un contexto seguro, para ejercer el portapapeles real como en localhost. No se simula la API ni se cambia
  la seguridad de la aplicación.
- Todas las utilidades de Playwright comparten el mismo origen; no quedan respaldos divergentes
  entre configuración, preparación y limpieza.

El README muestra un recorrido de comandos Docker: preparar una instancia de pruebas nueva,
arrancar sus servicios, ejecutar Playwright y retirar **sólo** sus recursos. Las dependencias se
instalan en contenedores; el fallo del runner conserva su código de salida. El comando predeterminado
no utiliza una dirección externa ni convierte `BASE_URL` en una forma implícita de tocar la instalación
local. Un destino externo es una opción avanzada explícita y no forma parte de este recorrido.

La preparación parte del esquema y configura por la API sólo la instalación aislada. Incluye un
caso de navegador de los cinco pasos desde vacío **antes de sellarla**, seguido de la suite usual.
Las pruebas crean sus cuentas y fixtures, sin exigir el seeder opcional del usuario. Los endpoints
requieren la misma protección que en el producto, sin puertas exclusivas para pruebas.

Los informes, trazas y capturas se conservan en una carpeta de resultados identificada como de
pruebas; retirar la base desechable no borra esas evidencias. No se promete ejecución simultánea de
varias suites en este alcance. Una pasada nueva comienza con base y archivos de prueba nuevos,
también después de una pasada interrumpida; la limpieza nunca busca recursos por un prefijo amplio.

## 7. Documentos y archivos afectados

Este documento coordina el cambio; las normas de cada área siguen en su documento actual.
Las enmiendas específicas quedan registradas en los documentos afectados y describen la implementación.

| Documento actualizado | Qué decide o refleja | Archivos afectados |
| --- | --- | --- |
| `docs/ambientes.md` | Esquema con un comando, seeder conservando configuración y suite desechable | `dev.yml`, `tests.yml`, `scripts/dev-seed.sh`, migraciones y `tests/e2e/**` |
| `docs/arquitectura.md` | Red y volúmenes aislados, hosts admitidos y separación OIDC | Compose, `frontend/angular.json`, cableado del backend |
| `docs/modules/auth.md` | Descubrimiento, emisor público y conexión interna de Keycloak | `backend/modules/auth/**`, `backend/shared/config/**`, `keycloak.yml`, `config/keycloak/**` |
| `docs/modules/settings.md` | Nuevo campo interno, lectura, escritura y prueba de conexión | `backend/modules/settings/**`, migración de `keycloak_settings`, pantalla de Configuración |
| `docs/primer-arranque.md` | Campo interno opcional en el paso de entrada y cobertura desde cero | DTO y servicios de setup, `setup-page.*`, caso Playwright |
| `docs/interfaz-y-experiencia.md` | Campo y ayudas de Keycloak en las dos pantallas, textos en ambos idiomas | Configuración, asistente y diccionarios de i18n |
| `README.md` | Recorrido vacío, ejemplos opcionales y pruebas locales reproducibles | Instrucciones y ejemplos de comandos |
| `docs/README.md` y `AGENTS.md` | Estado as-built y correspondencia documental | Índice y tabla de cobertura |

Los documentos afectados se actualizaron en este trabajo, con la enmienda fechada en su cabecera.

## 8. Criterios de aceptación

1. Desde base nueva: `up` y `run --rm migrate` dejan cero cuentas y cero tickets; el navegador
   muestra el asistente y al terminar se puede entrar con `admin` / `admin`.
2. El comando de esquema se puede repetir: conserva datos, ajustes y sello. SQL inválido hace
   fallar el comando y no se anuncia éxito.
3. Alta local: el correo llega a Mailpit, su enlace abre esta instalación y permite establecer
   contraseña e iniciar sesión con la cuenta nueva.
4. Seeder sobre instalación sellada: compara antes y después las cinco tablas de configuración
   y la marca; se conservan mientras se recrean los ejemplos previstos y sus adjuntos.
5. Seeder sobre base nueva o asistente incompleto: deja una demo local utilizable y sellada, sin
   dominio externo; repetirlo no vuelve a cambiar la configuración de la demo ya sellada.
6. Keycloak local: contenedor opcional, prueba de conexión, entrada con una persona del reino,
   vuelta exacta y fragmento borrado. Sin contenedor, Local y la cuenta de fábrica siguen funcionando.
7. OIDC: emisor incorrecto, reino equivocado, endpoints fuera del prefijo y redirecciones externas
   se rechazan; se demuestra que el backend usa la dirección interna y el navegador la pública.
8. Suite aislada: asistente desde vacío, cuentas locales y correos en PC y móvil; LDAP y Keycloak
   cuando se solicita el perfil. Servicios opcionales ausentes se indican como casos omitidos.
9. Una pasada fallida o interrumpida no cambia la instalación del usuario y no deja datos de
   pruebas que contaminen la siguiente. Los informes se conservan y los fallos no se convierten en éxito.
10. Comandos y montajes se prueban en Linux. Si no hay máquinas Windows o macOS disponibles,
    se declara esa limitación y no se afirma haber verificado esos sistemas.

## 9. Fuera de alcance

Producción y su script, publicación del dominio, etiqueta de versión, cambios de roles o del flujo
de tickets, migraciones de versiones futuras, sustitución del motor de IA, PKCE, gestión de cuentas
de servicio de Keycloak y ejecución concurrente de suites. Tampoco se cambia la contraseña de fábrica
ni se obliga a instalar Go o Node en la máquina.

## 10. Registro del repaso

El responsable respondió **1A, 2A, 3A** al repaso del documento:

1. **Seeder:** completa y sella la demo cuando la base es nueva o el asistente está incompleto.
   Se aceptó para poder probar inmediatamente, con el aviso de que sustituye los pasos incompletos.
   En una instalación sellada conserva la configuración, como fija la sección 4.
2. **Keycloak:** la dirección interna es opcional y visible tanto en Configuración como en el
   asistente. Se aceptó poder configurarla sin editar archivos y rechazar el descubrimiento si
   el `issuer` falta o no coincide, como fija la sección 5.
3. **Suite:** retira exclusivamente los datos y recursos de su instancia desechable, conserva
   los informes y permite incluir LDAP y Keycloak mediante un perfil opcional. Se eligió que
   cada pasada comience limpia, como fija la sección 6.

No quedan decisiones de diseño abiertas. Después del repaso, el responsable aprobó explícitamente
el documento con «Lo apruebo». La implementación y su verificación quedaron completas, y el documento se cierra como `as-built`.

## 11. Verificación de la implementación

Comprobaciones realizadas en copias aisladas de Linux, sin modificar datos de desarrollo o producción:

- Esquema desde cero: 0 cuentas, 0 tickets y sello nulo. Repetir la migración conserva las
  11 cuentas, los 25 tickets y el hash de las cinco tablas de configuración. SQL inválido
  devuelve código 3 y no imprime el mensaje de éxito.
- Asistente, SMTP de Mailpit, admin/admin, alta local, enlace de contraseña y entrada real
  desde localhost verificados en navegador. La API del asistente sellado devuelve 409.
- Seeder sobre instalación sellada: los cinco ajustes conservan exactamente el mismo hash
  antes y después; se recrean 11 cuentas, 25 tickets y 5 adjuntos. Sobre base nueva y sello nulo
  deja método local, URL localhost y sello puesto.
- Keycloak real: autorización pública en localhost, tráfico interno del backend, prueba y
  guardado desde Configuración, acceso de una persona del reino y fragmento eliminado al volver.
- Selector vacío, campos visibles y rechazo 422 sin cambiar el método local verificados en
  PC y móvil. Las 220 pruebas de unidad del frontend pasan. Go: `go test ./...` y `go vet ./...` pasan.
- Suite con directorios: 214 casos, 193 aprobados, 20 omitidos y uno interrumpido por una
  recarga de Angular durante un ajuste de formato. Se conserva su informe, con salida 1.
  En una instalación nueva y estable, el caso interrumpido pasó en PC y móvil: 2 aprobados
  y 24 de directorio omitidos correctamente al no levantar los servicios. Entre ambas
  comprobaciones quedan 194 casos distintos verificados y 20 omisiones previstas en la suite completa.
  No se oculta ni se convierte en éxito la salida fallida de la primera pasada.
- LDAP presente: el mismo endpoint de detección devuelve 200; sin servicio los casos se omiten.
- Las pasadas nuevas comienzan con `down -v` exclusivo de `tests.yml`; la retirada de los
  contenedores y volúmenes conserva los informes y capturas. Las fuentes no reciben escrituras
  del runner y no se usan volúmenes, adjuntos ni correos del entorno del usuario.

No se han ejecutado estas comprobaciones en Windows o macOS. Producción permanece fuera de alcance.

Hallazgo durante la verificación sin servicios opcionales: `directorio.spec.ts` declaraba
`respondeElDirectorio`, pero no lo utilizaba antes de ejecutar sus casos. La documentación
prometía omitirlos cuando LDAP estuviera ausente. Se aplica la detección en `beforeAll`,
conforme al comportamiento ya aprobado en las secciones 6 y 8; no cambia el producto.

## 12. Arranque limpio y valores iniciales del asistente

### 12.1 Hallazgos

1. En una base nueva, `backend` espera a que PostgreSQL esté saludable, pero no a que exista el
   esquema. Por eso alcanza a consultar `installation_settings` y `ai_insights`, escribe errores
   `42P01` y queda escuchando aunque `/setup` todavía no puede funcionar.
2. El formulario de `/setup` inicializa el idioma en `es` y la fila nueva de
   `installation_settings` también nace con `language = 'es'`. Elegir otro idioma sólo cambia el
   valor que se guardará: no cambia los textos del asistente ni los demás textos de la aplicación.
3. Mailpit forma parte de `dev.yml`, pero una instalación vacía recibe los campos SMTP vacíos. El
   README obliga a copiar a mano los valores locales que el entorno ya conoce.

### 12.2 Arranque de desarrollo

`docker compose -f dev.yml up -d --build` aplica automáticamente `v1.0.0.sql` antes de iniciar el
backend. El servicio `migrate` deja de estar oculto tras un perfil; espera a la base saludable,
termina con éxito y sólo entonces habilita el arranque de `backend` mediante
`service_completed_successfully`.

La migración sigue siendo transaccional e idempotente. Se puede repetir con
`docker compose -f dev.yml run --rm migrate`, pero deja de ser un paso obligatorio del recorrido
normal. Si falla, `backend` no arranca y el error queda en `migrate`; no aparece un backend saludable
que en realidad carece de tablas. Esto se aplica sólo a `dev.yml`: no cambia el despliegue ni las
migraciones de producción. `seed` continúa separado, explícito, opcional y con la misma semántica
destructiva para los datos de ejemplo.

### 12.3 Idioma del primer arranque

Una instalación nueva nace con inglés como idioma de instalación. Al entrar por primera vez a
`/setup`, toda esa pantalla se presenta en inglés, aunque el navegador prefiera español o conserve
una elección anterior de otra instalación.

El selector **The installation language** sigue ofreciendo English y Español. Cambiarlo actualiza
inmediatamente todos los textos de la aplicación, incluidos los cinco pasos, ayudas, botones,
avisos y resumen de `/setup`; también actualiza `html[lang]` y recuerda la elección en el navegador.
Al guardar el primer paso, el mismo valor queda persistido como idioma global de la instalación. Si
se vuelve a un asistente empezado, manda el idioma ya
guardado en esa instalación.

El cambio del valor inicial de `installation_settings.language` de `es` a `en` afecta a bases nuevas.
No modifica automáticamente el idioma global de instalaciones existentes.

### 12.4 Valores editables de Mailpit en desarrollo

Mientras la instalación no esté sellada y no haya correo guardado, la API de `/api/setup` puede
devolver valores iniciales procedentes del entorno, exclusivos del formulario de primer arranque.
`dev.yml` declara:

| Campo | Valor inicial de desarrollo |
| --- | --- |
| Servidor | `mail` |
| Puerto | `1025` |
| TLS | no |
| Usuario y contraseña | vacíos |
| Nombre del remitente | `Catalina Support` |
| Correo del remitente | `no-responder@catalina-support.local` |

Son sugerencias editables: el usuario puede cambiar cualquier campo antes de probar o guardar. Sólo
se persisten al guardar el paso 4. No son un segundo mecanismo de configuración SMTP ni un respaldo
para enviar correo después de instalar: el módulo `mail` continúa leyendo exclusivamente la
configuración guardada. En producción estas variables no se declaran y el formulario permanece
vacío. Si el asistente ya tiene correo guardado, siempre se devuelve lo guardado y no los valores
iniciales.

### 12.5 Documentos y código de la implementación

| Documento | Cambio al cerrar como `as-built` |
| --- | --- |
| `docs/prueba-local.md` | Un solo comando de arranque, seeders opcionales y valores locales iniciales |
| `docs/ambientes.md` y `docs/arquitectura.md` | Orden real de servicios y fallo de migración |
| `docs/primer-arranque.md` | Inglés inicial, traducción inmediata y valores editables de Mailpit |
| `docs/interfaz-y-experiencia.md` | Comportamiento del selector de idioma durante `/setup` |
| `README.md` | Quitar el segundo comando obligatorio y explicar que Mailpit ya aparece precargado |
| `docs/README.md` y `AGENTS.md` | Estado y recuento de las enmiendas afectadas |

El código implementado se limita a `dev.yml`, la carga de configuración del backend y el servicio de
primer arranque, el valor inicial de la migración, `setup-page.*` y sus pruebas. No cambia los
seeders, los métodos de acceso, el correo de producción, el flujo posterior al asistente ni el
despliegue de producción.

### 12.6 Criterios de aceptación

1. Con un volumen nuevo, `docker compose -f dev.yml up -d --build` aplica el esquema y después
   arranca `backend`; sus logs no contienen consultas a tablas inexistentes.
2. Si la migración falla, `backend` no arranca. Repetirla sobre una base preparada conserva datos,
   ajustes y sello. No se cargan ejemplos.
3. Una base nueva abre `/setup` en inglés y muestra English seleccionado. Elegir Español traduce en
   el acto toda la pantalla, cambia `html[lang]` y se conserva al avanzar, recargar y volver.
4. Una instalación o cuenta existente conserva su idioma. Reanudar un asistente empezado usa el
   idioma que ya se guardó.
5. En desarrollo, el paso 4 muestra los seis valores de Mailpit de la tabla; todos se pueden editar,
   probar y guardar. Una edición reaparece al recargar.
6. En producción, y en cualquier entorno sin valores iniciales, el correo empieza vacío. Los valores
   iniciales nunca sustituyen una configuración guardada ni se usan directamente para enviar.
7. El README se comprueba desde una base nueva. Las pruebas de backend, frontend y el recorrido de
   `/setup` en PC y móvil pasan dentro de contenedores.

### 12.7 Decisiones registradas antes de escribir

El responsable confirmó el 2026-10-04:

1. El esquema se aplica automáticamente al levantar desarrollo y `backend` espera su terminación.
2. `/setup` empieza enteramente en inglés; cambiar el selector traduce toda la aplicación de
   inmediato y recuerda la elección.
3. Desarrollo precarga todos los valores locales de Mailpit. Son valores por defecto y el usuario
   puede modificarlos.

### 12.8 Registro del repaso

1. **Fallo de la migración:** el responsable confirmó que `backend` debe permanecer detenido hasta
   que `migrate` termine correctamente. `database`, `frontend` y Mailpit pueden quedar levantados
   para diagnosticar el fallo y repetir la migración; no se presenta el backend como saludable.
2. **Idioma al reanudar:** el responsable confirmó que, después de guardar el primer paso, manda el
   idioma guardado en esa instalación aunque el navegador recuerde otro. `/setup` lo aplica a toda
   la interfaz en cuanto carga el estado.
3. **Correo al reanudar:** el responsable confirmó que los valores iniciales de Mailpit sólo llenan
   un correo todavía vacío. Después de guardar el paso 4, siempre mandan los valores persistidos por
   el usuario y los valores iniciales de desarrollo no vuelven a restaurarse.

El repaso no deja decisiones de diseño abiertas. El responsable aprobó explícitamente la enmienda
con «Lo apruebo». Implementada y verificada, esta sección y el documento quedan `as-built`.

### 12.9 Verificación de la implementación

Verificado el 2026-10-04 en Linux:

- Una copia aislada de `dev.yml`, con volumen nuevo, ejecutó `migrate`, esperó su salida correcta y
  arrancó después el backend. Sus logs no contienen `42P01`: leyó las dos tablas y quedó escuchando.
- `GET /api/setup` devolvió `language: "en"` y los valores de Mailpit aprobados. La base tenía cero
  usuarios y cero tickets; `installed_at` seguía nulo. Repetir `migrate` terminó correctamente,
  conservó esos datos y dejó el valor por defecto de la columna en `'en'`.
- `go test ./...` pasó en todos los paquetes; `go vet ./...` y `gofmt -d` no informaron problemas.
- El frontend pasó sus 220 pruebas en 21 ficheros.
- La instalación desechable de Playwright recorrió `/setup` en PC y móvil antes de sellar. Comprobó
  inglés aunque el navegador pidiera español, traducción inmediata al elegir Español y prioridad
  del idioma guardado al reanudar. La suite sin servicios opcionales terminó con 175 casos en verde
  y 39 omitidos.

## 13. Backend escribible fuera del código en la suite aislada

### 13.1 Hallazgo reproducido

Con una instancia nueva, `docker compose -f tests.yml run --rm e2e` se detiene antes de Playwright:

```text
error mounting ... backend_tmp ... to /app/tmp:
create mountpoint for /app/tmp: read-only file system
```

`tests.yml` monta el código como `./backend:/app:ro` y después intenta montar un volumen hijo en
`/app/tmp`. En este runtime OCI el segundo montaje necesita crear su punto dentro de `/app`, que ya
es de sólo lectura. La suite no alcanza el backend ni el navegador.

El volumen existe sólo para Air: `.air.toml` compila `./tmp/server`. La suite aislada no necesita
recarga en caliente porque el código no cambia durante una pasada.

### 13.2 Solución propuesta

El servicio `backend` de `tests.yml`:

- conserva `./backend:/app:ro`;
- define `command: ["go", "run", "."]`, de modo que compila y ejecuta una vez sin Air;
- elimina `backend_tmp:/app/tmp` y la declaración `backend_tmp` de los volúmenes;
- conserva `backend_go_mod` y `backend_go_build`, que guardan dependencias y caché fuera de `/app`;
- conserva los volúmenes de logs, adjuntos y los demás servicios de la suite.

`go run` escribe su binario temporal fuera del árbol de código. El backend sigue esperando la
migración completada y el runner sigue esperando su salud antes de instalar y probar.

### 13.3 Alcance

Cambian `tests.yml` y la descripción de la suite en `docs/prueba-local.md`, `docs/ambientes.md`,
`docs/arquitectura.md`, `docs/README.md` y `AGENTS.md`.

Quedan fuera `dev.yml`, `.air.toml`, las imágenes, producción, el código Go, las pruebas y sus datos.
Desarrollo continúa usando Air para recargar; la suite aislada arranca el backend una sola vez.

### 13.4 Criterios de aceptación

1. `docker compose -f tests.yml down -v` seguido de la suite crea el backend sin errores de montaje.
2. El bind `/app` permanece de sólo lectura y no existe un volumen con destino `/app/tmp`.
3. El backend espera `migrate`, compila con `go run .`, conecta con la base y responde salud.
4. La preparación recorre la instalación vacía y Playwright comienza sus casos.
5. Al terminar, `down -v` elimina todos los volúmenes de la instancia, sin dejar `backend_tmp`.

### 13.5 Decisión registrada y registro del repaso

El responsable confirmó el 2026-10-06 proponer `go run .`, conservar el código de sólo lectura y
retirar `backend_tmp`, sin cambiar desarrollo ni producción.

En el repaso del mismo día confirmó que:

1. la suite no necesita recarga con Air: compila y arranca una sola vez;
2. `backend_go_mod` y `backend_go_build` son las únicas cachés de Go que se conservan;
3. `down -v` no debe dejar ningún recurso adicional.

El repaso queda cerrado sin decisiones abiertas. El responsable aprobó explícitamente los tres
documentos con «Apruebo» el 2026-10-06; esa aprobación habilitó modificar `tests.yml`.

### 13.6 Verificación de la implementación

Verificado el 2026-10-06 en Linux, directamente con el `tests.yml` definitivo y sin archivos de
sobreescritura:

- `docker compose -f tests.yml --profile directory down -v` dejó una instancia limpia y
  `docker compose -f tests.yml up -d --build database backend frontend mail` creó y arrancó el
  backend sin el error OCI de `/app/tmp`;
- la configuración resuelta conserva `/app` como sólo lectura, usa `command: [go, run, .]` y no
  declara ningún montaje ni volumen `backend_tmp`;
- la migración terminó antes del backend, el backend conectó con PostgreSQL y el runner alcanzó su
  comprobación de salud, preparó una instalación vacía y comenzó Playwright en PC y móvil;
- `docker compose -f tests.yml run --rm e2e` terminó con salida 0: 177 casos aprobados y 39 omitidos
  por los servicios y perfiles opcionales ausentes, en 13,3 minutos;
- el `down -v` final eliminó los contenedores, la red y todos los volúmenes de la instancia; un
  `docker compose -f tests.yml ps -a` posterior quedó vacío.

La enmienda y el documento quedan `as-built`.

## 14. Recorrido manual aislado del catálogo real

### 14.1 Preparación y evidencia

El recorrido parte de una instalación vacía y recursos Compose con nombres exclusivos. Antes de
descargar se anotan memoria y disco disponibles y se contrastan los metadatos del catálogo con la
fuente oficial. El informe incluye los comandos exactos, versiones de Docker e imagen, tiempos,
memoria, checksums y resultados del navegador; no incluye credenciales ni binarios.

### 14.2 Secuencia

1. Crear el entorno desechable, aplicar la migración y levantar el administrador local vacío.
2. Recorrer `/setup` en PC y móvil hasta el quinto paso.
3. Descargar, activar y probar 1.5B; sellar la instalación y verificar una generación en español y
   otra en inglés.
4. Abrir `/settings`, descargar y activar 3B y repetir tres generaciones por idioma.
5. Terminar sólo `llama-server` con cada modelo activo y medir la recuperación, hasta cinco minutos.
6. Descargar y verificar 7B sin activarlo; comprobar que la falta de RAM se explica en la interfaz.
7. Volver a 1.5B, borrar 3B y 7B desde la aplicación cuando proceda y ejecutar la limpieza total del
   proyecto desechable.

### 14.3 Criterios de aceptación

- Los tres archivos coinciden con los metadatos oficiales y su checksum; una discrepancia detiene
  ese modelo y abre una enmienda.
- 1.5B y 3B quedan activos y sanos, generan tres respuestas válidas por entrada y recuperan salud en
  cinco minutos o menos después de terminar el hijo.
- Cada respuesta cumple el contrato objetivo de JSON, campos, idioma y longitud.
- `/setup` y `/settings` comunican descarga, activación, salud, recursos y errores en PC y móvil.
- El 7B se valida sin forzar swap ni intentar activarlo con memoria insuficiente.
- La limpieza no deja contenedores, redes, volúmenes, bases ni modelos del entorno temporal.

### 14.4 Alcance y decisiones

Este bloque valida lo ya construido. No prueba GPU, servicios externos ni producción y no conserva
modelos reales. Cualquier defecto que requiera código vuelve primero al ciclo documental dentro del
bloque 5. El responsable eligió 1A–4A y 5A–9A; el repaso eligió 10A–14A. No quedan decisiones
abiertas.

### 14.5 Resultado

El recorrido terminó con los tres checksums oficiales, 1.5B y 3B sanos y 24 respuestas válidas. La
recuperación tardó aproximadamente 4 s y 9 s respectivamente. El rechazo seguro del 7B y su toast se
comprobaron en PC y móvil. Se usaron exclusivamente los proyectos
`catalina-support-block5` y `catalina-support-ai-block5`; al cerrar se eliminaron también su volumen
de modelos, base, cachés, archivos y red.

## 15. Recorrido de formato y accesibilidad

1. Levantar una instalación vacía y recorrer `/setup` en español e inglés, claro y oscuro, en PC,
   tableta y móvil.
2. Reiniciar la instancia desechable, aplicar el seeder opcional y entrar con los papeles necesarios
   para alcanzar sesión, armazón, configuración, correo, usuarios y todas las vistas de tickets.
3. Registrar cada hallazgo con prioridad, contexto, reproducción y una medida observable; completar
   el inventario antes de corregir.
4. Corregir primero en el componente compartido cuando el patrón se repita y ejecutar las pruebas de
   todos sus consumidores.
5. Repetir ambas lenguas, los dos temas de fábrica y las tres resoluciones sobre cada pantalla
   afectada; ejecutar las pruebas comunes de contraste para los ocho temas.
6. Cerrar sólo cuando no quede ningún hallazgo abierto, pasen las pruebas y la limpieza deje vacío el
   proyecto temporal.

El recorrido aplica WCAG 2.2 AA donde corresponde: teclado y foco, estructura semántica, nombre y
estado accesibles, contraste y objetivos táctiles. No conserva capturas base, no añade navegadores,
no altera flujos ni toca producción.

### 15.1 Resultado

El recorrido de `/setup` completó 48 combinaciones. Tras aplicar el seeder se revisaron, por cada
idioma, 150 vistas autenticadas de los cuatro papeles y 18 vistas públicas. El único hallazgo real
fue el desbordamiento del filtro de papeles de `/users` en inglés y móvil; la corrección compartida
pasó PC, tableta y móvil, claro y oscuro, con Administrador y Soporte.

Pasaron 220 pruebas unitarias, 18 casos afectados con 4 omisiones previstas y la suite canónica con
180 casos aprobados y 38 omisiones previstas. No quedó ningún hallazgo abierto. La instancia
`catalina-support-block6` se eliminó con sus contenedores, red y volúmenes; desarrollo y producción
no participaron en el recorrido.

## 16. Propuesta del bloque 8: incorporación desde el repositorio público

### 16.1 Objetivo y fuente de verdad

El recorrido representa a una persona que llega por primera vez al repositorio público. Parte de un
clon nuevo de la rama predeterminada `main` y usa únicamente los comandos, direcciones, cuentas y
explicaciones del README que contiene ese clon. No usa este documento, conocimiento interno, el
árbol de `dev` ni comandos improvisados para completar un paso.

Una instrucción ausente, incorrecta o que exija conocimiento no publicado es un hallazgo, aunque el
equipo sepa cómo continuar. Se puede investigar después de registrar el punto exacto, pero no se
cuenta el recorrido como aprobado mediante ese atajo.

La ejecución real se hace en Linux con Docker y Docker Compose. Windows y macOS no se simulan y no
se presentan como verificados. La dirección principal es `http://127.0.0.1:11001`; nginx, TLS y el
dominio local no forman parte del camino de primera prueba.

### 16.2 Preparación limpia y datos desechables

Los datos de desarrollo actuales fueron declarados desechables por el responsable (15B). Antes del
clon se detienen los compose de desarrollo, IA, directorio y Keycloak y se eliminan sus contenedores,
redes y volúmenes, incluido el volumen de modelos. También se limpian `_files` y `_logs` de
desarrollo. No se crea una copia de respaldo de esos recursos.

Se eliminan las imágenes construidas específicamente para Catalina Support, para que `--build`
demuestre una construcción nueva. No se purgan imágenes base compartidas, cachés globales de Docker
ni recursos ajenos al proyecto. El repositorio de trabajo y su historial no se borran ni se usan
como fuente del clon público.

El clon vive en una carpeta temporal fuera del repositorio de trabajo. Al terminar se eliminan sus
contenedores, redes, volúmenes, modelo, adjuntos, logs, imágenes propias y la carpeta temporal. Un
inventario final debe demostrar que no quedan recursos de Catalina Support creados por el recorrido.
Producción no se detiene, modifica ni consulta.

### 16.3 Recorrido vacío con correo e IA reales

Se sigue el README desde el clon: comprobación de requisitos, construcción y arranque. Una base
nueva debe llevar a `/setup` sin seeders. Se completan sus cinco pasos con entrada local, región y
dirección local, Mailpit y el modelo local predeterminado.

El modelo se descarga desde la interfaz, se activa y se prueba de verdad. El criterio no compara una
frase exacta: exige que el motor quede sano, que Catalina Support obtenga un resumen válido y que
«Mejorar con IA» devuelva un borrador revisable. Una respuesta vacía, un error oculto o una conexión
que sólo pasa salud no cumplen.

Tras sellar la instalación se entra con `admin` / `admin`, se crea una cuenta local usando el flujo
documentado, se abre su correo real en Mailpit y se establece su contraseña. Con esa cuenta se crea
un ticket; Soporte o Desarrollo ejercita el resumen y la mejora de redacción según sus permisos.
El recorrido comprueba que `/setup` ya no se puede reabrir.

### 16.4 Seeders y métodos de entrada opcionales

Después del recorrido vacío se ejecuta exactamente el comando opcional del README. Debe terminar
correctamente, conservar sello, correo, IA y los demás ajustes de la instalación, y dejar once
cuentas de ejemplo, veinticinco tickets y cinco adjuntos utilizables. Se comprueba una entrada de
Usuario, Soporte y Desarrollo con las credenciales publicadas.

AD y Keycloak se levantan después, uno por uno, con sus comandos documentados. Desde Configuración
se prueba y activa AD, entra una persona del directorio y se confirma la creación o actualización de
su cuenta al primer acceso. Después se prueba y activa Keycloak y entra una persona del reino. La
cuenta de fábrica debe poder entrar durante ambos métodos. Al terminar se restaura el método local y
se comprueba de nuevo una cuenta local.

### 16.5 Tiempos y evidencia

El informe separa:

- **tiempo activo**, dedicado a leer, escribir, elegir opciones y corregir datos;
- **esperas automáticas**, construcción de imágenes, arranque, descarga del modelo y generación.

Se anotan sistema, arquitectura, versiones de Docker y Compose, commit público probado, comandos
copiados del README, resultado de cada hito y consumo final de disco del modelo. No se registran
tokens, contraseñas nuevas, cuerpos privados, cabeceras ni contenido completo de tickets. Las
capturas o logs temporales se conservan sólo mientras sirvan para diagnosticar un hallazgo y se
eliminan con el entorno.

No hay un límite de minutos que convierta por sí solo el recorrido en fallo: los tiempos sirven para
detectar esperas mal explicadas y comparar mejoras futuras. Un proceso detenido sin progreso o sin
mensaje sí es un hallazgo.

### 16.6 Criterios de cierre y tratamiento de hallazgos

El bloque se cierra sólo si:

1. un clon de `main` completa el camino sin instrucciones externas al README;
2. la instalación comienza vacía y `/setup` termina y queda sellado;
3. Mailpit completa el alta local y el modelo real ejecuta resumen y mejora de redacción;
4. el seeder conserva la configuración y deja 11 cuentas, 25 tickets y 5 adjuntos;
5. local, AD y Keycloak permiten sus entradas y al final vuelve a funcionar local;
6. los tiempos quedan separados y la limpieza final no deja recursos del recorrido;
7. no queda ningún hallazgo abierto.

Una errata o instrucción documental se corrige en el documento que la contiene y se repite desde el
punto anterior estable. Un hallazgo de comportamiento, interfaz, API, datos, permisos o flujo se
reporta y vuelve al ciclo de propuesta y aprobación antes de cambiar código. No se agrupan arreglos
no aprobados dentro de esta auditoría.

### 16.7 Decisiones y exclusiones

El responsable eligió 1A–5A: instalación pública, recursos nuevos, vacío antes de ejemplos, modelo
local real y métodos opcionales después de local. Eligió 6A–10A: `main`, sólo README, recorrido
funcional completo, tiempos separados y tratamiento documental estricto de hallazgos. En el repaso
eligió 11A–14A y 15B: puerto 11001, generación real sin texto fijo, conteos y conservación tras el
seeder, AD seguido de Keycloak y datos actuales desechables sin respaldo. No quedan decisiones
abiertas.

Quedan fuera producción, nginx, TLS, dominio, etiqueta `v1.0.0`, proveedores externos, costos,
GPU, rendimiento comparativo, Windows, macOS y cambios funcionales no aprobados.

## 17. Propuesta de reanudación después del hallazgo de idioma

La primera ejecución pública probó el commit
`afeae10190a87b0df92941f53d356bfc013453f4` sobre Docker 29.5.2 y Compose 5.1.4. El clon tardó
0,95 s, el entorno se construyó en 40,14 s y el administrador de IA en 4,14 s. `/setup` comenzó en
inglés, Mailpit apareció como `mail:1025` y su prueba respondió 200. El modelo predeterminado mostró
licencia, checksum y requisitos; descargó 1,04 GiB en 55,98 s, ocupó 1,1 GB, quedó sano y completó
la generación del asistente en 3,91 s.

El recorrido selló la instalación, entró con la cuenta de fábrica, creó por correos reales de
Mailpit una cuenta de Usuario y otra de Soporte y abrió `CS-2026-0001` desde la interfaz. Los dos
resúmenes quedaron listos en 16,37 s. «Última acción» salió en inglés; «Motivo» salió en español aun
cuando el contrato de API decía `language: "en"`. La consulta repetida del detalle también produjo
logs GORM `record not found` para el interno inexistente. Conforme a 16.6 se detuvo antes de seeders,
AD y Keycloak y no se modificó código.

El entorno se limpió por completo: contenedores, redes, volúmenes, modelo, imágenes construidas,
clon y herramientas temporales. El inventario final conservó únicamente recursos previos, incluidas
las imágenes y contenedores detenidos de producción, que no se tocaron.

Tras implementar y verificar `docs/modules/ai.md`, sección 15, y `docs/modules/tickets.md`, sección
14, no se continúa desde un punto intermedio. Se publica primero la corrección en `main` y se repite
todo el bloque 8 desde una clonación y recursos nuevos: instalación vacía, IA local real, alta por
Mailpit, resúmenes y redacción, seeders, AD, Keycloak, retorno a local, tiempos y limpieza. El cierre
sigue exigiendo cero hallazgos abiertos.

## 18. Propuesta de reanudación después del hallazgo de disponibilidad TCP

La segunda ejecución pública probó el commit `80e01a75c7bc3b284778bcd214850efb0690714b` con Docker 29.5.2 y Compose 5.1.4. El clon terminó en 2,21 s. El primer `up -d --build` tardó 63,87 s y falló antes del backend: `database` figuró saludable, pero `migrate` recibió `connection refused` al conectar por TCP a `database:11003`. No se inició IA ni se alcanzó `/setup`.

El fallo se atribuye al healthcheck por socket durante la transición entre el servidor temporal de inicialización de PostgreSQL y el servidor definitivo. Se detuvo conforme a 16.6 y se retiraron contenedores, redes, volúmenes, imágenes propias y clon; producción no se consultó ni modificó.

Después de implementar y verificar las secciones nuevas de `docs/ambientes.md` y `docs/arquitectura.md`, la corrección se publicará primero en `main`. Entonces se repetirá nuevamente todo el bloque 8 desde recursos y clon nuevos; no se continúa desde el fallo. Además de los criterios existentes, la pasada comprobará que una base nueva no presenta la carrera, que desarrollo y pruebas migran automáticamente, y que la salud distingue en una comprobación aislada el esquema pendiente de una base inaccesible.

Antes de la repetición pública pasan las pruebas automatizadas de los tres estados de salud y una
ejecución desechable del servicio manual de producción. En el repaso el responsable eligió 14A–18A:
publicar y montar las migraciones, permitir que el comando manual levante la base, devolver JSON
estable, cubrir los tres estados con pruebas y explicar el orden también en el README. Se conservan
6B, 7A, 8A y 9A, junto con 10A–13A. No quedan decisiones abiertas; falta aprobar explícitamente la
propuesta antes de modificar código o configuración.

## 19. Reanudación aprobada después del hallazgo de logs

La tercera ejecución pública probó el commit
`2f7c6468555fdf23efec08716d42755ccefa8b31` con Docker 29.5.2 y Compose 5.1.4. El clon tardó 1,30 s,
el entorno 47,87 s y el administrador de IA 4,15 s. Desarrollo migró automáticamente sin errores de
tablas; `/setup` comenzó vacío, en inglés y con Mailpit precargado. El modelo 1.5B descargó
1.117.320.736 bytes en 49,25 s, quedó sano en 14,62 s y completó la prueba del asistente en 3,04 s.

Se selló la instalación, se completaron altas locales mediante correos reales, se creó
`CS-2026-0001`, la mejora de redacción devolvió texto y los dos resúmenes quedaron listos con
`language=en`. El seeder conservó configuración e IA y dejó 11 cuentas, 25 tickets y 5 adjuntos.
Entraron cuentas de Usuario, Soporte y Desarrollo; después funcionaron AD con creación automática,
Keycloak mediante su vuelta OIDC real, la cuenta de fábrica y el retorno final al método local.

La inspección final encontró `record not found` con nivel de error al buscar la categoría
normalizada antes de crearla y al buscar por correo e identidad externa antes de crear cuentas de
AD y Keycloak. Los resultados funcionales eran correctos, pero contradicen el requisito de una
operación sin ruido. Se detuvo el cierre y se limpió por completo el entorno, incluido el modelo de
1,1 GB; producción permaneció intacta.

Tras aprobar, implementar, probar e integrar las secciones propuestas de arquitectura, usuarios y
tickets, se repetirá todo el bloque 8 desde otro clon público de `main`. Además de los siete
criterios de 16.6, el examen de logs debe mostrar cero falsos errores por ausencias esperadas y las
pruebas deben demostrar que un error SQL real conserva su nivel. El responsable eligió 1A–4A; no
quedan decisiones abiertas.

## 20. Cuarta pasada pública completada

La repetición final probó desde un clon limpio el commit público de `main`
`b0898922f6d584d9b3349f2de82adc6046e82315`, con Docker 29.5.2 y Compose 5.1.4. El comando exacto
del README, `docker compose -f dev.yml up -d --build`, construyó y levantó una base nueva en
29,13 s. La migración automática terminó correctamente y la salud respondió con base disponible.
`/setup` comenzó sin sello, en inglés, con entrada local y los valores editables de Mailpit.

El administrador de IA levantó en 3,36 s. Desde el asistente se descargó el modelo predeterminado
`qwen2.5-1.5b-instruct`: 1.117.320.736 bytes en 42,8 s, con checksum válido. La activación tardó
5,06 s y la generación obligatoria del asistente, 2,49 s. La instalación quedó sellada con inglés,
`Europe/Berlin`, dirección local, Mailpit y el motor local sano.

La cuenta de fábrica entró y creó por el flujo normal una cuenta de Usuario y una de Soporte. Los
dos correos de invitación llegaron a Mailpit, sus enlaces permitieron establecer las contraseñas y
ambas cuentas iniciaron sesión. Se creó `CS-2026-0001`; «Mejorar con IA» devolvió un borrador en
12,22 s y los resúmenes «Motivo» y «Última acción» llegaron a `listo` en inglés.

El seeder opcional se ejecutó exactamente como indica el README y tardó 4,69 s. Conservó el sello,
el idioma, la región, la dirección, el correo y la IA; dejó sus 11 cuentas, 25 tickets y 5 adjuntos.
Entraron las cuentas publicadas de Usuario, Soporte y Desarrollo. Las dos cuentas creadas antes del
seeder también se conservaron.

Después se levantaron los servicios opcionales. AD respondió a su prueba, se activó y autenticó a
una persona del directorio, que se aprovisionó al entrar. Keycloak importó el reino, respondió a su
prueba y completó la vuelta OIDC real con el token en el fragmento; su persona también se
aprovisionó. La cuenta de fábrica entró durante ambos métodos. Por último se restauró `local` y una
cuenta local volvió a entrar. El conteo final fue de 15 cuentas —11 del seeder, 2 locales creadas en
el recorrido y 2 aprovisionadas por AD y Keycloak—, 25 tickets y 5 adjuntos.

La inspección completa de los logs del backend no encontró `ERROR:`, `record not found`, nivel
`[error]`, pánicos ni errores fatales. La interfaz se abrió además con Chromium en un viewport móvil:
la instalación sellada redirigió `/setup` a la entrada, conservó `lang=en` y la cuenta local llegó a
`/tickets`. No se editaron archivos dentro del clon.

Al terminar se bajaron `keycloak.yml`, `active-directory.yml`, `ai.yml` y `dev.yml` con sus
volúmenes. Se retiraron la base, el modelo de 1,1 GB, cachés, redes, contenedores, clon y herramientas
temporales del recorrido. Los contenedores detenidos de producción se conservaron sin iniciarlos ni
modificarlos. Con cero hallazgos abiertos, el bloque 8 queda cerrado.

## 21. Recorrido de prioridades

Sobre una instalación desechable se crea un ticket sin cambiar la selección y se comprueba `medium`.
Después se crean o modifican casos de los cuatro niveles, se verifica que principal e interno lean
el mismo valor y que los permisos coincidan con la matriz. Las bandejas se prueban en PC y móvil con
filtro múltiple, URL recargable y orden crítica→baja con antigüedad como desempate.

El recorrido marca un ticket como crítico y confirma un único correo a todo Soporte activo, con los
campos mínimos y el idioma global. Luego baja y vuelve a subir la prioridad y comprueba que no haya
otro correo. Finalmente interrumpe SMTP, cambia otro ticket a crítico y confirma prioridad e
historial guardados junto al aviso visible de que el correo falló. El seeder debe mostrar los cuatro
niveles con una distribución realista.

El recorrido automatizado del 2026-10-14 cubre en PC y móvil la corrección de prioridad, la
persistencia de filtros y orden en la dirección, el historial, la lectura sin edición del Usuario y
el único correo en Mailpit. El fallo de SMTP se comprueba en backend contra PostgreSQL real: la
prioridad y su historial permanecen guardados y la respuesta lleva el aviso que presenta el toast.
Dos aplicaciones consecutivas del seeder dejaron la misma distribución: 2 críticos, 6 altos, 11
medios y 6 bajos.

## 22. Verificación de la nueva portada

Después de reescribir ambos README se comprobará que una persona puede seguir el camino principal
en inglés y encontrar el equivalente completo en español; los comandos, puertos, credenciales y
advertencias deben coincidir. Se verificarán todos los enlaces relativos y los enlaces públicos al
sitio, la demo y Ko-fi, sin automatizar accesos ni pagos externos.

Para la imagen pública se ejecutará conscientemente el seeder sobre el entorno local, aceptando su
limpieza documentada de tickets, categorías, etiquetas, cuentas `@demo.com` y adjuntos de
desarrollo. Antes de capturar se confirmarán marca de fábrica, tema con color institucional, vista
doble, `Summary`, `Last action` y ausencia de información personal. La evidencia no mostrará la
cuenta de fábrica, secretos, correos privados ni identificadores de infraestructura.

La validación no tocó producción, `demo.catalina-support.com`, DNS, el sitio, Ko-fi ni el script
externo de restauración. Todos los enlaces relativos existen, `git diff --check` terminó limpio y
los Compose de desarrollo e IA son válidos. `prod.yml` confirmó su protección esperada: no puede
evaluarse sin el archivo privado `config/env/prod.env`, requisito que el README explica en vez de
ocultar.

La interfaz comprobó en 1920 × 1200 el ticket principal y el interno completos, y en el detalle del
mismo ticket mostró resultados reales en `Summary` y `Last action` junto a `Improve with AI`. La
vista doble no renderiza ese panel lateral; se registró el desajuste respecto de la propuesta y se
resolvió con dos capturas reales complementarias, sin componer ni maquillar la UI.

## 23. Dashboard y reportería

Una instalación desechable generará tickets principales e internos en límites de periodo y zona,
con respuestas, resoluciones, reaperturas, cierres, prioridades y responsables variados. Las pruebas
contrastarán los agregados con SQL conocido, el periodo anterior, los umbrales configurables y los
valores pendientes. Recorrerán Soporte, Desarrollo, Administrador combinado, fábrica y Usuario
rechazado.

En PC y móvil comprobarán filtros persistentes en URL, gráficos accesibles, actualización manual,
tabla progresiva, apertura de ticket y CSV completo con duraciones humana/minutos y celdas
neutralizadas. Producción, dominio, correos y etiqueta `v1.0.0` quedan fuera.

La verificación ejecutada usó PostgreSQL local y los usuarios del seeder. Soporte vio ambos tipos,
Desarrollo sólo internos, Usuario recibió 403 y Administrador accedió a Dashboard y Configuración.
Se comprobó un filtro recargable, apertura de fila, CSV completo, vista móvil y restauración de los
umbrales. Las pruebas automatizadas cubren servicio, repositorio real, contrato HTTP frontend,
fronteras y compilación. No se tocó producción ni el dominio.

## 24. Aceleración del motor local

La instalación desechable verifica migración y valor inicial Auto, ejecución CPU con el Compose
normal, atribución de dispositivo, saludo, reinicio y persistencia. Simula capacidad NVIDIA para
probar argumentos, ausencia, respaldo de Auto, rechazo explícito, bloqueo concurrente y rollback.
La suite portable pasó en CPU. El manifiesto CUDA y el overlay se resolvieron correctamente, pero
la descarga de la imagen NVIDIA se detuvo tras validar el manifiesto (sus capas rondan 2 GB) porque
este host no tiene Toolkit ni GPU; la pasada real queda para un servidor Linux compatible.

PC y móvil recorren `/setup` y `/settings` en ambos idiomas. Se inspecciona que API, toast y logs no
expongan UUID, bus, memoria ni trazas. Al terminar se restaura Auto, se baja la variante opcional y
se conserva el modelo compartido sólo durante el entorno de prueba. Producción, dominio y etiqueta
siguen fuera.

## 25. Verificación de los modos SMTP

Las pruebas unitarias y de integración cubren: STARTTLS correcto y rechazo cuando no se anuncia;
TLS directo correcto; `plain` rechazado sin aceptación; aceptación anulada al abandonar el modo;
rechazo si una sola resolución es pública o IPv6; admisión de `10/8`, `172.16/12`, `192.168/16` y
`127/8`; y conexión contra la dirección ya validada sin segunda resolución.

La migración se prueba con `smtp_secure=true` y `false`, esperando respectivamente `tls_direct` y
`starttls`. El entorno de desarrollo verifica aparte que Mailpit queda `plain`, sin usuario ni
contraseña y con aceptación local explícita. Ninguna prueba autoriza esa excepción en producción.

El recorrido del paso 4 comprueba los tres textos en español e inglés, la advertencia, la
confirmación persistida mientras el modo siga activo, su borrado al cambiar y los toasts. La prueba
SMTP continúa sin `MAIL FROM`, `RCPT TO` ni `DATA`. Una comprobación con relay falso reproduce
`AUTH LOGIN` sin STARTTLS: STARTTLS debe rechazarlo y `plain` sólo debe aceptarlo con confirmación y
destino privado. Usuario, contraseña y cuerpo no aparecen en API, logs ni capturas.

La ejecución terminó con las regresiones focalizadas de `mail` y `settings` aprobadas, 247 pruebas
Angular aprobadas, fronteras y compilación de producción correctas. La suite backend completa
conserva un fallo ajeno y dependiente de datos sembrados en el repositorio de reportería; los
paquetes afectados por SMTP pasaron. La prueba manual real autenticó contra el relay y confirmó en
el toast que no se envió ningún correo.

## 26. Verificación del correo real de prueba

Las pruebas unitarias atraviesan la transacción SMTP completa (`MAIL FROM`, `RCPT TO` y `DATA`) y
comprueban que el remitente y el destinatario son exactamente los escritos. Cubren texto fijo en
español e inglés, secreto guardado cuando el campo está vacío, secreto temporal sin persistencia,
rechazo del `From:`, rechazo del destinatario y fallo durante `DATA`.

El endpoint público admite cinco intentos para una pareja IP–destinatario dentro de quince minutos,
rechaza el sexto con `429`, separa otra IP u otro destinatario y vuelve a admitir al vencer la
ventana. También se comprueban el `409` tras el sello y que una dirección inválida no inicia SMTP.
La ruta de Configuración exige capacidad administrativa y no hereda el carácter público de `/setup`.

En PC y móvil se verifica un único botón, el destino no persistido, el carácter opcional y el toast
**El servidor aceptó el correo para {destinatario}**. Mailpit permite confirmar asunto, cuerpo,
idioma y cabeceras sin salir del entorno. La comprobación manual contra el relay institucional debe
usar un buzón controlado y registrar aceptación y recepción como hechos distintos. Los logs se revisan
para asegurar que sólo contienen destinatario parcialmente oculto y ningún secreto o cuerpo.

## 27. Verificación de actualizaciones y migraciones

La implementación se verificará en instalaciones desechables, sin tocar producción, el dominio ni
la base habitual. La matriz mínima cubre:

1. Base nueva: crea el respaldo sólo al detectar trabajo, aplica `v1.0.0.sql`, registra checksum y
   fecha, y arranca backend y frontend después del ejecutor.
2. Base existente sin historial: reaplica idempotentemente `v1.0.0.sql`, conserva datos y registra
   la versión sin duplicados.
3. Reinicio sin pendientes: no ejecuta SQL, no crea respaldo nuevo y arranca normalmente.
4. Actualización con dos versiones pendientes: las ordena por SemVer, crea un único respaldo y
   confirma cada transacción en orden.
5. Error SQL: revierte íntegramente sólo esa versión, no la registra y deja la aplicación detenida
   con un mensaje que identifica el archivo sin exponer secretos.
6. Respaldo inválido o imposible: no inicia ninguna migración ni aplicación.
7. Dos ejecutores concurrentes: el segundo espera el advisory lock, vuelve a leer el historial y no
   repite trabajo.
8. Checksum cambiado: la versión actual se reaplica únicamente con `Development=true`; una versión
   publicada o histórica bloquea el arranque.
9. Archivo futuro, nombre inválido, versión duplicada o base más nueva que el código: todos se
   rechazan antes de modificar el esquema.
10. Archivo `*_dev.sql`: nunca lo descubre el ejecutor; continúa aplicándose sólo por el seeder.
11. `update.sh`: exige `main`, árbol limpio y avance rápido; un fallo de descarga o compilación deja
   sirviendo la instalación anterior y no modifica el despliegue activo.
12. Recuperación: una migración compatible permite levantar temporalmente la versión anterior; no
   se ejecuta restauración automática. Se valida aparte el procedimiento manual con una copia
   desechable.

También se comprueban retención de catorce días, nombre del respaldo con versiones, logs completos
sin DSN ni credenciales, archivos SQL sin transacciones propias, compilación del comando Go,
configuración resuelta de los tres Compose y las suites existentes. README, runbook y arquitectura
quedan actualizados a `as-built`.

Se verificaron los casos de base nueva, adopción existente, segundo arranque, rollback, checksum,
nombres estrictos, exclusión `_dev` y rechazo de `update.sh` fuera de `main`. La exclusión mutua
queda implementada con un advisory lock sobre una conexión exclusiva; el recorrido con varias
versiones se ejercerá cuando el responsable abra la siguiente. La suite backend conserva únicamente
el fallo ambiental conocido del informe sin tickets sembrados.
