package entity

import (
	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/services/filter/consts"
)

type (
	Filter struct {
		ID            int               `json:"id"`
		Name          string            `json:"name"`
		Hint          *string           `json:"hint"`
		Type          consts.FilterType `json:"type"`
		Values        []string          `json:"values"`
		SelectedValue *string           `json:"selected_value,omitempty"`
	}

	UserFilters struct {
		UserID int             `json:"user_id"`
		Values []*FilterValues `json:"values"`
	}

	UserFilter struct {
		UserID   int    `json:"user_id"`
		FilterID int    `json:"filter_id"`
		Value    string `json:"value"`
	}

	FilterValues struct {
		FilterID int    `json:"filter_id"`
		Value    string `json:"value"`
	}
)

func MakeStorageFilterToEntity(filter *ent.Filter) *Filter {
	return &Filter{
		ID:     filter.ID,
		Name:   filter.Name,
		Hint:   filter.Hint,
		Type:   consts.FilterType(filter.Type),
		Values: filter.Values,
	}
}

func MakeStorageFilterSliceToEntity(filters []*ent.Filter) []*Filter {
	result := make([]*Filter, len(filters))

	for _, filter := range filters {
		filterEntity := MakeStorageFilterToEntity(filter)
		result = append(result, filterEntity)
	}

	return result
}

func MakeStorageUserFilterToEntity(userFilter *ent.UserFilter) *UserFilter {
	return &UserFilter{
		UserID:   userFilter.UserID,
		FilterID: userFilter.FilterID,
		Value:    userFilter.Value,
	}
}
