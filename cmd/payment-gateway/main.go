package main

import (
	"context"
	"log"

	app "github.com/itua234/payment-bridge/internal/app"
)

func main() {
	ctx := context.Background()

	application, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	//Run the Server
	if err := application.Run(); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
