package commonHandler

import (
	core "calc_rest_api/internal/app/core"
	logger "calc_rest_api/internal/app/logger"
	"net/http"

	"github.com/labstack/echo/v4"
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

// @Summary Sum of numbers
// @Description Calculate the sum of a list of numbers
// @Tags Calculator
// @Accept json
// @Produce json
// @Param data body DataRequest true "Data Request"
// @Success 200 {object} SumResponse
// @Failure 400 {object} BadRequest
// @Router /sum [post]
func Sum(c echo.Context) error {

	logger := logger.GetLogger()
	var data DataRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, BadRequest{Error: "Invalid input"})

	}

	sum := 0.0

	for _, number := range data.Numbers {
		sum += number
	}

	SafeMap := core.NewSafeMap()
	SafeMap.Set(data.UUID, sum)

	if value, ok := SafeMap.Get(data.UUID); ok {
		logger.Info("Retrieved from SafeMap: %s = %f\n", data.UUID, value)
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
func Multiply(c echo.Context) error {
	logger := logger.GetLogger()
	var data DataRequest
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, BadRequest{Error: "Invalid input"})
	}

	multiply := 1.0

	for _, number := range data.Numbers {
		multiply *= number

	}
	SafeMap := core.NewSafeMap()
	SafeMap.Set(data.UUID, multiply)

	if value, ok := SafeMap.Get(data.UUID); ok {
		logger.Info("Retrieved from SafeMap: %s = %f\n", data.UUID, value)
	}
	response := MultiplyResponse{Multiply: multiply, UUID: data.UUID}
	return c.JSON(http.StatusOK, response)
}
