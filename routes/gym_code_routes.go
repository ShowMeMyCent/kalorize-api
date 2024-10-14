package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/controllers"
)

func RouteGymCode(apiv1 *echo.Group, db *gorm.DB) {
	gymCodeController := controllers.NewGymCodeController(db)

	// Route to generate a new gym token
	apiv1.POST("/admin/generate-gym-token", gymCodeController.GenerateGymToken)

	// JWT middleware to authenticate the user
	tokenController := controllers.NewTokenController(db)
	apiv1.Use(tokenController.CheckTokenMiddleware())

	// Additional routes for gym code can be added here
}
