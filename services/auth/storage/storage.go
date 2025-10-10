package storage

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/ent/user"
	"github.com/citizenkz/core/services/auth/entity"
)

type storage struct {
	client *ent.Client
	log    *slog.Logger
}

type Storage interface {
	CreateUser(ctx context.Context, req *entity.RegisterRequest) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, userID int) (*entity.User, error)
	UpdateUser(ctx context.Context, req *entity.UpdateRequest) (*entity.User, error)
	UpdateUserPassword(ctx context.Context, password string) (*entity.User, error)
	UpdateUserEmail(ctx context.Context, email string) (*entity.User, error)
}

func New(client *ent.Client, log *slog.Logger) Storage {
	return &storage{
		client: client,
		log:    log,
	}
}

func (s *storage) CreateUser(ctx context.Context, req *entity.RegisterRequest) (*entity.User, error) {
	user, err := s.client.User.Create().
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetNillableBirthDate(req.BirthDate).
		SetPassword(req.Password).
		SetEmail(req.Email).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to save user", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageUserToEntity(user), nil
}

func (s *storage) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := s.client.User.Query().
		Where(
			user.Email(email),
		).First(ctx)
	if err != nil {
		s.log.Error("failed to get user email", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageUserToEntity(user), nil
}

func (s *storage) GetUserByID(ctx context.Context, userID int) (*entity.User, error) {
	user, err := s.client.User.Get(ctx, userID)
	if err != nil {
		s.log.Error("failed to get user by id", slog.String("error", err.Error()))
		return nil, err
	}
	
	return entity.MakeStorageUserToEntity(user), nil
}

func (s *storage) UpdateUser(ctx context.Context, req *entity.UpdateRequest) (*entity.User, error) {
	userID, err := s.client.User.Update().
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetNillableBirthDate(req.BirthDate).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update user", slog.String("error", err.Error()))
		return nil, err
	}

	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		s.log.Error("failed to get user by id", slog.String("error", err.Error()))
	}

	return user, nil
}

func (s *storage) UpdateUserPassword(ctx context.Context, password string) (*entity.User, error) {
	userID, err := s.client.User.Update().
		SetPassword(password).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update user's password", slog.String("error", err.Error()))
		return nil, err
	}

	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		s.log.Error("failed to get user by id", slog.String("error", err.Error()))
	}

	return user, nil
}

func (s *storage) UpdateUserEmail(ctx context.Context, email string) (*entity.User, error) {
	userID, err := s.client.User.Update().
		SetEmail(email).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update user's email", slog.String("error", err.Error()))
		return nil, err
	}

	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		s.log.Error("failed to get user by id", slog.String("error", err.Error()))
	}

	return user, nil
}
