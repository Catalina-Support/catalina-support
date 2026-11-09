# Cuentas de demo

[English](../en/demo-accounts.md) · **Español**

Los seeders de desarrollo (`seed-en` y `seed-es`) crean estas **11 cuentas** con la misma contraseña
de desarrollo: **`123123123`**. Son credenciales de demo; nunca las uses en producción.

| Nombre | Correo | Papel |
| --- | --- | --- |
| Usuario Uno | `user1@demo.com` | Usuario |
| Usuario Dos | `user2@demo.com` | Usuario |
| Usuario Tres | `user3@demo.com` | Usuario |
| Usuario Cuatro | `user4@demo.com` | Usuario |
| Usuario Cinco | `user5@demo.com` | Usuario |
| Soporte Uno | `support1@demo.com` | Soporte |
| Soporte Dos | `support2@demo.com` | Soporte |
| Soporte Tres | `support3@demo.com` | Soporte |
| Desarrollador Uno | `dev1@demo.com` | Desarrollo |
| Desarrollador Dos | `dev2@demo.com` | Desarrollo |
| Desarrollador Tres | `dev3@demo.com` | Desarrollo |

La cuenta de fábrica **`admin` no forma parte del seed** y su contraseña no se publica. En desarrollo
usa `ADMIN_PASSWORD` de `config/env/dev.env`. En producción usa `ADMIN_PASSWORD` de
`config/env/prod.env`; ese archivo no se versiona y no tiene contraseña por defecto.
