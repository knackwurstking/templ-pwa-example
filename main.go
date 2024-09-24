package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"os"
	"templ-pwa-example/components"

	"github.com/a-h/templ"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

const (
	ErrorCodeOK = 0
)

//go:embed public
var public embed.FS

// NOTE: go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwindcss -m

func main() {
	flags := newFlags()
	flags.Parse()

	if flags.html != "" {
		// TODO: Render html to destination

		os.Exit(ErrorCodeOK)
	}

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

	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatalln(err)
	}
}

type flags struct {
	html string
}

func newFlags() *flags {
	return &flags{
		"",
	}
}

func (flags *flags) Parse() {
	flag.StringVar(&flags.html, "html", flags.html,
		"generate html to destination and exit")

	flag.Usage = func() {
		fmtHeader := color.New(color.Bold, color.Underline, color.FgWhite)
		fmtHeader.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
	}

	flag.Parse()
}
