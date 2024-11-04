package middleware

import (
	"context"

	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// Defining interface for TokenRefresher to allow for mocking in tests
// Only contains methods used by this middleware
type TokenRefresher interface {
	RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error)
}

type UserProvider interface {
	GetUserByID(userID uuid.UUID) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}
