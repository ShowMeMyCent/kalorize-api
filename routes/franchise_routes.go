package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/controllers"
	"kalorize-api/app/controllers/admin"
)

func RoutesFranchise(apiv1 *echo.Group, db *gorm.DB) {

	// Initialize controllers
	franchiseAdminController := admin.NewFranchiseController(db)
	franchiseController := controllers.NewFranchiseController(db)

	// Apply JWT middleware
	tokenController := controllers.NewTokenController(db)
	apiv1.Use(tokenController.CheckTokenMiddleware())

	// Route to get all franchises
	apiv1.GET("/franchises", franchiseController.GetAllFranchises)

	// Route to get a franchise by ID
	apiv1.GET("/franchises/:id", franchiseController.GetFranchiseById)

	// Route to get a franchise by name
	apiv1.GET("/franchises/name/:name", franchiseController.GetFranchiseByName)

	// Route for admin
	apiv1.POST("/admin/franchises", franchiseAdminController.CreateFranchiseWithMakanan)
	apiv1.PUT("/admin/franchises/:id", franchiseAdminController.UpdateFranchise)
	apiv1.DELETE("/admin/franchises/:id", franchiseAdminController.DeleteFranchise)
}
