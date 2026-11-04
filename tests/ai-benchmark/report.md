# Banco de pruebas de modelos locales

Este informe se genera desde `results/*.json`. Cada respuesta se contrasta con el texto que
recibió el modelo: JSON, idioma, palabras, hechos esperados, nombres/números inventados y
revisión humana de la frase.

## Resumen

| Modelo | Casos | Correctos | Latencia media | Memoria del motor |
| --- | ---: | ---: | ---: | ---: |
| Qwen2.5-1.5B · 2 hilos (actual) | 8 | 7/8 | 4008 ms | 1,191 GiB |
| Qwen2.5-1.5B · 4 hilos | 8 | 6/8 | 2434 ms | 1,191 GiB |
| Qwen2.5-1.5B · 8 hilos | 8 | 7/8 | 1954 ms | 1,191 GiB |
| Qwen2.5-1.5B · 16 hilos | 8 | 7/8 | 1892 ms | 1,191 GiB |
| Qwen2.5-1.5B · 32 hilos | 8 | 7/8 | 2522 ms | 1,191 GiB |
| Gemma 3 1B · 2 hilos | 8 | 5/8 | 4391 ms | 880 MiB |
| Qwen3-0.6B · razonamiento por defecto | 8 | 4/8 | 11291 ms | 717 MiB |
| Qwen3-0.6B · razonamiento desactivado | 8 | 0/8 | 1367 ms | — |
| Qwen3-1.7B · razonamiento por defecto | 8 | 3/8 | 26297 ms | 1,576 GiB |
| Qwen3-1.7B · razonamiento desactivado | 8 | 5/8 | 4096 ms | — |
| Qwen2.5-0.5B · 2 hilos | 8 | 3/8 | 1649 ms | 493 MiB |

## Cada respuesta

## Entradas entregadas al modelo

El texto siguiente es el que recibió cada modelo, además del mensaje de sistema y de la
exigencia de formato. Se reproduce para poder contrastar cada respuesta con su entrada.

### es-summary-password

```text
[TARGET_NUMBER] CS-2026-0101
[TARGET_STATE] en progreso
[TARGET_REASON] Desde ayer, al intentar entrar en la aplicación, aparece «usuario o contraseña incorrectos». No he cambiado la contraseña.
[TARGET_LAST_EVENT] estado
[TARGET_LAST_ACTOR] Soporte Uno
[TARGET_LAST_STATE] en progreso
[TARGET_LAST_COMMENT] 

Requester: Ana Pérez
Created by: Ana Pérez
Main assignee: Soporte Uno

MAIN: CS-2026-0101 (en progreso)
Subject: No puedo entrar en la aplicación
Description: Desde ayer, al intentar entrar en la aplicación, aparece «usuario o contraseña incorrectos». No he cambiado la contraseña.

Comments:
- Soporte Uno (comentario): ¿Puedes probar desde otro navegador?

Category: Software
Priority: medium
```

### es-last-router

```text
[TARGET_NUMBER] CS-2026-0102
[TARGET_STATE] resuelto
[TARGET_REASON] La conexión se corta cada pocos minutos desde esta mañana.
[TARGET_LAST_EVENT] resuelto
[TARGET_LAST_ACTOR] Juan Carlos Estevez Ayala
[TARGET_LAST_STATE] resuelto
[TARGET_LAST_COMMENT] Reiniciamos el router y, por el momento, la conexión parece funcionar correctamente.

Requester: Ana Pérez
Created by: Ana Pérez
Main assignee: Juan Carlos Estevez Ayala

MAIN: CS-2026-0102 (resuelto)
Subject: La conexión se corta
Description: La conexión se corta cada pocos minutos desde esta mañana.

Comments:
- Juan Carlos Estevez Ayala (comentario): Reiniciamos el router y, por el momento, la conexión parece funcionar correctamente.

Category: Red
Priority: high
```

### en-summary-vpn

```text
[TARGET_NUMBER] CS-2026-0103
[TARGET_STATE] en progreso
[TARGET_REASON] The VPN client shows "certificate expired" when connecting from home. The user needs a renewed certificate.
[TARGET_LAST_EVENT] estado
[TARGET_LAST_ACTOR] Support One
[TARGET_LAST_STATE] en progreso
[TARGET_LAST_COMMENT] 

Requester: John Smith
Created by: John Smith
Main assignee: Support One

MAIN: CS-2026-0103 (en progreso)
Subject: VPN certificate expired
Description: The VPN client shows "certificate expired" when connecting from home. The user needs a renewed certificate.

Comments:
- Support One (comment): Please send a screenshot of the error.

Category: Network
Priority: high
```

### en-last-escalation

```text
[TARGET_NUMBER] INT-CS-2026-0103
[TARGET_STATE] escalado
[TARGET_REASON] The session service returns 502 when signing in; Support already checked the network.
[TARGET_LAST_EVENT] escalado
[TARGET_LAST_ACTOR] Marta Ruiz
[TARGET_LAST_STATE] escalado
[TARGET_LAST_COMMENT] The session service returns 502; this needs a code change.

Requester: John Smith
Created by: Support One
Internal created by: Marta Ruiz

INTERNAL: INT-CS-2026-0103 (escalado)
Subject: VPN certificate expired
Escalation reason: The session service returns 502 when signing in; Support already checked the network.

Category: Network
Priority: high
```

### es-summary-internal

```text
[TARGET_NUMBER] INT-CS-2026-0104
[TARGET_STATE] en progreso
[TARGET_REASON] El servidor de sesiones devuelve 502 al iniciar sesión; Soporte ya comprobó la red.
[TARGET_LAST_EVENT] estado
[TARGET_LAST_ACTOR] Desarrollo Uno
[TARGET_LAST_STATE] en progreso
[TARGET_LAST_COMMENT] 

Requester: Luis Mora
Created by: Soporte Dos
Internal created by: Soporte Dos
Internal assignee: Desarrollo Uno

INTERNAL: INT-CS-2026-0104 (en progreso)
Subject: No puedo iniciar sesión
Escalation reason: El servidor de sesiones devuelve 502 al iniciar sesión; Soporte ya comprobó la red.

Category: Software
Priority: critical
```

### en-last-internal-resolved

```text
[TARGET_NUMBER] INT-CS-2026-0105
[TARGET_STATE] resuelto
[TARGET_REASON] The report service returns duplicated rows after the last deployment.
[TARGET_LAST_EVENT] resuelto
[TARGET_LAST_ACTOR] Dev One
[TARGET_LAST_STATE] resuelto
[TARGET_LAST_COMMENT] Updated the session configuration and restarted the service.

Requester: Laura Díaz
Created by: Support Two
Internal created by: Support Two
Internal assignee: Dev One

INTERNAL: INT-CS-2026-0105 (resuelto)
Subject: Duplicated rows in the report
Escalation reason: The report service returns duplicated rows after the last deployment.

Category: Software
Priority: high
```

### es-summary-printer

```text
[TARGET_NUMBER] CS-2026-0106
[TARGET_STATE] nuevo
[TARGET_REASON] La impresora de la segunda planta no imprime desde esta mañana; enciende pero los trabajos quedan en cola.
[TARGET_LAST_EVENT] creado
[TARGET_LAST_ACTOR] Carlos Núñez
[TARGET_LAST_STATE] nuevo
[TARGET_LAST_COMMENT] 

Requester: Carlos Núñez
Created by: Carlos Núñez

MAIN: CS-2026-0106 (nuevo)
Subject: La impresora no imprime
Description: La impresora de la segunda planta no imprime desde esta mañana; enciende pero los trabajos quedan en cola.

Category: Impresoras
Priority: medium
```

### en-last-user-reply

```text
[TARGET_NUMBER] CS-2026-0107
[TARGET_STATE] en espera
[TARGET_REASON] The application closes unexpectedly when opening the reports section.
[TARGET_LAST_EVENT] comentario
[TARGET_LAST_ACTOR] Ana Pérez
[TARGET_LAST_STATE] en espera
[TARGET_LAST_COMMENT] I tried from another browser and the error persists.

Requester: Ana Pérez
Created by: Ana Pérez
Main assignee: Support One

MAIN: CS-2026-0107 (en espera)
Subject: The application closes in Reports
Description: The application closes unexpectedly when opening the reports section.

Comments:
- Ana Pérez (comment): I tried from another browser and the error persists.

Category: Software
Priority: high
```

### Qwen2.5-1.5B · 2 hilos (actual)

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 4892 ms | 18 | sí | es | 2/2 | — | Desde ayer, al intentar entrar en la aplicación, aparece 'usuario o contraseña incorrectos'. No he cambiado la contraseña. |
| es-last-router | 4474 ms | 13 | sí | es | 3/3 | — | Juan Carlos Estevez Ayala reinició el router y la conexión parece funcionar correctamente. |
| en-summary-vpn | 4112 ms | 16 | sí | en | 3/3 | — | The VPN client shows 'certificate expired' when connecting from home. The user needs a renewed certificate. |
| en-last-escalation | 4685 ms | 22 | sí | en | 3/3 | — | Marta Ruiz escalated the issue due to the session service returning 502 when signing in, which was already checked on the network. |
| es-summary-internal | 4172 ms | 23 | sí | es | 3/3 | — | El usuario Luis Mora no puede iniciar sesión debido a un servidor de sesiones que devuelve 502; la red fue comprobada por Soporte. |
| en-last-internal-resolved | 3398 ms | 23 | sí | en | 2/3 | — | Dev One updated the session configuration and restarted the service to resolve the duplicated rows in the report service after the last deployment. — **La frase es correcta; el evaluador no aceptaba «resolve» como «resolved».** |
| es-summary-printer | 3596 ms | 14 | sí | es | 2/2 | — | La impresora no imprime desde esta mañana; enciende pero los trabajos quedan en cola. |
| en-last-user-reply | 2733 ms | 10 | sí | en | 2/2 | — | Ana Pérez tried from another browser and the error persists. |

### Qwen2.5-1.5B · 4 hilos

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 3051 ms | 18 | sí | es | 2/2 | — | Desde ayer, al intentar entrar en la aplicación, aparece 'usuario o contraseña incorrectos'. No he cambiado la contraseña. |
| es-last-router | 2560 ms | 13 | sí | es | 3/3 | — | Juan Carlos Estevez Ayala reinició el router y la conexión parece funcionar correctamente. |
| en-summary-vpn | 2608 ms | 22 | sí | en | 3/3 | — | John Smith created a ticket for a VPN client showing a 'certificate expired' error when connecting from home, requesting a renewed certificate. |
| en-last-escalation | 2238 ms | 12 | sí | en | 1/3 | — | The session service returned 502 when signing in, requiring a code change. |
| es-summary-internal | 2822 ms | 30 | sí | es | 3/3 | — | El usuario Luis Mora no puede iniciar sesión debido a un servidor de sesiones que devuelve 502 al iniciar sesión, y el Soporte ya confirmó que la red está correcta. |
| en-last-internal-resolved | 2045 ms | 23 | sí | en | 2/3 | — | Dev One updated the session configuration and restarted the service to resolve the duplicated rows in the report service after the last deployment. |
| es-summary-printer | 2213 ms | 18 | sí | es | 2/2 | — | La impresora de la segunda planta no imprime desde esta mañana; enciende pero los trabajos quedan en cola. |
| en-last-user-reply | 1937 ms | 16 | sí | en | 2/2 | — | Ana Pérez tried from another browser and the error persisted, indicating the issue is not browser-related. |

### Qwen2.5-1.5B · 8 hilos

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 2471 ms | 18 | sí | es | 2/2 | — | Desde ayer, al intentar entrar en la aplicación, aparece 'usuario o contraseña incorrectos'. No he cambiado la contraseña. |
| es-last-router | 2112 ms | 13 | sí | es | 3/3 | — | Juan Carlos Estevez Ayala reinició el router y la conexión parece funcionar correctamente. |
| en-summary-vpn | 1893 ms | 17 | sí | en | 3/3 | — | John Smith requested a renewed certificate for the VPN client showing 'certificate expired' when connecting from home. |
| en-last-escalation | 2107 ms | 20 | sí | en | 3/3 | — | Marta Ruiz escalated the ticket due to the session service returning 502 when signing in, which needed a code change. |
| es-summary-internal | 2425 ms | 30 | sí | es | 3/3 | — | El usuario Luis Mora no puede iniciar sesión debido a un servidor de sesiones que devuelve 502 al iniciar sesión, y Soporte ya confirmó que la red está en orden. |
| en-last-internal-resolved | 1221 ms | 10 | sí | en | 1/3 | — | The report service returned duplicated rows after the last deployment. — **Repite el motivo en vez de la última acción (Dev One actualizó la configuración y reinició el servicio).** |
| es-summary-printer | 2017 ms | 18 | sí | es | 2/2 | — | La impresora de la segunda planta no imprime desde esta mañana; enciende pero los trabajos quedan en cola. |
| en-last-user-reply | 1382 ms | 10 | sí | en | 2/2 | — | Ana Pérez tried from another browser and the error persisted. |

### Qwen2.5-1.5B · 16 hilos

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 2534 ms | 18 | sí | es | 2/2 | — | Desde ayer, al intentar entrar en la aplicación, aparece 'usuario o contraseña incorrectos'. No he cambiado la contraseña. |
| es-last-router | 2037 ms | 13 | sí | es | 3/3 | — | Juan Carlos Estevez Ayala reinició el router y la conexión parece funcionar correctamente. |
| en-summary-vpn | 1848 ms | 16 | sí | en | 3/3 | — | The VPN client shows 'certificate expired' when connecting from home. The user needs a renewed certificate. |
| en-last-escalation | 1931 ms | 18 | sí | en | 3/3 | — | Marta Ruiz escalated the issue to a code change after the session service returned 502 when signing in. |
| es-summary-internal | 2458 ms | 27 | sí | es | 3/3 | — | El usuario Luis Mora no puede iniciar sesión debido a un servidor de sesiones que devuelve 502 al iniciar sesión; la red ha sido comprobada por Soporte. |
| en-last-internal-resolved | 1233 ms | 10 | sí | en | 1/3 | — | The report service returned duplicated rows after the last deployment. |
| es-summary-printer | 1838 ms | 18 | sí | es | 2/2 | — | La impresora de la segunda planta no imprime desde esta mañana; enciende pero los trabajos quedan en cola. |
| en-last-user-reply | 1260 ms | 10 | sí | en | 2/2 | — | Ana Pérez tried from another browser and the error persisted. |

### Qwen2.5-1.5B · 32 hilos

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 2802 ms | 18 | sí | es | 2/2 | — | Desde ayer, al intentar entrar en la aplicación, aparece 'usuario o contraseña incorrectos'. No he cambiado la contraseña. |
| es-last-router | 2105 ms | 13 | sí | es | 3/3 | — | Juan Carlos Estevez Ayala reinició el router y la conexión parece funcionar correctamente. |
| en-summary-vpn | 3644 ms | 22 | sí | en | 3/3 | — | John Smith created a ticket about the VPN client showing a 'certificate expired' error when connecting from home, needing a renewed certificate. |
| en-last-escalation | 2336 ms | 20 | sí | en | 3/3 | — | Marta Ruiz escalated the ticket due to the session service returning 502 when signing in, which needed a code change. |
| es-summary-internal | 2967 ms | 26 | sí | es | 3/3 | — | El usuario Luis Mora no puede iniciar sesión debido a un servidor de sesiones que devuelve 502 al iniciar sesión; el Soporte ya comprobó la red. |
| en-last-internal-resolved | 1443 ms | 10 | sí | en | 1/3 | — | The report service returned duplicated rows after the last deployment. |
| es-summary-printer | 3051 ms | 18 | sí | es | 2/2 | — | La impresora de la segunda planta no imprime desde esta mañana; enciende pero los trabajos quedan en cola. |
| en-last-user-reply | 1832 ms | 16 | sí | en | 2/2 | — | Ana Pérez tried from another browser and the error persisted, indicating the issue is not browser-specific. |

### Gemma 3 1B · 2 hilos

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 4852 ms | 26 | sí | es | 1/2 | — | Ana Pérez reporta que no puede acceder a la aplicación debido a errores de autenticación, solicitando asistencia del Soporte Uno para diagnosticar y resolver el problema. — **No menciona la contraseña; dice «errores de autenticación».** |
| es-last-router | 4428 ms | 16 | sí | es | 3/3 | — | Juan Carlos Estevez Ayala reinició el router y, por el momento, la conexión parece funcionar correctamente |
| en-summary-vpn | 4631 ms | 33 | sí | en | 3/3 | — | John Smith is requesting a renewed VPN certificate due to an expired certificate displayed on the VPN client when connecting from home, and Support One is tasked with investigating and resolving the issue. |
| en-last-escalation | 4565 ms | 28 | sí | en | 3/3 | — | Marta Ruiz escalated the issue due to the session service returning a 502 error when signing in, requiring a code change based on the previous event and comment. |
| es-summary-internal | 4900 ms | 30 | sí | es | 3/3 | — | El solicitante Luis Mora ha reportado que no puede iniciar sesión debido a un error 502 en el servidor de sesiones, y el soporte ya ha descartado problemas de red. |
| en-last-internal-resolved | 3784 ms | 22 | sí | en | 2/3 | — | The Dev One team updated the session configuration and restarted the service to resolve the duplicated rows issue after the last deployment. — **Atribuye la acción a «Dev One team»; el evaluador no acepta «resolve» como «resolved».** |
| es-summary-printer | 4289 ms | 23 | sí | es | 2/2 | — | Carlos Núñez solicitó una impresora nueva en la segunda planta, que no imprime desde esta mañana debido a problemas de cola de trabajo. — **Inventa que se «solicitó una impresora nueva», que no está en el texto.** |
| en-last-user-reply | 3677 ms | 19 | sí | en | 1/2 | — | Ana Pérez attempted to open the reports section from a different browser, but the application continued to close unexpectedly. — **Resume bien; el evaluador no acepta «attempted» como «tried».** |

### Qwen3-0.6B · razonamiento por defecto

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 13082 ms | 0 | no | desconocido | 0/2 | — | (sin texto) — **JSON truncado: se queda a mitad de frase.** |
| es-last-router | 11236 ms | 23 | sí | es | 3/3 | — | El ticket CS-2026-0102 ha sido resuelto, con la acción de Juan Carlos Estevez Ayala reiniciando el router, por el momento asegurando la conexión. |
| en-summary-vpn | 10515 ms | 24 | sí | en | 3/3 | — | CS-2026-0103 (en progreso) Subject: VPN certificate expired Description: The VPN client shows 'certificate expired' when connecting from home. The user needs a renewed certificate. |
| en-last-escalation | 12151 ms | 0 | no | desconocido | 0/3 | — | (sin texto) — **JSON truncado antes de cerrar.** |
| es-summary-internal | 10944 ms | 25 | sí | es | 3/3 | Ticket | Ticket INT-CS-2026-0104 está en progreso, con motivo de que el servidor de sesiones devuelve 502 al iniciar sesión, y el soporte ya comprobó la red. |
| en-last-internal-resolved | 10195 ms | 16 | sí | en | 1/3 | — | The report service was updated and restarted after the last deployment, resolving the duplicated rows issue. |
| es-summary-printer | 11486 ms | 20 | sí | es | 2/2 | — | La impresora de la segunda planta no imprime desde esta mañana, con trabajos en cola en la planta de segunda. |
| en-last-user-reply | 10716 ms | 10 | sí | en | 2/2 | — | Ana Pérez tried from another browser and the error persists. |

### Qwen3-0.6B · razonamiento desactivado

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 1654 ms | 7 | sí | desconocido | 0/2 | — | Ana Pérez, ¿puedes probar desde otro navegador? — **Responde con una pregunta del comentario, no con un resumen.** |
| es-last-router | 1649 ms | 10 | sí | es | 1/3 | — | La conexión se cortó y se reinició por el momento. |
| en-summary-vpn | 1505 ms | 3 | sí | es | 0/3 | — | CS-2026-0103 (en progreso) — **Devuelve sólo el número y el estado; no resume.** |
| en-last-escalation | 1207 ms | 3 | sí | desconocido | 1/3 | — | Marta Ruiz escalado |
| es-summary-internal | 1556 ms | 20 | sí | es | 2/3 | — | Luis Mora no puede iniciar sesión debido a un error en el servidor de sesiones; el estado es en progreso. |
| en-last-internal-resolved | 1019 ms | 8 | sí | en | 1/3 | — | Updated the session configuration and restarted the service. |
| es-summary-printer | 1289 ms | 14 | sí | es | 1/2 | — | Carlos Núñez necesita que la impresora de la segunda planta imprima desde esta mañana. |
| en-last-user-reply | 1055 ms | 9 | sí | en | 0/2 | — | The application closes unexpectedly when opening the reports section. |

### Qwen3-1.7B · razonamiento por defecto

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 26912 ms | 23 | sí | es | 2/2 | — | Ana Pérez está teniendo problemas al intentar entrar en la aplicación por 'usuario o contraseña incorrectos', ya que no ha cambiado la contraseña. |
| es-last-router | 28924 ms | 0 | no | desconocido | 0/3 | — | (sin texto) — **Sin contenido: el razonamiento agota el tope y no llega a escribir.** |
| en-summary-vpn | 27253 ms | 20 | sí | en | 3/3 | — | The requester needs a renewed certificate for the VPN client to resolve the 'certificate expired' error when connecting from home. |
| en-last-escalation | 27182 ms | 0 | no | desconocido | 0/3 | — | (sin texto) — **Sin contenido: el razonamiento agota el tope y no llega a escribir.** |
| es-summary-internal | 25906 ms | 0 | no | desconocido | 0/3 | — | (sin texto) — **Sin contenido: el razonamiento agota el tope y no llega a escribir.** |
| en-last-internal-resolved | 25529 ms | 0 | no | desconocido | 0/3 | — | (sin texto) — **Sin contenido: el razonamiento agota el tope y no llega a escribir.** |
| es-summary-printer | 23516 ms | 20 | sí | es | 2/2 | — | La impresora de la segunda planta no imprime, enciende pero los trabajos quedan en cola; el requester es Carlos Núñez. |
| en-last-user-reply | 25157 ms | 0 | no | desconocido | 0/2 | — | (sin texto) — **Sin contenido: el razonamiento agota el tope y no llega a escribir.** |

### Qwen3-1.7B · razonamiento desactivado

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 7019 ms | 47 | sí | es | 2/2 | — | El ticket indica que el usuario Ana Pérez está enfrentando el problema de que al intentar entrar en la aplicación se muestra el mensaje «usuario o contraseña incorrectos», y no ha cambiado la contraseña. El soporte está en progreso y ha solicitado una prueba en otro navegador. |
| es-last-router | 4801 ms | 16 | sí | es | 3/3 | — | Juan Carlos Estevez Ayala reinició el router y, por el momento, la conexión parece funcionar correctamente |
| en-summary-vpn | 4111 ms | 17 | sí | en | 3/3 | — | John Smith needs a renewed VPN certificate to resolve the 'certificate expired' error when connecting from home. |
| en-last-escalation | 3799 ms | 13 | sí | en | 1/3 | — | The session service returns 502 when signing in; Support already checked the network. |
| es-summary-internal | 3514 ms | 14 | sí | es | 3/3 | — | El servidor de sesiones devuelve 502 al iniciar sesión; Soporte ya comprobó la red. |
| en-last-internal-resolved | 2626 ms | 10 | sí | en | 2/3 | — | Dev One updated the session configuration and restarted the service |
| es-summary-printer | 3492 ms | 17 | sí | es | 2/2 | — | La impresora de la segunda planta no imprime, mostrando que enciende pero los trabajos quedan en cola. |
| en-last-user-reply | 3409 ms | 22 | sí | en | 2/2 | — | Ana Pérez reported that the application closes unexpectedly when opening the reports section, and tried from another browser but the error persists. |

### Qwen2.5-0.5B · 2 hilos

| Caso | Lat. | Pal. | JSON | Idioma | Hechos | Inventado | Respuesta / nota |
| --- | ---: | ---: | :---: | :---: | :---: | :---: | --- |
| es-summary-password | 2285 ms | 11 | sí | es | 1/2 | Para | Para Ana Pérez, la aplicación no puede entrar en la progreso. |
| es-last-router | 1733 ms | 4 | sí | es | 0/3 | — | La conexión se corta — **Repite el motivo («La conexión se corta»), no la última acción.** |
| en-summary-vpn | 1730 ms | 16 | sí | en | 3/3 | — | The VPN client shows 'certificate expired' when connecting from home. The user needs a renewed certificate. |
| en-last-escalation | 1590 ms | 10 | sí | en | 1/3 | — | The session service returns 502; this needs a code change. — **Repite el motivo, no la escalación.** |
| es-summary-internal | 1793 ms | 20 | sí | es | 3/3 | — | Luis Mora, Desarrollo Uno, en progreso, el servidor de sesiones devuelve 502 al iniciar sesión; Soporte ya comprobó la red. |
| en-last-internal-resolved | 1228 ms | 10 | sí | en | 1/3 | — | The report service returns duplicated rows after the last deployment. — **Repite el motivo, no la última acción.** |
| es-summary-printer | 1568 ms | 6 | sí | es | 2/2 | — | Carlos Núñez: La impresora no imprime. |
| en-last-user-reply | 1266 ms | 9 | sí | en | 0/2 | — | The application closes unexpectedly when opening the reports section. — **Repite el motivo, no la respuesta del usuario.** |

## Lectura

- **Qwen2.5-1.5B con 2 hilos** es el mejor equilibrio actual: 7/8 correctos y 3,4-4,9 s.
- **Subir a 8 hilos** baja la latencia a 1,2-2,4 s sin cambiar el modelo ni el prompt. La
  máquina tiene 32 núcleos y el motor usa 2; es la mejora de velocidad más clara y con menos
  riesgo. 16 y 32 hilos no mejoran: el cuello deja de ser la CPU.
- **Gemma 3 1B** es válida en formato y ocupa menos memoria, pero no es más rápida y comete
  errores de contenido (la impresora «nueva»).
- **Qwen3-0.6B** con razonamiento por defecto tarda 10-13 s y trunca el JSON; sin razonamiento
  baja a 1-1,6 s pero falla casi todos los casos por contenido e idioma.
- **Qwen3-1.7B** con razonamiento por defecto tarda 25-29 s y deja respuestas vacías; sin
  razonamiento mejora, pero ocupa 1,576 GiB (por encima del techo de 1,5 GB) y sigue fallando
  la última acción en dos casos.
- **Qwen2.5-0.5B** es rapidísimo (1,2-2,3 s) y cabe en 493 MiB, pero en los casos de última
  acción repite el motivo y no cuenta lo que pasó.

**Conclusión:** no hay un cambio de modelo que sea a la vez más rápido y de calidad igual o
mejor. La mejora recomendada es mantener Qwen2.5-1.5B y subir los hilos de 2 a 8, midiendo en
desarrollo antes de tocar producción. Si se quiere explorar un modelo, el siguiente paso es
Qwen3-1.7B sin razonamiento y con el techo de memoria en 2 GiB, pero necesita más pruebas.
