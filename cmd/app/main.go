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
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/api/v1/sum", handlers.Sum)
	e.POST("/api/v1/multiply", handlers.Multiply)

	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Fatal("Error starting server:", err)
		}
		fmt.Println("Server started on :8080")
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Корректное завершение работы сервера
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Error shutting down server:", err)
	}

	fmt.Println("Server gracefully stopped")
}
