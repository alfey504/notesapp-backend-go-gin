package utils

import (
	"errors"
	"fmt"
	"notes-appapi/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(username string) (string, error) {

	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(10 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = username

	secretKey, err := config.GetEnvVariableNonFatal(config.JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKeyInterface := []byte(secretKey)
	tokenString, err := newToken.SignedString(secretKeyInterface)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenFromRequest(c *gin.Context) (string, error) {
	if token := c.Query("token"); token != "" {
		return token, nil
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1], nil
	}
	return "", errors.New("Unable to find token in Header or body")
}

func ValidateToken(c *gin.Context) error {
	tokenString, err := ExtractTokenFromRequest(c)
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing method: %v ", token.Header["alg"])
		}
		secretKey, err := config.GetEnvVariableNonFatal(config.JWT_SECRET_KEY)
		if err != nil {
			return nil, err
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	return nil // returning nill when token is valid
}

func ValidateTokenAndGetUsername(c *gin.Context) (string, error) {
	tokenString, err := ExtractTokenFromRequest(c)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected Signing method: %v ", token.Header["alg"])
		}
		secretKey, err := config.GetEnvVariableNonFatal(config.JWT_SECRET_KEY)
		if err != nil {
			return nil, err
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		username := fmt.Sprintf("%v", claims["user"])
		return username, nil
	}

	return "", fmt.Errorf("Unable to find username")

}
