.PHONY: build

BINARY_NAME=temple-pwa-example

generate-pwa-assets:
	npm install
	npx pwa-assets-generator

generate-html:
	make clean
	rm -r dist/* && \
		go mod tidy && \
		go run . -html ./dist && \
		cp -r public/* dist/

build:
	make clean
	go mod tidy && \
       	templ generate && \
		go generate && \
		go build -ldflags="-w -s" -o ${BINARY_NAME}

dev:
	make clean
	go mod tidy && \
		templ generate --watch --cmd="go generate" &\
		templ generate --watch --cmd="go run ."

clean:
	go clean
