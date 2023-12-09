package conversion

import "capstone/features/order"

func OrderItems(menuID uint, quantity int, subTotal float32) *order.OrderItem {
	return &order.OrderItem{
		MenuID:   menuID,
		Quantity: quantity,
		SubTotal: subTotal,
	}
}

func OrderCreateRequest(req []order.OrderItem, total float32) *order.Order {
	return &order.Order{
		Orders:     req,
		TotalPrice: total,
	}
}
