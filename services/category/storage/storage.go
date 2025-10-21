package storage

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/ent/category"
	"github.com/citizenkz/core/services/category/entity"
)

type storage struct {
	client *ent.Client
	log    *slog.Logger
}

type Storage interface {
	CreateCategory(ctx context.Context, req *entity.CreateRequest) (*entity.Category, error)
	GetCategory(ctx context.Context, id int) (*entity.Category, error)
	ListCategories(ctx context.Context, req *entity.ListRequest) ([]*entity.Category, int, error)
	UpdateCategory(ctx context.Context, req *entity.UpdateRequest) (*entity.Category, error)
	DeleteCategory(ctx context.Context, id int) error
}

func New(client *ent.Client, log *slog.Logger) Storage {
	return &storage{
		client: client,
		log:    log,
	}
}

func (s *storage) CreateCategory(ctx context.Context, req *entity.CreateRequest) (*entity.Category, error) {
	category, err := s.client.Category.Create().
		SetName(req.Name).
		SetNillableDescription(req.Description).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to save category", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageCategoryToEntity(category), nil
}

func (s *storage) GetCategory(ctx context.Context, id int) (*entity.Category, error) {
	category, err := s.client.Category.Get(ctx, id)
	if err != nil {
		s.log.Error("failed to get category", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageCategoryToEntity(category), nil
}

func (s *storage) ListCategories(ctx context.Context, req *entity.ListRequest) ([]*entity.Category, int, error) {
	query := s.client.Category.Query()

	// Apply search filter
	if req.Search != "" {
		query = query.Where(
			category.Or(
				category.NameContains(req.Search),
				category.DescriptionContains(req.Search),
			),
		)
	}

	// Get total count
	total, err := query.Count(ctx)
	if err != nil {
		s.log.Error("failed to count categories", slog.String("error", err.Error()))
		return nil, 0, err
	}

	// Apply pagination
	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset > 0 {
		query = query.Offset(req.Offset)
	}

	categories, err := query.All(ctx)
	if err != nil {
		s.log.Error("failed to list categories", slog.String("error", err.Error()))
		return nil, 0, err
	}

	result := make([]*entity.Category, 0, len(categories))
	for _, c := range categories {
		result = append(result, entity.MakeStorageCategoryToEntity(c))
	}

	return result, total, nil
}

func (s *storage) UpdateCategory(ctx context.Context, req *entity.UpdateRequest) (*entity.Category, error) {
	category, err := s.client.Category.UpdateOneID(req.ID).
		SetName(req.Name).
		SetNillableDescription(req.Description).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update category", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageCategoryToEntity(category), nil
}

func (s *storage) DeleteCategory(ctx context.Context, id int) error {
	err := s.client.Category.DeleteOneID(id).Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete category", slog.String("error", err.Error()))
		return err
	}

	return nil
}
