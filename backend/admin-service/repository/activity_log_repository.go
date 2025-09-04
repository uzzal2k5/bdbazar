package repository

import (
	"admin-service/models"
	"gorm.io/gorm"
)

type ActivityLogRepository interface {
	Create(log *models.ActivityLog) error
	GetAll() ([]models.ActivityLog, error)
	GetByID(id uint) (*models.ActivityLog, error)
}

type activityLogRepository struct {
	db *gorm.DB
}

func NewActivityLogRepository(db *gorm.DB) ActivityLogRepository {
	return &activityLogRepository{db: db}
}

func (r *activityLogRepository) Create(log *models.ActivityLog) error {
	return r.db.Create(log).Error
}

func (r *activityLogRepository) GetAll() ([]models.ActivityLog, error) {
	var logs []models.ActivityLog
	err := r.db.Order("created_at desc").Find(&logs).Error
	return logs, err
}

func (r *activityLogRepository) GetByID(id uint) (*models.ActivityLog, error) {
	var log models.ActivityLog
	err := r.db.First(&log, id).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}
