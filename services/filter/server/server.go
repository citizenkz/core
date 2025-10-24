package server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/citizenkz/core/services/filter/entity"
	"github.com/citizenkz/core/services/filter/usecase"
	"github.com/citizenkz/core/utils/json"
	"github.com/citizenkz/core/utils/jwt"
	"github.com/go-chi/chi/v5"
)

type server struct {
	log     *slog.Logger
	usecase usecase.UseCase
}

type Server interface {
	List(w http.ResponseWriter, r *http.Request)
	SaveUserFitlers(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func New(log *slog.Logger, usecase usecase.UseCase) Server {
	return &server{
		log:     log,
		usecase: usecase,
	}
}

func (s *server) List(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()

	rawLimit := queryParam.Get("limit")
	rawOffset := queryParam.Get("offset")

	// Default values
	limit := 100
	offset := 0

	if rawLimit != "" {
		var err error
		limit, err = strconv.Atoi(rawLimit)
		if err != nil {
			s.log.Error("failed to strconv.Atoi", slog.String("error", err.Error()))
			json.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	if rawOffset != "" {
		var err error
		offset, err = strconv.Atoi(rawOffset)
		if err != nil {
			s.log.Error("failed to strconv.Atoi", slog.String("error", err.Error()))
			json.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}

	// Try to parse token from header (optional)
	token, _ := jwt.ParseTokenFromHeader(r)

	req := &entity.ListRequest{
		Token:       token,
		SearchQuery: queryParam.Get("search"),
		Limit:       limit,
		Offset:      offset,
	}

	resp, err := s.usecase.List(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.List", slog.String("error", err.Error()))
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

func (s *server) SaveUserFitlers(w http.ResponseWriter, r *http.Request) {
	token, err := jwt.ParseTokenFromHeader(r)
	if err != nil {
		s.log.Error("failed to jwt.ParseTokenFromHeader", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	req := &entity.SaveFilersRequest{}
	if err := json.ParseJSON(r, req); err != nil {
		s.log.Error("failed to json.ParseJSON", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	req.Token = token

	resp, err := s.usecase.SaveUserFilters(r.Context(), req)
	if err != nil {
		s.log.Error("failed to usecase.SaveUseFilters", slog.String("error", err.Error()))
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

func (s *server) Create(w http.ResponseWriter, r *http.Request) {
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
		s.log.Error("failed to json.WriteJson", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}

func (s *server) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		s.log.Error("failed to strconv.Atoi", slog.String("error", err.Error()))
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
		s.log.Error("failed to json.WriteJson", slog.String("error", err.Error()))
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}
}
