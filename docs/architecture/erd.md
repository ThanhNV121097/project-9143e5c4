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
