package libs

import (
	"fmt"
	"garavel/internal/configs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey []byte

func init() {
	secretKey = []byte(configs.Env("JWT_KEY", ""))
}

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub": strconv.FormatUint(uint64(userID), 10),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(configs.EnvInt("JWT_TTL", "24")) * time.Hour).Unix(), // 24 hours expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	userID := claims["sub"].(string)
	return userID, nil
}
