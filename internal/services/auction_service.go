package services

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

var _ AuctionService = (*auctionService)(nil)

// TODO: Add a cleanup function to remove expired auctions from the activeAuctions map
// TODO: Concurrency handling for distributed systems
type auctionService struct {
	activeAuctions     sync.Map
	auctionRepo        repositories.AuctionRepository
	userService        UserService
	shareHolderService ShareHolderService
}

func NewAuctionService(auctionRepo repositories.AuctionRepository, userService UserService, shareHolderService ShareHolderService) AuctionService {
	return &auctionService{
		auctionRepo:        auctionRepo,
		userService:        userService,
		shareHolderService: shareHolderService,
		activeAuctions:     sync.Map{},
	}
}

func (s *auctionService) GetAuctionByID(auctionID uuid.UUID) (*models.Auction, error) {
	return s.auctionRepo.GetAuctionByID(auctionID)
}

func (s *auctionService) CreateAuction(marketInstanceID uuid.UUID, marketChannelID uuid.UUID) (*models.Auction, error) {
	auction := &models.Auction{
		AuctionID:        uuid.New(),
		MarketInstanceID: marketInstanceID,
		MarketChannelID:  marketChannelID,
	}
	auction, err := s.auctionRepo.CreateAuction(auction)
	return auction, err
}

func (s *auctionService) IsAuctionExists(auctionID uuid.UUID) bool {
	_, err := s.auctionRepo.GetAuctionByID(auctionID)
	return err == nil
}

// StartAuction starts an auction
func (s *auctionService) StartAuction(auctionID uuid.UUID, duration time.Duration) (*models.Auction, error) {
	auction, err := s.auctionRepo.GetAuctionByID(auctionID)
	if err != nil {
		return nil, err
	}

	auction.Status = "open"
	auction.StartTime = time.Now()
	duration = duration * time.Second
	auction.EndTime = auction.StartTime.Add(duration)

	s.activeAuctions.Store(auction.AuctionID, auction)

	// Start a goroutine to handle auction expiration
	go s.handleAuctionExpiration(auction)

	return auction, nil
}

// handleAuctionExpiration handles the expiration of an auction
func (s *auctionService) handleAuctionExpiration(auction *models.Auction) {
	// Create a ticker that ticks every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Calculate the duration until the auction ends
	duration := time.Until(auction.EndTime)

	// Loop until the auction ends
	for remaining := duration; remaining > 0; remaining = time.Until(auction.EndTime) {
		// Log the remaining time
		minutes := int(remaining.Seconds()) / 60
		seconds := int(remaining.Seconds()) % 60
		log.Printf("%s %s %d:%02d", auction.MarketInstanceID, auction.MarketChannelID, minutes, seconds)

		// Wait for the next tick
		<-ticker.C
	}

	auction.Status = "closed"
	// Remove currency from user
	s.finalizeAuction(auction)
}

// finalizeAuction finishes an auction
func (s *auctionService) finalizeAuction(auction *models.Auction) error {
	// Update the auction status
	err := s.auctionRepo.UpdateAuction(auction)
	if err != nil {
		return fmt.Errorf("failed to update auction: %w", err)
	}

	// Subtract currency from user
	s.userService.SubtractCurrency(auction.HighestBidderID, auction.HighestBid)
	// Create share holder for the user if they don't already have one
	s.shareHolderService.CreateShareHolderForChannel(auction.HighestBidderID, auction.MarketChannelID)
	return nil
}

func (s *auctionService) PlaceBid(instanceID uuid.UUID, auctionID uuid.UUID, userID uuid.UUID, amount float64) error {
	value, exists := s.activeAuctions.Load(auctionID)
	if !exists {
		return errors.New("auction not found")
	}

	auction := value.(*models.Auction)

	if auction.Status == "closed" || time.Now().After(auction.EndTime) {
		return errors.New("auction is closed")
	}

	if amount <= auction.HighestBid {
		return errors.New("bid amount must be greater than the current highest bid")
	}

	// Verify user has enough balance
	if !s.userService.ValidateUserBalance(userID, amount) {
		return fmt.Errorf("user does not have enough balance")
	}

	auction.HighestBid = amount
	auction.HighestBidderID = userID

	s.activeAuctions.Store(auction.AuctionID, auction)
	s.auctionRepo.UpdateAuction(auction)

	return nil
}
