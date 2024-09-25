package main

import (
	"context"
	"embed"
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"templ-pwa-example/components"

	"github.com/a-h/templ"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

//go:embed public
var public embed.FS

// NOTE: go:generate npx tailwindcss build -i static/css/style.css -o static/css/tailwindcss -m

func main() {
	flags := newFlags()
	flags.Parse()

	// greet := components.Greet("Lizzy The Cat", 2)
	index := components.Index()

	if flags.html != "" {
		err := os.MkdirAll(flags.html, os.ModeSticky|os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create directories %s: %v", flags.html, err)
		}

		f, err := os.Create(filepath.Join(flags.html, "index.html"))
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
		}

		err = index.Render(context.Background(), f)
		if err != nil {
			log.Fatalf("Failed to write output file: %v", err)
		}

		os.Exit(0)
	}

	e := echo.New()

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
