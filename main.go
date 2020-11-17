package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "atbmarket.comfaceapp/app_logger"
	"atbmarket.comfaceapp/config"
	"atbmarket.comfaceapp/handlers"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	conf := config.GetConfig()
	// Config logger

	//Load .env file (for dev environment)
	if env := conf.IsDev; env != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			log.Logger.Info("No .env files found. Using real environment")
		}

	}

	wait := time.Second * 2
	log.Logger.Info("Current environment", zap.String("Config", fmt.Sprintf("%v", conf)))
	router := handlers.NewRouter()
	srv := &http.Server{
		Addr:    ":" + conf.TCP_Port,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
		log.Logger.Info("Listen is started")
	}()
	// Listen for os sygnals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	//Grasefun shutdown
	log.Logger.Info("Star server Shutdown on system sygnal")
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)

}
