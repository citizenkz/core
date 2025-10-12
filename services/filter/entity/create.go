package entity

import "github.com/citizenkz/core/services/filter/consts"

type (
	CreateRequest struct {
		Name   string            `json:"name"`
		Type   consts.FilterType `json:"type"`
		Hint   *string           `json:"hint"`
		Values []string          `json:"values"`
	}

	CreateResponse struct {
		Filter Filter `json:"filter"`
	}
)
