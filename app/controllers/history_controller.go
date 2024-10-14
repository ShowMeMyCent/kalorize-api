// history_controller.go
package controllers

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	service "kalorize-api/app/services"
	"kalorize-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HistoryController struct {
	service  service.HistoryService
	validate validator.Validate
}

type HistoryService interface {
	GetAllHistories() utils.Response
	GetHistoryById(id int) utils.Response
	CreateHistory(history models.History) utils.Response
	UpdateHistory(history models.History) utils.Response
	DeleteHistory(id int) utils.Response
	GetHistoryCSV(c echo.Context) utils.Response // Optional: If you want to export histories as CSV
}

func NewHistoryController(db *gorm.DB) HistoryController {
	service := service.NewHistoryService(db)
	controller := HistoryController{
		service:  service,
		validate: *validator.New(),
	}
	return controller
}

func (c *HistoryController) GetAllHistories(ctx echo.Context) error {
	response := c.service.GetAllHistories()
	return ctx.JSON(http.StatusOK, response)
}

func (c *HistoryController) GetHistoryById(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("idHistory"))
	response := c.service.GetHistoryById(id)
	return ctx.JSON(http.StatusOK, response)
}

func (c *HistoryController) CreateHistory(ctx echo.Context) error {
	history := new(models.History)
	if err := ctx.Bind(history); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Messages: "Invalid input", Data: err.Error()})
	}
	response := c.service.CreateHistory(history)
	return ctx.JSON(http.StatusCreated, response)
}

func (c *HistoryController) UpdateHistory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("idHistory"))
	history := new(models.History)
	if err := ctx.Bind(history); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.Response{StatusCode: http.StatusBadRequest, Messages: "Invalid input", Data: err.Error()})
	}
	history.IdHistory = id
	response := c.service.UpdateHistory(history)
	return ctx.JSON(http.StatusOK, response)
}

func (c *HistoryController) DeleteHistory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("idHistory"))
	response := c.service.DeleteHistory(id)
	return ctx.JSON(http.StatusOK, response)
}
