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

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Debug("Error loading .env file")
	}

	configFile := os.Getenv("CONFIGFILE")
	if configFile == "" {
		configFile = "config.yaml"
	}

	logFile := os.Getenv("LOGFILE")
	if logFile == "" {
		logFile = "app.log"
	}

	log, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Error opening log file:", err)
	}
	defer log.Close()

	logrus.SetOutput(log)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	config, err := config.LoadConfig(configFile)
	if err != nil {
		logrus.Fatal("Error opening config file:", err)
	}

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/api/v1/sum", handlers.Sum)
	e.POST("/api/v1/multiply", handlers.Multiply)

	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	if err := e.Start(address); err != nil {
		logrus.Fatal("Error starting server:", err)
	}
	logrus.Info("Server started on : ", address)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	logrus.Info("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Printf("checking for graceful shutdown...\n")

	if err := e.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("Server Shutdown Failed:%+v", err)
	}

	logrus.Info("Server exited gracefully")
}
