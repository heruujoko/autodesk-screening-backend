package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SuccessResponse(ctx echo.Context, data interface{}) error {
	responseData := map[string]interface{}{
		"success": true,
		"data":    data,
	}

	return ctx.JSON(http.StatusOK, responseData)
}

func ErrorResponse(ctx echo.Context, error interface{}) error {
	responseData := map[string]interface{}{
		"success": false,
		"error":   error,
	}

	return ctx.JSON(http.StatusBadRequest, responseData)
}
