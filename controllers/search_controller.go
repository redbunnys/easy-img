package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Search(c *fiber.Ctx) error {
	return c.Render("search/search", fiber.Map{
		"Title": "搜索",
	})
}
