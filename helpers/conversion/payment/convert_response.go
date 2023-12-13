package conversion

import (
	"capstone/features/payment"
	"capstone/features/payment/dto"
)

func PaymentCreateResponse(res *payment.Payment, token, redirectURL string) *dto.ResCreatePayment {
	return &dto.ResCreatePayment{
		ID:          res.ID,
		OrderID:     res.OrderID,
		Token:       token,
		RedirectURL: redirectURL,
	}
}

func PaymentResponse(res *payment.Payment) *dto.ResPayment {
	return &dto.ResPayment{
		ID:            res.ID,
		OrderID:       res.OrderID,
		PaymentStatus: res.PaymentStatus,
		PaymentMethod: res.PaymentMethod,
		UpdatedAt:     res.UpdateAt,
	}
}
