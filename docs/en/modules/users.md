# Users: the module

[Español](../es/modules/users.md) · **English**

> **State:** as-built
> **Last updated:** 2026-10-30
>
> **Proposal of 2026-10-30, approved, implemented and verified.** Section 18 creates a requester account from the creation of a
> ticket when Support types an email that has no account. The local account is born without a password
> and without a creation email; with AD it is created from the directory and, if the person is not there, the ticket
> is rejected. Administrative permissions and entry rules do not change.
>
> **Amendment approved, implemented and verified on 2026-10-20.** Section 16 reserves the entire
> module and the administrative AD search to the Administrator capability. Support without Admin stops
> administering accounts; the ticket-specific people search boxes remain outside this module.
>
> **Amendment approved, implemented and verified on 2026-10-20.** It amends the previous decision that
> prevented creating directory accounts from `/users`. With AD as the current method, Support and
> Administrator may search by user or email, select a person and create an active account with the
> directory data. Keycloak keeps automatic creation on first access
> as long as there is no service account able to query users. The owner closed
> 1A–15A, with 6B.
>
> **Implemented on 2026-10-18.** The module stores an operational role and an independent administrative
> capability, with auditing and a persistent notice to the affected account.
>
> **Finding fixed and verified from the third public pass, 2026-10-13.** Automatic creation from AD and
> Keycloak worked, but the checks by email and by external identity logged the
> normal absence as a GORM error. Section 12 proposes keeping the same creation or link without
> that false error. The owner chose and approved the general scope 1A–4A; the tests passed
> against real PostgreSQL.
>
> **Amended on 2026-10-07**, in accordance with the approved proposal of 2026-10-06: `users.language`
> disappears from the migration, the model, the APIs and the screens. The profile keeps first name and
> last names; emails use the global language. Backend and frontend were verified.
>
> **Becomes as-built on 2026-09-26**: the module **is complete** —backend, the three screens and the three
> actions that ask the directory, with the five amendments below—, so this document
> describes what the code does. The header stayed at `aprobado` because of a slip of form, not
> because anything was missing.
>
> **Amended on 2026-09-25**: the three actions that depended on the directory —the ones
> the owner decided that same day and were left undone— **are done**, with the scope he
> set: **reactivating an AD account asks the directory** whether it still knows that person, and
> **the Keycloak one is not reactivated by hand** (it is explained that it comes back on its own on entry, because asking
> would require a service account in the realm); and **creation and the change of origin toward the directory
> are not done by hand**, with a message that says that instead of talking about the directory. With this, **this
> document has nothing pending**.
>
> Approved by the owner on 2026-09-23, after reviewing it in the form of questions. It fixes
> the `users` module: the accounts table, the endpoints, the creation rules, the role
> change, deactivation and the own profile.
>
> It is implemented together with the other one: you cannot have accounts without being able to enter, or enter without accounts.
>
> **Amended on 2026-09-25**, when finishing the AD path of `auth`: point 4 of section 5 said
> that its three rejections would be removed "in the same change that brings the check", and **the
> check is already here**. The owner decided the three on 2026-09-25: **reactivating a directory
> account asks the directory** (if it finds it, it is reactivated), and **manual creation and the change
> of origin toward `ad` are not done**, with a message that says that instead of talking about the directory.
> The three remain **decided and not done**, to implement them together as soon as the module is resumed
> (owner's decision: Keycloak is closed first). With this **the doubt is closed** in
> section 7 about whether the `users.directoryMayReturn` notice could be exercised: it is exercised, and there is an interface test
> in `docs/ambientes.md`, section 9.3.
>
> **Amended on 2026-09-24**, when finishing the module: section 9 debuts **`users.selfDeactivation`**
> (nobody deactivates themselves), section 8 fixes that **the factory account has no profile** and
> responds 404, and section 5 records that **everything that requires asking the directory** —the creation of
> directory accounts, their reactivation and the change of origin toward `ad` or `keycloak`— is rejected
> while that path does not exist. The three decisions are the owner's, of 2026-09-24.
>
> **Amended on 2026-09-24**, when building the screens, with five decisions of the
> owner that come from a finding: **Support can change the name and last names of another
> account, but does not see anyone's record** —decision 2 of this document—, so that permission had no
> place to be exercised. It remains like this: **Support edits them inline, from the list itself** (decision 15);
> **creation is a dialog** over the list (16); **origin is offered with its three values, and `ad` and
> `keycloak` disabled** with their note (17); **the list is a table on PC and cards on mobile, with filter
> chips** (18); and **the record and the profile are two routes**, `/users/{id}` and `/profile` (19). The
> five are written where they belong, in `docs/interfaz-y-experiencia.md`, section 3.6: **this
> document tells the module, not how it looks**. What does hold here is that **decision 2 does not
> change** —Support still does not see anyone's record— and that its editing limit is now exercised in the
> list.
>
> **Implementation status (2026-09-24): the module is complete**, backend and frontend. Creation, the
> resending of the link, **the list** with filters and pagination, **the record**, **the changes** (with
> Support's limit), **the own profile**, **deactivate**, **reactivate** and **the change of origin**,
> tested end to end, plus **its three screens** —the list, the record and the profile—, which
> `docs/interfaz-y-experiencia.md`, section 3.6, fixes. What depends on the directory is rejected while that
> path does not exist (section 5, point 4). **That path already exists and that point is done** (fifth
> amendment, above).
>
> **Amended on 2026-09-23**, when implementing the local path of `auth`: the error table had no
> **HTTP code** —just as happened to `mail.md`— nor any key for an invented origin or language. `users.origin.unknown` and the now-retired `users.language.unknown` were added, with their codes, approved
> by the owner: 404 what does not exist, 409 the repeated email, 422 what is not valid and **403
> `users.role.notAllowed`**, which is a permission and not a typo.

## 1. Scope of this document

It tells **how the `users` module is built**: the accounts table, the endpoints with their request and
response shapes, the creation rules, the role change, deactivation and the own profile.

**It does not repeat** what is already decided elsewhere:

| What | Where it is |
| --- | --- |
| What each role can do | `docs/usuarios-y-permisos.md` (the matrix) |
| How entry works, session and passwords | `docs/usuarios-y-permisos.md` and, in detail, `docs/modules/auth.md` |
| The module's screens | `docs/interfaz-y-experiencia.md` |
| The endpoint list, in summary | `docs/modules/tickets.md`, section 5 |

It is **approved**, so it enables writing code (Rule 0). Section 11
collects what I proposed.

## 2. The accounts table

A single table, `users`, with column names in English and **values in Spanish**, which is the same
criterion the ticket states already follow.

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | |
| `name` | `text` | First name. Required |
| `last_name` | `text` | Last names, in a single field. Required |
| `email` | `text` | Required and **unique case-insensitively** |
| `password_hash` | `text` | **Null** in directory accounts: there the password is not ours |
| `role` | `text` | `usuario`, `soporte`, `desarrollo` or `administrador` |
| `origin` | `text` | `local`, `ad` or `keycloak` |
| `external_id` | `text` | The directory identifier (Keycloak `sub` or AD GUID/account). Null in local ones |
| `language` | `text` | `es` or `en`. It is what decides **in which language emails are written to them** |
| `is_active` | `boolean` | True at birth |
| `last_login_at` | `timestamptz` | When they last entered |
| `created_at`, `updated_at` | `timestamptz` | |

**Constraints** (in the database, not only in the code):

- `email` **unique over `lower(email)`**: `Ana@…` and `ana@…` are the same person.
- `role` and `origin` only admit their values; a typo in a role cannot reach the database.
- **`password_hash` comes and goes with the origin**: required if the origin is `local`, and **null** if it is
  `ad` or `keycloak`. Thus the rule "if the origin is directory, there is no local password" does not depend
  on the code remembering it.
- **`external_id` unique per origin** when it is not null: two accounts cannot point to the same
  directory user.

**There is no `deleted_at`**: deletion does not exist (`docs/usuarios-y-permisos.md`, section 9). An account
is deactivated and remains readable by the tickets it has.

**The factory administrator is not here.** It is the only account that has no row: it lives in the
configuration (`docs/usuarios-y-permisos.md`, section 8). That is why this table **does not need a
username column**: everyone enters with their email.

## 3. How someone is identified

- The **email** is the identifier of people, and it is also what matches a local account
  with the directory one.
- The **factory administrator** is identified with `admin`, and is not a row of this table.
- Nobody else has a username or alias: if two people share an email, they are not two people.

## 4. The endpoints

These are the module's, and they are the only ones: the `auth` list lives in `docs/modules/auth.md`, and
`docs/modules/tickets.md` keeps the ticket ones.

| Method and route | What it does | Who |
| --- | --- | --- |
| `GET /api/users` | Paginated list, with filters by role, origin and state, and search by name or email | Support and Administrator |
| `POST /api/users` | Creates an account | Support (only `usuario`) and Administrator |
| `GET /api/users/{id}` | The detail of an account | Administrator |
| `PATCH /api/users/{id}` | Changes name, last names, role, email, language or state | Administrator; Support only the name and last names |
| `PATCH /api/users/me` | Changes **your** name, your last names and your language | Any authenticated account |
| `POST /api/users/{id}/deactivate` | Deactivates | Support and Administrator |
| `POST /api/users/{id}/activate` | Reactivates. In an AD account, only if the directory still knows it; in a Keycloak one, it is not done from here | Support and Administrator |
| `POST /api/users/{id}/origin` | Changes the origin of the account | Administrator |
| `POST /api/users/{id}/reset-password` | Launches a reset: the user receives the email | Support and Administrator |

- **State-changing actions have their own route** (`/deactivate`, `/activate`) instead of
  being a `PATCH` with a field: that way nobody can be deactivated by accident while editing their
  name, and the request log makes clear what happened. That is why the `PATCH` **does not** change the
  state: if it receives it, it responds `users.stateHasItsOwnAction` and touches nothing. The table above names
  state among what the `PATCH` changes; the reason prevails, and state goes through its actions.
- **`/origin` and `/reset-password` would not exist if they were not dangerous**: they are the only two
  actions that can leave someone unable to enter, and that is why they have their own route and their own permission.
- The **detail** of an account is only seen by an Administrator: Support needs the list to search for
  whoever reports, not anyone's record.

## 5. The creation rules

1. **Required**: name, last names, email, role and origin. The first two because they are people
   and whoever creates the account knows them; the email, because it is what they enter with.
2. **The email is validated** in form and in uniqueness, case-insensitively. If it already exists, the error
   says so clearly: it is a creation, not a mystery.
3. **According to the origin**:
   - `local`: the creation link is generated and sent by email. The account **cannot enter until
     it sets the password**, because it does not have any.
   - `ad`: when it is the current method, the person is searched and selected in AD; the server
     queries them again and creates or links the account without a local password. Generic creation continues
     rejecting directory data written by hand.
   - `keycloak`: it is not created by hand because there is no service account to search for people;
     the account is born or updated on entry.
4. **What we ask the directory**, and why the same thing is not asked in the two
   paths:

   | Action | What is done |
   | --- | --- |
   | **Reactivate** an **AD** account | **Asks the directory** whether it still knows that person, with the same search used to enter. If it finds them, the account is reactivated; if they do not appear, `users.directory.notFound`. **Giving access back to someone who is no longer there would be giving back an access they cannot use**, because their password is from there |
   | **Reactivate** a **Keycloak** account | **It is not reactivated by hand**, and the reason is given with `users.directory.activatesItself`: its path **enters on its own and on entry it is reactivated**. Asking Keycloak whether that person is still there would require a **service account with read permission inside the realm**, and that is not asked of whoever administers it for a button that fixes itself (owner's decision, 2026-09-25) |
   | **Create** an `ad` or `keycloak` account | **It is not done by hand** (`users.origin.byDirectory`): whoever is in the directory enters alone and their account is born, linked or brought up to date on that access |
   | **Change the origin** of an account to `ad` or `keycloak` | **It is not done by hand** (`users.origin.byDirectory`), for the same reason: the origin of a directory account is set by the directory, when that person enters through its path |

   The **change of origin toward `local`** does work: it removes the link with the directory and, if the
   account is left without a password, sends the creation link.

   **If it cannot be checked** —the directory does not respond, or this installation has none
   **configured**—, `users.directory.unavailable` (503) is answered and the reason remains in the log: the
   failure is not that of whoever requests the reactivation.

   **Whoever asks the directory is `auth`** (amendment of 2026-09-25): this module declares the
   question it needs —"does the directory still know this person?"— and `auth` answers it, which
   is the module that speaks the protocol and **the one that has the configuration stored**. Thus the
   check is done with the same service account and the same filter used to enter, and there
   are no two copies of the directory configuration. The connection is made in `main.go`
   (`docs/modules/settings.md`, section 5.8).

5. **A deactivated account is not resent the link** (`users.accountInactive`): it would set
   a password to continue without being able to enter. It is reactivated first.
6. **Who can create what**: Support only creates accounts with role `usuario`; handing out `soporte`,
   `desarrollo` or `administrador` is an Administrator's business
   (`docs/usuarios-y-permisos.md`, section 3).
7. **Creation and email are done by two modules.** `users` creates the account; the link and the email are
   issued by `auth`, which is the one that knows about passwords and tokens. `users` does not touch `auth`'s tables:
   it asks its service.
8. **If the creation email does not go out**, the account exists and another can be launched:
   `POST /api/users/{id}/reset-password` sends the same link. It is the way out for whoever did not receive the
   first one, and there is no need to delete and create anything again.

## 6. Changing the role

- An **Administrator** does it, and it is an explicit action with its own record.
- **It changes instantly**: the role does not travel inside the token, it is read from the account on every request.
- The four roles remain the same: what changes is which one a person has. No new
  roles are created nor are permissions edited from the application.

## 7. Deactivating and reactivating

- **Support and Administrator** do it. Deactivating deletes nothing: the account stops being able to enter and
  the tickets it has keep telling its history with its name.
- **A directory account is not deprived of access from here** (rule 6 of
  `docs/usuarios-y-permisos.md`): if it is still active in AD or Keycloak, it enters again and is reactivated
  on its own. The interface says so when deactivating, so that nobody is left with the feeling of having done it.
- **Reactivating an AD account checks with the directory first**: if it is no longer there, it is not
  reactivated, because its password is from there and it could not enter. **The Keycloak one is not reactivated by
  hand**: its path reactivates it on its own on entry, and the screen explains it instead of offering the button
  (point 4 of section 5).
- **Nobody deactivates themselves** (`users.selfDeactivation`): deactivating yourself is staying out at
  that instant and whoever does it cannot return, because another person is needed to reactivate them. The screen
  does not offer that button on oneself either (owner's decision, 2026-09-24).
- **Deactivating blocks instantly**: the next request from that account is no longer valid, even if its
  token is still valid until it expires. What it does not do is erase the token from that
  person's browser: whoever tries something will receive a notice and will go to the entry screen.

## 8. The own profile

- Anyone can change **their name, their last names and their language** from `PATCH /api/users/me`. Nothing
  more: neither their email, nor their role, nor their origin, which are someone else's decisions.
- **Support's limit is for editing others**: Support only changes the name and last names of
  another account, but **of their own they change the same as anyone**, their language included. Applying that
  limit to their own profile would leave them unable to change the language of their emails.
- **Changing the language changes the language of the emails they receive**: it is the only personal setting that the
  server needs to know, and that is why it lives in the account and not in the browser. The theme, on the other hand, stays
  in the browser, because the server does not care.
- The profile is opened from the sidebar, with the other controls
  (`docs/interfaz-y-experiencia.md`, section 3.1).
- **The factory account has no profile**: it is not in the accounts table, so there is no row to
  change and **404 `users.notFound`** is answered (owner's decision, 2026-09-24). Its name and
  its password live in the installation configuration, and the profile screen is not shown
  to it.

## 9. The module's errors

With keys, which the frontend translates (`docs/interfaz-y-experiencia.md`, section 8):

| Key | When | Code |
| --- | --- | --- |
| `users.notFound` | The account does not exist | **404** |
| `users.email.duplicated` | The email is already in use | **409** |
| `users.accountInactive` | The link is resent to a deactivated account: it must be reactivated first | **409** |
| `users.selfDeactivation` | Someone tries to deactivate **their own** account | **403** |
| `users.stateHasItsOwnAction` | The state is changed through the general `PATCH`, instead of with `/deactivate` or `/activate` | **422** |
| `users.email.invalid` | The email does not have an email shape | **422** |
| `users.name.required` | The name or the last names are missing | **422** |
| `users.origin.unknown` | The origin is not `local`, `ad` or `keycloak` | **422** |
| `users.directory.notFound` | The directory **no longer knows** that person, so access is not given back | **422** |
| `users.origin.byDirectory` | An attempt is made to create a directory account, or to change its origin to `ad` or `keycloak`, **by hand** | **422** |
| `users.directory.activatesItself` | An attempt is made to reactivate a **Keycloak** account by hand, which reactivates itself on entering through its path | **422** |
| `users.directory.unavailable` | It could not be checked with the directory, or this installation has none connected | **503** |
| `users.origin.inUse` | The origin of an active directory account cannot be changed | **422** |
| `users.role.notAllowed` | Support tries to create a role that is not theirs; **the role is none of the four**; or someone tries to change a datum that is not theirs (another's email or role, or their own) | **403** |

- **`users.role.notAllowed` is a 403 and not a 422** because it is a permission: whoever receives it has asked for
  something that does not correspond to them, and has not written anything wrong.
- **An invented role uses the same key as a role that is not theirs**: in both cases the role cannot be
  set, and separating them would add a key that changes what no screen does.

## 10. The module's emails

`users` **does not send emails**: it asks `auth`, which issues the link, and `auth` asks
`mail`, which has the text and sends. The two it can trigger are the creation and the reset, and both
carry the same link. The text goes in the installation's global language (`docs/modules/mail.md`).

## 11. What was decided when reviewing this document

| # | Decision | It remained like this |
| --- | --- | --- |
| 1 | **`last_login_at`** | Yes: knowing when someone last entered is what allows deciding with judgment whom to deactivate |
| 2 | **The detail of an account** | Only Administrator. Support sees the list, which is what it needs to search for whoever reports |
| 3 | **One email, one account** | Yes, and the comparison **ignores case** |
| 4 | **The keys that were missing** | `users.origin.unknown`, and an invented role uses `users.role.notAllowed`. The old language key disappeared with the personal preference |
| 5 | **The code of each key** | 404, 409, 422 and 403, each one where it belongs |
| 11 | **The profile of the factory account** | It responds **404 `users.notFound`** and the screen is not shown to it: it is not in the accounts table |
| 12 | **Nobody deactivates themselves** | **403 `users.selfDeactivation`**: deactivating yourself is staying out without being able to return |
| 14 | **State is not changed by the `PATCH`** | The document contradicted itself: the `PATCH` table includes the state, but the same document explains that deactivation has its own route so it does not happen by accident. The reason prevails: the `PATCH` responds `users.stateHasItsOwnAction` |
| 13 | **Changing the origin, today** | What can be checked is checked (not to an active directory account) and **switching to `ad` or `keycloak` is rejected** while the directory query does not exist; switching to `local` does work |
| 6 | **Resending the link to a deactivated account** | No: 409 `users.accountInactive`, because it would be promising an entry that is not going to happen |
| 7 | **The creation of directory accounts** | Since 2026-10-20, AD is queried from `/users` when it is the current method; Keycloak keeps the creation on entry. The generic API does not accept directory data written by hand |
| 20 | **Reactivating a directory account** | **Asks the directory** whether it still knows it: if it finds it, it is reactivated; if not, `users.directory.notFound`. Decided on 2026-09-25 and **done that same day**, with a scope fixed when implementing it: **in Keycloak it is not asked** —it would require a service account in the realm— and its manual reactivation is rejected saying it comes back on its own on entry |
| 21 | **The message of manual creation and of the change of origin** | `users.origin.byDirectory` still protects the generic API and the change of origin. The preparation from AD uses its own endpoints and data queried again on the server |
| 8 | **The email address in the log** | Yes, by owner's decision: it is what allows investigating a "it did not arrive". Never passwords, tokens or the email content |
| 15 | **Where Support edits** | **Inline, from the list**, because it does not see anyone's record (decision 2). The finding: without a record, its permission to change name and last names had nowhere to be exercised |
| 16 | **The creation** | **A dialog over the list**, not a screen: it is six fields, and on saving you return to where you were |
| 17 | **The origin on the screen** | The three, with `ad` and `keycloak` **disabled and with their note**: it is visible that they are planned and something that is going to fail today is not offered |
| 18 | **The list** | **Table on PC and cards on mobile**, with filter chips (role, origin, state and search) and pagination. Fixed order, by last names and name |
| 19 | **The record and the profile** | **Two routes**: `/users/{id}`, which only an Administrator opens, and `/profile`, which opens from your name in the menu. The factory account has no profile |

## 12. What this document enables

With `docs/modules/users.md` approved, the next one is **`docs/modules/auth.md`**, which already has its
decisions made (stateless token of 10 hours, AD by search with a service account, links with
fingerprint and single use, test containers for the directory). After that, `users` and `auth` are
implemented together: you cannot have accounts without being able to enter, or enter without accounts.

## 13. Global language, without personal preference

The migration removes `users.language`. The creation, list, detail, session and profile DTOs do not
accept it or return it; `users.language.unknown` disappears. AD and Keycloak no longer assign the installation
language when creating an account. `PATCH /api/users/me` changes only name and last names, and
`PATCH /api/users/{id}` does not admit language either.

The list, creation, record and profile screens remove language fields, labels and filters. Dates and
texts are formatted with the global language delivered by `settings`. The creation and password change
emails also use that global value, without reading the account.

Roles, origins, permissions, directories, deactivation and passwords do not change. The owner
confirmed the scope on 2026-10-06. The review added no personal exceptions or migration of
values. It was implemented and verified on 2026-10-07.
## 12. Implementation: searches prior to automatic creation

The first access via AD or Keycloak first searches for an account by email and by the
origin/external identifier pair. That both searches are empty is the normal case that leads to creating
the account; it will not be logged as an error. The methods will keep their current results of
found, absent and real error, and the service will keep the current rules to create, link
or update.

The tests exercise a new identity, an already linked one and a real database error, and capture the
logger to confirm that only the first absence remains silent. They do not change login,
the synchronized data, the roles or the permissions.

`FindByEmail` and `FindByExternalID` now distinguish zero rows with `RowsAffected`; `FindByID` applies
the same rule to the absent-account contract. The integration test confirmed the two
empty searches without error events and an invalid SQL query still visible. The full backend suite and
`go vet` passed on 2026-10-13.

## 14. Administrative capability of an operational account

`users.role` admits only `usuario`, `soporte` and `desarrollo`. `is_admin boolean not
null default false` is added, restricted to Support or Development. The existing rows with
`role = administrador` are migrated to `role = soporte, is_admin = true`.

A table of its own records in technical English the affected account, the actor —including the factory
identity—, the action `granted` or `revoked`, the instant and the moment when the affected person closed
the notice. Closing the notice does not erase the history.

List, record, own profile and `auth/me` expose `isAdmin`. Creation and editing present separately the
operational role and the capability. Granting or revoking requires an Administrator other than the affected
account, only admits Support or Development, keeps the capability when switching between those roles and
records the event. Support without administration keeps its current scope.

Deactivating or revoking the capability from the last ordinary Administrator is allowed because the
factory account remains. The API provides the signal needed for the interface to warn and confirm before
executing. The affected account queries and closes only its notices; the full history remains in the
administrative record.

AD and Keycloak keep creating as User and never derive or erase `is_admin`. The factory account
remains outside `users`, is a pure Administrator and has no record, profile or operational role.

Migration, filters, contracts, seeder and tests stopped interpreting Admin as an operational role.
The development database confirmed the column and the events table; the module's tests passed on
2026-10-18.

## 15. Creation from Active Directory in `/users`

The previous decision "directory accounts are not created by hand" is amended. Automatic
creation on first access is kept, but when **AD is the current entry method** an account may also
be prepared from `/users`:

1. Support or Administrator searches by user name or email. The search starts with three
   characters, admits partial matching and returns at most twenty people.
2. Each result shows full name, user name and email. Only people who
   meet the installation's base filter appear; a disabled account in AD does not appear.
3. A person is selected and the role is assigned. Support can only choose User; Administrator
   keeps the current model of operational role and administrative capability.
4. The server queries the selected identifier again before creating. The data sent
   by the browser does not replace name, last names, email or GUID from the directory.
5. The account is born active, with origin `ad`, without a local password and without a creation email. It can enter
   immediately while AD is the current method.

`GET /api/users/directory?q=...` offers the authenticated search and
`POST /api/users/directory` creates or links the selection. Both keep the permissions of
`POST /api/users`. The LDAP query is executed by `auth` behind an interface declared by `users`;
`users` does not read configuration or import internal code from `auth`.

If the pair `ad/external_id` already exists, it is not duplicated: the identifiable conflict is returned. If
the email belongs to a local account, the API requires `confirmLink: true`; the first response
explains that the local password will stop working. On confirming, the account identifier is kept, its local password is removed, its data is updated from AD and the change is recorded
in the administrative audit. A collision with another directory origin is not converted
automatically.

Keycloak remains outside these endpoints: without a service account with query permission
there is no correct way to search for people. Its accounts continue to be created, linked or
updated on first access. The interface explains it and does not simulate a query that does not exist.

The automated verification covers the search/creation contract, the options by method and the
escaped composition of the filter. The real verification confirmed entry with user name and with email;
the migration confirmed the audit table. The backend queries the selection again and keeps
the validations of permissions, duplicate, link confirmation and absence of a creation email.

## 16. Exclusive administration by Admin capability

List, record, local creation, search and creation from AD, editing, administrative capability, resending the
link, activation, deactivation and change of origin require `is_admin`. The factory account and
any account combined with Admin capability keep all the functions. The Support role by itself
stops authorizing any administrative `/api/users/**` endpoint and any `/users` screen.

This restriction does not remove the people search boxes that other modules need to assign,
watch or mention within tickets. Those contracts still belong to their operational flow,
return only the minimum data and do not allow administering accounts or querying AD.

Frontend, routes and controllers apply the same rule. The navigation tests verify the
access of a pure Administrator and of Support/Development + Admin, and the absence of administrative
entries for an unauthorized role; the Admin middleware protects the API.

## 17. A single accessible control to create accounts

> **Proposal of 2026-10-23.** The list renders simultaneously two controls with the accessible name "New account" for its responsive variants. The owner chose to include the
> fix via 4A and closed its verification with 5A. No open decisions remain; explicit approval
> is missing.

The adaptation between table and cards may keep different presentations, but at each width
there must be a single button exposed with the name **New account**. The hidden variant receives no
focus and does not remain in the accessibility tree. Visual position, action, permissions,
dialog and responsive behavior are kept.

The test uses queries by role and accessible name, without `.first()`, and confirms a single button on PC
and mobile.

## 18. Requester account from ticket creation

> **Approved, implemented and verified on 2026-10-30.** When Support opens a ticket for someone who has no account, the
> ticket creation needs a `usuario` account that identifies them. This section fixes how `users` creates it
> without granting access and without opening a new administrative path.

### 18.1 Contract

- **Who requests it**: only the `tickets` module, when Support creates a ticket with
  `requesterEmail` and there is no account. The contract is declared in `tickets` and implemented by `users`,
  like the rest of the accounts that appear in a ticket; `main.go` connects them. It returns the account
  and whether it was **created now** or reused.
- **The account is normal**: role `usuario`, origin `local`, active, without a password and **without a creation
  email**. It has no administrative capability and cannot enter until someone gives it access from
  Users or until it enters via AD or Keycloak with the same email.
- **`Create` is not reused**: the normal creation sends the link to set a password, and here
  precisely access is not wanted. The validations of name, last names and
  email are reused, but the path is its own and sends no email.
- **If the account already exists active**, it is returned and not touched. If it exists deactivated, it is rejected; it is not
  reactivated from the creation of a ticket.
- **With AD as the current method**: the person is searched in the directory by their exact email and
  created or linked with `CreateFromDirectory` (role `usuario`, no admin, directory data). If they are not
  in the directory, it is rejected and nothing is created. It does not fall back to a local account: with AD, a local
  requester account could not enter and the person should already be in the directory. If the
  account is **created** (not if an existing one is only linked), the ticket flow sends the
  `auth.accountCreated` email.
- **With Keycloak**: there is no people search in the realm, so nothing is created or linked.
  It is rejected with a key that says the account has to exist in Keycloak or enter once.
- **With the local method**: the account is created without access. If later the installation switches to AD or
  Keycloak, that account is linked by email on the first entry, like any local account.
- **Duplicates**: the database is still the one that guarantees the unique email. If two requests try to
  create the same account, one wins and the other re-reads the created account; a duplicate error is not returned
  to the ticket.
- **Auditing**: the identifier of the created account and the fact that it came from the creation of a
  ticket are recorded in the log. Names, email and the ticket content are not recorded.
- **Permissions**: it does not change who administers Users. The created account appears in the list and from
  there the link can be resent, it can be deactivated or its role changed according to the current rules.
- **Version and migration**: the new email requires template 14, which enters in `v1.0.1.sql`; the
  application and schema version becomes **1.0.1**. `v1.0.0.sql` is not rewritten, since it is already
  applied.

### 18.2 Why a ticket without an account is not created

The ticket stores `requester_id` as a mandatory foreign key and all the reading, the notices and the
requester's permissions depend on that account. Adding a requester without an account would force
changing the model, the lists, the emails and the permission matrix. The minimum account without access
solves the case without touching any of that.

### 18.3 Verification

- Service: creating the local account from the creation of a ticket leaves `origin=local`, `role=usuario`,
  `is_active=true` and `password_hash` null; no creation email is sent.
- Service: an existing active account is reused; an inactive one is rejected; a local creation
  without a name or last names is rejected.
- Service: with AD, a person from the directory is created or linked with their data and without a local
  password; a person who is not in AD rejects the ticket.
- Service: with Keycloak nothing is created and the required-account key is returned.
- The database guarantees the unique email; if the creation collides with another request, the service re-reads the
  created account and the ticket reuses it.
- Interface: the ticket shows the created account as requester and Support as creator.
