package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

// Constants
const (
	// Key for AES encryption/decryption (must be 16, 24, or 32 bytes for AES)
	EncryptionKey = "thisis32bitlongpassphraseimusing" // Example for AES-256

	// Nonce size used in AES-GCM (not a constant in the strict sense but good to define)
	NonceSize = 12
)

func ParseDataEmail(bearerToken string) (email string, err error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kalorize"), nil
	})

	if err != nil {
		log.Printf("Error parsing JWT token: %v", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		emailClaim := claims["Email"]
		if emailClaim == nil {
			err := fmt.Errorf("email claim is missing in JWT token")
			log.Printf("Error: %v", err)
			return "", err
		}
		email = emailClaim.(string)
	}
	return email, err
}

func ParseDataFullname(bearerToken string) (email string, err error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("kalorize"), nil
	})

	if err != nil {
		log.Printf("Error parsing JWT token: %v", err)
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		emailClaim := claims["Fullname"]
		if emailClaim == nil {
			err := fmt.Errorf("email claim is missing in JWT token")
			log.Printf("Error: %v", err)
			return "", err
		}
		email = emailClaim.(string)
		
	}
	return email, err
}

func ParseDataId(bearerToken string) (id int, err error) {
	// Parse the JWT token
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key
		return []byte("kalorize"), nil
	})

	if err != nil {
		log.Printf("Error parsing JWT token: %v", err)
		return id, err
	}

	// Check if the token's claims are valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println("Claims are valid")

		// Ensure "IdUser" claim exists and is of type float64 (common in JSON decoding)
		idClaim, ok := claims["IdUser"].(string)
		if !ok {
			err := fmt.Errorf("id claim is missing or not a number in JWT token")
			log.Printf("Error: %v", err)
			return id, err
		}

		// Convert the float64 to int
		idUser, _ := Decrypt(idClaim, EncryptionKey)
		id, _ = strconv.Atoi(idUser)
		log.Printf("IdUser claim: %d", id)
	} else {
		err := fmt.Errorf("claims are not of type jwt.MapClaims or token is invalid")
		log.Printf("Error: %v", err)
		return id, err
	}

	return id, nil
}

func Encrypt(plainText string, key string) (string, error) {
	// Convert the key to a byte slice
	keyBytes := []byte(key)
	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("error creating AES cipher: %v", err)
	}

	// Create a GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating GCM cipher: %v", err)
	}

	// Generate a nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("error generating nonce: %v", err)
	}

	// Encrypt the plain text
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	// Encode to base64 for easy storage
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts base64-encoded encrypted text using AES
func Decrypt(cipherText string, key string) (string, error) {
	// Convert the key and cipher text to byte slices
	keyBytes := []byte(key)
	cipherTextBytes, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", fmt.Errorf("error decoding base64: %v", err)
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", fmt.Errorf("error creating AES cipher: %v", err)
	}

	// Create a GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("error creating GCM cipher: %v", err)
	}

	// Extract the nonce and the encrypted text
	nonceSize := gcm.NonceSize()
	if len(cipherTextBytes) < nonceSize {
		return "", fmt.Errorf("cipher text too short")
	}
	nonce, cipherTextBytes := cipherTextBytes[:nonceSize], cipherTextBytes[nonceSize:]

	// Decrypt the cipher text
	plainTextBytes, err := gcm.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", fmt.Errorf("error decrypting text: %v", err)
	}

	return string(plainTextBytes), nil
}
