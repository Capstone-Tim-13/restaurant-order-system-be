package repository

import (
	"capstone/features/payment"
	"capstone/helpers"

	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	db            *gorm.DB
	snapClient    snap.Client
	coreAPIClient coreapi.Client
}

func NewPaymentRepository(db *gorm.DB, snapClient snap.Client, coreAPIClient coreapi.Client) payment.Repository {
	return &paymentRepositoryImpl{db: db, snapClient: snapClient, coreAPIClient: coreAPIClient}
}

func (r *paymentRepositoryImpl) Save(payment *payment.Payment) (*payment.Payment, error) {
	result := r.db.Create(payment)
	if result.Error != nil {
		return nil, result.Error
	}

	return payment, nil
}

func (r *paymentRepositoryImpl) FindAll() ([]payment.Payment, error) {
	var payments []payment.Payment

	result := r.db.Find(&payments)
	if result.Error != nil {
		return nil, result.Error
	}

	return payments, nil
}

func (r *paymentRepositoryImpl) FindById(id int) (*payment.Payment, error) {
	var payment payment.Payment

	result := r.db.Where("id = ?", id).First(&payment)
	if result.Error != nil {
		return nil, result.Error
	}

	return &payment, nil
}

func (r *paymentRepositoryImpl) FindOrderById(orderId int) (*payment.Order, error) {
	var order payment.Order

	result := r.db.Where("id = ?", orderId).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (r *paymentRepositoryImpl) Update(updatePayment *payment.Payment) (*payment.Payment, error) {
	result := r.db.Save(updatePayment)
	if result.Error != nil {
		return nil, result.Error
	}
	return updatePayment, nil
}

func (r *paymentRepositoryImpl) Delete(id int) error {
	result := r.db.Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *paymentRepositoryImpl) SnapRequest(paymentId string, total int64) (string, string) {
	snapResponse, err := helpers.CreateSnapRequest(r.snapClient, paymentId, total)
	if err != nil {
		return "", ""
	}

	return snapResponse.Token, snapResponse.RedirectURL
}

func (repository *paymentRepositoryImpl) CheckTransaction(paymentId string) (string, error) {
	var status string

	transactionStatusResp, err := repository.coreAPIClient.CheckTransaction(paymentId)
	if err != nil {
		return "", err
	} else {
		if transactionStatusResp != nil {
			status = helpers.TransactionStatus(transactionStatusResp)
			return status, nil
		}
	}

	return status, err
}
