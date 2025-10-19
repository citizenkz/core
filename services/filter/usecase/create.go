package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/filter/entity"
)

func (u *usecase) Create(ctx context.Context, req *entity.CreateRequest) (*entity.CreateResponse, error) {
	filter, err := u.storage.Create(ctx, req)
	if err != nil {
		u.log.Error("failed to storage.Create", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.Create: %w", err)
	}
	
	return &entity.CreateResponse{
		Filter: *filter,
	}, nil
}
