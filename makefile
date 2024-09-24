.PHONY: build

BINARY_NAME=templ-pwa-example

generate-pwa-assets:
	npm install
	npx pwa-assets-generator

generate-html:
	make clean
	rm -rf dist/* && \
		go mod tidy && \
		go run . -html ./dist && \
		cp -r public/* dist/

build:
	make clean
	go mod tidy && \
       	templ generate && \
		go generate && \
		go build -ldflags="-w -s" -o ${BINARY_NAME}

# FIXME: Watcher spits an error: "too many open files"
dev:
	make clean
	go mod tidy && \
		templ generate --watch --cmd="go generate" &\
		templ generate --watch --cmd="go run ."

clean:
	go clean
