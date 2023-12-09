package repository

import (
	"capstone/features/order"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) order.Repository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Save(newOrder *order.Order) (*order.Order, error) {
	result := r.db.Create(newOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	return newOrder, nil
}

func (r *OrderRepositoryImpl) FindAll() ([]order.Order, error) {
	var orders []order.Order

	result := r.db.Preload("Orders").Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (r *OrderRepositoryImpl) FindById(id int) (*order.Order, error) {
	var order order.Order

	result := r.db.Preload("Orders").Where("id = ?", id).First(&order)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (r *OrderRepositoryImpl) Delete(id int) error {
	result := r.db.Table("orders").Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepositoryImpl) FindMenu(menuID []int) (bool, []float32) {
	var price []float32

	result := r.db.Select("price").Table("menus").Where("id IN ?", menuID).Pluck("price", &price)
	if result.Error != nil {
		return false, nil
	}

	if len(price) == len(menuID) {
		return true, price
	}

	return false, nil
}
