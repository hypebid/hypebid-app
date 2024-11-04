package auth

import (
	"context"

	"golang.org/x/oauth2"
)

type OAuthManager interface {
	GetAuthURL() (string, error)
	ValidateState(state string) bool
	RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error)
	Exchange(ctx context.Context, code string) (*oauth2.Token, error)
	GetTwitchUserData(accessToken string) (*TwitchUser, error)
}
