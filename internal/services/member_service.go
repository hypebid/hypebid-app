package services

import (
	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
)

var _ MemberService = (*memberService)(nil)

type memberService struct {
	memberRepo repositories.MemberRepository
}

func NewMemberService(memberRepo repositories.MemberRepository) *memberService {
	return &memberService{memberRepo: memberRepo}
}

func (s *memberService) CreateMember(marketInstanceID uuid.UUID, userID uuid.UUID) error {
	member := &models.Member{
		MarketInstanceID: marketInstanceID,
		UserID:           userID,
	}

	return s.memberRepo.CreateMember(member)
}

func (s *memberService) GetAllMembersForInstance(marketInstanceID uuid.UUID) ([]models.Member, error) {
	members, err := s.memberRepo.GetAllMembersForInstance(marketInstanceID)
	if err != nil {
		return nil, err
	}
	return members, nil
}
