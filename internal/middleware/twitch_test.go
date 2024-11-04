package middleware

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) UpdateUser(user *models.User) (*models.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

type MockOauthManager struct {
	mock.Mock
}

func (m *MockOauthManager) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	args := m.Called(refreshToken)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*oauth2.Token), args.Error(1)
}

// func TestRequireTwitchAuth(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		setupContext   func() context.Context
// 		setupMocks     func(*MockUserService, *MockOauthManager)
// 		expectedStatus int
// 	}{
// 		{
// 			name: "Valid user with valid token",
// 			setupContext: func() context.Context {
// 				user := &models.User{
// 					UserID:         uuid.New(),
// 					TokenExpiresAt: time.Now().Add(time.Hour),
// 				}
// 				return ContextWithUser(context.Background(), user)
// 			},
// 			setupMocks: func(us *MockUserService, om *MockOauthManager) {
// 				us.On("GetUserByID", mock.Anything).Return(&models.User{
// 					TokenExpiresAt: time.Now().Add(time.Hour),
// 				}, nil)
// 			},
// 			expectedStatus: http.StatusOK,
// 		},
// 		{
// 			name: "No user in context",
// 			setupContext: func() context.Context {
// 				return context.Background()
// 			},
// 			setupMocks:     func(us *MockUserService, om *MockOauthManager) {},
// 			expectedStatus: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			userService := &MockUserService{}
// 			oauthManager := &MockOauthManager{}
// 			tt.setupMocks(userService, oauthManager)

// 			middleware := NewTwitchMiddleware(&config.Config{}, oauthManager, userService)

// 			handler := middleware.RequireTwitchAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				w.WriteHeader(http.StatusOK)
// 			}))

// 			req := httptest.NewRequest(http.MethodGet, "/", nil)
// 			req = req.WithContext(tt.setupContext())
// 			w := httptest.NewRecorder()

// 			handler.ServeHTTP(w, req)

// 			assert.Equal(t, tt.expectedStatus, w.Code)
// 			userService.AssertExpectations(t)
// 			oauthManager.AssertExpectations(t)
// 		})
// 	}
// }
