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
	benefitServer "github.com/citizenkz/core/services/benefit/server"
	benefitStorage "github.com/citizenkz/core/services/benefit/storage"
	benefitUsecase "github.com/citizenkz/core/services/benefit/usecase"
	categoryServer "github.com/citizenkz/core/services/category/server"
	categoryStorage "github.com/citizenkz/core/services/category/storage"
	categoryUsecase "github.com/citizenkz/core/services/category/usecase"
	childServer "github.com/citizenkz/core/services/child/server"
	childStorage "github.com/citizenkz/core/services/child/storage"
	childUsecase "github.com/citizenkz/core/services/child/usecase"
	filterServer "github.com/citizenkz/core/services/filter/server"
	filterStorage "github.com/citizenkz/core/services/filter/storage"
	filterUsecase "github.com/citizenkz/core/services/filter/usecase"
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

	filterStorage := filterStorage.New(s.log, client)
	filterUsecase := filterUsecase.New(s.log, filterStorage, s.cfg)
	filterServer := filterServer.New(s.log, filterUsecase)

	categoryStorage := categoryStorage.New(client, s.log)
	categoryUsecase := categoryUsecase.New(s.log, categoryStorage, s.cfg)
	categoryServer := categoryServer.New(s.log, categoryUsecase)

	benefitStorage := benefitStorage.New(client, s.log)
	benefitUsecase := benefitUsecase.New(s.log, benefitStorage, s.cfg)
	benefitServer := benefitServer.New(s.log, benefitUsecase)

	childStorage := childStorage.New(client, s.log)
	childUsecase := childUsecase.New(s.log, childStorage, s.cfg)
	childServer := childServer.New(s.log, childUsecase)

	router.Route("/api/v1", func(apiRouter chi.Router) {
		apiRouter.Route("/auth", func(authRouter chi.Router) {
			authRouter.Post("/login", userServer.HandleLogin)
			authRouter.Post("/register", userServer.HandleRegister)
			authRouter.Get("/profile", userServer.HandleGet)
			authRouter.Put("/password", userServer.HandleUpdatePassword)
			authRouter.Put("/email", userServer.HandleUpdateEmail)
			authRouter.Delete("/profile", userServer.HandleDelete)
			authRouter.Post("/forget-password", userServer.HandleForgetPassword)
			authRouter.Post("/forget-password/confirm", userServer.HandleForgetPasswordConfirm)
		})
		apiRouter.Route("/filter", func(filterRouter chi.Router) {
			filterRouter.Post("/", filterServer.Create)
			filterRouter.Post("/save", filterServer.SaveUserFitlers)
			filterRouter.Get("/", filterServer.List)
		})
		apiRouter.Route("/category", func(categoryRouter chi.Router) {
			categoryRouter.Post("/", categoryServer.HandleCreate)
			categoryRouter.Post("/list", categoryServer.HandleList)
			categoryRouter.Get("/{id}", categoryServer.HandleGet)
			categoryRouter.Put("/{id}", categoryServer.HandleUpdate)
			categoryRouter.Delete("/{id}", categoryServer.HandleDelete)
		})
		apiRouter.Route("/benefit", func(benefitRouter chi.Router) {
			benefitRouter.Post("/", benefitServer.HandleCreate)
			benefitRouter.Post("/list", benefitServer.HandleList)
			benefitRouter.Get("/{id}", benefitServer.HandleGet)
			benefitRouter.Put("/{id}", benefitServer.HandleUpdate)
			benefitRouter.Delete("/{id}", benefitServer.HandleDelete)
		})
		apiRouter.Route("/child", func(childRouter chi.Router) {
			childRouter.Post("/", childServer.HandleCreate)
			childRouter.Post("/list", childServer.HandleList)
			childRouter.Get("/{id}", childServer.HandleGet)
			childRouter.Put("/{id}", childServer.HandleUpdate)
			childRouter.Delete("/{id}", childServer.HandleDelete)
			childRouter.Post("/filters", childServer.HandleSaveFilters)
		})
	})

	s.log.Debug("server running", slog.String("address", fmt.Sprintf("localhost:%d", s.cfg.Port)))
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.cfg.Port), router)
}
