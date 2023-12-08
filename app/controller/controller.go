package controller

import (
	"github.com/gofiber/fiber/v2"
	controlllerpage "github.com/nguyenduclam1711/investment-management/app/controller/page"
)

type Controller func(renderPage controlllerpage.RenderPage)

func runControllers(app *fiber.App, slices []Controller) {
	renderPage := controlllerpage.RenderPageWrapper(app)
	for _, controller := range slices {
		controller(renderPage)
	}
}

func Handler(app *fiber.App) {
	controllers := []Controller{
		controlllerpage.Index,
		controlllerpage.Test,
	}
	runControllers(app, controllers)
}
