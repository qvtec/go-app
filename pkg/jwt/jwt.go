package jwt

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type MapClaims = jwt.MapClaims

type JWTManager struct {
	secretKey []byte
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{
		secretKey: []byte(secretKey),
	}
}

func (manager *JWTManager) GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(manager.secretKey)
}

func (manager *JWTManager) VerifyToken(bearer string) (jwt.Claims, error) {
	if bearer == "" {
		return nil, errors.New("invalid Authorization header")
	}

	tokenString := strings.TrimPrefix(bearer, "Bearer ")
	if tokenString == bearer {
		return nil, errors.New("invalid Authorization Bearer header")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return manager.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims, nil
}
