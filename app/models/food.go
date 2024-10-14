package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Makanan struct {
	IdMakanan     int    `json:"id_makanan" gorm:"column:id_makanan;primary_key;autoIncrement"`
	Nama          string `json:"nama" gorm:"column:nama;type:varchar(255);"`
	Foto          string `json:"foto" gorm:"column:foto;type:varchar(255);"`
	Kalori        int    `json:"kalori" gorm:"column:kalori;type:int;"`
	Protein       int    `json:"protein" gorm:"column:protein;type:int;"`
	Bahan         string `json:"bahan" gorm:"column:bahan;type:text;"`
	CookingStep   string `json:"cooking_step" gorm:"column:cooking_step;type:text;"`
	ListFranchise string `json:"franchise" gorm:"column:franchise;type:text;"`
}

func (m *Makanan) TableName() string {
	return "makanans"
}

type TimeWrapper struct {
	time.Time
}

func (tw *TimeWrapper) Scan(value interface{}) error {
	if value == nil {
		tw.Time = time.Time{}
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan time")
	}

	tw.Time = t
	return nil
}

func (tw TimeWrapper) Value() (driver.Value, error) {
	return tw.Time, nil
}
