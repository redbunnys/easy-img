package controllers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Image 渲染图片上传页面
func Image(c *fiber.Ctx) error {
	return c.Render("image/image", MergeWithShared(c, fiber.Map{
		"Title":  "Image",
		"Active": "Image",
	}), "layouts/main")
}

// Upload 处理图片上传
func Upload(c *fiber.Ctx) error {
	// 获取上传的文件
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file uploaded",
		})
	}

	// 检查文件大小（5MB）
	if file.Size > 5*1024*1024 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File too large",
		})
	}

	// 检查文件类型
	contentType := file.Header.Get("Content-Type")
	allowedTypes := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/gif":  true,
	}
	if !allowedTypes[contentType] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid file type",
		})
	}

	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	// 保存文件
	err = c.SaveFile(file, fmt.Sprintf("./public/uploads/%s", filename))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save file",
		})
	}

	// 返回文件URL
	fileURL := fmt.Sprintf("/public/uploads/%s", filename)
	return c.JSON(fiber.Map{
		"url": fileURL,
	})
}
