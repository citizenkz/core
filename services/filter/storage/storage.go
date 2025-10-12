package storage

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/ent/filter"
	"github.com/citizenkz/core/ent/userfilter"
	"github.com/citizenkz/core/services/filter/entity"
)

type storage struct {
	log    *slog.Logger
	client *ent.Client
}

type Storage interface {
	Create(ctx context.Context, req *entity.CreateRequest) (*entity.Filter, error)
	Update(ctx context.Context, req *entity.UpdateRequest) (*entity.Filter, error)
	Get(ctx context.Context, id int) (*entity.Filter, error)
	List(ctx context.Context, req *entity.ListRequest) ([]*entity.Filter, error)
	GetUserFilter(ctx context.Context, userID, filterID int) (*entity.UserFilter, error)
	UpdateUserFilters(ctx context.Context, userID int, filterID int, value string) (*entity.UserFilter, error)
	CreateUserFilters(ctx context.Context, userID int, filterID int, value string) (*entity.UserFilter, error)
	DeleteFilter(ctx context.Context, filterID int) error
}

func New(log *slog.Logger, client *ent.Client) Storage {
	return &storage{
		log:    log,
		client: client,
	}
}

func (s *storage) Create(ctx context.Context, req *entity.CreateRequest) (*entity.Filter, error) {
	createdFilter, err := s.client.Filter.Create().
		SetName(req.Name).
		SetNillableHint(req.Hint).
		SetValues(req.Values).
		SetType(filter.Type(req.Type.String())).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to create filter", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageFilterToEntity(createdFilter), nil
}

func (s *storage) Update(ctx context.Context, req *entity.UpdateRequest) (*entity.Filter, error) {
	updatedFilter, err := s.client.Filter.UpdateOneID(req.ID).
		SetName(req.Name).
		SetNillableHint(req.Hint).
		SetValues(req.Values).
		SetType(filter.Type(req.Type.String())).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update filter", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageFilterToEntity(updatedFilter), nil
}

func (s *storage) Get(ctx context.Context, id int) (*entity.Filter, error) {
	filter, err := s.client.Filter.Get(ctx, id)
	if err != nil {
		s.log.Error("failed to get filter", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageFilterToEntity(filter), nil
}

func (s *storage) List(ctx context.Context, req *entity.ListRequest) ([]*entity.Filter, error) {
	filterQuery := s.client.Filter.Query()
	if req.SearchQuery != "" {
		filterQuery = filterQuery.Where(
			filter.NameContains(req.SearchQuery),
		)
	}
	
	filters, err := filterQuery.All(ctx)
	if err != nil {
		s.log.Error("failed to list filters", slog.String("error", err.Error()))
		return nil, err
	}
	
	return entity.MakeStorageFilterSliceToEntity(filters), nil
}

func (s *storage) GetUserFilter(ctx context.Context, userID, filterID int) (*entity.UserFilter, error) {
	userFilters, err := s.client.UserFilter.Query().
		Where(
			userfilter.UserIDEQ(userID),
			userfilter.FilterIDEQ(filterID),
		).First(ctx)
	if err != nil {
		s.log.Error("failed to get user filters", slog.String("error", err.Error()))
		return nil, err
	}
	return entity.MakeStorageUserFilterToEntity(userFilters), nil
}

func (s *storage) CreateUserFilters(ctx context.Context, userID int, filterID int, value string) (*entity.UserFilter, error) {
	userFilter, err := s.client.UserFilter.Create().
		SetUserID(userID).
		SetFilterID(filterID).
		SetValue(value).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to get user filters", slog.String("error", err.Error()))
		return nil, err
	}
	
	return entity.MakeStorageUserFilterToEntity(userFilter), nil
}

func (s *storage) UpdateUserFilters(ctx context.Context, userID int, filterID int, value string) (*entity.UserFilter, error) {
	_, err := s.client.UserFilter.Update().
		Where(
			userfilter.UserIDEQ(userID),
			userfilter.FilterIDEQ(filterID),
		).
		SetUserID(userID).
		SetFilterID(filterID).
		SetValue(value).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update user filters", slog.String("error", err.Error()))
		return nil, err
	}

	userFilter, err := s.GetUserFilter(ctx, userID, filterID)
	if err != nil {
		s.log.Error("failed to update user filters", slog.String("error", err.Error()))
		return nil, err
	}
	
	return userFilter, nil
}

func (s *storage) DeleteFilter(ctx context.Context, filterID int) error {
	err := s.client.Filter.DeleteOneID(filterID).
		Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete filter", slog.String("error", err.Error()))
		return err
	}

	return nil
}
