package service

import (
	"capstone/features/order"
	"capstone/features/order/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/Order"
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

func (s *OrderServiceImpl) Create(ctx echo.Context, req dto.CreateOrder) (*order.Order, error) {
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

	// var resOrderItems []dto.OrderItems
	// for i, item := range result.Orders {
	// 	orderItem := conversion.OrderItemsResponse()
	// }
	

	return result, nil
}

func (s *OrderServiceImpl) FindAll(ctx echo.Context) ([]order.Order, error) {
	result, err := s.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *OrderServiceImpl) FindById(id int) (*order.Order, error) {
	result, err := s.orderRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return result, nil
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