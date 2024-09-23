package main

import (
	"log"
	"templ-pwa-example/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	greet := components.Greet("Lizzy The Cat", 2)
	handler := templ.Handler(greet)

	e.GET("/", echo.WrapHandler(handler))

	if err := e.Start(":8080"); err != nil {
		log.Fatalln(err)
	}
}
