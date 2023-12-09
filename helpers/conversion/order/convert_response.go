package conversion

import (
	"capstone/features/order"
	"capstone/features/order/dto"
)

func OrderItemsResponse(res order.OrderItem) dto.ResOrderItems {
	return dto.ResOrderItems{
		ID:       res.ID,
		OrderID:  res.OrderID,
		MenuID:   res.MenuID,
		Quantity: res.Quantity,
		SubTotal: res.SubTotal,
	}
}


func OrderResponse(res order.Order, orderItems []dto.ResOrderItems) dto.ResOrder {
	return dto.ResOrder{
		ID:         res.ID,
		Orders:     orderItems,
		TotalPrice: res.TotalPrice,
	}
}