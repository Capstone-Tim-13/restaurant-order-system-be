package feedback

import (
	"capstone/features/user"
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	ID       uint `gorm:"PrimaryKey"`
	UserID   uint
	OrderID  uint
	Rating   int
	Review   string
	User     user.User      `gorm:"foreignKey:UserID"`
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}
