# PWA (templ) Example

> Just some boilerplate code.

<!--toc:start-->

- [PWA (templ) Example](#pwa-templ-example)
  - [Getting Started](#getting-started)
  - [Files to edit](#files-to-edit)
  - [Build Android App](#build-android-app)

<!--toc:end-->

## Getting Started

Project structure

- `components/` — templ components
- `assets/` — Only used for generating file icons with `capacitor-assets` command
- `public/` — Files available to the public
- `.gitignore` — Some stuff are not worthy of being committed
- `main.go` — The entry point to our application
- `Makefile` — A runner and builder script
- `pwa-assets.config.js` — PWA configuration for the `make generate-pwa-assets`
    script (`@vite-pwa/assets-generator`)

Install the templ cli

```build
go install github.com/a-h/templ/cmd/templ@latest
```

Initialize

```bash
make init
```

Build

```bash
make generate
make build
```

Build html and store in "./dist/"

```bash
make build-html
```

Build for Android

```bash
make android-init
make android-generate-assets
make android-build
```

## Files to edit

- [makefile](makefile)
  - Edit: `BINARY_NAME=`
  - ... [@TODO]
- ... [@TODO]

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

- Change PWA app icons background to black
