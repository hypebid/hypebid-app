package database

import (
	"fmt"
	"log"

	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabase initializes the database connection and performs any necessary setup
func SetupDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	// Database connection code here
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Create extensions
	err = createExtensions(db)
	if err != nil {
		return nil, err
	}

	// Create enum types
	err = CreateEnumTypes(db)
	if err != nil {
		return nil, err
	}

	// Run migrations and create enum types
	err = runMigrations(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createExtensions(db *gorm.DB) error {
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		log.Fatalf("Error creating extension uuid-ossp: %v", err)
	}

	return nil
}

// runMigrations performs all necessary database migrations
func runMigrations(db *gorm.DB) error {

	// DB Updates / Migrations here
	// // Rename the DurationWeeks to DurationDays
	// err := RenameDurationWeeksToDurationDays(db)
	// if err != nil {
	// 	return err
	// }

	// // Drop the DurationWeeks column
	// err = DropDurationWeeksColumn(db)
	// if err != nil {
	// 	return err
	// }

	// Add the UserID column to the MarketInstance table
	// err := AddUserIDToMarketInstance(db)
	// if err != nil {
	// 	return err
	// }

	// Add the Role, JoinedAt, and Status fields to the members table
	// err := AddFieldsToMembers(db)
	// if err != nil {
	// 	return err
	// }

	var err error

	// // Add the 'active' value to the market_status enum
	// err = AddActiveToMarketStatus(db)
	// if err != nil {
	// 	return err
	// }

	// Update the Auction MarketChannelID column name
	// err = UpdateAuctionMarketChannelID(db)
	// if err != nil {
	// 	return err
	// }

	// err = UpdateTradeStructure(db)
	// if err != nil {
	// 	return err
	// }

	// Create the TwitchUsers table
	// err = CreateTwitchUsersTable(db)
	// if err != nil {
	// 	return err
	// }

	// // Add the TwitchUserID to the User table
	// err = AddTwitchUserToUser(db)
	// if err != nil {
	// 	return err
	// }

	// Drop the UserID column from the TwitchUsers table
	// err = DropUserIDFromTwitchUsersTable(db)
	// if err != nil {
	// 	return err
	// }

	// Allow null password hash and twitch id in the users table
	// err = AllowNullPasswordHashInUsersTable(db)
	// if err != nil {
	// 	return err
	// }

	// err = AllowNullTwitchIDInUsersTable(db)
	// if err != nil {
	// 	return err
	// }

	// Add the Auth fields to the users table
	// err = AddAuthFieldsToUsersTable(db)
	// if err != nil {
	// 	return err
	// }

	// Remove the Auth fields from the TwitchUsers table
	// err = RemoveAuthFieldsFromTwitchUsersTable(db)
	// if err != nil {
	// 	return err
	// }

	// Add the Currency field to the Members table
	// err = AddCurrencyToMembersTable(db)
	// if err != nil {
	// 	return err
	// }

	// Add the UserLogin and Name fields to the MetricEntry table
	// err = AddUserLoginToMetricEntry(db)
	// if err != nil {
	// 	return err
	// }

	// Auto-migrate your models
	err = db.AutoMigrate(
		&models.Channel{},
		&models.User{},
		&models.UserShare{},
		&models.Metric{},
		&models.MetricDataPoint{},
		&models.MarketInstance{},
		&models.Member{},
		&models.Boost{},
		&models.Auction{},
		&models.Trade{},
		&models.TradeItem{},
		&models.MarketChannel{},
		&models.ShareHolder{},
		&models.TwitchUser{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database schema: %v", err)
	}

	return nil
}

// createEnumTypes creates the necessary enum types in the database
func CreateEnumTypes(db *gorm.DB) error {

	// check if the enum types already exist before creating them
	statements := []string{
		"DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'market_status') THEN CREATE TYPE market_status AS ENUM ('not_started', 'in_auction', 'trading', 'completed'); END IF; END $$;",
		"DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'boost_type') THEN CREATE TYPE boost_type AS ENUM ('2x', '4x', '8x'); END IF; END $$;",
		"DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'auction_status') THEN CREATE TYPE auction_status AS ENUM ('open', 'closed'); END IF; END $$;",
		"DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'trade_status') THEN CREATE TYPE trade_status AS ENUM ('pending', 'accepted', 'rejected'); END IF; END $$;",
		"DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'trade_direction') THEN CREATE TYPE trade_direction AS ENUM ('offer', 'request'); END IF; END $$;",
	}

	for _, stmt := range statements {
		if err := db.Exec(stmt).Error; err != nil {
			// Can check for other errors, should already be checking if the type exists
			return err
		}
	}

	return nil
}
