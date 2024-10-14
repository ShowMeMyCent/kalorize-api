package controllers

import (
	"kalorize-api/app/services"
	"kalorize-api/utils"

	vl "github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type QuestionnaireController struct {
	questionnaireService services.QuestionnaireService
	validate             vl.Validate
}

func NewQuestionnaireController(db *gorm.DB) QuestionnaireController {
	service := services.NewQuestionnaireService(db)
	controller := QuestionnaireController{
		questionnaireService: service,
		validate:             *vl.New(),
	}
	return controller
}

func (controller *QuestionnaireController) FillQuestionnaire(c echo.Context) error {
	type payload struct {
		IdUser       int `json:"idUser" validate:"required"`
		Umur         int `json:"umur"`
		BeratBadan   int `json:"beratBadan"`
		TinggiBadan  int `json:"tinggiBadan"`
		JenisKelamin int `json:"jenisKelamin"`
		FrekuensiGym int `json:"frekuensiGym"`
		TargetKalori int `json:"targetKalori"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	var questionnairePayload = utils.UserRequest{
		IdUser:       payloadValidator.IdUser,
		Umur:         payloadValidator.Umur,
		BeratBadan:   payloadValidator.BeratBadan,
		TinggiBadan:  payloadValidator.TinggiBadan,
		JenisKelamin: payloadValidator.JenisKelamin,
		FrekuensiGym: payloadValidator.FrekuensiGym,
		TargetKalori: payloadValidator.TargetKalori,
	}

	response := controller.questionnaireService.FillQuestionnaire(questionnairePayload)
	return c.JSON(response.StatusCode, response)
}
