package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/services"
	"net/http"
	"strconv"
)

// FranchiseController struct
type FranchiseController struct {
	Service  services.FranchiseServiceImpl
	validate validator.Validate
}

func NewFranchiseController(db *gorm.DB) FranchiseController {
	service := services.NewFranchiseService(db)
	controller := FranchiseController{
		Service:  service,
		validate: *validator.New(),
	}
	return controller
}

// GetAllFranchises handles the request to get all franchises
func (fc *FranchiseController) GetAllFranchises(c echo.Context) error {
	response := fc.Service.GetAllFranchises()
	return c.JSON(response.StatusCode, response)
}

// GetFranchiseById handles the request to get a franchise by ID
func (fc *FranchiseController) GetFranchiseById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Invalid franchise ID"})
	}
	response := fc.Service.GetFranchiseById(id)
	return c.JSON(response.StatusCode, response)
}

// GetFranchiseByName handles the request to get franchises by name
func (fc *FranchiseController) GetFranchiseByName(c echo.Context) error {
	name := c.Param("name")
	response := fc.Service.GetFranchiseByName(name)
	return c.JSON(response.StatusCode, response)
}
