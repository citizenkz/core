package server

import (
	"log/slog"

	"github.com/citizenkz/core/services/filter/usecase"
)

type server struct {
	log     *slog.Logger
	usecase usecase.UseCase
}

type Server interface{}

func New(log *slog.Logger, usecase usecase.UseCase) Server {
	return &server{
		log:     log,
		usecase: usecase,
	}
}
