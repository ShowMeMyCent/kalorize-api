package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GymOwnerRoute(apiv1 *echo.Group, db *gorm.DB) {
	// gymOwnerController := controllers.NewGymOwnerController(db)

	// apiv1.POST("/gym-owner", gymOwnerController.CreateGymOwner)
	// apiv1.GET("/gym-owner", gymOwnerController.GetGymOwner)
	// apiv1.GET("/gym-owner/:id", gymOwnerController.GetGymOwnerById)
	// apiv1.PUT("/gym-owner/:id", gymOwnerController.UpdateGymOwner)
	// apiv1.DELETE("/gym-owner/:id", gymOwnerController.DeleteGymOwner)
}
