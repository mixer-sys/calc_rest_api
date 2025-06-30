// @title		   Calculator REST API
// @version	   1.0
// @description   A simple REST API for performing basic arithmetic operations
// @host		   localhost:8080
// @BasePath	   /api/v1
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "calc_rest_api/api/openapi-spec/v1"

	config "calc_rest_api/internal/app/config"
	handlers "calc_rest_api/internal/app/handlers"
	"calc_rest_api/internal/app/logger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.yaml"
	}
	config, err := config.LoadConfig(configFile)
	if err != nil {
		logrus.Fatal("Error opening config file:", err)
	}

	logger := logger.GetLogger()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/api/v1/sum", handlers.Sum)
	e.POST("/api/v1/multiply", handlers.Multiply)

	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	if err := e.Start(address); err != nil {
		logger.Fatal("Error starting server: %+v", err)
	}
	logger.Info("Server started on %s: ", address)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	logger.Info("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Info("checking for graceful shutdown...\n")

	if err := e.Shutdown(shutdownCtx); err != nil {
		logger.Error("Server Shutdown Failed:%+v", err)
	}

	logger.Info("Server exited gracefully")
}
