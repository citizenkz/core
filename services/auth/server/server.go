package server

import (
	"log/slog"
	"net/http"

	"github.com/citizenkz/core/services/auth/entity"
	"github.com/citizenkz/core/services/auth/usecase"
	"github.com/citizenkz/core/utils/json"
	"github.com/citizenkz/core/utils/jwt"
)

type server struct {
	log     *slog.Logger
	usecase usecase.UseCase
}

type Server interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleRegister(w http.ResponseWriter, r *http.Request)
	HandleGet(w http.ResponseWriter, r *http.Request)
}

func New(log *slog.Logger, usecase usecase.UseCase) Server {
	return &server{
		log:     log,
		usecase: usecase,
	}
}

func (s *server) HandleLogin(w http.ResponseWriter, r *http.Request) {
	req := &entity.LoginRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := s.usecase.Login(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Login", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJson", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleRegister(w http.ResponseWriter, r *http.Request) {
	req := &entity.RegisterRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := s.usecase.Register(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Register", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJson", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleGet(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		s.log.Error("failed to jwt.ParseTokenFromHeader", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	req := &entity.GetRequest{
		Token: token,
	}

	resp, err := s.usecase.GetProfile(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.GetProfile", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJson", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}
