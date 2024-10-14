// history_repository.go
package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/utils"
)

type HistoryRepository interface {
	FindAll() ([]models.History, error)
	FindById(id int) (*models.History, error)
	Create(history *models.History) (*models.History, error)
	Update(history *models.History) (*models.History, error)
	Delete(id int) error
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db}
}

func (r *historyRepository) FindAll() ([]models.History, error) {
	var histories []models.History
	if err := r.db.Preload("Breakfast").Preload("Lunch").Preload("Dinner").Find(&histories).Error; err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return histories, nil
}

func (r *historyRepository) FindById(id int) (*models.History, error) {
	var history models.History
	if err := r.db.Preload("Breakfast").Preload("Lunch").Preload("Dinner").First(&history, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return &history, nil
}

func (r *historyRepository) Create(history *models.History) (*models.History, error) {
	if err := r.db.Create(history).Error; err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return history, nil
}

func (r *historyRepository) Update(history *models.History) (*models.History, error) {
	var existingHistory models.History

	// Find the existing record by primary key
	if err := r.db.First(&existingHistory, history.IdHistory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = utils.Error(err, utils.ErrDatabase, err.Error())
			return nil, fmt.Errorf("record with id_history %d not found", history.IdHistory)
		}
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}

	// Update the record
	if err := r.db.Model(&existingHistory).Updates(history).Error; err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}

	return &existingHistory, nil
}

func (r *historyRepository) Delete(id int) error {
	if err := r.db.Delete(&models.History{}, id).Error; err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}
	return nil
}
