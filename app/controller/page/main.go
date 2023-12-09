package controlllerpage

import "github.com/gofiber/fiber/v2"

func Main(app *fiber.App) {
	controllers := []controller{
		renderIndexPage,
		renderTestPage,
	}
	runControllers(app, controllers)
}
