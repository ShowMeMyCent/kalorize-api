package services

import (
	"kalorize-api/app/models"
	"kalorize-api/app/repositories"
	"kalorize-api/utils"
)

type userService struct {
	repo repositories.UserRepository
}

type UserService interface {
	GetAllUsers() utils.Response
	GetUserById(id int) utils.Response
	CreateUser(user models.User) utils.Response
	UpdateUser(user models.User) utils.Response
	DeleteUser(id int) utils.Response
	GetUserByEmail(email string) utils.Response
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) GetAllUsers() utils.Response {
	var response utils.Response
	users, err := s.repo.GetUser()

	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "success"
	response.Data = users

	return response
}

func (s *userService) GetUserById(id int) utils.Response {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return utils.Response{
			StatusCode: 404,
			Messages:   "User not found",
			Data:       nil,
		}
	}
	return utils.Response{
		StatusCode: 200,
		Messages:   "User retrieved successfully",
		Data:       user,
	}
}

func (s *userService) CreateUser(user models.User) utils.Response {

	userEncrypt, _ := utils.HashPassword(user.Password)

	user.Password = userEncrypt

	err := s.repo.CreateNewUser(user)

	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to create user",
			Data:       nil,
		}
	}
	return utils.Response{
		StatusCode: 201,
		Messages:   "User created successfully",
		Data:       user,
	}
}

func (s *userService) UpdateUser(user models.User) utils.Response {
	userEncrypt, _ := utils.HashPassword(user.Password)

	user.Password = userEncrypt

	err := s.repo.UpdateUser(user)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to update user",
			Data:       nil,
		}
	}
	return utils.Response{
		StatusCode: 200,
		Messages:   "User updated successfully",
		Data:       user,
	}
}

func (s *userService) DeleteUser(id int) utils.Response {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to delete user",
			Data:       nil,
		}
	}
	return utils.Response{
		StatusCode: 200,
		Messages:   "User deleted successfully",
		Data:       nil,
	}
}

func (s *userService) GetUserByEmail(email string) utils.Response {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return utils.Response{
			StatusCode: 404,
			Messages:   "User not found",
			Data:       nil,
		}
	}
	return utils.Response{
		StatusCode: 200,
		Messages:   "User retrieved successfully",
		Data:       user,
	}
}
