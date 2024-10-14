package admin

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/services/admin"
	"net/http"
)

type MakananController struct {
	makananService admin.MakananService
	validate       validator.Validate
}

func NewMakananController(db *gorm.DB) MakananController {
	service := admin.NewMakananService(db)
	controller := MakananController{
		makananService: service,
		validate:       *validator.New(),
	}
	return controller
}

func (controller *MakananController) CreateMakanan(c echo.Context) error {
	var makanan models.Makanan
	if err := c.Bind(&makanan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Invalid request"})
	}

	if err := controller.validate.Struct(makanan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Validation failed"})
	}

	response := controller.makananService.CreateMakanan(makanan)
	return c.JSON(response.StatusCode, response)
}

func (controller *MakananController) UpdateMakanan(c echo.Context) error {
	var makanan models.Makanan
	if err := c.Bind(&makanan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Invalid request"})
	}

	if err := controller.validate.Struct(makanan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Validation failed"})
	}

	response := controller.makananService.UpdateMakanan(makanan)
	return c.JSON(response.StatusCode, response)
}
