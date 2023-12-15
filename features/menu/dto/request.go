package dto

type ReqMenuCreate struct {
	Image       string  `json:"image" form:"image"`
	Name        string  `json:"name" form:"name" validate:"required,min=1"`
	CategoryID  uint    `json:"categoryid" form:"categoryid" validate:"required"`
	Description string  `json:"description" form:"description" validate:"required,min=1"`
	Price       float32 `json:"price" form:"price" validate:"required"`
}

type ReqMenuUpdate struct {
	Image       string  `json:"image" form:"image"`
	Name        string  `json:"name" form:"name" validate:"required,min=1"`
	CategoryID  uint    `json:"categoryid" form:"categoryid" validate:"required"`
	Description string  `json:"description" form:"description" validate:"required,min=1"`
	Price       float32 `json:"price" form:"price" validate:"required"`
}

type ReqUpdateStatus struct {
	Status string `json:"status" form:"status"`
}
