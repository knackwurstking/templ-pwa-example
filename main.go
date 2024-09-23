package main

import (
	"embed"
	_ "embed"
	"log"
	"templ-pwa-example/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

//go:embed static/*
var static embed.FS

// NOTE: go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwindcss -m

func main() {
	e := echo.New()

	// greet := components.Greet("Lizzy The Cat", 2)
	index := components.Index()

	handler := templ.Handler(index)

	e.GET("/", echo.WrapHandler(handler))

	// TODO: Serve static files

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
