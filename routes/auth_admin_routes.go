package routes

import (
	"kalorize-api/app/controllers/admin"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAdminAuth(apiv1 *echo.Group, db *gorm.DB, signingKey string) {
	authController := admin.NewAuthController(db, signingKey)

	apiv1.POST("/admin/login", authController.Login)
	apiv1.GET("/admin/user", authController.GetUser)
	apiv1.POST("/admin/logout", authController.Logout)
	apiv1.POST("/admin/refresh", authController.Refresh)
}
