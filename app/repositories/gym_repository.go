package repositories

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/utils"
)

type dbGym struct {
	Conn *gorm.DB
}

func (db *dbGym) GetGym() ([]models.Gym, error) {
	var gym []models.Gym
	err := db.Conn.Find(&gym).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return gym, err
}

func (db *dbGym) GetGymByGymName(gymName string) (models.Gym, error) {
	var gym models.Gym
	err := db.Conn.Where("nama LIKE ?", "%"+gymName+"%").First(&gym).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return models.Gym{}, err
	}
	return gym, nil
}

func (db *dbGym) GetGymById(idGym int) (models.Gym, error) {
	var gym models.Gym
	err := db.Conn.Where("id_gym = ?", idGym).First(&gym).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return models.Gym{}, err
	}
	return gym, err
}

type GymRepository interface {
	GetGym() ([]models.Gym, error)
	GetGymByGymName(gymName string) (models.Gym, error)
	GetGymById(idGym int) (models.Gym, error)
}

func NewDBGymRepository(conn *gorm.DB) *dbGym {
	return &dbGym{Conn: conn}
}
