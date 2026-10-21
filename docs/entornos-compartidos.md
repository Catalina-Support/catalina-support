# Entornos compartidos

> **Estado:** as-built
> **Última actualización:** 2026-10-21

## 1. Alcance

Catalina Support puede presentar y proteger una instalación pública o compartida sin incorporar al
repositorio público su topología ni sus valores operativos. Todas estas capacidades son opt-in por
variables de entorno; si no se configuran, el comportamiento de una instalación normal no cambia.

Este documento cubre únicamente mecanismos reutilizables del producto. Dominios, cuentas concretas,
Compose, nginx, calendarios, ranuras, secretos y procedimientos de una instalación pertenecen a su
propio repositorio operativo.

## 2. Presentación pública

`SHARED_ENVIRONMENT=true` marca la instalación como compartida. La marca pública entrega el aviso y
la referencia de restauración configurados mediante `SHARED_ENVIRONMENT_NOTICE` y
`SHARED_ENVIRONMENT_RESET`. El frontend enseña el aviso antes de entrar y de forma persistente en el
armazón autenticado. En este modo oculta los formularios de recuperación y cambio de contraseña.

## 3. Cuentas compartidas

`PROTECTED_EMAIL_SUFFIX` protege las cuentas cuyo correo termina en ese sufijo. El backend no emite
recuperaciones para ellas y rechaza restablecer o cambiar su contraseña. La cuenta de fábrica queda
fuera porque no pertenece a la base. Vacío desactiva la regla.

## 4. Cuotas

Los límites configurables son:

- `LIMIT_LOGIN_PER_MINUTE`, por IP;
- `LIMIT_AI_PER_MINUTE` y `LIMIT_AI_PER_DAY`, por cuenta;
- `LIMIT_COMMENTS_PER_MINUTE`, por cuenta;
- `LIMIT_ATTACHMENT_BYTES_PER_DAY`, por IP y tamaño declarado;
- `LIMIT_EXPORTS_PER_HOUR`, por IP.

Cero o ausencia desactiva cada límite. Los contadores viven en memoria y una denegación responde 429
con `Retry-After`; este mecanismo está diseñado para una sola réplica. Un proxy público sigue siendo
responsable de limitar tamaño de petición y de entregar al backend la IP real confiable.

## 5. Análisis de adjuntos

`ANTIVIRUS_ADDRESS` conecta el protocolo INSTREAM de clamd. El servicio lee el archivo completo antes
de escribirlo en el almacén y sólo acepta la respuesta limpia. Un virus, una respuesta desconocida,
un error de lectura o la caída del analizador rechazan la subida; sin dirección configurada se
conserva el flujo normal del producto.

## 6. Verificación

Las pruebas unitarias cubren cuotas habilitadas y deshabilitadas, coincidencia del sufijo protegido y
respuestas limpia, infectada e indeterminada de ClamAV. La suite backend, `go vet`, las 253 pruebas del
frontend y la compilación de producción pasan con estas capacidades.
