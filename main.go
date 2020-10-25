package main

import (
	"autodesk-go/config"
	"autodesk-go/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"autodesk-go/identity"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	config.GetConfig()
	RegisterRoute(e)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type HealthStatus struct {
	Version string `json:"version"`
}

// Handler
func Health(c echo.Context) error {
	healthStatus := HealthStatus{
		Version: "0.1.0",
	}
	return utils.SuccessResponse(c, healthStatus)
}

func RegisterRoute(e *echo.Echo) {
	e.GET("/", Health)

	identity.RegisterRoute(e)
}
