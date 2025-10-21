package entity

import (
	"github.com/citizenkz/core/ent"
)

type (
	Category struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description *string `json:"description,omitempty"`
	}
)

func MakeStorageCategoryToEntity(category *ent.Category) *Category {
	return &Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
