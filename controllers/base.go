package controllers

import "github.com/gofiber/fiber/v2"

// MergeWithShared 合并特定页面数据与共享数据
func MergeWithShared(c *fiber.Ctx, pageData fiber.Map) fiber.Map {
	// 获取共享数据
	shared := c.Locals("shared").(fiber.Map)

	// 合并数据
	for k, v := range shared {
		pageData[k] = v
	}

	return pageData
}

// BaseController 基础控制器
type BaseController struct {
	ctx    *fiber.Ctx
	page   string
	data   fiber.Map
	layout string
}

// New 创建新的控制器实例
func New(c *fiber.Ctx) *BaseController {
	return &BaseController{
		ctx:    c,
		data:   fiber.Map{},
		layout: "layouts/main", // 默认布局
	}
}

// Page 设置渲染的页面
func (b *BaseController) Page(page string) *BaseController {
	b.page = page
	return b
}

// Layout 设置布局
func (b *BaseController) Layout(layout string) *BaseController {
	b.layout = layout
	return b
}

// Title 设置页面标题
func (b *BaseController) Title(title string) *BaseController {
	b.data["Title"] = title
	return b
}

// Active 设置当前激活的导航项
func (b *BaseController) Active(active string) *BaseController {
	b.data["Active"] = active
	return b
}

// CSS 设置CSS文件
func (b *BaseController) CSS(file string) *BaseController {
	b.data["CssFile"] = file
	return b
}

// JS 设置JS文件
func (b *BaseController) JS(file string) *BaseController {
	b.data["JsFile"] = file
	return b
}

// Data 添加数据
func (b *BaseController) Data(key string, value interface{}) *BaseController {
	b.data[key] = value
	return b
}

// Bulk 批量添加数据
func (b *BaseController) Bulk(data fiber.Map) *BaseController {
	for k, v := range data {
		b.data[k] = v
	}
	return b
}

// Render 执行渲染
func (b *BaseController) Render() error {
	// 获取共享数据
	shared := b.ctx.Locals("shared").(fiber.Map)

	// 合并数据
	for k, v := range shared {
		if _, exists := b.data[k]; !exists {
			b.data[k] = v
		}
	}

	return b.ctx.Render(b.page, b.data, b.layout)
}
