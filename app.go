package main

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
	"github.com/nguyenduclam1711/investment-management/database"
	"github.com/nguyenduclam1711/investment-management/router"
)

func main() {
	database.ConnectDB()
	engine := html.New("web/template", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// use default compress
	app.Use(compress.New())
	app.Use(logger.New())

	app.Static("/static", "web/static")

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
