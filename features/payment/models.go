package payment

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID            uint `gorm:"PrimaryKey"`
	OrderID       uint
	PaymentStatus string `gorm:"type:enum('pending','cancel','success','challenge','failed'); default:'pending'"`
	PaymentMethod string
	CreateAt      time.Time      `gorm:"autoCreateTime"`
	UpdateAt      time.Time      `gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt `gorm:"index"`
}

type Order struct {
	ID         uint        
	TotalPrice float32
	Status     string        
	UpdateAt   time.Time      
	DeleteAt   gorm.DeletedAt 
}
