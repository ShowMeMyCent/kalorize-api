package repositories

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
)

type dbKodeGym struct {
	Conn *gorm.DB
}

type KodeGymRepository interface {
	GetKodeGymByKode(kode string) (models.GymCode, error)
	CreateNewKodeGym(kodeGym models.GymCode) error
	UpdateKodeGym(kodeGym models.GymCode) error
	DeleteKodeGym(idKodeGym int) error
	GetKodeGymById(idKodeGym int) (models.GymCode, error)
	GetIDFromKode(kode string) (int, error)
}

func (db *dbKodeGym) GetKodeGymByKode(kode string) (models.GymCode, error) {
	var kodeGym models.GymCode
	err := db.Conn.Where("kode_gym = ?", kode).First(&kodeGym).Error
	if err != nil {
		kodeGym.IdKodeGym = 0
	}
	return kodeGym, err
}

func (db *dbKodeGym) GetIDFromKode(kode string) (int, error) {
	var kodeGym models.GymCode
	err := db.Conn.Where("kode_gym = ?", kode).First(&kodeGym).Error
	return kodeGym.IdKodeGym, err
}

func (db *dbKodeGym) CreateNewKodeGym(kodeGym models.GymCode) error {
	return db.Conn.Create(&kodeGym).Error
}

func (db *dbKodeGym) UpdateKodeGym(kodeGym models.GymCode) error {
	return db.Conn.Save(&kodeGym).Error
}

func (db *dbKodeGym) DeleteKodeGym(idKodeGym int) error {
	return db.Conn.Delete(&models.GymCode{}, idKodeGym).Error
}

func (db *dbKodeGym) GetKodeGymById(idKodeGym int) (models.GymCode, error) {
	var kodeGym models.GymCode
	err := db.Conn.Where("id_kode_gym = ?", idKodeGym).First(&kodeGym).Error
	return kodeGym, err
}

func NewDBKodeGymRepository(conn *gorm.DB) *dbKodeGym {
	return &dbKodeGym{Conn: conn}
}
