package usecase

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/services/category/entity"
	"github.com/citizenkz/core/services/category/storage"
)

type usecase struct {
	log     *slog.Logger
	storage storage.Storage
	cfg     *config.Config
}

type UseCase interface {
	Create(ctx context.Context, req *entity.CreateRequest) (*entity.CreateResponse, error)
	Get(ctx context.Context, req *entity.GetRequest) (*entity.GetResponse, error)
	List(ctx context.Context, req *entity.ListRequest) (*entity.ListResponse, error)
	Update(ctx context.Context, req *entity.UpdateRequest) (*entity.UpdateResponse, error)
	Delete(ctx context.Context, req *entity.DeleteRequest) (*entity.DeleteResponse, error)
}

func New(log *slog.Logger, storage storage.Storage, cfg *config.Config) UseCase {
	return &usecase{
		log:     log,
		storage: storage,
		cfg:     cfg,
	}
}

func (u *usecase) Create(ctx context.Context, req *entity.CreateRequest) (*entity.CreateResponse, error) {
	category, err := u.storage.CreateCategory(ctx, req)
	if err != nil {
		u.log.Error("failed to create category", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.CreateResponse{
		Category: category,
	}, nil
}

func (u *usecase) Get(ctx context.Context, req *entity.GetRequest) (*entity.GetResponse, error) {
	category, err := u.storage.GetCategory(ctx, req.ID)
	if err != nil {
		u.log.Error("failed to get category", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.GetResponse{
		Category: category,
	}, nil
}

func (u *usecase) List(ctx context.Context, req *entity.ListRequest) (*entity.ListResponse, error) {
	categories, total, err := u.storage.ListCategories(ctx, req)
	if err != nil {
		u.log.Error("failed to list categories", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.ListResponse{
		Categories: categories,
		Total:      total,
	}, nil
}

func (u *usecase) Update(ctx context.Context, req *entity.UpdateRequest) (*entity.UpdateResponse, error) {
	category, err := u.storage.UpdateCategory(ctx, req)
	if err != nil {
		u.log.Error("failed to update category", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.UpdateResponse{
		Category: category,
	}, nil
}

func (u *usecase) Delete(ctx context.Context, req *entity.DeleteRequest) (*entity.DeleteResponse, error) {
	err := u.storage.DeleteCategory(ctx, req.ID)
	if err != nil {
		u.log.Error("failed to delete category", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.DeleteResponse{
		Success: true,
	}, nil
}
