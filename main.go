package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/torvald2/faces/app_logger"
	"github.com/torvald2/faces/config"
	"github.com/torvald2/faces/handlers"
	"go.uber.org/zap"
)

func main() {

	// Config logger
	conf := config.GetConfig()

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
