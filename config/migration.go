package config

import (
	"gorm.io/gorm"
	"kalorize-api/app/models"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserAdmin{})
	db.AutoMigrate(&models.Token{})
	db.AutoMigrate(&models.GymCode{})
	db.AutoMigrate(&models.UsedCode{})
	db.AutoMigrate(&models.Gym{})
	db.AutoMigrate(&models.Makanan{})
	db.AutoMigrate(&models.MealSet{})
	db.AutoMigrate(&models.Franchise{})
	db.AutoMigrate(&models.History{})
	db.AutoMigrate(&models.FranchiseMakanan{})
}
