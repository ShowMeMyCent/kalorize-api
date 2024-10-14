package models

import (
	"time"
)

type UsedCode struct {
	IdUseCode int       `json:"id_use_code" gorm:"column:id_use_code;primaryKey;autoIncrement"`
	IdGym     int       `json:"id_gym" gorm:"column:id_gym;index"`
	IdUser    int       `json:"id_user" gorm:"column:id_user;index"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at;type:datetime;"`
}

func (UsedCode) TableName() string {
	return "used_codes"
}
