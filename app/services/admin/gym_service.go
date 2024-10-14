package admin

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories/admin"
	"kalorize-api/utils"
)

type gymService struct {
	repo admin.GymRepository
}

type GymService interface {
	CreateGym(gym models.Gym) utils.Response
	UpdateGym(gym models.Gym) utils.Response
	DeleteGym(id int) utils.Response
}

func NewGymService(db *gorm.DB) GymService {
	return &gymService{repo: admin.NewDBGymRepository(db)}
}

func (s *gymService) CreateGym(gym models.Gym) utils.Response {
	var response utils.Response
	err := s.repo.CreateNewGym(gym)

	response = utils.BuildResponse(err)
	response.Data = gym

	return response
}

func (s *gymService) UpdateGym(gym models.Gym) utils.Response {
	var response utils.Response
	err := s.repo.UpdateGym(gym)

	response = utils.BuildResponse(err)
	if err != nil {
		response.Data = gym
	}

	return response
}

func (s *gymService) DeleteGym(id int) utils.Response {
	var response utils.Response
	err := s.repo.DeleteGym(id)

	response = utils.BuildResponse(err)
	response.Data = nil

	return response
}
