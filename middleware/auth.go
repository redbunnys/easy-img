package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New()

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取session
		sess, err := Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "获取session失败",
			})
		}

		// 检查用户是否已登录
		auth := sess.Get("authenticated")
		username := sess.Get("username")
		log.Printf("username: %v", username)
		if auth == nil {
			return c.Redirect("/login")
		}

		return c.Next()
	}
}
