# SRS — Greeting

Module: `greeting`
Design: [View the approved design](http://localhost:8080/design/9143e5c4-7d25-4d3d-a023-6407563c80b9)
Design system: `design/design-system.md`

> One file per module, at `docs/{module}/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

The greeting module lets any visitor view and update the single greeting for "Hello World Acceptance 8". It proves the full product path: persisted PostgreSQL value, Go API read/write, and Next.js page rendering. Without this module, the project cannot meet its acceptance goal of showing and persisting a database-backed greeting.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Visitor | Any person opening the public page; no sign-in exists | View the stored greeting, edit it in the text field, and save a new stored greeting |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Show persisted greeting
- Edit persisted greeting

**Out of scope** — name what a reader would reasonably expect here and say
where it lives instead. This section prevents the same argument twice.

- Sign-in and permissions — deliberately not built; project has one public visitor role.
- Navigation and additional sections — deliberately not built; approved structure has one centered greeting section only.
- External services — deliberately not built; project uses no external service.
- Animation — deliberately not built; approved design has no motion or transition.
- Visible loading, empty, or error states — deliberately not built; approved design shows only the default state.

## 4. Functional requirements

### 4.1 Show persisted greeting

**Requirement GREETING-001 — Display stored greeting**

*As a* Visitor, *I want to* see the stored greeting when I open the page, *so that* the page reflects the current persisted value.

Behaviour:

1. Visitor opens the greeting page.
2. The page obtains the current stored greeting through the product API.
3. The page displays that greeting as the only large heading in the centered greeting section.
4. On first project use, the stored greeting is `Hello, World!`.
5. The heading preserves the stored greeting text exactly, including punctuation, capitalization, and spacing.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/show-persisted-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-001 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Stored greeting has never been changed | Visitor opens the page | The heading text is exactly `Hello, World!` |
| AC-2 | Stored greeting is `Good morning` | Visitor opens the page | The heading text is exactly `Good morning` |
| AC-3 | Stored greeting is `Hello,   World!` | Visitor opens the page | The heading text preserves the three spaces between words |
| AC-4 | Page is displayed | Visitor inspects the default screen | The screen contains one centered greeting section with one large heading, one text input labelled `Greeting`, and one `Save` button |
| AC-5 | Page is displayed | Visitor views the page | The page has white background, black text, and no navigation or extra sections |

**Failure, boundary and permission behaviour** — the part most often skipped
and most often the source of bugs. Every case this function actually has needs a
defined outcome; "should not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Visitor is only reading the greeting | Not applicable: this function has no visitor input |
| Boundary | Stored greeting is long | Heading wraps within the page width and does not create horizontal page scroll |
| Not found | Stored greeting row is absent | Not applicable: seeded greeting exists by product requirement; recovery belongs to service design, not an approved screen state |
| Not permitted | Visitor opens the page | Not applicable: all visitors may view the greeting and no sign-in exists |
| Conflict | Stored greeting changes during display | Not applicable: approved design has no live update state; a later reload shows the latest stored value |
| Upstream failure | API or database cannot provide the greeting | No additional page state is specified: approved design has no error screen; API error envelope belongs to service contract |

**Data touched** — the fields this function reads and writes, in product terms.
The physical schema is TL's job in `docs/architecture/erd.md`; this is the list
that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Seeded as `Hello, World!`; displayed exactly as stored; no maximum length decided in product scope |

### 4.2 Edit persisted greeting

**Requirement GREETING-002 — Save updated greeting**

*As a* Visitor, *I want to* edit the greeting and save it, *so that* the new greeting remains after reload.

Behaviour:

1. Visitor sees the current greeting in the heading and the same current greeting in the text field.
2. Visitor edits the text field.
3. Visitor submits with the `Save` button or native form submit.
4. If the submitted value is non-empty after trimming leading and trailing whitespace, the product saves the trimmed value through the API.
5. After a successful save, the heading and text field both show the saved value.
6. After page reload, the saved value is still shown.

**Acceptance criteria** — each is proved by at least one test case in
`docs/greeting/test-cases/edit-persisted-greeting.md`, through the story plan that cites it
(`SC-1 [GREETING-002 AC-1]`). Given/When/Then, no compound
conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Page shows `Hello, World!` | Visitor enters `Hello, Pipeline!` and selects `Save` | The heading text becomes exactly `Hello, Pipeline!` |
| AC-2 | Visitor has saved `Hello, Pipeline!` | Visitor reloads the page | The heading text is exactly `Hello, Pipeline!` |
| AC-3 | Page shows `Hello, World!` | Visitor enters `  Trim me  ` and selects `Save` | The saved heading text is exactly `Trim me` |
| AC-4 | Page shows `Hello, World!` | Visitor submits an empty text field | Focus returns to the text field and the heading remains `Hello, World!` |
| AC-5 | Page is displayed at `520px` viewport width or below | Visitor views the form | The text field and `Save` button are stacked, full width, and each `48px` tall |
| AC-6 | Page is displayed above `520px` viewport width | Visitor views the form | The text field and `Save` button appear in one row with `12px` gap and each `48px` tall |
| AC-7 | Input or `Save` button has keyboard focus | Visitor tabs through controls | Focused control has a black `3px` outline with `3px` offset |

**Failure, boundary and permission behaviour** — the part most often skipped
and most often the source of bugs. Every case this function actually has needs a
defined outcome; "should not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Submitted greeting is empty after trimming | Nothing is saved, heading stays unchanged, text field receives focus, and no visible error message appears |
| Boundary | Submitted greeting is long | Saved greeting wraps in the heading and does not create horizontal page scroll |
| Not found | Stored greeting row is absent during save | Not applicable: seeded greeting exists by product requirement; recovery belongs to service design, not an approved screen state |
| Not permitted | Visitor submits a greeting | Not applicable: all visitors may save and no sign-in exists |
| Conflict | Two visitors save different greetings | Last successful save is the value shown after reload |
| Upstream failure | API or database cannot save the greeting | No additional page state is specified: approved design has no error screen; API error envelope belongs to service contract |

**Data touched** — the fields this function reads and writes, in product terms.
The physical schema is TL's job in `docs/architecture/erd.md`; this is the list
that document has to satisfy.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Trim leading and trailing whitespace before save; reject empty trimmed value; preserve internal whitespace, punctuation, and capitalization |

## 5. Screens

The design is the source of truth for appearance; this section maps functions
onto it so nothing in the design is unaccounted for and nothing specified here
is missing from the design.

List only the states the approved design actually shows. A screen the design
draws once, with no variant for waiting, for no data, or for a failure, has
exactly **one** state and its name is `default`. That is not a placeholder and
not an invented state: it is what "this screen has one appearance" is called,
and it is the correct and complete answer for a static screen. Writing
`loading`, `empty` or `error` for a screen whose design has no such variant
invents work, and the reviewer will reject it.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting page | `<section class="greeting-section" aria-labelledby="greeting-heading">` | GREETING-001, GREETING-002 | default |

## 6. Non-functional requirements

Only what is real for this module. Delete rows that do not apply rather than
inventing a number nobody will check.

| Area | Requirement |
|---|---|
| Accessibility | Text input has accessible label `Greeting`; input and button are keyboard reachable; focus indicator is visible; text/background contrast is at least 4.5:1 |
| Responsive | Page works at viewport widths from `320px` upward with no horizontal page scroll; at `520px` and below the input and button stack and remain `48px` tall |
| Localisation | User-facing fixed copy is English: `Greeting` and `Save`; greeting text is user-controlled and preserved as stored |
| Privacy | No personal data is required or stored by this module; only greeting text is stored |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL persistence, for storing the single greeting across reloads.
- **Depends on:** Go API, for reading and saving the greeting between frontend and persistence.
- **Depends on:** Next.js frontend, for rendering the approved page and handling visitor input.
- **Assumption:** One shared greeting exists for all visitors. If false, scope changes to per-visitor storage and needs sign-in or another identity rule.
- **Assumption:** Last successful save wins for concurrent visitors. If false, scope changes to conflict detection and a designed conflict state.

| Open question | Proposed default | Who decides |
|---|---|---|
| — | No open questions | — |

## 8. Traceability

Every plan item in this module appears exactly once, and every requirement id
traces to a test case. A gap in this table is a gap in the build.

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Show persisted greeting | GREETING-001 | `test-cases/show-persisted-greeting.md` |
| Edit persisted greeting | GREETING-002 | `test-cases/edit-persisted-greeting.md` |
