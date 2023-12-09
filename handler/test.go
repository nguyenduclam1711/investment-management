package handler

import "github.com/gofiber/fiber/v2"

func TestPage(c *fiber.Ctx) error {
	return c.Render("pages/test", fiber.Map{})
}
