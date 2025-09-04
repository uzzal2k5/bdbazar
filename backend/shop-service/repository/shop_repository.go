package repository

import (
	"shop-service/models"
	"gorm.io/gorm"
	"time"
)

type ShopRepository interface {
	Create(shop *models.Shop) error
	GetByID(id uint) (*models.Shop, error)
	Update(shop *models.Shop) error
	Delete(id uint) error
	ListAll() ([]models.Shop, error)
	SearchByName(name string) ([]models.Shop, error)

	CountByOwner(ownerID uint) (int64, error)
	CountByOwnerAndApproved(ownerID uint, approved bool) (int64, error)
	CountByOwnerAndBlocked(ownerID uint, blocked bool) (int64, error)
	CountRecentByOwner(ownerID uint, days int) (int64, error)
}

type shopRepo struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepo{db}
}

func (r *shopRepo) Create(shop *models.Shop) error {
	return r.db.Create(shop).Error
}

func (r *shopRepo) GetByID(id uint) (*models.Shop, error) {
	var shop models.Shop
	if err := r.db.First(&shop, id).Error; err != nil {
		return nil, err
	}
	return &shop, nil
}

func (r *shopRepo) Update(shop *models.Shop) error {
	return r.db.Save(shop).Error
}

func (r *shopRepo) Delete(id uint) error {
	return r.db.Delete(&models.Shop{}, id).Error
}

func (r *shopRepo) ListAll() ([]models.Shop, error) {
	var shops []models.Shop
	err := r.db.Where("is_approved = ? AND is_blocked = ?", true, false).Find(&shops).Error
	return shops, err
}

func (r *shopRepo) SearchByName(name string) ([]models.Shop, error) {
	var shops []models.Shop
	err := r.db.Where("name ILIKE ? AND is_approved = ? AND is_blocked = ?", "%"+name+"%", true, false).Find(&shops).Error
	return shops, err
}

func (r *shopRepo) CountByOwner(ownerID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Shop{}).Where("owner_id = ?", ownerID).Count(&count).Error
	return count, err
}

func (r *shopRepo) CountByOwnerAndApproved(ownerID uint, approved bool) (int64, error) {
	var count int64
	err := r.db.Model(&models.Shop{}).
		Where("owner_id = ? AND is_approved = ?", ownerID, approved).
		Count(&count).Error
	return count, err
}

func (r *shopRepo) CountByOwnerAndBlocked(ownerID uint, blocked bool) (int64, error) {
	var count int64
	err := r.db.Model(&models.Shop{}).
		Where("owner_id = ? AND is_blocked = ?", ownerID, blocked).
		Count(&count).Error
	return count, err
}

func (r *shopRepo) CountRecentByOwner(ownerID uint, days int) (int64, error) {
	var count int64
	threshold := time.Now().AddDate(0, 0, -days)
	err := r.db.Model(&models.Shop{}).
		Where("owner_id = ? AND created_at >= ?", ownerID, threshold).
		Count(&count).Error
	return count, err
}
