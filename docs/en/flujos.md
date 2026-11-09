# Flows

[Español](../es/flujos.md) · **English**

> **State:** as-built
> **Last updated:** 2026-10-30
>
> **Proposal of 2026-10-30, approved, implemented and verified.** Section 3.2 allows Support to open the ticket on behalf of a
> person without an account, creating it in the same step. The local account does not receive access; with AD the person is required to
> be in the directory. The journey and the emails do not change.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** Section 14 adds the single notice to all active
> Support when a ticket reaches critical priority for the first time. No code is modified
> before implementing it. The owner explicitly approved it.
>
> **Amended on 2026-10-07**, in accordance with the approved proposal of 2026-10-06: all emails in
> the journeys use the global language in force when they are sent. Recipient groups are formed
> once, without consulting or separating personal preferences. The journeys, recipients and
> data of each notice do not change.
>
> **Amended on 2026-09-26**, when the ticket lists were split: where these journeys said "the
> Support queue" or "the Development queue" there are now **two places** —**My tickets** (mine) and
> the lists of everything, **Main tickets** and **Internal tickets**—, and **assigning a ticket to yourself is done
> inside the ticket** and not from the list (`docs/modules/tickets.md`, decisions 50 to 53). **The
> journeys and the emails do not change**: what changes is where each thing is seen.
>
> **It becomes as-built on 2026-09-25**: the step-by-step journeys and their emails **are implemented and
> verified** —the entire journey of a ticket, with escalation, the return to Support and the
> seven notices, has its interface case—. The status changes, not the content.
>
> Approved by the owner on 2026-09-22, after reviewing it in the form of questions while it was
> being written. The review uncovered a gap that was not in this document but in the two previous ones —the
> **return to Support** that `propósito-y-alcance.md` promised and had no door in
> `modules/tickets.md`— and from there came **rule 8**, the change to **rule 5** and the **seventh notice**
> by email.
>
> With this document, **the product chain is complete**: `auth`, `users` and `tickets` can be
> implemented in that order. `ambientes.md` remains for deployment and testing.

## 1. Scope of this document

It tells **step by step** what happens from the moment someone asks for something until the ticket is closed, who
does each thing, what is seen on the screen, what email goes out and what remains in the history. It is the
document that is read to build the screens and the emails.

It takes as written the permissions (`docs/usuarios-y-permisos.md`), the states, the transitions and the
endpoints (`docs/modules/tickets.md`) and the product (`docs/propósito-y-alcance.md`). **It does not repeat them**: if
something contradicts, those three prevail.

It is **as-built**: it describes the journeys and emails that the application executes today.

## 2. How to read it

Each flow carries the same, in the same order:

| | |
| --- | --- |
| **Who** | The roles involved |
| **Step by step** | What happens, in order |
| **What is seen** | What changes on the screen |
| **Email** | What notice goes out and to whom |
| **History** | What is recorded |

A rule that applies to all: **the email is a notice, not the action**. If a notice does not go out, what
the flow describes has already happened (docs/arquitectura.md, section 9).

## 3. Creating a ticket

**Who**: the user, or Support on their behalf.

### 3.1 The user creates it

1. They enter and press "New ticket".
2. They write the **subject** and **description**, **choose the category** —it is mandatory: without it there is no ticket
   (`docs/modules/tickets.md`, decision 65)—, add the **tags** they want and attach what they have
   (screenshots, above all). Support can correct the category later; the tags too.
3. They save.

**What is seen**: the ticket appears in the requester's "My tickets" with its number (`ACME-2026-0042`), in
state `nuevo`, and **already assigned to the technician whose shift it was** —so that technician also has it
in **their** "My tickets", and the others see it in **Main tickets**, which is the list of
everything there is (2026-09-26).

**Email**: **to the assigned technician**, "New ticket: ACME-2026-0042", with the subject and the link. If there is no
active technician, the ticket remains without an assignee and **nobody is notified**: it appears in
**Main tickets** so that someone picks it up, and it remains in the log.

**The distribution is by shifts** among the active technicians, and it also applies to the tickets that
Support creates on behalf of another (`docs/modules/tickets.md`, section 3.3). It can be reassigned by hand
later.

**History**: `creado`, with the user as actor.

### 3.2 Support creates it on behalf of another

The same, with two differences: Support **chooses the requester** (the person being helped) and
the ticket is left with **creator** and **requester** different.

**If the person has no account**, Support also writes their first and last names. With the local
method a `usuario` account is created **without a password and without a sign-up email**: it serves to identify
and notify them, but **it does not give them access**. With AD the person must be in the directory and the account
is created or linked from there; with Keycloak they must already exist in the realm. The ticket and the account
are created in the same step (`docs/modules/tickets.md`, section 35; `docs/modules/users.md`, section
18).

**What is seen**: the ticket appears in **Main tickets** —the list of everything there is—, in the "My
tickets" of the technician whose shift it was and **also in the requester's "My tickets"** when
they have access, because it is their ticket: they will see it and can reply.

**Email**: to Support, as in 3.1. **The requester is not notified that it has been created**: they did not
create it, and an email saying "you have created a ticket" would be noise. They do receive the later notices
of the ticket, which is what the product says.

## 4. Triage and attention in Support

**Who**: Support.

1. Support sees the new ticket —in **Main tickets** if it is not theirs, or in **My tickets** if it was
   their shift— and **reads it**.
2. **They assign it from the ticket itself**, with the assignee selector in its side panel (or they comment
   and ask the user first).
3. They move the ticket to **`en progreso`** while they work on it.

From there three paths open, and **Support chooses**:

### 4.1 It is resolved without escalating

4. Support moves the ticket to **`resuelto`** and writes what was done.
5. Support —or the user— **closes it**.

**Email**: to the user, "Your ticket ACME-2026-0042 is resolved" and then "…it has been closed".

### 4.2 Something is needed from the user

4. Support moves the ticket to **`en espera`** and asks in the conversation.
5. The user replies when they can.
6. Support **resolves from `en espera`** (there is no need to go through `en progreso`) or **escalates**.

**Email**: to the user, the change to `en espera` ("we need something from you"). It is the only notice that
exists to ask for something: there are no reminders or deadlines.

**The state does not change by itself when the user writes**: the ticket remains in `en espera` until
Support moves it. A comment is not always a reply —sometimes it is "thanks" or a new question—, and taking
something as answered when it is not is worse than moving it by hand.

### 4.3 It cannot be resolved: it is escalated

See section 5.

### 4.4 The ticket should not have existed

If it is a duplicate, a test or something created by mistake, Support —or the user, if it is theirs— **closes it
without resolving it**. There is no `cancelado` state: the history says it was closed without going through
`resuelto`, and that is all the difference that is needed.

## 5. Escalation to Development

**Who**: Support escalates. Development receives.

1. Support **escalates** the ticket and **writes the reason**: what was tested, what was ruled out and what is
   needed. **The reason is mandatory**: without it, escalation is "you pass it on".
2. The system **creates the internal ticket** (`INT-ACME-2026-0042`) with that reason, in state `nuevo`.
3. The main ticket moves to **`escalado`**.

**What is seen**: the main ticket remains in Support's lists, now in `escalado`; the internal one appears in
**Internal tickets** —and in Development's "My tickets" when it has an assignee—. As this is the first
escalation, the internal ticket is created.

**Email**: to Development, "Escalated: INT-ACME-2026-0042", with the reason.

**History**: `escalado` in both tickets, with Support as actor.

**How the internal one is read**: Development reads the main one (the conversation with the user included)
but **does not write in it**. Everything it has to say goes in the internal one, and Support decides
what it passes on to the user.

## 6. Work in Development

**Who**: Development.

1. They see the internal one —in **Internal tickets** or in **My tickets** if it is already theirs—, **assign it to themselves from
   the ticket itself** and move it to **`en progreso`**.
2. If, on looking at it, they think **it is not a code change** —it is something Support can resolve—, they
   **close it without resolving it** with a comment that explains it. **The main ticket returns to
   Support's hands** (rule 8), and Support receives the same "it has returned to your queue" notice.
3. If they need something from Support (a test, a piece of data, a confirmation), they move the internal one to
   **`en espera`** and ask for it in the conversation.

**What is seen**: the main ticket moves to **`en progreso`** while the internal one is `en espera`: what is
awaited is Development's business, and the user does not have to see the ticket as "escalated"
forever.

**Email**: **to Support**, "Development needs something in INT-ACME-2026-0042". It is the only notice that
warns inward at level 2, and without it the ticket can sit waiting for someone who does not
know it is their turn.

4. When Support replies, the internal one returns to `en progreso`.
5. Development **resolves** the internal one, with a comment that explains what was done. That comment is
   read by Support, not the user: it lives in the internal ticket.

**What is seen**: the main ticket returns to **`en progreso`**, in Support's queue. The case is not closed
from Development: **it returns to level 1**, which is who talks to the user.

**Email**: **to Support**, "Ticket ACME-2026-0042 has returned to your queue". Without this notice, a case
that Development has already finished can remain unexplained to the user.

6. **Support reads the internal one**, writes in the main one what was done —in the user's words, not technical
   notes— and **resolves** the main one.

**Email**: to the user, "Your ticket ACME-2026-0042 is resolved". And from there, Support or the
user closes it.

**Why the last step belongs to Support**: Development has no contact with the user, so the
explanation that they read cannot come from the internal ticket. It is rule 5, and it is the reason
Development does not resolve the main ticket.

## 7. Closing and reopening

### 7.1 Closing

**Who**: Support or the requester.

1. The main ticket is `resuelto`.
2. Support —or the user— **closes it**.

**Email**: to the user, "Your ticket has been closed".

**The internal one is not closed by itself.** If it is still open in `resuelto`, Support or Development **closes it by
hand**. It is a housekeeping task, not an effect of the main ticket, and that is why it can remain pending without
anything breaking.

### 7.2 Reopening

**Who**: Support or the requester, on a `cerrado` ticket.

1. It is reopened: the main ticket returns to **`en progreso`**.
2. `resolved_at` and `closed_at` are cleared: the screen cannot say "resolved on…" with a date
   from the previous time. The history preserves when each thing happened.

**To escalate again you must reopen first**: a closed ticket is not escalated.

## 8. Re-escalation

**Who**: Support, on a main ticket in `resuelto` (or already reopened).

1. Support **escalates again**, with its reason again.
2. The **same internal ticket** —not a new one— returns to **`en progreso`**, even if it was `resuelto`
   or `cerrado`.
3. The main ticket moves to `escalado`.

**Email**: to Development, the escalation notice **again**: the work has reappeared, and that
is what the notice means.

## 9. What was decided when reviewing this document

When writing it a gap appeared that was not in this document, but in the two previous ones: the synchronization rules
**had no door for the "return to Support"** that `docs/propósito-y-alcance.md`
promises —the case in which Development looks at the internal one and concludes that it is not a code change—, and
they did not say either who explains to the user what was done. Both things forced amending
`docs/propósito-y-alcance.md` and `docs/modules/tickets.md`.

| # | Decision | It was left like this |
| --- | --- | --- |
| 1 | **The return to Support** | Closing the internal one without resolving it **returns the main ticket to `en progreso`**. It is the **new rule 8** |
| 2 | **Who resolves the main ticket** | **Support, always**, and with its explanation for the user. **Rule 5** changed: the resolved internal one returns the main ticket to `en progreso`, not to `resuelto` |
| 3 | **Commenting on a closed ticket** | You cannot: to keep talking you must reopen it |
| 4 | **Deadline to reopen** | None: a closed ticket can be reopened no matter how long it has been |
| 5 | **The ticket created on behalf of another** | The requester sees it in "My tickets" and does not receive a notice of its creation: they did not create it |
| 6 | **The format of the emails** | Plain text, with the number, the subject and the link |
| 7 | **The notice to Support when the ticket returns** | **Seventh notice**: it covers the two paths by which the main ticket returns to Support (Development resolved the internal one, or returned it without resolving it) |
| 8 | **Assigning the internal one to oneself** | Development can assign the internal ones to itself, and Support can assign them to someone |
| 9 | **Where tickets are assigned** | **Inside the ticket**, in its side panel: in the main one Support reassigns to any technician, and in the internal one Support and Development to any developer. **The queue does not assign** (owner decision, 2026-09-26) |

## 10. What is NOT automated

None of the following exists, and it is written so that it is not taken for granted:

- **Deadline reminders**: if a ticket has been `en espera` for days, nobody is reminded by email.
- **Automatic escalation by time**: a ticket does not go up to Development because it has taken long.
- **Automatic closing**: a `resuelto` ticket does not close itself after a few days.
- **Notice to Development for each Support comment**: only for escalation.
- **Automatic reopening** when the user replies: if the ticket is closed, it must be reopened.

## 11. Flows that are not here

- **Accounts**: sign-up, first access, forgotten and password reset, and the coexistence of the three
  entry methods. They are in `docs/usuarios-y-permisos.md`, sections 4 to 9.
- **Deployment and testing**: they go in `docs/ambientes.md`.

## 12. What this document enables

With this document approved, **the product chain is complete**: `auth`, `users` and `tickets`
can be implemented in that order, and the screens and the emails have their journey written. `ambientes.md`
remains for deployment and testing, which is the last one.

## 13. Single language in all notices

The notices and emails of the six journeys do not change for now, nor do their recipients or data. Only
the choice of template changes: each send uses the global language that is active when it is sent.
When there are several people, a single group is sent and it is no longer separated by account language.

A global change only takes effect after confirming the eleven translations, so that no
journey can send a mixture of languages. The owner confirmed it on 2026-10-06. The review
confirmed that an open session adopts the change on its next request and that no flow keeps
a personal language. The rule was implemented and verified on 2026-10-07.

## 14. First arrival at critical priority

Creating a critical ticket or changing an open one from low, medium or high to critical generates a single
notice during the entire life of the main ticket. It is received by all active Support accounts. Lowering the
priority and returning to critical does not repeat the sending; saving critical over critical either.

Priority and history are saved first and then the sending is attempted. If the email fails, the API
confirms the change and also communicates the failure of the notice so that the person does not believe it was sent.
There is no automatic retry or new queue. This notice does not assign, escalate, change state or alter the
existing notices.
