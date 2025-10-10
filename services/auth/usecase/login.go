package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/auth/entity"
	"github.com/citizenkz/core/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := u.storage.GetUserByEmail(ctx, req.Email)
	if err != nil {
		u.log.Error("failed to storage.GetUserByEmail", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.GetUserByEmail: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		u.log.Error("failed to bcrypt.CompareHashAndPassword", slog.String("error", err.Error()))
		return nil, fmt.Errorf("incorrect password")
	}

	token, err := jwt.Generate(ctx, user.ID, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.Generate", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &entity.LoginResponse{
		Token:   token,
		Profile: *user,
	}, nil
}
