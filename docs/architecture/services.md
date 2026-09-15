# Service contracts

Backend routes omit deployment `/api` prefix.

## Shared errors

All non-success responses use:

```json
{"error":{"code":"MALFORMED_REQUEST","message":"Request body is malformed."}}
```

| HTTP | Code | Message | When |
|---|---|---|---|
| 400 | `MALFORMED_REQUEST` | `Request body is malformed.` | Bad JSON, wrong type, unknown field |
| 422 | `VALIDATION_FAILED` | `Greeting must not be empty.` | Trimmed greeting is empty |
| 500 | `INTERNAL` | `Internal server error.` | Query or unexpected server failure |
| 503 | `UNAVAILABLE` | `Service unavailable.` | Database dependency unavailable |

## Greeting

### `GET /v1/greeting`

Returns current shared greeting. Auth: none; public Visitor access.

Success `200`:

```json
{"greeting":"Hello, World!"}
```

Returns `503 UNAVAILABLE` when database cannot be reached; `500 INTERNAL` for other failed queries.

### `PUT /v1/greeting`

Replaces shared greeting. Request JSON must contain only string field `greeting`.

```json
{"greeting":"Hello, Pipeline!"}
```

Server trims leading and trailing whitespace, rejects empty result, preserves internal whitespace, punctuation, and capitalization. Last successful write wins.

Success `200`:

```json
{"greeting":"Hello, Pipeline!"}
```

Errors use shared envelope: `400 MALFORMED_REQUEST`, `422 VALIDATION_FAILED`, `503 UNAVAILABLE`, or `500 INTERNAL`.

## Health

### `GET /healthz`

Returns `200` with plain `ok` only after migrations and database probe succeed. Returns `503 UNAVAILABLE` otherwise. This endpoint is operational, not versioned product API.

## Edit persisted greeting implementation decision

`PUT /v1/greeting` is unauthenticated because product has one shared public Visitor role. It updates only `greetings.text` for `id = 1`, sets `updated_at` to current time, and returns resulting trimmed text. Use one parameterized update statement; concurrent successful requests have no version check, so last completed write wins.

Mock review: `GreetingResponse` is sound and matches contract response. Its `getMockGreeting` and `saveMockGreeting` localStorage functions are temporary UI behavior. Backend integration replaces those imports with API client calls; no component props or JSON fields change.

## Show persisted greeting implementation decision

`GET /v1/greeting` is unauthenticated because product has one shared public Visitor role. It performs one parameterized query for `greetings.text` where `id = 1`, returning `200` `{ "greeting": string }`. The UI mock shape is sound and already matches this contract; backend integration replaces `getGreeting` mock import with API client call only.

No request body. Errors use shared envelope: `503 UNAVAILABLE` only when database connection is unavailable; failed query after connection returns `500 INTERNAL`. Missing row also returns `500 INTERNAL`: bootstrap seed is required invariant and no product error state exists.
