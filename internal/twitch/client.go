package twitch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hypebid/hypebid-app/internal/config"
)

const (
	baseURL    = "https://api.twitch.tv/helix"
	authURL    = "https://id.twitch.tv/oauth2"
	apiTimeout = 10 * time.Second
)

type client struct {
	cfg *config.Config
}

func NewClient(cfg *config.Config) *client {
	return &client{cfg: cfg}
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type FollowersResponse struct {
	Total      int        `json:"total"`
	Data       []Follower `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type Follower struct {
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	UserLogin  string `json:"user_login"`
	FollowedAt string `json:"followed_at"`
}

type UserResponse struct {
	Data []TwitchUser `json:"data"`
}

type TwitchUser struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
	CreatedAt       string `json:"created_at"`
}

func (c *client) GetAccessToken() (string, error) {
	fmt.Println("clientId: ", c.cfg.ClientID)
	fmt.Println("clientSecret: ", c.cfg.ClientSecret)
	url := fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&grant_type=client_credentials", c.cfg.ClientID, c.cfg.ClientSecret)
	fmt.Println("url:", url)

	resp, err := http.PostForm(url, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Print the raw response
	fmt.Println("Raw response:", string(body))

	var tokenResp AccessTokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v, body: %s", err, string(body))
	}

	fmt.Printf("Token Response: %+v\n", tokenResp)
	fmt.Println("HTTP Status:", resp.Status)
	return tokenResp.AccessToken, nil
}

func (c *client) GetUserByLogin(accessToken, login string) (string, error) {
	fmt.Println("login: ", login)
	url := fmt.Sprintf("%s/users/?login=%s", apiURL, login)
	fmt.Println("url:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Client-ID", c.cfg.ClientID)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Print the raw response
	fmt.Println("Raw response:", string(body))
	fmt.Println("HTTP Status:", resp.Status)

	var userResp UserResponse
	err = json.Unmarshal(body, &userResp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v, body: %s", err, string(body))
	}

	return userResp.Data[0].ID, nil
}

func (c *client) GetUsersByLogin(accessToken string, logins []string) ([]TwitchUser, error) {
	url := fmt.Sprintf("%s/users?", apiURL)
	for _, login := range logins {
		url += fmt.Sprintf("login=%s&", login)
	}
	url = url[:len(url)-1] // Remove the trailing '&'

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Client-ID", c.cfg.ClientID)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userResp UserResponse
	err = json.Unmarshal(body, &userResp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v, body: %s", err, string(body))
	}

	return userResp.Data, nil
}

func (c *client) GetFollowerCount(accessToken, broadcasterID string) (int, error) {
	// fmt.Println("twitchAPIURL: ", apiURL)
	// fmt.Println("broadcasterID: ", broadcasterID)
	url := fmt.Sprintf("%s/channels/followers?broadcaster_id=%s", apiURL, broadcasterID)
	// fmt.Println("url:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Client-ID", c.cfg.ClientID)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Print the raw response
	fmt.Println("Raw response:", string(body))

	var followersResp FollowersResponse
	err = json.Unmarshal(body, &followersResp)
	if err != nil {
		return 0, fmt.Errorf("error unmarshalling response: %v, body: %s", err, string(body))
	}

	return followersResp.Total, nil
}
