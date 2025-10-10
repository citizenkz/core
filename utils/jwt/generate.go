package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Generate(ctx context.Context, user_id int, secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
