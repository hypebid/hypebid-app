package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/hypebid/hypebid-app/internal/config"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

var _ OAuthManager = (*oauthManager)(nil)

type oauthManager struct {
	config *oauth2.Config
	states map[string]time.Time
	mu     sync.RWMutex
}

func NewOAuthManager(cfg *config.Config) *oauthManager {
	fmt.Println("cfg.HostURL: ", cfg.HostURL)
	manager := &oauthManager{
		config: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  fmt.Sprintf("%s/api/v1/auth/twitch/callback", cfg.HostURL), // Use HostURL for production
			Scopes: []string{
				"user:read:email",
			},
			Endpoint: twitch.Endpoint,
		},
		states: make(map[string]time.Time),
	}

	go manager.cleanupStates()

	return manager
}

func (m *oauthManager) generateStates() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	state := hex.EncodeToString(b)

	m.mu.Lock()
	m.states[state] = time.Now().Add(10 * time.Minute)
	m.mu.Unlock()

	return state, nil
}

func (m *oauthManager) GetAuthURL() (string, error) {
	state, err := m.generateStates()
	fmt.Println("state: ", state)
	if err != nil {
		return "", err
	}

	return m.config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("force_verify", "true"), // Forces re-authentication
	), nil
}

func (m *oauthManager) ValidateState(state string) bool {
	m.mu.RLock()
	expiry, exists := m.states[state]
	m.mu.RUnlock()

	fmt.Println("m.states: ", m.states)
	fmt.Println("expiry: ", expiry)
	fmt.Println("exists: ", exists)
	fmt.Println("time.Now().After(expiry): ", time.Now().After(expiry))

	if !exists || time.Now().After(expiry) {
		return false
	}

	m.mu.Lock()
	delete(m.states, state)
	m.mu.Unlock()

	return true
}

func (m *oauthManager) cleanupStates() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		now := time.Now()
		m.mu.Lock()
		for state, expiry := range m.states {
			if now.After(expiry) {
				delete(m.states, state)
			}
		}
		m.mu.Unlock()
	}
}

// RefreshToken handles token refresh when the access token expires
func (m *oauthManager) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}
	newToken, err := m.config.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func (m *oauthManager) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return m.config.Exchange(ctx, code)
}

type TwitchUser struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func (m *oauthManager) GetTwitchUserData(accessToken string) (*TwitchUser, error) {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", m.config.ClientID)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Data []TwitchUser `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("no user data returned")
	}

	return &result.Data[0], nil
}
