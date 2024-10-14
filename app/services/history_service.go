package services

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories"
	"kalorize-api/utils"
)

type HistoryService interface {
	GetAllHistories() utils.Response
	GetHistoryById(id int) utils.Response
	CreateHistory(history *models.History) utils.Response
	UpdateHistory(history *models.History) utils.Response
	DeleteHistory(id int) utils.Response
}

type historyService struct {
	repo repositories.HistoryRepository
}

func NewHistoryService(db *gorm.DB) HistoryService {
	return &historyService{repo: repositories.NewHistoryRepository(db)}
}

func (s *historyService) GetAllHistories() utils.Response {
	var response utils.Response

	histories, err := s.repo.FindAll()
	response = utils.BuildResponse(err)
	response.Data = histories

	return response
}

func (s *historyService) GetHistoryById(id int) utils.Response {
	var response utils.Response

	history, err := s.repo.FindById(id)
	response = utils.BuildResponse(err)

	if history == nil {
		response.StatusCode = 404
		response.Messages = "History not found"
		response.Data = nil
		return response
	}

	response.Data = history

	return response
}

func (s *historyService) CreateHistory(history *models.History) utils.Response {
	var response utils.Response

	newHistory, err := s.repo.Create(history)
	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	result, err := s.repo.FindById(newHistory.IdHistory)

	response = utils.BuildResponse(err)

	response.Data = result

	return response
}

func (s *historyService) UpdateHistory(history *models.History) utils.Response {
	var response utils.Response

	updatedHistory, err := s.repo.Update(history)

	response = utils.BuildResponse(err)

	result, err := s.repo.FindById(updatedHistory.IdHistory)

	response = utils.BuildResponse(err)

	response.Data = result

	return response
}

func (s *historyService) DeleteHistory(id int) utils.Response {
	var response utils.Response

	err := s.repo.Delete(id)
	response = utils.BuildResponse(err)

	response.Data = nil
	return response
}
