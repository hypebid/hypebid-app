package twitch

type Client interface {
	GetAccessToken() (string, error)
	GetUserByLogin(accessToken, login string) (string, error)
	GetUsersByLogin(accessToken string, logins []string) ([]TwitchUser, error)
	GetFollowerCount(accessToken, broadcasterID string) (int, error)
}
