# Catalina-Support Documentation

[Español](../es/README.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-05
>
> **QA report added on 2026-11-05.** The comprehensive report for version 1.0.1 is in
> `informe_qa_20261105.md`, with the eleven findings closed and the full verification on WSL2.
> **Coordinated change approved, implemented and verified on 2026-10-24:** `prueba-local.md` cleans
> up the Playwright suite so that it describes exclusively the current product, creates all its data
> without seeders and uses a deterministic fake AI engine for summaries and writing. The owner chose
> 1A–4A. The final pass ran 228 cases: 198 passed and 30 were skipped only because of the explicit
> AD, Keycloak and diagnostics profiles.
>
> **Coordinated change approved, implemented and verified on 2026-10-23:** `primer-arranque.md`, `modules/tickets.md`,
> `modules/users.md`, `modules/auth.md`, `modules/settings.md`, `interfaz-y-experiencia.md` and
> `prueba-local.md` record
> the findings from the ephemeral walkthrough: a local 503 that blocks the AI step, the general list
> out of sync, the duplicated create control, the outdated SMTP preparer and the verification of
> consecutive password links. The owner chose 1A–4A and closed the review with 1A–5A, extending the
> protection to Settings. The implementation was covered by 255 Angular tests, the production build
> and focused walkthroughs on desktop and mobile.
>
> **Coordinated proposal of 2026-10-22:** `modules/tickets.md`, `modules/ai.md` and
> `interfaz-y-experiencia.md` widen the four inboxes with 16 px margins and fixed ends, and harden
> assisted writing through a structured prompt, validation, one retry and rejection of still
> deficient proposals. The owner closed decisions 1A–4A; it was approved and implemented on
> 2026-10-22. Focused backend, 252 Angular tests, the build and a real browser at 1920, 1440 and
> 390 px passed; the test with a real model was limited because the environment has no AI engine
> configured.
>
> **Coordinated proposal of 2026-10-21:** `arquitectura.md`, `ambientes.md`, `prueba-local.md` and
> `proyecto-abierto.md` separate the English and Spanish examples, fix the public demo in English and
> coordinate six real screenshots with the website. The single Spanish seeder and its current
> references were recorded as a finding. The owner closed the review with 1A–5A, approved the
> proposal on 2026-10-22 and the implementation was verified that same day.
>
> **Coordinated change approved, implemented and verified on 2026-10-20:** `ambientes.md`,
> `arquitectura.md` and `prueba-local.md` replace the manual production migration with an automatic,
> versioned, transactional Go runner, with history, checksum, lock and prior backup. The owner
> closed and approved decisions 1A–25A. Production was not deployed.
>
> **Coordinated change approved, implemented and verified on 2026-10-20:** `modules/tickets.md`,
> `usuarios-y-permisos.md`, `modules/users.md` and `interfaz-y-experiencia.md` unify the general
> lists under **Tickets**, reserve Users and Classification for the Admin capability and set
> Development's private comments on main tickets. Decisions 1A–25A were closed with
> 6B, 12B, 14B, 20B and 25B. The focused backend suites, the 249 frontend tests,
> the boundaries and the build passed; the global Go suite keeps the known environmental failure of
> reports because the local database does not contain its seeded tickets.
>
> **Coordinated change approved, implemented and verified on 2026-10-20:** `modules/auth.md`,
> `modules/users.md`, `modules/settings.md` and `interfaz-y-experiencia.md` fix binary `objectGUID`,
> accept username or email when signing in through AD and allow searching for and creating an
> active AD account from `/users`. The search is partial from three characters, returns up to
> twenty people and keeps the base filter that excludes disabled accounts. Keycloak keeps automatic
> provisioning until a service account for queries is available. The owner closed and approved
> 1A–15A, with 6B. Real access by username and email was verified, the focused backend suites and
> the frontend tests and build passed; the global Go suite keeps an environmental failure of reports
> because the local database does not contain its seeded tickets.
>
> **Coordinated change approved, implemented and verified on 2026-10-20:** `modules/mail.md`, `modules/settings.md`,
> `primer-arranque.md`, `arquitectura.md` and `prueba-local.md` turn the SMTP test into an optional
> real send to an ephemeral address, with a real `From:`, public rate limit, fixed text and
> confirmation of acceptance without promising delivery. The owner closed and approved 1A–14A; the
> focused tests, the 248 frontend tests, boundaries and build passed.
>
> **Coordinated change approved, implemented and verified on 2026-10-19:** `modules/mail.md`,
> `modules/settings.md`, `primer-arranque.md`, `arquitectura.md` and `prueba-local.md` distinguish
> STARTTLS, direct TLS and unencrypted SMTP. The last one requires administrative acceptance, only
> allows private IPv4 or loopback and does not appear through automatic downgrade; Mailpit is the
> only explicit development exception. The owner closed and approved 1A–10A; the real relay
> authenticated without sending mail and the focused tests, the 247 frontend tests, boundaries and
> build passed.
>
> **Coordinated change approved, implemented and verified on 2026-10-19:** `modules/tickets.md`, `usuarios-y-permisos.md`,
> `interfaz-y-experiencia.md`, `modules/ai.md` and `modules/mail.md` add private comments
> on the main ticket, protected attachments, audit, internal mail and summaries separated by audience.
> The previous rule that sent all internal content to the escalated ticket is identified as a
> finding. The owner closed and approved 1A–26A; backend, 247 Angular tests, the production build
> and the real Support/User walkthrough passed.
>
> **Coordinated change approved, implemented and verified on 2026-10-19:** `modules/tickets.md`, `usuarios-y-permisos.md` and
> `interfaz-y-experiencia.md` redefine **My tickets** as currently assigned or watched tickets,
> separated by level, with `All · Assigned · Watching` in the first row. They exclude opening,
> comments and unassigned tickets; removing the watch removes membership unless assigned. The
> Admin capability keeps the operational role. The owner approved the document; the frontend and
> backend tests passed and the flow was verified with a Development + Admin account.
>
> **Coordinated correction approved, implemented and verified on 2026-10-19:** `interfaz-y-experiencia.md` fixes the double
> selection in the sidebar through a single active entry, exact matches and explicit ownership
> for auxiliary screens. The owner chose 1A–4A. It does not change routes, permissions, labels,
> icons or themes. The build and the 246 frontend tests passed; the browser confirmed a single
> selection in all entries and the ownership of the auxiliary routes.
>
> **Coordinated correction approved, implemented and verified on 2026-10-19:** `modules/tickets.md` fixes the
> contamination of GORM clauses when changing priority from an internal ticket. The fix preserves the
> transaction and the already approved behavior; it does not change permissions, API, interface or data. The
> isolated regression and the entire backend suite passed.
>
> **Coordinated correction approved, implemented and verified on 2026-10-19:** `propósito-y-alcance.md`,
> `interfaz-y-experiencia.md` and `modules/dashboard.md` turn the Dashboard into purely statistical
> reporting, fix transparent bars in the eight themes, group the trend without overlapping
> dates, reduce the filters and replace CSV with a localized six-sheet XLSX. The real review
> also found that the current screen lists tickets against the new scope. The owner closed
> 1A–4A, proposed XLSX in 5, chose 6A–14A, corrected 15A and chose 16A. The specific Dashboard and
> reporting tests in backend passed, 241 frontend tests, boundaries, build, the eight themes and the
> structural validation of the XLSX.
>
> **Coordinated proposal of 2026-10-18:** `modules/ai.md`, `modules/settings.md`,
> `primer-arranque.md`, `arquitectura.md`, `ambientes.md` and `prueba-local.md` add Auto/CPU/NVIDIA
> local acceleration, an optional Compose variant and a transactional switch with safe backup.
> The owner closed decisions 1A–25A; no decisions remain open.
>
> **Dashboard implemented and verified on 2026-10-18:** `modules/dashboard.md` is born and scope,
> permissions, tickets, settings, architecture, interface and local testing are updated. It includes
> operational reporting, filters, trends, times, thresholds and CSV. The owner closed 1A–45A, with 24A
> extended to both duration formats and 26A excluding the User. Real PostgreSQL, service tests,
> 241 frontend tests, boundaries, build and a visual desktop/mobile walkthrough passed.
>
> **Interface change of 2026-10-17, approved, implemented and verified:**
> `interfaz-y-experiencia.md` expands vertically the operational tickets and users tables,
> compacts their headers and keeps 16 px of top, bottom and inner spacing. The 1A–15A review
> was closed. The 238 frontend tests, the boundaries and the build passed; the browser
> confirmed both tables on desktop and that the ticket cards remain intact on mobile.
>
> **Interface change of 2026-10-17, approved, implemented and verified:**
> `interfaz-y-experiencia.md` reduces the sidebar to 208/56 px, places the logo and the name in a
> centered vertical header, separates the single collapse control above the logo when it is
> collapsed and renames the joint screen as Classification/Clasificación. The 1A–14A review was
> closed. The 238 frontend tests, the boundaries and the build passed; the browser
> confirmed desktop expanded and collapsed, the mobile drawer, the centered brand and Classification.
>
> **Public landing page closed on 2026-10-17:** `proyecto-abierto.md`, `ambientes.md` and
> `prueba-local.md` document the bilingual README, the official site, the shared demo, Ko-fi and
> two complementary real screenshots of the dual view and the AI. The relative links, the restored
> local brand and the development and AI Compose files were verified; neither production
> nor any public environment was touched.
>
> **Coordinated correction of 2026-10-16, approved, implemented and verified:**
> `modules/tickets.md` returns to Support the visible Esperando → En progreso transition that the
> backend already supports. `modules/ai.md` prevents an old generation from overwriting the current
> one, adds a deferred and finite automatic recovery and retries the existing errors once.
> It starts from three failures reproduced during the real walkthrough of `CS-2026-0028` to `CS-2026-0030`.
> The owner chose 1A–4A. 238 frontend tests, build, boundaries and the AI tests
> against a real repository passed; the browser ran the return in `CS-2026-0011`. The backend suite
> keeps the previous, unrelated failure of priority from an internal ticket. The manual AI actions
> do not return.
>
> **Change of 2026-10-16, approved, implemented and verified:** `interfaz-y-experiencia.md` distinguishes the
> `en espera` state by role across the UI: `Waiting for you` / `Esperándote` for the User and
> `Waiting` / `Esperando` for Support, Development and Administrator. It does not change the API, database or
> emails. 76 focused tests, the boundaries and the frontend build passed.
>
> **Coordinated proposal of 2026-10-15, pending documentary approval:**
> `interfaz-y-experiencia.md` and `modules/tickets.md` make **Action/Acción** visible above the
> **Open/Abrir** column, centered and with a minimum width. The UI confirmed that the cell was inside the
> table, but its `sr-only` header caused the visual separation. The later review defines
> order, width hierarchy, truncation, limited classification, intermediate and mobile variants and
> tests in both languages. The final review adds accessible keyboard aids, enumerated hidden
> labels, response to the container's real width and preservation of the existing sorting. No
> decisions remain open and no code has been modified yet.
>
> **Coordinated change of 2026-10-15, approved, implemented and verified:** `modules/ai.md`,
> `modules/tickets.md`, `interfaz-y-experiencia.md` and `arquitectura.md` turn Motivo/Reason into a
> stable Summary/Resumen, remove the retry buttons and endpoint, and migrate the affected contract to
> `summary` and `escalationReason`. The owner chose 1A–20B, accepted full regeneration even
> with a provider and left no decisions open. The migration was applied only in development; the
> specific tests, the 231 frontend tests, their boundaries and the build passed. The full
> backend suite keeps the previous failure of priority from an internal ticket, unrelated to this change.
>
> **Coordinated correction of 2026-10-15, approved, implemented and verified:** `modules/ai.md`,
> `modules/tickets.md` and `arquitectura.md` fix the last action that copied the reason, deliver
> complete functional context of main and internal, link event and comment, validate actor and
> action and regenerate the existing rows in the background. The owner chose 1A–14A and approved
> the documents explicitly; no decisions remain open.
>
> **Coordinated correction of 2026-10-15, approved, implemented and verified:**
> `interfaz-y-experiencia.md` sets all dropdowns to 42 px to match the Main/Both/Internal switch.
> `ambientes.md` confirms that `ng serve` already rebuilds and reloads with `no-cache` responses,
> so no redundant flags or polling are added. The owner chose 1A–5A.
> The 232 frontend tests, the build and the real measurement of 42 px passed; no decisions remain
> open.
>
> **Adjustment of 2026-10-15, approved, implemented and verified:**
> `interfaz-y-experiencia.md` records that 52 px were served correctly but turned out
> too large. All dropdowns now use 48 px, keeping typography and padding.
> The 232 frontend tests, the build and the measurement in a new tab passed; no decisions remain
> open.
>
> **Correction of 2026-10-15, approved, implemented and verified:**
> `interfaz-y-experiencia.md` records that 44 px were indeed applied, but were still perceived
> as too thin. All dropdowns now use 52 px without enlarging the typography. The
> 232 frontend tests, the build and the real measurement passed; it was not necessary to restart the
> containers and no decisions remain open.
>
> **Implementation of 2026-10-15, approved and verified:** `interfaz-y-experiencia.md` aligns all
> dropdowns with the 44 px minimum height and the spacing of text fields, with no compact variants
> and at all screen sizes. The 232 frontend tests, the build
> and the real review on desktop and mobile passed; no decisions remain open.
>
> **Coordinated implementation closed on 2026-10-15, with explicit approval:** `modules/ai.md`,
> `modules/tickets.md` and `interfaz-y-experiencia.md` reject meaningless summaries, keep
> the previous text on failures, measure real paraphrases and prevent the AI refresh from destroying
> local drafts or proposals. It starts from the findings reproduced in `CS-2026-0026`; the 1A–15A
> review was closed, with 8B. The owner's later correction removed the permanent help
> from the comment editor and the neutral ready-proposal state. 231 frontend tests,
> the build, the affected backend tests and the real walkthrough `CS-2026-0027` passed; no decisions
> remain open.
>
> **Coordinated implementation of 2026-10-14, approved and verified:** `modules/ai.md`,
> `modules/tickets.md` and `interfaz-y-experiencia.md` remove the modal and the five tones of
> «Mejorar con IA». Generation will be professional and immediate, with an inline proposal per editor,
> application or discard and a different paraphrase on regeneration. The 230 frontend tests,
> the production build, the affected backend tests and the real walkthrough without publishing passed;
> the only failure in the full backend suite is the unrelated priority case recorded in those
> documents. No decisions remain open.
>
> **Coordinated implementation of 2026-10-14, approved and verified:** `modules/tickets.md` and
> `interfaz-y-experiencia.md` remove the internal column that remains empty in **Los dos** and make
> each conversation fill its half completely, keeping 24 px between panels and without touching the
> individual views. The comparison will appear from 1152 px. The owner chose 1A–6A, 7B and
> 8B. The 229 frontend tests, the build and the real walkthrough at 1800, 1152 and 1151
> px passed; no decisions remain open.
>
> **Coordinated implementation of 2026-10-14, approved and verified:** `modules/tickets.md` and
> `interfaz-y-experiencia.md` turn the simultaneous view into two 50/50 panels with their own scroll,
> a visible common header and full usable width; mobile and portrait tablet keep one ticket at a time.
> The owner chose 1A–4A. The 228 frontend tests, the build and the real walkthrough
> on desktop and portrait tablet passed; no decisions remain open.
>
> **Coordinated implementation of 2026-10-14, approved and verified:** `primer-arranque.md`,
> `modules/settings.md` and `interfaz-y-experiencia.md` show the local catalog automatically and
> reserve the **Refrescar catálogo** button to refresh without hiding the cards. The 228
> frontend tests, the build and the real `/setup` walkthrough passed.
>
> **Coordinated implementation of 2026-10-14, approved and verified:** `primer-arranque.md`,
> `modules/settings.md` and `interfaz-y-experiencia.md` replace all native confirmations with
> the themed dialog and add to step 1 the eight themes and the institutional color with preview.
> Backend, 227 frontend tests, build and the real walkthrough on desktop and mobile passed.
>
> **Coordinated implementation of 2026-10-14, approved and verified:** `primer-arranque.md`,
> `modules/ai.md`, `modules/settings.md` and `interfaz-y-experiencia.md` define clickable logos without
> a file name, activation and deactivation control from each card, a greeting attributed by
> toast and a common manual close policy with a maximum duration of ten seconds. The owner
> closed 1A, 2A and corrected 3A to the common ten-second limit; no decisions remain open.
>
> **Coordinated implementation of 2026-10-14, approved and verified:** `primer-arranque.md`, `modules/ai.md`,
> `modules/settings.md` and `interfaz-y-experiencia.md` return to `as-built` after showing and controlling the
> model download, activating and testing the chosen one, managing inactive models and loading the two
> logos from step 1. The finding —the previous interface did not show progress or prevent
> advancing during the download— was fixed after explicit approval. The isolated administrator,
> the entire backend, the 228 frontend tests and the real 3B model walkthrough passed.
>
> **Priorities implemented for 1.0.0, 2026-10-14:** nine documents return to `as-built` after
> adding low, medium, high and critical, with permissions, history, order, filter, single email and
> tests. The owner closed 1A–25A with corrections of its own to 3A, 13A and 18A and approved
> explicitly the coordinated proposal.
>
> **Block 8 closed, 2026-10-13:** `prueba-local.md` and `ambientes.md` record the fourth full
> pass from the public repository. Empty installation, AI, mail, seeders, AD, Keycloak and return
> to local worked; the logs ended clean and the disposable environment was removed.
>
> **Finding corrected and verified from the third public walkthrough, 2026-10-13:** `arquitectura.md`,
> `modules/settings.md`, `modules/auth.md`, `modules/mail.md`, `modules/users.md` and
> `modules/tickets.md` remove the false GORM errors for expected
> absences. `ambientes.md` and `prueba-local.md` keep the approved fourth full repeat. The
> owner approved 1A–4A; backend and `go vet` passed against real PostgreSQL.
>
> **Coordinated proposal of 2026-10-10:** `proyecto-abierto.md`, `arquitectura.md`,
> `modules/settings.md` and `interfaz-y-experiencia.md` prepare the publication under AGPL-3.0-only,
> a CLA without assignment, a brand policy, a community profile, an audit and a link to the source. The
> owner closed 1A–25A. The repository remains private and the license, the history,
> the GitHub configuration and the code are not modified until the documents are approved.
>
> **First finding closed on 2026-10-10:** the audit found two critical vulnerabilities and
> four high ones in the Angular lockfile. The owner approved 26A; Angular was left at 22.2.2, the
> audit ended clean and the tests, the build and the affected walkthrough passed.
>
> **Second finding closed on 2026-10-10:** `go-logs v1.0.0`, a direct backend dependency, did not
> publish a license. The owner approved 27A and authorized the specific change: `go-logs v1.0.1` was
> published under MIT and Catalina Support already resolves and verifies that licensed artifact.
>
> **Third finding closed on 2026-10-10:** `proyecto-abierto.md` added a versioned
> inventory of 29 production dependencies, their full texts, reproducible generation,
> CI check and a copy next to the artifacts, in accordance with 28A–37A.
>
> **Fourth amendment closed on 2026-10-10:** the real CLA test confirmed the full flow and
> revealed the additional public identifiers of the action's fixed schema. The owner
> chose and approved 38A; `proyecto-abierto.md` became `as-built` with the format and the evidence.
>
> **Metadata and links closed on 2026-10-10:** `arquitectura.md`, `modules/settings.md` and
> `interfaz-y-experiencia.md` return to `as-built`. Backend, 227 frontend tests and 12 isolated
> cases on desktop and mobile passed; 39A aligns the collapsed menu with its real behavior. The disposable
> environment was left clean.
>
> **GitHub configuration approved on 2026-10-11:** `proyecto-abierto.md` specifies presentation,
> community channels, squash and protection of `dev` and `main` in accordance with 40A–54A. Visibility is not
> part of this amendment; the application and verification of the settings is still pending.
>
> **Publication correction approved on 2026-10-11:** GitHub is already public and the default
> branch `main`, behind, exposes the previous MIT license. `proyecto-abierto.md` records
> 56A–65A: temporary landing page on `dev`, a second bundle, a limited rewrite of the three branches,
> preservation of the CLA and later synchronization through a PR. The implementation is still pending.
>
> **Sequence corrected and approved on 2026-10-11:** 66A keeps `main` as the main branch
> and `dev` as development. The rewrite will publish `main` first to fix MIT and then the
> other branches and the PR. The implementation is still pending.
>
> **Block 7 implemented and verified on 2026-10-08.** `arquitectura.md` closes documentary
> coherence, automates the frontend boundaries and leaves it decided that there will be no
> logging system in the browser. The frontend, build and backend suites passed; production was not touched.
>
> **Updated on 2026-10-08 (sixth time):** the visibility of «Mejorar con IA» was implemented.
> The detail returns `capabilities.aiWriting`; the interface only mounts the button and the modal when it is
> true. Settings is distinguished from health and AI remains mandatory. Backend, 225
> frontend tests and Playwright on desktop and mobile passed; the three documents return to **as-built**, the
> disposable environment was left clean and production was not touched.
>
> **Updated on 2026-10-08 (fifth time):** `modules/ai.md`, `modules/tickets.md` and
> `interfaz-y-experiencia.md` move to **propuesta** to hide «Mejorar con IA» when there is no
> active configuration. The owner chose 1A–3A: configuration and not health, capability inside the
> detail and update on reload. No decisions were left open and the proposal was approved.
> During implementation it was confirmed that mandatory AI prevents opening tickets in a sealed
> installation without configuration. The owner decided to keep that rule: the false state will be tested in
> backend and component, and Playwright will keep the configured walkthrough. The correction was approved
> and was implemented.
>
> **Updated on 2026-10-08 (fourth time):** the approved writing assistant was implemented.
> Support and Development can improve descriptions and comments with five tones, review the
> proposal and apply it to the draft without publishing. The backend derives the recipient and applies the
> ticket permissions before delegating to AI. Backend, 224 frontend tests and Playwright
> on desktop and mobile passed; the four documents return to **as-built** and production was not touched.
>
> **Updated on 2026-10-08 (third time):** `modules/ai.md`, `modules/tickets.md`,
> `usuarios-y-permisos.md` and `interfaz-y-experiencia.md` move to **propuesta** to add the
> AI writing assistant chosen in 1A–19A. It improves Support and Development drafts with
> minimal context, mandatory review and without persisting prompts or results. No decisions remain
> open; the owner approved the four documents explicitly on 2026-10-08.
> When preparing the code a contradiction with the modular boundary appeared. The owner chose
> 20A: the screen will call `/api/tickets/{number}/writing/improve` and `tickets` will delegate to `ai`.
> The owner approved this correction explicitly; `modules/ai.md` and `modules/tickets.md`
> return to **aprobado** for implementation.
> During implementation it was found that the 60-second HTTP deadline would cut off an authorized engine
> allowed to take 240. The owner chose 21A: only the writing route will have 270 seconds. The
> two documents return to **aprobado** after their explicit approval.
>
> **Updated on 2026-10-08 (second time):** block 6 closed the formatting review of all
> screens in both languages, eight themes and three resolutions. In the shared component it fixed the
> only real finding, an overflow of `/users` in English and mobile. 220 unit tests
> and the full Playwright passed (180 passed and 38 expected skips); zero open findings remained.
> The four documents return to **as-built** and production remained outside the walkthrough.
>
> **Updated on 2026-10-08:** block 5 verified the source and checksum of the three models, fixed
> the 1.5B→3B change and completed the real validation. Activation uses recoverable RSS, measures again,
> restores the origin on failures and rejects 7B without interrupting 3B; the interface shows 2/4/6 GiB and the
> error with required/available memory. `ai-manager`, backend, 220 frontend tests and
> the affected walkthrough on desktop and mobile passed. The cleanup left no `block5` resources. The six documents
> of the block return to **as-built** and production remains postponed until version 1.0.0 closes.
>
> **Updated on 2026-10-07 (seventh time):** block 5 is opened. `modules/ai.md`, `ambientes.md` and
> `prueba-local.md` move to **propuesta** to validate the real GGUF files 1.5B, 3B and 7B in a disposable
> environment. The owner chose 1A–9A and closed the review with 10A–14A: resources, times,
> contract and recovery are measured; `/setup` and `/settings` are tested on desktop and mobile; 7B is not activated for
> lack of RAM; 3B and 7B are removed first, and the final cleanup also destroys 1.5B along with
> all temporary resources. No decisions remain open; the three documents were
> approved explicitly on 2026-10-08 and validation can begin.
>
> **Updated on 2026-10-07 (sixth time):** documentary coherence review, with no behavior
> changes. The current summaries are aligned with the five steps of `/setup`, the global
> language, the five single-row configuration tables and the two AI auxiliary tables, and the local engine
> that is only brought up for that modality. References that still described the
> four already implemented and verified blocks as pending are also closed. All their documents remain
> **as-built** and production remains postponed until version 1.0.0 closes.
>
> **Updated on 2026-10-07 (fifth time):** block 4 detects differences between the approved
> architecture and the current AI clients/administrator. `arquitectura.md`, `ambientes.md` and
> `modules/ai.md` moved to **propuesta** to close redirects with credentials, supervision of
> `llama-server`, effective resources and isolated tests of the four adapters. The owner
> chose 1A, 2A, 3A and 4A; the review closed 5A, 6A, 7A and 8A with no open issues. Only the
> documentary approval is missing, granted explicitly on 2026-10-07. The block returns to **as-built** after
> passing backend, local administrator, 218 frontend tests and 180 executed E2E walkthroughs.
>
> **Updated on 2026-10-07 (fourth time):** when opening block 3 a contradiction was found in
> `primer-arranque.md`: the approved section swapped Mail and AI compared with the rest of the document,
> the code and the tests. The owner corrected the order: Mail stays in step 4 and AI in
> step 5. The final review decided to trust the activation saved at sealing, gather privacy
> and cost into a single confirmation and resume local downloads. The document returns to
> **as-built** after verifying services, 218 unit tests and 180 executed E2E walkthroughs.
>
> **Updated on 2026-10-07 (third time):** `flujos.md` returns to **as-built**. The eight ticket
> notices and the three account emails query the current global language when sent, without separating
> recipients by personal preferences. Production remains postponed until 1.0.0 closes.
>
> **Updated on 2026-10-07 (second time):** `usuarios-y-permisos.md`, `modules/auth.md` and
> `modules/users.md` return to **as-built**. The `users.language` column, the API fields and the
> session state disappear; the account emails query the global language. Production remains
> postponed until version 1.0.0 closes.
>
> **Updated on 2026-10-07:** the coordinated language proposal is corrected. Only the eleven
> mail templates are translated; they can always be edited manually and only an external provider
> requires confirming the economic estimate. Existing summaries remain visible in their
> original language; new or recalculated ones use the current global language. The correction was
> approved explicitly and enables implementation.
>
> **Updated on 2026-10-06 (third time):** the coordinated amendments of
> `modules/ai.md`, `modules/settings.md`, `primer-arranque.md`, `usuarios-y-permisos.md`,
> `modules/auth.md`, `modules/users.md`, `modules/mail.md`, `modules/tickets.md`, `flujos.md`,
> `interfaz-y-experiencia.md`, `arquitectura.md` and `ambientes.md` move to **propuesta**. They define mandatory AI —local,
> own server or provider— managed from
> the interface, and a single language for interface, emails and AI. The review was closed with no open
> decisions and the owner approved the proposal explicitly. Implementation can begin.
>
> **Updated on 2026-10-06 (second time):** `prueba-local.md`, `ambientes.md` and
> `arquitectura.md` return to **as-built**. The correction keeps the read-only test code,
> starts its backend once with `go run .` and removes `backend_tmp`. The canonical run
> without optional profiles ended with 177 passed cases and 39 skipped, exit 0; `down -v` removed
> all resources.
>
> **Updated on 2026-10-06:** `interfaz-y-experiencia.md`, `modules/settings.md` and
> `primer-arranque.md` return to **as-built**. The amendment replaces the general notices located at
> the beginning of `/settings` and `/setup` with a shared toast, visible from the current position. The
> permanent notices remain in their section. The review was closed, the owner approved the
> three documents and the implementation passed 220 unit tests, the build and the affected
> Playwright walkthroughs on desktop and mobile.
>
> **Updated on 2026-10-05 (third time)**: `modules/settings.md` and
> `interfaz-y-experiencia.md` return to `as-built` after explicitly showing in `/settings` the
> selected zone, with the same read-only field already used in `/setup`. It does not change persistence,
> the API or the database. The review was closed, the owner approved both
> proposals explicitly and the tests passed on desktop and mobile.
>
> **Updated on 2026-10-05 (second time)**: `interfaz-y-experiencia.md` returns to `as-built` after
> finding that `/mail` places its back arrow on the right although the approved rule and the view's own
> comment say it goes on the left. The inventory confirms that the other three
> headers with this control already comply. The scope does not include flow links inside the content;
> the review was closed, the owner approved the amendment explicitly and the change passed the tests
> on desktop and mobile.
>
> **Updated on 2026-10-05**: `primer-arranque.md` returns to `as-built`: step 3 of `/setup`
> shows the selected region in a read-only field before the search. The amendment was
> reviewed, approved, implemented and verified on desktop and mobile. `interfaz-y-experiencia.md` records the
> same visual rule.
>
> **Updated on 2026-10-04**: `prueba-local.md` returns to `as-built` after implementing and verifying
> the automatic migration before the backend, `/setup` initially in English and the
> editable Mailpit values in development.
>
> **Updated on 2026-10-03**: `prueba-local.md` is added, implemented and verified on Linux, for the
> local walkthrough of the empty startup, the schema with one command, the examples that keep
> configuration, local Keycloak and the isolated suite. Its amendment §5.1.b was also approved explicitly.
>
> **Updated on 2026-10-02**: review of the README for whoever tests the product: empty installation
> first and optional examples, by correction of the owner; `admin`/`admin` locally and
> `ADMIN_PASSWORD` in production. The outdated pending item of the wizard in
> `arquitectura.md` (7 amendments) is corrected and the real limits of the current deployment are reported in
> `ambientes.md` (19 amendments). No code is modified and production is not opened.
>
> **Updated on 2026-10-01 (seventh time)**: **the example data is seeded with a one-off container**,
> `docker compose -f dev.yml run --rm seed-en|seed-es`, **the same on Linux, macOS and Windows** —it leaves
> out the problem of the bash script, which does not run on Windows, and with it the copy of the
> attachments—. The service lives in `dev.yml` **with its own profile**, so **`up -d` does not start it**, and
> **the examples are optional** (the schema remains mandatory). The container runs **the usual
> script, `scripts/dev-seed-en.sh` or `scripts/dev-seed-es.sh`**, which also still works on Linux and macOS outside the
> container. `ambientes.md` is amended (**18 amendments**: sections 3.1, 3.3, 5 and 9.3) and
> `arquitectura.md` (**6 amendments**: section 10.3), and the `README.md` removes the Windows notice and the one about
> the attachments. Checked on 2026-10-01: **two consecutive passes** of the command («Listo: 11 cuentas, 25
> tickets y 5 adjuntos» both times), 11 accounts and 25 tickets in the database, **5 attachments in `_files/`**
> and `up -d` without starting the service.
>
> **Updated on 2026-10-01 (sixth time)**: **the Angular development server forwards `/api` to the
> backend**, so **you enter `http://127.0.0.1:11001` with `admin`/`admin` without nginx** —the
> walkthrough of whoever downloads the project and only has Docker—. `frontend/proxy.conf.json` is added
> and declared in `serve.options.proxyConfig` of `frontend/angular.json`, with destination
> **`http://backend:11002`** (the service inside the environment network) and **without rewriting the path**.
> It is **development only**: the `build` does not see it and **in production nginx keeps forwarding**. The
> `README.md` removes the «pendiente de decisión» notice. Amended: `arquitectura.md` (**5
> amendments**: section 9, 10, 11 and the row of 13, and nginx is no longer necessary to
> enter) and `ambientes.md` (**17 amendments**: section 3.1 and 9.3). Checked on 2026-10-01 through
> 11001: `GET /api/auth/methods` and `GET /api/health` → **200 JSON**, `POST /api/auth/login` with
> `admin`/`admin` → **200** with token, and Playwright against `http://127.0.0.1:11001` → **14 cases
> green** (sign in and out with the factory account).
>
> **Updated on 2026-10-01 (fifth time)**: **the schema and the examples can be applied without `bash`**.
> The local steps of the `README.md` remain as a walkthrough for whoever arrives from outside —clone,
> bring up, apply the schema, sign in with `admin`/`admin`—, with the reason for each step, and **the two
> Docker commands that work the same in PowerShell, CMD and bash**. `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh` is stated
> as **what it is, for Linux and macOS** (and as what it is: **extra content, not a requirement to
> enter**). `ambientes.md` moves to **16 amendments** (its section 3.2 adds the two commands and
> sections 3.1, 3.3 and 5 separate the two ways) and `arquitectura.md` to **4** (its section 10.3, which
> began with `up -d` without the schema, already includes it). Checked on 2026-10-01 on a new database
> (`catalina_support_prueba`, deleted afterwards): the schema leaves 18 tables, `GET /api/auth/methods` and
> `POST /api/auth/login` with `admin`/`admin` respond **200 without example data**, and the same file
> applied twice in a row gives **0 errors**.
>
> **Updated on 2026-10-01 (fourth time)**: the `README.md` makes clear, in the local steps, that
> **on a new database `./scripts/dev-seed-en.sh` or `./scripts/dev-seed-es.sh` is what creates the schema** —not only the examples—, that
> **until the application is run it has no tables** and that the symptoms are
> `relación "installation_settings" does not exist` and `relación "ai_insights" does not exist`; and
> **reorders the steps** (environment, then schema and data, and only then open). `ambientes.md`
> moves to **15 amendments**: its sections 3.1, 3.3 and 5 say what the schema applies in development and
> **when** —before opening the application— and point to the `README.md` instead of repeating it.
>
> **Updated on 2026-10-01 (third time)**: **the test directory and Keycloak leave `dev.yml`**
> and move to **`active-directory.yml`** and **`keycloak.yml`**, each with **its own command**, and **the
> `auth` profile disappears**. The suite of the directory paths moves to **two steps** —bringing up each
> service with its file and running the normal suite—, and both enter the network
> **`catalina-support-dev`**, which **`dev.yml` owns** with a fixed name (without `external`). Amended
> are `arquitectura.md` (**3 amendments**), `ambientes.md` (**14**) and `modules/auth.md` (**7**).
> And the `README.md` already states **the development factory password** —the one in `config/env/dev.env`, which
> is versioned—, making clear that **the production one is `ADMIN_PASSWORD` and is not written**.
>
> **Updated on 2026-10-01 (second time)**: the `README.md` adds **the table of the eleven development
> example accounts** —name, email, role and password—, and `ambientes.md` moves to **13
> amendments** by **no longer repeating that list** and pointing to the table in the `README.md` (section 3.3),
> so that the data lives in a single place.
>
> **Updated on 2026-10-01**: the **three optional elements** —the AI engine, the test
> directory and Keycloak— are left with **one Docker command each** and stated to be optional:
> `ambientes.md` moves to **12 amendments** (section 9.3: the directory and Keycloak share the profile
> `auth` but **can be brought up separately** by naming the service, without activating the profile
> —checked—, and the usual command `--profile auth up -d` still brings up both),
> `modules/ai.md` to **2 amendments** (the engine resources are corrected: **1500m of memory**, **no
> CPU limit** and **~1,44 GiB measured**, where it said «950 MB / 1,09 GB» and «2 CPU»; and the
> **expected warning of the `ai_modelos` volume** is explained) and `arquitectura.md` to **2 amendments** (the row of the measured
> engine, in section 13, with the same figures: **1500 MiB limit**, **~1,44 GiB —98% of the
> limit—** and **no CPU limit**).
>
> **Updated on 2026-09-30 (third time)**: the two module documents affected by the wizard tests
> are up to date: `modules/mail.md` moves to **6 amendments** —it adds **`Probar`**, which
> connects and authenticates without sending any mail— and `modules/settings.md` to **8 amendments** —it declares
> **`ProberDeCorreo`**, which the wizard reuses for step 4—.
>
> **Updated on 2026-09-30 (second time)**: **the first-start wizard tests what it asks for**
> —its two public endpoints `POST /api/setup/entry/test` and `POST /api/setup/mail/test`, with the same
> seal lock, and the buttons of steps 2 and 4—, and the mail test **connects and authenticates
> without sending any mail**. `primer-arranque.md` moves to **1 amendment**
> (`docs/primer-arranque.md`, section 3.1).
>
> **Updated on 2026-09-30**: **the AI engine is configured from the Settings screen**
> —its address and its model, with its «Probar la conexión» button—, with the environment as backup
> (`docs/modules/settings.md`, decision 16, and `docs/modules/ai.md`, decision 21). `modules/ai.md`
> moves to **as-built**.
>
> **Updated on 2026-09-27**: `modules/ai.md` is born, the sixth module, with the two fields that
> the AI engine writes in the ticket list.
>
> **Updated on 2026-09-26 (third time)**: the text and code attachments —a `.sql` could not be
> sent— and **the image and video capped at 480 × 360** that open in the viewer when clicked
> (`docs/modules/tickets.md`, decisions 54 and 55).
>
> **Updated on 2026-09-26 (second time)**: the ticket lists become **My tickets** (what is
> mine), **Main tickets** and **Internal tickets**, and «Nuevo ticket» and
> «Asignármelo» are removed from the menu and the inbox (owner decisions; `docs/modules/tickets.md`, decisions 50 to 53).
>
> **Updated on 2026-09-26**: **how an amendment is counted** is fixed (section "The amendment log and how it is counted")
> and the numbers in the table are set with that rule, because until today
> some added up and others did not; and `modules/mail.md` and `modules/users.md` move to **as-built**, which was
> what corresponded to them since their code exists.

Index of the project documentation. **Rule 0** of `AGENTS.md` is mandatory: no code is
written without an approved document first.

## Documents

| Document | State | Covers |
| --- | --- | --- |
| `README.md` (this file) | **as-built** | Coordination of the comprehensive Playwright cleanup |
| `informe_qa_20261105.md` | **as-built** | Comprehensive QA for 1.0.1: closed findings, WSL2, ISO dates and tests |
| `proyecto-abierto.md` | **propuesta** | English demo and bilingual screenshots of the official site |
| `prueba-local.md` | **as-built** | Current, isolated suite with no historical contracts |
| `arquitectura.md` | **propuesta** | Demo data separated by language |
| `propósito-y-alcance.md` | **as-built** (3 previous amendments) | Purely statistical Dashboard, with no evaluation or SLA |
| `usuarios-y-permisos.md` | **as-built** | Cumulative permissions and privacy of comments on the main ticket |
| `flujos.md` | **as-built** | Single notice to all Support when a ticket first becomes critical |
| `modules/settings.md` | **as-built** | Local catalog unavailable without blocking Settings |
| `modules/mail.md` | **as-built** | Internal notice of a private comment without sensitive content |
| `modules/auth.md` | **as-built** | Verification of consecutive password links |
| `modules/users.md` | **as-built** | A single accessible control to create accounts |
| `interfaz-y-experiencia.md` | **as-built** | Consistent states in settings, inboxes and creation |
| `primer-arranque.md` | **as-built** (8 previous amendments) | Missing local catalog without a false server-down |
| `ambientes.md` | **propuesta** | Explicit `seed-en` and `seed-es` services and commands |
| `modules/ai.md` | **as-built** | Verifiable professional writing and rejection of deficient proposals |
| `modules/tickets.md` | **as-built** | Consistency when switching between Main and Internal |
| `modules/dashboard.md` | **as-built** | Aggregated Dashboard, metrics, charts and localized XLSX |
| `entornos-compartidos.md` | **as-built** | Generic capabilities for shared public installations |

The product chain **is complete**: seven documents that cover what is built, how it
behaves, how it looks and how it is deployed. In addition, **each module has its document**, which is written
right before implementing it, in this order:

1. `modules/mail.md` — **complete**: backend and the editor.
2. `modules/settings.md` — **the Settings screen is complete**: the brand, the color, the
   prefix, the assignment and the installation language.
3. `modules/auth.md` — **complete**: the three sign-in paths, done and verified.
4. `modules/users.md` — **complete**: backend, the three screens and the three directory actions.
5. `modules/tickets.md` — **complete**: backend and screens, with the comment attachments.

## Conventions

### Mandatory header

Every document in `docs/` starts like this:

```markdown
# Título

> **Estado:** propuesta | aprobado | as-built
> **Última actualización:** AAAA-MM-DD
```

- `propuesta`: written, pending approval. No code in that area is touched.
- `aprobado`: approved by the project owner. It enables implementation.
- `as-built`: the code exists and the document describes what it does today.

### The amendment log and how it is counted

Below the header, each document carries **its amendment log**, from the most recent to the
oldest: **one entry per change**, with its date and what changed and why.

**An amendment is a log entry that begins with `Enmendado el …`.** That is what the
number shown in the table above and in `AGENTS.md` says, and it can be checked by counting those
entries in the document header. **`Pasa a as-built el …` does not count**: it changes nothing in the
document, it changes its state.

Two things are written separately, and are not amendments:

- **The project owner's correction**, when it changed what was proposed, is noted **inside the entry**
  («es una corrección del responsable», «yo proponía… y decidió…»).
- **What was discovered during implementation** also goes inside its entry, saying what was found.

The log is a memory of **why** a document is the way it is. Entries are not crossed out or deleted:
if something is corrected, the new entry is added and what has been corrected is stated.

### Names and folders

- **Product documents** live in `docs/`, and are named in Spanish, in lowercase, without accents
  or spaces: `usuarios-y-permisos.md`.
- **Module documents** live in `docs/modules/` and are named **like the module**:
  `modules/auth.md`, `modules/users.md`. The module name prevails there, even if it is in English, because
  it is the name its folder has in the code.
- A module document tells **how** that module is built. What is a product decision —what
  each role can do, how tickets behave— lives in the `docs/` documents, and the module document
  does not repeat it: it states it in writing and points to it.

### Correspondence with the code

The document ↔ code table lives in `AGENTS.md`. When a document moves to `as-built`, its
row must be added or updated there in the same change.

### Writing rules

- Spanish, short sentences, no filler and no marketing flourishes.
- Each document states **what is decided**, **what is pending** and **what is out of scope**.
- What is not decided is marked as pending; it is not filled in with assumptions.
- When a document describes something already implemented, it describes what the code really does,
  not what was intended.
