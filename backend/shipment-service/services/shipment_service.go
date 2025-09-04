package services

import "shipment-service/models"

type ShipmentService interface {
    Create(*models.Shipment) error
    GetByOrderID(uint) (*models.Shipment, error)
    GetBySeller(uint) ([]models.Shipment, error)
    UpdateStatus(uint, string) error
    Delete(uint) error
}

type shipmentService struct {
    repo ShipmentService
}

func NewShipmentService(r ShipmentService) ShipmentService {
    return &shipmentService{repo: r}
}

func (s *shipmentService) Create(shipment *models.Shipment) error {
    shipment.Status = "pending"
    return s.repo.Create(shipment)
}

func (s *shipmentService) GetByOrderID(orderID uint) (*models.Shipment, error) {
    return s.repo.GetByOrderID(orderID)
}

func (s *shipmentService) GetBySeller(sellerID uint) ([]models.Shipment, error) {
    return s.repo.GetBySeller(sellerID)
}

func (s *shipmentService) UpdateStatus(id uint, status string) error {
    return s.repo.UpdateStatus(id, status)
}

func (s *shipmentService) Delete(id uint) error {
    return s.repo.Delete(id)
}
