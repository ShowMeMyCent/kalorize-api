package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/controllers"
)

func RouteHistory(apiv1 *echo.Group, db *gorm.DB) {
	historyController := controllers.NewHistoryController(db)

	// JWT middleware to authenticate the user
	tokenController := controllers.NewTokenController(db)
	apiv1.Use(tokenController.CheckTokenMiddleware())

	// Route to get all histories
	apiv1.GET("/histories", historyController.GetAllHistories)

	// Route to get a history by ID
	apiv1.GET("/histories/:idHistory", historyController.GetHistoryById)

	// Route to create a new history
	apiv1.POST("/histories", historyController.CreateHistory)

	// Route to update an existing history
	apiv1.PUT("/histories/:idHistory", historyController.UpdateHistory)

	// Route to delete a history
	apiv1.DELETE("/histories/:idHistory", historyController.DeleteHistory)
}
