# PWA (templ) Example

<!--toc:start-->

- [PWA (templ) Example](#pwa-templ-example)
  - [Getting Started](#getting-started)
    - [Installing the templ cli](#installing-the-templ-cli)
    - [Project structure](#project-structure)
    - [Generate templ files](#generate-templ-files)
    - [Generate pwa assets](#generate-pwa-assets)

<!--toc:end-->

## Getting Started

### Installing the templ cli

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Project structure

- `components/` — templ components.
- `db/` — Database access code used to access the spending logs.
- `handlers/` — HTTP handlers.
- `static/` — Files that are available to the public.
- `services/` — Services used by the handlers.
- `.gitignore` — Some stuff are not worthy of being committed.
- `Dockerfile` — Container configuration to run the application with the
  glorious Docker.
- `Makefile` — A runner and builder script.
- `main.go` — The entry point to our application.

### Generate templ files

```bash
templ generate
```

### Generate pwa assets

```bash
make generate-pwa-assets
```
