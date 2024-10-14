package admin

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/utils"
)

type FranchiseRepository interface {
	CreateFranchise(franchise models.Franchise) error
	UpdateFranchise(franchise models.Franchise) error
	DeleteFranchise(id int) error
	CreateFranchiseWithMakanan(franchise models.Franchise) (models.Franchise, error)
}

type dbFranchise struct {
	Conn *gorm.DB
}

func NewDbFranchise(conn *gorm.DB) *dbFranchise {
	return &dbFranchise{Conn: conn}
}

func (db *dbFranchise) CreateFranchise(franchise models.Franchise) error {

	err := db.Conn.Create(&franchise).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

func (db *dbFranchise) UpdateFranchise(franchise models.Franchise) error {
	// Start a transaction
	tx := db.Conn.Begin()

	// Find the existing franchise by ID
	var existingFranchise models.Franchise
	if err := tx.Where("id_franchise = ?", franchise.IdFranchise).First(&existingFranchise).Error; err != nil {
		tx.Rollback()
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	// Update the franchise fields
	if err := tx.Model(&existingFranchise).Updates(map[string]interface{}{
		"nama_franchise":      franchise.NamaFranchise,
		"longitude_franchise": franchise.LongitudeFranchise,
		"latitude_franchise":  franchise.LatitudeFranchise,
		"telepon":             franchise.NoTeleponFranchise,
		"foto":                franchise.FotoFranchise,
		"email":               franchise.EmailFranchise,
		"password":            franchise.PasswordFranchise,
		"lokasi":              franchise.LokasiFranchise,
	}).Error; err != nil {
		tx.Rollback()
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	// Clear existing associations
	if err := tx.Model(&existingFranchise).Association("Makanan").Clear(); err != nil {
		tx.Rollback()
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	// Add new associations
	if err := tx.Model(&existingFranchise).Association("Makanan").Append(franchise.Makanan); err != nil {
		tx.Rollback()
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (db *dbFranchise) DeleteFranchise(id int) error {
	err := db.Conn.Delete(&models.Franchise{}, id).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

func (db *dbFranchise) CreateFranchiseWithMakanan(franchise models.Franchise) (models.Franchise, error) {
	// Start a transaction
	tx := db.Conn.Begin()

	// Create the franchise record
	if err := tx.Create(&franchise).Error; err != nil {
		tx.Rollback()
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return franchise, utils.Error(err, utils.ErrDatabase, "failed to create franchise record")
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return franchise, utils.Error(err, utils.ErrDatabase, "failed to commit transaction")
	}

	// Return the created franchise and no error
	return franchise, nil
}
