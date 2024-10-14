package controllers

import (
	"kalorize-api/app/models" // Make sure to import the models package
	"kalorize-api/app/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TokenController struct {
	tokenService services.TokenService
	validate     *validator.Validate
}

func NewTokenController(db *gorm.DB) TokenController {
	service := services.NewTokenService(db)
	controller := TokenController{
		tokenService: service,
		validate:     validator.New(),
	}
	return controller
}

func (controller *TokenController) GetAllTokens(c echo.Context) error {
	tokens, err := controller.tokenService.GetAllTokens()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tokens)
}

func (controller *TokenController) CreateToken(c echo.Context) error {
	var token models.Token
	if err := c.Bind(&token); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := controller.validate.Struct(token); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := controller.tokenService.CreateToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, "Token created successfully")
}

func (controller *TokenController) UpdateToken(c echo.Context) error {
	var token models.Token
	if err := c.Bind(&token); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := controller.validate.Struct(token); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := controller.tokenService.UpdateToken(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Token updated successfully")
}

func (controller *TokenController) DeleteToken(c echo.Context) error {
	idToken := c.Param("idToken")
	err := controller.tokenService.DeleteToken(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Token deleted successfully")
}
