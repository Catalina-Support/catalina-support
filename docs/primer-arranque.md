# Primer arranque: la instalación desde cero

> **Estado:** as-built
> **Última actualización:** 2026-10-20
>
> **Enmienda aprobada, implementada y verificada el 2026-10-20.** La sección 19 incorpora el icono
> opcional del sitio al paso 1. Las decisiones 1A–14A están cerradas; pasaron 253 pruebas frontend,
> las fronteras y la compilación de producción.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-20: mensaje real desde el paso 4.** La sección 18
> incorpora un destinatario efímero y convierte la prueba SMTP en un envío opcional, limitado y con
> texto fijo en el idioma global. El responsable cerró y aprobó 1A–14A; pasaron las pruebas
> focalizadas, las 248 del frontend, fronteras y compilación.
>
> **Enmienda aprobada, implementada y verificada el 2026-10-19: elección segura del transporte
> SMTP.** La sección 17 incorpora tres modos en el paso 4, con aceptación explícita y límite de red
> privada para el modo sin cifrar. El responsable cerró y aprobó 1A–10A; la prueba real autenticó
> contra el relay interno sin enviar correo.
>
> **Aprobado, implementado y verificado el 2026-10-18.** La sección 16 incorpora Auto, CPU y NVIDIA
> al paso 5 cuando el motor vive en este servidor.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección
> 15 muestra automáticamente el catálogo del paso 5 y reserva su botón para refrescar sin ocultar
> las fichas existentes.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** La sección
> 14 añade tema y color institucional al primer paso y sustituye las confirmaciones de modelos por
> el diálogo temático.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** El paso 1 abre la carga al pulsar cada logo
> sin mostrar el archivo elegido; el paso 5 administra la activación desde las fichas y permite
> quedar temporalmente sin modelo, manteniendo bloqueado **Siguiente**.
>
> **Implementado y verificado el 2026-10-14, con aprobación explícita del responsable.** El paso 1 incorpora los dos huecos
> opcionales del logo y el paso 5 hace visible el progreso real de la descarga, permite cancelarla y
> reanudarla, activa y prueba automáticamente el modelo elegido y no deja avanzar hasta que esté
> activo y saludable. El responsable decidió además probar manualmente el modelo con un saludo que
> incluya la fecha y hora de la zona de la instalación, y administrar en `/setup` los modelos
> inactivos descargados. La sección 12 registra el hallazgo, el alcance cerrado y la verificación.
>
> **Implementado y verificado el 2026-10-07.** `/setup` exige los cinco pasos en el orden aprobado,
> prueba y activa la IA antes de sellar, muestra una sola confirmación de privacidad y posible costo
> para proveedores externos y conserva las descargas parciales de modelos locales para reanudarlas.
> El cierre confía en la activación de IA guardada y no repite una llamada que podría tener costo.
> Pasaron las pruebas de servicios, las 218 pruebas unitarias del frontend y el recorrido aislado
> completo de Playwright: 180 casos ejecutados y 38 omisiones previstas, en PC y móvil.
>
> **Corrección del responsable el 2026-10-07.** Al iniciar la implementación se encontró que la
> sección 11 aprobada colocaba la IA antes del correo, mientras el resto del documento, `AGENTS.md`,
> el código y las pruebas conservaban Correo en el paso 4 e IA en el paso 5. El responsable eligió
> explícitamente ese orden. Esta propuesta corregida alinea todo el documento antes de cerrar la
> implementación; no cambia el contenido de los pasos ni abre producción. El responsable aprobó
> explícitamente el documento corregido el 2026-10-07.
>
> **Enmendado el 2026-10-06**, con aprobación explícita del responsable. Los resultados de guardar y
> probar una conexión se mostraban antes del asistente y podían quedar fuera de la ventana. Ahora usan el mismo toast de
> `/settings` para que el resultado permanezca visible desde el paso actual, sin mover los avisos
> informativos propios de cada paso. Se implementó y verificó en PC y móvil sin cambiar el flujo.
>
> **Enmendado el 2026-10-05**, con aprobación explícita del responsable. En el paso 3,
> la zona elegida sólo se distingue por el resaltado dentro de una lista larga y puede quedar fuera
> de la parte visible. El responsable pidió un campo que indique explícitamente la región
> seleccionada y aceptó la propuesta de mostrarlo antes del buscador, como sólo lectura, con el
> identificador IANA exacto y la hora con su desfase debajo. Se implementó y se verificó en PC y
> móvil dentro de la instalación desechable de Playwright.
>
> **Enmendado el 2026-10-04**, conforme a `docs/prueba-local.md` aprobado: una instalación nueva
> empieza en inglés aunque el navegador recuerde español; el selector traduce todo `/setup` al
> instante y, al reanudar, manda el idioma guardado. En desarrollo, un correo todavía vacío recibe
> los valores editables de Mailpit; lo guardado siempre tiene prioridad y producción sigue vacía.
>
> **Enmendado el 2026-10-03**, conforme a `docs/prueba-local.md` aprobado: El paso de Keycloak
> admite `internalIssuer` opcional. La preparación de Playwright sobre una base nueva recorre los
> cinco pasos en PC y móvil antes de sellar, comprueba Mailpit y la IA y termina en móvil.
> Esta verificación automatizada sustituye la limitación histórica del punto 3 de las desviaciones.
>
> **Enmendado el 2026-09-30**: **el asistente prueba lo que pide**, como prometía la §3 y la §6. Se
> añaden dos endpoints públicos más —`POST /api/setup/entry/test` y `POST /api/setup/mail/test`— con
> el mismo candado del sello (**409** `setup.alreadyInstalled`), el paso 2 y el paso 4 estrenan su
> botón de «Probar la conexión» y la prueba del correo **comprueba la conexión y la autenticación sin
> mandar ningún correo**, porque en el asistente no hay destinatario. Los datos que llegan se validan
> igual que al guardar, con las mismas claves. Con esto **quedan corregidas las desviaciones 1 y 1.b**
> del bloque de abajo. Cuenta como **1 enmienda** (`docs/README.md`).
>
> **Escrito a petición del responsable** (2026-09-30) y **aprobado e implementado el mismo día**: quería
> poder probar **la vista de primer arranque en otra máquina**, y que esa vista **pida toda la
> información necesaria para arrancar** —dominio, dirección, puerto, correo, método de entrada— sin
> tener que editar archivos en el servidor.
>
> **Lo que se desvió de esta propuesta, contado aquí** (Regla 0):
> 1. **Ni la prueba de conexión del directorio o Keycloak (paso 2) ni la del correo (paso 4) estaban en
>    el asistente** al implementarlo. La API de `/api/setup/**` no tenía endpoints de prueba y
>    **todavía no hay sesión** con la que probar: los botones de «Probar la conexión» de Configuración
>    piden ser Administrador, y aquí no hay ninguno. Los datos se guardaban y se validaban —el método
>    tiene que estar configurado para poder terminar—, y **la prueba se hacía después**. **Corregido el
>    2026-09-30**: los dos endpoints existen y son públicos, con el sello como único candado, y el
>    paso 2 y el paso 4 traen su botón (§3.1).
>
> 1.b. **La prueba del correo saliente desde el asistente no se hacía.** El paso 4 guardaba y validaba,
>    pero no probaba nada: para mandar un correo de prueba hay que tener ya el correo configurado y un
>    destinatario, y en el asistente no hay ninguno. **Corregido el 2026-09-30**: el paso 4 tiene su
>    botón y prueba **la conexión y la autenticación**, no el envío, que es lo que se puede comprobar
>    sin destinatario (§3.1).
> 2. **Desviación histórica, corregida el 2026-10-07:** el asistente no tenía un paso propio para
>    probar y activar la IA. Ahora el paso 5 guarda una activación válida antes del resumen y el
>    backend rechaza terminar si falta.
> 3. **Limitación histórica, corregida el 2026-10-03 y ampliada el 2026-10-07 en `tests.yml`:** el
>    recorrido completo se ejecuta sobre una instalación desechable en PC y móvil. Comprueba los
>    cinco pasos, las conexiones, el resumen, el sello y el candado posterior.
> 4. **Los caminos de la API son** `POST /api/setup/installation`, `/entry`, `/entry/test`, `/location`,
>    `/mail`, `/mail/test`, `/ai` y `/finish`.

## 1. Qué problema resuelve

Hoy una instalación nueva **arranca ya configurada** con valores de fábrica, y la única puerta es **la
cuenta de fábrica**, cuya contraseña está en `config/env/*.env`. Para dejarla en serio hay que:

1. editar el archivo de entorno en el servidor (correo, contraseñas, dirección);
2. entrar con la cuenta de fábrica;
3. recorrer **Configuración** —que tiene ocho tarjetas— para poner el nombre, el idioma, el método de
   entrada, la región y la dirección.

**Funciona**, y está documentado en `docs/ambientes.md`. Pero **no se puede probar sin editar un
archivo**, y quien llega de fuera no sabe por dónde empezar. Esta propuesta añade **una vista de
instalación** que se enseña **sólo la primera vez** y pide lo necesario, en orden, con sus pruebas.

## 2. Cómo sabe la aplicación que no está instalada

**Un sello en la propia fila de la configuración**: `installation_settings.installed_at`
(`timestamptz`, **nulo** mientras nadie haya terminado el asistente). Va **en `v1.0.0.sql`** —la
versión sigue abierta— con **relleno para las instalaciones que ya existen**: una base que ya estaba
configurada **queda sellada** con la fecha de su última actualización, así que el asistente **no
aparece** en producción por el simple hecho de actualizar.

**Una instalación está «sin instalar» sólo si el sello es nulo.** Eso hace que el asistente:

- **se enseñe** en una instalación recién creada, y
- **no se pueda volver a recorrer** en una ya instalada, ni por la pantalla ni por la API
  (**409** `setup.alreadyInstalled`), que es lo que impide que alguien reescriba la configuración de
  una instalación en marcha.

## 3. Qué pregunta, y en qué orden

**Cinco pasos**, con su línea de avance y un resumen final. Cada paso guarda al avanzar, así que
cerrar el navegador a medias no pierde lo hecho: al volver, el asistente sigue donde estaba.

| Paso | Qué pide | Por qué |
| --- | --- | --- |
| **1 · La instalación** | **El idioma global**, como primer control, y **el nombre** | Es lo que se lee en la interfaz, la entrada, el menú, los correos y los nuevos resúmenes de IA |
| **2 · Cómo se entra** | **El método** —local, Active Directory o Keycloak— y **sus datos**, **con su prueba**: el botón comprueba lo que corresponda al método elegido. **La cuenta de fábrica se explica siempre**, sea cual sea el método: cómo se llama y **de dónde sale su contraseña** (`ADMIN_PASSWORD`, del entorno) | Es la puerta: sin esto no entra nadie, y la de fábrica entra con **cualquiera** de los tres métodos. **La contraseña no se pide aquí**: sigue en el entorno, y esa regla —que la puerta de fábrica **no dependa de la base**— no se toca |
| **3 · Dónde está** | **La región horaria** —de la lista con buscador— y **la dirección pública** —esquema, host y puerto—, con **el aviso si no es https** | Deciden cómo se leen las fechas y a dónde apuntan los enlaces de los correos |
| **4 · El correo** | **El servidor saliente**: host, puerto, TLS, usuario, contraseña y remitente, **con su prueba** | Sin él no sale ningún correo: ni un alta, ni un restablecer, ni un aviso |
| **5 · Inteligencia artificial** | **Motor local, servidor propio o proveedor externo**. Se elige el modelo o conexión, se aceptan licencia y privacidad cuando correspondan y se ejecuta **Probar y activar** | La IA es obligatoria y diferencial del producto; no se sella la instalación sin una generación válida en el idioma global |

Una instalación nueva abre el asistente completo en **inglés** y muestra English como idioma. El
selector del paso 1 cambia de inmediato todos los textos y `html[lang]`, y recuerda la elección. Si
el paso ya se guardó, al volver manda el idioma persistido en la instalación.

**Y el resumen**, antes de terminar: muestra la modalidad, el proveedor o servidor, el modelo y si
hay credencial, sin enseñar secretos. En proveedor externo repite el aviso de privacidad aceptado.
El botón **Terminar la instalación** sólo se habilita con los cinco pasos completos y pone el sello.

**Los tres métodos se pueden elegir desde el principio**, sin esperar a que estén configurados: en
Configuración un método sin sus datos se ofrece apagado, y aquí **no puede ser**, porque si no habría
forma de rellenar los campos del que se quiere usar. Elegirlo abre sus campos, y al avanzar se
comprueba que estén completos.

### 3.1 Las pruebas de conexión del asistente

Los dos pasos que configuran algo que hay que comprobar traen su acción de prueba. En la entrada se
llama **Probar la conexión**; en correo, **Enviar correo de prueba**. Ninguna guarda: mandan lo escrito
y permiten corregirlo antes de avanzar.

| Paso | Qué prueba | Qué **no** hace |
| --- | --- | --- |
| **2 · Cómo se entra** | **El directorio (AD) o Keycloak, según el método elegido**: con AD abre la conexión y valida la cuenta de servicio; con Keycloak lee el documento del reino. Con el método **local** no hay nada que conectar y el botón no se enseña. | No guarda nada, no valida la contraseña de nadie —eso sólo se sabe cuando alguien entra— y no toca la base. |
| **4 · El correo** | La transacción completa: conexión, autenticación, `MAIL FROM`, `RCPT TO` y `DATA`, con el remitente y el destinatario efímero escritos. | No guarda los datos ni afirma entrega final al buzón. |

**Las contesta quien ya sabía hacerlo**, sin duplicar nada: el directorio y Keycloak los prueba el
módulo `auth` —el mismo que contesta a los botones de Configuración— y el correo lo prueba el módulo
`mail`, con la conexión que usa para enviar. El módulo `settings` declara lo que necesita
(`Prober`, `ProberDeCorreo`) y `main.go` se lo da, como manda la regla de modularidad
(`docs/arquitectura.md`, sección 4).

**Los datos que llegan se validan igual que al guardar**, con las mismas claves: el método tiene que
ser uno de los tres y estar configurado, el directorio y Keycloak tienen que estar completos, y el
correo necesita servidor, puerto y remitente con arroba. La contraseña o el secreto vacíos usan **lo
que ya esté guardado**, que es lo que permite probar sin volver a escribir un secreto que no sale
nunca por la API. Cuando la prueba no llega, la pantalla lee `settings.directory.unreachable`,
`settings.keycloak.unreachable` o `setup.mail.unreachable`, según el caso, y el detalle técnico queda
en el log.

## 4. Lo que el asistente **no** pregunta, y por qué

- **Claves comerciales, compra de créditos y GPU.** El asistente acepta una credencial que ya
  exista y permite elegir un modelo del catálogo local, pero no contrata servicios ni configura el
  hardware del servidor.
- **La contraseña de la cuenta de fábrica.** Sigue siendo `ADMIN_PASSWORD` del entorno
  (`docs/usuarios-y-permisos.md`, sección 8). El asistente **cuenta de dónde sale** en lugar de
  pedirla, porque si la contraseña viviera en la base, la puerta de fábrica **dejaría de ser
  independiente** del resto de la configuración.
- **El certificado y el dominio de nginx.** Son del servidor: van en el README, no en la aplicación.

## 5. El correo saliente se muda de sitio

Hoy el SMTP se lee del **entorno**; el paso 4 lo pide en la pantalla, así que pasa a **vivir en la
base**, como el directorio y Keycloak (`docs/modules/settings.md`). Se hace con el patrón que ya está
en el repositorio: **el módulo `mail` declara lo que necesita** (`Instalacion`, una interfaz con el
SMTP resuelto) y **`main.go` se lo da** desde el módulo de configuración. **Y las variables `SMTP_*` se retiran del entorno**, por la misma razón que se retiraron las del
directorio y Keycloak (corrección del responsable, 2026-09-30): **un dato que vive en la base no se
configura en dos sitios**. En desarrollo, `SETUP_MAIL_*` aporta únicamente sugerencias editables de
Mailpit (`mail:1025`, sin TLS ni credenciales y con el remitente local) mientras la instalación está
sin sellar y no tiene correo guardado. Guardar el paso las persiste; después siempre manda la base.
Producción no declara esas sugerencias y empieza vacía. El módulo `mail` nunca las lee para enviar.

## 6. Dónde vive, y con qué se prueba

- **La pantalla**: una ruta propia, `/setup`, **sin el armazón del menú** —no hay nada que navegar
  todavía— y con la marca de fábrica.
- **La API**: `/api/setup/**`, en el módulo `settings`, **disponible sólo mientras el sello sea
  nulo**. Guardar dos veces la misma instalación se rechaza con **409** y su clave. Además de los cinco
  caminos que guardan y sellan, están **las dos pruebas** del paso 2 y del paso 4 —`POST
  /api/setup/entry/test` y `POST /api/setup/mail/test`—, que reciben los datos escritos, **no tocan la
  base**, contestan `{"status":"ok"}` —correo añade `recipient`— o la clave del fallo, y **con la instalación sellada contestan
  también 409** `setup.alreadyInstalled` (§3.1).
- **Las pruebas**: las unitarias de cada paso y **las dos de los endpoints de prueba** —con dobles, sin
  salir a la red: el 409 con el sello puesto, la validación de los datos malos y que con datos buenos
  se llama a la prueba del módulo—, la del remitente del correo —que entrega el mensaje fijo— y **la
  que más importa en la suite de interfaz**: que **una instalación ya configurada no
  enseña el asistente y su API lo rechaza**. El recorrido completo de los cinco pasos se ejecuta
  automáticamente sobre una instalación desechable, en PC y móvil, antes de la suite normal.

## 7. Lo que queda fuera, y hay que decirlo

- **Comprar o aprovisionar un proveedor de IA**: se configura una cuenta y credencial existentes.
- **Reinstalar**: una vez sellada, la configuración se cambia en **Configuración**. Para volver a
  empezar de cero hay que **borrar la base** y aplicar el esquema otra vez, y eso es del servidor.
- **Mientras el asistente esté sin terminar, la instalación no se expone**: cualquiera que llegue a
  una instalación sin sellar puede configurarla. El README lo dirá en su paso a paso: **primero se
  configura, después se abre el dominio**.

## 8. Qué habilita este documento

Aprobar esta propuesta habilita: la columna del sello en `v1.0.0.sql` con su relleno, la ruta
`/setup` con sus cinco pasos y su resumen, los endpoints `/api/setup/**` con su candado, el correo
saliente en la base —y fuera del entorno—, y las pruebas de las tres capas. **Nada de eso se
escribe hasta que el responsable lo apruebe.**

## 9. Región seleccionada visible en el paso 3

### 9.1 Hallazgo

En **Where it is / Dónde está**, la lista marca la zona elegida con fondo y peso de texto. Esa marca
no basta: la lista es larga, tiene desplazamiento propio y la opción seleccionada puede no estar en
la parte visible. Después de elegir, el buscador se limpia y no queda un dato estable que diga cuál
es la región actual. La hora y el desfase ayudan a validarla, pero no nombran la zona.

### 9.2 Comportamiento

Antes del buscador aparece un campo de sólo lectura:

- etiqueta **Selected time zone** en inglés y **Región horaria seleccionada** en español;
- valor IANA exacto, por ejemplo `America/Guayaquil`;
- al abrir el paso muestra la zona que ya tiene el asistente —`UTC` en una instalación nueva o la
  persistida si se está reanudando—;
- al elegir una opción de la lista, el campo cambia inmediatamente;
- su texto se puede seleccionar y copiar, pero no editar: una zona sólo entra al estado al elegirla
  de la lista válida;
- la hora actual y el desfase continúan debajo de la lista, como hoy.

El buscador, la lista, el resaltado de la opción, la dirección pública y la forma de guardar el paso
no cambian. El campo ocupa el ancho disponible tanto en PC como en móvil y conserva una etiqueta
visible y asociada para lectores de pantalla.

### 9.3 Alcance de la implementación

La implementación modifica `frontend/src/app/core/pages/setup-page.html`, añade los dos textos al
diccionario estricto ES/EN y amplía la preparación Playwright de la instalación para comprobar el
valor inicial y el cambio de zona. `docs/interfaz-y-experiencia.md` recoge también el comportamiento
final.

Quedan fuera la API, la base de datos, la validación de zonas, la pantalla general de Configuración
y los otros tres pasos de `/setup`.

### 9.4 Criterios de aceptación

1. Al abrir el paso 3 se ve el campo con `UTC` o con la zona guardada que corresponda.
2. Elegir `America/Guayaquil` actualiza el campo en el acto y mantiene debajo su hora y desfase.
3. El campo no admite escritura, pero permite seleccionar y copiar su texto.
4. La búsqueda y la lista siguen funcionando con teclado y lector de pantalla, sin identificadores
   duplicados.
5. El comportamiento se comprueba en PC y móvil dentro de la instalación desechable de Playwright.

### 9.5 Decisión registrada

El responsable confirmó la recomendación el 2026-10-05: campo de sólo lectura antes del buscador,
identificador IANA exacto y hora/desfase conservados debajo.

### 9.6 Registro del repaso

El responsable confirmó el 2026-10-05:

1. El texto se selecciona y se copia con el comportamiento normal del campo; no se añade un botón
   de copiar.
2. Esta última observación se limita al paso 3 de `/setup`; la pantalla general de Configuración no
   cambia.

El repaso no deja decisiones abiertas. El responsable aprobó explícitamente esta enmienda con
«Sí». La implementación respeta el alcance aprobado y el documento vuelve a `as-built`.

### 9.7 Verificación

- Las 216 pruebas unitarias del frontend pasan.
- La preparación Playwright de una instalación vacía pasa en PC y móvil: comprueba `UTC` en el
  recorrido nuevo, `America/Guayaquil` al reanudar, la actualización inmediata después de elegir y
  el atributo de sólo lectura.
- Los cuatro casos de `arranque-instalacion.spec.ts` pasan en PC y móvil después de sellar esa
  instalación.

## 10. Resultados visibles durante el primer arranque

### 10.1 Regla

El mensaje general situado antes del asistente se sustituye por el toast compartido definido en
`docs/interfaz-y-experiencia.md`, sección 15. Lo usan los resultados transitorios de:

- fallar al guardar un paso o avanzar; cuando el guardado sale bien, el asistente avanza sin añadir
  un mensaje redundante;
- probar Active Directory o Keycloak en el paso 2;
- probar el correo saliente en el paso 4;
- errores generales al cargar, guardar, probar o terminar la instalación.

Los avisos que explican el paso actual permanecen dentro del contenido: la cuenta de fábrica, la
dirección sin HTTPS y cualquier información que deba poder consultarse mientras se rellenan los
campos. Los errores de un campo permanecen asociados al campo cuando exista esa presentación.

### 10.2 Alcance y criterios de aceptación

Esta enmienda visual no cambia los pasos, sus endpoints públicos, el sello, las validaciones, el
idioma inicial, la persistencia ni el resumen final.

1. El resultado de las pruebas de los pasos 2 y 4 se ve en la ventana actual en PC y móvil.
2. Un fallo al avanzar se muestra sin llevar la ventana al principio y permanece hasta cerrarse o
   ser sustituido.
3. Cualquier toast desaparece a los 10 segundos como máximo y puede cerrarse antes.
4. Los avisos de contexto siguen dentro del paso al cambiar de idioma o de ancho de pantalla.
5. Las pruebas del asistente comprueban una conexión correcta, una fallida, el cierre manual y la
   posición fija del toast.

### 10.3 Decisión y repaso

El responsable confirmó el 2026-10-06 que `/setup` use el mismo componente y las mismas reglas que
`/settings`. El repaso común quedó cerrado en `docs/interfaz-y-experiencia.md`, sección 15.7, sin
decisiones abiertas.

### 10.4 Implementación y verificación

`setup-page` muestra su señal general mediante `app-toast`; sus avisos explicativos permanecen dentro
del paso correspondiente. La preparación Playwright de la instalación vacía comprobó en PC y móvil
que la prueba del correo muestra el toast fijo dentro de la ventana y que su botón **Cerrar** lo
retira. El recorrido vigente conserva los cinco pasos, sella la instalación y termina en la entrada.

## 11. Implementación: cinco pasos con idioma global e IA obligatoria

El orden nuevo es:

1. **Idioma e instalación:** el idioma es el primer control y cambia inmediatamente todos los textos
   del asistente; después se pide el nombre. Ese idioma será único para interfaz, fechas, correos e IA.
2. **Cómo se entra:** conserva local, AD y Keycloak y sus pruebas actuales.
3. **Dónde está:** conserva zona horaria y dirección pública.
4. **Correo saliente:** prueba opcionalmente la transacción completa con un destinatario efímero.
5. **Inteligencia artificial:** local, servidor propio o proveedor. En local se elige el catálogo,
   se comprueban recursos, se muestra fuente, licencia, tamaño y checksum, se acepta la licencia de
   esa versión, se descarga con progreso y se activa. En externo se escriben conexión y secreto, y
   se confirma qué datos saldrán del servidor.

El paso 5 sólo queda completo después de **Probar y activar** con una generación artificial válida
en el idioma global. Una descarga local incompleta **se reanuda desde el estado guardado** y el
asistente permanece en el paso 5; cerrar o recargar el navegador no convierte una descarga parcial
en modelo instalado ni obliga a repetir lo ya descargado. El botón de terminar permanece
deshabilitado si falta IA, correo o cualquiera de los pasos anteriores.

Una activación válida ya guardada basta para terminar. **Terminar no repite la prueba de IA**: si el
motor deja de responder entre la activación y el sello, la instalación puede completarse y la caída
posterior se trata como cualquier indisponibilidad del motor, dejando los trabajos pendientes. Así
una interrupción transitoria no bloquea una instalación correctamente configurada.

El resumen final muestra modalidad, proveedor o servidor, modelo y si la credencial está puesta,
nunca el secreto. Para un proveedor externo, antes de probar se presenta **una sola confirmación**
que reúne el aviso de privacidad y el posible costo; no se piden dos aceptaciones ni se repite la
confirmación para la misma activación. El resumen recuerda que fue aceptada. Al terminar se sella en
una sola operación como hoy; un fallo no deja un sello parcial.

Quedan fuera pedir claves comerciales, comprar créditos y configurar GPU. La contraseña de fábrica
continúa en `ADMIN_PASSWORD`. La propuesta sustituye la afirmación actual de que el asistente sólo
comprueba automáticamente un motor opcional y la exclusión de elegir modelo desde la aplicación.

El responsable confirmó el alcance el 2026-10-06. El repaso añadió la aceptación de licencia e
integridad del catálogo y confirmó que los costos sólo afectan a recuperaciones posteriores, que en
proveedor quedan bajo decisión del Administrador. El 2026-10-07 corrigió explícitamente el orden:
Correo permanece en el paso 4 e IA ocupa el paso 5. En el repaso final del mismo día eligió que una
activación guardada permita sellar aunque el motor caiga después, una sola confirmación conjunta de
privacidad y costo para proveedores externos, y reanudar las descargas locales incompletas. No
quedan decisiones abiertas.

## 12. Marca y administración visible de modelos

### 12.1 Hallazgos

La implementación no cumple dos afirmaciones vigentes. El catálogo entrega `downloadedBytes`, pero
`/setup` y `/settings` sólo consultan una vez y enseñan «Descargando» sin barra, cantidad ni avance.
Tampoco existe la cancelación que promete `docs/interfaz-y-experiencia.md`. Además, el paso 5 deja
pulsar **Siguiente** durante una descarga y entonces intenta probar un motor que todavía no está
listo, por lo que muestra el error genérico de motor inaccesible. Se reportan antes de tocar código.

El paso 1 sólo pide idioma y nombre. Los logos se pueden cargar después desde Configuración, aunque
forman parte de la misma identidad inicial. El responsable pidió llevar sus dos huecos al paso 1.

### 12.2 Paso 1: dos logos opcionales

Después del idioma y el nombre aparecen los huecos **Claro** y **Oscuro**, con la misma vista previa,
validación, reemplazo, caída al otro hueco y vuelta al archivo de fábrica que Configuración. Los dos
son opcionales y nunca bloquean **Siguiente**. Cada archivo se guarda al subirlo, el encabezado adopta
la marca inmediatamente y recargar conserva lo guardado.

Las rutas viven bajo `/api/setup/**`, sólo funcionan con el sello nulo y después responden `409
setup.alreadyInstalled`. Conservan PNG, JPEG, WebP y SVG, el máximo de 1 MB y 2000 píxeles de lado.

### 12.3 Paso 5: descarga y activación

Sólo hay una descarga a la vez. La ficha muestra barra, porcentaje y bytes descargados y totales,
actualizados automáticamente. **Cancelar** detiene la transferencia y conserva el parcial;
**Reanudar** continúa desde él. Otra acción permite eliminar definitivamente un parcial. Cerrar o
recargar el navegador no detiene el trabajo que ya quedó encargado al servidor.

Al completar el archivo elegido, el servidor valida tamaño, cabecera y checksum, lo activa, espera
salud y ejecuta la prueba automática vigente. La ficha distingue descarga, validación, activación,
prueba, activo y error. Un fallo conserva el modelo instalado, explica la etapa y ofrece
**Reintentar activación**. **Siguiente** permanece deshabilitado hasta que el elegido esté activo y
saludable y dice qué falta, en vez de provocar una prueba prematura.

Los modelos instalados inactivos se pueden eliminar con confirmación. El activo no ofrece esa
acción: primero se activa otro. La misma regla vale en `/settings`.

### 12.4 «¡Pruébame!» / “Try me!”

El modelo activo y saludable ofrece un botón **¡Pruébame!** en español y **Try me!** en inglés. El
backend toma el instante actual, lo expresa en la zona IANA de la instalación y se lo entrega al
modelo para que salude en el idioma global e incluya esa fecha y hora. No se confía en que el modelo
conozca el reloj ni se envían tickets, personas o archivos.

La ficha conserva el saludo, el instante autoritativo y la zona hasta repetir la prueba, cambiar de
modelo o abandonar la pantalla. Un fallo se queda dentro de la ficha, no desactiva el modelo y
permite reintentar. Esta prueba manual no sustituye la prueba automática que habilita **Siguiente**.

### 12.5 Alcance y cierre

Se conserva el catálogo 1.5B/3B/7B, una sola descarga simultánea, un solo modelo activo y los
permisos existentes. No entran modelos personalizados, producción, el dominio ni la etiqueta
`v1.0.0`.

Las pruebas cubren servicios, controladores y componentes, y un recorrido desechable en PC y móvil:
progreso, cancelación, recarga, reanudación, trabajo en segundo plano, activación, fallo y reintento,
saludo con zona e instante, borrado seguro y ambos logos. El repaso del responsable cerró todas las
decisiones: progreso completo y cancelación; activo y saludable para avanzar; activación automática
en segundo plano; una descarga; fallo conservado; logos inmediatos y opcionales; borrado sólo de
inactivos; y saludo visible dentro de la ficha. No quedan decisiones abiertas.

## 13. Logos pulsables y control desde las fichas

En el paso 1, pulsar la vista previa clara u oscura abre el selector correspondiente. El archivo se
guarda inmediatamente y sólo se ve el logo actualizado: no se presenta su ruta ni su nombre. El
control conserva nombre accesible, foco visible y operación con teclado; restaurar sigue siendo una
acción separada.

En el paso 5 desaparece el selector superior del modelo local. Cada modelo instalado muestra su
estado y **Activar** o **Desactivar** en la propia ficha. Se permite dejar el motor sin activo, pero
en ese estado **Siguiente** permanece bloqueado y explica que hace falta activar un modelo saludable.
El saludo manual se presenta por toast, atribuido al nombre del modelo, durante un máximo de diez
segundos y con cierre manual. No repite la fecha y hora que ya dice el modelo ni enseña la zona.

Las pruebas recorren ambos logos y los estados sin activo, activo y saludable en PC y móvil. No
cambian los cinco pasos, el sello, los límites de logo, el catálogo ni la prueba automática.

## 14. Apariencia inicial y confirmaciones propias

El paso 1 reúne nombre, idioma, tema, color institucional y los dos logos. Idioma y tema aparecen
juntos; el tema ofrece las ocho opciones existentes y cambia en vivo mediante la preferencia local
del navegador. Debajo, el color usa el control y la vista previa de Configuración. El color elegido
se conserva al cambiar de tema y viaja junto con nombre e idioma cuando se pulsa **Siguiente**.

El contrato del paso 1 añade `primaryColor`. La vista previa usa el contrato existente de resolución
del color sin persistir; guardar reutiliza la misma validación hexadecimal y actualiza la marca
pública. Recargar conserva el color ya guardado, mientras el tema sigue siendo una elección de ese
navegador y no una columna nueva.

Eliminar un modelo o un parcial abre `app-dialogo`, presentado con el tema visible. Cancelar, Escape
o cerrar no ejecutan la petición; confirmar usa el botón peligroso y conserva las reglas actuales de
borrado.

Las pruebas verifican los ocho temas, color y vista previa, persistencia al avanzar, recarga, diálogo
en temas claro y oscuro, teclado, PC y móvil. No cambia el orden ni el número de pasos.

El recorrido real confirmó el catálogo completo de ocho temas, la persistencia local de la elección,
la vista previa resuelta del color y el diseño móvil a 390 px. Avanzar guardó el color junto con el
nombre y el idioma; el diálogo de borrado se abrió y cerró con Escape sin eliminar el modelo.

## 15. Catálogo visible al entrar al paso 5

Entrar al paso 5 con modalidad local solicita inmediatamente el catálogo y presenta sus tres fichas.
Volver de otra modalidad a **En este servidor** hace lo mismo, sin exigir una pulsación adicional.

El botón se llama **Refrescar catálogo** / **Refresh catalog** y sólo repite la consulta. Mientras
trabaja conserva el catálogo anterior, muestra su estado de trabajo y evita peticiones duplicadas.
Si falla, mantiene las fichas conocidas y muestra el error mediante el toast compartido.

Las pruebas verifican entrada directa, regreso a local, refresco exitoso y fallido, PC y móvil. No
cambian el quinto paso, el sello, la API ni la administración de cada modelo.

El recorrido real avanzó desde el paso 1 con los valores guardados y encontró las tres fichas ya
cargadas en el paso 5. El refresco manual terminó en el sitio sin ocultarlas ni alterar modelos.

## 16. Aceleración en el paso de IA

El paso 5 muestra Auto, CPU y NVIDIA únicamente al elegir **En este servidor**, con Auto como valor
inicial. Activar un modelo reinicia y prueba usando el modo elegido; **Siguiente** continúa bloqueado
hasta tener modelo saludable. La ficha y **Try me!** muestran modo efectivo y nombre comercial del
dispositivo sin identificadores técnicos.

Auto cae una vez a CPU con aviso. NVIDIA explícito sin la variante Compose rechaza la activación y
explica cómo levantar `ai-nvidia.yml`; no sella ni guarda una configuración inservible. Reanudar el
asistente reconstruye modo configurado, modo efectivo y operación `restarting` si existe.

## 17. Transporte SMTP en el paso 4

El paso 4 sustituye la casilla **Conexión cifrada (TLS)** por una selección de tres modos:
**STARTTLS**, **TLS directo** y **Sin cifrado**. STARTTLS es la opción normal del puerto 587 y exige
la extensión antes de autenticar; TLS directo corresponde normalmente al 465. La pantalla no elige
un modo a partir del puerto ni degrada automáticamente una conexión segura.

**Sin cifrado** muestra antes de probar o avanzar que usuario, contraseña y contenido viajarán en
claro. Requiere una aceptación administrativa que se conserva sólo mientras siga seleccionado ese
modo. Además, el backend exige que todas las direcciones resueltas sean IPv4 privadas RFC1918 o
loopback. El toast común cuenta un rechazo por modo, aceptación o destino sin revelar credenciales.

La prueba conecta, autentica y envía el mensaje técnico. El remitente será una dirección
institucional propia de Catalina; el asistente no propone ni copia la de SIGOC. Las sugerencias de
desarrollo cambian Mailpit a `plain` sin credenciales y con su excepción local explícita.

## 18. Correo real de prueba en el paso 4

Esta enmienda reemplaza la última frase de la sección 17 y la fila del paso 4 en §3.1: el asistente
sí pide una dirección **Enviar prueba a** y el botón único **Enviar correo de prueba** ejecuta un
envío SMTP completo con lo que está escrito. El campo se usa sólo en esa petición; no se guarda ni
aparece en el resumen.

El mensaje fijo usa el idioma global y explica que valida Catalina-Support, que no requiere acción
y que no debe responderse. El resultado confirma que el servidor aceptó el mensaje para la dirección
indicada, sin prometer que llegó a la bandeja. No probar sigue permitiendo avanzar y sellar.

Como el endpoint existe antes de iniciar sesión, limita cada pareja IP–destinatario a cinco intentos
en quince minutos, responde `429` al excederlos y permanece inaccesible después del sello. Los logs
ocultan parcialmente el destinatario y excluyen secreto y contenido.

## 19. Icono opcional en el paso 1

Después de los dos logos, el paso 1 presenta la tarjeta **Icono del sitio** definida en
`docs/interfaz-y-experiencia.md`, sección 50. Subirlo o restaurarlo usa las rutas públicas de
`/api/setup/brand/site-icon`, disponibles sólo mientras la instalación no esté sellada, y actualiza
el favicon de todas las pestañas abiertas inmediatamente.

El icono no es obligatorio y no participa en **Siguiente**, en la validación del paso ni en el sello
final. Si se omite, la pestaña usa el logo claro personalizado y, si éste tampoco existe, el logo
claro de fábrica. Los errores permanecen en el paso mediante el toast común y no descartan los demás
datos ya escritos.
