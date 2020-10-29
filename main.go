package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Config logger
	logger, _ := zap.NewProduction()
	undo := zap.ReplaceGlobals(logger)
	defer undo()
	//Load .env file (for dev environment)
	err := godotenv.Load()
	if err != nil {
		zap.L().Info("No .env files found. Using real environment")
	}

}
