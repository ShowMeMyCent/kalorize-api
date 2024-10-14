package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Database   DatabaseConfig `mapstructure:"database"`
	SigningKey string         `mapstructure:"signingkey"`
}

func InitDB() (db *gorm.DB, signingkey string, err error) {
	viper.SetConfigFile("prod.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		return nil, "", fmt.Errorf("Error reading config file: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, "", fmt.Errorf("Error unmarshalling config: %w", err)
	}

	host := config.Database.Host
	port := config.Database.Port
	dbname := config.Database.DBName
	username := config.Database.Username
	password := config.Database.Password

	signingkey = config.SigningKey

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, signingkey, err
}
