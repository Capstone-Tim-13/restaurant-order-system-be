package dto

type OrderItems struct {
	MenuID   uint `json:"menu_id" form:"menu_id" validate:"required,min=1"`
	Quantity int  `json:"quantity" form:"quantity" validate:"required,gte=1"`
}

type CreateOrder struct {
	OrderItems []OrderItems `json:"order_items" form:"order_items" validate:"required,min=1"`
}

type ReqUpdateStatus struct {
	Status string `json:"status" form:"status" validate:"required,min=1"`
}

type ReqUpdateOrderItem struct {
	ID       uint `json:"id" form:"id" validate:"required,min=1"`
	Quantity int  `json:"quantity" form:"quantity" validate:"required,gte=1"`
}
