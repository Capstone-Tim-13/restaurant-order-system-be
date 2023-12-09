package dto

type OrderItems struct {
	MenuID   uint `json:"menu_id" form:"menu_id"`
	Quantity int  `json:"quantity" form:"quantity"`
}

type CreateOrder struct {
	OrderItems []OrderItems `json:"order_items" form:"order_items"`
}

type ReqUpdateStatus struct {
	Status string `json:"status" form:"status"`
}

type ReqUpdateOrderItem struct {
	ID       uint `json:"id" form:"id"`
	Quantity int  `json:"quantity" form:"quantity"`
}
