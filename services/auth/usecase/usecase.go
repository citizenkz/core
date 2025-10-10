package usecase

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/services/auth/entity"
	"github.com/citizenkz/core/services/auth/storage"
)

type usecase struct {
	log     *slog.Logger
	storage storage.Storage
	cfg     *config.Config
}

type UseCase interface{
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error)
	GetProfile(ctx context.Context, req *entity.GetRequest) (*entity.GetResponse, error)
	Update(ctx context.Context, req *entity.UpdateRequest) (*entity.UpdateResponse, error)
	UpdateEmail(ctx context.Context, req *entity.UpdateEmailRequest) (*entity.UpdateEmailResponse, error)
	UpdatePassword(ctx context.Context, req *entity.UpdatePasswordRequest) (*entity.UpdatePasswordResponse, error)
	ForgetPassword(ctx context.Context, req *entity.ForgetPasswordRequest) (*entity.ForgetPasswordResponse, error)
	ForgetPasswordConfirm(ctx context.Context, req *entity.ForgetPasswordConfirmRequest) (*entity.ForgetPasswordConfirmResponse, error)
}

func New(log *slog.Logger, storage storage.Storage, cfg *config.Config) UseCase {
	return &usecase{
		log:     log,
		storage: storage,
		cfg:     cfg,
	}
}
