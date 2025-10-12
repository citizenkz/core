package entity

import "github.com/citizenkz/core/services/filter/consts"

type (
	UpdateRequest struct {
		ID     int               `json:"id"`
		Name   string            `json:"name"`
		Type   consts.FilterType `json:"type"`
		Hint   *string           `json:"hint"`
		Values []string          `json:"values"`
	}

	UpdateResponse struct {
		Filter Filter `json:"filter"`
	}
)
