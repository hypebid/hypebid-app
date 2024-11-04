package repositories

import (
	"github.com/google/uuid"
	"github.com/hypebid/hypebid-app/pkg/models"
	"gorm.io/gorm"
)

type metricRepository struct {
	db *gorm.DB
}

func NewMetricRepository(db *gorm.DB) *metricRepository {
	return &metricRepository{db: db}
}

func (r *metricRepository) CreateMetric(metric *models.Metric) error {
	return r.db.Create(metric).Error
}

func (r *metricRepository) GetMetricsForChannel(channelID uuid.UUID) ([]models.Metric, error) {
	var metrics []models.Metric
	err := r.db.Where("channel_id = ?", channelID).Find(&metrics).Error
	if err != nil {
		return nil, err
	}
	return metrics, nil
}

func (r *metricRepository) GetMetricByNameForChannel(channelID uuid.UUID, metricName string) (*models.Metric, error) {
	var metric models.Metric
	err := r.db.Where("name = ? AND channel_id = ?", metricName, channelID).First(&metric).Error
	if err != nil {
		return nil, err
	}
	return &metric, nil
}
