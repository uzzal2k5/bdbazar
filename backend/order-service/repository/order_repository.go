package repository

import (
	"order-service/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByBuyerID(buyerID uint) ([]models.Order, error)
	GetBySellerID(sellerID uint) ([]models.Order, error)
	UpdateStatus(orderID uint, status string) error
	DeleteOrder(orderID string) error  // <-- Ensure this line exists
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepo{db}
}

func (r *orderRepo) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepo) GetByBuyerID(buyerID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("buyer_id = ?", buyerID).Find(&orders).Error
	return orders, err
}

func (r *orderRepo) GetBySellerID(sellerID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("seller_id = ?", sellerID).Find(&orders).Error
	return orders, err
}

func (r *orderRepo) UpdateStatus(orderID uint, status string) error {
	return r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status).Error
}

func (r *orderRepo) DeleteOrder(orderID string) error {
    return r.db.Delete(&models.Order{}, "id = ?", orderID).Error
}