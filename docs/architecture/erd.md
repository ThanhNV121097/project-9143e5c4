# ERD

## `greetings`

Single shared row holds current public greeting.

| Column | Type | Constraints |
|---|---|---|
| `id` | `smallint` | Primary key; must equal `1` |
| `text` | `text` | `NOT NULL`; trimmed value must not be empty |
| `updated_at` | `timestamptz` | `NOT NULL`; defaults to current time |

Migration seeds `(1, 'Hello, World!')` with conflict-safe insert. No relationships: project has one shared greeting and no users.

## `schema_migrations`

Migration runner owns this technical table.

| Column | Type | Constraints |
|---|---|---|
| `version` | `text` | Primary key; migration filename |
| `applied_at` | `timestamptz` | `NOT NULL`; defaults to current time |

`schema_migrations` has no domain relationship. It makes migration re-runs no-ops.

## Edit persisted greeting migration plan

Mock `GreetingResponse` is `{ "greeting": string }`; this matches existing API response exactly. Mock localStorage persistence is UI-only and must be replaced by `PUT /v1/greeting`; no frontend response-shape change.

Forward migration creates `schema_migrations`, creates `greetings` with `id smallint PRIMARY KEY CHECK (id = 1)`, `text text NOT NULL CHECK (btrim(text) <> '')`, and `updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP`, then inserts `(1, 'Hello, World!')` using `ON CONFLICT (id) DO NOTHING`. No foreign keys or secondary indexes: reads and writes target primary key `id = 1`.

Backward migration drops `greetings`; then drops `schema_migrations` only if rolling back migration runner bootstrap. This destroys saved shared greeting, so use only before production data matters. Forward is safe on populated databases: table creation is additive and seed insert does not overwrite existing row.

## Show persisted greeting design

Mock module `code/frontend/lib/mock/show-persisted-greeting.ts` defines `GreetingResponse` as `{ "greeting": string }` and returns seeded text. Shape is sound: it matches the existing `GET /v1/greeting` response and heading needs one non-null string. Backend integration replaces only mock import with API client; no component props or JSON fields change.

No schema changes: this story reads `greetings.text` from single row `id = 1`. Primary key serves query `SELECT text FROM greetings WHERE id = 1`; no secondary index or foreign key applies.

Migration: none for this read-only story. It depends on shared greeting bootstrap migration above. Forward is safe on populated databases because it makes no change; backward has no story-specific step.
