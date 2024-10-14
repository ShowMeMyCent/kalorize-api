package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() (*echo.Group, *echo.Echo) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://kalorize-api.fly.dev", "*"},
	}))
	apiv1 := e.Group("/api/v1")
	e.Debug = true
	return apiv1, e
}
