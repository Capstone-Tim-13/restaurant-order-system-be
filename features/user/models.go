package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"PrimaryKey"`
	Username  string         `gorm:"unique;not null"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	NoHp      string         `gorm:"unique;not null"`
	BirthDate string         `gorm:"type:date;not null"`
	CreateAt  time.Time      `gorm:"autoCreateTime"`
	UpdateAt  time.Time      `gorm:"autoUpdateTime"`
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}
