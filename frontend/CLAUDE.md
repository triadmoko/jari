---
description:
alwaysApply: true
---

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Claude Context MCP — WAJIB

> **ATURAN KERAS**: Setiap kali perlu menjelajahi atau memahami kode di repo ini, **WAJIB** gunakan Claude Context MCP terlebih dahulu. DILARANG langsung memakai `grep`, `find`, Bash, atau spawn Agent Explore sebelum mencoba Claude Context.

### Langkah wajib sebelum eksplorasi codebase:

1. Load tool via `ToolSearch` dengan query: `select:mcp__claude-context__search_code,mcp__claude-context__get_indexing_status`
2. Jalankan `get_indexing_status` untuk memastikan index siap
3. Gunakan `search_code` dengan natural-language query dan `path` = absolute path repo ini

### Tool yang tersedia:

- **`search_code`** — Semantic search dengan `query` natural language; set `path` ke absolute path repo ini. Gunakan untuk menemukan fitur, route, use case, handler, entity, dan call site.
- **`get_indexing_status`** — Cek apakah workspace sudah terindeks dan siap; jika belum, jalankan `index_codebase` dulu.
- **`index_codebase`** — Index atau re-index repo ini setelah perubahan besar atau saat search gagal karena belum terindeks.

### Fallback:

Hanya boleh fallback ke `grep`/`find`/Bash/Agent **jika** Claude Context MCP tidak tersedia (server down/disabled) atau hasil search benar-benar tidak cukup setelah beberapa query.

## Commands

Package manager: **bun** (bun.lock). The root `Dockerfile` builds the frontend in an `oven/bun:1.2-alpine` stage. Use `bun install`, `bun run <script>`.

- `bun run dev` — start vite dev server (proxies `/api` and `/swagger` to the Go backend at `localhost:8080`, so run `make dev` in the repo root too)
- `bun run build` — static SPA build (output in `build/`, no Node/Bun runtime needed). From the repo root, `make web-build` also copies it into `internal/shared/delivery/http/webui/dist`, which Gin embeds (`go:embed`) and serves with an `index.html` fallback
- `bun run preview` — preview production build
- `bun run check` — `svelte-kit sync` + `svelte-check` against `tsconfig.json` (this is the type-check; there is no separate `tsc` step)
- `bun run lint` / `bun run format` — Prettier check / write
- `bun run swag-local` — regenerate `src/lib/api/api.ts` from the local backend (`localhost:8080/swagger/doc.json`); you can also point `-p` at `../docs/swagger.json`. Run this whenever the backend contract changes; do not hand-edit `src/lib/api/api.ts`.
- `bun run test` — vitest (unit tests live next to the code as `*.test.ts`)

## Architecture

### Stack

SvelteKit 2 + Svelte 5 with **runes mode forced on** for all non-`node_modules` files (see `compilerOptions.runes` in `vite.config.ts`). TypeScript strict. Tailwind CSS **v4** via `@tailwindcss/vite` — there is no `tailwind.config.*`; tokens are declared as CSS custom properties + `@theme inline` in `src/routes/layout.css`. shadcn-svelte (style `vega`, base color `mist`, icon library `tabler`) provides UI primitives in `src/lib/components/ui/`. **This is a SPA**: `@sveltejs/adapter-static` with `fallback: 'index.html'` (configured in `vite.config.ts`) and `ssr = false` in `src/routes/+layout.ts`. There is no server runtime, so **never add `+page.server.ts`, `+layout.server.ts`, `hooks.server.ts`, `$env/*/private` or `src/lib/server/`** — they cannot work. Everything runs in the browser.

Path alias: `$lib` → `src/lib` (set in `vite.config.ts`; there is no `svelte.config.js`).

Config: `VITE_API_BASE_URL` (see `.env.example`) is the API origin, inlined at build time. Leave empty for same-origin (production and dev proxy).

### Feature layout (DDD-ish)

Domain logic lives under `src/lib/features/<feature>/` with four sub-folders:

- `domain/` — types and repository **interfaces** (e.g. `auth.types.ts` declaring `AuthRepository`). No I/O.
- `application/` — pure use-cases and constants (`login.use-case.ts`, `auth.constants.ts`). Validation + orchestration; takes a repository as a parameter.
- `infrastructure/` — repository implementations. Use `api-*.client.ts` (or `*.repository.ts`) built on the shared `api` from `$lib/api/client`, which attaches the bearer token automatically. May also include a `mock-*.repository.ts` for offline/dev.
- `components/` — feature-specific Svelte components.

Each feature exposes a single barrel `index.ts` (constants, types, use-cases, repo).

When adding a feature, mirror this shape and keep use-cases free of HTTP/SDK calls — they should depend on the domain interface only.

### Authentication & sessions

- Session `{ token, user, expiresAt }` is stored in `localStorage` under `jari_session` (`src/lib/features/auth/infrastructure/session.storage.ts`; reads/writes are try/catch-guarded and expired or malformed data is dropped). TTL is a **sliding 12h** window, refreshed on each authenticated API request.
- Reactive state is the `session` store in `application/session.svelte.ts` (`session.login(token, user)`, `session.logout()`, `session.token`, `session.user`). `login` receives `token` and `user` (incl. `role`) from `POST /api/v1/auth/login`.
- Route guards run client-side in `src/routes/+layout.ts` via the pure `resolveGuard()` in `application/route-guard.ts`: `PROTECTED_PREFIXES = ['/admin', '/profile']` redirect to `/auth/login?redirectTo=…` when logged out; `GUEST_ONLY_PATHS = ['/auth/login']` redirect to `AUTH_ROUTES.home` when logged in. This is UX only — the backend enforces auth.
- After login use `postLoginTarget(session, redirectTo)`: a safe `redirectTo` wins, otherwise admins → `/admin`, others → `/`. Always pass `redirectTo` through `safeRedirectTarget` (blocks open redirects).
- `localStorage` is readable by any XSS, so keep dependencies lean and never render untrusted HTML.

### API client

`src/lib/api/api.ts` is **auto-generated** by `swagger-typescript-api` from the backend Swagger doc — regenerate, don't edit. `src/lib/api/client.ts` exports the singleton `api` (`new Api(...)`); its axios interceptors add `Authorization: Bearer <token>` and, on a 401 for an authenticated session, clear it and `goto` the login page. Generated paths already include `/api/v1`, so the base URL is only the origin.

### Routes

`(guest)` is a layout group for the public landing page. `admin/` and `profile/` are protected by the client-side guard in `+layout.ts`. All data fetching happens in the browser through the shared `api` client. Unknown paths are served `index.html` by the backend (and by `adapter-static`'s fallback) and resolved by the client router.

## Conventions

- **User-facing strings are Indonesian** (see `auth.constants.ts` `AUTH_MESSAGES`, error messages in repositories). Match the existing tone when adding new strings.
- Prettier enforces tabs, single quotes, no trailing comma, `printWidth: 100`. The `prettier-plugin-tailwindcss` reads classes from `src/routes/layout.css` for class sorting. Run `bun run format` before committing.
- Imports use explicit `.js` extensions on relative paths (TS `rewriteRelativeImportExtensions` is on) — follow the existing pattern.
- Repository methods return discriminated unions (`{ ok: true, ... } | { ok: false, code, message }`) instead of throwing. Map HTTP statuses to domain codes (`UNAUTHORIZED`, `FORBIDDEN`, `INVALID_CREDENTIALS`, `VALIDATION_ERROR`, `UNKNOWN_ERROR`) with localized messages — keep new repos consistent with this.
