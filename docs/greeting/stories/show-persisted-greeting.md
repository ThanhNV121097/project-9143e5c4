# Story — Show persisted greeting

Module: `greeting`
Plan item: Show persisted greeting
Requirement: GREETING-001 — Display stored greeting

## User story

As a Visitor, I want to see the stored greeting when I open the page, so that the page reflects the current persisted value.

## In scope

- Read the single shared greeting from PostgreSQL through the Go API.
- Render the stored greeting in the Next.js page as the only large heading in the centered greeting section.
- Seeded first-use greeting displays as `Hello, World!`.
- Preserve stored greeting text exactly when displayed, including punctuation, capitalization, and internal spacing.
- Keep approved minimal page structure and style for the default greeting screen.
- Ensure long greeting text wraps within page width without horizontal page scroll.

## Out of scope

- Editing or saving greeting text; covered by `Edit persisted greeting` / GREETING-002.
- Per-visitor greetings, sign-in, roles, or permissions; product has one public Visitor role.
- Navigation, extra sections, animation, loading states, empty states, and visible error states; approved design excludes them.
- Live updates when another visitor changes the greeting; reload behaviour belongs to save story and API contract.
- Recovery UI for absent greeting row, API failure, or database failure; service contract handles API error envelope and approved screen has no failure state.

## UI scope

Screen: Greeting page, default state only.

This story owns visible display of `<section class="greeting-section" aria-labelledby="greeting-heading">` from approved design: one centered section on white background, one large black heading showing stored greeting, one text input labelled `Greeting`, and one blue `Save` button below it. Input and button are visible as part of approved default screen but save interaction is out of scope for this story.

## Acceptance criteria

- SC-1 [GREETING-001 AC-1]: Given stored greeting has never been changed, when a Visitor opens the page, the heading text is exactly `Hello, World!`.
- SC-2 [GREETING-001 AC-2]: Given stored greeting is `Good morning`, when a Visitor opens the page, the heading text is exactly `Good morning`.
- SC-3 [GREETING-001 AC-3]: Given stored greeting is `Hello,   World!`, when a Visitor opens the page, the heading text preserves the three spaces between words.
- SC-4 [GREETING-001 AC-4]: Given the page is displayed, when a Visitor inspects the default screen, the screen contains one centered greeting section with one large heading, one text input labelled `Greeting`, and one `Save` button.
- SC-5 [GREETING-001 AC-5]: Given the page is displayed, when a Visitor views the page, the page has white background, black text, no navigation, and no extra sections.
- SC-6 [GREETING-001 boundary]: Given the stored greeting is long, when a Visitor opens the page, the heading wraps within the page width and the page has no horizontal scroll.

## Dependencies

- PostgreSQL stores one shared greeting row seeded as `Hello, World!`.
- Go API exposes read access for the current greeting per architecture service contract.
- Next.js page can call the API through configured frontend environment.
- No external accounts or credentials required.
- No prior story must land first; this story may use mock data during UI stage and replace it with API data during backend stage.

## Notes for downstream stages

- Actor is `Visitor` from SRS.
- Fixed UI copy is English: `Greeting` and `Save`.
- Greeting text is user-controlled stored data and must display exactly as returned by API for this read story.
- Backend recovery for absent seeded row and upstream failures belongs to service design, not an added visible UI state.
