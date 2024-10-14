package services

import (
	"gorm.io/gorm"
	"kalorize-api/app/repositories"
	"kalorize-api/utils"
)

type gymService struct {
	repo repositories.GymRepository
}

type GymService interface {
	GetAllGyms() utils.Response
	GetGymByName(name string) utils.Response
	GetGymById(id int) utils.Response
}

func NewGymService(db *gorm.DB) GymService {
	return &gymService{repo: repositories.NewDBGymRepository(db)}
}

func (s *gymService) GetAllGyms() utils.Response {
	var response utils.Response
	gyms, err := s.repo.GetGym()

	response = utils.BuildResponse(err)
	response.Data = gyms

	return response
}

func (s *gymService) GetGymByName(name string) utils.Response {
	var response utils.Response
	gym, err := s.repo.GetGymByGymName(name)
	response = utils.BuildResponse(err)
	response.Data = gym

	return response
}

func (s *gymService) GetGymById(id int) utils.Response {
	var response utils.Response
	gym, err := s.repo.GetGymById(id)

	response = utils.BuildResponse(err)
	response.Data = gym

	return response
}
