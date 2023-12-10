package conversion

import (
	"capstone/features/payment"
	"capstone/features/payment/dto"
)

func paymentCreateRequest(req dto.CreatePayment) *payment.Payment {
	return &payment.Payment{
		OrderID: req.OrderID,
	}
}