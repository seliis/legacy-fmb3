package page

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

type Menu struct {
	Menu string
}

func SetRoutes(app *fiber.App) {
	var Menus [6]Menu
	Menus[0].Menu = "aviators"
	Menus[1].Menu = "training"
	Menus[2].Menu = "document"
	Menus[3].Menu = "campaign"
	Menus[4].Menu = "database"
	Menus[5].Menu = "signs-in"

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home/home", fiber.Map{
			"Menus": Menus,
		}, "page")
	})

	checkPage := func(n string) bool {
		var found bool = false
		filepath.Walk("./page", func(p string, i os.FileInfo, e error) error {
			if i.Name() == n+".html" {
				found = true
			}
			return nil
		})
		return found
	}

	app.Get("/:page", func(c *fiber.Ctx) error {
		p := c.Params("page")
		for _, arr := range Menus {
			if arr.Menu == p && checkPage(p) {
				fmt.Println("passed")
				return c.Render(p+"/"+p, fiber.Map{
					"Menus": Menus,
				}, "page")
			}
		}
		return c.Render("wrong/wrong", fiber.Map{
			"Menus": Menus,
		}, "page")
	})
}
