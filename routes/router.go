package routes

import (
	"log"
	"navpage/controllers"
	"navpage/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Routes() *fiber.App {
	// 初始化模板引擎
	engine := html.New("./views", ".html")
	engine.Reload(true) // 开发模式下启用模板重载
	engine.Debug(true)  // 启用调试模式

	// 创建 Fiber 实例
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// 添加中间件来记录请求
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("请求路径: %s", c.Path())
		return c.Next()
	})

	// 添加模板数据中间件
	app.Use(middleware.TemplateDataMiddleware())

	// 静态文件服务
	app.Static("/public", "./public")

	// 图片上传相关路由 - 放在认证中间件之前
	app.Post("/api/upload", controllers.Upload)

	app.Get("/login", controllers.Login)
	app.Post("/login", controllers.LoginPost)
	app.Get("/logout", controllers.Logout)

	app.Get("/", controllers.Image)

	// 需要认证的路由
	app.Use(middleware.AuthMiddleware())
	app.Get("/dashboard", controllers.Home)
	app.Get("/about", controllers.About)

	// 启动服务器
	log.Printf("服务器启动在 http://localhost:3000")
	return app
}
