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
	Numbers []float64 `json:"numbers"`
	UUID    string    `json:"uuid"`
}

type SumResponse struct {
	Sum  float64 `json:"sum"`
	UUID string  `json:"uuid"`
}

type MultiplyResponse struct {
	Multiply float64 `json:"sum"`
	UUID     string  `json:"uuid"`
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

func (s *SafeMap) Set(key string, value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *SafeMap) Get(key string) (float64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.data[key]
	return value, ok
}

func sum(c echo.Context) error {
	var data DataRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})

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

func multiply(c echo.Context) error {
	var data DataRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
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
