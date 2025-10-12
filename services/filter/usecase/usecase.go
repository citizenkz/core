package usecase

import (
	"log/slog"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/services/filter/storage"
)

type usecase struct {
	log     *slog.Logger
	storage storage.Storage
	cfg     *config.Config
}

type UseCase interface{}

func New(log *slog.Logger, storage storage.Storage, cfg *config.Config) UseCase {
	return &usecase{
		log:     log,
		storage: storage,
		cfg:     cfg,
	}
}
