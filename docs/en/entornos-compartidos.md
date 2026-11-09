# Shared environments

[Español](../es/entornos-compartidos.md) · **English**

> **Status:** as-built
> **Last updated:** 2026-10-21

## 1. Scope

Catalina Support can present and protect a public or shared installation without adding its
topology or operational values to the public repository. All these capabilities are opt-in via
environment variables; if they are not configured, the behavior of a normal installation does not
change.

This document covers only reusable product mechanisms. Domains, specific accounts, Compose, nginx,
calendars, slots, secrets, and procedures of an installation belong to its own operational
repository.

## 2. Public presentation

`SHARED_ENVIRONMENT=true` marks the installation as shared. The public branding delivers the notice
and the reset reference configured through `SHARED_ENVIRONMENT_NOTICE` and
`SHARED_ENVIRONMENT_RESET`. The frontend shows the notice before login and persistently in the
authenticated shell. In this mode it hides the password recovery and change forms.

## 3. Shared accounts

`PROTECTED_EMAIL_SUFFIX` protects accounts whose email ends with that suffix. The backend does not
issue recoveries for them and rejects resetting or changing their password. The factory account is
excluded because it does not belong to the database. Empty disables the rule.

## 4. Quotas

The configurable limits are:

- `LIMIT_LOGIN_PER_MINUTE`, per IP;
- `LIMIT_AI_PER_MINUTE` and `LIMIT_AI_PER_DAY`, per account;
- `LIMIT_COMMENTS_PER_MINUTE`, per account;
- `LIMIT_ATTACHMENT_BYTES_PER_DAY`, per IP and declared size;
- `LIMIT_EXPORTS_PER_HOUR`, per IP.

Zero or absence disables each limit. Counters live in memory and a denial responds 429 with
`Retry-After`; this mechanism is designed for a single replica. A public proxy remains responsible
for limiting request size and delivering the real trusted IP to the backend.

## 5. Attachment scanning

`ANTIVIRUS_ADDRESS` connects clamd's INSTREAM protocol. The service reads the full file before
writing it to storage and only accepts the clean response. A virus, an unknown response, a read
error, or the analyzer going down reject the upload; without a configured address, the product's
normal flow is preserved.

## 6. Verification

Unit tests cover quotas enabled and disabled, protected suffix matching, and clean, infected, and
indeterminate ClamAV responses. The backend suite, `go vet`, the frontend's 253 tests, and the
production build pass with these capabilities.
