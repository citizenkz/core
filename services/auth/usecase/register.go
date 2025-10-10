package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/auth/entity"
	"github.com/citizenkz/core/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (u *usecase) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Error("failed to bcrypt.GenerateFromPassword", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to bcrypt.GenerateFromPassword: %w", err)
	}

	req.Password = string(hashedPassword)

	user, err := u.storage.CreateUser(ctx, req)
	if err != nil {
		u.log.Error("failed to storage.CreateUser", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.CreateUser: %w", err)
	}

	token, err := jwt.Generate(ctx, user.ID, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.Generate", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.Generate: %w", err)
	}
	
	return &entity.RegisterResponse{
		Profile: *user,
		Token: token,
	}, nil
}

