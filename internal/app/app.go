package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/itua234/payment-gateway/internal/config"
	"github.com/itua234/payment-gateway/internal/database"
)

// Application holds the dependencies for the HTTP server
type Application struct {
	Config *config.Config
	Router *gin.Engine
}

// New initializes the application and its dependencies
func New(ctx context.Context) (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	if err := database.Connect(cfg.DB); err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}
	fmt.Println("Database connected.")

	return &Application{
		Config: cfg,
	}, nil
}

// Run starts the HTTP server
func (a *Application) Run() error {
	log.Printf("Server starting on port :%s", a.Config.Port)
	return a.Router.Run(":" + a.Config.Port)
}
