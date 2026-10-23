# Catalina Support

[Español](README.es.md) · [Website](https://catalina-support.com) · [Live demo](https://demo.catalina-support.com) · [Support on Ko-fi](https://ko-fi.com/juanky1990)

**A quiet, open-source help desk for teams that need two clear support levels—not enterprise clutter.**

```text
requester  →  technical support (level 1)  →  development (level 2)
```

Catalina Support keeps the conversation simple. Users report an issue, Support handles what it
can, and only filtered cases—with their context—reach Development through a linked internal ticket.
The two conversations stay separate, while authorized staff can review them side by side.

AI is useful here, not decorative: a local or configurable engine summarizes each ticket's current
situation and latest action, and helps Support and Development improve drafts. People still review,
choose and publish every response.

[Try the live demo](https://demo.catalina-support.com) · [Install locally](#run-it-locally) · [Deploy on a server](#production-deployment)

![A main ticket and its internal ticket shown side by side](docs/assets/catalina-support-side-by-side.png)

The same ticket keeps AI assistance visible and reviewable: **Summary**, **Last action** and
**Improve with AI** all support the team without publishing anything automatically.

![AI-generated Summary and Last action beside the Improve with AI draft action](docs/assets/catalina-support-ai-insights.png)

## Why Catalina Support

- A focused two-level workflow: User → Support → Development.
- Main and internal tickets with independent conversations and a side-by-side view.
- AI-generated **Summary** and **Last action**, plus professional draft improvement.
- A local AI engine, your own compatible server, or an external provider.
- Categories, tags, priority, attachments, mentions, watchers and an auditable timeline.
- Local accounts, Active Directory or Keycloak—one sign-in method at a time.
- Editable email templates, installation branding, time zone and public URL.
- English and Spanish UI, eight accessible themes and responsive screens.
- Containers for the application, database, test mailbox and optional local AI.

## Try the live demo

Open **[demo.catalina-support.com](https://demo.catalina-support.com)**. It is a public, shared
environment: availability is not guaranteed and its data is restored periodically, so your changes
will not be permanent.

> **Do not enter personal, confidential or production information, and do not upload sensitive
> files.** Other visitors may see shared demo content.

All demo accounts below use the password **`123123123`**. The administrator account is deliberately
not published.

<details>
<summary>Show the 11 demo accounts</summary>

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

</details>

A good first tour:

1. Sign in as `user1@demo.com` and create a ticket from **My tickets**.
2. Sign in as `support1@demo.com`, open it under **Main tickets**, comment, change its state and
   escalate it with useful context.
3. Sign in as `dev1@demo.com`, open the linked ticket under **Internal tickets**, work on it and
   resolve it. The requester never sees this internal conversation.

## Run it locally

The recommended path starts with an empty installation and walks through `/setup`. Sample data is
optional and comes afterwards.

### Requirements

- Git.
- Docker with Docker Compose running.
- A browser.
- Internet access on the first run to download images, dependencies and the selected AI model.

You do **not** need Go, Node.js, PostgreSQL or nginx installed on the host. Assigning at least 4 GB
of RAM to Docker is a practical starting point. The default local model downloads about 1.1 GB and
does not require a GPU.

### 1. Clone and start

```bash
git clone https://github.com/Catalina-Support/catalina-support.git catalina-support
cd catalina-support
docker compose -f dev.yml up -d --build
docker compose -f ai.yml up -d --build
docker compose -f dev.yml ps
```

The development stack waits for PostgreSQL, applies the schema and then starts the backend. It does
not load sample data. Follow compilation if the application is not ready yet:

```bash
docker compose -f dev.yml logs -f frontend backend
```

`Ctrl+C` stops following logs; the services keep running.

### 2. Complete first setup

Open **[http://127.0.0.1:11001](http://127.0.0.1:11001)**. A fresh database opens `/setup`, not the
sign-in page.

| Step | Local value |
| --- | --- |
| Installation | Any name and language; changing the language updates the UI immediately |
| Sign-in | **Local** |
| Location | Your time zone and `http://127.0.0.1:11001` as the public URL |
| Outgoing email | Host `mail`, port `1025`, no TLS or credentials; the remaining editable defaults are already filled in |
| Artificial intelligence | **Local**; open the catalog, accept the 1.5B model license, download it, activate it and run the required test |

The email test checks connection and authentication without sending a message. AI is required: setup
can finish only after a valid generation. Then sign in with the development factory account
**`admin` / `admin`**, create users and open their invitation links in
**[Mailpit](http://127.0.0.1:11004)**.

The development password is defined in `config/env/dev.env`. Production uses its own
`ADMIN_PASSWORD`; never use `admin` there. The factory account remains the recovery door even when
the installation uses AD or Keycloak.

| Service | URL |
| --- | --- |
| Application | [http://127.0.0.1:11001](http://127.0.0.1:11001) |
| Test mailbox | [http://127.0.0.1:11004](http://127.0.0.1:11004) |
| Backend health | [http://127.0.0.1:11002/api/health](http://127.0.0.1:11002/api/health) |

### 3. Optional sample data

To load the same English accounts and realistic tickets used by the demo:

```bash
docker compose -f dev.yml run --rm seed-en
```

This creates 11 accounts, 25 tickets and 5 attachments. It also **deletes existing development
tickets, categories, tags, `@demo.com` accounts and development attachments except branding**.
Use `seed-es` instead for Spanish examples. Run either command only when you intentionally want to restore the
examples.

The seeder preserves a completed installation's branding, language, public URL, email, sign-in
method, directories, AI, prefix and assignment settings. On an unfinished installation it prepares
and seals a usable local demo. The 11 accounts and shared password are listed in
[Try the live demo](#try-the-live-demo); they are development credentials only.

### Stop and resume

```bash
docker compose -f dev.yml stop
docker compose -f dev.yml up -d
```

The database and attachments remain. `docker compose -f dev.yml down` removes containers and the
network but keeps volumes; adding `-v` also removes the database. Do not run `seed-en` or `seed-es` merely to restart:
it resets the sample content.

### Troubleshooting

| Symptom | Check |
| --- | --- |
| Docker does not respond | Run `docker compose version` and make sure Docker is running; Linux users also need daemon access. |
| A port is busy | Development uses ports 11001–11004. Stop the conflicting service. |
| The web page is not ready | Follow `docker compose -f dev.yml logs -f frontend backend`; detached startup does not wait for Angular and Go compilation. |
| The backend stays down after migration | Read `docker compose -f dev.yml logs migrate`, fix the error, run `docker compose -f dev.yml run --rm migrate`, then start the backend again. |
| `/setup` appears | This is expected for a new database without sample data. |
| Demo accounts cannot sign in | Load the sample data and select **Local** as the sign-in method. The factory account can always switch it back. |
| Mail does not reach Mailpit | Confirm `mail` is running and SMTP points to `mail:1025`, without TLS or credentials. |
| The model catalog does not respond | Read `docker compose -f ai.yml logs -f ai`; the AI manager must be running. |

Development publishes its ports on every host interface. Use it only on a trusted machine and
network; do not expose this stack as production.

## Optional development services

Start the main development stack before these services.

### Local AI

```bash
docker compose -f ai.yml up -d
docker compose -f ai.yml logs -f ai
```

No model is downloaded automatically. `/setup` and **Settings → AI engine** show source, license,
size, checksum and RAM before downloading the 1.5B, 3B or 7B model. The selected model can be
activated, tested and removed from the UI. See the [AI module documentation](docs/modules/ai.md).

### Active Directory test service

```bash
docker compose -f active-directory.yml up -d
```

This is an OpenLDAP directory with test identities, not an organization's real Active Directory.
Its local configuration and credentials are documented in the
[authentication module](docs/modules/auth.md).

### Keycloak test service

```bash
docker compose -f keycloak.yml up -d
docker compose -f keycloak.yml logs -f keycloak
```

It imports the repository's test realm and supports a real local OIDC round trip. Configuration and
test credentials are in the [authentication module](docs/modules/auth.md). Neither directory's test
credentials belong in production.

## Production deployment

Production is an operator-managed deployment, not the development stack copied to a public server.
Read the complete [deployment runbook](docs/ambientes.md) before starting. It records environment-
specific assumptions and known limitations that must be resolved for your server.

At minimum:

1. Use a dedicated host with Docker Compose, Bash, a reverse proxy and a valid HTTPS certificate.
2. Create `config/env/prod.env` from `config/env/prod.env.example` with unique database credentials,
   `ADMIN_PASSWORD`, `TOKEN_SECRET`, AI administration and encryption secrets. Never commit it.
3. Provide persistent storage for PostgreSQL, `_files`, `_logs` and local AI models when used.
4. Build and publish from a clean, reviewed revision. The bundled script publishes to
   `/srv/catalina-support`; inspect and adapt that assumption before running it.
5. Let the one-shot migration service apply signed-off version migrations before the application
   starts. It never discovers `*_dev.sql` and never runs the seeder.
6. Complete `/setup` before exposing the service, configure real SMTP and replace every test value.
7. Put the application behind HTTPS, restrict direct container ports and verify CSP and health.
8. Back up and test restoration of both the database **and attachments**. The bundled database
   backup does not include `_files`.

For the first deployment, build and start the stack:

```bash
./scripts/prod-build.sh
```

For later published updates, use `./scripts/update.sh` from a clean `main` branch. It only accepts a
fast-forward from `origin/main`, builds before interrupting the active version, creates and verifies
a database backup when migrations are pending, and starts the application only after they succeed.
Applied filenames and SHA-256 checksums live in `schema_migrations`; released migrations are
immutable. Verify the exact domain, nginx configuration, exposed ports, artifact paths and recovery
plan for your installation; the files include assumptions from the reference server.

## Tests

Unit tests run inside the development containers:

```bash
docker compose -f dev.yml exec backend go test ./...
docker compose -f dev.yml exec backend go vet ./...
docker compose -f dev.yml exec frontend npm test -- --watch=false
```

The browser suite uses an isolated disposable stack and never touches your development database:

```bash
docker compose -f tests.yml --profile directory down -v
docker compose -f tests.yml up -d --build database backend frontend mail
docker compose -f tests.yml run --rm e2e
docker compose -f tests.yml --profile directory down -v
```

To include AD and Keycloak, start the second command with `--profile directory` and add `ldap
keycloak`. Reports, traces and screenshots remain in `tests/e2e/resultados/`. More detail is in the
[local testing guide](docs/prueba-local.md).

## Contributing and security

Catalina Support follows **documentation before code**. Start with [AGENTS.md](AGENTS.md) and the
[documentation index](docs/README.md), then read the [contribution guide](CONTRIBUTING.md) and
[code of conduct](CODE_OF_CONDUCT.md). Substantial contributions require the
[Contributor License Agreement](CLA.md).

Use Issues for concrete defects and proposals, Discussions for open questions, and GitHub's private
security reporting for vulnerabilities. See the [security policy](SECURITY.md) and
[support channels](SUPPORT.md).

## Support the project

If Catalina Support saves you time, consider supporting its development on
**[Ko-fi](https://ko-fi.com/juanky1990)**. Donations are optional and do not change the software's
features, support terms or license.

## License

Catalina Support is free software licensed under the
[GNU Affero General Public License v3.0 only](LICENSE). You may use, study, modify and redistribute
it, including commercially. If you provide a modified version over a network, its users must be able
to obtain the corresponding source code under the license.

Also see the [copyright notice](NOTICE), [trademark policy](TRADEMARKS.md) and
[third-party notices](THIRD_PARTY_NOTICES.md).
