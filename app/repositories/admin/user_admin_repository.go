package admin

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
	"kalorize-api/utils"
)

type dbUser struct {
	Conn *gorm.DB
}

func NewDBUserRepository(conn *gorm.DB) *dbUser {
	return &dbUser{Conn: conn}
}

// UserRepository defines the interface for user-related data operations.
type UserRepository interface {
	GetUser() ([]models.UserAdmin, error)
	GetUserById(id int) (models.UserAdmin, error)
	CreateNewUser(user models.UserAdmin) error
	UpdateUser(user models.UserAdmin) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (models.UserAdmin, error)
}

// GetUser retrieves all users from the database.
func (db *dbUser) GetUser() ([]models.UserAdmin, error) {
	var users []models.UserAdmin
	err := db.Conn.Find(&users).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return users, err
}

// GetUserById retrieves a user by ID.
func (db *dbUser) GetUserById(id int) (models.UserAdmin, error) {
	var user models.UserAdmin

	err := db.Conn.First(&user, id).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return user, err
	}

	return user, nil
}

// CreateNewUser creates a new user in the database.
func (db *dbUser) CreateNewUser(user models.UserAdmin) error {
	err := db.Conn.Create(&user).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

// UpdateUser updates an existing user.
func (db *dbUser) UpdateUser(user models.UserAdmin) error {
	err := db.Conn.Save(&user).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

// DeleteUser deletes a user by ID.
func (db *dbUser) DeleteUser(id int) error {
	err := db.Conn.Delete(&models.UserAdmin{}, id).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

// GetUserByEmail retrieves a user by email.
func (db *dbUser) GetUserByEmail(email string) (models.UserAdmin, error) {
	var user models.UserAdmin
	err := db.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return models.UserAdmin{}, err
	}
	return user, err
}
