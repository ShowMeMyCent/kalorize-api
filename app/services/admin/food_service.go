package admin

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories/admin"
	"kalorize-api/utils"
)

type makananService struct {
	makananRepo admin.MakananRepository
}

func NewMakananService(db *gorm.DB) MakananService {
	return &makananService{makananRepo: admin.NewDBMakananRepository(db)}
}

func (service *makananService) CreateMakanan(makanan models.Makanan) utils.Response {
	var response utils.Response
	err := service.makananRepo.CreateMakanan(makanan)
	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = makanan
	return response
}

// UpdateMakanan updates a Makanan record and returns a response
func (service *makananService) UpdateMakanan(makanan models.Makanan) utils.Response {
	var response utils.Response

	// Validate input
	if makanan.IdMakanan == 0 {
		response.StatusCode = 400
		response.Messages = "ID Makanan tidak boleh kosong"
		response.Data = nil
		return response
	}

	// Update Makanan
	err := service.makananRepo.UpdateMakanan(makanan)
	if err != nil {
		errCode, errMessage := utils.ErrorCodeAndMessage(err)
		response.StatusCode = errCode
		response.Messages = errMessage
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Makanan berhasil diperbarui"
	response.Data = makanan
	return response
}

type MakananService interface {
	CreateMakanan(makanan models.Makanan) utils.Response
	UpdateMakanan(makanan models.Makanan) utils.Response
}
