package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/investment-management/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.IndexPage)
	app.Get("/test", handler.TestPage)
}
