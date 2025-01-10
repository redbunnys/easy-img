package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.Render("home/index", MergeWithShared(c, fiber.Map{
		"Title":   "首页",
		"Active":  "home",
		"Content": "这是首页内容",
	}), "layouts/main")
}

func About(c *fiber.Ctx) error {
	return c.Render("home/about", MergeWithShared(c, fiber.Map{
		"Title":   "关于我们",
		"Active":  "about",
		"Content": "这是关于页面内容",
	}), "layouts/main")
}
