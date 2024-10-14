package routes

import (
	"kalorize-api/app/controllers"
	"kalorize-api/app/controllers/admin"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteMakanan(apiv1 *echo.Group, db *gorm.DB) {
	makananAdminController := admin.NewMakananController(db)
	makananController := controllers.NewMakananController(db)

	//JWT middleware to authenticate the user
	tokenController := controllers.NewTokenController(db)
	apiv1.Use(tokenController.CheckTokenMiddleware())

	// Existing routes
	apiv1.GET("/makanan", makananController.GetAllMakanan)
	apiv1.GET("/makanan/csv", makananController.GetMakananCSV)
	apiv1.GET("/makanan/:makananId", makananController.GetMakananById)

	// Route for Admin
	apiv1.POST("/admin/makanan", makananAdminController.CreateMakanan)    // Route to create a new Makanan
	apiv1.PUT("/admin/makanan/:id", makananAdminController.UpdateMakanan) // Route to update an existing Makanan
}
