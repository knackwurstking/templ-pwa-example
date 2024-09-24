# PWA (templ) Example

<!--toc:start-->

- [PWA (templ) Example](#pwa-templ-example)
  - [Getting Started](#getting-started)
    - [Installing the templ cli](#installing-the-templ-cli)
    - [Project structure](#project-structure)
    - [Generate templ files](#generate-templ-files)
    - [Generate PWA assets](#generate-pwa-assets)
    - [Start the dev server](#start-the-dev-server)
    - [Build](#build)
  - [How I use this repo](#how-i-use-this-repo)
  - [Build Android App [work-in-progress]](#build-android-app-work-in-progress)
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

### Generate PWA assets

```bash
make generate-pwa-assets
```

### Start the dev server

```bash
make dev
```

### Build

```bash
make
```

## How I use this repo

- Clone with `git clone https://github.com/knackwurstking/templ-pwa-example`
- Copy all the files from this repo to the new project and let's go

## Build Android App

- <https://capacitorjs.com/docs/getting-started>
- <https://capacitorjs.com/docs/guides/splash-screens-and-icons>

```bash
# Install dependencies
npm i @capacitor/core
npm i -D @capacitor/cli

# Initialize
npx cap init
# Name: Templ PWA Example
# Package ID: templpwaexample.knackwurstking.com
# Web asset directory: dist

# Add the android platform
npm i @capacitor/android
npx cap add android
npx cap sync android

# Generate Android icons
npm install @capacitor/assets --save-dev
# NOTE: The "./assets/" directory will be used for this
npx capacitor-assets generate --android

# Open android studio and edit or build the app
npx cap open android
```

> Edit [capacitor.config.json](capacitor.config.json) to your needs.

## TODO

- [x] Add a script for generate html
  - [x] Add some stuff to allow a capacitor android build
- [ ] Change the app icon
- [ ] Test PWA
  - [x] macOS, tested with Chrome (http + https)
    - [x] online
    - [x] offline
  - [ ] android (https only)
    - [ ] online
    - [ ] offline
  - [ ] iOS (https only)
    - [ ] online
    - [ ] offline
