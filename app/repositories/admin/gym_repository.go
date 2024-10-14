package admin

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/utils"
)

type dbGym struct {
	Conn *gorm.DB
}

func (db *dbGym) CreateNewGym(gym models.Gym) error {
	err := db.Conn.Create(&gym).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}
	return nil
}

func (db *dbGym) UpdateGym(gym models.Gym) error {
	// Attempt to update the gym record where id_gym matches
	result := db.Conn.Where("id_gym = ?", gym.IdGym).Updates(&gym)

	// Check if the query resulted in any rows being updated
	if result.RowsAffected == 0 {
		// No gym found with the given id_gym, return a "not found" error
		return utils.Error(nil, utils.ErrDatabase, "Gym not found")
	}

	// If there is a database error, return it
	if result.Error != nil {
		return utils.Error(result.Error, utils.ErrDatabase, result.Error.Error())
	}

	// If everything went well, return nil
	return nil
}

func (db *dbGym) DeleteGym(idGym int) error {
	err := db.Conn.Delete(&models.Gym{}, idGym).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}
	return nil
}

type GymRepository interface {
	CreateNewGym(gym models.Gym) error
	UpdateGym(gym models.Gym) error
	DeleteGym(idGym int) error
}

func NewDBGymRepository(conn *gorm.DB) *dbGym {
	return &dbGym{Conn: conn}
}
