package service

import (
	"capstone/features/order"
	"capstone/features/order/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/order"
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type OrderServiceImpl struct {
	orderRepository order.Repository
	Validate        *validator.Validate
}

func NewOrderService(orderRepository order.Repository, validate *validator.Validate) order.Service {
	return &OrderServiceImpl{orderRepository: orderRepository, Validate: validate}
}

func (s *OrderServiceImpl) Create(ctx echo.Context, req dto.CreateOrder) (*dto.ResOrder, error) {
	err := s.Validate.Struct(req)
	if err != nil {
		fmt.Println("validation error", err.Error())
		return nil, helpers.ValidationError(ctx, err)
	}

	var menuID []int
	for _, order := range req.OrderItems {
		menuID = append(menuID, int(order.MenuID))
	}

	findMenu, prices := s.orderRepository.FindMenu(menuID)
	if !findMenu {
		return nil, fmt.Errorf("menu not found")
	}

	var subTotal float32
	var totalPrice float32
	var orderItems []order.OrderItem

	for i, item := range req.OrderItems {
		subTotal = prices[i] * float32(item.Quantity)
		orderItem := conversion.OrderItems(item.MenuID, item.Quantity, subTotal)
		orderItems = append(orderItems, *orderItem)
		totalPrice += subTotal
	}

	var NewData = conversion.OrderCreateRequest(orderItems, totalPrice)
	result, err := s.orderRepository.Save(NewData)
	if err != nil {
		return nil, err
	}

	var resOrderItems []dto.ResOrderItems
	for _, item := range result.Orders {
		orderItem := conversion.OrderItemsResponse(item)
		resOrderItems = append(resOrderItems, orderItem)
	}

	var resOrder = conversion.OrderResponse(*result, resOrderItems)

	return &resOrder, nil
}

func (s *OrderServiceImpl) FindAll(ctx echo.Context) ([]dto.ResOrder, error) {
	result, err := s.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var resOrders []dto.ResOrder

	for _, data := range result {
		var resOrderItems []dto.ResOrderItems
		for _, item := range data.Orders {
			orderItem := conversion.OrderItemsResponse(item)
			resOrderItems = append(resOrderItems, orderItem)
		}

		var resOrder = conversion.OrderResponse(data, resOrderItems)
		resOrders = append(resOrders, resOrder)
	}

	return resOrders, nil
}

func (s *OrderServiceImpl) FindById(id int) (*dto.ResOrder, error) {
	result, err := s.orderRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	var resOrderItems []dto.ResOrderItems
	for _, item := range result.Orders {
		orderItem := conversion.OrderItemsResponse(item)
		resOrderItems = append(resOrderItems, orderItem)
	}

	var resOrder = conversion.OrderResponse(*result, resOrderItems)

	return &resOrder, nil
}

func (s *OrderServiceImpl) Delete(id int) error {
	result, _ := s.orderRepository.FindById(id)
	if result == nil {
		return fmt.Errorf("order not found")
	}

	err := s.orderRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when delteing : %s", err)
	}

	return nil
}

func (s *OrderServiceImpl) UpdateStatus(id int, Status string) error {
	result, _ := s.orderRepository.FindById(id)
	if result == nil {
		return fmt.Errorf("order not found")
	}

	result.Status = Status
	_, err := s.orderRepository.Update(result)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderServiceImpl) UpdateOrderItems(updateOrderItem dto.ReqUpdateOrderItem) (*dto.ResOrder, error) {
	orderItem, err := s.orderRepository.FindOrderItemById(int(updateOrderItem.ID))
	if err != nil {
		return nil, err
	}

	newSubtotal := (orderItem.SubTotal / float32(orderItem.Quantity)) * float32(updateOrderItem.Quantity)
	orderItem.Quantity = updateOrderItem.Quantity
	orderItem.SubTotal = newSubtotal

	err = s.orderRepository.UpdateOrderItem(orderItem)
	if err != nil {
		return nil, err
	}

	totalPrice, err := s.orderRepository.CalculateTotalPrice(int(orderItem.OrderID))
	if err != nil || totalPrice == 0 {
		return nil, err
	}

	order, err := s.orderRepository.FindById(int(orderItem.OrderID))
	if err != nil {
		return nil, err
	}

	order.TotalPrice = totalPrice
	result, err := s.orderRepository.Update(order)
	if err != nil {
		return nil, err
	}

	var resOrderItems []dto.ResOrderItems
	for _, item := range result.Orders {
		orderItem := conversion.OrderItemsResponse(item)
		resOrderItems = append(resOrderItems, orderItem)
	}

	var resOrder = conversion.OrderResponse(*result, resOrderItems)

	return &resOrder, nil
}
