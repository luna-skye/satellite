package routes

import (
	"github.com/gofiber/fiber/v2"
)


func IndexPage(c *fiber.Ctx) error {
  return c.Render("views/index", fiber.Map{}, "layouts/main")
}
