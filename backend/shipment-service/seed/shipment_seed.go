package seed

import (
    "time"

    "gorm.io/gorm"
    "shipment-service/models"
)

func SeedShipments(db *gorm.DB) error {
    shipments := []models.Shipment{
        {
            OrderID:      1001,
            SellerID:     201,
            BuyerID:      301,
            Address:      "123 Dhaka Street, Dhaka, Bangladesh",
            Status:       "pending",
            TrackingCode: "BD123456789",
            CreatedAt:    time.Now(),
            UpdatedAt:    time.Now(),
        },
        {
            OrderID:      1002,
            SellerID:     202,
            BuyerID:      302,
            Address:      "456 Chittagong Avenue, Chittagong",
            Status:       "shipped",
            TrackingCode: "BD987654321",
            CreatedAt:    time.Now().Add(-24 * time.Hour),
            UpdatedAt:    time.Now().Add(-12 * time.Hour),
        },
        {
            OrderID:      1003,
            SellerID:     203,
            BuyerID:      303,
            Address:      "789 Sylhet Road, Sylhet",
            Status:       "delivered",
            TrackingCode: "BD112233445",
            CreatedAt:    time.Now().Add(-72 * time.Hour),
            UpdatedAt:    time.Now().Add(-48 * time.Hour),
        },
    }

    for _, shipment := range shipments {
        if err := db.Create(&shipment).Error; err != nil {
            return err
        }
    }

    return nil
}


