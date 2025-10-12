package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/services/filter/entity"
)

func (u *usecase) SaveUserFilters(ctx context.Context, req *entity.SaveFilersRequest) (*entity.SaveFilterResponse, error) {
	userFilters := &entity.UserFilters{
		UserID: req.UserID,
	}
	newFilterValues := make([]*entity.FilterValues, 0)
	for _, filterValue := range req.FilterValues {
		filter, err := u.storage.GetUserFilter(ctx, req.UserID, filterValue.FilterID)
		if err != nil && ent.IsNotFound(err) {
			u.log.Error("failed to storage.GetUserFilter", slog.String("error", err.Error()))
			return nil, fmt.Errorf("failed to storage.GetUserFilter: %w", err)
		}
		switch {
		case ent.IsNotFound(err):
			userFilter, err := u.storage.CreateUserFilters(ctx, req.UserID, filterValue.FilterID, filterValue.Value)
			if err != nil {
				u.log.Error("failed to storage.GetUserFilter", slog.String("error", err.Error()))
				return nil, fmt.Errorf("failed to storage.GetUserFilter: %w", err)
			}
			newFilterValues = append(newFilterValues, &entity.FilterValues{
				FilterID: userFilter.FilterID,
				Value:    userFilter.Value,
			})
		case err == nil:
			userFilter, err := u.storage.UpdateUserFilters(ctx, filter.UserID, filter.FilterID, filterValue.Value)
			if err != nil {
				u.log.Error("failed to storage.GetUserFilter", slog.String("error", err.Error()))
				return nil, fmt.Errorf("failed to storage.GetUserFilter: %w", err)
			}
			newFilterValues = append(newFilterValues, &entity.FilterValues{
				FilterID: userFilter.FilterID,
				Value: userFilter.Value,
			})
		}
	}
	userFilters.Values = newFilterValues

	return &entity.SaveFilterResponse{
		UserFilters: *userFilters,
	}, nil
}
