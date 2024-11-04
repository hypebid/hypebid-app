package services

import (
	"github.com/google/uuid"
	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"
)

var _ MetricService = (*metricService)(nil)

type metricService struct {
	metricRepo     repositories.MetricRepository
	channelService ChannelService
}

func NewMetricService(metricRepo repositories.MetricRepository, channelService ChannelService) *metricService {
	return &metricService{metricRepo: metricRepo, channelService: channelService}
}

func (s *metricService) CreateMetric(metric *models.Metric) error {
	return s.metricRepo.CreateMetric(metric)
}

func (s *metricService) GetMetricsForChannel(channelID uuid.UUID) ([]models.Metric, error) {
	return s.metricRepo.GetMetricsForChannel(channelID)
}

func (s *metricService) GetMetricByNameForChannel(channelName string, metricName string) (*models.Metric, error) {
	channel, err := s.channelService.GetChannelByName(channelName)
	if err != nil {
		return nil, err
	}
	return s.metricRepo.GetMetricByNameForChannel(channel.ChannelID, metricName)
}
