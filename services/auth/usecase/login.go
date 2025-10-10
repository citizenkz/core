package usecase

import (
	"context"

	"github.com/citizenkz/core/services/auth/entity"
)

func (u *usecase) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	return nil, nil
}
