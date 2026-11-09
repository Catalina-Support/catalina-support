# Interface and experience

[Español](../es/interfaz-y-experiencia.md) · **English**

> **State:** as-built
> **Last updated:** 2026-11-05
>
> **Owner's decision, 2026-11-05.** Dates are read in **ISO `AAAA-MM-DD hh:mm`**, in 24-hour
> format and in the same order in Spanish and English. On mobile and in vertical tablet the double
> view switch keeps **Main** and **Internal**; “Both” is not offered because two columns do not fit.
> `/setup` still returns to step 1 on reload and does not promise to remember the position.
>
> **QA correction, 2026-11-03.** The sidebar footer repeated “Administrator” three times
> when the name, the role and the capability matched. The capability is only added when it contributes
> something different from the role.
>
> **Coordinated change approved, implemented and verified on 2026-10-22.** Section 51 reduces the horizontal margin of the
> four inboxes to 16 px, removes their maximum width and fixes Ticket and Action at the ends with inner
> breathing room. The owner closed and approved the review with 1A–4A; the implementation matches
> `modules/tickets.md` and `modules/ai.md`.
>
> **Amendment approved, implemented and verified on 2026-10-20.** Section 50 adds the **Site icon** card and defines
> the immediate update of the favicon in all tabs. The owner closed 1A–14A and
> corrected the fallback: without a custom icon the effective light logo is used. 253 frontend tests passed,
> as did the boundaries and the production build.
>
> **Amendment approved, implemented and verified on 2026-10-20.** Section 49 replaces the two
> general entries with **Tickets**, reserves Users and Classification for Admin and reviews the
> complete visible terminology of the ticket area.
>
> **Amendment approved, implemented and verified on 2026-10-20.** The creation dialog of `/users`
> incorporates Active Directory search by user or email when AD is the current method,
> selectable results and explicit confirmation when linking a local account. Keycloak is
> shown disabled, explaining that the account is created on first access. Decisions 4A–15A, with
> 6B, closed by the owner.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 47 adds the **Public / Private** selector to the
> main editor, visually distinguishes private notes and guarantees their total absence from the
> User's view. The owner closed and approved 1A–26A; the real flow was verified with both roles.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 46 brings `All · Assigned · Watching` to the
> first row of **My tickets** and opens on `All`. The list represents exclusively current assignment or
> watching, separated by level and not expanded by the Admin capability. The owner
> closed 4A and 6A–9A and approved the document.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 45 corrects the active state of the sidebar.
> The finding is that `/tickets` is currently evaluated as a prefix and remains selected at the same time as
> **Main tickets**, **Internal tickets** or **Classification**. The owner chose 1A–4A:
> a single active entry, soft background without shadow, creation and detail belonging to **My tickets**,
> and the user and mail records belonging to **Users** and **Settings**. The resolution
> function is covered by unit tests; the build and the 246 frontend tests pass.
> The browser confirmed a single selection in all entries and in ticket creation and detail, user
> record and mail editor.
>
> **Amendment approved, implemented and verified on 2026-10-19.** Section 43 corrects the
> composition of the operational role and the administrative badge in the user list. The
> owner chose 1A–4A: the role occupies
> the first row and the **Admin / Administrator** badge a second one, left-aligned and with
> a small, constant vertical separation, both in the desktop table and in the cards
> of narrow screens. The 243 frontend tests passed, as did its boundaries and the build. A
> Playwright regression was added for both PC and mobile; the isolated environment did not get to run it because of a
> failure mounting its `node_modules` volume, unrelated to this interface.
>
> **Correction approved, implemented and verified on 2026-10-19.** Section 44 removes the table and the cards
> from the Dashboard, reduces its filters, corrects the transparent colour tokens in the eight themes,
> prevents date overlap and replaces CSV with an aggregated six-sheet XLSX. The owner
> closed 1A–4A, proposed XLSX in 5, chose 6A–14A, corrected the name in 15A and chose 16A. No open
> decisions remain. The 241 frontend tests passed, as did the boundaries and the build; the browser
> confirmed the bars in the eight themes, the four filters and the absence of individual detail.
>
> **Implemented and verified on 2026-10-18.** Section 44 defines the accessible navigation and presentation
> of the new Dashboard.
>
> **Implemented on 2026-10-18.** Section 43 adapts accounts, navigation and notices to the operational role
> plus administrative capability.
>
> **Amendment approved, implemented and verified on 2026-10-17.** Section 42 takes advantage of the height
> available in the operational ticket and user tables: 16 px outer top and bottom, 12 px
> between header and table, a 42 px action row and a stable container down to the bottom edge.
> The finding is that today the shell reserves 32 px at the top and 40 px at the bottom, the inbox adds 24 px between
> blocks and only tickets calculate the remaining height; Users keeps the normal page flow. The
> owner closed the review with 1A–15A; no open decisions remain. The 238 frontend tests passed,
> as did its boundaries and the build. The browser confirmed in the ticket and
> User inboxes the compact header, the use of the available height and the preservation of the mobile cards.
>
> **Amendment approved, implemented and verified on 2026-10-17.** Section 41 reduces and
> recomposes the sidebar: 208 px open and 56 px collapsed, centred vertical brand, larger logo and
> a single collapse control at the top. It also adopts **Classification / Clasificación**
> as the short name of the joint category and tag screen. The starting finding is that
> this section still documented and the implementation still uses 240/64 px and a horizontal
> header that shifts the logo when collapsed. The owner closed the review with 1A, 2A, 3B, 4C,
> 5A–14A; no open decisions remain. The 238 frontend tests passed, as did its boundaries and the
> build. The browser confirmed open and collapsed desktop, mobile drawer, the single control,
> the centred brand and Classification with its two cards.
>
> **Amendment approved, implemented and verified on 2026-10-16.** Section 40 distinguishes the
> visible label of `en espera` according to who is viewing the ticket: the User will read `Waiting for you` /
> `Esperándote`, while Support, Development and the Administrator will read `Waiting` / `Esperando`.
> The difference will be applied throughout the interface, without changing the technical value, the API, the database
> or the emails. The owner chose 1A–4A in the previous review; no open decisions
> remain. 76 focused tests passed, as did the frontend boundaries and the build.
>
> **Amendment approved, implemented and verified on 2026-10-16.** Section 39 turns the table
> of the inboxes into the scrollable area that occupies the remaining height, fixes its
> headers vertically, repairs the Ticket separator and removes the duplicated tooltips. The owner chose
> 1A–13A. As findings, section 37 prohibited fixing the header vertically and section 36.2
> combined the custom tooltip with `title`; both contracts were corrected. Nine focused tests passed,
> as did the boundaries and the build. The browser verified PC, tablet and mobile, scrolling in
> both axes, loading 20→27, internal restoration, a single help only when truncated and the eight themes.
> **Compatibility correction of 2026-10-16:** Action/Acción declares `text-left` directly on
> its `th`; relying on the row's inheritance allowed Firefox to apply its native alignment.
>
> **Amendment approved, implemented and verified on 2026-10-16.** Section 38 replaces the
> visible pagination of the inboxes with progressive loading on scroll, shortens the English user state to `Waiting`
> and left-aligns the Action/Acción header. The owner chose
> 1A–10A and no open decisions remained. Seven focused tests passed and the frontend build
> did too. The browser confirmed 20 initial items and 27 after reaching the end, without pagination,
> on PC, tablet and mobile; it also confirmed `Waiting`, both alignments and the restoration of the
> 27 items and the position when returning from the detail.
>
> **Amendment approved, implemented and verified on 2026-10-15.** Section 37 keeps the
> Ticket column visible during horizontal scrolling of the wide and intermediate tables.
> The owner chose 1A–7A: only Ticket is fixed, only on the horizontal axis, with its current
> width, opaque theme background and right border; the eight themes will be checked. During preparation
> it was detected that section 36 prevented the scrolling that this decision needs. The owner
> corrected the contract with 1A–3A: from 768 px a table with native scrolling is kept, with no
> additional notice, and only below that does it switch to cards. No open decisions remain. The specific
> test and the build passed. In the browser, the
> table appears at 768 px and the cards at 767 px; when scrolling 165.5 px, the second column moved
> but Ticket kept exactly its position. The eight themes showed an opaque background and a 1 px border.
>
> **Amendment approved, implemented and verified on 2026-10-15.** Section 36 makes
> the header of the last column of the PC inboxes visible. The real check confirmed
> that “Open” does live inside a `td`, but its `th` was visually hidden and therefore seemed
> to be outside the table. The owner chose 1B, 2A, 3A, 4A, 5A, 6A, 7A–10A and, in the full review
> of this table, 1A–12A, 13B+13C, 14A–24A. He expressly corrected the order of the columns.
> No open decisions remained. The specific test and the frontend build passed; the
> browser confirmed eleven columns at wide width, nine at intermediate width and cards before
> horizontal scrolling appears.
>
> **Amendment approved, implemented and verified on 2026-10-15.** Section 35 changes
> Reason/Motivo to Summary/Resumen in all views, removes the two retry buttons and leaves
> discrete states only in the details panel. The owner chose 1A–20B and no open decisions remain.
> The interface was left without manual actions, shows `—` when there is no current result and reserves the discrete
> state for the details panel. The 231 frontend tests passed, as did its boundaries and the build.
>
> **Correction approved, implemented and verified on 2026-10-15.** Section 34
> matches all dropdowns to the real 42 px of the Main/Both/Internal switch. The owner
> chose 1A–5A: fixed, global height, without modifying the switch or the environment's reload
> configuration. The 232 frontend tests and the build passed; a clean browser measured 42 px,
> a 14 px font and 12 × 8 px padding. No open decisions remain.
>
> **Adjustment implemented and verified on 2026-10-15, with explicit approval from the owner.** Section 33
> records that the 52 px were checked correctly in a clean browser, but
> turned out to be too large. The owner chose 1A, 2A and 3A: all dropdowns now use
> 48 px, with the current padding and typography. The 232 frontend tests passed, as did the
> build and the measurement in a new tab; no open decisions remain.
>
> **Correction implemented and verified on 2026-10-15, with explicit approval from the owner.**
> Section 32 records that the 44 px of section 31 did reach the browser, but were still
> perceived as too thin. The owner chose 1A, 2A and 3A: all dropdowns now use
> 52 px and keep the current typography. The 232 frontend tests passed, as did the build
> and the real measurement; no open decisions remain.
>
> **Implemented and verified on 2026-10-15, with explicit approval from the owner.** Section
> 31 aligns the height and spacing of all dropdowns with the text fields and buttons.
> The owner chose 1A–5A; the 232 frontend tests passed, as did the build and the real
> review on PC and mobile. No open decisions remain.
>
> **Implemented and verified on 2026-10-15, with explicit approval from the owner.** Section 30 makes visible the quality of
> the proposals and the errors of each summary without losing drafts during refresh. The
> owner closed 1A–15A, with 8B, and corrected the proposal by removing two redundant aids from the
> editor and the review panel; no open decisions remain.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section 29 replaces the modal and its five
> tones with immediate professional generation and independent inline review in each editor.
> The owner closed 1A–15A with 3B, 6B, 7A, 10B and 11B; no open decisions remain.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section
> 28 corrects the internal space
> that the double view still reserves for a hidden details panel and makes each conversation use its
> full half and enables comparison from 1152 px. The owner chose 1A–6A, 7B and
> 8B; the review is closed with no open decisions.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section
> 27 gives the two tickets in the double view independent scrolling and width, without narrowing them
> on small screens.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section
> 26 makes the local catalogue visible from the moment the screen opens and turns its button into a
> refresh without visual jump.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section
> 25 removes the browser's native confirmations and adds to step 1 the complete theme selector
> and the institutional colour control, reusing the existing components and rules.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section 24 turns each model
> card into the only activation or deactivation point, moves the greeting to a toast that identifies
> the model and turns the logo preview itself into the upload control. It also amends the
> general toast policy: all of them allow manual closing and last at most ten seconds.
>
> **Implemented and verified on 2026-10-14, with explicit approval from the owner.** Section 23 corrects the local catalogue that
> today only shows “Downloading”, blocks the wizard's premature advance, adds the dated manual
> test and brings the two logos to the first step. The behaviour was closed with the owner before
> writing this proposal.
>
> **Amendment implemented for 1.0.0, 2026-10-14.** Section 22 defines selection, accessible chip,
> in-place editing, multiple filtering and priority order on PC and mobile. No code is modified
> before implementing it. The owner explicitly approved it.
>
> **Correction approved and verified on 2026-10-10:** the review of section 21 found that the
> document promised to keep the source link in the collapsed menu, while the implementation
> and its test hide it until it is expanded again. The owner chose 39A: document the
> current behaviour, without adding another icon to the collapsed bar. The owner explicitly approved
> 39A.
>
> **Amendment implemented and verified on 2026-10-10: source and licence visible.** Section 21
> incorporates a discrete link in the entry and the menu, fed by the public brand. It passed the
> component tests and the isolated journey on PC and mobile.
>
> **Amendment implemented and verified on 2026-10-08: hide help when AI is not configured.** Section
> 20 makes the button and its modal depend on the capability received with the ticket
> detail. The owner approved 1A–3A and confirmed keeping AI mandatory.
>
> **Correction approved and implemented on 2026-10-08.** There is no real ticket journey without AI configured:
> the current mandatory nature leads to Settings first. The owner decided to keep it. The
> false state will be tested in the component and the browser journey will verify the
> configured state. The final tests respect this boundary.
>
> **Amendment implemented and verified on 2026-10-08: “Improve with AI” modal.** Section 19
> describes the button with an icon, the simple editor, the five tones, the review and its states.
> It passed unit tests and the real journey on PC and mobile, in the disposable environment.
>
> **Amendment implemented and verified on 2026-10-08 (block 6).** The only visual loose end recorded in
> `docs/arquitectura.md` was a formatting review without scope or closure criteria. Section 18
> defines the complete review chosen by the owner in 1A–14C. What exists continues
> as-built; the owner explicitly approved this amendment on 2026-10-08.
>
> **Finding corrected and verified on 2026-10-08.** Block 5 found different RAM figures between
> the selectors and the administrator, and a generic error when a model does not fit. Section 17
> aligns the figures and defines the toast chosen in 17A and 22A. The correction was explicitly approved
> on 2026-10-08.
>
> **Amendment proposed on 2026-10-06.** What exists continues as-built. Section 16 defines the
> global language as the first control and the experience for configuring, downloading and activating the mandatory
> AI. The review was closed, the proposal was approved and the implementation was verified.
>
> **Amended on 2026-10-06**, with explicit approval from the owner. The results of saving and
> testing connections in `/settings` and `/setup` were shown in a notice at the top of the document. On a long
> screen, the result falls outside the visible part and it seems that the action did not respond. The
> owner confirmed a shared toast, fixed in the window, for those results. The notices that
> explain a condition of the screen remain inside their section. It was implemented and verified on
> PC and mobile.
>
> **Amended on 2026-10-05**, with explicit approval from the owner. Settings had the same
> visual problem that step 3 of `/setup` had: a long list can hide the highlighted option.
> The owner confirmed applying in `/settings` the same read-only field before the search,
> with the exact IANA identifier, immediate update and the time and offset below. It was
> implemented and verified on PC and mobile.
>
> **Amended on 2026-10-05**, with explicit approval from the owner. When reviewing all the views it was
> found that the header navigation controls are already on the left in ticket creation and
> detail and in the user record, but `/mail` places its return to
> Settings on the right, even though this same document has said since 2026-09-29 that it should go on the left.
> The owner confirmed unifying the header controls as an icon before the title and
> that flow links inside the content keep their presentation. `/mail` was corrected and the
> position and the return to Settings were verified on PC and mobile.
>
> **Amended on 2026-10-05**, in accordance with the approved `docs/primer-arranque.md`: step 3 of `/setup`
> shows before the search a read-only field with the selected zone. It shows the
> exact IANA identifier, updates immediately when choosing from the list and allows selecting and
> copying the text; the time and offset remain below. It was verified on PC and mobile.
>
> **Amended on 2026-10-04**, in accordance with the approved `docs/prueba-local.md`: `/setup` is the initial
> exception to the browser language. A new installation presents itself entirely in English; its selector
> changes all texts, `html[lang]` and the remembered value instantly. On resuming, the language
> already saved in the installation governs.
>
> **Amended on 2026-10-03**, in accordance with the approved `docs/prueba-local.md`: The Keycloak
> forms in Settings and first start incorporate the optional internal address, separated from the
> public issuer, with ES/EN help. The Settings selector allows choosing methods that are still
> empty to show their fields; the selection is a draft until successfully saved (amendment
> §5.1.b approved).
>
> **Amended on 2026-09-30 (ninth time)**: **the Settings cards were ordered and each one
> carries what its title says**. They had been out of place since the previous change: “Numbering and
> dispatch” contained only the language, and the prefix and dispatch lived in “Time zone and public
> address”. Now there are six, in this order: **The installation** (name and language) · **The brand** (logo and
> colour) · **Authentication method** · **Numbering and dispatch** · **Time zone and public
> address** · **The emails** (`docs/modules/settings.md`, section 5.11).
>
> **Amended on 2026-09-29 (eighth time)**: **the installation's time zone** and
> **the public address** enter **Settings** (decisions 13 to 15 of `docs/modules/settings.md`). The zone is chosen from a
> **list of zones with a search** and the screen shows **the time there and its offset**; the
> address is written with its scheme, its host and —if needed— its port, with **`localhost` allowed**
> and **a notice, without blocking anything, when it is not https**.
>
> **Amended on 2026-09-29 (seventh time)**, at the owner's request, while using the product: (1) the
> **ticket number governs** —larger, and it grows on wide screens— and **the “The internal ticket” banner is removed**;
> (2) **icon-only buttons are shrunk** (from 44 px to 36, which is still above what
> can be tapped); (3) in **Settings**, “How to sign in” is renamed **“Authentication
> method”**; (4) the **“Settings” button of the mail screen** is replaced by **the back
> button** —the arrow, on the left—, as in the ticket lists; and (5) in **Users**, the button
> to send the link is renamed **“Change password”** and **is only shown on local accounts**
> (decision 90 and `docs/usuarios-y-permisos.md`).
>
> **Amended on 2026-09-29 (sixth time)**, at the owner's request, while using the product: (1) in the
> ticket header, **your team's number comes first and with more weight** —Development sees the
> **internal** first; Support and the Administrator, the **main** one— and **the links “View the internal ticket”
> and “View the main ticket” are removed**, because that is what the switch is for (decision 87); (2) **“On hold”
> is only confirmed**, without asking for text, and **closing still requires the comment** (decision 88); and (3) the
> **list of whom to tag appears next to the cursor**, where you are writing, instead of above the
> comment block (decision 89).
>
> **Amended on 2026-09-29 (fifth time)**, at the owner's request, while using the product: (1) in the
> summary tables, **the “Ticket” column shows only the number** and **the category and tags are
> grouped in a new column, “Classification”** (decision 86); (2) **the assignee is reassigned with a
> button next to their name**, which opens a modal —the loose “Assign to” field disappears—; (3) **the
> back buttons go on the left and are only an icon** (a square arrow from the project); (4) the
> **workspace is wider** on large screens —`px-5`, or `px-10` from tablet up—, without reaching
> full width; and (5) **the engine's two summaries refresh themselves** while they are
> pending, after **any** change (before, only after pressing “Summarise again”, and the field was
> left in “Generating…” until reload: the owner reported it as a hang). **The mention, by
> the way, looks like a mention**: with its `@` in front and with a theme colour that distinguishes it.
>
> **Amended on 2026-09-29 (fourth time)**, when it was clarified who tags (decision 85): the **tag modal**
> is available to **both teams**, each from the ticket in which they work —Support in the
> main one, Development in the internal one—, and **the tags belong to the main ticket** (the internal one inherits them):
> changing them from either one changes the same ones. That is what “classifying belongs to both
> teams” means, and it **does not** touch the rule that Development does not write the main ticket's text.
>
> **Amended on 2026-09-29 (third time)**, when it was verified that two decisions contradicted each other: **the
> catalogue is curated** (decision 84). Writing a new tag in a ticket **no longer creates it**: whoever
> tags **chooses from the existing ones** —the details panel modal only offers those in the catalogue— and **new
> names are introduced by the Administrator** on the catalogue screen. To anyone else, adding a
> tag that does not exist answers **422** with its notice, instead of introducing it unintentionally.
>
> **Amended on 2026-09-29 (second time)**, at the owner's request: **the user does not add tags**
> (decision 82) —**creation does not carry that field** and **the tag line is not shown in the details panel**—,
> because tagging belongs to **Support and Development**; and **tagging is done in a modal**: **all the tags
> in the catalogue** with how many tickets each one has, a **search**, **checkboxes** to select and
> deselect several at once —the ticket's own appear checked— and **“Save” at the end** (decision 82). The
> **catalogue is maintained only by the Administrator** (decision 83), so on the catalogue screen the
> buttons to create, rename and retire **are not shown to Support**. And **signing out of the application is
> confirmed**: the sign-out button in the menu opens its window before closing the session.
>
> **Amended on 2026-09-29**, **owner's correction** to decision 78: the ticket dropdown
> **is not a list of actions, it is for changing state**, and it shows **the list of states** —the
> same in main and internal tickets—, with **the current state as the first option**, checked and
> not selectable, and **only the states to which the ticket can move now** (decision 80). Choosing
> one opens **a single window** that says which one it moves to and, **if that transition needs text, asks for it
> there**: the reason when escalating, what was done when resolving, the question when putting it on hold and **the
> comment when closing, which is mandatory** (decision 81). And the actions that had a name **are those
> same states**: in an internal ticket, “On hold” returns the case to Support without closing it and “Closed” without
> resolving it returns it and closes it; in a main ticket, “On hold” is asking the user and “Closed”
> from closed is reopening it.
>
> **Amended on 2026-09-28 (sixth time)**, at the owner's request: **the entry, square**
> (decision 79). The logo box did not measure the same as the form, and the language and theme each measured
> their own thing —**measured: logo 136 px, form 329, controls 256, language 104 and theme 140**—. The
> cause was twofold: the logo carried **its own maximum width** (288 and 320 px) and, above all, **the entry
> column centres its children**, so a component **shrinks to the width of its content**
> even if inside it takes 100%. The width is now requested **on the element** —the logo, the card and the
> controls—, and the two dropdowns **split in half**. Measured afterwards: **380 = 380 = 380**, and
> 184 each dropdown. There is a test that measures it.
>
> **Amended on 2026-09-28 (fifth time)**, at the owner's request: **the ticket actions, in a
> single dropdown** (decision 78). They were four loose buttons in the details panel —“Ask the user”,
> “Escalate”, “Resolve”, “Close” in the main one; “Ask Support for something”, “It is not a code change”,
> “Resolve” and “Close” in the internal one— and the owner did not see the point of them loose. Now it is **a
> dropdown of actions**, named **by what they do** —not by the state they lead to, which remains
> in its chip—, and **those that ask for something (the reason, the question) still open their window**. It lives **in
> the top row when a single ticket is shown** and **in the header of each column when both are shown**.
> And **“Escalate to Development” becomes “Escalate”**.
>
> **Amended on 2026-09-28 (fourth time)**, at the owner's request: **windows open centred**
> (decision 77). They did not: the native `<dialog>` centres itself with the `margin: auto` of the browser's
> stylesheet, and **Tailwind's preparation sets `margin: 0` on everything**, so they appeared in the corner
> —measured: `x: 0, y: 0`—. The `auto` is restored, with maximum height and its own scrolling so that
> a tall window does not overflow. And **the window's ✕ is called “Close”**, not “Cancel”: it was called
> the same as the cancel button inside, and they were two buttons with the same name.
>
> **Amended on 2026-09-28 (third time)**, at the owner's request: **the ticket icons are the
> project's own** (decision 76). The copy-number one was the character `⧉` —a square with another behind it,
> which at 16 px reads as a square— and the edit pencil, `✏`, which **depending on the machine's font appears
> as a colour emoji**: the two looked alike. Now they are **custom SVGs** —like the menu ones and without
> libraries—: **two overlapping sheets** for copying, a **check** to confirm it and an **inclined pencil**
> for editing. The button receives them **by name**, not by their character.
>
> **Amended on 2026-09-28 (second time)**, at the owner's request: in the lists, **the module name,
> the filters and “New ticket” go in a single row** (decision 75), which was what was missing to
> avoid spending a header line: the table gains that height. At a short width they distribute themselves —the
> name, the filters and the button on the right if they fit, and whatever does not fit, below—.
>
> **Amended on 2026-09-28**, at the owner's request, while using the ticket: **everything is edited where it is
> shown** (decision 74). In the details panel, the **category** is a dropdown on its own line, the
> **tags** are chips with their cross and a small field to add, the **watchers** carry their
> **“Add”** with the people search in addition to each one's cross, and the **subject** and
> **description** carry their pencil which turns them **in place** into a field with “Save” and
> “Cancel”. **The separate forms go away** —the “change classification” one and the edit-ticket
> dialog—, which was what made them impossible to find. Also: the ticket number and state
> **were shown twice** —painted by the page header and by the ticket's own header— and
> now only the page's paints them; and in the editor you can **tag by typing `@`**, with the
> list filtered **by name or by email** as you type.
>
> **Amended on 2026-09-27 (fifth time)**, at the owner's request, while using it: **the list carries a
> single “Search” button**, which opens a modal with all the criteria —text, view, type, state,
> category and tag—, and in its header remains **the summary of what is filtered** with its clear
> button; the three rows of filters ate the screen and **the category one overflowed** (section 3.3).
> The title **“What can be done”** of the details panel's actions card **is removed**, because it was
> misunderstood. The **ticket view** becomes **the role's view** —main for Support, internal
> for Development—, “Both” shows **only the two conversations** without the details panels, and the
> **“Retire” button of a category is only shown to the Administrator**. The catalogue screen now
> has **two halves with the same weight**, with tags able to be **created, renamed and retired**.
>
> **Amended on 2026-09-27 (fourth time)**, at the owner's request: in Settings, **each
> sign-in method shows only its configuration** —“Application accounts” shows none, the
> directory appears with the organisation and the realm with Keycloak—. Before, both were seen at once
> (section 3.8).
>
> **Amended on 2026-09-27 (third time)**, at the owner's request: **copying the number** of a
> ticket is an **icon-only button**, and when copied **it changes to a green check for a few seconds**. The
> component inventory (section 6.3) records the two things that were needed: a **fourth form**
> of button —the green confirmation one— and that a button can be **only an icon**, with its mandatory
> accessible name.
>
> **Amended on 2026-09-27 (second time)**, when **categories and tags** entered
> (`docs/modules/tickets.md`, decisions 64 to 68): **creation asks for the category** —it is mandatory: without
> it the ticket is not created— and accepts **tags** in lowercase with hyphens and suggestions; the
> **list** shows them as **chips on the row, below the number** (and on mobile cards),
> with **category and tag chips to filter**; the **details panel** lets those who can edit the ticket
> change them; and the Administrator and Support menu introduces the **“Categories and tags”**
> screen, which is where the catalogue is maintained.
>
> **Amended on 2026-09-27**, when **mentions and watchers** entered (`docs/modules/tickets.md`,
> decisions 58 to 63): the editor introduces a **“Tag”** button —and typing `@` is enough— that puts the
> name of a technician or developer **inside the comment, highlighted**; that makes them a
> **watcher** of the ticket, which **is not the same as the assignee**. “My tickets” carries a new chip
> —currently **All · Assigned · Watching**, according to section 46—, the details panel has its list of **“Watchers”** with its cross to
> remove them (Support and Development), and removing someone **is recorded in the timeline**.
>
> **Amended on 2026-09-27**, when the owner asked for **two new columns in the ticket list**,
> written by an artificial intelligence engine (`docs/modules/ai.md`): **“Reason”** —what the
> ticket is about— and **“Last action”** —what happened last—, both **in the language of the viewer**.
> And with them, **two names change**: the “Number” column is renamed **“Ticket”** and “Last
> update” becomes **“Update date”** (section 3.3).
>
> **Amended on 2026-09-26**, when the owner asked that **images and video not take the full
> width**: in the description and comments they are seen **at most at 480 × 360** —a full-screen
> screenshot ate the conversation— and **when clicked they open in the large viewer**, which is
> where they are seen whole and downloaded. **Video becomes a thumbnail with its play button**,
> and plays in the viewer (section 6.5). It is decision 55 of `tickets.md`.
>
> **Amended on 2026-09-26**, when the owner asked that **the inbox be “mine”** and that there be
> **two new modules to see everything**: the Inbox becomes **My tickets** —its current semantics
> are in section 46—, and **Main tickets** and **Internal tickets** are born, only for
> Support and Development (sections 3.2, 3.3 and 3.7). **“New ticket” leaves the menu** —it was confused with
> a module— and stays where it is created, in the inbox button. **“Assign to me” disappears from the
> inbox**: reassignment is done inside the ticket, which is where you see who it is being passed to. The
> inbox rows say **“Open”** and not “Open the ticket”, the text “What has been waiting longest
> goes first…” is unnecessary —the rule remains, silent—, and **the language becomes a dropdown like the
> theme one**, with its width and without a visible label (sections 3.1 and 6.3).
>
> **Amended on 2026-09-26**, when implementing the editor with attachments: **pasting text pastes
> without formatting** —with line breaks—, and **a pasted file does enter as an attachment**, because pasting
> the HTML of an email would bring tags that the backend rejects and the entire comment would be lost because of
> something the person did not write (section 6.5). And **the gap of a file that did not upload is shown
> marked**, in the place where it was going.
>
> **Amended on 2026-09-26**, when the owner asked that **attachments be seen inside the text** and
> that **the preview also be seen while writing**: the writing box becomes a **formatted editor** (section 6.5), the
> **viewer** grows and gains **Download** and **Open in a tab**, and
> **the file upload piece is removed**, because attaching is no longer a separate box.
>
> **Amended on 2026-09-25**, when the owner asked that **the system version be
> seen**: in the sign-out row of the sidebar, right-aligned, and **also in the footer of the
> entry screen** (sections 3.1 and 6.4). **It is his choice and not my recommendation**: I proposed
> only the menu, and he decided that it should also be seen before signing in, which is where it serves to say which version
> is being looked at when reporting a bug.
>
> **Becomes as-built on 2026-09-25**, and **amended the same day**: the **card of how to sign in** on
> the Settings screen —the method, the two configurations and their two connection tests— and what
> the entry screen offers according to the method (section 4.4). The shape of the screens that this
> document fixes is all built.
>
> **Amended on 2026-09-25**: the installation has **its own name**, configurable, and
> **it is what everyone reads** —the entry screen, the sidebar and **the browser
> tab**— instead of the factory text (section 6.4 and `docs/modules/settings.md`, section 5.7).
> **The owner corrected two things when asking for it**: the emails and their sender do not change.
>
> **Amended on 2026-09-25**, when implementing the three actions that depended on the
> directory: on the user screens, **reactivation is only offered where it can work** —on an AD
> account yes, and on a Keycloak one no, with what happens explained instead— and **the origin note**
> distinguishes what happens: AD can be queried from creation when it is the current method;
> Keycloak continues creating the account on sign-in (`docs/modules/users.md`, section 15).
>
> Approved by the owner on 2026-09-22, after reviewing it in the form of questions while it was
> being written: **there is no top bar** —the sidebar carries the product, the role's options and,
> at the bottom and separated, the controls (profile, theme and sign-out)— and the `administrador` role **sees
> tickets read-only**. It fixes the principles (intuitive before beautiful), the shape of the
> application, the experience of the four roles, the main/internal double view, the visual
> language with **Tailwind v4 and custom components**, the eight themes and what is expected on each
> device.
>
> With this document, **the documentation is complete**: seven documents that cover what is
> built, how it behaves, how it looks and how it is deployed.
> **Amended on 2026-09-23**, when writing the session shell: it adds **where the interface
> language comes from** (principle 6: from the browser, and if it does not request either of the two, English).
> It is an **owner's correction**: I had proposed Spanish by default
> (`docs/modules/auth.md`, decision 25).
>
> **Amended on 2026-09-24**, when building the shell: section 3.1 says that **the
> collapsed state of the menu is remembered** and that **the menu shows only what exists** —the entries of
> table 3.2 arrive with their screens—, and 6.4 records the logo in the menu.
>
> **Amended on 2026-09-24**, when the institution's logo entered: section 6.4 says **where
> the brand is shown** (entry, sidebar and Settings) and **how any logo is shown**
> —whole, without cropping and in a box with the theme's border—, so that one with an opaque
> background works the same as one with transparency. The logo and its replacement belong to the `settings` module
> (`docs/modules/settings.md`).
>
> **Amended on 2026-09-24**, when building the mail editor: section 3.8 fixes
> **its screen**, with the two languages side by side, the placeholders that are inserted when clicked and a preview
> that **renders the backend** with the sample data from the test, so that what is seen is
> what will go out.
>
> **Amended on 2026-09-24**, when building the ticket screens: section 3.7 fixes
> **the four routes** and who sees each one, and records that **the detail is the same screen for the
> four roles** and that **the prefix and dispatch are configured in Settings**, not on a screen
> of `tickets`. The table in section 3.2 introduces **“New ticket” for Support** —three entries—:
> they can create a ticket on behalf of another person and that permission needed its door (**owner's
> correction**). Escalation is written **in a dialog**, because the reason is mandatory and the action
> is deliberate, and attachments are seen **with a preview** when they are images or PDF.
>
> **Amended on 2026-09-24**, when tickets began: section 3.3 adds the **type
> chip** (Main · Internal · All) in the Administrator inbox, which is the only role that
> sees both ticket types.
>
> **Amended on 2026-09-24**, when building the user screens: section 3.6
> fixes **the three `users` screens** —the list, the record and the own profile—, their addresses and
> who sees each one. The five decisions belong to the owner, from 2026-09-24, and come from a finding:
> **Support can change another account's first and last names, but does not see anyone's record**, so
> that permission is exercised **inline, from the list itself**. Section 3.1 also records that **the
> menu measures the visible area and scrolls internally** —it measured the document height, and on mobile it left
> the sign-out button off screen—, and 6.3 that the card **chooses the level of its title**,
> because a screen with three `h1` is not navigable with a screen reader. Both were found by the Playwright
> layer, not by the unit tests.
>
> **Amended on 2026-09-23**, when setting up the theme: section 6.2 says **how it is implemented**
> (an attribute on `html`, the palettes written once with `light-dark()`, the choice remembered in
> the browser and the theme applied by our code **because the production CSP does not allow inline
> scripts**). The switch has **two options, light and dark**: **“automatic” is neither shown
> nor selectable**, because following the system is not having chosen. The six fixed palettes remain for the
> complete shell, when their names and colours are approved.

## 1. Scope of this document

It fixes **how it looks, how it is navigated and what each role feels**: the structure of the screens, the
visual language, the behaviour on each device and the experience of the four roles.

It **does not repeat** what each role can do (`docs/usuarios-y-permisos.md`), nor the flows
(`docs/flujos.md`), nor the data model (`docs/modules/tickets.md`). When this document and those do not
match, those govern: here only how it is presented is decided.

It is **approved**, so it enables writing code (Rule 0). Section 11
records what I proposed.

## 2. The principles, in order

The order matters: when two clash, the one higher up wins.

1. **Intuitive before beautiful.** If you have to choose between something understood without explanation and
   something that looks better, the former wins. A product that has to be explained is a product that is
   used badly.
2. **No noise** (the project's criterion, `docs/propósito-y-alcance.md`): what does not contribute is not
   there. No decorative animations, no counters nobody looks at, no notices closed without
   being read.
3. **Each person sees their work, not the entire application.** The user does not need to know that an
   internal ticket exists, nor does a developer need to see user administration.
4. **The state is read at a glance.** Look at the list and know what is waiting for whom.
5. **Nothing important is done by accident.** Closing, escalating or deleting are confirmed; everything else
   is not.
6. **Language is global.** A new installation opens `/setup` in English and its first control can
   change the whole screen. Afterwards the language saved by the installation always governs, also
   before signing in and in open sessions; there is no selector or per-account preference.
7. **What fails is explained.** An error says what happened and what to do, in the language of the reader.

## 3. The shape of the application

### 3.1 The shell

```text
┌───────────┬───────────────────────────────────────────────┐
│  producto │                                               │
│           │                                               │
│  opciones │  contenido                                    │
│  del papel│                                               │
│           │                                               │
│  ───────  │                                               │
│  controles│                                               │
└───────────┴───────────────────────────────────────────────┘
```

- **There is no top bar.** The sidebar carries everything, and the content gains the full width: in
  a help desk the screen fills with lists, conversations and attachments, and a top strip
  is space taken away from what matters.
- **The sidebar has three zones**, visually separated:
  1. **Top**: the product —**the installation logo and name**—. It is the only thing that does not move.
  2. **Middle**: the **role options** (section 3.2). It is what changes from one person to another.
  3. **Bottom, separated by a line**: the **controls**, which are the same for everyone —**your
     name (and your profile)**, **the language**, **the theme** and **sign out**—. It is the same distribution used by
     Calibyou, and it avoids mixing “what I do here” with “who I am”.
- **Language and theme are the same control, one below the other** (owner's decision,
  2026-09-26): both are **dropdowns**, they take **the same width** and **neither carries the label
  visible** —it reads “Spanish” or “Light”—, because the value already says what it is. The language is changed
  **by choosing any value from the list**, and the accessible name remains there, hidden, for those who navigate
  with a screen reader. Before, the language was two segmented buttons and the theme a dropdown: two
  forms for the same thing, and two different widths in the same column.
- **The system version is read in the sign-out row, right-aligned** (2026-09-25): it is a
  piece of installation data, not an option, so **it is not a button nor does it occupy its own row**: it goes with the
  control that is already in that row. In the collapsed menu it is lost with the text, which is correct: a
  loose number next to an icon says nothing. **On mobile it is read inside the drawer**, which is where
  that row lives.
- On PC the menu can be **collapsed** and a strip of icons remains; on tablet it will collapse by itself. **The
  state is remembered** in the browser, like the theme: those who work with the menu collapsed do not have to
  collapse it every morning.
- **It measures 13 rem (208 px)** and **3.5 rem (56 px) collapsed**. The reduction gives more width to
  the content without narrowing the touch targets. The installation name is shown complete,
  centred and in as many lines as needed below the logo; it does not widen the menu and the header can
  grow vertically.
- **The menu shows what exists**: each entry arrives **with its screen**, not before —a menu that
  leads nowhere is worse than a short menu—, and it is not shown to those it does not belong to. The user sees
  a single option, **My tickets**, and that is all they need: to ask and follow their own.
- **The menu measures the visible area, not the document, and scrolls internally if its content does not fit**
  (owner's decision, 2026-09-24, when building the user screens). The drawer measured the
  document height: on a mobile with a browser bar, the visible area is 839 px of 890, so **the
  sign-out button was below what is seen**, impossible to tap; and on PC, with a long list, the bar stretched to the whole
  content and its controls ended up at the end of the page. On
  PC the bar goes **stuck to the top** while the content scrolls, which is what section 7 calls
  “fixed”.
- **On mobile there is a minimal top strip**, with the menu button and the product name: without
  it there would be nowhere to put the trigger. Everything else lives in the menu, which opens as a drawer.

### 3.2 Menu entries by role

| Role | Menu options |
| --- | --- |
| **User** | My tickets |
| **Support** | My tickets · Main tickets · Internal tickets · **Classification** · Users |
| **Development** | My tickets · Main tickets · Internal tickets |
| **Administrator** | Inbox (read-only) · **Classification** · Users · Settings |

**Five entries at most** in the busiest case —Support— and four for the Administrator, and **the
controls do not count**: the profile, **the language**, the theme and sign out go at the bottom and are the same for everyone.
**They were four until 2026-09-27**, when the screen now called **“Classification”** entered with the categories: it is a
ticket screen —the catalogue is ticket data, not installation configuration— and that is why
it lives in the menu and not inside Settings (`docs/modules/tickets.md`, decision 64). If any role
were to need more options, it would be a sign that something has become complicated elsewhere.

**“New ticket” is no longer in the menu** (owner's decision, 2026-09-26): it was confused with a
module, and creating a ticket is an action of the inbox, not a place to go. The button lives **inside
the inbox** —at the top and also when it is empty—, so **nothing is lost**: the route
`/tickets/new` still exists and is still the same screen.

**Support has five** because, in addition to its own, it is the only role —with Development— that sees **everything
there is**, and now that is seen in two places at once: main and internal. Before, everything was together
in the inbox, which is what stopped working when the inbox became “mine”.

### 3.3 The inboxes

**Three inboxes and one screen.** The same list screen serves all three —the content
and the title change—, because separating them into three screens would triple the table, the filters and the pagination
to change one condition (`docs/modules/tickets.md`, section 8).

| Inbox | Who | What it shows |
| --- | --- | --- |
| **My tickets** (`/tickets`) | User, Support and Development | User: their requested main tickets. Teams: assigned or currently watched, limited to their operational level |
| **Main tickets** (`/tickets/main`) | Support and Development | **All** main tickets |
| **Internal tickets** (`/tickets/internal`) | Support and Development | **All** internal tickets |
| **Administrator inbox** (`/tickets`) | Administrator | Both types, with their chip and read-only |

- **“Mine” is a current relationship** (updated on 2026-10-19): for Support and Development it is what is
  assigned or currently watched. Opening or commenting does not add the ticket, and ceasing to watch it
  removes it unless it remains assigned. Support sees only main tickets here and Development only internal ones.
- **The list is a table** on PC and **cards** on mobile, where a table of eight columns cannot be read.
  The columns are: **Ticket** · **Subject** · **Reason** · **Last action** · **Classification** ·
  **State** · **Requested by** · **Assignee** · **Update date** (the two names at the
  ends and the two middle columns are from 2026-09-27: “Number” was what is read as a ticket, and
  “Last update” is a date, not an event; **“Classification” is from 2026-09-29** and carries the
  category and tags, decision 86 — and **“Ticket” keeps only the number**—).
- **Everything is edited where it is shown** (owner's decision, 2026-09-28): the category, the
  **tags**, the watchers, the subject and the description **are changed in the place where they are read**,
  not in a separate form. The distribution, for **the text and the category**, is the usual one: **the
  requester in their ticket and Support**; Development does not write the main ticket and the Administrator touches
  nothing, so **they are not shown the pencil**. **The only thing Development does do in a main ticket is follow whoever follows it**: the
  watcher list —add and remove— belongs to **both teams** (decision 63), because following a
  ticket is not writing it. **And tags also belong to both** (decision 85): they are added, changed
  and removed **from the ticket each one works in** —Support in the main one, Development in the
  internal one—, and since **they belong to the main ticket** (the internal one inherits them, decision 67), changing them from either
  of the two changes the same ones. The interface test found it when checking that we were not leaving them anything —the interface does not offer what the backend would reject—. Each control says what it
  does and **the change is saved on confirmation**, never while typing.
- **Tag by typing `@`** (owner's decision, 2026-09-28): in the editor, typing `@` opens
  the list **filtered by what is still being typed**, and **the name or the email** works —the search
  shows **both**, which is what makes it possible to distinguish two people with the same name—. The mention
  **is read with the name**, as until now, and the “Tag” button is still there for those who do not
  remember the shortcut.
- **Sending the link asks first** (owner's decision, 2026-09-28): the link is the sign-up one
  for those who have never signed in and the **recovery** one for those who already have a password —the one they had
  stops working—, so the screen **asks first with a window**, and the notice says **which of the
  two it is** and how long it lasts: 24 hours for sign-up, 1 hour for recovery. No browser `confirm()` is used,
  as in the rest of the actions that ask.
- **In the ticket header, the internal number** (owner's decision, 2026-09-28): when opening
  an internal ticket, that row showed only the main number and it was not known which of the two was being
  viewed. It goes in muted text, because it is the reference and not the title, and **the back button is an icon with its accessible name “Back”**
  —it said “Back to the inbox”, and you return to the inbox from many places—.
- **Watchers are also added by hand** (owner's decision, 2026-09-28): in addition to
  arriving through a mention, they can be **added from their list** with the people search and **removed**
  with their cross. Whoever is added this way **is the same** as whoever arrived through a mention: all are
  watchers, and the list does not distinguish where each one came from.
- **The details panel's actions card has no title** (owner's decision, 2026-09-27): it said
  “What can be done”, which **is misunderstood** —it seems a promise of what the system will do, and it is
  only where the buttons are—. The buttons still say what they do, which is what was needed.
- **The ticket view, by role** (owner's decision, 2026-09-27): when opening a ticket,
  **Support sees the main one** and **Development the internal one**; both can switch to **Main · Both
  · Internal** whenever they want, and **opening an internal number shows the internal one** even if you are
  Support. **“Both” shows only the two conversations** —each ticket's column—: with the four
  columns (two conversations and two details panels) “it looks very saturated”. To act on one, that one is chosen.
- **Copying the number is an icon-only button** (owner's decision, 2026-09-27): the
  clipboard, and **when copied it changes to a green check for two seconds** and returns by itself. **The
  drawing is the project's** —two overlapping sheets—, and it is not confused with the edit pencil
  (2026-09-28). The
  accessible name **changes with the state** (“Copy the number” → “Copied”), so whoever does not see the
  colour does hear the confirmation; the green comes from the theme (`--exito`), not written by hand.
- **The row carries its classification chips** (2026-09-27, **moved on 2026-09-29**, decision 86):
  **a category chip** —the ticket's *what it is*— and those of its tags. In the **table** they live in their
  own column, **“Classification”**; in the **mobile cards**, where there are no columns, they remain
  below the number. A chip takes what its name measures and is read at a glance. **The category is not
  clipped**: it is one word or two.
- **A single search button, and everything inside** (owner's decision, 2026-09-27): the list carries
  **“Search”**, and the modal it opens has **all the criteria** —the text, the view, the type, the
  state, the category and the tag— with its **search** and **clear** buttons inside. In its place
  nothing remains but **the summary of what is filtered** —“3 filters”— with its clear button:
  the list looks clean and **it is understood why few tickets appear**. Before there were **three rows of
  filters** for four criteria, and the category one **overflowed the screen** when the catalogue
  grew; a modal does not grow, it fills.
- **And both (category and tag) enter the search**: searching “network” finds the category “Network”
  and the tag `red`, which is what they were requested for.
- **“Reason” and “Last action” are writings of the AI engine** (`docs/modules/ai.md`), and they are shown
  **clipped to two lines**, because they are a summary for scanning the list and not the ticket text: the
  full text is read in the details panel. While the engine has not answered, they say **“Generating…”**; if there is no
  engine or it failed, they say so and **nothing is invented** (decisions 2 and 8 of `docs/modules/ai.md`).
- **Both appear on the mobile cards**, as two more lines: on mobile is where knowing what a ticket is about
  without opening it is most appreciated.
- **The row says “Open”**, not “Open the ticket” (2026-09-26): in an actions column the word
  “ticket” is unnecessary, and the whole row already speaks of a ticket.
- **There is no “Assign to me”** (2026-09-26). It was removed from the table and the cards: **reassignment is
  done inside the ticket**, in its details panel, where you see who it is being passed to, and with the usual rules
  (`docs/modules/tickets.md`, decision 50). The list shortcut assigned without more, and that is what
  left reassignment to another colleague undone.
- The filters are **visible chips** above the list, not a hidden form: state, text and
  nothing else.
- **“My tickets” carries the selector in its first row —All · Assigned · Watching—** for Support
  and Development (updated on 2026-10-19), and **starts on “All”**. “All” is the union, without
  duplicates, of current assignment and watching; the visible selector does not count as search.
- **The level is not chosen in “My tickets”**: Support sees only main tickets and Development only internal ones,
  even when the account accumulates the Admin capability. **The user does not carry the selector** and the
  pure Administrator keeps their general read-only inbox with the type selector.
- **The two “everything” inboxes do not carry a type chip** for the reason above: “Main tickets”
  shows main tickets and “Internal tickets”, internal ones. They do carry the **state** chip and the search.
- The **state goes with colour and text**: colour helps scan the list, the text says what it is.
  Never colour alone (section 8).
- Default order: **what has been waiting longest, first**, within what is not
  closed. Closed tickets do not get in the way. **The rule remains written here and is no longer written on the screen**
  (2026-09-26): the order is seen in the list, and a sentence explaining what is already seen is noise.

### 3.4 The ticket detail: the central piece

It is the screen where time is spent, so it takes most of the design.

```text
┌──────────────────────────────────────────────────────────────┐
│  ACME-2026-0042 · asunto                    [estado]         │
├──────────────────────────────┬───────────────────────────────┤
│  conversación                │  ficha: solicitante,          │
│  (comentarios, adjuntos,     │  responsable, fechas,         │
│   historial intercalado)     │  adjuntos                     │
│                              │  acciones                     │
├──────────────────────────────┴───────────────────────────────┤
│  escribir un comentario…                       [Enviar]      │
└──────────────────────────────────────────────────────────────┘
```

- **A single timeline**: people's comments and what the system did, in order.
  The history is not a separate tab: if the state changed, it is seen there, among the messages.
- **The details panel and the actions go on the right** on PC: who requested it, who is handling it, since when, and
  the state, assignment and escalation buttons. On mobile they go below, collapsed.
- **The ticket number always visible and with a copy button**: it is what is said aloud and
  what is pasted into an email.
- The comment box is **always visible** at the end of the conversation. Writing is the
  most frequent action in the system; hiding it behind a button is a beginner's mistake.
- A **closed ticket does not accept comments** (`docs/flujos.md`): in its place the explanation of
  why appears, and **reopening is chosen in the actions dropdown** (decision 78), which is where all
  the ticket actions live; there are not two controls for the same thing.

### 3.5 The user does not read internal jargon

The user **never** sees the internal ticket or the word `escalado`. And their states are told in
plain language:

| Internal state | What the user sees |
| --- | --- |
| `nuevo` | Received |
| `en progreso` | In progress |
| `en espera` | Waiting for you |
| `escalado` | In progress *(the user has no reason to know it has gone to Development)* |
| `resuelto` | Resolved |
| `cerrado` | Closed |

### 3.6 The user screens: the list, the record and the profile

They are **the first product screens**, and there are three, each with its address:

| Screen | Route | Who sees it |
| --- | --- | --- |
| The account list | `/users` | Support and Administrator |
| An account's record | `/users/{id}` | Administrator |
| The own profile | `/profile` | Anyone who has signed in, except the factory account |

- **The list is a table on PC** —name, email, role, origin, state and last sign-in, plus the row's
  actions— and **cards on mobile**: a table of six columns cannot be read on a phone, and
  it is not dragged sideways (section 7).
- **The filters are visible chips above the list**, like those of the inboxes (section 3.3): role,
  origin and state, plus search by name or email. Nothing hidden behind a button.
- **The order is fixed —surnames and first name— and the list is paginated**, with the total visible. **Sorting
  by column is not offered**: that belongs to the inbox table, which is where it is needed.
- **Creation is a dialog over the list**, not a screen: six fields and a button. On saving, the
  list refreshes with the new account in its place, without leaving where you were.
- **Support edits the first and last names inline, from the list itself** (owner's decision,
  2026-09-24). Support **does not see anyone's record** (`docs/modules/users.md`, section 4,
  decision 2), and that permission must be exercisable somewhere: the finding was precisely that,
  without a record, there was nowhere to do it. The Administrator changes the same from the record.
- **The record** shows the account —name, email, role, origin, state and when they last signed
  in— and **its actions**: change the name, email, role and language, deactivate or reactivate,
  change the origin and **send the password link again**. Every sensitive action says what is going
  to happen before doing it (section 4.4).
- **Reactivation is offered only where it can work**, and it is not a cosmetic detail: on an
  **AD** account the button asks the directory before returning access, and on a **Keycloak** one it is not
  offered —it is explained that **it comes back by itself on sign-in**, which is what actually happens— because its path has
  no way to check who is still there without a service account in the realm
  (`docs/modules/users.md`, section 5, point 4). **The interface does not offer what cannot be done.**
- **The deactivate button is not offered on yourself** (`docs/modules/users.md`, section 7):
  deactivating yourself is being left out instantly and with no way back.
- **The origin is offered with its three values**: `ad` is enabled only when it is the current method and
  opens the directory search; `keycloak` remains disabled and explains that the account is created
  on first sign-in (`docs/modules/users.md`, section 15).
- **Your own profile opens from your name**, in the menu's controls zone (section 3.1), and
  carries **your first name, your surnames and your language**, plus the link to change the password. The language is
  changed here and in the controls: it is the same setting, and it is the one that decides in which language
  emails are written to you (`docs/modules/users.md`, section 8).
- **The factory account has no profile**: it is not in the account table, so **its name in the
  menu leads nowhere** and the screen is not shown to it.

### 3.7 The ticket screens: the inbox, creation and the detail

Four routes for the lists, and **the detail is the same screen for the four roles**: what
changes is what each one can do inside, not the screen.

| Screen | Route | Who sees it | What it is called in the menu |
| --- | --- | --- | --- |
| **My tickets** | `/tickets` | User, Support and Development | **My tickets** |
| **Main tickets** | `/tickets/main` | Support and Development | **Main tickets** |
| **Internal tickets** | `/tickets/internal` | Support and Development | **Internal tickets** |
| **The Administrator inbox** | `/tickets` | Administrator | **Inbox** |
| Creation | `/tickets/new` | User, Support and Development | — (it is entered from the inbox button) |
| Detail | `/tickets/{numero}` | Whoever can see it | — |
| **Classification** | `/tickets/categories` | Support (consults the catalogue) and Administrator (maintains it) | **Classification** |
| The prefix and dispatch | `/settings` | Administrator | **Settings** |

- **A single list screen with four names**: it is the role that decides what is seen inside, and the
  document already calls it differently by role. The name is the only thing that changes.
- **The four addresses belong to the `tickets` module** and none of them asks the backend for anything new: the lists
  are already distinguished by their type and by whose ticket it is (`docs/modules/tickets.md`, section 5).
- **The catalogue screen has two halves with the same weight** (owner's decision,
  2026-09-27): **categories** and **tags**, each with its list, its ticket count and its
  buttons. Tags can be **created**, **renamed** —which applies to all tickets that carry
  it— and **retired** —which removes them from them, warning first—; the retire button **is only
  shown to the Administrator** and the create and rename ones to Support and the Administrator.
- **The catalogue screen belongs to `tickets` and not to Settings** (2026-09-27): it lists the
  categories with their name and **how many tickets each one has**, leaves **create and rename** to Support and
  the Administrator, and **retire or restore** only to the Administrator —retiring a category in use
  is allowed and **does not touch the tickets that have it**: it stops being offered on creation—. Below, **the list of
  tags that are being used**, with their count, which is what makes people see that `red-wifi` already
  exists and not write `wifi-red`.
- **The Administrator keeps the usual inbox**, with its type chip and read-only
  (owner's decision, 2026-09-26): **they do not have their own tickets**, so “My tickets” means
  nothing to them, and the two “everything” inboxes would give them the same as they already have, split in
  two.
- **Creation is a screen and not a dialog**, unlike creating an account: here a long
  text is written and files are attached, and that needs room. On mobile it is a full screen.
- **The addresses were fixed by the backend before the screen**: the emails that already go out carry
  `/tickets/ACME-2026-0042`, so that route is not chosen here, it is respected.
- **The detail of an internal ticket is the same detail**, with its number (`INT-…`) and its conversation:
  there is no separate screen for internal tickets.
- **The numbering prefix and dispatch are configured in Settings**, not on a
  `tickets` screen: they live in the installation and their endpoints belong to `settings`. A screen of the `tickets` module
  that called `/api/settings` would break the hard modularity rule. `tickets.md`, section
  8, listed them as its own screen and it is corrected (2026-09-24).

### 3.8 The mail editor

It lives in the `mail` module, at `/mail`, and **Settings leads to it**: it is a screen of another
module on purpose, because calling `/api/mail/**` from `settings` would break the hard modularity
rule (section 4.4).

| Screen | Route | Who sees it |
| --- | --- | --- |
| The mail editor | `/mail` | Administrator |

- **On the left, the fourteen emails**; on the right, the one being edited: its **subject**, its
  **body** and its **placeholders**, with the essential ones marked.
- **The two languages go side by side** —the same email in Spanish and English, to see that they say the
  same— **and on mobile they stack**: two columns do not fit on a phone.
- **The body is edited with formatting, not by writing tags**: bold, italic, list and link
  **over the selected text**, and the HTML is written by the application. It also **accepts HTML written by
  hand**, because whoever administers the installation knows what they are doing.
- **Placeholders are inserted when clicked**, at the point where the cursor is.
- **The preview goes beside it, and the backend renders it** with sample data: it is not an
  approximation, it is the email that will go out. Below, the **“send me a test”** button, which sends
  the real email to the address of whoever asks for it.
- **Restoring the factory text** is in each template, and the notice that an essential placeholder is missing
  **does not prevent saving**: there are legitimate notices without a link.

## 4. The four approaches

### 4.1 User: ask and follow

What a user does is **create a ticket and reply**. Everything else is superfluous.

- **On signing in**: their tickets, open ones first. A large, visible **New ticket** button.
- **Create**: subject, description and attachments. Three fields, none with jargon. If a screenshot
  can be dragged onto the window, better than looking for it.
- **Follow**: the conversation, with their comments and Support's. When Support asks them something, the
  state says it —“Waiting for you”— and the email has warned them.
- **Close and reopen**: both, in the ticket's **actions dropdown**, and closing asks
  first (`docs/usuarios-y-permisos.md`: the user also closes).
- **What they do not see**: private comments on the main ticket, the internal ticket, the user list, anything
  administrative. Their world is their tickets.

### 4.2 Support: attend and escalate

Their problem is not understanding the application, it is **not losing sight of what is theirs to do**.

- **On signing in**: the **inbox**, and within it what demands attention first — `nuevo`,
  `en espera` and `escalado`. One glance and you know where to start.
- **Assign to yourself**: one click from the list, without opening the ticket.
- **Attend**: comment, move the state, ask the user for information (a button that says
  **“Ask the user”**, not “move to `en espera`”), resolve or **escalate**.
- **Escalate**: a form with **the reason as the main** and mandatory field. It is what
  turns an escalation into work and not a “you pass it on”; the interface has to make clear that
  without that text it cannot be escalated.
- **The double view** (section 5) is theirs: they are the one who translates between the user and Development.
- **Users**: creation, deactivation and password reset, without leaving the application.

### 4.3 Development: understand and resolve

Development arrives at a ticket that already comes filtered, and **does not speak with the user**. Their screen has
to make that clear.

- **On signing in**: their **internal tickets** inbox.
- **Context**: they read the **main** one (the conversation with the user) and the **internal** one where they work.
  The double view is exactly this.
- **Actions**: assign to yourself, **ask Support for something** (“I need something from Support”), and **resolve** with
  the explanation of what was done.
- **Return**: a button that says **“It is not a code change”** and returns the case to Support with a
  comment. The interface does not call it “close without resolving”: it says what it means.
- **What they do not do**: write on the main ticket, comment as if they were Support, or write to
  the user. The interface **does not offer** those actions, instead of offering them and denying them
  afterwards.

### 4.4 Administrator: keep the system running

Their work is once at the beginning, not every day.

- **The first time, they land on start-up**: a screen that groups the **installation
  language**, the **appearance** (theme and primary colour) and the **numbering prefix**, to
  leave it ready at once instead of looking for it in three places. Afterwards they enter wherever they want, and
  that same screen remains in the menu as **Settings**.
- **On signing in**: the user list, which is what they are going to touch.
- **Users**: creation (with the role and origin), deactivate, reactivate, change the origin and **send a
  password reset**. Every sensitive action says what is going to happen before doing it.
- **Numbering**: the prefix, with a notice that **it does not change numbers already issued** — it is the
  doubt anyone has before touching it.
- **And now also how to sign in** (2026-09-25): the **sign-in method** card —the three in a
  dropdown, with the explanation of the chosen one— and **the configuration that applies according to the chosen
  method**, with its **“Test connection”** button. **Each panel appears with its method and no other**
  (owner's decision, 2026-09-27): with **“Application accounts”** neither of the
  two is shown —there is nothing to configure—, with the **organisation** the directory one appears, and with
  **Keycloak** the realm one. **What is not being used is not work to be done**, and the two
  configurations at once invited touching the wrong one. **The options of an
  unconfigured method are shown but cannot be chosen**, and that also applies to what is being
  written: as soon as the server and the search base are filled in, the option turns on without
  saving first. The fields of the two secrets **always arrive empty** and their help says whether one is
  set: the screen never shows a secret, because it does not receive it (`docs/modules/settings.md`,
  section 5.8).
- **Settings already has everything of its own** (2026-09-24 and 2026-09-25): **the installation name**
  —the field, with its help and its save button; it is seen throughout the application as soon as it is saved—,
  **the brand** —the two logo slots, with their preview, upload and return to the factory one—, **the institutional colour** with **preview of how
  it looks** before saving, the **global installation language** —the first control of the screen— and the
  **numbering prefix** with the notice that **it does not change numbers already issued**, plus the
  **dispatch and notices** of each ticket type. Everything on this screen comes from a single endpoint,
  the `settings` module's, plus the logo ones.
- **Emails are edited on a `mail` module screen**, to which Settings links: each
  template with its subject and body, the language beside it —both are edited—, **the available
  placeholders visible** (`{{nombre}}`, `{{enlace}}`, `{{numero}}`…), a **preview** and a
  **“send me a test”** button. A placeholder that does not exist is warned about **on saving**, not when the
  email has gone out. `docs/modules/mail.md`, decision 15.
- **It is a screen of another module on purpose**: calling `/api/mail/**` from `settings` would break
  the hard modularity rule (`docs/arquitectura.md`, section 4), so the screen lives where
  its data lives and Settings only leads to it.
- **The body is edited with formatting**, not by writing tags: bold, italic, list and link
  over the selected text, and the HTML is generated by the application. See decision 1 of `mail.md`.
- **They see everything, without touching anything**: they can look at any ticket —including internal ones and the double
  view— but **they do not comment, do not move states, do not assign and do not close**. It is the price of being the role
  that can do everything in configuration: where they do not work, they watch.

## 5. The double view: main and internal

Requested on purpose, and it is the piece that makes the two-ticket model usable. **It only exists for
Support and Development**, because the user does not see the internal one.

- **When it appears**: when the ticket has an internal one and the role can see it. If there is no internal one, there is
  nothing to compare and nothing appears.
- **How it is controlled**: a switch —**Main**, **Both**, **Internal**— in the header of the
  detail. On PC it has three positions; “Both” is the default value when the internal one exists,
  because whoever is there wants both things.
- **On PC**: two columns, each ticket's conversation and its details panel. It is comparing, and comparing is reading
  in parallel.
- **On mobile and vertical tablet**: two columns do not fit, so the same switch keeps
  **Main** and **Internal** and changes between one and the other; **“Both” is not offered**. **Never two
  narrow columns**: it is worse than seeing them separately.
- **How not to confuse them**: the internal one carries a different background and its `INT-…` number clearly visible. If the
  two look the same, someone is going to write in the wrong ticket.
- **The actions belong to the ticket being touched**: in “Both” each column has its details panel and its
  buttons, and the comment box is the one of the column where you are. Writing on the main ticket from
  Development is not an option that is offered (read-only).

## 6. The visual language

### 6.1 The base

**Tailwind CSS v4 with custom components.** The look is written with utilities and **the components
are built here**, not brought from a library.

What that means, put clearly:

- **The look is exactly the one decided**, without fighting the styles of a foreign library
  or its theme system.
- **The complex components are ours**: table, dialog, menu, tabs, file upload,
  notices. Each one has to be built **and made accessible** (section 8). It is the price, and it is a
  real price: writing a button is not the same as a dialog that traps focus and returns it when
  it closes.
- **Closed inventory of components** (section 6.3): what is not on that list is discussed before
  building it, instead of appearing by surprise.

Calibyou uses Tailwind 3; here **4** is used (the current one), which defines colours as CSS variables
and fits better with the theme system of section 6.2.

**Modern, sober and without ornaments.** The trend of these years —flat interfaces, large typography,
lots of white space, soft borders, almost non-existent shadows— coincides with what is needed
here: nothing that competes with the data.

| | |
| --- | --- |
| **Typography** | A single system family (the one the machine already uses), with a fixed size scale. No downloaded fonts: they load late and on a help desk they contribute nothing |
| **Spacing** | A 4 and 8 pixel scale, without loose values |
| **Borders** | Small radius and thin borders; a border separates better than a shadow |
| **Shadows** | Only on what really floats: dialogs and menus |
| **Colour** | Neutral for almost everything and **one accent colour** for actionable things. Colour is used to signal, not to decorate. Colours are not written by hand in components: they come from the themes (section 6.2) |
| **Icons** | A single library, always with text beside it in important actions |
| **Motion** | Almost none: short transitions when opening a dialog or collapsing the menu, and nothing more |

**The colour of states** (always accompanied by its text, never alone):

| State | Intention |
| --- | --- |
| `nuevo` | Draws attention: it has just arrived |
| `en progreso` | Neutral: someone has it |
| `en espera` | Amber: blocked, waiting for someone |
| `escalado` | Different from everything else: it is with another team |
| `resuelto` | Soft green: done, pending closure |
| `cerrado` | Muted grey: out of sight |

### 6.2 The themes: eight, and only the colours change

The shape of the controls **never changes**: a button is the same button in all themes. What
changes is the palette. That is what makes eight themes not eight different interfaces.

| | How many | What they are |
| --- | --- | --- |
| **Light** | 3 | Three fixed light palettes |
| **Dark** | 3 | Three fixed dark palettes |
| **Customisable** | 1 light and 1 dark | **The primary colour** is changed: the institutional one, to the taste of whoever administers |

- **How it is implemented**: colours are **CSS variables** (`--color-primario`, `--color-fondo`,
  `--color-texto`, the state colours…). A theme is a set of values for those variables, and
  no component knows which theme is active. Adding a ninth theme is writing nine lines.
- **The two customisable ones are the defaults.** A new installation starts with the customisable theme
  **light or dark according to what the operating system asks for**, always with the institutional
  colour: so the first thing everyone sees is the house's face, without anyone having to
  configure anything.
- **The primary colour is set by the administrator**, and only by them: it is the institution's colour, and if each
  person changed it, it would stop being so. What each person can change is **the theme**, choosing
  among the eight from the header.
- **Choosing is optional**: whoever touches nothing keeps the installation's. Each person's choice
  is remembered **in their browser**, which is the simple thing: no need to touch the database or
  add a profile screen.
- Each theme's palette has to **pass the contrast** of section 8, including the state
  colours. A beautiful theme that cannot be read does not enter.

**How it is implemented today** (2026-09-23, with what is built):

- **An attribute on the `html` element**: `data-theme="claro"`, `data-theme="oscuro"` or **none**,
  which is “follow the system”. The palettes are written **once** with CSS's `light-dark()`
  function and the attribute only changes `color-scheme`: so there are not two colour lists that can
  become unsynchronised, and adding a new palette is a few lines, as the table above says.
- **The choice is remembered in the browser** (`localStorage`), which is the simple thing and what is already
  decided: it does not touch the database or need a profile screen.
- **The theme is applied by our code, never by a `script` written in `index.html`**: the production
  CSP is `script-src 'self'`, without `'unsafe-inline'`, so an inline script **would not
  work in production** and widening the CSP for this would be trading security for convenience. It is
  applied before the application starts, from `main.ts`.
- **The switch shows two options: light and dark. Nothing more.** **“Automatic” is neither shown nor
  chosen** (owner's decision, 2026-09-23): following the system is **not having chosen**, which is the
  factory state, and that is why a third option is not needed, which would also be the only way to
  go back. As long as nobody touches anything, the switch **shows checked what is being seen**
  —light or dark according to the system—, not an option that does not exist; as soon as someone chooses, they keep
  what they chose and the system stops governing. If the system changes with the page open and nobody
  has chosen, the application changes with it, without reloading.
- **The six fixed palettes** arrive with the complete shell, when their names and
  colours are approved. As long as the shell header does not exist, the switch lives **next to the language one**.
- **The factory palette follows the operating system** until someone chooses something else, which is
  decision 5 of the record.

### 6.4 The brand: the logo and the installation name

The brand **belongs to the installation**, not to each person, and an Administrator changes it from
Settings (`docs/modules/settings.md`, sections 5 and 5.7). They are **the logo and the name**, and they are
shown together and always the same:

| Where | The logo | The name |
| --- | --- | --- |
| **Entry screen** | Above the form, centred, **up to 128 px on mobile and 160 px on tablet and above**: it is the first thing seen | Below the card title, where the form is: it is the first thing read |
| **First start** | Next to the header, **80 px high**: it occupies the same visual block as the title and its explanation without competing with the step's content | The name is not shown here yet: it is being configured |
| **Sidebar** | At the top, next to the name, **up to 32 px high** | At the top, next to the logo |
| **Browser tab** | — | It is the tab title, on all screens |
| **Settings** | The two slots with their preview and the upload and return-to-factory buttons | The field, with its help: if left empty, the factory name returns |

**And with the brand goes the system version**, which is not configured: the code writes it
(`docs/modules/settings.md`, section 5.9).

| Where | The version |
| --- | --- |
| **Sidebar** | In the sign-out row, **right-aligned**, in small, muted text |
| **Entry screen** | **Below the controls**, centred and in small, muted text: it is the footer of the screen |
| **Nowhere else** | It does not go in emails, nor in the tab title, nor on product screens |

- **It is written `v1.0.0`, with the `v`**: that is how a version is read, and without it a loose `1.0.0`
  looks like anything else.
- **It is not a link or a button**: it leads nowhere. It is read and that is it.
- **On the entry screen it does make sense** (owner's decision, 2026-09-25): it is where
  someone who cannot sign in looks to say which version they are using, and the entry screen is the
  only thing they see.
- **The build date is not shown**: it would change on every build without saying anything new, and two
  installations with the same code would say different things.

- **The name does not carry the screen's**: the tab is called like the installation and nothing more (all
  routes carried the same text, so they did not say where you were).
- **It is counted in characters and does not exceed 60**, which is what fits in the menu and in the entry
  card; the field cuts it while typing and the backend checks it the same.

**The system has to work for any logo**, and that is solved in how it is shown, not by asking
for a specific file:

- **It is not cropped and not stretched**: it is shown whole, with its proportion, within the maximum height.
- **It goes in a box with rounded corners and a theme border.** Thus a logo with an **opaque
  background** —the factory one, for example, is a black square— looks like a deliberate piece, and one
  **with transparency** integrates into the background. Neither is assumed.
- **In the controls bar, the theme selector goes without a visible label** (owner's
  decision, 2026-09-24): the dropdown already shows the theme that is set, so the word
  “Theme” only took up space. **The accessible name remains there**, hidden from view: a screen
  reader has to be able to say what the control is for (section 8).
- **There are two slots, light and dark, and both are optional** (section 6.2 and
  `docs/modules/settings.md`, section 5.1): if there is only one, it is used in the eight themes. **And the two
  factory logos are also two, one per theme** (2026-09-26), so an installation without its own logo
  reads just as well in light as in dark.
- **It does not go in emails**: emails are HTML without images (`docs/modules/mail.md`).

### 6.3 Inventory of custom components

What has to be built, and nothing more:

`button` (**four forms** —primary, secondary, danger and the green confirmation one— **and it can
be only an icon**, which then needs its accessible name; **the icon is one of the project's**
—`app-icono`, `frontend/src/app/shared/components/icono.ts`, custom SVGs, never a character or a
library—) · `text field` · `text area` · `selector` · `checkbox` · `state label` ·
`table` (with sorting and pagination) · `card` (the list on mobile) · `dialog` · `menu` ·
`segmented switch` (the double view and the filters) · **`formatted text editor`, with the
attachments inside** · `attachment` (with preview) · `comment` (with edited and deleted marks) ·
`notice` (success, error, information) · **`toast`** (fixed result in the window) · `empty state` ·
`loading indicator` · `sidebar` (with its three zones) · `mobile bar`.

Twenty pieces. There were twenty before **`file upload with drag and drop` was removed**
(2026-09-26); when it was integrated into the editor there were nineteen, and the toast approved on 2026-10-06 brings
the inventory back to twenty. Attaching stopped being a separate box —now **it is attached inside
the text**, with the editor— and having two places where the same thing is done was the noise this product
avoids. What it did —drag, choose and filter by extension with its notice— **the editor does**,
with the three doors of section 6.5.

Each one is built **once**, with its states (normal, hover, focus, disabled, loading) and its
keyboard behaviour.

**Done today**: the button, the text field, the **text area**, the selector, the segmented
switch, the card, the **notice**, the **toast**, the **dialog**, the **state label**, the **attachment with its
preview**, the **formatted editor —with the attachments inside, section 6.5—**, the sidebar with
its three zones and the mobile bar. What is missing —the comment as a separate piece and the loading
indicator— arrives with the screens that use it.

- **The dialog is a native `<dialog>`** (2026-09-24): it closes with **Escape**, traps focus and
  **returns it to where it was** without having to program it. It is what account creation uses.
- **The card chooses the level of its title** (2026-09-24): `h1` when the card *is* the screen —the
  session ones— and `h2` when it goes below the title of a product screen. Without that, the user
  screen had three `h1`, and a page like that is not navigable with a screen reader (section 8).
- **Attaching has three doors and a single rule** (2026-09-25, and since 2026-09-26 with the editor):
  **paste**, **drag** and **the button**, and all three **put the file where the cursor is**. The
  same is checked —the closed list of extensions— and the same is warned about what is not valid
  (`docs/modules/tickets.md`, sections 2.3 and 4).
- **The selector can show an option that cannot be chosen** (2026-09-24), with its help text
  below: that is how directory origins are shown today, which are planned and do not yet
  work (section 3.6).
- **All dropdowns keep a native `<select>`**, but draw their indicator with the text
  colour and separate it from the right border. The text reserves that space, so it never ends up
  below the arrow.

### 6.5 The writing box: a formatted editor with attachments inside

**Writing a ticket or a comment is not filling in a text field**: what is written carries
formatting and **the files attached are seen inside the text**, in the place where they are put
(`docs/modules/tickets.md`, section 2.3). A normal text box cannot carry an image inside,
so **the writing box is a formatted editor**.

- **The editor bar**: bold, italic, underline, strikethrough, bulleted list, numbered list,
  link and **Attach**. The buttons act **on what is selected** —it is the same rule as the
  mail editor, section 3.8— and **Attach puts the file where the cursor is**.
- **Three doors for attaching, and all three put the file in its place**: **paste** (`Ctrl+V`),
  **drag it over the editor** and **the button**. What is in the closed list of
  extensions is accepted, and what is not is rejected with the usual notice.
- **When pasting text, the text enters without formatting** (2026-09-26): what is pasted remains as plain text
  and **line breaks are respected**. Pasting the HTML of an email or Word would bring tags that are not
  on the allowlist, and the comment would be rejected entirely because of something the person did not write
  —a `style`, a `<font>`—, which is the worst way to lose a text. A **file** pasted (a
  screenshot, what is copied from the explorer) does enter, and enters as an attachment. It is the same rule of
  `docs/modules/tickets.md`, section 2.3.
- **A link is requested with its address and checked before putting it in**: it has to start with `http`,
  `https` or `mailto`. What does not, is not put in and the reason is given, instead of rejecting the comment on
  saving (it is the same rule the backend applies).
- **What is going to go out is seen**: the image and the video **are really seen** while writing, not as
  a text that represents them. A PDF or a Word is seen as **its link**, with its name.
- **Calling someone is written in the comment**: the editor has a **“Tag”** button —and
  typing `@` opens the same—, which shows **the list of active technicians and developers** and puts the
  chosen name **where the cursor is**, highlighted as a **mention**. Whoever reads it sees who is
  being called, and that person becomes a **watcher** of the ticket: they appear in their “Watching” and in the
  watcher list of the details panel. **Only Support and Development can tag** (the user does not see the list of
  technicians), and whoever is tagged **is notified by email**. In the details panel, each watcher carries their
  cross to remove them, and removing them **is recorded in the timeline**.
- **With a cap, and the click opens them large**: in the text —writing and reading— the image and video
  are seen **at most at 480 × 360** (decision 55, of 2026-09-26), without deforming: what it cannot do is
  let a full-screen screenshot eat the conversation. **The cap is `min(100%, 480px)` and not
  `480px`**: a `max-width` only limits, so with the fixed value a 480 image **overflowed the
  screen of a 412 px mobile** and widened the entire document —the Playwright layer found it,
  because clicking a button on the same card began to fail with horizontal scrolling—. Now it shrinks with
  the screen instead of overflowing it. **When clicked they open in the viewer**,
  larger, with **Download** and **Open in a tab**. For that, the frame of what is seen is **a
  real `button`** —it is reached with the tab key, it is pressed with Enter and it says what it is—, and the video carries
  **its play button** on top: an image that is clicked and does not say it is clicked is an image that
  nobody discovers with a screen reader (section 8).
- **Video plays in the viewer**, not in the text: in the conversation **a thumbnail** is seen
  —without controls— and clicking it opens the viewer, which is where it is seen large, played and
  downloaded. **While writing it does carry controls**, because there is no viewer there and what is being
  done is checking what has just been attached.
- **The player controls are put there by the screen**: what is saved is
  `<video data-adjunto="grabacion.mp4">`, without `controls`, because **that attribute is not on the backend's
  allowlist** (`backend/modules/tickets/services/body.go`): with it, the entire text would be
  rejected. **The size is not saved either**: it is how it is painted. It is said here because **the finding
  was exactly the opposite**: the example in `docs/modules/tickets.md`, section 2.3, wrote
  `<video data-adjunto="x.mp4" controls>`, which is a text that the backend would have rejected; the
  document was corrected in the same work.
- **A file that was named and did not upload leaves its gap marked**: if an upload fails after
  publishing the comment (decision 44 of `tickets.md`), the place where the file was going **is seen**, with its
  name and a notice, instead of leaving a mute gap that looks like a screen failure. When the
  file uploads, the gap becomes the image.
- **Removing an attachment is removing it from the text**: it is deleted as anything that has been
  pasted is deleted. And **a file that is not in the text is not uploaded**, so there is no need to remember
  to remove it in two places: there is one.
- **The separate file box disappears** from these forms (`app-archivos` is removed, and the
  inventory of section 6.3 reflects it): having two places where the same thing is attached was exactly the
  noise this product avoids.
- **What is written is counted in text, not in HTML**: the tags and addresses of the
  attachments are the editor's and the screen's business, and **nobody sees them**. That includes the inbox
  search, which **removes the tags before searching** so that searching “p” does not find all the
  tickets with an image.
- **And what is sent carries only what the backend admits**: an attachment goes out as its `data-adjunto` and
  **nothing more** —neither the preview's `blob:` address, nor its `class`, nor a `style`—, because the backend's
  sanitiser **rejects the entire text** with `422 tickets.body.notAllowed` as soon as something appears
  that is not on its list. When that happens, the screen says it in Spanish (“El texto lleva algo
  que no se admite…”), it does not show the key. **Video controls are not saved either**: the
  screen puts them there when painting it.
- **When reading, the same**: the conversation paints the text with formatting and, in its place, the image or
  the video player; the PDF and the other files, as links. **Clicking an image or a PDF opens
  the viewer** —the modal, larger, with **Download** and **Open in a tab**, decision 48 of
  `tickets.md`—, and clicking a Word or Excel **downloads it**.

## 7. Multi-device

**The priority is the PC**: it is where work is done, where there are two columns and where the complete
inboxes are. Tablet and mobile have to **really work**, not be a promise.

| | Mobile (< 768) | Tablet (768–1200) | PC (> 1200) |
| --- | --- | --- | --- |
| **Menu** | Button in the header | Collapsed, opens if wanted | Fixed, collapsible |
| **Lists** | Cards | Reduced table | Full table |
| **Detail** | One column; the details panel goes below | One wider column | Conversation + details panel on the right |
| **Double view** | Switch, one at a time | Switch, one at a time | The two columns |
| **Administration** | Consult and simple actions | Everything | Everything |

- **Touch targets of 44 px** minimum on mobile: a button that cannot be hit with a finger is a
  button that does not exist.
- **No horizontal scrolling**: if a table does not fit, it becomes cards, it is not
  dragged sideways.
- **What is done on mobile** (decided, section 11): **everything except administration and the double
  view**, which are what is left for the PC.

## 8. Accessibility and trust

- **Sufficient contrast** in text and state colours, in light and dark.
- **Everything can be used with the keyboard**: focus always visible, coherent tab order and nothing that
  only works with the mouse.
- **Never colour alone** to say something: the state carries its text and notices carry their icon.
- **Real labels** on fields, not sample texts that disappear when typing.
- **Errors say what happened and what to do**, in the language of the reader. A “Error 422” is not a
  message. The backend sends a **key** (`auth.invalidCredentials`) and the interface translates it: that is why
  having a dictionary of screens is not enough —**each error key needs its text in both
  languages**, and a key without text is an error without a message.
- **Dialogs close with Escape** and do not steal focus without returning it.
- **Attachments are announced** with their name and size before downloading them.

**Since the components are ours, this does not come built in.** What has to be done well once and
reused:

| Component | What cannot be forgotten |
| --- | --- |
| **Dialog** | Focus inside on opening, returned to where it was on closing, Escape to close and the background unable to be touched |
| **Menu** | Arrows to move, Enter to choose, Escape to close |
| **Segmented switch** | It is announced as a group of options, not as loose buttons |
| **Table** | Real headers, announced sorting and keyboard navigation through the rows |
| **File upload** | It can be done without a mouse, and each file says what happened to it |
| **Notices** | They announce themselves to screen reader users, and do not steal focus |

## 9. The states that are not screens

Almost nobody designs this and it is what is suffered most:

- **Empty list**: it says why it is empty and what to do. “You have no tickets yet” with the create
  button beside it, not a blank table.
- **Empty inbox**: “There is nothing pending” — and let it be true, without inventing a motivational message.
- **Loading**: a discreet indicator and only if it takes time; no loading screens for half a second.
- **No permission**: if someone reaches a screen that is not theirs, they are told, not shown a
  blank screen.
- **Something failed**: what was being attempted, what happened and what they can do. And the way back always visible.
- **No connection or server down**: a clear notice, without technicalities.

## 10. Small details that are noticed a lot

- **The ticket number, always visible and copyable.**
- **The last update date** in the list and in the detail: what has been stopped for days is seen.
- **Dates** are read `AAAA-MM-DD hh:mm`, in 24-hour format and in the same order in both languages;
  the ISO `T` or `Z` from the backend is never shown.
- **No automatic refreshes or live counters** (there is no real time, `docs/modules/tickets.md`): the
  list is updated on entry and with a refresh button. A counter that does not update itself is
  worse than not having it.
- **Image and PDF attachments are previewed** inside the ticket; the rest are downloaded
  (`docs/modules/tickets.md`, section 4).
- **Emails are HTML, with an automatic text version** (`docs/modules/mail.md`): same tone
  as the interface, without jargon and with inline styles, without external images or third-party fonts.

## 11. What was decided when reviewing this document

| # | Decision | It was left like this |
| --- | --- | --- |
| 1 | **The visual base** | **Tailwind CSS v4 with custom components** (20 pieces, section 6.3), not a library |
| 2 | **The double view** | Two columns on PC with **“Both” by default** when the internal one exists; switch on mobile |
| 3 | **What is done from mobile** | **Everything except administration and the double view** |
| 4 | **The themes** | **Eight: 3 light + 3 fixed dark, plus 1 light and 1 dark customisable**, which are the defaults. **The primary colour is set by the administrator**; each person chooses a theme. Only colours change, never the shape |
| 5 | **Light and dark mode** | Yes from 1.0.0, following the system until someone chooses |
| 6 | **Accessibility** | **Requirement**: contrast, keyboard and focus on all screens, with AA as reference, and checked when building each component |
| 7 | **The theme, today** | Light and dark (the two customisable factory ones), with the mechanism ready for the six fixed palettes. The factory one follows the system, which is the same as **not having chosen** |
| 10 | **“Automatic” is not chosen or shown** | Owner's decision (2026-09-23): the switch has two options and while nobody chooses **it shows checked what is being seen**. Following the system is the factory state, not a third option |
| 11 | **The system governs live** | If it changes from light to dark with the page open and nobody has chosen, the application changes with it without reloading |
| 8 | **Where the switch lives while there is no header** | Next to the language one, in the entry and at the start: it moves to the header controls when the shell exists |
| 9 | **The theme is not applied with an inline script** | The production CSP is `script-src 'self'`: it is applied from `main.ts`, before starting |
| 13 | **A comment without text is read through its attachments** | **Owner's choice, 2026-09-25**, when deciding that a comment can be only a file: in the conversation it appears with its author, its date and its files, **without a blank gap** that looks like a screen failure. And **if a file does not upload after publishing the comment**, it warns which one is missing and it can be retried without writing it again (`docs/modules/tickets.md`, decisions 41 and 44) |
| 12 | **The system version** | **Owner's choice, 2026-09-25**, and **not the one I recommended**: it is read in the sign-out row of the sidebar, **right-aligned**, and **also in the footer of the entry screen**. It is a datum, not a button or its own row, and **collapsed it is not shown**. I proposed only the menu; he decided that it should also be seen before signing in, which is where it serves to say which version is being looked at |

## 12. What this document enables

With `docs/interfaz-y-experiencia.md` approved, the documentation covers **what is built, how it
behaves and how it looks**. The code can begin with `auth` (`docs/modules/tickets.md`), and its screens
—sign in, recover password— are designed already with these rules, instead of retouching them afterwards.

## 13. Back navigation on the left

### 13.1 Finding and inventory

The rule approved on 2026-09-29 already says that back buttons go on the left and are an
icon-only arrow with an accessible name. The inventory of views with that header control
found:

| View | Destination | Position found | Complied with the rule |
| --- | --- | --- | --- |
| `/mail` | Settings | To the right of the title | No |
| `/tickets/new` | Ticket list | To the left of the title | Yes |
| `/tickets/:número` | Ticket list | To the left of the number | Yes |
| `/users/:id` | User list | To the left of the name | Yes |

The comment in `correos-page.html` itself also claimed that the arrow was to the left
of the title, but the `justify-between` structure placed it after the title and at the right
edge. It was an inconsistency between the document, the comment and what was painted.

### 13.2 Rule

Every control that **returns from a secondary view and lives in its header** appears as the first
element of that header, immediately before the title or main identifier. It is an arrow
of only an icon, keeps an accessible name and maintains its current destination.

This rule applies to the four inventoried controls. Three already complied; the implementation
reorders the header of `/mail` to place its link to Settings before the title and
description block. It does not change the destination, the accessible text, the icon or the size of the control.

Out of scope:

- the links that are part of a screen's flow, such as **Back to sign in**;
- the **Back to my work** button of the no-permission view;
- actions with another purpose that use the word “back”, such as restoring a logo or an email;
- adding new buttons to views that today have no back navigation;
- the backend, the routes and the permissions.

### 13.3 Documents and code

- This document defines the general rule and records the finding.
- `docs/README.md` and `AGENTS.md` reflect the final state of the amendment.
- `frontend/src/app/modules/mail/correos-page.html` places the arrow before the title and
  description block.
- The E2E test of the mail editor checks that the arrow is the first element of the header,
  before the title, and that it still returns to Settings. The suite runs it on PC and mobile.

`docs/modules/mail.md` does not change: it describes the editor's functions and does not fix the composition of its
header. Moreover, its functional behaviour does not change.

### 13.4 Acceptance criteria

1. In `/mail`, the arrow appears immediately to the left of the title **The emails**.
2. At PC and mobile widths, the control remains before the title when the header is distributed over
   more than one line.
3. Its accessible name remains **Back / Volver** and its destination remains `/settings`.
4. The three controls that already comply keep their position and destination.
5. The excluded links and actions do not change.

### 13.5 Recorded decision

The owner confirmed on 2026-10-05 the recommendation: unify only the header navigation
controls as an icon to the left of the title. Flow links inside the content remain
out of scope.

### 13.6 Review record

The owner confirmed on 2026-10-05 that:

1. the arrow would remain as the first element of the header;
2. the title and its description would remain together, in a block to the right of the arrow;
3. no additional exception would be added to the rule nor would the scope be widened.

The review is closed with no open decisions. The owner explicitly approved the proposal
with “Yes”. The implementation respects the approved scope and the document returns to `as-built`.

### 13.7 Verification

- The 216 frontend unit tests pass.
- The four E2E cases of the mail editor pass: two on PC and two on mobile.
- The test checks that the arrow is the first element of the header, that its horizontal position
  remains before the title and that pressing it returns to `/settings`.

## 14. Selected time zone visible in Settings

### 14.1 Finding

The **Time zone and public address** card of `/settings` offers the same search and the same
scrollable list as `/setup`. The chosen option is only recognised by the background and text weight
inside that list; if it falls outside the visible part, the screen does not name in a stable way the
zone that will be saved. The time and offset make it possible to check it, but they do not say which one it is.

### 14.2 Behaviour

The pattern approved for `/setup` is reused exactly:

- a read-only field before the search;
- label **Selected time zone / Región horaria seleccionada** according to the interface language;
- exact IANA identifier as the value;
- zone saved on loading the screen and immediate update when choosing another;
- selectable and copyable text, without direct editing;
- time and offset kept below the list.

The field's value is the form's current draft. Choosing a zone changes it and enables the existing
saving; saving, receiving an error or reloading follow the current Settings rules. The search, the list, the
highlighting, the public address and the save button do not change.

### 14.3 Scope

- `docs/modules/settings.md` defines the behaviour inside the time zone card.
- This document defines its presentation, consistent with `/setup`.
- `docs/README.md` and `AGENTS.md` reflect the final state of the amendment.
- `frontend/src/app/core/pages/settings-page.html` shows the field and the existing Settings tests
  check its behaviour.
- The current ES/EN texts are reused; no translation keys are added.

Out of scope are the API, the database, zone validation, `/setup`, the public address and the
other Settings cards.

### 14.4 Acceptance criteria

1. On loading `/settings`, the field shows the saved zone.
2. Choosing `America/Guayaquil` updates the field immediately before saving.
3. The field is read-only and allows selecting and copying the identifier.
4. The search, the list, the highlighting, the time and the offset keep their behaviour.
5. The unit tests check the loaded value, the immediate change and the read-only
   attribute; Playwright checks the presence of the field on PC and mobile.

### 14.5 Recorded decision

The owner confirmed on 2026-10-05 reusing exactly the `/setup` pattern, without changes to
saving, the API or the database.

### 14.6 Review record

The owner confirmed on 2026-10-05 that `/settings` will use the same technical identifier, the
same appearance and the same behaviour as `/setup`, without additional differences.

The review is closed with no open decisions. The owner explicitly approved the two
proposals with “Yes”. The implementation respects the approved scope and both documents return to
`as-built`.

### 14.7 Verification

- The 216 frontend unit tests pass; the Settings ones check the initial value,
  the immediate change and the read-only attribute.
- The four affected E2E journeys pass: two on PC and two on mobile.
- Playwright checks the saved value, `America/Guayaquil` after choosing, the read-only
  attribute and that the field belongs to the time zone and address card.

## 15. Action results in toast messages

### 15.1 Finding

`/settings` and `/setup` have a single screen message placed before their content. Saving
a card or testing a connection from a lower section updates that message, but the window
does not return to the top. The action does respond and the text does exist; it is outside the visible part.
This explains why a failed connection test seemed to show no result.

### 15.2 Component and behaviour

A shared **toast** component is added for the transitory results of an action:

- it remains fixed in the window, at the top right on PC and with the available width, leaving a margin, on
  mobile;
- only one is shown at a time and the most recent result replaces the previous one;
- all forms disappear after **10 seconds at most**;
- all can be closed earlier or replaced by a later action;
- all include a close button with a translated accessible name;
- success and information are announced as status; error and attention are announced as alert, without moving
  focus or shifting the content;
- it keeps the four visual forms of the existing notice —success, error, information and attention— and
  uses the colours, icons and contrast of the active theme.

The timer starts again when a different result arrives. Destroying the screen cancels
the pending timer, so that an old message does not act after navigating.

### 15.3 What is a toast and what remains inside the page

The results of asynchronous actions that today feed the general message of
`/settings` or `/setup` use toast: save, upload or restore, advance in the wizard and test a connection.
Their loading or validation failures returned by the API also use it when they do not belong to a specific
field.

Notices that describe a state or a consequence for as long as that
condition exists remain inside their section, including the public address without HTTPS, the explanation of the factory account
and the informative notices of a step. Errors associated with a field keep their place
next to the field if a screen presents them that way.

### 15.4 Scope

- This document defines the shared component and its accessibility.
- `docs/modules/settings.md` defines its use in `/settings`.
- `docs/primer-arranque.md` defines its use in `/setup`.
- `docs/README.md` and `AGENTS.md` reflect the state of the amendment.
- The implementation lives in `frontend/src/app/shared/components/**` and in the two affected pages.

Out of scope are the backend, its responses and codes, persistence, validation rules,
the messages of the other screens and a global notification service for the whole application.

### 15.5 Acceptance criteria

1. Testing a connection from a lower card or step shows its success or error inside the visible
   part of the window, without scrolling it to the top.
2. Saving, uploading, restoring and advancing show their results through the same component.
3. All forms disappear after 10 seconds at most and can be closed earlier.
4. A second result replaces the first and restarts the corresponding timer.
5. The close button works with the keyboard, has an accessible name in Spanish and English and the message
   keeps the appropriate semantic announcement.
6. Persistent context notices remain inside their section.
7. The unit tests cover replacement, closing and duration; Playwright checks the visible
   position of the connection results in `/settings` and `/setup`, on PC and mobile.

### 15.6 Recorded decision

The owner confirmed on 2026-10-06 proposing the shared component for `/settings` and
`/setup`: fixed at the top right on PC and adapted to the mobile width, one message at a time, success with
automatic closing after 10 seconds for all forms, manual closing and context notices still
integrated into the page.

### 15.7 Review record

The owner confirmed on 2026-10-06 that:

1. all the transitory results described for `/settings` and `/setup` use the toast;
2. there is no queue: the new result replaces the previous one;
3. all forms disappear automatically after 10 seconds at most and can also
   be closed or replaced earlier.

The review is closed with no open decisions. The owner explicitly approved the three
documents with “I approve” on 2026-10-06. The implementation respects the approved scope.

### 15.8 Implementation and verification

- `app-toast` wraps the shared notice, remains fixed in the window and emits manual or
  automatic closing; when destroyed it cancels its timer.
- `app-aviso` accepts an optional close button. Integrated notices do not receive it and keep their
  presentation; error and attention use `alert`, success and information use `status`.
- The 220 frontend unit tests pass; four new cases check closing, semantics,
  duration, persistence and timer restart.
- The production build finishes correctly.
- The 20 selected Playwright cases pass on PC and mobile; 4 are skipped because the optional
  directory and AI services were not running. The empty-installation preparation
  also checked the `/setup` toast at both widths.

## 16. Implementation: global language and guided AI configuration

### 16.1 Language is the first control

In `/setup` and `/settings`, Español/English appears before any other datum. In the wizard
it changes all texts when chosen. In Settings it previews the screen, but a banner
explains that the global change has not yet been applied while emails are being translated and reviewed.

All personal language selectors disappear: entry, shell controls, profile, creation
and user record. The space freed in the menu is not filled with another control. The theme continues
to be personal.

### 16.2 AI modality selector

The card and the step use three large options with a name and explanation: **On this server**,
**On another server** and **External provider**. Only the fields of the chosen option are shown. The
external modality includes a persistent notice with the data that will go out and an acceptance checkbox
before the **Test and activate** button.

For provider, OpenAI, Claude, DeepSeek and OpenAI-compatible are shown. The model is a
dropdown of suggestions with **Another model…**; when chosen, the free identifier appears. Secrets
use a password field, indicate whether one is saved and never show dots that could
be confused with their real value.

### 16.3 Local catalogue and progress

Each model shows name, download size, recommended RAM and state: not installed, downloading,
installed, active or error. The current 1.5B appears first and marked as recommended. Before a
download, available disk and RAM are shown and it is blocked with an explanation if they are not enough.

The card links the official source and shows licence and checksum. Before the first download of a
version the licence opens and an explicit acceptance is requested; the screen says who accepted it and
when. A change of version, licence or checksum requests it again.

The download has a progress bar, downloaded/total bytes and a cancel action. Reloading keeps
the real progress returned by the server. Activating shows the stages “loading”, “testing” and “active”.
Inactive models have **Delete** with confirmation; the active one never offers that action.

On mobile the cards stack, the bar takes the width and the actions remain below. No state
depends only on colour; progress and result are announced accessibly. Brief results
use the existing toast, while download, requirements and active state remain inside the
card.

### 16.4 Review of the language change

The review presents the fourteen templates as a list of pending items. Opening one shows origin and
draft side by side on PC and stacked on mobile, with subject, body and placeholders. If a custom
destination already exists, it appears as a third version and the Administrator chooses to keep it, use the
draft or combine them in the editor. It is not replaced silently. **Apply language** is only enabled
when all twelve are valid; cancelling discards the drafts and returns to the persisted language.

After confirming, the entire interface adopts the new language. Previous summaries are kept and
remain visible in the language in which they were written. New ones and those requested again are
generated in the new global language; there is no mass regeneration when it changes.

The economic warning appears only before translating the templates with an external
provider and shows estimated requests, tokens and cost. Rejecting it opens the same review with
manually editable destinations. Manual editing remains available at all times from Settings.

Sessions that were already open receive the language and configuration version in the next
authenticated response. The interceptor updates the global signal and repaints texts and dates without
closing the session or polling periodically. Until that tab makes another request it keeps the
view it already had.

### 16.5 Proposed acceptance criteria

Playwright will cover the three modes, hidden secret, privacy notice, failed test that keeps
the previous configuration, download/progress/resumption, licence, insufficient resources,
activation and deletion, review with a custom destination, commercial confirmation, global change,
updating an open session and absence of personal selectors, on PC and mobile.

### 16.6 Review record

The owner closed the review on 2026-10-06. They chose a mandatory encryption key per environment,
comparison of three versions so as not to overwrite custom templates, commercial recovery under
the Administrator's decision, visible licence and integrity and updating the language in open sessions
without signing out or polling. No open decisions remain; proposal approved.

## 17. Coherent memory when activating models

The selectors and cards show minimums of **2 GiB, 4 GiB and 6 GiB** for 1.5B, 3B and 7B. They do not mix
decimal GB with GiB nor do they present an estimate lower than the gate the administrator applies.

When a model does not fit, the toast remains until closed and says which one was attempted, how much
RAM it requires and how much was available, with values formatted in GiB. The previous model continues
marked “Active · Engine available”. The message works the same in `/setup` and `/settings`, on PC and
mobile, and does not recommend forcing swap.

The real check on PC and Pixel 7 showed 2.00/4.00/6.00 GiB. When trying the 7B, the toast remained
visible with 6.00 GiB required and between 5.06 and 5.08 GiB available, while the 3B continued marked
“Active · Engine available”.

## 18. Full formatting review

### 18.1 Scope

Block 6 reviews all existing screens: session, `/setup`, shell, settings, mail,
users and tickets. It includes empty states, lists with content, forms, dialogs, toasts,
attachments, double views and visible errors. Spanish and English are checked in Chromium at
**1440×900**, **768×1024** and **412×915**.

The review covers spacing, alignment, hierarchy, legibility, cut-offs, overflows, responsive
adaptation, visible focus, keyboard navigation, accessible names, semantics, contrast and touch
areas according to the applicable criteria of **WCAG 2.2 AA**. The eight themes keep their automated
contrast test; the factory light and dark themes also receive the detailed visual journey.
The six fixed themes are checked through their variables and common tests.

### 18.2 Inventory and priority

Each finding is recorded before being corrected with screen, language, theme, resolution, way of
reproducing it and measurable evidence. Priority means:

- **critical**: blocks a task;
- **high**: prevents using it on a device or without a mouse;
- **medium**: hinders reading, understanding or consistency;
- **minor**: perceptible cosmetic defect.

The chosen closure criterion is **zero open findings of any level**. If a change appears
that affects behaviour, scope, data, permissions or flow, that correction stops and another
documentary cycle opens; it is not reclassified as formatting to introduce it in this block.

### 18.3 Correction and evidence rule

When the defect repeats, the shared component or style is corrected and all its
uses are verified. A local adjustment is reserved for a composition specific to a single screen. Reproducible
measurements and assertions are preferred over reference screenshots: width, position, overflow,
focus order, accessible name, contrast and target size are automated when they provide a stable
signal. The manual journey remains documented with its result.

No visual screenshot library is introduced, no Firefox or WebKit is added, no journeys are
redesigned and backend, API, persistence, permissions and product rules do not change. Production remains
parked until the closure of version 1.0.0.

### 18.4 Decisions and review

The owner chose 1A–4A: all screens; formatting, accessibility and responsive; PC, tablet and
mobile; complete inventory and correction by priority. They chose 5A–9A: eight themes with a detailed
journey in the two factory ones; both languages; Chromium; measurable evidence; and common correction for
repeated defects. They closed the review with 10A–13A and 14C: applicable WCAG 2.2 AA, four priorities,
disposable data with seeders, no functional changes and zero open findings. No open
decisions remain.

### 18.5 As-built result

The inventory went through the authenticated screens of the four roles and the public screens in
Spanish and English, with the factory light and dark themes and the three approved resolutions. `/setup`
was also checked empty in its four navigable states: there were **48 combinations** of the wizard,
**150 authenticated views and 18 public ones per language**. Legitimate WCAG exceptions for links
in text, controls with sufficient separation and email content were not recorded as defects.

One medium finding was found: the role switch of `/users` measured 424 px in a 412 px window
when it was in English. The shared `Conmutador` component now limits its width and splits
its options over several lines; the correction was verified at the three sizes, both themes and with the
Administrator and Support roles. No flow or contract changed.

The closure left **zero open findings**. The 220 frontend unit tests passed, 18 affected
Settings cases with 4 expected skips and the complete Playwright suite with 180 passed cases
and 38 expected skips. The eight themes kept their measured contrast.

## 19. Implementation: “Improve with AI” modal

### 19.1 Entry and content

The description and comment editors allowed to Support and Development incorporate, next to their
controls, a button with an icon and text **“Improve with AI”**. On mobile it can take its own row so as
not to compress the editor. It does not appear for User or Administrator.

The button requires a non-empty draft and opens a modal with a copy in a simple `textarea`: without
bold, italic, mentions or attachments. Editing that copy does not touch the main editor. Below, one
of five tones is chosen: professional, friendly, brief, empathetic or technical. A note indicates that the
configured engine will process the text and, when it is not local, that it will leave the installation according to
the existing administrative acceptance.

### 19.2 Generation and review

“Improve” starts a single request and remains disabled while it responds. The result replaces
the copy inside the modal so it can be reviewed and edited. From there you can **“Use this text”**,
**“Generate again”** or **“Cancel”**. Regenerating uses the visible content and allows changing the
tone. Using the text replaces the main editor's draft as plain text with line
breaks; it still does not save or publish anything.

Cancelling or closing keeps the main editor intact. If it is closed during a request, a late
response is discarded. A failure remains inside the modal, keeps its content and offers
retry or cancel; it is not shown far from the context through a toast.

### 19.3 Accessibility, adaptation and tests

The modal uses the existing shared pieces, has an accessible title and description, traps and
restores focus, closes with Escape when there is no other pending confirmation and announces loading,
error and result. At 412 px it does not create horizontal scrolling and its actions can split over
several rows.

The unit tests cover the route contract, permissions by role, empty text, invalid tone and
editor, minimal prompt context, response, error keeping the draft and application
without automatic saving. The Playwright journey checks on PC and mobile the modal, the tone, the
review, the local application and the absence of the control for the User. The previous general closure of
the interface keeps both languages and the factory light and dark themes covered.

### 19.4 Decisions and scope

The owner chose 1A–19A and added that the AI receive the name of the person who requires support
on the main ticket or of the Support person on the internal one, to use it only when the wording
requires it. No open decisions remain. The proposal does not introduce new design components,
formats, attachments, persistence, quotas or administrative configuration.

## 20. Implementation: visibility of “Improve with AI”

The details panel only mounts the “Improve with AI” buttons and the shared modal when the detail returns
`capabilities.aiWriting: true`. With `false` or with the field absent, it does not leave a disabled control or
a notice taking up space: the aid simply does not appear. This rule is added to the non-empty draft,
which continues to disable the button once the capability exists.

The screen does not query configuration, health or providers on its own. On reloading the ticket
it receives the capability again and reflects any administrative change. A configured but
temporarily down AI keeps the button; the error only appears if the person tries to generate, inside
the modal and keeping the text.

The component tests cover true, false and absent capability. Playwright verifies on PC and
mobile that a configured installation shows the button and keeps the review journey. The case
without configuration remains in backend and component: a real browser would be directed to Settings
before reaching this view, in accordance with mandatory AI. The modal, its
tones, the editors, the administrative configuration and the global blocking are not modified. The owner
chose 1A, 2A and 3A and confirmed keeping it mandatory; no open decisions remain.

## 21. Implementation: “Source code and licence”

The entry screen shows in its footer a link **“Source code and licence”**, next to the
version without competing with the form. The sidebar does not repeat it: its lower zone keeps
only the account, the theme, the language, sign out and the version.

The entry reads `license` and `sourceUrl` from the public brand. The link opens the corresponding
source in another tab with the usual protections. A development build will indicate it in the text associated
with the version. If an earlier backend does not deliver a valid URL, the
control will not appear. No alternative URL is hard-coded in the frontend.

Accessible name, focus, contrast, absence of overflow and destination were checked on PC and
mobile. It does not add a legal screen, a modal window or an administrative control. The scope and
the brand policy live in `docs/proyecto-abierto.md`; the contract, in
`docs/modules/settings.md`, section 14.

The original verification of 2026-10-10 covered both locations. The textual adjustment of
2026-10-15 removed the menu repetition at the owner's request and kept the link in the
entry; `marca.spec.ts` now expressly checks both conditions.

## 22. Visible and operable priority

Creation offers the four priorities with medium selected. In lists and detail a chip appears
with translated text and icon; colour and icon help, but the meaning never depends on them. The
detail allows Support and, from the authorised internal view, Development to change the value in
place with Save and Cancel. The User and the Administrator see it without an editor. Closed does not offer the
control.

The search modal allows selecting several priorities. The inboxes sort by factory order critical,
high, medium and low, and within each group they show the oldest ticket first. The person can
change the order. Filter and order are reflected in URL parameters and survive reload, return and
shared links. Table and mobile cards show the same datum without creating horizontal scrolling.

The history writes the translated previous and new labels, the author and the date. If the critical
email fails, a persistent attention toast confirms that the priority was saved and that the
notice did not go out; it does not offer to save again or claim that the operation failed completely.

## 23. Progress, local test and logos in `/setup`

The model card shows an accessible native bar, percentage and `downloaded / total`. It is
updated without pressing “View and refresh the catalogue” while a transitory phase exists and stops
querying when finishing or leaving the screen. The text names download, validation, activation and
test; no state depends only on colour. On mobile the bar takes the width and the actions
stack.

Only one download is started. During it **Cancel** appears and the explanation that the partial file is
kept. Afterwards **Resume** and **Delete partial download** appear. The installed inactive model
offers **Delete model** after confirmation; the active one does not offer it. An activation failure keeps
the card and adds **Retry activation**.

In `/setup`, **Next** is disabled until the selected model is active and
healthy. Next to the button it says whether download, validation, activation or testing is missing. The generic
engine error is not used as validation of an action that the screen itself let run too early.

The healthy active model offers **Try me!** or **¡Pruébame!**. While it works it keeps the button with its
accessible state. Afterwards it shows inside the card the model's greeting and, separately, the date,
time and IANA zone that the server supplied. The result does not use toast because it must be readable and
comparable; it disappears when repeating, changing the active one or leaving. The error occupies the same place and allows
retrying.

Step 1 places, after language and name, the two light and dark logo cards that already exist
in Settings. Each shows a preview, factory or custom state, upload and restore. The
preview and the header logo change on saving. Neither variant is mandatory.

The component and Playwright tests cover complete texts in both languages, keyboard, accessible
announcements, PC, mobile, reload, cancellation, resumption, background, error, deletion, dated greeting and the
two logos. General toasts continue to be reserved for brief results; progress, greeting and
operational errors remain in their card.

## 24. Direct actions on cards, visible greeting and upload from the logo

### 24.1 Findings

The current implementation forces returning to the selector located before the catalogue to choose the
model, even though each card already contains its other actions and states. The test greeting remains inside
a long card and can go unnoticed. In the logos, a separate native control shows the
file name and does not make clear that the preview is the place being replaced.

The current policy of section 15 also does not match the new decision: success and information
last five seconds, while error and attention can remain indefinitely.

### 24.2 Models

The top selector of the local model disappears in `/setup` and `/settings`. Each installed card
includes its own **Activate** or **Deactivate** button and a textual indicator of whether it is active. The card
is the only source of selection and keeps download, progress, health, testing and deletion visible.

Deactivating the only active model is allowed. The engine is left without a model; `/setup` disables
**Next** and explains that one must be activated, while `/settings` shows that condition without preventing
administering the rest. Deactivating does not delete the file, so it can be activated again from the
same card. The active model still does not offer deletion.

### 24.3 Greeting and toast

**Try me!** / **¡Pruébame!** shows the result through the shared toast. The message identifies
the name of the model that responded and contains only its greeting. The supplied date and time are already part
of the generated greeting, so they are not repeated, and the IANA zone is not presented. The previous
result stops taking up space inside the card.

The common policy of all `app-toast` becomes: manual closing always available and automatic
disappearance after **10 seconds at most**, including the error and attention forms. A new result
continues to replace the previous one and restarts the timer. Persistent notices that
describe a state remain inside their section and do not become toast.

### 24.4 Logos

The preview of each light or dark logo is an accessible button. Pressing it opens the file
selector; on choosing a valid one, it is saved and the image is updated. The file control is
visually hidden and no path or file name is shown. A separate action is kept to
return to the factory logo and the accessible text indicates which variant will be replaced.

### 24.5 Tests and scope

The component and browser tests cover keyboard and screen reader, activation and
deactivation from each card, absence of the top selector, blocking of **Next**, toast of the
greeting with the model name, closing and ten-second limit for the four forms, and uploading
both logos by pressing their preview on PC and mobile.

The 1.5B/3B/7B catalogue, the exclusion of downloads, the licences, the file limits,
the external modalities and production do not change. The owner chose 1A, 2A and corrected 3A to a maximum of
ten closable seconds for all toasts. No open decisions remain.

The implementation passed the 227 frontend tests and the production build. In the real browser
the cards without a top selector, deactivation and reactivation, blocking of
**Next**, attribution of the greeting without repeating the clock or showing the zone, disappearance of the toast after ten
seconds and accessible logo controls without file name or path were verified.

## 25. Themed confirmations and appearance in step 1

### 25.1 Findings

The four deletion confirmations of the local model administrator use the browser's
`confirm()`. That window belongs to the system, does not receive the theme variables and in a dark theme
appears white. It is the only product confirmation that still avoids the custom component.

The first step allows deciding language, name and logos, but does not present the selector of the eight
themes or the institutional colour that already exist in the shell and in Settings. The installation can
be customised, but not seen complete while it is being configured.

### 25.2 A single dialog for all confirmations

Every application confirmation uses `app-dialogo`, never the browser's `confirm()`. The rule
includes the four current deletions of installed model or partial download in `/setup` and
`/settings`, and any later confirmation. The dialog inherits the active theme's variables,
uses the danger button to confirm and the secondary one to cancel, traps focus, returns focus
to the trigger and closes with Escape. The text remains translated and specifies which resource is deleted.

### 25.3 Theme and colour in step 1

Next to the language selector appears the same complete selector as the shell: eight themes grouped into
light and dark. It changes the interface immediately and saves the preference in the browser through
`ThemeService`; no preference is added to the installation nor an “Automatic” option.

The same block incorporates the institutional colour control from Settings: free selector, hexadecimal
value and preview resolved by the backend. The change is seen before saving. Name, language
and colour are saved together when pressing **Next**. Changing the theme keeps the chosen colour; only the
two factory themes use that accent, while the six fixed themes keep their palette.

### 25.4 Tests and scope

The tests cover absence of `confirm()`, opening, cancellation and confirmation of the dialog in the
eight themes; focus, Escape and ES/EN texts; selection and memory of the eight themes; preview and
persistence of the colour when advancing; preservation of the colour when changing theme; and presentation on PC
and mobile.

The eight palettes, the accessible derivation of the colour, the logo limits, the model
catalogue, the permissions and production do not change. The owner chose 1A, 2A and 3A for confirmations and theme, and
1A, 2A and 3A for colour. No open decisions remain.

The implementation replaces the four existing native calls and leaves zero product
`confirm()`. The 227 frontend tests and the production build pass. In a real browser
theme memory after reloading, presentation at 390 px, dialog in dark theme, closing
with Escape and return of focus to the trigger were verified, without executing the deletion.

## 26. Local catalogue always visible

### 26.1 Finding

In `/setup` and `/settings`, the **Local models** heading appears before querying the catalogue,
but its cards remain hidden until **View and refresh the catalogue** is pressed. The action mixes two
responsibilities: revealing content that should be available from the start and updating it.

### 26.2 Loading and refresh

On opening the AI step of `/setup`, the AI card of `/settings` or switching to local modality,
the screen automatically queries the catalogue. The cards, the available space and their states
are presented without a prior action by the person.

The button is renamed **Refresh catalogue** / **Refrescar catálogo**. Pressing it queries the real
state again, shows **Refreshing…** / **Refrescando…**, is temporarily disabled and
keeps the current cards visible until the response arrives. A correct response replaces
the data in place; an error keeps the previous catalogue and is communicated through the common toast.

### 26.3 Tests and scope

The tests cover automatic loading, switching to local modality, manual refresh without hiding cards,
blocking double presses, error keeping data and presentation on PC and mobile, in `/setup` and
`/settings`.

The 1.5B/3B/7B catalogue, downloads, activation, licences, API, permissions and production do not change.
The owner chose 1A, 2A and 3A. No open decisions remain.

The implementation passed the 228 frontend tests and the production build. The real browser
confirmed in `/setup` that the three cards are already present when entering step 5 and remain
visible after pressing **Refresh catalogue**.

## 27. Two independent ticket panels

### 27.1 Findings

The **Both** view places main and internal inside the same vertical flow. Scrolling
belongs to the page, so moving one conversation also moves the other and the
reference between events is lost. In addition, the detail keeps the global maximum of 88 rem even when there is width
available around it, leaving an unused strip on large screens.

At widths below the two-column point, the current implementation can stack both tickets,
even though the documented rule already asked to see one at a time. The mismatch is corrected explicitly.

### 27.2 PC: own height and scrolling

On PC, **Both** takes the entire useful width of the content area and divides it into two equal columns.
The general header —back, numbers, state and switch— remains visible. Below, each ticket
takes the remaining height of the window and creates its own vertical scrolling region. The wheel,
the gesture, Page Up/Page Down and the bars act only on the column that has the pointer or
focus.

Each panel keeps its header and its comment form inside the scrollable content. The
editor appears at the end of its conversation; it does not remain fixed nor permanently reduce the reading
area. The columns are not resizable and split the space 50/50.

### 27.3 Narrow widths

On mobile and vertical tablet both panels are not stacked. The switch shows Main or Internal,
one at a time, using the full width. The **Both** option only produces the two columns when the
width allows comparing them; when entering or narrowing the window, the role's ticket or the
last individual side chosen is kept.

Individual views keep their current width and scrolling. The change to full width and internal
scroll belongs only to comparison on PC.

### 27.4 Tests and scope

The component and browser tests cover independent scrolling without moving the neighbouring column,
general header visible, equal columns without an unused strip, editors at the end of each
flow, keyboard, wide PC and a single view on mobile and vertical tablet.

Content, permissions, states, actions, API, data model and production do not change. The owner
chose 1A, 2A, 3A and 4A. No open decisions remain.

The implementation passed the 228 frontend tests and the production build. In a real browser,
at 1800 px, both panels measured 728 px, scrolling the main one did not move the internal one
and the common header kept its position. At 900 px the **Both** option disappeared and a
single ticket remained at full width.

## 28. Content at the full width of each panel

### 28.1 Finding

Section 27 correctly widened the outer container and split the useful area into two halves,
but each `app-ticket-vista` keeps inside the `2fr 1fr` grid of the individual view. In
**Both**, the details panel of the right column is hidden by a previous decision, although its fraction continues
to be reserved. That is why the conversation takes only two thirds of each half and wide gaps remain
between both tickets and on the right. The documentation said that the panels used the width,
but did not record this internal reservation; the mismatch is reported before correcting it.

### 28.2 Approvable layout

When `soloLaConversacion` is active, the internal grid will have a single column that takes
100% of the panel. Header, subject, reason, description, comments, history and comment editor
participate in the same wide flow; no intermediate maximum nor a phantom column is added.

The separation between main and internal remains at 24 px. The change is limited to **Both**: the individual
forms keep their current conversation and details panel grid, their widths and their
behaviour. Nor does the independent scroll incorporated in section 27 change.

Image and video previews keep the 480 × 360 limit: the new space does not
enlarge them artificially. The editor takes the new width, but keeps its current height.

The comparison goes from 1280 to 1152 px. From that width **Both** and its two halves appear; below
that, one ticket at a time continues to be shown. The change allows using the comparison on small
laptops without taking it to the typical width of a vertical tablet.

### 28.3 Tests and scope

The component tests will check that `soloLaConversacion` selects a full column and
that an individual view keeps `2fr 1fr`. The browser will verify on PC that messages and editors
reach the useful width of each half, that only the 24 px space between panels remains and that no horizontal
scroll appears. The independent scrolling check will be repeated and the limits of
1151 and 1152 px will be tested, in addition to a view at 900 px.

Texts, content, permissions, states, actions, API, data model and production do not change. The
owner chose 1A, 2A, 3A, 4A, 5A, 6A, 7B and 8B. No open decisions remain.

The implementation passed the 229 frontend tests and the production build. In a real browser,
at 1800 px, each panel measured 728 px and its content 720 px, without horizontal overflow. At
1152 px two panels of 404 px with 396 px of content were shown; at 1151 px **Both** disappeared
and a single view remained. Independent scrolling and individual views were preserved.

## 29. Implementation: “Improve with AI” without a modal

### 29.1 Immediate generation

The **Improve with AI** button remains next to each authorised editor and requires text. Pressing it starts
the request in that place, shows activity and always uses a professional tone. It does not open a modal, does not ask for
prior confirmation and does not show a tone selector. The notice about processing by external
engines remains only in Settings, in accordance with the current administrative acceptance.

Description and comment manage independent states: each one can keep its own
proposal without expelling the other's. During the wait the editor remains available.

### 29.2 Review panel

The response appears below the editor that requested it, in a region with an accessible title. It is plain
read-only text, visible and selectable; no **Copy** button is added. Its actions are
**Use proposal** and **Keep original**. The first replaces the local draft and closes the
panel; the second discards the output, hides the panel and leaves the original intact.

With a current proposal, the button becomes **Generate another proposal**. The new request starts from the
current draft and asks for a professional paraphrase different from the last proposal. Only the
most recent output is kept.

If the person modifies the draft after starting the request, both texts remain, but
the proposal is marked as outdated and **Use proposal** is disabled. It can be copied
by selection, the original kept or another proposal generated from the new draft. Thus a late
response never overwrites recent writing.

### 29.3 Errors, accessibility and tests

An error does not open a panel or alter the draft: it uses the common toast, closable and with a maximum of ten
seconds, and returns the button to its initial state. Loading and the result are announced; focus does not
jump when the panel appears and the actions can split on mobile without horizontal overflow.

The tests cover independent description and comment, loading, selectable result,
application, discard, outdated draft, different regeneration, error through toast, PC and mobile, both
languages and absence of the control due to permissions or configuration. The modal and the
five tones tests are removed. No persistence, automatic saving or additional confirmation is added.

The implementation passed the 230 frontend tests and the production build. The real journey
with Support confirmed inline generation without a modal or tones, accessible region, discard without touching
the draft and blocking of **Use proposal** when the text changes. No comment was published
during verification.

## 30. Implementation: visible quality and refresh without loss

### 30.1 Reason and last action

During a regeneration, each block keeps the previous summary with an accessible indication that
it is outdated and shows activity without replacing it with a placeholder. If it fails, a specific message appears for
**Reason** or **Last action** and a **Retry** button in that same block. When there was never valid text,
the error takes its place; `...`, `—` or an invalid response is not shown as successful content.

Polling and retries update those two blocks in place. They do not empty fields, do not close
editors or remove AI proposals. Focus and selection remain where they were.

### 30.2 Warning of an imperfect proposal

A proposal that still does not reach the minimum quality after the retry appears in the inline
panel with a persistent warning. It remains visible, selectable and can be applied by
express decision; **Use proposal** is not blocked by this warning. The panel distinguishes that
insufficient quality from the outdated state due to a draft change, which does prevent applying it until
regenerating.

The warning does not use toast because it must remain next to the text it qualifies. Network or
engine errors continue to use the common toast. The tests cover accessible reading, keyboard, PC, mobile,
both languages, long texts, intact draft and visual difference between warning, outdatedness
and error.

### 30.3 Redundant texts that disappear

The comment editor does not show the permanent help that lists formatting, dragging, pasting,
file selection and a comment made only of attachments; the controls and the three attachment
doors remain the same and are sufficient by themselves. The English text
**“With formatting, and with the files inside: drag them here, paste them with Ctrl+V or look for
them. A comment can carry only files.”** and its Spanish equivalent are removed.

A valid proposal also does not show the neutral state **“The proposal is ready to review.”** or its
Spanish equivalent: the title, the text and the actions already explain the state. The messages that do change
a decision are kept —loading, outdated draft, insufficient quality and error— and their
accessible announcements do not depend on that visible sentence.

The real verification of 2026-10-15 used `CS-2026-0027`: it showed the two complete summaries, kept
the comment draft and its proposal literally during a retry of **Reason**, showed
the inline warning when the second proposal remained too similar and kept
**Use proposal** enabled. The sidebar stopped repeating the “Source code and licence” link;
the entry still shows it.

## 31. Dropdowns with the same presence as the other controls

### 31.1 Finding

The shared selector component currently has a minimum height of 36 px and less inner
spacing than the text fields and buttons, which use 44 px. Although several native selectors already
reach 44 px, the difference in the common component makes the lists look thinner and without
harmony with the rest of the form. The document did not fix a common height for this control; the
mismatch is recorded before changing it.

### 31.2 Visual rule applied

All application dropdowns, both the shared component and native `select`s,
have a minimum height of **44 px** and the same horizontal and vertical inner spacing as the
text fields: 12 px at the sides and 8 px at the top and bottom. The rule applies in forms,
filters, tables, dialogs, settings, wizard and sidebar.

There is no compact variant: PC, tablet and mobile use the same height. Width,
typography, colours, border, focus, icon, options and current behaviour are kept; this change does not alter
the content or the number of elements each list shows.

### 31.3 Tests and scope

The component tests check the height and spacing of the shared selector. The frontend review
confirmed that no native selectors remain below 44 px and the browser verified a
sample of forms, filters and sidebar on PC and mobile, with their light and dark themes.

Texts, data, permissions, navigation, API, data model and production do not change. The owner
chose 1A, 2A, 3A, 4A and 5A. No open decisions remain.

The implementation passed the 232 frontend tests and the production build. In a real browser,
the eight visible dropdowns of Settings measured 44 px and showed 8 px of vertical
spacing and 12 px horizontal, both at 1440×900 and at 390×844. The check included theme,
language, sign-in method, dispatch, notices and AI, and was repeated with the light and dark themes.

## 32. Correction: 52 px dropdowns

### 32.1 Later finding

The implementation of section 31 is present: the controls measured 44 px after
reloading and no longer use the previous 36 px. However, the owner's visual review in the
ticket detail showed that **Assign to**, the state and **Theme** are still perceived as too
thin compared with the other controls. The reduced screenshot accentuated that impression, but
the measurement applied does not resolve the visual presence sought either.

### 32.2 Correction applied

All dropdowns move to a minimum height of **52 px**. The measurement applies to the shared component
and to any native `select`, without exceptions for forms, dialogs, filters, tables,
wizard, Settings or sidebar. PC, tablet and mobile use the same height.

The `text-sm` typography is kept, along with the colours, borders, focus, icon, width, options and
current behaviour. Spacing can grow only as needed to centre the content within
the 52 px; the text is not enlarged nor is a compact variant introduced.

### 32.3 Tests and scope

The component test requires 52 px and checks that the typography remains `text-sm`. The
code review covers the native selectors. The browser measures the controls actually served,
without depending on a resized screenshot.

Texts, data, permissions, navigation, API, data model and production do not change. The owner
chose 1A, 2A and 3A. No open decisions remain.

The implementation passed the 232 frontend tests and the production build. A new tab
against the development server measured the theme selector at **52 px**, with a 14 px font
and 8 px vertical spacing. The server had already rebuilt automatically; it was not necessary to
restart the containers.

## 33. Final adjustment of dropdowns to 48 px

### 33.1 Visual finding

The 52 px of section 32 loaded and measured correctly in a different browser. The
first tab kept the previous style in cache, so it was not a failure of the container or
of the rebuild. Once the real change was seen, the owner determined that 52 px take up
too much space.

### 33.2 Adjustment applied

All dropdowns have a minimum height of **48 px**, as a midpoint between the 44 px that
were perceived as thin and the oversized 52 px. The measurement is still global: shared component
and native `select`s, forms, dialogs, filters, tables, headers, wizard,
Settings and sidebar, on PC, tablet and mobile.

`text-sm`, the 12 px horizontal spacing and the 8 px vertical one are kept, in addition to colours,
borders, focus, icon, width, options and behaviour. No compact variant is created.

### 33.3 Tests and scope

The component test requires 48 px and keeps the typography and padding assertions. The
review also covers the native selector, and a clean browser measures the served result to
avoid confusing cache with the current version.

Texts, data, permissions, navigation, API, data model and production do not change. The owner
chose 1A, 2A and 3A. No open decisions remain.

The implementation passed the 232 frontend tests and the production build. A new tab
measured the served selector at **48 px**, with a 14 px font, 12 px horizontal spacing and
8 px vertical. The frontend rebuilt automatically and it was not necessary to restart containers.

## 34. Dropdowns equal to the segmented switch

### 34.1 Reference measurement

The Main/Both/Internal switch combines 36 px buttons, 2 px outer padding per side and a
1 px border per side: its outer height is **42 px**. The 48 px dropdowns approved in
section 33 are 6 px above and look misaligned when both controls share a header.

### 34.2 Normalisation applied

All dropdowns use a **fixed height of 42 px**, equal to the outer height of the switch.
The rule is global and reaches the shared component and the native `select`s in forms,
dialogs, filters, tables, headers, wizard, Settings and sidebar, on PC, tablet and
mobile. The switch does not change.

`text-sm`, the 12 px horizontal spacing and the 8 px vertical one are kept, in addition to colours,
borders, focus, icon, width, options and behaviour. There is no compact variant or automatic
growth by context.

### 34.3 Tests and scope

The component test requires a fixed height of 42 px and keeps the typography and
padding assertions. The review also covered the only native selector. A clean browser measured the
shared dropdown at **42 px**, with a 14 px font, 12 px horizontal padding and 8 px vertical;
it matches the 42 px outer height derived from the Main/Both/Internal switch.

The implementation passed the 232 frontend tests and the production build.

Texts, data, permissions, navigation, API, data model, switches and production do not change. The
owner chose 1A, 2A and 3A. No open decisions remain.

## 35. Implementation: automatic Summary without buttons

The PC lists, the mobile cards and the details panel replace **Reason/Motivo** with
**Summary/Resumen**. Last action/Última acción keeps its name. The details panel card keeps
both texts, but removes **Try Reason again**, **Try Last action again** and their Spanish
equivalents. No substitute action appears: both fields are automatic results of the ticket.

Only the details panel communicates `Generating…/Generando…` or a brief error next to the affected field. It keeps
the layout and does not use toast for a background process. In lists and mobile cards, a non-current result
is represented only with `—`; no chips, icons or repeated messages are added. The
screen reader receives the new name and the state when it exists.

The tests go through both languages, PC and mobile, ready, pending and error state, absence of all
the buttons and consistency of the label in the details panel, table and card. Permissions, navigation,
theme, layout and controls unrelated to insights do not change.

## 36. Ordered and legible ticket summary table

### 36.1 Reproduced finding

In the desktop inbox, the **Open/Abrir** link is technically inside the table and occupies
the eleventh cell of each row. However, the header of that column exists only for screen
readers through `sr-only`. The screenshot and the measurement at 1440×900 confirmed that the table
ends exactly after that cell, but the absence of a visible label makes the button
seem placed outside the summary table.

### 36.2 Agreed columns and order

The wide table keeps the eleven columns, but orders them by the operational reading sequence:

1. **Ticket**;
2. **Priority/Prioridad**;
3. **State/Estado**;
4. **Subject/Asunto**;
5. **Classification/Clasificación**;
6. **Summary/Resumen**;
7. **Last action/Última acción**;
8. **Requested by/Solicitante**;
9. **Assignee/Responsable**;
10. **Update date/Fecha actualización**;
11. **Action/Acción**.

Summary receives the largest flexible space; then Subject and Last action. Subject, Summary and Last
action occupy at most two lines with ellipsis and expose the full text through the
native help on hovering. Rows grow only to what those two lines require and
all content is aligned at the top.

**Ticket** always remains on a single line, without truncation and with enough width for the number.
Priority and State adjust their width to the longest chip and centre both header and content.
Requested by and Assignee allow up to two lines and show `—` when there is no value. Update date
uses one line for the date and another for the time within a compact column.

Classification keeps category and tags together. It shows at most two lines of chips; if
there are more tags than fit, a focusable `+N` indicator reports how many remain hidden. On
hovering or focusing it, it enumerates their names; its accessible name also enumerates them. It does not let
the row grow without limit.

The help for Subject, Summary and Last action with the full text opens both on hovering
and on focusing the cell with the keyboard; truncation does not turn that information into something exclusive
to the mouse.

### 36.3 Visible and compact Action

The last column shows the visible header **Action** in English and **Acción** in Spanish. The
header is left-aligned, like the other titles; the buttons remain centred within
their cells. The column uses only the minimum width necessary
for the button and its spacing; it does not receive a fixed width or steal additional space from subject, Summary or
Last action.

The button still says **Open/Abrir** and keeps its link, focus, size and permissions. The ticket
number remains an independent link. The entire row does not become clickable, so as not to
open a ticket when selecting text or interacting with other content.

The header remains on a single line and does not offer sorting: it describes an action, not a sortable
datum.

### 36.4 Intermediate width and mobile

At intermediate width, Requested by and Assignee are hidden first. These remain visible:

`Ticket · Priority · State · Subject · Classification · Summary · Last action · Update date · Action`.

When those nine columns stop fitting legibly, the table allows native horizontal
scrolling and keeps its measurements; it does not compress the controls. Only below 768 px does it switch to cards.
The real available width still decides whether the eleven columns or the intermediate variant appear,
so it also responds to opening or collapsing the sidebar and not only to the window.

The mobile card adopts the same conceptual order: Ticket, Priority, State, Subject,
Classification, Summary, Last action, Requested by, Assignee, Update date and Open. The button
**Open/Abrir** remains at the end. Missing people show `—` on the card too.

### 36.5 Tests and scope

The tests check both languages; order and number of columns; Ticket without wrapping; priority and
state centred, Action header on the left and its button centred; two-line texts with complete accessible content; limited chips and
`+N` counter; dashes for missing people; date and time on two lines; Action header
visible, not sortable and compact; the nine intermediate columns; timely switch to cards; and the
mobile order. The browser will go through wide desktop, intermediate width and mobile in English and Spanish.
Priority and Update date keep only the sorting already offered by the inbox; changing their
position does not add interactive headers or duplicate the existing sorting controls.

Data, API, permissions, navigation, summary content, filters, ticket order,
themes and production do not change. The visual order of the columns does change, as does the responsive point at which
the table gives way to cards, in accordance with this proposal.

## 37. Ticket fixed during horizontal scrolling

In any presentation that still uses a table —wide or intermediate—, the first column,
**Ticket**, remains visible when scrolling the content horizontally. Thus, each datum keeps the
unambiguous reference to the ticket in its row. Priority and Action are not fixed and the entire row remains
not a link.

The body cells are fixed exclusively horizontally. From section 39, the Ticket
header is the fixed corner of both axes and also remains at the top edge when going down a
long list. The column keeps exactly the width agreed in section 36; it does not gain space or
force the number to be truncated, which continues on a single line.

Both the `th` and the `td` of Ticket use the fully opaque background of the theme surface,
a sufficient stacking level and a right border. The scrolled content never becomes transparent
nor is drawn over the number. The border marks where the fixed area ends without introducing a new
shadow.

From 768 px the table remains available even if its content is wider than the container and
offers native horizontal scrolling. No gradient, text or icon is added to announce it: the
partially visible content itself indicates that it continues. Below 768 px cards are used,
which need neither fixing nor lateral scrolling.

The cards do not change: they have no horizontal scrolling or a column to fix. Nor do
data, routes, permissions, order, filters, pagination or the behaviour of the number link change.

The tests cover wide and intermediate tables, real horizontal scrolling, number always visible,
header fixed vertically, opaque background, right separator, focus and activation of the link. The
visual check goes through the eight themes to confirm that the background and border use their variables
and keep contrast.

## 38. Progressive loading of the inboxes

The four inboxes stop showing the total, `Page/Página`, `Previous/Anterior` and `Next/Siguiente`.
The first response brings 20 tickets. When the person approaches the end of the table or the
cards, an invisible marker automatically requests the next block before they reach
the last ticket.

While that block arrives, the already visible tickets remain available and a discreet indicator
at the foot reports that it is loading. Each received block is added in the order delivered by
the backend, without replacing or duplicating the previous ones. An accessible live region announces how many new
tickets were incorporated. When there is no other page left, the list ends without an additional message.

If progressive loading fails, everything already loaded is kept. At the foot the error appears with an action
**Retry/Reintentar**; there is no automatic retry or blocking of the list. A single block can be
in progress: several crossings of the marker do not produce parallel requests.

Changing search, filters, view or order cancels the previous continuity, empties its results,
returns to the start and requests the first block of the new criterion. The same behaviour is used in
table and cards, on PC, tablet and mobile.

When opening a ticket and returning with the browser history during the same browsing session,
the inbox recovers its criteria, the accumulated blocks and the vertical position. That memory is
ephemeral: reloading or closing the tab starts again from the first block. Going to another inbox does not
mix its results with the current one.

The English label of the `en espera` state is presented according to the role in accordance with section 40. In
the table, **Action/Acción** is left-aligned like the other headers; the Open/Abrir button
remains centred in its cell.

The tests cover appending without duplicates, exclusion of concurrent requests, silent end,
failure and retry, restart when changing criteria, accessible announcement, table and cards, restoration
on return, discard after reload, both languages, the text `Waiting` and the two different alignments
of header and button. The browser check goes through PC, tablet and mobile.

## 39. Table of available height and fixed header

From 768 px, the inbox header —title, search, filter summary and **New ticket/Nuevo
ticket**— remains outside scrolling. The table takes all the remaining height of the view, with
a minimum of 240 px, and its own container manages vertical and horizontal scrolling.
After the last row it keeps 16 px of inner space so that the content does not end stuck
to the edge. Below 768 px the cards keep the normal page scrolling.

The title row remains fixed at the top edge of the container while the
tickets are scrolled. It keeps its current background, typography and bottom border, with an opaque surface and the necessary
stacking level; it does not add a shadow. The **Ticket** header cell is fixed on both axes and
is drawn above both the other headers and the scrolled cells.

The right separator of the Ticket column stops depending on the border of the fixed cell: a 1 px line,
independent and superimposed on its right end, keeps the theme colour and remains
visible throughout horizontal scrolling. Its thickness is not increased nor is a shadow introduced.

Subject, Summary, Last action and the `+N` counter keep a single help of their own, accessible with
cursor and keyboard. The `title` attribute that duplicated that help is removed. The tooltip is shown only
when the content is really truncated; when focused with the keyboard it remains visible while it
keeps focus and closes when moving it. The full text remains available without depending on the mouse.

Progressive loading observes the end of the internal scrollable area and requests the next block on
approaching, with the same anticipation, exclusion of requests and error recovery of
section 38. The horizontal bar belongs to the container and remains at its bottom edge, without requiring
reaching the last row to scroll the columns.

The tests cover calculation of the remaining height and a minimum of 240 px, 16 px at the bottom, table versus
cards, combined vertical and horizontal scroll, fixed headers, Ticket corner on both axes,
persistent separator, absence of `title`, tooltip only with overflow, focus and closing, and progressive
loading inside the new container. The real check goes through PC, tablet and mobile, and confirms
the eight themes for the background, the borders and the separator.

## 40. `en espera` state according to who consults it

The technical value remains `en espera`. Its visible label depends on the role of the person
viewing the ticket:

| Role | English | Spanish |
| --- | --- | --- |
| User | `Waiting for you` | `Esperándote` |
| Support | `Waiting` | `Esperando` |
| Development | `Waiting` | `Esperando` |
| Administrator | `Waiting` | `Esperando` |

The distinction is applied coherently throughout the interface: inbox table, mobile cards,
ticket header and details panel, history and filters shown to each role. The User only sees the
main ticket; no variant is introduced for the internal one.

The selectors with which Support or Development change the state keep the operational label
`Waiting` / `Esperando`. They do not try to name the expected person there: the context and the text
requested when executing the transition already explain whether the User or Support is awaited.

Out of scope are the technical value `en espera`, the API contracts, the database, the
transition rules, the permissions and the email texts. No personal preference is added nor
is a new translation persisted.

The tests must cover the four roles in both languages and check the label in inbox,
card, details panel, history and filters. They must also confirm that the technical selector uses
`Waiting` / `Esperando` and that the backend continues to receive `en espera`.

### 40.1 Review record

The owner decided on 2026-10-16:

1. apply the difference throughout the interface;
2. treat the Administrator as a technical role;
3. keep the operational label in the state selectors;
4. limit the change to the interface, without modifying API, database or emails.

The review is closed with no open decisions. The owner explicitly approved the proposal
on 2026-10-16. The implementation respects the approved scope; 76 focused tests passed, as did the
frontend boundaries and the build.

## 41. Narrower sidebar and vertical brand

The sidebar keeps its three zones and its same permissions, routes and controls, but makes better
use of the view's width. On desktop and in the mobile drawer it measures **208 px open**. On desktop
collapsed it measures **56 px**; the mobile drawer does not offer the collapsed state. The links keep their touch
area, their icons, the active state and visible focus.

The open header stops being a row. The logo takes a frame of **80 × 80 px**, centred, and the
installation's full name appears below, also centred, in as many lines as it needs.
The name does not widen the menu or get truncated. The collapse or close button takes the top right
corner in a reserved space: it does not push or off-centre the logo.

In the collapsed state, the only button to expand is placed at the top. The **40 × 40 px** logo is
centred below, and a single separator closes the entire brand block. The second expand button that today
is repeated in the lower zone is removed. The brand does not share a row with the control and therefore
does not shift to the left. The name is visually hidden in this state.

Collapsed links keep the name through a tooltip on hovering and through accessible text
for assistive technologies. The reduction does not eliminate or replace any entry.

The joint screen `/tickets/categories` and its entry are renamed **Classification** in
English and **Clasificación** in Spanish. Inside it keeps two clearly titled cards
**Categories / Categorías** and **Tags / Etiquetas**. It is not split into two routes or two modules, and its
data, API, permissions and operations do not change.

### 41.1 Finding and scope

The previous contract fixed 240 px open and 64 px collapsed, and the current implementation places logo,
name and control in a row. In the collapsed state, 64 px are not enough for logo and button without
shifting the brand. In addition, `Categories and tags` does not fit comfortably in the proposed width although
the two subjects already form a single classification in lists and details panels.

Out of scope are the other menu labels, the order of entries, navigation, roles,
persistence of the collapsed state, drawer behaviour, logo formatting and
the administration of categories or tags.

### 41.2 Tests

The tests must cover 208 px open and 56 px collapsed, logo of 80 and 40 px, institutional name
of several lines, a single expand button, absence of lateral shift of the brand,
tooltip and accessible name in collapsed links, and the new name in both languages. The
browser check will go through desktop and mobile, menu open and collapsed, short and long name,
the eight themes and the roles that see Classification.

### 41.3 Review record

The owner decided on 2026-10-17:

1. keep a single screen and call it Classification / Clasificación;
2. use 208 px open and 56 px collapsed, also 208 px in the mobile drawer;
3. show the logo at 80 × 80 px open and 40 × 40 px collapsed;
4. show the full name, centred and without a line limit;
5. reserve the top right corner to collapse or close without moving the brand;
6. put the only expand control above the logo when the menu is collapsed;
7. keep a single separator below the entire brand block;
8. remove the duplicated lower control and keep tooltip and accessible text in the entries;
9. also rename the screen title and keep Categories and Tags inside.

The review is closed with no open decisions. The owner explicitly approved the proposal
on 2026-10-17. The implementation respects the approved scope; the 238 frontend tests passed,
as did the boundaries and the build. The real check covered desktop open and collapsed and mobile
drawer, and confirmed the centred brand, the single control and Classification with Categories and Tags.

## 42. Available height for the operational tables

From **768 px**, the ticket inboxes and the user list take the available height of the
window. The shell leaves **16 px at the top and 16 px at the bottom** on these screens. The header and the table are
distributed in a column limited in height by the view; the table grows with `flex` to reach the
bottom margin even when it has few results.

Each list's header keeps its hierarchy and content, but reduces wasted space:
the separation between the header and the table is **12 px**. In tickets, title, Search/Buscar,
filter summary, clear and New ticket/Nuevo ticket share a single row of **42 px**, centred
vertically when the width allows. The buttons in that row use 42 px, like the normalised
dropdowns; the title keeps its size and weight.

In Users the title, its description and the complete filter panel are kept. The description
continues below the title with a smaller separation, and only the panel's outer margins are reduced:
fields, switches and internal targets keep their current measurements. New account/Nueva
cuenta uses the same 42 px as the header. The table now manages the remaining height like the ticket
one.

The two containers have a minimum of **240 px** to remain usable in short windows. After the last row
they keep **16 px of inner space**, even with few results. The outer border ends 16 px from the bottom
edge of the window; it does not adjust to the last row or change height with the number of results.

Below 768 px the current card and page flow is kept. Touch controls are not compacted nor is a
fixed-height scrollable region forced.

### 42.1 Finding and scope

The current shell applies 32 px at the top and 40 px at the bottom on desktop. The ticket inbox
adds 24 px between its header and the table and calculates a view height that subtracts 72 px. Users does not
calculate the remaining height: its filters and table grow with the document. This produces uneven
margins and leaves vertical area unused.

The change is limited to the ticket inboxes and the user list. Out of scope are the catalogues,
the mail editor, the records, forms, settings, embedded tables, rows, columns,
order, filters, progressive loading, pagination and API contracts.

### 42.2 Tests

The tests must check from 768 px: 16 px outer top and bottom, 12 px between header and
table, 42 px header controls, a table with a 240 px minimum and 16 px inside after the last
row. They must also cover tickets with and without a creation action, Users with description and filters,
few and many results, short and tall windows, and absence of change below 768 px. The
real check will go through desktop, tablet and mobile in both languages and with the menu open and collapsed.

### 42.3 Review record

The owner decided on 2026-10-17:

1. apply the pattern to the ticket inboxes and the user table;
2. leave 16 px at the top and bottom and 12 px between header and table;
3. use 42 px in the header buttons, without reducing the title;
4. keep a single compact row in tickets when the width allows;
5. keep the container to the bottom edge even with few records;
6. keep the description and filter panel in Users, reducing only outer margins;
7. keep a minimum table of 240 px and 16 px inside after the last row;
8. apply the available height from 768 px and keep the current mobile flow.

The review is closed with no open decisions. The owner explicitly approved the proposal
on 2026-10-17. The implementation respects the agreed scope: it applies the available height only to the
ticket inboxes and Users from 768 px, compacts their headers and keeps the mobile flow. The
238 frontend tests, the boundaries and the build passed; the real check confirmed the
two tables on desktop and the ticket cards on mobile.

## 43. Interface for the cumulative Administrator

A combined account sees a single menu with the union of its operational role's entries and those of
Administrator, without duplicates or personality selector. On signing in it lands on the inbox of its
operational role; the factory account keeps the administrative start.

The **Admin / Administrator** badge appears with the role in the user list, record, own profile
and account block at the foot of the menu. In the user list, both the desktop table and the
cards of narrow screens present the operational role in the first row and the badge in a
second row, left-aligned and separated by a small, constant space; they never compete
for the width of the same line. It is not repeated in tickets or comments and does not depend only on colour.

Creation and record present an operational role selector and an independent Administrator checkbox. The
checkbox is only enabled for Support or Development and only another Administrator modifies it. Switching
between both roles does not uncheck it.

When removing or deactivating the last ordinary Administrator, a window consistent with the theme reports that the
factory account will be the only administrative route and requires confirmation, without preventing the action.

The next identity load shows the affected person a bilingual notice for the grant or
removal. It can be closed and, as a deliberate exception to the ten-second toasts, does not disappear
automatically. The record shows the history with actor, action and date; a global screen is not born.

Selector, checkbox, badge, confirmation, notice and history are verified with keyboard, screen
reader, PC and mobile. The 240 frontend tests and their boundaries passed on 2026-10-18.

## 44. Dashboard and reporting

Dashboard/Panel is the first menu entry for Support, Development and Administrator and its initial
destination. User does not see it. At the top it presents six indicators compared with the previous period;
in the centre, a line of entries/closures and bars of state, priority, category and assignee;
it does not present tables, cards or links to individual tickets below.

The period, type, state and priority filters live in the URL and adapt to the width without hiding
their value. Refreshing keeps the last
visible result. Each chart has a title, legend, accessible values and a textual alternative;
colour is never the only difference. The bars use `bg-primario`, `text-primario`,
`text-primario-texto` and `bg-exito`, never the non-existent English aliases. On mobile the blocks
stack. The series is daily up to 31 days and weekly from 32; it only labels dates that fit and keeps
the exact date of each point accessibly.

The only visible export is presented as a button with a download icon to the left of
“XLSX” and generates `report_AAAA-MM-DD_AAAA-MM-DD.xlsx`, localised and divided into six aggregated sheets,
without embedded charts or individual data. The verification will go through
desktop and mobile, the eight themes and both languages, and will check dimensions and computed colours,
temporal labels, absence of tickets and workbook content.

## 45. Single active state of the sidebar

The sidebar shows **a single active entry**. The routes of its lists match exactly:
`/tickets`, `/tickets/main`, `/tickets/internal` and `/tickets/categories` are not activated by
prefix nor can they light up simultaneously. The active entry keeps the background
`bg-superficie-suave`, the main text and the current medium weight; it does not add shadow, border or a
second decorative signal. The link also exposes `aria-current="page"`.

Screens that do not have their own entry keep the agreed functional belonging:

- `/tickets/new` and `/tickets/:numero` activate **My tickets / Inbox**;
- `/users/:id` activates **Users**;
- `/mail` activates **Settings**;
- the other routes without their own entry do not select any module.

The rule is identical with the menu open, collapsed or in the mobile drawer and does not depend on the role or on whether the
account accumulates Administration. A pure function resolves the URL's belonging and its tests
go through the visible and auxiliary routes. The E2E test checks that there is exactly one link
with `aria-current="page"` and that it keeps the soft background. The manual journey with a combined
account confirmed all the entries and the four classes of auxiliary route.

### 45.1 Decision record

The owner chose on 2026-10-19:

1. exact match and a single active entry (1A);
2. current soft background, without shadow (2A);
3. creation and detail inside **My tickets / Inbox** (3A);
4. user record inside **Users** and mail editor inside **Settings** (4A).

Out of scope are order, labels, icons, permissions, routes, initial
destinations, the theme palette and any API or data change. No open decisions remain;
the owner explicitly approved the proposal and the implementation respects the agreed scope.

## 46. Header and views of “My tickets”

The first row of **My tickets** contains, next to the title and the existing actions, an always-visible
selector with three options: **All · Assigned · Watching** / **Todos · Asignados · Observo**.
The initial view is **All**, which combines assigned and watched tickets without duplicates. Changing the
option updates the list and its address; it does not open a modal or count as a search filter.

Support sees only main tickets and Development only internal ones. An operational account with
Admin capability keeps exactly that presentation. When ceasing to watch an unassigned ticket,
it disappears when the inbox is updated; if it was open in the detail, the details panel does not close or get
interrupted and on returning the current list is reflected.

No view of unassigned or historical participation is offered. **Main tickets** and
**Internal tickets** keep their filters and their general content. The tests cover the three
selectors, the combination without duplicates, both operational roles, a combined account and the
removal of a watcher.

## 47. Public and private comments

In the main ticket's comment editor, Support sees a **Public · Private** / **Público ·
Privado** selector, with `Public` selected on opening. Development sees the same control limited to `Private`,
since they cannot respond publicly to the User. The selector does not appear in the internal ticket or for the User.
The editors that accompany state movements do not show it and always publish publicly.

The team distinguishes a private note through the **Private / Privado** chip next to the author and a
subtle background derived from the theme; it does not depend only on colour. The User sees no gap, padlock, counter,
history entry or sign that it exists. Private attachments are rendered only inside
that note for the authorised audience.

The author can edit the comment and toggle its visibility in both directions from the existing edit
action. The screen immediately updates conversation, attachments, search and
summaries; it does not open a confirmation when turning a public comment private. The technical history shows
who changed **Public → Private** or **Private → Public** and when, without repeating the body.

When escalating, the note is not duplicated in the internal thread. In double view it remains in the main
ticket's column; if only the internal one is viewed, it is consulted by changing view. The tests go through the eight
themes, both languages, PC and mobile, keyboard and screen reader, in addition to checking that the User's DOM
does not contain the comment or its attachments.

## 48. Search for and create an Active Directory account

The creation dialog of `/users` keeps **Application accounts** for local creation. When AD
is the current method it adds **Organisation directory**. Keycloak appears disabled with
the explanation “The account is created on its first access”; if the current method is not AD, its search
is not offered.

The directory view starts with a single field **User or email**. Before three characters it does not
query and shows a brief help. Afterwards it waits for the typing pause, cancels the previous
request and presents up to twenty results with full name, user and email. It has different states
for searching, no matches and directory unavailable; keyboard and screen reader
can go through and select each result.

After selecting, the role is chosen with the current limits: Support only User and
Administrator the roles/capability they already administer. Name, surnames and email are not edited,
because they come from AD. Create returns to the list, keeps filters and shows the result in the common
toast. It does not send a creation email.

In the Spanish interface, the visible account labels say **Rol** —in creation, the filters,
the columns and the profile— as it is the most direct term for whoever administers users. The internal
model and permissions do not change.

If the email already corresponds to a local account, the themed dialog explains before confirming that
that account will be kept, will move to AD and its local password will stop working. Cancelling changes
nothing. The confirmation is not reused for another person nor does it survive closing the dialog.

The automated tests verify the options according to the method, that Keycloak does not appear
queryable and the search and creation contract. The complete frontend suite (249 cases), its
boundaries and the production build passed on 2026-10-20.

## 49. Unified navigation and administrative modules

The Support and Development menu replaces **Main tickets** and **Internal tickets** with a
single **Tickets** entry. **My tickets** remains independent. On entering, the screen opens
**Main** and presents a segmented selector **Main / Internal**; on mobile that control
becomes a dropdown so as not to compress the two labels.

The selector writes `type=main|internal` in the URL and responds to browser history. Each side
keeps its own filters while the screen remains mounted. It does not show counters. The old
routes disappear without redirection, and the menu always links to `/tickets?type=main`.

In the main view, Development sees an editor labelled **Internal comment**. Everything published
there is private automatically: no visibility selector or blocked control appears. Support
and Development can read it; the User cannot infer it. Support keeps the current public/private selector in its own
comments.

**Classification** and **Users** only appear when the identity has the Administrator capability.
A combined account orders its menu as Panel, My tickets, Tickets, Classification, Users and
Settings. The pure Administrator keeps their administrative navigation and does not receive Tickets.

The visible terminology of the area is reviewed completely in Spanish and English: module names,
headers, filters, buttons, aids and states. The concepts **My tickets**,
**main ticket** and **internal ticket** are kept; the review removes mixtures and old names, it does not change the
model. The 249 frontend tests, its boundaries and the production build passed on
2026-10-20; navigation covers pure and combined roles and absence of unauthorised entries.

## 50. Site icon and favicon

The brand block of `/setup` and Settings adds a square **Site icon** card, apart
from the light and dark logos. The preview fits the whole image centred, without cropping or
deforming it. Pressing it opens the file selector; the help accepts PNG, JPEG, WebP and SVG, up to 1 MB and
2000 pixels per side, with a minimum of 32 × 32. The **Restore** action only appears when there is a
custom icon and explains that the light logo will be returned to.

The tab favicon updates when an upload or restoration finishes, without reloading. If there is no
custom icon, it also updates when the light logo changes or is restored. The same update
is propagated to the other open Catalina Support tabs; a tab opened later obtains
the effective value when reading the public brand. The link carries the version stamp so as not to reuse
a previous image from the cache.

The icon is decorative because the tab name already identifies the application. Uploading keeps
focus, shows the result in the common toast and does not block other parts of Settings. It is tested
in the eight themes, both languages, keyboard, PC and mobile, including the three fallback levels and the
synchronisation between tabs. The icon of installed shortcuts is not changed.

## 51. Useful width of the ticket inboxes

The four ticket inboxes take the available width inside the shell and leave 16 px on each
side on desktop. This exception replaces for those screens the general 40 px and its limit of
88 rem; it does not widen forms or reading pages. The sidebar keeps its widths, collapse
control and current state.

The table distributes the freed space among its descriptive columns. Ticket remains fixed to the
left and Action to the right when there is scrolling. The first column adds air next to the
left edge and the last one next to the right so that links and buttons do not look stuck to the
frame. Separators, backgrounds, focus and contrast use the tokens of each theme. On mobile the
cards remain.

The browser is checked with the menu open and collapsed, 1440 and 1920 px, an intermediate width with
scroll, user zoom and mobile. The visual review requires that the final button be visible or remain
fixed during scrolling, that it does not cover text and that Summary and Last action keep their
legibility.

### 51.1 Decision record

The owner clarified on 2026-10-22 that the wasted space is to the left and right of the
table, not in the menu. They chose the four inboxes (1A), 16 px (2A), fixed Action (3A) and added inner
padding at both ends. No open decisions remain.

### 51.2 Implementation and verification

The implementation leaves 16 effective px through compensation localised in the inbox, without changing
the shell or other pages. Action received 6 rem after the first visual review detected
that the previous minimum width clipped the header and button. The second Chromium pass confirmed
the complete table without overflow at 1920 and 1440 px, and the cards at 390 px. The 252 Angular
tests passed, as did the focused regression, the boundaries and the build into a clean directory.

## 52. Consistent states in settings, inboxes and creation

> **Coordinated proposal of 2026-10-23.** The owner chose 1A, 2A and 4A and closed the review with
> 1A–5A. No open decisions remain; explicit approval is pending.

A failure of the local AI catalogue is presented inside its card and keeps the alternatives of step 5
operable. It does not open the global server-down notice while the backend responds.
The card allows trying again. Settings offers the same treatment.

When changing Main or Internal, the previous table disappears and loading takes its place. The button,
the URL and the rows are shown together again when the current query finishes.

In `/users`, each width exposes a single **New account** button to focus and to the accessibility tree.
The responsive variants do not produce duplicate controls for screen readers or tests.
