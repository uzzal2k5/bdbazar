package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"product-service/models"
	"product-service/services"
)

type ProductController struct {
	Service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{Service: service}
}

// 🔐 Helper to extract user info from context
func getAuthUser(c *gin.Context) (uint, string, error) {
	uidRaw, ok := c.Get("id")
	if !ok {
		return 0, "", errors.New("user ID not found in context")
	}
	uid, ok := uidRaw.(uint)
	if !ok {
		return 0, "", errors.New("invalid user ID type in context")
	}

	roleRaw, ok := c.Get("role")
	if !ok {
		return 0, "", errors.New("role not found in context")
	}
	role, ok := roleRaw.(string)
	if !ok {
		return 0, "", errors.New("invalid role type in context")
	}

	return uid, role, nil
}

// ✅ Create Product (Seller only)
func (productController *ProductController) CreateProduct(contxt *gin.Context) {
	var product models.Product
	if err := contxt.ShouldBindJSON(&product); err != nil {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, role, err := getAuthUser(contxt)
	if err != nil {
		contxt.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if role != "seller" {
		contxt.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can create products"})
		return
	}

	product.SellerID = userID
	if err := productController.Service.CreateProduct(&product, role); err != nil {
		contxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contxt.JSON(http.StatusCreated, gin.H{"message": "Product created", "product": product})
}

// 🔍 Get All Products (Public)
// GetAll supports pagination and RBAC filtering
func (productController *ProductController) GetAll(contxt *gin.Context) {
    userID, role, err := getAuthUser(contxt)
    if err != nil {
        userID = 0
        role = "public"
    }
    // Pagination params
	offset, _ := strconv.Atoi(contxt.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(contxt.DefaultQuery("limit", "10"))

	products, err := productController.Service.GetAll(role, userID, offset, limit)
	if err != nil {
		contxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	contxt.JSON(http.StatusOK, products)
}

// 🔍 Get Product by ID (Public)
func (productController *ProductController) GetByID(contxt *gin.Context) {
    id64, err := strconv.ParseUint(contxt.Param("id"), 10, 32)
    if err != nil {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	id := uint(id64)

    userID, role, err := getAuthUser(contxt)
	if err != nil {
		// For public access fallback
		userID = 0
		role = "public"
	}

	product, err := productController.Service.GetByID(id, role, userID)
	if err != nil {
		contxt.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	contxt.JSON(http.StatusOK, product)
}

// ✏️ Update Product (Seller only, must own product)
func (productController *ProductController) UpdateProduct(contxt *gin.Context) {
    id64, err := strconv.ParseUint(contxt.Param("id"), 10, 32)
	if err != nil {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	id := uint(id64)

	var product models.Product
	if err := contxt.ShouldBindJSON(&product); err != nil {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, role, err := getAuthUser(contxt)
	if err != nil {
		contxt.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

    existing, err := productController.Service.GetByID(id, role, userID)
	if err != nil {
		contxt.JSON(http.StatusForbidden, gin.H{"error": "Cannot access this product"})
		return
	}

    if role == "seller" && existing.SellerID != userID {
		contxt.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own products"})
		return
	}
	product.ID = id
    product.SellerID = existing.SellerID // maintain original sellerID

    if err := productController.Service.UpdateProduct(&product, role, userID); err != nil {
		contxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contxt.JSON(http.StatusOK, gin.H{"message": "Product updated", "product": product})
}

// ❌ Delete Product (Seller only, must own product)
func (productController *ProductController) DeleteProduct(contxt *gin.Context) {
    id64, err := strconv.ParseUint(contxt.Param("id"), 10, 32)
	if err != nil {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	id := uint(id64)


	userID, role, err := getAuthUser(contxt)
	if err != nil {
		contxt.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

    existing, err := productController.Service.GetByID(id, role, userID)
	if err != nil {
		contxt.JSON(http.StatusForbidden, gin.H{"error": "Cannot access this product"})
		return
	}

	if role == "seller" && existing.SellerID != userID {
		contxt.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own products"})
		return
	}

	if err := productController.Service.DeleteProduct(id, role, userID); err != nil {
		contxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	contxt.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}


// 🔧 Adjust Stock (internal use only, no auth)
func (productController *ProductController) AdjustStock(contxt *gin.Context) {
	var payload struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"` // negative = decrease, positive = increase
	}

	if err := contxt.ShouldBindJSON(&payload); err != nil {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err error
	if payload.Quantity < 0 {
		err = productController.Service.DecreaseStock(payload.ProductID, -payload.Quantity)
	} else {
		err = productController.Service.IncreaseStock(payload.ProductID, payload.Quantity)
	}

	if err != nil {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contxt.JSON(http.StatusOK, gin.H{"message": "Stock adjusted successfully"})
}



// 🔍 Search Product by ID or Name (Public)
func (productController *ProductController) SearchProduct(contxt *gin.Context) {
	query := contxt.Query("q")
	if query == "" {
		contxt.JSON(http.StatusBadRequest, gin.H{"error": "Missing query param 'q'"})
		return
	}

	// Extract auth info from context
    // 	userID := contxt.GetUint("userID")
	shopID := contxt.GetUint("shopID")
	role := contxt.GetString("role")

    userID, exists := contxt.Get("userID")
	if exists {
		_ = userID // TODO: Use this if needed for logging/audit/etc.
	}
	// Try to parse query as ID first
	if id, err := strconv.Atoi(query); err == nil {
		product, err := productController.Service.GetByID(uint(id), role, shopID)
		if err == nil {
			contxt.JSON(http.StatusOK, []models.Product{*product})
			return
		}
	}

	// If not found by ID or not numeric, search by name
	products, err := productController.Service.SearchByName(query)
	if err != nil || len(products) == 0 {
		contxt.JSON(http.StatusNotFound, gin.H{"error": "No product found matching query"})
		return
	}

	contxt.JSON(http.StatusOK, products)
}
