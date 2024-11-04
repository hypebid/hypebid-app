package services

import (
	"time"

	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"
)

type DailyMetricStats struct {
	Date                 time.Time
	AverageFollowerCount int
}

type FollowerStats struct {
	ChannelName   string
	StartTime     time.Time
	EndTime       time.Time
	DailyMetrics  []repositories.DailyMetric
	MaxFollowers  int
	MinFollowers  int
	OverallGrowth int
}

type metricDataPointService struct {
	metricDataPointRepo repositories.MetricDataPointRepository
	metricService       MetricService
}

func NewMetricDataPointService(metricDataPointRepo repositories.MetricDataPointRepository, metricService MetricService) *metricDataPointService {
	return &metricDataPointService{metricDataPointRepo: metricDataPointRepo, metricService: metricService}
}

func (s *metricDataPointService) CreateMetricDataPoint(metricDataPoint *models.MetricDataPoint) error {
	return s.metricDataPointRepo.CreateMetricDataPoint(metricDataPoint)
}

// func (s *metricDataPointService) GetMetricDataPointsForMetric(metricID uuid.UUID) ([]models.MetricDataPoint, error) {
// 	return s.metricDataPointRepo.GetMetricDataPointsForMetric(metricID)
// }

func (s *metricDataPointService) GetRecentFollowerCount(channelName string) ([]models.MetricDataPoint, error) {
	metric, err := s.metricService.GetMetricByNameForChannel(channelName, "follower_count")
	if err != nil {
		return nil, err
	}

	metricDataPoints, err := s.metricDataPointRepo.GetRecentMetricDataPointsForChannelMetric(metric.MetricID)
	if err != nil {
		return nil, err
	}

	return metricDataPoints, nil
}

// Get the average value of a metric for a channel over a given number of days
func (s *metricDataPointService) GetFollowerStats(channelName string, days int) (*FollowerStats, error) {
	metric, err := s.metricService.GetMetricByNameForChannel(channelName, "follower_count")
	if err != nil {
		return nil, err
	}

	// Calculate time range
	endTime := time.Now().UTC()
	startTime := endTime.AddDate(0, 0, -days)

	// Get daily metrics from repository
	dailyMetrics, err := s.metricDataPointRepo.GetDailyMetricsForPeriod(metric.MetricID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	// Calculate meta information
	var maxFollowers, minFollowers int
	if len(dailyMetrics) > 0 {
		maxFollowers = dailyMetrics[0].AverageFollowerCount
		minFollowers = dailyMetrics[0].AverageFollowerCount
	}

	for _, metric := range dailyMetrics {
		if metric.AverageFollowerCount > maxFollowers {
			maxFollowers = metric.AverageFollowerCount
		}
		if metric.AverageFollowerCount < minFollowers {
			minFollowers = metric.AverageFollowerCount
		}
	}

	stats := &FollowerStats{
		ChannelName:   channelName,
		StartTime:     startTime,
		EndTime:       endTime,
		DailyMetrics:  dailyMetrics,
		MaxFollowers:  maxFollowers,
		MinFollowers:  minFollowers,
		OverallGrowth: maxFollowers - minFollowers,
	}

	return stats, nil
}
