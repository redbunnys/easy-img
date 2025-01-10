package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

// 这里可以定义您的基础模型
type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt int64
	UpdatedAt int64
} 