package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"gitlab.com/M.darvish/funtory/configs"
	"strconv"
	"time"
)

// GetAuthToken GenerateToken for authentication with jwt token
func GetAuthToken(username string) (string, error) {
	configs.Setup()
	secretKey := []byte(viper.GetString("jwt.secret_key"))
	return generateToken(username, secretKey)
}

// GetResetPwToken GenerateToken for reset password with jwt token
func GetResetPwToken(username string) (string, error) {
	configs.Setup()
	secretKey := []byte(viper.GetString("jwt.reset_pw_key"))
	return generateToken(username, secretKey)
}

// generateToken GenerateToken with jwt token
func generateToken(username string, secretKey []byte) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	// Set the expiration time for the token
	expirationTimeStr := viper.GetString("jwt.expiration_time")

	expirationTime, err := strconv.Atoi(expirationTimeStr)
	if err != nil {
		return "", err
	}

	expirationDuration := time.Duration(expirationTime) * time.Hour
	claims["exp"] = time.Now().Add(expirationDuration).Unix()

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidationAuthToken verify the token for valid expiration and secret key
func ValidationAuthToken(tokenStr string) (string, error) {
	configs.Setup()
	secretKey := []byte(viper.GetString("jwt.secret_key"))
	return checkToken(tokenStr, secretKey)
}

// ValidationResetPwToken verify the token for valid expiration and reset password key
func ValidationResetPwToken(tokenStr string) (string, error) {
	configs.Setup()
	secretKey := []byte(viper.GetString("jwt.reset_pw_key"))
	return checkToken(tokenStr, secretKey)
}

// checkToken verify the token
func checkToken(tokenStr string, secretKey []byte) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing")
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}

	// Check the expiration time
	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return "", errors.New("invalid token")
	}

	// Get the username claim from the token
	username, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("userName claims not found")
	}

	return username, nil
}
