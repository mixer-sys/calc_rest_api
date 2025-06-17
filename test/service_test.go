package service_test

import (
	"bytes"
	commonHandler "calc_rest_api/pkg/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/sum", bytes.NewBufferString(`{"numbers":[1.5,2.0,3.0],"uuid":"123e4567-e89b-12d3-a456-426614174000"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, commonHandler.Sum(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"sum":6.5,"uuid":"123e4567-e89b-12d3-a456-426614174000"}`, rec.Body.String())
	}
}

func TestMultiply(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/multiply", bytes.NewBufferString(`{"numbers":[1.5,2.0,3.0],"uuid":"123e4567-e89b-12d3-a456-426614174000"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, commonHandler.Multiply(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, `{"multiply":9.0,"uuid":"123e4567-e89b-12d3-a456-426614174000"}`, rec.Body.String())
	}
}
