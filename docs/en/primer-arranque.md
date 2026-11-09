# First start: installation from scratch

[Español](../es/primer-arranque.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-04
>
> **Correction of 2026-11-04.** The wizard text no longer promises to remember the step on reload:
> the data from each step is preserved, but `/setup` returns to step 1. The sentence was adjusted to
> say what it actually does.
>
> **Amendment approved, implemented and verified on 2026-10-20.** Section 19 adds the optional site
> icon to step 1. Decisions 1A–14A are closed; 253 frontend tests passed, as did the boundaries and
> the production build.
>
> **Amendment approved, implemented and verified on 2026-10-20: a real message from step 4.** Section 18
> adds an ephemeral recipient and turns the SMTP test into an optional, rate-limited send with
> fixed text in the global language. The project owner closed and approved 1A–14A; the focused
> tests, the 248 frontend tests, boundaries and build passed.
>
> **Amendment approved, implemented and verified on 2026-10-19: safe choice of the SMTP
> transport.** Section 17 adds three modes to step 4, with explicit acceptance and a private-network
> limit for the unencrypted mode. The project owner closed and approved 1A–10A; the real test
> authenticated against the internal relay without sending mail.
>
> **Approved, implemented and verified on 2026-10-18.** Section 16 adds Auto, CPU and NVIDIA
> to step 5 when the engine lives on this server.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Section
> 15 automatically shows the step 5 catalog and reserves its button for refreshing without hiding
> the existing cards.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Section
> 14 adds theme and institutional color to the first step and replaces the model confirmations with
> the themed dialog.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Step 1 opens the upload when each logo
> is clicked without showing the chosen file; step 5 manages activation from the cards and allows
> staying temporarily without a model, keeping **Siguiente** disabled.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the project owner.** Step 1 adds the two optional
> logo slots and step 5 makes the real download progress visible, allows cancelling and
> resuming it, activates and tests the chosen model automatically and does not allow advancing until it is
> active and healthy. The project owner also decided to test the model manually with a greeting that
> includes the date and time of the installation's zone, and to manage downloaded inactive models in
> `/setup`. Section 12 records the finding, the closed scope and the verification.
>
> **Implemented and verified on 2026-10-07.** `/setup` requires the five steps in the approved order,
> tests and activates AI before sealing, shows a single privacy and possible cost confirmation
> for external providers and keeps partial downloads of local models so they can be resumed.
> Completion relies on the saved AI activation and does not repeat a call that could have a cost.
> The service tests, the 218 frontend unit tests and the full isolated Playwright run passed:
> 180 cases executed and 38 expected skips, on PC and mobile.
>
> **Correction by the project owner on 2026-10-07.** When implementation started it was found that the
> approved section 11 placed AI before mail, while the rest of the document, `AGENTS.md`,
> the code and the tests kept Mail in step 4 and AI in step 5. The project owner explicitly
> chose that order. This corrected proposal aligns the whole document before closing
> implementation; it does not change the content of the steps nor open production. The project owner
> explicitly approved the corrected document on 2026-10-07.
>
> **Amended on 2026-10-06**, with explicit approval from the project owner. The results of saving and
> testing a connection were shown before the wizard and could end up outside the window. They now use the same toast as
> `/settings` so that the result remains visible from the current step, without moving the
> informational notices specific to each step. It was implemented and verified on PC and mobile without changing the flow.
>
> **Amended on 2026-10-05**, with explicit approval from the project owner. In step 3,
> the chosen zone is only distinguished by the highlight within a long list and may fall outside
> the visible part. The project owner asked for a field that explicitly indicates the selected
> region and accepted the proposal to show it before the search box, as read-only, with the
> exact IANA identifier and the time with its offset below. It was implemented and verified on PC and
> mobile inside the disposable Playwright installation.
>
> **Amended on 2026-10-04**, in accordance with the approved `docs/prueba-local.md`: a new installation
> starts in English even if the browser remembers Spanish; the selector translates all of `/setup` at
> once and, when resuming, the saved language prevails. In development, an email that is still empty receives
> the editable Mailpit values; what is saved always takes precedence and production remains empty.
>
> **Amended on 2026-10-03**, in accordance with the approved `docs/prueba-local.md`: The Keycloak step
> accepts an optional `internalIssuer`. The Playwright preparation on a new database goes through the
> five steps on PC and mobile before sealing, checks Mailpit and the AI and ends on mobile.
> This automated verification replaces the historical limitation of point 3 of the deviations.
>
> **Amended on 2026-09-30**: **the wizard tests what it asks for**, as §3 and §6 promised. Two
> more public endpoints are added —`POST /api/setup/entry/test` and `POST /api/setup/mail/test`— with
> the same seal lock (**409** `setup.alreadyInstalled`), step 2 and step 4 debut their
> "Probar la conexión" button and the mail test **checks the connection and authentication without
> sending any mail**, because there is no recipient in the wizard. The data that arrives is validated
> the same way as when saving, with the same keys. With this, **deviations 1 and 1.b are corrected**
> from the block below. It counts as **1 amendment** (`docs/README.md`).
>
> **Written at the request of the project owner** (2026-09-30) and **approved and implemented the same day**: the project owner wanted
> to be able to test **the first-start view on another machine**, and for that view to **ask for all the
> information needed to start** —domain, address, port, mail, sign-in method— without
> having to edit files on the server.
>
> **What deviated from this proposal, told here** (Rule 0):
> 1. **The directory or Keycloak connection test (step 2) and the mail test (step 4) were not in
>    the wizard** when it was implemented. The `/api/setup/**` API had no test endpoints and
>    **there is still no session** with which to test: the "Probar la conexión" buttons in Settings
>    require being an Administrator, and here there is none. The data was saved and validated —the method
>    has to be configured in order to finish—, and **the test was done afterwards**. **Fixed on
>    2026-09-30**: both endpoints exist and are public, with the seal as the only lock, and
>    step 2 and step 4 bring their button (§3.1).
>
> 1.b. **The outgoing mail test from the wizard was not performed.** Step 4 saved and validated,
>    but tested nothing: to send a test email you need the mail already configured and a
>    recipient, and in the wizard there is neither. **Fixed on 2026-09-30**: step 4 has its
>    button and tests **the connection and authentication**, not the send, which is what can be checked
>    without a recipient (§3.1).
> 2. **Historical deviation, corrected on 2026-10-07:** the wizard had no step of its own to
>    test and activate AI. Now step 5 saves a valid activation before the summary and the
>    backend refuses to finish if it is missing.
> 3. **Historical limitation, corrected on 2026-10-03 and extended on 2026-10-07 in `tests.yml`:** the
>    full run executes on a disposable installation on PC and mobile. It checks the
>    five steps, the connections, the summary, the seal and the subsequent lock.
> 4. **The API paths are** `POST /api/setup/installation`, `/entry`, `/entry/test`, `/location`,
>    `/mail`, `/mail/test`, `/ai` and `/finish`.

## 1. What problem it solves

Today a new installation **starts already configured** with factory values, and the only door is **the
factory account**, whose password is in `config/env/*.env`. To set it up for real you have to:

1. edit the environment file on the server (mail, passwords, address);
2. sign in with the factory account;
3. go through **Settings** —which has eight cards— to set the name, the language, the sign-in
   method, the region and the address.

**It works**, and it is documented in `docs/ambientes.md`. But **it cannot be tested without editing a
file**, and someone coming from outside does not know where to start. This proposal adds **an installation view**
that is shown **only the first time** and asks for what is needed, in order, with its tests.

## 2. How the application knows it is not installed

**A seal in the configuration row itself**: `installation_settings.installed_at`
(`timestamptz`, **null** while nobody has finished the wizard). It goes **in `v1.0.0.sql`** —the
version is still open— with **backfill for installations that already exist**: a database that was already
configured **is sealed** with the date of its last update, so the wizard **does not
appear** in production just because of upgrading.

**An installation is "not installed" only if the seal is null.** That makes the wizard:

- **show itself** on a newly created installation, and
- **not be runnable again** on one already installed, neither through the screen nor through the API
  (**409** `setup.alreadyInstalled`), which is what prevents someone from rewriting the configuration of
  a running installation.

## 3. What it asks, and in what order

**Five steps**, with their progress line and a final summary. Each step saves as you advance, so
closing the browser halfway does not lose what was done: on returning, the wizard continues where it was.

| Step | What it asks for | Why |
| --- | --- | --- |
| **1 · The installation** | **The global language**, as the first control, and **the name** | It is what is read in the interface, the login, the menu, the emails and the new AI summaries |
| **2 · How you sign in** | **The method** —local, Active Directory or Keycloak— and **its data**, **with its test**: the button checks whatever corresponds to the chosen method. **The factory account is always explained**, whichever the method: what it is called and **where its password comes from** (`ADMIN_PASSWORD`, from the environment) | It is the door: without it nobody gets in, and the factory account gets in with **any** of the three methods. **The password is not requested here**: it stays in the environment, and that rule —that the factory door **does not depend on the database**— is not touched |
| **3 · Where it is** | **The time zone** —from the list with a search box— and **the public address** —scheme, host and port—, with **the notice if it is not https** | They decide how dates are read and where the email links point |
| **4 · The mail** | **The outgoing server**: host, port, TLS, user, password and sender, **with its test** | Without it no email goes out: neither an account creation, nor a reset, nor a notice |
| **5 · Artificial intelligence** | **Local engine, own server or external provider**. The model or connection is chosen, license and privacy are accepted when applicable and **Probar y activar** is run | AI is mandatory and a differentiator of the product; the installation is not sealed without a valid AI generation in the global language |

A new installation opens the full wizard in **English** and shows English as the language. The
step 1 selector immediately changes all the texts and `html[lang]`, and remembers the choice. If
the step was already saved, on returning the language persisted in the installation prevails.

**And the summary**, before finishing: it shows the modality, the provider or server, the model and whether
there is a credential, without showing secrets. For an external provider it repeats the accepted privacy notice.
The **Terminar la instalación** button is only enabled with the five steps complete and sets the seal.

**The three methods can be chosen from the start**, without waiting for them to be configured: in
Settings a method without its data is offered disabled, and here **that is not possible**, because otherwise there would be
no way to fill in the fields of the one to be used. Choosing it opens its fields, and on advancing
it is checked that they are complete.

### 3.1 The wizard's connection tests

The two steps that configure something that has to be checked bring their test action. For entry it is
called **Probar la conexión**; for mail, **Enviar correo de prueba**. Neither saves: they send what was written
and allow correcting it before advancing.

| Step | What it tests | What it does **not** do |
| --- | --- | --- |
| **2 · How you sign in** | **The directory (AD) or Keycloak, depending on the chosen method**: with AD it opens the connection and validates the service account; with Keycloak it reads the realm document. With the **local** method there is nothing to connect and the button is not shown. | It saves nothing, validates nobody's password —that is only known when someone signs in— and does not touch the database. |
| **4 · The mail** | The full transaction: connection, authentication, `MAIL FROM`, `RCPT TO` and `DATA`, with the sender and the ephemeral recipient as written. | It does not save the data nor claim final delivery to the mailbox. |

**They are answered by whoever already knew how to do it**, without duplicating anything: the directory and Keycloak are tested by the
`auth` module —the same one that answers the Settings buttons— and mail is tested by the `mail`
module, with the connection it uses to send. The `settings` module declares what it needs
(`Prober`, `ProberDeCorreo`) and `main.go` supplies it, as the modularity rule requires
(`docs/arquitectura.md`, section 4).

**The data that arrives is validated the same way as when saving**, with the same keys: the method has to
be one of the three and be configured, the directory and Keycloak have to be complete, and
mail needs a server, port and sender containing an @. An empty password or secret uses **what
is already saved**, which is what allows testing without retyping a secret that never leaves
the API. When the test does not reach its target, the screen reads `settings.directory.unreachable`,
`settings.keycloak.unreachable` or `setup.mail.unreachable`, as appropriate, and the technical detail is left
in the log.

## 4. What the wizard does **not** ask for, and why

- **Commercial keys, credit purchases and GPU.** The wizard accepts a credential that already
  exists and lets you choose a model from the local catalog, but does not contract services or configure the
  server hardware.
- **The factory account password.** It is still `ADMIN_PASSWORD` from the environment
  (`docs/usuarios-y-permisos.md`, section 8). The wizard **says where it comes from** instead of
  asking for it, because if the password lived in the database, the factory door **would stop being
  independent** of the rest of the configuration.
- **The certificate and the nginx domain.** They belong to the server: they go in the README, not in the application.

## 5. Outgoing mail moves house

Today SMTP is read from the **environment**; step 4 asks for it on the screen, so it goes on to **live in the
database**, like the directory and Keycloak (`docs/modules/settings.md`). This is done with the pattern that is already
in the repository: **the `mail` module declares what it needs** (`Instalacion`, an interface with the resolved SMTP)
and **`main.go` gives it to it** from the configuration module. **And the `SMTP_*` variables are removed from the environment**, for the same reason
that the directory and Keycloak ones were removed (correction by the project owner, 2026-09-30): **a piece of data that lives in the database is not
configured in two places**. In development, `SETUP_MAIL_*` provides only editable Mailpit suggestions
(`mail:1025`, without TLS or credentials and with the local sender) while the installation is
unsealed and has no saved mail. Saving the step persists them; afterwards the database always prevails.
Production does not declare those suggestions and starts empty. The `mail` module never reads them to send.

## 6. Where it lives, and what it is tested with

- **The screen**: its own route, `/setup`, **without the menu shell** —there is nothing to navigate
  yet— and with the factory brand.
- **The API**: `/api/setup/**`, in the `settings` module, **available only while the seal is
  null**. Saving the same installation twice is rejected with **409** and its key. Besides the five
  paths that save and seal, there are **the two tests** of step 2 and step 4 —`POST
  /api/setup/entry/test` and `POST /api/setup/mail/test`—, which receive the written data, **do not touch the
  database**, answer `{"status":"ok"}` —mail adds `recipient`— or the failure key, and **with the installation sealed they also answer
  409** `setup.alreadyInstalled` (§3.1).
- **The tests**: the unit tests of each step and **the two for the test endpoints** —with doubles, without
  going out to the network: the 409 with the seal set, the validation of bad data, and that with good data
  the module's test is called—, the one for the mail sender —which delivers the fixed message— and **the
  one that matters most in the interface suite**: that **an already configured installation does not
  show the wizard and its API rejects it**. The full run of the five steps executes
  automatically on a disposable installation, on PC and mobile, before the normal suite.

## 7. What is left out, and it has to be said

- **Buying or provisioning an AI provider**: an existing account and credential are configured.
- **Reinstalling**: once sealed, the configuration is changed in **Settings**. To start again
  from scratch you have to **delete the database** and apply the schema again, and that belongs to the server.
- **While the wizard is unfinished, the installation is not exposed**: anyone who reaches
  an unsealed installation can configure it. The README will say so in its step by step: **first you
  configure, then you open the domain**.

## 8. What this document enables

Approving this proposal enables: the seal column in `v1.0.0.sql` with its backfill, the route
`/setup` with its five steps and its summary, the `/api/setup/**` endpoints with their lock, outgoing mail
in the database —and out of the environment—, and the tests of the three layers. **None of that is
written until the project owner approves it.**

## 9. Selected region visible in step 3

### 9.1 Finding

In **Where it is / Dónde está**, the list marks the chosen zone with background and font weight. That mark
is not enough: the list is long, has its own scroll and the selected option may not be in
the visible part. After choosing, the search box is cleared and there is no stable piece of data saying which
is the current region. The time and offset help validate it, but they do not name the zone.

### 9.2 Behavior

Before the search box a read-only field appears:

- label **Selected time zone** in English and **Región horaria seleccionada** in Spanish;
- exact IANA value, for example `America/Guayaquil`;
- on opening the step it shows the zone the wizard already has —`UTC` on a new installation or the
  persisted one if resuming—;
- on choosing an option from the list, the field changes immediately;
- its text can be selected and copied, but not edited: a zone only enters the state by choosing it
  from the valid list;
- the current time and the offset continue below the list, as today.

The search box, the list, the option highlight, the public address and the way of saving the step
do not change. The field takes the available width both on PC and mobile and keeps a visible label
associated for screen readers.

### 9.3 Implementation scope

The implementation modifies `frontend/src/app/core/pages/setup-page.html`, adds the two texts to the
strict ES/EN dictionary and extends the Playwright preparation of the installation to check the
initial value and the zone change. `docs/interfaz-y-experiencia.md` also records the final
behavior.

Out of scope are the API, the database, zone validation, the general Settings screen
and the other three steps of `/setup`.

### 9.4 Acceptance criteria

1. On opening step 3 the field is shown with `UTC` or with the corresponding saved zone.
2. Choosing `America/Guayaquil` updates the field at once and keeps its time and offset below.
3. The field does not accept typing, but allows selecting and copying its text.
4. Search and the list keep working with keyboard and screen reader, without duplicate
   identifiers.
5. The behavior is checked on PC and mobile inside the disposable Playwright installation.

### 9.5 Recorded decision

The project owner confirmed the recommendation on 2026-10-05: read-only field before the search box,
exact IANA identifier and time/offset kept below.

### 9.6 Review record

The project owner confirmed on 2026-10-05:

1. The text is selected and copied with the field's normal behavior; no copy button is added.
2. This last observation is limited to step 3 of `/setup`; the general Settings screen does not
   change.

The review leaves no open decisions. The project owner explicitly approved this amendment with
"Yes". The implementation respects the approved scope and the document returns to `as-built`.

### 9.7 Verification

- The 216 frontend unit tests pass.
- The Playwright preparation of an empty installation passes on PC and mobile: it checks `UTC` in the
  new run, `America/Guayaquil` when resuming, the immediate update after choosing and
  the read-only attribute.
- The four cases of `arranque-instalacion.spec.ts` pass on PC and mobile after sealing that
  installation.

## 10. Visible results during first start

### 10.1 Rule

The general message placed before the wizard is replaced by the shared toast defined in
`docs/interfaz-y-experiencia.md`, section 15. It is used by the transient results of:

- failing to save a step or advance; when the save succeeds, the wizard advances without adding
  a redundant message;
- testing Active Directory or Keycloak in step 2;
- testing outgoing mail in step 4;
- general errors when loading, saving, testing or finishing the installation.

The notices that explain the current step remain inside the content: the factory account, the
address without HTTPS and any information that must be consultable while filling in the
fields. A field's errors remain associated with the field when that presentation exists.

### 10.2 Scope and acceptance criteria

This visual amendment does not change the steps, their public endpoints, the seal, the validations, the
initial language, the persistence or the final summary.

1. The result of the step 2 and step 4 tests is seen in the current window on PC and mobile.
2. A failure when advancing is shown without taking the window to the top and remains until closed or
   replaced.
3. Any toast disappears after 10 seconds at most and can be closed earlier.
4. The context notices remain inside the step when changing language or screen width.
5. The wizard tests check a successful connection, a failed one, manual closing and the
   fixed position of the toast.

### 10.3 Decision and review

The project owner confirmed on 2026-10-06 that `/setup` uses the same component and the same rules as
`/settings`. The common review was closed in `docs/interfaz-y-experiencia.md`, section 15.7, with no
open decisions.

### 10.4 Implementation and verification

`setup-page` shows its general signal through `app-toast`; its explanatory notices remain inside
the corresponding step. The Playwright preparation of the empty installation checked on PC and mobile
that the mail test shows the fixed toast inside the window and that its **Cerrar** button removes it. The current run keeps the five steps, seals the installation and ends at the login.

## 11. Implementation: five steps with global language and mandatory AI

The new order is:

1. **Language and installation:** the language is the first control and immediately changes all the wizard
   texts; then the name is requested. That language will be the same for interface, dates, emails and AI.
2. **How you sign in:** it keeps local, AD and Keycloak and their current tests.
3. **Where it is:** it keeps time zone and public address.
4. **Outgoing mail:** it optionally tests the full transaction with an ephemeral recipient.
5. **Artificial intelligence:** local, own server or provider. For local, the catalog is chosen,
   resources are checked, source, license, size and checksum are shown, the license for
   that version is accepted, it is downloaded with progress and activated. For external, connection and secret are written, and
   it is confirmed which data will leave the server.

Step 5 is only complete after **Probar y activar** with a valid AI generation
in the global language. An incomplete local download **resumes from the saved state** and the
wizard remains on step 5; closing or reloading the browser does not turn a partial download
into an installed model nor force repeating what was already downloaded. The finish button remains
disabled if AI, mail or any of the previous steps is missing.

A valid activation already saved is enough to finish. **Finishing does not repeat the AI test**: if the
engine stops responding between activation and the seal, the installation can be completed and the subsequent
outage is treated like any unavailability of the engine, leaving the jobs pending. Thus
a transient interruption does not block a correctly configured installation.

The final summary shows modality, provider or server, model and whether the credential is set,
never the secret. For an external provider, before testing **a single confirmation** is presented
that brings together the privacy notice and the possible cost; two acceptances are not requested nor is the
confirmation repeated for the same activation. The summary recalls that it was accepted. On finishing, the seal is set in
a single operation as today; a failure does not leave a partial seal.

Out of scope are asking for commercial keys, buying credits and configuring GPU. The factory password
continues in `ADMIN_PASSWORD`. The proposal replaces the current statement that the wizard only
automatically checks an optional engine and the exclusion of choosing a model from the application.

The project owner confirmed the scope on 2026-10-06. The review added the acceptance of license and
catalog integrity and confirmed that costs only affect later recoveries, which for a
provider remain under the Administrator's decision. On 2026-10-07 they explicitly corrected the order:
Mail remains in step 4 and AI occupies step 5. In the final review of the same day they chose that a
saved activation allows sealing even if the engine falls afterwards, a single joint confirmation of
privacy and cost for external providers, and resuming incomplete local downloads. No
open decisions remain.

## 12. Brand and visible model administration

### 12.1 Findings

The implementation does not fulfill two current statements. The catalog delivers `downloadedBytes`, but
`/setup` and `/settings` only query once and show «Descargando» without a bar, amount or progress.
Nor does the cancellation promised by `docs/interfaz-y-experiencia.md` exist. In addition, step 5 allows
pressing **Siguiente** during a download and then tries to test an engine that is not yet
ready, so it shows the generic inaccessible-engine error. They are reported before touching code.

Step 1 only asks for language and name. The logos can be uploaded later from Settings, although
they are part of the same initial identity. The project owner asked to bring their two slots to step 1.

### 12.2 Step 1: two optional logos

After the language and the name the **Claro** and **Oscuro** slots appear, with the same preview,
validation, replacement, fallback to the other slot and return to the factory file as in Settings. Both
are optional and never block **Siguiente**. Each file is saved when uploaded, the header adopts
the brand immediately and reloading keeps what was saved.

The routes live under `/api/setup/**`, only work with the null seal and afterwards answer `409
setup.alreadyInstalled`. They keep PNG, JPEG, WebP and SVG, the maximum of 1 MB and 2000 pixels per side.

### 12.3 Step 5: download and activation

There is only one download at a time. The card shows a bar, percentage and downloaded and total bytes,
updated automatically. **Cancelar** stops the transfer and keeps the partial file;
**Reanudar** continues from it. Another action allows permanently deleting a partial file. Closing or
reloading the browser does not stop the work already entrusted to the server.

On completing the chosen file, the server validates size, header and checksum, activates it, waits for
health and runs the current automatic test. The card distinguishes download, validation, activation,
test, active and error. A failure keeps the installed model, explains the stage and offers
**Reintentar activación**. **Siguiente** remains disabled until the chosen one is active and
healthy and says what is missing, instead of causing a premature test.

Inactive installed models can be deleted with confirmation. The active one does not offer that
action: first another one is activated. The same rule applies in `/settings`.

### 12.4 «¡Pruébame!» / “Try me!”

The active and healthy model offers a **¡Pruébame!** button in Spanish and **Try me!** in English. The
backend takes the current instant, expresses it in the installation's IANA zone and gives it to the
model so that it greets in the global language and includes that date and time. It does not rely on the model
knowing the clock nor are tickets, people or files sent.

The card keeps the greeting, the authoritative instant and the zone until the test is repeated, the
model is changed or the screen is left. A failure stays inside the card, does not deactivate the model and
allows retrying. This manual test does not replace the automatic test that enables **Siguiente**.

### 12.5 Scope and closure

The 1.5B/3B/7B catalog is kept, a single simultaneous download, a single active model and the existing
permissions. Custom models, production, the domain and the tag
`v1.0.0` are not included.

The tests cover services, controllers and components, and a disposable run on PC and mobile:
progress, cancellation, reload, resumption, background work, activation, failure and retry,
greeting with zone and instant, safe deletion and both logos. The project owner's review closed all
decisions: full progress and cancellation; active and healthy to advance; automatic activation
in the background; one download; failure kept; immediate and optional logos; deletion only of
inactive ones; and greeting visible inside the card. No open decisions remain.

## 13. Clickable logos and control from the cards

In step 1, clicking the light or dark preview opens the corresponding selector. The file is
saved immediately and only the updated logo is seen: its path or name is not presented. The
control keeps an accessible name, visible focus and keyboard operation; restoring remains a separate
action.

In step 5 the top selector of the local model disappears. Each installed model shows its
status and **Activar** or **Desactivar** in the card itself. It is allowed to leave the engine without an active model, but
in that state **Siguiente** remains blocked and explains that a healthy model must be activated.
The manual greeting is presented by toast, attributed to the model name, for a maximum of ten
seconds and with manual closing. It does not repeat the date and time the model already says nor show the zone.

The tests go through both logos and the states without active model, active and healthy on PC and mobile. They do not
change the five steps, the seal, the logo limits, the catalog or the automatic test.

## 14. Initial appearance and own confirmations

Step 1 brings together name, language, theme, institutional color and the two logos. Language and theme appear
together; the theme offers the eight existing options and changes live through the browser's local
preference. Below, the color uses the control and preview from Settings. The chosen color
is kept when changing theme and travels together with name and language when **Siguiente** is pressed.

The step 1 contract adds `primaryColor`. The preview uses the existing color resolution contract
without persisting; saving reuses the same hexadecimal validation and updates the public
brand. Reloading keeps the already saved color, while the theme remains a choice of that
browser and not a new column.

Deleting a model or a partial file opens `app-dialogo`, presented with the visible theme. Cancel, Escape
or closing do not execute the request; confirm uses the dangerous button and keeps the current deletion
rules.

The tests verify the eight themes, color and preview, persistence when advancing, reload, dialog
in light and dark themes, keyboard, PC and mobile. The order and number of steps do not change.

The real run confirmed the full catalog of eight themes, the local persistence of the choice,
the resolved color preview and the mobile design at 390 px. Advancing saved the color together with the
name and the language; the delete dialog opened and closed with Escape without deleting the model.

## 15. Catalog visible on entering step 5

Entering step 5 in local mode immediately requests the catalog and presents its three cards.
Returning from another modality to **En este servidor** does the same, without requiring an extra click.

The button is called **Refrescar catálogo** / **Refresh catalog** and only repeats the query. While it
is working it keeps the previous catalog, shows its working status and avoids duplicate requests.
If it fails, it keeps the known cards and shows the error through the shared toast.

The tests verify direct entry, return to local, successful and failed refresh, PC and mobile. They do not
change the fifth step, the seal, the API or the administration of each model.

The real run advanced from step 1 with the saved values and found the three cards already
loaded in step 5. The manual refresh finished in place without hiding them or altering models.

## 16. Acceleration in the AI step

Step 5 shows Auto, CPU and NVIDIA only when **En este servidor** is chosen, with Auto as the initial
value. Activating a model restarts and tests using the chosen mode; **Siguiente** remains blocked
until there is a healthy model. The card and **Try me!** show the effective mode and the commercial name of the
device without technical identifiers.

Auto falls back once to CPU with a notice. Explicit NVIDIA without the Compose variant rejects activation and
explains how to bring up `ai-nvidia.yml`; it does not seal or save an unusable configuration. Resuming the
wizard rebuilds the configured mode, the effective mode and the `restarting` operation if it exists.

## 17. SMTP transport in step 4

Step 4 replaces the **Conexión cifrada (TLS)** checkbox with a selection of three modes:
**STARTTLS**, **TLS directo** and **Sin cifrado**. STARTTLS is the normal option for port 587 and requires
the extension before authenticating; direct TLS normally corresponds to 465. The screen does not choose
a mode from the port nor automatically downgrade a secure connection.

**Sin cifrado** shows before testing or advancing that user, password and content will travel in
the clear. It requires an administrative acceptance that is kept only while that mode remains
selected. In addition, the backend requires that all resolved addresses be RFC1918 private IPv4 or
loopback. The common toast reports a rejection per mode, acceptance or destination without revealing credentials.

The test connects, authenticates and sends the technical message. The sender will be Catalina's own institutional address; the wizard does not propose or copy SIGOC's. Development
suggestions change Mailpit to `plain` without credentials and with its explicit local exception.

## 18. Real test email in step 4

This amendment replaces the last sentence of section 17 and the step 4 row in §3.1: the wizard
does ask for an **Enviar prueba a** address and the single **Enviar correo de prueba** button performs a
full SMTP send with what is written. The field is used only in that request; it is not saved nor
does it appear in the summary.

The fixed message uses the global language and explains that it validates Catalina-Support, that it requires no action
and that it must not be replied to. The result confirms that the server accepted the message for the indicated
address, without promising that it reached the inbox. Not testing still allows advancing and sealing.

Since the endpoint exists before signing in, it limits each IP–recipient pair to five attempts
in fifteen minutes, answers `429` when exceeded and remains inaccessible after the seal. The logs
partially hide the recipient and exclude secret and content.

## 19. Optional icon in step 1

After the two logos, step 1 presents the **Icono del sitio** card defined in
`docs/interfaz-y-experiencia.md`, section 50. Uploading or restoring it uses the public routes of
`/api/setup/brand/site-icon`, available only while the installation is not sealed, and updates
the favicon of all open tabs immediately.

The icon is not mandatory and does not participate in **Siguiente**, in the step validation or in the final
seal. If it is omitted, the tab uses the custom light logo and, if that does not exist either, the
factory light logo. Errors remain in the step through the common toast and do not discard the other
data already written.

## 20. Local catalog unavailable without a false server-down

> **Proposal of 2026-10-23.** The ephemeral run confirmed that the absence of the local model
> manager returns 503 on entering step 5 and the global interceptor confuses it with a complete
> backend outage. The project owner chose 1A and closed the review with 2A and 5A. No open
> decisions remain; explicit approval is pending.

Entering step 5 does not presuppose that the local engine manager is available. If the catalog
query fails, the wizard keeps the three AI options and shows the result inside the local card. The
Administrator can retry or choose their own server or an external provider without reloading or losing what was written.

The card offers **Volver a intentar** without changing mode or clearing the fields. The same rule
applies in Settings, through the contract described in `modules/settings.md`.

The global **No hay conexión con el servidor** screen is reserved for a general loss of the
backend. A 503 from `/api/setup/ai/models` does not activate it by itself. Testing and saving the option
finally chosen remains mandatory.

The regression opens step 5 with a healthy backend and the local manager absent, confirms the local notice,
switches to their own server and finishes the installation using the real test endpoint.
