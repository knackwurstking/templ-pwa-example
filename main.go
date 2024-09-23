package main

import (
	"embed"
	"io/fs"
	"log"
	"templ-pwa-example/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

//go:embed public
var public embed.FS

// NOTE: go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwindcss -m

func main() {
	e := echo.New()

	// greet := components.Greet("Lizzy The Cat", 2)
	index := components.Index()

	handler := templ.Handler(index)

	e.GET("/", echo.WrapHandler(handler))

	fsys, err := fs.Sub(public, "public")
	if err != nil {
		log.Fatalln(err)
	}
	e.StaticFS("", fsys)

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
