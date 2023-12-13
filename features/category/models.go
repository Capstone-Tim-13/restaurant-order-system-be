package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID       uint           `gorm:"PrimaryKey"`
	Name     string         `gorm:"unique;Not Null" `
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}
