package menu

import (
	"capstone/features/category"
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID          uint              `gorm:"PrimaryKey"`
	Image       string            `gorm:"Not Null"`
	Name        string            `gorm:"unique;not null"`
	CategoryID  uint              `gorm:"Not Null"`
	Category    category.Category `gorm:"foreignKey:CategoryID"`
	Description string
	Price       float32
	Status      bool
	CreateAt    time.Time      `gorm:"autoCreateTime"`
	UpdateAt    time.Time      `gorm:"autoUpdateTime"`
	DeleteAt    gorm.DeletedAt `gorm:"index"`
}
