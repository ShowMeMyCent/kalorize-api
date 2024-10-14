package models

import (
	"time"
)

type History struct {
	IdHistory     int       `json:"id_history" gorm:"column:id_history;primary_key;auto_increment;"`
	IdUser        int       `json:"id_user" gorm:"column:id_user;type:char(16);"`
	IdBreakfast   int       `json:"id_breakfast" gorm:"column:id_breakfast;type:int;"`
	IdLunch       int       `json:"id_lunch" gorm:"column:id_lunch;type:int;"`
	IdDinner      int       `json:"id_dinner" gorm:"column:id_dinner;type:int;"`
	TotalProtein  int       `json:"total_protein" gorm:"column:total_protein;type:int(11);"`
	TotalKalori   int       `json:"total_kalori" gorm:"column:total_kalori;type:int(11);"`
	TanggalDibuat time.Time `json:"tanggal_dibuat" gorm:"column:tanggal_dibuat;type:datetime;"`
	Breakfast     Makanan   `json:"breakfast" gorm:"foreignKey:IdBreakfast;references:IdMakanan"`
	Lunch         Makanan   `json:"lunch" gorm:"foreignKey:IdLunch;references:IdMakanan"`
	Dinner        Makanan   `json:"dinner" gorm:"foreignKey:IdDinner;references:IdMakanan"`
}
