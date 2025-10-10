package jwt

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func ParseUserID(ctx context.Context, tokenString string, secret string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if uid, ok := claims["user_id"].(float64); ok {
			return int(uid), nil
		}
		return 0, errors.New("user_id not found in token claims")
	}

	return 0, errors.New("invalid token")
}
