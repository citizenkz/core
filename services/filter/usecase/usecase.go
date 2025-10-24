package usecase

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/services/filter/entity"
	"github.com/citizenkz/core/services/filter/storage"
)

type usecase struct {
	log     *slog.Logger
	storage storage.Storage
	cfg     *config.Config
}

type UseCase interface {
	List(ctx context.Context, req *entity.ListRequest) (*entity.ListResponse, error)
	SaveUserFilters(ctx context.Context, req *entity.SaveFilersRequest) (*entity.SaveFilterResponse, error)
	Create(ctx context.Context, req *entity.CreateRequest) (*entity.CreateResponse, error)
	Delete(ctx context.Context, req *entity.DeleteRequest) (*entity.DeleteResponse, error)
}

func New(log *slog.Logger, storage storage.Storage, cfg *config.Config) UseCase {
	return &usecase{
		log:     log,
		storage: storage,
		cfg:     cfg,
	}
}
