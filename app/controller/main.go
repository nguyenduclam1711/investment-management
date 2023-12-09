package controller

import (
	"github.com/gofiber/fiber/v2"
	controlllerpage "github.com/nguyenduclam1711/investment-management/app/controller/page"
)

func Main(app *fiber.App) {
	controlllerpage.Main(app)
}
