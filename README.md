# PWA (templ) Example

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
- `Makefile` — A runner and builder script to run the templ thing alongside
  tailwindcss, and it has the build commands.
- `main.go` — The entry point to our application.

### Generate templ files with

```bash
templ generate
```
