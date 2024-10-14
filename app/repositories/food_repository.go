package repositories

import (
	"kalorize-api/app/models"
	"kalorize-api/utils"

	"gorm.io/gorm"
)

type dbMakanan struct {
	Conn *gorm.DB
}

func (db *dbMakanan) GetAllMakanan() ([]models.Makanan, error) {
	var makanans []models.Makanan
	err := db.Conn.Find(&makanans).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return makanans, err
}

func (db *dbMakanan) GetMakananById(id int) (models.Makanan, error) {
	var makanan models.Makanan
	err := db.Conn.Where("id_makanan = ?", id).First(&makanan).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
	}
	return makanan, err
}

type MakananRepository interface {
	GetAllMakanan() ([]models.Makanan, error)
	GetMakananById(id int) (models.Makanan, error)
}

func NewDBMakananRepository(conn *gorm.DB) *dbMakanan {
	return &dbMakanan{Conn: conn}
}
