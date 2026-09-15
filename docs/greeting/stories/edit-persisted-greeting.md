# Story — Edit persisted greeting

Module: `greeting`
Plan item: Edit persisted greeting
Requirement: `GREETING-002 — Save updated greeting`

## User story

As a Visitor, I want to edit the greeting and save it, so that the new greeting remains after reload.

## In scope

- Greeting page default state only.
- Existing single greeting shown in heading and text field.
- Visitor can change text field value and submit with `Save` button or native form submit.
- Non-empty submitted value is trimmed at leading and trailing whitespace before save.
- Successful save updates both heading and text field.
- Saved value persists in PostgreSQL through Go API and remains after reload.
- Empty trimmed submit saves nothing, keeps existing heading, and returns focus to text field with no visible error message.
- Responsive form layout and focus styling from approved design.

## Out of scope

- Initial greeting display rules owned by story `Show persisted greeting`, except this story depends on that behaviour being present before edit.
- Sign-in, permissions, per-visitor greeting, ownership, and history; product has one public Visitor role and one shared greeting.
- Navigation, extra sections, loading state, empty state, and visible error state; approved design has one centered default section only.
- Animation or transition; approved design has no motion.
- External services; none are used.
- Conflict detection or optimistic locking; SRS uses last successful save wins.
- Recovery UI for missing seeded row or API/database failure; service contract may return errors, but no approved visible failure state exists.

## UI scope

Screen: Greeting page, section `<section class="greeting-section" aria-labelledby="greeting-heading">`.

This story touches the one approved section only:

- Large greeting heading updates after successful save and wraps long saved text without horizontal page scroll.
- Text input is labelled `Greeting`, starts with the current greeting, accepts visitor edits, and shows saved value after successful save.
- `Save` button submits the form.
- At `520px` viewport width or below, input and button stack, fill container width, and remain `48px` tall.
- Above `520px`, input and button sit in one row with `12px` gap and remain `48px` tall.
- Keyboard focus on input or button uses black `3px` outline with `3px` offset.
- No additional UI states or messages are introduced.

## Acceptance criteria

- SC-1 [GREETING-002 AC-1]: Given the page shows `Hello, World!`, when Visitor enters `Hello, Pipeline!` and selects `Save`, then the heading text becomes exactly `Hello, Pipeline!`.
- SC-2 [GREETING-002 AC-2]: Given Visitor has saved `Hello, Pipeline!`, when Visitor reloads the page, then the heading text is exactly `Hello, Pipeline!`.
- SC-3 [GREETING-002 AC-3]: Given the page shows `Hello, World!`, when Visitor enters `  Trim me  ` and selects `Save`, then the saved heading text is exactly `Trim me`.
- SC-4 [GREETING-002 AC-4]: Given the page shows `Hello, World!`, when Visitor submits an empty text field, then focus returns to the text field and the heading remains `Hello, World!`.
- SC-5 [GREETING-002 AC-5]: Given the page is displayed at `520px` viewport width or below, when Visitor views the form, then the text field and `Save` button are stacked, full width, and each `48px` tall.
- SC-6 [GREETING-002 AC-6]: Given the page is displayed above `520px` viewport width, when Visitor views the form, then the text field and `Save` button appear in one row with `12px` gap and each `48px` tall.
- SC-7 [GREETING-002 AC-7]: Given the input or `Save` button has keyboard focus, when Visitor tabs through controls, then the focused control has a black `3px` outline with `3px` offset.

## Dependencies

- Story `Show persisted greeting` must provide current stored greeting on page load and seed default `Hello, World!`.
- PostgreSQL persistence must store the single shared greeting across reloads.
- Go API must expose read and save behaviour for greeting text.
- Next.js frontend must render approved Greeting page and handle form input.
- Approved design system tokens must exist in frontend globals before component styling.
- No external accounts, credentials, or stakeholder data required.

## Notes for delivery

- Preserve internal whitespace, punctuation, and capitalization in saved value.
- Trim only leading and trailing whitespace before save.
- Reject empty trimmed value in UI before save attempt when possible; backend must still enforce same rule at API boundary.
- Last successful save wins for concurrent visitors.
