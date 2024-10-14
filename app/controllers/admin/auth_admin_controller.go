package admin

import (
	services "kalorize-api/app/services/admin"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthController struct {
	authService services.AuthService
	validate    vl.Validate
}

func NewAuthController(db *gorm.DB, signingKey string) AuthController {
	service := services.NewAuthService(db, signingKey)
	controller := AuthController{
		authService: service,
		validate:    *vl.New(),
	}
	return controller
}

func (controller *AuthController) Refresh(c echo.Context) error {
	type payload struct {
		RefreshToken string `json:"refreshToken"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.authService.Refresh(payloadValidator.RefreshToken)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) Login(c echo.Context) error {
	type payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.authService.Login(payloadValidator.Email, payloadValidator.Password)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) GetUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	response := controller.authService.GetLoggedInUser(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) Logout(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	response := controller.authService.Logout(token)
	return c.JSON(response.StatusCode, response)
}
