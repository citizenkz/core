package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/services/filter/entity"
)

func (u *usecase) Delete(ctx context.Context, req *entity.DeleteRequest) (*entity.DeleteResponse, error) {
	err := u.storage.DeleteFilter(ctx, req.ID)
	if err != nil {
		u.log.Error("failed to storage.DeleteFilter", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to storage.DeleteFilter: %w", err)
	}

	return &entity.DeleteResponse{
		IsDeleted: true,
	}, nil
}
