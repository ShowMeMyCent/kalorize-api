package repositories

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/utils"
)

type FranchiseRepository interface {
	GetAllFranchises() ([]models.Franchise, error)
	GetFranchiseById(id int) (models.Franchise, error)
	GetFranchiseByName(name string) ([]models.Franchise, error)
}

type dbFranchise struct {
	Conn *gorm.DB
}

func NewDbFranchise(conn *gorm.DB) *dbFranchise {
	return &dbFranchise{Conn: conn}
}

func (db *dbFranchise) GetAllFranchises() ([]models.Franchise, error) {
	var franchises []models.Franchise
	err := db.Conn.Preload("Makanan").Find(&franchises).Error
	return franchises, err
}

func (r *dbFranchise) GetFranchiseById(id int) (models.Franchise, error) {
	var franchise models.Franchise

	err := r.Conn.Preload("Makanan").First(&franchise, id).Error
	if err != nil {
		return franchise, err
	}

	return franchise, nil
}

func (db *dbFranchise) GetFranchiseByName(name string) ([]models.Franchise, error) {
	var franchises []models.Franchise
	err := db.Conn.Preload("Makanan").Where("nama_franchise LIKE ?", "%"+name+"%").Find(&franchises).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}

	return franchises, err
}
