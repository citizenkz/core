package storage

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/ent/child"
	"github.com/citizenkz/core/ent/childfilter"
	"github.com/citizenkz/core/services/child/entity"
)

type storage struct {
	client *ent.Client
	log    *slog.Logger
}

type Storage interface {
	CreateChild(ctx context.Context, userID int, req *entity.CreateRequest) (*entity.Child, error)
	GetChild(ctx context.Context, userID, childID int) (*entity.Child, error)
	ListChildren(ctx context.Context, userID int, req *entity.ListRequest) ([]*entity.Child, int, error)
	UpdateChild(ctx context.Context, userID int, req *entity.UpdateRequest) (*entity.Child, error)
	DeleteChild(ctx context.Context, userID, childID int) error
	SaveChildFilters(ctx context.Context, childID int, filters []entity.FilterValueRequest) error
	GetChildFilters(ctx context.Context, childID int) ([]*entity.ChildFilter, error)
}

func New(client *ent.Client, log *slog.Logger) Storage {
	return &storage{
		client: client,
		log:    log,
	}
}

func (s *storage) CreateChild(ctx context.Context, userID int, req *entity.CreateRequest) (*entity.Child, error) {
	c, err := s.client.Child.
		Create().
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetBirthDate(req.BirthDate).
		SetUserID(userID).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to create child", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.Child{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		BirthDate: c.BirthDate,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
	}, nil
}

func (s *storage) GetChild(ctx context.Context, userID, childID int) (*entity.Child, error) {
	c, err := s.client.Child.
		Query().
		Where(child.ID(childID), child.UserID(userID)).
		Only(ctx)
	if err != nil {
		s.log.Error("failed to get child", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.Child{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		BirthDate: c.BirthDate,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
	}, nil
}

func (s *storage) ListChildren(ctx context.Context, userID int, req *entity.ListRequest) ([]*entity.Child, int, error) {
	query := s.client.Child.
		Query().
		Where(child.UserID(userID))

	// Get total count
	total, err := query.Count(ctx)
	if err != nil {
		s.log.Error("failed to count children", slog.String("error", err.Error()))
		return nil, 0, err
	}

	// Apply pagination
	children, err := query.
		Limit(req.Limit).
		Offset(req.Offset).
		All(ctx)
	if err != nil {
		s.log.Error("failed to list children", slog.String("error", err.Error()))
		return nil, 0, err
	}

	result := make([]*entity.Child, len(children))
	for i, c := range children {
		result[i] = &entity.Child{
			ID:        c.ID,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			BirthDate: c.BirthDate,
			UserID:    c.UserID,
			CreatedAt: c.CreatedAt,
		}
	}

	return result, total, nil
}

func (s *storage) UpdateChild(ctx context.Context, userID int, req *entity.UpdateRequest) (*entity.Child, error) {
	c, err := s.client.Child.
		UpdateOneID(req.ID).
		Where(child.UserID(userID)).
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetBirthDate(req.BirthDate).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update child", slog.String("error", err.Error()))
		return nil, err
	}

	return &entity.Child{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		BirthDate: c.BirthDate,
		UserID:    c.UserID,
		CreatedAt: c.CreatedAt,
	}, nil
}

func (s *storage) DeleteChild(ctx context.Context, userID, childID int) error {
	// First delete all child filters
	_, err := s.client.ChildFilter.
		Delete().
		Where(childfilter.ChildID(childID)).
		Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete child filters", slog.String("error", err.Error()))
		return err
	}

	// Then delete the child
	err = s.client.Child.
		DeleteOneID(childID).
		Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete child", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (s *storage) SaveChildFilters(ctx context.Context, childID int, filters []entity.FilterValueRequest) error {
	// Delete existing filters for this child
	_, err := s.client.ChildFilter.
		Delete().
		Where(childfilter.ChildID(childID)).
		Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete existing child filters", slog.String("error", err.Error()))
		return err
	}

	// Create new filters
	for _, f := range filters {
		_, err := s.client.ChildFilter.
			Create().
			SetChildID(childID).
			SetFilterID(f.FilterID).
			SetValue(f.Value).
			Save(ctx)
		if err != nil {
			s.log.Error("failed to create child filter", slog.String("error", err.Error()))
			return err
		}
	}

	return nil
}

func (s *storage) GetChildFilters(ctx context.Context, childID int) ([]*entity.ChildFilter, error) {
	filters, err := s.client.ChildFilter.
		Query().
		Where(childfilter.ChildID(childID)).
		All(ctx)
	if err != nil {
		s.log.Error("failed to get child filters", slog.String("error", err.Error()))
		return nil, err
	}

	result := make([]*entity.ChildFilter, len(filters))
	for i, f := range filters {
		result[i] = &entity.ChildFilter{
			ID:       f.ID,
			ChildID:  f.ChildID,
			FilterID: f.FilterID,
			Value:    f.Value,
		}
	}

	return result, nil
}
