package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTAccessToken(id int, fullname, email, key string) (string, error) {
	userId, _ := Encrypt(strconv.Itoa(id), EncryptionKey)
	emailEnc, _ := Encrypt(email, EncryptionKey)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"IdUser":   userId,
		"Fullname": fullname,
		"Email":    emailEnc,
		"exp":      time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+1, 0, 0, 0, 0, time.Now().Location()).Unix(),
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error(), err
	}
	return tokenString, err
}

func GenerateJWTRefreshToken(id int, fullname, email, key string) (string, error) {
	userId, _ := Encrypt(strconv.Itoa(id), EncryptionKey)
	emailEnc, _ := Encrypt(email, EncryptionKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"IdUser":   userId,
		"Fullname": fullname,
		"Email":    emailEnc,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error(), err
	}
	return tokenString, err
}
