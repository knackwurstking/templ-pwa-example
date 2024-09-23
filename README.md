# PWA (templ) Example

<!--toc:start-->

- [PWA (templ) Example](#pwa-templ-example)
  - [Getting Started](#getting-started)
    - [Installing the templ cli](#installing-the-templ-cli)
    - [Project structure](#project-structure)
    - [Generate templ files](#generate-templ-files)
    - [Generate pwa assets](#generate-pwa-assets)
  - [How I use this repo](#how-i-use-this-repo)
  - [TODO](#todo)

<!--toc:end-->

## Getting Started

### Installing the templ cli

```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### Project structure

- `components/` — templ components
- `public/` — Files available to the public
- `.gitignore` — Some stuff are not worthy of being committed
- `main.go` — The entry point to our application
- `Makefile` — A runner and builder script
- `pwa-assets.config.js` — PWA configuration for the `make generate-pwa-assets`
  script (`@vite-pwa/assets-generator`)

### Generate templ files

```bash
templ generate
```

### Generate pwa assets

```bash
make generate-pwa-assets
```

## How I use this repo

- Clone with `git clone https://github.com/knackwurstking/templ-pwa-example`
- Copy all the files from this repo to the new project and let's go

## TODO

- [ ] Test if the app is working without the service worker file
  - [ ] android
    - [ ] online
    - [ ] offline
  - [ ] iOS
    - [ ] online
    - [ ] offline
  - [ ] macOS, tested with Chrome
    - [x] online
    - [ ] offline
- [ ] Change the app icon
