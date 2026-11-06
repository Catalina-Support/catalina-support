# settings

[Español](../es/modules/settings.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-05
>
> **NVIDIA hardware verified, 2026-11-05.** The Auto/CPU/NVIDIA selection was tested with an RTX
> A2000 12 GB on Docker Desktop WSL2: the administrator reported the effective mode `nvidia`, the model
> remained active and the ticket showed the generated summaries. Linux with NVIDIA Container Toolkit remains
> the other supported path; native Windows with Windows containers does not.
>
> **Amendment approved, implemented and verified on 2026-10-20: configurable site icon.** Section 25 adds a
> favicon of its own, independent of the two logos, with administration in `/setup` and Settings,
> immediate update and fallback to the effective light logo. The project owner closed the review with
> 1A–14A and corrected 4A: without a custom icon, the custom light logo is used or, in its absence, the
> factory one. The settings module, 253 frontend tests, boundaries and build passed; the migration and
> the public endpoint were verified in the local environment.
>
> **Amendment approved, implemented and verified on 2026-10-20.** It states that the configured LDAP
> filter is the baseline security condition and that the application adds the escaped search by
> `sAMAccountName` or `mail`; the `objectGUID` identifier is normalized as a canonical GUID. It does not add
> Keycloak administration credentials: their manual creation is out of scope (6B).
>
> **Amendment approved, implemented and verified on 2026-10-20: real test email.** Section 23 makes
> `/setup` and Settings request an ephemeral recipient and test the written data through a real
> send. The test still saves no configuration and is optional. The project owner closed and
> approved 1A–14A; focused backend, 248 frontend tests, boundaries and build passed.
>
> **Amendment approved, implemented and verified on 2026-10-19: explicit SMTP transport.** Section
> 22 replaces the encryption boolean with three modes and stores acceptance of the unencrypted SMTP
> risk separately. The project owner closed and approved 1A–10A; migration, API and screen were
> verified against Mailpit and the real internal relay.
>
> **Approved, implemented and verified on 2026-10-18.** Section 21 adds Auto/CPU/NVIDIA selection,
> effective state, device, visible fallback and restoration on failure.
>
> **Implemented and verified on 2026-10-18.** Section 20 adds the four reporting delay thresholds,
> manageable from Settings.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Section
> 19 separates the automatic loading of the local catalog from its manual refresh and keeps the data
> while it refreshes.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Section
> 18 extends step 1 with the institutional color and requires model confirmations to use the
> shared, themed dialog.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Section 17 adds parallel contracts for
> local deactivation, attributes the greeting to the model and simplifies logo loading without exposing the
> file name.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Section 16 defines the parallel
> `/setup` and `/settings` contracts for progress, cancellation, activation, testing and deletion of
> models, and allows managing both logos during first start without lowering the seal lock.
>
> **Amendment implemented and verified on 2026-10-13: absence without a false error.** The five single
> rows keep `settings.notFound` if their migration is missing, but the query no longer logs that
> expected absence as `record not found`. Real SQL errors remain visible. It is part of the 1A–4A scope
> approved for the third finding of block 8.
>
> **Amendment implemented and verified on 2026-10-10: source metadata.** Section 14 adds the
> license and corresponding source to the public brand. The URL belongs to the build and is not
> an editable setting of the installation.
>
> **Finding corrected and verified on 2026-10-08.** The local model administrator returned any
> activation failure as “not available”, so Settings cannot explain the required and available memory
> found in block 5. Section 13 proposes keeping that data through to
> the toast. Decisions 15A–22A are closed and were explicitly approved on 2026-10-08.
>
> **Amendment proposed on 2026-10-06.** What exists remains as-built. The language becomes the
> first setting and governs the whole installation; the AI card goes from optional URL/model to
> a mandatory configuration with a local engine, own server or provider. The proposal is
> detailed in section 12. The review was closed and the proposal was approved.
>
> **Amended on 2026-10-06**, with explicit approval from the project owner. The results of
> Settings actions were drawn in a notice at the top of a long screen. They are now
> presented in the shared toast of `docs/interfaz-y-experiencia.md`, so that a connection
> test or a save is seen from the card where it ran. Permanent notices remain in
> place. It was verified on PC and mobile; backend and persistence do not change.
>
> **Amended on 2026-10-05**, with explicit approval from the project owner. The region chosen in
> Settings is only distinguished by highlighting within the list and may fall outside the
> visible part, the same finding already corrected in `/setup`. The project owner confirmed reusing in
> `/settings` the read-only field with the exact IANA identifier, before the search box, with
> immediate update and the time and offset below. It was implemented and verified in unit tests and in Playwright, on PC and mobile. Persistence does not change.
>
> **Amended on 2026-10-04**, in accordance with the approved `docs/prueba-local.md`: the initial value of
> `installation_settings.language` is `en` for new databases. The first-start service accepts
> `SETUP_MAIL_*` suggestions, but only returns them if the installation is not sealed and has no
> SMTP saved; they never configure sending.
>
> **Amended on 2026-10-03**, in accordance with the approved `docs/prueba-local.md`: `keycloak_settings.internal_issuer` is added,
> exposed as `internalIssuer`, empty by default.
> Settings and first start show the optional internal address next to the public issuer,
> with ES/EN texts; saving and testing use both fields. Amendment §5.1.b of `prueba-local.md`
> explicitly approved: choosing AD/Keycloak allows completing empty fields; the current method
> changes only when valid data is saved, and the backend keeps 422 for incomplete methods.
>
> **Amended on 2026-09-30**: the module declares **`ProberDeCorreo`**, its fourth test interface
> —along with the directory, Keycloak and AI ones—, to **test outgoing email without sending any email**.
> It is implemented by the `mail` module and **reused by the first-start wizard API** for the test
> of step 4 (`docs/primer-arranque.md`, section 3.1). It is counted where the interfaces that
> this module declares live (section 4) and next to the connection tests (section 5.8).
>
> **Amended on 2026-09-30**: **the four configuration tables were reviewed against the real schema**
> (section 3), at the project owner's request, and **columns were missing**: the
> `installation_settings` table did not include `entry_method`, `time_zone`, `public_app_url`, `installed_at` or
> the eight outgoing email ones (`smtp_host`, `smtp_port`, `smtp_security_mode`, `smtp_insecure_ack`,
> `smtp_user`, `smtp_password`, `smtp_from_name` and `smtp_from_email`); `ai_url` and `ai_model` were
> there, in a joint line. The
> four tables now have **each column, its type, whether it allows null and its default value**,
> checked against `\d` of the database and against `backend/migrations/v1.0.0.sql`, which **match**. The
> review also uncovered two sentences that contradicted the schema —outgoing email was still counted as
> an environment variable (section 2 and section 5.7)—, a wrong section reference (section 5.11)
> and a couple of old notes that took the sender as an environment variable (decision 9 and the
> name amendment): all corrected or marked as superseded.
>
> **Amended on 2026-09-30**: **the AI engine is configured from the screen** (decision 16). Its
> address and model leave the environment and enter the closed list of what is configurable (section 2),
> with **its card** in Settings (section 5.13) and **its «Test connection» button**, which
> asks the engine's health check. The `ai` module **reads the configuration on every
> request**; the `AI_URL` and `AI_MODEL` variables **remain as a fallback**, like `PUBLIC_APP_URL`
> (decision 14). The three connection tests answer **`{"status":"ok"}`**, and the engine's **tests
> only the address that arrives**: empty is «there is nothing to test».
>
> **Two jobs were agreed with the project owner as separate work** (2026-09-30). The first,
> **reviewing the schema tables of section 3 and leaving them the same as the schema**, **is already done**
> (amendment above): they were checked against the database and against the migration, and the missing columns
> were documented. The second, **extracting a `SettingsService`**: **is already done** (2026-09-30). The
> calls to `/api/settings/**` live in `frontend/src/app/core/services/settings.service.ts`, and the
> screen only uses the service: **no `HttpClient` remains in the component**. **The brand and the
> color** —reading the public brand, uploading and restoring the logo and previewing the color— are served by
> `frontend/src/app/core/services/brand.service.ts`, which belongs to whoever owns the resource
> (`/api/settings/brand`), and Settings uses it **only** for that: the rest continues through `SettingsService`.
>
> **Amended on 2026-09-27**, at the project owner's request: in «Authentication method», **each method shows
> only its configuration** —«Application accounts» shows none, the directory appears with the
> organization and the realm with Keycloak—, instead of both at once (decision 10-bis).
>
> **Amended on 2026-09-26**, when the project owner **replaced the factory brand files**:
> the **favicon** and **two logos**, one for light themes and one for dark ones. That turns the
> factory brand into **two files per theme** —before it was a single one, used in all eight—, so
> section 5.1, 5.6 and decision 4 change. **The factory files are not configured from
> the screen**: they are replaced in `frontend/public/`, and that is done by whoever maintains the application
> (project owner decision, 2026-09-26).
>
> **Amended on 2026-09-25**, when the project owner asked that **the system version be visible**:
> it is a backend constant that travels in the public brand and is read in the sidebar and
> in the footer of the entry screen (section 5.9, decision 12). **He chose the entry footer, which was not
> my recommendation**: I proposed only the menu.
>
> **Amended on 2026-09-25**, when the project owner asked that **the organization directory and
> Keycloak be configured from the screen** instead of through environment variables, and that **the
> installation use one method at a time**: they enter the closed list of what is configurable (section
> 2), they debut **two single-row tables** —`directory_settings` and `keycloak_settings`— and their section
> (5.8). **Secrets are stored in the database and never leave through the API**, and the
> `LDAP_*` and `OIDC_*` variables **disappear from the environment** (decision 10). This amends
> `docs/usuarios-y-permisos.md` (section 5) and `docs/modules/auth.md` (decision 28), which described the
> three paths as coexisting.
>
> **Amended on 2026-09-25**, when the project owner asked that **the name of the institution, company or team be configurable**
> instead of the text «Catalina Support»: it enters the closed list of what is configurable (section 2),
> it debuts its own column —in `v1.0.0.sql`, along with the rest of the
> schema— and its section (5.7), which says where it is read and what **does not** change. **The
> project owner corrected two things in my proposal**: emails remain as they are and the sender
> is still the one from the environment variable (decision 9; **superseded on 2026-09-30**: the sender moved
> to `smtp_from_name`/`smtp_from_email` in the table, section 5.12).
>
> Written when the project owner asked that **the installation logo be replaceable** from the
> Settings screen: the logo, the language, the institutional color, the numbering prefix and
> the ticket assignment **are the installation configuration**, and their owner is this module, which
> until now had no document.
>
> **Implemented on 2026-09-24** (the brand and the institutional color, which is what was approved for
> this stretch): the two tables, the six endpoints, logo validation by content, the public
> endpoint with version and SVG isolation, the logo on the entry screen and the institutional
> color applied to the two factory themes.
>
> **And the shell and the Settings screen, too** (2026-09-24): the sidebar with the logo
> in its product area, and the screen with **the brand** —the two slots with their preview, upload and
> return to the factory one— and **the institutional color** with **a preview requested from the backend** before
> saving. The prefix, the assignment and the language will arrive with their consumers.
>
> **Approved by the project owner on 2026-09-24**, after answering the seven review questions: the
> logo limit (1 MB and 2000 pixels), what the installation language does (it does not govern the
> interface), that the institutional color only affects the two factory themes, the two optional logo
> slots, **the amendment to `docs/modules/tickets.md`** (`ticket_number_settings` becomes
> `ticket_settings`), that the logo does not go in emails and that its endpoint is public.

## 1. Scope of this document

It describes **what is configurable in a Catalina Support installation**, where that configuration lives,
how the other modules read it, and **the brand**: the installation logo, how it is uploaded, where it is
stored, which one is shown in each theme and how to return to the factory one.

**It does not repeat** what is already decided elsewhere:

| What | Where it is |
| --- | --- |
| The themes and the primary color | `docs/interfaz-y-experiencia.md`, section 6.2 |
| The numbering prefix and the ticket assignment | `docs/modules/tickets.md`, sections 2.2 and 3.3 |
| The email templates | `docs/modules/mail.md` (they belong to the `mail` module, not here) |
| Who can change the configuration | `docs/usuarios-y-permisos.md`, section 3: **only an Administrator** |
| Where the screen is | `docs/interfaz-y-experiencia.md`, section 5 |
| The three entry paths, and how each one signs in | `docs/modules/auth.md`, sections 5.1 to 5.4. **Here it only says where they are configured** |

## 2. What is configurable, and what is not

**The list is closed.** What is not here is not configured from the application: it belongs either to the server
(the environment variables), to each person, or it does not exist.

| What | Values | Factory default |
| --- | --- | --- |
| **The installation name** | text, **up to 60 characters**; empty returns the factory one | `Catalina Support` |
| **Installation language** | `es` or `en` | `en` |
| **Institutional primary color** | a color, in hexadecimal | the institutional blue |
| **The brand**: the logo | up to **two** files: one for light themes and one for dark themes, **both optional** | none: the two factory logos are used, one per theme |
| **The numbering prefix** | uppercase letters and digits, 2 to 8 characters | `CS` |
| **The assignment of new tickets** | `ninguna` or `por_turnos`, **for each ticket type** | main `por_turnos`; internal `ninguna` |
| **The new-ticket notification** | `a_nadie`, `a_todo_el_equipo` or `al_asignado`, **for each ticket type** | main `al_asignado`; internal `a_nadie` |
| **The entry method** | `local`, `ad` or `keycloak`. **One at a time** | `local` |
| **The organization directory** | server, port, TLS, service account **with its password**, where it searches, the filter and four attributes | empty: there is no AD path |
| **Keycloak** | realm issuer, client, **its secret** and the redirect address | empty: there is no Keycloak path |
| **The time zone** | an **IANA zone name** (`America/Guayaquil`, `Etc/GMT+5`…), from the list of zones with a search box. It decides **how dates are read** —interface and emails—; stored dates remain in UTC (decision 15) | `UTC` |
| **The public address** | `http` or `https`, host and optional **port**; **`localhost` is valid**. It is the base for email links and the Keycloak callback (decision 14) | empty, and then `PUBLIC_APP_URL` is used |
| **The AI engine** | local mode, own server or provider; model, URL and authentication where applicable. It is tested before activation | local with Qwen2.5 1.5B selected; the installation does not finish until a valid option is activated |
| **The outgoing email** | server, port, encryption, user, **its password**, the sender name and address. **It has no card in Settings**: it is requested in the first-start wizard (`docs/primer-arranque.md`, section 5) and lives in `installation_settings` | empty: no email configured |

- **The prefix and the assignment were already decided** in `docs/modules/tickets.md`; here it only says
  **where they live**, which was what was missing there.
- **Values travel in English and without spaces** (`por_turnos`, `al_asignado`) even though the interface
  shows them in Spanish: it is a contract value, not a text, and the texts are provided by the
  frontend dictionary (`docs/interfaz-y-experiencia.md`, section 8).
- **The installation language is global** and the installation uses it for the interface, emails and
  new summaries. Accounts do not store their own language.
- **The institutional color only changes the two factory themes** (Light and Dark), which are the
  customizable ones. The six fixed themes carry their own accent, approved on 2026-09-23
  (`docs/interfaz-y-experiencia.md`, section 6.2).
- **The directory and Keycloak are no longer environment variables** (project owner decision,
  2026-09-25): in a delivered installation, configuring an AD by asking someone to edit a file on
  the server and restart the container is not configuring. **Their secrets are stored in the database** —it is
  the only place where an administrator can change them without SSH access— and **they are shown only
  once: when they are written**. What the API returns is whether one is set, never the value: section 5.8
  establishes that and a test checks it.
- **The entry method is chosen, and the other two are turned off.** It can be selected to write its fields, but one that is not
  configured cannot be saved: that would leave the installation with no door for everyone except the factory account, and
  the backend rejects it with its key (section 7).

**What is NOT configured here**, so that this module does not end up as the drawer for everything:

- **The mail server is not configured from the Settings screen**: its data lives in
  `installation_settings` (`smtp_*`) and is requested in the first-start wizard
  (`docs/primer-arranque.md`, section 5). **The `SMTP_*` environment variables were removed**
  (project owner correction, 2026-09-30): what remains of the environment are `AI_URL`/`AI_MODEL` and
  `ADMIN_PASSWORD`, and nothing else.
- **The email templates**: they belong to the `mail` module, and their editor lives on the same screen.
- **The roles and their permissions**: the four roles are fixed (`docs/propósito-y-alcance.md`).
- **Each person's theme**: each person chooses it, and it is remembered in their browser; what the
  installation sets is **the primary color**, which only appears in the two customizable themes.
- **Ticket statuses, attachment limits and the password policy**: they are
  product rules, not configuration.

## 3. Where it lives: five single-row tables and two AI auxiliary tables

As was already decided for numbering, **no key-value «settings»**: such a table ends up
being the drawer where everything goes and it loses types and constraints. Each table has **a single
row**, with real columns.

**`installation_settings`** — the installation's own configuration. **The columns follow the schema
order** (`\d installation_settings`) and the default values are those of the migration:

| Column | Type | Null | Default | What it is |
| --- | --- | --- | --- | --- |
| `id` | `integer` | no | `1` | Primary key with `CHECK (id = 1)`: the whole table is **a single row** |
| `language` | `text` | no | `'en'` | The global language (`es` or `en`) of the interface, emails and new summaries; during `/setup` it also governs that interface |
| `primary_color` | `text` | no | `'#1d4ed8'` | The institutional color, in hexadecimal `#rrggbb`; it only affects the two factory themes |
| `logo_light` | `text` | yes | — (null) | **The file name** of the logo for light themes, or null if there is no custom logo |
| `logo_dark` | `text` | yes | — (null) | The one for the dark theme, or null |
| `site_icon` | `text` | yes | — (null) | The internal name of the custom tab icon; null falls back to the effective light logo |
| `updated_at` | `timestamptz` | no | `now()` | When the row was last changed |
| `updated_by_id` | `bigint` | yes | — (null) | Who changed it: `users(id)` with `ON DELETE SET NULL`; null if it was the seed |
| `installation_name` | `text` | no | `'Catalina Support'` | The name of the institution, company or team; it replaces «Catalina Support» where that used to be read. **It does not allow empty** and does not exceed 60 characters |
| `entry_method` | `text` | no | `'local'` | **The installation entry method**: `local`, `ad` or `keycloak`. **One at a time**, and it can be changed (decision 11) |
| `time_zone` | `text` | no | `'UTC'` | The time zone (IANA name) with which dates are read, in the interface and in emails. **Stored dates remain in UTC**, so changing it does not move any ticket (decision 15) |
| `public_app_url` | `text` | no | `''` | The public address —scheme, host and port—: base for email links and the Keycloak callback. Empty uses `PUBLIC_APP_URL` (decision 14) |
| `installed_at` | `timestamptz` | yes | — (null) | **The installation seal**: when the first-start wizard was completed; null is «not installed» (section 5.12) |
| `smtp_host` | `text` | no | `''` | The outgoing mail server; empty is «not configured» |
| `smtp_port` | `text` | no | `''` | The port of the outgoing mail server |
| `smtp_security_mode` | `text` | no | `'starttls'` | The transport: `starttls`, `tls_direct` or `plain` |
| `smtp_insecure_ack` | `boolean` | no | `false` | The acceptance in force of the `plain` mode; outside it, it is always false |
| `smtp_user` | `text` | no | `''` | The user the sending authenticates as |
| `smtp_password` | `text` | no | `''` | The sending password: **it never leaves through the API** |
| `smtp_from_name` | `text` | no | `''` | The sender name of the emails |
| `smtp_from_email` | `text` | no | `''` | The sender address of the emails |
| `ai_url` | `text` | no | `''` | A column inherited from the previous configuration; the current source is `ai_settings` |
| `ai_model` | `text` | no | `''` | A column inherited from the previous configuration; the current source is `ai_settings` |

- **The table stores the file name, not the file.** The logo is a file and lives on
  disk (section 5): putting it in the database would bloat every backup and every query.
- **Neither logo variant is mandatory**: only one can be uploaded, and the application
  makes do with it (section 5).

**`directory_settings`** — the organization directory, for the AD path:

| Column | Type | Null | Default | What it is |
| --- | --- | --- | --- | --- |
| `id` | `integer` | no | `1` | A single row, like the others |
| `host` | `text` | no | `''` | The directory server; **empty is «this installation has no directory»** |
| `port` | `text` | no | `'389'` | Its port |
| `use_tls` | `boolean` | no | `false` | Credentials travel through it: in production, yes |
| `bind_dn` | `text` | no | `''` | The service account with which people are searched |
| `bind_password` | `text` | no | `''` | The password of that account: **it never leaves through the API** |
| `search_base` | `text` | no | `''` | Where it searches |
| `user_filter` | `text` | no | `'(objectClass=person)'` | The baseline condition that limits the accounts admitted |
| `attr_username` | `text` | no | `'sAMAccountName'` | The username attribute |
| `attr_email` | `text` | no | `'mail'` | From which attribute the email comes |
| `attr_name` | `text` | no | `'givenName'` | From which one, the first name |
| `attr_last_name` | `text` | no | `'sn'` | From which one, the last names |
| `attr_id` | `text` | no | `'objectGUID'` | From which one, the identifier |
| `updated_at` | `timestamptz` | no | `now()` | When the row was changed |
| `updated_by_id` | `bigint` | yes | — (null) | Who changed it: `users(id)` with `ON DELETE SET NULL` |

**`keycloak_settings`** — the realm, for the Keycloak path:

| Column | Type | Null | Default | What it is |
| --- | --- | --- | --- | --- |
| `id` | `integer` | no | `1` | A single row |
| `internal_issuer` | `text` | no | `''` | Optional internal address of the same realm; empty uses the public issuer |
| `issuer` | `text` | no | `''` | The realm address, **the one the browser sees**; empty is «this path does not exist» |
| `client_id` | `text` | no | `''` | The realm's confidential client |
| `client_secret` | `text` | no | `''` | Its secret: **it never leaves through the API** |
| `redirect_uri` | `text` | no | `''` | Where Keycloak returns: the `auth` callback route |
| `updated_at` | `timestamptz` | no | `now()` | When the row was changed |
| `updated_by_id` | `bigint` | yes | — (null) | Who changed it: `users(id)` with `ON DELETE SET NULL` |

- **The factory values are set where possible**, so as not to force everything to be written: port
  `389`, the filter `(mail=%s)`, the attributes `mail`, `givenName`, `sn` and `objectGUID` and the
  empty directory. What cannot be guessed —the server, the search base, the realm— is set by
  whoever administers.
- **A half-configured path is rejected on save** (`settings.directory.incomplete`,
  `settings.keycloak.incomplete`): a directory without a search base searches for no one, and a realm without a
  callback cannot return.

**`ticket_settings`** — what configures ticket behavior:

| Column | Type | Null | Default | What it is |
| --- | --- | --- | --- | --- |
| `id` | `integer` | no | `1` | A single row, like the previous one |
| `number_prefix` | `text` | no | `'CS'` | The numbering prefix: `^[A-Z0-9]{2,8}$` |
| `main_assignment` | `text` | no | `'por_turnos'` | The assignment of main tickets: `ninguna` or `por_turnos` |
| `main_notification` | `text` | no | `'al_asignado'` | The new main ticket notification: `a_nadie`, `a_todo_el_equipo` or `al_asignado` |
| `internal_assignment` | `text` | no | `'ninguna'` | The internal assignment: same as the main one |
| `internal_notification` | `text` | no | `'a_nadie'` | The new internal ticket notification: same as the main one |
| `updated_at` | `timestamptz` | no | `now()` | When the row was changed |
| `updated_by_id` | `bigint` | yes | — (null) | Who changed it: `users(id)` with `ON DELETE SET NULL` |

> **Finding, amended in `docs/modules/tickets.md`** (approved on 2026-09-24). That document describes
> the assignment and notification per ticket type (section 3.3) and says the prefix lives here, but
> **it named a `ticket_number_settings` table that only collected the prefix**: the assignment did not fit
> anywhere. That table is renamed **`ticket_settings`** and collects the five things, which is what
> is described above. The amendment is recorded in the change log of `tickets.md` and its migration
> table was corrected.

**A constraint that cannot be forgotten**: with an assignment of `ninguna`, the `al_asignado` notification makes no
sense —there is no one to notify— and the database rejects it, **in both ticket types**:
(`CHECK (main_assignment = 'por_turnos' OR main_notification <> 'al_asignado')` and the same with
`internal_assignment`/`internal_notification`), just as the screen does, which does not offer that
combination.

The fifth single-row table is **`ai_settings`**: mode, provider, URL, model, authentication,
encrypted credential, privacy confirmation and last test. Its details and API are in §12.3.
`ai_model_acceptances` and `ai_recovery_batches` are auxiliary tables with several rows: they record licenses and
recovery batches, so they are not part of the single-row pattern.

**The migration** (`backend/migrations/v1.0.0.sql`) creates the five tables and **seeds their single row**
with the factory values of section 2, so that the application starts without going through the configuration
screen. **The installation name is one more column of that same table**, and the
accounts and tickets used in development go in `backend/migrations/v1.0.0_dev_en.sql` or `backend/migrations/v1.0.0_dev_es.sql`,
the sample data file applied only in development with `scripts/dev-seed-en.sh` or `scripts/dev-seed-es.sh`
(`docs/ambientes.md`, sections 3.3 and 5).

## 4. How whoever needs it reads it

**No one reads another module's tables** (modularity rule, `docs/arquitectura.md`, section 4). What
exists is **a service function**, which is what `docs/modules/tickets.md` already decided:

| Who | What it asks for | What for |
| --- | --- | --- |
| `tickets` | The prefix and the assignment for that ticket type | Number a ticket and assign it when it is created |
| `users` | It does not query language | Accounts no longer store a personal preference |
| `auth` | **`Access()`: the entry method and the two configurations, with their secrets** | Sign in: decide where entry happens and talk to the directory or the realm |
| `ai` | **`AI()`: mode, provider, address, model, authentication, credential and global language** | Write with the configuration in force on each request |
| The frontend (`core`) | All the configuration | The Settings screen, the logo, the primary color and the fields of the two paths |

- **No cache**: it is a row read by its primary key, and a cache that fell behind
  would make a prefix change go unnoticed. When it is saved, it is saved, and the next read sees it.
  **And with how sign-in works, it must be that way**: changing the method or the directory has to
  take effect **on the next attempt**, without restarting anything.
- **The `settings` module does not notify anyone**: the others read when they need to. There are no events or
  subscriptions to maintain.
- **What `auth` needs is declared by `auth`** —an interface with a single method— and this
  module implements it; what this module needs to test a connection is declared by this module and implemented by the
  modules that know how to speak the protocol: **`Prober`** (directory and Keycloak, implemented by `auth`),
  **`ProberDeIA`** (the engine, implemented by `ai`) and **`ProberDeCorreo`** (outgoing email, implemented by
  `mail`). The wiring is connected in `main.go`, which is the only place where the modules know each other
  (`docs/arquitectura.md`, section 4).

## 5. The brand: the logo and the name

**The logo belongs to the installation, not to a person.** It is shown on the entry screen, in the
sidebar and on the Settings screen, and **an Administrator can replace it** with their
institution's.

### 5.1 Two slots, both optional

A logo works for a light background or for a dark one, almost never for both: one with dark ink
**is not readable on a dark background**, and vice versa. That is why there are **two slots** —and that is why
**the factory brand is also two files**, one per theme:

| Slot | What it is for | Mandatory? |
| --- | --- | --- |
| **Light** | The logo shown in **light themes** (Paper, Mist, High contrast, Light) | No |
| **Dark** | The one shown in **dark themes** (Graphite, Night, Sepia, Dark) | No |

**Which one is shown**, in this order:

1. The one from the slot matching the theme being viewed.
2. If that slot is empty, **the one from the other slot** (better a logo designed for the other background than
   none).
3. If there is none, **the factory logo for that theme**, which the application ships with.

Thus, whoever uploads a single logo sees it in all eight themes, and whoever uploads both sees the one that
corresponds. **And if they have uploaded none, they see the factory one for the theme they have set**,
which is what makes the entry screen readable in the dark theme just as in the light one.

**The two factory logos are application files**, not installation files: they live in
`frontend/public/` (`logo-catalina-support-light.png` and `logo-catalina-support-dark.png`) and **they are not
configured from the screen**. Changing them means replacing the file and rebuilding the frontend,
and that is done by whoever maintains the application, not by an administrator.

### 5.2 What can be uploaded

| What | Rule |
| --- | --- |
| **Formats** | `png`, `jpg`, `jpeg`, `webp` and `svg` |
| **Maximum size** | **1 MB** |
| **Maximum side** | **2000 pixels** wide and high (SVG has no pixels: its `viewBox` is measured) |
| **Name** | Whatever the file comes with does not matter: **the application sets the name** |

- **The content is validated, not the extension**: the first bytes of the file are checked. A `.png` that
  is something else inside is rejected.
- **SVG is served isolated**: it is downloaded with `Content-Security-Policy: sandbox` and
  `X-Content-Type-Options: nosniff`, so that **it cannot execute anything** if someone opens its address
  directly. Inside the application it is used as an image, where an SVG cannot do anything.
- **Nothing is cropped or stretched**: the logo is shown whole, with its proportion, inside a
  rounded-corner box with a theme border. A logo with an opaque background —like the factory
  one— looks like a solid piece, and one with transparency blends in; both look good without touching the
  file. **Neither the background nor the shape is assumed**, because the system has to work for
  any logo.

### 5.3 Where it is stored

- **On disk**, in `_files/brand/`, which is the folder that already exists for files (ticket
  attachments) and which **is included in the backup** (`docs/ambientes.md`).
- **The application generates the name** (`logo-claro-20260924T181500.png`): thus there are no two files
  with the same name, a path cannot be slipped in, and it is known when each one is from.
- **The new one is saved before deleting the old one**: if something fails halfway, the previous logo remains and
  not an installation without a logo.
- If the file in the table is no longer on disk —someone deleted it by hand— the application **behaves
  as if that slot were empty** and shows the appropriate one. It does not fail: the brand is not a
  critical function.

### 5.4 How it is served

The logo is seen **on the entry screen, before anyone has signed in**, so its endpoint is
**public**: it carries no session. It shows the institution's logo and nothing else.

**It is served with a version**: the address carries the stamp of the last time it was changed
(`/api/settings/brand/logo?theme=claro&v=20260924T181500`). Without that, a browser that already had the old
logo in its cache would keep showing it after replacement, and the administrator would believe that the
change did not work. With the stamp, the response can be cached forever; without it, it is requested every
time.

### 5.5 Replacing it, and returning to the factory one

- **Uploading** a logo for a variant **replaces** the one in that slot, without asking and without
  deleting the other slot.
- **Removing** the custom logo from a variant returns that slot to the factory logo.
- **None of this touches the factory logo**: the file shipped with the application is never overwritten,
  so it is always possible to return to it.
- The screen shows **how each slot is**: with its thumbnail, the file name, its size and
  when it was changed, or «factory» if it is empty.

### 5.6 Where it is shown, and where it is not

| Where | How |
| --- | --- |
| **Entry screen** | Above the form, centered, with a maximum height of 128 px on mobile and 160 px from tablet up |
| **Sidebar** | At the top, next to the product name, with a maximum height of 32 px |
| **Settings** | The brand: the two slots with their preview, upload and return to the factory one |

- **It does not go in emails**: `docs/modules/mail.md` decided that emails are HTML **without images**,
  and a logo per email would require serving it externally or attaching it to every send.
- **The browser icon (favicon) is an application file** (`frontend/public/favicon.ico`),
  **not a configurable image**: a favicon must be an `.ico` with several sizes inside, and that
  is a matter for whoever maintains the application, not for the Settings screen. It is changed by replacing
  the file, like the two factory logos.

### 5.7 The installation name

**It is what this installation is called**, and it replaces the text «Catalina Support» that the
application shipped with: the institution, the company or the team. It is configured on the Settings screen,
**next to the logo**, which is how it was requested: customizing the brand is the logo **and** the name.

| Where it is read | How |
| --- | --- |
| **Entry screen** | Below the title, where the form is: it is the first thing read on arrival |
| **Sidebar** | At the top, next to the logo |
| **Browser tab** | It is the tab title, on all screens |

- **It travels in the public brand** (`GET /api/settings/brand`), like the logo: all three things are
  read **before anyone has signed in**, so they cannot depend on a session.
- **It is counted in characters, not bytes** («Ayuntamiento de Ávila» is 20): what is measured is what
  is seen, and the limit of 60 is set where it is shown, which is a narrow menu.
- **Empty returns the factory one**, and the field itself says so: it is the way to undo the change without
  being left without a name, which the database does not allow either.
- **It does not change emails** (project owner decision, 2026-09-25): subjects and texts remain
  as they are, and **the sender comes from `smtp_from_name` and `smtp_from_email`**, which live in the table
  itself (section 5.12); the installation name does not touch them. Customizing emails is another matter
  —the `mail` editor is for that— and is not mixed with this.
- **The tab title does not carry the route's**: all routes carried the same text, so they
  did not say which screen you were on. The installation sets it, in a single place (`TitleStrategy`).

### 5.8 Authentication method: the method and the two configurations

**The installation signs in through one method at a time**, and that is chosen here: `local` —accounts live in
this database—, `ad` —the password is the organization directory's— or `keycloak` —you sign in from
the realm screen—. The other two **are turned off**, even if configured: it is not that they do not
work, it is that the installation does not offer them.

| Method | What the entry screen offers | Who can sign in |
| --- | --- | --- |
| `local` | The email and password form, and the link to recover it | The database accounts |
| `ad` | **The same form** —what changes is who answers—, without the recovery link | Anyone in the directory |
| `keycloak` | The button that leads to the realm, and **no form** | Anyone with an account in the realm |

- **The factory account always signs in**, whatever the method: it is the only door that cannot be
  closed, because without it choosing the wrong method would leave the installation with no one to change it. With
  the method set to Keycloak there is no form, so **its door is a discreet link** on the entry
  screen («Sign in as administrator») that gets it in. With the other two methods, the same form
  works for it.
- **With the method set to `ad` or `keycloak`, local accounts —Support and Development included— cannot
  sign in.** That is the consequence of «one at a time» and it is stated where it belongs
  (`docs/usuarios-y-permisos.md`, section 5): an installation that moves to AD is left with the directory
  as its only door, and going back is choosing `local` again.
- **An empty method can be selected to fill in its fields, but it cannot be activated without being configured**: the backend rejects the save
  (`settings.method.notConfigured`, 422). That is what prevents leaving the installation with the factory
  account as its only door through an oversight.
- **The two paths are tested before being saved**, with their button: the test lets the `auth` module
  connect with the service account —without validating anyone's password— or read the realm
  document, and **it does not touch the database**. That is what avoids saving a configuration with which no one can
  sign in and discovering it the next day.
- **Secrets never leave through the API.** What `GET /api/settings` returns is
  `directory.passwordSet` and `keycloak.secretSet`: two booleans that say whether one is set. When
  saving, **empty means «do not change it»**, not «delete it»: the screen leaves the field blank and
  only whoever wants to change it writes in it.
- **The secret is in the database, and that has a consequence that must be stated**: it goes inside the database
  backup (`docs/ambientes.md`, section 6), so **a database backup is one more secret**. That is the price of being able to change it
  from the screen, and it is what the project owner decided on 2026-09-25.
- **Changing the method or the directory takes effect on the next attempt**, without restarting anything: the entry
  path reads the database on each attempt (section 4). That is what makes this configuration and not
  deployment.
- **Tests do not save**: you can test with what is on screen and then decide. In email,
  the button is called **Send test email** because it completes a real send; in the other
  integrations it keeps **Test connection**.
- **The first-start wizard reuses these tests**, without duplicating them: `POST
  /api/setup/entry/test` calls **`Prober`** for the directory or Keycloak, and `POST
  /api/setup/mail/test` calls **`ProberDeCorreo`** for outgoing email. They are public —there is still no
  session— and they are protected by **the installation seal**, like the rest of `/api/setup/**`, not
  by an Administrator permission. The email one sends the technical message to the ephemeral recipient and
  applies its public limit (`docs/primer-arranque.md`, section 3.1).

### 5.9 The system version

**The installation says which version it is**, and it is read in two places: **in the exit row of the
sidebar, aligned to the right**, and **in the footer of the entry screen**
(`docs/interfaz-y-experiencia.md`, sections 3.1 and 6.4).

- **It is not configurable, and that is why it is not in the closed list of section 2**: it is a fact about the
  software, not about the institution. **It is written in a single place in the backend** —`shared/version`, a
  constant— and it is changed when a version is closed, together with the git tag.
- **It travels in the public brand** (`GET /api/settings/brand`), which is what the application requests on
  startup: thus the menu and the entry screen have it **without a second call** and without a session,
  because the entry is painted before anyone signs in. **A product version is not a secret**,
  and it says nothing about anyone.
- **It is shown as `v1.0.0`**: the `v` is added by the interface, and the value the backend stores and publishes
  is `1.0.0`.
- **No build date**: it would change on every build, and two installations with the same code
  would say different things. What needs to be known —which version is installed— is said by the number.
- **The application version and the database schema version are the same thing**: `1.0.0` is the version
  of this repository, which is the one named by the migration `v1.0.0.sql`. When there is a `v1.1.0.sql`,
  there will be a `1.1.0` version to state.

### 5.10 The time zone and the public address

**They belong to the whole installation, not to numbering**, so they live in **their own card** of the
screen, with **their own button** (decisions 14 and 15). The `PUT` sends **the whole configuration** —it is
a document, not a patch—, so the only thing that changes is which button requests it and with what message
it is confirmed.

**The region** is chosen from **a list with a search box**, it is not typed:

- Before the search box, a read-only field explicitly shows the zone the form has, with its exact
  IANA identifier —for example, `America/Guayaquil`—. On load it shows the saved zone and when another
  is chosen it updates immediately, even if **Save** has not yet been pressed. The text can be selected
  and copied, but not edited.
- The list is **the browser's zone catalog** (`Intl.supportedValuesOf('timeZone')`), **plus the
  fixed offsets** `Etc/GMT±1…12`, which that catalog leaves out and are the way to say «UTC-5» without
  daylight saving time. If the browser does not provide the function —or returns empty—, it falls back to **a short
  handwritten list** with the usual zones: the screen is never left with nothing to choose.
- **Below, the chosen one is explained**: the time it is there **right now** and its offset from
  UTC. They are calculated with `Intl.DateTimeFormat` (`timeZoneName: 'shortOffset'`) inside a `try`, so
  that an invalid zone **does not break the screen**: nothing is shown.
- **The zone name is saved**, and the server checks it against the zone database **embedded
  in the binary** (`time/tzdata`): a zone that does not exist is rejected with
  `settings.timeZone.unknown`.
- **It is read by whoever has to write a date**: the interface, through the public brand (`timeZone`), and
  emails, through what the mail module asks this one. **Stored dates remain in UTC**, so
  changing it does not move any ticket; and emails that carry a date —the password-changed
  notice— and the editor preview are written in it.

The field label reuses the strict texts that already exist: **Región horaria seleccionada**
in Spanish and **Selected time zone** in English. It is the same pattern as step 3 of `/setup`; on both
screens the value only changes when a valid option from the list is chosen. The project owner confirmed in
the 2026-10-05 review that the technical identifier and the appearance are also reused, with no
visual or functional differences.

**The public address** is written with its scheme, its host and **its port if needed**, and
**`localhost` is valid**: it is what allows testing the installation locally and using it seriously without touching the
server. The server validates it —`http` or `https`, with a host— and rejects anything else with
`settings.publicUrl.invalid`; the database has its own limit for the same. It is **the base of the links
that go out in emails** and of **the Keycloak callback**, and **if there is none** —neither in the
configuration nor in `PUBLIC_APP_URL`, which remains as a fallback— the link **fails with its key**
(`auth.publicUrl.missing`) instead of sending a broken link.

**And if it is not https, the screen warns about it** (`configuracion.direccionInsegura`): the password and the
session travel unencrypted. **It blocks nothing** —the project owner chose it that way—, because the installation
has to be testable locally; the warning appears in Settings and not on the entry screen.

### 5.11 The order of the screen

**Each card carries what its title says** (project owner decision, 2026-09-30). The screen had
two incoherences that were visible when using it: **«Numbering and assignment» only had inside it the language
of the installation**, and **«Time zone and public address» had inside it the prefix and the whole
assignment**. It was ordered like this, from what the installation **is** to how it **enters**, **works** and **notifies**:

| # | Card | What it contains |
| --- | --- | --- |
| 1 | **The installation** | **The global language as the first control** and the name. One button: **Save the installation** |
| 2 | **The brand** | **The logo and the institutional color**, which are the institution's identity and were in two cards for no reason. The logo is saved when uploaded; the **Save the brand** button is for the color |
| 3 | **Authentication method** | The method and its two configurations, with their connection tests (section 5.8) |
| 4 | **Numbering and assignment** | The prefix and the assignment of the two ticket types (sections 2.2 and 3.3 of `docs/modules/tickets.md`) |
| 5 | **Time zone and public address** | The time zone and the address (section 5.10) |
| 6 | **The AI engine** | Mode, provider, connection or local catalog, model and **Test and activate** (sections 5.13 and 12) |
| 7 | **Emails** | The link to the editor, which belongs to the `mail` module |

**The safeguard is in an interface test**: it looks at the titles in order and checks that the name and
language are together, that the logo and color are together, that the prefix and assignment are in
numbering, that the region does not carry numbering inside and that the AI engine has its own. It is what
prevents it from being misplaced again when a field is added.

### 5.12 The installation seal, and the first-start wizard

**The installation is configured before it has a door, and it is done only once**
(`docs/primer-arranque.md`). What sustains it is **a seal**: `installation_settings.installed_at`,
**null** while no one has completed the wizard.

- **Null is «not installed»**: the application leads to `/setup` and the API accepts configuration.
- **Set is «installed»**: `/setup` leads to the entry and **the whole** wizard API answers
  **409** `setup.alreadyInstalled`. That is what prevents rewriting the configuration of a running installation,
  and it is the only write in the application **without an account behind it**: what protects it is not
  a permission, it is the seal.
- **The seal is set on completion** and saved with `installed_at = now()`.
- **And it does not appear in production because of an upgrade**: the migration backfill seals the rows that already
  existed, and it only does so the first time the column is born (see the backfill block in
  `v1.0.0.sql`).

**The five steps** save one at a time —the installation, how entry works, where it is, email and AI—, and
each one validates **only its own**, because the rest is not set yet. The state they return
—`GET /api/setup` and the responses of each step— is what allows **continuing where it was left**, and **it
carries no secret**: about passwords it only says whether one is set.

**Outgoing email lives in this table** (section 5 of the wizard document) and the `mail` module
reads it on every send through the interface it declares. **There is no fallback in the environment**: the
`SMTP_*` variables were removed, as were those of the directory and Keycloak (project owner
correction, 2026-09-30).

### 5.13 The AI engine

**Its address and model are configured from the screen** (decision 16), with **its card** and its
«Test connection» button. It is the engine that writes the «Reason» and the «Last action» of each
ticket (`docs/modules/ai.md`), and it remains **optional**: without it the help desk works entirely,
with those two fields without text.

- **The address is validated like the public one** —`http` or `https`, with a host— and **empty is valid**: it is «this
  installation has no engine», not an error. With something written that is not an address, the key is
  `settings.aiUrl.invalid`, specific to this card so that whoever reads it knows where the problem is.
- **The model is free text**: the engine is the one that names it, and the installation does not carry a catalog of
  models. Empty means «the environment's».
- **The «Test connection» button tests the address that is on screen**, and **only that one**: when opening
  the screen the field comes filled with what is in the database, so **empty means it has been deleted on
  purpose** and then nothing is tested —`settings.aiUrl.invalid` is returned, «there is nothing to
  test»— instead of falling back to the saved address. It asks **the engine's health check**
  (`<address>/health`) with a short timeout, **saves nothing** and answers `{"status":"ok"}` —the same
  as the directory and Keycloak tests— or its error key (`settings.aiUrl.invalid`,
  `settings.ai.unreachable`); the reason is left in the log.
- **The `ai` module reads the configuration on each request**, not at startup: the address and the model
  come from the database every time a draft is to be written or health is to be checked, so **changing them takes effect without
  restarting anything**. It is the same rule as the entry method (section 4).
- **If the installation has no engine, the environment acts as a fallback** (`AI_URL`, `AI_MODEL`): the two
  variables remain in the environment files so as not to break an installation that already had it
  set that way, with the note that the source is the database, like `PUBLIC_APP_URL` (decision 14). The
  fallback is resolved **in the wiring**, not in the AI module: `settings.AI()` says whether there is an engine, and
  if there is none the module uses its own.

## 6. The endpoints

| Method and route | What it does | Who |
| --- | --- | --- |
| `GET /api/settings` | The installation configuration, with the brand status | Administrator |
| `PUT /api/settings` | Changes it. The prefix **does not touch numbers already issued** | Administrator |
| `POST /api/settings/brand/logo?variant=claro\|oscuro` | Uploads or replaces the logo for that slot (`multipart`) | Administrator |
| `DELETE /api/settings/brand/logo?variant=claro\|oscuro` | Removes the custom one and returns to the factory one | Administrator |
| `GET /api/settings/brand/logo?theme=claro\|oscuro` | **The file that applies** for that theme. **Public** | Anyone |
| `POST /api/settings/brand/site-icon` | Uploads or replaces the tab icon (`multipart`) | Administrator |
| `DELETE /api/settings/brand/site-icon` | Removes the custom icon and returns to the effective light logo | Administrator |
| `GET /api/settings/brand/site-icon` | The custom icon or the custom light logo. **Public** | Anyone |
| `GET /api/settings/brand?color=%23rrggbb` | **Preview**: the same resolved colors, but **without saving anything**. It is what allows showing how it will look before changing the house color | Anyone |
| `GET /api/settings/brand` | **Public.** The **installation name**, **the system version**, the already resolved color and whether there is a custom logo: it is what the application reads to paint itself before anyone signs in | Anyone |
| `POST /api/settings/directory/test` | **Tests** the directory configuration that arrives in the body, **without saving it**: it connects with the service account and closes it. It does not validate anyone's password | Administrator |
| `POST /api/settings/keycloak/test` | **Tests** the Keycloak configuration that arrives: it reads the realm document and checks that it says where its entry screen is | Administrator |
| `POST /api/settings/ai/test` | **Tests** the AI engine: it asks `<address>/health`. The address arrives **only** in the body (`url`); empty is «there is nothing to test» and it answers with its key. It answers `{"status": "ok"}` | Administrator |

- **`GET` and `PUT /api/settings` were already decided** in `docs/modules/tickets.md`, section 5, and they are
  administrator-only. Here the brand status is added to the response.
- **The logo does not travel inside `PUT /api/settings`**: it is a file, it goes by `multipart` on its own
  route, like ticket attachments. **The name does**: it is text, and it goes with the rest of the
  configuration.
- **`GET .../logo` with nothing uploaded answers 404**, and the frontend shows the factory logo. It is not an
  error: it is normal in a new installation.
- **`PUT /api/settings` also carries the method and the two configurations**, and `GET` returns them
  **without the secrets**. The two test endpoints are separate and **by `POST` with the configuration in
  the body**, because they test **what is on screen** and not what is saved: if they read the database,
  they would not serve the only purpose they serve.
- **The three tests answer `{"status":"ok"}`** —the same shape in the directory, Keycloak and the
  engine, so that the screen does not have to know which one it is testing— or their error key, which is what
  the screen needs to say what happened without the backend telling its details: the reason —it does not
  answer, the bind is invalid, the realm does not exist— is left in the log, which is where whoever administers reads it.

## 7. The error keys

| Key | When | Code |
| --- | --- | --- |
| `settings.name.tooLong` | The installation name exceeds 60 characters | **422** |
| `settings.language.unknown` | The language is neither `es` nor `en` | **422** |
| `settings.primaryColor.invalid` | The color is not a six-digit hexadecimal | **422** |
| `settings.numberPrefix.invalid` | The prefix does not comply with `^[A-Z0-9]{2,8}$` | **422** |
| `settings.assignment.unknown` | The assignment value is neither `ninguna` nor `por_turnos` | **422** |
| `settings.notification.unknown` | The notification value is none of the three | **422** |
| `settings.notification.withoutAssignment` | An `al_asignado` notification is requested with a `ninguna` assignment | **422** |
| `settings.logo.format` | The file is not `png`, `jpg`, `jpeg`, `webp` or `svg` | **422** |
| `settings.logo.invalid` | The content is not a valid image, or exceeds the maximum side | **422** |
| `settings.logo.tooBig` | It weighs more than 1 MB | **413** |
| `settings.logo.variantUnknown` | The requested variant is neither `claro` nor `oscuro` | **404** |
| `settings.method.unknown` | The entry method is neither `local`, `ad` nor `keycloak` | **422** |
| `settings.method.notConfigured` | A method that is not yet configured is chosen | **422** |
| `settings.directory.incomplete` | The directory is missing the server or the search base | **422** |
| `settings.keycloak.incomplete` | The realm is missing the issuer, the client or the callback | **422** |
| `settings.directory.unreachable` | The directory connection test has failed | **422** |
| `settings.keycloak.unreachable` | The test with the realm has failed | **422** |
| `settings.aiUrl.invalid` | The AI engine address is not `http`/`https` with a host, or there is none to test | **422** |
| `settings.ai.unreachable` | The AI engine test has failed: it does not answer or answers badly | **422** |

## 8. What was decided when reviewing this document

| # | Decision | It was left like this |
| --- | --- | --- |
| 1 | **The logo limit** | **1 MB and 2000 pixels per side**. A logo over 1 MB adds nothing at 160 px high, and one over 2000 px only bloats it |
| 2 | **What the installation language does** | **Corrected by section 12**: it is global and governs the interface, emails and new summaries; it does not belong to accounts |
| 3 | **The institutional color** | The factory one is the blue already set, and **it only affects the two factory themes**: the six fixed ones carry their own accent |
| 4 | **Two logo slots, both optional** | Yes, and if there is only one it is used in all eight themes. **The factory brand is two files, one per theme** (project owner decision, 2026-09-26): the ones shipped with the application, replaceable but **not configurable from the screen** |
| 5 | **`ticket_number_settings` becomes `ticket_settings`** | The amendment to `docs/modules/tickets.md`: the assignment did not fit in any table |
| 6 | **The logo does not go in emails** | It was already decided in `docs/modules/mail.md` (HTML without images) |
| 7 | **The logo endpoint is public** | It has to be: the entry shows it before signing in. It only returns the institution's brand |
| 8 | **What enters the Settings screen now** | **The brand and the institutional color**, which are what its users already have. The language arrives with account creation, and the prefix and assignment with tickets |
| 9 | **The installation name** | It is configured, and it replaces «Catalina Support» in the application. **Project owner decision, 2026-09-25**, with two corrections from him when it was proposed: **emails are not touched** —neither their subjects nor their text— and **the sender still comes from `SMTP_FROM_NAME`**, which was then from the environment (**superseded on 2026-09-30**: the sender moved to `smtp_from_name` and `smtp_from_email` in `installation_settings`, section 5.12). The 60-character limit and empty returning the factory one are my proposals, applied and recorded here (section 5.7) |
| 10-bis | **Each method shows only its configuration** | **Project owner decision, 2026-09-27**: in «Authentication method», the directory panel appears **only** with the organization method and the Keycloak one **only** with Keycloak; with «Application accounts» neither of the two is shown. **It does not change what is saved** —the `PUT` still sends both configurations and saved secrets are not touched— nor the «unconfigured method» notice, which still follows the view of the chosen method: it changes **what is shown**, which was showing work that was not called for |
| 10 | **The directory and Keycloak are configured from the screen** | **Project owner decision, 2026-09-25**: both paths leave the environment and enter the closed list of what is configurable, with **two single-row tables** and their section (5.8). **Secrets are stored in the database and never returned by the API** —only `passwordSet` and `secretSet`—, with a «Test connection» button that saves nothing. Accepted consequence and **confirmed by the project owner on 2026-09-25**: **a database backup is one more secret**, with the machine's permissions and without leaving it (`docs/ambientes.md`, section 6) |
| 15 | **The installation region decides how dates are read** | **Project owner decision, 2026-09-29**: the installation has **a time zone** —chosen from a **list of zones (IANA) with a search box**, and the screen shows the time there and its offset— and **it decides how all dates are read: those of the interface and those inside emails**. **Stored dates are not touched**: they remain in UTC, so **changing the zone does not move any ticket**, it only changes how they are read. And it is not an ornament: until today dates were painted with the **browser** zone of whoever was looking, so the same action was read at different times depending on where they were. **The zone name is saved**, not an offset: thus daylight saving time, where it exists, is respected on its own. The public brand carries it (`timeZone`), because the entry is painted before signing in |
| 14 | **The public address: email links and the Keycloak callback** | **Project owner decision, 2026-09-29**: the installation has **a public address** —scheme (`http` or `https`), host and port, **also `localhost`**— and it is **the base of the links that go out in emails** (account creation, password reset, notices) **and the Keycloak redirect address**. The application is still navigated **relatively**, so it works through any address used to reach it: that is what allows testing it locally and using it seriously without touching the server. Until today that was **an environment variable** (`PUBLIC_APP_URL`), mandatory in production; it moves to living in the database, like the directory and Keycloak, and the variable remains **as a fallback** for an installation that already has it set |
| 13 | **The warning that there is no https, in Settings** | **Project owner decision, 2026-09-29**: if the public address **is not https**, the Settings screen **warns** —the session and the password travel unencrypted— **and blocks nothing**: the installation works the same. The project owner chose that the warning appears **only in Settings** (and not also on the entry screen, which was the proposal): whoever decides how it is exposed is whoever configures it |
| 12 | **The system version** | **Project owner decision, 2026-09-25**, with **a choice of his that was not my recommendation**: it is **a backend constant** that travels **in the public brand**, and it is read **in the sidebar —in the exit row, on the right— and also in the footer of the entry screen**. I proposed only the menu; he decided that it should also be seen before signing in, which is where it serves to say which version is being viewed. **Only the number, with its `v`, without a build date** |
| 11 | **Testing a failing connection answers 422** | **Confirmed by the project owner on 2026-09-25**, as implemented: what is being tested is **the data on screen**, so a failure is «that configuration does not work» and not «the service is down». The reason —it does not answer, the bind is invalid, the realm does not exist— is left in the log |
| 11 | **One entry method at a time** | **Project owner decision, 2026-09-25**, when asked whether the three paths coexisted: **one at a time**, changeable afterwards, and **the factory account can always sign in** —its discreet door on the entry screen is what he approved—. Consequence stated and accepted: with the method set to `ad` or `keycloak`, local accounts cannot sign in, and that amends section 5 of `docs/usuarios-y-permisos.md` |
| 16 | **The AI engine is configured from the screen** | **Project owner decision, 2026-09-30**: **the engine address and model move to Settings**, with **their card** and **their «Test connection» button** (section 5.13), instead of living in the environment. The `ai` module **reads the configuration on each request** and, when nothing is set, **uses `AI_URL`/`AI_MODEL` as a fallback**, which is why they remain in the environment files. The test asks `<address>/health` with a short timeout and **saves nothing**; **it tests only the address provided** —empty is «there is nothing to test»— and **the three connection tests answer `{"status":"ok"}`**. The new keys are `settings.aiUrl.invalid` and `settings.ai.unreachable`. Changing `AI_PALABRAS` and `AI_ESPERA_SEGUNDOS`, which remain environment settings, is **out of scope** |

## 9. The institutional color, resolved

The administrator chooses **one** color, but the application needs four values, because the same color
is not read the same on a light background as on a dark one. The backend calculates them and sends them already
resolved:

| Value | What it is |
| --- | --- |
| `colors.light` | The color for light themes: the chosen one, **darkened** if it does not reach 4.5:1 on white |
| `colors.dark` | The one for dark themes: the chosen one, **lightened** if it does not reach 4.5:1 on the dark background |
| `colors.onLight`, `colors.onDark` | The button text color: white or almost black, **whichever gives the most contrast** |

- **Only lightness moves**, not hue: the institution's color still looks like what
  whoever chose it picked.
- The four values appear in the public brand endpoint, and the application puts them in CSS
  variables. **The six fixed themes do not use them**: they carry their own accent.
- It is checked with difficult colors —white, black, yellow, red— in the module tests:
  with any of them, the four values are readable.
- **The screen asks the backend for the preview**, with the chosen color and without saving it: resolving it in
  the frontend would mean having the same rule in two places, and one day they would say different things. The preview
  shows **what is really going to be seen**, including the adjustment if the color is not readable.

## 10. What this document enables

With `docs/modules/settings.md` approved, the following can be done:

1. **The complete brand**: the table, logo upload with its validation, the public endpoint with
   version, SVG isolation, and returning to the factory one.
2. **The logo on the entry screen and in the sidebar** —the menu arrives with the application
   shell—, and on the Settings screen, with its two slots.
3. **The prefix and assignment read from where they live**, which is what `tickets` needs to
   number and assign when its module arrives.
4. **Authentication method, configurable from the screen**: the method, the organization directory and
   Keycloak, with their two connection tests and their secrets stored without leaving through the API. It is what
   allows delivering the installation without asking anyone to edit a file on the server.
5. **The AI engine, configurable from the screen** (decision 16): its address and model, with its
   connection test, read by the `ai` module on every request. The environment remains as a fallback.

## 11. Visible results in Settings

### 11.1 Finding and rule

The screen places its general message before the first card. Since Settings contains seven
cards, the result of **Test connection** in the AI engine —and that of any action located
below— can end up outside the window. The answer is not missing: its location hides it.

All transient results that today use that general message move to the shared toast of
`docs/interfaz-y-experiencia.md`, section 15:

- saving the installation, the brand, authentication, numbering, the region or the AI engine;
- uploading or restoring a logo;
- testing Active Directory, Keycloak or the AI engine;
- failing to load the configuration or receiving a general validation or server error.

Context notices that belong to a card remain in it. In particular, the warning
that a public address without HTTPS exposes the session and the password does not become a toast.

### 11.2 Scope and acceptance criteria

The amendment only changes the presentation in `settings-page.*` and reuses the shared component.
It does not modify requests, endpoints, result texts, HTTP codes, validation, persistence or the
order of the cards.

1. The response of the three connection tests is shown fixed in the window from any point
   on the page.
2. The other general results use the same presentation, without keeping a second notice at the
   top of Settings.
3. The toast complies with duration, replacement, closing and accessibility defined in the
   interface document.
4. Notices specific to each card remain visible while their condition holds.
5. The page tests cover at least one successful and one failed response; Playwright tests
   the AI engine from the bottom on PC and mobile without scrolling the window to the top.

### 11.3 Decision and review

The project owner confirmed on 2026-10-06 that `/settings` and `/setup` should be included in the proposal, and that
transient results should be distinguished from permanent notices. The common review was closed in
`docs/interfaz-y-experiencia.md`, section 15.7: all the results described use the toast, there is no
queue and all forms disappear automatically within ten seconds at most. No open decisions remain, nor any
decisions specific to the backend.

### 11.4 Implementation and verification

`settings-page` shows its general signal through `app-toast`; the conditional notices of the
cards continue as `app-aviso`. Requests, texts and error handling did not change.
The original policy was replaced in section 17: any form remains long enough
to be read, accepts **Close** and disappears automatically after ten seconds.

## 12. Implementation: global language and mandatory AI

### 12.1 Settings order

The **language is the first control in `/settings`**, before the name. It only offers Español and English and
is no longer described as the initial value for new accounts: it governs the interface, dates, emails and AI texts
of the whole installation. Choosing another language previews the current interface, but does not persist the
change until the translation and confirmation of the emails are completed (§12.4).

The language of account creation, user detail and profile and the entry and menu selectors disappear. The
global language must be available before signing in; it travels in the public brand configuration and
governs over any old browser or account preference.

### 12.2 AI card

The card starts with three mutually exclusive options: **Local**, **Own server** and **Provider**.
In local mode it shows the catalog, download, progress, status, requirements, installed models, active
model and deletion of inactive ones. On an own server it asks for URL, model and authentication —none, Bearer,
header with key or username/password—. In provider mode it offers OpenAI, Claude, DeepSeek and OpenAI-compatible,
with suggested models and a free identifier.

The written secret only travels to the backend. Reading returns `credentialSet`; it never returns
the encrypted credential. For an external connection, the card shows what ticket text will leave the
server and requires the Administrator's confirmation before **Test and activate**.

The test performs a complete artificial generation. Saving and activating are a single operation:
it keeps the previous configuration until the new one authenticates, accepts the model and produces the
valid object in the global language. A failure is shown by toast and does not change the configuration.

### 12.3 Persistence and proposed API

The configuration no longer fits in `ai_url` and `ai_model`. A single-row AI configuration is created,
owned by `settings`, with mode, provider, URL, model, authentication, header name,
encrypted credential, privacy confirmation, who confirmed it, configuration date and update
date. The public response never includes that row.

The authenticated Settings API delivers the state without secrets and adds specific operations
to test/activate, query catalog/progress, download, activate and delete models. Local actions
go through the interface that `settings` declares and the model administrator implements; ticket texts
remain the responsibility of `ai`.

The credential master key is independent of `TOKEN_SECRET`. `AI_CREDENTIAL_KEY` must
contain 32 random bytes encoded in base64: if it is missing or invalid the backend does not start; if
it was replaced and the existing encryption no longer opens, the application enters «AI requires
configuration» and allows the factory account to replace the credential. It never logs or returns
the unreadable value.

### 12.4 Atomic language change

`settings` coordinates, but does not translate templates. It asks `mail` for drafts of the fourteen templates
starting from the currently active version. `mail` protects HTML and placeholders, `ai` translates only the text
and `mail` validates that no new placeholders appear and existing ones are not altered. If the destination was not
customized, the screen compares source and draft. If it already had changes, it shows three
versions: active source, AI draft and existing customized destination. The Administrator chooses
one, or combines and corrects the text; it is never silently overwritten.

On confirmation, a transaction saves the fourteen destinations and the global language. If a translation is missing,
a placeholder changed or any write fails, nothing changes. Existing summaries are not part of
this operation: they are preserved and remain visible with their source language; only new ones or
those requested again use the new global language.

With an external provider, the screen asks for financial confirmation exclusively to translate the
templates. If it is rejected, it does not call the provider and allows manually editing the fourteen destinations; manual
editing is always available. With a local engine or own server it generates drafts without that
financial confirmation.

Each authenticated response reports `X-Catalina-Language` and `X-Catalina-Settings-Version`. The version
grows when a global change is applied. The interceptor adopts the language on receiving the next response
and reformats the interface through the global signal, without logging out or querying in the background.

### 12.5 Existing updates and scope

A row already sealed without tested AI is marked as pending. The Administrator or the factory
account can only complete the AI card; the rest of the accounts do not enter the product. `/setup` is not
unsealed nor is its public API reopened.

The persistence of the installation and AI, the DTOs and services of `settings`, the public brand and
`settings-page` change. Prices, quotas, commercial accounts and key administration in each provider's
portal are out of scope.

### 12.6 Recorded decisions and review

The project owner confirmed on 2026-10-06 the decisions recorded in `docs/modules/ai.md`, and corrected
the scope of the language: it will be global, it will be the first item in Settings, it will not belong to
accounts and the change will only take effect after reviewing and saving the email translations. The review added the mandatory 32-byte key, the three-version review when the
destination was already customized, the cost confirmation only before translating templates with a
provider and the alternative of always editing them by hand, in addition to updating open sessions through language and version headers. No open decisions remain;
the correction of 2026-10-07 was explicitly approved.

## 13. Structured memory error of the local engine

The local model administrator client distinguishes lack of memory from a crash. It reads a limited
and structured response with `key`, `requiredBytes` and `availableBytes`; it validates non-negative values and never
returns the raw body. The service exposes a typed error and the `/setup` and
`/settings` controllers respond with the same key and both numbers.

The other failures keep their current translation. A rejected activation does not save or mark as
tested the requested configuration. If the model administrator restored the previous model, the response
is still an error for the new attempt, but the subsequent catalog shows the previous one active and
healthy.

The contract was verified end to end: the 7B answered 422 with a stable key,
6,442,450,944 bytes required and the measured availability; `/settings` kept both numbers through
the toast and the catalog kept showing the 3B active and healthy.

## 14. Implementation: license and source in the public brand

`GET /api/settings/brand` delivers two read-only fields: `license`, with the SPDX identifier
`AGPL-3.0-only`, and `sourceUrl`, with a validated HTTPS address. They are not database columns and cannot
be changed in `/settings`: they belong to the build, like `version`.

A published version will build `sourceUrl` toward the exact tag of its version. The development
build points to the repository and unambiguously exposes its condition through
`development`. The same public response feeds the entry and, once authenticated, the menu. It does not
contain secrets or information specific to an installation.

Tests cover both fields, the development indicator, rejection of a non-HTTPS URL and
frontend compatibility when `sourceUrl` is missing. Correspondence with an exact tag
will be verified when a version is published, work deferred until 1.0.0 is closed. The policy and
publication are defined in `docs/proyecto-abierto.md`; the presentation is defined in
`docs/interfaz-y-experiencia.md`, section 21. Persistence, permissions and the other brand settings
do not change.

The implementation also delivers `development`, which allows presenting `v1.0.0 · dev` without inferring it
from the URL. The 2026-10-10 verification passed all controller, DTO and settings service tests. The isolated
brand walkthrough checked the public contract and its consumption on PC and mobile.
## 15. Implementation: missing settings row

Reads of installation, tickets, directory, Keycloak and AI query at most one row and
distinguish zero rows through `RowsAffected`. They preserve `ErrSettingsNotFound`, which the services
translate to the existing public key. A connection, SQL or scan error still propagates; the global logger,
the API and the five tables were not modified.

## 16. Wizard contracts for brand and models

`/api/setup/ai/models/**` and `/api/settings/ai/models/**` expose the same operations: catalog with
phase and progress, start or resume, cancel, partial delete, retry activation, test the
active one and delete an inactive installed one. The first family keeps the null seal as its only
lock and answers `409 setup.alreadyInstalled` afterwards; the second requires Administrator. License
acceptance preserves actor, version and checksum.

The download request records that the chosen one must be activated and tested. The service coordinates the
work even if the browser request disappears. The catalog distinguishes the operational state from
the tested configuration; it does not present a merely downloaded file as usable. The manual test
resolves `now` through the service clock and the `installation_settings` zone, calls the engine
with artificial content and returns greeting, instant and zone without saving the text.

The active one cannot be deleted. A partial and an installed model are different resources: cancelling
preserves the first, deleting a partial removes it and deleting a model only accepts an inactive one. Public
errors distinguish concurrent download, incompatible state, active model and test failure,
without returning bodies from the internal administrator.

Step 1 adds `POST` and `DELETE` for logo under `/api/setup/brand/logo`, with light or dark variant,
`multipart`, current validations and the seal lock. It reuses the brand service; it does not create another
table or another folder. The public brand changes version when saving or restoring so that the
header requests the image again.

Tests cover identical contracts, authorization or seal, incompatible states, injected clock and zone,
absence of personal data, logo validation and version update. It does not change
the persistence of the installation or the existing authenticated logo API.

## 17. Deactivate, attribute the test and load from the preview

The `/api/setup/ai/models/**` and `/api/settings/ai/models/**` families add a parallel operation
to deactivate the active local model. The first keeps the seal lock and the second requires
Administrator. It only accepts the active identifier; it does not delete the installed one or modify external
models. The subsequent catalog contains no active model until another is activated.

The manual test response extends the contract with `modelId` and `modelName`, in addition to greeting,
instant and zone. `/setup` uses the absence of a healthy active model to block **Next**. The
local configuration no longer depends on a separate selector: activating a card updates the
effective configuration and deactivating it leaves the local engine without a usable selection.

Logo upload and restore keep the current APIs, formats, size and permissions. Only
their presentation changes: the preview triggers the `multipart`, the file name does not form part
of the interface and the public brand still updates its version on completion.

Tests cover permissions or seal, deactivation and reactivation, rejection of the wrong model,
attributed greeting contract and upload/restore regression for both variants. The data model does not change
and no persistence is added.

The full backend suite passes and the frontend compiles with its 227 tests. The browser confirmed
the public contract without a top selector, deactivation and reactivation, the attributed greeting and
the two file controls hidden behind their accessible previews.

## 18. Initial color and integrated confirmations

`POST /api/setup/installation` accepts `primaryColor` along with `name` and `language`. It applies the same
normalization, `#rrggbb` validation, contrast derivation and public brand update as
Settings. The `/api/setup` response returns the saved color to rebuild the step when
resuming. The preview reuses the current endpoint and remains without persistent effects.

The theme is not part of the contract or the database: it continues in `ThemeService` and browser
storage. The six fixed themes do not consume the institutional color; the two factory ones do.

Deletion confirmations do not create a new API or change HTTP codes. They only replace
`confirm()` with `app-dialogo` before calling the existing contracts. Tests cover
color validation and persistence under the seal, the resumable response and the absence of delete
calls when cancelling.

The service normalizes the hexadecimal to lowercase, rejects invalid values and returns the color in
the resumable state. The settings tests, the 227 frontend tests and the build pass;
the browser walkthrough cancelled via Escape without issuing the destructive request.

## 19. Automatic query of the local catalog

`/setup` and `/settings` consume the existing catalog contracts when showing local
administration, not after an explicit action. No endpoint is added and the response does not change: the adjustment is
one of lifecycle and frontend presentation.

Manual refresh keeps the last visible response in memory until the new one completes.
The button does not trigger two simultaneous queries and a failure does not erase valid data already
known. The automatic query follows the same current locks: null seal in `/setup` and Administrator in
`/settings`.

Tests cover one automatic query per entry, manual refresh, exclusion of double
requests and preserving the catalog on error. It does not change persistence, permissions or contracts.

The component test handles the automatic query, checks that the card remains during
refresh and that the disabled button does not open a second request. The suite ends with 228 passing cases
and the frontend compiles for production.

## 20. Reporting thresholds

The Reporting card stores elapsed hours by priority: critical 4, high 24, medium 72 and low 120
as initial values. Only an Administrator can change them. They must be positive integers and
saving uses the common toast. The Dashboard reads the current values on each query; changing them does not
rewrite tickets or history and does not create an SLA, notice or email.

The API and the screen share the validation. The real test temporarily modified the critical value,
confirmed the toast and restored it; PostgreSQL ended at 4/24/72/120.

## 21. Local engine acceleration

The **On this server** card offers Auto, CPU and NVIDIA; Auto is selected by default. The control only
appears for the local engine. Next to the active model, the configured mode and effective execution are read,
for example `Auto · NVIDIA · RTX 3060` or `Auto · CPU`. The automatic test and **Try me!** include
that same attribution in addition to the greeting.

Saving with an active model restarts and tests the process. During `restarting`, incompatible changes are blocked.
A success uses the common toast; the Auto-to-CPU fallback adds a brief notice; missing NVIDIA
explains that `ai-nvidia.yml` must be started and does not save. If the new mode fails, the screen
keeps the previous one after its healthy restoration.

The API does not return UUID, PCI bus, memory or traces. The build and the 241 frontend tests
verify both languages and the contract; the administrator tests cover missing NVIDIA and the
Auto fallback. The visual pass and real NVIDIA hardware were completed on 2026-11-05 with Docker
Desktop WSL2 and an RTX A2000: effective mode `nvidia`, active model and real summaries.

## 22. Configuration of the three SMTP modes

`installation_settings` replaces `smtp_secure boolean` with `smtp_security_mode text`, limited to
`starttls`, `tls_direct` or `plain`, and adds `smtp_insecure_ack boolean NOT NULL DEFAULT false`.
The migration translates `true` to `tls_direct` and `false` to `starttls`. Only the Mailpit development
data declares `plain` and true acceptance; there is no equivalent exception in production.

Step 4 shows an explicit selection with the three alternatives. When choosing `plain`, it presents an
unambiguous warning that credentials and messages travel unencrypted and requires confirmation before
testing or saving. Switching to a secure mode clears the acceptance; returning to `plain` requires accepting it
again. The `/api/setup/mail/test` contract, the final save and the wizard read
carry `securityMode` and `insecureAccepted`; it stops carrying `secure` when the coordinated migration finishes.

The backend rejects `plain` without acceptance and also when the host resolves any address other than
RFC1918 IPv4 or loopback. The same check is repeated on test and on send; screen validation is not enough.
Secrets keep their current rule: they are stored, never returned by the API
and an empty value preserves the existing one where applicable.

The configuration asks for an institutional address of Catalina's own as sender. It does not copy the
sender or credentials from another application.

## 23. Recipient and test send

Step 4 and the administrative card incorporate **Send test to** and replace **Test connection**
with **Send test email**. The field is not part of `installation_settings`: it only
accompanies that request, it is not recovered on reload and it does not alter the recipients of functional emails.
The action is optional and its status does not participate in the requirements to advance or seal.

Both screens send the visible values to the backend, even if they have not been saved. On the administrative
screen, an empty password means reusing the persisted secret; a new password is used for the attempt and is not
persisted. The rest of the validations, including `plain` consent, apply exactly as in a normal send.

`POST /api/setup/mail/test` remains public only before the seal and is limited to five attempts per
IP–recipient pair in fifteen minutes. The authenticated card has the equivalent action
under `/api/settings/**`, protected by administrative capability. Success returns the accepted recipient
to build the precise toast, without claiming final delivery.

## 24. Baseline condition and Active Directory user attribute

`directory_settings` adds `attr_username`, with `sAMAccountName` as the initial value for Active
Directory; the OpenLDAP data uses `uid`. `attr_email` remains the email and `attr_id` the
stable identifier.

`user_filter` now represents exclusively the **baseline security condition**: object classes, organizational
scope and exclusion of disabled accounts. It does not receive text written by a
person. For exact entry and administrative search, `auth` combines that condition with another one
generated and escaped over `attr_username` and `attr_email`.

The open migration of `v1.0.0.sql` updates the known development values and the current installation
without preserving an ambiguous template. A legacy configuration that still contains `%s` is
rejected with a specific key and must be saved again as a baseline condition; it is never executed by silently
reinterpreting it. The card explains the new meaning and tests connection,
service credentials and syntactic validity of the filter, without searching or exposing people.

No fields or secrets are added for Keycloak administration. Until an authorized service
account is available, Settings keeps OIDC for access and `/users` keeps account creation there on
first sign-in.

## 25. Configurable site icon

The brand incorporates a **Site icon** independent of `logo_light` and `logo_dark`. It is optional and
persists in the new nullable `site_icon` column of `installation_settings`; it stores only the
internal file name. Uploading another replaces it transactionally and deletes the previous one
after saving the row. Restoring leaves the column null. Both actions record actor and date
through the existing audit fields.

It accepts PNG, JPEG, WebP and SVG with content validation, the 1 MB maximum and the 2000-pixel
maximum already applied to logos. In addition, both sides must measure at least 32 pixels. The image is not
cropped or stretched: the browser fits it whole, centered, and preserves its transparency. A
rectangular file is valid.

The public cascade is unique: custom icon; if it does not exist, custom light logo; if that does not exist either,
`logo-catalina-support-light.png`. Therefore, restoring the icon or changing/restoring the light logo
can change the effective favicon immediately. The old `favicon.ico` remains only as a static
fallback before the application starts or if the public brand cannot be read.

The public brand adds `siteIcon.filled` and a stamp that changes when modifying the icon or any of
the images that can serve as fallback. `GET /api/settings/brand/site-icon` returns the custom
icon or the custom light logo, with its real type, `nosniff`, SVG isolation and the same
versioned cache policy as the logos. If neither exists it answers 404 and the frontend uses the factory
light logo. It does not generate `.ico`: current browsers accept the four decided formats.

Settings adds `POST` and `DELETE /api/settings/brand/site-icon`, reserved for Administrator and with
the same `multipart` and logo limit. `/setup` offers the parallel routes before the seal. The
configuration response describes whether the custom slot is filled, but never exposes internal paths.
The existing format, content and size keys are reused; an image smaller than 32 × 32 uses
`settings.siteIcon.tooSmall` with a 422 response.

Tests will cover validation, replacement and file cleanup, fallback at its three levels,
cache, isolated SVG, permissions, seal lock and favicon update when changing both the
icon and the light logo. No manifest, installable icons or `apple-touch-icon` are added: the scope
is exclusively the browser tab.

## 26. Local catalog unavailable in Settings

> **Proposal of 2026-10-23.** The review expressly extended to `/settings` the correction discovered
> in `/setup`. The project owner chose 1A, 2A and 5A. No open decisions remain; approval is pending.

If the local model administrator does not respond, Settings shows the failure inside the
local card, offers **Try again** and keeps own server and external provider usable. That 503 does not
by itself activate the global server-down screen. Failures when testing or
saving the chosen mode keep their current toast messages.

**QA correction, 2026-11-03.** The 503 from `POST /api/mail/language/drafts` —AI unavailable
for translation— also does not activate the global server-down screen: it is treated as a failure of an
AI dependency, leaves the notice inside Settings and keeps **Edit manually** accessible.

The focused regression covers Settings on PC and mobile with the general backend healthy and the local
administrator absent. Endpoints, secrets, permissions, AI mandatory status and effective state do not change.
