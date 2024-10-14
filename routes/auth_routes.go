package routes

import (
	"kalorize-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAuth(apiv1 *echo.Group, db *gorm.DB, signingKey string) {
	authController := controllers.NewAuthController(db, signingKey)

	apiv1.POST("/login", authController.Login)
	apiv1.POST("/logout", authController.Logout)
	apiv1.POST("/register", authController.Register)
	apiv1.GET("/user", authController.GetUser)
	apiv1.POST("/refresh", authController.Refresh)
}
