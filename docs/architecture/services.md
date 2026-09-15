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

Returns current shared greeting.

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
