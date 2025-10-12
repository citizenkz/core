package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/filter/entity"
)

func (u *usecase) List(ctx context.Context, req *entity.ListRequest) (*entity.ListResponse, error) {
	filters, err := u.storage.List(ctx, req)
	if err != nil {
		u.log.Error("failed to storage.ListFilters", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.List: %w", err)
	}

	if req.UserID != nil {
		for i, filter := range filters {
			userfilter, err := u.storage.GetUserFilter(ctx, *req.UserID, filter.ID)
			if err != nil {
				u.log.Error("failed to storage.GetUserFilter", slog.String("error", err.Error()))
				return nil, fmt.Errorf("failed to storage.GetUserFilter: %w", err)
			}

			filter.SelectedValue = &userfilter.Value
			filters[i] = filter
		}
	}

	return &entity.ListResponse{
		Filters: filters,
	}, nil
}
