package database

import (
	"github.com/hypebid/hypebid-app/pkg/models"

	"gorm.io/gorm"
)

func RenameDurationWeeksToDurationDays(db *gorm.DB) error {
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='market_instances' AND column_name='DurationWeeks')").Scan(&exists).Error
	if err != nil {
		return err
	}

	if exists {
		// Proceed to rename the column
		return db.Exec("ALTER TABLE market_instances RENAME COLUMN DurationWeeks TO duration_days").Error
	}

	// If the column does not exist, return nil or handle as needed
	return nil
}

// DropDurationWeeksColumn checks if the column exists before dropping it
func DropDurationWeeksColumn(db *gorm.DB) error {
	// Check if the column exists
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name='market_instances' AND column_name='duration_weeks')").Scan(&exists).Error
	if err != nil {
		return err
	}

	if exists {
		// Proceed to drop the column
		return db.Exec("ALTER TABLE market_instances DROP COLUMN duration_weeks").Error
	}

	// If the column does not exist, return nil or handle as needed
	return nil
}

func AddUserIDToMarketInstance(db *gorm.DB) error {
	statements := []string{
		// Step 1: Add the column allowing NULL values
		"ALTER TABLE market_instances ADD COLUMN user_id uuid;",

		// Step 2: Update existing rows with a default value (e.g., a specific user ID or NULL)
		"UPDATE market_instances SET user_id = '57a783e0-fc75-4933-9866-63224cc345b5' WHERE user_id IS NULL;",

		// Step 3: Alter the column to be NOT NULL
		"ALTER TABLE market_instances ALTER COLUMN user_id SET NOT NULL;",
	}

	for _, stmt := range statements {
		if err := db.Exec(stmt).Error; err != nil {
			// Can check for other errors, should already be checking if the type exists
			return err
		}
	}

	return nil
}

func AddFieldsToMembers(db *gorm.DB) error {
	// Add the Role, JoinedAt, and Status fields to the members table
	err := db.Migrator().AddColumn(&models.Member{}, "Role")
	if err != nil {
		return err
	}

	err = db.Migrator().AddColumn(&models.Member{}, "JoinedAt")
	if err != nil {
		return err
	}

	err = db.Migrator().AddColumn(&models.Member{}, "Status")
	if err != nil {
		return err
	}

	// Optionally, you can set default values for existing records
	// This assumes that the Role and Status fields are nullable
	// If they are not nullable, you may need to handle existing records differently
	err = db.Exec("UPDATE members SET Role = 'participant', Status = 'active' WHERE Role IS NULL OR Status IS NULL").Error
	if err != nil {
		return err
	}

	return nil
}

// AddActiveToMarketStatus adds the 'active' value to the market_status enum type
func AddActiveToMarketStatus(db *gorm.DB) error {
	// Execute the SQL command to add the new value to the enum
	return db.Exec("ALTER TYPE market_status ADD VALUE 'active'").Error
}

func UpdateTradeStructure(db *gorm.DB) error {
	// Drop the trades table if it exists
	var exists bool
	err := db.Raw("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'trades')").Scan(&exists).Error
	if err != nil {
		return err
	}
	if exists {
		if err := db.Exec("DROP TABLE trades").Error; err != nil {
			return err
		}
	}

	// Recreate the trades table
	err = db.Exec(`
        CREATE TABLE IF NOT EXISTS trades (
            trade_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
            market_instance_id uuid,
            initiator_id uuid,
            recipient_id uuid,
            status trade_status DEFAULT 'pending',
            created_at timestamp DEFAULT CURRENT_TIMESTAMP,
            completed_at timestamp,
            boost_id uuid
        )
    `).Error
	if err != nil {
		return err
	}

	// Create the new trade_items table
	err = db.Exec(`
        CREATE TABLE IF NOT EXISTS trade_items (
            trade_item_id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
            trade_id uuid REFERENCES trades(trade_id),
            market_channel_id uuid,
            market_instance_id uuid,
            share_count integer,
            currency numeric(14,2),
            direction trade_direction,
            created_at timestamp DEFAULT CURRENT_TIMESTAMP
        )
    `).Error
	if err != nil {
		return err
	}

	// Add indexes for better query performance
	indexes := []string{
		"CREATE INDEX IF NOT EXISTS idx_trade_items_trade_id ON trade_items(trade_id)",
		"CREATE INDEX IF NOT EXISTS idx_trade_items_market_channel_id ON trade_items(market_channel_id)",
		"CREATE INDEX IF NOT EXISTS idx_trades_initiator_id ON trades(initiator_id)",
		"CREATE INDEX IF NOT EXISTS idx_trades_recipient_id ON trades(recipient_id)",
	}

	for _, idx := range indexes {
		if err := db.Exec(idx).Error; err != nil {
			return err
		}
	}

	return nil
}

// drop trade items table
func DropTradeItemsTable(db *gorm.DB) error {
	return db.Exec("DROP TABLE trade_items").Error
}

func AddTwitchUserToUser(db *gorm.DB) error {
	return db.Exec("ALTER TABLE users ADD COLUMN twitch_user_id VARCHAR(255) REFERENCES twitch_users(id)").Error
}

func CreateTwitchUsersTable(db *gorm.DB) error {
	return db.Exec(`
    CREATE TABLE IF NOT EXISTS twitch_users (
        id VARCHAR(255) PRIMARY KEY,
        user_id UUID UNIQUE REFERENCES users(user_id),
        login VARCHAR(255) UNIQUE,
        display_name VARCHAR(255),
        type VARCHAR(50),
        broadcaster_type VARCHAR(50),
        description TEXT,
        profile_image_url TEXT,
        offline_image_url TEXT,
        view_count INTEGER,
        email VARCHAR(255),
        created_at VARCHAR(255),
        access_token TEXT,
        refresh_token TEXT,
        token_expires_at TIMESTAMP
    )
`).Error
}

func DropUserIDFromTwitchUsersTable(db *gorm.DB) error {
	return db.Exec("ALTER TABLE twitch_users DROP COLUMN user_id").Error
}

func AllowNullPasswordHashInUsersTable(db *gorm.DB) error {
	return db.Exec("ALTER TABLE users ALTER COLUMN password_hash DROP NOT NULL").Error
}

func AllowNullTwitchIDInUsersTable(db *gorm.DB) error {
	return db.Exec("ALTER TABLE users ALTER COLUMN twitch_id DROP NOT NULL").Error
}

func AddAuthFieldsToUsersTable(db *gorm.DB) error {
	// Add AccessToken, RefreshToken, TokenExpiresAt, and LastLoginAt to the users table
	if err := db.Migrator().AddColumn(&models.User{}, "AccessToken"); err != nil {
		return err
	}
	if err := db.Migrator().AddColumn(&models.User{}, "RefreshToken"); err != nil {
		return err
	}
	if err := db.Migrator().AddColumn(&models.User{}, "TokenExpiresAt"); err != nil {
		return err
	}
	if err := db.Migrator().AddColumn(&models.User{}, "LastLoginAt"); err != nil {
		return err
	}
	return nil
}

func RemoveAuthFieldsFromTwitchUsersTable(db *gorm.DB) error {
	statements := []string{
		"ALTER TABLE twitch_users DROP COLUMN access_token;",
		"ALTER TABLE twitch_users DROP COLUMN refresh_token;",
		"ALTER TABLE twitch_users DROP COLUMN token_expires_at;",
	}

	for _, stmt := range statements {
		if err := db.Exec(stmt).Error; err != nil {
			// Can check for other errors, should already be checking if the type exists
			return err
		}
	}

	return nil
}

func AddCurrencyToMembersTable(db *gorm.DB) error {
	return db.Exec("ALTER TABLE members ADD COLUMN currency numeric(14,2) DEFAULT 0").Error
}

func AddUserLoginToMetricEntry(db *gorm.DB) error {
	statements := []string{
		"ALTER TABLE metric_entries ADD COLUMN user_login VARCHAR(255);",
		"ALTER TABLE metric_entries ADD COLUMN name VARCHAR(255);",
	}

	for _, stmt := range statements {
		if err := db.Exec(stmt).Error; err != nil {
			// Can check for other errors, should already be checking if the type exists
			return err
		}
	}

	return nil
}
