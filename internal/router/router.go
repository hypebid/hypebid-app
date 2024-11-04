package router

import (
	"github.com/hypebid/hypebid-app/internal/auth"
	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/handlers"
	"github.com/hypebid/hypebid-app/internal/middleware"
	"github.com/hypebid/hypebid-app/internal/repositories"
	"github.com/hypebid/hypebid-app/internal/services"
	"github.com/hypebid/hypebid-app/internal/twitch"
	"github.com/hypebid/hypebid-app/internal/validator"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// NewRouter creates a new router with all the routes defined
func NewRouter(cfg *config.Config, db *gorm.DB, oauthManager auth.OAuthManager) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	// r.Use(middleware.Logger)    // Log requests
	// r.Use(middleware.Recoverer) // Recover from panics

	// handlerCfg := handlers.BaseHandlerConfig{}

	userRepo := repositories.NewUserRepository(db)
	marketInstanceRepo := repositories.NewMarketInstanceRepository(db)
	memberRepo := repositories.NewMemberRepository(db)

	channelRepo := repositories.NewChannelRepository(db)
	metricRepo := repositories.NewMetricRepository(db)
	metricDataPointRepo := repositories.NewMetricDataPointRepository(db)

	auctionRepo := repositories.NewAuctionRepository(db)
	marketChannelRepo := repositories.NewMarketChannelRepository(db)
	shareHolderRepo := repositories.NewShareHolderRepository(db)
	twitchRepo := repositories.NewTwitchRepository(db)
	twitchClient := twitch.NewClient(cfg)

	// auctionValidator := validator.NewAuctionValidator(auctionRepo)
	channelValidator := validator.NewChannelValidator(channelRepo)
	marketInstanceValidator := validator.NewMarketInstanceValidator(marketInstanceRepo)
	marketChannelValidator := validator.NewMarketChannelValidator(marketChannelRepo)
	shareHolderValidator := validator.NewShareHolderValidator(shareHolderRepo)
	userValidator := validator.NewUserValidator(userRepo)
	auctionValidators := validator.NewAuctionValidators(auctionRepo, channelValidator, marketInstanceValidator, userValidator)
	marketChannelValidators := validator.NewMarketChannelValidators(marketChannelRepo, channelValidator, marketInstanceValidator, shareHolderValidator)
	shareHolderValidators := validator.NewShareHolderValidators(shareHolderRepo, marketChannelValidator)

	userService := services.NewUserService(userRepo, twitchRepo)
	marketInstanceService := services.NewMarketInstanceService(marketInstanceRepo)
	memberService := services.NewMemberService(memberRepo)

	channelService := services.NewChannelService(channelRepo)

	metricService := services.NewMetricService(metricRepo, channelService)
	metricDataPointService := services.NewMetricDataPointService(metricDataPointRepo, metricService)

	marketChannelService := services.NewMarketChannelService(marketChannelRepo, marketChannelValidators)
	shareHolderService := services.NewShareHolderService(shareHolderRepo, marketChannelService, userService, shareHolderValidators)

	auctionService := services.NewAuctionService(auctionRepo, userService, shareHolderService)

	authHandler := handlers.NewAuthHandler(cfg, oauthManager, userService)
	twitchMiddleware := middleware.NewTwitchMiddleware(cfg, oauthManager, userService)

	// Auth Routes
	r.Route("/api/v1/auth", func(r chi.Router) {
		r.Get("/twitch/login", authHandler.TwitchLoginHandler())
		r.Get("/twitch/callback", authHandler.TwitchCallbackHandler())
	})

	// Protected routes that require Twitch auth
	// Get Twitch data or update something here
	r.Group(func(r chi.Router) {
		r.Use(twitchMiddleware.RequireTwitchAuth)
		// Add your protected routes here
		// r.Get("/api/v1/auth/twitch/callback", authHandler.TwitchCallbackHandler())
	})

	// Twitch API Routes
	r.Route("/api/v1/twitch", func(r chi.Router) {
		r.Get("/followers", handlers.FollowersHandler(cfg))
		r.Get("/bulk-followers", handlers.BulkFollowersHandler(cfg))
		r.Get("/users", handlers.TwitchUsersHandler(cfg))
		r.Post("/users", handlers.TwitchUserHandler(cfg, userService))
	})

	// HypeBid API Routes
	r.Route("/api/v1/", func(r chi.Router) {
		// User Management
		r.Post("/users/register", handlers.RegisterUserHandler(userService))
		r.Post("/users/login", handlers.LoginUserHandler(userService))
		r.Post("/users/logout", handlers.LogoutUserHandler(cfg))
		r.Get("/users/{userId}/instances", handlers.GetUserInstancesHandler(marketInstanceService, userService))

		// Currency Management
		r.Post("/users/{userId}/currency", handlers.AddCurrencyHandler(userService))

		// Instance Management
		r.Post("/instances", handlers.CreateInstanceHandler(marketInstanceService, userService))
		r.Get("/instances", handlers.GetAllInstancesHandler(marketInstanceService))

		// r.Post("/instances/{instanceId}/join", handlers.JoinInstanceHandler(marketInstanceRepo, userRepo))
		r.Post("/instances/{instanceId}/start", handlers.StartInstanceHandler(marketInstanceService))
		r.Get("/instances/{instanceId}", handlers.GetInstanceHandler(marketInstanceService))

		// Instance Member Management
		r.Post("/instances/{instanceId}/users", handlers.AddUserToInstanceHandler(marketInstanceService, userService, memberService))
		r.Get("/instances/{instanceId}/users", handlers.GetAllMembersForInstanceHandler(memberService))

		// Channel Management
		r.Post("/channels", handlers.RegisterChannelHandler(channelService, twitchClient))
		r.Get("/channels", handlers.GetAllChannelsHandler(channelService))

		// Channel Data
		// recent follower count
		// r.Post("/channels/{login}/followers", handlers.FollowerCountHandler())
		// average follower count for day parameter, capped at 365 days
		r.Get("/channels/{login}/followers/average/{days}", handlers.AverageFollowerCountHandler(metricDataPointService))
		// last 15 minutes follower count
		r.Get("/channels/{login}/followers/recent", handlers.RecentFollowerCountHandler(metricService, metricDataPointService))

		// Market Channel Management
		r.Post("/instances/{instanceId}/channels", handlers.CreateMarketChannelHandler(marketChannelService, marketInstanceService, channelService))
		r.Get("/instances/{instanceId}/channels", handlers.GetMarketChannelsHandler(marketChannelService))

		// Shareholder Management
		r.Post("/instances/{instanceId}/channels/{marketChannelId}/shareholders", handlers.CreateShareHolderHandler(shareHolderService, marketChannelService, userService))
		// r.Get("/instances/{instanceId}/channels/{marketChannelId}/shareholders", handlers.GetShareHoldersHandler(shareHolderService))

		// Auction Management
		r.Post("/instances/{instanceId}/auctions", handlers.CreateAuctionHandler(auctionService, auctionValidators))
		r.Get("/instances/{instanceId}/auctions", handlers.GetAllAuctionsHandler(auctionService))
		r.Get("/instances/{instanceId}/auctions/current", handlers.GetCurrentAuctionHandler(cfg))
		r.Post("/instances/{instanceId}/auctions/{auctionId}/start", handlers.StartAuctionHandler(auctionService, auctionValidators))
		r.Post("/instances/{instanceId}/auctions/{auctionId}/bids", handlers.PlaceBidHandler(auctionService, auctionValidators, userService))
		r.Get("/instances/{instanceId}/auctions/{auctionId}", handlers.GetAuctionHandler(auctionService, auctionValidators))
		r.Get("/instances/{instanceId}/auctions/{auctionId}/history", handlers.GetAuctionHistoryHandler(cfg))

		// Trading Management
		r.Post("/instances/{instanceId}/trades", handlers.CreateTradeHandler(cfg))
		r.Get("/instances/{instanceId}/trades", handlers.GetPendingTradesHandler(cfg))
		r.Post("/instances/{instanceId}/trades/{tradeId}/{action}", handlers.ManageTradeHandler(cfg))
		r.Post("/instances/{instanceId}/trades/{tradeId}/boost", handlers.BoostTradeHandler(cfg))

		// Boost Management
		r.Get("/instances/{instanceId}/boosts/{userId}", handlers.GetBoostsForUserHandler(cfg))
		r.Get("/instances/{instanceId}/boosts/{userId}/active", handlers.GetActiveBoostsForUserHandler(cfg))
	})

	return r
}
