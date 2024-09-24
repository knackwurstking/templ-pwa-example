.PHONY: build

BINARY_NAME=templ-pwa-example

clean:
	go clean

init:
	make clean
	npm install && \
		go mod tidy

generate:
	go mod tidy && \
		templ generate && \
		go generate

generate-assets:
	npx pwa-assets-generator

build:
	go mod tidy && \
       	templ generate && \
		go generate && \
		go build -ldflags="-w -s" -o ${BINARY_NAME}

run:
	go mod tidy && \
		templ generate --cmd="go generate" & templ generate --cmd="go run ."

build-html:
	rm -rf dist/* && \
		make generate && \
		go run . -html ./dist && \
		cp -r public/* dist/

android-init:
	npm install @capacitor/core && \
		npm install -D @capacitor/cli && \
		npm install @capacitor/android && \
		npm install @capacitor/assets --save-dev && \
		npx cap init && \
		npx cap add android

android-generate-assets:
	npx capacitor-assets generate --android

android-build:
	make android-generate-assets && \
		npx cap sync android && \
		echo 'Run "npx cap open android"'
