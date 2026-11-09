# Open project, contributions and trademarks

[Español](../es/proyecto-abierto.md) · **English**

> **Status:** as-built
> **Last updated:** 2026-10-22
>
> **Tenth amendment approved, implemented and verified on 2026-10-22:** section 18 coordinates the English demo and the new
> bilingual evidence of the official site.
>
> **Ninth amendment implemented on 2026-10-20: repository migration.** The source code was
> moved to the `Catalina-Support` organization, under
> `Catalina-Support/catalina-support`. Public and operational links were updated without changing
> the license, the contribution rules or the publication model.
>
> **Eighth amendment implemented and verified on 2026-10-17: public front page, site and
> sustainability.** Section 17 defines the bilingual front page of the repository, the link to the
> official site, the public demonstration and the invitation to support the project on Ko-fi.
> `README.md`, `README.es.md` and two complementary real screenshots were published. The demo, the
> external site, production and the private restore script were not modified.
>
> **Seventh amendment implemented and verified on 2026-10-13: rules for a single maintainer.** Section 16 separates `main` and `dev`, keeps PRs and the technical protections, and sets zero approvals while only one maintainer exists. The owner chose 1A, 2A and 3A, closed the review with 1A–4A and explicitly approved the document. The rulesets were read again and PR `#7` was merged normally, without administrative bypass.
>
> **Open publication closed on 2026-10-12.** The accessible history of `main`, `dev` and
> `cla-signatures` expresses AGPL-3.0-only; PR `#2` published the current tree on `main` through
> squash; and the presentation, the community channels, the private reports and the branch rules
> were applied and read again from GitHub. Production, the domain and the `v1.0.0` tag
> remain deferred.
>
> **Sixth amendment approved on 2026-10-11.** The repository was made public before the planned
> closing and `main`, still behind, showed the previous MIT license as the main license. Sections
> 7, 8 and 15 record the finding and the correction chosen in 56A–65A. The proposal remains
> explicitly approved; production, domain and version remain out of scope.
>
> **Correction approved by the owner on 2026-10-11.** The main branch must remain
> `main` and `dev` must keep its development role. Decision 66A replaces the temporary branch
> change of 56A and 65A. Section 15 incorporates the corrected and explicitly approved sequence.
>
> **Implementation finding of 2026-10-11.** The current README already contained a notice about the
> copies received under MIT, contrary to 59C. It is removed from the README before rewriting; the
> traceability remains only in section 15, in accordance with 64A.
>
> **Fifth amendment approved on 2026-10-11.** Section 14 specifies the community configuration
> and the branch rules that are missing on GitHub. The owner closed 40A–54A. Visibility continues
> out of scope. The owner explicitly approved the amendment.
>
> **Fourth amendment approved and implemented on 2026-10-10.** The real CLA test revealed that the
> action's fixed schema keeps more public identifiers than the approved description. The
> owner chose 38A: accept and document the real format. Section 13 delimits that data
> and records the evidence. The owner explicitly approved the amendment.
>
> **Second amendment approved and implemented on 2026-10-10.** The direct dependency
> `github.com/juanky-estevez/go-logs v1.0.0` did not publish a license. The owner chose 27A,
> approved the amendment and authorized the specific publication: `go-logs v1.0.1` contains MIT and Catalina
> Support already resolves that version. Section 11 keeps the evidence.
>
> **Third amendment approved and implemented on 2026-10-10.** The frontend generated third-party
> notices when compiling, but the repository did not keep an equivalent and verifiable
> inventory for the backend. The owner closed 28A–37A; section 12 records the versioned file, the
> full texts, their automatic check and the evidence.
>
> **First amendment approved and implemented on 2026-10-10.** The audit found six npm warnings:
> two critical and four high. The owner chose and approved 26A; Angular was updated in a
> coordinated way to 22.2.2, the lockfile was regenerated and section 10 keeps the evidence.

## 1. Objective

Catalina Support will be published as an open source project so that anyone can
examine it, install it, test it, modify it and contribute. The community edition will remain
the complete product: it will have no artificial limits on users, tickets or features.

Publication primarily seeks that the project can be evaluated as a product and as engineering
work. It also keeps future sustainability paths: hosted service, managed installation and
support, consulting and commercial licenses for organizations that cannot use
AGPL. None of those offerings is part of this work.

The repository lives in the `Catalina-Support` organization, under
`Catalina-Support/catalina-support`. The move from the personal account does not change these rules.

## 2. Code license

The code published from this change onward will use **GNU Affero General Public License,
version 3 only (`AGPL-3.0-only`)**. The root `LICENSE` file, the README statement and the
repository metadata must match. New files that need an SPDX identifier
will use `AGPL-3.0-only`.

The change does not revoke permissions already granted: whoever has received a copy under MIT may
continue to use that copy under MIT. The new license applies to the version published after the
change and to its subsequent contributions.

The initial copyright holder will be **Juan C. Estevez**. The AGPL allows commercial use, modification and
redistribution; it requires keeping its notices and offering the corresponding source code, also
when a modified version is used to provide the program over a network. The README will explain
this consequence in plain language and will link the full text of the license, without replacing it.

## 3. Contributions and CLA

The repository will include a complete community profile:

- `CONTRIBUTING.md`, with the documentation-before-code cycle, environment setup, tests,
  commit format and review process;
- `CODE_OF_CONDUCT.md`, based on Contributor Covenant and with a private contact channel;
- `SECURITY.md`, with supported versions and the use of GitHub Private Vulnerability Reporting;
- `SUPPORT.md`, which separates usage help, proposals and security reports;
- issue, feature request and pull request templates;
- `CLA.md` and the minimal record of acceptances;
- `TRADEMARKS.md`, with the policy described in section 4.

General discussions will live in GitHub Discussions; specific defects and work, in
Issues. Vulnerabilities will not be published as an issue: they will be sent through Private Vulnerability
Reporting.

The CLA will follow the Harmony license model: each contributor keeps their rights and grants
Juan C. Estevez sufficient permissions to distribute the contribution under AGPL and also under
other licenses, including commercial or proprietary licenses. It is not a transfer of ownership.
It will be required for code, graphic assets and substantial documentation; not for opening
issues or for trivial typo fixes.

Each pull request subject to the CLA will have an automatic check. A GitHub Action pinned to a
commit hash, not to a mutable tag, will record in the repository itself the public fields
described in section 13. The record will be public. The contribution will not be merged while
acceptance is missing. Before adopting a commercial license, independent legal review of the
final text and of the rights chain is recommended.

## 4. Trademarks

`TRADEMARKS.md` will separate the permissions of the code license from the permissions on the
Catalina Support name and logos.

An official unmodified build may keep the name and logos. An official installation that changes its
name, logo or color from Settings may also keep the attribution to Catalina Support: that
customization, foreseen by the product, does not turn it into a
fork.

A fork or a modified build may nominatively explain that it is "based on
Catalina Support", but must use a different name and visual identity and may not
suggest that it is official, sponsored or approved. Any additional use requires written
authorization. The policy will allow truthful references in articles, comparisons and links.

## 5. Visible attribution and corresponding source

The entry screen shows **"Source code and license"** as a discreet and accessible link,
separate from the main actions. It leads to the repository and explains `AGPL-3.0-only`. Since
2026-10-15 it is no longer repeated at the bottom of the side menu, by textual correction of the owner.

The public brand already delivered by the backend will add `license` and `sourceUrl`, so that both views
use the same source. For a published version, `sourceUrl` will point to the tag that corresponds
exactly to `version`. For a development build it will point to the repository and identify
the build as development. The frontend will not duplicate those values in constants.

The URL cannot be changed from Settings: it identifies the distributed code, not the installation.
A derived build must modify its build metadata to link its own corresponding
source. If the data is missing for compatibility with an older backend, the interface will not
show an invented link.

## 6. Audit before publishing

The repository will remain private while a specific result is prepared and reviewed. The audit
will cover the current tree and the entire Git history, searching for secrets, credentials, personal
data, large files and references to real infrastructure. It will also check dependencies,
third-party notices, generated files, container permissions and that an external person can
follow the README from an empty installation.

Two findings are already known:

1. eighteen commits contain `<personal email removed>` as the author email;
2. versioned documents and configurations contain domains and paths from the real deployment.

Before publishing, the affected authors will be rewritten with GitHub's `noreply` email. This
operation will change all descendant hashes; it will be done once, after finishing the
changes and saving a recoverable private reference. The resulting history will be verified before
removing that reference.

The infrastructure's own domains, paths and data will be replaced by examples and variables.
The real values will remain only in unversioned configuration. The documentation will continue
to be reproducible and will clearly distinguish examples from values that the operator must choose.

**Local evidence of 2026-10-10.** Gitleaks reviewed separately the exact set of files that
will be published and the previous 19 commits. The only four tree warnings were two keys
declared as development, the same key in the isolated test environment and an unsigned JWT that
tests the rejection of `alg=none`; they were allowed by exact path and content. Two historical
fingerprints of those same values are recorded by commit, file, rule and line. The production
secrets and tokens found in the working directory live exclusively in
`config/env/prod.env` and `tmp/`, both ignored and absent from the index. With those narrow exceptions,
the publishable set and the history ended without findings. The largest historical blob is a logo
of 603,822 bytes; there are no unexpected large files.

The private backup
`/tmp/catalina-support-before-public-history.bundle` was created and verified, SHA-256
`ca61d18cac089c71faf576040df4e39ec0533fde9629b61888ef3404156a9`. Afterwards `main` and `dev` were rewritten a
single time: all personal authors and committers now use the public
`noreply` email, the real domains became `example.com` and the real paths became
`/srv/catalina-support`. The resulting local branches are `main` at `b9df02f` and `dev` at
`1260932` before incorporating the Gitleaks fingerprint update. The remote publication of
these branches and the removal of the old local references remain for after the
final verification of the rewritten history.

If an additional secret or sensitive datum appears, publication will stop and the finding will be
presented before cleaning or rewriting the history. It will not be assumed that deleting the current file
removes it from Git.

## 7. GitHub preparation and publication

Before the visibility change, the following will be prepared and verified:

1. license, notices, community documents and templates;
2. automatic CLA and branch protection to require its checks together with the tests;
3. Issues, Discussions and Private Vulnerability Reporting;
4. repository description, topics and link;
5. clean audit of the tree and of the rewritten history;
6. local installation from scratch following only the README;
7. source and license links on the entry screen and the menu;
8. a final comparison between the local repository and what will become public.

Publication will be of a **development version, prior to 1.0**. No candidate tag or `v1.0.0` is created now.
The README will state this without presenting stability that does not yet exist.

The visibility change was planned as the last step and required explicit approval of the
prepared result. The repository was made public before that closing; section 15 records
the deviation and its correction. Publication does not authorize deploying production, opening the domain,
publishing a version, accepting donations or offering a commercial plan.

## 8. Verification and closing

To move this document to `as-built`, the following must be recorded:

- the review of the license text, CLA, trademarks and community files;
- the result of scanning the tree and all history revisions;
- the confirmation that the personal email and real infrastructure no longer remain;
- the test of the CLA flow in a trial pull request;
- the backend and frontend tests affected by `license` and `sourceUrl`;
- the PC and mobile walkthrough of the link on the entry screen and the menu;
- the clean installation run from the README;
- `git diff --check` and the final state of the documents.

Visibility was to remain private during these verifications. As it was opened earlier, the
priority becomes correcting the branch and the license that GitHub presents, verifying the public
branches again and completing the community configuration. The closing will keep the inventory of changes,
findings and evidence.

## 9. Decisions and review

The owner chose 1A–5A: AGPL-3.0-only, separate trademark policy, lightweight CLA, audit before
publishing and complete community edition. He chose 6A–10A: initial personal ownership, CLA
verifiable per pull request, scope of substantial contributions, private security reporting and
stopping on sensitive findings. He chose 11A–15A: visible link, nominative trademark use, personal
repository, complete community tools and final approval before changing visibility.

In the review he chose 16A–20A: CLA without transfer and with relicensing permission, local record pinned by
hash, source corresponding to the version, complete community set and publication prior to
1.0. Finally he chose 21A–25A: rewriting the personal email, replacing real infrastructure,
minimal public CLA record, metadata served by the backend and attribution kept in official
customizations. No open decisions remain.

## 10. Implemented amendment: vulnerable Angular dependencies

The tree audit found six warnings in the current lockfile: two critical and four high. The
`@angular/router` warning affects server-side rendering, which Catalina Support does not use, but the
version remains within the vulnerable range. The critical warnings come through the build
tool and `piscina`; the others affect transitive CLI tools and source maps.
All have a fix available.

The frontend will update the Angular packages in a coordinated way to **22.2.2** and regenerate
`package-lock.json` with the package manager declared by the project. No major update will be
applied, no warnings will be disabled and no audit exceptions will be added. The lockfile diff
will be reviewed to confirm that it only resolves the new family and its transitives.

The closing criteria will be:

1. `npm audit` without high or critical vulnerabilities;
2. the boundaries, the unit tests and the complete frontend build passing;
3. the affected Playwright walkthrough on PC and mobile passing;
4. no functional changes beyond compatibility with Angular 22.2.2;
5. recording of the real result in this document before returning to `as-built`.

The owner chose 26A. The update is integrated into the audit prior to publishing; it does not change
the license, the trademark policy, the CLA or the separate decision on the visibility of the
repository. No open decisions remain.

**Evidence of 2026-10-10.** All direct Angular packages were pinned at 22.2.2 and the
lockfile was regenerated with the npm declared by the project. `npm audit --audit-level=high` ended
with zero vulnerabilities; the 227 frontend unit tests, the complete build and the 12
affected Playwright cases on PC and mobile passed. The disposable environment was removed at the end.

## 11. Implemented amendment: `go-logs` license

The dependency audit found that `github.com/juanky-estevez/go-logs v1.0.0`, used by the
backend for its four logging levels, contains no license file or statement. Its
published tree includes code, tests and generic configuration; it contains no credentials. Since no
explicit redistribution permission exists, Catalina Support will not be published
depending on that version.

The `go-logs` repository, also owned by Juan C. Estevez, will include an **MIT** license and
a 2026 copyright notice. Its API and behavior will not change. After verifying its
tests, the **`v1.0.1`** tag will be created and published, whose content will include the license. `v1.0.0` is not
retagged and its history is not rewritten.

Catalina Support will update `go.mod` and `go.sum` to `github.com/juanky-estevez/go-logs v1.0.1`,
download the module from its origin and verify that the resolved artifact contains `LICENSE`. The
complete backend suite and `go vet ./...` must pass without behavior or log format
changes.

The `go-logs` tag and push are external publications and irreversible in practice. The
specific commit of the external repository will be presented first and final approval will be requested to
publish them. If the repository has its own contribution instructions, they will be applied before
preparing that commit.

The owner chose 27A. This amendment only fixes the license chain; it does not incorporate the
module's code into Catalina Support or change its AGPL policy. No open decisions remain.

**Evidence of 2026-10-10.** The published commit of `go-logs` is
`6c13674c9f9c5794302fcae002409fc30e51836d`, with author and committer at the public `noreply` email;
`main` and the annotated tag `v1.0.1` point to that change. The module passed `go test ./...` and
`go vet ./...` in its container. Catalina Support updated `go.mod` and `go.sum` to `v1.0.1`, verified
the MIT `LICENSE` inside the downloaded artifact and passed the whole backend suite and `go vet ./...`
again.

## 12. Implemented amendment: third-party notices

The audit confirmed that the Angular build produces `3rdpartylicenses.txt`, while the
backend does not generate or keep an equivalent inventory. The reviewed direct dependencies have
licenses compatible with the planned publication; `go-logs`, which was the exception without explicit
permission, was fixed in section 11. The remaining gap is traceability: a person who
examines the repository does not have a single and reproducible list of the third-party components that are
distributed.

The repository will include **`THIRD_PARTY_NOTICES.md`**, generated deterministically and
versioned. It will include production dependencies and the tools whose components are
distributed with the application; it will exclude dependencies used only for development or
tests. The set will be calculated with the locked versions and the graph reached by the
production build, not only with the directly declared dependencies. Each entry in the
index will identify the component, the resolved version, the license and a stable link to its origin.
The notices and full texts that the licenses require to be kept will live, also versioned,
in **`third_party_licenses/`**; the compact index will link each component to its local text.

A repository script will regenerate the file from the locked dependency sources, without
resolving new versions. It will regenerate both the index and `third_party_licenses/`. CI will run
the same generator in check mode and will fail if any of the versioned files
differs, if an included dependency does not declare a license or if a required notice cannot be
kept. Thus, a change in `go.mod`, `go.sum`, `package.json` or `package-lock.json` cannot be merged
leaving the inventory behind.

The frontend build will continue to produce its own file inside the artifact. The root inventory
does not replace it: it offers review on GitHub and additionally covers the backend. External
images and base images are outside this file: the project currently references or builds them
locally, but does not publish its own images. If it later distributes images, that work must
audit and deliver the notices of their layers.

A missing, unknown or incompatible license will stop both generation and CI. There will be no
exception list that turns that error into a warning. The implementation does not change
third-party licenses, does not incorporate their code into the repository and does not alter the behavior of the
application. `scripts/prod-build.sh` will copy the index and the licenses folder next to the published
backend and frontend artifacts; a deployment prepared by the script will keep
the notices even if separated from the Git clone.

The owner chose 28A–30A: versioned file, production scope and automatic check
in CI. In the first review he chose 31A, 32B, 33B and 34A: real graph over locked versions,
images outside the current inventory, compact index with links and blocking of unknown
licenses. Choice 33B corrected the presentation proposed by the agent. In the final review
he chose 35A–37A: the required texts remain in a separate and versioned folder, generated and
checked together with the index, and both are copied to the artifacts. Thus the index remains compact
without depending on external links to comply with the notices. No open decisions remain.

**Evidence of 2026-10-10.** `scripts/third-party-notices.go` walked the normal
backend build graph and the entries not marked as development in `package-lock.json`: it generated
29 components —19 Go and 10 npm— with version, license, source and local text. A second
run in `--check` mode produced no differences. The negative test added content to a generated
text and the check failed naming exactly that file; afterwards the documented script
regenerated it and it passed again. The generator passed `gofmt` and `go vet`; the scripts passed `bash -n` or
`sh -n`, and `git diff --check` passed. `npm ci` resolved 311 packages from the lockfile with zero
vulnerabilities.

CI runs the check in its own job with Go and Node. `scripts/prod-build.sh` copies the index
and the texts into each artifact after verifying that it exists. That copy was not run against the
real deployment: production remains parked until the closing of 1.0 and `docs/ambientes.md` keeps
the previous finding that prevents presenting the complete walkthrough today as reproducible from scratch.

## 13. Implemented amendment: CLA record and test

The action keeps one entry per acceptance with `name`, the account's public numeric
identifier (`id`), the comment's public identifier (`comment_id`), its date (`created_at`), the
repository's public identifier (`repoId`) and the pull request number (`pullRequestNo`). The
accepted version is expressed through the versioned path `signatures/version1/cla.json`. No
email, civil name, address, company or additional free text is stored.

These identifiers allow linking the acceptance to the specific account, comment, repository
and change even if the user changes their display name. The format belongs to the adopted
action; Catalina Support documents it instead of maintaining its own legal processor. The record
lives in the `cla-signatures` branch, which must remain writable for the automation.

The owner chose 38A after learning of the difference from the minimal record described
before the test. The discarded alternative was replacing or adapting the action to reduce the
schema. No open decisions remain in this amendment.

**Evidence of 2026-10-10.** The temporary PR `#1` exercised the request, the acceptance through the
exact comment, the automatic commit in `cla-signatures` and the `CLA / cla` check; the latter and
the three CI jobs ended green. The repeated attempts during diagnosis left
three equivalent entries; the first valid acceptance was kept and the file returned to
a single entry. The PR was closed without merging, and its branch and temporary file were deleted. The
`dev` branch did not receive the test content. The workflow must exist in the default branch
`main` for GitHub to process comments; the finding was fixed there and a later
synchronization of the PR confirmed the fourth check in green.

## 14. Approved GitHub community configuration

### 14.1 Presentation and channels

The repository will keep `main` as the default branch and will show this description in English:

> Self-hosted, AI-powered help desk for Support and Development teams.

It will use the topics `helpdesk`, `ticketing-system`, `self-hosted`, `artificial-intelligence`, `golang`,
`angular`, `postgresql` and `open-source`. The website field will remain empty while the domain and
production remain deferred.

Issues, Discussions and Private Vulnerability Reporting will be enabled. Issues will receive defects and
specific proposals; Discussions, open questions and ideas; security reports will follow
the private channel linked from `SECURITY.md` and the issue template.

### 14.2 Integration and cleanup

The repository will allow only **squash merge**, require linear history and automatically
delete the source branch after merging a pull request. It will not require signed commits
until accessible guidance and configuration exist for contributors. It will not add
`CODEOWNERS` while only one maintainer exists.

Approvals will be dismissed when new commits arrive. Before merging, the pull request
branch must be up to date with its base and all conversations must be
resolved.

### 14.3 Protection of `dev` and `main`

Both branches will require a pull request, one approval and these four checks:

- `CI / backend`;
- `CI / frontend`;
- `CI / third-party-notices`;
- `CLA / cla`.

They will not allow deletion or force push and will block direct push during normal work. The
owner may use the rule's manual bypass, necessary while he is the only maintainer and
cannot approve his own pull request. The exception does not replace the usual flow: it is reserved
for administrative tasks, recovery or a blockage that cannot be resolved through a PR.

The `cla-signatures` branch will remain outside these rules and writable by GitHub Actions, because
the bot records acceptances there. Protecting it would prevent the flow already verified in PR `#1`.

### 14.4 Decisions, verification and scope

The owner chose 40A–44A: English description, focused topics, empty website, protection of both
branches and owner bypass. In the second review he chose 45A–49A: default `main`, squash
only, approval renewed after changes, up-to-date branch and automatic cleanup. In the final
review he chose 50A–54A: linear history, no mandatory signing or `CODEOWNERS`, four checks on both
branches and direct push blocked except manual bypass. No open decisions remain.

The implementation was verified by reading the effective configuration again, not just trusting the
form. GitHub keeps `main` as default, the description and the eight agreed topics,
the empty website, Issues, Discussions and private vulnerability reports. It only allows squash,
offers to update the branch, automatically deletes source branches and keeps secret scanning and its
protection during push.

The active ruleset `Protect main and dev` (identifier `24942297`) reaches exactly
`refs/heads/main` and `refs/heads/dev`: it prevents deletion and force push, requires linear history, PR, one
approval, renewal after new commits, resolved conversations, up-to-date branch and the checks
`backend`, `frontend`, `third-party-notices` and `cla`. The administrator role can only skip the rule
within a PR. `cla-signatures` remains outside the ruleset and keeps its single signature.

This amendment does not change product files, production, domain, donations or version
tags. The repository is public; that publication was verified when closing section 15.

## 15. Executed correction of the early publication

### 15.1 Public finding

GitHub's public API confirmed `visibility: public` with `main` as the default branch. In the
initial finding that branch was seventeen commits behind `dev`, had a commit of its own from the
CLA workflow and still showed the previous MIT license at the root and in its README. GitHub
therefore identified the public repository as MIT while `dev` already contained AGPL-3.0-only, the
community files, the third-party notices and the current application.

The earlier exposure is not presented as reversible: rewriting the references that GitHub still
offers does not withdraw permissions from a copy that someone might have obtained under MIT. The goal of the
correction was for the current repository and its accessible history to coherently express the
license chosen for Catalina Support.

### 15.2 Main branch and scope of the rewrite

`main` remained the default branch throughout the operation and `dev` kept its development
role. To shorten the MIT exposure, the corrected history of `main` was published first.
Its tree remained the previous stable version until the pull request was merged, but its root
license and its accessible history already expressed AGPL-3.0-only.

The second private bundle
`/tmp/catalina-support-before-license-history-rewrite.bundle` was created and verified, with SHA-256
`b787e3d6c455e1feeb9c51ff3d3a664a9d1115506bc423ff8f4c12b6c543a100`. Afterwards `main`, `dev` and
`cla-signatures` were rewritten in a single operation. In the commits prior to the
approved change, only the following were replaced:

- Catalina Support's own root `LICENSE`;
- the project's MIT statement in the historical root `README.md`.

MIT licenses of dependencies, `go-logs`, notices or third-party texts were not altered. The current
signature of `cla-signatures` was kept with its data; only the identifiers of the
commits that contain it and of its ancestors changed.

### 15.3 Publication and verification

The three branches were published with `--force-with-lease`, comparing the expected remote
identifiers beforehand, in this order: `main`, `dev` and `cla-signatures`. Their rewritten tips were
`6d3946d`, `59dba84` and `2b248f2`; there were no Catalina Support tags to rewrite. Afterwards
all reachable revisions were checked and it was confirmed that:

1. no root `LICENSE` keeps MIT;
2. no root README declares MIT as the Catalina Support license;
3. AGPL-3.0-only remains at the tips of `dev` and `cla-signatures`;
4. third-party licenses and notices did not change;
5. the CLA signature is still present only once;
6. the secret scan of the tree and of the rewritten history is still clean;
7. authorship, domains and paths remain sanitized.

Gitleaks examined the 32 reachable commits and found no secrets; the two fixed
development values kept an exact exception with their new fingerprints. PR `#2` passed backend,
frontend, third-party notices and CLA and was merged via squash as
`afeae10190a87b0df92941f53d356bfc013453f4`. No direct content push to `main` was made; at the
end, `main` and `dev` published the same tree `47df75a40000e9202b0dc9e20e46593c239e8279`,
with `main` still default. GitHub recognizes AGPL-3.0 and no specific notice was added to the README
or to `NOTICE`.

### 15.4 Decisions and scope

The owner chose 56A: temporarily use `dev` as the front page and then synchronize `main` through a
pull request. He chose 57B, 58A, 59C and 60A: rewrite the public history, merge with squash, not
add a separate notice and fix the license first. In the review he chose 61A–65A: the three
branches and only the project's own license, signature kept, second bundle, traceability in this document and
default branch change before rewriting. The owner later corrected that sequence
with 66A: `main` remains the main branch and its corrected history is published first. 66A replaces
56A and 65A; no open decisions remain.

The operation created no tag, did not declare version 1.0.0 finished, did not touch production or the
domain and did not change third-party licenses. The community configuration of section 14 was
applied after `main` and `dev` reached the same tree.

## 16. Implementation: branch rules for a single maintainer

### 16.1 Finding

The single ruleset `Protect main and dev` requires one approval for `main` and `dev`. The only maintainer is also the author of his pull requests, and GitHub prevents him from approving his own change. Therefore PR `#7`, with `backend`, `frontend`, `third-party-notices` and `cla` passing, was blocked. A review comment does not satisfy the rule and the administrative bypass precisely evades the protection that was meant to be kept.

### 16.2 Applied configuration

The shared ruleset was replaced by two active rulesets, one for `refs/heads/main` and another for
`refs/heads/dev`. Both require:

- integration through pull request and only by squash;
- zero mandatory approvals while the repository has a single maintainer with write access;
- `backend`, `frontend`, `third-party-notices` and `cla` passing;
- branch up to date with the base and conversations resolved;
- linear history;
- prohibition of deletion and force push.

The separation does not relax one branch against the other today: it allows `main` to require an
approval again without altering `dev` when there is another person with write access able to review.
Until then, the checks and the PR are the objective and reproducible gate; an impossible
approval is not forced nor is the bypass used as a normal path. `dev` will keep zero approvals
also when there are more collaborators, unless another future proposal justifies changing it.

The administrator role bypass remains limited to pull requests for recovery or administrative tasks. It does not enable direct push and is not part of the usual flow. `cla-signatures` remains outside both rulesets so that the CLA workflow can record signatures.

### 16.3 Verification

The previous configuration was read and temporarily kept before modifying GitHub. The active rulesets are `Protect main` (`24942297`), which reaches only `refs/heads/main`, and `Protect dev` (`24963846`), which reaches only `refs/heads/dev`. Both keep exactly the four checks, zero approvals, strict update, resolved conversations, squash, linear history and protection against deletion and force push. The administrator bypass continues to be limited to pull requests.

The general configuration was read again: `main` continues as the default branch, only squash is enabled and GitHub deletes the source branch after merging. With those rules, PR `#7` moved to `CLEAN` and was merged normally into `dev`, without `--admin`, through squash commit `6ad3a1b1288cd743da7ecc07e1a42ce801938d44`. It was not promoted to `main`: that promotion belongs to the already approved walkthrough of block 8. When a second maintainer with write access is added, a new amendment will be opened to
require at least one approval on `main`; it is not activated automatically because GitHub does not offer a
reliable condition based on the number of collaborators. `dev` will continue to rely on the checks and on
resolved conversations, without mandatory approval.

### 16.4 Scope and decisions

Product code, application permissions, license, CLA, visibility, production,
domain or version do not change. The owner chose 1A, 2A and 3A: zero approvals on both branches during the
single-maintainer stage, separate rulesets and full preservation of checks and protections. In
the review he chose 1A, 2A, 3A and 4A: immediate application, administrative bypass only for emergencies,
a new proposal to require review on `main` when another maintainer arrives and `dev` without mandatory
approval. No open decisions remain.

## 17. Bilingual front page, public demonstration and support for the project

### 17.1 Objective and hierarchy of the front page

`README.md` will be the public front page in English and will link, from the beginning, to a complete
translation in `README.es.md`. Both files will have the same scope and will be updated together when
capabilities or instructions change. Reading will start with a brief and friendly explanation of the
product and three visible actions: **try the demo**, **install Catalina Support** and **support its
development on Ko-fi**.

The front page will invite using the software before going into technical detail. It will explain the
two-tier flow and present AI as a local or configurable aid that summarizes tickets and improves
drafts, without attributing autonomous decisions to it. A real screenshot will show at the same time the main
and the internal ticket, with `Summary` and `Last action` visible. The screenshot will be taken after
restoring the local sample data, using factory logo and colors and without sensitive data.

### 17.2 Official site, demo and shared credentials

The front page will link directly to [catalina-support.com](https://catalina-support.com) as the official
site and to [demo.catalina-support.com](https://demo.catalina-support.com) through **Try the live
demo**. The site is developed in another repository: this work does not create or publish its sources and does not
mark it as future or incomplete. The README remains self-sufficient to evaluate and install the
project even if the external site is unavailable.

The demo is presented as a shared public environment, without guarantee of availability or
permanence. It expressly warns not to enter personal, confidential or
production data and not to upload sensitive files. It publishes in a secondary section the eleven local
seeder accounts and their common password, grouped by User, Support and Development; **it does not publish
or offer the `admin` account**. A short walkthrough proposes creating as User, attending and escalating as
Support and resolving as Development.

The demo data will be restored periodically through a server operation. The README
only explains that changes may disappear: the script, its implementation, its frequency and the
infrastructure that runs it are outside this repository and this proposal.

### 17.3 Installation and linked documentation

The main local path remains an empty installation that goes through `/setup`; the seeder remains
as an optional quick walkthrough. The README will keep requirements, copyable commands, addresses,
local values, accounts, stop, resume and diagnosis, but ordered as a progressive guide.

The **Production deployment** section will be brief but operational. It will require HTTPS, own secrets,
persistence, real mail, backups, manual migration and absolute absence of seeders; it will refer to the
detailed runbook in `docs/ambientes.md`. It will not hide the current limitations or findings of that
runbook nor run production as part of this review.

Any reference to a Markdown or legal file will be written as a link. The files
that exist will be linked: `LICENSE`, `THIRD_PARTY_NOTICES.md`, `SECURITY.md`, `CONTRIBUTING.md`,
`CODE_OF_CONDUCT.md` and the `docs/README.md` index. No missing documents are invented.

### 17.4 Sustainability, scope and decisions

The initial invitation will be discreet and the final section **Support the project** will say: “If Catalina
Support saves you time, consider supporting its development on Ko-fi”, linking to
[ko-fi.com/juanky1990](https://ko-fi.com/juanky1990). It does not condition features, support or license on
a donation.

The owner chose 1A–5A; 6A, 7A, 9A and 10A; corrected 8A to publish the eleven local accounts
except `admin`; chose 11A, 13A–15A and corrected 12A to link the site immediately without marking it
as future; chose 16A–20A; and closed the review with 21A, 23A–25A and 22C. Therefore, the front page will not
show a notice about the pre-1.0 state, although the production guide will keep only
verifiable statements and the current technical cautions.

The documentation implementation was limited to `README.md`, the new `README.es.md`, two versioned
screenshots and the coherence updates of these documents. The real split view does not include
the AI results side panel; to avoid misrepresenting the interface, two complementary evidences
were published: the main/internal view and the detail of the same ticket with `Summary`, `Last
action` and `Improve with AI`. No product code, permissions, seeders, deployment, DNS,
demo, external site, restore script, production or version tag was modified.

## 18. English demo and bilingual site screenshots

The public demo is restored with the English samples and presented in English. It keeps the accounts,
common password, warnings and deletion policy already published; only the language of the
seeded content and the visible names of the accounts change. The public documentation uses `User`, not
`Requester`, to name that role and its accounts.

The official site obtains its six screenshots again from separate real environments: three with
interface and samples in English and three with interface and samples in Spanish. Images are not translated nor
is data fabricated visually. The shots keep the same functional walkthrough and do not show personal
data, secrets or content outside the seeders.

The renewal of the site material does not redesign its composition or change the product. The two
screenshots belonging to the repository READMEs remain outside this amendment unless a textual
reference needs correcting to reflect the new commands or names.

The six shots keep composition, visual size and narrative function; only the localized
content is replaced and, where applicable, its intrinsic dimensions. The owner closed the
review with 1A–5A and no open decisions remain.

The real verification ran both services on disposable databases. Each variant produced eleven
accounts, 25 tickets, 123 history rows, 27 comments, seven internal tickets and five attachments. A new
database adopted `en` with `seed-en` and `es` with `seed-es`; a sealed installation kept its
settings when changing the set. The resolved Compose configuration contains only `seed-en` and
`seed-es`, and the runner tests confirm that it ignores both development files.
