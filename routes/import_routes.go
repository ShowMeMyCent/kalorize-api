package routes

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteImportDatabase(apiv1 *echo.Group, db *gorm.DB) {
	apiv1.POST("/import", func(c echo.Context) error {
		sqlFilePath := "kalorize.sql"
		sqlContent, err := os.ReadFile(sqlFilePath)
		if err != nil {
			log.Fatalf("Error reading SQL file: %v", err)
		}
		execute, err := db.DB()
		if err != nil {
			log.Fatalf("Error executing SQL file: %v", err)
		}
		success, err := execute.Exec(string(sqlContent))
		if err != nil {
			log.Fatalf("Error executing SQL file: %v", err)
		}
		return c.JSON(200, success)

	})
}
