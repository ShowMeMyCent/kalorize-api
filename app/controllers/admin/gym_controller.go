package admin

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/services/admin"
	"kalorize-api/utils"
	"net/http"
	"strconv"
)

type GymController struct {
	gymService admin.GymService
	validate   validator.Validate
}

// NewGymController creates a new GymController
func NewGymController(db *gorm.DB) GymController {
	service := admin.NewGymService(db)
	controller := GymController{
		gymService: service,
		validate:   *validator.New(),
	}
	return controller
}

// CreateGym handles the request to create a new gym
func (ctrl *GymController) CreateGym(c echo.Context) error {
	var gym models.Gym
	if err := c.Bind(&gym); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid request body",
			Data:       nil,
		})
	}

	// Validate the gym model
	if err := ctrl.validate.Struct(gym); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Validation error",
			Data:       err.Error(),
		})
	}

	response := ctrl.gymService.CreateGym(gym)
	return c.JSON(response.StatusCode, response)
}

// UpdateGym handles the request to update an existing gym
func (ctrl *GymController) UpdateGym(c echo.Context) error {
	var gym models.Gym
	if err := c.Bind(&gym); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid request body",
			Data:       nil,
		})
	}

	// Validate the gym model
	if err := ctrl.validate.Struct(gym); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Validation error",
			Data:       err.Error(),
		})
	}

	response := ctrl.gymService.UpdateGym(gym)
	return c.JSON(response.StatusCode, response)
}

// DeleteGym handles the request to delete a gym by ID
func (ctrl *GymController) DeleteGym(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("idGym"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid gym ID",
			Data:       nil,
		})
	}

	response := ctrl.gymService.DeleteGym(id)
	return c.JSON(response.StatusCode, response)
}
