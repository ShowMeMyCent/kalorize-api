package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/services"
	"kalorize-api/utils"
	"net/http"
	"strconv"
)

type GymController struct {
	gymService services.GymService
	validate   validator.Validate
}

// NewGymController creates a new GymController
func NewGymController(db *gorm.DB) GymController {
	service := services.NewGymService(db)
	controller := GymController{
		gymService: service,
		validate:   *validator.New(),
	}
	return controller
}

// GetAllGyms handles the request to get all gyms
func (ctrl *GymController) GetAllGyms(c echo.Context) error {
	response := ctrl.gymService.GetAllGyms()
	return c.JSON(response.StatusCode, response)
}

// GetGymByName handles the request to get a gym by name
func (ctrl *GymController) GetGymByName(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Name query parameter is required",
			Data:       nil,
		})
	}

	response := ctrl.gymService.GetGymByName(name)
	return c.JSON(response.StatusCode, response)
}

// GetGymById handles the request to get a gym by ID
func (ctrl *GymController) GetGymById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("idGym"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid gym ID",
			Data:       nil,
		})
	}

	response := ctrl.gymService.GetGymById(id)
	return c.JSON(response.StatusCode, response)
}
