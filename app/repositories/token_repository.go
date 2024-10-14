package repositories

import (
	"kalorize-api/app/models"
	"kalorize-api/utils"

	"gorm.io/gorm"
)

type DbToken struct {
	Conn *gorm.DB
}

func (db *DbToken) GetToken() ([]models.Token, error) {
	var token []models.Token
	err := db.Conn.Find(&token).Error

	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}

	return token, err
}

func (db *DbToken) GetTokenByUserEmail(email string, accessToken string) (*models.Token, error) {
	var token models.Token
	err := db.Conn.Where("email = ? AND access_token = ?", email, accessToken).First(&token).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return &token, nil
}

func (db *DbToken) CreateNewToken(token models.Token) error {
	err := db.Conn.Create(&token).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

func (db *DbToken) UpdateToken(token models.Token) error {
	err := db.Conn.Save(&token).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

func (db *DbToken) DeleteToken(accessToken string) error {
	// Perform the delete operation where access_token matches
	return db.Conn.
		Where("access_token = ?", accessToken).
		Delete(&models.Token{}).Error
}

type TokenRepository interface {
	GetToken() ([]models.Token, error)
	GetTokenByUserEmail(email string, idToken string) (*models.Token, error)
	CreateNewToken(token models.Token) error
	UpdateToken(models.Token) error
	DeleteToken(idToken string) error
}

func NewDBTokenRepository(conn *gorm.DB) *DbToken {
	return &DbToken{Conn: conn}
}
