package dto

type ResOrderItems struct {
	ID       uint    `json:"id"`
	OrderID  uint    `json:"orderId"`
	MenuID   uint    `json:"menuId"`
	Quantity int     `json:"quantity"`
	SubTotal float32 `json:"subTotal"`
}

type ResOrder struct {
	ID         uint            `json:"id"`
	Orders     []ResOrderItems `json:"order_items"`
	TotalPrice float32         `json:"totalPrice"`
}
