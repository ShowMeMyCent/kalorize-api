package controllers

import (
	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/services/admin"
)

type GymCodeController struct {
	gymService admin.GymCodeService
	validate   vl.Validate
}

// NewGymController creates a new GymController
func NewGymCodeController(db *gorm.DB) GymCodeController {
	service := admin.NewGymCodeService(db)
	controller := GymCodeController{
		gymService: service,
		validate:   *vl.New(),
	}
	return controller
}

func (controller *GymCodeController) GenerateGymToken(c echo.Context) error {
	type payload struct {
		Uid int `json:"uid" validate:"required"`
	}
	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}
	response := controller.gymService.GenerateKodeGym(payloadValidator.Uid)
	return c.JSON(response.StatusCode, response)
}
