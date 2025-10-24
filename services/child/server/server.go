package server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/citizenkz/core/services/child/entity"
	"github.com/citizenkz/core/services/child/usecase"
	"github.com/citizenkz/core/utils/json"
	"github.com/citizenkz/core/utils/jwt"
	"github.com/go-chi/chi/v5"
)

type server struct {
	log     *slog.Logger
	usecase usecase.UseCase
}

type Server interface {
	HandleCreate(w http.ResponseWriter, r *http.Request)
	HandleGet(w http.ResponseWriter, r *http.Request)
	HandleList(w http.ResponseWriter, r *http.Request)
	HandleUpdate(w http.ResponseWriter, r *http.Request)
	HandleDelete(w http.ResponseWriter, r *http.Request)
	HandleSaveFilters(w http.ResponseWriter, r *http.Request)
}

func New(log *slog.Logger, usecase usecase.UseCase) Server {
	return &server{
		log:     log,
		usecase: usecase,
	}
}

func (s *server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		s.log.Error("failed to jwt.ParseTokenFromHeader", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	req := &entity.CreateRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Token = token

	resp, err := s.usecase.Create(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Create", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
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

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.log.Error("failed to strconv.Atoi", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &entity.GetRequest{
		ID:    id,
		Token: token,
	}

	resp, err := s.usecase.Get(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Get", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleList(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		s.log.Error("failed to jwt.ParseTokenFromHeader", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	req := &entity.ListRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Token = token

	resp, err := s.usecase.List(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.List", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		s.log.Error("failed to jwt.ParseTokenFromHeader", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.log.Error("failed to strconv.Atoi", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &entity.UpdateRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.ID = id
	req.Token = token

	resp, err := s.usecase.Update(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Update", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleDelete(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		s.log.Error("failed to jwt.ParseTokenFromHeader", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.log.Error("failed to strconv.Atoi", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &entity.DeleteRequest{
		ID:    id,
		Token: token,
	}

	resp, err := s.usecase.Delete(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Delete", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleSaveFilters(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		s.log.Error("failed to jwt.ParseTokenFromHeader", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	req := &entity.SaveFiltersRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Token = token

	resp, err := s.usecase.SaveFilters(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.SaveFilters", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := json.WriteJSON(w, http.StatusOK, resp); err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}
