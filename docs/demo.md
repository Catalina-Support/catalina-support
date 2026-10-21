# Demostración pública alojada

> **Estado:** aprobado
> **Última actualización:** 2026-10-21
>
> **Corrección del responsable, 2026-10-21:** la infraestructura específica de la demo no se
> publicará en este repositorio. Vivirá en el repositorio privado
> `Catalina-Support/demo`, disponible localmente en `/root/dev/catalina-support/demo`. Este proyecto
> público sólo incorporará capacidades genéricas configurables que también puedan usar otras
> instalaciones.

## 1. Propósito

`https://demo.catalina-support.com` será una instalación pública y compartida de Catalina Support
para evaluar el producto sin instalarlo. Tendrá las once cuentas, los veinticinco tickets y los
cinco adjuntos del seeder. No es producción, no contiene información real y no ofrece SLA.

La demo se desplegará en el servidor actual, pero será un entorno independiente de desarrollo y
producción. Su base, adjuntos, configuración y sesiones se restaurarán todos los días a las 03:00 de
`America/Guayaquil`.

## 2. Hallazgo que origina la propuesta

`docs/ambientes.md`, sección 25, presenta la demo como destino público, pero dice expresamente que
no es un entorno administrado por el repositorio y deja el script, el calendario y la operación
fuera de él. Esa descripción ya no alcanza para levantarla de forma repetible y contradice el nuevo
objetivo. La implementación aprobada deberá sustituir esa regla y actualizar, sin hacerlo en
silencio, `ambientes.md`, `arquitectura.md`, `prueba-local.md`, `proyecto-abierto.md`,
`interfaz-y-experiencia.md`, `modules/auth.md`, `modules/users.md`, `modules/ai.md` y
`modules/settings.md`.

## 3. Repositorios y aislamiento

La solución se divide en dos repositorios:

- `Catalina-Support/catalina-support`, público, contiene el producto y capacidades genéricas:
  aviso configurable de entorno compartido, cuotas configurables, protección configurable de
  cuentas compartidas e integración opcional con un analizador antivirus;
- `Catalina-Support/demo`, privado y clonado en `/root/dev/catalina-support/demo`, contiene todos
  los valores y mecanismos propios de `demo.catalina-support.com`.

El repositorio privado fija un commit exacto del público y construye ese código sin mantener un fork
ni aplicar parches privados. El binario desplegado debe corresponder al commit público registrado;
si necesita una capacidad nueva, primero se incorpora y verifica como función genérica en el
repositorio público.

La demo tendrá:

- Compose propio para la aplicación y otro para su motor de IA;
- archivo de entorno no versionado y ejemplo sin secretos;
- proyectos, contenedores, redes, volúmenes, puertos y directorios exclusivos;
- raíz de despliegue `/srv/catalina-support-demo`;
- base PostgreSQL, adjuntos, logs, Mailpit y modelo separados de desarrollo y producción;
- sólo HTTPS expuesto al público; PostgreSQL, Mailpit, administrador de IA y motor permanecerán en
  redes internas.

Los nombres y rutas no reutilizarán el prefijo `prod`. Ningún comando de demo aceptará por defecto
el archivo `prod.env`, los volúmenes de producción ni `/srv/catalina-support`.

Compose, nginx, cron, ranuras, límites concretos, cuentas protegidas, monitor y secretos pertenecen
al repositorio privado. No se copiarán al público ni siquiera como ejemplos con la topología real.

## 4. Publicación y actualización

La fuente de una publicación será un commit exacto de `main` que el repositorio privado registre en
su manifiesto. Su script obtiene el repositorio público, exige un árbol limpio, verifica el commit,
construye los artefactos, registra ambas revisiones y actualiza únicamente la demo. La actualización
será manual; un push a cualquiera de los dos repositorios no despliega automáticamente.

El nginx del host publicará `demo.catalina-support.com`, redirigirá HTTP a HTTPS y usará un
certificado de Let's Encrypt con renovación automática. La dirección pública guardada en la
instalación será `https://demo.catalina-support.com`.

La demo no tendrá SLA. Un monitor HTTPS externo consultará cada cinco minutos una ruta de salud y
avisará por correo al administrador. El proveedor, destinatario y credenciales del monitor son
configuración operativa externa al repositorio.

## 5. Estado inicial conocido

El estado restaurable se construirá, no se leerá de un volcado con secretos. El proceso:

1. crea una base y un directorio de adjuntos vacíos en una ranura inactiva;
2. aplica todas las migraciones versionadas con el ejecutor normal;
3. fija nombre y marca de fábrica, inglés global, zona `America/Guayaquil`, dirección pública HTTPS,
   entrada local y sello de instalación;
4. configura Mailpit interno y la IA local dedicada;
5. instala y activa el modelo local 1.5B;
6. aplica los mismos datos del seeder: once cuentas, veinticinco tickets y cinco adjuntos;
7. verifica esquema, conteos, acceso, salud, correo interno e IA;
8. conmuta el tráfico sólo después de completar todas las verificaciones.

El modelo descargado pertenece a la demo y persiste entre restauraciones para no descargarlo cada
día. La configuración efectiva sí vuelve al valor conocido. La contraseña administrativa procede
de una variable secreta y nunca se publica; las cuentas públicas mantienen la contraseña documentada
del seeder.

## 6. Restauración diaria y rollback

`scripts/demo-reset.sh` se ejecutará diariamente a las 03:00 con
`CRON_TZ=America/Guayaquil`. Usará un bloqueo para impedir dos restauraciones simultáneas y dos
ranuras alternas para base y adjuntos.

Mientras prepara la ranura inactiva, la activa sigue atendiendo. Después de verificar la nueva, el
script muestra una ventana breve de mantenimiento, invalida las sesiones anteriores, conmuta nginx
de forma atómica y retira la ranura vieja. Si cualquier paso previo falla, no conmuta: conserva el
último estado funcional, registra el error y dispara el aviso operativo. No reintenta
indefinidamente.

La restauración elimina tickets, cuentas añadidas, cambios administrativos, contraseñas cambiadas,
marca, correo, configuración, sesiones, cuotas y adjuntos de visitantes. No se crean copias de la
base desechable.

## 7. Correo

La demo usa Mailpit en una red interna. Sirve para que el producto complete sus operaciones, pero su
interfaz y API no se publican. No se configura SMTP real y ningún mensaje sale del servidor.

La prueba de correo puede entregar en ese Mailpit privado. La interfaz no promete al visitante que
podrá abrir el buzón.

## 8. IA

La IA es obligatoria, local y dedicada a la demo. Usa el modelo 1.5B y no llama a proveedores
externos. El administrador del motor, su token, red y volumen no se exponen públicamente.

Los límites de la demo son cinco solicitudes de IA por minuto y cincuenta por día para cada cuenta.
Los rechazos deben explicar cuándo se recupera la capacidad sin revelar contadores de otras
personas. La restauración diaria reinicia la cuota.

## 9. Acceso y cuentas públicas

El único método de entrada es local. Se publican las once cuentas del seeder; `admin` y su contraseña
permanecen secretos.

En modo demo, las once cuentas públicas no pueden cambiar ni recuperar su contraseña. El backend
aplica la regla, no sólo la interfaz, y responde con un mensaje que explica que la credencial es
compartida. La cuenta administrativa secreta conserva ambos recorridos. No se cambia el
comportamiento de una instalación que no esté en modo demo.

## 10. Aviso visible

La marca pública indicará que la instalación está en modo demo y la hora de la próxima restauración.
La interfaz mostrará en todas las pantallas un banner accesible y adaptable que diga:

- que el entorno es compartido;
- que se restaura diariamente a las 03:00 de `America/Guayaquil`;
- que no deben ingresarse datos personales, confidenciales ni sensibles.

El aviso no se puede ocultar de forma permanente. La pantalla de entrada lo muestra antes de
autenticarse.

## 11. Límites y adjuntos

Además de los límites propios del producto, la demo aplica:

| Recurso | Límite |
| --- | --- |
| Entrada | 10 intentos por minuto e IP |
| IA | 5 solicitudes por minuto y 50 por día y cuenta |
| Comentarios | 30 por minuto y cuenta |
| Adjuntos | límites de extensión y tamaño existentes, y 100 MiB diarios por IP |
| Exportaciones | límite por IP para impedir ejecuciones continuas |

Los valores exactos de exportación se fijarán en la implementación a partir del costo medido, sin
ser menos restrictivos que diez exportaciones por hora e IP.

Todo adjunto nuevo pasa primero a cuarentena y se analiza con un servicio antivirus dedicado. Sólo
un resultado limpio permite moverlo al almacén definitivo y crear o conservar su referencia. Un
resultado infectado, indeterminado o una caída del antivirus rechaza la subida y elimina el temporal.
Los adjuntos sembrados forman parte del repositorio y no necesitan reanalizarse en cada restauración.

## 12. Logs, privacidad y observación

La demo conserva logs técnicos durante siete días mediante rotación. Los logs no incluyen cuerpos
de tickets o comentarios, contraseñas, tokens, cookies, contenido de adjuntos ni claves del motor.
No se conservan copias de la base ni de adjuntos de visitantes después de la restauración.

La salud distingue aplicación disponible, migración fallida y dependencia no preparada. La
comprobación pública no revela versiones de servicios, secretos, rutas internas ni datos de uso.

## 13. Reparto de archivos previsto

El repositorio público puede crear o modificar, dentro del alcance aprobado:

- backend para aviso de entorno compartido, cuotas configurables, analizador antivirus opcional y
  protección configurable de cuentas compartidas;
- frontend para el banner genérico, mensajes de límites y ocultación de acciones bloqueadas;
- pruebas unitarias, de integración y de navegador del comportamiento anterior;
- los documentos enumerados en la sección 2.

El repositorio privado `Catalina-Support/demo` contendrá:

- `demo.yml` y el Compose dedicado de IA;
- plantilla de entorno sin secretos y referencia a la ubicación privada de los valores reales;
- configuración nginx del vhost;
- scripts de construcción, publicación, restauración, rollback y rotación;
- cron, ranuras, límites concretos y lista de cuentas públicas protegidas;
- configuración de ClamAV y del monitor;
- manifiesto con el commit público desplegado;
- pruebas operativas que no revelen secretos ni topología en sus resultados.

Los secretos reales, el correo del monitor, la configuración de DNS y el certificado emitido no se
versionan ni siquiera en el repositorio privado. Se mantienen como archivos del host con permisos
restringidos o en el almacén de secretos que se adopte.

## 14. Verificación previa a abrir el dominio

Antes de publicar tráfico se comprobará:

1. que ningún nombre, volumen, puerto o archivo de entorno apunta a desarrollo o producción;
2. que el manifiesto privado fija un commit existente de `main` y el artefacto corresponde a él;
3. que una instalación vacía llega al estado sembrado conocido;
4. que las once cuentas entran y `admin` no aparece publicado;
5. que cambio y recuperación de contraseña se bloquean sólo para las cuentas públicas;
6. que Mailpit, base, IA, antivirus y puertos internos no responden desde Internet;
7. que la IA genera texto y respeta ambos límites;
8. que un adjunto limpio se acepta y uno de prueba EICAR se rechaza y elimina;
9. que el banner aparece en entrada y sesión, en PC y móvil, en inglés;
10. que HTTP redirige a HTTPS y el certificado es válido;
11. que una restauración correcta reemplaza todos los cambios y sesiones;
12. que una restauración fallida conserva la ranura anterior;
13. que los logs rotan a siete días y no contienen contenido sensible;
14. que el monitor detecta una caída controlada y entrega el aviso;
15. que desarrollo y producción no cambiaron durante todo el recorrido.

El dominio sólo se abre después de completar el recorrido. La evidencia registra commit, hora,
conteos, ranura activa y resultados sin copiar secretos.

## 15. Fuera de alcance

Esta propuesta no:

- despliega ni abre la producción de Catalina Support;
- crea la etiqueta `v1.0.0`;
- publica `admin`, Mailpit, PostgreSQL, el motor o su administrador;
- envía correo real ni usa proveedores externos de IA;
- promete disponibilidad, conservación de datos o soporte a visitantes;
- reutiliza secretos, volúmenes o copias de producción;
- publica los archivos operativos, la topología o los valores concretos de la demo;
- mantiene un fork privado o parches sobre el código público;
- cambia permisos funcionales fuera del modo demo.

## 16. Registro de decisiones

El responsable eligió el 2026-10-20: 1A aislamiento propio; 2A restauración diaria; 3A Mailpit
privado; 4A restauración completa; 5A IA local 1.5B; 6A administrador secreto; 7A 03:00 de
`America/Guayaquil`; 8A mismo servidor con recursos exclusivos; 9A actualización manual desde
`main`; 10A límites contra abuso; 11A banner persistente; 12A sin copias y logs por siete días; 13A
inglés; 14A acceso local; 15A nginx y Let's Encrypt; 16A sin SLA, con monitor; 17A conmutación
atómica; 18A límites recomendados; 19A monitor externo; 20A construcción determinista; 21B bloqueo
de contraseña; 22A bloqueo sólo para cuentas públicas; 23A antivirus; y 24A conservación del último
estado funcional ante fallos.

El 2026-10-21 el responsable eligió la separación híbrida 1A: capacidades genéricas en el producto
público e infraestructura específica en el repositorio privado `Catalina-Support/demo`, disponible
en `/root/dev/catalina-support/demo`.

No quedan decisiones abiertas. La implementación requiere aprobación explícita de este documento.
