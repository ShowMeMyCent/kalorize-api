package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"kalorize-api/utils"
	"net/http"
	"strings"
)

// JWTMiddleware provides JWT validation middleware
func (controller *TokenController) JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get signingKey from context
			signingKey, ok := c.Get("signingKey").(string)
			if !ok {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "signing key not found"})
			}

			// JWT middleware configuration
			jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
				SigningKey:  []byte(signingKey),
				Claims:      &jwt.MapClaims{},
				TokenLookup: "header:Authorization",
				AuthScheme:  "Bearer",
				ErrorHandlerWithContext: func(err error, c echo.Context) error {
					return c.JSON(http.StatusUnauthorized, map[string]string{"message": "invalid or missing token"})
				},
			})

			// Use the configured JWT middleware
			return jwtMiddleware(next)(c)
		}
	}
}

// CheckTokenMiddleware validates the token from the database against the JWT token
func (controller *TokenController) CheckTokenMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get signingKey from context
			signingKey, ok := c.Get("signingKey").(string)
			if !ok {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"statusCode": http.StatusUnauthorized, "messages": "signing key not found"})
			}

			// Extract the token from the Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"statusCode": http.StatusUnauthorized,
					"message":    "Missing or invalid authorization header",
				})
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			// Parse the token
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				// Ensure that the signing method is HMAC
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(signingKey), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"statusCode": http.StatusBadRequest, "messages": "invalid token or expired"})
			}

			// Get JWT claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"statusCode": http.StatusBadRequest, "messages": "invalid token claims"})
			}

			emailClaim, ok := claims["Email"]
			if !ok {
				// Jika klaim tidak ada, kirimkan response unauthorized
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"statusCode": http.StatusUnauthorized,
					"messages":   "email claim not found in token",
				})
			}

			// Decrypt the user ID from the claims
			email, err := utils.Decrypt(emailClaim.(string), utils.EncryptionKey)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"statusCode": http.StatusUnauthorized, "messages": "invalid token"})
			}

			// Retrieve the token from the service layer using the email
			storedToken, err := controller.tokenService.GetTokenByUserEmail(email, tokenStr)
			if err != nil || storedToken == nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"statusCode": http.StatusBadRequest, "messages": "token not found"})
			}

			// Compare the token in the database with the JWT token from the Authorization header
			if tokenStr != storedToken.AccessToken {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{"statusCode": http.StatusBadRequest, "messages": "invalid token"})
			}

			// If everything is valid, proceed to the next handler
			return next(c)
		}
	}
}
