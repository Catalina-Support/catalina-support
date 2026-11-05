# Informe de QA — Catalina Support 1.0.1

> **Estado:** as-built
> **Última actualización:** 2026-11-05
> **Fecha de la pasada:** 2026-10-31, actualizada el 2026-11-05
> **Entorno:** desarrollo local (`docker compose -f dev.yml`), datos de ejemplo `seed-es`, Docker Desktop WSL2
> **Interfaz:** Chrome 156 visible en el host, controlado con Playwright desde `tmp/qa`
> **URL:** `http://127.0.0.1:11001`
> **Cuentas:** `admin/admin`, `user1@demo.com`, `support1@demo.com`, `dev1@demo.com` (contraseña de demo `123123123`)
> **Alcance:** sólo el entorno de desarrollo; no se ejecutó nada fuera de él.

## 1. Resumen ejecutivo

Se recorrieron las pantallas principales, los cuatro papeles, el flujo completo de un ticket
(creación, asignación, comentario, escalado, cierre interno y vuelta a Soporte), el alta de un
ticket, la recuperación de contraseña con el correo real del buzón de pruebas, el editor de correos,
los ocho temas, el cambio de idioma y la vista móvil.

**Los once hallazgos originales están cerrados.** Los tres primeros eran bloqueantes (bucle de
redirección, 503 de correos y etiquetas de plantillas); los demás eran de claridad, historial, menú,
seed o decisiones de producto. Después de corregirlos se repitió la suite completa en WSL2:
backend, 269 pruebas de frontend, 200 E2E aprobadas y AD/Keycloak sin fallos.

## 2. Hallazgos

| ID | Severidad | Hallazgo | Estado |
| --- | --- | --- | --- |
| QA-01 | **Crítica** | Bucle infinito de redirección con `aiRequired` y IA sin configurar | **Corregido** |
| QA-02 | **Alta** | El 503 de traducción de correos se trata como «servidor caído» y bloquea Configuración | **Corregido** |
| QA-03 | **Alta** | El editor de correos dice «diez», no etiqueta `auth.accountCreated` y conserva textos de 10 plantillas | **Corregido** |
| QA-04 | Media | El historial del usuario muestra dos veces «En curso» para el mismo movimiento | **Corregido** |
| QA-05 | Media | El chip de la cabecera del interno muestra el estado del principal | **Corregido** (PR #46) |
| QA-06 | Media | Con la IA caída, los resúmenes se quedan en «Generando…» y no hay un estado final claro | **Corregido** |
| QA-07 | Baja | El menú lateral repite «Administrador» tres veces | **Corregido** |
| QA-08 | Baja | `seed-es` deja cuentas de pruebas antiguas y la lista de Usuarios queda sucia | **Corregido** |
| QA-09 | Informativa | En inglés las fechas usaban el formato local `en-GB` | **Resuelto**: ISO `AAAA-MM-DD hh:mm` en ambos idiomas |
| QA-10 | Informativa | En móvil no existe «Los dos» | **Decidido**: se mantienen Principal e Interno |
| QA-11 | Baja | `/setup` no recuerda el paso al recargar | **Corregido el texto**; se decidió no persistir la posición |

## 3. Detalle

### QA-01 — Crítico: bucle de redirección con IA obligatoria sin configurar

**Qué se vio.** En una instalación recién sembrada con `seed-es`, el login de `admin` devuelve 200,
`/api/auth/me` devuelve 200 y la aplicación se queda en `/login` con el renderer al 100 % de CPU. No
llega a pintar el panel. Lo mismo ocurre con una cuenta que no es administradora.

**Causa.** `GET /api/setup` responde `aiRequired: true` y `ai.mode: ""`:

```json
{ "installed": true, "aiRequired": true, "aiAvailable": false, "ai": { "mode": "", "tested": false } }
```

La guarda `sesionActiva` (`frontend/src/app/core/guards/sesion.guard.ts`) redirige a
`/settings?ai=required` y compara el `state.url` completo con `/settings`:

```ts
if ((await setup.estadoDeInstalacion()).aiRequired &&
    state.url !== '/settings' && state.url !== '/forbidden') {
  return usuario?.isAdmin
    ? router.createUrlTree(['/settings'], { queryParams: { ai: 'required' } })
    : router.createUrlTree(['/forbidden'], { queryParams: { ai: 'required' } });
}
```

En la siguiente navegación el `state.url` es `/settings?ai=required`, no `/settings`, así que vuelve
a redirigir. El bucle no termina. Para una cuenta no administradora ocurre lo mismo con
`/forbidden?ai=required`.

**Reproducción.**

1. Base de desarrollo vacía y `docker compose -f dev.yml run --rm seed-es`.
2. Entrar como `admin/admin`.
3. La URL se queda en `/login`; el panel no aparece y Chrome consume CPU.

**Impacto.** Cualquier instalación sellada cuyo `ai_settings` quede vacío (o cuya credencial no se
pueda descifrar) deja fuera a todos los papeles. El seed de desarrollo lo provoca de forma
sistemática. Es el hallazgo más grave de la pasada.

**Recomendación.**

- Comparar **la ruta**, no la URL con query: `state.url.split('?')[0]`, o comprobar la URL del
  `UrlTree` de destino.
- No volver a redirigir si ya se está en la pantalla de destino con el mismo motivo.
- El seed debe configurar una IA válida antes de sellar, o sellar sin `aiRequired`; hoy sella una
  demo que no se puede usar.
- Añadir una prueba de regresión: instalación sellada + `aiRequired` + login de admin y de usuario,
  esperando una sola redirección y la pantalla correspondiente.
- Verificar también el caso de credencial de IA ilegible (`AIConfigurationRequired`), que produce el
  mismo estado.

**Evidencia.** Logs del backend con `entrada correcta de la cuenta de fábrica`; captura del login
colgado; consulta a `/api/setup`; `sesion.guard.ts`. Para poder seguir con el resto de la QA se
configuró la IA local directamente en la base de desarrollo (`mode=local`, `tested_at=now()`), lo
que **no es un arreglo del producto**.

### QA-02 — Alto: el 503 de traducción de correos se confunde con «servidor caído»

**Qué se vio.** En Configuración, cambiar el idioma de Español a English y pulsar «Guardar la
instalación» produce:

```text
POST /api/mail/language/drafts -> 503 Service Unavailable
GET  /api/settings/ai/models  -> 503 Service Unavailable
```

La aplicación muestra el aviso global **«No hay conexión con el servidor»** y un estado
«Cargando…», y bloquea la pantalla con el componente `app-sin-servidor`. El botón «Editar
manualmente» queda detrás del overlay y no se puede pulsar. El idioma no cambia.

**Causa.** El interceptor (`api.interceptor.ts`) trata cualquier 502/503 como servidor caído salvo
`esCatalogoLocalDeIA(peticion.url)`. La ruta `/api/mail/language/drafts` no está excluida, aunque el
503 significa «la IA no está disponible para traducir», no «el backend no responde». El overlay
global tapa la tarjeta y las acciones.

**Impacto.** Un fallo de una dependencia opcional (IA) deja la aplicación entera como si estuviera
caída y bloquea una tarea administrativa. El usuario no puede elegir el camino manual, que la propia
pantalla ofrece.

**Recomendación.**

- Diferenciar el 503 de una dependencia de IA del 503 de nginx/backend; no activar
  `sin-servidor` para endpoints de IA, traducción o correo.
- Mostrar un aviso local en la tarjeta: «No se pudo generar la traducción; puedes editar los
  borradores a mano», y dejar «Editar manualmente» siempre accesible.
- Añadir un E2E con el motor de IA apagado que compruebe que Configuración sigue usable y que el
  camino manual cambia el idioma.
- Revisar el texto de `mail.template.unknownKey` y los de la tarjeta de correos.

**Evidencia.** `tmp/qa/capturas/20-settings-en.png` (estado «Guardando…»), `tmp/qa/capturas/22-language-flow.png`,
respuestas 503 capturadas, y el overlay de `app-sin-servidor` descrito en el intento de QA.

### QA-03 — Alto: textos de «diez correos» y clave cruda de la plantilla 14

**Qué se vio.** El editor `/mail` muestra 14 plantillas, pero el texto dice:

> «Los diez correos que manda la aplicación…»

Y en la lista aparece la clave cruda **`auth.accountCreated`** como primera entrada, seguida de
«Alta de una cuenta» (que es la etiqueta de `auth.invitation`). La plantilla nueva no tiene etiqueta
traducida.

También quedan textos de «diez» en:

- `frontend/src/app/core/i18n/es.ts` (tarjeta de Configuración y cabecera del editor).
- `frontend/src/app/core/i18n/en.ts` (`The ten emails…`).
- `mail.template.unknownKey`: «That is not one of the ten emails».

**Impacto.** El editor de correos es una pantalla de administración; ver una clave interna y un
recuento incorrecto rompe la confianza y la coherencia con la versión 1.0.1, que añadió la plantilla
14. No es un fallo funcional de envío, pero sí de contenido y traducción.

**Recomendación.**

- Cambiar «diez» por «catorce» en español y «ten» por «fourteen» en inglés.
- Añadir la etiqueta de `auth.accountCreated` en `es.ts` y `en.ts` (por ejemplo, «Alta de cuenta
  desde un ticket» / «Account created from a ticket»).
- Revisar el mapa de etiquetas del editor para que ninguna clave llegue a la vista.
- Añadir una prueba que recorra las 14 claves y exija etiqueta y asunto en ambos idiomas.
- Actualizar `mail.template.unknownKey` y cualquier texto/documentación que siga diciendo diez.

**Evidencia.** `tmp/qa/capturas/61-correos.png`; consulta a `mail_templates` con 28 filas y la fila
`auth.accountCreated` correcta en la base (el problema es la interfaz, no la plantilla).

### QA-04 — Media: dos «En curso» en el historial del usuario

**Qué se vio.** En el ticket `CS-2026-0026`, la vista del usuario muestra dos veces:

> «Soporte Uno lo pasó a «En curso» · 31/10/2026, 20:40»

La base tiene dos filas de historial: `en progreso` y `escalado`. Para el usuario, el estado
`escalado` se traduce también como «En curso», así que dos movimientos distintos se leen igual.

**Impacto.** El usuario ve un historial que parece duplicado o un cambio repetido sin motivo. No
pierde datos, pero confunde y hace dudar de la trazabilidad.

**Recomendación.**

- Diferenciar el texto del historial para `escalado` cuando se muestra al usuario (por ejemplo,
  «El ticket pasó al equipo técnico») o agrupar movimientos consecutivos que el usuario ve igual.
- Añadir una prueba de historial con `en progreso` seguido de `escalado`.

### QA-05 — Media: chip del interno con el estado del principal (PR #46)

**Qué se vio.** Con Desarrollo en el interno `INT-CS-2026-0026`, la cabecera muestra el número del
interno pero el chip «En progreso» del principal, mientras el interno está cerrado. Es exactamente
el fallo que corrige el PR #46, que **todavía no está fusionado** en `main`/`dev`.

**Impacto.** Confirmación de que el PR #46 sigue siendo necesario. No es un hallazgo nuevo del
código actual: es el comportamiento de la rama sin el parche.

**Recomendación.** Fusionar el PR #46 y repetir este caso; ya tiene prueba de unidad y E2E.

### QA-06 — Media: «Generando…» indefinido y estados de IA poco claros

**Qué se vio.** Con el motor de IA caído:

- La ficha del ticket pasa de «Sin motor de IA» a «Generando…» después de un cambio de estado y se
  queda así.
- `GET /api/settings/ai/models` y `POST /api/mail/language/drafts` devuelven 503.
- «Mejorar con IA» sigue visible en el editor; al pulsarlo no hay una explicación clara en la ficha.
- Configuración muestra «El administrador local de modelos no está disponible», correcto, pero el
  resto de la pantalla puede quedar en estado de carga por QA-02.

**Impacto.** El usuario no sabe si el resumen está en cola, si va a fallar o si no hay motor. El
estado «Generando…» eterno es peor que un error explícito.

**Recomendación.**

- Poner un estado final «Sin motor de IA» o «No se pudo generar» con reintento manual, y no dejar
  «Generando…» indefinidamente.
- Diferenciar «motor no configurado» de «motor configurado pero caído».
- Añadir una prueba de UI con IA apagada que verifique el texto final y que el resto del ticket
  sigue usable.

### QA-07 — Baja: «Administrador» tres veces en el menú

**Qué se vio.** En el pie del menú lateral:

```text
Administrador
Administrador
Administrador
```

Son el nombre de la cuenta, el papel y la capacidad administrativa. En una cuenta cuyo nombre es
«Administrador» los tres textos coinciden y parecen un error de duplicación.

**Recomendación.** Mostrar nombre y papel, y la capacidad sólo cuando aporte (por ejemplo, «Admin»
una vez), o combinar «Administrador · Admin». No es un fallo funcional, pero afecta a la claridad.

### QA-08 — Baja: cuentas de pruebas antiguas en la lista de Usuarios

**Qué se vio.** `seed-es` dice dejar 11 cuentas, pero la base tiene 18: 11 `@demo.com` y 7 cuentas
`solicitante-*` / `ad-alta-*` de pasadas anteriores. La lista de Usuarios las muestra mezcladas.

**Impacto.** Ruido en desarrollo y en las capturas; no afecta a datos reales. El seed sólo borra las
cuentas `@demo.com`.

**Recomendación.** Que el seed limpie también las cuentas de pruebas conocidas o que documente que
la lista puede contener restos. Para QA, usar una base limpia.

### QA-09 — Resuelta: fechas ISO en los dos idiomas

Las fechas se leen ahora **`AAAA-MM-DD hh:mm`**, en 24 horas y con el mismo orden en español y en
inglés. Se centralizó en `frontend/src/app/shared/fechas.ts` y se conserva la zona horaria de la
instalación. Decisión del responsable, 2026-11-05; PR #61.

### QA-10 — Decidida: «Los dos» no aparece en móvil

En 390 px el conmutador muestra sólo **Principal** e **Interno**. El responsable decidió mantenerlo
así: no caben dos columnas y apilar dos conversaciones largas invita a confundir el ticket activo.
El documento (`docs/interfaz-y-experiencia.md`) ya lo dice de forma explícita.

### QA-11 — Corregida: `/setup` no recuerda el paso al recargar

El asistente guarda los datos de cada paso, pero al recargar vuelve siempre al paso 1. El texto de
cabecera prometía «puedes cerrar y seguir donde lo dejaste». Los campos sí se conservan, así que no
se pierde información; lo que no se conserva es la posición. **Corregido el texto** y **decidido no
persistir la posición**: no merece una migración ni estado en el navegador.

### NVIDIA — funcionando en WSL2

Docker Desktop usa ahora el kernel `6.6.87.2-microsoft-standard-WSL2`, con la RTX A2000 12 GB:

- `docker run --gpus all nvidia/cuda:... nvidia-smi` devuelve la GPU.
- El administrador de IA arranca con `AI_ACCELERATION_CAPABILITIES=cpu,nvidia`, detecta
  `NVIDIA RTX A2000 12GB`, activa `qwen2.5-1.5b-instruct` y lo deja `active`/`healthy`.
- `nvidia-smi` muestra `/llama-server` usando **1.347 MiB**.
- Respuesta en caliente: **0,41–0,43 s** y **~140 tok/s**; CPU en el mismo entorno:
  **2,59–2,82 s** y **~24 tok/s**. La GPU es unas **6 veces** más rápida.
- **Prueba de punta a punta:** el backend contestó `{"status":"ok"}` contra `http://ai:8080`; el
  ticket `CS-2026-0026` generó sus dos resúmenes y la ficha los mostró en el navegador
  (`tmp/qa/capturas/110-ticket-gpu-wsl2.png`).

El seed configuraba la IA con `http://catalina_support_ai:8080`, un nombre que no existe en la red de
Compose. Se corrigió a `http://ai:8080` (PR #55). `seed-es` y `seed-en` se volvieron a ejecutar:
**11 cuentas, 25 tickets** y la URL correcta.

El producto documenta `ai-nvidia.yml` para **Linux con NVIDIA Container Toolkit y Windows con Docker
Desktop WSL2**; Windows nativo, AMD/ROCm y multi-GPU quedan fuera (PR #59).

## 4. Lo que sí funcionó

- **Login local** con `admin/admin`, usuarios, Soporte y Desarrollo.
- **Panel** con filtros, tarjetas, gráficos, XLSX y actualización.
- **Usuarios**: listado, filtros, búsqueda, ficha y acciones visibles.
- **Configuración**: instalación, marca, método de entrada, numeración, reportería, zona horaria,
  dirección pública, IA, correo y enlace al editor; los campos y ayudas se leen.
- **Clasificación**: categorías y etiquetas, renombrar/retirar.
- **Flujo completo de ticket**:
  - Usuario crea `CS-2026-0026` con asunto, categoría, prioridad y descripción.
  - Soporte lo reasigna a Soporte Uno, comenta y lo pasa a En progreso.
  - Soporte lo escala con motivo; se crea `INT-CS-2026-0026` y se avisa a Desarrollo.
  - Desarrollo comenta en el interno y lo cierra; el principal vuelve a En progreso y se avisa a
    Soporte.
  - El usuario ve sólo la conversación pública, no el interno.
- **Validaciones**: asunto/descripción obligatorios y categoría obligatoria, con mensajes claros.
- **Recuperación de contraseña**: el correo llega a Mailpit, el enlace con token en el fragmento
  abre la pantalla, se guarda la contraseña y se entra con ella.
- **Editor de correos**: 14 plantillas, vista previa en español e inglés, marcadores y botones.
- **Asistente `/setup` completo**: cinco pasos en una instalación vacía, correo de prueba entregado
  en Mailpit, proveedor externo simulado (`openai` + `ai-test`), sellado y entrada con `admin/admin`.
- **Idioma inglés**: no aparecieron claves crudas en las pantallas revisadas; los textos estáticos
  están traducidos. Los datos de ejemplo siguen en español, que es lo esperado.
- **Móvil**: sin desbordamiento horizontal en panel, usuarios, configuración, ticket y
  clasificación; el menú se abre como cajón; la ficha baja debajo de la conversación.
- **Tema oscuro**: cambia el fondo y el atributo `data-theme`; vuelve a claro correctamente.

## 5. Cobertura y límites

| Área | Estado |
| --- | --- |
| Login local, sesión, logout | Probado |
| Panel y filtros | Probado |
| Usuarios (lista, filtros, ficha) | Probado; alta completa no creada para no dejar más datos |
| Configuración (lectura y formularios) | Probado completo; cambio de idioma y pruebas de conexión incluidos |
| Editor de correos | Probado en lectura, vista previa y envío de prueba |
| Tickets: alta, asignación, comentario, estados, escalado, cierre interno | Probado |
| Historial y visibilidad usuario/interno | Probado |
| Recuperación de contraseña | Probado de extremo a extremo |
| Inglés y español | Probado en pantallas principales |
| Móvil 390 px y escritorio 1440 px | Probado |
| Temas | Los ocho recorridos por la suite E2E |
| AD y Keycloak | Probados con `tests.yml --profile directory`: 17 aprobadas, 7 omitidas previstas, 0 fallos |
| Asistente `/setup` | **Probado completo**: cinco pasos, correo real en Mailpit, proveedor externo simulado, sellado y login |
| NVIDIA | **Funcionando y soportado en WSL2**: Docker Desktop con kernel WSL2, RTX A2000 12 GB, `--gpus all`, administrador en `nvidia`, modelo activo y resúmenes reales. Documentado para Linux y Windows WSL2 (PR #59). |
| Adjuntos y límites de subida | Probados en la suite E2E completa |
| Accesibilidad con lector de pantalla | Revisión visual y de foco básica, no auditoría WCAG completa |

## 6. Recomendaciones vigentes

1. Mantener la suite completa como puerta de cada cambio de frontend, backend, migraciones o Compose;
   los comandos y resultados están en la sección 10.
2. Mantener el informe en `docs/` como evidencia de la pasada; no se versionan los scripts auxiliares
   de `tmp/qa`.

## 7. Evidencia

Las capturas están en `tmp/qa/capturas/`. Las principales:

- `10-panel.png`, `10-usuarios.png`, `10-configuracion.png`, `10-clasificacion.png`,
  `10-ticket-escalado.png`.
- `20-settings-en.png`, `21-dashboard-en.png`, `21-ticket-en.png`, `22-language-flow.png`.
- `30-mobile-tickets.png`, `31-mobile-ticket.png`, `42-mobile-ticket-en.png`,
  `70-mobile-*.png`, `71-mobile-menu.png`.
- `50-nuevo-ticket.png`, `51-ticket-creado.png`, `52-support-ticket.png`,
  `53-support-flow.png`, `54-dev-cerrado.png`, `57-user-final.png`.
- `60-password-reset.png`, `61-correos.png`, `72-tema-oscuro.png`, `73-validacion-ticket.png`.
- `110-ticket-gpu-wsl2.png`, `111-fechas-iso-movil.png`.

Los scripts de la pasada están en `tmp/qa/*.mjs`. El ticket `CS-2026-0026` y sus correos quedan en
la base de desarrollo como evidencia del flujo. La contraseña de `user5@demo.com` se restauró al
valor de demo y el idioma de la instalación se devolvió a español.

## 9. Correcciones aplicadas

| ID | Estado | Corrección |
| --- | --- | --- |
| QA-01 | **Corregido** (PR #48) | La guarda compara la ruta sin query y no repite la redirección por `aiRequired`. |
| QA-02 | **Corregido** (PR #48) | El 503 de `mail/language/drafts` se trata como dependencia de IA y no activa el aviso global. |
| QA-03 | **Corregido** (PR #48) | El editor dice catorce, etiqueta `auth.accountCreated` y actualiza los textos. |
| QA-04 | **Corregido** (PR general) | El usuario lee «pasó el ticket al equipo técnico» en vez de un segundo «En curso». |
| QA-05 | **Corregido** (PR #46) | La cabecera del interno muestra el estado del interno. |
| QA-06 | **Corregido** (PR general) | Tras agotar las relecturas, la ficha muestra «Sin motor de IA» y no «Generando…». |
| QA-07 | **Corregido** (PR general) | La capacidad administrativa sólo se repite cuando aporta algo distinto del papel. |
| QA-08 | **Corregido** (PR general) | El seed borra todas las cuentas y deja las once de ejemplo; verificado con `seed-es` y `seed-en`. |
| QA-09 | **Resuelto** (PR #61) | Fechas ISO `AAAA-MM-DD hh:mm` en español e inglés, con la zona de la instalación. |
| QA-10 | **Decidido** | En móvil se mantienen Principal e Interno; «Los dos» no se ofrece. |
| QA-11 | **Corregido y decidido** | El texto no promete recordar el paso; no se persiste la posición. |

**Verificación de las correcciones:** pruebas de unidad focalizadas, **269 pruebas del frontend**,
**todo el backend** y la **suite E2E completa**: **200 aprobadas, 30 omitidas y 0 fallos**, primero
en Hyper-V y después en **WSL2** (12,7 min, PC y móvil). El bucle QA-01, el overlay QA-02, la
etiqueta QA-03, el historial QA-04, el menú QA-07, el seed QA-08 y las fechas ISO se comprobaron
también contra la interfaz de desarrollo.

### PR y ramas

- **PR #48:** QA-01/02/03, fusionado en `main`.
- **PR #46:** chip del interno, fusionado en `main`.
- **PR #50:** QA-04/06/07/08 y estabilización E2E, fusionado en `main` (`6d56292`).
- **PR #49:** sincronización de `dev`, fusionado (`be5b014`).
- **PR #47:** modelos e hilos automáticos, fusionado en `main` (`8b062c4`).
- **PR #51:** IA local en el seed, fusionado en `main` (`a788d38`).
- **PR #52:** sincronización final de `dev`, fusionado (`b6af21c`).
- **PR #53:** ajuste del texto de reanudación de `/setup`, fusionado (`f8ec1dc`).
- **PR #54:** sincronización final de `dev`, fusionado (`7765934`).
- **PR #55:** URL del motor en el seed (`ai:8080`), fusionado (`13a7b8a`).
- **PR #56:** sincronización final de `dev`, fusionado (`552f99a`).
- **PR #57:** corrección de la configuración del correo, fusionado (`238dccf`).
- **PR #58:** sincronización de `dev`, fusionado (`d6e714c`).
- **PR #59:** formalización de NVIDIA en Windows WSL2, fusionado (`35bd837`).
- **PR #60:** sincronización final de `dev`, fusionado (`97b0d59`).
- **PR #61:** fechas ISO en la interfaz, fusionado (`c02009e`).
- **PR #62:** sincronización final de `dev`, fusionado (`3d8a5ca`).
- `main` y `dev` tienen el mismo árbol y no quedan PRs abiertos.

**Cobertura adicional cerrada:** AD y Keycloak con el perfil `directory` (**17 aprobadas, 7
omitidas previstas, 0 fallos**), **`/setup` completo** con proveedor externo simulado y correo real
en Mailpit, **seed con IA local** y **NVIDIA funcionando en Docker Desktop WSL2** con resúmenes
reales de la aplicación.

## 10. Verificación final en WSL2

Con Docker Desktop usando WSL2 se repitió la verificación completa:

| Capa | Resultado |
| --- | --- |
| Backend | `go test ./...` correcto |
| Backend | `go vet ./...` correcto |
| Frontend | **269 pruebas aprobadas**, 28 archivos |
| Fronteras del frontend | Correctas |
| E2E completa | **200 aprobadas, 30 omitidas, 0 fallos** (12,7 min) |
| AD y Keycloak | **17 aprobadas, 7 omitidas, 0 fallos** |
| Fechas ISO | Navegador: `2026-11-03 04:00`; móvil con Principal/Interno y sin «Los dos» |

El entorno aislado se eliminó con `down -v`; sólo quedó el entorno de desarrollo.



## 11. Estado del entorno

- No se ejecutó nada fuera del entorno de desarrollo.
- Docker Desktop usa **WSL2** (`6.6.87.2-microsoft-standard-WSL2`) como único backend activo.
- El entorno de desarrollo se reconstruyó desde cero: **4 contenedores** de `dev.yml` y el motor
  de IA con GPU (**5 en total**), **6 volúmenes**, **11 cuentas y 25 tickets del seed**, más el
  ticket de prueba `CS-2026-0026`.
- La IA local está en `mode=local`, `base_url=http://ai:8080`, `model=qwen2.5-1.5b-instruct` y el
  motor NVIDIA está activo y saludable.
- El Docker Engine que quedó instalado dentro de Ubuntu WSL está **detenido, deshabilitado y
  enmascarado**; no se usa.
- El disco antiguo de Hyper-V (`D:\_docker_data\DockerDesktop.vhdx`, ~49 GB) fue eliminado. En
  `D:\_docker_data` sólo queda `wsl`.
- Los datos reales de Docker Desktop WSL2 se movieron a `D:\_docker_data\wsl` y
  `C:\Users\jcestevez\AppData\Local\Docker\wsl` quedó como **junction** a esa carpeta. El VHDX
  activo es `D:\_docker_data\wsl\disk\docker_data.vhdx` (~19,9 GB) y el entorno arrancó intacto
  después del movimiento.
