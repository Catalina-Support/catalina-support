# Demo accounts

[Español](../es/demo-accounts.md) · **English**

The development seeders (`seed-en` and `seed-es`) create these **11 accounts** with the same
development password: **`123123123`**. They are demo credentials only; never use them in production.

| Name | Email | Role |
| --- | --- | --- |
| User One | `user1@demo.com` | User |
| User Two | `user2@demo.com` | User |
| User Three | `user3@demo.com` | User |
| User Four | `user4@demo.com` | User |
| User Five | `user5@demo.com` | User |
| Support One | `support1@demo.com` | Support |
| Support Two | `support2@demo.com` | Support |
| Support Three | `support3@demo.com` | Support |
| Developer One | `dev1@demo.com` | Development |
| Developer Two | `dev2@demo.com` | Development |
| Developer Three | `dev3@demo.com` | Development |

The factory administrator **`admin` is not part of the seed** and its password is not published.
In development it uses `ADMIN_PASSWORD` from `config/env/dev.env`. In production it uses
`ADMIN_PASSWORD` from `config/env/prod.env`; that file is not versioned and has no default password.
