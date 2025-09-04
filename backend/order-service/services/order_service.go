package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"order-service/repository"
	"order-service/models"
)

type OrderService struct {
	Repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	order.Status = "pending"
	return s.Repo.Create(order)
}

func (s *OrderService) GetOrdersByBuyer(buyerID uint) ([]models.Order, error) {
	return s.Repo.GetByBuyerID(buyerID)
}

func (s *OrderService) GetOrdersBySeller(sellerID uint) ([]models.Order, error) {
	return s.Repo.GetBySellerID(sellerID)
}

func (s *OrderService) MarkAsShipped(orderID uint) error {
	return s.Repo.UpdateStatus(orderID, "shipped")
}


func adjustProductStock(productID uint, quantity int) error {
	url := "http://product-service:8082/api/products/adjust-stock"
	payload := map[string]interface{}{
		"product_id": productID,
		"quantity":   quantity,
	}
	jsonValue, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("stock update failed: %s", string(body))
	}
	return nil
}

func (s *OrderService) DeleteOrder(id string) error {
    return s.Repo.DeleteOrder(id)
}
