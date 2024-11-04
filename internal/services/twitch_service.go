package services

import (
	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/twitch"
)

var _ TwitchService = (*twitchService)(nil)

type twitchService struct {
	client twitch.Client
}

func NewTwitchService(cfg *config.Config) TwitchService {
	return &twitchService{client: twitch.NewClient(cfg)}
}

func (s *twitchService) GetAccessToken() (string, error) {
	return s.client.GetAccessToken()
}

func (s *twitchService) GetUserByLogin(accessToken, login string) (string, error) {
	return s.client.GetUserByLogin(accessToken, login)
}

func (s *twitchService) GetUsersByLogin(accessToken string, logins []string) ([]twitch.TwitchUser, error) {
	return s.client.GetUsersByLogin(accessToken, logins)
}

func (s *twitchService) GetFollowerCount(accessToken, broadcasterID string) (int, error) {
	return s.client.GetFollowerCount(accessToken, broadcasterID)
}
