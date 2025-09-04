package repository

import (
    "shipment-service/models"

    "gorm.io/gorm"
)

type ShipmentRepository interface {
    Create(*models.Shipment) error
    GetByOrderID(uint) (*models.Shipment, error)
    GetBySeller(uint) ([]models.Shipment, error)
    UpdateStatus(uint, string) error
    Delete(uint) error
}

type shipmentRepo struct {
    db *gorm.DB
}

func NewShipmentRepository(db *gorm.DB) ShipmentRepository {
    return &shipmentRepo{db}
}

func (r *shipmentRepo) Create(s *models.Shipment) error {
    return r.db.Create(s).Error
}

func (r *shipmentRepo) GetByOrderID(orderID uint) (*models.Shipment, error) {
    var shipment models.Shipment
    err := r.db.Where("order_id = ?", orderID).First(&shipment).Error
    return &shipment, err
}

func (r *shipmentRepo) GetBySeller(sellerID uint) ([]models.Shipment, error) {
    var shipments []models.Shipment
    err := r.db.Where("seller_id = ?", sellerID).Find(&shipments).Error
    return shipments, err
}

func (r *shipmentRepo) UpdateStatus(id uint, status string) error {
    return r.db.Model(&models.Shipment{}).Where("id = ?", id).Update("status", status).Error
}

func (r *shipmentRepo) Delete(id uint) error {
    return r.db.Delete(&models.Shipment{}, id).Error
}
