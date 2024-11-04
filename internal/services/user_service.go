package services

import (
	"errors"
	"fmt"
	"log"

	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/pkg/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

var _ UserService = (*userService)(nil)

type userService struct {
	userRepo   repositories.UserRepository
	twitchRepo repositories.TwitchRepository
}

func NewUserService(userRepo repositories.UserRepository, twitchRepo repositories.TwitchRepository) *userService {
	return &userService{userRepo: userRepo, twitchRepo: twitchRepo}
}

func (s *userService) CreateUser(username, email, password string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	existingUser, _ := s.GetUserByEmail(email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	user := &models.User{
		UserID:       uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: new(string),
		Currency:     0, // Default currency value
	}

	*user.PasswordHash = string(hashedPassword)

	user, err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *userService) FindOrCreateTwitchUser(twitchData *models.TwitchUser, token *oauth2.Token) (*models.User, error) {
	// First try to find an existing user by Twitch ID
	user, err := s.userRepo.GetUserByTwitchID(twitchData.ID)
	if err == nil && user != nil {
		log.Printf("Found existing user with Twitch ID: %s", twitchData.ID)

		// Update the associated Twitch user data
		// user.AccessToken = token.AccessToken
		// user.RefreshToken = token.RefreshToken
		// user.TokenExpiresAt = token.Expiry

		// user, err = s.userRepo.UpdateUser(user)
		// if err != nil {
		// 	log.Printf("Error updating user data: %v", err)
		// 	return nil, fmt.Errorf("failed to update user: %w", err)
		// }

		// Change this to update the user with the token
		err = s.twitchRepo.UpdateTwitchUser(twitchData)
		if err != nil {
			log.Printf("Error updating Twitch user data: %v", err)
			return nil, fmt.Errorf("failed to update Twitch user: %w", err)
		}

		return user, nil
	}

	// No existing user found, check if Twitch user exists independently
	existingTwitchUser, err := s.twitchRepo.GetTwitchUserByID(twitchData.ID)
	if err == nil && existingTwitchUser != nil {
		log.Printf("Found Twitch user but no associated user account, updating Twitch data")
		err = s.twitchRepo.UpdateTwitchUser(twitchData)
		if err != nil {
			log.Printf("Error updating existing Twitch user: %v", err)
			return nil, fmt.Errorf("failed to update Twitch user: %w", err)
		}
	} else {
		if err != nil {
			log.Printf("Error getting Twitch user: %v", err)
		}

		// create twitch user from twitch data
		err = s.twitchRepo.CreateTwitchUser(twitchData)
		if err != nil {
			log.Printf("Error creating Twitch user: %v", err)
			return nil, fmt.Errorf("failed to create Twitch user: %w", err)
		}
	}

	// Create new user with Twitch data
	log.Printf("Creating new user from Twitch data")
	// Update this to create the user with the token
	return s.CreateUserFromTwitch(twitchData)
}

func (s *userService) CreateUserFromTwitch(twitchData *models.TwitchUser) (*models.User, error) {
	// create new user
	user := &models.User{
		UserID:   uuid.New(),
		Username: twitchData.Login,
		Email:    twitchData.Email,
		TwitchID: &twitchData.ID,
	}

	user, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) LinkTwitchAccount(userID uuid.UUID, twitchData *models.TwitchUser, token *oauth2.Token) error {
	user, err := s.userRepo.GetUserByTwitchID(twitchData.ID)
	if err != nil {
		return err
	}

	// Check if this Twitch account is already linked to another user
	existingTwitchUser, _ := s.twitchRepo.GetTwitchUserByID(twitchData.ID)
	if existingTwitchUser != nil && user.UserID != userID {
		return errors.New("twitch account already linked to different user")
	}

	// Update/Create TwitchUser
	user.TwitchID = &twitchData.ID
	user.AccessToken = token.AccessToken
	user.RefreshToken = token.RefreshToken
	user.TokenExpiresAt = token.Expiry

	if existingTwitchUser != nil {
		return s.twitchRepo.UpdateTwitchUser(twitchData)
	}

	return s.twitchRepo.CreateTwitchUser(twitchData)
}

// TODO: Implement this
func (s *userService) UnlinkTwitchAccount(twitchID string) error {
	user, err := s.userRepo.GetUserByTwitchID(twitchID)
	if err != nil {
		return err
	}

	user.TwitchID = nil

	_, err = s.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) ValidateUserBalance(userID uuid.UUID, amount float64) bool {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return false
	}

	if user.Currency < amount {
		return false
	}

	return true
}

// Does this need to return the user?
func (s *userService) AddCurrency(userID uuid.UUID, amount float64) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	user.Currency += amount

	user, err = s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) SubtractCurrency(userID uuid.UUID, amount float64) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	user.Currency -= amount

	user, err = s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserCurrency(userID uuid.UUID) (float64, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return 0, err
	}

	return user.Currency, nil
}

func (s *userService) SetUserCurrency(userID uuid.UUID, amount float64) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	user.Currency = amount

	user, err = s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) ValidateUserForBid(userID uuid.UUID, amount float64) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if user.Currency < amount {
		return errors.New("insufficient funds")
	}

	return nil
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	user, err := s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
