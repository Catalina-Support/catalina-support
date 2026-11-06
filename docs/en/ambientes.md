# Environments: deployment and testing

[Español](../es/ambientes.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-05
>
> **WSL2 support verified, 2026-11-05.** The NVIDIA variant was tested on Windows 11 with Docker Desktop
> using the WSL2 backend, an RTX A2000 12 GB and driver 596.51. Docker sees the GPU with `--gpus all`,
> the administrator detects it, activates the model and real summaries are generated. It requires Docker
> Desktop with WSL2 and the GPU visible; native Windows containers are not supported.
>
> **Correction of 2026-11-04: production mail does not go in `prod.env`.** It lives in
> `installation_settings.smtp_*` and is configured from step 4 of `/setup` or from Settings → Mail
> with the factory account. Section 4.5 kept the old instruction to edit `prod.env` and restart the
> backend; it has now been corrected.
>
> **Correction of 2026-11-04.** The seeder configures the example local AI when the installation is not
> sealed. Before, it left it sealed with `aiRequired` and an empty `ai_settings`, and users could not
> get past «No permission» until an administrator configured it.
>
> **QA correction, 2026-11-03.** The development seeder deletes all accounts before creating the
> eleven examples. Before, it left alive the requester accounts created by previous tests
> (`solicitante-*`, `ad-alta-*`) and the Users list ended up dirty.
>
> **Amendment proposed on 2026-10-21:** section 28 offers explicit seeders per language and fixes
> English for the public demo restart.
>
> **Amendment approved, implemented and verified on 2026-10-20: self-sufficient upgrade and versioned
> migrations.** Section 27 replaces the manual migration in production and the absence of history. The
> owner closed and approved decisions 1A–25A. The runner passed unit tests, adoption of an existing
> database, a new database, a second start, rollback and the production build; production was not
> deployed.
>
> **Approved, implemented and verified on 2026-10-18.** `ai-nvidia.yml` was added for Linux with
> NVIDIA Container Toolkit without altering the portable CPU deployment. **Extended and verified on
> 2026-11-05** for Windows with Docker Desktop WSL2.
>
> **Amendment implemented and verified on 2026-10-17: public guide and hosted demo.** Section 25
> defines how the README summarizes development and production and presents `demo.catalina-support.com`.
> The development and AI Compose files were verified; production deliberately keeps its requirement
> of an unversioned `config/env/prod.env`. No public environment was deployed or modified.
>
> **Clarification approved and verified on 2026-10-15.** Section 24 records that development already
> uses `ng serve` with a volume, live reload and `no-cache` responses. The logs confirmed detection,
> rebuild and notification to clients; the owner chose 4A and 5A: do not duplicate default options or
> add polling. Forced reload remains the recovery for a disconnected tab. The environment
> configuration was not modified.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** Section 23 adds priority, its stable order and the
> non-blocking critical notice to the three test layers. Production remains out. Production is not
> modified. The owner approved it explicitly.
>
> **Fourth public pass completed, 2026-10-13.** Section 22 records that the four documented
> environments were brought up from the public `main`, completed the approved walkthrough and were
> removed with their volumes. No false errors or audit resources remained; production was neither
> started nor modified.
>
> **Third approved resumption of block 8, 2026-10-13.** The public walkthrough completed all
> functional milestones, but the final log review found false GORM errors on normal absences. Section
> 21 records the evidence, the cleanup and the full repeat chosen by the owner in 4A. The fix passed
> against real PostgreSQL and the fourth pass is pending its arrival in `main`.
>
> **Finding fixed and verified on 2026-10-13: TCP availability and manual migration.** The second
> public pass of block 8 found a race between the socket healthcheck and the TCP connection of
> `migrate`. Section 20 keeps automatic migration in development and tests, adds the manual production
> command and publishes its SQL. The owner closed 1B, 2A, 3A, 4A, 5B, 6B, 7A–9A, correction 10A–13A,
> review 14A–18A and explicitly approved the proposal. The fix passed in disposable instances; the
> full public walkthrough remains to be repeated.
>
>
> **Finding cleaned up and resumption proposed, 2026-10-13.** Block 8 stopped because of an AI
> response in another language and removed all its temporary resources. Section 19 requires a new full
> pass from `main` after integrating the correction explicitly approved on 2026-10-13. Production was
> left intact.
>
> **Block 8 proposal, 2026-10-12: clean public environment.** Section 18 defines the approved removal
> of development data, the temporary clone of `main`, the evidence and the final cleanup of the
> walkthrough defined in `docs/prueba-local.md`, section 16. The owner closed 1A–15B and explicitly
> approved the proposal on 2026-10-12; production is out.
>
> **Amendment implemented and verified on 2026-10-08 (block 6).** Section 17 defines the environment and the matrix of
> verification of the full format review. Existing content remains as-built and production stays
> parked. The owner explicitly approved this amendment on 2026-10-08.
>
> **Finding fixed and verified on 2026-10-08.** The real instance was stopped keeping only its
> isolated volumes after reproducing the 1.5B→3B rejection. Section 16 defines the tests of correction
> 15A–22A and the resumption with the already verified GGUF files. Development and production are not
> touched; the correction was explicitly approved on 2026-10-08.
>
> **Finding and amendment proposed on 2026-10-07 (block 5).** The isolated tests cover the engine with
> fake processes and GGUF files, but validating the real catalog without altering development is still
> missing. Section 15 defines a disposable Compose project, reproducible measurements and full
> cleanup. Decisions 1A–14A were closed by the owner. The proposal was explicitly approved on
> 2026-10-08; existing content remains as-built while the validation runs.
>
> **Finding and amendment proposed on 2026-10-07.** Block 4 needs to complete the isolated tests
> promised for adapters, redirects, container limits and supervision of the local engine. Existing
> content remains as-built. Section 14 records the correction chosen by the owner (4A). The review
> ended with 5A, 6A, 7A and 8A, with no open decisions. The owner explicitly approved the amendment,
> which was implemented and verified on 2026-10-07.
>
> **Implemented and verified on 2026-10-07.** Providers and the administrator were tested with fake
> servers, processes, files and limits, without commercial providers or real GGUF files. The canonical
> pass ended with 180 E2E cases passed and 38 expected skips, on PC and mobile; the instance was
> removed with its volumes on closing.
>
> **Amendment proposed on 2026-10-06.** Existing content remains as-built. Section 13 incorporates AI
> as an installation requirement, separates its local engine per environment and adds the encryption
> and administration variables. The review was closed, the proposal was approved and the
> implementation was verified.
>
> **Amended on 2026-10-06**, explicitly approved by the owner, implemented and verified: the isolated
> suite runs its backend with `go run .`, keeps `./backend:/app:ro` and removes the nested
> `backend_tmp:/app/tmp` mount. The canonical walkthrough ended with exit 0 and the cleanup removed
> all the resources of the instance.
>
> **Amended on 2026-10-04**, in accordance with the approved section 12 of `docs/prueba-local.md`:
> `dev.yml` runs the migration automatically and `backend` waits for it to finish correctly. A new
> local installation is ready with a single `up -d --build`, without example data; the seeder remains
> separate and optional. The mail step receives Mailpit's editable suggestions.
>
> **Amended on 2026-10-03**, in accordance with the approved `docs/prueba-local.md`: The current
> walkthrough is an empty installation with `docker compose -f dev.yml up -d --build` and `docker
> compose -f dev.yml run --rm migrate`, the wizard and admin/admin. The seeder is optional: it keeps
> the five settings tables of a sealed installation; an unsealed installation receives the local demo
> and becomes sealed. The tests use `tests.yml` exclusively; the previous commands on the development
> environment are replaced by the procedure below.
>
> **Amended on 2026-10-02**: the README prioritizes the empty installation, by correction of the
> owner, with the optional examples. The schema alone does not seal a new database: first `/setup` is
> completed and then you log in with `admin`/`admin` locally; in production the password is
> `ADMIN_PASSWORD`. The examples restore the public address of the development domain, which must be
> adjusted to use links on localhost. The build context of the production script is reported as a
> blocker to reproduce it from scratch; no code is changed and nothing is deployed. Checked in an
> isolated Linux copy: schema on a new database (0 accounts and 0 tickets), the five wizard steps in
> the browser, SMTP and AI test, login with `admin`/`admin`, 409 block after sealing, account creation
> and opening the link from the mailbox. Afterwards, the seeder left 11 example accounts, 25 tickets
> and 5 attachments; browser logins as factory, User, Support and Development worked. These checks
> have not been run on Windows or macOS. For the container's browser the internal IP of the frontend
> was used: the host `frontend` is rejected by Angular. The default suite still uses the development
> domain.
>
> **Amended on 2026-10-01 (seventh time)**: **the example data is seeded with a single-use container**,
> `docker compose -f dev.yml run --rm seed-en|seed-es`, **the same on Linux, macOS and Windows** —which is
> what leaves out the bash script, which cannot be run on Windows, and with it the copying of the
> attachments—. The service lives in `dev.yml` **with its own profile**, so **`up -d` does not start it**:
> it is requested on purpose, and **the examples are optional** (the schema is still mandatory). The
> container runs **the usual script, `scripts/dev-seed-en.sh` or `scripts/dev-seed-es.sh`** —a single truth, and the script is still
> valid on Linux and macOS outside the container—: the script uses the `SEED_EN_CONTENEDOR` flag to tell
> whether it talks to the database over the environment network with the image's `psql` or by entering
> the container with `docker compose exec`. Sections 3.1, 3.3, 5 and 9.3 are up to date. Checked on
> 2026-10-01 with **two passes in a row** of the command («Listo: 11 cuentas, 25 tickets y 5 adjuntos»
> both times), with the eleven accounts and the 25 tickets in the database, **the 5 attachments in
> `_files/`** and `up -d` without starting the service.
>
> **Amended on 2026-10-01 (sixth time)**: **the Angular development server forwards `/api` to the
> backend** (`frontend/proxy.conf.json`, declared in `serve.options.proxyConfig` of
> `frontend/angular.json`), with destination **`http://backend:11002`** —the backend service name
> inside the environment network— and **without rewriting the path**. It is a `serve` option: **the
> production `build` does not change** and **in production nginx still forwards**. Section 3.1 says
> so: **with the environment up and the schema applied you log in at `http://127.0.0.1:11001` with
> `admin`/`admin`**, without nginx in front. Checked on 2026-10-01 through port 11001 (`GET
> /api/auth/methods` and `GET /api/health` → **200 JSON**; `POST /api/auth/login` with
> `admin`/`admin` → **200** with token) and with Playwright against `http://127.0.0.1:11001`: **the
> login and logout walkthrough, 14 cases green**.
>
> **Amended on 2026-10-01 (fifth time)**: **the schema and the examples can be applied without `bash`**,
> with Docker commands that work **the same in PowerShell, CMD and bash**: the file is copied to the
> container with `docker compose -f dev.yml cp` and passed to `psql` with `-f` and `-v ON_ERROR_STOP=1`
> (section 3.2, which was the only place that only showed the `<` redirection, a shell one). The
> `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh` script **is stated for what it is: for Linux and macOS**, and
> whoever does not have `bash` applies its two files with those two commands (sections 3.1, 3.3 and 5).
> The local startup steps, with the reason for each one, live in the `README.md`. Checked on
> 2026-10-01: on a new database (`catalina_support_prueba`) the two commands of section 3.2 leave
> **18 tables** and the application responds (`GET /api/auth/methods` and `POST /api/auth/login` with
> `admin`/`admin`, **200**); the same file applied **twice in a row** gives `ERROR lines: 0` and ends
> in `COMMIT`, and the development database, after reapplying it, still has `12 usuarios, 25 tickets y 22 plantillas`.
>
> **Amended on 2026-10-01 (fourth time)**: it is now stated **what applies the schema in development and
> when**. Section 3.3 says what the script always did and the documentation did not tell: **on
> a new database `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh` is what creates the schema** —its step 1, `Aplicando el
> esquema…`, applies `v1.0.0.sql`—, **not only the example data**, and it is run **after bringing the
> environment up and before opening the application**: the application **does not create tables on
> startup** and migrations are not applied by themselves, so opening it first leaves `relación "installation_settings" does
> not exist` and `relación "ai_insights" does not exist`. Section 3.1 notes it in one line, section 5 adds the
> moment in its table and names the symptom, and **the local startup steps are not repeated here**:
> they live in the `README.md`, «Try it locally: empty installation». Checked on the development database: the
> seeder ends with `Listo: 11 cuentas, 25 tickets y 5 adjuntos` on a database that already has the
> schema, and `v1.0.0.sql` applied **twice in a row** gives no error (`ERROR lines: 0` in both
> passes).
>
> **Amended on 2026-10-01 (third time)**: the test directory and Keycloak **leave `dev.yml`** and move to
> **`active-directory.yml`** and **`keycloak.yml`**, each with **its own command**, and **the `auth`
> profile disappears**. Both join the **`catalina-support-dev`** network, which **`dev.yml` owns**
> (fixed name, no `external`), and that is why **the environment is brought up first** and the
> service afterwards. Section 3.3 is up to date (the two services no longer live in `dev.yml`), 9.3
> (the three commands and the new suite command, in two steps) and the case table (they need their
> service up, not a profile).
>
> **Amended on 2026-10-01 (second time)**: section 3.3 **stops listing the eleven example accounts**
> —and repeating their password— and **points to the table in the `README.md`**, which is where they
> are detailed with their **name, email, role and password** («Try it locally: empty installation»):
> the same data in two documents drifts apart with time, so it lives in one and the other references
> it instead of copying it.
>
> **Amended on 2026-10-01**: the **three optional elements** are each brought up with their own
> command and that is how it is stated (section 9.3): the AI engine (`docker compose -f ai.yml up -d`, section
> 3.4), the test directory (`docker compose -f dev.yml up -d ldap`) and Keycloak
> (`docker compose -f dev.yml up -d keycloak`). The directory and Keycloak **share the `auth` profile**
> but **can be brought up separately** by naming the service, **without activating the profile** —checked by
> bringing up only one and seeing in `docker compose -f dev.yml ps` that the other does not appear—; the usual command
> `docker compose -f dev.yml --profile auth up -d` **still brings both up at once** and is the one used
> by the suite with the method in AD or Keycloak and by the interface tests. All three are optional.
>
> **Amended on 2026-09-30 (sixth time)**: **the production deployment is parked** until the product is
> finished, so **this runbook is not rewritten now**: section 4 carries a **note** saying that its flow
> **changed on 2026-09-30** —the paths in `prod.yml` became relative to the file itself,
> `scripts/prod-build.sh` publishes the compose and `config/` into the deployment folder and runs the
> compose from there, and the ports are no longer bound to `127.0.0.1`— and that **it will be brought
> up to date when the deployment is resumed**. What is in the section describes the deployment
> **already done** on 2026-09-25.
>
> **Amended on 2026-09-30 (fifth time, and corrected the same day)**: **the network shared with the AI
> engine is created by the environment**, and the engine **joins it**. The containers of `dev.yml`, `prod.yml` and
> `ai.yml` join the `catalina-support-ai` network. At first it was declared **external in all three**
> —"someone has to create it"— and that **broke startup on a new machine**: `docker compose -f dev.yml up -d` failed with «network catalina-support-ai declares as external, but could not be found»,
> and the README did not say so. It was fixed with a script, and **the owner corrected it**: a `.sh`
> leaves Windows out.
>
> **Now there is no script**: **`dev.yml` and `prod.yml` own the network** —with its fixed name and no `external`—
> and **`ai.yml` declares it external**, because the engine is brought up **afterwards**, with the environment running, and
> **it cannot try to recreate it** (if all three believed they owned it, starting the engine with the environment
> up fails: it was tried and it failed). So `docker compose -f dev.yml up -d` **on its own works on a
> new machine** —verified by taking everything down and bringing up only the development one: it creates it— and the engine can
> be added afterwards —verified with the environment running: it starts healthy and the backend sees it—. It works
> the same on Linux, macOS and Windows.
> **Amended on 2026-09-30 (fourth time)**: the migration **`v1.0.0.sql`** also brings **the installation
> seal** (`installation_settings.installed_at`) and **the outgoing mail** (`smtp_*`), which now
> lives in the database **and stops living in the environment**: the `SMTP_*` variables are removed from the environment and from the
> example files (correction of the owner, 2026-09-30). **An installation that had them set
> enters the mail once** in the installation view or in Settings. **The seal is filled in automatically only in
> installations that already existed** —and only the first time the column is born—, so a running installation
> **does not see the wizard** because of updating, and a freshly created database **is born unsealed**, which is
> "not installed": opening it shows `/setup` (`docs/primer-arranque.md`). In development **the example file
> seals the installation**, so the wizard does not appear on every start or in the tests.
>
> **Amended on 2026-09-29 (third time)**: **`PUBLIC_APP_URL` stops being mandatory**: the public address
> of the installation moves to **Settings** (`docs/modules/settings.md`, decision 14) and the
> variable remains **as a fallback** for an installation that already had it set. A new installation
> no longer needs that variable to start; without a configured address, **emails with a link fail
> with their key** instead of sending a broken link. The two fields —**the time zone** and **the public address**— go **in `v1.0.0.sql` itself**: the
> version is still open at **1.0.0** and no new file is created for this.
>
> **Amended on 2026-09-28 (second time)** — **owner's decision** on seeing the environment full of
> test data: **after each pass of the interface layer `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh` is run**,
> so that development keeps the eleven example accounts and the 25 tickets and **nobody finds
> test tickets in the lists**. The tickets created by the suite **cannot be deleted from it** —a
> ticket is not deleted, and the test container does not talk to the database—, so the restart is done by
> the seeder. **A warning is given before each pass**: the seeder **also deletes whatever is at hand in
> development**, and the owner has accepted that with the cost stated.
>
> **Amended on 2026-09-28**: the interface layer **turns off its accounts when it finishes**
> (`tests/e2e/limpiar.ts`, the closing of `playwright.config.ts`, section 9.3). Each case creates the
> accounts it needs, and **they were left on in development**: the owner found
> **six hundred active support and development people, all named «Prueba Automática»**, in the
> "Assign to" dropdown and in the search for whom to tag. Now they are turned off —**not deleted**,
> which is the rule for accounts— by asking **through the API**, which is what the application uses: the
> test container does not talk to the database. The **tickets** left by the pass are not touched —a
> ticket is not deleted—: the environment is left as new with `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh`.
>
> **Amended on 2026-09-27 (second time)**: the interface layer **returns the installation to local accounts
> before starting** (`tests/e2e/preparar.ts`, the prior preparation of `playwright.config.ts`),
> so that a pass cut halfway with the method in AD or Keycloak does not make the next one fail with
> "invalid credentials" (section 9.3). The **factory account** does it, and it always logs in.
>
> **Amended on 2026-09-27**, with **the AI engine** entering (`docs/modules/ai.md`): a new container,
> **in its own compose (`ai.yml`)** and **shared by development and production**, with the network
> `catalina-support-ai` that both backends join (section 3.4). **It is not mandatory**: without it,
> the two fields it writes are left without text and everything else works the same. And the example data
> file **also deletes the summaries**, which go by ticket number and would stick to the new
> tickets (section 3.3).
>
> **Amended on 2026-09-26**, when checking the update of the migration with the
> attachments inside the text: the condition that protects it **was not the one needed**. It said "does not
> start with `<`", and since the editor writes a body **that starts with its text and carries tags
> inside**, reapplying the file **escaped the attachment references and left them as literal text**
> (section 5). Corrected to **"has neither `<` nor `&`"**, which is what only the plain text from before
> satisfies, and **checked by applying the whole file twice** on development: `UPDATE 0` in both
> conversions.
>
> **Amended on 2026-09-26**, when the owner changed **the domain of the example addresses**: the seeder
> accounts, the people in the test directory and those in the realm go from `@ejemplo.com` to
> **`@demo.com`** (section 3.3). **It is a development change and nothing more**: the seeder is not
> applied in production, the directory and realm people belong to the test services, and **the mail
> sender does not change** (in production it must be a domain of the house, with its SPF and its DKIM, or the
> mail does not arrive). The addresses filled in by the unit tests are not touched: there the domain
> does not matter.
>
> **Amended on 2026-09-25**, when doing **the first production deployment**:
> `scripts/prod-build.sh` **is written** and is the one that builds and publishes (section 4.2), the
> **production containers are up and verified from the inside** —health, brand, schema applied
> **only once** and the factory account logging in with the password from `config/env/prod.env`—,
> and **the domain still responds 503** until it is opened (section 4.5). **Outgoing mail remains
> pending**, and with it account creation: it is what is most missing. In addition, the backup of
> the **production** database enters the `cron` (section 6) and **a bug in the backup script is fixed**: its production branch
> looked for `prod.yml` and the environment among the artifacts, and they live in the repository.
>
> **Becomes as-built on 2026-09-25**, and **amended the same day**: the **organization directory and
> Keycloak stop being environment variables** and are configured from the screen —they live in the database
> and their secrets go inside the database backup, section 3.2—, and **the example data leaves
> both paths set with the method in `local`** (section 3.3). The production deployment is still
> what is missing from this document.
>
> Approved by the owner on 2026-09-22, after reviewing it in the form of questions while it was being
> written. It fixes the deployment (a script with the same shape as Calibyou's), the migrations
> —one file per version, with three digits—, the backups and the three test layers.
>
> With this document, **the documentation chain is complete**: there are no more documents planned, and
> whatever is missing is decided in the document of the corresponding area.

## 1. Scope of this document

It tells **how it is deployed and how it is tested**: the production runbook, the migrations, the backups
and the three test layers.

**It does not describe the architecture or the containers**: that is in `docs/arquitectura.md` (sections
2, 10 and 11, with the ports, the images and the vhosts). It is not repeated here, and when something
changes in one place it is changed there.

It is **approved**, so it enables writing code (Rule 0). Section 10
collects what I proposed.

## 2. The two environments at a glance

| | Development | Production |
| --- | --- | --- |
| **Domain** | `dev.catalina-support.example.com` | `support.example.com` |
| **Ports** | frontend 11001, backend 11002, database 11003 | frontend 21001, backend 21002, database 21003 |
| **Code** | Mounted as a volume: it is edited and the container reloads | **Built artifacts** in `/srv/catalina-support` |
| **Reload** | Angular and `air` reload on save | It has to be deployed |
| **Database** | Compose volume, it can be wiped without thinking | Real data: **the migration is applied only once**, when 1.0.0 closes |
| **Status today** | **Up and verified** | **Postponed** until 1.0.0 closes (see `docs/arquitectura.md`, section 13) |

While the deployment is postponed, the production domain responds 503 with a notice that it is
under development. It is not a failure: it is what was decided.

## 3. Development

### 3.1 Bring up, look and stop

```bash
docker compose -f dev.yml up -d
docker compose -f dev.yml ps
docker compose -f dev.yml logs -f backend
docker compose -f dev.yml down          # los volúmenes se conservan
```

Login is at `https://dev.catalina-support.example.com`. **And also, without nginx, at
`http://127.0.0.1:11001`**: the development server forwards `/api` to the backend
(`frontend/proxy.conf.json` → `http://backend:11002`), so with the environment up **port 11001 opens**. `migrate` waits for PostgreSQL, applies the schema and `backend` waits for it to finish correctly.
On a new database `/setup` appears; after completing the
wizard you log in with **`admin`/`admin`**. The examples are optional and seal the installation
(`README.md`, «Try it locally: empty installation»).
Ports 11001 and 11002 remain published to debug without going through nginx.

**Bringing the environment up leaves the database ready**: the application does not create tables on its own; the
single-use `migrate` service applies them before Compose allows `backend` to start. If the SQL fails,
the backend is left stopped and the diagnosis lives in `docker compose -f dev.yml logs migrate`. The
example data also applies the schema in its first step, but it is still optional.

### 3.2 The database

- The schema is applied **repeatedly** without fear: `v1.0.0.sql` is transactional and idempotent
  (`docs/arquitectura.md`, section 7). **The file exists, is applied and creates 18 tables.**
- `docker compose -f dev.yml up -d --build` runs the single-use service automatically:
  it waits for PostgreSQL's health, uses its client in Docker and ends with an error if the SQL fails.
  `backend` depends on it finishing correctly.
- It can also be repeated manually; it does not delete data, does not load examples and does not reopen the wizard:

```bash
docker compose -f dev.yml run --rm migrate
```

- On Linux and macOS, moreover, it can be done without copying, with the shell redirection —which **does not
  exist in PowerShell** and that is why it is not the form above—:

```bash
docker compose -f dev.yml exec -T database \
  psql -v ON_ERROR_STOP=1 -U catalina_support -d catalina_support -p 11003 \
  < backend/migrations/v1.0.0.sql
```

- **Emptying the development database** is as simple as deleting the volume:

```bash
docker compose -f dev.yml down
docker volume rm catalina_support_dev_database
docker compose -f dev.yml up -d
```

The next `up -d` applies the migration automatically again. This deletes the data and the local
configuration: do it only if you want to start another installation.

### 3.3 The example data

**The example data is applied on purpose**, after bringing the environment up. **It is not a
requirement to log in** —`up` already leaves the schema ready, the wizard is completed and then you log in
with `admin`—, but it is what leaves the environment with content. The application does not create the tables: the
`migrate` service of
Compose does it and finishes before the backend starts (section 3.2). The local startup steps
are in the `README.md`, «Try it locally: empty installation».

**The form on the three systems —Linux, macOS and Windows— is one Docker command**: a container
that starts, seeds and leaves.

```bash
docker compose -f dev.yml run --rm seed-en|seed-es
```

**It does not start with `up -d`**, because the `seed` service of `dev.yml` has **its own profile**: it is requested
on purpose, and **whoever wants the empty project —only with the schema— does not run this command**.
**The examples are a choice; the schema is mandatory.**

The container runs **the usual script, `scripts/dev-seed-en.sh` or `scripts/dev-seed-es.sh`** —a single truth about how it is
seeded—: it applies the schema, applies `v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql` and **copies the attachment files** to
`_files/`, which is the step an SQL seeder cannot do. Its mounts are **read-only** except
`_files/`, which is the only thing the script writes on the machine.

**On Linux and macOS the same script still exists outside the container**, and it does the same:

```bash
./scripts/dev-seed-en.sh  # o ./scripts/dev-seed-es.sh
```

The script talks to the database **by entering the container with `docker compose exec`** when it runs on
the machine, and **with the image's `psql` over the environment network** when it runs inside the container:
its `SEED_EN_CONTENEDOR` flag distinguishes it, and that is why there are not two ways of seeding that could
drift apart.

It leaves **eleven accounts, 25 tickets and 5 attachments** —tested on 2026-10-01 with **two passes in a row** of the
Docker command («Listo: 11 cuentas, 25 tickets y 5 adjuntos» both times), also checking in
the database the eleven accounts and the 25 tickets and in `_files/` the five files—. The `migrate` service of
section 3.2 applies the schema, but **they do not copy the attachments**: that is what the
container is for, which is step 3 of the script.

Its **step 1** (`Aplicando el esquema…`) applies `backend/migrations/v1.0.0.sql` —transactional and
idempotent, so on a database that already has it it can be run again without harm—, **step 2**
applies `v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql` and **step 3** **copies the attachment files** from `config/seed/` to
`_files/`, in the folder of its ticket: an SQL seeder cannot create files, and an attachment without its
file cannot be opened.

It leaves **eleven accounts** —**the table with their name, email, role and password is in the
`README.md`**, in «Try it locally: empty installation», and **it is not repeated here**— and **25 tickets**
spread across all states, with assignments, reassignments, comments from the three roles,
**seven with an internal ticket** (one waiting for Development, another returned to Support, others resolved or
closed), real attachments and **the dates spread over the last three months**, so that the
inbox does not come out all from the same day.

**The addresses are `@demo.com`** (owner's decision, 2026-09-26), like those of the
people in the test directory and the realm: in development everything is fake, and a single domain
—the same in all three places— is what means you do not have to think about which one applies.

**The eleven log in with the same password**, the one in the `README.md` table (owner's decision,
2026-09-25): in development you log in and out many times a day, and a password that
is typed with one hand is what keeps testing from becoming a chore.

- **It complies with the product policy**, which asks for **8 characters** minimum —lowered from 12 on
  2026-09-25—, so the seeder **does not skip any rule**: it writes the `bcrypt` already made because an
  SQL file cannot call the application, and the result is the same as if the password had
  been set from the screen.
- **The interface tests do not use these accounts**: they create their own with the password set by
  `tests/e2e/ayudas.ts`. Changing the one for the eleven does not touch any test.

**It deletes to leave a known state**: any tickets there might be —including those created by hand
while testing—, **all the accounts** except the eleven examples and the attachments that are not the brand. It is a
development seeder, not a migration, and that is why it is applied when wanted and not at startup.

**The demo is born with the local AI configured** (`mode=local`, engine URL from `ai.yml` and model
`qwen2.5-1.5b-instruct`) when the installation is not sealed yet. Without that, the seed left
`aiRequired` active with an empty `ai_settings` and users only saw «No permission» until an
administrator configured the AI. The engine may not be running: in that case the summaries say
«No engine», but the rest of the application is used normally. An already sealed installation
keeps its configuration.

**Only when the installation is not sealed**, it leaves the two directory paths configured
and activates the `local` method (approved amendment of 2026-10-03): it writes the directory configuration pointing to the `ldap` brought up by
`active-directory.yml` and the Keycloak one pointing to the test realm, **leaving the local method**. A sealed installation keeps its settings. Thus
the environment starts by logging in with the eleven example accounts and **changing method is one click in
Settings**, which is exactly what has to be testable by hand. The two test services live
each in its own file (`active-directory.yml` and `keycloak.yml`): with them up, «Test
connection» answers yes; without them, it answers no, which is the truth.

- **The cleanup deletes all accounts** before creating the eleven examples: that way the
  `solicitante-…` or `ad-alta-…` from previous tests do not survive, which dirtied the list and the search. The
  sealed installation keeps its configuration; the accounts are only development data.

### 3.4 The local AI engine: its own compose and resources per environment

The engine that writes the **reason** and the **last action** of tickets (`docs/modules/ai.md`) **does not
live in `dev.yml` or in `prod.yml`**, but in **`ai.yml`**. It is only brought up when the local
mode is chosen:

```bash
docker compose -f ai.yml up -d --build
docker compose -f ai.yml ps
docker compose -f ai.yml logs -f ai
docker compose -f ai.yml down       # el volumen del modelo se queda
```

**Why separate**: the administrator downloads, activates and supervises `llama-server` without accessing the Docker
socket. `dev.yml` creates `catalina-support-ai-dev` and `prod.yml` creates
`catalina-support-ai-prod`; `ai.yml` joins the corresponding network through `AI_NETWORK_NAME`.
Its projects and volumes also carry a name per environment (`AI_PROJECT_NAME` and
`AI_MODELS_VOLUME`), so development and production do not share models or secrets. The
backend reaches it by service name and the administrator does not publish ports to the
machine.

- **The model is downloaded from `/setup` or `/settings`** to the environment volume; the download can be
  resumed and only `down -v` deletes the volume. The interface offers the 1.5B, 3B and 7B models with their
  disk and RAM estimates.
- **AI is mandatory to install**, but the local mode is not: a self-hosted server or an external provider can also be chosen. If the configured engine falls afterwards, the summaries
  remain pending and the help desk keeps working.
- **The network is created by the environment**, so **the engine can be brought up later, with the environment
  running**: `docker compose -f dev.yml up -d` works on a new machine without creating anything by hand.
- Production operation remains parked until the 1.0.0 version closes. When it is resumed,
  Compose will run with the production names and secrets documented in section 13.

### 3.5 Where what is not code lives

| Folder | What it stores | In git? |
| --- | --- | --- |
| `_logs` | The `go-logs` files, one file per day | The folder yes, the `.log` files no |
| `_files` | The ticket attachments | The folder yes, the files no |

Both are volumes mounted in the backend: they survive `down` and are copied from the machine without
entering the container.

## 4. Production

> **Note (2026-09-30):** this flow **changed** —the volume paths in `prod.yml` became
> relative to the file itself, `scripts/prod-build.sh` publishes the compose and `config/` into the
> deployment folder and the deployment is run from there, and the ports are no longer bound to `127.0.0.1`—, but
> **the section is not brought up to date yet**: the production deployment **remains parked** until the
> product is finished. What is read below describes the deployment **already done** on 2026-09-25, not the
> state of today. It will be brought up to date when the deployment is resumed.

### 4.1 How it is put together

«Thin» containers: the runtime lives in base images and **the code is built artifacts** in
`/srv/catalina-support` (frontend, backend, `_logs`, `_files`). The `builder_*` of `prod.yml`
exist to build and publish, and **they are not brought up with `up`**: they are under the `build` profile.

### 4.2 The deployment, step by step

> **Finding of 2026-10-02:** the following steps keep the record of the historical deployment;
> they are not today a reproducible walkthrough from scratch. `prod-build.sh` uses the compose published in the
> artifacts folder to build, but it does not copy there the sources the Dockerfiles ask for.
> The output path remains fixed at `/srv/catalina-support`. The behavior fix
> requires an approved proposal and production remains parked.


A script, `scripts/prod-build.sh`, with the same script as Calibyou's. **It is written and tested**
(on 2026-09-25 the two artifacts were built and published with it). What it does:

1. **Warns if the working tree has uncommitted changes**: an artifact built from a dirty
   tree does not correspond to any commit, and that has to be known before and not after.
2. **Rotates the previous artifact** of each component to `<name>_prev`, **by copy and not by
   move**, because the directory is mounted in the containers: renaming it would leave them
   pointing to the old inode.
3. **Builds the builder image and publishes** with
   `docker compose -f prod.yml --profile build build builder_<name>` and then
   `… --profile build run --rm builder_<name>`.
   - **Both steps are mandatory, and the first is the one that gets forgotten**: `run` only builds the
     image if it **does not exist**, so from the second deployment on it would reuse the one from the previous
     time and **would publish the old artifact without saying anything**. It was discovered on 2026-09-26, when
     replacing the factory logos: the published artifact still brought the old ones, and the
     closed domain (503) was what prevented it from being noticed in production.
4. **Verifies that the artifact exists** before continuing: `frontend/index.html` and
   `backend/catalina-support` executable. If missing, it stops without deploying.
5. **Copies the third-party notices**: `THIRD_PARTY_NOTICES.md` and `third_party_licenses/` are left
   inside each built artifact. This copy is implemented and its real check remains
   for the final production review, which is still parked.
6. **Writes `BUILD_INFO`** in `/srv/catalina-support` with the commit, the branch, whether the tree
   was dirty, the date and which artifacts were built. It is the answer to "what is deployed
   right now?" without guessing.
7. **Deploys**: `docker compose -f prod.yml up -d` and **restarts the backend**. The frontend is not
   restarted: the nginx container reads the static files from disk on every request.

```bash
./scripts/prod-build.sh                 # construye, publica y despliega
./scripts/prod-build.sh --no-deploy     # sólo construir y publicar
./scripts/prod-build.sh --only backend  # un solo componente
```

- **An artifact built from a dirty tree does not correspond to any commit**, and that is why the
  script warns **and writes it in `BUILD_INFO`**: the first deployment was done that way, with the tree
  uncommitted, and that is stated there so that nobody confuses it with the published version.
- **The `prod.yml` and the production environment live in the repository**, and what is in
  `/srv/catalina-support` is the built code. That is what `prod.yml` mounts, and it is what
  any script that talks to production has to know (it was fixed that way in
  `scripts/backup-db.sh`, section 6).

### 4.3 After deploying

1. **Apply the migrations** that apply (section 5).
2. **Check that it started**, in this order and without skipping any:

```bash
docker compose -f prod.yml ps                                  # los tres arriba y la base sana
curl -s https://support.example.com/api/health       # {"database":"ok","status":"ok"}
curl -sI https://support.example.com/ | head -1      # 200 y la CSP de producción
```

3. **Look at the logs** of the backend, which is where `go-logs` tells what has happened:

```bash
docker compose -f prod.yml logs --tail=50 backend
tail -f /srv/catalina-support/_logs/catalina-support_prod_$(date +%Y%m%d).log
```

### 4.4 Rolling back

If the deployment goes wrong, **the previous artifact is still there**: `<name>_prev` is copied over
`<name>`, the backend is restarted and you go back to the previous version. It is automated in
section 4.2 (the rotation) and manually in the worst case:

```bash
cp -a /srv/catalina-support/backend_prev/. /srv/catalina-support/backend/
docker compose -f prod.yml restart backend
```

**What cannot be undone with this is a migration**: once applied, going back requires
writing another one that reverts it. That is why the migration is applied **after** checking that the new
code starts, and not before.

### 4.5 The first time: done on 2026-09-25, and what is missing

It is a normal deployment, with three things that only happen once. **Two are done and verified**:

| Step | Status |
| --- | --- |
| **1. The migration `v1.0.0.sql`, only once** | **Done**: 13 tables, the 20 mail templates and the settings row with the sign-in method in `local`, and **zero accounts** —in production none are seeded—. In development it has been applied many times; in production, that one and no more |
| **2. The `admin` account logs in** with the password from `config/env/prod.env` | **Done and checked**: it logs in with the `administrador` role, and with the wrong password it answers `auth.invalidCredentials`. The password lives **only** in that file, which is not versioned (`docs/usuarios-y-permisos.md`, section 8) |
| **3. Create Support and Development accounts** from the application | **Pending, and it depends on mail**: account creation sends a link by email, and **production mail is unconfigured**. While it stays that way, **nobody can set their password** and the only account that logs in is the factory one |

**Outgoing mail is configured from the application**, not in `config/env/prod.env`: in step 4
of `/setup` if the installation is not sealed yet, or in **Settings → Mail** with the factory
account if it already is. The data is tested with «Send test email» and saved in
`installation_settings.smtp_*`; there is no need to restart the backend. `prod.env` still contains
`ADMIN_PASSWORD`, the AI key and the PostgreSQL data, but **no `SMTP_*` variable**.
As soon as there is mail, step 3 is done from the application in a minute.

**The vhost and the certificate are already in place** since 2026-09-22, and **the domain still responds
503** on purpose: 1.0.0 is not closed yet and the owner may change things before
closing it. **Opening it is a two-line change** in `config/nginx/catalina-support-prod.conf` —the
notice block is replaced by the `proxy_pass` to the frontend, which is commented right below—,
copy it to `/etc/nginx/conf.d/` and reload. What is gained and what is lost is clear: the domain
starts showing the real application, with mail still pending.

## 5. The migrations

**One file per version**, and the name is the version that is published: **three digits**,
`v1.0.0.sql`, `v1.1.0.sql`, `v1.2.0.sql`… The file and the git tag that marks that version
carry the same number, so by looking at the tag you know which migration applies.

**And an example data file per version, with `_dev` behind the name** (`v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql`):
that is where the accounts and tickets you work with in development go. **That file is never applied
in production**, and that is why it carries the suffix clearly visible.

| Moment | What is applied |
| --- | --- |
| **In development, at the start** | `docker compose -f dev.yml up -d --build` applies `v1.0.0.sql` automatically before the backend. `run --rm migrate` allows repeating it. The examples are **separate** (section 3.3) |
| **Until 1.0.0** | `backend/migrations/v1.0.0.sql`, as many times as needed in development |
| **When closing 1.0.0** | The same file, **only once** in production |
| **After 1.0.0** | One file per version, applied **in order**, without skipping any |

**One file per version and not one per change**: the file tells the complete history of that
version, is reviewed at once and applied at once. One file per change makes the
deployment piecemeal and fills `migrations/` with scraps.

**Each version brings its schema file and, if needed, its examples file**, and each one says what it is for:

| File | What it carries | Where it is applied |
| --- | --- | --- |
| `v1.0.0.sql` | The tables, the indexes and their constraints **and what the application needs to work**: the two settings rows —the installation one and the tickets one— and the twenty mail templates. Without the settings row neither the language nor the name can be read; without a template no mail goes out | **Development and production** |
| `v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql` | The **example data**: eleven accounts (five requesters, three Support and three Development), 25 tickets with their history —assignments, reassignments, comments, escalations with their internal ticket, attachments— and the numbering counter | **Only development**, with `docker compose -f dev.yml run --rm seed-en|seed-es` (`scripts/dev-seed-en.sh` or `scripts/dev-seed-es.sh` on Linux and macOS) |

**The script also deletes the AI engine summaries** (`ai_insights`): the table is kept **by ticket
number**, and the example data **re-issues the same numbers**, so an old row would
stick to a new ticket and show the reason of another (`docs/modules/ai.md`, section 4).

**The separation is not for form**: what is in `_dev` **deletes and recreates** the working
content —tickets included—, so it cannot be something applied just by deploying. In
development it is applied when you want to return to a known environment.

**A migration can also change data, and this one does**: on 2026-09-26 the text of
tickets and comments **went from plain text to HTML** (`docs/modules/tickets.md`, section 2.3),
so `v1.0.0.sql` **converts what already exists** in its update section: it escapes the
characters HTML eats (`&`, `<`, `>`), changes line breaks for `<br>` and **wraps the
result in a paragraph** (`<p>…</p>`).
   - **The condition is "has neither `<` nor `&`", and not "does not start with `<`"** (corrected on
     2026-09-26, in the same work that put the attachments inside the text). With the old condition
     —"this text is not HTML yet", asked by its first letter—, a body written by
     the editor **started with its text and carried tags inside** («Mira lo que me
     sale.<img data-adjunto="captura.png">»), so it came back in, was escaped whole and the
     attachment references were left as literal text: **applying the file twice broke
     what the editor writes**, which is the opposite of what this section promises. Now only
     what can be plain text from before is converted, because an editor text **always carries `&`
     or `<`** —it escapes `&` and `<` when saving them, and inserts tags as soon as there is formatting or an
     attachment— and an already converted text starts with `<p>`.
     - What the condition leaves out, and it is said: an old text that **carried a `<` or an `&` by
       hand** is not touched and is read without formatting. **The `<p>` is still there** —it is how a text is read as
       a paragraph—, but it is no longer what sustains idempotence.
     - **Checked by applying the whole file twice** on development: the second pass leaves
       both conversions at **`UPDATE 0`**, and a body with an image inside —the one the
       editor writes— **is not touched**. It is the check that has to be repeated if this section is touched
       again.
   It is a
one-way conversion —the text is read the same as before, now with formatting—, it is applied
on its own when applying the file and **in production it touches nothing**, because there is no ticket there yet.
What has to be kept in mind is that **backups before that day have the plain text**: an old
restore would look the same, without formatting, and the migration has to be applied again to bring it
up to date.

**With a control table**: `schema_migrations` records version, file, SHA-256 and date. The files
are still idempotent, but only the pending ones are applied; the checksum prevents changing a
published version. The exception is the current version while `Development=true`, which can be
reapplied while it is being built.

And to know which version is published there is no need to guess: **the git tag and the
`BUILD_INFO`** of the deployment say it (section 4.2). Applying too much is harmless; applying too little is seen
right away, because the application will ask for a table that does not exist —`installation_settings`, `ai_insights`…—
and the database will answer `relación "…" does not exist` (section 3.3).

Never `AutoMigrate`, never a manual `ALTER`, never a step that is not in a file in the
repository.

## 6. The backups

**What cannot be rebuilt is the database**, so that is what is backed up on its own: one dump
a day, with **fourteen days of retention**, in the server's `cron`.

```bash
./scripts/backup-db.sh            # el entorno de desarrollo
./scripts/backup-db.sh prod       # producción
```

What it does, in this order (and if anything fails, **it does not leave a half-made backup**: it deletes what it has
written and exits with an error):

1. **Checks that the database is up.** Without it there is nothing to copy.
2. **Dumps the database** with `pg_dump -Fc` **inside the container** —PostgreSQL's compressed
   format, which is restored with `pg_restore`— and leaves it in the backups folder with the date in the
   name: `/srv/catalina-support/backups/catalina_support_20260925_033000.dump`.
3. **Checks that the dump can be read** (`pg_restore --list`). A backup that cannot be
   restored is not a backup, and that has to be known today and not the day it is needed.
4. **Deletes backups older than fourteen days** and says how many remain.

**It does not copy `_files`**, and it is on purpose (correction of the owner, 2026-09-25): **the attachments of
the tickets and the documents people upload are never deleted**, and copying them is a matter of manual
support when needed. The script takes care of the database and nothing else. When the files have to be taken
away —to another server, to a disk, to a review—, they are copied from the machine like any other
folder, respecting what is inside:

```bash
cp -a /srv/catalina-support/source/_files/. /srv/catalina-support/backups/_files/
```

**And a warning that is worth having in front of you on the day of a real restore**: in the database go the
**rows** of the attachments, and the files go separately. Restoring only the database leaves an installation with
attachments that **cannot be opened**, because their file is not there. When a real rollback is needed,
both things are restored: the database **and** `_files`.

**And the second warning, confirmed by the owner on 2026-09-25**: since the organization
directory and Keycloak are configured from the screen, **their secrets —the service account
password and the client secret— live in the database**, so **a database backup is one more
secret**: whoever has it can talk to the directory and to the realm. It is the price of being able to change them
without going in through SSH, and that is why the backups live where they live, with the machine's permissions and without leaving
it.

**Where the backups live**: `/srv/catalina-support/backups/`, outside the repository and with the
backups of the other projects on the machine. The script writes what it does to standard output, and
the `cron` collects it in `copias.log`, right there:

```cron
30 3 * * * /srv/catalina-support/source/scripts/backup-db.sh dev  >> /srv/catalina-support/backups/copias.log 2>&1
45 3 * * * /srv/catalina-support/source/scripts/backup-db.sh prod >> /srv/catalina-support/backups/copias.log 2>&1
```

**Both databases are backed up**, and at different times so they do not compete: development at 3:30 and
**production at 3:45**. Production was backed up and checked on 2026-09-25 —the dump can
be read and carries its thirteen tables—, and since that day it is in the `cron`.

**The backups of the two databases have the same name**, because the database has the same name in both environments
(`catalina_support`): what distinguishes them is the date and time in the name. If some day you have to
know which one a backup is from, look at the content —production's has no accounts and development's
has the eleven examples—. It is stated here **instead of inventing a different name**: renaming the
file would also force touching the restore, and it adds nothing that the content does not say.

**Restoring**, when needed: a database is created and the dump is loaded into it. With `--no-owner`, because the
owner of the tables of a backup does not have to be the owner of the database where it is restored.

```bash
docker compose -f dev.yml exec database psql -U catalina_support -d postgres -p 11003 \
  -c "CREATE DATABASE restauracion"
docker compose -f dev.yml exec database pg_restore -U catalina_support -d restauracion -p 11003 \
  --no-owner < /srv/catalina-support/backups/catalina_support_20260925_033000.dump
```

**What is tested and what is not.** On 2026-09-25 it was checked, against the development environment: that
the script leaves its backup, that the backup **is restored** into a new database —**11 tables and the same
records** as the real one, counted one by one— and that retention deletes what is older than fourteen
days without touching the rest. **What is not tested is in production**, because there is no production: the
script has its branch for `prod` written and never used, and on deployment day it is tested there.

## 7. The environment variables

Each group is documented by whoever uses it, and here it is only said where it lives:

| Group | Who documents it |
| --- | --- |
| `ENVIRONMENT`, `PROJECT_NAME`, `LOGS_FOLDER`, `TIMEZONE` (the `go-logs` ones) | `docs/arquitectura.md`, section 8 |
| `APP_PORT`, `POSTGRES_HOST`, `PGPORT`, `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` | `docs/arquitectura.md`, section 10 |
| `FILES_PATH` | `docs/modules/tickets.md`, section 2.3 |
| `ADMIN_PASSWORD` | `docs/usuarios-y-permisos.md`, section 8 |
| `TOKEN_SECRET` | `docs/usuarios-y-permisos.md`, section 6 |
| `KEYCLOAK_ADMIN_PASSWORD` | The password of the **development Keycloak** administrator (only `dev.yml` reads it, not the backend) |

**And two things that stopped being environment variables on 2026-09-25**: the **organization
directory** and **Keycloak** (`LDAP_*` and `OIDC_*`). Now they live in the database —`directory_settings` and
`keycloak_settings`— and are configured from the Settings screen, with its «Test
connection» button (`docs/modules/settings.md`, section 5.8). **None remains per environment for those two
paths**: if `LDAP_*` or `OIDC_*` are seen in an environment file, they are leftovers from a previous version
and nobody reads them. And a consequence that has to be kept in mind: **the secrets of the two paths
go inside the database backup**, so that backup is one more secret (section 6).

**And outgoing mail is not an environment variable either**: it lives in `installation_settings.smtp_*` and
is configured in the first-start view or in Settings; the `SMTP_*` variables were removed
(`docs/arquitectura.md`, section 9). In development it is left set by
`backend/migrations/v1.0.0_dev_en.sql` or `backend/migrations/v1.0.0_dev_es.sql`, pointing to the test mailbox of `dev.yml` (`mail`, web at
`127.0.0.1:11004`).

Rules that apply to both environments:

- **Only what changes per environment is a variable.** What is worth the same in both is written as a
  constant in the code.
- **`config/env/dev.env` is versioned** because it has no real secrets: they are credentials of
  local containers.
- **`config/env/prod.env` is not versioned**, and none of its values appears in the documentation or
  in the code. There is a template (`prod.env.example`) with the list of what is needed.
- **If a variable stops having a consumer, it is deleted**: from the code, from the envs, from the compose files and from
  the documentation, in the same change.
- **Adding a mandatory variable is not enough to write it in the env**: the containers read their
  environment **when they are created**, so one that has been up for hours does not have it and the backend will refuse to
  start —which is what it should do—. The service has to be recreated, not just restarted:

```bash
docker compose -f dev.yml up -d backend      # recrea el contenedor con el entorno nuevo
```

  Checked on 2026-09-23 when adding `TOKEN_SECRET`: with the old container, the backend answered
  «no se puede firmar la sesión: falta TOKEN_SECRET» and nginx returned 502 until it was recreated.
- **The same applies to the frontend dependencies**: the Angular development server
  **caches the Tailwind result**, so one that had been up before installing
  `tailwindcss` keeps serving the sheet **without any generated utility** —the page looks like
  text without styles— until it is restarted. It happened on 2026-09-23, with the server having been up for a day.
  If the interface looks unstyled after touching dependencies or the
  Tailwind configuration:

```bash
docker compose -f dev.yml restart frontend
```

## 8. nginx and the certificates

- The vhosts are **copies** of `config/nginx/` in `/etc/nginx/conf.d/`: if one is changed in the
  repository, it has to be copied again and reloaded (the commands are in `AGENTS.md`).
- The certificates **renew themselves** (certbot with the `webroot` plugin, one certificate per
  domain). Renewal needs no intervention because both vhosts serve
  `/.well-known/acme-challenge/` on port 80.
- **Check that it is still alive**, from time to time and always before a deployment:

```bash
sudo certbot certificates | grep -A2 catalina
echo | openssl s_client -connect support.example.com:443 \
  -servername support.example.com 2>/dev/null | openssl x509 -noout -dates
```

## 9. The tests

Three layers, from the cheapest to the slowest. All three run **in the containers**.

### 9.1 Backend: `go test ./...`

What is tested, in order of importance:

1. **The ticket state machine**: which transitions exist, who can do each one and, above
   all, **the eight synchronization rules** (`docs/propósito-y-alcance.md`). It is the most
   delicate logic in the system and the one that is hardest to see at a glance: a ticket that stays in `escalado` or
   a main ticket that does not return to Support does not jump out until someone looks at it.
2. **The numbering**: that two tickets created at the same time do not get the same number, and that the number is
   immutable even if the prefix changes.
3. **The permissions**: that each role can only do its own, tested in the services and not only in the
   interface.
4. **The resolution of the origin of an account** on login: `local` against the database, `ad` against the
   directory, and the coexistence rules (`docs/usuarios-y-permisos.md`, section 5).

```bash
docker compose -f dev.yml exec backend go test ./...
docker compose -f dev.yml exec backend go vet ./...
```

### 9.2 Frontend: `npm test`

Vitest, which is what Angular 22 brings. For what has logic: services, guards, and the
components that calculate something. **Appearance is not tested**: that is what the tests of 8.3 are for.

```bash
docker compose -f dev.yml exec frontend npm test -- --watch=false
```

### 9.3 Interface: Playwright in a disposable installation

From the root, without starting or modifying the development environment:

```text
docker compose -f tests.yml --profile directory down -v
docker compose -f tests.yml up -d --build database backend frontend mail
docker compose -f tests.yml run --rm e2e
docker compose -f tests.yml --profile directory down -v
```

To include AD and Keycloak, replace the second command with:

```text
docker compose -f tests.yml --profile directory up -d --build database backend frontend mail ldap keycloak
```

The runner keeps Playwright's exit code. After a failure the final cleanup is also run; the screenshots, traces and report remain in `tests/e2e/resultados/`. Each pass begins by
removing exclusively the volumes of this test project, also after an interruption.
No seeders are loaded: the accounts and fixtures are created through the application's normal APIs.

The Chromium origin is `http://frontend.localhost:11001`: Docker declares the alias and the browser
resolves it to the `frontend` service with `--host-resolver-rules`. The localhost suffix enables the real
clipboard without modifying the application's security. AI remains absent by design.

The suite covers local accounts, mail, settings, permissions and tickets on PC and mobile.
The directories are tested only once because the protocol does not depend on width. The AI cases
require an engine and are skipped in this environment without AI. The preparation also verifies the
wizard before sealing, with screenshots of both widths. The history of the previous findings
remains in the header amendments; their commands against development are no longer the procedure.

Verified on Linux on 2026-10-03: suite with directories, 193 passed, 20 skipped and one
access cancelled during the frontend reload (exit 1, report kept). On a new database,
without directories and with stable files, that case passed on PC and mobile: 2 passed
and 24 skipped due to missing services, exit 0. The detection of a missing LDAP was fixed when
finding that its check function was not called. Go and vet pass; frontend: 216 tests.
These checks have not been run on Windows or macOS.

### 9.4 Isolated backend startup without nested mount

The as-built amendment of `docs/prueba-local.md`, section 13, keeps the commands of this section.
Only how the `backend` service of `tests.yml` starts internally changes: `go run .` instead of the
Air command of the image, without `backend_tmp`. This avoids writing inside the read-only code bind
and does not change the preparation, the profiles, the reports or the cleanup.

The verification always starts from `down -v`, rebuilds the services, waits for the backend's health,
runs at least the affected walkthroughs and ends with another `down -v`. The final compose must
work directly, without temporary override files.

The review confirmed that there will be no reload, temporary volumes or additional caches. There are no
open decisions of the runbook itself.

Verified on Linux on 2026-10-06 with the final compose: from a full cleanup, the
services `database`, `migrate`, `backend`, `frontend` and `mail` started without the OCI error; the suite
without optional profiles ended with 177 cases passed and 39 skipped, exit 0, and the final `down -v`
removed containers, network and volumes. `docker compose -f tests.yml ps -a` was left empty.

## 10. What was decided when reviewing this document

| # | Decision | It stayed like this |
| --- | --- | --- |
| 1 | **The three test layers** | Yes: backend unit tests, frontend unit tests and the five walkthroughs with **Playwright** against development |
| 2 | **The backend test database** | A separate database, `catalina_support_test`, in the same container: the tests empty it and do not touch development data |
| 3 | **Backups** | Daily `pg_dump`, with retention (14 days) and a script in `scripts/`, scheduled in the server's `cron`. **Correction of the owner (2026-09-25)**: the script **does not copy `_files`** —attachments and documents are copied by hand when Support asks for it—, and **attachments are never deleted**. It is in section 6 |
| 4 | **Continuous integration** | **None** in 1.0.0: the tests are run by hand in the containers, as part of the work |
| 5 | **The migrations** | One file per version, three digits, SemVer order, history and checksum in `schema_migrations`; the runner applies pending ones before starting the application |
| 6 | **The example data, separate** | **Owner's decision, 2026-09-25**: 1.0.0 closes with **two files** —`v1.0.0.sql`, with the tables and what the application needs to work in both environments, and `v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql`, with eleven accounts and 25 tickets that **are only applied in development**—. What the application needs to start (the two settings rows and the twenty mail templates) **is not example data** and stays in the first one, because in production there is nobody to seed it another way |
| 7 | **The cleanup of the logs** | **Manual**: `go-logs` does not rotate and no script is added to it |

## 11. What does NOT go into 1.0.0

- **Continuous integration and automatic deployment**: all by hand, on purpose.
- **A staging environment**: development and production, and nothing in between.
- **Monitoring and alerts**: nobody warns that the site has gone down. Finding out is a matter of looking.
- **High availability**: one server, no replicas or load balancing.
- **Deployment without cutting the service**: restarting the backend cuts whatever is in progress, and that is accepted.
- **Backups outside the server**: the script leaves them on the server; taking them out is separate.
- **Automatic log rotation** (neither by size nor by age): cleanup is manual, and the
  `_logs` folder grows one file per day. It is written down so that it does not surprise anyone.
- **Sending logs to a centralized server.**

## 12. What this document enables

With `ambientes.md` approved, **the document chain is complete** and implementation can begin, in the order fixed by `docs/modules/tickets.md`: `auth`, `users` and `tickets`, with their tests. The
`scripts/prod-build.sh` and the backup one are written then, not before.

## 13. Implemented operation for mandatory AI

### 13.1 Variables and secrets

- `AI_CREDENTIAL_KEY`: 32 random bytes encoded in base64, different from `TOKEN_SECRET` and
  mandatory when starting the backend. Development has a versioned local value. Each production
  generates its own, for example with `openssl rand -base64 32`, and stores it in `config/env/prod.env`.
- `AI_MANAGER_TOKEN`: shared secret between backend and local administrator. It is mandatory when
  the local mode is used and it never reaches the browser.
- `AI_URL` and `AI_MODEL` stop being the fallback of an empty configuration: the tested database is the
  single source. `AI_PALABRAS` and the maximum time can be kept as operational limits.

The absence or invalid format of `AI_CREDENTIAL_KEY` prevents the backend from starting. A valid key
that does not decrypt an existing credential leaves the installation in «AI requires configuration» so
that the factory account replaces it. PostgreSQL backups carry the encryption, but not the master
key; both pieces are backed up separately.

### 13.2 Local engine per environment

`ai.yml` is parameterized with the project and the environment network, without global shared names.
Development and production have different model volumes. Bringing it up downloads nothing until
`/setup` or `/settings` requests it; the 1.5B appears selected by default. On a machine that cannot
keep both engines, the runbook says to shut one down before bringing up the other.

The usual operation remains Compose: bring up, see logs and take down without deleting models. The
interface manages files inside the volume, not containers. An incomplete download is left with a
partial name, is resumed and is never activated. Deleting the volume remains a manual and destructive
action outside the application.

### 13.3 New start and upgrade

The main application can start before the local engine to show `/setup`. If local is chosen,
the administrator must be reachable and complete download/activation before sealing. If a provider or
self-hosted server is chosen, `ai.yml` is not brought up.

On an upgrade, the migration keeps an existing configuration that passes the test generation. An
installation without it remains restricted until an Administrator or the factory account configures
it; no credential is invented nor is the local container assumed to exist.

### 13.4 Tests

Unit and integration tests use fake HTTP servers for providers and tiny test GGUF files for the
administrator. Playwright offers a fake service inside `tests.yml`, so that `/setup` can
mandatorily complete the AI step without internet, commercial credentials or gigabyte downloads. It
also tests failure and recovery.

Manual verification of the real catalog downloads and measures each model separately and records size,
RAM, load time and generation in `docs/modules/ai.md` before closing as as-built. It is not tested
against commercial accounts of the project.

Local operation recovers accumulated jobs automatically. A commercial provider requires the
Administrator to confirm each recovery batch after seeing its estimated volume; the tests
verify that no call of the batch happens before that confirmation.

### 13.5 Out of scope

Buying credits, spending limits, GPU, model volume backups and high availability of the engine are not
automated. The GGUF files can be downloaded again; the mandatory backup remains focused on PostgreSQL
and the attachments. The owner decided this scope on 2026-10-06. The review closed generation and
commercial recovery, the mandatory key and license acceptance; no open decisions remain; proposal
explicitly approved.

## 14. Isolated verification implemented for block 4

Unit tests bring up distinct `httptest` servers for OpenAI, DeepSeek, OpenAI-compatible and Claude. They verify route, body, authentication, language, valid response, response too large and rejection of redirects. They do not use public DNS, internet, tokens or commercial accounts.

The local administrator admits substitutes only inside its tests for executable, clock/wait, memory and download. That is how restart with growing wait, cancellation when changing model, cgroup v1/v2, remaining disk, invalid partial, resumption, checksum and redirect limit are verified. The test GGUF files are minimal and temporary.

The supervisor test uses replaceable waits and checks the series 1, 2, 4… 60 seconds without waiting real time. The fake download allows up to five HTTP/HTTPS redirects, rejects the sixth and any other scheme. A file with the correct size and an incorrect checksum is deleted before the next attempt.

`tests.yml` keeps its fake compatible server for `/setup`. Playwright does not simulate process death or download real models; those conditions remain in deterministic Go tests. The closing runs the full backend, `ai-manager`, the frontend tests and Playwright on PC and mobile, and ends with `down -v --remove-orphans`. The counts are noted with the real result. Production remains parked until the 1.0.0 version closes.

The owner closed the review on 2026-10-07 with 5A, 6A, 7A and 8A. No open decisions remain in this amendment.

## 15. Disposable environment used for real models

The validation uses its own names for project, network and volume, plus a disposable database for the browser walkthroughs. It does not change the configuration, the model volume or the data of the usual development. The application and the local administrator run with the same code and images that are being validated.

The walkthrough checks `/setup` and `/settings` on PC and mobile. In `/setup` it downloads, activates and tests the default model before sealing. In `/settings` it verifies catalog, active/healthy states, switching between 1.5B and 3B, generation test and error messages. The 7B is downloaded and validated, but the interface must prevent its activation when the required RAM is not met.

The measurement separates `llama-server` memory and total container memory. The times are taken from the request to health and around each generation request. To test the supervisor only the child process is terminated; the administrator must stay alive and recover the model in no more than five minutes.

On closing, only the report and the documented commands are kept. Containers, network, database, attachments, logs and model volume of the temporary project are removed. A later inventory must show that no resource remains with their names. Development and production remain intact; production stays parked until the 1.0.0 version closes.

## 16. Resumption completed after the memory finding

Unit tests replace RSS, effective memory, process and health readings. They verify that an impossible destination does not stop the origin and that any failure after the stop restores it healthy. Backend and frontend test the structured error and its toast with required/available.

Then only the administrator image of the project `catalina-support-ai-block5` is rebuilt and `catalina_support_ai_models_block5` is reused. 3B is activated, memory, times, twelve generations and recovery are measured. The rejection of 7B is checked keeping 3B healthy and `/setup` or `/settings` are repeated only where the fix changes the result, on PC and mobile.

At the end the approved cleanup is run on the explicit names of block 5. `down -v` is not used on the usual test project. The GGUF files are not downloaded again and production remains parked.

The resumption passed the full suites of `ai-manager` and backend, 220 frontend unit tests and the affected `/settings` walkthrough on PC and mobile. The isolated environment was removed at the end and the final inventory left no containers, networks or volumes with `block5`.

## 17. Isolated verification of block 6

The review uses `tests.yml` and disposable data. `/setup` is walked from an empty installation; afterwards, an instance restarted with the optional seeders provides lists, states, attachments, users and varied tickets for the rest of the screens. Development data is neither read nor modified and production is not used.

The matrix combines Spanish and English, factory light and dark, and 1440×900, 768×1024 and 412×915. The six fixed themes pass their common variable and contrast checks. Chromium is the execution browser; the code still uses compatible HTML and CSS, but Firefox and WebKit are not part of the closing criterion of this block.

The inventory is obtained first, without fixing while walking a screen. Then it is fixed by priority and the affected matrix is repeated. The stable checks enter the component or Playwright tests; the manual review records what cannot be expressed with a reliable assertion. Closing requires frontend unit tests, Playwright on the approved matrix, `git diff --check` and zero open visual or accessibility findings.

The final cleanup uses the explicit names of the disposable project and removes containers, network, database, attachments, logs and caches. No deployment is run, the domain is not opened and the known finding of `prod-build.sh` is not fixed in this block.

The real run used the Compose project `catalina-support-block6`. It first walked `/setup` on an empty database and then applied the official seeder: 11 accounts, 25 tickets and 5 attachments. The reproducible audit found and allowed fixing a single overflow in `/users`; the affected repeat and the 220 unit tests passed. The canonical run ended with **180 E2E cases passed and 38 expected skips**. The isolated project was removed with its volumes and production remained untouched.

## 18. Block 8 proposal: environment for the public walkthrough

The test uses the names that come with `dev.yml`, `ai.yml`, `active-directory.yml` and `keycloak.yml`, because changing names or mounting an alternative compose would stop validating the README. Therefore it is not run in parallel with the usual development.

Before starting, the inventory of containers, networks, volumes and images of Catalina Support is recorded.
The known compose files are taken down with their volume-removal options; the local model volume is also removed. `_files` and `_logs` of development are disposable in this block and are emptied. The cleanup is limited to explicit names and paths of the project: it never uses broad prefixes nor touches resources of other applications.

The temporary clone comes from the public URL and is pinned to the commit that `main` publishes at the start. The only changes admitted inside that folder are expected effects of the documented commands: log directories, attachments and ignored artifacts. Any manual edit of the clone invalidates the pass and forces it to be repeated.

The evidence records versions, commit, commands, active and automatic times, health states, counts and the inventory before and after. It keeps no secrets or the model. On closing, the four compose files are taken down, volumes, own images and the temporary folder are removed, and it is verified that no resource created by the walkthrough remains. The previous development is not restored: its data was declared disposable by the owner.

The original pass validated Linux, Docker Compose, access through `127.0.0.1:11001`, Mailpit and the local engine. Production, nginx, certificates, domain and external providers remain out. The verification of 2026-11-05 extended coverage to **Windows with Docker Desktop WSL2**, including the NVIDIA GPU; macOS remains out.

## 19. Proposal of environment to repeat block 8

The stopped pass removed the `catalina-support-dev` project, the AI project, its networks and volumes, the 1.1 GB of the model, the built images and the temporary folders. The subsequent inventory showed only resources previous to the audit; the three stopped containers, the volume and the production images remained intact.

The repeat runs section 18 in full again after the fix reaches `main`. It does not reuse model, database, caches of Catalina Support or clone. It keeps the same boundaries: explicit names, no global prune, no backup of the disposable data and no touching production. A new functional finding stops and cleans the pass again before any code change.

## 20. TCP availability and history of the manual migration

The second pass of block 8 on `main` (`80e01a75c7bc3b284778bcd214850efb0690714b`) stopped at the first `docker compose -f dev.yml up -d --build`. The healthcheck declared PostgreSQL healthy while its temporary initialization server was answering on the local socket; `migrate`, which connects over TCP, started during the switch to the definitive server and received `connection refused`. The environment and the clone were removed completely before proposing changes.

The PostgreSQL healthchecks in `dev.yml`, `tests.yml` and `prod.yml` will check `127.0.0.1` and the configured port. That way they only enable dependencies when the same TCP path that other containers use accepts connections. `migrate` does not incorporate its own waits or retries: development and tests keep their automatic migration and the backend continues to depend on it finishing correctly.

## 21. Approved environment for the fourth pass of block 8

The third pass tested from `main` an empty installation, Mailpit, the real 1.5B model, summary and
writing, seeder, AD, Keycloak and return to local. The final review found the false errors
described in `docs/arquitectura.md`; by the zero-findings rule it is not considered closed.

All its containers, networks, volumes, model, own images, clone and temporary files were
removed. The final inventory kept only the three production containers that already
existed and that were neither consulted nor modified.

After integrating the fix into `main`, section 18 will be repeated in full from another clone and
new resources. The final inspection will require that successful walkthroughs do not produce false
absence errors and that real errors remain visible. Times, counts, cleanup
and all the already approved exclusions are kept.

The manual migration originally described in this section was replaced on 2026-10-20 by the
runner of section 27. This section is kept as a record of the TCP finding that originated the
`127.0.0.1` healthchecks; it no longer describes the current production operation.

The seeder is still a voluntary action and applies the idempotent schema first. The example data and the future policy of one file per version do not change. When there is a later migration, the service and the runbook will be updated through their own documentation cycle.

The automated verification reproduces a new database over TCP in development and tests. The
production verification inspects the resolved Compose configuration and runs the manual service
against a disposable database, without touching the postponed deployment.

### 20.1 Implementation and verification

The three healthchecks use `127.0.0.1`. A disposable instance of `tests.yml` created PostgreSQL from
zero, ran `migrate` once, started the backend after its correct exit and answered health
`200`. A temporary and isolated copy of `prod.yml` ran the manual command, confirmed
`installation_settings` and removed its container, network and volume. Without the profile, the resolved
production configuration does not include `migrate`; with it, it does. `bash -n` and the validation of the three
Compose files ended clean. Production was neither started nor modified.

## 22. Result of the fourth pass of block 8

The public commit `b0898922f6d584d9b3349f2de82adc6046e82315` correctly brought up `dev.yml`,
`ai.yml`, `active-directory.yml` and `keycloak.yml` from a new clone. Development applied the
migration automatically; AI downloaded and ran the 1.5B model; AD and Keycloak authenticated and
provisioned accounts; and the return to the local method was checked. The detailed counts and times
live in `docs/prueba-local.md`, section 20.

The final log review was empty for SQL errors, expected absences, panics and fatal errors. The cleanup
took down the four projects by their Compose files and removed their volumes and
networks. No global prune was used. The stopped production resources were preserved and
`prod-build.sh`, the production migration, nginx, the domain and the 1.0.0 tag were not run.

## 23. Priority verification

Backend covers validation, authorization, transaction, single notice, closed ticket and SMTP failure against
PostgreSQL. Frontend covers selector, accessible chip, editing, URL and failure toast. Playwright
walks creation, correction by Support, reading by the User, inboxes and Mailpit on PC and mobile; the
permissions of Development on the internal and the preservation when SMTP is interrupted are checked in
backend. The boundaries suite must remain clean.

The migration `v1.0.0.sql` is applied twice on a new database. The seeder is also repeated and
keeps its optional and idempotent nature. Production is not run, the domain is not opened and the
version tag is not created.

The verification of 2026-10-14 ran the full backend suite, 228 frontend tests and their
build, plus the E2E walkthrough on PC and mobile. The full pass had 181 cases passed and 38
expected skips; a live reload during the review interrupted an unrelated Settings
case, which later passed in isolation on a freshly created instance. The focused
priority walkthrough checked creation, editing, history, reloadable URL, reading by the User and single notice.
The migration was reapplied twice without changes and the seeder twice without duplicates; 25
example tickets remained —2 critical, 6 high, 11 medium and 6 low— and 24 templates.

## 24. Frontend reload in development

The `frontend` service mounts `./frontend` in `/app` and runs `ng serve` with the
development configuration. Live reload and HMR updates when they apply are default
behaviors of the Angular server; they are not duplicated as flags in `dev.yml`. Neither is polling
enabled: the current watcher detects volume changes correctly.

The check of 2026-10-15 observed in the logs **Changes detected**, a full rebuild and
**Page reload sent to client(s)**. Both `/` and `styles.css` responded `Cache-Control: no-cache`,
and the project does not register a Service Worker. This confirms that restarting containers is not part of the
normal flow for changes to components or already available Tailwind utilities.

A tab that was suspended or lost the connection to the development server may not receive
the reload notice. In that case the browser's forced reload is used or a new tab is opened.
Restarting `frontend` is reserved for dependency or Tailwind configuration changes, in accordance with
the rule already documented in section 7. No caches, polls or new processes are added.

## 25. README guide

The bilingual README will keep the empty installation through `/setup` as the main local path and
will present the seeder as a separate option to explore example data. The commands and values
come from this document and from the current Compose files; the rewrite does not change containers,
ports, migrations or configuration.

The short production guide will link this runbook and list, without replacing it: HTTPS, unversioned
secrets, persistent storage, real outgoing mail, verified backups, manual migration,
build/publish and health checks. It will state that example data is never
applied in production. The findings and limits already recorded in section 4 remain valid; a
friendlier cover does not turn them into implemented capability nor authorize running production.

After approval, local development was restored with the existing seeder —a destructive
operation already documented— and the name Catalina Support, the two factory logos and
the institutional color `#1d4ed8` were checked. The public evidence consists of a side-by-side view without sensitive data
and a second screenshot of the same ticket where `Summary`, `Last action` and `Improve with AI` are read.
Production was not used or altered to obtain them.

## 26. NVIDIA variant of the local engine

`ai.yml` continues to be the CPU deployment and works through Docker on Linux, macOS and Windows.
`ai-nvidia.yml` is an optional overlay for **Linux and Windows with Docker Desktop WSL2**: it uses
the engine's CUDA image, reserves a visible GPU and keeps the same GGUF volume, network and token.

Requirements:

- **Linux:** compatible NVIDIA driver, NVIDIA Container Toolkit and the NVIDIA runtime checked
  with `docker run --gpus all ... nvidia-smi`.
- **Windows:** Docker Desktop with the **WSL2 backend**, NVIDIA driver with WSL2 support and the GPU
  visible inside Docker. Native Windows containers are not used; the engine remains Linux.
- In both cases it is brought up with `docker compose -f ai.yml -f ai-nvidia.yml up -d`. The CPU compose
  remains the portable one for Linux, macOS and Windows.

The UI does not install drivers or restart Compose. If NVIDIA is requested without that variant, the
administrator answers absent capability and the interface shows the deployment command. Auto uses
CPU in the normal Compose and NVIDIA when the variant is available. Production only changes when
its operator decides to use that variant; this work does not deploy it.

**Real verification of 2026-11-05:** Docker Desktop WSL2, RTX A2000 12 GB, `effectiveAcceleration:
nvidia`, model `qwen2.5-1.5b-instruct` active and healthy, `/llama-server` using ~1.3 GiB and about
140 tok/s warm. The full suite in that environment ended with 200 E2E passed, 30 skipped and
0 failures; AD/Keycloak, 17 passed and 7 skipped.

## 27. Self-sufficient upgrade and migrations per version

The installable distribution will have a single official upgrade path: `scripts/update.sh`.
The script requires the `main` branch, a clean tree and a relationship that can advance fast-forward
from `origin/main`; if any condition is not met, it stops without changing files. It runs
`git fetch`, prepares and builds the new version outside the active installation and keeps intact
the version that is serving if it fails before migrating. It does not create merges, does not discard changes and does not
install a different branch. The version continues to be the constant indicated by the owner;
today it remains at `1.0.0`.

The deployment incorporates a single-run container built with a small command in
Go. Backend and frontend only start after that container finishes correctly. On each
start, the runner takes a PostgreSQL advisory lock, reads exclusively files whose name
is `vMAJOR.MINOR.PATCH.sql`, validates their SemVer order and only considers those that do not exceed the application
version. A file with a later version makes the package inconsistent and blocks
startup. The `*_dev.sql` files are never production migrations and the seeder remains
separate, optional and explicit.

`schema_migrations` records one row per version with version, file name, SHA-256 and application
date. There is only one file per version. The runner owns the transaction: each schema SQL
stops including `BEGIN` and `COMMIT`; running its content and recording the row form a
single transaction. If it fails, it reverts that version completely, does not record it, stops the application and
writes the file and the error without revealing secrets. To adopt control in existing databases,
`v1.0.0.sql` is run again idempotently and recorded; an incomplete database is not inferred nor marked
as valid.

While `shared/version.Development` is true, the only file that can change is the one of the
current version. If its checksum changes, the runner applies it again idempotently and updates
the record inside the same transaction. When publishing the version (`Development=false`), all its
files are immutable: any different checksum blocks startup. A database that records
a version higher than the application also blocks startup; automatic downgrade is not
admitted.

Before the first pending migration —and only then— a `pg_dump -Fc` is created, checked
that it can be read and saved in `/srv/catalina-support/backups/` with the source and
target versions in the name. It uses the current fourteen-day retention. If the backup fails, no SQL is executed. This protection belongs to the migration path itself, so it also operates if someone
starts Compose directly and not only from `update.sh`.

Published migrations are backward compatible: removing or renaming a structure is
postponed at least to a version later than the one in which it stopped being used. Thus, if the migration
finishes but the new application does not start, the operator can temporarily restore the previous
version while diagnosing. The database is not restored automatically: a restore that
overwrites data requires a human decision and the documented procedure of section 6.

During preparation the previous version keeps serving. The interruption begins, after
building everything, right before publishing the artifacts and running backup and
migration; the new application is only activated after they finish correctly. The upgrade record includes
start and end, previous and new versions, backup path, migrations and result. It excludes
credentials, sensitive SQL content and secret variables.

This implementation replaces the manual migration of the `migrate` profile, the manual order of the README
and the detection of pending schema as the normal flow. It keeps health as
a diagnostic defense, the existing daily backup, the separate development data and the policy
of one file per version. The user will be able to upgrade a published installation by synchronizing
`main` through the script and rebuilding the containers; later restarts will only validate the
history and will not create backups nor repeat migrations without changes.

The real verification adopted the existing development database and left a single valid backup; a
second run did not create another. A disposable database was born from zero, recorded a 64-character
checksum and was also left with no work on the second start. A deliberately broken migration
created its backup, failed, reverted the test table and kept the previous checksum. The runner's unit tests passed and the backend suite passed except the already known environmental failure of reporting
without seeded tickets. Both production binaries built and the runtime confirmed `pg_dump 18`.

## 28. Explicit examples per language

`dev.yml` exposes two single-run services: `seed-en` and `seed-es`. Their public commands
are `docker compose -f dev.yml run --rm seed-en` and `docker compose -f dev.yml run --rm seed-es`;
neither starts with `up`. On Linux and macOS `./scripts/dev-seed-en.sh` and
`./scripts/dev-seed-es.sh` are also offered. Docker remains the portable path for PowerShell, CMD and bash.

The periodic restart of `demo.catalina-support.com` uses `seed-en`, keeps its schedule and its current
credentials and leaves both the interface and the examples in English. Seeders remain
forbidden in production and outside development or demonstration deployments expressly
prepared for them.

README, runbook and cross references publish the explicit services and scripts per language.
The files `v1.0.0_dev_en.sql` and `v1.0.0_dev_es.sql` are the two data paths. The change does not affect production
deployments, TLS, DNS or the frequency of the public restart.

The previous `seed` service is withdrawn: it does not remain as an alias. The public restart replaces its
invocation with `seed-en` without changing schedule or credentials. The owner closed the review with
1A–5A and no open decisions remain.

The real verification ran both services on disposable databases. Each variant produced eleven
accounts, 25 tickets, 123 history rows, 27 comments, seven internals and five attachments. A
new database adopted `en` with `seed-en` and `es` with `seed-es`; a sealed installation kept its
settings when changing the set. The resolved Compose configuration contains only `seed-en` and
`seed-es`, and the runner tests confirm that it ignores both development files.
