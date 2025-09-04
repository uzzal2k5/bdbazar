package services

import (
	"errors"
	"product-service/models"
	"product-service/repository"
)

type ProductService interface {
	CreateProduct(product *models.Product, role string) error
	GetAll(role string, sellerID uint, offset int, limit int) ([]models.Product, error)
	GetByID(id uint, role string, sellerID uint) (*models.Product, error)
	UpdateProduct(product *models.Product, role string, sellerID uint) error
	DeleteProduct(id uint, role string, sellerID uint) error

	DecreaseStock(productID uint, quantity int) error
	IncreaseStock(productID uint, quantity int) error
	CheckAvailability(productID uint, quantity int) (bool, error)
	SearchByName(name string) ([]models.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

// RBAC helper function
func isAdmin(role string) bool {
	return role == "admin"
}

func isVendor(role string) bool {
	return role == "seller"
}

// CreateProduct allows only admin to create
func (s *productService) CreateProduct(product *models.Product, role string) error {
	if !isAdmin(role) {
		return errors.New("unauthorized: only admin can create products")
	}
	return s.repo.Create(product)
}

// GetAll accessible to all roles, with pagination
func (s *productService) GetAll(role string, sellerID uint, offset, limit int) ([]models.Product, error) {
	return s.repo.GetAll(role, sellerID, offset, limit)
}

// GetByID accessible to all roles, checks seller ownership for seller
func (s *productService) GetByID(id uint, role string, sellerID uint) (*models.Product, error) {
	return s.repo.GetByID(id, role, sellerID)
}

// UpdateProduct allows only admin or product owner to update
func (s *productService) UpdateProduct(product *models.Product, role string, sellerID uint) error {
	if !isAdmin(role) && product.SellerID != sellerID {
		return errors.New("unauthorized: only admin or owner can update products")
	}
	return s.repo.Update(product, role, sellerID)
}

// DeleteProduct allows only admin or product owner to delete
func (s *productService) DeleteProduct(id uint, role string, sellerID uint) error {
	return s.repo.Delete(id, role, sellerID)
}

// DecreaseStock decreases stock quantity (no role check here)
func (s *productService) DecreaseStock(productID uint, quantity int) error {
	product, err := s.repo.GetByID(productID, "admin", 0) // admin role to bypass ownership check
	if err != nil {
		return err
	}
	if product.Quantity < quantity {
		return errors.New("not enough stock available")
	}
	product.Quantity -= quantity
	return s.repo.Update(product, "admin", 0)
}

// IncreaseStock increases stock quantity (no role check here)
func (s *productService) IncreaseStock(productID uint, quantity int) error {
	product, err := s.repo.GetByID(productID, "admin", 0)
	if err != nil {
		return err
	}
	product.Quantity += quantity
	return s.repo.Update(product, "admin", 0)
}

// CheckAvailability verifies product availability
func (s *productService) CheckAvailability(productID uint, quantity int) (bool, error) {
	product, err := s.repo.GetByID(productID, "admin", 0)
	if err != nil {
		return false, err
	}
	return product.Quantity >= quantity, nil
}

func (s *productService) SearchByName(name string) ([]models.Product, error) {
    return s.repo.SearchByName(name)
}