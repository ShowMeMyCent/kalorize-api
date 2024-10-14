package services

import (
	"gorm.io/gorm"
	"kalorize-api/app/repositories"
	"kalorize-api/utils"
)

type FranchiseServiceImpl interface {
	GetAllFranchises() utils.Response
	GetFranchiseById(id int) utils.Response
	GetFranchiseByName(name string) utils.Response
}

type franchiseService struct {
	repo repositories.FranchiseRepository
}

func NewFranchiseService(db *gorm.DB) FranchiseServiceImpl {
	return &franchiseService{repo: repositories.NewDbFranchise(db)}
}

func (s *franchiseService) GetAllFranchises() utils.Response {
	var response utils.Response
	result, err := s.repo.GetAllFranchises()

	response = utils.BuildResponse(err)
	response.Data = result
	return response
}

func (s *franchiseService) GetFranchiseById(id int) utils.Response {
	var response utils.Response
	result, err := s.repo.GetFranchiseById(id)

	response = utils.BuildResponse(err)
	response.Data = result
	return response
}

func (s *franchiseService) GetFranchiseByName(name string) utils.Response {
	var response utils.Response
	result, err := s.repo.GetFranchiseByName(name)

	response = utils.BuildResponse(err)
	response.Data = result
	return response
}
