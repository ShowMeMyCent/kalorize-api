package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/controllers"
	"kalorize-api/app/controllers/admin"
)

func RouteGym(apiv1 *echo.Group, db *gorm.DB) {
	gymAdminController := admin.NewGymController(db)
	gymController := controllers.NewGymController(db)

	// JWT middleware to authenticate the user
	tokenController := controllers.NewTokenController(db)
	apiv1.Use(tokenController.CheckTokenMiddleware())

	// Route to get all gyms
	apiv1.GET("/gyms", gymController.GetAllGyms)

	// Route to get a gym by ID
	apiv1.GET("/gyms/:idGym", gymController.GetGymById)

	// Route For Admin

	apiv1.POST("/admin/gyms", gymAdminController.CreateGym)
	apiv1.PUT("/admin/gyms/:idGym", gymAdminController.UpdateGym)
	apiv1.DELETE("/admin/gyms/delete/:idGym", gymAdminController.DeleteGym)
}
