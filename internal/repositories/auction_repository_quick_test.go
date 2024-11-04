package repositories

import (
	"testing"

	"github.com/hypebid/hypebid-app/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAuction_Quick(t *testing.T) {
	if !testing.Short() {
		t.Skip("skipping quick tests in normal mode")
	}

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	db.AutoMigrate(&models.Auction{})

	// Quick tests here...
}
