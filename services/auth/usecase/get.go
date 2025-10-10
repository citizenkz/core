package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/auth/entity"
	"github.com/citizenkz/core/utils/jwt"
)

func (u *usecase) GetProfile(ctx context.Context, req *entity.GetRequest) (*entity.GetResponse, error) {
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.ParseUserID: %w", err)
	}
	
	user, err := u.storage.GetUserByID(ctx, userID)
	if err != nil {
		u.log.Error("failed to storagee.GetUserByID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.GetUserByID: %w", err)
	}
	
	return &entity.GetResponse{
		Profile: *user,
	}, nil
}
