package routes

import (
	"kalorize-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteQuestionnaire(apiv1 *echo.Group, db *gorm.DB) {
	questionnaireController := controllers.NewQuestionnaireController(db)
	apiv1.PUT("/questionnaire", questionnaireController.FillQuestionnaire)
}
