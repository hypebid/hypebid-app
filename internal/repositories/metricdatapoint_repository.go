package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/hypebid/hypebid-app/pkg/models"
	"gorm.io/gorm"
)

var _ MetricDataPointRepository = (*metricDataPointRepository)(nil)

type metricDataPointRepository struct {
	db *gorm.DB
}

func NewMetricDataPointRepository(db *gorm.DB) *metricDataPointRepository {
	return &metricDataPointRepository{db: db}
}

func (r *metricDataPointRepository) CreateMetricDataPoint(metricDataPoint *models.MetricDataPoint) error {
	return r.db.Create(metricDataPoint).Error
}

func (r *metricDataPointRepository) GetMetricDataPointsForMetric(metricID uuid.UUID) ([]models.MetricDataPoint, error) {
	var metricDataPoints []models.MetricDataPoint
	err := r.db.Where("metric_id = ?", metricID).Find(&metricDataPoints).Error
	if err != nil {
		return nil, err
	}
	return metricDataPoints, nil
}

// Last 24 hours
func (r *metricDataPointRepository) GetRecentMetricDataPointsForChannelMetric(metricID uuid.UUID) ([]models.MetricDataPoint, error) {
	now := time.Now()
	oneDayAgo := now.Add(-24 * time.Hour)

	var metricDataPoints []models.MetricDataPoint
	err := r.db.Where("metric_id = ? AND recorded_at BETWEEN ? AND ?", metricID, oneDayAgo, now).Find(&metricDataPoints).Error
	if err != nil {
		return nil, err
	}
	return metricDataPoints, nil
}

type DailyMetric struct {
	Date                 string `json:"date"`
	AverageFollowerCount int    `json:"averageFollowerCount"`
}

// Get the average value of a metric for a channel over a given number of days
func (r *metricDataPointRepository) GetDailyMetricsForPeriod(metricID uuid.UUID, startTime, endTime time.Time) ([]DailyMetric, error) {
	var metrics []DailyMetric

	rows, err := r.db.Raw(`
        SELECT 
            DATE(recorded_at) as date,
            ROUND(AVG(value)) as average_follower_count
        FROM metric_data_points 
        WHERE metric_id = ? 
        AND recorded_at >= ? 
        AND recorded_at <= ?
        GROUP BY DATE(recorded_at)
        ORDER BY DATE(recorded_at)
    `, metricID, startTime, endTime).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var metric DailyMetric
		var date time.Time
		if err := rows.Scan(&date, &metric.AverageFollowerCount); err != nil {
			return nil, err
		}
		metric.Date = date.Format(time.RFC3339)
		metrics = append(metrics, metric)
	}

	return metrics, nil
}
