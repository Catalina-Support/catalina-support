# auth

[Español](../es/modules/auth.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-03
>
> **QA correction, 2026-11-03.** The `sesionActiva` guard compared the full URL with
> `/settings` and `/forbidden`; because the redirect adds `?ai=required`, it redirected again in a loop and
> left the application unable to load. It now compares **the path without the query** and does not repeat the redirect if
> it is already on the destination screen.
>
> **Amendment approved, implemented and verified on 2026-10-20.** It fixes sign-in against a real Active
> Directory: it accepts a username or email, keeps the installation's base filter and
> converts a binary `objectGUID` to its canonical GUID representation before querying or persisting it.
> The reproduced finding is that AD did accept the credentials, but PostgreSQL rejected the GUID bytes
> as UTF-8. Decisions 1A–15A, except 6B, were closed by the owner.
>
> **Implemented on 2026-10-18.** Identity resolves the operational role and administrative capability separately,
> keeping the factory account as a pure Administrator.
>
> **Amendment implemented and verified on 2026-10-13: absent token without a false error.** Looking up a
> nonexistent invitation or recovery keeps `auth.token.used` without logging
> `record not found`. Real SQL failures are not silenced. It is part of 1A–4A approved after the
> third pass of block 8.
>
> **Amended on 2026-10-07**, in accordance with the approved proposal of 2026-10-06: identity, session and
> account responses no longer carry a personal language. The three account emails consult the global
> language when they are sent, without changing any of the three sign-in paths.
>
> **Amended on 2026-10-03**, in accordance with the approved `docs/prueba-local.md`: Keycloak distinguishes a public
> issuer and an optional `internalIssuer`. Discovery, token and UserInfo use the internal address;
> authorization remains public. The advertised issuer is required to match the configured one, and endpoints must belong to the
> same origin and realm, and HTTP redirects are rejected. The configuration is read on every attempt.
> The local compose file advertises `http://127.0.0.1:11006/sso` by default; nginx is not needed.
>
> **Amended on 2026-10-01**: **the two test services leave `dev.yml`**. OpenLDAP and Keycloak
> now live each in **its own file** —`active-directory.yml` and `keycloak.yml`—, with **its
> own command**, and **the `auth` profile disappears**. Both join the
> **`catalina-support-dev`** network, which **`dev.yml` owns** (fixed name, no `external`), so **the
> environment is brought up first** and then the service; the backend reaches them by their service name
> (`ldap`, `keycloak`). Section 11 and decision 1 are brought up to date. The ports, container names
> and test credentials do not change.
>
> **Amended on 2026-09-25**, when the owner asked that **the directory and Keycloak be
> configured from the screen** and that **the installation sign in through one method at a time**: the data for
> the two paths **leave the environment and live in the database** —`directory_settings` and
> `keycloak_settings`, `docs/modules/settings.md` section 5.8—, `auth` reads them **on every attempt** through
> an interface it declares itself, and the method that is set decides where sign-in happens: the other
> two are turned off, with **the factory account always in** through its own door. The
> `LDAP_*` and `OIDC_*` variables **disappear**, and with this decisions 28 and 34 and
> section 5.4 are amended. **Consequence stated and accepted**: with the method set to `ad` or `keycloak`, local
> accounts —Support and Development included— do not get in (section 5.5).
>
> Approved by the owner on 2026-09-23, after reviewing it in the form of questions. It fixes
> the `auth` module: the session token, the three sign-in paths, the passwords and their
> links, the factory account, the account emails, the endpoints, the six shell screens (section 9)
and the containers with which it is tested.
>
> It is implemented together with the other one: there can be no accounts without being able to sign in, nor sign-in without accounts.
>
> **Amended on 2026-09-23** (when implementation began, section 13, step 1): the table in
> section 10 had no key for the **session token** —only for the email
> links—, so the middleware could not explain why it rejects a request. `auth.session.invalid`
> and `auth.session.expired` are added, `auth.forbidden` is added for the 403 of
> permissions (which had no key in any document: `docs/usuarios-y-permisos.md`, section 10, says
> where the permission is checked, but not with which key it is answered), and the HTTP status code
> that accompanies each key is fixed. The owner approved the decisions on 2026-09-23.
>
> **Implementation status (2026-09-23): the local path is done and verified.** There are
> sign-in (local account and factory account), `me`, `logout`, forgot, reset and change
> password, with their links and their three emails, tested end to end against development.
> **The six shell screens are done** and sign-in works from the browser (step 4).
>
> **Status (2026-09-25): the three paths are done and verified** and **this document is
> complete**. The local one, the AD one and the Keycloak one, all three against their real services (the
> test directory and the test realm live in the repository, in `config/ldap/` and
> `config/keycloak/`), with the five cases of section 5.4 checked one by one on each directory
> path and with the account that remains in the database. For the Keycloak path, in addition, the complete
> journey has been checked in a real browser: the button, the Keycloak screen, the return, the
> token in the fragment and the fragment erased.
>
> **Amended on 2026-09-25**, when implementing the Keycloak path. The
> four new error keys are fixed (section 10) and **how the browser reaches Keycloak in development**:
> through the same domain as the application and through a path (`/sso/`), which is what makes **the
> issuer a single one** and not two addresses that could disagree (section 11). Decisions
> 36 to 40 record what was decided when implementing it.
>
> **Background of 2026-09-25, amended on 2026-10-20**, when implementing the AD path:
> I brought him the link as a question with two options and he answered what he really wants —
> «with AD or with Keycloak I would like users to be able to sign in directly, **without the manual
> account creation process** by an administrator or technician». Direct entry is retained; since 2026-10-20
> an AD account can also be prepared by querying it from `/users`. What a person still decides
> is the **role**, because no directory says that. It is recorded as a correction in the
> record (31 to 35), together with the trigger for linking and what the sign-in screen does not change.
>
> **Amended on 2026-09-23**, when writing the local path: the input field,
> what sign-in returns, the link route and where its token travels were fixed. **The owner corrected a
> proposal**: I proposed the route in Spanish and he decided that **frontend routes always go in
> English**, even if the product language is Spanish. What he said is applied and recorded as a
> correction in the record (21).
>
> **Amended on 2026-09-23**, when writing the shell: it fixes **where the interface
> language comes from** and states **what the six screens are**, which until then were just a number.
> **The owner corrected another proposal**: I proposed signing in in Spanish by default and he decided
> that **the browser language is taken and, if it is neither of the two, English is used**, with the
> switch always visible. It is recorded as a correction in the record (24 to 26).

## 1. Scope of this document

It describes **how the `auth` module is built**: the link token table, the session token, the
passwords, the three sign-in paths with their mechanics, the factory account, the account
emails, the endpoints, the shell screens and the containers with which it is tested.

**It does not repeat** what is already decided elsewhere:

| What | Where it is |
| --- | --- |
| What each role can do | `docs/usuarios-y-permisos.md`, section 3 |
| The coexistence rules of the three accesses | `docs/usuarios-y-permisos.md`, section 5 |
| How the session travels and what it implies | `docs/usuarios-y-permisos.md`, section 6 |
| The password policy and its three ways of changing them | `docs/usuarios-y-permisos.md`, section 7 |
| The factory account and why it is not in the database | `docs/usuarios-y-permisos.md`, section 8 |
| The accounts table | `docs/modules/users.md`, section 2 |
| The screens and the visual language | `docs/interfaz-y-experiencia.md` |

It is **approved**, so it enables writing code (Rule 0). Section 12
records what I proposed.

## 2. The link token table

The creation and recovery links **are not signed tokens**: they are random values that are
stored, with their hash, in a module table.

**`password_tokens`**

| Column | Type | Notes |
| --- | --- | --- |
| `id` | `bigserial` | |
| `user_id` | `bigint` | Who it belongs to |
| `purpose` | `text` | `alta` or `recuperacion` |
| `token_hash` | `text` | **The hash**, never the token |
| `expires_at` | `timestamptz` | Creation: 24 hours. Recovery: 1 hour |
| `used_at` | `timestamptz` | When it was used. Null until it has been used |
| `created_by_id` | `bigint` | Who requested it. **Null if the person requested it themselves** |
| `created_at` | `timestamptz` | |

Decisions that go with the table:

- **The `sha256` of the token is stored, not the token.** If someone reads the database they cannot reconstruct
  any link; and comparing is as simple as calculating the hash of what arrived.
- **One link, one use.** When it is used, `used_at` is set and it stops being valid.
- **The token is 32 random bytes** in base64url. It is not signed or derived from anything: it is not needed,
  because it lives in a table.
- **Requesting a new one deletes the previous ones for that account and that purpose.** If someone requests three
  emails in a row, only the last one is valid; keeping the other two would add nothing, because they can no longer
  be used. The log records that they were requested.
- **`created_by_id` distinguishes who requested it**: if it is null, the person requested it from «I have forgotten my
  password»; if it has a value, Support or an administrator launched it. It is the only difference between the
  two cases, because the email is the same.

## 3. The session token

- **Signed with HMAC-SHA256** using `TOKEN_SECRET`, an environment variable that is **mandatory in
  production**: the backend does not start without it, as with `ADMIN_PASSWORD` and `PUBLIC_APP_URL`.
- **Content**: `sub` —the account identifier, or `admin` for the factory account—, `iat` and `exp`.
  **Nothing else**: not the role, not the name, not the language. All of that is read from the account on every
  request (`docs/usuarios-y-permisos.md`, section 6).
- **A 10-hour lifetime**, with no sliding renewal: `exp` is set when signing and is not touched again.
- **It travels in `Authorization: Bearer …`**, and is stored in the browser's `localStorage`.
- **Where the code lives**: signing and validating belongs to `backend/shared/auth`, **not** to the module. The reason
  is the modularity rule: the middleware that protects routes lives in `shared/middleware` and
  **`shared` cannot import from `modules`**. So the token mechanics are cross-cutting and the
  `auth` module uses them just as the middleware does.
- On every protected request, the middleware: validates the signature and the date, extracts `sub`, **reads the
  account** and leaves in the context who it is, what role it has and in what language they read. If the account is
  disabled or does not exist, it responds **401**.
- **The middleware does not know the accounts table**: it lives in `shared/middleware` and `shared` cannot
  import from `modules`, so **it receives a function that loads the identity** and `main.go` connects it
  to the `users` service. It is the same reason why the token mechanics are in `shared`:
  to keep the modularity rule without duplicating code.
- **What this cannot do, and it is worth knowing**: since there is no session table, **changing the
  password does not invalidate already issued tokens**. Whoever had one will stay signed in until it
  expires, up to 10 hours. It is the price of not having stored sessions, and that is why the lifetime is
  hours and not days.

## 4. The passwords

- **Hashed with `bcrypt`**, with cost 12. A password is never stored or logged in plain text, neither in
  the log nor in an email.
- **Policy**: minimum **8 characters** —lowered from 12 by the owner on 2026-09-25—, with no
  expiration and no security questions (`docs/usuarios-y-permisos.md`, section 7). **Nothing else**: no
  uppercase, no numbers, no mandatory symbols. Every added rule makes the password people end up choosing worse.
- **The password is checked before spending the link** (2026-09-25): a single-use link that
  is spent when the password is not valid leaves the person in a loop —they request another one and start over—,
  and the error they just read already told them what they had to change. The text is checked first
  —and hashed— and only then is the link spent, with which the account is found. **It was found by a
  test that began as a password policy test**: when the minimum was lowered from 12 to 8, the whole
  path was tested, and with a short password the link was lost.
- **The factory password is compared in constant time** against `ADMIN_PASSWORD`
  (`subtle.ConstantTimeCompare`), because there is no hash to compare: the variable is the source of
  truth and this prevents response time from revealing partial matches.
- **Changing one's own password requires the current one**, except when arriving from a recovery
  or creation link, which is precisely the case in which it is not known.
- **A password is never given over the phone or by email** (`docs/usuarios-y-permisos.md`,
  section 7): both paths end in a link, and the user chooses.

## 5. The three sign-in paths

### 5.0 Where sign-in happens: one method at a time

**The installation signs in through one method at a time**, and an administrator chooses it from Settings: without
that, having a configured directory would mean that local passwords stop working without anyone
having decided it. What is set decides **where sign-in happens**, and the other two paths are not
served even if they are configured (`docs/modules/settings.md`, section 5.8).

| Method set | Where sign-in happens | What the sign-in screen offers |
| --- | --- | --- |
| `local` | Section 5.1: the database password | The form and the link to recover it |
| `ad` | Section 5.2: the organization directory | **The same form**, without the link to recover it |
| `keycloak` | Section 5.3: the return through the browser | The button that leads to the realm, and **no form** |

1. **The factory account always signs in, and it comes first**: its email is `admin` and its password lives in
   the configuration. It is the only door that cannot be closed —otherwise, choosing the wrong method would leave the
   installation with no one able to change it— and **with the method set to `keycloak` there is no form**, so
   its door is a discreet link («Sign in as administrator») that brings it up.
2. Then, the method that is set. Without installation configuration —a newly
   installed installation, which does not yet have its row— the method is `local`.
3. **An account belongs to one path and not two**: a Keycloak account does not sign in through the directory,
   even if the method set is the AD one, and an AD account does not sign in with a local password.

**And with the method set to `ad` or `keycloak`, local accounts do not get in.** It is not a side effect
that can be avoided: it is what «one method at a time» means, and that is why it is stated here and in
`docs/usuarios-y-permisos.md`, section 5. Returning to `local` brings them all back.

### 5.1 Email and password (local)

**It is the factory path**, and with the method set to `local` it is the only one served.

1. The account is searched by email —and by username, for the factory account—, case-insensitively.
   **What arrives is a single field, `email`**, which also accepts `admin` for the factory account: the
   screen asks for «email» and does not force anyone to know that the factory account is different.
2. If it is `ad` or `keycloak`, **nothing is compared here**: it is resolved by its path (5.2 and 5.3).
3. The password is compared with its `bcrypt` and, if all goes well, the token is signed, `last_login_at` is updated
   and it is returned.

**A single message for a wrong password and a nonexistent account**: `auth.invalidCredentials` for the wrong
password, for the account that does not exist and for the one that does not yet have a password set. Saying
«that account does not exist» is giving away the list of who works here.

**A deliberate exception**: if the account exists and is **disabled**, the message says so
(`auth.accountInactive`), because that person cannot fix it alone and needs to call Support.
It is information given to someone who **has already proven they know the password**… except in the case of a
directory account, where there is no password to prove: there the same message is given as in
`auth.invalidCredentials`.

### 5.2 Active Directory

**The sign-in screen does not change for AD**: it is the same email and password field, and the one
that decides where to validate is the backend, which is the only one that knows what kind of account each one is
(section 5.1).

Someone who is in the directory **signs in directly**, with the password they already have, and their account is born
on that first access with what the directory says about them (section 5.4). In addition, Support or an
Administrator can prepare it beforehand from `/users`, always querying AD; they do not write its data
by hand or set a password. There the **role** of the account is decided, because the directory knows
nothing about that.


**How it is validated**: with a read-only **service account**. The person is searched for and
validated against the result, because in a real directory organizational units do not follow a
predictable pattern and building the full name by hand fails the day someone changes
department.

1. The backend connects to the directory **with the service account**.
2. It searches for the person by their username or email, combined with the configured base condition.
3. If they do not appear, `auth.invalidCredentials`.
4. If they appear, **it connects with their full name and their password**. That attempt is the validation: if
   the directory accepts it, they are who they say they are.
5. It reads their attributes and **updates or creates the account** (5.4).

**Where it is configured**: on the Settings screen, **not in the environment** (amendment of
2026-09-25). The data lives in the `directory_settings` table, an administrator changes it from there and
`auth` **reads it on every attempt**, so the change takes effect without restarting anything
(`docs/modules/settings.md`, section 5.8). These are:

| Data | What it is for |
| --- | --- |
| `host`, `port` | The directory. **An empty `host` means «this installation has no directory»** |
| `use_tls` | Encrypt the connection. **Yes in production**: credentials travel through it |
| `bind_dn`, `bind_password` | The service account. **The secret is stored in the database and is never returned by the API** |
| `search_base` | Where to search |
| `user_filter` | Base condition that limits the people admitted; it does not contain `%s` |
| `attr_username`, `attr_email`, `attr_name`, `attr_last_name`, `attr_id` | Which attribute is each piece of data |

- **The AD path is offered if the method set is `ad`** (section 5.0), and **that
  method cannot be chosen unless the directory is configured**: the backend rejects it with
  `settings.method.notConfigured` before leaving the installation without a door.
- **Each method shows its configuration and only its own** (owner's decision, 2026-09-27): in
  Settings, «Application accounts» shows no panel, the directory appears when choosing the
  organization and the realm when choosing Keycloak (`docs/interfaz-y-experiencia.md`, section 3.8).
- **The configuration is tested before saving it**: the «Test connection» button in Settings
  asks `auth` to open the connection with the service account and close it. **It does not validate anyone's
  password** —that is only known when someone signs in— and it does not touch the database.

**If the directory does not respond**, the error is its own (`auth.directory.unavailable`) and the screen
says that sign-in is not possible **because of a directory problem**, not that the password is wrong. They are
two different things and confusing them sends people looking for a password they have not lost.

### 5.3 Keycloak

**The callback is received by the backend**, not the frontend. This way there is **a single session** to maintain —ours—
and the frontend does not know that Keycloak exists.

1. The browser goes to `GET /api/auth/keycloak/start`.
2. The backend redirects to Keycloak with a **signed `state` that expires in 5 minutes**. It is not
   stored anywhere: it is signed with `TOKEN_SECRET` and validated on return. It is the way to
   protect itself without cookies.
3. Keycloak authenticates and returns to `GET /api/auth/keycloak/callback?code=…&state=…`.
4. The backend validates the `state`, exchanges the code for the Keycloak token and reads the `sub`, the email
   and the name.
5. It updates or creates the account (5.4), signs **our** token and redirects to the frontend with it **in the
   URL fragment** (`#token=…`), which does not travel to the server or remain in the logs.
6. The frontend stores it and **erases the fragment** from the address bar.

**Where it is configured**: like the directory, **on the screen and not in the environment** (amendment of
2026-09-25): the issuer, the client, its secret and the callback live in the `keycloak_settings` table
(`docs/modules/settings.md`, section 5.8). **The four go together**: without issuer, without client or without
callback, the configuration is half done and is rejected when saving (`settings.keycloak.incomplete`), and
**the path is offered if the method set is `keycloak`** (section 5.0). The secret is stored in the
database and is never returned by the API.

**How it is built** (as-built, and what was decided when doing it):

- **OIDC is spoken with the standard library**: the realm's **discovery document** is requested
  (`/.well-known/openid-configuration`), the code is exchanged against its `token_endpoint` and the
  person is read from its `userinfo_endpoint`. **No dependency is added** for this, and the document is
  requested **once and remembered**: it changes when the realm changes.
- **The `state` is a random value with the date, signed with `TOKEN_SECRET`** (HMAC-SHA256). It is not
  stored anywhere: what prevents an old return from being valid is **its five-minute
  expiry**, and since the date is inside what is signed, stretching it invalidates the signature. A `state` from another
  secret, a tampered one or an invented one are not valid, and **without a valid `state` nothing is exchanged**.
- **The exchange is a request from the server to Keycloak**, with the client secret. A 4xx from the exchange is
  «that return is not valid» (`auth.oidc.rejected`); a network failure or a 5xx is «Keycloak does not respond»
  (`auth.oidc.unavailable`). They are different things and they are answered differently.
- **Three things are read from Keycloak**: the `sub`, which is the **external identifier** and takes precedence over the
  email; the email, without which no account is created —email is what sign-in uses and what
  identifies tickets—; and the first name and last names, with the username as the last
  resort so that an account never ends up without a name.
- **The return leaves the person on the sign-in screen**, with the token in the fragment
  (`/login#token=…`) or with **the failure key** (`/login#error=auth.directory…`): it is a browser
  navigation in the middle of a journey, so even what goes wrong has to end somewhere
  where what happened can be explained. The screen **reads the fragment, uses it and erases it** with
  `replaceState`, so that the token does not remain in the address bar or in the history.
- **What is not done**: there is no sign-out in Keycloak (our session is ours and it is left the way
  it is always left), the direct password flow is not used —it is turned off in the
  client— and **neither the code nor the tokens are logged**: only the failure key goes to the log.
- **And a consequence of the above that must be known**: since the session is not closed in Keycloak,
  **its session remains alive after leaving the application**. Pressing the button again, Keycloak
  recognizes the person and **does not ask them anything**: it returns alone with a new code and signs them in. It is the
  normal behavior of an SSO, and the interface test treats it as one more case, not as a
  failure.
- **One thing that can be added and has not been added**: PKCE. With a confidential client and the exchange
  in the backend, the secret already protects the exchange; PKCE would additionally protect against interception of the
  code. It remains **proposed**, not done, and if it is done, the `code_verifier` would travel inside the signed
  `state` itself, which already exists for that.

### 5.4 What happens to the account, in both cases

When someone signs in through AD or Keycloak, and following the already approved coexistence rules:

| Situation | What is done |
| --- | --- |
| No account exists with that email | It is **created** with role `usuario`, directory origin and its external identifier |
| A local account exists with that email | It is **linked**: it becomes a directory account and **its local password stops working** |
| An account from that directory already exists | It is **updated**: first name, last names and email, which is what the directory dictates |
| It was disabled and the directory has it active | It is **reactivated** (coexistence rule 6). In a directory there is no «is active» that can simply be read: on sign-in, what matters is that the directory **accepts its credentials**, and that is what is checked |
| The email changed in the directory | The account **remains the same**: the external identifier dictates, not the email |

**The directory also dictates the data.** First name, last names and email are updated on every
sign-in: if someone changes their surname at the company, this finds out on its own and nobody has to notify
Support.

**When the directory is consulted, which is what decides which of the five rows applies.** The order
matters, and since the amendment of 2026-09-25 **the method that is set is the first thing that decides**:

1. **The factory account** (`admin`), which is not in the table: it is compared with `ADMIN_PASSWORD`. It comes
   before anything, with any method.
2. With the method set to **`local`**, the password is compared with the one in the database and **the
   directory is not asked**: something already known at home is not sent over the network. A directory account **has no
   password here**, so it does not get in with this method.
3. With the method set to **`ad`**, **the directory is always asked**, without looking at the local password
   first: it is the installation's path. If the directory accepts those credentials and their email was that of a
   local account, that account becomes a directory account at that very moment and its local password
   stops working —that is the **link**—. Nobody has to notify Support, nor does Support have to
   touch the account.
4. With the method set to **`keycloak`**, this endpoint does not sign in anyone other than the factory account: the
   path is the return through the browser (section 5.3).

**An account disabled here and active in the directory signs in again** (rule 6). It also applies
to linking: if the local account was disabled, it is reactivated when linked, because from then
on the directory dictates its access. When disabling a directory account the interface
already warns that it may come back (`users.directoryMayReturn`), and this is exactly that.

**If the directory does not respond or rejects the credentials**, the response depends on **whether it is known
where that person is from**, and that is known if there is a directory account with that email:

| Who signs in | What is answered |
| --- | --- |
| An account that **is already from AD**, and the directory does not respond | `auth.directory.unavailable` (503): it is known where that person is from, so it says the problem is the directory and not their password |
| An account that **is already from AD**, and the directory rejects those credentials | `auth.directory.rejected` (401): it is their password, and the password belongs there |
| An email that **has no account**, or a local account | `auth.invalidCredentials` (401), and the failure remains in the log. Here **it is not known** whether that person is from the directory, and telling a stranger how the house network is doing helps no one |

**And the screen does not change because of any of this**: both messages travel as keys and are read translated
in the same form as always (section 5.2).

### 5.5 What is left out when choosing a method

Choosing a method **turns the other two off**, and it must be said in full because it has consequences that are
noticed the next day:

| With the method set to… | Who can sign in | Who is left out |
| --- | --- | --- |
| `local` | The accounts in the database, with their password | Anyone who only has an account in the directory or in the realm |
| `ad` | Anyone who is in the directory | **Local accounts, Support and Development included**; and anyone who only has a Keycloak account |
| `keycloak` | Anyone who has an account in the realm | **Local accounts, Support and Development included**; and anyone who only has an AD account |

- **No one loses their account**: what is turned off is the door, not the account. Returning to `local`, everyone
  gets in again with what they had.
- **A local account cannot be sent the recovery link while the method is not
  `local`**: the link is set through a public endpoint, but signing in with that new password would
  be useless. It is a consequence to be aware of, not a failure.
- **The factory account is the exception and that is why it is there**: it is the one that can set the method back
  where it was without touching the database by hand.

## 6. The factory account

`admin` **is not a row**: it lives in the configuration
(`docs/usuarios-y-permisos.md`, section 8). Here is its mechanics:

- When signing in with `admin`, **the table is not searched**: the password is compared with `ADMIN_PASSWORD`
  in constant time. If it matches, a token is signed with `sub = "admin"`.
- The middleware recognizes that `sub` and gives it the factory identity: role `administrador`, reading
  everything and writing only in the configuration. It does not read any account because there is none.
- **It cannot change its password from the application or receive emails**: there is nowhere to store it
  and no email address to write to. Changing it means editing the variable and restarting.
- It is the only account that never appears as the author of anything.

## 7. The account emails

**`auth` does not send emails: it asks `mail` for them**, which is where the texts live. Here it only says
when they are requested and what they carry; the text, the placeholders and the sending belong to `docs/modules/mail.md`.

Three, and no more. They are **something different from the seven ticket notices**
(`docs/propósito-y-alcance.md`):

| Email | When | Link expiry |
| --- | --- | --- |
| **Creation** | When creating a `local` account | 24 hours |
| **Recovery** | When the person requests it, or when Support or an administrator launches it | 1 hour |
| **Change notice** | When a password changes: when completing creation, when using a link, when changing it from inside | It carries no link |

- **In the installation's global language**: `mail` chooses the template for that language.
- **HTML, with an automatic text version**, and the `mail` templates with their placeholders.
- **The expiry text is set by `auth`**, not `mail`: the templates carry the placeholder
  `{{caducidad}}` («this link expires in …») and the one that knows whether it is 24 hours or 1 is this module, which is
  the one that issues the link. That is why `auth` passes `24 horas` or `1 hora`… in the global language, which is
  the same one the template uses: `24 hours`, `1 hour`.
- **The address from which it was changed comes from `X-Real-IP`**, which is the header with which nginx
  passes the real address of the caller: inside the container, `RemoteAddr` is the proxy's.
- **The change notice says when and from where** it was changed, and ends with «if it wasn't you,
  tell Support». It is the only way for someone to find out that their account has been accessed.
- They are sent **in the background** and **an email failure does not bring down the action**
  (`docs/arquitectura.md`, section 9). If the change notice does not go out, the password is already changed.

## 8. The endpoints

| Method and route | What it does | Who |
| --- | --- | --- |
| `POST /api/auth/login` | Signs in with email and password. Returns the token and the session data | Anyone |
| `GET /api/auth/methods` | **Public.** How this installation signs in: `{"method": "local", "local": true, "ad": false, "keycloak": false}`. With no data about anyone | Anyone |
| `GET /api/auth/keycloak/start` | Starts Keycloak access: **redirects** to its sign-in screen with a signed `state` | Anyone |
| `GET /api/auth/keycloak/callback` | Returns from Keycloak: **redirects** to the sign-in screen with the token —or with the failure key— in the fragment | Keycloak (the browser) |
| `POST /api/auth/logout` | Signs out. Since the token cannot be revoked, here it is only logged | Authenticated |
| `GET /api/auth/me` | Who I am: identifier, first name, last names, email, role, origin and language | Authenticated |
| `POST /api/auth/password/forgot` | Sends the recovery link. **It responds the same whether the account exists or not** | Anyone |
| `POST /api/auth/password/reset` | Sets the password with the token from the link | Anyone with a link |
| `POST /api/auth/password/change` | Changes one's own password, with the current one | Authenticated, `local` account |

- **`forgot` always responds the same**, even if the email does not exist: if it said «that account does not
  exist», the endpoint would serve to find out who works here.
- **`logout` invalidates nothing** (there is no session table) but it is logged: it is the way to
  know when someone said they were leaving.
- **`me` is what the frontend uses on startup** to know whether the session is still alive and with what role.

What they return, when it goes well:

| Endpoint | Returns |
| --- | --- |
| `POST /api/auth/login` | `{"token": "…", "expiresAt": "…", "user": { … }}` |
| `GET /api/auth/keycloak/start` | **302** to Keycloak, with `client_id`, `redirect_uri`, `response_type=code`, `scope` and `state`. Without Keycloak configured, **404 `auth.oidc.notConfigured`** |
| `GET /api/auth/keycloak/callback` | **302** to `/login`, with `#token=…` if sign-in succeeded and `#error=<clave>` if not. **It never responds with the token in the body**: it travels in the fragment, which does not reach the server |
| `POST /api/auth/logout` | `{"status": "ok"}`: there is nothing to revoke |
| `GET /api/auth/methods` | `{"method": "local", "local": true, "ad": false, "keycloak": false}`. It says **the method that is set and which of the three paths is open** —only one, and the factory account signs in either way—. **It looks at the configuration, not health**: it says `ad: true` with the directory configured even if it does not respond at that moment, and what happens then is explained on sign-in |
| `GET /api/auth/me` | `{"user": { … }}`, **the same object** that sign-in returns |
| `POST /api/auth/password/forgot` | `{"status": "ok"}` **whether or not the account exists** |
| `POST /api/auth/password/reset` | `{"status": "ok"}` |
| `POST /api/auth/password/change` | `{"status": "ok"}` |

And the `user` object, which is the same in the three places where it appears:

| Field | What it is |
| --- | --- |
| `id`, `name`, `lastName`, `email` | Who they are |
| `role`, `origin`, `language` | Their role, where their account comes from and in what language they read |
| `factory` | Whether it is the factory account, which is not in the database |

## 9. The frontend: the shell

All of this lives in `frontend/src/app/core`, which is the exception already foreseen for the session.

| Piece | What it does |
| --- | --- |
| **Session service** | Stores the token, exposes the identity with signals and says whether there is a session |
| **Interceptor** | Sets the `Authorization` header on every request and **handles the 401**: it deletes the token and takes you to the sign-in screen |
| **Route guard** | Without a session, you do not get in: to the sign-in screen |
| **Sign-in screen** | **What it offers depends on the method set** (section 5.0): the email and password form with «I forgot my password» (`local`), the same form without that link and with a notice that the password is the organization's (`ad`), or only the button that leads to the realm (`keycloak`), with the discreet door of the factory account. **At the bottom, in the footer, the system version is shown** (`docs/modules/settings.md`, section 5.9) |
| **Forgot screen** | Asks for the email and confirms that it has been sent, **without saying whether it existed** |
| **Set password screen** | It is the link's screen: it arrives with the token in the address and asks for the new password twice |
| **Change password screen** | For someone already inside: current, new and repeat |

- **The interceptor reads the token on every request**, it does not keep it in memory: this way, signing out in one
  tab leaves the others without a session as soon as they request something.
- **A 401 is not handled on each screen**: it is handled in one place, and the sign-in screen explains
  that the session has expired.
- **The session service asks how sign-in works** (`GET /api/auth/methods`) and **adopts the token that
  the Keycloak return brings**, which the sign-in screen passes to it after reading it from the fragment.
  With that response the screen decides what to render: **the form if the open path is local or
  the AD one**, and **the Keycloak button only if the open path is its own** —offering a button that
  leads to an error is worse than not offering it—. If the backend does not answer, it stays with the local path,
  which is the one that always exists.
- **And with the method set to `keycloak`, the sign-in screen shows the factory account door**:
  a discreet link in the footer that brings up the email and password form. Without it, an installation whose
  method is chosen wrongly would be left with no one able to change it back (section 5.0).
- **The Keycloak button is not a form submit**: it is taking the browser to
  `/api/auth/keycloak/start`, because the one that authenticates is Keycloak and you have to leave with real
  navigation. It is named the same in the inventory —a secondary button, with its text in both languages— and
  since it sends nothing, it cannot be triggered with the form's Enter key.
- The token is stored in `localStorage` (`docs/usuarios-y-permisos.md`, section 6) and **it is not shown
  on any screen**.
- The texts of these six screens are in **Spanish and English**, like everything else, and **the
  startup language comes from the browser**: if the browser requests `es` or `en` —or any of their
  variants, such as `es-MX`— sign-in happens in that language; **if it requests any other one, sign-in happens in English**,
  which is the installation's second language and the one more people can read. The switch is visible on
  the sign-in screen, and what someone chooses is remembered in their browser. Inside the
  application the installation's global language dictates, which also decides the language of the emails
  (section 7).
- **The texts live in `core/i18n/es.ts` and `core/i18n/en.ts`**, with the same shared type: **a
  key that is missing in one language does not compile**. It is the only way for the rule «every key needs
  its text in both languages» not to depend on someone remembering it, and that is why the error texts
  go in the same dictionary as the screen texts.

**The six shell screens**, which are these and no more:

| # | Screen | What it is |
| --- | --- | --- |
| 1 | Sign-in | **What it offers depends on the method set** (section 5.0): the form with «I forgot my password», the same form with the notice that the password is the organization's, or only the Keycloak button with the factory account door |
| 2 | Forgot | Asks for the email and confirms that it has been sent, **without saying whether it existed** |
| 3 | Set password | The link's screen: it arrives with the token in the address and asks for the new password twice |
| 4 | Change password | For someone already inside: current, new and repeat |
| 5 | **No permission** | Someone reaches something their role cannot reach: they are told, they are not shown a blank screen |
| 6 | **Server down** | The application cannot talk to its backend: a clear notice, without technical jargon |

- **The last two are not «auth screens»**, they are shell states
  (`docs/interfaz-y-experiencia.md`, section 9): they live in `core` and all modules use them. The 5th
  is triggered by a 403 and also by a guard; the 6th, by any response saying the server is not there
  (or no response, which is what happens when it goes down completely).

The routes of the first four, **in English and lowercase** (corrected by the owner on
2026-09-23: routes always go in English even if the product language is Spanish):

| Screen | Route |
| --- | --- |
| Sign-in | `/login` |
| Forgot | `/forgot-password` |
| Set password (the email link) | `/set-password` |
| Change password (from inside) | `/change-password` |
| No permission | `/forbidden` |
| Server down | No route of its own: it is a notice placed **on top of** whatever is on screen, because it can happen at any moment and on any screen |

- **The email link points to `/set-password` with the token in the fragment**:
  `…/set-password#token=…`. The fragment **is not sent to the server**, so the token does not end up in the
  nginx access log or in any log, and the screen erases it from the address bar when
  using it —just as with the Keycloak return (decision 13)—.
- **A single screen serves both links**, the creation one and the recovery one: in both cases
  you arrive without knowing the password and leave with one set. The text changes, not the mechanism.

## 10. The error keys

| Key | When |
| --- | --- |
| `auth.invalidCredentials` | Wrong password, account that does not exist or account without a password set |
| `auth.accountInactive` | The account exists and is disabled (local accounts) |
| `auth.directory.unavailable` | The directory does not respond |
| `auth.directory.rejected` | The directory has rejected the credentials |
| `auth.token.expired` | The email link has expired |
| `auth.token.used` | The link was already used, or a newer one replaced it |
| `auth.password.tooShort` | Fewer than 8 characters. **It does not spend the link**: you can try again with the same one |
| `auth.password.wrong` | The current password is not the one that was entered |
| `auth.password.notLocal` | Attempt to change the password of a directory account |
| `auth.oidc.state` | The Keycloak access `state` is not valid or has expired |
| `auth.oidc.notConfigured` | This installation does not have Keycloak, and its path was requested manually |
| `auth.oidc.rejected` | Keycloak has not accepted the return: expired code, already used, from another installation, or someone who cancelled on its screen |
| `auth.oidc.unavailable` | Keycloak does not respond: its realm could not be read, nor the code exchanged, nor who they are |
| `auth.oidc.noEmail` | The Keycloak account does not bring an email, and without an email the account cannot be created here |
| `auth.session.invalid` | The `Authorization` header is missing, the token is not valid or its account no longer exists or is disabled |
| `auth.session.expired` | The session token has gone past its 10 hours |
| `auth.forbidden` | The account has signed in correctly, but its role is not enough for that action |

And with what HTTP status code each one travels, so that the frontend does not have to guess it:

| Key | Code |
| --- | --- |
| `auth.session.invalid`, `auth.session.expired` | **401** |
| `auth.forbidden` | **403** |
| `auth.invalidCredentials`, `auth.accountInactive`, `auth.directory.rejected`, `auth.password.wrong`, `auth.password.notLocal`, `auth.password.tooShort`, `auth.token.expired`, `auth.token.used` | **401** |
| `auth.directory.unavailable`, `auth.oidc.unavailable` | **503**, because the failure is the server's and not that of the person typing the password |
| `auth.oidc.notConfigured` | **404**: in this installation that path does not exist |
| `auth.oidc.state`, `auth.oidc.rejected`, `auth.oidc.noEmail` | **In the Keycloak return there is no HTTP status code that applies**: the browser has to end up somewhere, so they travel in the URL fragment and the sign-in screen translates them |

- **`auth.forbidden` is cross-cutting**: the check lives in `backend/shared/authz`
  (`docs/usuarios-y-permisos.md`, section 10) and all modules use it, but the key is a single one
  and the frontend translates it once. It is the only key in this table that is not born in the `auth` module.
- **If reading the account fails** —the database does not respond—, the request is not rejected because of the
  session: it is logged with `go-logs` and **503 with `error interno`** is returned, just like the
  panic middleware. A server failure is not an invalid session.

**`auth.session.invalid` covers the three cases on purpose** (no header, signature that is not valid and
account that is no longer there or is disabled): the frontend does the same in all three —delete the token and
take you to the sign-in screen—, and separating them would only add keys that do not change any screen.

## 11. What it is tested with

With **two test services**, each in **its own file** and **optional** —the environment does not bring them
up—: `active-directory.yml` (OpenLDAP) and `keycloak.yml` (Keycloak). **First the environment is brought up**,
which is what creates the development network, and then the service you want:

```bash
docker compose -f dev.yml up -d                  # el entorno (crea la red `catalina-support-dev`)
docker compose -f active-directory.yml up -d     # el directorio de pruebas (AD/LDAP)
docker compose -f keycloak.yml up -d             # Keycloak
```

| Service | What it is | What it is tested with |
| --- | --- | --- |
| **OpenLDAP** | A directory with **three** test people and its configuration **in the repository** (`config/ldap/01-personas.ldif`) | The AD path: search, validate and read attributes. **It is done and verified** |
| **Keycloak** | With **a realm imported from the repository** (`config/keycloak/realm-catalina-support.json`) and its three test people | The OIDC path from start to finish. **It is done and verified** |

The three people in `config/ldap/01-personas.ldif` are set on purpose so that each case in
section 5.4 can be triggered by hand: one that **has no account** in the application (automatic
creation), another that is not used in any test and a third **whose email is that of a local account**
(the link).

**A warning to keep in mind**: OpenLDAP **is not Active Directory**. It is spoken to in the
same way —search, connect, validate—, so the code is tested for real, but **what is specific to AD is not covered**:
neither its cross-domain references, nor its particular attributes, nor its password rules. When there is a real AD to point to, it will be tested against it
and whatever comes up will be noted.

**The directory container is not stopped and started**: its startup scripts are not idempotent
—when started again it tries to copy the LDIFs into a folder that is no longer there and dies—, so a
`stop` leaves the AD path down with the container in `Exited`. It is brought up again by recreating it, and
since the directory stores nothing in a volume, **it returns to the LDIF content from the repository**:

```bash
docker compose -f active-directory.yml up -d --force-recreate ldap
```

That is convenient for testing (the directory is always as the repository says) and it must be kept
in mind: **whatever is changed by hand inside the container is lost when it is recreated**. The interface tests for the AD path skip themselves if the directory does not respond
(`tests/e2e/specs/directorio.spec.ts`), so the entire suite can be run without it.

**And the tests for the two paths set the method and put it back** (amendment of 2026-09-25): since
the installation signs in through only one (section 5.0), testing the AD or Keycloak one requires choosing it
first. Each suite sets it at the start and **leaves it at `local` at the end, whatever happens**, with the
factory account, which always signs in. Whether the service responds is asked **through the «Test
connection» button in Settings**, which is what a person asks and does not depend on the method set.

**The three people in the Keycloak realm** are set for the same reasons as those in the
directory: one that **has no account** here (automatic creation), another that signs in again and a
third **whose email is that of a local account** (the link). The realm also brings the **client**
with which the application signs in: confidential —with a secret, because the exchange is done by the backend—, with
the code flow on, **the direct password one off** and the only callback address that
is admitted.

**How the browser reaches Keycloak.** Locally it is published only on `127.0.0.1:11006`,
with issuer `http://127.0.0.1:11006/sso/realms/catalina-support`. The backend connects through
`internalIssuer = http://keycloak:8080/sso/realms/catalina-support`; the advertised issuer keeps
the public identity. Authorization is public and only backend traffic uses the internal
address. The client allows exact callbacks to `http://127.0.0.1:11001/api/auth/keycloak/callback`
and `http://frontend.localhost:11001/api/auth/keycloak/callback` for the isolated suite, without wildcards.
No nginx is needed. `KEYCLOAK_PUBLIC_URL` allows configuring the advertised origin.

The `directory` profile of `tests.yml` brings up its own LDAP and Keycloak, without publishing ports
or using the user's services. There the browser and backend reach `http://keycloak:8080/sso`,
so the internal address can remain empty. The three realm users are kept.

**One more query, which is not sign-in but lives here**: «do you still know this person?». The
`users` module declares it, which is the one that uses it —so as not to return access to an AD account that is no longer
in the directory—, and the directory client of this module fulfills it with **the same service
account and the same filter** with which sign-in is done: the question «are they still there?» has to be answered
the same as «can they sign in?» (`docs/modules/users.md`, section 5, point 4). **The Keycloak path does not offer it**:
asking it would require a service account in the realm, and that is why manual reactivation
of those accounts is rejected by saying that they return on their own when signing in.

## 12. What was decided when reviewing this document

| # | Decision | It was left like this |
| --- | --- | --- |
| 1 | **The test directories** | OpenLDAP and Keycloak, each in its file (`active-directory.yml` and `keycloak.yml`), optional and with its own command. **Changed on 2026-10-01**: before they were services of `dev.yml` behind the `auth` profile, which disappears |
| 2 | **The Keycloak return** | It is received by **the backend**, which delivers our token to the frontend |
| 3 | **The account data** | The directory **updates first name, last names and email** on every sign-in |
| 4 | **Failed attempts** | **No account lock**: nginx limits and each failure goes to the log |
| 5 | **Access audit** | **Only in the logs**, with no table |
| 6 | **Links** | Requesting a new one **deletes the previous ones** |
| 7 | **Change notice** | Yes: when a password changes, its owner receives an email |
| 8 | **The session token** | Stateless, 10 hours, `sub` with the identifier, the account is read on every request |
| 9 | **The transport** | `Authorization` header, token in `localStorage`, no cookies |
| 10 | **The cost of `bcrypt`** | 12 |
| 11 | **The token mechanics** | It lives in `shared/auth`: the middleware cannot import from `modules` |
| 12 | **`auth.accountInactive`** | Only for local accounts: directory accounts have no password to prove |
| 13 | **The Keycloak token** | It returns **in the URL fragment**, and the frontend erases it when storing it |
| 14 | **The attempt log** | Account, result, IP and date. **Never the password**, not even the wrong one |
| 15 | **The session token keys** | `auth.session.invalid` and `auth.session.expired`, apart from the link ones, which stay as they were |
| 16 | **The code of each error** | 401 in all except `auth.directory.unavailable`, which is 503 |
| 17 | **The 403 key** | `auth.forbidden`, a single one for all modules, checked in `shared/authz` |
| 18 | **Failure reading the account** | It is not a 401: 503 with `error interno`, because the failure is the server's |
| 19 | **The input field** | Only one, `email`, which also accepts `admin` for the factory account |
| 20 | **What sign-in returns** | `token`, `expiresAt` and a `user` object; `me` returns the same `user` object, so as not to have two ways of describing someone |
| 21 | **The link route** | `/set-password`. **Owner's correction**: I proposed the route in Spanish and he decided that **frontend routes always go in English**, even if the product language is Spanish |
| 22 | **The link token** | It travels in the **fragment** (`#token=…`), which the browser does not send to the server, and the screen erases it from the address when using it |
| 23 | **A single screen for both links** | Creation and recovery arrive at `/set-password`, and the expiry is written by `auth` in the global language |
| 24 | **The six screens** | The four with routes plus **no permission** (`/forbidden`) and **server down** (a notice on top, with no route) |
| 25 | **The startup language** | **Owner's correction**: I proposed Spanish by default and he decided that **the browser's** is taken and, if it is neither `es` nor `en`, sign-in happens **in English**, with the switch visible |
| 26 | **Where the texts live** | `core/i18n/es.ts` and `core/i18n/en.ts` with a shared type: **a key that is missing does not compile**, and the error texts go in the same dictionary |
| 27 | **How the screen knows what paths exist** | **A public endpoint**, `GET /api/auth/methods`, which says whether this installation has AD and Keycloak. The Keycloak button is shown **only if its path exists**: offering a button that leads to an error is worse than not offering it |
| 28 | **Without a configured directory there is no AD path** | **Amended on 2026-09-25**: it is no longer «without `LDAP_HOST`», which was an environment variable, but **without a directory configured in the database** —and the method also has to be `ad`—. There is still no separate switch: an installation without a directory does not have to turn anything off |
| 29 | **The AD path first, and Keycloak afterwards** | They are implemented separately, each with its test container and its verification: this way a failure is attributed to one path and not two |
| 30 | **The Keycloak button, on the sign-in screen** | The Keycloak return delivers the token **to the same screen where everything started** (`/login#token=…`), which stores it and erases the fragment (decision 13) |
| 31 | **Signing in without manual creation** | **Owner's correction** (2026-09-25), when bringing him the link as a question: «with AD or with Keycloak I would like users to be able to sign in directly, without the manual account creation process by an administrator or technician». Someone who is in the directory signs in and their account is born, linked or brought up to date on its own. What an administrator still assigns is the **role**: the directory does not say it |
| 32 | **When linking is triggered** | When the local password of an account **is not valid** and the directory knows that email: it is linked on that same sign-in and the local password stops working. That is what makes the row «a local account exists with that email» happen without anyone notifying Support |
| 33 | **The sign-in screen changes little because of AD** | **Amended on 2026-09-25**: the form **is the same** and there is no button of its own, but it now does carry a notice —«this installation signs in with the organization's accounts»— and **the «I forgot my password» link is removed**, because with the method set to AD that password is not from here and the link would send an email that is useless for signing in. **With Keycloak the screen does change completely**: there is no form |
| 34 | **`methods` does not probe the directory** | **Amended on 2026-09-25**: it says the method that is set, not anyone's health. Probing the directory on every screen load would cost one connection per visit and a slow directory would delay sign-in; the honest notice arrives on sign-in, which is when it matters, and **the connection test is in Settings**, which is where it is configured |
| 36 | **The issuer is the one the browser sees** | In development. Historical decision, superseded by the local amendment of 2026-10-03: Keycloak was published on a path of the same domain as the application (`/sso/`) and the backend uses that same issuer. The gain is that there is **a single address** and that `iss` always matches; the loss is one network round trip from the backend to its own domain on each exchange, which is irrelevant next to having two addresses that could disagree |
| 37 | **The `state` is stored nowhere** | It is a random value with its date, signed with `TOKEN_SECRET`: what prevents an old return from being valid is its five-minute expiry, and the date is inside what is signed. No cookies and no table of pending states, which is the same decision as with the session |
| 38 | **The return leaves the person on the sign-in screen, also when it fails** | With the token (`#token=…`) or with the failure key (`#error=…`). It is a browser navigation: even what goes wrong has to end up somewhere where what happened can be explained. The screen reads the fragment and **erases it with `replaceState`** |
| 39 | **Four new error keys** | `auth.oidc.notConfigured` (404: this installation does not have that path), `auth.oidc.rejected` (the return is not valid: expired code, used, from another installation, or someone who cancelled), `auth.oidc.unavailable` (503: Keycloak does not respond) and `auth.oidc.noEmail` (the directory account does not bring an email, and without an email there is no account here). `auth.oidc.state`, which already existed, is the fifth |
| 40 | **PKCE, proposed and not done** | With a confidential client and the exchange in the backend, the secret already protects the exchange; PKCE would additionally protect against interception of the code. **It remains proposed**: if it is done, the `code_verifier` would travel inside the signed `state`, which already exists. It has not been done because it was not in the approved document and it changes nothing that is visible |
| 44 | **With the method set to AD, recovering the password is not offered** | **Confirmed by the owner on 2026-09-25**: with the method set to AD the link disappears, because that password is not from here and an email with a link to set it would be a lie —signing in with it would not work—. With the local method it stays as it was, and with the Keycloak one it is not offered either |
| 45 | **The factory account door is discreet** | **Confirmed by the owner on 2026-09-25**: a text link in the footer of the sign-in screen, which brings up the form when the method does not have it. It is not offered to those who sign in every day, but it is where it has to be looked for, and it is what prevents choosing the wrong method from closing the installation |
| 41 | **One sign-in method at a time** | **Owner's decision, 2026-09-25**: the three paths **do not coexist**; the one that is set is the only one served, and it can be changed whenever desired. **The factory account always signs in**, on its own and with its discreet door on the sign-in screen when the method has no form. Consequence stated and accepted: with the method set to `ad` or `keycloak` **local accounts do not get in**, and that amends section 5 of `docs/usuarios-y-permisos.md`, which described the three paths coexisting |
| 42 | **The directory and the realm are read on every attempt** | **A consequence of the previous decision, and its price**: changing the method or the directory takes effect on the next attempt, without restarting anything, in exchange for three primary-key reads on every sign-in. The Keycloak client is built per attempt with the current configuration; the document is remembered only during that attempt |
| 43 | **An account belongs to one path, and not two** | A Keycloak account **does not sign in through the directory** even if the method set is the AD one, and an AD account does not sign in with a local password with the method set to `local`. What the method changes is **where sign-in happens**, not the nature of each account |
| 35 | **An account disabled here returns through the directory** | It also applies to linking: if the local account was disabled, it is reactivated when linked. It is rule 6, and it is what makes the notice `users.directoryMayReturn` true. With this **the doubt that remained open** in `docs/modules/users.md` about whether that notice could be exercised **is closed**: it is exercised, and there is an interface test |

## 13. What this document enables

With `docs/modules/auth.md` approved, **`auth` and `users` can now be implemented**, and it is convenient to
do them together: there can be no accounts without being able to sign in, nor sign-in without accounts. The order I
propose is the one that makes each step testable on its own:

0. **`mail`**, which is what `auth` depends on to send the creation email: without it, creation cannot
   be completed. It goes first.
1. `shared/auth` (signing and validation) and the middleware, with their tests. **Done.**
2. The `password_tokens` table and the accounts, in `v1.0.0.sql`. **Done.**
3. The complete local path: creation, link, sign-in, change and reset. **Done and verified.**
4. The frontend shell: service, interceptor, guard and the six screens. **Done.**
5. The test containers and the AD path. **Done and verified** (2026-09-25).
6. The Keycloak path. **Done and verified** (2026-09-25). **With this, this document has
   nothing pending.**

## 14. Authentication without a personal language

The shared identity, the sign-in responses and `me`, and the session state stop including
`language`. The frontend gets the language from the public brand and does not change it again when receiving an
account. The sign-in screen removes its selector; it is painted directly in the global language.

Authenticated responses include the `X-Catalina-Language` and
`X-Catalina-Settings-Version` headers. The existing interceptor observes them and, when the version changes,
updates the global language signal. This does not change the token, does not close the session and does not add polling:
an open tab adopts the change on its next request.

The account creation, forgot-password and password-change emails consult the global language when sent. The expiry
text is formed with that same language. They do not change the token, lifetime, passwords, factory
account, AD, Keycloak, permissions or routes. The proposal was confirmed by the owner on
2026-10-06 and in the review he chose to update open sessions with the language and
version headers, without closing the session. It was implemented and verified on 2026-10-07.
## 15. Implementation: missing password link

`FindByHash` limits the query to one row and uses `RowsAffected` to distinguish a token that does not
exist. It keeps the same domain error for unknown, replaced or already retired links;
it does not change expiry, single use, public responses or any of the three sign-in methods.

## 16. Identity with cumulative permissions

The token still carries only `sub`. When resolving each request, the ordinary identity incorporates
`role` and `isAdmin`: the first is User, Support or Development and the second adds the
administrative functions. The factory identity keeps `factory = true` and `isAdmin = true`, but has no
operational role.

Management operations ask for the administrative capability and ticket operations for the operational
role. `GET /api/auth/me` publishes both dimensions and the pending administrative notices. The
person can close them through the endpoint owned by `users`.

No token renewal: the grant or withdrawal takes effect on the next request. The three sign-in methods
remain intact and AD or Keycloak synchronization keeps `isAdmin` as local data.
The authentication, identity and authorization tests passed on 2026-10-18.

## 17. Real Active Directory, username or email and `objectGUID`

The finding reproduced on 2026-10-20 is not a rejected password. Active Directory found
`jcestevez` and accepted his credentials; then it returned `objectGUID` as sixteen bytes and the
backend converted them directly to `string`. PostgreSQL rejected that string because it was not UTF-8.
Attempts with email did not find the person because the installation was filtering only by
`sAMAccountName`.

AD sign-in accepts **username or email**. The configured base condition still
limits the allowed universe —including the exclusion of disabled accounts— and the application
adds an escaped condition that compares the identifier with the configured user attribute
and with the email attribute. Received text will not be interpolated as an LDAP filter. The exact
sign-in search and the partial administrative search will share the same composition and the same
attributes; only exact or partial matching will change.

When `attr_id` is `objectGUID`, the binary value will be read and converted respecting Active
Directory's byte order to the canonical lowercase GUID
`xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`. Any other identifier attribute will keep its text,
which maintains OpenLDAP and custom configurations. An empty value or an `objectGUID` that does not
have sixteen bytes will not be persisted as garbage: the DN will be used as a fallback, as it is today.

Normalization is applied before searching by `external_id`, linking or creating. Tests cover
a real binary GUID, a textual identifier, short username, email, characters that need
LDAP escaping, a disabled account and a rejected password. No log will include passwords or the
filter with unmasked data.

## 18. Consecutive password-setting links

> **Verification proposal of 2026-10-23.** During the walkthrough, reusing a tab for
> consecutive links seemed to keep the previous result. The owner chose 4A: first
> it is reproduced in isolation and only fixed if it fails. He closed the behavior with 4A and the matrix
> with 5A. No decisions remain open; explicit approval is missing.

The same tab can navigate from one valid `/set-password#token=…` link to a different one. Each
change must start a new load, clear the previous success or error, validate the new token and
erase the fragment after adopting it. The state of one link is never reused for another.

The reproduction uses two accounts and two valid consecutive links on the same page. If the
behavior is already correct, only the regression is kept and it is recorded as not
reproduced. If it fails, only the lifecycle of that screen is fixed; tokens,
expiry, API, passwords and emails do not change.
