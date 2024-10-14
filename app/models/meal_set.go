package models

import (
	"time"
)

type MealSet struct {
	IdUser         int       `json:"id_user" gorm:"column:id_user;type:char(36);"`
	IdMakanan      int       `json:"id_makanan" gorm:"column:id_makanan;type:int(36);"`
	JumlahKalori   int       `json:"jumlah_kalori" gorm:"column:jumlah_kalori;type:int;"`
	JumlahProtein  int       `json:"jumlah_protein" gorm:"column:jumlah_protein;type:int;"`
	TanggalMealSet time.Time `json:"tanggal_meal_set" gorm:"column:tanggal_meal_set;type:timestamp;"`
}
