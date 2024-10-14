package admin

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories"
	"kalorize-api/app/repositories/admin"
	"kalorize-api/utils"
	"strings"
)

type authService struct {
	authRepo   admin.UserRepository
	tokenRepo  repositories.TokenRepository
	signingKey string
}

func (service *authService) Login(email, password string) utils.Response {
	var response utils.Response
	if email == "" || password == "" {
		response.StatusCode = 400
		response.Messages = "email dan password tidak boleh kosong"
		response.Data = nil
		return response
	}

	if !utils.IsEmailValid(email) {
		response.StatusCode = 400
		response.Messages = "Email kamu tidak valid"
		response.Data = nil
		return response
	}

	user, err := service.authRepo.GetUserByEmail(email)
	if err != nil {
		response.StatusCode = 401
		response.Messages = "Email kamu belum terdaftar"
		response.Data = nil
		return response
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		response.StatusCode = 401
		response.Messages = "Password kamu salah"
		response.Data = nil
		return response
	}
	AccessToken, err := utils.GenerateJWTAccessToken(user.IdUser, user.Fullname, user.Email, service.signingKey)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	refreshToken, err := utils.GenerateJWTRefreshToken(user.IdUser, user.Fullname, user.Email, service.signingKey)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	token := models.Token{
		IdToken:      uuid.New(),
		AccessToken:  AccessToken,
		RefreshToken: refreshToken,
		Email:        user.Email,
	}
	err = service.tokenRepo.CreateNewToken(token)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token creation failed"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "success"
	response.Data = map[string]interface{}{
		"accessToken":  AccessToken,
		"refreshToken": refreshToken,
		"role":         user.Role,
		"userId":       user.IdUser,
	}
	return response
}

func (service *authService) GetLoggedInUser(bearerToken string) utils.Response {
	var response utils.Response
	var firstname, lastname string
	id, err := utils.ParseDataId(bearerToken)
	if id != 0 && err == nil {
		user, err := service.authRepo.GetUserById(id)
		if err != nil {
			response.StatusCode = 500
			response.Messages = "User tidak ditemukan"
			response.Data = nil
			return response
		}
		names := strings.Split(user.Fullname, " ")
		if len(names) == 1 {
			firstname = names[0]
			lastname = names[0]
		} else {
			firstname = names[0]
			lastname = names[len(names)-1]
		}
		response.Data = map[string]interface{}{
			"idUser":       user.IdUser,
			"firstName":    firstname,
			"lastName":     lastname,
			"email":        user.Email,
			"jenisKelamin": user.JenisKelamin,
			"umur":         user.Umur,
			"role":         user.Role,
			"foto":         user.FotoUrl,
			"noTelepon":    user.NoTelepon,
		}
		response.StatusCode = 200
		response.Messages = "success"
		return response
	} else {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}
}

func (service *authService) Logout(bearerToken string) utils.Response {
	var response utils.Response
	_, err := utils.ParseDataEmail(bearerToken)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token parsing failed"
		response.Data = nil
		return response
	}

	err = service.tokenRepo.DeleteToken(bearerToken)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token deletion failed"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = nil
	return response
}

func (service *authService) Refresh(refreshToken string) utils.Response {
	var response utils.Response
	userId, err := utils.ParseDataId(refreshToken)
	if userId == 0 || err != nil {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}
	user, err := service.authRepo.GetUserById(userId)
	if err != nil {
		response.StatusCode = 401
		response.Messages = "User tidak ditemukan"
		response.Data = nil
		return response
	}
	AccessToken, err := utils.GenerateJWTAccessToken(user.IdUser, user.Fullname, user.Email, service.signingKey)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	refreshToken, err = utils.GenerateJWTRefreshToken(user.IdUser, user.Fullname, user.Email, service.signingKey)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	token := models.Token{
		IdToken:      uuid.New(),
		AccessToken:  AccessToken,
		RefreshToken: refreshToken,
		Email:        user.Email,
	}
	err = service.tokenRepo.CreateNewToken(token)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token creation failed"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = map[string]interface{}{
		"accessToken":  AccessToken,
		"refreshToken": refreshToken,
		"role":         user.Role,
		"userId":       user.IdUser,
	}
	return response
}

type AuthService interface {
	Login(username, password string) utils.Response
	GetLoggedInUser(bearerToken string) utils.Response
	Logout(bearerToken string) utils.Response
	Refresh(refreshToken string) utils.Response
}

func NewAuthService(db *gorm.DB, signingKey string) AuthService {
	return &authService{
		authRepo:   admin.NewDBUserRepository(db),
		tokenRepo:  repositories.NewDBTokenRepository(db),
		signingKey: signingKey,
	}
}
