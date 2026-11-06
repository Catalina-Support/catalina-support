# Tickets: data model, transitions, and modules

[Español](../es/modules/tickets.md) · **English**

> **Status:** as-built
> **Last updated:** 2026-11-03
>
> **QA correction, 2026-11-03.** The user history misrepresented the move to `escalado`: it was
> translated as "En curso", the same as `en progreso`, and looked like a duplicate movement. Now the
> user reads "moved the ticket to the technical team", with no internal jargon.
>
> **Proposal of 2026-10-30, approved, implemented, and verified.** Section 35 allows Support to create a ticket on behalf of a
> person without an account, creating the account in the same submission. The local account is born without access; with AD
> the person is required to exist in the directory. Permissions for the other roles and the
> ticket model do not change.
>
> **Correction approved, implemented, and verified on 2026-10-30.** Section 34 corrects the
> ticket card: manual assignment does not depend on dispatch and, in tickets, the operational role
> prevails even when the account has administrative capability. The correction reaches the
> interface, which hid the reassign button on internal tickets from Support and treated Support and
> Development with Admin capability as read-only. API, data model, and backend permissions do not
> change.
>
> **Coordinated change approved, implemented, and verified on 2026-10-22.** Section 32 corrects the horizontal utilization of
> the four ticket queues. The owner chose 1A–4A in the review: keep the menu,
> reduce the outer margin to 16 px, keep Ticket and Action fixed at the ends, and give inner
> breathing room to the first and last column. The owner approved the scope and the implementation
> matches `interfaz-y-experiencia.md` and `modules/ai.md`.
>
> **Correction approved, implemented, and verified on 2026-10-20.** Section 27.3 corrects the discrepancy between the approved
> matrix and the comment flow: when the requester replies on a main ticket `en espera`,
> comment, return to `en progreso`, and history form a single transaction. The owner closed
> decisions 1A–8A; no open decisions remain. The ticket service tests passed;
> the full suite only retains the known environmental reporting failure when there are no seeded
> tickets.
>
> **Amendment approved, implemented, and verified on 2026-10-20.** Section 31 merges the general lists
> into a single **Tickets** module, with Main as the initial view and Internal as the alternative
> view. The owner closed 1A–25A except 6B, 12B, 14B, 20B, and 25B.
>
> **Amendment approved, implemented, and verified on 2026-10-19.** Section 30 adds public and private comments to the
> main ticket. It expressly replaces the rule that referred all internal conversation to the internal
> ticket: main tickets that are not escalated also need team coordination. The owner
> closed and approved decisions 1A–26A. The full backend and frontend suites passed, along with the
> build and a real check with Support and User.
>
> **Amendment approved, implemented, and verified on 2026-10-19.** Section 29 redefines **My tickets** as a strict
> personal queue: assigned or watched, without including tickets because they were opened or
> commented on and without including unassigned ones. Support keeps only main tickets and Development only
> internal ones; the general lists retain access to the rest. Admin capability does not alter the
> operational role. The owner closed decisions 1B and 3A–9A and approved the document.
>
> **Correction approved, implemented, and verified on 2026-10-19.** The test
> `TestDesarrolloCambiaLaPrioridadDesdeElInterno` found that `ChangePriority` reuses a
> GORM query with a lock: after locating the internal ticket by `number`, the search for the main ticket
> mistakenly keeps that same predicate and ends in `tickets.notFound`. Section 18 already establishes
> the correct behavior. A new locked query is created for each table, a
> single transaction is preserved, and the regression demonstrates both the update of the
> main ticket and the returned internal card. Permissions, HTTP contract, interface, emails, and
> data model do not change. The isolated regression and the entire backend suite passed.
>
> **Implemented and verified on 2026-10-18.** Section 28 defines the data sources and temporal
> semantics that the new Dashboard consumes without changing the ticket flow.
>
> **Correction approved, implemented, and verified on 2026-10-16.** Section 27 fixes the gap
> verified in the browser that prevented Support from returning a main ticket from "Waiting" to
> "In progress". The owner chose 1A. The 238 frontend tests, its boundaries, and the
> build passed; the browser executed the real transition on `CS-2026-0011` and confirmed the history and
> the regeneration of Last action. Permissions, email, and the other transitions do not change.
>
> **Amendment approved, implemented, and verified on 2026-10-16.** Section 26 adapts incremental
> consumption to the new scrollable table container without changing API, data, or permissions. The
> owner chose 1A–13A. It is coordinated with `interfaz-y-experiencia.md` and corrects the previous
> prohibition on vertically fixing the header. Nine focused tests, the boundaries, and the
> build passed; the browser confirmed loading 20→27 and restoration of the internal scroll.
>
> **Amendment approved, implemented, and verified on 2026-10-16.** Section 25 defines the incremental
> loading of the four queues on top of the existing pagination, without changing the API. The
> owner chose 1A–10A and no open decisions remained. The tests cover appending without
> duplicates, error and retry, text, and alignment; the build passed. The browser verified the
> real sequence 20→27, the silent end, and restoration when returning within the same session.
>
> **Amendment approved, implemented, and verified on 2026-10-15.** Section 24 declares
> the Ticket column fixed during horizontal scrolling in the table variants. The
> owner chose 1A–7A and, after detecting that the previous contract suppressed that scrolling,
> corrected the threshold with 1A–3A: scrollable table from 768 px and cards only below that. It is
> coordinated with `interfaz-y-experiencia.md`, does not change module contracts, and leaves no open
> decisions. The specific test and the build passed; the real check confirmed the threshold,
> keyboard scrolling, the fixed position, and the eight themes.
>
> **Amendment approved, implemented, and verified on 2026-10-15.** Section 23 makes
> the opening column explicit in the PC table without changing routes, permissions, or the module's
> response. The owner chose 1B, 2A, 3A, 4A, 5A, 6A, 7A–10A and, in the comprehensive review of this
> table, 1A–12A, 13B+13C, 14A–24A. He expressly corrected the order of the columns. It is coordinated with
> `interfaz-y-experiencia.md` and leaves no open decisions. The specific test and the
> build passed; the real check covered the full table, the intermediate table, and cards.
>
> **Amendment approved, implemented, and verified on 2026-10-15.** Section 22 limits
> Summary to the original approach, renames `reason` to `escalationReason`, and removes the manual
> retry route. The owner chose 1A–20B with the technical-name corrections; it is coordinated with
> `modules/ai.md`, `interfaz-y-experiencia.md`, and `arquitectura.md`. No open decisions remain and
> the implementation matches the approved contract. The migration was applied only in development; the specific AI and
> ticket tests, the 231 frontend tests, and their build passed. The full backend suite retains the previous failure of priority from an internal ticket, unrelated to this change.
>
> **Correction approved, implemented, and verified on 2026-10-15.** Section 21 explicitly
> links the comment of a movement with its history entry, builds the full context of main and
> internal tickets, and guarantees that the target state is the one of the displayed ticket. It is
> coordinated with `modules/ai.md` and `arquitectura.md`; the owner chose 1A–14A, approved the three
> documents, and no open decisions remain.
>
> **Implemented and verified on 2026-10-15, with explicit approval from the owner.** Section 20 preserves previous summaries
> when a regeneration fails and prevents its refresh from destroying local work. The owner closed
> 1A–15A, with 8B, and removed the redundant aids from the editor and the valid proposal; no open
> decisions remain.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section 19 removes the tones and the
> prior modal from "Improve with AI", keeping a local review by editor before replacing the
> draft. The owner closed 1A–15A with 3B, 6B, 7A, 10B, and 11B; no open decisions remain.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section
> 18 removes from the simultaneous view the internal column that remained reserved even when its card was hidden and lowers the comparison width
> to 1152 px. The owner chose 1A–6A, 7B, and 8B; no open decisions remain.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section
> 17 fixes the simultaneous view of the main and internal tickets so that each conversation has its own scroll
> and uses the full width.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** Section 16 incorporates the shared priority,
> its permissions, history, filters, order, and sample data according to 1A–25A. It removes priority
> from the historical exclusion in section 7. The owner explicitly approved it.
>
> **Second noise finding corrected and verified, 2026-10-13.** Although the absence of the internal ticket was already
> corrected, the seeder showed the same false error when checking a normalized category
> before creating it. Section 15 extends the rule to all expected absences in the ticket
> repository. The owner chose and approved 1A–4A; the suite passed against real PostgreSQL.
>
> **Finding corrected and implementation verified on 2026-10-13: language and expected query.** Block 8 found
> a summary in another language and that reading a main ticket that was not escalated prints `record not found` as if
> it were an error. Section 14 coordinates the AI contract with `docs/modules/ai.md` and silences only
> that normal absence. The owner chose 1A–12A, with 10B, and explicitly approved the proposal
> on 2026-10-13.
>
> **Amendment implemented and verified on 2026-10-08: writing capability in the detail.** Section
> 13 makes the detail decide whether help is offered according to role, permission, and a configured AI.
> The owner approved 1A–3A and the correction that keeps AI mandatory.
>
> **Correction approved and implemented on 2026-10-08.** The implementation review found that Playwright cannot
> open a ticket on a sealed installation without AI: the mandatory block acts first. The
> owner kept AI mandatory. The expected evidence is corrected without changing the contract or
> allowing the product without a configured engine. The final evidence follows this correction.
>
> **Amendment implemented and verified on 2026-10-08: AI writing improvement.** Section 12
> describes the editors that can request it and how `tickets` gives `ai` the minimum authorized
> context. Decisions 1A–21A were approved; backend, frontend, and Playwright passed on PC and
> mobile without touching production.
>
> **Correction proposed on 2026-10-08.** The route initially assigned to `/api/ai/**` contradicted
> the rule requiring a ticket screen to use its own API. The owner chose 20A:
> `tickets` publishes the operation, validates, and delegates to `ai`. It required documentary approval before
> implementing; the owner explicitly approved it.
>
> **Correction proposed on 2026-10-08.** The global HTTP timeout of 60 seconds is less than the 240
> seconds allowed by the engine. The owner chose 21A: this route sets 270 seconds and no
> other one changes. The owner explicitly approved the adjustment.
>
> **Amendment proposed on 2026-10-06.** What exists remains as-built. Section 11 adapts the
> summaries and notices to the global language: the API returns one text per field and emails are no longer
> grouped by account language. The review was closed and the proposal was approved.
>
> **Amended on 2026-09-27 (third time)**, at the owner's request, upon using it: **tags
> can now be maintained** —they are created, **renamed** (and the change applies to **all** tickets
> that carry it) and **retired** (and removed from those tickets, asking first)—, because on the
> screen there was no way to create or correct them and **categories took all the
> spotlight**. They still **emerge on their own** when written in a ticket: there is no need to register them to
> use them. And the **two halves of the screen have the same weight** (decision 72).
>
> **Amended on 2026-09-27 (second time)**, at the owner's request, upon using it: the list has **a
> single search button** that opens a modal with all criteria, with a summary of what is filtered
> (section 5 and `docs/interfaz-y-experiencia.md`, section 3.3) —the three filter rows ate up the
> screen and **the category row overflowed**—; and the **ticket view when opening it** becomes the one of the
> role of the person looking, with "Both" showing **only the two conversations** (section 5).
>
> **Amended on 2026-09-27**, when the owner requested **categories and tags** to group and
> search: **`ticket_categories`** is born —"Network", "Software x", "Licenses"…— with which **every ticket
> is born classified** (there is never one without a category, and the database itself guarantees it) and which can be
> corrected later, and **`ticket_tag_names`**
> —optional, **lowercase and hyphenated** (`red-wifi`),
> with suggestions from those that already exist—. **The category and tags belong to the main ticket** and the
> internal ticket inherits them when read. Both serve to **filter** (chips) and to **search** (they enter the
> same text box). The catalog is maintained by **Support and the Administrator**, and **retiring** a
> category is **only the Administrator's** job. Decisions 64 to 68.
>
> **Amended on 2026-09-27**, when the owner requested **being able to tag other technicians and
> developers** so they become participants in the ticket: the concept of **watcher** is born, which **is not
> the same as the assignee** —the ticket has **at most one assignee** (still optional,
> as it was) **and many watchers**—. Tagging is done **by writing in the comment** (`@` or the editor
> button), the name is **highlighted as a mention**, and whoever reads it sees who is being called. Watchers
> **appear in their "My tickets" with their own chip** —currently All · Assigned ·
> Watching, according to section 29—, are seen
> in the card as **"Watchers"** and **receive the same notices as the assignee**. Tagging belongs to
> **Support and Development**. Decisions 58 to 63, and a **new email template** (the eleventh).
>
> **And the owner added, in the same batch**, that **a watcher can be removed** and that **any technician or developer
> can do it**: that requires **storing watchers in their own
> table** —the remover is not the author of the comment and cannot rewrite it— and leaves a rule that must
> be said out loud: **removing prevails over what was already written**, so the mention remains in the
> comment but **does not add them again**; it returns only if someone mentions them **again** (decision
> 63). Removing someone **is recorded in the ticket history**.
>
> **Amended on 2026-09-27**, when the owner requested **two new columns in the list** written by
> the `ai` module —**"Reason"** and **"Last action"**, in Spanish and English—, and **two names that
> change**: "Number" becomes **"Ticket"** and "Last update", **"Update date"**
> (decisions 56 and 57). This module **does not change schema**: it is `ai` that stores its texts in its
> own table, and here they are only **requested** when moving the ticket and **read in bulk** to paint the
> list. The only new endpoint is the regenerate button, and it belongs to this module.
>
> **Amended on 2026-09-26**, when the owner requested **two things about attachments**: that **text and code
> files** can be sent —a `.sql` did not go in, and that is what is sent when the case is
> a query that fails or a deployment that does not start: **forty-two extensions** are added
> (decision 54)—, and that **images and video do not take the full width** of the conversation, but rather a
> **cap of 480 × 360** that when clicked opens **in the large viewer** to see it whole and download it,
> with video as a **thumbnail with its play button** (decision 55).
>
> **Amended on 2026-09-26**, when the owner requested that **the queue be "mine"** and that it be possible to
> see **everything** from its own place: `GET /api/tickets` debuts the **`mine`** filter —what is assigned to
> me or what I currently watch (decision 51, replaced by section 29)—, **the Queue becomes My
> tickets** with a type chip for Support and Development, and the `tickets` module gains **two lists**:
> **Main tickets** and **Internal tickets**, for both teams (decision 52). **"Assign to me"
> leaves the queue** —reassignment is done inside the ticket, which is where you see who it is being passed
> to, and with the usual rules— and rows say **"Open"** (decision 53). **No changes to the
> data model, transitions, or permissions**: the only new thing in the backend is the filter.
>
> **Amended on 2026-09-26**, when the owner requested that **attachments be seen inside the text**,
> in the place where they are written and in the order in which they are loaded: the text of a ticket and of a
> comment **cease to be plain text and become formatted HTML**, and **inside goes everything that is
> attached** —images and videos as visible elements, PDF and other files as
> links—. **It looks the same when reading as when writing**: the preview is also there while editing,
> which is what was requested. **Four video extensions** are added (`mp4`, `webm`, `mov`, `avi`) and
> **five decisions** (45 to 49). **Emails are not touched**: none carries the text of the ticket or of
> the comments.
>
> **It is the most expensive decision in this document**, and it is taken knowingly: formatted text opens
> the door to hand-written HTML, and therefore it enters with a **tag whitelist**, with **no
> address stored in the text** —attachments are referenced by name and the address is
> resolved when painting— and with the queue search **removing tags before searching**.
>
> **Becomes as-built on 2026-09-26**: what this document describes is what the code does, backend
> and screens.
>
> Approved by the owner on 2026-09-22 **after three question-based reviews**: the
> first fixed the six model and API decisions (the two tables, the identifiers, the
> `svg`, re-escalation, the prefix screen, and 9999); the second, the missing transitions, ticket editing, the three ways to touch a password, and six technical gaps
> —the most serious, that 25 MB attachments would never have arrived because of the nginx limit—; and the
> third, when writing `flujos.md`, the **return to Support** (rule 8) and who resolves the
> main ticket (rule 5, which changed).
>
> It is the third document in the chain. With it, **`auth`, `users`, and `tickets` are ready to
> be implemented**, in that order. `flujos.md` comes later and does not block starting with `auth`.
>
> **Amended on 2026-09-25**, when writing the backup script, with an **owner decision**:
> attachments no longer live loose at the root of `_files/` and move to the **folder of their ticket**
> (`_files/2026/CS-2026-0042/`), with the internal ticket in its own. Section 2.3 explains how it ends up, and the
> file name is still generated: what changes is where it is stored, not how it is called.
>
> **Amended on 2026-09-24**, when starting to implement it, with the question-based review prior
> to code. Five things:
>
> 1. Section 0 said that the matrix of `docs/usuarios-y-permisos.md` was missing the row for
>    **closing the internal ticket**. **It was already there** —the 2026-09-22 amendment added it, as
>    the header of this same document recounts—. It is a typo and is corrected; nothing changes.
> 2. The endpoint table in section 5 had `GET` and `PUT /api/settings`, which **are not from this
>    module**: they live in `docs/modules/settings.md`, which also already lists them. They are removed from here.
> 3. Section 3.3 did not say **among whom** an internal ticket is dispatched when the Administrator
>    activates its dispatch. It is fixed: among active users with the `desarrollo` role, with the same
>    shift criterion. **It is an owner correction**: internal tickets are born **without dispatch** and
>    dispatch is **configurable by the Administrator**, which is the value the database already brings.
> 4. The matrix debuts the row for **reopening the internal ticket** (Support): the transition existed in
>    section 3.2 and was not visible in the matrix.
> 5. The Administrator queue carries a **type chip** (Main · Internal · All), because it sees
>    both and needs to be able to separate them (`docs/interfaz-y-experiencia.md`, section 3.3).
>
> And four others that came out **when writing the service** and that the model had not resolved: the reason
> for a **re-escalation** goes as a comment on the internal ticket (30), a **closed ticket is not edited** (31),
> the two internal notices go to one person and not to the whole team (32), and the assignee of an internal ticket
> comes from Development (33). All are from the owner, from 2026-09-24, and are in the fourth review.
>
> **Amended on 2026-09-24**, when building the screens: section 8 stops listing
> "the numbering prefix screen" as a screen of this module —it is configured in
> **Settings**, with the rest of the installation, and its endpoints belong to `settings`—, and the document
> collects the four decisions that came up when making them (the fifth review): **"New ticket" for
> Support** —its menu debuts the entry, because it can create on behalf of another person and that permission had no
> door; owner correction—, escalation **in a dialog**, attachments **with
> preview**, and that **the detail is a single screen for the four roles**.
>
> **Amended on 2026-09-24, when writing `docs/modules/settings.md`** (approved by the owner the
> same day): this table was called `ticket_number_settings` and **only collected the prefix**, when
> dispatch and the notices of section 3.3 did not fit anywhere. It is renamed
> **`ticket_settings`** and collects the five things: the prefix, assignment, and the notice for main and internal
> tickets.
>
> **Implementation status (2026-09-24): the module is complete**, backend and screens. The backend
> —the eight tables, numbering, dispatch, endpoints, and the eight emails— is tested
> end to end; the screens —the queue, the create form, and the detail with the double view— are done and
> tested in a browser, on PC and mobile. The seven tables in
> `v1.0.0.sql`, race-free numbering, shift dispatch, the thirteen endpoints, the seven email
> notices, and **78 end-to-end checks** against the development environment,
> reading emails from the test mailbox: creation with its dispatch, triage, escalation, the two
> synchronization rules, the return, re-escalation, closing, reopening, comments
> (edit, delete), attachments (upload, download, the invalid extension, the svg, the 25 MB), and the
> permissions of the four roles. **The screens are missing.**

## 0. Closing the internal ticket, and a gap in an approved document

The eight synchronization rules in `docs/propósito-y-alcance.md` **do not say who closes the
internal ticket**. It was decided that **it is closed by hand**, and not as a consequence of closing the
main ticket:

- **Development or Support** closes it, when they consider it finished.
- **Closing the internal ticket does not change the state of the main ticket.**
- **Re-escalating the main ticket reopens it**, even if it was closed (rule 3 of
  `docs/propósito-y-alcance.md`).

That **does not contradict** the approved document, so it does not need to be amended. And the missing permission
—**closing the internal ticket**— **is already in the matrix** since the 2026-09-22 amendment
(`docs/usuarios-y-permisos.md`, section 3): this section said the opposite and is corrected
(2026-09-24). When implementing it, the row for **reopening the internal ticket** was also added, which is the other
half of the same thing.

## 1. Scope of this document

It fixes **the data model**, **the transitions of the two lifecycles** with who can make each
movement and its effects, **the closed list of modules and endpoints**, and **attachments**. It does not define
the step-by-step flows with their emails (`flujos.md`), nor does it repeat permissions: it takes them as written in
`docs/usuarios-y-permisos.md`.

It is **approved**: with it and the two previous documents, `auth`, `users`, and `tickets` can be
implemented in that order (Rule 0 of `AGENTS.md`). Section 9 collects what was decided in the two
reviews, including the decisions that corrected what I proposed.

## 2. The data model

### 2.1 Two tables: `tickets` and `internal_tickets`

Main and internal tickets live in **separate tables**. It has a consequence that must be
accepted in exchange and that is seen in section 2.3: **comments, attachments, and history
point to one table or the other**, so each of those three tables carries **two columns** —one
per type— and a constraint requiring exactly one to be filled. In exchange, each table
has only its own columns, with none for padding.

**`tickets`** — main tickets

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | Internal key. The API does not expose it: it exposes the `number` |
| `number` | `text` | `ACME-2026-0042`. **Unique** |
| `number_year` | `int` | Year of the number, for sorting and filtering |
| `number_seq` | `int` | Sequential **per year** |
| `subject` | `text` | Subject. Required |
| `description` | `text` | Description. Required |
| `state` | `text` | `nuevo`, `en progreso`, `en espera`, `escalado`, `resuelto`, `cerrado` |
| `requester_id` | `bigint` | The person who needs the attention |
| `created_by_id` | `bigint` | The person who created it. May not be the requester |
| `assignee_id` | `bigint` | Assignee. **Set when the ticket is dispatched**, and remains null only if there was no active technician |
| `subject_edited_at` | `timestamptz` | Null until someone edits the subject |
| `description_edited_at` | `timestamptz` | Null until someone edits the description |
| `resolved_at` | `timestamptz` | When it was resolved. Cleared on reopening |
| `closed_at` | `timestamptz` | When it was closed. Cleared on reopening |
| `created_at`, `updated_at` | `timestamptz` | |

**`internal_tickets`** — internal tickets

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | |
| `ticket_id` | `bigint` | Its main ticket. **Unique**: this is what guarantees one internal ticket per main ticket |
| `number` | `text` | `INT-ACME-2026-0042`. **Unique** |
| `state` | `text` | `nuevo`, `en progreso`, `en espera`, `resuelto`, `cerrado` |
| `escalation_reason` | `text` | Why it was escalated. Required, and not edited |
| `created_by_id` | `bigint` | The person who escalated it |
| `assignee_id` | `bigint` | Assignee in Development |
| `resolved_at` | `timestamptz` | |
| `closed_at` | `timestamptz` | |
| `created_at`, `updated_at` | `timestamptz` | |

Decisions that go with the tables:

- **The internal ticket repeats nothing from the main ticket.** Not the requester, not the subject, not the description:
  its content is **the reason for escalation**, and everything else is read from its main ticket. This is where
  the two tables behave better than one: no empty or padding columns.
- **The number is immutable**: it is stored whole in `number`, and is not recomposed from the current
  prefix. If the prefix changes tomorrow, old tickets keep theirs.
- **There is no `deleted_at`** in either of the two: the action of deleting a ticket does not exist.
- The internal ticket stores **its own number** (`INT-…`) instead of computing it on the fly, because it appears in
  all screens and emails.
- **The subject and description can be edited** (the requester and Support, on the main ticket), and each
  one leaves its mark: `subject_edited_at` and `description_edited_at`, one per field to know **what**
  was changed without having to read the history. The previous text is not stored, as with
  comments.
- **The reason for escalation of the internal ticket is not edited.** It is the justification for why it was escalated, and
  rewriting it is rewriting history.

### 2.2 Numbering, without races

The per-year sequence is kept in a table with one row per year:

```sql
ticket_number_counters (year int primary key, last_number int not null)
```

And it is incremented **in the same transaction** that creates the ticket, with a single atomic statement:

```sql
INSERT INTO ticket_number_counters (year, last_number) VALUES (2026, 1)
ON CONFLICT (year) DO UPDATE SET last_number = ticket_number_counters.last_number + 1
RETURNING last_number;
```

Why this way and not with `MAX(number) + 1`: two tickets created at the same time would read the same maximum and
produce the same number. Why not a `SEQUENCE` per year: it forces creating objects in the database every
January 1, and that is a manual step that gets forgotten.

`number_seq` is filled to **four digits** (`0042`), but **is not truncated**: from 9999 on, the
number grows to five digits and remains unique and sortable.

**Only main tickets consume the sequence.** The internal ticket does not spend a number: its own is the
one of its main ticket with `INT-` in front.

The prefix allows **uppercase letters and digits, from 2 to 8 characters** (`^[A-Z0-9]{2,8}$`): `CS` and `ACME`
are valid, `acme-2026` is not. It is validated when storing it, not when using it.

**The prefix is changed by an Administrator from the application**, it is not server configuration, and it
lives with the rest of the installation settings: **the `settings` module**
(`docs/modules/settings.md`). `tickets` reads it from there through its service when numbering a ticket, just
as it reads the dispatch settings.

It is a table with **a single row** and not an "adjustments" table with key and value: such a table ends up being
the drawer where everything goes. With one row you see at a glance **what** is configurable, and adding another
setting requires touching the schema, which is exactly the conversation that should be had.

### 2.3 Comments, attachments, and history

The three tables that follow point to **a main ticket or an internal ticket**: they carry `ticket_id` and
`internal_ticket_id`, and the constraint that **exactly one** is filled. It is the price of
having two ticket tables, and it is paid here and only here.

**`ticket_comments`**

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | |
| `ticket_id` | `bigint` | The main ticket, or null |
| `internal_ticket_id` | `bigint` | The internal ticket, or null. Never both |
| `author_id` | `bigint` | The person writing |
| `body` | `text` | **Formatted HTML, and it may be empty when the comment brings attachments.** It carries inside the attachments that were placed (amendment of 2026-09-26) |
| `edited_at` | `timestamptz` | Null until someone edits it |
| `deleted_at` | `timestamptz` | Null while it remains alive. **Logical** deletion |
| `created_at` | `timestamptz` | |

- **They can be edited, but only by whoever wrote it**, and the comment is marked as
  edited (`edited_at`). Nobody edits what another person said.
- **The previous version is not stored.** The mark says *that* it was edited, not *what* it said before; if
  the previous text were needed, it would be another document.
- **The author can delete their own, and the mark remains.** The row does not disappear: the conversation
  shows that there was a comment there, with its author and date, and **the text is emptied**. It is emptied on
  purpose: if someone pasted a password or data they should not have, deleting must really delete, not hide.
- **Deleting a comment does not delete its attachments**: they remain attached to the ticket and are downloaded like
  any other. An attachment can be the proof of something, and whoever deletes their text is not asking
  for the file to disappear.
- **The original rule did not contemplate main tickets without escalation.** Previously everything the
  User should not read was referred to the internal ticket; the proposal in section 30 replaces that conclusion,
  because a main ticket that is never escalated also needs private coordination.
- **A comment can be only an attachment** (decision 41, from the owner): "look at this" with the
  screenshot and without writing anything is something that happens every day, and forcing a sentence to be written
  to send a screenshot is noise. **It is what forces `body` to stop being required**, and the reason
  is stated in full in the decision: when the comment is created, its files **have not yet been
  uploaded** —they go later, with their `commentId`—, so the backend cannot know whether they will arrive.
  **The rule that does hold is held on the screen**: a comment that has
  **neither text nor files** is not sent, and that is said before sending it.
- **And the reverse**: editing a comment **does** still require text (decision 42). Leaving a comment
  blank through the back door, after it has been written, is not the same as being born without text: what
  is wanted in that case is to delete it, and that is what deletion is for.
- **A comment without text is read through its attachments**: in the conversation it appears with its author, its
  date, and its files, and **without a blank gap** that looks like a screen failure.

**`ticket_attachments`**

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | |
| `ticket_id` | `bigint` | The main ticket, or null |
| `internal_ticket_id` | `bigint` | The internal ticket, or null. Never both |
| `comment_id` | `bigint` | Null if the attachment goes in the initial description |
| `uploaded_by_id` | `bigint` | |
| `filename` | `text` | Original name, as the person uploaded it |
| `stored_name` | `text` | **Path on disk**, relative to the files folder and generated: `2026/CS-2026-0042/ab12….png` |
| `content_type` | `text` | |
| `size_bytes` | `bigint` | |
| `created_at` | `timestamptz` | |

- Files **are not stored in the database**, but on disk, in the folder indicated by
  `FILES_PATH`. In containers it is a mounted volume (`_files`), like `_logs`: attachments
  survive restarts and can be copied from the machine. `dev.yml` and `prod.yml` mount it and
  `config/env/*.env` defines the variable.
- **Each ticket has its folder**: `_files/<año>/<número>/`, with the ticket year and its number as
  shown (`_files/2026/CS-2026-0042/ab12….png`). **The internal ticket has its own**, because its
  number is its own: `_files/2026/INT-CS-2026-0042/`. Thus, looking at the disk, you know which ticket each
  file belongs to and what was uploaded in the internal conversation, without opening the application (owner decision,
  2026-09-25: before they were loose files at the root of `_files/`, and there was no way to know whose
  each one was).
- **The file name is still generated** (32 characters and its extension): the original can be
  repeated, contain accents, or try to escape the folder. **The original stays in the database** to
  show it, and the folder number is also cleaned before use: nothing that comes from outside enters
  a path without going through there.
- **Attachments uploaded before the folders** live loose in `_files/` and **are still displayed
  the same way**: `stored_name` is a relative path, and a loose name resolves to the root of the files
  folder. **None was moved**: in development they were deleted and started over, as the
  owner requested, and in production there is nothing yet.
- If deleting a file leaves the folder of its ticket empty, it is removed too. It only happens in
  failure paths —a file uploaded by mistake, a row that could not be saved—, because **deleting
  already-uploaded attachments is not part of 1.0.0** (section 7).
- **Upload limit in nginx**: the vhosts carry `client_max_body_size 30m;`. Without it, nginx uses
  **1 MB** by default and any attachment over one megabyte would be rejected with a 413 before reaching
  the backend. The 30 leaves margin over the 25 MB of the file.
- **They are downloaded through the API, never through a static route**: the attachment of an internal ticket cannot
  be reachable by the requester, and that is only guaranteed if each download goes through the
  permission check.
- Limits: **25 MB per file** and **no quantity limit**. What is a closed list is
  the **extension**: `pdf`; images `png`, `jpg`, `jpeg`, `gif`, `webp`; **video `mp4`, `webm`, `mov`,
  `avi`** (added on 2026-09-26, decision 47); Office `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`,
  `odt`, `ods`, `odp`; **text and code** `txt`, `csv`, `log`, `md`, `css`, **`sql`, `json`, `xml`,
  `yml`, `yaml`, `ini`, `conf`, `cnf`, `sh`, `bash`, `bat`, `ps1`, `py`, `js`, `ts`, `java`, `php`,
  `go`, `cs`, `rb`, `pl`, `kt`, `rs`, `swift`, `c`, `h`, `cpp`, `hpp`, `vue`, `jsx`, `tsx`,
  `htaccess`, `env`, `properties`, `diff`, `patch`, `bak`, `old`** (added on 2026-09-26,
  decision 54); and archives and backups `zip`, `rar`, `tar`, `gz`, `tgz`, `7z`. Whatever is not in the
  list is rejected: "anything" is not accepted and then looked at.
- **Text and code are stored with their type and always downloaded**: they are never rendered inline,
  so it does not matter what they contain. It is the same that already applied to Office and `zip`, and it is
  what allows admitting a `.sh` or a `.sql` without opening a hole: **the only thing shown inside the
  ticket is images, video, and PDF** (`sePuedePrevisualizar`), and everything else downloads as
  `application/octet-stream` with `attachment`, which is what prevents a file from opening by itself.
- **Text can also name a person, and it does so by identifier**: a mention is
  `<span data-mencion="12">María Pérez</span>`, and **never by the hand-written name**. It is the same
  idea as `data-adjunto`: the name shown is decoration and the identifier is the data, so
  if someone changes their surnames —or there are two people with the same name— the mention still
  points to whom it pointed (decision 58).
- **The user is not rejected for text they did not write** (2026-09-27): if they edit the description
  of a ticket that Support opened on their behalf and that description **already carried a mention**, the body is
  saved —they have not tagged anyone, they have only kept what was there—. What is rejected is that
  **a new mention appears** in a body of theirs.
- **Only active technicians and developers can be tagged** and **only Support and Development can
  do it**: when saving the comment, the module **checks one by one** the mentions against the active
  accounts of those two roles, and if any is invalid **it rejects the entire comment** with
  `tickets.mention.notAllowed` —just as the sanitizer rejects HTML it does not allow: a text that is
  half-saved is worse than one that is rejected—. The user does not tag: they do not see the list of technicians nor
  do they need to (decision 59).
- **`svg` is left out on purpose**, even though it is an image: an SVG can carry code inside and
  opens in the browser. The other image formats do not.
- **No antivirus** in 1.0.0: it is said so that no one takes it for granted.

**`ticket_categories`**

The installation's category catalog. **It belongs to `tickets`** and not to `settings`: they are ticket
data, and its screen is a screen of this module (decision 64).

| Column | Purpose |
| --- | --- |
| `id` | Key |
| `name` | The name, as read: "Network", "Software x" |
| `normalized` | The name in lowercase and without accents, **unique**: it is what prevents the catalog from ending up with "Network" and "network" |
| `active` | Whether it is offered when creating a ticket. **Retiring is deactivating** (decision 66) |
| `created_by_id`, `created_at` | Who created it and when |
| `deactivated_at` | When it was retired, if it was retired |

**`tickets.category_id`**: new column, **`NOT NULL`** with a foreign key to `ticket_categories`
(decision 65). The migration brings it with the **"General"** category already set on whatever exists.

**The tag catalog** (decision 72): the name of each tag, so it can be created without being
carried by any ticket, renamed all at once, and counted.

| Column | Purpose |
| --- | --- |
| `id` | Key |
| `tag` | The name, already normalized: lowercase, hyphenated, without accents. **The normalized form is what is stored**, unlike the name of a category —which is a name that is read—: a tag is a key used to group |
| `normalized` | The same normalized name, **unique**: it is what prevents `red-wifi` and `Red-Wifi` at the same time |
| `created_by_id` | Who created it, **null** if it was born from an updated database or from an earlier tag |
| `created_at` | When |

**`ticket_tags`**

**Which ticket carries which tag**: the bridge table between the ticket and the catalog.

| Column | Purpose |
| --- | --- |
| `id` | Key |
| `ticket_id` | The ticket. Foreign key with **cascade deletion**, like comments |
| `tag_id` | The tag, **foreign key to the catalog with cascade deletion**: retiring a tag deletes its row from the catalog and its links go with it, in a single movement |
| `created_by_id`, `created_at` | Who put it on that ticket and when |

**`ticket_observers`**

The watchers of a ticket: **who follows the ticket without attending to it**. It goes apart from the ticket
table because there are **many per ticket** and because the list must be able to change without touching anyone's text.

| Column | Purpose |
| --- | --- |
| `id` | Key |
| `ticket_id` / `internal_ticket_id` | **One of the two**, and only one: each thread has its watchers, just as it has its comments |
| `account_id` | Who watches. **Foreign key to `users`** |
| `added_by_id` | Who added them —the author of the mention—, or who put them if at some point added by hand |
| `created_at` | When |

- **`UNIQUE (ticket_id, account_id)` and `UNIQUE (internal_ticket_id, account_id)`**: mentioning the same person twice
  does not put them twice. With `NULL` involved —a null column is always different from
  another— the two partial unique indexes are needed, one per side, as in comments.
- **It is deleted for real**: there is no deletion mark. Removing a watcher is removing them, and what remains
  —that they were removed, by whom, and when— is in the history.

**`ticket_history`**

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | |
| `ticket_id` | `bigint` | The main ticket, or null |
| `internal_ticket_id` | `bigint` | The internal ticket, or null. Never both |
| `actor_id` | `bigint` | Who caused the change. Null if the system did it |
| `event` | `text` | `creado`, `editado`, `asignado`, `estado`, `escalado`, `resuelto`, `cerrado`, `reabierto` |
| `from_state`, `to_state` | `text` | Null when the change is not of state |
| `detail` | `text` | Free text: who it was assigned to, why it was escalated… |
| `created_at` | `timestamptz` | |

History is **what the system did**; comments are **what people said**. The
screen timeline is the mixture of the two ordered by date, and history is not
written by anyone by hand: each transition generates it.

#### The text of a ticket and of a comment is HTML, and attachments go inside

**What someone writes in a ticket description or in a comment is formatted text**, and
**the files they attach live inside that text**, in the place where they put them, in the order in which
they loaded them (decision 45, from the owner, 2026-09-26).

```text
Hola mi PC no funciona, me aparece estos errores:

[captura1.png]          <- esto es una imagen, y se ve la imagen

[captura2.png]          <- y esta otra

Solicito su ayuda por favor
Gracias
```

**What is seen in each place** (decision 46):

| What is attached | Inside the text | In the list at the end |
| --- | --- | --- |
| **Image** (`png`, `jpg`, `jpeg`, `gif`, `webp`) | **The image inline**, with the width of the conversation | Same as the inline ones, with their buttons |
| **Video** (`mp4`, `webm`, `mov`, `avi`) | **An inline player**: the text stores `<video data-adjunto="grabacion.mp4"></video>` and **the screen adds the controls when painting it** —the text does not carry them, because an extra attribute is rejected— | |
| **PDF** | A **link** with the file name: clicking it opens **the viewer** | |
| **Everything else** (Word, Excel, text, archives…) | A **link**: clicking it **downloads it** | |

**And it looks the same when writing as when reading** (the owner requested it this way): while writing, each
image and each video **is really seen** in its place, not as text representing them. A normal text
box cannot carry an image inside, so **the writing box is a formatted editor**
(`docs/interfaz-y-experiencia.md`, section 6.5), with its button bar —bold, italic,
list, link— **and the attach button**, which puts the file where the cursor is.

**The four rules that keep this from becoming a hole**:

1. **Tag whitelist.** The backend **does not store any HTML**: paragraphs, line breaks,
   bold, italic, underline, strikethrough, lists, links, images, and videos are accepted, and nothing else. A `<script>`,
   an `onerror`, an `<iframe>`, or a `<style>` **are rejected** —the key is `tickets.body.notAllowed`—,
   and what arrives outside the list is not silently cleaned: it is said, because a text that is saved
   halfway is worse than one that is rejected. **The reference of an attachment is mandatory**: an `<img>`,
   a `<video>`, or an `<a>` without it are rejected, because a gap is nothing. **HTML comments
   do pass**, which are not tags and do not execute: the browser inserts an `<!--StartFragment-->` when
   pasting formatted text, and rejecting it would break pasting.
2. **No address is stored inside the text.** An attachment is referenced **by its name**
   —`<img data-adjunto="captura1.png">`— and **the address is resolved when painting**, with the list of
   attachments of the ticket that has already been read. Thus **there is no way to insert an image from outside** (a
   counter that notifies a third party that someone opened the ticket), nor a giant `data:`, nor an
   address that expires. A `src` that is not the one of an attachment of **that** ticket is rejected.
3. **The text is not stored with provisional addresses.** While writing, the editor shows
   the file from the browser (`blob:`); that **does not reach the database**: what is stored is the
   reference by name, and the `blob:` is dropped on save.
4. **The queue search removes tags before searching.** The queue searches by number,
   subject, and text (`ILIKE`), and searching for "p" or "img" over raw HTML would find **all** tickets
   with an image. The query **removes the tags** before comparing. What is left out, and it is said:
   a hand-written `&amp;` is not found by searching for "&", and **the search does not look inside
   comments**, only the subject and description, as before.

**Attachments that the text does not name are still shown at the end**, in their list, as before
(decision 49). It is what makes **tickets and comments that already exist look the same** and ensures that
nothing attached before this decision is lost. It is also where those uploaded without
being able to be placed fall.

**And an attachment inside the text cannot be deleted by itself**: removing it from the text **does not delete the
file** —it remains in the list at the end, because an attachment can be the proof of something (section 2.3)—.
For it to disappear, it would be necessary to be able to delete attachments, which remains **outside 1.0.0** (section 7).

**And there is a place where the text writes itself**: the **reason for a re-escalation**, which decision 30
sends as a comment on the internal ticket. That reason is plain text —the escalation dialog has no editor—,
so **it is escaped when writing it** (`&`, `<`, `>`), or a reason with a `<` would be read as a
tag. The `escalation_reason` of the internal ticket remains plain text and stays as it is.

**The loading order**, which is what is behind all this:

1. When attaching —by pasting, dragging, or with the button— **the file is uploaded** and the editor shows it
   while writing continues.
2. When sending, **the text is stored with the reference by name** already in its place.
3. If a file **does not upload**, the rule is that of decision 44: the comment remains published, it is
   said which file is missing, and the button uploads it **to the same comment**, without writing it again. What
   changes is that now the gap of the missing file **is also seen in the text**, as a box
   saying that file did not upload.

**Ticket creation needs nothing special**, and that is the advantage of referencing by name: the
ticket is created **with the references already in place** and its files are uploaded afterwards, hanging from the
ticket; when painting it, each reference is resolved with the attachment that has that name. The person sees a
single create button, and the application does not save the text again along the way.

### 2.3.1 Watchers

**A ticket has at most one assignee and many watchers**, and they are not the same: the assignee **is
who attends to it** —and remains **optional**, as it was: a ticket is born without an assignee when there
is no active technician to dispatch it to—, and watchers **follow the ticket without attending to it**
(decision 60, owner correction).

| | Assignee | Watcher |
| --- | --- | --- |
| How many | **At most one** | **As many as needed** |
| Who | A technician (main) or a developer (internal), by shift or by hand | Any active technician or developer, **tagged** |
| What they do | Attend: comment, move the state, close | Watch: **the same as anyone else who sees the ticket**, they gain no permissions |
| Where they come from | From dispatch, or from the card | **From the text** —from the mentions in the comments of that thread— **or by hand, from the ticket card** (decision 74) |
| Email | The dispatch one and the ticket notices | **The same staff notices as the assignee**, and the one that they were tagged |
| Where they see it | In "Assigned to me" | In **"Watching"**, inside "My tickets" |

- **Being a watcher grants no permissions**: a watcher sees the ticket because they could already see it —both teams
  see main and internal tickets—, and still cannot comment where it is not their place (Development does not
  write on the main ticket, `docs/usuarios-y-permisos.md`, rule 3).
- **Watchers of an internal ticket belong to the internal ticket** and those of the main ticket to the main ticket: each thread has
  its own, because each is read separately (section 2.1).
- **They are stored in their table** (`ticket_observers`), and are not inferred from the text on the fly: **removing a
  watcher can be done by any technician or developer**, and the remover is not the author of the
  comment that named them, so the list must be able to live without the text (decision 63).
- **They are added when writing the comment**: when saving a comment —new or edited— **those mentioned
  again are added** (in a new comment, all; in an edited one, those that were not before
  in that same comment) and **those no longer mentioned by any comment in the thread are removed**. With that,
  both things that were decided hold: **deleting the mention by editing the comment removes the
  watcher**, and **removing by hand prevails** —someone removed is not brought back by a later save, only
  by a new mention, which is a deliberate call—.
- **Adding by hand leaves its entry in the history** (`observador_anadido`, with the name of who
  enters), as does removing (`observador`). **And only when it is really added**: repeating it leaves no other.
- **Removing is an action of the card**, with the same rule as tagging: Support and Development. And
  **it remains in the history** ("María removed Juan from the watchers"), like everything that happens in the
  ticket.
- **"New mention" means really new**: a **new comment** that again names
  someone who was removed **does add them again** —it is a deliberate call, and that is why it counts—; what does not
  bring them back is **saving the same text again** (an edition of the comment that already carried it).
- **Someone removed remains named in the comment**: nobody's text is rewritten. The mention
  is read as what it is —a call that was made— and the watcher list says who follows it
  now.

### 2.3.2 Categories and tags

**Two different things, with two different jobs** (decisions 64 to 68):

| | **Category** | **Tags** |
| --- | --- | --- |
| How many | **Exactly one**: there is no ticket without a category | **As many as needed**, or none |
| Where they come from | From a closed **catalog** maintained by **only the Administrator** (decision 83) | From a **catalog** maintained by **only the Administrator** (decisions 83 and 84): **they no longer emerge when written** |
| Who sets it | **Whoever opens the ticket** —the user included—, and **Support can correct it** later. **Development cannot**: it does not write on the main ticket | **Support and Development**, each from the ticket they work on: **Support on the main ticket and Development on the internal ticket** (decision 85). They belong to the main ticket, and the internal ticket inherits them |
| What it is for | **Grouping**: "Network", "Software x", "Licenses"… it is the *what it is* of the ticket | **Nuancing**: what does not deserve its own category |
| How it is written | As named in the catalog | **Lowercase and with a hyphen instead of spaces**: `red-wifi`, `licencia-office`. No accents or uppercase, and the editor keeps it that way while writing |
| Where it lives | Ticket column, with a foreign key to the catalog | Its table, many per ticket |
| Where it is seen | **A chip on the ticket line**, below the number, and in the card | The same chips, next to it |

**The category belongs to the case, not to each half**: it is stored in the main ticket and **the internal ticket inherits it when
read**, as it inherits the subject (decision 67). The same for tags.

**"There is no ticket without a category" is guaranteed by the database, not the screen** (decision 65): the column is
`NOT NULL` with a foreign key, so not even a manual query can leave a ticket unclassified. For
that to hold from day one, **the migration creates the "General" category and sets it on everything
that already exists** —in production there are no tickets, in development there are—, and **the catalog cannot be left
without any active category**: retiring the last one is rejected with a key.

**Retiring a category is deactivating it** (decision 66), as an account is deactivated: it stops being
offered when creating a ticket, and **tickets that already have it keep it and continue to show it**.
Nothing is deleted, and that is why **there cannot be a ticket without a category** —deleting the category of a
ticket in use would leave it without one—.

**The tag catalog is maintained only by the Administrator** (decisions **83** and **84**, which
correct 72 and 73): they can be **created** on the screen, **renamed** —and renaming **changes the
tag on all tickets that carry it**, not on one— and **retired** —and retiring it **removes it from those
tickets**, with a warning beforehand of how many are affected—, and **all three things belong to the Administrator**.
**They no longer emerge when written**: putting a tag on a ticket that is not in the catalog is **rejected**
(`tickets.etiqueta.desconocida`), so **whoever tags chooses from those that exist** —which is what the
card modal does— and **new names are debuted by the Administrator** on that screen.

**And who puts tags on a ticket is both teams** (decision **85**): **Support on the
main ticket and Development on the internal ticket**, each from the ticket they work on. Tags **belong
to the main ticket** and the internal ticket **inherits them** (decision 67), so changing them from either one
changes **the same ones**.

**Tags are normalized on save**: lowercase, spaces become hyphens, and anything that is not
an unaccented letter, number, or hyphen is dropped. **Existing ones are suggested** while writing,
so they do not end up being five ways of saying the same thing. **They emerge on their own** —a tag exists the
first time someone writes it in a ticket— and since **2026-09-27 they also have a catalog**
(decision 72): its name lives in **`ticket_tag_names`** and `ticket_tags` says **which ticket carries it**.
Thus they can be **created without anyone carrying them**, **renamed** —and the change applies to all tickets
that carry it, because the name is in a single place— and **retired** —which removes them from those tickets—.
**A single source of truth**: the name in the catalog, and the bridge table for the rest.

**And both serve the same purpose: finding** (decision 68). The category is filtered with its chip, the
tags with theirs, and **both enter the search box**: searching for "network" finds the
tickets of the "Network" category and those carrying the `red` tag.

### 2.4 Constraints

- `tickets.state` only admits the six states of the main ticket; `internal_tickets.state`, the five
  of the internal ticket. **The internal ticket cannot be `escalado`**: the internal ticket *is* the escalation.
- `internal_tickets.ticket_id` is **unique**: a single internal ticket per main ticket, guaranteed by the database
  and not by the code.
- In `ticket_comments`, `ticket_attachments`, and `ticket_history`, **exactly one** of the two
  destination columns is filled.
- `internal_tickets.escalation_reason` cannot be empty: escalating without saying why is what
  turns an escalation into a "pass it on yourself".

### 2.5 Indexes

| Index | Purpose |
| --- | --- |
| `tickets(number)` unique | Search by number, which is how people talk about a ticket |
| `internal_tickets(number)` unique | The same, for internal tickets |
| `internal_tickets(ticket_id)` unique | The internal ticket of each main ticket, and only one |
| `tickets(state)` | Support queues |
| `internal_tickets(state)` | Development queue |
| `tickets(requester_id)` | "My tickets" |
| `tickets(assignee_id)` | The tickets of a person |
| `ticket_comments(ticket_id, created_at)` and `(internal_ticket_id, created_at)` | The conversation in order |
| `ticket_history(ticket_id, created_at)` and `(internal_ticket_id, created_at)` | The timeline in order |
| `ticket_attachments(ticket_id)` and `(internal_ticket_id)` | The attachments of a ticket |
| **`ticket_tags(tag_id)`** | **Filter by tag** and count how many tickets carry it (section 2.3.2) |
| **`ticket_tags(ticket_id, tag_id)`** unique | The same tag is not repeated on the same ticket |
| **`ticket_tag_names(normalized)`** unique | So the catalog does not end up with `red-wifi` and `Red-Wifi` |
| **`ticket_categories(normalized)`** unique | So the catalog does not end up with "Network" and "network" |
| **`tickets(category_id)`** | Group and count the tickets of a category |

Text search uses `ILIKE` on subject and description, **and since 2026-09-27 also on the
name of the category and the tags** of the ticket (decision 68). It is enough for the volume of an
internal help desk; when it stops being enough, full-text indexes are adopted with their own
change.

## 3. Transitions

### 3.1 Main ticket

| From | To | Who | Effects |
| --- | --- | --- | --- |
| — | `nuevo` | User (their own) or Support (on behalf of another) | The number is assigned; **it is dispatched by shifts**; `creado` history; email **to the assigned technician** |
| `nuevo` | `en progreso` | Support | When assigning it to someone or when starting |
| `nuevo` | `escalado` | Support | Creates the internal ticket. Requires a reason. Email to Development |
| `en progreso` | `en espera` | Support | Support asks the user for information. **Email to the user** |
| `en espera` | `en progreso` | Support or the user (upon replying) | |
| `en progreso` | `escalado` | Support | Creates the internal ticket. Requires a reason. Email to Development |
| `en espera` | `escalado` | Support | Same as the previous one |
| `en progreso` | `resuelto` | Support | Resolved **without escalating**, or after Development resolved the internal ticket. Email to the user |
| `en espera` | `resuelto` | Support | What was expected arrived. **There is no need to go through `en progreso`** |
| `nuevo`, `en progreso`, or `en espera` | `cerrado` | Support or the user (their own) | **Close without resolving**: duplicates, tests, tickets created by mistake |
| `escalado` | `en progreso` | System | Caused by the internal ticket when it moves to `en espera` (rule 2) |
| `escalado` | `en progreso` | System | Caused by the internal ticket when it is resolved: **the case returns to Support** (rule 5), which explains to the user what was done and resolves it |
| `resuelto` | `escalado` | Support | Re-escalate without reopening: the internal ticket returns to `en progreso` |
| `resuelto` | `cerrado` | Support or the user | **Does not touch the internal ticket**: that one is closed separately (section 3.2). **Email to the user** |
| `cerrado` | `en progreso` | Support or the user | Reopen. Escalating again requires going through here |

Four rules that come out of the table:

- **A closed ticket is not escalated.** To escalate it again, it is first reopened.
- **Nobody jumps from `nuevo` to `resuelto`**: if it was resolved, someone worked on it.
- **Closing without resolving is possible**, from any open state, and the difference remains in the
  history: there is no `cancelado` state because "six states and not one more", and because a ticket
  closed without resolution does not behave differently from one closed after resolving it.
- **On reopening, `resolved_at` and `closed_at` are cleared**, so the screen does not say "resolved on"
  with an old date. The history preserves when each thing happened.

### 3.2 Internal ticket

| From | To | Who | Effects |
| --- | --- | --- | --- |
| — | `nuevo` | System | Created by the escalation of the main ticket, with its reason |
| `nuevo` | `en progreso` | Development | When assigning it or when starting |
| `en progreso` | `en espera` | Development | Development needs something from Support. The main ticket moves to `en progreso` (rule 2). **Email to Support** |
| `en espera` | `en progreso` | Support or Development | Support replies |
| `en progreso` | `resuelto` | Development | **The main ticket returns to `en progreso`**, in the hands of Support (rule 5). **Notice to Support** |
| `en espera` | `resuelto` | Development | What was expected arrived. **There is no need to go through `en progreso`** |
| `nuevo`, `en progreso`, or `en espera` | `cerrado` | Development or Support | **Close without resolving** (for example, escalated by mistake). **Returns the main ticket to `en progreso`** (rule 8). **Notice to Support** |
| `resuelto` | `cerrado` | Development or Support | **By hand**, when they consider it finished. Changes nothing about the main ticket |
| `cerrado` | `en progreso` | Support | Reopen the internal ticket: **it is escalating again**. The main ticket moves to `escalado` (rule 3) |

The internal ticket **has no `escalado`**: the internal ticket *is* the escalation. And its closing is **manual**, not
a consequence of the main ticket: it is closed when Development or Support considers it finished, and re-escalating
returns it to `en progreso` even if it was closed.

**Closing the internal ticket without resolving it returns the main ticket to `en progreso`** (rule 8): it is the
return to Support. If it were left in `escalado`, the ticket would remain in the state that means
"Development has it" without anyone from Development looking at it.

**And resolving the internal ticket does not return the case to the user by itself**: the main ticket moves to
`en progreso` so that Support writes what was done. The explanation the user reads cannot be
the escalation reason or Development's notes.

### 3.3 Dispatch of new tickets

A ticket is not born in the air: **it is dispatched by shifts** among the active Support technicians.

- **Only active technicians** with the `soporte` role enter. A deactivated technician does not receive tickets.
- **The shift is decided by who has gone longest without receiving one**: technicians are ordered by the date
  of their last assigned ticket and the first one is chosen. Whoever has never received one goes first.
- **Thus there is no need for a shift pointer** to store and that can become desynchronized: the list of
  tickets already says who got the last one. If a new technician joins tomorrow, they receive the
  next one; if one is deactivated, the shift continues without skipping anyone forever.
- **The account is kept by `tickets`**, which is who has the table; the list of technicians is requested from
  `users` through its service, which is how modules cross (`docs/arquitectura.md`, section 4).
- **If there is no active technician**, the ticket remains **without an assignee** and stays in the Support
  queue. There is no one to notify, so the notice is not sent and it is recorded in the log.
- **Tickets created by Support also go through the shift**, even if they are created by the person who will
  attend to them.
- **Reassigning is always possible**: dispatch decides the starting point, not the destination.

**And all this is configurable**, because there is no single sensible way to dispatch. The administrator
chooses, **for each ticket type**, two things separately:

| Setting | Values | What it means |
| --- | --- | --- |
| **Assignment** | `ninguna` or `por turnos` | Whether the ticket is born without an assignee or dispatched |
| **Notice** | `a nadie`, `a todo el equipo`, or `al asignado` | Who receives the new ticket notice |

Both are independent on purpose: you can dispatch without notifying —technicians look at their queue—
or notify the whole team without dispatching, which is how a ticket that nobody has taken yet works.
With `asignación = ninguna`, the `al asignado` notice makes no sense and the screen does not offer it.

The default values, which are the ones decided here:

| | Assignment | Notice |
| --- | --- | --- |
| **Main tickets** | `por turnos` | `al asignado` |
| **Internal tickets** | `ninguna` | `a nadie` |

Internal tickets are born without dispatch and without notice: Development team communication is internal and
they are the ones watching their queue.

**Among whom each type is dispatched** (clarified on 2026-09-24, when implementing it): the shift for
**main** tickets is kept among active users with the `soporte` role, and for **internal** tickets,
if the Administrator activates it, among active users with the `desarrollo` role. Each
type keeps **its own account**: whoever has gone longest without receiving one of that type goes first. The
dispatch of internal tickets **comes turned off** (`ninguna`) and it is the Administrator who turns it on
(owner correction, 2026-09-24); the list of people for each role is given by `users` to
`tickets` through its service, like the rest.

## 4. Attachments: whose they are

Attachments follow the same rule as comments, which is what the approved matrix says: **if
you can comment on a ticket, you can attach to it**. It is not a new permission, it is the application of
the one that already exists.

| Ticket | Who attaches | Who downloads |
| --- | --- | --- |
| Main | The requester and Support | The requester, Support, and Development (Development read-only) |
| Internal | Support and Development | Support and Development |

The requester **never** reaches an attachment of the internal ticket: neither through the API nor through a file route,
because no such route exists.

**Three ways to attach**, and all three lead to the same place (decision 43): **paste** (`Ctrl+V`), with the
cursor **in the writing box or in the files box**; **drag and drop** onto the box; and **choose the file**
with the button. At a help desk half of the attachments are screenshots, and pasting them is what everyone does without thinking.

- **When pasting, whatever comes in the paste is accepted**, whether a screenshot or a file copied from the
  explorer, **if it is in the closed list of extensions**; what is not is rejected **with the same
  notice** as when choosing it with the button. One rule for the three ways.
- **Pasting only works with focus where writing happens**: in the comment box and in the files
  box. Pasting anywhere else on the screen **does not** add anything, so that pasting does not
  surprise anyone.
- **Files go with the comment**: first the comment is published and then its files are uploaded,
  one by one, with their `commentId`. **If one fails** (decision 44), the comment **remains**
  —it is written and whoever wrote it has seen it, and deleting it would leave a gap in the conversation—, the
  box is emptied, and **which file is missing** is announced with a button to **retry** without writing the comment again. It is what prevents the duplicate comment, which is what happened when the
  failure was counted as if nothing had happened.

**How they are served**: the types that can be previewed —**images** (`png`, `jpg`, `jpeg`,
`gif`, `webp`) and **videos** (`mp4`, `webm`, `mov`, `avi`)— are served **inline** so they can be seen
inside the ticket, and so is **PDF**, because the viewer is what shows it. **Everything else is always served as a download**, with `application/octet-stream` and `Content-Disposition: attachment`.

The case that forces this rule is **`svg`**, which **is not admitted**: it can carry code inside and the
browser executes it when opening it, and since 2026-09-26 images are shown **inline** inside the
text, which is exactly what cannot be done with a file that executes. The same applies to
any format that the browser interprets instead of showing: what is not in the closed list
is rejected.

## 5. Modules and their endpoints

Six modules, and no more:

| Module | What it covers |
| --- | --- |
| `auth` | Sign in, sign out, who I am, passwords |
| `users` | Accounts, roles, deactivation, origin |
| `tickets` | Main, internal, comments, attachments, and their history |
| `mail` | Email templates, their editing, and sending |
| `settings` | Installation settings: language, appearance, numbering, and dispatch |
| `ai` | The two ticket summaries: **reason** and **last action**, in Spanish and English (`docs/modules/ai.md`). **Without its own endpoints**: it is reached through `tickets` |

The two ticket types live in **a single module** (`tickets`): they are the same thing told in two
ways, and separating them would force two modules to share tables, which is exactly what the modularity
rule forbids.

### Those of other modules, in their place

The endpoints of `auth`, `users`, and `mail` **are not listed here**: each one details them in its document
(`docs/modules/auth.md`, `docs/modules/users.md`, and `docs/modules/mail.md`), and this section keeps
those of `tickets`. One list per module, and none repeated: two lists of the same contract end up
saying different things.

### `tickets`

| Method and route | What it does | Who |
| --- | --- | --- |
| `GET /api/tickets` | Queue and lists. Filters: `type`, `state`, `q` (text or number), **`mine`** (only mine), and pagination | According to the matrix: each one sees what corresponds to them |
| `POST /api/tickets` | Creates a main ticket. `requesterEmail` optional, to create it on behalf of another | User, Support, Development |
| `GET /api/tickets/assignees` | Who can be the assignee of each ticket type | Support, Development, and Administrator |
| `GET /api/tickets/{number}` | Detail, with comments, attachments, and history | Whoever can see it |
| `POST /api/tickets/{number}/comments` | Comments. **The body may be empty**: its files are uploaded later, so here it is not known whether they will arrive (decision 41) | Whoever can comment |
| `PATCH /api/tickets/{number}/comments/{id}` | Edits a comment. Only its author; leaves the edited mark | The comment author |
| `DELETE /api/tickets/{number}/comments/{id}` | Deletes a comment: empties the text and leaves the deleted mark | The comment author |
| `POST /api/tickets/{number}/attachments` | Uploads an attachment (`multipart`) | Whoever can comment |
| `GET /api/tickets/{number}/attachments/{id}` | Downloads, with permission check | Whoever can see the ticket |
| `PATCH /api/tickets/{number}` | Edits the subject, description, or category, with its mark. **If the body brings only `tags`**, it is the tags door: **Development can use it from an internal number**, and the change goes **to the main ticket**, which is whose tags they are (decision 85) | The requester and Support, on the main ticket. **Only tags**: also Development, from the internal ticket |
| `POST /api/tickets/{number}/assign` | Assigns an assignee | Support; Development on internal tickets |
| `POST /api/tickets/{number}/state` | Moves the state: `en progreso`, `en espera`, `resuelto`, `cerrado`, with `{ "state": "…", "comment": "…" }` —**the state is also accepted in `to`**, which is the name of this table, and both are valid—. **When closing, the comment is mandatory** (decision 81): without it it responds **422** with `tickets.cierre.sinComentario`, and with it **it remains as a comment of the conversation** | Whoever can make that movement (section 3) |
| `POST /api/tickets/{number}/escalate` | Escalates. Requires a reason. Creates the internal ticket or reopens it | Support |
| `POST /api/tickets/{number}/reopen` | Reopens a closed main ticket | Support or the requester |
| `GET /api/tickets?view=…` | In addition to the usual filters, **`view`**: `assigned` (assigned to me), `watching` (those I currently watch), or empty (the union of both). It is the "My tickets" selector. **`view` nuances "mine", it does not replace it**: `mine=1` remains the switch, and `assigned`/`watching` imply it by themselves. The general lists send neither `mine` nor `view` and continue returning everything (section 29) |
| `GET /api/tickets/categories` | The catalog: active ones for everyone —the create form needs them— and retired ones for those who maintain them |
| `POST /api/tickets/categories` | Creates a category | Support and Administrator |
| `PATCH /api/tickets/categories/{id}` | Renames it | Support and Administrator |
| `POST /api/tickets/categories/{id}/state` | Retires it or puts it back | **Only the Administrator** |
| `POST /api/tickets/tags` | Creates a tag | Support and Administrator |
| `PATCH /api/tickets/tags/{tag}` | **Renames it on all tickets that carry it** | Support and Administrator |
| `DELETE /api/tickets/tags/{tag}` | **Retires it and removes it from tickets** | **Only the Administrator** |
| `GET /api/tickets/tags?q=…` | The tags that already exist —**at most ten**, by use— to **suggest** them while writing. **Includes catalog ones that no ticket carries**, with `tickets: 0` | Anyone who has signed in |
| `GET /api/tickets/tags?all=1` | **The entire catalog**, without a cap and by name: it is the listing of the maintenance screen, and without it **a newly created tag that no ticket carried was not seen** (found by the test of the person who made the catalog, with ten tags already in use). With `all`, `q` is ignored: it is the catalog, not a search | Anyone who has signed in |
| `POST /api/tickets/{number}/observers` | **Adds a watcher by hand**, with `{ "accountId": 12 }` in the body. Responds **201** with `{ "ticket": { …, "observers": [ … ] } }`, and **is idempotent**: adding someone who already watches is not an error and leaves no other history entry | Support and Development |
| `DELETE /api/tickets/{number}/observers/{id}` | Removes a watcher. Responds `{ "status": "ok" }` —**not** the card—, and the screen reloads on its own | Support and Development |
| `POST /api/tickets/{number}/insights` | Requests the **reason** and **last action** from the engine again (`docs/modules/ai.md`) | Support and Development |

- **The API is addressed by number, not by `id`.** `ACME-2026-0042` is what people say out loud
  and what appears in emails; the internal `id` does not leave the database.
- **`mine` is "mine", and it is asked in the backend, not on the screen** (updated in section
  29). For Support and Development a ticket is mine if **it has been assigned to me** or if **I currently watch
  it**; requesting or commenting is not enough. The filter is applied to the query —with an
  `EXISTS` over the current watchers— and not by bringing everything to discard it afterwards: the list is paginated, and filtering on the screen would give pages of different
  sizes. **A user still sees only their own**, with or without `mine` (section 3 of
  `docs/usuarios-y-permisos.md`).
- **The internal ticket is read with the same endpoint**, passing its number (`INT-ACME-2026-0042`): it is a
  ticket, and does not need its own route.
- **The only deletion that exists is that of a comment, and it is logical**: the row remains with its
  mark. A ticket is not deleted, and neither is an attachment.
- **The requester of a ticket created on behalf of another is identified by their email**
  (`requesterEmail`), and not by their identifier: email **is** the identifier of people in
  this product (`docs/usuarios-y-permisos.md`, section 3), and it is what Support has at hand when
  someone calls. Thus their screen does not need to read the list of accounts, which belongs to another module.
- **`GET /api/tickets/assignees` exists so that the screen does not have to call `/api/users`**:
  a frontend module only talks to its own API (`docs/arquitectura.md`, section 4), so the
  list of possible assignees is given by this module, which already asks `users` for the active technicians for
  the shift. It returns the two lists —Support and Development— and each screen uses its own.
- **`GET` and `PUT /api/settings` are not in this list on purpose**: they belong to `settings` and live in
  `docs/modules/settings.md`, section 7, which is where they are read. Here they were repeated and have been
  removed (2026-09-24): two lists of the same contract end up saying different things.
- **The settings live in `settings`**, and the frontend Settings screen lives in
  `core` because it belongs to the installation and not to a module: it is the exception recorded in
  `docs/arquitectura.md`, section 4. Only an Administrator sees it.
- The lists are paginated. The error format is the only one of the API: **`{"error": "clave"}`**, with
  a **key** and not a assembled text, defined in `shared/httpx`. The frontend decides how it is said and
  in which language, because the interface is in Spanish and English
  (`docs/interfaz-y-experiencia.md`). It is the same criterion as in Calibyou, where errors travel
  as `reception.documents.template.unknownField`.

### The notice that you have been tagged

It is **template number eleven** (`ticket.mentioned`), the only one that debuts this function, and it goes **to each
mentioned person**, once per comment that mentions them:

| | |
| --- | --- |
| **When** | When saving a comment with a **new** mention |
| **To whom** | To whoever has just been mentioned, and **not** to whoever was already mentioned before (mentioning the same thing twice does not send two emails) |
| **What it carries** | The number, the subject, and the link —like the other notices: **never the text of the ticket or of the comment** |
| **Subject** | "You have been tagged in {{numero}}: {{asunto}}" and its English version |

And **watchers also receive the same staff emails as the assignee** (decision 61): the
dispatch one, the escalation one, the "Development needs something" one, and the "the ticket returns to Support" one.

## 6. The migration

The whole schema enters **`backend/migrations/v1.0.0.sql`**, the single, transactional, and
idempotent file commanded by `docs/arquitectura.md`. New tables:

| Table | What it stores |
| --- | --- |
| `tickets` | Main tickets |
| `internal_tickets` | Internal tickets, one per main ticket |
| `ticket_comments` | The conversation of both types |
| `ticket_attachments` | The attachments of both types |
| `ticket_history` | What the system did in both types |
| `ticket_number_counters` | The sequence of each year |
| `ticket_settings` | One row: the numbering prefix **and dispatch and notices per ticket type** |
| **`ticket_categories`** | **The category catalog** of the installation, with its name, its normalized name (unique), and whether it is active (section 2.3.2) |
| **`ticket_tag_names`** | **The tag catalog**: the normalized and unique name of each one |
| **`ticket_tags`** | **Which ticket carries which tag**: the bridge table with the catalog, unique per ticket and tag, with who put it |

The update **also puts the tags that already existed into `ticket_tag_names`** —those that were
as text in `ticket_tags`— and links the tickets with their `tag_id` before removing the old column, so
a database with earlier tags updates itself (decision 72).

And **`tickets` gains `category_id`**, `NOT NULL` with a foreign key: **every ticket is born classified**. The
migration **creates "General" and sets it on everything that already exists** before applying the `NOT NULL`, so
a database with half-done tickets updates itself (decision 65).

The migration **seeds the prefix with `CS`**, so that the first ticket can be created without going
through the settings screen. If it is not liked, it is changed before creating the first one.

The migration **does not seed any account**: the factory administrator account (`admin`) is created by the
**backend startup** from `ADMIN_PASSWORD`, because the migration cannot read environment
variables (`docs/usuarios-y-permisos.md`, section 8). And there is no more sample data: tickets are
created by people.

## 7. What does NOT enter 1.0.0

- **Deleting another person's comment** and **recovering a deleted one**: the emptied text is not stored.
- **Editing or deleting attachments** already uploaded.
- **Version history**: the previous text of a comment or of an edited subject is not stored.
- **Editing the escalation reason** of an internal ticket.
- **Deleting tickets.**
- **Antivirus** and analysis of attachment content.
- **External storage** (S3 and similar).
- **Full-text search** with dedicated indexes: `ILIKE` is used until it hurts.
- **Exporting** tickets to CSV or PDF.
- **Predefined responses** and templates.
- **Custom fields** other than category and priority.
- **Real time** (websockets) and live counters.
- **SLA, metrics, and reports.**
- **Ticket merging and duplication.**

## 8. Frontend modules

| Module | Screens |
| --- | --- |
| `tickets` | **My tickets** (mine, with type chip for Support and Development), **Main tickets**, **Internal tickets**, the Administrator queue, the ticket detail (conversation, attachments, history), ticket creation, and the escalation form |
| `users` | User list, creation, deactivation, and origin change |
| — (`core`) | Sign in, sign out, password recovery, and session guards, as commanded by `docs/usuarios-y-permisos.md` |

**The four lists are the same screen** with a different filter and a different name: separating them would
duplicate the table, filters, and pagination to change one condition. **Their routes are four**
—`/tickets`, `/tickets/main`, `/tickets/internal`, and the Administrator's, which is `/tickets`—, and the
two of "everything" **are only offered to Support and Development** (2026-09-26).

**The numbering prefix is not a `tickets` screen**, although it was listed here: it is
configured in **Settings**, with the rest of the installation, and its endpoints belong to `settings`. A
screen of this module that called `/api/settings` would break the hard modularity rule
(`docs/arquitectura.md`, section 4). Corrected on 2026-09-24.

## 9. What was decided in the reviews

### First review: the model and the API

| # | Decision | It remained this way |
| --- | --- | --- |
| 1 | **Tables** | **Two tables** (`tickets` and `internal_tickets`), with two destination columns in comments, attachments, and history |
| 2 | **Identifiers** | `bigserial` inside; the API is addressed by the ticket number |
| 3 | **`svg`** | **It is admitted**, but always as a download, never inline. **Corrected on 2026-09-26**: the code never admitted it —it is not in the closed list of extensions, neither in the backend nor on the screen—, and section 2.3 says so from its first amendment. This row and section 4 said the opposite, and are corrected: **`svg` is not admitted**, and even less now that images are shown inline inside the text, which is exactly what cannot be done with a file that the browser executes |
| 4 | **Re-escalating from `resuelto`** | Directly, without forcing a reopen. From `cerrado`, it must be reopened first |
| 5 | **The number prefix** | **Administration screen**, with a row in `ticket_number_settings` |
| 22 | **The prefix table, renamed** | Amended on 2026-09-24: it becomes `ticket_settings` and also collects dispatch and notices, which had no table |
| 6 | **The sequence passes 9999** | The number grows to five digits; it is never repeated or truncated |

### Second review: transitions, editing, and gaps

| # | Decision or gap | It remained this way |
| --- | --- | --- |
| 7 | **Closing without resolving** | It is possible, from any open state (duplicates, tests). There is no `cancelado` state |
| 8 | **Resolving from `en espera`** | Allowed, in both ticket types, without going through `en progreso` |
| 9 | **Editing the ticket** | Subject **and** description, with a mark per field (`subject_edited_at`, `description_edited_at`). The escalation reason is not edited |
| 10 | **Passwords** | The `local` user changes it whenever they want; when registering they create it from the email; an administrator can **launch a reset**, which also goes by email |
| 11 | **nginx upload limit** | `client_max_body_size 30m;` in both vhosts: without it, nginx cut at 1 MB |
| 12 | **Attachments on disk** | `FILES_PATH` variable and `_files` volume in `dev.yml` and `prod.yml`, like `_logs` |
| 13 | **Email** | New section in `docs/arquitectura.md` (section 9): `shared/mail` with `net/smtp`, `SMTP_*` variables, and an email failure **does not** bring down the action |
| 14 | **Prefix format** | Uppercase and digits, from 2 to 8 (`^[A-Z0-9]{2,8}$`) |
| 15 | **Login attempts** | `limit_req` zone for login and for requesting the recovery email, in both vhosts |
| 16 | **Dates on reopening** | `resolved_at` and `closed_at` are cleared on reopening; the history preserves when each thing happened |

### Third review: the return and closing (when writing `flujos.md`)

| # | Decision | It remained this way |
| --- | --- | --- |
| 17 | **The return to Support** | Closing the internal ticket without resolving it **returns the main ticket to `en progreso`**: it is **rule 8**, new in `docs/propósito-y-alcance.md` |
| 18 | **Who resolves the main ticket** | **Support**, always, and with its explanation for the user. **Rule 5** changes: when Development resolves the internal ticket, the main ticket returns to `en progreso` instead of remaining resolved |
| 19 | **Commenting on a closed ticket** | It cannot be done: to keep talking it must be reopened |
| 20 | **Deadline to reopen** | None |
| 21 | **The email format** | Plain text, with the number, subject, and link |

The gaps of the second review (11 to 15) are applied in the vhosts, in `dev.yml`/`prod.yml`, in
`config/env/*.env`, and in `docs/arquitectura.md`, not only written here.

### Fourth review: what was missing when starting to implement it (2026-09-24)

| # | Decision or finding | It remained this way |
| --- | --- | --- |
| 23 | **The row that already existed** | Section 0 said that the matrix was missing "closing the internal ticket": it had been there since 2026-09-22. Section corrected |
| 24 | **The repeated endpoints** | `GET` and `PUT /api/settings` leave the `tickets` table: they belong to `settings`, which already lists them |
| 25 | **The internal shift** | Among active users with the `desarrollo` role, with its own account per type. **Owner correction**: internal tickets are born **without dispatch**, and the Administrator turns dispatch on |
| 26 | **Reopening the internal ticket** | New row in the matrix: **Support**. It was the missing half of "closing the internal ticket" |
| 27 | **The Administrator queue** | **Type chip** (Main · Internal · All): it sees both and needs to separate them |
| 28 | **The work breakdown** | **The entire backend first** —tables, numbering, dispatch, endpoints, and the seven emails— and the screens afterwards, against an API that no longer moves |
| 29 | **The emails** | They go out **in the same batch as the backend**: the notice is part of the transition that causes it, and they are tested by reading the mailbox |
| 30 | **The reason for a re-escalation** | It goes **as a comment on the internal ticket**, written by whoever escalates: the original reason is not rewritten —it is the justification of the first time— and the new one remains where Development reads it |
| 31 | **Editing a closed ticket** | **It is not edited**: it is reopened first. It is the same rule already decided for comments, and it does not leave two ways of saying the same thing |
| 32 | **Who the two internal notices go to** | "It has returned to your queue" to the **assignee of the main ticket**, and "Development needs something" to **whoever escalated**. If that account is no longer active, all active Support is notified instead: it is what makes "your queue" mean something and not an email for everyone |
| 33 | **The assignee of an internal ticket** | Comes from **Development**: active users with that role, who are also among whom the shift dispatches if the Administrator turns it on |

### Fifth review: the screens (2026-09-24)

| # | Decision or finding | It remained this way |
| --- | --- | --- |
| 34 | **"New ticket" for Support** | Its menu debuts the entry, and there are three: Support can create a ticket on behalf of another person and that permission had no door. **Owner correction** |
| 35 | **Escalation** | **A dialog**, with the reason as the main and mandatory field, like account creation |
| 36 | **Attachments** | **With preview** for images and PDFs; everything else is always downloaded, which is what the backend already did |
| 37 | **The prefix and dispatch** | They are configured in **Settings**, not on a `tickets` screen: their endpoints belong to `settings` and the hard modularity rule forbids it. Section 8 of this document listed them as its screen and is corrected |
| 38 | **The detail** | **A single screen for the four roles**: what changes is what each one can do inside, not the screen |
| 39 | **Creating on behalf of another** | The requester is identified **by their email** (`requesterEmail`), not by their identifier: it is what Support has at hand and prevents their screen from reading the account list |
| 45 | **Attachments go inside the text** | **Owner decision, 2026-09-26**: what is written in a ticket description and in a comment **is formatted text**, and **the files attached live inside it**, in the place where they were put and in the order in which they were loaded. **It is the most expensive decision in the document**: the text stops being plain text, and with it coexist the tag whitelist, the search that removes tags, and the migration that converts what already exists |
| 46 | **What is seen of each attachment** | **Owner decision, 2026-09-26**: **images and videos inline** —with the width of the conversation and a player with controls—; **PDF as a link**, which when clicked opens the viewer; **everything else as a link**, which when clicked downloads. **The preview is also seen while editing**, which is what was requested: the writing box is a formatted editor, not a text box |
| 47 | **Four video extensions** | **Owner choice, 2026-09-26**, and **broader than my recommendation**: `mp4` and `webm` are admitted —which any browser plays— **and also `mov` and `avi`**, which are camera formats and which many browsers **cannot play**. Consequence stated and accepted: if the browser cannot, the player does not start and the viewer's download button remains. **The 25 MB cap does not grow**, so they are a few seconds of video |
| 48 | **The viewer** | **Owner decision, 2026-09-26**: the modal that already existed —image or PDF inside— becomes **larger** and carries **Download** and **Open in a tab**. It is also the way to see an image at full size, because inline it goes with the width of the conversation |
| 49 | **What the text does not name, at the end** | **Owner decision, 2026-09-26**: an attachment that the text does not reference **is still shown in its list, at the end**, as before. It is what makes **tickets that already exist look the same** and ensures that nothing attached before this change is lost |
| 41 | **A comment can be only an attachment** | **Owner choice, 2026-09-25**, and **not the one I recommended**: I proposed continuing to require text; he decided that only the file can be sent. **Consequence stated and accepted**: `body` **stops being required** in the table, because when the comment is created its files have not yet been uploaded and the backend cannot know whether they will arrive; the rule "text **or** files" is held on the screen, which is who knows it |
| 42 | **Editing a comment still requires text** | **Owner decision, 2026-09-25**: the back door is closed. Leaving blank a comment that was already written is not the same as being born without text, and what is wanted in that case is to delete it |
| 43 | **Pasting, dragging, and choosing the file** | **Owner decision, 2026-09-25**: the three ways lead to the same place and share the closed list of extensions and its notice. **When pasting, whatever comes is accepted** —screenshot or copied file— if it is of an admitted type, and **pasting only works with focus in the writing box or in the box**, so that it does not add files unintentionally |
| 44 | **If an upload fails, the comment remains** | **Owner decision, 2026-09-25**: the comment is already published, so it is not deleted; **which file** did not upload is announced and **retrying** to the same comment is possible. Before, a failure there left the box with the text again and the most likely outcome was two identical comments |
| 50 | **Who reassigns, and from where** | **Owner decision, 2026-09-26**: **reassignment is done inside the ticket** —in its card, with the assignee selector and its button— and **"Assign to me" disappears from the queue**, in the table and in the cards. The rules do not change: on the main ticket **Support** reassigns, to **any active technician**; on the internal ticket **Support and Development** reassign, to **any active developer**. Thus `support1` passes their ticket to `support2` from the ticket, which is where you see who it is being passed to; the list shortcut assigned without more, and that was what left reassignment undone |
| 56 | **Two columns written by the AI engine** | **Owner decision, 2026-09-27**: the list debuts **"Reason"** —what the ticket is about— and **"Last action"** —what was the last thing that happened—, **written by the `ai` module** in Spanish and English (`docs/modules/ai.md`). **The texts live in the `ai` table**, not in the ticket one: here they are only requested and read, and the tickets module **does not change schema** |
| 57 | **The two header names** | **Owner decision, 2026-09-27**: **"Number" becomes "Ticket"** —it is what the column shows and what the thing inside is called— and **"Last update" becomes "Update date"**, which says it is a date and not an event. The right column becomes the date, and the event is told by **"Last action"** |
| 54 | **Text and code are also attached** | **Owner decision, 2026-09-26**, when a `.sql` could not be sent: the closed list remains **closed** and **text and code** are added —`sql`, `json`, `xml`, `yml`/`yaml`, `ini`, `conf`, `cnf`, `sh`, `bash`, `bat`, `ps1`, `py`, `js`, `ts`, `java`, `php`, `go`, `cs`, `rb`, `pl`, `kt`, `rs`, `swift`, `c`, `h`, `cpp`, `hpp`, `vue`, `jsx`, `tsx`, `htaccess`, `env`, `properties`, `diff`, `patch`, `bak`, `old`— **and archives** (`tar`, `gz`, `tgz`, `7z`). **They are never rendered inline**: they are served as `text/plain` and downloaded, which is what allows admitting a `sh` without opening a hole. **The alternative he discarded** was opening the entire list with a blacklist: more convenient, and it lets in anything no one has thought of |
| 55 | **Image and video, with a cap and to the viewer** | **Owner decision, 2026-09-26**: in the description and in comments images and videos are seen **at most at 480 × 360**, without deforming, and **when clicked they open in the large viewer** —with **Download** and **Open in a tab**—, which is where they are seen whole. **Video becomes a thumbnail with its play button** and is played **in the viewer**: before, it sent its player inside the conversation and the viewer only appeared when the browser could not play it. **What is stored does not change** —the reference without `controls`—, and the width is not stored data either: it is how it is painted. **The cap is written `min(100%, 480px)`**, and not `480px`: a maximum only limits, so with the fixed value a 480 px image **overflowed the screen on a 412 mobile** and the document widened to 530 —found by the Playwright layer, because clicking a button on the same card stopped working— |
| 64 | **Categories and tags belong to `tickets`** | **Owner decision, 2026-09-27**: the category catalog lives in this module and its screen is **a ticket screen** —"Categories and tags"—, not a Settings section: they are ticket data, and a frontend module only talks to its own API (`docs/arquitectura.md`, section 4). **It is maintained by Support and the Administrator**, and **retiring** a category is **only the Administrator's** job |
| 65 | **There is no ticket without a category, and the database guarantees it** | **Owner decision, 2026-09-27**: **one category per ticket, mandatory**. It is implemented with the **`NOT NULL`** column and a foreign key, not with a screen check, and the migration **creates "General" and sets it on everything that already exists**. The catalog **cannot be left without any active category**: retiring the last one is rejected |
| 66 | **Retiring a category is deactivating it** | **Owner decision, 2026-09-27**, over mine of forbidding deleting the one in use: **retiring is deactivating** —it stops being offered when creating a ticket and **tickets that have it keep it**—, as an account is deactivated. **No row is deleted**, and it is what makes "there is no ticket without a category" always fulfillable |
| 67 | **The category belongs to the main ticket; the internal ticket inherits it** | **Owner decision, 2026-09-27**: the category and tags are **of the case**, they are stored in the main ticket and **the internal ticket inherits them when read**, as it inherits the subject. Classifying the same case twice —one per half— would be asking for extra work |
| 68-quater | **The catalog count counts main tickets, and the filter respects the list** | **Clarified on 2026-09-27**, because it looked like a mismatch: the number shown by the catalog next to each category is **the main tickets** that have it, and **filtering by that category in the internal list returns the internal tickets** —which inherit it from their main ticket—. That is, **it is not the same question**: "how many tickets have this category" (the main ones: the category is theirs) and "which tickets appear when filtering by it in this list" (those of that list, and the internal ticket inherits it so it can be found). It is left this way, and **it is said here** so that no one reads it as a failure |
| 68-ter | **The keys of the catalog and of tags** | `tickets.category.required` (422), `tickets.category.notFound` (404), `tickets.category.inactive` (422), `tickets.category.name.required` (422), `tickets.category.duplicate` (409), `tickets.category.lastActive` (409), `tickets.category.forbidden` (403), `tickets.tag.required` (422), and `tickets.tag.tooLong` (422) —**there is no `tag.invalid`**: normalization turns what is invalid into empty, and that is `tag.required`, so a key that cannot come out is not included—. All with their text in Spanish and English (`docs/interfaz-y-experiencia.md`, section 8) |
| 69 | **A single search button, with a modal** | **Owner correction, 2026-09-27**, upon using it: "I already have three rows for four criteria, and I can have more". The criteria —text, view, type, state, category, and tag— live in **a modal** opened by the "Search" button, and **at the top of the list there remains a summary** of what is filtered with its remove button: a clean list and, at the same time, it is understood why few tickets appear. The overflow that caused it —the category row overflowing the screen— **is also fixed** (section 5) |
| 70 | **The ticket view, by role** | **Owner correction, 2026-09-27**: when opening a ticket, **Support sees the main ticket** and **Development the internal ticket**, and both can switch to Main · Both · Internal whenever they want. **If you open an internal number, the internal ticket is seen** even if you are Support: it is what whoever clicks that link expects. And **"Both" shows only the conversations** —the two left columns—, without the cards: with the four columns "it looks very saturated" |
| 71 | **"What can be done" goes away** | **Owner correction, 2026-09-27**: the title of the card's actions card **was misinterpreted** and is removed. The buttons still say what they do, which is what was needed |
| 72 | **Tags are maintained, with the same weight as categories** | **Owner correction, 2026-09-27**: on the screen there was no way to **create or correct** a tag and **categories took all the spotlight**. **A screen with two equal halves** is chosen —not two modules—, and tags can be **created**, **renamed**, and **retired**. Renaming **applies to all tickets** that carry it, and retiring **removes it from them**, warning beforehand how many are affected: a tag lives in many tickets, and changing one cannot change only one. **They still emerge on their own** when written in a ticket |
| 73 | **Who maintains tags** —**corrected by 83: only the Administrator maintains them** — | **Historical proposal, 2026-09-27**, initially implemented on 2026-09-28: creating and renaming corresponded to Support and the Administrator; retiring, only the Administrator. Decision 83 replaced it and expressly closed it: today the entire catalog is maintained only by the Administrator |
| 80 | **The state control, and the mandatory comment when closing** | **Owner correction, 2026-09-29**, over 78: the dropdown **is not a list of actions**, it is **for changing the ticket state**, and **the list of states** is presented on main and internal tickets. And **closing is always confirmed in a window and asks for a mandatory comment**, also from the backend (see 81). The list shows **the states the ticket can move to now**, with **the current one as the first option** (marked, and not selectable), and **the one that needs a text asks for it in that same window**: the reason when escalating, what was done when resolving, the question when putting it on hold, and **the comment when closing**. **Named actions do not disappear: they are states** — verified by whoever implemented it in the code itself—: on an internal ticket, **"On hold" is returning the case to Support without closing it** (rule 2 of section 3.3) and **"Closed" without resolving is returning it and closing it** (rule 8), which is what "I need something from Support" and "It is not a code change" did; and on a main ticket, **"On hold" is asking the user** and **"Closed" from closed is reopening it**, which is what "Ask the user" and "Reopen" did |
| 82 | **Tags belong to Support and Development, and the create form does not ask for them** | **Owner decision, 2026-09-29**: **the user does not set tags** —nor sees them in their ticket—: they are a way of **classifying and filtering** used by **Support and Development**. So **the create form has no tags field** (the category does, which is mandatory and chosen by whoever opens the ticket) and **the tags line of the card is not shown to the user**. And tagging is **better**: in the card, a button opens **a modal with all catalog tags** —with how many tickets each one carries—, a **search**, and **checkboxes to mark and unmark several at once**; those the ticket already carries appear checked and **"Save" applies everything at once** (and "Cancel" changes nothing) |
| 90 | **The number rules, and the "internal" sign is unnecessary** | **Owner decision, 2026-09-29**: in the ticket header, **the number is more visible** —it is the largest text of the row, and it grows on wide screens— and **the "The internal ticket" sign** next to it **disappears**: with the internal number and its state you already know which one is being looked at, and the sign was noise. **And icon-only buttons are shrunk**: they measured 44 px and the card pencils weighed too much; now 36, which is still above what can be tapped |
| 89 | **The list of whom to tag, next to the cursor** | **Owner decision, 2026-09-29**: when writing `@` in a comment, the list of people **appears next to the cursor** —where writing is happening— and not above the comment block, which is where it was and where the eye does not find it. The list is placed **below the line being written**, and if it does not fit there, **above**, always within the screen. The **"Tag" button** still opens the same list in its usual place, for those who do not use the shortcut |
| 87 | **The header shows your team's ticket first** | **Owner decision, 2026-09-29**: in the ticket header, **the most relevant number is that of the ticket the viewer works with**: for **Development, the internal ticket** —and the main one remains as a reference, dimmed—; for **Support and the Administrator, the main ticket** —and the internal one, as a reference—. **And the links "View the internal ticket" and "View the main ticket"** disappear from the header of each column: to switch from one to the other there is already **the switcher** "Main · Both · Internal", which is the place where what is seen is chosen |
| 88 | **"On hold" is only confirmed; closing still asks for the reason** | **Owner decision, 2026-09-29**, which corrects part of 80: moving the ticket to **"On hold"** —on the main ticket, which was "Ask the user"; on the internal ticket, which was "I need something from Support"— **is confirmed in its window and no longer asks for any text**. **Closing still requires the comment** (decision 81): it is the only state that asks for it. The reason when escalating and what was done when resolving **are still requested**, which are the ones really needed to understand the ticket later |
| 86 | **The lists: the number alone, and a classification column** | **Owner decision, 2026-09-29**, looking at the queue: in the summary tables, the **"Ticket"** column **shows only the number** (the classification chips dirtied it), **"State"** keeps the state, and **category and tags are grouped in a new column: "Classification"**. It is where they are read together, which is what they are: how the ticket is classified |
| 85 | **Tags are set by both teams, and belong to the main ticket** | **Owner decision, 2026-09-29**, upon clarifying 82: tags **are applied to the main ticket** and **the internal ticket inherits them** (decision 67), and **both Support and Development can put them, change them, and remove them** on tickets. On the screen, each from **the ticket they work on** —Support on the main ticket, Development on the internal ticket—, and **changing them from either one changes those of the main ticket**, because they are the same: there are not two sets of tags. **What does not change**: Development does not write the main ticket text (subject, description, or category), and the catalog is still curated by the Administrator (decision 84) |
| 84 | **The catalog is curated: tags do not emerge when written** | **Owner decision, 2026-09-29**, when it was verified that 72 and 83 contradicted each other: until today, writing a new tag in a ticket **created it in the catalog** —measured: Support put `etiqueta-nueva-de-soporte` on a ticket, it responded **200**, and it appeared in the catalog—, so "only the Administrator creates tags" was not fulfilled. Now **tagging a ticket with a name that is not in the catalog is rejected** (`422`, `tickets.etiqueta.desconocida`) for anyone who is not an Administrator: **whoever tags chooses from those that exist** —which is what the card modal does— and **new names are debuted by the Administrator in the catalog**. The Administrator **does not touch tickets** (looks without buttons), so their path is the catalog screen, and from there the tag becomes available to everyone |
| 83 | **The catalog is maintained only by the Administrator** | **Owner decision, 2026-09-29**, which **closes the doubt still open in 73**: **creating, renaming, and retiring** —both **categories** and **tags**— is **only the Administrator's** job. Support and Development **use** the catalog —classify tickets and filter with it— but **do not change it**. It is 73 corrected: there it was proposed that Support could create and rename, and the owner has decided the opposite |
| 81 | **A ticket is not closed without saying why** | **Owner decision, 2026-09-29**: closing —whether done by Support, Development, or the requester— **requires a comment**, and **the backend** requires it, not only the screen: it is requested in the confirmation window and stored **as a comment of the conversation**, so whoever reads it sees why it was closed. Through the API, closing without a comment responds **422** with `tickets.cierre.sinComentario`. It is the same rule as the mandatory category: what cannot be missing is guaranteed by the server |
| 78 | **A single control to move the ticket, instead of four buttons** | **Owner decision, 2026-09-28**: the loose buttons of the card —"Ask the user", "Escalate", "Resolve", "Close" on the main ticket, and "I need something from Support", "It is not a code change", "Resolve", and "Close" on the internal ticket— **are brought together in a dropdown**, so that the card is not a wall of buttons. It is decided **by actions and not by states** (`docs/interfaz-y-experiencia.md`, section 8): the dropdown shows **what each one does** —"Start", "Ask the user", "Escalate", "Resolve", "Close", "Reopen"—, and **the state is still seen in its chip**. Choosing one executes it, and **if that action asks for something** —the reason, the question, what was done— **its window appears as before**; those that ask keep asking. **Where it lives**: with a single ticket in view, **in the top row**, next to "Main · Both · Internal"; with **both** in view, **in the header of each column**, because each ticket has its state and from a single control it would not be known which one is being changed. **And "Escalate to Development" becomes "Escalate"**: the word —and the "Escalated" state— already say which team it belongs to |
| 77 | **The internal header, and centered windows** | **Owner correction, 2026-09-28**: (1) in the ticket header, **also the internal number** —opening an internal ticket showed only that of the main ticket— and **the button says "Back"**, not "Back to the queue"; (2) **the windows appeared in the corner**, because the Tailwind preparation sets `margin: 0` on everything and the native `<dialog>` centers itself with `margin: auto`: it is restored, with maximum height and its own scrolling; and (3) **the window's ✕ is called "Close"**, not "Cancel", which was the same name as the button inside |
| 76 | **The ticket icons are the project's** | **Owner correction, 2026-09-28**: the copy-number button and the edit pencil "are not easily distinguished". They were two characters —`⧉` and `✏`—, and the second one **comes out as a colored emoji** depending on the machine's font. Now they are **the project's own SVG** from `app-icono` —two overlapping sheets to copy, a check to confirm it, and a slanted pencil to edit—, and **`app-boton` receives the icon by name**, not by its character. It is what the menu already did: its own drawings, without icon libraries |
| 75 | **A list header, in a single row** | **Owner decision, 2026-09-28**: the **module name**, the **filters**, and **"New ticket"** go **in the same row** —they were two lines—, to gain that height for the ticket table. It applies to the three lists (`Mis tickets`, `Tickets principales`, and `Tickets internos`); at a short width they distribute themselves |
| 74 | **Editing where it is shown, and watchers added by hand** | **Owner correction, 2026-09-28**, upon using the ticket: he could not find where to change the **category** or the **tags** —they were in a separate form, inside the actions card—, and the subject and description, in a dialog. From here on, **each piece of data carries its control where it is read**: the category with its dropdown, tags with their chips and their field to add, watchers with their "Add" —**new**: until now they were only reached through a mention—, and the subject and description with their pencil in place. **Separate forms are removed.** Who can edit it does not change: the requester on their ticket and Support |
| 68-quinquies | **The tag catalog is requested whole, and the suggestions one remains capped** | **My correction, 2026-09-27**, upon finding the failure: the same route served two purposes with opposite needs —suggesting while writing (ten, by use) and maintaining the catalog (everything)—, and with ten tags already in use **a newly created one was not seen on the maintenance screen**, which is exactly what it was requested for. The route accepts **`all=1`**, and the complete listing goes **by name** (stable order: grouping by use would make the list dance when using or renaming) |
| 68-bis | **What the implementation fixed, and the document did not** | **Noted on 2026-09-27**, when closing the backend part: **a tag does not exceed 32 characters** (`tickets.tag.tooLong`) —`licencia-office-2026` fits, it is read in a chip and does not spoil the list—; **suggestions are at most 10**, by use and then by name; the catalog **shows retired ones only to Support and Administrator** (others receive only active ones, which is what the create form needs); **the factory "General" was created by no one** (`created_by_id` null) and tags do carry who set them; the category's foreign key is **`RESTRICT`** —a category with tickets is not deleted, and that is why **retiring is deactivating**— and tags go with **`CASCADE`**; the normalization of the **category** only lowercases, removes accents, and trims (inner spaces are respected, because it is a name that is read) while that of the **tag** also turns spaces into hyphens, joins repeats, and drops anything that is not a letter, number, or hyphen —which is why `red_wifi` is `redwifi`—; and the filter for **internal tickets** joins the main ticket table, because the category and tags are theirs |
| 68 | **Both serve to filter and to search** | **Owner decision, 2026-09-27**: category and tags have **their filter chips** in the list and, in addition, **enter the search box**: searching for "network" finds tickets of the "Network" category and those of the `red` tag. **Tags are free text in lowercase and with a hyphen** (`red-wifi`) —**owner correction** over "lowercase" alone: spaces become hyphens—, with **suggestions** from those that already exist, and without a catalog to maintain |
| 58 | **The mention is stored by identifier** | **Owner decision, 2026-09-27**: the text names a person as `<span data-mencion="12">María Pérez</span>`, **never by the hand-written name**: names are repeated and change, and a mention stored by name breaks by itself. It is the same idea as `data-adjunto`. The tag and the attribute enter the sanitiser's **whitelist**, with the same rule as the rest |
| 59 | **Who tags, and whom** | **Owner decision, 2026-09-27**: **Support and Development**, and **only active technicians and developers**. When saving the comment they are checked one by one, and an invalid one **rejects the entire comment** with `tickets.mention.notAllowed`. **The user does not tag**: they do not see the list of technicians nor need to |
| 60 | **Watcher is not assignee** | **Owner correction, 2026-09-27**, over my proposal, which spoke only of "participants": the ticket has **at most one assignee —the one who attends— and many watchers —those who follow it—**, and they are not the same. **The assignee remains optional**, as it was: the owner corrected his own correction upon seeing that a ticket without active technicians has no one to dispatch to. **Being a watcher grants no permissions**: it only makes the ticket be in your queue |
| 61 | **Watchers receive the same notices as the assignee** | **Owner decision, 2026-09-27**, and **not the one I recommended** —I proposed only the tagging notice—: whoever watches receives **the same staff emails** as the assignee, in addition to the one that they were tagged. **Consequence stated and accepted**: a ticket with three watchers sends four emails per movement. The notices that go **to the user** (we resolved it, we closed it, we are asking you something) **are not duplicated** to watchers: they belong to whoever opened the ticket |
| 62 | **How watched tickets are seen** | **Updated on 2026-10-19 by section 29**: "My tickets" shows in its first row **All · Assigned · Watching** and starts on **All**. The ticket card keeps its list of **"Watchers"**. |
| 63 | **Removing a watcher** | **Owner decision, 2026-09-27**: **they can be removed**, and **any technician or developer** can do it —not only whoever tagged them—. That requires **storing watchers in a table** (`ticket_observers`), because the remover is not the comment author and cannot rewrite it, and leaves the rule that must be said out loud: **removing prevails over what was already written** —the mention remains in the comment, as proof of the call that was made, but **does not add them again**; it returns only if someone mentions them **again**—. Removing someone **is recorded in the history** |
| 51 | **What "mine" is** | **Replaced on 2026-10-19 by section 29**: for Support and Development only the current assignment or watch counts. The decision number is kept here to explain the evolution, but opening and commenting no longer include the ticket. |
| 52 | **Three lists: mine, all main, and all internal** | **Updated on 2026-10-19**: **My tickets** is the personal queue; Support gathers its assigned or watched main tickets and Development its assigned or watched internal tickets. **Main tickets** and **Internal tickets** retain the authorized general access. Cumulative Admin does not alter the operational level; the pure Administrator keeps their general read-only queue. |
| 53 | **The texts and menu that go away** | **Owner decision, 2026-09-26**: the list rows say **"Open"** (`Open`) and not "Open the ticket" —in an actions column, the word "ticket" is superfluous—; **out goes the text** "What has been waiting longest goes first. Closed ones do not get in the way." —the rule remains, silent: the order is seen in the list—; **"New ticket" leaves the side menu**, because it was confused with a module, and stays **inside the queue**, which is where it is created; and **language becomes a dropdown** like the theme one, with the same width and without the visible label |
| 40 | **Who can be the assignee** | **This module** says it (`GET /api/tickets/assignees`), not `users`: the frontend screen can only talk to its own API |

## 10. What this document enables

Approving `docs/modules/tickets.md` closes the list of modules and leaves **`auth`, `users`, and `tickets`** ready
to be implemented, in that order:

1. `auth` — without a session there is nothing to protect.
2. `users` — Support needs to register people, and the prefix is configured.
3. `tickets` — the product.

`flujos.md` comes later: it describes the step-by-step journey of each case (creation, triage, escalation,
resolution, closing, and reopening) with its emails, and does not block starting with `auth`.

## 11. Implementation: summaries and notices in the global language

The `Resumen` contract and the list and detail DTOs stop carrying two texts. Each field carries
state, a single `text`, `language`, provider/model, and error key. The screen shows the preserved
text even if its language is older than the global one; during an individual regeneration it shows
pending.

The module continues rebuilding the text for `ai`, but also gives it the global language. A global
change does not touch existing summaries or create a batch. New ones and those recalculated
by a ticket movement or by "Summarize again" use the current global language.

Ticket notices stop reading `persona.Language` and stop grouping recipients by language. They ask
`mail` for a single send in the global language. Recipients, moments, permissions, states,
comments, attachments, and transitions do not change. The owner confirmed this scope on 2026-10-06.
On 2026-10-07 the owner corrected the scope: there will be no accumulated regeneration or cost
confirmation for summaries; commercial confirmation corresponds only to translating templates. No open
decisions remain; correction explicitly approved.

## 12. Implementation: improving drafts from a ticket

Support and Development see "Improve with AI" next to the description and comment editors they can
already use, both when creating content and when editing it. The User who requires support does not see it
and cannot invoke it through the API. The Administrator does not receive this capability either: their role remains
observation and configuration, without operational intervention in tickets.

`POST /api/tickets/{number}/writing/improve` receives the editor type (`description` or `comment`),
the draft, and the tone. Before delegating to `ai`, the backend applies
the existing permissions of the ticket and the field. For a main ticket it obtains the name of whoever
opened it; for an internal ticket it obtains the name of the Support person linked to the case. `tickets` does not
expose conversation or attachments to this journey and `ai` does not write directly to any ticket
table.

The controller extends the write deadline to 270 seconds only during this operation. Validation,
authorization, and reading of the recipient occur before calling the engine; a rejection does not
consume generation. The global limit for the other responses remains at 60 seconds.

Accepting a proposal only changes the local content of the editor. Saving the description or publishing
the comment remains a separate action, with the existing validations, sanitization, attachments,
mentions, and history. States, notices, editing permissions, and transitions do not change.

The full backend suite and the 224 frontend tests pass. Playwright runs the operation
on PC and mobile against the disposable compatible adapter: it confirms prior review, local application
without a published comment, and absence of the button for the User. Production was not started or changed.

## 13. Implementation: writing capability in the detail

The detail response incorporates `capabilities.aiWriting`, calculated by the backend for the requested
person and ticket. It is `true` only when the three conditions coincide: the role is
Support or Development, they can use at least one of the writing editors present in that view, and
an AI is configured. The absence or temporary failure of the engine does not change the existing permissions
or the authorization of `POST /api/tickets/{number}/writing/improve`.

`tickets` internally asks `ai` whether there is configuration, without importing the settings module
or publishing another route. The detail does not expose modality, provider, model, address, credential, or
health. The frontend takes this capability as the only signal to mount the button and the modal.

The capability is recalculated on each read of the detail. If the configuration changes while
the card is open, the screen adopts the change on its next normal reload; no polling or real-time connection is added. The service tests cover a missing adapter, false configuration,
and true configuration, in addition to the four roles and a closed ticket. A test of the AI
module confirms that a configured address preserves the capability even if it does not respond. The
Playwright journey is limited to the configured case, because the global block prevents entering tickets
without the mandatory AI. The owner chose 1A, 2A, and 3A and confirmed keeping that mandatory nature; no open
decisions remain.

The full backend suite and the 225 frontend tests pass on 2026-10-08. The configured
Playwright journey passes on PC and mobile. The disposable environment was removed at the end and
production was not touched.

## 14. Implementation: summaries in the global language and normal absence of the internal ticket

`tickets` continues giving `ai` the global language of the deployment and preserves literally the
content written by people. The scaffolding it builds —labels for subject, description,
state, attachments, and conversation, in addition to history events— uses that same language. No
language is added to the account, session, ticket, or client request.

The module only receives a ready summary when `ai` has validated its language. A clearly
different response does not appear in lists or details marked with the requested language: it remains pending during
the allowed retry or moves to `ai.invalid`. Previous texts are preserved and only change through
a normal movement or "Summarize again"; no mass regeneration is created.

Reading a main ticket without an internal ticket preserves exactly its functional result:
`InternalByTicket` returns "does not exist" without error. The query is formulated so that GORM does not
record that expected absence as an error. Real database errors continue to propagate and
be recorded; the logger level is not globally lowered nor are other queries silenced.

The tests verify a never-escalated main ticket with no error output, a real error that does
propagate, and an existing internal ticket that is returned. The public repetition of block 8 also confirms
that the normal detail does not fill the logs with `record not found`. API, data model,
permissions, states, transitions, and interface do not change. The owner chose 3A and the coordinated decisions
1A–12A, with 10B. No open decisions remain.
### 14.1 Implementation and verification

The summary context obtains the global language when building each assignment and localizes its labels, system actors, and known event names without modifying the content written by people. `InternalByTicket` uses a query that distinguishes zero rows through `RowsAffected`, so the normal absence does not trigger the `record not found` log; real errors continue to propagate. The service and repository tests, in addition to the full backend suite, passed on 2026-10-13.

## 15. Implementation: expected absences in tickets and catalogs

The rule already applied to the nonexistent internal ticket is extended to the other queries of the module whose
contract admits absence: searching for a ticket, comment, or attachment to respond with its domain
result, and checks of category or normalized name before creating or rejecting a
duplicate. Zero rows will be distinguished with `RowsAffected`; any real error will continue to propagate.

The tests preserve the current responses and check the logger level, including the
creation of a new category without `record not found` and an SQL failure that must indeed remain as
an error. The catalog, seeder, numbering, and no endpoint change.

The singular searches for tickets, comments, attachments, categories, and tag names now use
`Find` and `RowsAffected`. The test with real PostgreSQL confirmed an absent category without an
error event and a real SQL failure visible; the existing tests preserved the domain
results. The full backend suite and `go vet` passed on 2026-10-13.

## 16. Ticket priority

`tickets.priority` is mandatory and admits only `low`, `medium`, `high`, and `critical`, with
`medium` as the default value. The open migration `v1.0.0.sql` adds the column idempotently,
will fill any previous row with `medium`, and will apply the constraint. Although there are no
production tickets to convert, this sequence keeps the migration repeatable.

Creation accepts the priority proposed by the User or by Support when creating on their behalf. A route
of the module allows changing it on open tickets: Support from the main ticket and Development only if it
can access the internal one. The change updates the main ticket within a transaction and adds a
history event with the previous value, new value, author, and date. Saving the same value creates no event
or email. A closed ticket rejects the change.

The creation, list, and detail responses include the priority. The main and internal tickets expose the
same value; no column is added to the internal ticket. The lists accept several priorities as a filter and
an explicit order. The default order is critical, high, medium, and low; within the level, the
oldest first. Filters and order live in the URL. Search, states, assignment,
dispatch, and AI summaries do not change.

The first historical arrival at `critical`, including creation, requests the notice defined in
`docs/flujos.md`. A persistent mark on the main ticket avoids repeating it even if it is later lowered and
raised again. Sending occurs after saving: an SMTP failure is communicated to the actor without reverting
priority or history.

`v1.0.0_dev_en.sql` or `v1.0.0_dev_es.sql` distributes the 25 tickets realistically: mostly medium, some low and high, and
few critical. The tests cover invalid values, default value, permissions, closed ticket,
history, shared value, combined filters, stable order, and a single non-blocking notice.

### 16.1 Implementation and verification

Priority lives only in the main ticket and is projected onto the internal ticket. `PATCH
/api/tickets/{number}/actions/priority` applies the permissions by type, blocks closed tickets, and
records the change in the timeline. The `critical_notified_at` mark is set in the same
transaction as the first arrival at critical; the email is attempted after `COMMIT` and its failure
returns as a notice without undoing data. The full backend suite, the 228 frontend tests, the
build, and the focused E2E journey passed on 2026-10-14.

## 17. Layout of the simultaneous view

The `losdos` form does not alter data or contracts of the module. On PC it changes its container: it takes up all the
available width, uses two equal columns, and limits its height to the remaining space under the detail
header. Each instance of `app-ticket-vista` is an independent scrollable region and preserves
its header, timeline, and editor in the same flow.

Main and internal tickets do not synchronize position or content height. Reloading data preserves the chosen
form, but does not incorporate new persistence for scroll position. The `principal` and
`interno` forms continue with the current behavior.

Below the comparison width a single ticket is presented. Two conversations are not stacked nor
is each reduced to half a screen; the switcher remains the way to go from main to
internal. The tests verify the two regions, the 50/50 width, isolated scroll, and the regression of
individual views. API, permissions, history, and comments do not change.

The implementation keeps the common header outside the scrollable regions and maintains each
editor at the end of its conversation. The 228 frontend tests and the build passed; the
real browser journey verified equal columns, isolated scroll, and a single view at 900 px.

## 18. Full conversation inside each half

The `losdos` form already creates two equal outer regions, but each instance preserves the internal
grid of the individual card. Since `soloLaConversacion` hides the side details without removing
their column, each conversation uses two thirds of its region and leaves the remaining third empty.

With `soloLaConversacion`, the component will use a single column at 100%. Everything that belongs to the
conversation —header, subject, escalation reason, description, comments, history, and editor—
widens together. The space between the two instances remains 24 px and each region preserves
its own scrolling. Previews preserve their 480 × 360 cap and the editor only
grows horizontally.

Without `soloLaConversacion`, the component keeps `2fr 1fr`: the individual view does not change and continues
showing conversation and card. The tests will explicitly distinguish both modes and the real journey
will check width utilization, absence of horizontal overflow, and isolated scroll.
The `losdos` form will be available from 1152 px; at 1151 px or less a single view is kept. No
contracts or module behavior change.

The template applies a full column only with `soloLaConversacion`; without that input it still
uses `2fr 1fr`. The 229 frontend tests, the build, and the real journey at 1800, 1152, and
1151 px confirmed the full width, the absence of horizontal scroll, and the responsive regression.

## 19. Implementation: immediate improvement and local review

`POST /api/tickets/{number}/writing/improve` stops receiving `tone`. It receives `editor`, `text` and,
only when requesting another proposal, `previousSuggestion`. The last field must correspond to the
visible proposal of that editor and is not persisted. `tickets` preserves authorization, limits, access
to the editor, and contextual name resolution before delegating to `ai`.

Pressing **Improve with AI** immediately sends the current draft. The response writes no tables,
does not save the description, and does not publish the comment: it remains as a local read-only proposal
for that editor. Each open editor preserves its own proposal; generating in description does not
discard the comment one or vice versa.

**Use proposal** replaces only the local draft and closes the review. **Keep original** discards and
hides the proposal without touching the draft. There is no separate copy action:
the text is selectable and can be copied with the browser's normal capabilities. Saving or
publishing remains a separate action.

If the draft changes while generating or after receiving the result, the proposal remains
visible and selectable, but **Use proposal** is disabled as outdated. **Generate another
proposal** uses the current draft and delivers only the last output as wording to avoid. A failure
does not create a review or modify the editor; it returns to the initial state and is communicated through a toast.

The tests cover the DTOs without tone, independent proposal per editor, local application, discard,
selectable text, obsolescence when editing, different regeneration, error by toast, and absence of
automatic saving. Permissions, states, history, attachments, and transitions do not change.

The implementation passed the 230 frontend tests and the production build. In a real
browser, a Support account generated a proposal directly under the editor without a modal or selector,
kept it selectable when changing the draft, saw **Use proposal** disabled as outdated,
and discarded the output with **Keep original** without losing the text or publishing the comment. The
affected backend tests also pass; the only failure of the full suite remains the unrelated
priority case described in `modules/ai.md`.

## 20. Implementation: safe regeneration of summaries

Creating a ticket requests **Reason** and **Last action** as before. For the opening, the last action
text includes that the person opened the ticket and described the problem. Editing the subject or
description requests both fields again; comments and transitions preserve the current rules.

Each field distinguishes current text, outdated previous text, generation, and error. When requesting
a new version, the previous one is not immediately deleted: it remains available and is marked outdated
until a valid response is received. If both attempts fail, the field preserves that text and
shows its own error. Without previous text, it presents only the error. **Reason** and **Last action**
each have their **Retry** action, without forcing the other to regenerate.

The periodic summary query updates only its state and content. It does not rebuild the description editor,
the comment editor, the comment being edited, or their assisted writing components. The same rule
applies to manual retry: drafts, local attachments, and AI proposals survive intact. Saving or publishing
continues to be a separate action.

The tests cover creation, editing, independent update of each field, preservation of the
previous one, specific errors, retries, and absence of local loss during polling and manual
actions. Numbering, permissions, history, email, attachments, and ticket transitions do not change.
The editor preserves formatting, pasting, dragging, selection, and comments made up only of files,
but stops explaining them in a permanent paragraph. The valid proposal preserves title, content, and
actions without adding a neutral state that repeats that it is already ready.

The retry endpoint returns the newly persisted `pendiente` state, not the copy of the ticket
read before enqueuing. The screen then polls only the card to extract its `insights` and
merges them in memory; it does not emit the full reload of the detail. The `CS-2026-0027` journey
confirmed that the draft and the inline proposal survived a retry of **Reason**, and that a
description edit again produced both fields with updated content.

## 21. Implementation: linked context for summaries

### 21.1 Event and comment as a unit

When a movement carries a comment, both are saved in the same transaction as today, and in addition
the `ticket_history` entry expressly references the created `ticket_comment`. The
relationship is not inferred by proximity of dates nor is simply "the previous comment" used. Thus a closure can
be delivered to AI as a verifiable unit: who closed it, what state it came from, which one it reached, and what
explanation they wrote.

Movements prior to this relationship remain readable. If they have no persisted link,
the last action uses the event without attributing an old comment to it; the deterministic fallback says
actor and action. The sample data that does not include a closing comment does not invent one.

### 21.2 Coordinated reading of the main and internal tickets

The context builder always starts from the requested number. If it is internal, it publishes its state,
dates, assignee, comments, and history —not the state of the main ticket— and then adds all the
functional data of the main ticket. If it is main, it does the reverse and adds the internal ticket when it exists.
Both blocks are distinguished by number and type.

The last target event is calculated only over the chronology of the requested ticket. The service
delivers separately the chronology of the pair for context. It includes category, tags, priority,
requester, creator, assignees, watchers, attachment names, and dates, in addition to subject,
description, reason, comments, and history. The absence of an internal ticket is normal and generates no error.

### 21.3 Persistence, regeneration, and scope

`ticket_history` incorporates an optional reference to the associated comment. The movement service
receives the created identifier and stores it in the history entry within the same transaction.
The conversation query exposes that relationship to the context builder without changing the public
ticket response.

The regeneration of `ultima_accion` is requested in the background for all tickets whose context
version is earlier than the current one. It is idempotent, resumable, and does not block startup. New
movements continue requesting only the fields that already correspond to them; **Reason**
only changes when creating or editing the original problem, and the internal ticket reason remains immutable.

The service and repository tests cover atomicity, correct link, event without comment,
main ticket without internal ticket, different states between the pair, chronological order, and reconstruction after
restart. Transitions, mandatory comment when closing, permissions, email, numbering,
public DTO, and interface do not change.

The closing test confirms that comment and history share the identifier within the
transaction; the context test confirms that the closed internal ticket takes its own state, actor, and
linked comment. The affected tests pass. The full suite retains the previous and
unrelated priority failure from an internal ticket, unrelated to context or summaries.

## 22. Implementation: stable triggers and English contract for summaries

Creating a main ticket requests `summary` and `lastAction`. Editing its title or description requests
`summary` again; the edit also preserves the movement that updates `lastAction`. No other
action requests Summary. Creating the internal ticket requests its Summary from `escalationReason` and its
Last action; a future edit of `escalationReason` would request both fields again.

The public insights contract changes from `motivo` to `summary`. The internal field changes from
`reason` to `escalationReason` in DTOs, services, and frontend; the existing SQL column
`escalation_reason` keeps its English name. Previous aliases are neither accepted nor emitted because
1.0.0 is not yet published. The endpoint that retried an insight is removed: tickets only requests
generation as a consequence of a valid creation or update.

Reading distinguishes current result, generation, and error. In generation or error without a current
result it delivers empty text so that the interface shows `—`; it does not present the previous result
as current. The tests cover the trigger matrix, both ticket types, new JSON,
rejection of previous names, disappearance of the manual route, and continuity of local drafts
during automatic polling.

## 23. Visible contract of the summary table

The four queues share the same broad order:

`Ticket · Priority · State · Subject · Classification · Summary · Last action · Requested by · Assignee · Update date · Action`.

The data set and its meaning do not change. Ticket is not skipped or truncated; Subject, Summary, and
Last action have at most two lines and allow consulting the full text; Classification
limits its chips to two lines and summarizes the rest as `+N`; people allow two lines and use
`—` if missing; Update date separates date and time. Priority and State are centered; the Action header
is left-aligned and its button is centered. The three columns adjust their width to the content.

The full truncated text and the tags gathered in `+N` are accessible with cursor and keyboard.
`+N` lists the hidden tags also in its accessible name. Switching to cards depends on the
real width of the container, including the space left by the side menu. No new
ways of sorting Priority or Update date are incorporated.

Desktop queues keep a last column to open the ticket. Its visible contract
is **Action** in the header and **Open** in each row. The header is aligned to the left
and the content is centered; the column adjusts to the minimum width of the button instead of reserving a fixed
measure.

The action is not moved to the Ticket column nor eliminated: the number and the button remain two
explicit links to the same detail. Nor is the entire row turned into a link. In mobile
cards, the button remains at the end as today.

At intermediate width only Requested by and Assignee appear; the other nine columns remain. From
768 px the table preserves its measurements and admits native horizontal scrolling. Only below
that threshold does it change to cards, which show the same data in the same conceptual order and preserve
Open at the end.

The change is exclusively presentational. It does not modify endpoint, `routerLink`, authorization,
query, pagination, filters, order, Summary or Last action data, or persistent model. The
component tests verify order, responsive visibility, limits, missing values,
header in both languages, final cell, and mobile variant; the browser check covers wide PC,
intermediate width, and mobile and confirms that the button is perceived as part of the table.

## 24. Row identity during horizontal scrolling

The **Ticket** column is the fixed reference of each row in the wide and intermediate tables. Its
header and its cells remain at the leading edge when the other columns are scrolled horizontally.
They preserve the width, the unbroken text, and the existing link.

The table exists from 768 px and, if its columns exceed the available width, uses the container's native
horizontal scrolling without a gradient or additional notice. Only below that does it use cards.

The Ticket cells of the body are fixed only horizontally; their header is fixed on both axes according to
section 26. It does not include another column and does not change the card variant. The opaque surface and the
right separator come from the theme to prevent scrolled content from being seen behind the
number. API, authorization, navigation, data, and order do not change.

The interface test verifies the behavior in the eight themes and in the two table widths,
including focus and link activation.

## 25. Incremental consumption of lists

The paginated endpoint and its `page` and `perPage` parameters do not change. The screen requests pages of
20 items in sequence and accumulates them to present them as a single list. It only requests the
next one when it exists according to `total`, no other request is in progress, and the marker near the end
enters the visible area. The response is appended respecting the backend order and avoiding repeated ticket
numbers.

A change of filter, search, view, type, or order creates a new query: it discards the accumulation,
returns to `page=1`, and takes the scroll to the top. An error in a later page preserves the previous
pages and allows retrying exactly the failed one. The screen presents no pagination controls or text
and does not request pages after reaching the total.

Navigation memory lives only in the frontend and only during the current page session. For
each queue it preserves query, accumulated tickets, next page, and vertical position to
restore them when returning from the detail. It does not use persistent storage and is lost on
reload or closing the tab. It does not change the backend contract, permissions, or JSON response.

The table variant and the card variant share this logic. The indicator, the retryable error, and
the announcement of how many items were added are presentation; they generate no notifications or ticket
events.

## 26. Incremental loading inside the scrollable table

In the table variants, the list container becomes the visible root of incremental
loading. The marker for the next block is evaluated against its internal vertical scroll and
keeps the existing anticipation. Reaching it requests exactly the next page; several
crossings do not create parallel requests and the total continues to mark the silent end.

The table preserves the existing paginated response, blocks of 20, and deduplication by number. The
new available height, fixed headers, separator, and tooltips are presentation and do not
alter endpoint, parameters, order, filters, authorization, or JSON structure. In cards, the
marker continues to be observed against the normal page scroll.

When returning from a detail, the accumulated blocks and the position of the corresponding container are restored:
internal scroll in the table and page scroll in cards. Memory remains ephemeral and
is discarded on reload or closing the tab, as in section 25.

The tests verify the observation root in both variants, anticipated loading, exclusion of
concurrency, restoration of the correct scroll, and error/retry continuity. No backend contract
changes.

## 27. Implementation: return from Waiting to In progress

### 27.1 Reproduced finding

The comprehensive journey of 2026-10-16 left a main ticket in `en espera` and then entered as
Support. The state selector only offered escalate, resolve, or close, although the module contract
admits returning to `en progreso`. The cause is in the visible frontend matrix: it offers the
action of starting for Development from a new internal ticket and for Support from an internal ticket on hold, but
omits the same return for Support on the main ticket.

### 27.2 Implemented behavior

When a main ticket is in `en espera`, Support will see **In progress** and can perform that
transition. It preserves the confirmation and the same endpoint, history, email, and regeneration of Last
action that already correspond to the movement. The User continues not changing states;
Development continues not operating the main ticket; the Administrator remains read-only.

No states or routes are added, and the contextual translation of `en espera` does not change: the User continues
seeing **Waiting for you** and the teams **Waiting**. The tests cover the
action matrix of main and internal tickets, the absence of unauthorized options, and the real journey
Waiting → In progress with Support.

### 27.3 Requester response as an automatic transition

The matrix in section 3.1 already promises that the User returns the main ticket from `en espera` to
`en progreso` upon replying, but the built flow only saves the comment. That discrepancy is
corrected in the service, not through a new interface action.

When a new comment is created on the main ticket, the service locks and rereads the
main ticket within the same transaction. If it is still `en espera` and the author is exactly its
requester (`requester_id`), it saves the comment, changes the state to `en progreso`, and adds the
history entry attributed to that person. The three operations are confirmed or reverted together.
The lock guarantees that two simultaneous replies produce two comments, but a single
transition and a single history entry.

Any new comment from the requester counts as a reply, including the body that is initially
empty and will later receive attachments. Editing or deleting comments, uploading an attachment afterwards,
commenting as a watcher, Support, Development, or Administrator, and commenting on the internal ticket do not cause
the transition. Nor is a main ticket changed that has already stopped being `en espera`.

No additional email is sent for the automatic change: only the notices that
already correspond to the comment are preserved. The HTTP contract remains the same and returns the comment; the
screen rereads the detail through its normal flow to show the current state. The transition requests
**Last action** again only once and does not change permissions, states, or routes.

The tests against PostgreSQL cover the correct requester and the history actor, another
identity, the comment without a body, the internal thread, and later editing. The transaction includes
comment, update, and history, and the reread uses `SELECT ... FOR UPDATE`: therefore a failure
reverts the set and two concurrent replies can observe `en espera` only once. The
flow does not incorporate any additional email call.

#### 27.3.1 Decision record

The owner decided on 2026-10-20: a single transaction (1A), any new comment even if
empty (2A), history attributed to the requester (3A), no additional email (4A), exact match
with `requester_id` (5A), no transition when editing (6A), locking and concurrent rereading
(7A), and HTTP contract unchanged (8A). No open decisions remain.

## 28. Operational data for reporting

`tickets` preserves authority over states, priorities, assignees, categories, and history.
It exposes to the Dashboard module a read query with already authorized scope; Dashboard does not import
ticket repositories or reconstruct permissions. Existing events determine creation,
first human response, first resolution, reopening, and closing. Automatic events do not count
as a response. Main and internal tickets keep separate metrics.

Pending age is calculated in open intervals: it is born with the ticket, stops upon
resolution, and runs again if reopened. The backend returns UTC instants and applies the global zone
when delimiting periods and exporting. No state, transition, email, or AI summary is modified.

The read repository lives in `modules/tickets/repositories`, reconstructs the state at the cutoff, and
delivers typed facts to the Dashboard port. A test against PostgreSQL runs the real query
and validates types, numbers, and non-negative age.

## 29. Strict membership of "My tickets"

`/tickets` is a personal queue calculated with the current state. It contains only:

- tickets assigned to the signed-in account;
- tickets where that account currently appears as a watcher.

They do not enter because they were opened, requested, or commented on by that person, nor because they lack an
assignee. When a person is removed from watchers, the ticket disappears from **My tickets** on the
next update, unless it is still assigned to them. The mention and history remain, but
do not grant membership again.

The operational role sets the level: Support only gathers main tickets and Development only internal ones. The
Admin capability does not expand or transform this queue; Support + Admin continues as Support and
Development + Admin as Development. The other authorized tickets are consulted in **Main
tickets** and **Internal tickets**, which preserve their current scope.

The API keeps `view=assigned` and `view=watching`; the empty view combines both sets, without
duplicating a ticket that meets both conditions. Opening and participation through comments are removed from the semantics of `mine`. It does not change who can read a card found from the
general lists, assign, self-assign, or manage watchers.

### 29.1 Decision record

The owner decided on 2026-10-19: exclude unassigned (1B); exclude opening and comments
(3A and 9A); strictly separate main tickets for Support and internal tickets for Development (6A); open on
the combination of assigned and watched (7A); and preserve the operational role when there is
Admin capability (5A). He also expressly corrected that stopping watching removes the ticket from the
queue if it is not assigned. No open decisions remain.

## 30. Private comments on the main ticket

The main ticket admits comments with technical visibility `public` or `private`; the visible names
are **Public / Private** and **Público / Privado**. Every new comment starts as `public`. Support
can publish both types on any main ticket they can open. Development obtains a precise exception:
it can write `private` comments on the main tickets it accesses, but never a `public`
reply to the User. Comments that accompany state movements of the main ticket
always remain public.

A private comment is read by Support and Development with access to the main ticket, and the pure Administrator
can audit it read-only. The User does not receive
the row, a marker, the content, the mentions, or their attachments; the backend applies the exclusion in
reading, search, and download. Directly requesting a private attachment as a User responds 404.
The attachment dynamically inherits the visibility of its comment, so changing it also changes
its access without duplicating files.

Only the author, while retaining permission to comment, can edit the body or toggle
visibility. The change applies immediately in both directions. The history adds an entry with
actor, previous visibility, new visibility, and instant, without copying the content. Changing from public to private
cannot withdraw an email already delivered or erase the audited fact, but it presents no additional confirmation
or warning. Logical deletion preserves the current rules.

Mentions work the same: they validate active technicians or developers, add watchers, and
send their internal notice. In addition, a new private note notifies by email the assignee and the
watchers of the main ticket, without duplicates and without sending it to the author. The email contains
number, subject, author, and link; never body or attachments. The requester is not a recipient.

The User's search excludes private text; the authorized search of Support and Development does
consider it. If the main ticket is escalated, its private notes remain only on the main ticket:
they are not copied or projected into the internal conversation. Development can consult them by switching to
the main ticket view.

The model adds a mandatory visibility with `public` as the initial value for existing rows.
The creation API accepts the authorized visibility and editing allows changing body and/or
visibility; responses are filtered before building comments, attachments, watchers,
AI context, and search results. The tests must cover authorization by role, change in
both directions, history, search, download by URL, recipients without authorization, movements always
public, and total absence for the User.

### 30.1 Decision record

The owner decided on 2026-10-19: writing and reading for both teams (1A–2A), selector
**Public / Private** (3A corrected), permanence only on the main ticket when escalating (4C), attachments with
inherited privacy (5A), initial `public` (6A), bidirectional change by the author (7B and 11A), two
AI summaries (8A and 24A), internal notices (9A, 21A–23A), current mentions (10A), chip and background
(12A), auditing without content (13A), total invisibility for the User (14A), change without notice
(15C), Development private only (16A), Support with both types (17A), public movements (18A),
audience-based search (19A), others' attachments as 404 (20A), immediate AI regeneration (25A), and
administrative reading without editing (26A). No open decisions remain.

## 31. Unified general ticket list

Support, Development, and combined accounts that have one of those roles have a single
**Tickets** module. Inside it, **Main** or **Internal** is chosen; Main is always the initial view
when entering from the menu. **My tickets** remains separate and preserves its personal scope.

The canonical address is `/tickets?type=main|internal`. `/tickets` and any unknown `type`
are normalized to `?type=main`. Changing the option updates the URL and Back/Forward restores the view.
The old routes of the two lists are removed without redirects: no compatibility code is kept
and they respond 404.

Each view separately maintains search, filters, order, and loaded data while remaining on
the screen. The selector does not show quantities. Development can consult main tickets and comment
on them, but every comment of theirs is mandatorily born `private`; they receive no visibility selector
nor can they publish. Support and Development can read those private notes, in addition to the
combined accounts with those roles; the User never receives them.

The list contracts share screen and state, but the permissions and queries of
main and internal tickets remain explicit in the backend. The tests verify navigation,
combined permissions, and forced private visibility for Development.

## 32. Queues with useful width and legible ends

### 32.1 Finding

The four queues currently limit their content to 88 rem and the shell adds 40 px on each side on
desktop. On a wide screen both spaces accumulate: the table becomes narrower than
necessary and its final column may remain outside the capture even though physical width exists. This
does not match the full presentation the product needs.

### 32.2 Desktop contract

**My tickets** and the general, main, and internal views use all the width given to them by the
shell. The side menu preserves its current state and behavior; collapsing it is not a requirement
to obtain the new width. In these queues, the shell leaves 16 px on the left and right on
desktop and the 88 rem maximum is removed. The reading limits of forms and
other pages are not altered.

The table distributes the additional space among Subject, Classification, Summary, and Last action without
reducing typography or row height. From 1440 px, with the menu expanded and normal data,
it seeks to present the eleven columns without horizontal scrolling. Native scrolling remains
as a fallback when content, translation, zoom, or a smaller width does not allow it.

Ticket remains fixed at the leading edge and Action becomes fixed at the trailing edge during
horizontal scrolling. Both use opaque surfaces, theme separators, and stacking order
that prevents text from being seen behind. The first cell incorporates additional inner space with respect to the
left edge of the table and the last with respect to the right edge; that air is not obtained by creating
new outer margins or narrowing the central columns.

Below 768 px the current cards are maintained. Columns, data, truncation,
order, filters, permissions, routes, incremental loading, and actions do not change.

### 32.3 Verification

The tests cover the four queues, full width, 16 px margin, fixed ends, and outer padding
of Ticket and Action. The browser journey checks desktop with the menu open and
collapsed, scrolling at an intermediate width, zoom, visible focus, eight themes, and mobile cards.
It also confirms that Summary and Last action preserve two legible lines and that Action does not cover the
previous column.

### 32.4 Decision record

The owner corrected the initial premise on 2026-10-22: the problem was not collapsing the menu but the
sum of the content maximum and the side margins. He chose to apply the change to the four
queues (1A), 16 px sides (2A), fix Action to the right (3A), and keep the menu. He added
inner padding next to the left and right edges. No open decisions remain.

### 32.5 Implementation and verification

The queue compensates 24 px of the shell's desktop padding and leaves 16 px effective per side;
it removes the 88 rem maximum. The table uses minimums of 72 rem with people and 64 rem without them. Ticket
incorporates 16 px at the leading edge; Action measures 6 rem, incorporates 16 px at the trailing edge, and remains fixed to
the right with an opaque separator.

The 252 frontend tests, the boundaries, and the production build passed. The subsequent focused
regression of the final width passed its 10 tests. Chromium confirmed, with the menu open, 1680 px of
table within 1920 px and 1200 px within 1440 px, without horizontal scrolling; Action measured 96 px and
was seen complete. At 390 px the 20 visible cards were preserved. The visual review also checked
padding, borders, truncation, and focus at the ends.

## 33. Consistency when changing the type of the general list

> **Proposal of 2026-10-23.** The ephemeral journey verified that the switcher can mark
> Internal while the table keeps Main. The owner chose 2A and closed the review with
> 3A and 5A. No open decisions remain; explicit approval is pending.

In `/tickets`, **Main** and **Internal** form a single selection: URL, marked button, query,
and visible rows must always refer to the same type. When changing the option, any
previous request is invalidated, the table is emptied, and loading is shown until the new type is received. A late
response of the previous type cannot repopulate it.

Navigation separately preserves remembered filters and results, but only restores them after
checking that they belong to the current type. API, permissions, filters, order, incremental
loading, and detail routes do not change.

The regression deliberately delays a response, alternates both types, and checks URL, accessible
state of the switcher, and visible numbers after each change.

## 34. Manual assignment and operational role in the card

> **Approved, implemented, and verified on 2026-10-30.** The real journey found that Support cannot
> reassign an internal ticket from its card and that a Support or Development account with
> Admin capability becomes read-only on tickets. Section 3.3 and section 14 of
> `usuarios-y-permisos.md` already establish the opposite; the failure was in the interface. API,
> data model, and backend permissions do not change.

### 34.1 Finding

The card decides who can reassign with two rules that do not match the backend:

1. On an **internal ticket**, `puedeAsignar()` only shows it to Development, while the
   `Assign` service and decision 50 allow reassignment by **Support and Development**. Since
   internal tickets are born without dispatch by default, Support sees "No assignee" and has no way to set
   one.
2. `soloLectura` is calculated with **administrative capability**, not with the exclusive role. A
   `soporte` or `desarrollo` account with `isAdmin` true loses the reassign button —and also
   editing and closing—, although section 14 establishes that **in tickets the operational role
   prevails**. Only the pure Administrator looks without buttons.

The same rule is repeated in the state control (`acciones-del-ticket.ts`), where closing is
hidden with the same signal.

And, in addition, the ticket screen returned to the **default view on each reload**, not only when
opening the ticket. After reassigning an internal ticket, Support jumped to the main ticket and the
"Ticket assigned" notice disappeared. The view is chosen when opening the ticket and is preserved across reloads.

### 34.2 Contract

- **Manual assignment does not depend on dispatch.** It does not matter whether the ticket was born by shifts,
  without dispatch, or already assigned: whoever can reassign does so from the card, with the button next to the
  assignee and its selection window.
- **On the main ticket, Support reassigns**; **on the internal ticket, Support and Development reassign**, to
  any active person of the corresponding role. It is decision 50 and the backend matrix.
- **In tickets the operational role prevails.** Support + Administrator works as Support and
  Development + Administrator as Development. Administrative capability does not turn the card into
  read-only; the **pure Administrator** continues looking without buttons (section 14 of
  `usuarios-y-permisos.md`).
- **A closed ticket still needs to be reopened before reassigning.** It is the current behavior
  of the card and this change does not alter it: first it is reopened, then assigned. The backend
  does not change.
- **The chosen view is preserved on reload.** Assigning, commenting, or moving a ticket no longer returns
  to the default view: whoever was on the internal ticket remains on the internal ticket and sees their notice.
- Endpoints, server permissions, shift dispatch, email notices,
  and the candidate list do not change.

### 34.3 Verification

- View test: Support opens an internal ticket without an assignee, sees the button, chooses a
  developer, and the request goes to `POST /api/tickets/INT-…/assign`.
- View test: Support with Admin capability preserves the button and the actions of their role;
  the pure Administrator still lacks them.
- Service test: the `puedeAsignar` matrix covers main and internal tickets, the four roles, and
  cumulative Admin capability.
- Interface journey: Support reassigns an already escalated internal ticket on PC and mobile, and the ticket changes
  queue without losing the view or the notice.
- The 260 frontend tests, the boundaries, the build, the ticket service tests,
  and the interface journey on PC and mobile passed. The ticket repository suite retains the
  known environmental reporting failure when there are no seeded tickets; it does not touch this change.

## 35. Creation for a requester without an account

> **Approved, implemented, and verified on 2026-10-30.** The product already says that Support can open a ticket "for whoever does not
> have access to the platform" ([propósito-y-alcance.md](../propósito-y-alcance.md)), but creation
> required an active account and rejected an unknown email. This section makes that promise real without
> changing the ticket model or the permissions of the other roles.

### 35.1 Finding

`POST /api/tickets` accepts `requesterEmail`, but `solicitante()` looks for an active account by email
and, if it does not find it, responds `tickets.requester.unknown`. Support cannot create the ticket in the
same step: they must first go to Users, create the account, and return. With AD the path exists, but
also outside ticket creation.

The account needed is a normal account, with the `usuario` role; it is not a requester without an account
nor a ticket without `requester_id`. The ticket model does not change.

### 35.2 Contract

- **Only Support** can send `requesterEmail`; the current rule does not change. Creation adds
  `requesterName` and `requesterLastName`, which are only used when the account does not exist.
- **The account exists and is active**: it is used as is. The written names are ignored.
- **The account does not exist and the method is local**: a `usuario` account is created, origin `local`, active,
  **without a password and without a registration email**. It does not receive access to the platform; it serves to identify
  the requester and receive the ticket notices. Access can be given later from Users.
  `requesterName` and `requesterLastName` are mandatory.
- **The account does not exist and the method is AD**: the person is searched in the directory by their email.
  If they are there, they are created or linked with the directory data, origin `ad`, without a local password and with
  the `usuario` role; the form names do not replace those of the directory. If they are not there, it is
  rejected with `tickets.requester.notInDirectory` and **neither the ticket nor the account is created**.
- **The account does not exist and the method is Keycloak**: there is no people search in the realm, so
  it is rejected with `tickets.requester.accountRequired` and it is indicated that the person must exist in
  Keycloak or sign in once. A local account that could not enter through that method is not created.
- **The account exists but is deactivated**: it is rejected with `tickets.requester.inactive`; it is not
  reactivated from ticket creation.
- **The account and the ticket go together**: if the account cannot be created or linked, the ticket is not numbered nor
  saved. If the ticket fails afterwards, the account created in that attempt is preserved and can
  be reused; no half-done ticket is left.
- **No registration email is sent** to the locally created requester. The ticket notices remain
  as always: the requester receives the emails of their ticket and can reply when
  they have access.
- **If the account is created from AD**, in addition to the ticket, the new email
  `auth.accountCreated` is sent to the person —"Your account has been created in the help desk; sign in with your organization
  account"—. It carries no password link or token. If the email fails, the ticket and the account
  remain created and the failure is recorded in the log.
- The creation response carries `requesterCreated` and `requesterNotified` so that the screen can
  say what happened: account created for the ticket (without access) or account created from AD with its
  email sent.
- Shift dispatch, assignment, notices, transitions, and the internal ticket card do not change. The internal ticket continues to show the main ticket requester and the developer as
  assignee.

### 35.3 Interface

The Support creation form keeps "Who is the ticket for?" and adds **First name** and
**Last name** below the email, with a help text saying what will happen:

- if the account exists, the existing one is used;
- if it does not exist and the installation signs in with local accounts, it is created only for the ticket and **does not receive
  access**;
- if the installation signs in with AD, the person must be in the directory;
- if it signs in with Keycloak, the account must already exist in the realm.

The name fields are not required on the screen so as not to prevent the case of an account that already
exists; the backend requires them only when it has to create the local account. The error returns with its
key and is read in the form notice.

When the account is created, the ticket says so when opened: "Account created for the ticket" if it was
local, or "Account created from the directory; the notice has been sent" if it was AD and the email went out.
If the email fails, the notice says so without undoing the creation. The notice travels in the module service
and is consumed only once; it is not passed through the URL.

### 35.4 Verification

- Service: local account created without a password, without a registration email, and with the `usuario` role; the ticket
  remains with that account as requester and with Support as creator.
- Service: existing active account is reused; inactive account is rejected; AD found creates or
  links with the directory data; AD not found and Keycloak not found reject without
  creating a ticket.
- Service: if first name or last name is missing in local creation, it is rejected before numbering the ticket.
- Interface: Support creates a ticket for a new email and the card shows that person as
  requester and Support in "Opened by".
- Interface: with AD, the journey uses a person from the directory; with Keycloak, the attempt with an
  unknown email shows the error key and creates nothing.
- Email: the account created from AD receives `auth.accountCreated` once, without a password
  link; an existing or linked account does not receive it.
- Interface: after creating a local or AD account, creation navigates to the ticket and the notice appears there,
  once.
- The 260 frontend tests, the boundaries, the build, the ticket, user, and mail service tests,
  and the interface journey on PC and mobile passed. The `v1.0.1` migration
  left the 28 template rows in the isolated database.
