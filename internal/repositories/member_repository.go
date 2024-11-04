package repositories

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ MemberRepository = (*memberRepository)(nil)

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	return &memberRepository{db: db}
}

func (r *memberRepository) CreateMember(member *models.Member) error {
	return r.db.Create(member).Error
}

func (r *memberRepository) UpdateMember(member *models.Member) error {
	return r.db.Save(member).Error
}

func (r *memberRepository) DeleteMember(member *models.Member) error {
	return r.db.Delete(member).Error
}

func (r *memberRepository) GetAllMembersForInstance(instanceID uuid.UUID) ([]models.Member, error) {
	var members []models.Member
	err := r.db.Where("market_instance_id = ?", instanceID).Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) GetMemberByInstanceIDAndUserID(instanceID string, userID string) (*models.Member, error) {
	var member models.Member
	err := r.db.Where("instance_id = ? AND user_id = ?", instanceID, userID).First(&member).Error
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) GetMembersByInstanceIDAndStatus(instanceID string, status string) ([]*models.Member, error) {
	var members []*models.Member
	err := r.db.Where("instance_id = ? AND status = ?", instanceID, status).Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) GetMembersByInstanceIDAndRole(instanceID string, role string) ([]*models.Member, error) {
	var members []*models.Member
	err := r.db.Where("instance_id = ? AND role = ?", instanceID, role).Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}
