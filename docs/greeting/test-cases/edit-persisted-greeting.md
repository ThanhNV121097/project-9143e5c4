# Test cases — Edit persisted greeting

Module: `greeting`
Function: Edit persisted greeting
Story: `docs/greeting/stories/edit-persisted-greeting.md`
Requirement: `GREETING-002 — Save updated greeting`
Risk level: Medium. This function writes shared persisted state and has responsive and keyboard-accessibility requirements, but no sign-in, payments, personal data, or external services.

## Cases

**Scenario**: Save button updates heading to new greeting
**Given**: Stored greeting is `Hello, World!` and the page is open with the heading and `Greeting` text field both showing `Hello, World!`.
**When**: Visitor replaces the field value with `Hello, Pipeline!` and selects `Save`.
**Then**: The heading text becomes exactly `Hello, Pipeline!`, and the `Greeting` text field value is exactly `Hello, Pipeline!`.
Traces: SC-1 (GREETING-002 AC-1)
Check: interact_page

**Scenario**: Saved greeting remains after reload
**Given**: Visitor has saved `Hello, Pipeline!` successfully.
**When**: Visitor reloads the page.
**Then**: The heading text is exactly `Hello, Pipeline!`, and the `Greeting` text field value is exactly `Hello, Pipeline!`.
Traces: SC-2 (GREETING-002 AC-2)
Check: interact_page

**Scenario**: Save trims leading and trailing whitespace
**Given**: Stored greeting is `Hello, World!` and the page is open with the heading showing `Hello, World!`.
**When**: Visitor enters `  Trim me  ` in the `Greeting` text field and selects `Save`.
**Then**: The heading text becomes exactly `Trim me`, and the `Greeting` text field value is exactly `Trim me`.
Traces: SC-3 (GREETING-002 AC-3)
Check: interact_page

**Scenario**: Empty submit keeps existing greeting and focuses field
**Given**: Stored greeting is `Hello, World!` and the page is open with the heading showing `Hello, World!`.
**When**: Visitor clears the `Greeting` text field and submits the form.
**Then**: Focus is on the `Greeting` text field, the heading text remains exactly `Hello, World!`, the stored greeting remains `Hello, World!` after reload, and no visible error message is shown.
Traces: SC-4 (GREETING-002 AC-4)
Check: interact_page

**Scenario**: Mobile form controls stack full width at 520px
**Given**: Page is displayed at `520px` viewport width.
**When**: Visitor views the form.
**Then**: The `Greeting` text field and `button[type=submit]` are stacked vertically, each fills the form width, and each computed height is `48px`.
Traces: SC-5 (GREETING-002 AC-5)
Check: measure_styles

**Scenario**: Narrow mobile form controls stack full width at 320px
**Given**: Page is displayed at `320px` viewport width.
**When**: Visitor views the form.
**Then**: The `Greeting` text field and `button[type=submit]` are stacked vertically, each fills the form width, each computed height is `48px`, and the page has no horizontal scroll.
Traces: SC-5 (GREETING-002 AC-5), responsive NFR
Check: measure_styles

**Scenario**: Desktop form controls sit in one row with 12px gap
**Given**: Page is displayed at `521px` viewport width or wider.
**When**: Visitor views the form.
**Then**: The `Greeting` text field and `button[type=submit]` appear in one row, the computed gap between them is `12px`, and each computed height is `48px`.
Traces: SC-6 (GREETING-002 AC-6)
Check: measure_styles

**Scenario**: Text field focus has required outline
**Given**: Page is open and keyboard focus is before the form controls.
**When**: Visitor tabs to the `Greeting` text field.
**Then**: The focused text field has computed outline color `rgb(0, 0, 0)`, outline width `3px`, outline style not `none`, and outline offset `3px`.
Traces: SC-7 (GREETING-002 AC-7), accessibility NFR
Check: interact_page

**Scenario**: Save button focus has required outline
**Given**: Page is open and keyboard focus is on the `Greeting` text field.
**When**: Visitor tabs to the `Save` button.
**Then**: The focused `Save` button has computed outline color `rgb(0, 0, 0)`, outline width `3px`, outline style not `none`, and outline offset `3px`.
Traces: SC-7 (GREETING-002 AC-7), accessibility NFR
Check: interact_page

**Scenario**: Native form submit saves greeting
**Given**: Stored greeting is `Hello, World!` and the page is open with focus in the `Greeting` text field.
**When**: Visitor replaces the field value with `Hello from Enter` and presses Enter.
**Then**: The heading text becomes exactly `Hello from Enter`, and the `Greeting` text field value is exactly `Hello from Enter`.
Traces: GREETING-002 behaviour 3, story in scope native form submit
Check: interact_page

**Scenario**: Long saved greeting wraps without horizontal scroll
**Given**: Stored greeting is `Hello, World!` and the page is displayed at `320px` viewport width.
**When**: Visitor enters a long greeting of `Long greeting ` repeated 30 times and selects `Save`.
**Then**: The heading displays the full saved greeting, wraps within the page width, and the page has no horizontal scroll.
Traces: GREETING-002 boundary behaviour, responsive NFR
Check: interact_page

**Scenario**: Last successful save wins for two visitors
**Given**: Visitor A and Visitor B both have the page open with stored greeting `Hello, World!`.
**When**: Visitor A saves `First save`, then Visitor B saves `Second save`, and Visitor A reloads the page.
**Then**: Visitor A sees heading text exactly `Second save`, and the `Greeting` text field value is exactly `Second save`.
Traces: GREETING-002 conflict behaviour, assumption last successful save wins
Check: interact_page

**Scenario**: PUT greeting returns saved value
**Given**: Backend and database are available.
**When**: A client sends `PUT /v1/greeting` with JSON body `{"greeting":"Hello, Pipeline!"}`.
**Then**: The response status is `200`, and the response body is exactly `{"greeting":"Hello, Pipeline!"}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting trims outer whitespace and preserves inner whitespace
**Given**: Backend and database are available.
**When**: A client sends `PUT /v1/greeting` with JSON body `{"greeting":"  Hello,   Pipeline!  "}`.
**Then**: The response status is `200`, and the response body is exactly `{"greeting":"Hello,   Pipeline!"}` with the three internal spaces preserved.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: GET greeting returns last saved value
**Given**: `PUT /v1/greeting` has successfully saved `Hello, Pipeline!`.
**When**: A client sends `GET /v1/greeting`.
**Then**: The response status is `200`, and the response body is exactly `{"greeting":"Hello, Pipeline!"}`.
Traces: contract (GET /v1/greeting), SC-2 (GREETING-002 AC-2)
Check: fetch_url

**Scenario**: PUT greeting rejects empty trimmed value
**Given**: Backend and database are available and stored greeting is `Hello, World!`.
**When**: A client sends `PUT /v1/greeting` with JSON body `{"greeting":"   "}`.
**Then**: The response status is `422`, the response body is exactly `{"error":{"code":"VALIDATION_FAILED","message":"Greeting must not be empty."}}`, and a later `GET /v1/greeting` returns `{"greeting":"Hello, World!"}`.
Traces: contract (PUT /v1/greeting), GREETING-002 invalid input behaviour
Check: fetch_url

**Scenario**: PUT greeting rejects malformed JSON
**Given**: Backend and database are available.
**When**: A client sends `PUT /v1/greeting` with body `{` and `Content-Type: application/json`.
**Then**: The response status is `400`, and the response body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects wrong type
**Given**: Backend and database are available.
**When**: A client sends `PUT /v1/greeting` with JSON body `{"greeting":123}`.
**Then**: The response status is `400`, and the response body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting rejects unknown field
**Given**: Backend and database are available.
**When**: A client sends `PUT /v1/greeting` with JSON body `{"greeting":"Hello","extra":"ignored?"}`.
**Then**: The response status is `400`, and the response body is exactly `{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}`.
Traces: contract (PUT /v1/greeting)
Check: fetch_url

**Scenario**: PUT greeting failed query returns internal error envelope
**Given**: Backend is running and a greeting update query fails for a reason other than database connection refusal.
**When**: A client sends `PUT /v1/greeting` with JSON body `{"greeting":"Hello"}`.
**Then**: The response status is `500`, and the response body is exactly `{"error":{"code":"INTERNAL","message":"Internal server error."}}`.
Traces: contract (PUT /v1/greeting)
Check: manual

**Scenario**: PUT greeting database dependency unavailable returns unavailable envelope
**Given**: Backend is running and PostgreSQL refuses connections.
**When**: A client sends `PUT /v1/greeting` with JSON body `{"greeting":"Hello"}`.
**Then**: The response status is `503`, and the response body is exactly `{"error":{"code":"UNAVAILABLE","message":"Service unavailable."}}`.
Traces: contract (PUT /v1/greeting)
Check: manual
