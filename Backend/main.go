package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/utkarshg408/Portfolio_Tracker/Backend/internal/config"
	"github.com/utkarshg408/Portfolio_Tracker/Backend/internal/database"
	"github.com/utkarshg408/Portfolio_Tracker/Backend/internal/httpapi"
	"github.com/utkarshg408/Portfolio_Tracker/Backend/internal/market"
)

func main() {
	cfg := config.Load()
	logger := log.New(os.Stdout, "portfolio-api ", log.LstdFlags|log.LUTC)
	ctx := context.Background()
	db, err := database.Open(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Fatalf("database connection failed: %v", err)
	}
	defer db.Close()
	server := &http.Server{Addr: ":" + cfg.Port, Handler: httpapi.NewMux(db, market.NewFinnhubClient(cfg.FinnhubAPIKey)), ReadHeaderTimeout: 5 * time.Second}
	go func() {
		logger.Printf("listening on http://localhost:%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("server failed: %v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Printf("graceful shutdown failed: %v", err)
	}
}
