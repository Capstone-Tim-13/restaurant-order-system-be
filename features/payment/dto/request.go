package dto

type CreatePayment struct {
	OrderID       uint	`json:"order_ID" form:"order_ID"`
}