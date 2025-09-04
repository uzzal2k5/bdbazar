package repository

import (
	"errors"
	"product-service/models"

	"gorm.io/gorm"
)

// ProductRepository defines the contract for product data access
type ProductRepository interface {
	Create(product *models.Product) error
	GetAll(role string, sellerID uint, offset, limit int) ([]models.Product, error)
	GetByID(id uint, role string, sellerID uint) (*models.Product, error)
	SearchByName(name string) ([]models.Product, error)
	FilterProducts(filters map[string]interface{}, offset, limit int) ([]models.Product, error)
	Update(product *models.Product, role string, sellerID uint) error
	Delete(id uint, role string, sellerID uint) error
	IncreaseStock(productID uint, amount int) error
	DecreaseStock(productID uint, amount int) error
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new ProductRepository instance
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create inserts a new product into the database
func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// GetAll retrieves products with pagination, filtered by role
func (r *productRepository) GetAll(role string, sellerID uint, offset, limit int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.Model(&models.Product{})

	if role == "seller" {
		query = query.Where("seller_id = ?", sellerID)
	}

	if err := query.
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

// GetByID fetches a product by ID, checking role-based access
func (r *productRepository) GetByID(id uint, role string, sellerID uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	if role == "seller" && product.SellerID != sellerID {
		return nil, errors.New("unauthorized: vendor cannot access this product")
	}
	return &product, nil
}

// SearchByName performs a case-insensitive partial match
func (r *productRepository) SearchByName(name string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.
		Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").
		Find(&products).Error
	return products, err
}

// FilterProducts applies multiple optional filters
func (r *productRepository) FilterProducts(filters map[string]interface{}, offset, limit int) ([]models.Product, error) {
	var products []models.Product
	query := r.db.Model(&models.Product{})

	if name, ok := filters["name"].(string); ok && name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}
	if category, ok := filters["category"].(string); ok && category != "" {
		query = query.Where("category = ?", category)
	}
	if sellerID, ok := filters["seller_id"].(uint); ok && sellerID != 0 {
		query = query.Where("seller_id = ?", sellerID)
	}
	if minPrice, ok := filters["min_price"].(float64); ok {
		query = query.Where("price >= ?", minPrice)
	}
	if maxPrice, ok := filters["max_price"].(float64); ok {
		query = query.Where("price <= ?", maxPrice)
	}

	err := query.
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&products).Error

	return products, err
}

// Update modifies a product, enforcing role access
func (r *productRepository) Update(product *models.Product, role string, sellerID uint) error {
	if role == "seller" && product.SellerID != sellerID {
		return errors.New("unauthorized: vendor cannot update this product")
	}
	return r.db.Save(product).Error
}

// Delete removes a product, enforcing role access
func (r *productRepository) Delete(id uint, role string, sellerID uint) error {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		return err
	}
	if role == "seller" && product.SellerID != sellerID {
		return errors.New("unauthorized: vendor cannot delete this product")
	}
	return r.db.Delete(&product).Error
}

// IncreaseStock adds to the product quantity
func (r *productRepository) IncreaseStock(productID uint, amount int) error {
	return r.db.Model(&models.Product{}).
		Where("id = ?", productID).
		Update("quantity", gorm.Expr("quantity + ?", amount)).Error
}

// DecreaseStock subtracts from the product quantity if enough is available
func (r *productRepository) DecreaseStock(productID uint, amount int) error {
	return r.db.Model(&models.Product{}).
		Where("id = ? AND quantity >= ?", productID, amount).
		Update("quantity", gorm.Expr("quantity - ?", amount)).Error
}
