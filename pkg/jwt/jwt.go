package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type CustomClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWTAccessToken(userId string, isAdmin bool) (string, error) {

	err := godotenv.Load("local.env")
	if err != nil {
		return "", fmt.Errorf("some error occured. Err: %w", err)
	}

	var role string

	if isAdmin {
		role = "admin"
	}

	claims := &CustomClaims{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		},
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenString.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string, secretKey string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func GetClaimsValue(claims jwt.MapClaims) map[string]string {

	m := make(map[string]string)
	userId := claims["userId"].(string)
	role := claims["role"].(string)

	m["userId"] = userId
	m["role"] = role

	return m
}
