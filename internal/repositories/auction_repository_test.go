package repositories

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hypebid/hypebid-app/pkg/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// func TestMain(m *testing.M) {
// 	// Setup test container once for all tests
// 	ctx := context.Background()

// 	req := testcontainers.ContainerRequest{
// 		Image:        "postgres:latest",
// 		ExposedPorts: []string{"5432/tcp"},
// 		Env: map[string]string{
// 			"POSTGRES_DB":       "testdb",
// 			"POSTGRES_USER":     "test",
// 			"POSTGRES_PASSWORD": "test",
// 		},
// 		WaitingFor: wait.ForLog("database system is ready to accept connections"),
// 	}

// 	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
// 		ContainerRequest: req,
// 		Started:          true,
// 	})
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to start container: %v", err))
// 	}

// 	// Get the database connection
// 	port, _ := container.MappedPort(ctx, "5432")
// 	host, _ := container.Host(ctx)

// 	dsn := fmt.Sprintf("host=%s port=%s user=test password=test dbname=testdb sslmode=disable",
// 		host, port.Port())
// 	testDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to connect to database: %v", err))
// 	}

// 	// Run migrations
// 	err = testDB.AutoMigrate(&models.Auction{})
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to run migrations: %v", err))
// 	}

// 	// Run tests
// 	code := m.Run()

// 	// Cleanup
// 	container.Terminate(ctx)
// 	os.Exit(code)
// }

// Helper function to clean the database between tests
func cleanDB(t *testing.T) {
	t.Helper()
	testDB.Exec("TRUNCATE TABLE auctions CASCADE")
}

func TestCreateAuction(t *testing.T) {
	cleanDB(t)
	repo := NewAuctionRepository(testDB)

	auction := &models.Auction{
		AuctionID: uuid.New(),
		// Add other required fields
	}

	createdAuction, err := repo.CreateAuction(auction)
	assert.NoError(t, err)
	assert.NotNil(t, createdAuction)
	assert.Equal(t, auction.AuctionID, createdAuction.AuctionID)

	// Verify using raw SQL query to ensure it's really in the database
	var count int64
	testDB.Model(&models.Auction{}).Where("auction_id = ?", auction.AuctionID).Count(&count)
	assert.Equal(t, int64(1), count)
}

func TestGetAuctionByID(t *testing.T) {
	cleanDB(t)
	repo := NewAuctionRepository(testDB)

	auction := &models.Auction{
		AuctionID: uuid.New(),
		// Add other required fields
	}

	// Create directly in DB to test retrieval
	err := testDB.Create(auction).Error
	assert.NoError(t, err)

	// Test the get method
	fetchedAuction, err := repo.GetAuctionByID(auction.AuctionID)
	assert.NoError(t, err)
	assert.Equal(t, auction.AuctionID, fetchedAuction.AuctionID)
}

// Add more tests...
