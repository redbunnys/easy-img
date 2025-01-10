package middleware

import "github.com/gofiber/fiber/v2"

// TemplateDataMiddleware 处理所有模板共享的数据
func TemplateDataMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 获取 session
		sess, err := Store.Get(c)
		if err != nil {
			return err
		}

		// 创建共享数据
		sharedData := fiber.Map{
			"IsAuthenticated": sess.Get("authenticated") != nil,
			"Username":        sess.Get("username"),
		}

		// 将共享数据存储到上下文中
		c.Locals("shared", sharedData)

		return c.Next()
	}
}
