# Architecture overview

## Stack

| Part | Choice | Reason |
|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3, ESLint | Required UI stack; standalone output fits committed image. |
| Backend | Go 1.25, `net/http`, `database/sql`, pgx stdlib adapter | Small API needs no framework. |
| Database | PostgreSQL 16 | Greeting must survive reloads. |

## Layout

```text
code/backend/cmd/api/main.go    server, migration runner, health endpoint
code/backend/migrations/        ordered SQL migration pairs
code/frontend/app/              layout, composition root, global design tokens
code/frontend/components/       one default-export component per story
code/frontend/lib/              API clients; mock modules before API integration
```

`app/page.tsx` remains server-side composition root. Interactive story components begin with `"use client"`. Components use default function exports. Shared visual values use tokens from `app/globals.css`; CSS modules contain no hardcoded visual values or token fallbacks.

## Data and API flow

Frontend calls versioned backend routes through `NEXT_PUBLIC_API_URL`; deployment proxy removes `/api` before backend routing. Backend reads `DATABASE_URL`, applies ordered migrations tracked in `schema_migrations`, then serves traffic. `/healthz` returns 200 only after migrations and `SELECT 1` succeed. Greeting schema and API contract: `erd.md`, `services.md`.

## Decisions

| Decision | Rejected | Tradeoff |
|---|---|---|
| One shared greeting row | Per-visitor rows | Matches no-sign-in scope; no ownership/history. |
| SQL migrations at startup | External migration job | Empty runtime DB becomes usable; startup owns migration failure. |
| `net/http` routes | Web framework | Fewer dependencies; routing remains small. |
| Last successful save wins | Optimistic locking | Matches SRS; concurrent saves can overwrite. |

## Naming and compatibility

Go packages use lowercase names; exported Go symbols use PascalCase. SQL uses snake_case. Migration files use `YYYYMMDDHHMMSS_name.up.sql` and matching `.down.sql`. API JSON fields use lower camel case. API paths begin `/v1`, never `/api`. Migrations run in transactions except files containing `CREATE INDEX CONCURRENTLY`.

## Environment

| Service | Variables |
|---|---|
| Backend | `DATABASE_URL`, `PORT`, optional `APP_PORT` fallback |
| Frontend | `NEXT_PUBLIC_API_URL`, `API_ORIGIN` |
| Compose | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`; optional ports and resource limits |

Examples list every key and no secrets. Docker Compose supplies local defaults. Production provides database URL and public API URL through deployment environment.

## Run

Copy `.env.example` to `.env` if defaults need changes. Run `docker compose --profile local up --build`; open `http://localhost:3000`. Stop with `docker compose --profile local down`. Use `docker compose --profile local down -v` only when deleting local database data is intended.

## Rollout

First migration seeds `Hello, World!` idempotently. Schema changes require paired migrations and backwards-compatible API changes while old frontend images may remain active.
