package admin

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/services/admin"
	"net/http"
	"strconv"
)

// FranchiseController struct
type FranchiseController struct {
	Service  admin.FranchiseServiceImpl
	validate validator.Validate
}

func NewFranchiseController(db *gorm.DB) FranchiseController {
	service := admin.NewFranchiseService(db)
	controller := FranchiseController{
		Service:  service,
		validate: *validator.New(),
	}
	return controller
}

// CreateFranchise handles the request to create a new franchise
func (fc *FranchiseController) CreateFranchiseWithMakanan(c echo.Context) error {
	var request models.Franchise

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Invalid request payload"})
	}

	response := fc.Service.CreateFranchiseWithMakanan(request)

	return c.JSON(response.StatusCode, response)
}

// UpdateFranchise handles the request to update an existing franchise
func (fc *FranchiseController) UpdateFranchise(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Invalid franchise ID"})
	}
	var franchise models.Franchise
	if err := c.Bind(&franchise); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Invalid request payload"})
	}
	franchise.IdFranchise = id

	response := fc.Service.UpdateFranchise(franchise)
	return c.JSON(response.StatusCode, response)
}

// DeleteFranchise handles the request to delete a franchise
func (fc *FranchiseController) DeleteFranchise(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Invalid franchise ID"})
	}
	if err := fc.Service.DeleteFranchise(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"statusCode": http.StatusBadRequest, "message": "Failed to delete franchise"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Franchise deleted successfully"})
}
