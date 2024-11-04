package repositories

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ AuctionRepository = (*auctionRepository)(nil)

type auctionRepository struct {
	db *gorm.DB
}

func NewAuctionRepository(db *gorm.DB) *auctionRepository {
	return &auctionRepository{db: db}
}

func (r *auctionRepository) CreateAuction(auction *models.Auction) (*models.Auction, error) {
	err := r.db.Create(auction).Error
	return auction, err
}

func (r *auctionRepository) GetAuctionByID(auctionID uuid.UUID) (*models.Auction, error) {
	var auction models.Auction
	err := r.db.Where("auction_id = ?", auctionID).First(&auction).Error
	return &auction, err
}

func (r *auctionRepository) UpdateAuction(auction *models.Auction) error {
	return r.db.Save(auction).Error
}
