package storage

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/ent/attempt"
	"github.com/citizenkz/core/ent/user"
	"github.com/citizenkz/core/services/auth/entity"
	"github.com/google/uuid"
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
	UpdateUserPassword(ctx context.Context, userID int, password string) (*entity.User, error)
	UpdateUserEmail(ctx context.Context, userID int, email string) (*entity.User, error)
	DeleteUser(ctx context.Context, userID int) error
	CreateAttempt(ctx context.Context, email, otp string) (uuid.UUID, error)
	GetAttempt(ctx context.Context, attemptID uuid.UUID) (*ent.Attempt, error)
	DeleteAttempt(ctx context.Context, attemptID uuid.UUID) error
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
	user, err := s.client.User.UpdateOneID(req.ID).
		SetFirstName(req.FirstName).
		SetLastName(req.LastName).
		SetNillableBirthDate(req.BirthDate).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update user", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageUserToEntity(user), nil
}

func (s *storage) UpdateUserPassword(ctx context.Context, userID int, password string) (*entity.User, error) {
	user, err := s.client.User.UpdateOneID(userID).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update user's password", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageUserToEntity(user), nil
}

func (s *storage) UpdateUserEmail(ctx context.Context, userID int, email string) (*entity.User, error) {
	user, err := s.client.User.UpdateOneID(userID).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update user's email", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageUserToEntity(user), nil
}

func (s *storage) DeleteUser(ctx context.Context, userID int) error {
	err := s.client.User.DeleteOneID(userID).Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete user", slog.String("error", err.Error()))
		return err
	}

	return nil
}

func (s *storage) CreateAttempt(ctx context.Context, email, otp string) (uuid.UUID, error) {
	attempt, err := s.client.Attempt.Create().
		SetEmail(email).
		SetOtp(otp).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to create attempt", slog.String("error", err.Error()))
		return uuid.Nil, err
	}

	return attempt.ID, nil
}

func (s *storage) GetAttempt(ctx context.Context, attemptID uuid.UUID) (*ent.Attempt, error) {
	attempt, err := s.client.Attempt.Query().
		Where(attempt.ID(attemptID)).
		First(ctx)
	if err != nil {
		s.log.Error("failed to get attempt", slog.String("error", err.Error()))
		return nil, err
	}

	return attempt, nil
}

func (s *storage) DeleteAttempt(ctx context.Context, attemptID uuid.UUID) error {
	_, err := s.client.Attempt.Delete().
		Where(attempt.ID(attemptID)).
		Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete attempt", slog.String("error", err.Error()))
		return err
	}

	return nil
}
