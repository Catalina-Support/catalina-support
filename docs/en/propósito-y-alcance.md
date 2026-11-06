# Purpose and Scope

[Español](../es/propósito-y-alcance.md) · **English**

> **State:** as-built
> **Last updated:** 2026-10-19
>
> **Correction approved, implemented and verified on 2026-10-19.** The Dashboard is exclusively
> statistical: it does not list or link individual tickets. Detail remains in the existing
> inboxes and detail views. The individual-detail interface and API were removed.
>
> **Implemented and verified on 2026-10-18.** Section 12 adds the Dashboard and operational
> reporting without turning the product into an evaluation or SLA tool.
>
> **Implemented on 2026-10-18.** Section 11 allows adding Administrator to Support or Development
> without mixing their operational responsibilities.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** The owner adds ticket priorities `low`, `medium`,
> `high` and `critical`. This decision expressly corrects the previous exclusion in section 6.
> Section 10 records the scope agreed in decisions 1A–25A, with the owner's corrections to 3A, 13A
> and 18A. The owner explicitly approved it.
>
> **Becomes as-built on 2026-09-25**: what this document defines —the two teams, the ticket model,
> numbering and statuses— **is implemented** in the `tickets` module, backend and screens, and
> verified end to end. The state changes, not the content: what it says is still what the code does.
>
> **Amended on 2026-09-22**, when writing `docs/flujos.md`: **rule 5** changes
> —when Development resolves the internal ticket, the main ticket **returns to Support** instead of
> remaining resolved— and **rule 8** is added, the return to Support, along with the **seventh
> notice** (the ticket returns to the Support inbox), which prevents a case resolved by Development
> from being left unexplained to the user.
>
> **Amended on 2026-09-22**, when reviewing it after writing it: the **sixth email notice**, the
> repeated notice on re-escalation, the **25 MB** limit, the edited and deleted marks, and the
> `admin` account with which an installation is born.
>
> **Amended on 2026-09-22**, when detailing access: the fourth role `administrador`, user creation
> from the application, tickets created on behalf of someone else and the three ways to sign in.
>
> Approved by the owner on 2026-09-22, **amended three times that same day** —the three entries
> above— after reviewing it in the form of questions.
>
> It is the first document in the chain and **unblocks** `usuarios-y-permisos.md`,
> `modules/tickets.md` and `flujos.md`. **It does not enable writing code**: each area needs its
> approved document first (Rule 0 of `AGENTS.md`).

## 1. What problem it solves

Today a need arrives through whatever channel (WhatsApp, email, a call, a passing comment) and does
not end up anywhere. Two consequences:

1. **The trail is lost**: there is no single place to see what was requested, what was answered and
   how it ended.
2. **Development receives noise**: whoever writes code learns about problems without context and
   without knowing whether something has already been tried.

Catalina-Support exists for that and only for that: **that every need has a single place, an owner
and a history**, and that only what the first level cannot resolve reaches the second level.

## 2. For whom

Tickets are created by **the users who need attention**: someone requesting technical support or
requesting software development. Four roles, and no more:

| Role | What it does | Contact with the user |
| --- | --- | --- |
| **User** | Creates the ticket and replies when Support asks. Sees **only their tickets** and can close and reopen them. | The user themselves |
| **Technical Support** (level 1) | First contact. Tries to resolve the need and escalates what it cannot. Also creates users and **can create tickets on behalf of another user**, for those who have no access to the platform. | **Yes**: it is the visible face of the system |
| **Development** (level 2, escalation) | Acts when Support cannot resolve the need or when the user requests software development. Resolves and returns; **does not close tickets**. | **No**: it never talks to the user |
| **Administrator** | Does not handle tickets: manages accounts (user creation and deactivation). | — |

From the third column comes a rule that runs through the whole design:

> **Development has no contact with the user.** Everything the user needs to know goes through
> Technical Support; everything Development needs to ask also goes through Technical Support.

Development **does see who the user is** (name and email) to have context when resolving; what it
does not do is write to them. "No contact" means no dialogue, not blindly.

**Who creates a ticket is not always who needs it.** When Support creates it on behalf of a user,
the ticket stores the two things separately: the **requester** (the affected user) and **who created
it**. The requester is who receives the notices and who can close it.

## 3. The two levels, in one sentence

> **Support resolves everything that does not require touching code; anything that requires a code
> change is escalated to Development.**

From that come two rules that will be applied in the ticket and flow documents:

- Escalating is **not** "passing the buck": the ticket reaches level 2 with the reason for the
  escalation, what has already been tried and the minimum context so as not to start from zero.
- Nothing is lost between levels: the complete history is preserved and can be read in full.

## 4. The ticket model

This is the part with the most consequences, and that is why it is written here even though its
detail (transition table, endpoints and screens) belongs to `modules/tickets.md` and `flujos.md`.

### Two types of ticket, one single number

| Type | Numbering | Who sees it |
| --- | --- | --- |
| **Main ticket** | `XXX-YYYY-NNNN` | The user, Support and Development |
| **Internal ticket** (the escalation one) | `INT-XXX-YYYY-NNNN` — **the same number** as the main ticket, with the `INT-` prefix in front | Support and Development. **The user does not see it** |

- `XXX` is a **configurable prefix**, `YYYY` the year and `NNNN` a **per-year sequential number**.
- The internal ticket **is born when Support escalates** the main ticket. It is the face Development
  sees: the user does not have to find out it exists.
- The two share a number on purpose: looking at `INT-ACME-2026-0042`, anyone knows that its main
  ticket is `ACME-2026-0042`, without searching for anything.
- **There is a single internal ticket per main ticket**: it is created on the first escalation and,
  on subsequent escalations, it is reused (returns to `en progreso`). There is no `INT-…-1` or
  `INT-…-2`: the history of all escalations lives in the same ticket.
- The `XXX` prefix is **one per installation** and is configured by whoever administers it.
- Each ticket stores **requester** and **creator**, who may be different people: Support can open a
  ticket on behalf of a user without access. The requester is who receives the emails.

### States

| Type | States |
| --- | --- |
| Main ticket | `nuevo`, `en progreso`, `en espera`, `escalado`, `resuelto`, `cerrado` |
| Internal ticket | `nuevo`, `en progreso`, `en espera`, `resuelto`, `cerrado` |

Six statuses for the main ticket and five for the internal one, and **not one more**: each extra
status has to be explained, maintained on the screens and assigned to someone who can set it.

`en espera` means waiting for someone outside, and that someone is not the same in each type:

| Ticket | `en espera` = waiting for |
| --- | --- |
| Main | The **user**: Support asked them for information or a test |
| Internal | **Support**, which is Development's only interlocutor |

### How the two move at the same time

The rules that force the two tickets not to contradict each other:

1. When **escalating** a main ticket, the internal ticket is created (or the existing one is
   reopened) and the main ticket moves to `escalado`.
2. If the **internal ticket moves to `en espera`**, the main ticket moves to `en progreso`: what is
   being waited for is Development's business, not an open problem for the user.
3. If the **main ticket is escalated again**, the internal ticket returns to `en progreso`.
4. **Support can reopen the internal ticket** by escalating again: the same internal ticket, not a
   new one.
5. When the internal ticket is resolved, **the main ticket returns to `en progreso`**. Development has
   finished, but whoever closes the case with the user is Support: it writes in the main ticket what
   was done and resolves it.
6. If the **main ticket moves to `en espera`** (Support asked the user for information), the
   **internal ticket is not touched**: it stays where it is. What is being waited for belongs to the
   user, and Development has no reason to stop because of that.
7. **The user and Support close and reopen.** Development does not close tickets: it resolves them
   and returns them.
8. **Closing the internal ticket without resolving it returns the main ticket to `en progreso`**:
   it is the return to Support, for when Development sees that it is not a code change.

Reading rule: **the user never sees the internal ticket.** What they see is their main ticket, with
the notices Support decides to write to them.

## 5. What is included in 1.0.0

- **Creating a ticket** with subject, description and **attachments** (screenshots included). Only
  Support can create it **on behalf of another user**.
- **Automatic numbering** `XXX-YYYY-NNNN` with a configurable prefix and a per-year sequential
  number.
- **Three inboxes**: my tickets (user), Support inbox and Development inbox (which sees the internal
  tickets).
- **Conversation with history** in both ticket types, with author, date and **edited and deleted
  marks**.
- **Assignment**: a ticket has an owner (a person, not just a team).
- **Escalation** from main to internal with a reason, reopening of the internal ticket and return to
  Support.
- **The two lifecycles** of section 4, with their synchronization rules.
- **Email notices** for the seven events in the table below, and for no others.
- **The texts of the ten emails are edited from the application** (the seven ticket notices and the
  three account ones), with preview and a test to the email of whoever edits them: emails are the
  voice of the institution and their text is changed by whoever is in charge.
- **Automatic distribution of new tickets** among Support technicians, **by shift**, and the ability
  to reassign them afterwards.
- **Basic search and filters**: by status, by text in subject or description and by number.
- **Three ways to sign in**: email and password (with the account creation done by Administrator or
  Support), **Active Directory** (account created on first access) and **Keycloak** (same as AD).
  **Email** is the unique identifier of a person, and it is what matches the local account with the
  directory one.
- **Passwords, only with email and password**: when someone is created, an email is sent so they can
  set theirs, and whoever forgets it resets it through the same path. With Active Directory or
  Keycloak **there is no local password**: the domain manages it and the application does not
  intervene.
- **User management**: creation by Administrator or Support. Roles are not edited from the
  application.
- **Authentication and permissions** for the **four** roles, checked in the backend.
- **A new installation is born with the `admin` account**, which is the one that gets the system
  running: create Support and Development. Its password comes from configuration, so there is no
  need to create the first account by hand. Detail is in `docs/usuarios-y-permisos.md`.

The email notices are these, and no others:

| Event | To whom |
| --- | --- |
| The ticket is created | **The technician whose turn it is** |
| Support escalates to Development | To Development |
| Support asks the user for information (the main ticket moves to `en espera`) | To the user |
| **Development needs something from Support** (the internal ticket moves to `en espera`) | **To Technical Support** |
| **The ticket returns to Support** (Development resolved the internal ticket, or closed it without resolving) | **To Technical Support** |
| The ticket is resolved | To the user |
| The ticket is closed | To the user |

The escalation notice is sent **also when a ticket that already had an internal ticket is
re-escalated**: the work has appeared again, and that is what the notice means.

The **fourth** and the **seventh** notify inward, and they are the ones that prevent a ticket from
sitting idle waiting for someone who does not know it is their turn: the fourth when Development
needs something from Support, and the seventh when the ticket returns to the Support inbox —because
Development resolved it and it has to be explained to the user, or because it returned it without
resolving—.

All of the above is the minimum: if something fails, 1.0.0 is still useful; if creation,
conversation, escalation or statuses fail, there is no product.

## 6. What does NOT go into 1.0.0

- **SLA, deadlines, metrics and reports.** No response times or dashboards.
- **Knowledge base** and help articles.
- **Creating tickets from incoming email.** Email goes out, not in.
- **Live chat** and anything real-time beyond reloading the screen.
- **Custom fields** and configurable forms: the fields are the ones decided by the ticket document,
  period.
- **Integrations** (Slack, Teams, GitHub, conversation threads elsewhere).
- **Multi-tenant**, subdomain per client and isolation per product.
- **Mobile application.** The web must work on mobile, yes.
- **Attachments above 25 MB** or external storage: files go to disk on the server.
- **Custom roles and permission editing**: the four roles are fixed, no new roles are created and
  permissions are not changed from the application. User management is limited to account creation.
- **Custom fields** other than the category and priority closed for 1.0.0.

## 7. The criterion for deciding doubtful cases: «no noise»

When in the following documents a choice has to be made and this document does not resolve it, this
rule applies, which is the owner's rule and is worth more than any aesthetic preference:

- **Few statuses and few required fields.** If a field can be filled in wrongly or inferred, it is
  not required.
- **No automatisms that do not save real work**: no reminders, no calculated fields that nobody
  consults.
- **No step that exists only because a large product has it.** If it cannot be explained in one
  sentence why it is needed, it is not done.

## 8. How we will know that 1.0.0 is finished

When a real need travels the complete path without leaving the application: a user creates the
ticket, Support handles it, escalates it if needed, Development resolves it, Support closes it and
the history can be read in full months later. That, and not a task list, is the criterion for
deploying to production (postponed until then, see `docs/arquitectura.md`).

## 9. What this document enables

This document, already approved, fixes **what problem is solved, for whom, how tickets are numbered
and moved, and what is left out**, and unblocks the three that depend on it:

1. `usuarios-y-permisos.md` — the four roles, the two teams and what each can do.
2. `modules/tickets.md` — data model, transition table for the two lifecycles and the closed list of
   modules.
3. `flujos.md` — creation, triage, escalation, resolution, closure and reopening, with the email
   notices.

None of the three enables writing code on its own: each area needs its approved document before
anything is implemented (Rule 0 of `AGENTS.md`).

## 10. Operational priority for 1.0.0

Every main ticket will have a mandatory priority: low, medium, high or critical. The User proposes
it on creation and the default value is medium. Support can correct it; Development can also do so
when there is an internal ticket it has access to. Priority belongs to the main ticket and the
internal one always reads the same value. Support and Development can change it as many times as
necessary while the ticket is open; on closing it, it is frozen.

Priority serves to identify, filter and sort work. It does not change statuses, assignment,
permissions, shift distribution or response times. The only new automation is the email notice the
first time a ticket reaches critical, including its creation: it is sent to all active Support and
never blocks the operation if it fails.

## 11. Administering without leaving the team

1.0.0 allows a Support or Development person to also be Administrator. This does not create a third
level or mix the two existing ones: it preserves a single operational team and administration only
adds installation management.

User, Support and Development remain mutually exclusive. The factory account remains as a pure
Administrator and recovery door. A combined account works tickets from its level and also sees the
administrative tools. This avoids duplicate accounts and role changes, with traceability of who
grants or removes the capability.

## 12. Operational Dashboard

1.0.0 incorporates a Dashboard for Support, Development and Administrator. It summarizes load,
urgency, age, times and trends of the work that each role can already consult. It does not expand
the User's access, does not score people and does not generate automatisms. It only shows aggregated
statistics; it does not list, identify or link tickets. Detail continues exclusively in the inboxes
and detail views of the tickets module.
