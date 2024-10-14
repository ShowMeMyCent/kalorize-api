package admin

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	repositories "kalorize-api/app/repositories/admin"
	services "kalorize-api/app/services/admin"
	"kalorize-api/utils"
	"net/http"
	"strconv"
)

type UserAdminController struct {
	userService services.UserService
	validate    validator.Validate
}

func NewUserAdminController(db *gorm.DB) *UserAdminController {
	service := services.NewUserService(repositories.NewDBUserRepository(db))
	return &UserAdminController{
		userService: service,
		validate:    *validator.New(),
	}
}

// GetAllUsers handles the request to get all users.
func (ctrl *UserAdminController) GetAllUsers(c echo.Context) error {
	response := ctrl.userService.GetAllUsers()
	return c.JSON(response.StatusCode, response)
}

// GetUserById handles the request to get a user by ID.
func (ctrl *UserAdminController) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid user ID",
			Data:       nil,
		})
	}

	response := ctrl.userService.GetUserById(id)
	return c.JSON(response.StatusCode, response)
}

// CreateUser handles the request to create a new user.
func (ctrl *UserAdminController) CreateUser(c echo.Context) error {
	var user models.UserAdmin
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid request payload",
			Data:       nil,
		})
	}

	if err := ctrl.validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Validation failed",
			Data:       err.Error(),
		})
	}

	response := ctrl.userService.CreateUser(user)
	return c.JSON(response.StatusCode, response)
}

// UpdateUser handles the request to update an existing user.
func (ctrl *UserAdminController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid user ID",
			Data:       nil,
		})
	}

	var user models.UserAdmin
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid request payload",
			Data:       nil,
		})
	}

	user.IdUser = id

	if err := ctrl.validate.Struct(user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Validation failed",
			Data:       err.Error(),
		})
	}

	response := ctrl.userService.UpdateUser(user)
	return c.JSON(response.StatusCode, response)
}

// DeleteUser handles the request to delete a user by ID.
func (ctrl *UserAdminController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Invalid user ID",
			Data:       nil,
		})
	}

	response := ctrl.userService.DeleteUser(id)
	return c.JSON(response.StatusCode, response)
}

// GetUserByEmail handles the request to get a user by email.
func (ctrl *UserAdminController) GetUserByEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, utils.Response{
			StatusCode: http.StatusBadRequest,
			Messages:   "Email parameter is required",
			Data:       nil,
		})
	}

	response := ctrl.userService.GetUserByEmail(email)
	return c.JSON(response.StatusCode, response)
}
