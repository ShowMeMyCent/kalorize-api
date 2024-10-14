package admin

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/app/repositories"
	"kalorize-api/utils"
)

type gymCodeService struct {
	gymRepo     repositories.GymRepository
	gymKode     repositories.KodeGymRepository
	gymUsedCode repositories.UsedCodeRepository
}

func (gymCode *gymCodeService) GenerateKodeGym(idGym int) utils.Response {
	var response utils.Response
	Gym, err := gymCode.gymRepo.GetGymById(idGym)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Failed to get gym"
		response.Data = nil
		return response
	}
	var kodeGym = models.GymCode{
		IdKodeGym:   0, // Initialize with 0 or any other default value
		KodeGym:     utils.GenerateKodeGym(Gym.NamaGym),
		IdGym:       idGym,
		ExpiredTime: utils.GetExpiredTime(),
	}

	err = gymCode.gymKode.CreateNewKodeGym(kodeGym)

	if err != nil {
		response.StatusCode = 500
		response.Messages = "Failed to generate kode gym"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Success"
	response.Data = kodeGym
	return response
}

type GymCodeService interface {
	GenerateKodeGym(idGym int) utils.Response
}

func NewGymCodeService(db *gorm.DB) GymCodeService {
	return &gymCodeService{
		gymRepo:     repositories.NewDBGymRepository(db),
		gymKode:     repositories.NewDBKodeGymRepository(db),
		gymUsedCode: repositories.NewDBUsedCodeRepository(db),
	}
}
