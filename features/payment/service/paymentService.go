package service

import (
	"capstone/features/payment"
	"capstone/features/payment/dto"
	"capstone/helpers"
	conversion "capstone/helpers/conversion/payment"
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type PaymentServiceImpl struct {
	paymentRepository payment.Repository
	Validate          *validator.Validate
}

func NewPaymentService(paymentRepository payment.Repository, validation *validator.Validate) payment.Service {
	return &PaymentServiceImpl{paymentRepository: paymentRepository, Validate: validation}
}

func (s *PaymentServiceImpl) Create(ctx echo.Context, newPayment dto.CreatePayment) (*dto.ResCreatePayment, error) {
	err := s.Validate.Struct(newPayment)
	if err != nil {
		fmt.Println("validation error", err.Error())
		return nil, helpers.ValidationError(ctx, err)
	}

	var payment payment.Payment
	payment.ID = uint(helpers.GenerateRandomID())
	payment.OrderID = newPayment.OrderID

	result, err := s.paymentRepository.Save(&payment)
	if err != nil {
		return nil, err
	}

	order, err := s.paymentRepository.FindOrderById(int(result.OrderID))
	if err != nil {
		return nil, err
	}

	paymentId := strconv.Itoa(int(result.ID))
	token, url := s.paymentRepository.SnapRequest(paymentId, int64(order.TotalPrice))
	if token == "" || url == "" {
		return nil, fmt.Errorf("create transaction error")
	}

	resPayment := conversion.PaymentCreateResponse(result, token, url)

	return resPayment, nil
}

func (s *PaymentServiceImpl) FindAll(ctx echo.Context) ([]dto.ResPayment, error) {
	result, err := s.paymentRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var resPayments []dto.ResPayment

	for _, data := range result {
		resPayment := conversion.PaymentResponse(&data)
		resPayments = append(resPayments, *resPayment)
	}

	return resPayments, nil
}

func (s *PaymentServiceImpl) FindById(id int) (*dto.ResPayment, error) {
	result, err := s.paymentRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	var resPayment = conversion.PaymentResponse(result)

	return resPayment, nil
}

func (s *PaymentServiceImpl) Delete(id int) error {
	result, _ := s.paymentRepository.FindById(id)
	if result == nil {
		return fmt.Errorf("payment not found")
	}

	err := s.paymentRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when delteing : %s", err)
	}

	return nil
}

func (s *PaymentServiceImpl) Notification(notificationPayload map[string]any) error {
	paymentId, exist := notificationPayload["order_id"].(string)
	if !exist {
		return fmt.Errorf("invalid notification payload")
	}

	paymentType, exist := notificationPayload["payment_type"].(string)
	if !exist {
		return fmt.Errorf("invalid notification payload")
	}

	status, err := s.paymentRepository.CheckTransaction(paymentId)
	if err != nil {
		return err
	}

	paymentIdInt, _ := strconv.Atoi(paymentId)
	payment, err := s.paymentRepository.FindById(paymentIdInt)
	if err != nil {
		return err
	}

	payment.PaymentStatus = status
	payment.PaymentMethod = paymentType
	_, err = s.paymentRepository.Update(payment)
	if err != nil {
		return err
	}

	return nil
}
