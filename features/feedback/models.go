package feedback

import "capstone/features/user"

type Feedback struct {
	ID      uint `gorm:"PrimaryKey"`
	OrderID uint
	UserID  uint
	Rating  int
	Review  string    `gorm:"type:text"`
	User    user.User `gorm:"foreignKey:UserID"`
}
