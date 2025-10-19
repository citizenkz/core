package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/filter/entity"
	"github.com/citizenkz/core/utils/jwt"
)

func (u *usecase) List(ctx context.Context, req *entity.ListRequest) (*entity.ListResponse, error) {
	filters, err := u.storage.List(ctx, req)
	if err != nil {
		u.log.Error("failed to storage.ListFilters", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.List: %w", err)
	}

	if req.Token != ""{
		userID, err := jwt.ParseUserID(ctx, req.Token, u.cfg.JwtSecret)
		if err != nil {
			u.log.Error("failed to jwtp.ParseUserID", slog.String("error", err.Error()))
		}
		for i, filter := range filters {
			userfilter, err := u.storage.GetUserFilter(ctx, userID, filter.ID)
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
