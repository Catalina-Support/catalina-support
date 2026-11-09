# Users and permissions

[Español](../es/usuarios-y-permisos.md) · **English**

> **State:** as-built
> **Last updated:** 2026-10-20
>
> **Amendment approved, implemented and verified on 2026-10-20.** Section 18 reserves Users
> and Classification to the Administrator capability, without taking from the teams the catalogs
> needed to work on tickets. It also sets the access to the unified Tickets module.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 17 adds private comments on the main ticket
> for Support and Development, without granting Development public replies or exposing content to the
> User. The owner closed and approved 1A–26A; reading was checked with both roles.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 16 restricts **My tickets** to current
> assignment or watching, without opening, comments or unassigned tickets. The operating role sets
> the level even with the Admin capability.
>
> **Implemented and verified on 2026-10-18.** Section 15 sets the access and scope of the Dashboard
> without modifying the operating ticket permissions.
>
> **Implemented on 2026-10-18.** The previous model treated `usuario`, `soporte`,
> `desarrollo` and `administrador` as four mutually exclusive roles. Section 14 describes
> separating the operating role from the administrative capability and accumulating their permissions.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** The final section adds the agreed priority
> permissions: the User proposes on creation; Support corrects on the main ticket; Development corrects the
> shared value only from an internal ticket it has access to; the Administrator keeps read-only
> access. The owner approved it explicitly.
>
> **Amendment implemented and verified on 2026-10-08.** Only Support and Development can
> improve with AI the drafts of the editors they are already allowed. The backend enforces it and the
> interface hides it from the User and the Administrator; it grants no new permissions on tickets.
>
> **Amended on 2026-10-07**, in accordance with the approved proposal of 2026-10-06: language stopped
> belonging to the accounts. It is single for the whole installation and only the Administrator changes
> it from Settings. Data, API, session and screens already comply with this rule.
>
> **Amended on 2026-09-29 (second time)**: **the password is only changed on local accounts**
> (owner's decision). The button that sends the link is renamed **"Change password"** —that is what it
> does— and **it is only shown for accounts with origin `local`**: with Active Directory or Keycloak the
> password **is checked by the directory**, so offering it would be a false promise. **And the server
> guarantees it**, not just the screen: requesting the link of a directory account responds **422** with
> `users.password.notLocal`. It fits a rule that was already in the database: **a directory account cannot
> have a password** (`users_directory_has_no_password`).
>
> **Amended on 2026-09-29**: **the catalog is maintained only by the Administrator** —create, rename and
> retire, categories and tags— (decision 83 of `docs/modules/tickets.md`, owner's decision).
> **Support and Development use it** —they classify tickets and filter— but **they do not change it**. This corrects the two
> rows of the matrix and **decision 73**, which was my proposal and is now decided the other way round.
>
> **Amended on 2026-09-27 (third time)**: **tags are also maintained** —create and rename,
> Support and Administrator; **retire, only the Administrator**—, with the two new rows of the matrix
> (decision 73 of `docs/modules/tickets.md`). This was the rule initially implemented; the
> amendment of 2026-09-29 that appears just above replaced it with decision 83 and closed the matter:
> today only the Administrator maintains the whole catalog.
>
> **Amended on 2026-09-27 (second time)**, when **categories and tags** came in: the matrix
> debuts three rows —choosing them, maintaining the catalog and **retiring a category, which is only the
> Administrator's**—. **Every ticket is born with a category**: it is mandatory and it is chosen by whoever opens it.
>
> **Amended on 2026-09-27**, when **mentions and watchers** came in: the matrix debuts the two
> rows of **tagging** and **removing a watcher** (Support and Development), and the new rule 3 states what
> can be most confusing: **being a watcher is not being the assignee**. **The assignee is still
> optional**, as it was.
>
> **Amended on 2026-09-26**, when distributing the ticket lists: the row "see their own tickets"
> then defined one's own by assignment, opening and comments. **Section 16 replaced that
> rule on 2026-10-19** with current assignment or watching for the teams, and
> Support and Development debut **two lists where they see everything** (Main tickets and Internal tickets)
> in addition to their inbox. **No permission changes**: what changes is where each role sees what it
> could already see (`docs/modules/tickets.md`, decisions 51 and 52).
>
> **Amended on 2026-09-25**, when the owner lowered **the minimum password length
> from 12 to 8 characters** (section 7), which is the only rule of the policy.
>
> **Amended on 2026-09-25**, when the owner asked that **the installation sign in through one method at a
> time** (section 5): the three paths of section 4 **no longer coexist**. The one that is set is the
> only one attended, and with the method set to `ad` or `keycloak` **local accounts —Support and
> Development included— cannot get in**; the factory account can, always, because it is the only
> door that cannot be closed (section 8). The amendment is told in full in
> `docs/modules/settings.md`, section 5.8, and in `docs/modules/auth.md`, sections 5.0 and 5.5.
>
> **Amended on 2026-09-24**, when `tickets` began to be implemented: the matrix debuts the row of
> **reopening the internal ticket** (Support), which is the missing half of "closing the internal one": the
> transition exists in `docs/modules/tickets.md`, section 3.2 —`cerrado` → `en progreso`, and it is equivalent to
> escalating again— and here it was not seen.
>
> Approved by the owner on 2026-09-22, with the seven decisions of the former section 10
> confirmed as they were proposed. It unblocks `modules/tickets.md` and `flujos.md`, and with them the
> `auth` and `users` modules, which are the first that can be implemented. **It does not enable writing
> code on its own** (Rule 0 of `AGENTS.md`).
>
> **Amended on 2026-09-22** in two things, when writing `docs/modules/tickets.md`:
>
> 1. The action of **closing the internal ticket** was added to the matrix (Development and Support): it was
>    decided that its closing is manual, not a consequence of closing the main ticket.
> 2. The action of **configuring the numbering prefix** was added (only Administrator), and
>    section 6 no longer says that the prefix is server configuration: it is changed from the
>    application.
> 3. **Section 8** was added, the factory administrator account, which is entered by typing
>    `admin` and whose password comes from `ADMIN_PASSWORD` and is applied on every start. Section 4
>    records that it is the only account that is not identified by its email.
> 8. When starting `docs/modules/auth.md`: **the session stops travelling in a cookie** and becomes a token in
>    `localStorage` sent in the `Authorization` header. It is chosen this way so that the browser does not
>    send anything on its own, in exchange for an XSS being able to read the token; the compensations and the
>    rule of not adding third-party scripts are written in section 6.
> 7. When documenting the users module: the `administrador` role now **sees the tickets in read-only
>    mode**, internal ones included, and its factory account does not exist as a row.
> 6. When documenting the users module: the **role change** was added to the matrix (Administrator),
>    the **removal** was taken out of the account lifecycle —it is only deactivated—, and the factory account stopped being a
>    row: **it lives in the configuration**.
> 5. When starting `docs/modules/auth.md` the **session duration was changed to 10 hours** and
>    what a stateless token implies was written: **the session cannot be revoked on the server**, so
>    signing out deletes the token from the browser but does not invalidate it. The account is read on every
>    request, so deactivating someone **blocks them instantly**.
> 4. When reviewing `docs/modules/tickets.md` the **password reset** was added to the matrix, which
>    **Support and the administrators** can trigger, and section 7 records the three ways of touching a
>    password: creation, one's own change and the reset (which sends email and never assigns a
>    password by hand).

## 1. Scope of this document

It sets **who can do what**, **how one signs in**, **how accounts are managed** and the **coexistence
rules** among the three access methods. It does not define the ticket data model
(`modules/tickets.md`) nor the step-by-step flows (`flujos.md`).

It is **approved**, so it enables writing code (Rule 0 of `AGENTS.md`).
Section 11 lists the decisions I proposed, all of them confirmed.

## 2. The four roles

They come from `docs/propósito-y-alcance.md`. Here only how each one signs in is added:

| Role | What it is | How it signs in |
| --- | --- | --- |
| **User** | Whoever needs attention. Creates tickets and replies. | Any of the three paths |
| **Technical Support** | First contact. Attends, escalates and manages account creation. | Any of the three paths |
| **Development** | Second level. Resolves what has been escalated. | Any of the three paths |
| **Administrator** | Manages accounts. Does not attend tickets. | Any of the three paths |

The role is **one per account** and it is not changed from the screen: it is decided when the account is
created (section 9).

## 3. The matrix: what each role can do

| Action | User | Support | Development | Administrator |
| --- | :---: | :---: | :---: | :---: |
| See their own tickets —for a technician, the ones assigned or currently watched— | ✅ | ✅ | ✅ | — |
| See **all** main tickets | — | ✅ | ✅ (read-only, rule 2) | ✅ (read-only) |
| See **internal** tickets | — | ✅ | ✅ | ✅ (read-only) |
| Create their own ticket | ✅ | ✅ | ✅ | — |
| Create a ticket **on behalf of another user** | — | ✅ | — | — |
| Comment on the main ticket | ✅ (their own) | ✅ | — | — |
| Comment on the internal ticket | — | ✅ | ✅ | — |
| Assign an assignee | — | ✅ | ✅ (the internal ones) | — |
| Set `en progreso` / `en espera` | — | ✅ | ✅ (the internal ones) | — |
| **Escalate** to Development | — | ✅ | — | — |
| **Resolve** the internal ticket | — | — | ✅ | — |
| **Close** the internal ticket | — | ✅ | ✅ | — |
| **Reopen** the internal ticket *(it is escalating again)* | — | ✅ | — | — |
| Mark the main ticket as `resuelto` **without escalating** | — | ✅ | — | — |
| **Close and reopen** the main ticket | ✅ (their own) | ✅ | — | — |
| **Choose the ticket category** | ✅ (their own) | ✅ | — | — |
| **Add, change and remove the ticket tags** | — | ✅ (from the main ticket) | ✅ (from the internal one) | — |
| **Maintain the category catalog** (create, rename and retire) | — | — | — | ✅ |
| **Maintain tags** (create, rename and retire, on all tickets) | — | — | — | ✅ |
| **Retire a tag** from the catalog and from the tickets | — | — | — | ✅ |
| **Retire** a category from the catalog | — | — | — | ✅ |
| **Tag** another technician or developer (and thereby make them a watcher) | — | ✅ | ✅ | — |
| **Remove** a watcher | — | ✅ | ✅ | — |
| See the list of users (name and email) | — | ✅ | — | ✅ |
| Create a user | — | ✅ (`usuario` role only) | — | ✅ (any role) |
| Deactivate or reactivate an account | — | ✅ | — | ✅ |
| Change the **origin** of an account | — | — | — | ✅ |
| Change the **role** of an account | — | — | — | ✅ |
| Configure the **installation** (numbering, dispatch, notices, appearance and email texts) | — | — | — | ✅ |
| **Reset the password** of a user | — | ✅ | — | ✅ |

Five rules that the matrix cannot express and that go separately:

1. **The user only sees their own.** Not other people's tickets, nor the list of users, nor the internal ticket
   of their own incident: they see their main ticket and nothing else.
2. **"Mine" is calculated with the current relation** (updated on 2026-10-19): the User keeps
   the main tickets they requested; for Support and Development they are the tickets assigned to them or
   where they currently appear as watchers. Opening or commenting does not add a ticket, and no longer
   watching it removes it unless it is still assigned. Support gathers only main tickets and Development only
   internal ones, even if they also have the Admin capability. The **Main tickets** and
   **Internal tickets** lists contain the rest of the authorised scope. The pure Administrator keeps
   their general read-only inbox (`docs/modules/tickets.md`, section 29).
3. **Being a watcher is not being the assignee.** A ticket has **at most one assignee** —the one who
   attends it— and **many watchers** —those who follow it, because someone tagged them—. **Watching grants
   no permissions**: it does not let someone write who did not write, nor move states who did not move them; it only makes
   the ticket be in their inbox (`docs/modules/tickets.md`, decision 60).
4. **Development does not write on the main ticket.** It reads it for context, but everything it
   has to say goes on the internal ticket, and it is Support who decides what is passed on to the user.
5. **Support creates users, but only with the `usuario` role.** If Support could create an
   `administrador`, the highest role would stop being protected. Handing out `soporte`, `desarrollo` and
   `administrador` is an Administrator's business.

Closing the internal ticket is manual on purpose: Development or Support consider it finished, and
**closing it does not change the state of the main ticket**. Re-escalating the main ticket returns it to `en progreso`.
The detail is in `docs/modules/tickets.md`, section 3.2.

## 4. The three entry paths

| | Email and password | Active Directory | Keycloak |
| --- | --- | --- | --- |
| **Creation** | Administrator or Support do it | The same, validating that the person exists in the directory; or **automatic on the first sign-in** | Same as AD |
| **First sign-in** | The user receives an email to set their password | The account is created with the `usuario` role | Same as AD |
| **Password** | Stored by the application | **Managed by the domain**: the application does not intervene | Managed by Keycloak |
| **Identifier** | The email | Email + directory identifier | Email + Keycloak `sub` |

The only account that does **not** identify itself by its email is the factory one, which is entered by typing
`admin` (section 8). The sign-in form says "email or username" because of it.

Notes:

- **The email is the person's identifier**: with it one signs in, receives the notices and matches
  the local account with the directory one. It is unique across the whole system.
- From AD the **directory identifier** is also stored (the GUID or the account) and from Keycloak the
  **`sub`**. They are the ones that rule once the account is linked (rule 3 of section 5).
- AD is queried with **LDAP** using the credentials the user types: there is no other way of
  knowing who they are without asking the domain. Keycloak, through OIDC.
- **The three paths are not open at the same time** (amendment of 2026-09-25): the installation signs in through
  **only one**, the one that is set, and it is changed from Settings without restarting anything
  (`docs/modules/settings.md`, section 5.8). Where the two directory paths are configured —on the
  screen, not in the environment— is also there.

## 5. Coexistence rules

They are the ones that prevent the same person from ending up with two accounts, or with none.

0. **The installation signs in through one method at a time** (amendment of 2026-09-25). An administrator chooses
   `local`, `ad` or `keycloak` from Settings, and **the other two stay off**: it is not that they do not
   work, it is that the installation does not offer them. Changing it takes effect on the next sign-in, without
   restarting anything.
   - **With the method set to `ad` or to `keycloak`, local accounts do not get in** —Support and Development
     included—, because their password is from here and the installation's path is another one. **Nobody loses
     their account**: what is turned off is the door, and on returning to `local` everybody gets in again.
   - **The factory account always gets in**, whatever the method (section 8): it is the one that can
     put the method back where it was without touching the database by hand.
   - **A method that is not configured cannot be chosen**: the installation does not end up without a door
     because of an oversight, because the backend rejects it before saving it.
1. **The directory rules**, **when the directory is the installation's path**. If the person
   exists and is active in AD or Keycloak, they get in. The institutional configuration outweighs the
   local state of the account. And the other way around: with the method set to `local`, a directory account does not get in,
   because here there is no password of theirs to compare.
2. **It is linked by email, never duplicated.** If, on signing in for the first time through AD or Keycloak, an
   account with that email already exists, the access is linked to it: it keeps its roles, its tickets and
   its history.
3. **Once linked, the directory identifier rules**, not the email. If the email is changed in the directory,
   the account remains the same and the email is updated.
4. **One origin per account** (`local`, `ad`, `keycloak`), and it is set by the first path that is used. If the
   person exists in both directories, they get in through the one they use first and stay there.
5. **If the origin is `ad` or `keycloak`, there is no local password**, even if the account had one
   before. The origin rules.
6. **Local deactivation does not block a directory account.** If a user deactivated here
   is still active in AD or Keycloak, **they get back in and the account reactivates itself**. To really remove their
   access they must be deactivated **in the directory**. This rule does not apply to `local` accounts,
   which have no directory to ask: there deactivating does block, and reactivating is an explicit
   action.
7. **Changing the origin is an Administrator's action**, explicit and recorded. It never happens by
   signing in from somewhere else.

## 6. The session

- **The token travels in an `Authorization: Bearer …` header**, not in a cookie. The frontend stores it
  and attaches it to every request with an interceptor; **the browser sends nothing on its
  own**, which is what was being sought: everything that travels, travels because our code puts it there.
- **Where it is stored: `localStorage`.** It is a **conscious decision, contrary** to that of the sibling
  project Calibyou, where the token in `localStorage` was recorded as a security finding
  because **an XSS flaw can read it**. It is accepted in exchange for having nothing automatic in the
  browser, and it is compensated with what is already decided: a strict security policy in production
  (no inline or third-party scripts), no user content is painted as HTML —the
  comments are text and Angular escapes them— and `svg` files are forced to download.
- **The rule that goes with that**: **no third-party scripts are added to the application**. Any
  script that is loaded can read the token; an analytics tool, a support chat or a font with
  JavaScript would open through the door what we have just accepted through the window.
- **Duration of 10 hours**, without sliding renewal: one signs in once and it lasts the long working day.
  It is the simple thing, and what avoids session windows left open forever.
- **Inside the token goes the identifier of whoever has signed in**, and nothing else: **the account is read from the
  database on every request**, which is what allows knowing its role and whether it is still active. It costs one
  indexed query per request and in exchange **deactivating and changing role take effect
  instantly**.
- **Signing out deletes the token from the browser**, and that is enough because **the token cannot be
  revoked on the server**: if someone had copied the token, it would keep working for them until it
  expires. That is the price of not having a sessions table, and it is written here so that nobody takes it
  for granted.
- **The interceptor reads the token on every request**, it does not keep it in memory: thus, signing out in
  one tab leaves the others without a session as soon as they request something again, instead of leaving them
  working with a session that no longer exists.
- **There is nothing to protect against CSRF**: a request from another site cannot set the
  `Authorization` header, because it does not have the token. That is the advantage of the explicit token, and that is why
  the `Origin` check that had been planned disappears.
- The signing secret is an **environment variable** (`TOKEN_SECRET`), never a value in the code or
  in git. **The only setting changed from the application is the prefix of the ticket numbering**,
  and an Administrator changes it (it is in the matrix and in `docs/modules/tickets.md`, section 2.2).

## 7. Passwords, password changes and account emails

- **Minimum length 8, with no forced expiry and no security questions.** Expiry pushes
  people to invent variants of the same password, and security questions are a recovery path
  weaker than email. **The minimum was 12 and the owner lowered it to 8 on
  2026-09-25**: in an internal help desk the password is not what protects the system —what
  protects it is who has an account, and an account is deactivated in one click—, and a high minimum only achieves
  people writing the same phrase with a number behind it. **The factory credential** (section 8)
  remains the only exception, and for a practical reason: if it applied to it, the backend would not start
  in development.
- **It is the only rule**: no mandatory uppercase letters, numbers or symbols. A composition rule
  does not add real entropy and does add passwords written down on a piece of paper.
- They are stored **hashed** (`bcrypt`), never in the clear and never reversibly encrypted.
- **Creation**: an email is sent with a **one-time link** so that the user sets their
  password. **It expires after 24 hours.** Without that step the account is no use for signing in, so creation
  and password are the same thing.
- **Change it whenever one wants**: a `local` user can change their password while signed in, without
  having to pretend they forgot it. With AD or Keycloak, no: the domain manages it.
- **Forgot it**: the same link as creation, with **1 hour** of expiry, because anyone can request that
  email and it is better for the window to be short.
- **Support or an administrator can trigger a reset** for someone who cannot get in. Support
  is who attends the user who calls saying they cannot get in, so it is who can best
  trigger it. The user receives the same email: **nobody assigns them a password or tells it to them over the
  phone**, which is how they leak.
- If the origin is AD or Keycloak, **nothing is sent**: the password belongs to the domain (rule 5).
- The account emails —**creation, recovery and the notice that the password has changed**—
  are **something different from the ticket notices**: those are **seven** and are listed in
  `docs/propósito-y-alcance.md`. The detail of the three is in `docs/modules/auth.md`, section 7.
- Neither passwords, nor tokens, nor recovery links are written to the logs.

## 8. The factory administrator account

A new installation has nobody who can create users, so it is born with **a single
account**. It is the one that configures the system the first time: creating Technical Support, Development
and the rest of the people.

| | |
| --- | --- |
| **Identifier** | `admin`. A constant in the code: it is not an environment variable, because it never changes |
| **Password** | The `ADMIN_PASSWORD` environment variable |
| **Role** | `administrador` |
| **Email** | **None.** It is the only account that is not required to have one, and it does not matter: it does not need notices |
| **Where it lives** | **In the configuration, not in the database**: there is no row to create |

Details that matter:

- **In development** the password is `admin`, and it is in `config/env/dev.env`, which is versioned with
  the rest of the local container's credentials.
- **It does not exist as a row**: it is the only account that is not in the accounts table. One signs in with it
  because the application recognises `admin` and **compares the password with the environment variable**, not with
  anything stored in the database.
- **In production it is mandatory**: the backend **does not start** without `ADMIN_PASSWORD`, and its value lives
  only in `config/env/prod.env`, which is not versioned. **The production password is not written in
  this documentation**, which goes to the repository.
- **The password policy of section 7 does not apply to it**: it is a factory credential, and
  requiring a minimum from it would prevent starting in development.
- **The password is only changed in the configuration.** It cannot be changed from the application
  —there is nowhere to store it— nor recovered by email —it has none—. Changing it is editing the variable and
  restarting, and thus **access to the installation is never lost**. No "I forgot the
  administrator password".
- **Nobody seeds it**: there is nothing to seed. The rest of the accounts are created by a person,
  and this one exists from before the first start.
- **As it is not a row, it cannot appear in a ticket** as author, requester or assignee: the
  tickets point to accounts that exist. Its job is to get the system running, not to attend.
- **It sees everything read-only**, tickets included: it can check what is happening, but not comment,
  nor move states, nor assign. The `administrador` role is the highest in **configuration**
  (users, numbering and appearance) and **read-only** in everything else.
- **It always gets in, whatever the entry method** (amendment of 2026-09-25), and that is why it is
  the first thing checked on signing in. It is what makes choosing the wrong method not leave the installation
  without an exit: with the method set to `keycloak` the sign-in screen has no form, so **its
  door is a discreet link** ("Sign in as administrator") that gets it out
  (`docs/modules/auth.md`, sections 5.0 and 9).

## 9. Lifecycle of an account

| Moment | How it works |
| --- | --- |
| **Creation** | Administrator or Support from the application (Support only with the `usuario` role), or automatic on the first sign-in through AD or Keycloak. The factory account (section 8) is the exception: it exists from the first start |
| **First sign-in** | Through the email link if it is local; direct if it comes from the directory |
| **Email change** | An Administrator changes it. If the account is from a directory, the directory rules |
| **Deactivation** | Administrator or Support do it. It blocks **instantly**: the account is checked on every request, so the next one it makes is no longer valid. On directory accounts it does not block if it is still active there (rule 6) |
| **Reactivation** | Automatic if the directory says it is active; explicit (Administrator or Support) on `local` accounts |

That is why the accounts table **does not have `deleted_at`**: it would be a column nobody would use.
| **Removal** | **It does not exist**: there is no removal action. A deactivated account remains readable through the tickets it has, and that is enough |

## 10. Where the permission is checked

- **In the backend, always.** Authorisation lives in `backend/shared/authz` and is applied with
  middlewares over the routes; no controller decides permissions on its own.
- **In the frontend, only to avoid showing what cannot be done.** Hiding a button is not a
  security measure: the decision is made in the backend.
- Modules involved: **`auth`** (signing in, signing out, passwords) and **`users`** (accounts, roles,
  deactivation). In the frontend, the session lives in `core` —guards, interceptor and sign-out—,
  because it is the shell and not a module screen; it is the exception already foreseen in
  `docs/arquitectura.md`.

## 11. What this document enables

This document, already approved, sets the permissions, the access and the lifecycle of the accounts, and
unblocks:

1. `modules/tickets.md` — data model, transitions and closed list of modules, which can now say who
   sees and who moves each thing.
2. `flujos.md` — the step-by-step flows, with the email notices.

With `modules/tickets.md` and `flujos.md` approved, **`auth` and `users` are the first modules that can be
implemented**, because they do not depend on any pending product decision.

## 12. Language is not a permission or a personal preference

- All accounts see the application and receive emails in the global language of the installation.
- Creating, synchronising or editing an account neither receives nor returns a language. Changing one's own profile is
  limited to first name and last name.
- The selector disappears from sign-in, menu, profile, creation and record. Only the Administrator changes the
  global language, and the factory account can do it when an update requires configuring AI.
- The existing personal column is removed from the model and from the APIs; the global language is not copied into
  each account.
- The global language does not modify roles, permissions, entry method, sessions or lifecycle.

This proposal replaces only the rules that attributed language and email to each account. The
owner confirmed it on 2026-10-06. The review confirmed that open sessions adopt the global
language on their next request, without changing permissions or closing the session. The rule was implemented and
verified on 2026-10-07.

## 13. Implementation: permission to improve a draft with AI

| Action | User | Support | Development | Administrator |
| --- | --- | --- | --- | --- |
| Improve with AI a description or comment that it can already write | No | Yes | Yes | No |

The permission requires both conditions: an admitted role and a current permission over the ticket and the editor.
It does not allow seeing a ticket, conversation or field that the account could not consult before. The backend
checks it on every request and responds 403 when it fails; the frontend only avoids offering an action
that will be rejected. The factory account keeps the Administrator capabilities and does not obtain
this operating permission.

## Ticket priority

| Action | User | Support | Development | Administrator |
| --- | --- | --- | --- | --- |
| Propose priority on creation | Yes | Yes, when creating on behalf of a User | No | No |
| Change the priority of an open ticket | No | Yes | Yes, if it can access its internal ticket | No |
| See priority and history | Its tickets | Yes | Yes, on the tickets it accesses | Yes, read-only |

Development's change modifies the same field of the main ticket. The User sees the current value and the
history with the previous priority, the new one, the author and the date, including the real name of whoever changed it.
Closing freezes the priority for all roles.

## 14. Cumulative administration

Ordinary accounts keep exactly one operating role: `usuario`, `soporte` or
`desarrollo`. Only Support and Development can additionally receive the administrative capability. The
permissions accumulate, but on tickets the operating role always rules: Support + Administrator
works as Support and Development + Administrator works as Development. Administration does not
mix both levels.

The factory account remains a pure Administrator, with no operating role and no ticket work. The
ordinary account that today has the exclusive `administrador` role becomes Support + Administrator when
the migration is applied. A User can never receive the administrative capability.

Only an Administrator other than the affected account grants or withdraws that capability. Switching between
Support and Development keeps it. AD and Keycloak synchronise the directory data, but keep
the local capability. The server resolves both dimensions again on every request, so the
change applies from the next request and the interface reflects it when navigating or reloading.

The factory account counts as an available Administrator. Withdrawing the capability from or deactivating the last
ordinary Administrator is allowed, with a confirmation that warns that only the factory account will remain.

Every grant and withdrawal leaves an immutable record with actor, affected account, action and instant,
visible in that account's record. The affected person receives on the next identity load a
notice with no email that remains until they close it explicitly.

The matrix of section 3 still describes the operating permissions; its Administrator column now
describes an additional capability or the pure factory identity. When two permissions coincide
it is enough that one dimension grants it; when a rule restricts the ticket side, the operating
role rules. This expressly corrects the current assumption of four mutually exclusive roles.

The migration, the constraints and the authorisation tests were verified on 2026-10-18.

## 15. Dashboard and reporting

User has no access. Support consults main and internal tickets; Development only internal ones; a pure
Administrator consults everything. An operating account with the Admin capability accumulates the total
scope. Only an Administrator modifies the delay thresholds. Opening a row grants no new permissions: the ticket
record applies its normal authorisation again.

The backend resolves this scope on every endpoint. The real test confirmed 403 for User and the
forced internal type for Development; accounts with the administrative capability use the accumulated
scope defined in section 14.

## 16. Scope of "My tickets"

For Support and Development, **My tickets** is not an additional read grant but a
personal subset of what is already authorised: assigned to the account or currently watched by it.
Support gathers only main tickets and Development only internal ones. Opening, requesting or commenting does not add the
ticket; lacking an assignee does not either. Removing the watch removes membership, unless the
assignment is still current.

The Admin capability does not modify this rule because on tickets the operating role rules. Access to the
rest continues in the general lists and each record applies its normal permissions again. The permissions
to assign, self-assign, watch or remove watchers are not modified.

## 17. Visibility of main-ticket comments

| Action | User | Support | Development | Pure Administrator |
| --- | :---: | :---: | :---: | :---: |
| Read public comment on the main ticket | Its ticket | Yes | Yes | Yes, read-only |
| Read private comment on the main ticket | No | Yes | Yes | Yes, read-only |
| Create public comment on the main ticket | Yes, on its ticket | Yes | No | No |
| Create private comment on the main ticket | No | Yes | Yes | No |
| Edit the body of a comment | Only its own public one | Only its own | Only its own private one | No |
| Change Public / Private | No | Only its own | Only its own private one | No |

The cumulative Admin capability does not alter the operating role. The User cannot infer a private
note through conversation, searches, summaries, attachments, emails or differentiated responses:
the direct download returns 404. Development keeps the prohibition to reply publicly on
the main ticket; the only extension is to collaborate internally through `private`. Comments from
state movements follow the current public rules.

## 18. Administrative modules and unified Tickets

| Module or action | User | Support | Development | Administrator capability |
| --- | :---: | :---: | :---: | :---: |
| My tickets | According to its tickets | Yes | Yes | Only if it also has an operating role |
| Tickets · Main | — | Yes | Yes | Only if it also has an operating role |
| Tickets · Internal | — | Yes | Yes | Only if it also has an operating role |
| Users and AD query | — | — | — | Yes |
| Classification: maintenance | — | — | — | Yes |
| Consult categories/tags to work | According to the form | Yes | Yes | According to its operating role |

Support without the Admin capability loses the list, record, creation, editing, activation, deactivation and AD
search of accounts. It keeps only limited people searches for those who are part of ticket
flows. The administration routes and endpoints require the Admin capability; hiding the menu is not the
security measure.

Classification separates operating consultation from maintenance. Support and Development continue reading and
assigning existing categories and tags inside the ticket. Only the Admin capability accesses the
administrative screen and can create, edit, activate or retire catalog entries.

The pure Administrator keeps their read-only inbox, Users, Classification and Settings,
but does not receive the operating Tickets module. A combined account accumulates both groups and sees, in this
order: Dashboard, My tickets, Tickets, Classification, Users and Settings.
