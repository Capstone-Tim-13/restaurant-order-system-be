package order

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID     uint `gorm:"PrimaryKey"`
	// UserID uint
	Orders     []OrderItem `gorm:"foreignkey:OrderID"`
	TotalPrice float32
	Status     string `gorm:"type:enum('Sent','Accepted','Prepared','In Packing','Already Sent','Arrived'); default:'Sent'"`
	//User       user.User
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}

type OrderItem struct {
	ID       uint `gorm:"PrimaryKey"`
	OrderID  uint
	MenuID   uint
	Quantity int
	SubTotal float32
	CreateAt time.Time      `gorm:"autoCreateTime"`
	UpdateAt time.Time      `gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `gorm:"index"`
}
