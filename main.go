package main

import (
	"github.com/labstack/echo/v4"
	"autodesk-go/utils"
	"autodesk-go/config"

	"autodesk-go/identity"
)

func main() {
	e := echo.New()
	config.GetConfig()
	RegisterRoute(e)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type HealthStatus struct {
	Version    string     `json:"version"`
}

// Handler
func Health(c echo.Context) error {
	healthStatus := HealthStatus {
		Version: "0.1.0",
	}
	return utils.SuccessResponse(c,healthStatus)
}

func RegisterRoute(e *echo.Echo) {
	e.GET("/", Health)

	identity.RegisterRoute(e)
}
