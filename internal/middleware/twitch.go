package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/hypebid/hypebid-app/internal/auth"
	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/services"
	"github.com/hypebid/hypebid-app/pkg/models"
)

var (
	_ TokenRefresher = (auth.OAuthManager)(nil)
	_ UserProvider   = (services.UserService)(nil)
)

type TwitchMiddleware struct {
	cfg          *config.Config
	oauthManager TokenRefresher
	userService  UserProvider
}

func NewTwitchMiddleware(cfg *config.Config, oauthManager auth.OAuthManager, userService UserProvider) *TwitchMiddleware {
	return &TwitchMiddleware{
		cfg:          cfg,
		oauthManager: oauthManager,
		userService:  userService,
	}
}

// Helper function to create context with user
func ContextWithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, "user", user)
}

func GetUserFromContext(ctx context.Context) *models.User {
	user, ok := ctx.Value("user").(*models.User)
	if !ok {
		return nil
	}
	return user
}

func (m *TwitchMiddleware) RequireTwitchAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userContext := GetUserFromContext(r.Context())
		if userContext == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := m.userService.GetUserByID(userContext.UserID)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if user.TokenExpiresAt.Before(time.Now()) {
			newToken, err := m.oauthManager.RefreshToken(r.Context(), user.RefreshToken)
			if err != nil {
				http.Error(w, "Error refreshing token", http.StatusInternalServerError)
				return
			}

			user.AccessToken = newToken.AccessToken
			user.RefreshToken = newToken.RefreshToken
			user.TokenExpiresAt = newToken.Expiry

			user, err = m.userService.UpdateUser(user)
			if err != nil {
				http.Error(w, "Error updating user", http.StatusInternalServerError)
				return
			}
		}

		ctx := ContextWithUser(r.Context(), user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
