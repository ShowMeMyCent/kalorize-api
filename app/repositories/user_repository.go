package repositories

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
	GetUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	CreateNewUser(user models.User) error
	UpdateUser(user models.User) error
	DeleteUser(id int) error
	GetUserByEmail(email string) (models.User, error)
}

// GetUser retrieves all users from the database.
func (db *dbUser) GetUser() ([]models.User, error) {
	var users []models.User
	err := db.Conn.Find(&users).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return nil, err
	}
	return users, err
}

// GetUserById retrieves a user by ID.
func (db *dbUser) GetUserById(id int) (models.User, error) {
	var user models.User

	err := db.Conn.Preload("UsedCodes").
		Preload("Histories", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Breakfast").
				Preload("Lunch").
				Preload("Dinner")
		}).
		First(&user, id).Error

	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return user, err
	}

	return user, nil
}

// CreateNewUser creates a new user in the database.
func (db *dbUser) CreateNewUser(user models.User) error {
	err := db.Conn.Create(&user).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}
	return nil
}

// UpdateUser updates an existing user.
func (db *dbUser) UpdateUser(user models.User) error {
	err := db.Conn.Save(&user).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

// DeleteUser deletes a user by ID.
func (db *dbUser) DeleteUser(id int) error {
	err := db.Conn.Delete(&models.User{}, id).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return err
	}

	return nil
}

// GetUserByEmail retrieves a user by email.
func (db *dbUser) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.Conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		err = utils.Error(err, utils.ErrDatabase, err.Error())
		return models.User{}, err
	}
	return user, err
}
