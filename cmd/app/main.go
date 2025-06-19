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

	handlers "calc_rest_api/internal/app/handlers"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	LOGFILE := "app.log"
	CONFIGFILE := "config"
	CONFIGTYPE := "yaml"

	logFile, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Error opening log file:", err)
	}
	defer logFile.Close()

	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	viper.SetConfigName(CONFIGFILE)
	viper.SetConfigType(CONFIGTYPE)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Error reading config file: %s", err)
	}

	port := viper.GetString("server.port")
	host := viper.GetString("server.host")

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/api/v1/sum", handlers.Sum)
	e.POST("/api/v1/multiply", handlers.Multiply)

	address := fmt.Sprintf("%s:%s", host, port)
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
