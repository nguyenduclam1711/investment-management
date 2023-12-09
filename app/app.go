package app

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/template/html/v2"
	"github.com/nguyenduclam1711/investment-management/app/controller"
)

func App() *fiber.App {
	engine := html.New("app/view/templates", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// use default compress
	app.Use(compress.New())

	app.Static("/statics", "web/statics")

	// controllers handler
	controller.Main(app)

	return app
}
