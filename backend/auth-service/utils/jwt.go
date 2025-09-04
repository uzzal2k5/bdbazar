package utils

import (
	"time"
	"errors"
	"os"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)
var (
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
	errMissingSecret = errors.New("missing JWT secret")
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func init() {
	fmt.Println("JWT Key:", os.Getenv("JWT_SECRET"))
}


func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
        	ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
	}
    if len(jwtKey) == 0 {
        return "", errMissingSecret
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}


// ValidateJWT parses and validates the token
func ValidateJWT(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}