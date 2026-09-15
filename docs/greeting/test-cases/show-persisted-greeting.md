# Test cases — Show persisted greeting

Module: `greeting`
Function: Show persisted greeting
Story: `docs/greeting/stories/show-persisted-greeting.md`
Risk level: Medium. This story proves database-backed first paint and visible page structure; risk is data source/persistence and exact rendering, not complex interaction.

## Cases

**Scenario**: First load shows seeded greeting
**Given**: Stored greeting has never been changed and the seeded row contains `Hello, World!`
**When**: Visitor opens the greeting page
**Then**: The only large heading text is exactly `Hello, World!`
Traces: SC-1 (GREETING-001 AC-1)
Check: render_url

**Scenario**: First load shows changed stored greeting
**Given**: Stored greeting is `Good morning`
**When**: Visitor opens the greeting page
**Then**: The only large heading text is exactly `Good morning`
Traces: SC-2 (GREETING-001 AC-2)
Check: render_url

**Scenario**: Heading preserves internal spacing from storage
**Given**: Stored greeting is `Hello,   World!`
**When**: Visitor opens the greeting page
**Then**: The only large heading text preserves the three spaces between `Hello,` and `World!`
Traces: SC-3 (GREETING-001 AC-3)
Check: render_url

**Scenario**: Default screen contains one approved greeting section
**Given**: Page is displayed
**When**: Visitor inspects the default screen
**Then**: The screen contains exactly one centered `main > section`, exactly one large heading inside it, one text input with accessible label `Greeting`, and one button with visible text `Save`
Traces: SC-4 (GREETING-001 AC-4)
Check: render_url

**Scenario**: Default screen has no extra page structure
**Given**: Page is displayed
**When**: Visitor views the page
**Then**: No `nav` element is present, and `main` contains no section other than the greeting section
Traces: SC-5 (GREETING-001 AC-5)
Check: render_url

**Scenario**: Default screen uses approved minimal colors
**Given**: Page is displayed
**When**: Visitor views the page
**Then**: The page background is `#FFFFFF` and text color is `#000000`
Traces: SC-5 (GREETING-001 AC-5)
Check: measure_styles

**Scenario**: Long stored greeting wraps without horizontal page scroll
**Given**: Stored greeting is a long sentence repeated enough to exceed one line at `320px` viewport width
**When**: Visitor opens the greeting page at `320px` viewport width
**Then**: The heading wraps within the viewport, document horizontal scroll width is not greater than viewport width, and no horizontal page scroll is present
Traces: SC-6 (GREETING-001 boundary)
Check: measure_styles

**Scenario**: API returns current greeting success envelope
**Given**: Stored greeting is `Good morning`
**When**: Client requests `GET /v1/greeting`
**Then**: Response status is `200` and JSON body is exactly `{"greeting":"Good morning"}`
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: API returns seeded greeting before any change
**Given**: Stored greeting has never been changed and the seeded row contains `Hello, World!`
**When**: Client requests `GET /v1/greeting`
**Then**: Response status is `200` and JSON body is exactly `{"greeting":"Hello, World!"}`
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: API returns internal error envelope for failed greeting query
**Given**: Database is reachable, but the greeting query fails
**When**: Client requests `GET /v1/greeting`
**Then**: Response status is `500` and JSON body is exactly `{"error":{"code":"INTERNAL","message":"Internal server error."}}`
Traces: contract (GET /v1/greeting)
Check: fetch_url

**Scenario**: API returns unavailable envelope when database dependency is unavailable
**Given**: Backend is running, but the database dependency refuses connections
**When**: Client requests `GET /v1/greeting`
**Then**: Response status is `503` and JSON body is exactly `{"error":{"code":"UNAVAILABLE","message":"Service unavailable."}}`
Traces: contract (GET /v1/greeting)
Check: manual

## Coverage check

- SC-1 covered by: First load shows seeded greeting
- SC-2 covered by: First load shows changed stored greeting
- SC-3 covered by: Heading preserves internal spacing from storage
- SC-4 covered by: Default screen contains one approved greeting section
- SC-5 covered by: Default screen uses approved minimal style and no extra page structure
- SC-6 covered by: Long stored greeting wraps without horizontal page scroll
- Contract `GET /v1/greeting` success covered by API success cases
- Contract `GET /v1/greeting` `500 INTERNAL` covered by failed query case
- Contract `GET /v1/greeting` `503 UNAVAILABLE` covered by unavailable dependency case
- No visitor input, denied role, or malformed request exists for this read-only page function or `GET /v1/greeting`
