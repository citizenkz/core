package storage

import (
	"log/slog"

	"github.com/citizenkz/core/ent"
)

type storage struct {
	log *slog.Logger
	client *ent.Client
}

type Storage interface{}

func New(log *slog.Logger, client *ent.Client) Storage {
	return &storage{
		log: log,
		client: client,
	}
}
