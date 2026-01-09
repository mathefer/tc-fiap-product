package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mathefer/tc-fiap-product/docs"

	"github.com/mathefer/tc-fiap-product/internal/app"
)

// @title           Tc-Fiap-Product
// @version         1.0
// @description     Product microservice API
// @host            localhost:8081
// @BasePath        /
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capture system signals for graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	// Initialize the application using Uber FX
	app := app.InitializeApp()

	// Start the Uber FX lifecycle
	if err := app.Start(ctx); err != nil {
		log.Fatalf("Error while starting app: %v", err)
	}

	// Wait until the context is canceled
	<-ctx.Done()

	// Stop the Uber FX lifecycle
	if err := app.Stop(ctx); err != nil {
		log.Fatalf("Error while stopping app: %v", err)
	}
}
