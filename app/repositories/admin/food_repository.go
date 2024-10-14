package admin

import (
	"kalorize-api/app/models"
	"kalorize-api/utils"

	"gorm.io/gorm"
)

type dbMakanan struct {
	Conn *gorm.DB
}

func (db *dbMakanan) CreateMakanan(makanan models.Makanan) error {
	err := db.Conn.Create(&makanan).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}
	return nil
}

// UpdateMakanan updates an existing Makanan record in the database
func (repo *dbMakanan) UpdateMakanan(makanan models.Makanan) error {
	err := repo.Conn.Model(&models.Makanan{}).Where("id_makanan = ?", makanan.IdMakanan).Updates(makanan).Error

	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

type MakananRepository interface {
	CreateMakanan(makanan models.Makanan) error
	UpdateMakanan(makanan models.Makanan) error
}

func NewDBMakananRepository(conn *gorm.DB) *dbMakanan {
	return &dbMakanan{Conn: conn}
}
