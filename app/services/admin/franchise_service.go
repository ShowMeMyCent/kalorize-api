package admin

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories"
	"kalorize-api/app/repositories/admin"
	"kalorize-api/utils"
)

type FranchiseServiceImpl interface {
	CreateFranchise(franchise models.Franchise) utils.Response
	UpdateFranchise(franchise models.Franchise) utils.Response
	DeleteFranchise(id int) error
	CreateFranchiseWithMakanan(franchise models.Franchise) utils.Response
}

type franchiseService struct {
	repoAdmin admin.FranchiseRepository
	repo      repositories.FranchiseRepository
}

func NewFranchiseService(db *gorm.DB) FranchiseServiceImpl {

	return &franchiseService{repoAdmin: admin.NewDbFranchise(db),
		repo: repositories.NewDbFranchise(db)}
}

func (s *franchiseService) CreateFranchiseWithMakanan(franchise models.Franchise) utils.Response {
	var response utils.Response
	createdFranchise, err := s.repoAdmin.CreateFranchiseWithMakanan(franchise)
	response = utils.BuildResponse(err)
	response.Data = createdFranchise
	return response

}

func (s *franchiseService) CreateFranchise(franchise models.Franchise) utils.Response {
	var response utils.Response

	createdFranchise, err := s.repoAdmin.CreateFranchiseWithMakanan(franchise)
	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	getResult, err := s.repo.GetFranchiseById(createdFranchise.IdFranchise)

	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Makanan berhasil ditambahkan"
	response.Data = getResult

	return response
}

func (s *franchiseService) UpdateFranchise(franchise models.Franchise) utils.Response {
	var response utils.Response

	err := s.repoAdmin.UpdateFranchise(franchise)

	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	getMakanan, err := s.repo.GetFranchiseById(franchise.IdFranchise)

	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Makanan berhasil di update"
	response.Data = getMakanan

	return response
}

func (s *franchiseService) DeleteFranchise(id int) error {
	return s.repoAdmin.DeleteFranchise(id)
}
