package models

import "github.com/google/uuid"

type Token struct {
	IdToken      uuid.UUID `json:"id_item" gorm:"column:id_token;type:char(36);primary_key"`
	Email        string    `gorm:"index"`
	AccessToken  string    `json:"access_token" gorm:" column:access_token;type:mediumtext;"`
	RefreshToken string    `json:"refresh_token" gorm:" column:refresh_token;type:mediumtext;"`
}

func (t *Token) TableName() string {
	return "tokens"
}
