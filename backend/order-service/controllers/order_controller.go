package controllers

import (
	"net/http"
	"order-service/models"
	"order-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	Service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{Service: service}
}

// POST /api/orders
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract Buyer ID from JWT context
	buyerID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: no user_id"})
		return
	}
	order.BuyerID = buyerID.(uint)
	order.Status = "pending"

	if err := c.Service.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Order placed successfully",
		"order":   order,
	})
}

// GET /api/orders/buyer
func (c *OrderController) GetBuyerOrders(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	orders, err := c.Service.GetOrdersByBuyer(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch buyer orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// GET /api/orders/seller
func (c *OrderController) GetSellerOrders(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	orders, err := c.Service.GetOrdersBySeller(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch seller orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

// PUT /api/orders/:id/ship
func (c *OrderController) MarkOrderShipped(ctx *gin.Context) {
	orderIDStr := ctx.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Ensure only sellers can mark orders as shipped (role check optional)
	role := ctx.MustGet("role").(string)
	if role != "seller" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can mark orders as shipped"})
		return
	}

	if err := c.Service.MarkAsShipped(uint(orderID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to mark order as shipped"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order marked as shipped"})
}


func (oc *OrderController) DeleteOrder(c *gin.Context) {
    id := c.Param("id")
    err := oc.Service.DeleteOrder(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
