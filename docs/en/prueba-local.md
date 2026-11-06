# Local test: empty installation, examples and isolated verification

[Español](../es/prueba-local.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-05
>
> **Expansion verified on 2026-11-05.** Section 24 no longer leaves the GPU pending: the NVIDIA
> variant was tested on Docker Desktop WSL2 with an RTX A2000 12 GB, effective mode `nvidia`, active
> model and real summaries. Linux with the NVIDIA Container Toolkit remains the other path; native
> Windows containers are not used.
>
> **Full Playwright cleanup approved, implemented and verified on 2026-10-24.** Section
> 30 removes historical contracts and failures: every case exclusively checks the current product,
> creates its own data and uses a deterministic fake AI engine for summaries and writing. The
> project owner chose 1A–4A and approved the document. The final pass ended with 198 cases executed
> correctly and 30 omissions exclusive to the AD, Keycloak and diagnostic profiles.
>
> **Amendment approved, implemented and verified on 2026-10-22:** section 28 defines the equivalence and the real
> check of the two example sets.
>
> **Amendment approved, implemented and verified on 2026-10-20.** Section 27 verifies the
> self-sufficient update and the versioned migrations. The project owner closed and approved
> decisions 1A–25A. Production was not touched.
>
> **Amendment approved, implemented and verified on 2026-10-20: validating the test SMTP send.** Section 26
> covers ephemeral recipient, real `From:`, unsaved data, persisted secret, language, limit and
> accuracy of the result. The project owner closed and approved 1A–14A; the focused tests and the 248
> frontend tests passed, as did boundaries and build.
>
> **Amendment approved, implemented and verified on 2026-10-19: SMTP transport matrix.** Section
> 25 verifies the three modes, the absence of degradation and the private limit of the unencrypted
> mode. The project owner closed and approved 1A–10A; Mailpit, regressions and a real relay were
> checked.
>
> **Approved and implemented on 2026-10-18.** Section 24 records the portable verification of CPU,
> NVIDIA detection, Auto fallback and transactional restore without touching production.
>
> **Implemented and verified on 2026-10-18.** Section 23 adds verification of the Dashboard,
> its calculations, permissions, CSV and responsive behavior.
>
> **Amendment implemented and verified on 2026-10-17: bilingual README and public screenshots.** Section
> 22 records the check of the reordered path and of its Spanish equivalence. Local data and the
> factory branding were restored, and the double view and the AI panel were visually verified. No
> public environment was touched.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** Section 21 adds the local check of the
> four priorities, order, filters, history and the non-blocking critical notice. No code is modified
> before implementing it. The project owner explicitly approved it.
>
> **Fourth public pass completed, 2026-10-13.** Section 20 records the complete path
> from `main`: empty installation, real local model, mail, accounts, ticket, AI, example data,
> AD, Keycloak and return to local. The inspection ended with no false errors or open findings and
> all disposable resources were removed; production remained untouched.
>
> **Third resumption approved, 2026-10-13.** The path completed installation, AI, mail,
> examples, AD and Keycloak, but the final review found false `record not found` errors in
> searches whose absence is normal. Section 19 keeps the evidence and requires a fourth full pass
> after the correction. The project owner chose and approved 1A–4A; the correction has already passed
> against real PostgreSQL and remains to be integrated into `main`.
>
> **Second resumption approved on 2026-10-13.** Block 8 stopped before `/setup` because `migrate` received `connection refused` after a premature socket healthcheck. Section 18 records timings, cleanup and a full repetition from another corrected `main`. The project owner approved the correction and the plan; the repetition remains to be published and run.
>
>
> **Finding from block 8 and proposed resumption, 2026-10-13.** The first execution stopped on
> getting a "Motivo" in Spanish on an English deployment. Section 17 keeps the evidence,
> refers the correction to the AI and tickets documents and requires repeating the full path from
> `main`. The project owner chose 1A–12A, with 10B, and explicitly approved this amendment on 2026-10-13.
>
> **Block 8 proposal, 2026-10-12: public path from scratch.** Section 16 validates the
> experience of an external person from a clone of `main`, using exclusively the README, an
> empty installation, the real local model, the optional seeders and the three entry methods. The
> project owner closed 1A–15B and explicitly approved the proposal on 2026-10-12. Production, domain
> and the `v1.0.0` tag remain out.
>
> **Amendment implemented and verified on 2026-10-08 (block 6).** Section 15 adds the reproducible path of the
> full visual review. What exists continues as-built and the proposal requires approval before
> modifying frontend or tests; the project owner explicitly approved it on 2026-10-08.
>
> **Amendment implemented and verified on 2026-10-08 (block 5).** Section 14 contains the isolated manual path
> to validate the three real models, `/setup` and `/settings`, without touching development. The project owner
> closed decisions 1A–14A. What exists continues as-built; no models are downloaded and no code is changed
> until explicit approval is received, granted by the project owner on 2026-10-08.
>
> **Amended on 2026-10-06**, explicitly approved by the project owner, implemented and verified
> in the same work. On repeating the isolated suite it was found that Docker could not
> mount `backend_tmp:/app/tmp` inside the `./backend:/app:ro` bind: it tries to create the mount point
> under a read-only parent and the backend did not manage to start. The `tests.yml` backend now runs
> `go run .` once, keeps `/app:ro` and removes that volume. The canonical execution without
> optional services ended with 177 cases approved and 39 omitted, and the cleanup removed all
> resources of the instance.
>
> **Amended on 2026-10-04**, explicitly approved by the project owner and implemented in the same
> work. It corrects the noise of a new database —the backend queries tables before applying the
> schema—, makes the first start
> begin in English and its selector translate the whole interface instantly, and preloads in development
> the editable values of Mailpit. The three decisions were confirmed by the project owner before
> writing this proposal. The review was closed and the approval "I approve it" enables its
> implementation.
>
> **Proposal of 2026-10-03, explicitly approved by the project owner.** It gathers the project owner's
> decisions from the two rounds of questions. The first chose local first, schema with a Docker
> command, local tests and optional Keycloak on localhost. The second chose separating the public
> and internal addresses of Keycloak, isolating the tests and preserving the configuration when
> loading examples. **Previous correction by the project owner:** the empty installation is the main
> path; the agent had suggested starting with examples. This proposal keeps that correction.
> Approval received after the review: "I approve it". Implemented and verified;
> amendment §5.1.b also received explicit approval.

## 1. Problem and expected result

Anyone who downloads the project must be able to follow the README without the domain, nginx, certificates or prior data from this machine. First they bring up an empty installation, apply the schema with one command, complete the wizard and enter with the factory account. They can create their own accounts and tickets, load examples if they want and try AD or Keycloak as optional services.

The automated tests run against a separate installation. They do not change the accounts, tickets, configuration, mail or attachments of the installation that person is using.

This proposal covers **development and local evaluation**. It does not enable deploying production.

## 2. Findings that motivate the proposal

| Finding | Current evidence | Proposed change |
| --- | --- | --- |
| Applying only the schema requires copying SQL and calling `psql` manually | README and `dev.yml` | Single-use service for the schema |
| Playwright points to the project domain | `dev.yml`, `playwright.config.ts` and preparation | Its own Compose and addresses on its network |
| Angular rejects the name `frontend` | `frontend/angular.json` and check of 2026-10-02 | Explicitly allow that development host |
| The examples replace the user's settings | Last blocks of `v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql` | Separate example data and configuration |
| Keycloak advertises the original domain | `keycloak.yml` and imported realm | Local public address and independent internal connection |
| The OIDC client does not read or check `issuer` during discovery | `backend/modules/auth/services/keycloak.go` | Check the public issuer before using the endpoints |

The limits of the production script already reported remain out of scope. `prod-build.sh` is not run, the domain is not opened and the version tag is not created.

## 3. Main path: without examples

Current commands, from the repository root, the same in PowerShell, CMD and bash:

```text
docker compose -f dev.yml up -d --build
docker compose -f dev.yml run --rm migrate
```

`migrate` is the approved name for the service that only applies the current schema:

- It has its own profile and does not start with `up -d`.
- It waits for a healthy PostgreSQL, mounts the migration read-only and uses the PostgreSQL client
  inside the container. It does not need Docker, Go, Node or `psql` on the machine.
- It uses `ON_ERROR_STOP=1`, keeps the SQL transaction and idempotence and fails with a non-zero
  code if the migration fails. The success message only appears after completing it.
- It does not load examples, does not delete data, does not configure SMTP and does not seal a new database.
- Repeating it does not reopen a sealed installation or alter its settings.
- While 1.0.0 remains open it applies `v1.0.0.sql`; it does not introduce a general migration system.

When opening `http://127.0.0.1:11001`, a new database shows `/setup`. The README states the five steps, the local method, the Mailpit values (`mail`, `1025`, without TLS or credentials) and the AI, with its tests.
After finishing, one enters with `admin` / `admin`, creates accounts and opens their links at
`http://127.0.0.1:11004`. The production password remains `ADMIN_PASSWORD` from the environment.

The wizard flow is not changed and no access that bypasses it is added.

## 4. Optional examples that preserve the configuration

The command remains `docker compose -f dev.yml run --rm seed-en|seed-es`.

On a **sealed** installation, it does not modify `installation_settings`, `directory_settings`,
`keycloak_settings` or `ticket_settings`: it preserves name, language, branding, region, address,
SMTP, entry method, directories, AI, prefix and assignment. It preserves the seal.

On an **unsealed** database, it completes the local preparation of the examples with demo values:
local entry, address `http://127.0.0.1:11001`, Mailpit SMTP and settings needed for the example
accounts and tickets; it leaves the demo sealed. The README warns that, in this case, the seeder
replaces the wizard path and may replace its still incomplete steps.

The destructive semantics of the example data are maintained: tickets, categories,
tags, counters, `@demo.com` accounts and development attachment files are recreated; branding is preserved.
Preserving configuration does **not** mean preserving that data. The notice must appear before the
command, both in the README and in its output.

The eleven accounts keep their demo credentials. If the preserved installation uses AD or
Keycloak, the local example accounts do not enter until returning to Local with the factory account.
The method is not changed to make them enter automatically.

## 5. Keycloak: public identity and internal connection

### 5.1 Configuration

An **optional** internal address is added to the Keycloak block, both in Settings and in
the wizard's entry step. It lives in `keycloak_settings`, with an idempotent migration inside
the open version, and travels in the settings DTOs and in `config.OIDC`.

Approved contract name: `internalIssuer` in JSON, `internal_issuer` in PostgreSQL. It is a complete URL of the same realm, not another issuer nor a secret. Empty means using the public
addresses, preserving installations that already work. No environment variables are added to
replace a configuration that lives in the database.

| Data | Local in the machine's browser |
| --- | --- |
| Application and public address | `http://127.0.0.1:11001` |
| Public issuer | `http://127.0.0.1:11006/sso/realms/catalina-support` |
| Internal address of the same realm | `http://keycloak:8080/sso/realms/catalina-support` |
| Return | `http://127.0.0.1:11001/api/auth/keycloak/callback` |
| Test client and secret | Those of the versioned realm, documented as local credentials |

Keycloak remains in its own compose and on the environment network. The browser reaches the
11006 port directly; there is no need to add `/sso` to the Angular proxy or to install nginx.
The realm allows the exact local return. No redirect wildcards are added.
The address announced by the container must be configurable for the isolated test environment;
editing the realm or having a hosts file on the machine is not required.

### 5.1.b Approved amendment: configuring a method that is still empty

**Status of this amendment: explicitly approved by the project owner on 2026-10-03 ("I approve it").** Finding when verifying
Keycloak from an empty installation: Settings disables AD and Keycloak if their fields are
empty, but the fields are only shown when choosing that method. There is no way to fill them through the
interface. The wizard already allows that provisional choice.

Choosing AD or Keycloak is allowed in the Settings form even if its data is
empty. Choosing only shows the fields; **the working method does not change until saving**.
The backend keeps its validation: saving an incomplete method returns 422. The factory account
is preserved. No screens, endpoints or permissions are added. It was verified that an empty local
installation allows filling in and testing both methods without loading examples and that an incomplete save
does not change the persisted method. The approval enables the selector change and its regression test.

### 5.2 Connection and validation rules

- Discovery is obtained from `internalIssuer`, if configured; otherwise,
  from the public issuer.
- The advertised `issuer` must match the configured public issuer. Missing or different:
  the test and the access fail with a translated key, before redeeming credentials.
- The authorization URL sent to the browser remains public.
- Only for backend requests is the public prefix of the realm replaced by the internal one:
  discovery, code exchange and UserInfo reading. The rest of the path and query is preserved.
- The substitution must compare parsed URLs and path boundaries, not replace arbitrary text.
  The endpoints must belong to the same origin and expected realm; a similar host or
  prefix is not transformed nor is a redirect to another origin followed with secrets or tokens.
- HTTP and HTTPS, ports and realm paths are allowed; embedded credentials,
  fragments and queries are rejected in both configuration addresses. TLS validation is not disabled.
- Emptying the internal address restores the public connection. Changing the settings invalidates the
  remembered discovery, as corresponds to a configuration read on every attempt.
- The confidential client, the signed `state` and the session in the return fragment are maintained.
  No PKCE or new ID token validation is added in this work.

Both fields are explained with simple text and examples, in Spanish and English. The internal
URL is not shown to the user who enters through Keycloak; the administrator is the one who configures it.

## 6. Tests in an isolated installation

`tests.yml` is implemented at the root, with a different project from the development and production ones:

- Its own containers, networks and volumes, without global names that match the current ones.
- Source mounted only to build or run; attachments, logs, dependencies and database in
  their own volumes. It does not mount `_files/`, `_logs/` or the local installation volumes.
- Frontend, backend, PostgreSQL and Mailpit without ports published on the machine.
- LDAP and Keycloak optional through a directory test profile, inside this network.
  They do not reuse the services or realm of the local installation.
- It does not enter the local engine network or download models. It uses `ai-test`, a fake compatible server
  inside its own network, to complete and test the mandatory AI without real providers.
- The browser uses `http://frontend.localhost:11001`, explicitly allowed in Angular. Mail links
  use that same base in the exclusive test configuration.
- The isolated Keycloak advertises `http://keycloak:8080/sso`; client, issuer, internal URL and return
  match the addresses accessible to the container's browser.
- Chromium resolves `frontend.localhost` to the `frontend` service only inside the isolated test
  browser; the localhost suffix provides a secure context, to exercise the real clipboard as on localhost. The API is not simulated nor is the
  application's security changed.
- All Playwright utilities share the same origin; no divergent fallbacks
  remain between configuration, preparation and cleanup.

The README shows a Docker command path: prepare a new test instance,
start its services, run Playwright and remove **only** its resources. Dependencies are
installed in containers; the runner failure keeps its exit code. The default command
does not use an external address nor turn `BASE_URL` into an implicit way of touching the local
installation. An external destination is an explicit advanced option and is not part of this path.

The preparation starts from the schema and configures through the API only the isolated installation. It includes a
browser case of the five steps from empty **before sealing it**, followed by the usual suite.
The tests create their accounts and fixtures, without requiring the user's optional seeder. The endpoints
require the same protection as in the product, without doors exclusive to tests.

The reports, traces and screenshots are kept in a results folder identified as test
results; removing the disposable database does not delete those evidences. Simultaneous execution of
several suites is not promised in this scope. A new pass begins with a new database and test files,
also after an interrupted pass; cleanup never searches for resources by a broad prefix.

## 7. Documents and files affected

This document coordinates the change; the rules of each area remain in its current document.
The specific amendments are recorded in the affected documents and describe the implementation.

| Updated document | What it decides or reflects | Affected files |
| --- | --- | --- |
| `docs/ambientes.md` | Schema with one command, seeder preserving configuration and disposable suite | `dev.yml`, `tests.yml`, `scripts/dev-seed-en.sh` or `scripts/dev-seed-es.sh`, migrations and `tests/e2e/**` |
| `docs/arquitectura.md` | Isolated network and volumes, allowed hosts and OIDC separation | Compose, `frontend/angular.json`, backend wiring |
| `docs/modules/auth.md` | Discovery, public issuer and internal connection of Keycloak | `backend/modules/auth/**`, `backend/shared/config/**`, `keycloak.yml`, `config/keycloak/**` |
| `docs/modules/settings.md` | New internal field, reading, writing and connection test | `backend/modules/settings/**`, migration of `keycloak_settings`, Settings screen |
| `docs/primer-arranque.md` | Optional internal field in the entry step and coverage from scratch | Setup DTO and services, `setup-page.*`, Playwright case |
| `docs/interfaz-y-experiencia.md` | Keycloak field and helps on both screens, texts in both languages | Settings, wizard and i18n dictionaries |
| `README.md` | Empty path, optional examples and reproducible local tests | Instructions and command examples |
| `docs/README.md` and `AGENTS.md` | as-built status and documentary correspondence | Index and coverage table |

The affected documents were updated in this work, with the amendment dated in their header.

## 8. Acceptance criteria

1. From a new database: `up` and `run --rm migrate` leave zero accounts and zero tickets; the browser
   shows the wizard and on finishing one can enter with `admin` / `admin`.
2. The schema command can be repeated: it preserves data, settings and seal. Invalid SQL makes
   the command fail and no success is announced.
3. Local sign-up: the mail reaches Mailpit, its link opens this installation and allows setting
   a password and signing in with the new account.
4. Seeder on a sealed installation: compares the five configuration tables before and after
   and the branding; they are preserved while the expected examples and their attachments are recreated.
5. Seeder on a new database or incomplete wizard: leaves a usable and sealed local demo, without
   external domain; repeating it does not change the configuration of the already sealed demo again.
6. Local Keycloak: optional container, connection test, entry with a person from the realm,
   exact return and deleted fragment. Without the container, Local and the factory account keep working.
7. OIDC: incorrect issuer, wrong realm, endpoints outside the prefix and external redirects
   are rejected; it is demonstrated that the backend uses the internal address and the browser the public one.
8. Isolated suite: wizard from empty, local accounts and mail on PC and mobile; LDAP and Keycloak
   when the profile is requested. Missing optional services are indicated as omitted cases.
9. A failed or interrupted pass does not change the user's installation and does not leave test
   data that contaminates the next one. Reports are preserved and failures do not become success.
10. Commands and mounts are tested on Linux. If no Windows or macOS machines are available,
    that limitation is stated and those systems are not claimed to have been verified.

## 9. Out of scope

Production and its script, publishing the domain, version tag, changes to roles or the ticket
flow, migrations of future versions, replacement of the AI engine, PKCE, Keycloak service account
management and concurrent execution of suites. Nor is the factory password changed
nor is installing Go or Node on the machine required.

## 10. Review record

The project owner answered **1A, 2A, 3A** to the document review:

1. **Seeder:** completes and seals the demo when the database is new or the wizard is incomplete.
   It was accepted in order to test immediately, with the notice that it replaces the incomplete steps.
   On a sealed installation it preserves the configuration, as section 4 fixes.
2. **Keycloak:** the internal address is optional and visible both in Settings and in the
   wizard. Configuring it without editing files was accepted, as was rejecting discovery if
   the `issuer` is missing or does not match, as section 5 fixes.
3. **Suite:** removes exclusively the data and resources of its disposable instance, preserves
   the reports and allows including LDAP and Keycloak through an optional profile. It was chosen that
   each pass begins clean, as section 6 fixes.

No design decisions remain open. After the review, the project owner explicitly approved
the document with "I approve it". The implementation and its verification were completed, and the document is closed as `as-built`.

## 11. Verification of the implementation

Checks made on isolated Linux copies, without modifying development or production data:

- Schema from scratch: 0 accounts, 0 tickets and null seal. Repeating the migration preserves the
  11 accounts, the 25 tickets and the hash of the five configuration tables. Invalid SQL
  returns code 3 and does not print the success message.
- Wizard, Mailpit SMTP, admin/admin, local sign-up, password link and real entry
  from localhost verified in the browser. The sealed wizard's API returns 409.
- Seeder on a sealed installation: the five settings preserve exactly the same hash
  before and after; 11 accounts, 25 tickets and 5 attachments are recreated. On a new database and null seal
  it leaves the local method, localhost URL and seal set.
- Real Keycloak: public authorization on localhost, backend internal traffic, test and
  save from Settings, access by a person from the realm and fragment removed on return.
- Empty selector, visible fields and 422 rejection without changing the local method verified on
  PC and mobile. The 220 frontend unit tests pass. Go: `go test ./...` and `go vet ./...` pass.
- Suite with directories: 214 cases, 193 approved, 20 omitted and one interrupted by an
  Angular reload during a formatting adjustment. Its report is kept, with exit 1.
  On a new and stable installation, the interrupted case passed on PC and mobile: 2 approved
  and 24 directory ones correctly omitted when the services were not started. Between both
  checks there remain 194 distinct cases verified and 20 foreseen omissions in the full suite.
  The failed output of the first pass is neither hidden nor turned into success.
- LDAP present: the same detection endpoint returns 200; without the service the cases are omitted.
- New passes begin with `down -v` exclusive to `tests.yml`; removing the
  containers and volumes preserves the reports and screenshots. The sources receive no writes
  from the runner and no volumes, attachments or mail from the user's environment are used.

These checks have not been run on Windows or macOS. Production remains out of scope.

Finding during verification without optional services: `directorio.spec.ts` declared
`respondeElDirectorio`, but did not use it before running its cases. The documentation
promised to omit them when LDAP was absent. Detection is applied in `beforeAll`,
in accordance with the behavior already approved in sections 6 and 8; the product does not change.

## 12. Clean start and initial wizard values

### 12.1 Findings

1. On a new database, `backend` waits for PostgreSQL to be healthy, but not for the
   schema to exist. That is why it manages to query `installation_settings` and `ai_insights`, writes `42P01`
   errors and remains listening even though `/setup` cannot work yet.
2. The `/setup` form initializes the language as `es` and the new `installation_settings`
   row is also born with `language = 'es'`. Choosing another language only changes the
   value that will be saved: it does not change the wizard texts or the other application texts.
3. Mailpit is part of `dev.yml`, but an empty installation receives empty SMTP fields. The
   README requires manually copying the local values that the environment already knows.

### 12.2 Development start

`docker compose -f dev.yml up -d --build` automatically applies `v1.0.0.sql` before starting the
backend. The `migrate` service is no longer hidden behind a profile; it waits for the database to be healthy,
finishes successfully and only then enables the start of `backend` through
`service_completed_successfully`.

The migration remains transactional and idempotent. It can be repeated with
`docker compose -f dev.yml run --rm migrate`, but it is no longer a mandatory step of the normal
path. If it fails, `backend` does not start and the error remains in `migrate`; a healthy backend
that actually lacks tables does not appear. This applies only to `dev.yml`: it does not change the deployment or the
production migrations. `seed` remains separate, explicit, optional and with the same destructive
semantics for the example data.

### 12.3 Language of the first start

A new installation is born with English as the installation language. On entering `/setup`
for the first time, that whole screen is presented in English, even if the browser prefers Spanish or keeps
a previous choice from another installation.

The **The installation language** selector still offers English and Español. Changing it updates
immediately all the application texts, including the five steps, helps, buttons,
notices and summary of `/setup`; it also updates `html[lang]` and remembers the choice in the browser.
On saving the first step, the same value is persisted as the installation's global language. If
one returns to a started wizard, the language already
saved in that installation prevails.

The change of the initial value of `installation_settings.language` from `es` to `en` affects new databases.
It does not automatically modify the global language of existing installations.

### 12.4 Editable Mailpit values in development

While the installation is not sealed and there is no saved mail, the `/api/setup` API can
return initial values coming from the environment, exclusive to the first-start form.
`dev.yml` declares:

| Field | Development initial value |
| --- | --- |
| Server | `mail` |
| Port | `1025` |
| TLS | no |
| User and password | empty |
| Sender name | `Catalina Support` |
| Sender email | `no-responder@catalina-support.local` |

They are editable suggestions: the user can change any field before testing or saving. They are only
persisted on saving step 4. They are not a second SMTP configuration mechanism nor a fallback
for sending mail after installing: the `mail` module continues to read exclusively the
saved configuration. In production these variables are not declared and the form remains
empty. If the wizard already has saved mail, the saved values are always returned and not the
initial ones.

### 12.5 Documents and code of the implementation

| Document | Change on closing as `as-built` |
| --- | --- |
| `docs/prueba-local.md` | A single start command, optional seeders and initial local values |
| `docs/ambientes.md` and `docs/arquitectura.md` | Real order of services and migration failure |
| `docs/primer-arranque.md` | Initial English, immediate translation and editable Mailpit values |
| `docs/interfaz-y-experiencia.md` | Behavior of the language selector during `/setup` |
| `README.md` | Remove the second mandatory command and explain that Mailpit already appears preloaded |
| `docs/README.md` and `AGENTS.md` | Status and count of the affected amendments |

The implemented code is limited to `dev.yml`, the backend configuration load and the first-start
service, the initial value of the migration, `setup-page.*` and its tests. It does not change the
seeders, the access methods, production mail, the flow after the wizard or the
production deployment.

### 12.6 Acceptance criteria

1. With a new volume, `docker compose -f dev.yml up -d --build` applies the schema and then
   starts `backend`; its logs contain no queries to nonexistent tables.
2. If the migration fails, `backend` does not start. Repeating it on a prepared database preserves data,
   settings and seal. No examples are loaded.
3. A new database opens `/setup` in English and shows English selected. Choosing Español translates
   the whole screen at once, changes `html[lang]` and is preserved when advancing, reloading and returning.
4. An existing installation or account preserves its language. Resuming a started wizard uses the
   language that was already saved.
5. In development, step 4 shows the six Mailpit values from the table; all can be edited,
   tested and saved. An edit reappears on reload.
6. In production, and in any environment without initial values, mail starts empty. The initial
   values never replace a saved configuration nor are they used directly to send.
7. The README is checked from a new database. The backend, frontend and `/setup` path tests on PC and
   mobile pass inside containers.

### 12.7 Decisions recorded before writing

The project owner confirmed on 2026-10-04:

1. The schema is applied automatically when starting development and `backend` waits for its completion.
2. `/setup` begins entirely in English; changing the selector translates the whole application
   immediately and remembers the choice.
3. Development preloads all the local Mailpit values. They are default values and the user
   can modify them.

### 12.8 Review record

1. **Migration failure:** the project owner confirmed that `backend` must remain stopped until
   `migrate` finishes correctly. `database`, `frontend` and Mailpit may stay up
   to diagnose the failure and repeat the migration; the backend is not presented as healthy.
2. **Language on resuming:** the project owner confirmed that, after saving the first step, the
   language saved in that installation prevails even if the browser remembers another. `/setup` applies it to the whole
   interface as soon as it loads the state.
3. **Mail on resuming:** the project owner confirmed that the Mailpit initial values only fill
   still empty mail. After saving step 4, the values persisted by
   the user always prevail and the development initial values are not restored again.

The review leaves no design decisions open. The project owner explicitly approved the amendment
with "I approve it". Implemented and verified, this section and the document remain `as-built`.

### 12.9 Verification of the implementation

Verified on 2026-10-04 on Linux:

- An isolated copy of `dev.yml`, with a new volume, ran `migrate`, waited for its correct output and
  then started the backend. Its logs contain no `42P01`: it read both tables and remained listening.
- `GET /api/setup` returned `language: "en"` and the approved Mailpit values. The database had zero
  users and zero tickets; `installed_at` was still null. Repeating `migrate` finished correctly,
  preserved that data and left the column's default value at `'en'`.
- `go test ./...` passed in all packages; `go vet ./...` and `gofmt -d` reported no problems.
- The frontend passed its 220 tests in 21 files.
- The disposable Playwright installation went through `/setup` on PC and mobile before sealing. It checked
  English even if the browser requested Spanish, immediate translation when choosing Español and priority
  of the saved language on resuming. The suite without optional services ended with 175 cases green
  and 39 omitted.

## 13. Writable backend outside the code tree in the isolated suite

### 13.1 Reproduced finding

With a new instance, `docker compose -f tests.yml run --rm e2e` stops before Playwright:

```text
error mounting ... backend_tmp ... to /app/tmp:
create mountpoint for /app/tmp: read-only file system
```

`tests.yml` mounts the code as `./backend:/app:ro` and then tries to mount a child volume at
`/app/tmp`. In this OCI runtime the second mount needs to create its point inside `/app`, which is already
read-only. The suite does not reach the backend or the browser.

The volume exists only for Air: `.air.toml` compiles `./tmp/server`. The isolated suite does not need
hot reload because the code does not change during a pass.

### 13.2 Proposed solution

The `backend` service of `tests.yml`:

- keeps `./backend:/app:ro`;
- defines `command: ["go", "run", "."]`, so that it compiles and runs once without Air;
- removes `backend_tmp:/app/tmp` and the `backend_tmp` declaration from the volumes;
- keeps `backend_go_mod` and `backend_go_build`, which store dependencies and cache outside `/app`;
- keeps the logs, attachments and the other suite service volumes.

`go run` writes its temporary binary outside the code tree. The backend continues to wait for the
completed migration and the runner continues to wait for its health before installing and testing.

### 13.3 Scope

`tests.yml` and the description of the suite change in `docs/prueba-local.md`, `docs/ambientes.md`,
`docs/arquitectura.md`, `docs/README.md` and `AGENTS.md`.

Out are `dev.yml`, `.air.toml`, the images, production, the Go code, the tests and their data.
Development continues to use Air to reload; the isolated suite starts the backend only once.

### 13.4 Acceptance criteria

1. `docker compose -f tests.yml down -v` followed by the suite creates the backend without mount errors.
2. The `/app` bind remains read-only and there is no volume with destination `/app/tmp`.
3. The backend waits for `migrate`, compiles with `go run .`, connects to the database and responds healthy.
4. The preparation goes through the empty installation and Playwright begins its cases.
5. At the end, `down -v` removes all volumes of the instance, leaving no `backend_tmp`.

### 13.5 Recorded decision and review record

The project owner confirmed on 2026-10-06 to propose `go run .`, keep the code read-only and
remove `backend_tmp`, without changing development or production.

In the review of the same day they confirmed that:

1. the suite does not need reload with Air: it compiles and starts only once;
2. `backend_go_mod` and `backend_go_build` are the only Go caches that are kept;
3. `down -v` must not leave any additional resource.

The review is closed with no open decisions. The project owner explicitly approved the three
documents with "I approve" on 2026-10-06; that approval enabled modifying `tests.yml`.

### 13.6 Verification of the implementation

Verified on 2026-10-06 on Linux, directly with the final `tests.yml` and without overwrite
files:

- `docker compose -f tests.yml --profile directory down -v` left a clean instance and
  `docker compose -f tests.yml up -d --build database backend frontend mail` created and started the
  backend without the OCI error of `/app/tmp`;
- the resolved configuration keeps `/app` as read-only, uses `command: [go, run, .]` and does not
  declare any `backend_tmp` mount or volume;
- the migration finished before the backend, the backend connected to PostgreSQL and the runner reached its
  health check, prepared an empty installation and began Playwright on PC and mobile;
- `docker compose -f tests.yml run --rm e2e` ended with exit 0: 177 cases approved and 39 omitted
  due to the missing optional services and profiles, in 13.3 minutes;
- the final `down -v` removed the containers, the network and all volumes of the instance; a
  subsequent `docker compose -f tests.yml ps -a` was empty.

The amendment and the document remain `as-built`.

## 14. Isolated manual path of the real catalog

### 14.1 Preparation and evidence

The path starts from an empty installation and Compose resources with exclusive names. Before
downloading, available memory and disk are noted and the catalog metadata is compared with the
official source. The report includes the exact commands, Docker and image versions, timings,
memory, checksums and browser results; it does not include credentials or binaries.

### 14.2 Sequence

1. Create the disposable environment, apply the migration and bring up the empty local administrator.
2. Go through `/setup` on PC and mobile up to the fifth step.
3. Download, activate and test 1.5B; seal the installation and verify one generation in Spanish and
   another in English.
4. Open `/settings`, download and activate 3B and repeat three generations per language.
5. Terminate only `llama-server` with each active model and measure recovery, up to five minutes.
6. Download and verify 7B without activating it; check that the lack of RAM is explained in the interface.
7. Return to 1.5B, delete 3B and 7B from the application when appropriate and run the full cleanup of the
   disposable project.

### 14.3 Acceptance criteria

- The three files match the official metadata and their checksum; a discrepancy stops
  that model and opens an amendment.
- 1.5B and 3B remain active and healthy, generate three valid responses per entry and recover health in
  five minutes or less after terminating the child.
- Each response meets the target contract of JSON, fields, language and length.
- `/setup` and `/settings` communicate download, activation, health, resources and errors on PC and mobile.
- 7B is validated without forcing swap or trying to activate it with insufficient memory.
- Cleanup leaves no containers, networks, volumes, databases or models of the temporary environment.

### 14.4 Scope and decisions

This block validates what is already built. It does not test GPU, external services or production and does not keep
real models. Any defect that requires code returns first to the documentary cycle within
block 5. The project owner chose 1A–4A and 5A–9A; the review chose 10A–14A. No decisions remain
open.

### 14.5 Result

The path ended with the three official checksums, 1.5B and 3B healthy and 24 valid responses. Recovery
took approximately 4 s and 9 s respectively. The safe rejection of 7B and its toast were
checked on PC and mobile. Only the projects
`catalina-support-block5` and `catalina-support-ai-block5` were used; on closing, their model volume, database, caches, files and network were also removed.

## 15. Format and accessibility path

1. Bring up an empty installation and go through `/setup` in Spanish and English, light and dark, on PC,
   tablet and mobile.
2. Restart the disposable instance, apply the optional seeder and enter with the roles needed
   to reach session, shell, settings, mail, users and all ticket views.
3. Record each finding with priority, context, reproduction and an observable measure; complete
   the inventory before correcting.
4. Correct first in the shared component when the pattern repeats and run the tests of
   all its consumers.
5. Repeat both languages, the two factory themes and the three resolutions on each affected
   screen; run the common contrast tests for the eight themes.
6. Close only when no finding remains open, the tests pass and the cleanup leaves the
   temporary project empty.

The path applies WCAG 2.2 AA where applicable: keyboard and focus, semantic structure, accessible name and
state, contrast and touch targets. It does not keep baseline screenshots, does not add browsers,
does not alter flows or touch production.

### 15.1 Result

The `/setup` path completed 48 combinations. After applying the seeder, 150 authenticated views
of the four roles and 18 public views were reviewed per language. The only real finding
was the overflow of the `/users` role filter in English and mobile; the shared correction
passed PC, tablet and mobile, light and dark, with Administrator and Support.

220 unit tests passed, 18 affected cases with 4 foreseen omissions and the canonical suite with
180 cases approved and 38 foreseen omissions. No finding remained open. The instance
`catalina-support-block6` was removed with its containers, network and volumes; development and production
did not take part in the path.

## 16. Block 8 proposal: onboarding from the public repository

### 16.1 Objective and source of truth

The path represents a person arriving for the first time at the public repository. It starts from a
new clone of the default branch `main` and uses only the commands, addresses, accounts and
explanations of the README contained in that clone. It does not use this document, internal knowledge, the
`dev` tree or improvised commands to complete a step.

A missing, incorrect instruction or one that requires unpublished knowledge is a finding, even if the
team knows how to continue. It can be investigated after recording the exact point, but the
path is not counted as approved through that shortcut.

The actual execution is done on Linux or on Windows with Docker Desktop WSL2, with Docker and Docker Compose.
macOS is not simulated and is not presented as verified. The main address is
`http://127.0.0.1:11001`; nginx, TLS and the local domain are not part of the first-test path.

### 16.2 Clean preparation and disposable data

The current development data were declared disposable by the project owner (15B). Before the
clone, the development, AI, directory and Keycloak composes are stopped and their containers,
networks and volumes are removed, including the model volume. The development `_files` and `_logs` are also
cleaned. No backup copy of those resources is created.

The images built specifically for Catalina Support are removed, so that `--build` demonstrates a
new build. Shared base images, global Docker caches and resources outside the project are not purged.
The working repository and its history are not deleted nor used as the source of the public clone.

The clone lives in a temporary folder outside the working repository. At the end its
containers, networks, volumes, model, attachments, logs, own images and the temporary folder are removed. A
final inventory must demonstrate that no Catalina Support resources created by the path remain.
Production is not stopped, modified or queried.

### 16.3 Empty path with real mail and AI

The README is followed from the clone: requirements check, build and start. A new database
must lead to `/setup` without seeders. Its five steps are completed with local entry, region and
local address, Mailpit and the default local model.

The model is downloaded from the interface, activated and truly tested. The criterion does not compare an exact
phrase: it requires the engine to remain healthy, that Catalina Support obtains a valid summary and that
"Mejorar con IA" returns a reviewable draft. An empty response, a hidden error or a connection
that only passes health do not comply.

After sealing the installation one enters with `admin` / `admin`, creates a local account using the documented
flow, opens its real mail in Mailpit and sets its password. With that account a ticket is created;
Support or Development exercises the summary and the writing improvement according to their permissions.
The path checks that `/setup` can no longer be reopened.

### 16.4 Seeders and optional entry methods

After the empty path, exactly the optional command of the README is run. It must finish
correctly, preserve seal, mail, AI and the other installation settings, and leave eleven
example accounts, twenty-five tickets and five usable attachments. An entry of
User, Support and Development is checked with the published credentials.

AD and Keycloak are brought up afterwards, one by one, with their documented commands. From Settings
AD is tested and activated, a person from the directory enters and the creation or update of
their account on first access is confirmed. Then Keycloak is tested and activated and a person from the realm enters. The
factory account must be able to enter during both methods. At the end the local method is restored and
a local account is checked again.

### 16.5 Timings and evidence

The report separates:

- **active time**, dedicated to reading, writing, choosing options and correcting data;
- **automatic waits**, image build, start, model download and generation.

System, architecture, Docker and Compose versions, public commit tested, commands
copied from the README, result of each milestone and final disk consumption of the model are noted. No
tokens, new passwords, private bodies, headers or complete ticket content are recorded. The
screenshots or temporary logs are kept only while they serve to diagnose a finding and are
removed with the environment.

There is no minute limit that by itself turns the path into a failure: the timings serve to
detect poorly explained waits and compare future improvements. A process stopped without progress or without
a message is a finding.

### 16.6 Closing criteria and handling of findings

The block is closed only if:

1. a clone of `main` completes the path without instructions external to the README;
2. the installation begins empty and `/setup` finishes and remains sealed;
3. Mailpit completes the local sign-up and the real model executes summary and writing improvement;
4. the seeder preserves the configuration and leaves 11 accounts, 25 tickets and 5 attachments;
5. local, AD and Keycloak allow their entries and at the end local works again;
6. the timings remain separated and the final cleanup leaves no resources of the path;
7. no finding remains open.

A typo or documentary instruction is corrected in the document that contains it and is repeated from the
previous stable point. A finding of behavior, interface, API, data, permissions or flow is
reported and returns to the proposal and approval cycle before changing code. Unapproved fixes
are not grouped within this audit.

### 16.7 Decisions and exclusions

The project owner chose 1A–5A: public installation, new resources, empty before examples, real local
model and optional methods after local. They chose 6A–10A: `main`, README only, complete functional
path, separated timings and strict documentary handling of findings. In the review
they chose 11A–14A and 15B: port 11001, real generation without fixed text, counts and preservation after the
seeder, AD followed by Keycloak and current data disposable without backup. No decisions remain
open.

Out are production, nginx, TLS, domain, the `v1.0.0` tag, external providers, costs,
GPU, comparative performance, Windows, macOS and unapproved functional changes.

## 17. Resumption proposal after the language finding

The first public execution tested commit
`afeae10190a87b0df92941f53d356bfc013453f4` on Docker 29.5.2 and Compose 5.1.4. The clone took
0.95 s, the environment was built in 40.14 s and the AI manager in 4.14 s. `/setup` began in
English, Mailpit appeared as `mail:1025` and its test responded 200. The default model showed
license, checksum and requirements; it downloaded 1.04 GiB in 55.98 s, occupied 1.1 GB, remained healthy and completed
the wizard generation in 3.91 s.

The path sealed the installation, entered with the factory account, created through real Mailpit
mails a User account and a Support account and opened `CS-2026-0001` from the interface. The two
summaries were ready in 16.37 s. "Última acción" came out in English; "Motivo" came out in Spanish even
though the API contract said `language: "en"`. The repeated query of the detail also produced
GORM `record not found` logs for the nonexistent internal. In accordance with 16.6 it stopped before seeders,
AD and Keycloak and no code was modified.

The environment was cleaned completely: containers, networks, volumes, model, built images,
clone and temporary tools. The final inventory kept only previous resources, including
the stopped production images and containers, which were not touched.

After implementing and verifying `docs/modules/ai.md`, section 15, and `docs/modules/tickets.md`, section
14, one does not continue from an intermediate point. The correction is published first in `main` and the whole
block 8 is repeated from a new clone and new resources: empty installation, real local AI, sign-up through
Mailpit, summaries and writing, seeders, AD, Keycloak, return to local, timings and cleanup. Closure
still requires zero open findings.

## 18. Resumption proposal after the TCP availability finding

The second public execution tested commit `80e01a75c7bc3b284778bcd214850efb0690714b` with Docker 29.5.2 and Compose 5.1.4. The clone finished in 2.21 s. The first `up -d --build` took 63.87 s and failed before the backend: `database` appeared healthy, but `migrate` received `connection refused` when connecting by TCP to `database:11003`. AI was not started nor was `/setup` reached.

The failure is attributed to the socket healthcheck during the transition between PostgreSQL's temporary initialization server and the final server. It stopped in accordance with 16.6 and containers, networks, volumes, own images and clone were removed; production was neither queried nor modified.

After implementing and verifying the new sections of `docs/ambientes.md` and `docs/arquitectura.md`, the correction will be published first in `main`. Then the whole block 8 will be repeated again from new resources and clone; one does not continue from the failure. In addition to the existing criteria, the pass will check that a new database does not present the race, that development and tests migrate automatically, and that health distinguishes in an isolated check the pending schema from an inaccessible database.

Before the public repetition, the automated tests of the three health states pass and a
disposable execution of the manual production service. In the review the project owner chose 14A–18A:
publish and mount the migrations, allow the manual command to bring up the database, return stable
JSON, cover the three states with tests and explain the order also in the README. 6B, 7A, 8A and 9A are kept,
together with 10A–13A. No decisions remain open; the proposal remains to be explicitly approved
before modifying code or configuration.

## 19. Approved resumption after the logs finding

The third public execution tested commit
`2f7c6468555fdf23efec08716d42755ccefa8b31` with Docker 29.5.2 and Compose 5.1.4. The clone took 1.30 s,
the environment 47.87 s and the AI manager 4.15 s. Development migrated automatically without table
errors; `/setup` began empty, in English and with Mailpit preloaded. The 1.5B model downloaded
1,117,320,736 bytes in 49.25 s, remained healthy in 14.62 s and completed the wizard test in 3.04 s.

The installation was sealed, local sign-ups were completed through real mails, `CS-2026-0001` was created,
the writing improvement returned text and the two summaries were ready with
`language=en`. The seeder preserved configuration and AI and left 11 accounts, 25 tickets and 5 attachments.
User, Support and Development accounts entered; afterwards AD worked with automatic creation,
Keycloak through its real OIDC return, the factory account and the final return to the local method.

The final inspection found `record not found` with error level when looking for the normalized category
before creating it and when searching by email and external identity before creating AD
and Keycloak accounts. The functional results were correct, but they contradict the requirement of an
operation without noise. Closing was stopped and the environment was cleaned completely, including the 1.1 GB
model; production remained untouched.

After approving, implementing, testing and integrating the proposed sections of architecture, users and
tickets, the whole block 8 will be repeated from another public clone of `main`. In addition to the seven
criteria of 16.6, the log examination must show zero false errors for expected absences and the
tests must demonstrate that a real SQL error keeps its level. The project owner chose 1A–4A; no
decisions remain open.

## 20. Fourth public pass completed

The final repetition tested from a clean clone the public commit of `main`
`b0898922f6d584d9b3349f2de82adc6046e82315`, with Docker 29.5.2 and Compose 5.1.4. The exact command
of the README, `docker compose -f dev.yml up -d --build`, built and brought up a new database in
29.13 s. The automatic migration finished correctly and health responded with the database available.
`/setup` began without a seal, in English, with local entry and the editable Mailpit values.

The AI manager came up in 3.36 s. From the wizard the default model
`qwen2.5-1.5b-instruct` was downloaded: 1,117,320,736 bytes in 42.8 s, with a valid checksum. Activation took
5.06 s and the mandatory wizard generation, 2.49 s. The installation was sealed with English,
`Europe/Berlin`, local address, Mailpit and the healthy local engine.

The factory account entered and created through the normal flow a User account and a Support one. The
two invitation mails reached Mailpit, their links allowed setting the passwords and
both accounts signed in. `CS-2026-0001` was created; "Mejorar con IA" returned a draft in
12.22 s and the "Motivo" and "Última acción" summaries reached `listo` in English.

The optional seeder was run exactly as the README indicates and took 4.69 s. It preserved the seal,
the language, the region, the address, the mail and the AI; it left its 11 accounts, 25 tickets and 5 attachments.
The published accounts of User, Support and Development entered. The two accounts created before the
seeder were also preserved.

Afterwards the optional services were brought up. AD responded to its test, was activated and authenticated
a person from the directory, who was provisioned on entering. Keycloak imported the realm, responded to its
test and completed the real OIDC return with the token in the fragment; its person was also
provisioned. The factory account entered during both methods. Finally `local` was restored and a
local account entered again. The final count was 15 accounts —11 from the seeder, 2 local ones created in
the path and 2 provisioned by AD and Keycloak—, 25 tickets and 5 attachments.

The complete inspection of the backend logs found no `ERROR:`, `record not found`, `[error]` level,
panics or fatal errors. The interface was also opened with Chromium in a mobile viewport:
the sealed installation redirected `/setup` to the entry, kept `lang=en` and the local account reached
`/tickets`. No files inside the clone were edited.

At the end `keycloak.yml`, `active-directory.yml`, `ai.yml` and `dev.yml` were brought down with their
volumes. The database, the 1.1 GB model, caches, networks, containers, clone and temporary tools of the
path were removed. The stopped production containers were preserved without starting or
modifying them. With zero open findings, block 8 is closed.

## 21. Priority path

On a disposable installation a ticket is created without changing the selection and `medium` is checked.
Then cases of the four levels are created or modified, it is verified that main and internal read
the same value and that the permissions match the matrix. The inboxes are tested on PC and mobile with
multiple filter, reloadable URL and critical→low order with age as tiebreaker.

The path marks a ticket as critical and confirms a single email to all active Support, with the
minimum fields and the global language. Then it lowers and raises the priority again and checks that there is no
other email. Finally it interrupts SMTP, changes another ticket to critical and confirms priority and
history saved along with the visible notice that the email failed. The seeder must show the four
levels with a realistic distribution.

The automated path of 2026-10-14 covers on PC and mobile the priority correction, the
persistence of filters and order in the address, the history, the User's read without editing and
the single email in Mailpit. The SMTP failure is checked in the backend against real PostgreSQL: the
priority and its history remain saved and the response carries the notice that the toast presents.
Two consecutive applications of the seeder left the same distribution: 2 critical, 6 high, 11
medium and 6 low.

## 22. Verification of the new cover

After rewriting both READMEs it will be checked that a person can follow the main path
in English and find the complete equivalent in Spanish; the commands, ports, credentials and
warnings must match. All relative links and the public links to the
site, the demo and Ko-fi will be verified, without automating external accesses or payments.

For the public image the seeder will be consciously run on the local environment, accepting its
documented cleanup of tickets, categories, tags, `@demo.com` accounts and development
attachments. Before capturing, factory branding, theme with institutional color, double
view, `Summary`, `Last action` and absence of personal information will be confirmed. The evidence will not show the
factory account, secrets, private emails or infrastructure identifiers.

The validation did not touch production, `demo.catalina-support.com`, DNS, the site, Ko-fi or the external
restoration script. All relative links exist, `git diff --check` ended clean and
the development and AI Compose files are valid. `prod.yml` confirmed its expected protection: it cannot
be evaluated without the private file `config/env/prod.env`, a requirement that the README explains instead of
hiding.

The interface checked at 1920 × 1200 the main ticket and the complete internal one, and in the detail of the
same ticket showed real results in `Summary` and `Last action` next to `Improve with AI`. The
double view does not render that side panel; the mismatch with the proposal was recorded and
resolved with two complementary real screenshots, without composing or dressing up the UI.

## 23. Dashboard and reporting

A disposable installation will generate main and internal tickets at period and zone boundaries,
with varied responses, resolutions, reopenings, closings, priorities and assignees. The tests
will compare the aggregates with known SQL, the previous period, the configurable thresholds and the
pending values. They will go through Support, Development, combined Administrator, factory and rejected
User.

On PC and mobile they will check filters persisted in the URL, accessible charts, manual refresh,
progressive table, ticket opening and complete CSV with human/minutes durations and
neutralized cells. Production, domain, mail and the `v1.0.0` tag remain out.

The verification executed used local PostgreSQL and the seeder users. Support saw both types,
Development only internal ones, User received 403 and Administrator accessed Dashboard and Settings.
A reloadable filter, row opening, complete CSV, mobile view and restoration of the
thresholds were checked. The automated tests cover service, real repository, frontend HTTP contract,
boundaries and build. Production and the domain were not touched.

## 24. Local engine acceleration

The disposable installation verifies migration and initial Auto value, CPU execution with the normal
Compose, device attribution, greeting, restart and persistence. It simulates NVIDIA capacity to
test arguments, absence, Auto fallback, explicit rejection, concurrent locking and rollback.
The portable suite passed on CPU. The CUDA manifest and the overlay resolved correctly. The
real pass was completed on 2026-11-05 on Docker Desktop WSL2: the GPU was exposed with `--gpus all`, the
manager detected `NVIDIA RTX A2000 12GB`, the model remained active and healthy and generated the two
summaries of a real ticket. The full E2E suite ended with 200 approved, 30 omitted and 0 failures.

PC and mobile go through `/setup` and `/settings` in both languages. It is inspected that API, toast and logs do not
expose UUID, bus, memory or traces. At the end Auto is restored, the optional variant is brought down and
the shared model is kept only during the test environment. Production, domain and tag
remain out.

## 25. Verification of the SMTP modes

The unit and integration tests cover: correct STARTTLS and rejection when it is not advertised;
correct direct TLS; `plain` rejected without acceptance; acceptance cancelled when leaving the mode;
rejection if a single resolution is public or IPv6; admission of `10/8`, `172.16/12`, `192.168/16` and
`127/8`; and connection against the already validated address without a second resolution.

The migration is tested with `smtp_secure=true` and `false`, expecting `tls_direct` and
`starttls` respectively. The development environment separately verifies that Mailpit remains `plain`, without user or
password and with explicit local acceptance. No test authorizes that exception in production.

The step 4 path checks the three texts in Spanish and English, the warning, the
confirmation persisted while the mode remains active, its deletion on changing and the toasts. The SMTP test
continues without `MAIL FROM`, `RCPT TO` or `DATA`. A check with a fake relay reproduces
`AUTH LOGIN` without STARTTLS: STARTTLS must reject it and `plain` must only accept it with confirmation and
private destination. User, password and body do not appear in API, logs or screenshots.

The execution ended with the focused regressions of `mail` and `settings` approved, 247
Angular tests approved, boundaries and production build correct. The full backend suite
keeps a failure unrelated and dependent on data seeded in the reporting repository; the
packages affected by SMTP passed. The real manual test authenticated against the relay and confirmed in
the toast that no email was sent.

## 26. Verification of the real test email

The unit tests go through the complete SMTP transaction (`MAIL FROM`, `RCPT TO` and `DATA`) and
check that the sender and the recipient are exactly those written. They cover fixed text in
Spanish and English, secret saved when the field is empty, temporary secret without persistence,
rejection of the `From:`, rejection of the recipient and failure during `DATA`.

The public endpoint admits five attempts for an IP–recipient pair within fifteen minutes,
rejects the sixth with `429`, separates another IP or another recipient and admits again when the
window expires. The `409` after the seal and the fact that an invalid address does not start SMTP are also checked.
The Settings route requires administrative capability and does not inherit the public character of `/setup`.

On PC and mobile a single button, the non-persisted destination, the optional character and the toast
**El servidor aceptó el correo para {destinatario}** are verified. Mailpit allows confirming subject, body,
language and headers without leaving the environment. The manual check against the institutional relay must
use a controlled mailbox and record acceptance and reception as distinct facts. The logs are reviewed
to ensure that they only contain a partially hidden recipient and no secret or body.

## 27. Verification of updates and migrations

The implementation will be verified on disposable installations, without touching production, the domain or
the usual database. The minimum matrix covers:

1. New database: creates the backup only on detecting work, applies `v1.0.0.sql`, records checksum and
   date, and starts backend and frontend after the runner.
2. Existing database without history: idempotently reapplies `v1.0.0.sql`, preserves data and records
   the version without duplicates.
3. Restart without pending work: does not run SQL, does not create a new backup and starts normally.
4. Update with two pending versions: orders them by SemVer, creates a single backup and
   confirms each transaction in order.
5. SQL error: fully reverts only that version, does not record it and leaves the application stopped
   with a message that identifies the file without exposing secrets.
6. Invalid or impossible backup: starts no migration or application.
7. Two concurrent runners: the second waits for the advisory lock, reads the history again and does not
   repeat work.
8. Changed checksum: the current version is reapplied only with `Development=true`; a published or
   historical version blocks startup.
9. Future file, invalid name, duplicate version or database newer than the code: all are
   rejected before modifying the schema.
10. `*_dev.sql` file: the runner never discovers it; it continues to be applied only by the seeder.
11. `update.sh`: requires `main`, clean tree and fast-forward; a download or build failure leaves
   the previous installation serving and does not modify the active deployment.
12. Recovery: a compatible migration allows temporarily bringing up the previous version; no
   automatic restore is run. The manual procedure with a disposable copy is validated separately.

Fourteen-day retention, backup name with versions, complete logs
without DSN or credentials, SQL files without their own transactions, build of the Go command,
resolved configuration of the three Compose files and the existing suites are also checked. README, runbook and architecture
remain updated to `as-built`.

The cases of new database, existing adoption, second start, rollback, checksum,
strict names, `_dev` exclusion and rejection of `update.sh` outside `main` were verified. Mutual exclusion
is implemented with an advisory lock on an exclusive connection; the path with several
versions will be exercised when the project owner opens the next one. The backend suite keeps only
the known environmental failure of the report without seeded tickets.

## 28. Verification of the bilingual examples

The test runs `seed-en` and `seed-es` on disposable and independent databases. It confirms in both the
eleven accounts, 25 tickets, 123 history rows, 27 comments, seven internal tickets, five attachments and
the same graph of identifiers, states, priorities, dates, relations and permissions. It also verifies
names, subjects, comments, categories, tags, reasons, file names and content
of attachments in the chosen language.

A new database adopts the seeder's language and remains sealed. An already sealed database preserves all its
settings when rebuilding the examples, even if the seeder of the other language is run. Changing the
platform language without running a seeder does not modify existing data.

The browser path opens both variants with their corresponding interface and checks that no
human content of the opposite language appears in the six screens intended for the web. The automated
tests compare quantities and structure between languages, download the five attachments and validate
their bytes. The two Compose files are also resolved, the affected suites pass and it is confirmed that the
production runner ignores both `*_dev_*.sql` files.

The regression also checks that the two public scripts go through the same internal runner and that
the generic `seed` service is no longer part of the resolved configuration. The project owner closed the
review with 1A–5A and no decisions remain open.

The real verification ran both services on disposable databases. Each variant produced eleven
accounts, 25 tickets, 123 history rows, 27 comments, seven internal tickets and five attachments. A
new database adopted `en` with `seed-en` and `es` with `seed-es`; a sealed installation preserved its
settings when changing sets. The resolved Compose configuration contains only `seed-en` and
`seed-es`, and the runner tests confirm that it ignores both development files.

## 29. Regressions of the ephemeral path without seeders

> **Coordinated change approved, implemented and verified on 2026-10-23.** The project owner chose
> 3A and 4A and closed the review with 1A–5A. The focused regressions passed; the full execution
> discovered additional historical contracts that are corrected in the proposal of section 30.

Playwright's official preparer uses the current step 4 contract: it fills **Enviar prueba a**,
selects SMTP without encryption for Mailpit, expressly accepts that transport and presses **Enviar correo
de prueba**. It keeps no names or assumptions from the old connection test.

The suite adds regressions to:

1. open step 5 without a local administrator, keep the other modes and finish with a self-hosted server;
2. alternate Main and Internal with out-of-order responses without mixing their rows;
3. expose a single **Nueva cuenta** button on PC and mobile;
4. use two consecutive password links in the same tab, correcting the product only if
   the reproduction fails.

The final pass starts from new volumes, does not run seeders, goes through the official preparer and
destroys containers and volumes. It does not touch persistent development, demo or production.

The focused regressions are run on PC and mobile. Then the full frontend and E2E suites pass
on empty volumes; a correction is not considered closed if only the case
that discovered it passes.

## 30. Full cleanup of E2E contracts

> **Approved, implemented and verified on 2026-10-24.** The project owner chose 1A–4A. No decisions
> remain open and there are no known failures in the current suite.

### 30.1 Validity rule

An E2E test describes only the behavior the product has today. The cases that still expect
previous routes, permissions, names, columns, states or flows are rewritten. Assertions, omissions and known failures
are not kept as historical memory: that memory belongs to Git and to the decision record, not to the executable suite.

A test is not removed only because its contract has changed. If the functionality is still current,
the case is adapted to the current contract and keeps user coverage. A case is only removed when
the functionality checked no longer exists or the same guarantee is covered identically by
another path.

### 30.2 Data isolation

The suite starts from an empty installation and does not run `seed-en` or `seed-es`. Each case creates through
API or interface all the accounts, tickets, categories, tags and relations it needs. It uses no
mail, identifiers, quantities or texts from the demo. The generated data carry unique
identifiers so that the execution order does not change the result.

The helpers reflect the current model: the operational role and the administrative capability travel
separately; entering leads to the Dashboard when appropriate; and each path navigates explicitly to the
screen it wants to test. No helper silently translates an old contract.

### 30.3 Deterministic fake AI engine

The static nginx server is replaced by a small and deterministic fake service. It recognizes
the request type without depending on the order of the tests and returns:

1. compatible health for the connection tests;
2. valid and stable JSON for **Motivo** and **Última acción**;
3. valid professional text for **Mejorar con IA**;
4. controllable error responses when a case needs to check unavailability or rejection.

The fake does not call the Internet, does not download models and does not share mutable state between tests. The
paths exercise frontend, backend and real adapter up to that boundary; they do not intercept the product
API from Playwright to fabricate a success.

### 30.4 Closing criterion

The cleanup ends when the **228 current tests** pass on PC and mobile from new volumes.
The only admitted omissions are those that require an explicit profile of Active Directory,
Keycloak or diagnostics; each one must declare that dependency in the case itself. No
omissions for debt remain, no tests marked with known failures and no comments that justify an
obsolete expectation.

The verification includes `--list`, the Angular tests, boundaries, production build and the full
E2E suite on an empty installation. At the end containers and volumes are destroyed. Demo,
persistent development, production, domain and publishing remain out of scope.

The final verification ran the 228 declared cases on PC and mobile: 198 passed and the 30 omitted
belong only to the explicit profiles of Active Directory, Keycloak and diagnostics. The
255 Angular tests and the frontend boundaries passed. The suite did not run seeders and the service
`ai-test` deterministically handled the connection test, the summaries and the assisted writing.
The full Go suite also passed. Its reactivation-by-comments tests create their own accounts and
tickets inside a transaction, explicitly configure the identity double and
revert that data at the end; they do not depend on contracts that another path might have left.
