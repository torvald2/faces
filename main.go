package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"atbmarket.comfaceapp/handlers"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Config logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()
	//Load .env file (for dev environment)
	if env := os.Getenv("ENVIRONMENT"); env != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			logger.Info("No .env files found. Using real environment")
		}

	}

	wait := time.Second * 2
	logger.Info("Current environment", zap.String("DB_CONNECTION_STRING", os.Getenv("DB_CONNECTION_STRING")))
	router := handlers.NewRouter()
	srv := &http.Server{
		Addr:    ":" + "8080",
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	// Listen for os sygnals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	//Grasefun shutdown
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)

}
