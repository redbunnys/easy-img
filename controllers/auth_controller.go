package controllers

import (
	"log"
	"navpage/middleware"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	log.Printf("正在渲染登录页面")
	err := c.Render("auth/login", fiber.Map{
		"Title": "登录",
	}, "layouts/main")
	if err != nil {
		log.Printf("渲染登录页面出错: %v", err)
		return err
	}
	return nil
}

func LoginPost(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// 获取session
	sess, err := middleware.Store.Get(c)
	if err != nil {
		log.Printf("获取session失败: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "获取session失败",
		})
	}
	if username == "admin" && password == "password" {
		sess.Set("authenticated", true)
		sess.Set("username", username) // 保存用户名到session
		if err := sess.Save(); err != nil {
			log.Printf("保存session失败: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "保存session失败",
			})
		}
		return c.Redirect("/dashboard")
	}

	return c.Render("auth/login", fiber.Map{
		"Title": "登录",
		"Error": "用户名或密码错误",
	}, "layouts/main")
}

func Logout(c *fiber.Ctx) error {
	// 清除 session
	sess, err := middleware.Store.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}

	sess.Delete("user_id")
	sess.Delete("username")
	if err := sess.Save(); err != nil {
		return c.Redirect("/login")
	}

	// 清除 cookie
	c.Cookie(&fiber.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // 设置为过期
		HTTPOnly: true,
	})

	return c.Redirect("/login")
}
