package controlllerpage

import "github.com/gofiber/fiber/v2"

type renderPageFunc func(path string, name string, bind interface{})

type controller func(renderPage renderPageFunc)

func renderPageWrapper(app *fiber.App) renderPageFunc {
	return func(path string, name string, bind interface{}) {
		app.Get("/"+path, func(c *fiber.Ctx) error {
			return c.Render("pages/"+name, bind, "layouts/main")
		})
	}
}

func runControllers(app *fiber.App, slices []controller) {
	renderPage := renderPageWrapper(app)
	for _, controller := range slices {
		controller(renderPage)
	}
}
