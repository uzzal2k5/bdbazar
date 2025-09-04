package controllers

import (
	"net/http"
	"shipment-service/models"
	"shipment-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShipmentController struct {
	svc services.ShipmentService
}

func NewShipmentController(svc services.ShipmentService) *ShipmentController {
	return &ShipmentController{svc: svc}
}

// Seller creates shipment info for an order
func (sc *ShipmentController) CreateShipment(c *gin.Context) {
	var shipment models.Shipment
	if err := c.ShouldBindJSON(&shipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)
	role := c.MustGet("role").(string)
	if role != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can create shipments"})
		return
	}

	shipment.SellerID = userID
	if err := sc.svc.Create(&shipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, shipment)
}

// Seller views all their shipments
func (sc *ShipmentController) GetShipmentsBySeller(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	shipments, err := sc.svc.GetBySeller(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipments)
}

// Get shipment by order ID (accessible to buyer or seller if authorized)
func (sc *ShipmentController) GetShipmentByOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order_id"})
		return
	}

	shipment, err := sc.svc.GetByOrderID(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shipment not found"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	role := c.MustGet("role").(string)

	if (role == "seller" && shipment.SellerID != userID) || (role == "buyer" && shipment.BuyerID != userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	c.JSON(http.StatusOK, shipment)
}

// Seller updates shipment status
func (sc *ShipmentController) UpdateStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shipment id"})
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := c.MustGet("role").(string)
	if role != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can update shipment"})
		return
	}

	if err := sc.svc.UpdateStatus(uint(id), req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}

// Seller deletes shipment
func (sc *ShipmentController) DeleteShipment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shipment id"})
		return
	}

	role := c.MustGet("role").(string)
	if role != "seller" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only sellers can delete shipments"})
		return
	}

	if err := sc.svc.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shipment deleted"})
}
