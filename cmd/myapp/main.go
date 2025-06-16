// @title		   Calculator REST API
// @version	   1.0
// @description   A simple REST API for performing basic arithmetic operations
// @host		   localhost:8080
// @BasePath	   /api/v1
package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "calc_rest_api/api/docs"
)

type DataRequest struct {
	Numbers []float64 `json:"numbers" example:"1.5,2.0,3.0"`
	UUID    string    `json:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type BadRequest struct {
	Error string `json:"error" example:"Bad request"`
}

type SumResponse struct {
	Sum  float64 `json:"sum" example:"6.3"`
	UUID string  `json:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type MultiplyResponse struct {
	Multiply float64 `json:"multiply" example:"6.1"`
	UUID     string  `json:"uuid" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type SafeMap struct {
	mu   sync.Mutex
	data map[string]float64
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]float64),
	}
}

func (safeMap *SafeMap) Set(key string, value float64) {
	safeMap.mu.Lock()
	defer safeMap.mu.Unlock()
	safeMap.data[key] = value
}

func (safeMap *SafeMap) Get(key string) (float64, bool) {
	safeMap.mu.Lock()
	defer safeMap.mu.Unlock()
	value, ok := safeMap.data[key]
	return value, ok
}

// @Summary Sum of numbers
// @Description Calculate the sum of a list of numbers
// @Tags Calculator
// @Accept json
// @Produce json
// @Param data body DataRequest true "Data Request"
// @Success 200 {object} SumResponse
// @Failure 400 {object} BadRequest
// @Router /sum [post]
func sum(c echo.Context) error {
	var data DataRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, BadRequest{Error: "Invalid input"})

	}

	sum := 0.0

	for _, number := range data.Numbers {
		sum += number
	}

	SafeMap := NewSafeMap()
	SafeMap.Set(data.UUID, sum)

	if value, ok := SafeMap.Get(data.UUID); ok {
		fmt.Printf("Retrieved from SafeMap: %s = %f\n", data.UUID, value)
	}
	response := SumResponse{Sum: sum, UUID: data.UUID}
	return c.JSON(http.StatusOK, response)
}

// @Summary Multiply of numbers
// @Description Calculate the multiply of numbers
// @Tags Calculator
// @Accept json
// @Produce json
// @Param data body DataRequest true "Data Request"
// @Success 200 {object} MultiplyResponse
// @Failure 400 {object} BadRequest
// @Router /multiply [post]
func multiply(c echo.Context) error {
	var data DataRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, BadRequest{Error: "Invalid input"})
	}

	multiply := 1.0

	for _, number := range data.Numbers {
		multiply *= number

	}
	SafeMap := NewSafeMap()
	SafeMap.Set(data.UUID, multiply)

	if value, ok := SafeMap.Get(data.UUID); ok {
		fmt.Printf("Retrieved from SafeMap: %s = %f\n", data.UUID, value)
	}
	response := MultiplyResponse{Multiply: multiply, UUID: data.UUID}
	return c.JSON(http.StatusOK, response)
}

func main() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/api/v1/sum", sum)
	e.POST("/api/v1/multiply", multiply)
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println("Server started on :8080")
}
