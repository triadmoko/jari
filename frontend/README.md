# sv

Everything you need to build a Svelte project, powered by [`sv`](https://github.com/sveltejs/cli).

## Creating a project

If you're seeing this, you've probably already done this step. Congrats!

```sh
# create a new project
bunx sv create my-app
```

To recreate this project with the same configuration:

```sh
# recreate this project
bun x sv@0.15.3 create --template minimal --types ts --add prettier vitest="usages:unit,component" tailwindcss="plugins:none" sveltekit-adapter="adapter:auto" eslint --install bun app
```

## Developing

Once you've created a project and installed dependencies with `bun install`, start a development server:

```sh
bun run dev

# or start the server and open the app in a new browser tab
bun run dev --open
```

## Building

To create a production version of your app:

```sh
bun run build
```

You can preview the production build with `bun run preview`.

> The build is a static SPA (`@sveltejs/adapter-static`, fallback `index.html`) in `build/`. In production it is embedded into the Go binary and served by Gin: run `make web-build` or `make build` from the repo root. During development run `make dev` (backend, :8080) next to `bun run dev` (Vite proxies `/api`).
