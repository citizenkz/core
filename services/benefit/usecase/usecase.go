package usecase

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/services/benefit/entity"
	"github.com/citizenkz/core/services/benefit/storage"
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
	benefit, err := u.storage.CreateBenefit(ctx, req)
	if err != nil {
		u.log.Error("failed to create benefit", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.CreateResponse{
		Benefit: benefit,
	}, nil
}

func (u *usecase) Get(ctx context.Context, req *entity.GetRequest) (*entity.GetResponse, error) {
	benefit, err := u.storage.GetBenefit(ctx, req.ID)
	if err != nil {
		u.log.Error("failed to get benefit", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.GetResponse{
		Benefit: benefit,
	}, nil
}

func (u *usecase) List(ctx context.Context, req *entity.ListRequest) (*entity.ListResponse, error) {
	benefits, total, err := u.storage.ListBenefits(ctx, req)
	if err != nil {
		u.log.Error("failed to list benefits", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.ListResponse{
		Benefits: benefits,
		Total:    total,
	}, nil
}

func (u *usecase) Update(ctx context.Context, req *entity.UpdateRequest) (*entity.UpdateResponse, error) {
	benefit, err := u.storage.UpdateBenefit(ctx, req)
	if err != nil {
		u.log.Error("failed to update benefit", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.UpdateResponse{
		Benefit: benefit,
	}, nil
}

func (u *usecase) Delete(ctx context.Context, req *entity.DeleteRequest) (*entity.DeleteResponse, error) {
	err := u.storage.DeleteBenefit(ctx, req.ID)
	if err != nil {
		u.log.Error("failed to delete benefit", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.DeleteResponse{
		Success: true,
	}, nil
}
