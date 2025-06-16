// @title		   Calculator REST API
// @version	   1.0
// @description   A simple REST API for performing basic arithmetic operations
// @host		   localhost:8080
// @BasePath	   /api/v1
package main

import (
	"fmt"

	_ "calc_rest_api/api/docs"

	commonHandler "calc_rest_api/pkg/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/api/v1/sum", commonHandler.Sum)
	e.POST("/api/v1/multiply", commonHandler.Multiply)
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println("Server started on :8080")
}
