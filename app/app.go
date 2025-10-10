package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/ent"

	_ "github.com/lib/pq"
)

type Server interface {
	Run() error
}

type server struct {
	cfg *config.Config
	log *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) Server {
	return &server{
		cfg: cfg,
		log: log,
	}
}

func(s *server) Run() error {
	entPsqlConnect := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		s.cfg.Database.Host,
		s.cfg.Database.Port,
		s.cfg.Database.User,
		s.cfg.Database.Name,
		s.cfg.Database.Password,
		s.cfg.Database.SSLMode,
	)
	
	client, err := ent.Open("postgres", entPsqlConnect)
    if err != nil {
        s.log.Error("failed opening connection to postgres", slog.String("error", err.Error()))
    }
    defer client.Close()

    if err := client.Schema.Create(context.Background()); err != nil {
        s.log.Error("failed creating schema resources", slog.String("error", err.Error()))
    }
	return nil
}
