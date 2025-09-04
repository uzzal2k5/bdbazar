package services

import (
	"shop-service/models"
	"shop-service/repository"
)

type ShopDashboard struct {
	TotalShops       int64 `json:"total_shops"`
	ApprovedShops    int64 `json:"approved_shops"`
	BlockedShops     int64 `json:"blocked_shops"`
	RecentShopsCount int64 `json:"recent_shops_count"` // shops created in last 7 days
}

type ShopService interface {
	CreateShop(shop *models.Shop) error
	GetShopByID(id uint) (*models.Shop, error)
	UpdateShop(shop *models.Shop) error
	DeleteShop(id uint) error
	ListShops() ([]models.Shop, error)
	SearchShops(name string) ([]models.Shop, error)
	GetShopDashboard(ownerID uint) (*ShopDashboard, error)
}

type shopService struct {
	repo repository.ShopRepository
}

func NewShopService(repo repository.ShopRepository) ShopService {
	return &shopService{repo}
}

func (s *shopService) CreateShop(shop *models.Shop) error {
	return s.repo.Create(shop)
}

func (s *shopService) GetShopByID(id uint) (*models.Shop, error) {
	return s.repo.GetByID(id)
}

func (s *shopService) UpdateShop(shop *models.Shop) error {
	return s.repo.Update(shop)
}

func (s *shopService) DeleteShop(id uint) error {
	return s.repo.Delete(id)
}

func (s *shopService) ListShops() ([]models.Shop, error) {
	return s.repo.ListAll()
}

func (s *shopService) SearchShops(name string) ([]models.Shop, error) {
	return s.repo.SearchByName(name)
}

func (s *shopService) GetShopDashboard(ownerID uint) (*ShopDashboard, error) {
	total, err := s.repo.CountByOwner(ownerID)
	if err != nil {
		return nil, err
	}

	approved, err := s.repo.CountByOwnerAndApproved(ownerID, true)
	if err != nil {
		return nil, err
	}

	blocked, err := s.repo.CountByOwnerAndBlocked(ownerID, true)
	if err != nil {
		return nil, err
	}

	recent, err := s.repo.CountRecentByOwner(ownerID, 7)
	if err != nil {
		return nil, err
	}

	return &ShopDashboard{
		TotalShops:       total,
		ApprovedShops:    approved,
		BlockedShops:     blocked,
		RecentShopsCount: recent,
	}, nil
}
