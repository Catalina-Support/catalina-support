# Dashboard and reporting

[Español](../es/modules/dashboard.md) · **English**

> **State:** as-built
> **Last updated:** 2026-10-19
>
> **Amendment approved, implemented and verified on 2026-10-19: direct reading of entries and
> closures.** The review of the
> chart confirmed that no closures are being lost: the ten closures in the local data
> occurred between August 3 and September 11, 2026 and therefore fall outside the
> queried period. The owner chose 1A–4A: keep the real date of the closing event,
> always show the value of each column —including zero—, use two series colours of their own and
> clearly distinct in the eight themes, and explain when the period contains no closure.
> The interface and the bilingual texts were implemented; two new tests cover periods with
> and without closures, the 243 frontend tests and the build passed, and the browser confirmed the
> permanent values and the colours computed in the eight themes. A second range, from August 1
> to September 15, confirmed real closures grouped by week and the correct absence of the warning.
>
> **Correction approved, implemented and verified on 2026-10-19.** The real review
> across the eight themes found that the data and sizes of the bars exist, but the main colour
> is transparent because this screen uses English tokens that the theme does not define. It also found
> a table and ticket cards that contradict the new decision that the Dashboard shows only
> statistics. The owner closed the review with 1A–4A, proposed XLSX in 5, chose 6A–8A, 9A–14A,
> corrected 15A to the name `report_AAAA-MM-DD_AAAA-MM-DD.xlsx` and chose 16A. No open decisions
> remain. The individual detail and its endpoint were removed, the tokens were corrected, the filters
> were reduced and the export is now the agreed six-sheet aggregated XLSX.
>
> **Implemented and verified on 2026-10-18, with explicit approval from the owner.** This
> document describes the first Dashboard and reporting module. The owner closed decisions
> 1A–45A, with the precision of 24A to export both duration formats and that of 26A to
> exclude the User entirely.

## 1. Purpose

The Dashboard offers an operational reading, not an evaluation of people or a contractual
SLA system. It makes it possible to detect load, urgency, age and trend only through aggregated
statistics. It does not list or open individual tickets and does not send notices or emails.

## 2. Access and scope

- User has no route, menu, API or access to the Dashboard.
- Support sees main and internal.
- Development sees internal only.
- A pure Administrator sees everything.
- Support or Development with the Admin capability accumulates the administrative scope and sees everything.

Dashboard/Panel appears before the inboxes and is the initial destination for Support, Development and
Administrator. User continues in My tickets.

## 3. Period, filters and refresh

It opens with the last thirty days in the installation's global time zone. It accepts only a date
range, type, status and priority. The filters live in the URL and govern indicators, charts and
XLSX. Data is queried on opening, when filters change and when Refresh is pressed; there is no polling or
real time. Category and assignee remain as breakdowns, not as filters.

## 4. Indicators

The header shows open, open critical, overdue, average first response time,
average resolution time and reopenings. Each one compares with the immediately preceding period
of equal length.

Open means any status other than closed. A missing value is presented with a dash;
a ticket awaiting a response says No response. Averages never turn absences into zero.

## 5. Time definitions

- First response on the main ticket: first human comment by Support.
- First response on the internal ticket: first human comment by Development.
- The average uses only tickets that have already received a response and shows the pending ones separately.
- Resolution: time from creation to the first resolution; the average includes the first
  resolutions that occurred in the period.
- Reopenings are counted separately.
- Entries use the creation date; closures use the date of the closing event.
- Age accumulates while the ticket is pending, stops on resolution and accumulates again
  from a reopening.

## 6. Configurable overdue

An open ticket is overdue when it exceeds, in running hours, the threshold of its priority. The initial
values are critical 4 h, high 24 h, medium 72 h and low 120 h. Only the Administrator modifies them in the
Reporting card in Settings; the teams only consult them. They are operational criteria, not a
formal SLA commitment or a compliance percentage.

## 7. Charts

Accessible bars and lines are used, never pie charts:

1. entries versus closures;
2. distribution by status;
3. distribution by priority;
4. top ten categories and Others;
5. open load by assignee, including No assignee.

The trend groups by day up to 31 days and by week from 32 days. The interface shows only the
date labels that fit without overlapping. In entries versus closures, each column always shows
its numeric value —including when it is zero— without requiring a pointer, focus or interaction; the pointer
tooltip and the accessible text also keep the exact date, series and value. The series
use two tokens of their own from the Dashboard: entries keeps the relation with the primary colour and
closed receives a clearly distinguishable secondary colour. Both tokens are defined and
checked in the eight themes; they do not rely only on the legend to be told apart. If all the closures
in the period are zero, the chart remains visible with its zeros and adds «No tickets closed in this
period» / «No se cerraron tickets en este periodo». The assignees chart measures load, not
productivity.

## 8. No ticket detail

The Dashboard contains no table, cards, progressive loading, numbers, subjects or links to tickets. The
operational detail continues exclusively in the inboxes and detail views of the tickets module.

## 9. XLSX workbook

The only export is an XLSX workbook of aggregated data with six sheets: Summary, Time series,
Status, Priority, Category and Assignee. Sheet names, headers and translatable values use
the global language. Summary includes the period, the four applied filters and the six indicators.
The other sheets contain exactly the visible aggregates; they add no hidden comparison or ticket
rows.

The sheets are formatted tables with a fixed header and filters, without embedded charts. Dates
are real cells with local formatting and add the ISO value when useful for integration. Durations
keep the human value and whole minutes. The file is named
`report_AAAA-MM-DD_AAAA-MM-DD.xlsx`, with the ends of the filtered period. The server neutralises
text that could be interpreted as a formula.

## 10. Interface states

Indicators and charts share loading, empty and error states without erasing the last valid result during
a manual refresh. On mobile they stack without adding a ticket list. The distributions
use the real tokens `primario` and `primario-texto`; the time series uses its specific
tokens for entries and closures. The values, the two colours against each other and their contrast with the background
are checked in the eight themes. Everything works with a keyboard, a screen reader and both languages.

## 11. Out of scope

It does not include individual productivity, goals, contractual SLA, alerts, emails, real time,
report scheduling, CSV, PDF, personal preferences, User access or any detail
of or access to individual tickets.

## 12. Required verification

The tests will cover calculations with a real PostgreSQL, permissions and cumulative scope, time zone,
previous periods, the four URL filters, daily/weekly grouping, thresholds, reopenings,
missing values, a localised and neutralised six-sheet XLSX, absence of tickets, responsive on
PC and mobile, visible bars in the eight themes, permanent values including zero, differentiation of
the two series and the bilingual warning when the period contains no closures.

The implementation keeps `/api/dashboard/summary` and adapts
`/api/dashboard/export` to the XLSX workbook; the UI does not consume detail and the contract
that only supported the individual list was removed. The `/dashboard` route, its menu entry and the
Reporting card in Settings remain. PostgreSQL reconstructs point-in-time states, active intervals
and temporal facts; the Dashboard service aggregates only authorised statistics.

The check walked through Support, Development, a rejected User and Administrator. The review
of 2026-10-18 added the findings of transparent colour in the eight themes, overlapping dates and
ticket detail outside the new scope. After implementing it, the Go services, the reporting
repository, the 241 frontend tests, the boundaries and the build passed. The browser confirmed the
absence of individual tickets, four filters, adaptive time labels and effective colour
in the eight themes; the downloaded XLSX contains six sheets and valid XML.
