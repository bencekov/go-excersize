package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bencekov/go-exercise/internal/logging"
	"github.com/bencekov/go-exercise/pkg/api"
	chi "github.com/go-chi/chi/v5"
)

func main() {
	logger := logging.NewLogger()

	logger.Infof("Setting up web server")

	chi_router := chi.NewMux()

	api.NewAPI(
		api.NewService(logger),
		logger,
	).RegisterEndpoints(chi_router)

	logger.Infof("Starting on port: %s", "8080")

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: chi_router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalf("Error: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	server.Shutdown(ctx)

	logger.Desugar().Sync()

	logger.Infof("Shutting down")
	os.Exit(0)
}
