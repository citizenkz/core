package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/services/child/entity"
	"github.com/citizenkz/core/services/child/storage"
	"github.com/citizenkz/core/utils/jwt"
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
	SaveFilters(ctx context.Context, req *entity.SaveFiltersRequest) (*entity.SaveFiltersResponse, error)
}

func New(log *slog.Logger, storage storage.Storage, cfg *config.Config) UseCase {
	return &usecase{
		log:     log,
		storage: storage,
		cfg:     cfg,
	}
}

func (u *usecase) Create(ctx context.Context, req *entity.CreateRequest) (*entity.CreateResponse, error) {
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.ParseUserID: %w", err)
	}

	child, err := u.storage.CreateChild(ctx, userID, req)
	if err != nil {
		u.log.Error("failed to storage.CreateChild", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.CreateChild: %w", err)
	}

	return &entity.CreateResponse{
		Child: child,
	}, nil
}

func (u *usecase) Get(ctx context.Context, req *entity.GetRequest) (*entity.GetResponse, error) {
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.ParseUserID: %w", err)
	}

	child, err := u.storage.GetChild(ctx, userID, req.ID)
	if err != nil {
		u.log.Error("failed to storage.GetChild", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.GetChild: %w", err)
	}

	return &entity.GetResponse{
		Child: child,
	}, nil
}

func (u *usecase) List(ctx context.Context, req *entity.ListRequest) (*entity.ListResponse, error) {
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.ParseUserID: %w", err)
	}

	children, total, err := u.storage.ListChildren(ctx, userID, req)
	if err != nil {
		u.log.Error("failed to storage.ListChildren", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.ListChildren: %w", err)
	}

	return &entity.ListResponse{
		Children: children,
		Total:    total,
	}, nil
}

func (u *usecase) Update(ctx context.Context, req *entity.UpdateRequest) (*entity.UpdateResponse, error) {
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.ParseUserID: %w", err)
	}

	child, err := u.storage.UpdateChild(ctx, userID, req)
	if err != nil {
		u.log.Error("failed to storage.UpdateChild", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.UpdateChild: %w", err)
	}

	return &entity.UpdateResponse{
		Child: child,
	}, nil
}

func (u *usecase) Delete(ctx context.Context, req *entity.DeleteRequest) (*entity.DeleteResponse, error) {
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.ParseUserID: %w", err)
	}

	err = u.storage.DeleteChild(ctx, userID, req.ID)
	if err != nil {
		u.log.Error("failed to storage.DeleteChild", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.DeleteChild: %w", err)
	}

	return &entity.DeleteResponse{
		Success: true,
	}, nil
}

func (u *usecase) SaveFilters(ctx context.Context, req *entity.SaveFiltersRequest) (*entity.SaveFiltersResponse, error) {
	userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
	if err != nil {
		u.log.Error("failed to jwt.ParseUserID", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to jwt.ParseUserID: %w", err)
	}

	// Verify the child belongs to the user
	_, err = u.storage.GetChild(ctx, userID, req.ChildID)
	if err != nil {
		u.log.Error("failed to storage.GetChild", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.GetChild: %w", err)
	}

	err = u.storage.SaveChildFilters(ctx, req.ChildID, req.Filters)
	if err != nil {
		u.log.Error("failed to storage.SaveChildFilters", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.SaveChildFilters: %w", err)
	}

	return &entity.SaveFiltersResponse{
		Success: true,
	}, nil
}
