package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/services"
	"strconv"
)

type MakananController struct {
	makananService services.MakananService
	validate       validator.Validate
}

func NewMakananController(db *gorm.DB) MakananController {
	service := services.NewMakananService(db)
	controller := MakananController{
		makananService: service,
		validate:       *validator.New(),
	}
	return controller
}

func (controller *MakananController) GetAllMakanan(c echo.Context) error {
	response := controller.makananService.GetAllMakanan()
	return c.JSON(response.StatusCode, response)
}

func (controller *MakananController) GetMakananById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("makananId"))
	response := controller.makananService.GetMakananById(id)
	return c.JSON(response.StatusCode, response)
}

func (controller *MakananController) GetMakananCSV(c echo.Context) error {
	response := controller.makananService.GetMakananCSV(c)
	return c.JSON(response.StatusCode, response)
}
