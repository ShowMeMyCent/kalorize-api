package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/controllers"
	admin "kalorize-api/app/controllers/admin"
)

func RouteUserAdmin(apiv1 *echo.Group, db *gorm.DB) {
	userAdminController := admin.NewUserAdminController(db)

	// JWT middleware to authenticate the user
	tokenController := controllers.NewTokenController(db)
	apiv1.Use(tokenController.CheckTokenMiddleware())

	// Existing routes
	apiv1.GET("/list/admin/user", userAdminController.GetAllUsers)
	apiv1.DELETE("/admin/user/admin/:id", userAdminController.DeleteUser)

	// New routes for creating, updating, and deleting User
	apiv1.POST("/admin/user", userAdminController.CreateUser)     // Route to create a new User
	apiv1.PUT("/admin/user/:id", userAdminController.UpdateUser)  // Route to update an existing User
	apiv1.GET("/admin/user/:id", userAdminController.GetUserById) // Route to delete a User
}
