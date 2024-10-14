package services

import (
	"encoding/csv"
	"gorm.io/gorm"
	"kalorize-api/app/repositories"
	"kalorize-api/formatter"
	"kalorize-api/utils"

	"github.com/labstack/echo/v4"
)

type makananService struct {
	makananRepo repositories.MakananRepository
	user        repositories.UserRepository
}

func NewMakananService(db *gorm.DB) MakananService {
	return &makananService{makananRepo: repositories.NewDBMakananRepository(db)}
}

func (service *makananService) GetAllMakanan() utils.Response {
	var response utils.Response
	makanan, err := service.makananRepo.GetAllMakanan()

	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}
	var formattedMakanan []formatter.MakananFormat
	for i := range makanan {
		formattedMakanan = append(formattedMakanan, formatter.FormatterMakananIndo(makanan[i]))
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = formattedMakanan
	return response
}

func (service *makananService) GetMakananCSV(c echo.Context) utils.Response {
	// response is .csv file generator
	wr := csv.NewWriter(c.Response())
	var response utils.Response
	makanan, err := service.makananRepo.GetAllMakanan()
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Internal server error"
		response.Data = nil
		return response
	}
	formattedMultiMakanan := formatter.FormatterMakananToMultiDimentionalArray(makanan)
	wr.WriteAll(formattedMultiMakanan)

	response.StatusCode = 200
	response.Messages = "success"
	return response
}

func (service *makananService) GetMakananById(id int) utils.Response {
	var response utils.Response
	makanan, err := service.makananRepo.GetMakananById(id)
	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	formattedMakanan := formatter.FormatterMakananIndo(makanan)

	response.StatusCode = 200
	response.Messages = "success"
	response.Data = formattedMakanan
	return response
}

type MakananService interface {
	GetAllMakanan() utils.Response
	GetMakananById(id int) utils.Response
	GetMakananCSV(c echo.Context) utils.Response
}
