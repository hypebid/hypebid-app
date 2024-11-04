package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/hypebid/hypebid-app/database"
	"github.com/hypebid/hypebid-app/internal/auth"
	"github.com/hypebid/hypebid-app/internal/config"
	"github.com/hypebid/hypebid-app/internal/router"
	"github.com/hypebid/hypebid-app/internal/tasks"
	"gorm.io/gorm"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	log.Println("Starting server...")
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	log.Println("Setting up database...")

	// Create a connection with retry logic
	var db *gorm.DB
	for i := 0; i < 5; i++ {
		log.Printf("Attempting database connection (attempt %d/5)...", i+1)
		db, err = database.SetupDatabase(cfg)
		if err == nil {
			log.Println("Database connected successfully")
			break
		}
		log.Printf("Database connection attempt failed: %v", err)
		if i < 4 {
			time.Sleep(5 * time.Second)
		}
	}
	if err != nil {
		log.Fatal("Failed to connect to database after 5 attempts:", err)
	}

	oauthManager := auth.NewOAuthManager(cfg)

	// Need a channel manager to track which logins are being tracked
	// recommended pool: Rubius,KaiCenat,caseoh_ElMariana,Jynxzi,PirateSoftware,Mictia00,FeirlyGab,brino,GUACAMOLEMOLLY
	// recommended pool 2: Caedrel,Papaplatte,TenZ,deepins02,alondrissa,Lacy,T2x2,zarbex,nicoleheart23,EDISON
	// logins := []string{"Rubius", "KaiCenat", "caseoh_ElMariana", "Jynxzi", "PirateSoftware", "Mictia00", "FeirlyGab", "brino", "GUACAMOLEMOLLY", "Lirik"}
	// test_logins := []string{"theconducter", "molson82", "theprimeagen", "ruhdacted", "xqc"} // Add the logins you want to track
	// tasks.StartFollowerUpdater(cfg, db, logins)
	tasks.StartFollowerUpdater(cfg, db, cfg.TrackedLogins)

	r := router.NewRouter(cfg, db, oauthManager)

	// Testing route
	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "base.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Add startup probe endpoint
	// r.Get("/startup", func(w http.ResponseWriter, r *http.Request) {
	// 	if db != nil {
	// 		// Try a simple query to verify DB connection
	// 		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	// 		defer cancel()

	// 		if err := db.WithContext(ctx).Raw("SELECT 1").Scan(&struct{}{}).Error; err != nil {
	// 			log.Printf("Database ping failed: %v", err)
	// 			http.Error(w, "Database not ready", http.StatusServiceUnavailable)
	// 			return
	// 		}
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// })

	log.Println("Health check route registered")
	// health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan

		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Server shutdown error: %v", err)
		}
		os.Exit(0)
	}()

	log.Printf("Server starting on port %s...", cfg.ServerPort)
	log.Println("Environment:", cfg.Environment)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
