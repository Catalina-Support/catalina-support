# mail

[Español](../es/modules/mail.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-03
>
> **QA correction, 2026-11-03.** The editor said «Los diez correos» even though there were fourteen, and the
> `auth.accountCreated` template appeared as a raw key because it had no label. The texts are corrected
> to fourteen, the label is added in Spanish and English and the language review messages and the
> unknown-template error key are updated.
>
> **Proposal of 2026-10-30, approved, implemented and verified.** Section 19 adds the `auth.accountCreated` email, template
> number 14, to notify the person whose account has just been created from a ticket. It goes into
> `v1.0.1.sql`; the application and the schema move to 1.0.1. The editor goes to twenty-eight rows.
>
> **Amendment approved, implemented and verified on 2026-10-20: real SMTP test send.** Section 17
> replaces the check that only connected and authenticated with a real message to an expressly
> written recipient. It uses the sender and the SMTP data being tested, without saving them, and
> confirms only that the server accepted the message. The project owner closed and approved 1A–14A;
> the focused tests, the 248 frontend tests, boundaries and the build passed.
>
> **Amendment approved, implemented and verified on 2026-10-19: three SMTP modes and an explicit
> insecure exception.** Section 16 distinguishes STARTTLS, direct TLS and unencrypted SMTP; the
> latter requires administrative acceptance, accepts only private destinations and is never enabled as
> automatic degradation. The project owner closed and approved 1A–10A. The real relay accepted `AUTH LOGIN`
> without sending email; the focused regressions, 247 frontend tests and the build passed.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 15 adds the editable private comment notice
> for the assignee and internal watchers, with no body or attachments and excluding the author. The
> project owner chose and approved 9A and 21A–23A.
>
> **Amendment proposed for 1.0.0, 2026-10-14.** Section 14 adds the twelfth editable email,
> the single critical-priority notice, in the two global languages. No code is modified before
> implementing it. The project owner approved it explicitly.
>
> **Amendment implemented and verified on 2026-10-13: missing template without a false error.** The
> lookup keeps `mail.template.notFound` without logging `record not found`; real SQL failures
> remain visible. It is part of 1A–4A, approved after the third walkthrough of block 8.
>
> **Amendment implemented and verified on 2026-10-13: validate the language of drafts.** Section 12 adopts
> the central control of `ai`: a clearly incorrect translation does not replace the draft and an
> external provider does not get an automatic retry. The project owner closed 1A–12A with 10B. The change stayed within what was explicitly approved by the project owner.
>
> **Amendment corrected on 2026-10-07.** What exists remains as-built. All sends will use the
> global language and a language change may generate with AI reviewable drafts of only the
> twelve templates, protecting HTML and markers. The project owner withdrew the mass regeneration of
> summaries on language change and confirmed that manual editing will always be available.
>
> **Amended on 2026-09-30**: the module debuts **`Probar`**, which **connects and authenticates against the mail
> server without sending any email**: it opens the connection with what it is given —host, port,
> TLS, user and password—, validates the credentials and closes. **There is no `MAIL FROM`, no `RCPT TO`, no
> `DATA`.** It uses **the same connection as the send** —the one that resolves direct encryption and STARTTLS—, so
> what is tested is what will be used when sending. It is consumed by the first-start wizard
> for step 4, where there is still no recipient (`docs/primer-arranque.md`, section 3.1).
>
> **Amended on 2026-09-30**: **outgoing mail is read from the installation settings**
> —the first-start view asks for it— and **the `SMTP_*` environment variables were removed**: a value that lives in
> the database is not configured in two places (correction of the project owner, 2026-09-30). The
> sender is resolved **on every send**, so changing it on the screen works without restarting anything
> (`docs/primer-arranque.md`, section 5).
>
> **Amended on 2026-09-27**, when **ticket mentions** came in (`docs/modules/tickets.md`,
> decision 58): **template number eleven** is born, `ticket.mentioned` —«Te han etiquetado en {{numero}}:
> {{asunto}}»—, which goes **to each mentioned person** when they are mentioned for the first time. **The emails
> still carry neither the ticket text nor the comment text**: number, subject and link, like the others.
>
> **Becomes as-built on 2026-09-26**: the module **has been complete since 2026-09-24** —the table with its
> twenty templates, the markers, the rendering, the sending, its five endpoints and the editor—, so
> this document describes what the code does. The header stayed at `aprobado` through a
> slip of form, not because anything was missing; the record of its implementation tells that story further down.
>
> Approved by the project owner on 2026-09-23, after reviewing it in the form of questions. It fixes the table of
> templates, the markers, **the twenty-four default texts** (twelve emails in Spanish and English),
> sending in HTML with a text version and the editor of the Settings screen.
>
> **It is the first module to be implemented**, because `auth` needs it to send the invitation email.
>
> **Implementation status (2026-09-24): the module is complete.** There are the table with its
> twenty templates, the markers, the rendering, the sending and **the five endpoints**, tested
> end to end against development, **and the editor**: the screen of the `mail` module that
> Settings links to (decision 15), with its two languages side by side, the markers that are inserted when
> clicked, the preview rendered by the backend and the test button.
>
> **Amended on 2026-09-23** (when its backend finished being implemented): the document still had a
> section 9 of «decisions to confirm» whose three decisions **were already in the body and already
> in the code** (the warning that does not prevent saving, the test with `[prueba]` and `PUBLIC_APP_URL`),
> so they have been moved to the review record (9, 10 and 11) and the section has disappeared: an approved document
> cannot be left with open decisions.
>
> **Amended on 2026-09-24**, when building the editor: the **preview** that the document asks for had no
> way of being done —the four endpoints were list, save, restore and send a test—,
> so it debuts a **fifth endpoint**, `POST …/{key}/{language}/preview`, which returns the email already
> rendered with the same sample data as the test. It is a `POST` and **accepts the draft**
> —`{subject, body}`—, not only what is saved: the preview serves to decide **before**
> saving, and it also warns about the invented marker at the same moment it is written. It is a decision of the project owner of
> 2026-09-24: the alternative was for the frontend to replace the markers on its own, and that is
> having two versions of the same rule. The other three decisions of that batch are in the record
> (17, 18 and 19).
>
> **Amended on 2026-09-23**, when writing the four endpoints: the document had
> **no error key** of the module, nor did it say what each endpoint returns when it goes well. They are added
> as section 8, with decisions 12 to 15 of the record, all approved by the project owner.

## 1. Scope of this document

It tells **how the `mail` module is built**: the template table, the markers, the fourteen
default emails with their texts, the sending, the endpoints and the editor of the Settings
screen.

**It does not repeat**: the permissions (`docs/usuarios-y-permisos.md`), which email is sent at each moment
(`docs/propósito-y-alcance.md` and `docs/flujos.md`) or the design of the screen
(`docs/interfaz-y-experiencia.md`).

It is **approved**, so it enables writing code (Rule 0). Section 9
collects what I proposed.

## 2. What the module does and does not do

**It does**: save the texts, validate them, fill them with data, send them over SMTP and let an
administrator edit them, preview them and send a test to themselves. The settings test delivers
a technical message to the ephemeral address written to also check the sender, the recipient and the
SMTP content (§17).

**It does not do**: decide **who** is notified. Whoever requests the send —`auth` or `tickets`— resolves the
recipients and passes them on. That is what keeps `mail` knowing nothing about roles, distributions or
tickets: **it receives a template, a language, some recipients and some data, and sends**.

## 3. The template table

**`mail_templates`**

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | |
| `key` | `text` | Which of the fourteen emails: `ticket.created`, `auth.recovery`… |
| `language` | `text` | `es` or `en` |
| `subject` | `text` | The current subject, with its markers |
| `body` | `text` | The current **HTML** of the body, with its markers |
| `default_subject` | `text` | The **factory** subject. It is never edited |
| `default_body` | `text` | The **factory** body. It is never edited |
| `updated_at` | `timestamptz` | When it was last edited |
| `updated_by_id` | `bigint` | Who edited it. Null while the factory text remains, and also when the factory account `admin` edits it, which is not in the accounts table |

**Unique per (`key`, `language`)**: twenty-eight rows, not one more —fourteen emails in two languages—.

- **The twenty-eight texts are seeded in the migration**, with `INSERT ... ON CONFLICT DO NOTHING`. That
  matters: being idempotent, **reapplying the migration does not overwrite what an administrator has
  edited**. If a factory text ever has to be changed, it will be with a new migration that says
  so, not by rewriting the seed.
- **Each row keeps both versions**: the one in use (`subject`, `body`) and the factory one
  (`default_subject`, `default_body`). It is what makes **restoring be copying two columns** instead
  of having the default texts in two places —the SQL and the code— that would end up saying
  different things. They cost a few kilobytes and remove a problem.
- **`updated_by_id` says whether someone has touched a text**: null means factory. It is the difference
  between «restore» and «there is nothing to restore».
- **Restoring** copies the factory columns over the current ones and leaves `updated_by_id` null.
- If a factory text ever has to be changed, it is done with a migration that updates
  `default_*` and, **only if nobody has touched it**, the current one too.

**The factory columns cannot be written from the API**: the `PUT` of a template only touches
`subject`, `body` and `updated_by_id`.

## 4. The markers

They are written with double braces and in lowercase: `{{numero}}`, `{{enlace}}`, `{{nombre}}`.

- **Each email has its list**, and the editor shows it next to the text.
- **It is validated on save**: a marker that is not in that email's list **is not saved**, and the
  warning says which one. It is half the module's job: an invented marker is discovered on save,
  not when the email has already gone out.
- **They are not required to all be there** (decision of the project owner, section 8). The editor **warns** if one
  of the essential ones is missing —the link, the number— but lets you save: there are legitimate cases for
  writing a notice without a link, and whoever edits knows what they are doing.
- **The data is always escaped** before entering the HTML, because the subject of a ticket is
  written by a person and may carry `<`, `>` or `&`. The escaping belongs to the code, not to the text: there is no
  way to write a template that skips this.
- **The links are built with the public address of the installation** —Settings— and
  `PUBLIC_APP_URL` remains **as a fallback**, never with a host written in the text: the same
  text has to work in development and in production.

## 5. The fourteen default emails

These are the texts that are seeded. **They are reviewed here, which is where they can be corrected without touching
code.**

### 5.1 The eight ticket notices

| Key | Who receives it |
| --- | --- |
| `ticket.created` | **To whoever the distribution decides** for new tickets (the assigned technician, or nobody, depending on the settings) |
| `ticket.escalated` | To whoever the distribution decides for the internal ticket |
| `ticket.waitingUser` | To the requester |
| `ticket.waitingSupport` | To the technician handling the ticket |
| `ticket.resolved` | To the requester |
| `ticket.closed` | To the requester |
| `ticket.backToSupport` | To the technician handling the ticket |
| `ticket.mentioned` | **To whoever has just been tagged** in a comment (`docs/modules/tickets.md`, decision 58) |

**And the watchers receive the same staff emails as the assignee** (decision 61 of
`tickets.md`): the distribution one, the escalation one, the «Development needs something» one and the «the ticket
goes back to Support» one. The ones that go **to the requester** are not duplicated for them.

**`ticket.created`** — markers: `{{numero}}`, `{{asunto}}`, `{{solicitante}}`, `{{enlace}}`

> **Subject (ES)**: Ticket nuevo: {{numero}}
> **Subject (EN)**: New ticket: {{numero}}
> **Body (ES)**: `<p>Se ha creado el ticket <strong>{{numero}}</strong>.</p><p>{{asunto}}</p><p>Lo ha pedido {{solicitante}}.</p><p><a href="{{enlace}}">Abrir el ticket</a></p>`
> **Body (EN)**: `<p>Ticket <strong>{{numero}}</strong> has been created.</p><p>{{asunto}}</p><p>Requested by {{solicitante}}.</p><p><a href="{{enlace}}">Open the ticket</a></p>`

**`ticket.escalated`** — markers: `{{numero}}`, `{{asunto}}`, `{{motivo}}`, `{{enlace}}`

> **Subject (ES)**: Escalado: {{numero}}
> **Subject (EN)**: Escalated: {{numero}}
> **Body (ES)**: `<p>El ticket <strong>{{numero}}</strong> se ha escalado a Desarrollo.</p><p>{{asunto}}</p><p><strong>Motivo:</strong> {{motivo}}</p><p><a href="{{enlace}}">Abrir el ticket interno</a></p>`
> **Body (EN)**: `<p>Ticket <strong>{{numero}}</strong> has been escalated to Development.</p><p>{{asunto}}</p><p><strong>Reason:</strong> {{motivo}}</p><p><a href="{{enlace}}">Open the internal ticket</a></p>`

**`ticket.waitingUser`** — markers: `{{nombre}}`, `{{numero}}`, `{{asunto}}`, `{{enlace}}`

> **Subject (ES)**: Necesitamos algo de ti: {{numero}}
> **Subject (EN)**: We need something from you: {{numero}}
> **Body (ES)**: `<p>Hola {{nombre}}:</p><p>Para seguir con tu ticket <strong>{{numero}}</strong> ({{asunto}}) necesitamos que nos contestes.</p><p><a href="{{enlace}}">Responder en el ticket</a></p>`
> **Body (EN)**: `<p>Hi {{nombre}},</p><p>To continue with ticket <strong>{{numero}}</strong> ({{asunto}}) we need you to get back to us.</p><p><a href="{{enlace}}">Reply in the ticket</a></p>`

**`ticket.waitingSupport`** — markers: `{{numero}}`, `{{asunto}}`, `{{enlace}}`

> **Subject (ES)**: Desarrollo necesita algo: {{numero}}
> **Subject (EN)**: Development needs something: {{numero}}
> **Body (ES)**: `<p>Desarrollo necesita algo de Soporte para seguir con el ticket <strong>{{numero}}</strong>.</p><p>{{asunto}}</p><p><a href="{{enlace}}">Abrir el ticket interno</a></p>`
> **Body (EN)**: `<p>Development needs something from Support to continue with ticket <strong>{{numero}}</strong>.</p><p>{{asunto}}</p><p><a href="{{enlace}}">Open the internal ticket</a></p>`

**`ticket.resolved`** — markers: `{{nombre}}`, `{{numero}}`, `{{asunto}}`, `{{enlace}}`

> **Subject (ES)**: Tu ticket {{numero}} está resuelto
> **Subject (EN)**: Your ticket {{numero}} is resolved
> **Body (ES)**: `<p>Hola {{nombre}}:</p><p>El ticket <strong>{{numero}}</strong> ({{asunto}}) está resuelto.</p><p>Si no es así, respóndenos en el ticket o ábrelo de nuevo.</p><p><a href="{{enlace}}">Ver el ticket</a></p>`
> **Body (EN)**: `<p>Hi {{nombre}},</p><p>Ticket <strong>{{numero}}</strong> ({{asunto}}) is resolved.</p><p>If that is not the case, reply in the ticket or reopen it.</p><p><a href="{{enlace}}">View the ticket</a></p>`

**`ticket.closed`** — markers: `{{nombre}}`, `{{numero}}`, `{{enlace}}`

> **Subject (ES)**: Tu ticket {{numero}} se ha cerrado
> **Subject (EN)**: Your ticket {{numero}} has been closed
> **Body (ES)**: `<p>Hola {{nombre}}:</p><p>El ticket <strong>{{numero}}</strong> está cerrado.</p><p>Si vuelve a pasar, puedes reabrirlo.</p><p><a href="{{enlace}}">Ver el ticket</a></p>`
> **Body (EN)**: `<p>Hi {{nombre}},</p><p>Ticket <strong>{{numero}}</strong> is closed.</p><p>If it happens again, you can reopen it.</p><p><a href="{{enlace}}">View the ticket</a></p>`

**`ticket.backToSupport`** — markers: `{{numero}}`, `{{asunto}}`, `{{enlace}}`

> **Subject (ES)**: {{numero}} ha vuelto a tu bandeja
> **Subject (EN)**: {{numero}} is back in your queue
> **Body (ES)**: `<p>El ticket <strong>{{numero}}</strong> ha vuelto a Soporte.</p><p>{{asunto}}</p><p><a href="{{enlace}}">Abrir el ticket</a></p>`
> **Body (EN)**: `<p>Ticket <strong>{{numero}}</strong> is back with Support.</p><p>{{asunto}}</p><p><a href="{{enlace}}">Open the ticket</a></p>`

**`ticket.mentioned`** — markers: `{{numero}}`, `{{asunto}}`, `{{enlace}}`

> **Subject (ES)**: Te han etiquetado en {{numero}}: {{asunto}}
> **Subject (EN)**: You were mentioned in {{numero}}: {{asunto}}
> **Body (ES)**: `<p>Te han etiquetado en el ticket <strong>{{numero}}</strong>.</p><p>{{asunto}}</p><p>Lo sigues desde tu bandeja, en «Observo».</p><p><a href="{{enlace}}">Abrir el ticket</a></p>`
> **Body (EN)**: `<p>You were mentioned in ticket <strong>{{numero}}</strong>.</p><p>{{asunto}}</p><p>You are following it from your inbox, under “Watching”.</p><p><a href="{{enlace}}">Open the ticket</a></p>`

### 5.2 The four account emails

| Key | Who receives it |
| --- | --- |
| `auth.invitation` | The person who has just been given an account |
| `auth.recovery` | The person who has requested a reset —or has had one launched for them— |
| `auth.passwordChanged` | The person whose password has changed |
| `auth.accountCreated` | The person whose account has just been created from a ticket and who can now sign in with their organization account |

**`auth.invitation`** — markers: `{{nombre}}`, `{{enlace}}`, `{{caducidad}}`

> **Subject (ES)**: Establece tu contraseña de Catalina Support
> **Subject (EN)**: Set your Catalina Support password
> **Body (ES)**: `<p>Hola {{nombre}}:</p><p>Ya tienes cuenta. Establece tu contraseña desde este enlace, que caduca en {{caducidad}}.</p><p><a href="{{enlace}}">Establecer mi contraseña</a></p>`
> **Body (EN)**: `<p>Hi {{nombre}},</p><p>Your account is ready. Set your password from this link; it expires in {{caducidad}}.</p><p><a href="{{enlace}}">Set my password</a></p>`

**`auth.recovery`** — markers: `{{nombre}}`, `{{enlace}}`, `{{caducidad}}`

> **Subject (ES)**: Restablecer tu contraseña
> **Subject (EN)**: Reset your password
> **Body (ES)**: `<p>Hola {{nombre}}:</p><p>Alguien ha pedido restablecer tu contraseña. Si has sido tú, hazlo desde este enlace; caduca en {{caducidad}}.</p><p>Si no has sido tú, no hagas nada: sin abrir el enlace, tu contraseña no cambia.</p><p><a href="{{enlace}}">Restablecer mi contraseña</a></p>`
> **Body (EN)**: `<p>Hi {{nombre}},</p><p>Someone asked to reset your password. If it was you, use this link; it expires in {{caducidad}}.</p><p>If it was not you, do nothing: your password does not change unless the link is opened.</p><p><a href="{{enlace}}">Reset my password</a></p>`

**`auth.passwordChanged`** — markers: `{{nombre}}`, `{{cuando}}`, `{{ip}}`

> **Subject (ES)**: Tu contraseña ha cambiado
> **Subject (EN)**: Your password has changed
> **Body (ES)**: `<p>Hola {{nombre}}:</p><p>Tu contraseña ha cambiado el {{cuando}}, desde la dirección {{ip}}.</p><p>Si no has sido tú, avisa a Soporte cuanto antes.</p>`
> **Body (EN)**: `<p>Hi {{nombre}},</p><p>Your password changed on {{cuando}}, from the address {{ip}}.</p><p>If it was not you, tell Support as soon as you can.</p>`

**`auth.accountCreated`** — markers: `{{nombre}}`, `{{enlace}}`

> **Subject (ES)**: Tu cuenta en la mesa de ayuda
> **Subject (EN)**: Your help desk account
> **Body (ES)**: `<p>Hola {{nombre}}:</p><p>Se ha creado tu cuenta en la mesa de ayuda. Ya puedes entrar con tu cuenta de la organización.</p><p><a href="{{enlace}}">Entrar en la mesa de ayuda</a></p>`
> **Body (EN)**: `<p>Hi {{nombre}},</p><p>Your help desk account has been created. You can now sign in with your organization account.</p><p><a href="{{enlace}}">Open the help desk</a></p>`

## 6. The send

- **`net/smtp` from the standard library**, with no dependencies (like Calibyou), with the
  outgoing mail settings that live in `installation_settings.smtp_*`
  (`docs/arquitectura.md`, section 9).
- **HTML with an automatic text version**: a two-part email is sent, and the text part is
  taken from the HTML by stripping the tags. A client that does not render HTML shows something readable instead of
  nothing.
- **Simple HTML**: text tags —paragraphs, bold, lists, links— and **inline styles**.
  No external images, no third-party fonts and no style sheets: mail clients do not
  load anything from outside, and loading it would leave a trace of who opens the email.
- **In the background**: the send does not block the request that triggered it
  (`docs/arquitectura.md`, section 9). An email takes as long as the SMTP server takes, and whoever creates
  a ticket does not have to wait for it.
- **A mail failure does not bring down the action**: if the ticket is created and the notice does not go out, the ticket
  exists. The failure goes to the log with `go-logs`, with the template key and the recipient.
- **Every send is logged** —key, language, recipient and result— and **in no
  table**; the recipient address is written, by decision of the project owner of 2026-09-23
  (`AGENTS.md`, rules for agents). Never the email content or the link it carries: same as the sign-in attempts (`docs/modules/auth.md`, section 12).
- **The language is decided by the recipient**, not by whoever triggers the email: if a team is notified
  to three people and one reads in English, that one receives the text in English. It is three sends, not one.
- **The settings test really sends**: it opens the same connection as the normal send, authenticates
  and completes `MAIL FROM`, `RCPT TO` and `DATA` with the sender and the recipient written. The message
  is fixed and does not belong to `mail_templates`; the password is never logged.

## 7. The endpoints

| Method and path | What it does | Who |
| --- | --- | --- |
| `GET /api/mail/templates` | The twenty-eight templates: key, language, subject and body | Administrator |
| `POST /api/mail/templates/{key}/{language}/preview` | The subject and the body **already rendered**, with sample data. Accepts the **draft** being written | Administrator |
| `PUT /api/mail/templates/{key}/{language}` | Saves a template. Validates the markers | Administrator |
| `POST /api/mail/templates/{key}/{language}/reset` | Restores the factory text | Administrator |
| `POST /api/mail/templates/{key}/{language}/test` | Sends a test **to the email of whoever requests it** | Administrator |

- **The test goes to your own email**, not to an address that is written: that way the endpoint cannot be used
  to send emails to third parties with the institution's template.
- The test carries sample data and says so in the subject (`[prueba]`), so that nobody confuses a test
  email with a real one.
- **There is no endpoint to send in general**: emails are sent by the system, not by a person.
- **The preview is rendered by the backend**, and the frontend does not replace it: marker substitution
  and escaping live in a single place (section 4), and the day they change, what is seen before saving has to change
  too. It is **the same engine** as the test email, with the same
  sample data: what is seen on the screen is what will go out.

## 8. The error keys and what each endpoint returns

The keys, with the HTTP code they travel with. The frontend translates them; none needs more
context than its own key:

| Key | When | Code |
| --- | --- | --- |
| `mail.template.notFound` | The template (key + language) is not in the table | **404** |
| `mail.template.unknownKey` | That key is not any of the fourteen emails | **404** |
| `mail.language.unknown` | The language is not `es` or `en` | **404** |
| `mail.subject.required` | The subject is empty | **422** |
| `mail.body.required` | The body is empty | **422** |
| `mail.marker.unknown` | The text uses a marker that email does not accept | **422** |
| `mail.marker.missing` | On sending, the value of a marker that the text uses is missing | **422**, and on the real send it only goes to the log |
| `mail.test.noEmail` | A test is requested from an account without email: it is the case of the factory account `admin`, the only one without an address | **422** |
| `mail.smtp.notConfigured` | There is no mail server configured | **503** |
| `mail.smtp.*` | The mail server does not respond or rejects the send | **503**, with the detail in the log and never in the response |

- **An SMTP failure is not explained in detail**: `mail.smtp.notConfigured` and `mail.smtp.*` are seen as
  they are, because they are installation settings and whoever sees them is an administrator. What never
  comes out is the server's response, which goes to the log.

What each endpoint returns when it goes well. **It always returns the affected template**, so that
the editor does not have to request it again or duplicate rules:

| Endpoint | Returns |
| --- | --- |
| `GET /api/mail/templates` | `{"templates": [ … ]}` with the twenty-eight |
| `PUT /api/mail/templates/{key}/{language}` | `{"template": { … }}` and, if an essential marker is missing, `"missing": ["enlace"]` next to it |
| `POST /api/mail/templates/{key}/{language}/reset` | `{"template": { … }}` with the factory text already in place |
| `POST /api/mail/templates/{key}/{language}/test` | `{"sentTo": "quien@demo.com"}` |
| `POST /api/mail/templates/{key}/{language}/preview` | `{"subject": …, "body": …, "text": …}` with the sample data already in place |

And **what each template carries**, which is what the editor needs to paint itself with a single call:

| Field | What it is |
| --- | --- |
| `key`, `language` | Which one it is |
| `subject`, `body` | Its text today |
| `edited` | Whether its text **is different** from the factory one. The text is compared, who touched it is not looked at: that way an edit by the factory account —which leaves no identifier— also counts as edited |
| `updatedAt` | When it was last edited |
| `markers` | The markers that email accepts, each with `name` and `essential` |
| `missing` | The essential ones that **do not** appear today in its text: it is the editor's warning |

- **The warning is computed by the backend**, not by the frontend: the list of what is essential in each
  email lives in a single place (section 4), and duplicating it in the frontend is the sure way for them to one day
  say different things.

## 9. What was decided when reviewing this document

| # | Decision | It was left like this |
| --- | --- | --- |
| 1 | **The editor** | A text field with buttons that wrap the selection (bold, italic, list, link), **preview** next to it, and **accepts hand-written HTML** |
| 2 | **The markers** | It is only validated that **there are no invented markers**; they are not required to all be there |
| 3 | **Restore** | Yes, each template can go back to its factory text |
| 4 | **The register** | Informal «tú», in both languages |
| 5 | **The format** | **HTML** with an automatic text version, with no images or external fonts |
| 6 | **The footer** | None: the content and the link |
| 7 | **The send record** | Only in the log, with no table |
| 8 | **The recipients** | They are resolved by **whoever requests the send**, not `mail`. Ticket distribution is a matter of `docs/modules/tickets.md` |
| 9 | **The missing warning** | When a template is missing an essential marker —the link, the number—, **it warns without preventing saving** |
| 10 | **The test email** | It carries sample data and the subject starts with `[prueba]`, so it is not confused with a real one |
| 11 | **`PUBLIC_APP_URL`** | It is the variable all the links come from, and it goes in the envs and in `docs/arquitectura.md` (section 9) |
| 12 | **The error keys** | Documented in section 8, with their code: 404 for what does not exist, 422 for what is not valid and 503 for the mail server |
| 13 | **What the endpoints return** | Always the affected template, so the editor does not make a second call |
| 14 | **The warning on save** | The backend returns it in `missing`: the rule of which marker is essential is not duplicated in the frontend |
| 15 | **The editor screen** | It belongs to the frontend `mail` module and Settings links to it: that way the hard modularity rule is respected without exceptions |
| 16 | **The test from the factory account** | It cannot be done: that account has no email and 422 is returned with `mail.test.noEmail`. The destination is still always that of whoever requests the test |
| 17 | **The preview** | **It is rendered by the backend**, with a fifth endpoint and the same sample data as the test: a single version of marker substitution and escaping |
| 18 | **The two languages** | **Two columns on PC** —the same email in Spanish and English, side by side— **and stacked on mobile**: two columns do not fit on a phone |
| 20 | **An email that does not go out stays in the log** | **Finding and correction of 2026-09-27**: notices are sent in the background, and the `Send` error **nobody was looking at** —the action that triggered it had already finished—, so a badly placed template or a downed mail server **did not warn anyone, not even the log**. It really happened: the tagging template existed in the migration, the closed list in `markers.go` did not know it, `Send` rejected it and **the email did not go out without leaving a trace**. Now the failure is logged, with the key, the language and **the recipients** —the address is logged, the email content never—. The closed list stays in sync with the catalog. |
| 19 | **The markers** | The ones for that email are shown, with the essential ones marked, and **they are inserted when clicked** at the point where the cursor is: nobody has to memorize `{{solicitante}}` or write it correctly |

## 10. What this document enables

With `docs/modules/mail.md` approved, **`mail` is the first module to be implemented**, because `auth`
needs it to send the invitation email. The work order I propose:

1. The `mail_templates` table and the twenty-eight texts seeded: twenty-six in `migrations/v1.0.0.sql`
   and the two rows of `auth.accountCreated` in `migrations/v1.0.1.sql`.
2. The filling of markers and the escaping, with their tests: it is the part that breaks most.
3. The send over SMTP with both versions, and the log entry.
4. The four endpoints.
5. The editor on the Settings screen.

## 11. Implementation: templates translated when the global language changes

The twenty-eight rows are kept —fourteen per language— as editable memory of both languages, but
the sends always choose the version of the global language. `Send` and `SendAsync` stop receiving the
recipient's language and resolve it from the installation on every send.

When a global change is requested, `mail` takes as source the fourteen customized templates of the active
language. Before sending them to `ai`, it replaces HTML tags and `{{…}}` markers with protected
elements. The AI translates only the text; `mail` rebuilds the HTML and rejects a draft if it loses,
adds, changes or invalidly moves a marker. No draft is saved yet.

If the destination keeps the factory text, the Administrator reviews source and draft. If the destination
was already customized, they review three versions: active source, AI draft and existing destination, and
can choose one or combine them in the editor. A customization is never replaced without showing it.
On confirming, a transaction saves the fourteen destination templates and the new global language. If one row or
validation fails, no template changes and neither does the language. The translation history does not create an
additional table: the two current versions remain and `updated_at`/`updated_by_id` say who
confirmed the last one.

The translation uses the active AI and therefore cannot start if the installation is in the required-AI state. No
recipients, addresses or real ticket data are sent: only subjects, bodies and template markers. Restoring and
manual editing remain available.

With an external provider, before translating the twenty-eight requests are shown —subject and body
of each template— and an estimate of tokens and cost. If the Administrator does not accept it, no
call is made: the fourteen destinations remain available for manual editing. **Apply language**
is enabled when the fourteen destination templates pass their validation, whether they come from the AI, from the existing
destination or from manual editing.

The project owner chose AI translation with a reviewable draft and atomic application on 2026-10-06.
In the review, they decided to keep and compare any customized destination, with explicit choice or mixing.
On 2026-10-07 they corrected that commercial confirmation is limited to these templates and that
rejecting it keeps the manual path. No open decisions remain; correction explicitly approved.

### 11.1 Verified implementation

`POST /api/mail/language/drafts` returns the fourteen comparisons without persisting; it protects tags and
markers and requires confirmation before the 22 calls of a provider. Manual mode returns the
editable destination without calling the AI. `POST /api/mail/language/apply` validates the eleven and, in a single
transaction, saves templates, language and global version. The manual walkthrough was verified on PC and mobile.

## 12. Implementation: language validated in each template draft

Every translated subject and body goes through the central validation described in `docs/modules/ai.md`,
section 15, in addition to the current protection of HTML and markers. The expected language is the new
global language of the deployment. The original content is delivered literally; only the instructions
and labels added by the application use the destination language.

If a clear discrepancy persists after the single retry allowed for `local` or `remote`,
that template keeps its previous draft and shows `ai.invalid`. In `provider` the first
breach produces the same error without another billable request. The other templates are not
partially applied: the Administrator can manually correct any draft and the existing
transaction requires that the eleven be valid before changing the global language.

The twenty-eight template rows, markers, commercial confirmation, cost estimate,
endpoints and manual editing do not change. The tests cover subject and body, preservation of the previous destination,
later editing and the absence of a second call to the provider. The project owner chose 6A, 7A, 10B,
11A and 12A; no open decisions remain.
### 12.1 Implementation and verification

The translation service reuses the central validator of `ai`; therefore it only delivers to the draft flow a text compatible with the destination language. Faced with `ai.invalid`, the existing flow does not partially apply the templates and keeps manual editing. The focused tests and the full backend suite passed on 2026-10-13.
## 13. Implementation: missing template

Reading by key and language queries at most one template and distinguishes zero rows by
`RowsAffected`. It keeps the same domain error and does not change the editor, the fourteen pairs, the global
language, the preview or the send.

## 14. Critical-priority notice

The critical-priority amendment added the twelfth key and at that time left twenty-four rows because of its English and Spanish versions. That key is editable, restorable, previewable and
can be sent as a test from the editor, like the others. It uses the current global language.

The email includes number, subject, requester name, priority and public link to the ticket. It does not
include description, comments or attachments. Its recipients are all the active Support
accounts, without duplicates. A failure returns information to the tickets module, but never undoes the
saved priority.

The implemented key is `ticket.priorityCritical`; it is part of the same editable,
restorable and previewable catalog as the other eleven. After adding the private notice, the migration leaves
twenty-eight rows; the backend and end-to-end tests checked a single message even if the ticket goes down and back to
critical.

## 15. Private comment notice

The catalog adds a thirteenth editable key in the two global languages to notify a
new private note on the main ticket. The recipients are the assignee and the current technical
watchers, active and without duplicates; the author is excluded even if they belong to those sets. The
requester is never a recipient.

The email contains number, subject, author name and link to the main ticket. It does not contain body,
fragments, mentions or attachments. It uses the public address and the same rules of send, preview,
restoration and test as the current catalog. A mention also keeps its specific notice,
but the deduplication of the private notice avoids two copies of this new template for the same account.
A mail failure does not revert the comment.

## 16. SMTP modes and controlled unencrypted authentication

The transport stops being summarized in an ambiguous boolean and adopts three closed modes:
`starttls`, `tls_direct` and `plain`. STARTTLS opens SMTP, requires the server to announce the extension and
upgrades to TLS before authenticating. Direct TLS negotiates TLS from the first byte. `plain` keeps SMTP
unencrypted for legacy internal relays that only announce `AUTH LOGIN`, like the one used by SIGOC.
There is no automatic fallback from a secure mode to `plain`.

`plain` can only authenticate when the administrator has chosen it and has expressly accepted
that user, password and content will travel unencrypted. The acceptance is stored next to the
settings, is annulled when leaving the mode and is requested again if it is reactivated. Before testing and
before each send the host is resolved: all its addresses must be private IPv4 RFC1918 or
loopback. The socket is opened against one of the already checked addresses, keeping the host name
for SMTP, so that a second resolution does not bypass the restriction. IPv6 and public destinations
are rejected in this mode.

The settings test uses `MAIL FROM`, `RCPT TO` and `DATA` with the chosen mode. The secure modes
keep certificate validation. `plain` does not change templates,
recipients, retries or failure handling.

The existing settings migrate `smtp_secure=true` to `tls_direct` and `false` to `starttls`; therefore,
an update never silently turns a real installation into plain text. The
only exception is Mailpit in the development data: it is declared as `plain`, without
credentials, with explicit development acceptance. That exception is not generalized to production.

The sender must be an institutional address of Catalina's own. This proposal does not reuse
SIGOC's identity, does not incorporate its credentials and does not fix a specific address that the organization must still
assign.

Closed decisions: 1A three modes; 2A explicit acceptance; 3A private network only; 4A own
sender; 5A safe migration; 6A RFC1918 and loopback; 7A all resolutions private; 8A acceptance persisted and revocable; 9A test without email; 10A Mailpit-only exception in development.

The real verification of 2026-10-19 connected to `relay.quito.gob.ec:587`, checked its private
destination, completed `AUTH LOGIN` and closed without `MAIL FROM`, `RCPT TO` or `DATA`. The STARTTLS mode rejected
the same relay for not announcing the extension and direct TLS correctly rejected its SMTP greeting.

## 17. Real test send of the SMTP settings

The settings test stops ending after authenticating: it sends a real message to the
address written in **Send test to**. There is a single button, **Send test email**, both in
step 4 of `/setup` and in the mail card of Settings. The recipient is mandatory for
that action, is validated as an email and is not saved in the database or in the browser.

The message uses exactly the server, mode, user, secret, sender name and sender email present in
the form. In Settings, changes not yet saved can be tested:
an empty secret reuses the saved one and a written one is used only for that test. That way it is also checked
that the relay authorizes the intended `From:`, not only the credentials.

The subject and the body are fixed, not editable, and are chosen with the global language of the installation.
They identify the installation, the test nature and that there is no need to reply or perform any
action. No rows are born in `mail_templates`: this technical message lives with the test service and does not
form part of the editor of functional notices.

A final `250` from the server produces the precise message **El servidor aceptó el correo para
{destinatario}**. The interface does not claim delivery to the inbox, reading or passing through later filters.
Rejections of sender, recipient or data keep their technical error on the server and offer
an operational explanation without credentials in the API.

During `/setup`, each combination of source IP and recipient allows five attempts in fifteen
minutes and the sixth responds `429`; the seal continues to close the endpoint completely. In
Settings the route requires administrative capacity. The logs note result, server, IP and
partially hidden recipient; never password, body or content of the SMTP session.

## 19. Account email created from a ticket

> **Approved, implemented and verified on 2026-10-30.** When Support opens a ticket for a person who is in AD and does not
> have an account here, the account is created from the directory and they have to be told. Template
> number fourteen is born, `auth.accountCreated`, which carries no password or token: the person signs in with their
> organization account.

- **Key:** `auth.accountCreated`, in Spanish and English.
- **Markers:** `{{nombre}}` and `{{enlace}}`; neither is a token. `{{enlace}}` is the public address
  of the installation, to open the sign-in.
- **When it is sent:** only when the ticket's creation **creates** a new account from AD. It is not
  sent if the account already existed or if only a local account has been linked; in those cases there is no
  creation to report.
- **When it is not sent:** in the local creation of a requester account without access. That account
  cannot sign in and the normal invitation email —with a password link— does not apply.
- **If the email fails:** the ticket and the account remain created; the failure stays in the log with the
  key, the language and the recipient, without content. The screen warns that the account was created,
  but not that the email has arrived.
- **Migration:** the two rows go into `v1.0.1.sql`, with `ON CONFLICT DO NOTHING`. `v1.0.0.sql` is not rewritten;
  the application and the schema move to **1.0.1**.
- **Editor:** the list goes from thirteen to fourteen emails and from twenty-six to twenty-eight rows. The
  marker catalog and the validation of language drafts are updated with the same
  key.
- **Tests:** the template exists in both languages, the rendering replaces `{{nombre}}` and
  `{{enlace}}`, the key catalog includes it and the test send does not use it. Creation from
  AD sends once; an existing or linked account sends nothing. The isolated database was left with
  28 rows and the two of `auth.accountCreated`, and the template validation passed.

Sending the test is optional and does not block progress or sealing. Closed decisions: 1A destination
field; 2B a single button that sends; 3B optional; 4A fixed text; 5A wizard and Settings; 6A
public limit; 7A ephemeral destination; 8A written sender; 9A precise SMTP acceptance; 10A five in
fifteen minutes; 11A data not yet saved; 12A saved or temporary secret; 13A global language; 14A
bounded technical logs.
