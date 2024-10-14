package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"kalorize-api/config"
	"kalorize-api/routes"
)

func main() {
	// Initialize the database and signing key
	db, signingKey, err := config.InitDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	config.AutoMigration(db)

	// Initialize the Echo instance and routes
	route, e := routes.Init()

	// Middleware to inject signingKey into the context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Set the signingKey in the context
			c.Set("signingKey", signingKey)
			return next(c)
		}
	})

	// Define the application routes
	routes.RouteAuth(route, db, signingKey)
	routes.RouteAdminAuth(route, db, signingKey)
	routes.RouteUser(route, db)
	routes.RouteMakanan(route, db)
	routes.RouteGymCode(route, db)
	routes.RouteQuestionnaire(route, db)
	routes.RouteUserAdmin(route, db)
	routes.RoutePhotoStatic(route)
	routes.RouteImportDatabase(route, db)
	routes.RouteGym(route, db)
	routes.RoutesFranchise(route, db)
	routes.RouteHistory(route, db)

	// Start the server
	port := 8080
	address := fmt.Sprintf("0.0.0.0:%d", port)
	e.Logger.Fatal(e.Start(address))
}
