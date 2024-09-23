.PHONY: build

BINARY_NAME=temple-pwa-example

generate-pwa-assets:
	npx pwa-assets-generator

build:
	make clean
	make generate-pwa-assets
	go mod tidy && \
       	templ generate && \
		go generate && \
		go build -ldflags="-w -s" -o ${BINARY_NAME}

dev:
	templ generate --watch --cmd="go generate" &\
	templ generate --watch --cmd="go run ."

clean:
	go clean
