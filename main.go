package main

import (
	"fmb/page"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/template/html"

	scss "github.com/wellington/go-libsass"
)

func makeCSS(dst string, src string) {
	// Make Ready Input and Output
	o, _ := os.Create(dst)
	f, _ := os.Open(src)

	// Get Compiler
	c, _ := scss.New(o, f)

	// Prepare for Find SCSS Directories
	var arr []string
	r := path.Dir(src)

	// Find SCSS Directories
	filepath.Walk(r, func(p string, i os.FileInfo, e error) error {
		if filepath.Ext(p) == ".scss" {
			d := path.Dir(p)
			if d != r {
				arr = append(arr, d)
			}
		}
		return nil
	})

	// Check Founds
	fmt.Println(arr)

	// Set Compile Option
	c.Option(
		scss.OutputStyle(
			scss.COMPRESSED_STYLE,
		),
		scss.IncludePaths(
			arr,
		),
	)

	// Do Compile
	if e := c.Run(); e != nil {
		log.Fatal(e)
	}
}

func main() {
	// Designate HTML Template Engine
	eng := html.New("./page", ".html")

	// Make Application Instance
	app := fiber.New(fiber.Config{
		Views: eng,
	})

	// Designate Statics
	app.Static("/", "./static")

	// Designate Favicon
	app.Use(favicon.New(
		favicon.Config{
			File: "./static/favicon.ico",
		},
	))

	// Set Routes of Application
	page.SetRoutes(app)

	// Make CSS
	makeCSS(
		"./static/miho.css",
		"./page/page.scss",
	)

	// Start Application Server
	app.Listen(":8080")
}
