package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/controllers"
)

func RouteUser(apiv1 *echo.Group, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	apiv1.POST("/users", userController.CreateUser) // Route to create a new User

	// JWT middleware to authenticate the user
	tokenController := controllers.NewTokenController(db)
	apiv1.Use(tokenController.CheckTokenMiddleware())

	// Existing routes
	apiv1.GET("/list/user", userController.GetAllUsers)
	apiv1.DELETE("/users/:id", userController.DeleteUser)

	// New routes for creating, updating, and deleting User
	apiv1.PUT("/users/:id", userController.UpdateUser)  // Route to update an existing User
	apiv1.GET("/users/:id", userController.GetUserById) // Route to delete a User
}
