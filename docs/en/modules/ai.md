# ai

[Español](../es/modules/ai.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-05
>
> **WSL2 support verified, 2026-11-05.** The NVIDIA variant in section 23 also works on Windows with
> Docker Desktop WSL2. It was tested with an RTX A2000 12 GB, driver 596.51, the
> `qwen2.5-1.5b-instruct` model active and real summaries generated. Linux with NVIDIA Container
> Toolkit remains the other supported path; native Windows with Windows containers does not.
>
> **QA correction, 2026-11-03.** When the engine does not answer and the card exhausts its re-reads,
> the summaries stop saying «Generating…» and show «No AI engine». Before, the pending state stayed
> visible forever.
>
> **Measurement and change of 2026-10-31.** Section 26 gathers the benchmark of models and threads.
> The current model is still the best balance. **The fixed `--threads 2` was removed**: the
> administrator lets `llama-server` choose and on this machine it uses **16 threads**, with a latency
> of 1.9 s against 4.0 s. The decision is implemented and covered by an administrator test; the
> verification with the real installation in development is still missing.
>
> **Correction approved, implemented and verified on 2026-10-22.** Section 25 records that the current
> contract still accepted a second insufficient proposal. The owner chose 1A–4A: structured prompt,
> 2–3 sentences of 25–60 words, contextual courtesy without a forced greeting, one retry and rejection
> of the second output if it is still deficient. The proposal was explicitly approved; only this scope
> is implemented, coordinated with `modules/tickets.md` and `interfaz-y-experiencia.md`.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 24 separates the public and
> internal versions of Summary and Last action so that private notes give the team context without
> leaking to the User. The owner chose and approved 8A, 24A and 25A.
>
> **Approved, implemented and verified on 2026-10-18.** Section 23 adds `auto|cpu|nvidia`
> acceleration, Auto recovery on CPU, transactional change and device attribution.
>
> **Correction approved, implemented and verified on 2026-10-16.** Section 22 starts from two failures
> of the comprehensive walkthrough: an old Last action could overwrite the recent one and two fields
> exhausted their attempts without autonomous recovery. The owner chose 2A, 3A and 4A. The current
> revision is protected in the database, there is a single deferred cycle and the sweep took once the
> errors of `CS-2026-0029`; the model failed again and they remained terminal, without a loop. The AI
> tests and the real repository passed. The backend suite keeps only the previous, unrelated failure
> of priority from an internal; manual buttons and periodic polling do not come back.
>
> **Amendment approved, implemented and verified on 2026-10-15.** Section 21 separates a stable
> `summary` of the original problem from `lastAction`, removes manual regeneration and migrates the
> affected contract to English technical names. The owner chose 1A–20B, corrected 5A to
> «Summary/Resumen», corrected 11A to require English in technical matter and chose 12B for
> `escalationReason`. No open decisions remain; no code is modified before approving the four
> coordinated documents. The migration was applied only in development; the specific AI and tickets
> tests, the 231 frontend tests and their build passed. The full backend suite keeps the previous
> failure of priority from an internal, unrelated to this amendment.
>
> **Correction approved, implemented and verified on 2026-10-15.** Section 20 starts from the failure
> reproduced in `INT-CS-2026-0016`, where «Last action» copied the reason even though the history
> ended in closure. The functional context includes principal and internal, fixes the target event,
> validates actor and action and has a deterministic fallback. The owner chose 1A–14A and explicitly
> approved the three coordinated documents; no open decisions remain.
>
> **Implemented and verified on 2026-10-15, with explicit approval from the owner.** Section 19 corrects
> the acceptance of summaries empty of meaning, protects the previous texts and measures the quality of
> professional paraphrases. The owner closed 1A–15A, with 8B; no open decisions remain.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section 18 sets the
> writing improvement in a professional tone, withdraws the tones from the contract and allows asking
> for a different paraphrase using only the previous proposal as text to avoid. The owner chose 1A–15A
> with 3B, 6B, 7A, 10B and 11B; no open decisions remain.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section 17 adds
> explicit deactivation without deleting the model and makes the manual test identify the model that
> produced the greeting.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section 16 completes
> the local administrator with an exclusive cancellable download, persistent phases, automatic
> activation and a manual test of the active model through a dated greeting. It does not change the
> catalogue or the real texts the module consumes.
>
> **Finding corrected and implementation verified on 2026-10-13: compliance with the global language.** The
> public walkthrough of block 8 got a «Reason» in Spanish with an English installation, even though the
> row was marked as `en`. Section 15 defines local validation, different treatment for engines without
> controlled cost and external providers, prompts in the deployment language and the resumption of the
> audit. The owner closed 1A–12A, with 10B. The proposal was explicitly approved and implemented on
> 2026-10-13.
>
> **Amendment implemented and verified on 2026-10-08: visibility by configuration.** Section 14
> distinguishes a configured engine from an available one and allows `tickets` to consult only that
> condition. The owner approved 1A–3A and confirmed keeping AI mandatory.
>
> **Correction approved and implemented on 2026-10-08.** During implementation it was confirmed that a
> sealed installation without AI configured cannot open tickets because AI is still mandatory. The
> owner decided to keep that rule. The false capability will be verified as a defensive contract,
> without removing the global block or promising a real walkthrough incompatible with it. The owner
> approved the correction and the tests reflect this boundary.
>
> **Amendment implemented and verified on 2026-10-08: writing assistant.** Section 13 describes the
> configured engine improving drafts from Support and Development, without storing prompts or replacing
> text without review. The owner approved decisions 1A–21A. Backend, frontend and the real walkthrough
> on PC and mobile passed; production remained outside the work.
>
> **Correction proposed on 2026-10-08.** While preparing the implementation it was found that publishing
> `/api/ai/**` for a tickets screen breaks the current modular boundary. The owner chose 20A: the
> public route belongs to `tickets`, which authorizes and delegates internally to `ai`. This
> documentary correction was explicitly approved by the owner.
>
> **Correction proposed on 2026-10-08.** Synchronous generation may use the 240 seconds that the engine
> client already allows, but the HTTP server cuts at 60. The owner chose 21A: extend only this
> operation to 270 seconds, without changing the rest of the routes. It required documentary approval;
> the owner approved it explicitly.
>
> **Finding corrected and verified on 2026-10-08.** The real validation of block 5 checked that 3B
> cannot replace 1.5B: the administrator requires 4 GiB free before stopping the process that occupies
> 1.80 GiB. It also found that the interface announces 1.5/2.8/5.5 GB while the catalogue requires
> 2/4/6 GiB. Section 12 incorporates 15A–22A. The owner approved the correction, which was implemented
> and validated with the three real GGUF files on 2026-10-08.
>
> **Finding and amendment proposed on 2026-10-07 (block 5).** The catalogue and the local administrator
> are implemented and pass tests with simulated files, but the real validation that this document
> requires for the GGUF files 1.5B, 3B and 7B does not exist yet. Section 11 defines the measurement in a
> disposable environment. The owner chose 1A–4A and 5A–9A; the review closed 10A–14A, with no open
> decisions. The owner explicitly approved the three documents on 2026-10-08.
>
> **Finding and amendment proposed on 2026-10-07.** The existing functional contract remains as-built,
> but four guarantees already promised by the architecture are missing: not forwarding credentials
> through redirects, supervising the local process, calculating effective resources and testing the four
> adapters. Section 10 incorporates options 1A, 2A, 3A and 4A chosen by the owner. The review added 5A,
> 6A, 7A and 8A and left no open decisions. The owner explicitly approved the amendment, which was
> implemented and verified on 2026-10-07.
>
> **Implemented and verified on 2026-10-07.** OpenAI, DeepSeek, compatible and Claude pass activation and
> summary tests against fake servers; redirects with credentials are rejected. The local engine restarts
> with growing waits, exposes health separately, respects cgroup and resumes partial downloads without
> keeping corrupt files.
>
> **Amendment proposed on 2026-10-06.** What exists remains as-built. The proposal turns AI into a
> mandatory part of the installation, allows choosing a local engine managed from the interface, a
> self-hosted server or an external provider, and generates each summary only in the global language. It
> replaces decisions 1, 2, 3, 8, 10, 11, 13, 14, 15, 17, 20 and 21 only on the points that this section
> identifies. The owner closed the functional decisions in three rounds and the later review. The owner
> explicitly approved the document and the implementation was verified by blocks.
>
> **Amended on 2026-10-01**: **the engine resources are corrected** so that they say what `ai.yml`
> really has, which is the source: the **memory cap is 1500m** (`mem_limit: 1500m`) and **there is no
> CPU cap** —the «2 CPU» statement that section 2 and decision 11 carried over came from nowhere in the
> compose file—; and **what is measured is ~1.44 GiB while running** —`docker stats` gives 1.441-1.448
> GiB and the process `VmRSS` 1.51 GB—, not the «~950 MB at rest and 1.09 GB with a long input» that had
> been noted before. The two old numbers are corrected in section 2 and in decision 11, and section 2
> also explains the **expected warning about the model volume** when starting the engine (that the
> volume `ai_modelos` appears with another project name: it breaks nothing and there is no need to
> «fix» it).
>
> **Amended on 2026-09-30**: the engine **is started separately** (`ai.yml`, like the test directory and
> Keycloak) and **is optional** —decision and correction of the owner—: without it the help desk works
> entirely, without the two summaries, and the first-start wizard **warns without blocking**, saying
> what to check (its container, and the model the first time). **Its address and its model are already
> configured from Settings** (decision 16 of `docs/modules/settings.md`), with its card and its «Test
> connection» button: this module reads them **on every request** through the `Configuracion` interface
> that it declares (section 5), and if nothing is set it uses `AI_URL` and `AI_MODEL` **as a fallback**
> (`AI_PALABRAS` and `AI_ESPERA_SEGUNDOS` still belong to the environment).
>
> **Becomes as-built on 2026-09-27**, the same day it was written: the module, the engine container, the
> `ai_insights` table, the two fields in the list and in the card, and their tests **are done and
> verified**. What was learned when putting it into operation —the engine takes as long as it takes, the
> model does not return the JSON unless it is required, and one in so many answers arrives half
> finished— is gathered in sections 2, 3 and 6 and in decisions 11 to 19.
>
> **Written on 2026-09-27**, when the owner asked for **two fields that summarize the ticket**: **«Reason»**
> —what the ticket is about— and **«Last action»** —what was the last thing that happened—, **both
> written by an artificial intelligence engine** that runs **in a separate container**, and **both in
> Spanish and in English**. He chose the five design decisions over my proposal, and **all the
> recommended ones went ahead**: a local engine with a small model, background generation with retry, a
> single call that returns both languages, recalculation with every ticket move, and the two fields in
> both ticket types (decisions 1 to 5).
>
> It is the **sixth backend module** and the only one that **has no screen of its own**: it is read
> inside the screens of `tickets`, which is who owns the tickets.
>
> **Approved by the owner on 2026-09-27.** It enables implementation.

## 1. What it does, in one sentence

When a ticket is born, is edited or moves —a comment, a status change, an escalation—, this module
**sends the ticket text to its own AI engine** and stores **two short write-ups**:

| Field | What it tells | Example |
| --- | --- | --- |
| **Reason** | **What the ticket is about**, in one line | «The user cannot sign in: the application has rejected their password since this morning» |
| **Last action** | **What happened last**, in one line | «The user was asked for a screenshot of the error; the ticket waits for their reply» |

**Both are stored in Spanish and in English**, and the screen shows the one for the language of whoever
is reading.

**They replace nothing**: they are a summary for reading the list at a glance. The ticket text is still
the ticket, and no decision of the system is taken from these fields.

## 2. The engine: its own container, and nothing leaves the server

| | |
| --- | --- |
| **Where** | A **`catalina_support_ai`** container, separate from the backend, with **llama.cpp** serving a model over HTTP |
| **Model** | **Qwen2.5-1.5B-Instruct**, quantized **Q4_K_M** (about 1.1 GB), chosen because it **fits in the memory of the original host**: it had no GPU and about 1.8 GB were free. The current development host does have an RTX A2000 and the NVIDIA variant is verified |
| **Image** | `ghcr.io/ggml-org/llama.cpp:server-b11206`, **with the build pinned** (and not the moving `:server` tag): a `pull` cannot change the engine underneath without anyone having decided it. The memory cap is **1500m** (`mem_limit: 1500m` in `ai.yml`) —**`ai.yml` does not limit the CPUs**— and it is what keeps it from eating the machine. What it uses, measured while running: **~1.44 GiB** —`docker stats` gives 1.441-1.448 GiB, **98% of the cap**, and the process `VmRSS` 1.51 GB—, just below it |
| **Startup** | `ai.yml` with a script (`config/ai/01-descargar-modelo.sh`) that **downloads the model only if it is missing** —to a temporary file, and checking the `GGUF` header before renaming it— and then does `exec` to the server. The `.gguf` lives in a named volume, so it survives `down`/`up` and **is not downloaded again** |
| **Network** | Its own network, **`catalina-support-ai`**, joined by the development and production backends. **A single container serves both**: two do not fit (decision 1) |
| **Port** | **None published on the machine**: the engine is not reached from outside, only through the internal docker network |
| **Internet** | The model **is downloaded once** when the container starts, and from then on **the engine works without internet access** |
| **Data** | The ticket text **does not leave the server**. That is why the engine is local and not a third-party API (decision 1) |

**And an expected warning when starting the engine**: `docker compose -f ai.yml up -d` may warn that the
volume `ai_modelos` «was created for project "catalina-support-ai" (expected "catalina-support")».
**It breaks nothing and there is no need to «fix» it**: the volume has a **fixed name** (`name: ai_modelos`
in `ai.yml`), so it is the same one even if the name of the project it was created with changes; that is
why the engine joins it and **the model is not downloaded again**.

**And it is slow, and that is said**: without a GPU, the engine reads the message at ~24 pieces per
second and writes at 5-9, so **each field takes between 12 and 24 seconds** and a long ticket —6,000
characters— takes around a minute and a half for the two. The second call for the same ticket is about
**seven times faster** (llama.cpp reuses the cached prefix). Two things come out of that, and they are
not decoration:

- **The client timeout is 240 seconds** (`AI_ESPERA_SEGUNDOS`), not the 30 of factory settings of any
  HTTP client: with less, exactly the long tickets would fail, and those are the ones that need it most.
- **The queue has two workers and the engine serves one at a time** (`--parallel 1`): the second worker
  queues inside the engine, and what cannot happen is that a wait for the engine blocks the creation of
  a ticket.

**«The engine is down» cannot be «the help desk is down»**: if the container is not there, the two
fields stay **without text and with their warning**, and everything else —opening tickets, commenting,
escalating, closing— works the same. That is decision 2 and it is the condition that rules over all
the others.

## 3. What it is sent and what it is asked for

**The text is assembled in `tickets`** —which is who owns the ticket— and is passed to this module
already made:

```text
Asunto: No puedo entrar a la aplicación
Descripción: Desde esta mañana me sale «usuario o contraseña incorrectos»…
Adjuntos: captura.png, informe.pdf

Conversación (de lo más antiguo a lo más reciente):
- user1 (comentario): Sigue fallando después de reiniciar.
- support1 (estado): en progreso
- support1 (comentario): ¿Puedes probar desde otro navegador?
```

- **The HTML is turned into text**: the engine receives no tags, no addresses and no `data-adjunto`.
- **Attachments go by their name**, never their content: saying «attachments: captura.png» gives the
  context that is needed, and sending the binary would ask for a model with vision and multiply the
  size.
- **The history goes in order, the most recent at the end**, and **if it has to be trimmed it is
  trimmed from the beginning**: the last thing that happened is exactly what has to be read. The cap is
  **6,000 characters** (decision 6).
- **The engine is asked for a JSON with the two languages** and nothing else (decision 3):

```json
{ "es": "…", "en": "…" }
```

- **And it has to be demanded in two ways at once** (measured, and it is not a detail): with
  `response_format: {"type": "json_object"}` in the body **and** with the demand written **at the end of
  the user message** —«Return EXACTLY this JSON object, with those two keys and no other field», followed
  by the template—. With only one of the two, a 1.5B model wraps the answer in a code block, invents its
  own keys or answers in a single language. Even so, the client **rescues the first JSON object** from
  the text (in case it comes wrapped) and **retries** if it is not valid.
- **The cap of pieces in the response leaves margin**: requested words × 4 + 160, and never more than
  half the model window: an answer cut in the middle of the JSON is an answer that is not valid.

- **What is asked is asked twice, with the same call and a different task**: the task of the «Reason»
  (what it is about) and the task of the «Last action» (what happened at the end). They are **two
  requests**, one per field, because they ask for different things (decision 7).
- **The temperature is low and the answer has a word cap**: a summary is sought, not literature. If the
  model returns something that is not the requested JSON, **it is retried once** and, if it fails again,
  the field stays in error (decision 8).

## 4. The table: `ai_insights`

One table, in `v1.0.0.sql` with its **catch-up** (1.0.0 is not closed, so a new version is not opened
for this):

| Column | What for |
| --- | --- |
| `id` | Key |
| `ticket_number` | The ticket number, `ACME-2026-0042` or `INT-…`, **exactly as it is read in the emails** |
| `kind` | `motivo` or `ultima_accion` |
| `state` | `pendiente`, `listo`, `error` or **`sin_motor`** (there is a configured engine and it does not answer, or there is no engine: the two are told apart on the screen) |
| `text_es`, `text_en` | The two write-ups, **in the same row** |
| `error_key` | The error key, when the state is `error` |
| `model` | The model that wrote it, to know what it was generated with |
| `requested_at`, `generated_at` | When it was requested and when it was written |
| `attempts` | How many attempts it has, which is what cuts the infinite retry. **It is counted also when it succeeds**: it is left said how many were needed |

- **One row per ticket and per field** (`UNIQUE (ticket_number, kind)`): the field is **rewritten** on
  every move, history is not accumulated. What matters is the latest, and storing every version would be
  filling the database with texts that nobody reads.
- **It is indexed by `ticket_number`**, which is how they are requested: the ticket list requests its own
  **in a block**, not one by one (decision 9). The index is redundant with `UNIQUE (ticket_number, kind)`
  —the first column already serves— and it stays because the document asks for it and it does not get in
  the way: a table with two rows per ticket is not going to grow because of this.
- **The table has real `CHECK`s**: `kind` only admits the two fields, `state` the four states and
  `attempts` cannot be negative. What the application must not write, the database does not accept
  either.

## 5. The contract between modules

**`tickets` declares what it needs and `ai` fulfils it without knowing it** —the same pattern it uses
with `users`—, and the one that joins them is `main.go`:

```go
// En tickets: lo que este módulo necesita del motor.
type Insights interface {
    // Pedir apunta un resumen y vuelve enseguida: no espera al motor.
    Pedir(entrada EntradaDeResumen)
    // De devuelve lo que hay, en bloque, para pintar la lista.
    De(numeros []string) (map[string]Resumen, error)
}

// En ai: el servicio que lo cumple.
func (s *Service) Pedir(entrada tickets.EntradaDeResumen)
func (s *Service) De(numeros []string) (map[string]tickets.Resumen, error)
```

**And on the other side, the engine**: `ai` declares what it needs from the installation settings and
`settings` fulfils it, **without either of the two importing the other**:

```go
// En ai: de dónde sale el motor que está puesto.
type Configuracion interface {
    AI() (AIDatos, error)
}

// El cableado lo conecta y traduce al tipo de cada módulo (main.go).
aiService.SetConfiguracion(instalacionesDeLaConfiguracion{settings: settingsService})
```

- **The configuration is read on every request**, not at startup: the address and the model from the
  database are the ones that rule every time a summary is queued, a job is resolved or the health of the
  engine is asked about. That is what makes changing them in Settings worth it without restarting
  anything.
- **If the configuration has no engine, the environment acts as a fallback** (`AI_URL`, `AI_MODEL`), and
  the module resolves it with its own settings: the configuration says **whether there is** an engine,
  not which one is the reserve (decision 21).
- **The connection test** (`Probar(url)`) is the one the Settings button asks for: it asks
  `<url>/health` with a three-second context and returns the error instead of a boolean. `settings`
  declares it in its `ProberDeIA` interface and this module fulfils it —just like `Disponible()`, which
  is the one the first-start wizard uses—.

- **The text is not stored**: the table has no text column on purpose —it would be copying the whole
  ticket into another table, and with it the data of the people—. That forces anyone who wants to request
  again a summary that was left half done on restart to **assemble the text again**: `ai` also declares
  `Fuente` (`TextoParaElMotor(numero)`) and **`tickets` fulfils it** with its method of the same name. It
  is the return path of the same `main.go` adapter.
- **`ai` does not import `tickets` and `tickets` does not import `ai`**: each one declares its side and
  `main.go` joins them (`docs/arquitectura.md`, section 4). It is the hard rule of modularity, and here it
  is what allows the engine to be turned off without touching the tickets module.
- **The `ai` module has no endpoints of its own.** What the user touches —asking for a field to be
  regenerated— enters through **`POST /api/tickets/{number}/insights`**, which belongs to `tickets`: the
  tickets screen can only talk to its own API, and the tickets module is the one that decides whom to ask.
  Without this rule, the frontend would have to call `/api/ai/**` from a tickets screen.
- **This module does not write in anyone else's tables**, and nobody writes in its own.

## 6. When it is requested, and what happens when it fails

| When | What is requested |
| --- | --- |
| The ticket is created | Reason and last action |
| The subject or the description is edited | Reason and last action |
| A comment is made | Last action |
| It changes status, is escalated, is reopened, is closed | Last action |
| Someone presses **«Regenerate»** in the card | The two |

- **Everything goes in the background** (decision 2): whoever creates a ticket **does not wait for the
  engine**. The creation responds when the ticket is saved, with the fields in **`pendiente`**.
- **An in-memory queue with few workers** (two), because **the engine is the expensive part**: without a
  queue, ten comments in a row would ask for ten summaries at once to a model that serves one at a time.
- **If the engine does not answer, it is retried three times** with growing waits (2 s and 4 s);
  afterwards the field stays in **`sin_motor`** with its key and **it is not retried on its own** —a loop
  of retries against a fallen container is noise—. It can be requested again with «Regenerate».
- **If it answers something that is not valid, it is also retried three times, right away** (without a
  wait: it is not a slow engine, it is a badly explained answer), and afterwards the field stays in
  **`error`** with `ai.invalid`. **Three and not one** because a 1.5B model is stochastic: measured, one
  in so many answers arrives half finished —only one language, or cut by the cap of pieces— and asking
  again is almost free. What there is not is an infinite loop.
- **When an answer is not valid, the log says its shape and never its content**: which keys arrived and
  how many characters each one had («keys: es=142 characters, en=missing»). Without that, «it answered
  something that is not valid» cannot be investigated; with that, it is known if English was missing, if
  it arrived empty or if it was cut in half. **The ticket text is never logged.**
- **One job per ticket and field**: asking for the same thing twice does not queue it twice; the new one
  replaces the one that was waiting, because the old text is no longer useful.
- **When the backend restarts, the queue is lost** and the fields that were `pendiente` stay that way:
  there is a **catch-up pass at startup** that asks again for what was left half done (decision 10).

## 7. Limits, and what it does NOT do

- **The content of attachments is not sent to it** —neither the binary of a PDF nor the pixels of a
  screenshot—: only their names.
- **It does not write in the ticket, does not comment, does not move statuses, does not send emails.** It
  only writes.
- **It decides nothing**: if the summary says one thing and the ticket says another, the ticket rules.
- **It is not used in the emails**: no notice carries these texts (the notices still use the number, the
  subject and the link, `docs/flujos.md`).
- **The English is not translated from the Spanish**: the two write-ups are requested at once, and the
  English text may say the same with other words. It is a summary, not an official document.
- **Quality depends on the model**: a 1.5B model writes short summaries correctly and **gets the nuances
  wrong**. It is the stated and accepted consequence of it fitting in this machine (decision 1), and that
  is why the fields **do not replace reading the ticket**.
- **One model per installation, not one per user or per language**: the engine does not know who is
  looking.

## 8. What was decided

| # | Decision | It stayed like this |
| --- | --- | --- |
| 1 | **Where the engine runs** | **Local container**, a **small model** (Qwen2.5-1.5B-Instruct Q4_K_M) **shared by development and production** over the `catalina-support-ai` network. **Decision of the owner, 2026-09-27**, and the recommended one: **nothing of the ticket leaves the server**, at the price of a more modest summary. **Two containers do not fit** —the original host had 1.8 GB free and the model occupies ~1.1 GB— and **at that time there was no GPU**. The alternative he discarded was an external API, better and more expensive in privacy. The current host does have a GPU and the NVIDIA variant was verified on WSL2 |
| 2 | **When it is generated** | **In the background and with retry**: the creation does not wait for the engine. This decision originally said that the application worked without configuring it; section 9 corrected it by making the initial configuration mandatory. A later fall of the engine does not prevent saving tickets: the field stays pending and is recovered when it returns |
| 3 | **The two languages** | **A single call** that asks for `{"es": …, "en": …}`: faster than two and without translation in between. If the JSON does not arrive, it is retried once |
| 4 | **When the last action is recalculated** | **With every move**: comment, status, escalation, reopening, closure, assignment and edition. That is what makes it useful: a text that is not updated stops being «the last action» as soon as someone writes |
| 5 | **Scope** | **The two ticket types**, and in their two lists. In an **internal**, the «Reason» is **the one of the escalation, which already exists** and is not generated; the «Last action» is. **The Administrator sees them** as he sees everything else, read-only, and **cannot regenerate them** (he does not write) |
| 6 | **How much text is sent** | Subject, description, **the conversation and the history**, with the HTML as text and **the attachments by their name**, trimmed to **6,000 characters** and **removing from the beginning**: the last thing that happened is what has to be read |
| 7 | **Two requests, not one** | «Reason» and «Last action» are requested separately, because they are different tasks: what the ticket is about, and what happened last. A single text for both things would come out worse in both |
| 8 | **What happens if the engine answers badly** | One retry, and afterwards the field stays in **`error`** with its key: **there is no loop**. The automatic retries of the queue (three, with growing wait) are for the fallen or slow engine, not for an answer that is not valid |
| 9 | **How they are read for the list** | **In a block** (`De(numeros)`), one query per page and not one per row: the ticket list is painted with a single question, like the accounts of the people |
| 11 | **The engine build is pinned** | `server-b11206` and not the moving tag: upgrading the engine is a decision, not something that happens one day when doing `pull`. And its arguments are written down: 4096 of context, two threads, one conversation at a time, 512 pieces of response, without web interface, **1500m of ceiling memory** —`ai.yml` does not limit the CPUs, corrected on 2026-10-01— and a health check with fifteen minutes of margin because of the first download |
| 12 | **The word cap is 40** | A column of a table is not a paragraph: with 60 words the «Reason» filled the cell. The screen, moreover, trims to two lines in the list and shows it whole in the card |
| 13 | **The engine timeout is 240 seconds** | Measured, a long ticket takes ~95 s for the two fields. With 30 s —the factory setting of the HTTP client— exactly the long tickets would fail. It is in `AI_ESPERA_SEGUNDOS` |
| 14 | **The JSON is demanded in two ways** | `response_format` in the body **and** the demand written at the end of the user message. With only one, the model wraps the answer, invents the keys or answers in a single language. The rescue of the first JSON object and the retry stay as a net |
| 15 | **The invalid response and the absent engine are treated differently** | `error` with `ai.invalid` and immediate retry for the first; `sin_motor` with `ai.unavailable` and growing wait for the second. **And they are compared with `errors.Is`**, not by the text: it is a finding of the tests —when enriching the error with its shape, the comparison stopped working and an invalid response would have been treated as a fallen engine— |
| 16 | **The text is not stored, and there is a way back** | `ai` declares `Fuente` and `tickets` fulfils it with `TextoParaElMotor`: it is what allows resuming at startup without copying the ticket into another table |
| 17 | **The catch-up does not touch the attempt counter** | If it reset it, restarting the backend would be the way to skip the cap of a fallen engine |
| 18 | **The repository tests go against real PostgreSQL**, inside a transaction that is undone | It is the only thing that really checks the `ON CONFLICT` of the table of one row per ticket and field. If there is no database, they are skipped |
| 19 | **The module depends on a store interface** | The repository fulfils it, and it is what allows testing the queue, the retries and the states without a database |
| 20 | **The queue has two workers, and the engine serves one at a time** | The engine has a single slot (`--parallel 1`), so the second worker queues inside it. Two stay on purpose: it is what prevents a ticket from being left waiting behind the queue of another in the application |
| 10 | **What happens when the backend restarts** | The queue lives in memory and **is lost**, so at startup **it is requested again** what was left in `pendiente` or with few attempts. What was in `error` **is not retried on its own**: that is what the button is for |
| 21 | **The engine comes from the configuration** | **Decision of the owner, 2026-09-30**: the address and the model are saved in Settings (decision 16 of `docs/modules/settings.md`) and this module reads them **on every request** through the `Configuracion` interface that it declares; the environment (`AI_URL`, `AI_MODEL`) remains **as a fallback**. The connection test —`Probar(url)` against `<url>/health`, with three seconds of cap— is consumed by `settings` through its `ProberDeIA` interface |

## 9. AI implemented: mandatory, local or external, in the global language

### 9.1 Findings that force amending what exists

1. AI is presented as a differentiator of the product, but today an installation can end up without
   configuring it and an empty `AI_URL`/`AI_MODEL` means «not integrated».
2. `ai.yml` downloads and serves a single fixed model; `/settings` only changes a text with the URL and
   the name, and cannot download, activate or delete models.
3. The fixed limit of 1500 MB prevents running the 3B and 7B local models chosen for the catalogue.
   Moreover, a single engine shared by development and production would allow one installation to change
   the model of the other.
4. The client only knows Chat Completions without authentication. OpenAI and DeepSeek can use that
   contract with credentials, but Claude needs an adapter for the Messages API.
5. Each request generates Spanish and English because the language belongs today to the account. When
   turning it into a global property, the second output is left over and lengthens the generation.
6. Jobs that reach `sin_motor` do not recover on their own. A mandatory AI tolerant to failures needs to
   resume them when the connection returns, without blocking the help desk.

### 9.2 What «mandatory» means

- A new installation **is not sealed** until there is an AI configuration tested through a valid real
  generation.
- An already sealed installation that is updated without valid AI enters the state **«AI requires
  configuration»**. The factory account or an Administrator reaches the AI configuration directly; the
  other accounts see that condition and do not enter the product. The public API of `/setup` is not
  reopened nor is the entry method touched.
- After configuring it, a fall of the engine **does not block** creating, commenting, moving or closing
  tickets. The summaries stay pending and the application shows the state to the Administrator.
- When the connection returns, a monitor detects and quantifies what is pending and what was left in
  `sin_motor`, and requeues it gradually in any of the three modes. It respects a maximum of attempts
  per cycle, growing waits and a single task per ticket and field; it does not create an unlimited loop.

### 9.3 The three modes

| Mode | Configuration | Adapter |
| --- | --- | --- |
| **Local** | Model from the catalogue installed in the environment's own volume | Chat Completions compatible with OpenAI, served by `llama.cpp` |
| **Self-hosted server** | Base URL, model and authentication: none, Bearer, key in a header or user/password | Chat Completions compatible with OpenAI |
| **Provider** | OpenAI, Anthropic Claude, DeepSeek or «compatible with OpenAI»; suggested model or written identifier; credential | OpenAI and DeepSeek: Chat Completions. Claude: Messages API. The generic one: Chat Completions |

The lists of external models are **suggestions**, not a closed list. An identifier can always be written
because providers change their catalogues without publishing Catalina Support. The URL of OpenAI, Claude
and DeepSeek comes from factory settings; the self-hosted server and the compatible provider ask for
theirs. The TLS check is not disabled.

### 9.4 Testing and activating is an atomic operation

The test stops consulting only `/health`. It uses an artificial text without data of people and demands
the same structured response as a real summary, in the global language and within the word limit. It also
checks authentication, model and format. Only a valid response allows activating.

When changing provider or model, the previous configuration keeps serving while the new one is tested. If
download, authentication, startup or generation fails, it is not replaced. The secret written in a failed
test is not persisted. The successful change saves provider and model and from that moment the new jobs
use them; the texts already ready keep in the table which provider, model and language produced them.

### 9.5 Local model administrator

`ai.yml` stops running the fixed script directly and starts an internal administrator as the main
process. That process supervises `llama-server`, but **does not use the Docker socket**. Its API only
lives on the internal network and demands `AI_MANAGER_TOKEN`, shared with the backend and absent from the
browser.

From `/setup` and `/settings` it is possible to:

- see the catalogue, which one is installed and which one is active;
- check space and memory before downloading;
- download to a resumable partial file, validate the size, the published sum and the GGUF header, and
  show progress;
- activate the validated file, restart only `llama-server`, wait for health and run the test generation;
- keep several models, one active, and delete only the inactive ones.

Each environment has its own administrator, volume and network. Development and production do not share a
model or configuration; on a small machine only the environment being worked with is started. The
container has a ceiling of 8 GB —not a reservation— and the administrator rejects an activation if the
available memory or disk does not meet the chosen profile.

### 9.6 Initial local catalogue

The three files are official, of the same family and Q4_K_M quantization. The current one remains as the
default.

| Model | Download | Engine memory | Visible recommendation |
| --- | ---: | ---: | --- |
| **Qwen2.5-1.5B-Instruct** | 1,117,320,736 bytes | 1,927,028 KiB RSS measured | 2 GiB available |
| **Qwen2.5-3B-Instruct** | 2,104,932,768 bytes | 3,566,068 KiB RSS measured | 4 GiB available |
| **Qwen2.5-7B-Instruct** | 4,683,073,632 bytes in two parts | not activated: the machine did not meet the minimum | 6 GiB available |

The process figures were measured with the 4096-token context used by Catalina Support. The disk
requirement is calculated live with what is already installed, the partial file and a margin, instead of
promising a fixed figure.

Sources verified on 2026-10-06: official repositories of
[1.5B](https://huggingface.co/Qwen/Qwen2.5-1.5B-Instruct-GGUF/tree/main),
[3B](https://huggingface.co/Qwen/Qwen2.5-3B-Instruct-GGUF/tree/main) and
[7B](https://huggingface.co/Qwen/Qwen2.5-7B-Instruct-GGUF/tree/main). Server compatibility
relies on the [official llama.cpp documentation](https://github.com/ggml-org/llama.cpp/blob/master/tools/server/README.md).

### 9.7 One language per installation

The task asks for a single object `{"text":"…"}` in the global language. It no longer generates `es` and
`en` in the same response. The table stores the text and the language that produced it; the ticket
responses return a single text. When the language changes, the texts already generated are kept and
remain visible with their original language. Only a move of the ticket or the individual action of
summarizing again replaces them, now in the new global language. There is no translation or regeneration
batch.

This reduces the output tokens and removes the failure «one language arrived and the other was missing».
It does not reduce the input text: the engine still reads the same ticket context.

### 9.8 Privacy and secrets

Before activating a self-hosted server or an external provider, the Administrator confirms that
**subject, description, comments, history and attachment names** will leave the server; when changing
the language, also the custom subjects and bodies of the templates it asks to translate. No files,
recipients, passwords, tokens or binaries leave. Who confirmed and when is recorded; during the wizard
the act of the installer is recorded without inventing an account, and in Settings the administrator
account is saved. It is not asked on every ticket.

The credentials are encrypted in an authenticated way before being stored, with a master key of the
environment's own. The API only returns whether a credential exists and never its value. The credential,
the body sent and the response of the provider are not recorded either. Losing or changing the master key
requires writing the credential again; no attempt is made to recover unreadable text.

`AI_CREDENTIAL_KEY` is mandatory when starting the backend in all environments and contains 32 random
bytes encoded in base64. Development carries a versioned local value; each production installation
generates and keeps its own. If it is missing or does not have the expected format, the backend does not
start. If a valid key was replaced and can no longer decrypt an existing credential, the backend does
start in the state «AI requires configuration» so that the factory account replaces it.

### 9.9 Recovery and cost

The recovery of pending or `sin_motor` summaries starts automatically in the three modes, including an
external provider: it does not create a new batch because the language changes and it respects the
normal caps of the queue. The economic confirmation is reserved exclusively for translating the fourteen
email templates. Before those twenty-eight requests, estimated tokens and cost are shown; if it is
rejected, the provider is not called and the templates are edited manually.

### 9.10 Licence and integrity of the local catalogue

Before downloading, each record shows official source, licence, size and checksum of the file. The first
download of each version requires accepting its licence and records model, version, checksum, date and
actor: «installer» in `/setup`, or the administrator account in `/settings`. If the version, licence or
checksum changes, acceptance is requested again. Acceptance does not replace the check of the downloaded
file.

### 9.11 Implementation and testing scope

`backend/modules/ai/**`, the contract that `tickets` consumes, the persistence of `ai_insights`, the
configuration that `settings` delivers, `ai.yml`, `config/ai/**` and the screens of `/setup` and
`/settings` change. The tests cover the four adapters with fake servers, encryption and concealment of
secrets, atomic change, interrupted and resumed download, activation, deletion, recovery after a fall,
prior confirmation of the commercial translation of templates, licence acceptance and preservation of
summaries when changing the language. No automatic test consumes commercial APIs.

Billing, quotas and purchase of provider credits are left out; GPU and layer selection; custom local
models outside the catalogue; attachments with vision; streaming; translation of tickets; and choosing a
model per user, team or ticket.

### 9.12 Record of decisions and review of the proposal

The owner decided on 2026-10-06: mandatory configuration with operation tolerant to failures; its own
step in `/setup`; three modes; OpenAI, Claude, DeepSeek and compatible with OpenAI; local catalogue
1.5B/3B/7B; complete administration from the interface; several models installed and one active;
encrypted secrets; external suggestions with free identifier; test through real generation and atomic
change; existing update blocked until configured; privacy confirmation; four authentications for a
self-hosted server; automatic local recovery and recovery confirmed by the Administrator when it can
generate charges; 8 GB ceiling; isolation per environment; and internal API protected with a token.

The review of 2026-10-06 added: mandatory encryption key and different per production; comparison of a
translation with the custom destination; commercial recovery stopped until the Administrator sees the
volume and accepts the possible cost; licence, source, size and checksum before downloading; and updating
open sessions by means of the language version. No open decisions remain. The proposal was explicitly
approved.

## 10. Implemented robustness of adapters and local engine

The external adapters reject any 3xx without following it. The rule applies in the activation and in
every summary. The four providers are covered by name, although OpenAI, DeepSeek and compatible share a
protocol; Claude keeps its own body and headers. Fake servers check that a token never reaches a second
server.

The administrator supervises `llama-server`: after an unexpected exit it restarts the active model
indefinitely with waits of 1, 2, 4… up to 60 seconds. Activating another model cancels that cycle and
waits for the previous process before starting the new one. The catalogue exposes `active` and `healthy`
separately and keeps the last operational error.

A resumption reserves only what is pending and 256 MiB of margin. An impossible partial file is
discarded; if its size is correct but format or checksum fail, it is also deleted to restart clean. RAM
is decided with the lowest value between the machine and the effective cgroup. These rules are tested
with simulated files and limits, without downloading real models or depending on the RAM of the machine.

Public downloads follow at most five redirects and only between HTTP/HTTPS addresses. Generation calls
with credentials follow none.

Catalogue, modes, authentication, prompts, queue, costs and public interface do not change. No Docker
socket or advanced supervisor configuration is added.

The owner closed the review on 2026-10-07 with decisions 5A to 8A. No open decisions remain in this
amendment.

## 11. Real validation of the local catalogue

### 11.1 Catalogue and capacity of the machine

Before transferring a GGUF, its URL, licence, name, size and checksum are compared with the official
publication. Any difference stops that model: it is recorded as a finding and requires an approved
amendment of the catalogue; a file is not trusted just because it matches the local data, nor are
metadata updated automatically.

The validation machine has about 5 GiB of RAM available and 68 GiB of free disk when opening the block.
The three models are downloaded and verified. 1.5B and 3B are activated; the 7B is not activated because
its profile requires 6 GiB. After collecting the evidence, 3B and 7B are deleted and no binary of the
disposable environment is kept.

### 11.2 Measurements

For 1.5B and 3B the following are recorded:

1. download bytes and time, final checksum and result of the GGUF validation;
2. time from requesting the activation until `active` and `healthy` are true;
3. memory of the `llama-server` process and of the container once healthy;
4. three generations per model, with individual time, minimum, average and maximum;
5. controlled killing of `llama-server` only, retry sequence and time to recover health, with a maximum
   of five minutes.

The inputs are two controlled fictional tickets, one in Spanish and one in English. A generation passes
if it returns valid JSON, respects the global language, completes `reason` and `lastAction` and meets the
length limits. The editorial appreciation is noted, but it does not change that objective result.

### 11.3 Result and defects

The report keeps reproducible commands and a table per model, without storing the GGUF files. A failure
of catalogue, download, activation, generation, resources, interface or recovery is documented before
being corrected. If it requires changing behaviour or code, an amendment is added within this block and
another explicit approval is requested. The validation does not authorize early corrections.

GPU, commercial provider, self-hosted server, prompt tuning, custom models, subjective comparison of
quality and production are left out. The owner chose 1A–4A and 5A–9A; in the review he chose 10A–14A. No
open decisions remain.

### 11.4 Result measured on 2026-10-08

The metadata and SHA-256 of the three models matched Hugging Face. The downloads took approximately 58 s
(1.5B), 113 s (3B) and 219 s (7B). The 1.5B became healthy in about 5 s and the change 1.5B→3B, including
the release of the previous process, took 28.46 s.

Each activated model produced twelve responses: Spanish and English, `reason` and `lastAction`, three
times per combination. The 24 were valid JSON, respected language, field and limit. In 1.5B the requests
took between 1.69 and 5.86 s, with an average of 3.62 s; in 3B, between 2.88 and 11.01 s, with an average
of 5.22 s. Two initial `lastAction` responses of the 1.5B were editorially imprecise, although they met
the objective contract; the prompt was not changed because that comparison was left out.

When only the child ended, the 1.5B recovered health in about 4 s and the 3B in 9 s, within the maximum
of five minutes. The 7B was verified but not started: it required 6 GiB and the recoverable capacity
observed with 3B active was approximately 5.07 GiB.

## 12. Safe change between models implemented

Before stopping the active model, the administrator adds the effective available memory and the real RSS
of the `llama-server` process. If that recoverable capacity does not reach the requirement of the
destination, it rejects the activation without interrupting the healthy model. That is how the 7B behaves
on the machine of block 5.

If the calculation allows trying it, it saves the previous model, stops its process and measures the
effective memory again. Only then does it start the destination. If memory is missing, the startup fails
or the new process does not reach health, it restores the previous one and waits for it to become healthy
again before answering. The persistent marker `active` changes only when the destination is healthy; a
failed activation does not turn the destination into the active model.

The lack of memory is returned with the requirement and the availability in bytes, plus a stable key.
Backend and frontend translate it to GiB to explain in the toast how much the model requires and how much
is available. The catalogue and the selectors show the same minimums applied by the administrator: 2, 4
and 6 GiB.

The tests cover: change 1.5B→3B using recoverable RSS; second insufficient measurement and restoration;
health failure and restoration; rejection of 7B without terminating 1.5B; and structured error with
memory required/available. The official GGUF files already verified are reused and the validation is
resumed from the activation of 3B. Production, GPU, forced swap and catalogue changes are left out. The
owner chose 15A–18A and closed the review with 19A–22A.

## 13. Implementation: writing assistant

### 13.1 Contract

The module does not publish a route of its own for this screen. `tickets` hands it internally the draft
already authorized, the contextual name, the ticket type, the language and a closed tone:
`professional`,
`friendly`, `brief`, `empathetic` or `technical`. The backend obtains from the database the ticket and the
contextual name: the person who requires support in the principal and the Support person in the internal.
The client does not send that name and cannot replace it.

The input to the engine contains only draft, tone, global language, contextual name and whether the ticket
is principal or internal. It does not include subject, previous description, conversation or attachments.
The prompt asks for clarity, spelling and the chosen tone; it forbids inventing information and demands
keeping language, names, numbers and technical details. The name is incorporated only when it comes out
natural.

The response is a single plain text with line breaks. It does not generate HTML, rich format, mentions or
attachments. An empty result, one that exceeds the field limit or one that is not valid text is rejected
with a stable key and never modifies the editor.

### 13.2 Authorization and data

Only Support and Development can call the endpoint. Besides the role, the backend demands access to the
ticket and permission to write in that specific editor; hiding the button does not replace this rule. An
empty draft or one that exceeds the current limit of the field is rejected before calling the engine.

The draft, the prompt and the result are not stored. The technical logs exclude content. The text is only
persisted later, through the normal saving of a description or the normal publishing of a comment. For a
self-hosted server or external provider, the administrative acceptance given when configuring AI is
reused; the modal briefly warns which engine will process the text and does not ask for a confirmation
per request.

### 13.3 Execution and errors

The same adapter configured for the summaries serves the improvement, whether local engine, self-hosted
server or provider. A person keeps a single active generation in the interface: the button is disabled
until a response is received. Closing the modal discards any later response, but does not promise to
cancel an operation that the engine is already processing nor introduces shared state between tabs.

The errors keep the draft inside the modal and allow retrying. «Generate another proposal» uses the text
visible at that moment and the chosen tone, including manual corrections. No quotas, history, usage
metrics or tables are added.

The route that waits for this generation extends its write deadline to **270 seconds**: it covers the 240
seconds configured for the engine and the margin of serialization and network. The exception is applied
with `ResponseController` only to this response; the other endpoints keep the global limit of 60 seconds.
If the engine exhausts its own deadline, `tickets.writing.unavailable` is returned while the HTTP
connection can still deliver the error.

### 13.4 Implemented scope

The owner chose 1A–19A and corrected the boundary with 20A and the deadline with 21A. Writing from an
empty field, sending the full ticket context, generating format, allowing the function to the User or
Administrator, storing requests, configuring tones, cancelling at the provider and changing the AI
configuration or catalogue are left out. No open decisions remain.

### 13.5 Verification

The full backend and frontend suite, and the Playwright walkthrough on PC and mobile, pass on 2026-10-08.
The walkthrough uses the compatible adapter of the disposable environment, checks that the proposal stays
in the modal, that applying it only changes the draft and that the User does not see the button.
Production was not started or modified.

## 14. Implementation: configuration condition for the writing help

The module offers internally to `tickets` a boolean response that indicates whether an effective AI
configuration exists. It re-reads the persisted configuration, including the internal resolution of
credentials, and keeps the environment fallback used by the generation itself. The final condition is
that the effective settings have an engine address. It does not return address, provider, model,
credentials or health status.

**Configured does not mean available.** A configured AI keeps the help visible even if the engine is
temporarily down; if generation is attempted, the modal shows the current error and keeps the draft. No
health check, periodic polling or provider call is added to decide whether to paint a button.

The condition is evaluated when building the ticket detail. Changing or withdrawing the configuration is
reflected the next time that detail is requested. Live updating, a public AI endpoint and adding
configuration information to the session or to the public branding are left out. The owner chose 1A, 2A
and 3A; no open decisions remain.

AI continues to be mandatory: a sealed installation that requires configuration does not access the
tickets and remains directed to Settings. `false` is a guarantee of the contract in the face of absence
of the adapter, invalid effective configuration, old responses or isolated tests; it does not open an
operating mode without AI.

## 15. Implementation: compliance with the global language in AI responses

### 15.1 Reproduced finding

On 2026-10-12, from a clean clone of the public `main`, an installation globally configured in English
created `CS-2026-0001` and generated its two fields with local Qwen2.5 1.5B. «Last action» came out in
English; «Reason» came out in Spanish while API and database declared it `language: "en"`. The schema,
the health, the JSON and the length were valid, so the current control stored a language statement that
it had not checked. The ticket text was in English, but part of the scaffolding delivered to the model
—labels and instructions— was still in Spanish.

### 15.2 Validation contract

The language is single and global per deployment. Each summary, writing improvement and template
translation is validated against the requested global language before accepting its text. The validation
is local to the backend: it does not send content to another service nor consume another generation. It
examines the complete response and only declares a discrepancy when it has enough confidence; names,
products, code, addresses, numbers and isolated technical terms do not turn a response into another
language. A short or ambiguous text without enough evidence continues to be valid.

The implementation incorporates a maintained detector limited to the two languages admitted by the
product, Spanish and English. Its licence and version go into the third-party notices. The detector does
not choose the language nor take it from an account, session or ticket: it compares exclusively with the
current global language that already accompanies each task.

### 15.3 Prompts and content

The instructions, the context labels and the event names that Catalina Support adds to the task are built
in the requested global language. The subject, the description, the comments, the draft, the placeholders
and any other content written by people are kept literally. No prior translation is introduced, no
additional calls are made to prepare the input and no names or technical data are altered.

### 15.4 Retries, cost and errors

A clear language discrepancy is treated as an invalid response. With the local engine or a self-hosted
server, a single immediate retry is made, reinforcing the instruction with the detected breach. The two
calls belong to the same work attempt of the queue: the internal retry does not separately exhaust the
persisted attempts nor creates two rows.

With the `provider` mode there is no automatic retry: a second request could have a cost and the owner
chose not to authorize it implicitly. The first incorrect text is discarded and the corresponding stable
error is returned. No «up to two requests» notice is added to `/setup` or `/settings`, because the paid
provider never receives the second automatic call.

If the allowed retry also breaches, the text is not stored nor replaces anything. A summary goes to
`error` with `ai.invalid`; «Improve with AI» keeps the draft and shows its error in the modal; and the
translation of a template keeps the previous draft for manual review or edition. The incorrect summaries
already stored are not swept in a batch: Support or Development can request «Summarize again», and all
future moves use the new contract.

### 15.5 Verification and scope

The unit tests cover Spanish, English, technical terms, ambiguous texts, clear discrepancy, successful
retry, second breach and absence of retry in `provider`. They also cover the three consumers: summaries,
writing and templates. The comprehensive check then repeats the whole block 8 from a clean clone of the
public `main`, with the real 1.5B model; a fake engine or an exact phrase are not enough.

Languages other than Spanish and English, translation of the original content, mass regeneration, quotas,
cost metrics, catalogue changes and production are left out. The owner chose 1A–9A, 10B, 11A and 12A. No
open decisions remain.
### 15.6 Implementation and verification

The backend uses `whatlanggo` 1.0.1 inside `ai` and accepts an ambiguous response when the detector does
not consider it reliable. A reliable discrepancy triggers a second generation only in `local` and
`remote`; `provider` returns `ai.invalid` after the first. The local summaries abandoned the old
bilingual contract and use the same global contract as the other adapters. The focused tests and the full
backend suite, including `go vet`, passed on 2026-10-13.

## 16. Observable cycle of the local model

The administrator accepts a single global download. It keeps the partial file when cancelled and
publishes for each model the bytes, the total and a stable phase: `idle`, `downloading`, `validating`,
`activating`, `testing`, `active` or `error`. Resuming reuses the partial file; deleting it is a different
action. The intention to activate stays on the server, so that closing the browser does not cut download,
validation, startup or test.

After validating the GGUF, the administrator automatically activates the requested model. Health and
artificial generation remain different requirements. A failure after the download keeps the file and
allows repeating the activation; it never deletes a valid model as a consequence of a RAM, startup or
response failure. Only an inactive installed model is deleted, and there is only one active.

The manual test receives from `settings` the language, instant and IANA zone already resolved. It
generates an artificial greeting that must include the supplied date and time. It returns text, instant
and zone, does not persist them and does not receive ticket context. A manual failure does not change the
active configuration.

The tests simulate progress, mutual exclusion, cancellation, resumption, deletion of the partial file,
continuation without a client, all the phases, failed and retried activation, and a greeting with an
injected clock. They do not download models nor call real providers. They do not change catalogue, ticket
prompts, queue, privacy, costs or external modes.

## 17. Deactivation and identity of the greeting

The local administrator admits deactivating the active model without deleting its GGUF. It stops the
process, clears the active selection and publishes all the records as inactive; afterwards any of the
installed models can be activated. The operation is idempotent for the given model and does not interfere
with a download in progress. There is no deactivation for the remote or provider modes through this local
catalogue contract.

The manual test also returns the identifier and the catalogue name of the model that answered. Those
data come from the active state that the administrator checked, not from text produced by the model. The
greeting continues without being persisted and without receiving data of tickets or people.

The tests cover stopping the active one, being left without a process, reactivating it, rejecting an
identifier that is not the active one and attributing the greeting to the checked model. They do not
change catalogue, installed files, download, validation or automatic activation after a download.

The administrator tests and the whole backend pass. The real walkthrough stopped the 1.5B, published the
two installed records as inactive, left the wizard blocked and started the same model healthy again; the
manual test correctly attributed its response.

## 18. Implementation: professional writing and a different paraphrase

### 18.1 Corrected contract

The previous contract received five tones. `friendly`, `brief`, `empathetic` and `technical` were
withdrawn, and `tone` is also withdrawn from the public request. Every improvement uses `professional`,
decided by the server; a client cannot replace it.

The first request delivers `editor` and `text`. A regeneration may add `previousSuggestion`, which
contains only the last proposal shown for that same draft. The backend uses the current original draft as
the source of truth and orders producing a professional paraphrase that keeps names, numbers, language and
details, but whose writing is clearly different from `previousSuggestion`. It does not chain the proposal
as a new source nor send proposals earlier than the last one, to avoid drift and limit the context.

The contextual name and the ticket type continue to be resolved on the server. No subject, conversation,
attachments, history or persistence are added. The result continues to be plain text.

### 18.2 Errors and tests

A failure returns the same stable key and does not change the draft. The interface decides to report it
through its common toast and can request the operation again from the initial state.

The tests cover absence of `tone`, professional tone in the prompt, first proposal without reference,
regeneration with the draft as source and only the last proposal as text to avoid, different response,
permissions, limits and errors. They do not change engine, mode, authorization, deadline, global language
or external data policy. The owner chose 7A, 14A and 15A; no open decisions remain.

The implementation removed `tone` from the public DTO and from the tickets service, fixed the
professional tone in the engine prompt and added `previousSuggestion` only for regenerating. The affected
tests of `ai/services` and `tickets/controllers` passed; the full backend suite keeps an unrelated and
reproducible failure in `TestDesarrolloCambiaLaPrioridadDesdeElInterno`, which returns `tickets.notFound`
and belongs to the priority flow, not to assisted writing.

## 19. Implementation: meaningful responses and verifiable paraphrases

### 19.1 Findings reproduced before intervening

The real walkthrough with `CS-2026-0026` found a mismatch between documentation and code: the engine
returned `...` for **Last action**, the backend stored it as `listo` and the interface presented it as a
valid summary. Creating, editing and retrying manually reproduced the same result. The current validation
only demands non-empty text and the language detector accepts as ambiguous what it cannot identify; that
is why a response made up only of punctuation goes through both doors. The literal example
`{"text":"..."}` of the prompt also makes it easier for a small model to copy the marker.

In the comment improvement, the first output only translated and lowercased the draft; the second changed
a single word. Both were accepted as professional and different. It was also checked that the retry of
summaries rebuilds the card and loses the local draft and proposal. These three behaviours contradict the
current sections 15 and 18 and are reported before touching code.

### 19.2 Valid summaries and preservation of the previous one

A summary is only valid if, after normalizing spaces, it contains significant letters or numbers. Empty,
only punctuation and absence markers such as `...`, `—` and `N/A` are rejected. The examples of the
prompt use a real phrase and never a marker that can be copied as a response.

The first invalid response is retried once with a reinforced instruction. If the second also fails, the
row stays in error with a key specific to the field. When a valid text already existed, it is kept and
marked as outdated instead of being replaced or deleted. In a new ticket, **Last action** must say that
the person opened the ticket and described the problem. Proper names are copied literally in any
language: `Usuario Uno` cannot become `User Uno`.

The local engine reproduced the invalid marker even with the reinforced retry during the verification of
`CS-2026-0027`. For that single recognizable case —a conversation whose only event is the opening— the
second failure uses a deterministic phrase in the global language that says that the ticket was opened
and the problem was described. It does not apply to tickets with comments or moves; in them a second
invalid response keeps the previous text and leaves the field in error.

### 19.3 Quality of the assisted writing

The output of **Improve with AI** must keep language, names, numbers and facts, begin with correct
writing and transform the draft appreciably even if it was already well written. The regeneration is also
compared with the previous proposal by means of a deterministic and proportional criterion: relevant
words changed and similarity limit, with specific tolerance for very short texts. Changing punctuation,
capital letters or a single incidental word is not enough.

A low-quality output is retried once. If it breaches again, it is delivered as a selectable and usable
proposal —decision 8B—, but accompanied by a stable warning; it is never presented silently as a verified
improvement. The draft continues to be the source of truth and the previous proposal continues to be only
the text to avoid.

### 19.4 Tests and scope

The tests cover markers, punctuation, retry, error per field, preservation of the previous text, proper
names, ticket opening, proportional similarity, short texts, second low-quality output and the local,
server and provider modes. They do not change permissions, catalogue, models, historical persistence,
global language or external data policy.

The real walkthrough of 2026-10-15 checked creation, edition, individual retry, rejection of `...`,
opening fallback, two proposals and similarity warning. The tests of `ai/services`, `ai/repositories`,
tickets controllers and the 231 frontend tests passed.

## 20. Implementation: stable reason and verifiable last action

### 20.1 Reproduced finding

In `INT-CS-2026-0016`, the card showed as **Reason** and **Last action** the same text: «The expiry notice
does not exist: it has to be added to the access module». The database confirmed that the only AI row,
`ultima_accion`, stayed `listo` with that text, even though the history of the internal ends with
`resuelto` and then `cerrado`. The value of **Reason** was not another AI summary: in an internal it is the
immutable escalation reason, as it should be.

There are three coordinated causes. The context of the internal publishes by mistake the status of the
principal; the escalation reason appears as description, reason and creation detail, while the final moves
lack detail; and the validation only demands significant text and language, without checking that the
output represents the last event or that it is different from the reason. The result contradicts decision
4: the last action must be recalculated with every move and tell the last thing that happened.

### 20.2 Complete and structured context

Each task receives an envelope with the **target ticket** and its linked pair. It includes all the
available functional content: numbers, subject, description, escalation reason, statuses, priority,
category, tags, requester, creator, assignees, observers, comments, history, dates and attachment names of
principal and internal. It does not include binaries, content extracted from attachments, credentials,
internal identifiers or technical metadata.

The envelope expressly separates base data, chronological conversation and **objective fact**. For
**Last action**, the objective fact is the last event of the displayed ticket, not the last event of its
pair. It carries event type, previous and new status, instant, exact visible name of the actor —or «the
system»— and, when it exists, the comment linked to that move. The pair provides context, but never
replaces that fact.

For **Reason**, the complete context serves to understand names and references, but the instruction is
limited to the original problem and need. It does not incorporate assignments, closures or later results.
In the internal, the reason continues to be the stored escalation reason and no other AI row is generated.

### 20.3 Budget and privacy

If the text exceeds the admitted window, the base data of the target, the last event and its linked
comment are never trimmed. Then the base data of the pair and the most recent chronology of both are kept;
the oldest elements are withdrawn first. The trimming keeps labels and order, and does not mix a partial
sentence with another.

The local engine keeps everything on the server. On a self-hosted server or provider, the current
administrative acceptance is reused, whose explanation must state that the functional content of the two
linked tickets may leave the installation. A second acceptance is not added nor is the configured mode
changed.

### 20.4 Hybrid generation and validation

The AI writes a natural sentence from the objective fact and its comment, for example: «Soporte Uno closed
the ticket because he solved the incident». The backend checks deterministically that the output keeps the
name of the actor and represents the expected class of action and final status. It also rejects a response
identical or materially equivalent to the reason when the reason is not the objective fact.

An invalid output is retried once with the mandatory fact reinforced. If it fails again, the field is made
correct through a deterministic fallback: visible actor + localized action + a safe and bounded version of
the linked comment. Without a comment, it ends after the action. The fallback is applied equally if the
engine is down; an explicit action does not depend on AI being available in order to be told. When the
last thing is an independent comment, the AI summarizes it and keeps the current error behaviour if it
does not achieve a valid output.

The language validation continues to be mandatory: a row marked `en` cannot store a clearly Spanish
sentence. Proper names are kept literally and are not translated.

### 20.5 Persistent regeneration and tests

All the existing `ultima_accion` rows with a previous context version are put into a bounded and
idempotent background queue. The version is persisted next to the result; restarting continues what is
pending and does not regenerate what is already updated. Startup and the normal use of tickets do not wait
for the engine. The existing reasons are not regenerated in this sweep.

The tests cover principal and internal, correct status of the target, context of the pair, closure with
and without comment, comment as the last fact, human actor and system, equivalence with the reason, retry,
fallback, language, context limit, provider with acceptance and resumable regeneration of all the last
actions. The browser verifies the reproduced case and another where the closing comment explains the
result. They do not change permissions, catalogue, active model, interface, production or the tag `v1.0.0`.

The migration was applied only to the development environment. The new tests of `ai/services`,
`ai/repositories` and ticket context pass, as do `go vet`, the 232 frontend tests and their build. The
tickets suite keeps only the previous and unrelated failure of priority from an internal, already
recorded. The background sweep regenerated the real case as: «Soporte Uno closed the ticket after
resolving the issue of adding the expiration warning to the access module», different from the reason and
consistent with the English global language and the visible closure.

## 21. Implementation: stable Summary and generation without manual actions

### 21.1 Two fields without overlap

**Summary/Resumen** briefly explains the original problem. In the principal it uses only title and
description; it is generated when the ticket is created and is generated again only when one of those two
texts changes. Comments, category, tags, priority, assignees and statuses do not alter it. In the internal,
the AI summarizes only `escalationReason`; it is generated when escalating and would be regenerated if in
the future editing that text is allowed. That way it remains stable and does not compete with **Last
action/Última acción**, which continues to be updated with every move of the target ticket.

The visible name changes throughout the interface to **Summary** in English and **Resumen** in Spanish. In
the affected technical scope English names are used: the persisted kind goes from `motivo` to `summary`,
the models and DTOs use `summary`, and the escalation reason goes from `reason` to `escalationReason`.
This decision expressly corrects the initial option 11A: `resumen` is not introduced as a technical
identifier. It is not a general translation of all the historical identifiers of the repository; it only
avoids adding or keeping Spanish names in this contract.

### 21.2 Automatic and finite recovery

The interface does not offer actions to regenerate Summary or Last action. The buttons, the manual retry
endpoint and their associated capabilities are removed. Each job keeps the current bounded automatic
attempts. When they are exhausted there is no infinite polling or periodic consumption: the field is
requested again on the next change that normally triggers that field.

While it is being generated or if it fails, the card shows a discreet state of its own for the field. The
lists and cards do not show additional states. If there is no current result, they show `—`; a failed
regeneration does not expose the previous text as if it were current either. The ticket and its other
actions never wait for the engine.

### 21.3 Migration and regeneration

The open migration changes the rows `ai_insights.kind = 'motivo'` to `summary` without keeping an alias.
Then all the existing Summary rows are regenerated in the background, including the internals, in any
configured mode. The owner expressly accepts that an external provider may generate cost during that
sweep. The queue is bounded, idempotent and resumable; production, domain and tag continue outside the
scope.

If a row does not manage to be regenerated, its presentation goes to `—` and keeps the error state until
the next relevant change; it is not restored by a button or by a periodic loop. The tests cover stable
sources, exclusion of moves, principal and internal, migration without compatibility, provider, restart,
attempt limits and absence of manual contracts.

## 22. Implementation: ordered results and finite automatic recovery

### 22.1 Reproduced findings

In the real walkthrough of 2026-10-16, `CS-2026-0028` reached resolved and then closed, but Last action
ended up showing a previous move caused by its internal. The queue replaces a job while it waits, but it
removes it from the map before calling the engine. If during that call a new request enters for the same
ticket and field, both can finish and the old one can save at the end, overwriting the correct revision.

In `CS-2026-0029`, Summary and Last action ended in error and the interface showed **Could not
summarize**. The behaviour matches the current contract —bounded attempts and no manual button—, but it
leaves a failed result indefinitely if no other action occurs that requests that field again. For a stable
Summary, not even an ordinary move recovers it.

### 22.2 Current revision per ticket and field

Each generation request receives a monotonic revision per `ticket + kind`. Marking the field pending and
assigning the revision form a single persistent write. The job carries that revision and all its final
writes —result, error and attributable counter— are conditioned on it still being the current one. If a
later request arrives while the engine serves the previous one, the previous one may finish, but its
output is discarded and never changes the visible state.

The revision is internal and uses English technical names; it does not appear in the DTO nor adds elements
to the interface. On restart, the source rebuilds only the current revision. A single wait per ticket and
field is kept and different tickets and Summary and Last action are not paralysed against each other.

### 22.3 Deferred and finite recovery

Exhausting the immediate attempts schedules a single deferred recovery for the current revision. The
pause avoids insisting on a momentarily saturated engine; when it expires, the context is rebuilt from the
current ticket and a new bounded cycle begins. If a newer functional request exists, it invalidates the
pending recovery. If the deferred cycle also fails, the field stays in terminal error until a relevant
creation or update produces a new revision.

The recovery is persisted to survive the restart, is resumed without duplicates and does not constitute
periodic polling. The endpoint and the manual buttons do not return. During the wait the card shows the
generation state already existing; if the second cycle is exhausted it shows the discreet error and the
lists continue using `—`. The ticket never waits for the engine.

The policy applies to local engine, self-hosted server and external provider. At most it adds a second
automatic cycle per revision; the limit is explicit to contain time and cost. The deterministic fallbacks
of Last action continue to resolve the events that can already be expressed safely without another call.

### 22.4 Single recovery of existing errors

When deploying this correction, the rows that are already in `error` or `sin_motor` and have exhausted
their previous cycle receive a single new revision, with rebuilt context. The sweep is persistent,
idempotent and resumable: it records its recovery version, does not take a row already migrated again and
does not block startup. It includes Summary and Last action of principal and internal and respects the
configured AI mode.

If that revision also exhausts the deferred recovery it stays in terminal error under the new rule.
Production, domain and tag `v1.0.0` are not touched; the sweep is verified only in the authorized
development environment and in a disposable installation.

### 22.5 Tests and out of scope

The tests force two concurrent responses of the same field and show that the old one cannot store a result
or error after the new one. They cover persisted revision, reconstruction after restart, deferred
recovery, invalidation by a later request, terminal limit, provider, idempotent sweep of existing errors
and absence of manual buttons or endpoints. The browser repeats the closure and error cases observed in
the three audit tickets.

They do not change the content or the sources of Summary and Last action, their permissions, language,
privacy, active model, catalogue, comment improvement or ticket transitions.

## 23. Configurable acceleration of the local engine

### 23.1 Finding

The current local engine is CPU only: the administrator starts `llama-server` with two threads, no GPU
layers, and `ai.yml` exposes no devices or NVIDIA runtime. The normal image remains the portable path for
Linux, macOS and Windows. Self-hosted server and external provider manage their hardware outside
Catalina-Support.

### 23.2 Modes and persistence

The local configuration incorporates the English technical value `accelerationMode` with `auto`, `cpu` or
`nvidia`, persisted in `ai_settings`. New installations and existing installations migrate to `auto`. Auto
prefers NVIDIA when the container has it available and falls back to CPU once if the GPU startup fails; it
keeps `auto` to try again on the next startup. CPU never tries GPU. Explicit NVIDIA fails without saving
if the infrastructure is not available and does not use a fallback.

The GGUF model is the same in all modes: changing acceleration does not download, duplicate or validate
the whole file again. NVIDIA downloads to GPU all the possible layers. With several GPUs it uses the first
one Docker exposes; choosing the device or spreading across several remains in the server configuration,
not in the interface.

### 23.3 Transactional change

With an active model, changing acceleration marks the operation as `restarting`, blocks another activation
or change, restarts the process and runs the automatic test. It only confirms when the new process is
healthy. If it fails, it restores the previous mode and process and confirms the error only after the
previous one is healthy again. The AI tasks remain pending and are recovered through the existing
mechanism.

Auto records in technical logs why it used CPU and returns a brief warning without traces. The state of
the catalogue and the tests report configured mode, effective mode and device as manufacturer plus
commercial name; they do not expose memory, UUID, PCI bus or other identifiers.

### 23.4 Distribution and scope

The current CPU image remains. A separate NVIDIA image is used through `ai-nvidia.yml`, an optional
Compose file that requires **Linux with an NVIDIA GPU and NVIDIA Container Toolkit**, or **Windows with
Docker Desktop WSL2 and a visible NVIDIA GPU**. The physical selection is limited with what Docker
exposes, including `NVIDIA_VISIBLE_DEVICES`. The API detects if that variant is not enabled and explains
the required command without saving NVIDIA.

AMD/ROCm, Apple Metal, **native Windows containers**, device selection from the UI and multi-GPU
distribution are left out. CPU continues to be supported through Docker on Linux, macOS and Windows.

### 23.5 Tests

The unit tests cover explicit rejection of an absent NVIDIA and fallback of Auto to CPU. The settings
suite, the 241 frontend tests, their build, the repeatable migration and the CPU build passed. Compose
resolved the NVIDIA overlay and confirmed the pinned CUDA image.

The real run was completed on 2026-11-05 in Docker Desktop WSL2: the GPU was exposed with `--gpus all`,
the administrator detected `NVIDIA RTX A2000 12GB`, activated `qwen2.5-1.5b-instruct`, `/llama-server`
used ~1.3 GiB and generated the two summaries of a real ticket. The full E2E suite passed with 200 cases
passed and 30 skipped. Production, the domain and the version tag were not touched.

## 24. Summaries by audience

The principal ticket keeps two independent versions of **Summary** and **Last action**:

- the `public` version uses only title, description, public comments and events visible to the User;
- the `internal` version adds private comments and their events for Support, Development and a pure
  Administrator.

The User receives exclusively the public version; no private content is sent to the model when generating
it. The teams receive the internal one. Both remain in the global language and keep the asynchronous
cycle, revisions, retries and protection against old results. Creating, editing, deleting or changing the
visibility of a note immediately requests the affected versions in the background; a change of visibility
regenerates both to remove any trace of their previous audience.

The contexts are built in the backend after applying visibility. The public version keeps the normal
ticket key and the internal one uses an opaque technical key that never leaves through the API. The module
tests, the build and the real walkthrough confirm that the User receives only the public version and does
not see the note used by the internal version.

## 25. Professional writing that is useful and verifiable

### 25.1 Finding

The current prompt only asks for professional tone, correctness and clarity. The validation measures
initial capital letter and lexical similarity, but does not demand courtesy, enough development or a next
step. If the second attempt is still weak, the backend delivers it with a warning. The real screenshot
showed a proposal almost identical to the draft, so the behaviour contradicts the appreciable
transformation promised in section 19.3.

### 25.2 Task for the model

The system asks for writing a professional, clear and natural support reply in the same language as the
draft. The output has 2–3 sentences and, as a target, 25–60 words. It must communicate what the technician
did or found, courteously request that the person check the result —including «please» or its natural
equivalent when appropriate— and explain the next step if the problem continues.

The model expands and reorganizes the writing, but keeps all the facts, names, numbers and technical
details. It cannot state that an action was completed, that a problem was solved or that a task will be
done if the draft does not support it. It does not add a greeting or farewell by default either; it uses
the addressee only when it fits naturally in the body. It returns exclusively plain text, without
explanation of its work.

The second attempt receives a reinforced instruction that points out each detectable breach: excessive
similarity, insufficient length, lack of professional structure or absence of a courteous request when the
message asks for verification. The original draft continues to be the only source of truth and the
previous proposal only serves as writing that must be avoided.

### 25.3 Validation and failure

The validation keeps language, meaningful content, initial capital letter and lexical difference, and adds
2–3 sentence structure and a sufficiently developed transformation. The range of 25–60 words guides the
prompt and the quality; a small deviation can be admitted when it avoids inventing content or producing
filler. Courtesy is evaluated when the response asks the user to try, confirm or provide information.

A first insufficient output is retried once. If the second continues to breach, the endpoint returns the
stable invalid-proposal error and does not show a usable proposal or a permissive warning. The draft
remains intact and the interface returns to the state from which another improvement can be requested.

### 25.4 Tests and scope

The tests use the same case in English and Spanish: a technician writes informally that he restarted the
printer and that it seems to work. The proposal must explain the action with professional writing, ask
please that it be verified and point out that it be reported if the failure reappears, without stating a
definitive resolution. Too similar text, short text, contextual courtesy, second failure, protected facts,
regeneration and the local, server and provider modes are also covered.

They do not change permissions, persistence, providers, configuration, privacy, ticket summaries or the
local review of the proposal.

### 25.5 Record of decisions

The owner decided on 2026-10-22 to apply validation and rejection after one retry (4A), besides expressly
improving the prompt. He chose 2–3 sentences of 25–60 words (3A), courtesy and next step without inventing
facts, and not forcing a greeting or farewell (4A of the final review). The same situation will be verified
in English and Spanish. No open decisions remain.

### 25.6 Implementation and verification

The bilingual prompt demands structure, length, contextual courtesy, next step and an express prohibition
against inventing facts. It includes a brief example in each language that shows the local models how to
substantially transform an informal draft without copying markers or inventing actions. The validator
checks 2–3 sentences, 20–70 words as a technical tolerance around the target of 25–60, initial capital
letter, lexical difference and courtesy when the draft asks to check or report. A false match was
corrected so that «we checked the server» is not interpreted as a request to the User. If the second
attempt fails, the service returns `ai.invalid`; the DTO and the interface no longer expose
`qualityWarning` nor allow applying a deficient proposal.

The focused suites of `ai/services`, tickets controllers and services passed. The direct tests validated
the same professional scenario in Spanish and English, besides the rejection after two poor responses.
The 252 frontend tests and the build also passed. A real call with `qwen2.5-1.5b-instruct` transformed the
informal English draft about restarting a printer into two professional sentences, with «Please», a
request for verification and a next step; the proposal passed the validator and was shown in the interface
for review.

## 26. Benchmark of models and threads

> **Decision of 2026-10-31.** The benchmark lives in `tests/ai-benchmark/` and does not touch the active
> engine or production. The conclusion is that the current model is still the best balance and that the
> speed improvement is in the threads, not in changing the model.

### 26.1 Method

Eight cases (summary and last action, Spanish and English, principal and internal) with the exact text that
`tickets` assembles. Each request uses:

- the same system message from `oficioDe`;
- the same user message from `pedirGlobalUnaVez`, including the demand of `{"text":"…"}`;
- `temperature: 0.2`, `max_tokens: 320` and `response_format: {"type":"json_object"}`;
- `llama-server` with `--ctx-size 4096 --parallel 1 --n-predict 512 --no-webui` and the threads that are
  indicated.

Each response is contrasted with the delivered text: JSON, language, cap of 40 words, expected facts,
proper names and numbers that do not appear in the input, and a human review of the sentence. The detail
per response is in `tests/ai-benchmark/report.md`.

### 26.2 Results

| Model and threads | Correct | Average latency | Engine memory |
| --- | ---: | ---: | ---: |
| Qwen2.5-1.5B · 2 threads (current) | 7/8 | 4.0 s | 1.191 GiB |
| Qwen2.5-1.5B · 4 threads | 6/8 | 2.4 s | 1.191 GiB |
| Qwen2.5-1.5B · 8 threads | 7/8 | 2.0 s | 1.191 GiB |
| Qwen2.5-1.5B · 16 threads | 7/8 | 1.9 s | 1.191 GiB |
| Qwen2.5-1.5B · 32 threads | 7/8 | 2.5 s | 1.191 GiB |
| Gemma 3 1B · 2 threads | 5/8 | 4.4 s | 880 MiB |
| Qwen3-0.6B · reasoning by default | 4/8 | 11.3 s | 717 MiB |
| Qwen3-0.6B · reasoning disabled | 0/8 | 1.4 s | 717 MiB |
| Qwen3-1.7B · reasoning by default | 3/8 | 26.3 s | 1.576 GiB |
| Qwen3-1.7B · reasoning disabled | 5/8 | 4.1 s | 1.576 GiB |
| Qwen2.5-0.5B · 2 threads | 3/8 | 1.6 s | 493 MiB |

### 26.3 Reading

- **The current model is not changed for a faster one**: none wins at once in speed and content. Qwen3
  with reasoning is much slower and truncates; without reasoning it improves latency but fails language,
  content or the last action. Gemma 3 1B is not faster and makes content errors. Qwen2.5-0.5B is very fast
  but repeats the reason instead of the last action.
- **The real improvement is in the threads**: the machine has 32 cores and `ai-manager` started
  `llama-server` with `--threads 2`. **That limit was removed**: `llama-server` uses its automatic value
  (`-1`) and on this machine it chooses **16 threads**, with 1.9 s on average against 4.0 s. 8 threads
  gave 2.0 s and 32 threads 2.5 s; the automatic one gets the physical cores right.
- **The last action is the fragile case**: even the current model sometimes repeats the reason instead of
  the final fact. It is not a problem of the model only; the benchmark has to be repeated and the task
  reviewed before changing anything.

### 26.4 Implementation

`ai-manager/main.go` no longer passes `--threads`: it lets `llama-server` choose. The test
`TestCommandLetsLlamaChooseTheHostThreads` checks that the arguments do not contain `--threads` and that
`--ctx-size 4096`, `--parallel 1` and `--n-predict 512` are still there. It was also checked that the
server, without the argument, records `n_threads = 16` on this machine.

### 26.5 Recommendation

1. Keep **Qwen2.5-1.5B** as the active model.
2. Verify in development that the automatic startup keeps health and latency; if the host changes cores,
   there is no need to touch code.
3. Repeat the benchmark at least three times per case and add anonymized real tickets before taking any
   other change to production.
4. Do not add Qwen3 to the catalogue yet: it needs reasoning disabled, a higher memory ceiling and task
   adjustments.
