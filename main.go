package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SumRequest struct {
	Numbers []float64 `json:"numbers"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
}

func sum(c echo.Context) error {
	var data SumRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})

	}

	sum := 0.0
	for _, number := range data.Numbers {
		sum += number
	}

	response := SumResponse{Sum: sum}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()

	e.POST("/sum", sum)
	e.Logger.Fatal(e.Start(":8080"))
	fmt.Println("Server started on :8080")
}
