package handler

import "github.com/gofiber/fiber/v2"

func IndexPage(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{
		"Title": "Index",
	})
}
