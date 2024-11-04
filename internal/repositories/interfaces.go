package repositories

import (
	"time"

	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

type AuctionRepository interface {
	CreateAuction(auction *models.Auction) (*models.Auction, error)
	GetAuctionByID(auctionID uuid.UUID) (*models.Auction, error)
	UpdateAuction(auction *models.Auction) error
}

type ChannelRepository interface {
	CreateChannel(channel *models.Channel) error
	GetAllChannels() ([]models.Channel, error)
	GetChannelByID(channelID uuid.UUID) (*models.Channel, error)
	GetChannelByName(name string) (*models.Channel, error)
}

type MarketChannelRepository interface {
	CreateMarketChannel(marketChannel *models.MarketChannel) (*models.MarketChannel, error)
	GetMarketChannelByID(marketChannelID uuid.UUID) (*models.MarketChannel, error)
	GetMarketChannelsByInstanceID(instanceId uuid.UUID) ([]models.MarketChannel, error)
}

type MarketInstanceRepository interface {
	CreateMarketInstance(instance *models.MarketInstance) error
	GetMarketInstanceByID(instanceID uuid.UUID) (*models.MarketInstance, error)
	GetMarketInstanceByNameAndUserID(name string, userID uuid.UUID) (*models.MarketInstance, error)
	GetAllMarketInstances() ([]models.MarketInstance, error)
	GetAllMarketInstancesByUserID(userID uuid.UUID) ([]models.MarketInstance, error)
	GetAllActiveMarketInstances() ([]models.MarketInstance, error)
	UpdateMarketInstance(instance *models.MarketInstance) error
	IsMarketInstanceExists(instanceID uuid.UUID) bool
}

type MemberRepository interface {
	CreateMember(member *models.Member) error
	UpdateMember(member *models.Member) error
	DeleteMember(member *models.Member) error
	GetAllMembersForInstance(instanceID uuid.UUID) ([]models.Member, error)
}

type MetricRepository interface {
	CreateMetric(metric *models.Metric) error
	GetMetricsForChannel(channelID uuid.UUID) ([]models.Metric, error)
	GetMetricByNameForChannel(channelID uuid.UUID, metricName string) (*models.Metric, error)
}

type MetricDataPointRepository interface {
	CreateMetricDataPoint(metricDataPoint *models.MetricDataPoint) error
	GetMetricDataPointsForMetric(metricID uuid.UUID) ([]models.MetricDataPoint, error)
	GetRecentMetricDataPointsForChannelMetric(metricID uuid.UUID) ([]models.MetricDataPoint, error)
	GetDailyMetricsForPeriod(metricID uuid.UUID, startTime, endTime time.Time) ([]DailyMetric, error)
}

type ShareHolderRepository interface {
	CreateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error)
	GetShareHolderByUserIDAndMarketChannelID(userID uuid.UUID, marketChannelID uuid.UUID) (*models.ShareHolder, error)
	UpdateShareHolder(shareHolder *models.ShareHolder) (*models.ShareHolder, error)
}

type TwitchRepository interface {
	CreateTwitchUser(twitchUser *models.TwitchUser) error
	GetTwitchUserByID(id string) (*models.TwitchUser, error)
	UpdateTwitchUser(twitchUser *models.TwitchUser) error
}

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	IsUserExists(userUUID uuid.UUID) bool
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uuid.UUID) (*models.User, error)
	GetUserByTwitchID(twitchID string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}
