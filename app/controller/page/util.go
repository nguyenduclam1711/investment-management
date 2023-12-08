package controlllerpage

import "github.com/gofiber/fiber/v2"

type RenderPage func(path string, name string, bind interface{})

func RenderPageWrapper(app *fiber.App) RenderPage {
	return func(path string, name string, bind interface{}) {
		app.Get("/"+path, func(c *fiber.Ctx) error {
			return c.Render("pages/"+name, bind, "layouts/main")
		})
	}
}
