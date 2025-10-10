package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/citizenkz/core/config"
	"github.com/citizenkz/core/ent"
	userServer "github.com/citizenkz/core/services/auth/server"
	userStorage "github.com/citizenkz/core/services/auth/storage"
	userUsecase "github.com/citizenkz/core/services/auth/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

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

func (s *server) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

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

	userStorage := userStorage.New(client, s.log)
	userUsecase := userUsecase.New(s.log, userStorage, s.cfg)
	userServer := userServer.New(s.log, userUsecase)

	router.Route("/api/v1", func(apiRouter chi.Router) {
		apiRouter.Route("/auth", func(authRouter chi.Router) {
			authRouter.Post("/login", userServer.HandleLogin)
			authRouter.Post("/register", userServer.HandleRegister)
			authRouter.Get("/profile", userServer.HandleGet)
		})
	})

	s.log.Debug("server running", slog.String("address", fmt.Sprintf("localhost:%d", s.cfg.Port)))
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.cfg.Port), router)
}
