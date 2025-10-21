package server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/citizenkz/core/services/category/entity"
	"github.com/citizenkz/core/services/category/usecase"
	"github.com/citizenkz/core/utils/json"
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
}

func New(log *slog.Logger, usecase usecase.UseCase) Server {
	return &server{
		log:     log,
		usecase: usecase,
	}
}

func (s *server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	req := &entity.CreateRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	resp, err := s.usecase.Create(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Create", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleGet(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.log.Error("failed to parse id", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &entity.GetRequest{
		ID: id,
	}

	resp, err := s.usecase.Get(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Get", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleList(w http.ResponseWriter, r *http.Request) {
	req := &entity.ListRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Set default limit if not provided
	if req.Limit == 0 {
		req.Limit = 10
	}

	resp, err := s.usecase.List(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.List", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.log.Error("failed to parse id", slog.String("error", err.Error()))
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

	resp, err := s.usecase.Update(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Update", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) HandleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.log.Error("failed to parse id", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req := &entity.DeleteRequest{
		ID: id,
	}

	resp, err := s.usecase.Delete(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.Delete", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = json.WriteJSON(w, http.StatusOK, resp)
	if err != nil {
		s.log.Error("failed to json.WriteJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}
