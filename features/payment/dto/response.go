package dto

import "time"

type ResCreatePayment struct {
	ID          uint   `json:"id"`
	OrderID     uint   `json:"order_id"`
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type ResPayment struct {
	ID            uint      `json:"id"`
	OrderID       uint      `json:"order_id"`
	PaymentStatus string    `json:"payment_status"`
	PaymentMethod string    `json:"payment_method"`
	UpdatedAt     time.Time `json:"payment_date"`
}
