package services

import (
	"time"

	"github.com/hypebid/hypebid-app/internal/twitch"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type AuctionService interface {
	GetAuctionByID(auctionID uuid.UUID) (*models.Auction, error)
	CreateAuction(marketInstanceID uuid.UUID, marketChannelID uuid.UUID) (*models.Auction, error)
	IsAuctionExists(auctionID uuid.UUID) bool
	StartAuction(auctionID uuid.UUID, duration time.Duration) (*models.Auction, error)
	PlaceBid(instanceID uuid.UUID, auctionID uuid.UUID, userID uuid.UUID, amount float64) error
}

type ChannelService interface {
	CreateChannel(name string, sharesTotal int) error
	GetAllChannels() ([]models.Channel, error)
	GetChannelByName(name string) (*models.Channel, error)
	GetChannelByID(channelID uuid.UUID) (*models.Channel, error)
	IsChannelNameTaken(name string) bool
	IsChannelExists(channelID uuid.UUID) bool
}

type MarketChannelService interface {
	InitializeMarketChannel(instanceId uuid.UUID, channelId uuid.UUID, totalShares int, sharePrice float64) (*models.MarketChannel, error)
	GetMarketChannelByID(marketChannelID uuid.UUID) (*models.MarketChannel, error)
	GetMarketChannelsByInstanceID(instanceId uuid.UUID) ([]models.MarketChannel, error)
}

type MarketInstanceService interface {
	CreateMarketInstance(name string, durationDays int, email string, userID uuid.UUID) (*models.MarketInstance, error)
	GetMarketInstanceByID(marketInstanceID uuid.UUID) (*models.MarketInstance, error)
	GetAllMarketInstances() ([]models.MarketInstance, error)
	StartMarketInstance(instanceID uuid.UUID) (*models.MarketInstance, error)
	GetAllMarketInstancesByUserID(userID uuid.UUID) ([]models.MarketInstance, error)
}

type MemberService interface {
	CreateMember(marketInstanceID uuid.UUID, userID uuid.UUID) error
	GetAllMembersForInstance(marketInstanceID uuid.UUID) ([]models.Member, error)
}

type MetricService interface {
	CreateMetric(metric *models.Metric) error
	GetMetricsForChannel(channelID uuid.UUID) ([]models.Metric, error)
	GetMetricByNameForChannel(channelName string, metricName string) (*models.Metric, error)
}

type MetricDataPointService interface {
	CreateMetricDataPoint(metricDataPoint *models.MetricDataPoint) error
	GetRecentFollowerCount(channelName string) ([]models.MetricDataPoint, error)
	GetFollowerStats(channelName string, days int) (*FollowerStats, error)
}

type ShareHolderService interface {
	GetShareHolderByUserIDAndMarketChannelID(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error)
	InitializeShareHolder(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error)
	UpdateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error)
	CreateShareHolderForChannel(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error)
	TransferShares(marketChannelID uuid.UUID, fromUserID uuid.UUID, toUserID uuid.UUID, shareCount int) error
}

type TwitchService interface {
	GetAccessToken() (string, error)
	GetUserByLogin(accessToken, login string) (string, error)
	GetUsersByLogin(accessToken string, logins []string) ([]twitch.TwitchUser, error)
	GetFollowerCount(accessToken, broadcasterID string) (int, error)
}

type UserService interface {
	CreateUser(username, email, password string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uuid.UUID) (*models.User, error)
	FindOrCreateTwitchUser(twitchData *models.TwitchUser, token *oauth2.Token) (*models.User, error)
	AddCurrency(userID uuid.UUID, amount float64) (*models.User, error)
	SubtractCurrency(userID uuid.UUID, amount float64) (*models.User, error)
	GetUserCurrency(userID uuid.UUID) (float64, error)
	SetUserCurrency(userID uuid.UUID, amount float64) (*models.User, error)
	ValidateUserBalance(userID uuid.UUID, amount float64) bool
	ValidateUserForBid(userID uuid.UUID, amount float64) error
	UpdateUser(user *models.User) (*models.User, error)
}
