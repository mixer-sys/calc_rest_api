package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type DataRequest struct {
	Numbers []float64 `json:"numbers"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
}

type MultiplyResponse struct {
	Multiply float64 `json:"sum"`
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
	str := ""
	for _, number := range data.Numbers {
		sum += number
		str += fmt.Sprintf("%f, ", number)
	}

	SafeMap := NewSafeMap()
	SafeMap.Set(str, sum)

	if value, ok := SafeMap.Get(str); ok {
		fmt.Printf("Retrieved from SafeMap: %s = %f\n", str, value)
	}
	response := SumResponse{Sum: sum}
	return c.JSON(http.StatusOK, response)
}

func multiply(c echo.Context) error {
	var data DataRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	multiply := 1.0
	str := ""
	for _, number := range data.Numbers {
		multiply *= number
		str += fmt.Sprintf("%f, ", number)
	}
	SafeMap := NewSafeMap()
	SafeMap.Set(str, multiply)

	if value, ok := SafeMap.Get(str); ok {
		fmt.Printf("Retrieved from SafeMap: %s = %f\n", str, value)
	}
	response := MultiplyResponse{Multiply: multiply}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.POST("/sum", sum)
	e.POST("/multiply", multiply)
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println("Server started on :8080")
}
