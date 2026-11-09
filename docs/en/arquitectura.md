# Architecture

[Español](../es/arquitectura.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-05
>
> **WSL2 support verified, 2026-11-05.** The CPU/NVIDIA boundary of section 28 also works on Windows
> with Docker Desktop WSL2: GPU visible with `--gpus all`, administrator in `nvidia` mode, active
> model and real summaries. Linux with NVIDIA Container Toolkit remains the other path; native
> Windows containers stay out.
>
> **Amendment approved, implemented and verified on 2026-10-22: explicit bilingual examples.** Section
> 32 separates the demonstration data into English and Spanish. The starting finding is that the only
> current seeder stores in Spanish all the content written by people, even when the installation and
> the interface are in English.
>
> **Amendment approved, implemented and verified on 2026-10-20.** Section 31 defines the Go migrations
> runner, its transactional record, checksums, locking and boundary with startup. The owner closed and
> approved decisions 1A–25A. Production was not deployed.
>
> **Amendment approved, implemented and verified on 2026-10-20: SMTP test with delivery to the relay.**
> Section 30 assigns to `settings` the entry, authorization and rate limiting of the test, and to
> `mail` the fixed message and the sending. The owner closed and approved 1A–14A; the boundaries,
> focused tests and compilation were verified.
>
> **Amendment approved, implemented and verified on 2026-10-19: insecure SMTP transport boundary.**
> Section 29 models three transports and encloses unencrypted SMTP behind a persisted acknowledgement
> and a private-destination check on every connection. The owner closed and approved 1A–10A; the
> boundaries and the real relay were verified.
>
> **Approved, implemented and verified on 2026-10-18.** Section 28 separates the CPU/NVIDIA images,
> persists the mode in `ai_settings` and makes the restart of the local administrator transactional.
>
> **Implemented and verified on 2026-10-18.** Section 27 adds `dashboard` as owner of aggregates and
> export, fed by a read contract from `tickets`.
>
> **Implemented on 2026-10-18.** Section 26 coordinates the migration from an identity with an
> exclusive role to an operational role plus administrative capability, keeping authorization in
> `shared/authz` and account management in `users`.
>
> **Amendment approved, implemented and verified on 2026-10-15.** Section 25 migrates the
> `motivo/reason` contract to `summary/escalationReason`, removes manual writing and versions a full
> regeneration of Summary. The owner chose 1A–20B, accepted the potential provider cost and left no
> open decisions. The open migration was applied only in development; the specific tests, the 231
> frontend tests, the boundaries and the compilation passed. The full backend suite keeps the previous
> internal-priority failure, unrelated to this amendment.
>
> **Amendment approved, implemented and verified on 2026-10-15.** Section 24 adds the persistent link
> between history and comment, a context version for idempotent regenerations and a bounded background
> job. It creates no modules and crosses no boundaries: `tickets` assembles the facts and `ai` writes
> and validates. The owner chose 1A–14A and explicitly approved the documents; no open decisions
> remain.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** Section 23 coordinates persistence, API,
> transaction, non-blocking mail and priority tests without creating modules or dependencies. No code
> outside this scope is modified. The owner explicitly approved it.
>
> **Finding corrected and verified from the third public walkthrough, 2026-10-13.** The installation,
> the seeder and the three entry methods worked, but the lookups prior to creating a category and a
> directory account wrote `record not found` at error level. Section 22 proposes distinguishing the
> expected absences in each repository, keeping all real SQL errors and also verifying the log level.
> The owner chose and approved 1A–4A; the backend suite and `go vet` also passed against real
> PostgreSQL.
>
> **Finding corrected and verified on 2026-10-13: health with pending schema.** Section 21
> distinguishes an unreachable database from an available database whose manual migration is missing,
> with `503`, a stable key and a single warning. The owner closed and approved the coordinated
> decisions on availability, production and tests. The backend suite, `go vet` and an empty PostgreSQL
> verified the implementation.
>
>
> **Amendment implemented and verified on 2026-10-13: local validation of the language.** Section 20
> keeps the language as the single configuration of the deployment, adds the detector inside `ai`,
> updates third-party notices and preserves the boundaries with `tickets` and `mail`. The
> implementation stayed within the proposal explicitly approved on 2026-10-13.
>
> **Amended on 2026-10-12: publication closed.** Section 19 no longer describes as pending the
> visibility, the community configuration and the integration of the current tree in `main`. The
> public repository, its AGPL history and its effective rules were verified; production, the domain
> and the tag `v1.0.0` remain postponed.
>
> **Amendment implemented and verified on 2026-10-10: preparation of the public repository.** Section
> 19 coordinates the full audit, the single rewrite of the history and the source metadata. GitHub
> visibility stays out until a separate final approval.
>
> **Block 7 implemented and verified, 2026-10-08.** Section 18 closes the two technical debts declared
> in section 12: automating the boundaries of the frontend modules and deciding that the browser will
> not have an application logging system. It includes an audit of all `as-built` documents, because
> the review found current statements that contradict the current code. The owner chose 1A–15A and
> explicitly approved the document. The checker, its tests, the integration with the normal commands
> and the documentation audit were completed.
>
> **Block 7 implementation finding, 2026-10-08.** The first run of the checker found that the `core`
> interceptor test imported the root route table only to create the `Router`, although none of its
> cases uses those routes. It is a `core → app` dependency contrary to the approved rule. It is fixed
> by providing an empty router in that test; it changes no product code or verified behavior.
>
> **Amendment implemented and verified on 2026-10-08 (block 6).** The “format review” of section 13
> was the only technical loose end without a definition. Section 17 coordinates it with interface and
> tests according to 1A–14C. What exists continues as-built; the owner explicitly approved the
> proposal on 2026-10-08.
>
> **Finding corrected and verified on 2026-10-08.** Block 5 found that local activation is not
> transactional when switching between models with tight memory. Section 16 defines preflight, second
> measurement and restoration of the healthy model according to 15A–22A. What exists continues
> as-built; the correction was explicitly approved on 2026-10-08.
>
> **Finding and amendment proposed on 2026-10-07.** While reviewing block 4 it was confirmed that the
> AI clients still follow redirects automatically, the local administrator does not supervise a
> `llama-server` that dies after starting and its disk and RAM calculations do not properly consider a
> partial download or the container limit. What exists continues as-built. Section 15 proposes
> correcting those differences; the owner chose 1A, 2A, 3A and 4A. The review closed decisions 5A, 6A,
> 7A and 8A and left no open matters. The owner explicitly approved the amendment, which was
> implemented and verified on 2026-10-07.
>
> **Implemented and verified on 2026-10-07.** The generation clients reject 3xx, limit the body and
> were checked against the four fake contracts. The local administrator supervises the child,
> publishes `active` and `healthy`, resumes according to pending bytes, removes corrupt partials and
> takes the lowest memory limit between machine and cgroup. The whole backend, `ai-manager`, 218
> frontend tests and the full Playwright passed (180 approved and 38 skipped).
>
> **Amendment proposed on 2026-10-06.** What exists continues as-built. Section 14 separates the local
> engine per environment, adds an administrator without the Docker socket, external adapters,
> encrypted secrets and a global language. The review was closed, the proposal was approved and the
> implementation was verified.
>
> **Amended on 2026-10-06**, explicitly approved by the owner, implemented and verified. The
> disposable instance keeps the read-only code bind, removes the child volume of `/app/tmp` and starts
> the test backend with `go run .`. The canonical run without optional services ended with 177
> approved cases and 39 skipped, and the cleanup left no resources.
>
> **Amended on 2026-10-06**, according to the approved amendment of
> `docs/interfaz-y-experiencia.md`: the frontend adds the shared toast of `/settings` and `/setup`. In
> that closure the full frontend and interface suites passed; the historical figures stay in the
> record of the corresponding block, not as a current inventory.
>
> **Amended on 2026-10-04**, according to the approved `docs/prueba-local.md`: in development,
> `backend` depends on `migrate` with `service_completed_successfully`; that way it never queries
> `installation_settings` or `ai_insights` before they exist. `migrate` no longer has a profile and
> the seeder keeps its own. The `SETUP_MAIL_*` suggestions only fill an empty `/setup`.
>
> **Amended on 2026-10-03**, according to the approved `docs/prueba-local.md`: Local startup approved
> in `prueba-local.md`: `dev.yml` adds a single-use `migrate` and removes the e2e runner. `tests.yml`
> owns a disposable instance without published ports or shared data; Angular accepts the `frontend`
> host. Keycloak announces localhost and connects through an independent internal address.
>
> **Amended on 2026-10-02**: review of the README walkthrough. The pending item of the wizard buttons
> is removed from §13: the public entry and mail tests exist, protected by the seal; AI availability
> is queried automatically. The full walkthrough of the empty wizard is still not automated in
> Playwright. A finding is also reported from reading the current deployment: `prod-build.sh` builds
> from the artifacts folder, but the Dockerfiles require the sources of the repository. No code is
> modified and production is not resumed.
>
> **Amended on 2026-10-01 (fifth time)**: **the example data is seeded with a single-use container**,
> `docker compose -f dev.yml run --rm seed-en|seed-es`, **the same on Linux, macOS and Windows**, and
> with **its own profile** so that **`up -d` does not start it**: the examples are **optional** and
> are requested on purpose. The container runs **the usual script, `scripts/dev-seed-en.sh` or
> `scripts/dev-seed-es.sh`** (a single truth), which also remains valid on Linux and macOS outside the
> container. Section 10.3 changes its block of examples for that command (docs/ambientes.md, section
> 3.3).
>
> **Amended on 2026-10-01 (fourth time)**: **the development server forwards `/api` to the backend**,
> so **11001 is enough to enter without nginx**. `frontend/proxy.conf.json` is added —an entry for
> `/api` with destination **`http://backend:11002`**, the backend service **inside the network of the
> environment**, because the one making the request is the frontend container, not the browser— and it
> is declared in `serve.options.proxyConfig` of `frontend/angular.json`. **It is development only**: it
> lives in `serve`, the `build` does not see it, and **in production nginx still forwards**. The
> frontend entry at `http://127.0.0.1:11001` stops depending on nginx; in the requirements table,
> nginx becomes **not necessary to enter** (yes to test the domain and the certificate). Checked on
> 2026-10-01 through 11001: `GET /api/auth/methods` and `GET /api/health` → **200 JSON**, `POST
> /api/auth/login` with `admin`/`admin` → **200** with token, and the entry and exit walkthrough with
> Playwright, **14 cases in green** against `http://127.0.0.1:11001`. The frontend rule does not
> change: it still calls **relative** routes.
>
> **Amended on 2026-10-01 (third time)**: **section 10.3 («Commands») debuts the schema step**. Its
> list started with `docker compose -f dev.yml up -d` and whoever followed it fell into the usual
> error: the application **does not create the tables on startup** and the migrations **do not apply
> themselves**, so on a new database the backend responds `relation "installation_settings" does not
> exist`. The **two Docker commands that work the same in PowerShell, CMD and bash** are added —copy
> `v1.0.0.sql` into the container and give it to `psql` with `-f`—, with the examples (`_dev`)
> afterwards and the notice that `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh` **is for
> Linux and macOS** (the full steps, in `README.md`).
>
> **Amended on 2026-10-01 (second time)**: **the test directory and Keycloak leave `dev.yml`** and move
> to **`active-directory.yml`** and **`keycloak.yml`**, each with **its own command**, and **the `auth`
> profile disappears**. Both enter the **`catalina-support-dev`** network, which **`dev.yml` owns**
> (fixed name, no `external`), so **the environment is started first** and then the service; the
> backend reaches them by their service name (`ldap`, `keycloak`). The structure, the development
> services table, the commands and section 13 are brought up to date.
>
> **Amended on 2026-10-01**: **the AI engine row of section 13 is corrected** so that it says the same
> as the README and `docs/modules/ai.md`, with what was measured: **1500 MiB cap** (`mem_limit: 1500m`
> in `ai.yml`), **~1.44 GiB running —98% of the cap—** and **no CPU cap** (the compose does not limit
> CPU; the "2 CPU" that the row said came from nowhere in the compose). The "1.09 GB with an input of
> 2,034 pieces" is removed, which could not be verified.
>
> **Amended on 2026-09-30**: **section 13 is brought up to date**. What has changed this month and was
> not accounted for: **the Settings screen is complete** —with **the AI engine** and its connection
> test, the time zone and the public address, and **the outgoing mail**—; there is **a first-start
> view** (`/setup`, then four steps and today five after the amendment of 2026-10-06, and the seal in
> `installation_settings.installed_at`, with its API at **409** afterwards); **the AI engine lives
> apart and is configured from Settings**; and **the counts** remain in Go **14 packages**; the
> frontend counts of that date were surpassed by the amendment of 2026-10-06. **The production
> deployment is postponed** until the product is finished (owner's decision, 2026-09-30), so it is not
> part of this update. What is missing is said below.
>
> Approved by the project owner on 2026-09-22 and **stood up the same day**: the container skeleton
> started with the two projects. What exists today is in section 13, and each module has its own
> document.

## 1. Scope of this document

It fixes the **stack**, the **shape of the repository**, the **containers** and the **modularity
rules** of Catalina-Support. It does not define product or flows: that lives in the documents of each
area (purpose and scope, users and permissions, tickets, flows).

This document is **as-built**: it describes what the code does today, and if the code changes, the
document changes in the same job (Rule 0 of `AGENTS.md`).

## 2. Stack

Versions checked on **2026-09-22** against their official source. Major versions are fixed; minor
ones are updated when appropriate without going through a new proposal.

| Layer | Technology | Version | Source of the check |
| --- | --- | --- | --- |
| Frontend | Angular | **22.1.7** (CLI 22.1.8) | `npm view @angular/core dist-tags` → `latest` |
| Frontend | Node.js | **24 LTS** (24.21.0) | `nodejs.org/dist/index.json` (Angular 22 requires `^22.22.3 \|\| ^24.15.0 \|\| >=26.0.0`) |
| Backend | Go | **1.27.1** | `go.dev/dl/?mode=json` |
| Backend | `net/http` (standard library) | Go 1.27.1 | owner's decision: no HTTP framework |
| Backend | GORM | **v1.31.2** + driver `gorm.io/driver/postgres` **v1.6.3** | `proxy.golang.org` |
| Backend | air (reload in development) | **v1.67.4** (`github.com/air-verse/air`) | `proxy.golang.org` |
| Backend | go-logs | **v1.0.0** (`github.com/juanky-estevez/go-logs`) | `proxy.golang.org` |
| Mail | `net/smtp` (standard library) | Go 1.27.1 | same as Calibyou: no new dependencies |
| Data | PostgreSQL | **18** (latest minor 18.6) | `postgresql.org/versions.json` → `current: true` |
| Runtime | Docker + Docker Compose | the machine's (29.5.2 checked) | `docker --version` |

Notes on the versions:

- **Angular 22** is the latest stable and the only one installed (`latest`); an earlier version is
  not fixed. The published `next` is 22.2.0-rc.0 and is **not** used.
- **Node 24 LTS** ("Krypton") is the current LTS line and the one Angular 22 recommends. Node 26
  works but is not LTS yet: it is not used.
- **PostgreSQL 18** is the major version marked as `current` by the PostgreSQL project, and the
  official image `postgres:18-alpine` is used.
- The `go-logs` package declares `go 1.23` as minimum, so it is compatible with Go 1.27.
- There is a sibling package `github.com/juanky-estevez/logs` (`@jcestevez/logs`) that is a
  **NestJS middleware** (it writes to the file system through Node). **It is not usable for the
  Angular frontend**, which runs in the browser: there is no logging package for the browser
  decided, and it remains pending (section 12).

## 3. General architecture

```text
navegador
  -> nginx (vhost; enruta /api/ al backend y el resto al frontend)
      -> contenedor frontend: Angular 22 (SPA)
          -> /api/<módulo>/...
      -> contenedor backend: Go (net/http) + GORM
          -> contenedor database: PostgreSQL 18
```

A single backend, a single database and a frontend, each in its container. There are no
microservices, no queues and no cache in this version: if they were needed, they are proposed in
their corresponding document.

## 4. Modularity (central rule)

The project is organized **by module**, and the module is the unit of everything: folder, endpoints,
tables and screens. The rule is explicit and verifiable:

> **A frontend module consumes only the endpoints of its homonymous backend module.**

Five concrete rules come out of that:

1. **One name, three places.** A module is called the same in the three places: backend folder,
   frontend folder and API prefix. Example: `users` →
   `backend/modules/users`, `frontend/src/app/modules/users` and `/api/users/...`.
2. **The frontend does not cross modules.** `frontend/src/app/modules/users` calls only
   `/api/users/**`. If a screen needs data from another module, the module that owns that data is
   called from the backend, not requested from another screen to a foreign endpoint.
3. **The backend does not cross tables.** A module writes only in its own tables. Whatever it needs
   from another module is requested through an exposed service function, never by reading its table.
4. **`shared` does not depend on `modules`.** The cross-cutting part (configuration, database
   connection, middlewares, utilities) lives in `shared` and knows no specific module. A module can
   use `shared`.
5. **`core` is the frontend exception.** The shell —`frontend/src/app/core`— can call the endpoints
   of the **session** and those of the **installation configuration** (the numbering prefix and the
   mail templates), and nothing else. It is a single rule with its reason: the session and the
   installation configuration belong to no module, and splitting them into three separate exceptions
   is what makes a rule be forgotten.
   **The permission does not force its use**: the screen that edits the mail templates belongs to
   the frontend `mail` module, not to `core` (decided on 2026-09-23, `docs/modules/mail.md`, decision
   15), because its data is its own and that way the exception does not become the norm.

Visual correspondence:

```text
frontend/src/app/modules/users  --HTTP-->  /api/users/**  -->  backend/modules/users  -->  tablas users*
```

Frontend calls are **always relative** (`/api/users`), never to an absolute host: the same code works
in development and in production because it is nginx that routes, and that way there is no CORS to
configure.

What is gained: finding anything about a module means looking at its folder, and deleting or changing
a module does not force a sweep of the rest of the repository. What is paid: the rule must be
respected by hand, because there is nothing that imposes it automatically (see section 12).

The closed list of modules is **not** decided here: each area document defines its own. `users` is
used in this document only as an example of the convention.

## 5. Backend (Go + `net/http` + GORM)

### Shape of the repository

```text
backend
├── go.mod
├── main.go                  # arranque: configuración, GORM, rutas, servidor
├── .air.toml                # recarga en caliente (sólo desarrollo)
├── migrations
│   └── v1.0.0.sql
├── modules
│   └── <módulo>
│       ├── controllers      # HTTP: leer petición, validar forma, responder
│       ├── services         # reglas del módulo
│       ├── repositories     # acceso a datos
│       └── dtos             # lo que entra y sale por la API
└── shared
    ├── auth                 # el token de sesión (firma y validación) y quién es quien llama
    ├── authz                # permisos: si el papel llega para lo que se pide
    ├── config               # variables de entorno
    ├── database             # conexión y configuración de GORM
    ├── httpx                # la forma única de las respuestas y de los errores
    ├── middleware           # autenticación, permisos, recuperación, logs
    └── version              # la versión del sistema, que sale en la marca pública
```

The style already used in the sibling project Calibyou is kept (`modules/<m>/controllers`,
`services`, `repositories`, `dtos`), because it is the one the owner already knows and maintains.

Today the repository has `main.go`, `.air.toml`, `modules/` with its **six modules** (`ai`, `auth`,
`mail`, `settings`, `tickets` and `users`), `migrations/` with `v1.0.0.sql` and `v1.0.0_dev_en.sql`
or `v1.0.0_dev_es.sql`, and
`shared/{auth,authz,config,database,httpx,middleware,version}`. `shared/httpx` was added when the
skeleton was created so that the shape of HTTP responses has a single owner from the start: today it
writes `{"error": "key"}`, with a **key** that the frontend translates into the language of the
reader (`docs/interfaz-y-experiencia.md`).

### `net/http` without a framework

The standard library router is used, which since Go 1.22 accepts **method and pattern with
parameters**:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /api/users", h.Listar)
mux.HandleFunc("GET /api/users/{id}", h.Obtener)   // r.PathValue("id")
```

Consequences that must be assumed:

- Middlewares are chained by hand (a function that wraps `http.Handler`); there is no
  `Use()`. It is solved with an explicit chain at startup.
- There is no binding or automatic JSON validation: each controller decodes, validates and responds;
  errors are written with a common helper of `shared` (single error response).
- In exchange: zero framework dependencies and updates without surprises.

### GORM

- A single connection created at startup and shared (`*gorm.DB` in `shared/database`).
- The models of the module live in its `repositories`/`dtos`; models are not shared between
  modules.
- **Explicit SQL with migrations** is preferred for the schema (section 7) and GORM for queries and
  writing; `AutoMigrate` is **not** used, so that the schema has a single source of truth.
- Soft delete (`gorm.DeletedAt`) when the module needs it, decided in its document.

### Development with air

- `air` (v1.67.4) reloads the backend on save. It is **installed in the development image** and is
  the container command: Go is not installed on the machine.
- Its configuration lives in `backend/.air.toml` (watches `**/*.go`, builds to a temporary binary
  and ignores `tmp/`).
- `air` is **development only**; the production container runs the compiled binary.

## 6. Frontend (Angular 22)

### Shape of the repository

```text
frontend
├── angular.json
├── package.json
├── .postcssrc.json          # Tailwind v4 (el plugin de PostCSS)
├── src
│   ├── main.ts
│   ├── index.html
│   ├── styles.css           # Tailwind y las variables de los temas
│   └── app
│       ├── app.config.ts    # providers globales (router, http, interceptor)
│       ├── app.routes.ts    # rutas del armazón; cada módulo aporta las suyas
│       ├── core             # armazón: sesión, i18n, interceptores, guardas, pantallas de la sesión
│       │   ├── components   # piezas del armazón (menú lateral, logo, controles, aviso de servidor caído)
│       │   ├── guards       # sin sesión no se entra
│       │   ├── i18n         # los dos diccionarios de textos y el servicio que los sirve
│       │   ├── interceptors # la cabecera de la sesión, el 401 y el «servidor caído»
│       │   ├── navigation.ts # las entradas del menú, por papel
│       │   ├── pages        # las seis pantallas de la sesión y la de Configuración
│       │   └── services     # sesión, salud del backend, disponibilidad
│       ├── shared           # UI y utilidades reutilizables, sin lógica de negocio
│       │   └── components   # los componentes propios (botón, campo, aviso, tarjeta…)
│       └── modules
│           └── <módulo>     # pages, components, services, models, <módulo>.routes.ts
```

### Conventions (inherited from modern Angular, not from old Angular)

- **Standalone components** (no `NgModule`), which is the CLI default since Angular 19. Angular
  modules are not created; "module" here means **feature folder**, not `NgModule`.
- **Signals** (`signal`, `computed`) for the screen state and `provideHttpClient` with `HttpClient`
  for the calls.
- **New control flow** (`@if`, `@for`, `@switch`) instead of `*ngIf`/`*ngFor`.
- **Lazy routes per module**: each module exposes its `<module>.routes.ts` and the root route loads
  it with `loadChildren`. A screen that is not visited is not downloaded.
- **One service per module** (`modules/<m>/services/<m>.service.ts`) that centralizes the calls to
  `/api/<m>/**`. No component calls `HttpClient` directly, and no service points to a foreign module
  (section 4).
- File and folder names in **English** (that is what the CLI generates: `user-list.ts`); interface
  texts and documentation, in Spanish.
- **Texts are not written inside the components**: they live in `core/i18n/es.ts` and
  `core/i18n/en.ts`, with the same shared type, so that **a key missing in one language does not
  compile**. The backend error texts go in the same dictionary (`docs/modules/auth.md`, decision 26).
- **Colors come from the theme variables** of `src/styles.css`, exposed to Tailwind with
  `@theme inline`: a component never carries a hand-written color, and changing the theme touches no
  component (`docs/interfaz-y-experiencia.md`, section 6.2). The palettes are written **once** with
  `light-dark()`, and the theme is chosen with an attribute on `html` (`data-theme`), or none to
  follow the system. **The theme is applied by `main.ts`, not by a script embedded in
  `index.html`**: the production CSP is `script-src 'self'`.
- **Routes also in English** (`/login`, `/set-password`), in lowercase and with hyphens: decided by
  the owner on 2026-09-23 (`docs/modules/auth.md`, decision 21), because an address is copied,
  pasted and shared, and in code it is not convenient to mix languages.

What once remained to be decided here **is already decided in other documents**: the styles and the
own components, in `docs/interfaz-y-experiencia.md` (Tailwind v4, closed inventory of twenty pieces);
session and permission management on the client, in `docs/modules/auth.md`, section 9 (the session
lives in `core`: service, interceptor, guard and six screens); and **SSR is not used**, which is a
decision and not an absence.

Today the repository has **the whole shell**: the session service, the interceptor that sets the
header and handles the 401, the route guard, the two text dictionaries, the **own components**
(button, field, notice, card, switch, selector and logo), the **six session screens** (`/login`,
`/forgot-password`, `/set-password`, `/change-password`, `/forbidden` and the server-down notice,
which has no route) and **the side menu with its three zones**.

The routes are in **two groups**: the session screens go full screen and everything else goes inside
the shell, with its menu. **The menu shows only what exists**, and what corresponds to each role: the
ticket lists, Users and, for an Administrator, their inbox and Settings.

**The full format review of the frontend is closed**. **The product screens —users, tickets and the
mail editor— are done**, and **the provisional home disappeared**: the root is an assignment —the
Administrator to the users, the rest to their inbox—.

## 7. Database (PostgreSQL 18)

- One database and one schema; isolation per client, if it ever exists, is decided in the users and
  permissions document.
- **A single path for the schema**: the files of `backend/migrations/` (today only `v1.0.0.sql`),
  transactional (`BEGIN` … `COMMIT`) and idempotent (`CREATE ... IF NOT EXISTS`,
  `INSERT ... WHERE NOT EXISTS`). They are applied with `psql -v ON_ERROR_STOP=1` inside the database
  container. No loose `ALTER` statements and no manual steps when deploying.
- While the project has no real data, that same file is updated instead of adding new migrations.
  **From 1.0.0 onwards it moves to one file per version** (`v1.1.0.sql`, `v1.2.0.sql`…) **plus one
  example data file per version with the `_dev` suffix** (`v1.0.0_dev_en.sql` or
  `v1.0.0_dev_es.sql`), which **is never applied in production**, applied in order: the full policy is
  in `docs/ambientes.md`, section 5.
- Table names in plural and in English, consistent with the module name (`users`, `tickets`); primary
  keys `id`; timestamps `created_at`, `updated_at`.
- The data lives in a **Docker volume**, never inside the container.

## 8. Logs (go-logs)

The backend uses `github.com/juanky-estevez/go-logs` v1.0.0 instead of `log`/`slog` and instead of
`fmt.Println`. Its API has exactly four functions:

```go
logs.LogInfo("servidor escuchando en :11002")
logs.LogSuccess("migración aplicada")
logs.LogWarning("intento de acceso sin token")
logs.LogError("no se pudo conectar a la base de datos")
```

Each call prints to the terminal in color **and** adds a line to the file of the day:

```text
[info] dev 2026-09-22 10:30:00 servidor escuchando en :11002
```

Environment variables it consumes:

| Variable | What for | Value |
| --- | --- | --- |
| `PROJECT_NAME` | first part of the file name | `catalina-support` |
| `ENVIRONMENT` | second part of the name and the text of each line | `dev` / `prod` |
| `LOGS_FOLDER` | folder of the files | `/logs` (volume `./_logs` mounted in the container) |
| `TIMEZONE` | offset used when dating | `UTC-5` |

The file is called `PROJECT_NAME_ENVIRONMENT_AAAAMMDD.log` (one file per day) and is written to
`_logs/`, which is a folder of the repository **outside git** and mounted in the container. Being a
volume, the logs survive restarts and are read from the machine without entering the container.

**Known limitations of the library** (they are accepted, and that is why they are written down):

- There are no configurable levels or structured format (it is not JSON) or fields: the message is a
  string and nothing more. Whatever is to be queried later must be written into the message.
- There is no `Debug` or traces: four levels, four functions.
- There is no rotation or deletion: one file per day, and they grow. Cleaning `_logs/` is a manual
  task.
- If it cannot write the file, **it does not fail**: it prints a warning in red and writes only to
  the terminal. A misconfigured `LOGS_FOLDER` loses the log without stopping the application, so the
  existence of the folder must be checked at startup.
- There is no request identifier or context: correlating a full trace requires including it by hand in
  each message.

Usage rules: no `fmt.Println` or `log` in the application code; startup records the loaded
configuration and the established connection; the errors the user sees are also recorded; **never**
are passwords, tokens or personal data recorded.

## 9. Mail

All mail from the application goes out through the **`mail` module**, with the **standard library**
(`net/smtp`), like Calibyou. `auth` and `tickets` **do not write or send mail**: they ask `mail` to
send, with the data inside, and `mail` puts the text.

| Who asks | What it asks for |
| --- | --- |
| `tickets` module | The **eight ticket notices** (`docs/propósito-y-alcance.md`) |
| `auth` module | The **three account mails**: sign-up, recovery and change notice (`docs/modules/auth.md`) |

**The texts are not in the code**: they are **editable templates** that live in the database, and an
administrator changes them from the Settings screen, with preview and a test to their own mail. The
detail is in `docs/modules/mail.md`.

**The outgoing mail lives in the database, not in environment variables**: the fields are
`installation_settings.smtp_*` (`smtp_host`, `smtp_port`, `smtp_security_mode`, `smtp_insecure_ack`,
`smtp_user`, `smtp_password`, `smtp_from_name`, `smtp_from_email`), they are requested in the
first-start view and changed from Settings. The `SMTP_*` environment variables **were removed**.

They are called `smtp_*` and not `SYSTEM_SMTP_*` as in Calibyou: there the prefix distinguishes the
system mail from that of each laboratory, and here there is only one.

**In development they point to the test mailbox** of `dev.yml`: the example data file
(`backend/migrations/v1.0.0_dev_en.sql` or `backend/migrations/v1.0.0_dev_es.sql`) leaves
`smtp_host='mail'` and `smtp_port='1025'` set, so everything that goes out is read at
`http://127.0.0.1:11004` and **nothing goes out to the internet** (section 10). Without it, the
sign-up of a local account could not be completed in development. If the configuration is missing,
the sending fails with a clear error and **it is not attempted another way**: there is no "backup"
mail.

**A mail failure does not bring down the action that caused it.** If the ticket is created and the
notice does not go out, the ticket exists: the action is already done and what fails is a consequence.
The failure is recorded with `go-logs` and the request responds fine. The opposite —undoing a ticket
because the mail did not go out— leaves the user looking at an error for something that did happen.

**The sending does not block the request**: it goes in the background with its time limit, so that
creating a ticket does not depend on how long the mail server takes. The price is that the failure is
known from the log and not from the screen, and it is a reasonable price for a notice.

**The mails carry just enough**: ticket number and subject, and the link when appropriate. Never
passwords, or tokens, or unnecessary personal data.

**The links are built with the public address of the installation**, which is configured on the
Settings screen: a host written by hand in a template cannot serve in development and in production
at the same time. `PUBLIC_APP_URL` remains **as a fallback** for an installation that already had it
set (`docs/modules/settings.md`, decision 14).

## 10. Containers and execution

Everything runs in containers, with the same shape as Calibyou: **one container for the frontend,
another for the backend and another for the database**, and nginx on the machine as the entry door.
Go and Node do not need to be installed on the machine: only Docker and nginx.

### Structure

```text
.
├── dev.yml                          # compose de desarrollo (los servicios; la red del entorno es suya
│                                    # y las de interfaz van detrás de un perfil)
├── prod.yml                         # compose de producción (imágenes construidas)
├── ai.yml                           # el motor local de IA, requerido sólo para la modalidad local
├── active-directory.yml             # el directorio de pruebas (OpenLDAP; opcional)
├── keycloak.yml                     # Keycloak de pruebas (opcional)
├── config
│   ├── dockerfiles                  # una imagen por servicio y por entorno
│   ├── seed                         # los archivos de los adjuntos del seeder de desarrollo
│   ├── env
│   │   ├── dev.env
│   │   ├── prod.env.example         # plantilla; prod.env NO se versiona
│   │   └── (prod.env)               # NO se versiona
│   ├── keycloak
│   │   └── realm-catalina-support.json  # el reino de pruebas, con su cliente y tres personas
│   ├── ldap
│   │   └── 01-personas.ldif         # las tres personas del directorio de pruebas
│   └── nginx
│       ├── catalina-support-dev.conf    # vhost de desarrollo
│       ├── catalina-support-prod.conf   # vhost de producción
│       └── frontend.prod.conf           # nginx del contenedor de frontend
├── backend
├── frontend
├── scripts                          # lo que se ejecuta en el servidor: las copias de la base
├── _logs                            # logs de go-logs (la carpeta se versiona, los .log no)
├── _files                           # adjuntos, por año y número de ticket (se versiona la carpeta,
│                                    # no los archivos)
└── docs
```

### Development services (`dev.yml`)

| Service | Base image | Command | Port |
| --- | --- | --- | --- |
| `frontend` | `node:24-alpine` | `npm start -- --host 0.0.0.0 --port 11001` | `11001` |
| `backend` | `golang:1.27-alpine` + air | `air -c .air.toml` | `11002` |
| `database` | `postgres:18-alpine` | `postgres` (`PGPORT=11003`) | `11003` |
| `mail` | `axllent/mailpit:latest` | `mailpit` | `11004` (its web; the SMTP is the internal `1025`) |
| `ldap` (**in `active-directory.yml`**, not in `dev.yml`) | `osixia/openldap:1.5.0` | The directory, with the people of `config/ldap/01-personas.ldif` seeded at startup | `11005` (LDAP in clear: **it is development**). Optional: **it is not started with `dev.yml`**, it is started with `docker compose -f active-directory.yml up -d` |
| `keycloak` (**in `keycloak.yml`**, not in `dev.yml`) | `quay.io/keycloak/keycloak:26.0` | `start-dev --import-realm`, with the realm of `config/keycloak/realm-catalina-support.json` | `11006` (HTTP: **it is development**). Optional: **it is not started with `dev.yml`**, it is started with `docker compose -f keycloak.yml up -d` |
| `ai` (**in `ai.yml`**, not in `dev.yml`) | llama.cpp with HTTP server | Serves the model **Qwen2.5-1.5B-Instruct Q4_K_M**, which is downloaded once to the volume the first time | **None**: the engine is not published on the machine, it is only spoken to through the `catalina-support-ai` network. It is shared by **development and production** (`docs/modules/ai.md`, decision 1) |
| `migrate` | `postgres:18-alpine` | Applies only `v1.0.0.sql`, with `ON_ERROR_STOP=1`; `backend` waits for its correct completion | None; it runs automatically with `up` and also accepts `run --rm migrate` |
| `seed-en`, `seed-es` | `postgres:18-alpine` | They run the explicit access and the shared runner | None; optional examples per language |

**The AI engine is not in the table above on purpose**: it lives in `ai.yml`, **shared by the two
environments** —two models do not fit on this machine— and the two backends enter its network
(`catalina-support-ai`). Its why, its arguments and its memory are in `docs/modules/ai.md`, section 2,
and how it is started in `docs/ambientes.md`, section 3.4.

**The test directory and Keycloak are not in `dev.yml` either**: each lives in its file
(`active-directory.yml` and `keycloak.yml`), they are started with their command **after** the
environment and they enter the **`catalina-support-dev`** network, which **`dev.yml` owns** with a
fixed name (no `external`): that way the backend reaches them by their service name (`ldap`,
`keycloak`) without them being in the compose of the environment. Their why and how they are started,
in `docs/ambientes.md`, section 9.3.

Details that matter:

- **Code mounted as a volume** (`./frontend:/app`, `./backend:/app`): it is edited on the machine and
  the container reloads (Angular with its development server, Go with `air`). The image is not
  rebuilt for each change.
- **Caches in named volumes** (the frontend `node_modules`, Go modules and build cache): they are
  neither lost when recreating the container nor mixed with files of the machine.
- **`_logs` mounted** in the backend (`./_logs:/logs`) so that the `go-logs` files are visible on the
  machine.
- **`_files` mounted** in the backend (`./_files:/files`, variable `FILES_PATH`), with **one folder
  per ticket** (`_files/2026/CS-2026-0042/`) and the brand one apart (`_files/brand/`), for the
  attachments of the tickets: they are files, not rows, so they survive restarts and are copied from
  the machine like any folder.
- **The backend waits for the database**: `depends_on` with `condition: service_healthy` and a
  `healthcheck` of `pg_isready`, so that it does not start against a database that does not yet accept
  connections.
- **Own ports**: `11001`-`11006` are used, different from those of Calibyou (`10001`-`10004`) so that
  the two projects can be up at the same time. They are published **without a listen address**, as
  they are in `dev.yml` and in `prod.yml`.
- **`ldap` is the test directory** (`active-directory.yml`): OpenLDAP with **three people** in an LDIF
  of the repository (`config/ldap/01-personas.ldif`), one for each case of the AD path. **It does not
  save anything in a volume**: it is seeded from the repository at startup, so it is always as the
  repository says, and **it is not stopped and started** —its startup scripts are not idempotent and
  it dies—: it is started again by recreating it (`docs/modules/auth.md`, section 11). And **it is not
  Active Directory**: it is spoken to the same way, so the code is really tested, but what is specific
  to AD is not covered.
- **`keycloak` is the test Keycloak** (`keycloak.yml`): the realm lives **in the repository**
  (`config/keycloak/realm-catalina-support.json`), with its confidential client and three people, and
  it is imported at startup: recreating the container leaves it as written. **The browser reaches it
  through the same domain as the application and through one path** (`/sso/`, a `location` of the
  development vhost): that way the OIDC issuer is a single one for the browser and for the backend,
  and there are no two addresses that can disagree (`docs/modules/auth.md`, section 11).
- **`mail` is the test mailbox** (Mailpit), decided by the owner on 2026-09-23: **without it the
  sign-up of a local account cannot be completed in development**, because the password is set from
  the link in the mail. It is **a single mailbox for everything**: it accepts any recipient and from
  any domain, it has no accounts or authentication to configure, and **it forwards nothing to the
  internet**, so in development no mail leaves the machine. The mails are read at
  `http://127.0.0.1:11004`, with their HTML and their text version, and they are filtered by recipient
  (`to:`), by subject (`subject:`) or by text, which is what serves as "inbox per account". It keeps
  the last 200 messages. In development the example data file leaves it set
  (`smtp_host='mail'`, `smtp_port='1025'`); in production the real mail server is set from Settings.
- **Variables in `config/env/dev.env`**, shared by the services, with the names of `POSTGRES_USER`,
  `POSTGRES_PASSWORD` and `POSTGRES_DB` so that the official PostgreSQL image takes them on its own.
  `config/env/prod.env` **is not versioned** (only a `prod.env.example`).

### nginx

nginx is the entry door in the two environments, with **one vhost per environment**, and it is
installed by copying the file of the repository (there are no symbolic links, as in Calibyou):

```bash
sudo cp config/nginx/catalina-support-dev.conf  /etc/nginx/conf.d/
sudo cp config/nginx/catalina-support-prod.conf /etc/nginx/conf.d/
sudo nginx -t && sudo systemctl reload nginx
```

| Environment | Domain | Vhost | Containers behind |
| --- | --- | --- | --- |
| Development | `dev.catalina-support.example.com` | `catalina-support-dev.conf` | frontend `11001`, backend `11002` |
| Production | `support.example.com` | `catalina-support-prod.conf` | frontend `21001`, backend `21002` |

- The frontend calls `/api/**` **relatively** and never to a host: the same code serves in development
  and in production without changes and without CORS.
- **The `/api` prefix is not rewritten**: the backend routes already carry it
  (`/api/<module>/...`). Calibyou does rewrite it because its backend exposes them without a prefix;
  here it does not, and that difference is commented in the two vhosts so that nobody "fixes" it.
- **TLS** with Let's Encrypt and the `webroot` plugin (`/var/www/certbot`), one certificate per domain,
  renewed by the certbot task that the package already brings (`certbot.timer` and
  `/etc/cron.d/certbot`). The two vhosts serve `/.well-known/acme-challenge/` on port 80 **before**
  the redirect to HTTPS, so the renewal does not need to touch anything.
- **`client_max_body_size 30m;`** in the two vhosts. Without this line nginx uses **1 MB** by default
  and any attachment larger than one mega is rejected with a 413 **before reaching the backend**: the
  25 MB limit of attachments (`docs/modules/tickets.md`) could not be met. The 30 leave margin.
- **`limit_req` zone for login**, as in Calibyou: it is the only public endpoint that receives
  credentials, and without a limit it can be probed with passwords at will. Only that route, because
  the rest of the traffic does not need it.
- **The development vhost does NOT include** `snippets/common-security.conf`, for the same reason as
  the one of Calibyou: that snippet returns 444 to routes with a segment that starts with a dot and
  applies `limit_req`/`limit_conn`, and that leaves the Angular development server blank (which
  internally is Vite). The production one **does** include it.
- The development CSP is permissive (`script-src 'self' 'unsafe-inline'`, `connect-src` with `ws:`)
  because the development server injects a script and opens a websocket for hot reload; the production
  one is strict (`script-src 'self'`).
- The Angular development server **rejects with 403** any `Host` it does not know: that is why
  `dev.catalina-support.example.com` is declared in `serve.options.allowedHosts` of `angular.json`.
  Without that line, the page responds 403 even if nginx is fine.
- **The development server forwards `/api` to the backend** (`frontend/proxy.conf.json`, declared in
  `serve.options.proxyConfig` of `frontend/angular.json`), with destination **`http://backend:11002`**
  —the service name of the backend **inside the network of the environment**, because the one making
  the request is **the frontend container**, not the browser— and **without rewriting the route** (the
  backend routes already carry `/api`, as in nginx). It is **development only**: it is an option of
  `serve`, the `build` does not see it and **in production nginx forwards**. It is what allows entering
  `http://127.0.0.1:11001` with only Docker and without nginx in front.
- The `limit_req` zone of login **is already set**, as in Calibyou, because `auth` is done and it is
  the only public endpoint that receives credentials (a few lines above).

### Production (`prod.yml`)

- Images **built** with two-stage Dockerfiles (compile and then run), not with the code mounted:
  - backend: `golang:1.27-alpine` compiles a static binary and the final image is `alpine` with the
    binary (no Go, no `air`, no source code).
  - frontend: `node:24-alpine` compiles with `npm ci && npm run build` and the static result is served
    by nginx inside the `frontend` container (port 21001).
- The production vhost is already installed and with a certificate. While the postponement lasts the
  domain responds **503 with a notice** («this version is still in development») instead of the 502
  that nginx gives when there is nothing behind: a 502 looks like a broken website. The response keeps
  the security headers and the strict CSP because the block does not declare its own `add_header` (if
  it did, nginx would stop inheriting those of the server). `/api/` still points to the backend that
  does not exist, so it returns 502: nobody consumes it during the pause and that way on deployment
  day only one block has to be restored, which is commented in the vhost itself.
- **The production deployment was done on 2026-09-25** with `prod.yml`, its Dockerfiles and
  `frontend.prod.conf`, and remained **verified from the inside**; today it is **postponed** until the
  product is finished (owner's decision, 2026-09-30). The deployment runbook is in
  `docs/ambientes.md`, section 4. What is not touched is the nginx and TLS part of the previous
  section: the certificates renew themselves and the vhost is already in place.

### Commands

All commands are run **on the machine**, against the containers: there is no Go or Node installed.
They are also in `AGENTS.md`.

```bash
docker compose -f dev.yml up -d --build         # esquema y servicios; no carga ejemplos
docker compose -f dev.yml ps                    # ver el estado
docker compose -f dev.yml logs -f backend        # seguir los logs (salida de go-logs)
docker compose -f dev.yml down                   # parar (los volúmenes se conservan)

# Repetición manual opcional del esquema idempotente; el `up` anterior ya lo ejecuta
docker compose -f dev.yml run --rm migrate
# Los datos de ejemplo, que son **opcionales** y **un extra para tener contenido, no un requisito
# para entrar**. **El comando de los tres sistemas** es un contenedor de un solo uso que corre el
# guion `scripts/dev-seed-en.sh` o `scripts/dev-seed-es.sh` (esquema, ejemplos y adjuntos); **no lo arranca `up -d`** (perfil
# propio). En Linux y macOS el mismo guion se puede correr fuera del contenedor
docker compose -f dev.yml run --rm seed-en|seed-es
# ./scripts/dev-seed-en.sh  # o dev-seed-es.sh

docker compose -f dev.yml exec backend go test ./...      # pruebas del backend
docker compose -f dev.yml exec backend go vet ./...       # análisis estático
docker compose -f dev.yml exec backend go mod tidy        # tras añadir dependencias

docker compose -f dev.yml exec frontend npm test -- --watch=false
docker compose -f dev.yml exec frontend npm run build

docker compose -f dev.yml exec database psql -U catalina_support -d catalina_support -p 11003

docker compose -f active-directory.yml up -d    # el directorio de pruebas (opcional; tras el entorno)
docker compose -f keycloak.yml up -d            # Keycloak (opcional; tras el entorno)
docker compose -f ai.yml up -d                  # motor requerido si se eligió IA local
```

In development you enter through **https://dev.catalina-support.example.com** (nginx → containers).
Ports 11001 and 11002 remain published to debug without going through nginx, and **11001 already
forwards `/api` to the backend** with the `serve` proxy (`frontend/proxy.conf.json`), so **you can also
enter `http://127.0.0.1:11001` without nginx**; 11002 is the backend directly.

## 11. Development environment

Requirements on the development machine (checked on 2026-09-22):

| Tool | Status | Mandatory? |
| --- | --- | --- |
| Docker | 29.5.2 installed | **Yes**: everything runs in containers |
| Docker Compose | included in Docker | **Yes** |
| nginx on the machine | 1.24 installed, **with the two vhosts installed and certified** | **No to enter**: 11001 forwards `/api` with its proxy. **Yes**, to test the domain and the certificate |
| Node.js | v24.19.0 installed | No: the frontend compiles in its container (v24.21.0 in the image) |
| Go | **not installed** | No: the backend compiles and runs in its container (1.27.1 in the image) |
| `psql` client | 18.4 installed | No: to query you enter the database container |

That Go and Node are not needed on the machine is deliberate: **Docker is enough to enter** —nginx is
only needed to test the domain and the certificate—, and the versions are fixed in the images (Go
1.27, Node 24), not in what each machine has installed.

## 12. Decided, pending and out of scope

**Decided** (the owner asked for it and this document records it):

- Angular in its latest version for the frontend.
- PostgreSQL in its latest version as the database.
- Go backend with `net/http`, GORM and `air` in development.
- `github.com/juanky-estevez/go-logs` for the backend logs.
- Modular organization, with the frontend module consuming the endpoints of its homonymous backend
  module.
- **Everything in containers**: one for the frontend, another for the backend and another for the
  database, with the structure of Calibyou (`dev.yml` / `prod.yml`, `config/{dockerfiles,env,nginx}`).
- **Node 24 LTS** and **no Go or Node installed on the machine**.
- **Domains**: `dev.catalina-support.example.com` in development and `support.example.com` in
  production, both with Let's Encrypt TLS and served by nginx on the machine (decided by the owner on
  2026-09-22).
- **The production deployment is postponed until version 1.0.0 is closed** (decided by the owner on
  2026-09-22). During development only the development environment is worked against.

**What was here and is already decided** (each in its document, which is where the decision lives):

| What | Where it was decided |
| --- | --- |
| Styles, component library and screens | `docs/interfaz-y-experiencia.md` (Tailwind v4 and twenty own components) |
| Authentication, session, roles and permissions | `docs/usuarios-y-permisos.md` and `docs/modules/auth.md` |
| Data model and ticket states | `docs/modules/tickets.md` |
| Closed list of modules and their endpoints | `docs/modules/tickets.md`: six modules, in the order `mail` → `auth` → `users` → `tickets`, and `settings` |
| Notifications and mail | `docs/modules/mail.md` and, for what is sent at each moment, `docs/flujos.md` |
| Deployment runbook, migrations and backups | `docs/ambientes.md` |
| What is tested and with what | `docs/ambientes.md`, section 9: three layers (`go test`, `npm test` and Playwright) |
**Closed in block 7** (section 18):

- **Frontend logs**: a logging system is not added. The browser only keeps the report of the fatal
  startup failure; the application does not record activity, bodies, credentials or functional errors.
- **Frontend boundaries**: an own checker, without new dependencies, will make tests and compilation
  fail on imports between modules, dependencies in the opposite direction between layers or calls from
  a module to a foreign API prefix.

**Out of scope of this version**: microservices, queues, distributed cache, SSR, multi-tenant and
mobile application.

## 13. Implementation status

What exists today in the repository and what has really been checked. The last pass is from
**2026-10-08**: full backend and frontend suites, production compilation of the frontend and automatic
modular boundaries. The last full visual matrix is also from 2026-10-08. Production operation remains
postponed and is not part of this evidence.

### Exists and is verified

| Element | Check |
| --- | --- |
| `dev.yml` with its services | The containers start and are healthy (`database` with `healthy`); the test directory and Keycloak start apart, with their file and their command (`docker compose -f active-directory.yml up -d`, `docker compose -f keycloak.yml up -d`), **after** the environment, which is the one that creates the network |
| Go backend in a container with `air` | Compiles, starts, connects to PostgreSQL and listens on 11002 |
| Hot reload of `air` | On changing a file it rebuilds and restarts, with an orderly shutdown of the previous process |
| `GET /api/health` | `200 {"database":"ok","status":"ok"}`, and 503 if the database does not respond |
| `go-logs` | Writes `_logs/catalina-support_dev_20260922.log` with `TIMEZONE` applied (`UTC-5`) |
| `go mod tidy` | Generates `go.sum`; the `go.mod` only carries the real dependencies |
| Angular 22 frontend | The development server serves `index.html` and the bundle on 11001, and **forwards `/api` to the backend** (`frontend/proxy.conf.json` → `http://backend:11002`): `GET /api/auth/methods`, `GET /api/health` and `POST /api/auth/login` with `admin`/`admin` respond **200 through 11001** (no nginx), and the enter and exit walkthrough passes **in a real browser** |
| The session shell in the browser | The six routes are served by nginx (200) and the Tailwind CSS arrives compiled. And **tested in a real browser** with Playwright, on PC and on mobile: you enter with the factory account, it is recognized, you change screens and you leave |
| Interface tests with Playwright | `docker compose -f tests.yml run --rm e2e`, with the `directory` profile when appropriate, walks the disposable instance on PC and mobile. The last full matrix finished correctly on **2026-10-08** and cleaned its resources; `docs/ambientes.md` keeps the scope and `docs/prueba-local.md` the historical evidence of each walkthrough. It includes installation from an empty database, session, brand, settings, users, tickets, mail, AD and Keycloak |
| Frontend tests | `npm test -- --watch=false` runs the boundary checker and its isolated cases first, and then the whole Angular suite. It finished correctly on **2026-10-08**. It covers global language, session, themes, shell, settings, first start, users, tickets, mail, attachments, toasts and writing help |
| Tailwind v4 and the themes | Installed and compiling: the served CSS carries **the generated utilities** and the theme variables, and **no component writes a color by hand**. The sources are declared by hand in `styles.css` (`@source './app'`), and there is a Playwright case that **fails if the sheet arrives without utilities** |
| The theme | **Eight themes** (two factory ones with `light-dark()` and the institutional color, and six fixed ones with their palette and their accent), chosen with an attribute on `html` and the chosen theme with an attribute on `html`. **Factory follows the system** —which is not having chosen, and that is why **«automatic» is not shown or chosen**—, the system rules live while nobody has chosen, the choice is remembered in the browser and is applied **before starting** from `main.ts` (not with an embedded script: the production CSP does not allow it). Tested in a browser **measuring the background color**, in light, in dark and with the choice beating the system |
| Production build of the frontend | `npm run build` → `dist/catalina-support/browser`, which is the path `prod.yml` expects |
| Development vhost | `https://dev.catalina-support.example.com` serves the application (200) and `GET /api/health` returns `{"database":"ok","status":"ok"}` **through nginx** |
| TLS of the two domains | Let's Encrypt certificates issued on 2026-09-22 (they expire on 2026-12-21), with automatic renewal configured by `webroot` |
| Security headers and CSP | Present in the development response (`nosniff`, `SAMEORIGIN`, `Referrer-Policy` and development CSP) |
| `v1.0.0.sql` migration | Applied with `psql -v ON_ERROR_STOP=1` **twice in a row** and without error: it creates **21 tables** —templates, accounts and tokens; the **five single-row configuration tables** (`installation_settings`, `ai_settings`, `directory_settings`, `keycloak_settings` and `ticket_settings`); two AI auxiliaries; the ten ticket ones; and `ai_insights`— with their indexes and constraints. It includes its idempotent update and the historical conversion of plain text to HTML without touching a body that already contains attachments. |
| The example data, with its script | `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh` applies the schema, applies `v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql` and **copies the attachment files** to `_files/`, in the folder of their ticket. It leaves **eleven accounts** (`user1`…`user5`, `support1`…`support3`, `dev1`…`dev3`, all with the test password) and **25 tickets** with their history: **123 history rows and 27 comments**, seven with an **internal ticket** (one waiting for Development, another resolved, another returned to Support), **five attachments that download identical** to those of the repository, and the numbering counter at 25. Checked by entering through the API with `user1` (sees their five tickets) and with `support1` (the whole inbox) |
| `shared/auth` | Signs and validates the session token (HS256, 10 hours, `sub`), takes the `Bearer` from the header and compares secrets in constant time. **20 tests in green**, including expiration, another secret, the `none` algorithm, the two empty secrets and the five of the OIDC `state` |
| `shared/authz` | Checks the role and responds 403 with `auth.forbidden`. **5 tests in green**, including the case of a route mounted without authentication |
| Authentication middleware | Reads the account **on every request**, leaves the identity in the context and responds 401 (`auth.session.invalid`/`auth.session.expired`) or 503 if the database does not respond. **8 tests in green** |
| `go test ./...` | The full backend suite finished correctly on **2026-10-08**. It covers authentication and authorization, mail, settings, users, tickets, AI, persistence and middleware; the records of each module keep the particular cases without duplicating here a volatile count |
| The four `mail` endpoints | Tested end to end against development, with a factory token signed by hand (there is no entry endpoint yet): without a header, with a token with garbage and with a token from another secret → **401 `auth.session.invalid`**; with a valid token → **200** and the 22 templates; invented marker → **422 `mail.marker.unknown`**; empty subject → **422 `mail.subject.required`**; saving without the link → **200** with `missing: ["enlace"]`; restoring → the factory text returns and `edited` becomes `false`; test from the factory account → **422 `mail.test.noEmail`**; key or language that do not exist → **404** |
| The account is read on every request | The middleware resolves it with the loader of `main.go`: the factory account has its own identity and any other is read **from the accounts table on every request**, so deactivating or changing a role take effect immediately |
| The log after the tests | Not a single match of `Bearer`, `eyJ` or the word `token`: tokens are not recorded |
| The local path of `auth`, end to end | Enter with the factory account and with a local account, `me`, `exit`, forget the password, set it from the mail link and change it from inside. **Tested with curl against development**, with the link read from the test mailbox |
| The AD path, end to end | **The five cases of section 5.4 of `docs/modules/auth.md`, checked one by one by curl** against the OpenLDAP of `active-directory.yml` and with the row that remains in the database: the **automatic sign-up** (a person from the directory without an account enters and their account is born with role `usuario`, origin `ad`, their external identifier and **no password**); the **linking** (a local account with the mail of someone in the directory enters with the password from there, moves to `ad` **keeping their role** and their local password stops working); the **update** (the surnames of the directory replace those from here); **the mail that changes in the directory** (I changed a person's `mail` with `ldapmodify` and they entered **the same account**, because the external identifier rules); the **reactivation** (deactivated here, the directory returns it on entering, with the `users.directoryMayReturn` notice already really exercised); and the **directory down** (an AD account → 503 `auth.directory.unavailable`; an unknown mail or a local account → 401 `auth.invalidCredentials`, and a local account with its good password still enters) |
| The password links | One link, one use: the same token twice → 401 `auth.token.used`; an invented token → 401; the sign-up one expires in 24 hours and the recovery one in 1, and they say so in the mail |
| The Keycloak path, by curl and in the browser | **By curl**, against the development Keycloak: `GET /api/auth/keycloak/start` → **302** to its entry screen with `client_id`, `redirect_uri`, `response_type`, `scope` and `state`; a return with an invented `state` → **302 to `/login#error=auth.oidc.state`**; and a return with a good `state` and a code that is not valid → **302 to `/login#error=auth.oidc.rejected`**, which is the proof that the exchange is really done against Keycloak. **In a real browser** (Playwright): the button, the Keycloak screen, the return, the session and **the fragment erased**. And **without Keycloak configured**: `methods` says `keycloak: false` and its route responds **404**, so that installation does not offer the button or have the path |
| Attachments in comments | **Three doors and a single rule**: paste, drag and choose with the button, with the same closed list of extensions and the same notice. Checked in the browser with **a real paste event**: the file appears in the list, the comment is published with its files, and **a comment without text** —empty body in the database— is read by its attachments **without leaving an empty paragraph**. And **the failure path**, cutting the upload in the browser: the notice says which file is missing, the comment remains published **only once** and the retry button uploads it to the same comment |
| Attachments inside the text | The formatted editor: **paste, drag and the button put the file where the cursor is**, the image **is seen while writing** and **after saving** with its size, a PDF comes out as **a link that opens the viewer** (larger, with **Download** and **Open in a tab**) and a Word **downloads**. Also checked: **what is saved is the text with its references and nothing more** (no `class`, no `src`, no `blob:`), **an unnamed attachment goes at the end**, **the search does not find by a tag**, and **editing** a comment with an image **shows the preview in the editor**. In the backend, `go test` covers the sanitizer: six forms of forbidden HTML give **422 `tickets.body.notAllowed`** without leaving a row |
| The two factory logos | That the application shows **the one that corresponds to the theme being viewed**: with a light theme it loads `logo-catalina-support-light.png` and with a dark one the `dark` one, checking that **the image really loads** (`naturalWidth`) and that they are two different files. Both come from the repository (`frontend/public/`), with a transparent background, and the case first removes any own logo and leaves the installation as it was |
| The password policy and the link | **Minimum 8 characters** (lowered from 12 on 2026-09-25), checked in the browser: a short password is rejected with its key **and does not spend the link** —it is tried again on the same screen and it enters—, and eight characters are valid. The case uncovered that the link was spent before checking the password, and it is corrected |
| The production deployment | **`scripts/prod-build.sh`** builds and publishes the two artifacts in `/srv/catalina-support` (rotates the previous one **by copy**, verifies that the artifact exists and stops if not, and writes `BUILD_INFO` with the commit, the branch, whether the tree was dirty and the date) and starts the containers with `prod.yml` restarting **only the backend**. The three production containers are up and healthy, and `admin` enters with the password of `config/env/prod.env`. **The domain still responds 503**: 1.0.0 is not closed. **This deployment is postponed** until the product is finished (owner's decision, 2026-09-30) |
| The system version | Published in the public brand (`GET /api/settings/brand` → `"version": "1.0.0"`) and read in the two places: **in the exit row of the side menu, right-aligned**, and **in the footer of the entry screen**. Checked in the browser on PC and on mobile, against what the API says —not against a number written in the test—, that it is text and not a link, that it is read over the menu background, and that **collapsed it is not shown**. And that without a brand no number is shown: the tab is not left nameless and the footer without a version, instead of inventing one |
| The installation name, by curl and in the browser | Changed by the API and **read before entering** (`GET /api/settings/brand`): a name of 61 characters → **422 `settings.name.tooLong`**, and a blank one → `Catalina Support` returns. And in the browser: the side menu, the tab and the entry screen show it, and when emptying the field the factory one returns |
| The database backups | `scripts/backup-db.sh dev`, really executed: it leaves the dump in `/srv/catalina-support/backups/`, **checks that it can be read** (`pg_restore --list`) and **it was restored to a new database**: 11 tables and the same records as the real one (tickets, accounts, templates and attachments, counted one by one). And the retention, tested with a backup from twenty days ago: it is deleted, and the one from three days ago and the one from today are not touched |
| The attachments, in their folder, by curl | A ticket really created and **its attachment uploaded through the API**: the file appears in `_files/2026/CS-2026-0001/<generated name>.png` and that **relative path** remains in the database; the attachment uploaded **to the internal** goes to its own, `_files/2026/INT-CS-2026-0001/`; and **both download identical** to the one uploaded, compared byte by byte (`cmp`) |
| The three directory actions, by curl | Against the real directory and realm: signing up an `ad` account by hand and changing the origin of a local account to `ad` and to `keycloak` → **422 `users.origin.byDirectory`** (and the change to `local` still works, **200**); reactivating a deactivated **AD** account → **200 and active**, because the directory finds it; reactivating an account **whose mail the directory no longer knows** → **422 `users.directory.notFound`**; and reactivating a **Keycloak** account by hand → **422 `users.directory.activatesItself`** |
| The accounts and their permissions | Local account sign-up → 201 and mail; repeated mail → 409; misspelled mail → 422; invented origin → 422; sign-up with origin `ad` or `keycloak` → 422 `users.directory.notFound`; Support creating `soporte` → 403 `users.role.notAllowed`; any user signing up → 403 `auth.forbidden` |
| Deactivating blocks instantly | With the account deactivated in the database, **the token it already had stops working** (401 `auth.session.invalid`) and on entering it is told that it is deactivated (`auth.accountInactive`) |
| The test mailbox | All development mail reaches it, wherever the address comes from, and it is filtered by recipient or subject. **Not a single mail goes out to the internet** |
| The log | It records correct entries, failed attempts with their key and their IP, and each mail with its recipient. **Not a single password, not a token**: checked by searching the files of `_logs` |
| The brand of the installation | The logo is uploaded, served and removed: a real 1254×1254 JPEG (203 KB) uploaded through the API, served **identical** to the original, with its version stamp; an SVG served with `Content-Security-Policy: sandbox`; a file that is not an image → 422; one of 1.6 MB → 413; and on removing it, the installation returns to the factory logo |
| The installation name | Configured from the Settings screen and **read in the three places**: the side menu, the browser tab and the entry screen —which is painted **before entering**, and that is why the name travels in the public brand—. With its limits: over 60 characters → **422 `settings.name.tooLong`**, and **empty returns the factory one**. Checked in a real browser |
| The institutional color | Tested with a color that is not readable (yellow): the backend derives the light theme one darkened and leaves the dark one, and the button text still contrasts. Tested in the browser |
| The logo on the entry screen | Playwright checks that **the image really loads** (`naturalWidth`), that it goes in the box with its border, that it measures between 120 and 160 px high, that an own logo replaces the factory one and that on removing it it returns |
| The shell with the side menu | Tested on PC and on mobile: the three zones with the logo at the top and the controls at the bottom, the collapse **that is remembered** on reload, the mobile drawer that opens and closes, and that **the session screens carry no menu** |
| The menu per role | A user **does not see Settings** in the menu, sees their own role, and if they type the address by hand the guard takes them to «no permission». Tested by really creating the account and entering with it |
| The Settings screen | A logo is uploaded **through the form**, it is seen in the preview and in the menu without reloading, and it returns to the factory one. The institutional color has **a preview requested from the backend** —a pure yellow is shown darkened for the light themes— and can be discarded |
| Authentication method, from the screen | The **entry method** is changed in Settings and **takes effect without restarting anything** (the entry path reads the database on every attempt): with `ad` the entry screen shows the notice that the password is the organization's and removes the link to recover it, and with `keycloak` there is no form and only the button remains, plus the **factory account door**. Checked in the browser, and **the suite leaves the method where it was**. **The secrets never leave**: `GET /api/settings` returns `passwordSet` and `secretSet` and no value, and saving with the empty field **keeps** the one that was there |
| The two connection tests | The «Test the connection» button of the two cards, against the real directory and realm: it answers yes when the service is there and **with its key when it is not** —tested by pointing the server to one that does not exist—. **It saves nothing**: on reloading, the configuration is as it was. It skips itself if the service is not started, and the whole suite runs without it |
| The time zone and the public address | In **Settings**, since 2026-09-29: **the time zone** is chosen from a list of zones (IANA) with a search engine —and the time it is there and its offset are shown— and **the public address** (scheme, host and port, `localhost` included), with the notice that **if it is not https the session and the password travel unencrypted**. **The zone decides how all dates are read**, in the interface and in the mails, and **the stored ones remain in UTC**: changing it moves no ticket. **The address is the basis of the mail links and of the Keycloak return**; `PUBLIC_APP_URL` remains as a fallback (`docs/modules/settings.md`, decisions 14 and 15) |
| The outgoing mail, in the database | The `SMTP_*` variables **were removed from the environment**: the SMTP lives in `installation_settings.smtp_*` and is configured in **step 4 of the first-start view** or, later, in **Settings**. Checked in development: the example data file leaves the test mailbox set (`smtp_host='mail'`, `smtp_port='1025'`) and the mail goes out through it (`docs/primer-arranque.md`, section 5) |
| The AI engine, configurable from the screen | **Its address and its model are configured in Settings**, with **its card and its «Test the connection» button** (it tests `<address>/health` with a short time and **saves nothing**): the `ai` module reads them **on every request**, and `AI_URL`/`AI_MODEL` remain **as a fallback**. The test skips itself if the engine is not started (`docs/modules/settings.md`, decision 16, section 5.13) |
| The first-start view (`/setup`) and the seal | In an **unsealed** installation —`installation_settings.installed_at` null— the application takes you to **`/setup`**, which asks in **five steps** for installation, entry, location, mail and AI, and at the end **seals** the installation: the view does not return and its API responds **409 `setup.alreadyInstalled`**. **An installation that was already configured is sealed on updating**, and in development **the examples file seals**, so that the wizard does not appear on every start. Checked in the browser and through the API, including the 409. `docs/primer-arranque.md` |
| The `users` endpoints | Tested end to end: the list with filters and search, the record, the changes, the own profile, deactivating and reactivating. And the limits: Support does not change the role or the mail of another (**403**) nor sees anyone's record, nobody deactivates themselves (**403 `users.selfDeactivation`**), the state is not changed by the `PATCH` (**422**), the invented origin (**422**) and moving to directory is rejected while its query does not exist (**422**) |
| The mail editor | The screen of the `mail` module: the fourteen mails on the left, **the two languages alongside** —stacked on mobile, measured by the test—, the body with its format buttons **wrapping the selection**, the markers that **are inserted when pressed**, the **preview that the backend renders** with the example data of the test, and the buttons to save, return to the factory one and **send me a test**, which is read in the mailbox. Tested end to end |
| The installation language | It is global for interface, mails and new summaries; the accounts do not keep their own language |
| The ticket screens | The **inbox** —a single one with three names, table on PC and cards on mobile, with state chips, search, pagination and the Administrator's type chip—, the **sign-up** with drag and drop, and the **detail**: a timeline with the comments and what the system did, the record with the actions —on the right on PC and below on mobile, measured by the test—, **the double view** when there is an internal one and the attachments **with preview**. Tested on PC and on mobile, including that **the user does not see the internal one**, that **Development does not write in the main one** and that **the Administrator is offered no button** |
| The assignment and the prefix, from Settings | The Settings screen already has the **prefix** —with the notice that it does not change the numbers already issued— and the **assignment and the notice** of each ticket type, with the «to the assignee» option deactivated when it is not assigned |
| The three user screens | The **list** (table on PC and cards on mobile, with the role, origin and state chips, the search and the pagination), **the record** and **the own profile**, tested on PC and on mobile. Support **edits inline** from the list and has nobody's record; deactivating themselves is not offered; the factory account has no profile or link to it, and its name in the menu is not a link |
| The menu reaches the bottom | The drawer **measures the visible area and scrolls inside**: on a mobile with a browser bar (839 px visible out of 890) the exit button was off screen and could not be pressed. Playwright found it, and with it was also corrected that on PC the bar stretched to the whole content |
| A single `h1` per screen | The cards of the product screens title with `h2`: the users one had three `h1` and the profile as many. The test said so when finding two headings with the same name where it expected one |

| The ten `tickets` tables | Created in `v1.0.0.sql` and applied **twice in a row** without error: the two ticket types, the conversation, the attachments, the history and the numbering counter, with their constraints (the internal one cannot be `escalado`, exactly one destination in what hangs from a ticket, the escalation reason cannot be empty) |
| The text and code extensions | Checked end to end through the screen: a `.sql` **is taken** —without the rejection notice—, **its reference is saved** in the comment, **it remains among the ticket attachments** and **on pressing it it downloads** with its name; and an `.svg` **is still rejected** before uploading it. The **forty-two** new extensions are the same in the two places —the editor and the backend—, and the editor checks it on its side with its unit tests |
| The `ai` module, with the real engine | **Checked end to end against the container**: on creating a ticket, `Pedir` returns without waiting for the generation and the summaries go from `pendiente` to `listo`. On the screen, the list and the record show reason and last action; the internal one keeps the escalation as its reason. AI is mandatory to complete an installation. If the configured engine falls afterwards, the jobs remain `sin_motor` or pending and the help desk keeps operating while the monitor tries to recover it |
| The AI engine, measured | The `catalina_support_ai` container (llama.cpp `server-b11206` + Qwen2.5-1.5B-Instruct Q4_K_M) has a **1500 MiB cap** (`mem_limit: 1500m` in `ai.yml`) and **no CPU cap** —the compose does not limit CPU—, **without publishing any port**. Measured while running: **~1.44 GiB**, **98% of the cap** (`docker stats` gives 1.441-1.448 GiB; the `VmRSS` of the process, 1.51 GB). **It takes between 12 and 24 seconds per field** and about 95 s on a ticket of 6,000 characters; the second call of the same ticket is ~7 times faster (prefix cache). Its health check answers in 325 ms while it is generating |
| The model's JSON, without trusting it | **With `response_format` and the demand written at the end of the message, 4 of 4 valid responses** with a ticket of 6,000 characters; with only one of the two ways, the model wraps the response in a code block, invents the keys or answers in a single language. The module rescues the first JSON object and retries; when it is not valid, the log says **the shape** of what arrived (keys and lengths) and never its content |
| The categories and the tags | **Checked by the three layers**: the migration creates «General» and puts it on everything that exists —simulating an old database, the ticket without a category ends in «General» and the column remains `NOT NULL`—; the sign-up **without category is rejected** (`tickets.category.required`) and **a user cannot create or remove** categories; dirty tags are saved normalized (`["Red Wifi","  Licencias  ","red-wifi","Ñoño_Ütil"]` → `['red-wifi','licencias','nonoutil']`); a ticket with 40 characters of tag is rejected; removing the last active one is rejected; `?category=` and `?tag=` filter; and **searching «red» finds the category «Red» and the tag `red`**. On the screen: the sign-up asks for the category, the tags are normalized when written, the chips come out on the line, there are filters for both and the catalog is maintained by whoever is allowed |
| The mentions and the observers | **Checked by the three layers**: the sanitizer accepts `<span data-mencion="12">` and rejects what is not an identifier; on commenting with a mention the observer is created, the ticket **enters the «I observe» of that person and not their «Assigned»**, removing the observer leaves their entry in the history and **does not return** when resending the same comment, a user who tries it receives `tickets.mention.notAllowed`, and the notice `ticket.mentioned` reaches the test mailbox. On the screen: the people search engine, the name highlighted inside the comment, the «Observers» line of the record with its cross and the view chip |
| The copy-number button | Icon only, and on pressing it a green check for two seconds. The test **measures the painted color** and compares it with the `--exito` of the theme, not with the class name —and there it came out that the button has a transition, so the first reading takes the color halfway— |
| The «mine» filter | `GET /api/tickets?mine=1` returns **exactly what the query says**: checked against the database with `support1` (**16 tickets**: 9 main ones assigned, open or commented by them and 7 internal ones escalated or commented), with `dev1` (**3 internal ones**) and with `user1` (**5**, all theirs: the user still sees only their own with or without `mine`) |
| The `tickets` backend, end to end | **78 checks against development**, reading the mails of the test mailbox: the sign-up with its **number** (`CS-2026-0001`, and the 9999 that grows to five digits), the **round-robin assignment** that alternates between the active technicians, the triage, the **escalation** with its mandatory reason, the synchronization rules (the internal one on hold and the internal one resolved return the main one to `en progreso`), the **return** on closing the internal one without resolving it, the **re-escalation** with the same internal one, the closure, the reopening with clean dates, the comments (editing and deleting by emptying the text), the attachments (upload, download, the extension outside the list, the `svg` and the 25 MB) and **the permissions of the four roles**, including that the user does not reach an internal ticket either by its number or by the path of an attachment |
| The seven `tickets` notices | They go out when the transition that causes them says so, and they have been read in the mailbox: to the assignee, to the user (on hold, resolved, closed), to Development (escalated, also on re-escalating) and to Support (Development needs something, and the ticket has returned to their inbox) |
| The sites that already existed | `example.com`, `dev.example.com` and `other-app.example.com` still respond 200 after installing the vhosts |

### Exists but is NOT verified

- **Reproducing the current deployment from scratch is blocked**: the script publishes the compose in
  `/srv/catalina-support` and uses that folder as the build context, but the Dockerfiles need
  `frontend/package*.json`, the frontend sources and the backend sources. The script does not publish
  those sources. Finding by reading on 2026-10-02; it has not been run against production. A proposal
  and approval are required before correcting its behavior.


- **The public production path, postponed**: the frontend container is up with its nginx, but **the
  domain still responds 503 on purpose** and **the deployment is not resumed until the product is
  finished** (owner's decision, 2026-09-30), so the public vhost and `config/nginx/frontend.prod.conf`
  serving the application have not been tested.

### Limitations and postponed work

- **Postponed on purpose** (owner's decision, 2026-09-30): **the production deployment and opening the
  domain are not resumed until the product is finished**. The deployment of 2026-09-25 is done and
  verified from the inside, and the domain remains at 503 on purpose (`docs/ambientes.md`, section 4).
- **A real Active Directory and Keycloak**: what is tested is OpenLDAP —which is spoken to the same
  way— and a development Keycloak, so the code is really tested, but what is specific to an AD —its
  cross-domain references, its particular attributes, its password rules— is not covered, nor the
  particularities of an organization's Keycloak (`docs/modules/auth.md`, section 11).
- **PKCE in the Keycloak path**, proposed and not done: the client is confidential and the secret
  already protects the exchange, so it was not in the approved document. If it is done, the
  `code_verifier` travels inside the signed `state` (`docs/modules/auth.md`, decision 40).

### Pending decision

No architectural decisions remain open. The production work is postponed by a decision already taken,
not pending design.

## Disposable test instance (`tests.yml`, 2026-10-03)

Project `catalina-support-tests`, with its own PostgreSQL, single-use migration, backend, frontend and
Mailpit. The backend waits for the migration to complete; the runner waits for the backend health. All
data, attachments, logs and dependencies have project volumes; the code enters read-only. There are no
published ports or shared AI network. LDAP and Keycloak enter only with the `directory` profile. The
Playwright reports are written to `tests/e2e/resultados/`, which is kept after
`docker compose -f tests.yml --profile directory down -v`.
The preparation, execution and cleanup walkthrough is in `docs/ambientes.md`, section 9.3.

### Corrected backend process (2026-10-06)

The code remains mounted at `/app:ro`. The suite replaces the image command —Air, useful for
development— with `go run .`, appropriate for an immutable execution. Go uses `backend_go_mod` and
`backend_go_build`; its temporary artifact stays outside `/app`. The mount `backend_tmp:/app/tmp` and
the volume `backend_tmp` disappear.

The shared development image, `.air.toml`, the migration dependency, the networks, the data and any
development or production container do not change.

The review confirmed a single compilation per pass, only the two Go caches already existing and a full
cleanup with `down -v`. No architectural decisions remain open.

The implementation was verified from new volumes: the migration finished before the backend, the
service compiled with `go run .`, responded health and allowed completing the 216 cases of the pass
without optional profiles —177 approved and 39 skipped—. The subsequent `down -v` removed all
resources and `docker compose -f tests.yml ps -a` was empty.

## 14. Implemented architecture: mandatory AI and a single global language

### 14.1 Components and dependencies

`tickets` keeps its summary interface and knows no providers. `ai` keeps the queue, normalization and
persistence and adds adapters for OpenAI/DeepSeek/OpenAI-compatible and Claude. `settings` owns the
configuration, encryption, privacy confirmation and required state. `mail` owns templates and
validation of their drafts. `main.go` connects the interfaces without new cross imports.

The local engine remains in `ai.yml`, but its process 1 is a small administrator that downloads models
and supervises a child `llama-server` process. It does not mount `/var/run/docker.sock`, does not
publish its API and requires `AI_MANAGER_TOKEN`. The backend can only request catalog, progress,
download, activation, test and deletion through the private network of the environment.

### 14.2 Local isolation

Development and production use different projects, networks and volumes. The container name stops
being global and the internal URL is resolved per service. The catalog is the same, but each volume
decides which files it keeps and which is active. The memory ceiling is 8 GB and the administrator
does the RAM/disk check before starting the model.

The default local model continues to be Qwen2.5-1.5B-Instruct Q4_K_M. `ai.yml` no longer fixes the
model path in `command`; the administrator chooses it after validating the catalog versioned in the
repository and the checksum fixed there.

### 14.3 Data

The detailed AI configuration leaves the two columns of `installation_settings` and moves to its own
row of `settings`. The credential is stored as authenticated encryption, with version, nonce and
ciphertext; the master key `AI_CREDENTIAL_KEY` lives only in the environment. `TOKEN_SECRET` is never
used to encrypt it. It is 32 random bytes encoded in base64 and the process refuses to start if it is
missing or has an invalid format. A valid key that does not open existing data puts the installation
in the required-configuration state.

`users.language` disappears. `installation_settings.language` remains as the only source. The public
brand includes that language to paint the entry and the shell before there is a session.

The installation also keeps an increasing version of its global configuration. All authenticated
responses publish `X-Catalina-Language` and `X-Catalina-Settings-Version`; the interceptor updates the
language signal when it observes a new version. The token is not invalidated, no polling is added and
a tab adopts the change on its next request.

`ai_insights` stops storing two texts per field: it stores one, its language, provider and model. On
changing the language, an update marks the texts of another language for regeneration without deleting
the ticket history from which the assignment is rebuilt.

### 14.4 Availability and updating

AI is a configuration requirement, not a synchronous dependency of each action. A configured
installation keeps saving tickets during an outage. A health monitor with increasing wait reactivates
the queue when it returns. With a local engine or own server it processes the backlog; with a provider
it leaves the batch paused until the Administrator sees the volume and confirms the possible cost. Only
the initial configuration state blocks entry.

Sealed installations are migrated without opening `/setup`. If there is no tested configuration, the
middleware allows health, brand, entry/exit and the minimal Settings routes for
Administrator/factory account; the other routes report that AI is missing.

### 14.5 Security and tests

The HTTP clients limit time, response body and redirects; they never record prompts, responses or
secret headers. The own server allows authentication without relaxing TLS. The adapters are tested
against fake servers and the administrator against small test files; CI and the local suite do not call
providers or download the real GGUF files. The catalog fixes source, license, size and checksum, and
records the acceptance per version before downloading.

This proposal affects backend, frontend, the open migration `v1.0.0.sql`, `ai.yml`, the environment
configuration and tests. It does not change public nginx, the authentication paths or tokens, or the
content of attachments.
It was decided by the owner on 2026-10-06. The review closed keys, custom templates, recovery costs,
licenses and propagation of the language to open sessions. No decisions remain open; proposal approved.

## 15. Implemented closure of block 4

### 15.1 Redirects and credentials

Generations and tests that carry Bearer, their own header, Basic or a Claude key **do not follow
redirects**. A 3xx response is treated as a failed connection and forces the final URL to be configured.
The rule applies to OpenAI, DeepSeek, Claude and OpenAI-compatible, when testing and when summarizing.

Public GGUF downloads can follow redirects because Hugging Face needs them, but with a finite maximum
and without Catalina Support credentials. Each hop can only use `http` or `https`; the client limits
connection and headers, and the download does not exceed the declared size.

### 15.2 Supervision of the local process

After activating a model, the administrator waits for the child process. If `llama-server` finishes
while that model remains active, it restarts it indefinitely with waits of 1, 2, 4, 8, 16, 32 and then
60 seconds between attempts. Only one supervisor exists per active model. Changing model or shutting
down the administrator cancels the previous supervisor; a termination caused by that change does not
count as a failure.

The catalog exposes `active` and `healthy` separately and shows the last operational error. The health
endpoint of the administrator still indicates that the administrator lives; the test generation
confirms that the model serves.

### 15.3 Effective disk and memory

To download, only the bytes missing from each file are added, counting its `.partial` files, plus
256 MiB of margin. A partial larger than the declared size is discarded and restarted from zero. If it
reaches the expected size but fails format or checksum, it is also deleted so that the next attempt
downloads a clean file.

Available memory is the lowest value between `MemAvailable` and the remaining capacity of cgroup v2 or
v1 when there is a finite limit. If there is no cgroup, `MemAvailable` is used. The model does not start
when its estimate exceeds that effective capacity.

### 15.4 Acceptance criteria

1. The four providers are tested with fake HTTP servers and without commercial network.
2. A 3xx response does not cause another request or resend credentials.
3. A response larger than the limit is rejected without growing indefinitely in memory.
4. Authentication reaches only the configured fake server.
5. The administrator resumes a valid partial and reserves only the pending bytes.
6. The simulated cgroup limits prevail over the machine memory.
7. The unexpected death of the child starts the increasing wait; activating another model cancels the
   previous supervisor without leaving two processes.

Configurable redirects, Docker socket, GPU, spending limits and calls to real providers are out.

### 15.5 Decisions and review

The owner chose on 2026-10-07: reject redirects in calls with credentials (1A), supervise and restart
the local engine (2A), measure pending disk and effective cgroup memory (3A) and cover everything with
isolated tests (4A). In the review they chose indefinite retry with a maximum wait of 60 seconds (5A),
separate `active` and `healthy` states (6A), delete corrupt partials (7A) and allow at most five
HTTP/HTTPS hops in public downloads (8A). No decisions remain open. Production and any real download
remain outside this block.

## 16. Transactional local activation implemented

The administrator treats the model change as an operational transaction: healthy source, preflight,
stop, real measurement, start and confirmation of the destination. It keeps the identifier and the
specification of the source until the destination is healthy. The `active` file represents the last
healthy model, not the last attempt.

The preflight adds effective available memory and RSS of the current child. It does not count the
memory of the administrator itself. After finishing the child it reads host and cgroup again; if the
requirement is not met, it restores the source. Any process or health error follows the same rollback
and waits for its health before responding.

The memory error contract carries a stable key, required bytes and available bytes. It exposes no PID,
paths or host details. The same rule serves `/setup` and `/settings` through the current adapter. No
model concurrency, memory reservation, GPU or forced use of swap is added.

The real test changed 1.5B→3B taking advantage of the recoverable RSS of the source. The 7B attempt
returned 6,442,450,944 required bytes and about 5.07 GiB available without interrupting the healthy 3B.
The deterministic tests also cover the insufficient second measurement and the health failure, with
restoration of the source before responding.

## 17. Block 6: full review of the frontend

The block removes the “format review” loose end through a complete inventory of all screens and their
total correction. The implementation remains in Angular, Tailwind and the own components; a repeated
pattern is solved in the shared layer and a single composition in its screen. No visual dependencies or
a screenshot comparison tool are added.

The architectural boundary is strict: only presentation, semantics, accessibility and responsive. API,
backend, database, permissions, states and walkthroughs do not change. A finding that needs one of
those changes is documented in another cycle before writing that code. The evidence uses the disposable
instance, Chromium and stable tests; production and the known reproduction defect of `prod-build.sh`
remain out.

The owner chose 1A–14C and approved the proposal without open decisions. The implementation
found a single overflow: the shared role selector in `/users` did not fit in 412 px with the English
texts. `Conmutador` now limits its width and distributes options over several lines.
The correction passed the three resolutions, both factory themes and the roles that use that list.

The final validation passed 220 frontend unit tests and the full Playwright: 180 approved cases and 38
expected skips. The eight themes kept their contrast measurement and no visual or accessibility findings
remain open. Backend, API, data, permissions and flows did not change.

## 18. Block 7 implemented: documentation coherence and frontend boundaries

### 18.1 Findings that motivate the block

The product and the writing help of the previous block are implemented, but the architectural inventory
does not describe the whole current state coherently:

- section 13 said that an installation worked entirely without AI configured, although AI is
  mandatory and the middleware limits a sealed installation that does not have it;
- it presented the full walkthrough of `/setup` as not automated, although the disposable suite already
  walks it;
- it kept as current figures and test dates from before blocks 5, 6 and the writing help;
- some historical records called a replaced and closed decision «pending»;
- section 12 declared logging and boundary automation unresolved, although the owner already decided
  both in the review of this block.

These findings are corrected visibly. History is not rewritten: decisions and amendments are kept, it
is marked which one replaced them and statements that appear to remain current are removed.

### 18.2 Boundary checker

The frontend adds `scripts/check-boundaries.mjs`, its own and versioned, without an additional package,
which reviews all the
TypeScript files of `frontend/src/app`, including their tests. It ignores dependencies and generated
artifacts. It applies these directions:

1. the root of `app` can compose `core`, `shared` and any module;
2. each folder of `modules/*` can depend on itself, on `core` and on `shared`, but not on another
   module;
3. `core` can depend on `shared`, but not on a module;
4. `shared` cannot depend on `core` or on a module.

In addition, each module can only consume its homonymous prefix under `/api`: `mail` uses `/api/mail`;
`tickets`, `/api/tickets`; and `users`, `/api/users`. The checker recognizes text literals and
interpolated templates. The cross-cutting routes of `core` remain outside that homonymous
correspondence because session, installation and configuration live there.

There is an explicit list of exceptions, initially empty. A future exception must indicate
origin, destination and reason, and changing it will require maintaining this document. Each violation
reports file, line, rule breached and expected correction; all are shown in one run and the process
ends with a non-zero code.

The command can be run separately and is part of both `npm test` and `npm run build`.
`scripts/check-boundaries.test.mjs` demonstrates a cross import, an inverted layer dependency, a
foreign API route, an interpolated template and a valid set.

### 18.3 Browser logging

No logging service, dependency or provider is added to the frontend. Functional errors are presented
through the existing components and messages. The only technical use of the console allowed is the
fatal failure that prevents Angular from starting, in `main.ts`. The boundary checker does not become a
general rule about console calls.

### 18.4 Documentation audit

All documents with `as-built` status are reviewed, not only this file. The correction:

- aligns the mandatory nature of AI and the behavior of an unconfigured installation;
- updates walkthroughs that are already automated;
- replaces volatile test inventories with the command, the date and the overall result of the last
  relevant verification;
- keeps the historical records, clearly marking the replaced decisions;
- updates `docs/README.md` and the correspondence table of `AGENTS.md`.

If the audit discovers a new difference between documentation and code that implies changing behavior,
it will be reported and another cycle will be opened. It will not be corrected as a silent part of this
block.

### 18.5 Verification and closure

Before returning the document to `as-built` these must pass:

1. the checker separately, including its isolated negative tests;
2. `npm test` and `npm run build`, demonstrating that both run the boundary;
3. the full backend suite, because the closure describes the whole repository;
4. a final search of documental states and of the recorded contradictions;
5. `git diff --check`.

The evidence is recorded with date, commands and overall result, without keeping breakdowns per package
or file that become obsolete with each new case.

### 18.6 Scope and decisions

The block modified `docs/`, `AGENTS.md`, the frontend scripts and its `package.json`. It did not change
API, backend, database, permissions, interface, user flows or production dependencies.

Production stays entirely out: mail is not configured, the domain is not opened, the `v1.0.0` tag is not
created, it is not deployed and the from-scratch reproduction of `prod-build.sh` is not corrected yet.
Those jobs remain postponed until the closure of version 1.0.0.

The owner chose 1A–5A, 6A–10A and 11A–15A on 2026-10-08 and explicitly approved the document.
During the implementation a `core → app` dependency appeared in the interceptor test; it was recorded in
the header and corrected using the empty router that test really needs.

The verification of 2026-10-08 finished correctly: the checker and its four isolated cases;
`npm test -- --watch=false`, which ran the boundaries first and then the **225 tests** of the frontend;
`npm run build`, also with the boundary first; and `go test ./...` for the whole backend.
The final audit found no documents in `propuesta` or `aprobado`, nor current contradictions of those
recorded in 18.1. `git diff --check` finished clean. No decisions remain open.

## 19. Implementation: publication boundary and source traceability

The opening is prepared on the complete repository, not as an export without history. Before
changing its visibility, the current tree and every Git revision are audited to detect secrets, personal
data, real infrastructure, binaries and third-party notices. The known findings and the stopping
procedure are in `docs/proyecto-abierto.md`, section 6.

The approved cleanup rewrites the commits with personal mail a single time and replaces real
infrastructure references with examples or unversioned configuration. As the rewrite changes
identifiers, the content is finished first, a recoverable private reference is kept, the new history is
verified and only then is it decided to withdraw the reference. A new sensitive finding stops the
process to inform the owner before another mutation.

The first local rewrite was run on 2026-10-10 after creating and verifying a private bundle. The `main`
and `dev` branches stopped containing the personal mail, the real domains and the real paths; the secret
scan only required two exact historical fingerprints for declared development values. The remote
references were replaced through `--force-with-lease`, verified against their new identifiers and the
previous local references were withdrawn after confirming the remote.

On discovering that the repository was already public with the historical MIT license in `main`, a
second private bundle was created and `main`, `dev` and `cla-signatures` were rewritten so that the own
root license and its historical declarations expressed AGPL-3.0-only. The operation kept third-party
licenses and the CLA signature, passed Gitleaks and published `main` first. PR `#2` then integrated the
current tree through a squash. Both recoverable bundles are kept outside the repository.

The visible traceability uses the public brand as the single source: the backend delivers version,
`AGPL-3.0-only` and URL of the corresponding code. The published versions link their tag; development
links the repository and identifies itself as such. These metadata belong to the compilation, not to the
database or the administrative configuration.

The community files and the CLA automation form a maintenance boundary at the root and `.github`; they
introduce no runtime dependencies. The actions are fixed by hash and their minimum permissions are
reviewed. The repository is public and GitHub recognizes AGPL-3.0. Production, the public domain of the
application and any version tag remain outside this implementation.

The coordinated verification of 2026-10-10 passed settings in the backend, the boundaries and the 227
frontend tests and the 12 isolated brand cases on PC and mobile. The independent test of the CLA had
confirmed its fourth check in green before. The E2E environment was removed with its volumes at the
end. On 2026-10-12 the description, the themes, Issues, Discussions, private vulnerability reports,
squash only, automatic cleanup and the effective ruleset of `main` and `dev` were also verified;
`cla-signatures` remained outside and writable by the action.

## 20. Implementation: boundary of the language validator

Language detection belongs to the `ai` module, next to the interpretation of the engine result.
`tickets` and `mail` deliver the global language and receive valid text or the existing stable error;
they do not import the library, do not know its confidence and do not duplicate rules. The detector works
inside the backend process and opens no endpoints, tables, jobs or external communication.

The chosen dependency must admit Spanish and English, expose enough confidence to distinguish a clear
discrepancy and have a license compatible with AGPL-3.0-only. Its version is fixed in `go.mod` and added
to `THIRD_PARTY_NOTICES.md` and `third_party_licenses/` through the current mechanism. The tests replace
the detector when they need to control each branch and also test a small set of real phrases to prevent
the adapter from inverting Spanish and English.

The requested language comes from the global configuration in each operation. It does not reappear in
identity, session or personal preferences. The retries remain inside `ai`: one for `local` and
`remote`, none for `provider`; therefore the interfaces between modules and the persisted attempts
counter of the queue do not change. The translation of context labels is done before crossing the
interface and the user content remains literal.

The cleanup of the expected `record not found` remains inside the `tickets` repository and does not
change the global GORM logger. Production, new observability and changes in the frontend dependency
rules remain out. The owner closed 1A–12A with 10B; no decisions remain open.
### 20.1 Implementation and verification

`whatlanggo` 1.0.1 remains enclosed in `backend/modules/ai/services`; no other module imports the
detector. Its MIT license was added to the generated third-party inventory. The queue keeps a single
persisted attempt for the internal language retry and automatically avoids a second call in `provider`
mode. The full backend suite and `go vet` passed on 2026-10-13.

## 21. Distinguishable health with pending schema

At startup the backend checks the existence of `installation_settings`, an essential table of
`v1.0.0.sql`. Compose starts it only after the automatic runner, but the check keeps a safe diagnosis if
someone runs the binary outside that order.

`GET /api/health` responds `503` with the stable object
`{"status":"error","errorKey":"database.schema_pending"}` while that table is missing. A disconnected
database or a real PostgreSQL error keeps the general classification of database unavailable; the two
diagnoses are not confused. The health queries do not repeat the message on every call. The interface
keeps the existing server-unavailable screen: the operational detail lives in health and logs, without
a new screen.

In the three environments backend and frontend wait for the automatic `migrate` service. The migrations
do not live inside the HTTP server: the independent Go command maintains `schema_migrations`, creates
the backup and finishes before enabling the application.

## 22. Implementation: expected absences without false errors

A query whose contract admits that there is no row is not a database failure. The repositories that
already translate that situation into a domain result will use a query that delivers zero rows and
`RowsAffected == 0`, without producing `gorm.ErrRecordNotFound`. This covers optional configuration,
templates and tokens searched by key, tickets or optional relationships, categories and accounts that
are checked before being created or linked.

The decision is taken in each repository method; `IgnoreRecordNotFoundError` is not configured globally.
A connection, permission, SQL, scan or any other GORM error is propagated and keeps its current level.
Nor is the public meaning of found, absent, duplicate or unauthorized changed.

The repository tests will cover at least one absence for each modified pattern and will check two
results: the existing domain value and that the GORM logger receives no error-level event. They will
also include a real SQL failure to prove that it was not silenced. No new observability system or
logging dependency is added.

API, data model, migrations, interface, permissions and business rules changes are out. The owner chose
to correct all the expected absences, do it query by query, test behavior and log, and then repeat the
full block 8 (1A–4A). No decisions remain open.

The implementation replaced the `First` reads that translated `gorm.ErrRecordNotFound` with
`Limit(1).Find`, an immediate check of the error and then `RowsAffected`. No `First` read remains in the
module repositories. The tests with real PostgreSQL confirmed the absence without an error event for
categories and for the two identity lookups; a nonexistent table still reached the logger. The full
backend suite and `go vet` passed on 2026-10-13.

The automated tests cover the three states: unreachable PostgreSQL, valid connection without
`installation_settings` and schema present. They also check the HTTP code, the JSON shape and that the
check does not write the same notice repeatedly.

### 21.1 Implementation and verification

Startup queries `to_regclass('public.installation_settings')` before building the modules. If the table
is missing, it records the notice once and mounts a minimal server: health returns the stable key and
the other API routes return `503`, so that the interface activates its existing server-unavailable
screen. After migrating, the backend is restarted to mount the full routes. The backend suite and
`go vet` passed. The test with an empty PostgreSQL confirmed that no queries to `installation_settings`
or `ai_insights` appear and that the real response is the agreed one.

## 23. Priority inside the tickets module

Priority belongs to the main `tickets` entity; no module, catalog or configurable table is born. The
database stores stable English technical values and the frontend translates their labels. The SQL
constraint, the domain type and the HTTP validation share the four values and the explicit operational
order.

The tickets service authorizes and saves priority plus history in a transaction. After the commit it
requests the critical notice from `mail` through the existing internal interface between modules. The
persistent first-notice mark is updated in the same operation that decides to send it, avoiding
duplicates on concurrent requests. The subsequent SMTP failure does not revert data or add queues or
retries.

The queries apply parametrized filters and an explicit `CASE` for the order, followed by creation date
and identifier for stability. The frontend keeps filters and order in the URL without per-account
preferences. The automatic boundaries remain the same: the tickets screen calls only `/api/tickets/**`
and `tickets` uses an interface to request the mail.

## 24. Verifiable context and versioned AI regeneration

Data ownership does not change. `tickets` reads main, internal, participants, classification, comments,
attachments and history and builds a structured internal contract. `ai` receives that contract, applies
budget, writes, validates and persists the summary. No `ai` repository reads ticket tables and no screen
calls a foreign prefix.

The open migration of `v1.0.0.sql` adds an optional reference from `ticket_history` to the comment that
accompanies the movement. The reference is null for movements without a comment and for historical data
that cannot be linked with certainty. The tickets transaction creates the comment, saves its identifier
in the history and changes the state as a single unit.

`ai_insights` adds a context version. New results are saved with the current version; the previous last
actions become detectable without a volatile global mark. At startup, the existing worker collects the
old rows in bounded batches, avoids duplicating them in the queue and continues after restarting. An
installation serves requests while it regenerates and a fallen engine does not prevent starting the
backend or working with tickets.

The backup of an explicit action is produced with structured data from `tickets`; that is why actor,
action and final state do not depend on interpreting free text. `ai` can improve the sentence and
summarize the associated comment, but only accepts the output if it preserves those anchors. For an
independent comment it continues using normal generation and its existing errors.

The tests cross the boundaries with fake contracts, real PostgreSQL and a simulated restart. They check
foreign keys, atomicity, batches, version, exclusion of duplicates, context budget, local/remote/provider
modalities and that the backend remains available during the sweep. Production, domain and tag remain
out of scope.

The open migration was applied only to development. The new version was persisted and the worker
regenerated the reproduced case without blocking the backend. The affected tests, `go vet`, the 232
frontend tests, the automatic boundaries and the compilation passed; the previous internal-priority
failure documented outside this change remains.

## 25. Implementation: Summary migration and automatic boundary

Modular ownership does not change: `tickets` decides when a change requires Summary or Last action,
builds the authorized source and calls the internal interface of `ai`; `ai` generates, validates and
persists. Removing the manual endpoint reduces the HTTP surface and creates no frontend call to
`/api/ai/**`.

The open migration turns the `motivo` discriminator of `ai_insights` into `summary`. API, DTO, Go and
TypeScript types use `summary`; the escalation reason uses `escalationReason`, while the
`escalation_reason` column needs no change. There is no double reading, alias or transitory
compatibility. The Summary context version is increased so that the worker regenerates all the rows
—also with a provider— in bounded and idempotent batches.

The attempts of each job still have a limit. No cron or infinite queue is born: after a failure, only a
relevant change queues the field again. The persisted state allows showing pending or error in the
record and empty text in the compact consumers. The tests cross migration, repository, restart,
provider, HTTP contracts and frontend boundaries. Production is not touched nor is the `v1.0.0` tag
created.

## 26. Cumulative and auditable authorization

`users` owns `users.is_admin`, the grant or withdrawal history and the closure of the notice. `auth`
authenticates and loads the current identity. `shared/authz` offers orthogonal checks of operational
role and administrative capability. The consumers declare which one they need; they do not read foreign
tables or rebuild the combination.

The frontend receives `role`, `isAdmin` and `factory`. The factory account has no operational role,
`isAdmin = true` and `factory = true`; ordinary accounts always have one of the three roles.

The open migration adds `users.is_admin`, restricts the persisted role to three values and converts any
persisted Administrator into Support + Administrator. It adds administrative events with an actor that
can represent the factory account and an independent read acknowledgement. It preserves identifiers,
sessions, origins and data, and updates seeder, contracts, repositories and tests.

The administrative routes require `isAdmin`; the operational ones keep their rules per role. The union
accumulates capabilities, but the main or internal side of the ticket depends only on the operational
role. The identity is read on every request, without permission cache or deferred revocation.

Event and capability change are written in a transaction. The pending notice derives from unclosed
events and not from an ephemeral toast. The directory synchronization does not touch the local column.

The tests cover restrictions, cumulative permissions and navigation; the migration was verified in
development PostgreSQL. The focused backend suite and the 240 frontend tests passed on 2026-10-18.

## 27. Dashboard module

`dashboard` is a module of its own in backend and frontend. `tickets` implements a read port with rows
and events; `dashboard` applies the scope of the identity, calculates aggregates, comparison, grouping,
delay and CSV. It exposes no ticket writing. `settings` keeps the global thresholds in its single-row
configuration.

The API separates summary/charts, incremental detail and full export. All receive the same normalized
filter and resolve role plus Admin capability on each request. PostgreSQL rebuilds states by cut, active
intervals and temporal facts; Dashboard aggregates the resulting rows on the server and the API does not
deliver raw history to the browser. The periods are converted from the global zone to UTC limits and the
responses include the chosen grouping. The CSV is streamed, escaped according to RFC 4180 and neutralizes
`=`, `+`, `-` and `@` prefixes in text controlled by people.

No chart library is added if the SVG and CSS of the project cover accessible bars and lines. The
automatic boundary will incorporate `/api/dashboard/**` as its own prefix and will forbid other frontend
modules from calling it.

The final location of the read repository under `modules/tickets/repositories` corrects the first
implementation version, which had placed it under Dashboard despite querying ticket tables. The focused
tests and the real PostgreSQL query verify that boundary. The final aggregation in Go is the deviation
documented with respect to the proposal of finishing it in SQL: it keeps a single semantics for summary,
detail and CSV without moving the history to the browser.

## 28. Local acceleration boundary

`settings` persists and validates `auto|cpu|nvidia`; `ai-manager` detects capacity, builds the
`llama-server` arguments, governs the process and reports effective execution. The backend does not
inspect host devices or run Docker. `ai.yml` keeps the CPU image and `ai-nvidia.yml` replaces only the
build and exposure necessary for CUDA, sharing network, token and model volume.

The NVIDIA variant is supported on **Linux with NVIDIA Container Toolkit** and on **Windows with Docker
Desktop WSL2**; in both cases Docker must see the GPU (`docker run --gpus all ... nvidia-smi`). The
engine remains Linux on the two paths: native Windows containers are not used. The verification of
2026-11-05 used an RTX A2000 12 GB, left the model active and healthy and generated the two summaries of
a real ticket.

The mode change is a serialized operation of the administrator. It keeps the specification and the
previous process until the new one passes health and generation; if it fails, it reactivates the previous
one before responding. The public state separates configuration from execution, limits the device to
manufacturer/name and never exposes sensitive inventory.

## 29. SMTP transport security boundary

`settings` owns the `starttls|tls_direct|plain` mode and the administrative acknowledgement; `mail` owns
applying them when connecting. STARTTLS is mandatory in its mode, direct TLS starts encrypted and
neither falls to `plain`. The insecure mode revalidates on each test and sending that all the
resolutions of the host belong to IPv4 RFC1918 or loopback.

To avoid a race between validation and connection, `mail` marks all resolved addresses as valid and
opens the socket against one of them without resolving again. The configured name is kept for the SMTP
dialogue. A later DNS change only affects the next operation, which repeats the full validation. Public
destinations and IPv6 are not admitted in `plain`.

The migration never degrades installations: the true boolean becomes direct TLS and the false one
STARTTLS. Mailpit is an exception declared in development data, without credentials; it does not open an
architectural exception for production. Logs and responses keep user, password and SMTP content out.

## 30. Boundary of the real test mail

`settings` receives the form data and the ephemeral recipient, decides whether the installation is open
or the identity has administrative capability, applies the limit of five attempts per IP and recipient
in fifteen minutes and resolves the saved secret when the field arrives empty. It does not save the
recipient or the form changes.

`mail` receives a complete SMTP configuration, sender, recipient, language and installation name. It
builds the fixed technical message and runs the same transport as normal sends, including the
restrictions of `plain`. The operation ends only after the final acceptance of the relay; it does not
interpret that acceptance as delivery to the mailbox.

The test message does not enter `mail_templates`, does not create a table or cross the boundary towards
the notices of `auth` or `tickets`. The structured log keeps result, server, IP and partially hidden
recipient; it excludes credentials, body and SMTP conversation.

## 31. Boundary of the migrations runner

The runner is a Go command independent of the HTTP server. It shares `shared/version` to know the
maximum version and the development state, but it does not import business modules or use their
repositories. It receives the PostgreSQL connection, the migrations directory and the backups directory
through the deployment configuration. It serves no network, creates no users and does not run the
seeder.

Its sequence is deterministic: it discovers strict names, parses and sorts SemVer, acquires an advisory
lock, inspects `schema_migrations`, rejects downgrades, future files and immutable checksums, identifies
pending work and, only if there is any, creates and validates the backup. Then it runs each file in its
own transaction and records version, file, SHA-256 and date inside that same transaction. It releases
the lock on finishing or failing. A second runner waits for the lock and recalculates the state when
acquiring it.

The development exception is limited to the file whose version matches the current constant: it can be
reapplied if the checksum changed because it is still idempotent. No historical file can mutate. On
changing `Development` to false, even the current file becomes sealed. The runner does not attempt to
interpret statements, fabricate rollbacks or continue after an error.

Compose expresses the dependency through correct completion of the migration service. HTTP availability
does not replace that dependency; the pending schema check remains as a defense if the binary is run
outside Compose. Logs and `BUILD_INFO` allow relating an update with commit and version, without
recording DSN or secrets. Destructive changes follow an expand–migrate–contract cycle between versions to
keep an operational recovery path.

## 32. Demonstration data in English and Spanish

The examples stop depending on a single `v1.0.0_dev.sql`. The distribution adds
`backend/migrations/v1.0.0_dev_en.sql` and `backend/migrations/v1.0.0_dev_es.sql`, in addition to the
public accesses `scripts/dev-seed-en.sh` and `scripts/dev-seed-es.sh`. Both languages keep the same
identifiers, ticket numbers, dates, priorities, persisted states, assignments, relationships and
quantities; only the visible content written by people and the example files change.

The English variant uses `User One` to `User Five`, `Support One` to `Support Three` and `Developer One`
to `Developer Three`. The Spanish one uses `Usuario Uno` to `Usuario Cinco`, `Soporte Uno` to `Soporte
Tres` and `Desarrollador Uno` to `Desarrollador Tres`. Subjects, comments, seeded summaries, categories,
tags, reasons, attachment names and content of those attachments belong to the same language. The
technical identifiers of roles, states and events are not translated in the database: the interface
remains responsible for its localized labels.

In a new database or without a sealed installation, the chosen seeder sets the corresponding language
and completes the installation. In an already sealed installation it keeps the current configuration.
Each execution rebuilds the examples in the chosen language, as the single seeder does today; changing
the language from Settings does not rewrite historical data.

The `*_dev_en.sql` and `*_dev_es.sql` files remain excluded from the production migrations runner. They
do not change the schema, the APIs, the permissions or the functional semantics of tickets. The
documents that still name the single file and command will be updated together with the implementation;
that gap is recorded as part of the finding and is not corrected before approving this proposal.

The owner closed the review with 1A–5A. The two public accesses call a shared internal runner that
explicitly receives language, SQL and attachment source; there is no ambiguous public command. The
structural equivalence between languages is mandatory and no decisions remain open.

The real verification ran both services on disposable databases. Each variant produced eleven accounts,
25 tickets, 123 history rows, 27 comments, seven internal ones and five attachments. A new database
adopted `en` with `seed-en` and `es` with `seed-es`; a sealed installation kept its settings when
changing the set. The resolved Compose configuration contains only `seed-en` and `seed-es`, and the
runner tests confirm that it ignores both development files.

