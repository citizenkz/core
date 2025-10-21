package entity

import (
	"github.com/citizenkz/core/ent"
)

type (
	Benefit struct {
		ID        int     `json:"id"`
		Title     string  `json:"title"`
		Content   string  `json:"content"`
		Bonus     string  `json:"bonus"`
		VideoURL  *string `json:"video_url"`
		SourceURL *string `json:"source_url"`
	}

	BenefitFilter struct {
		ID       int     `json:"id"`
		FilterID int     `json:"filter_id"`
		Value    *string `json:"value,omitempty"`
		From     *string `json:"from,omitempty"`
		To       *string `json:"to,omitempty"`
	}

	BenefitCategory struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description *string `json:"description,omitempty"`
	}

	BenefitWithFilters struct {
		ID         int                `json:"id"`
		Title      string             `json:"title"`
		Content    string             `json:"content"`
		Bonus      string             `json:"bonus"`
		VideoURL   *string            `json:"video_url"`
		SourceURL  *string            `json:"source_url"`
		Filters    []*BenefitFilter   `json:"filters,omitempty"`
		Categories []*BenefitCategory `json:"categories,omitempty"`
	}
)

func MakeStorageBenefitToEntity(benefit *ent.Benefit) *Benefit {
	return &Benefit{
		ID:        benefit.ID,
		Title:     benefit.Title,
		Content:   benefit.Content,
		Bonus:     benefit.Bonus,
		VideoURL:  benefit.VideoURL,
		SourceURL: benefit.SourceURL,
	}
}

func MakeStorageBenefitFilterToEntity(filter *ent.BenefitFilter) *BenefitFilter {
	return &BenefitFilter{
		ID:       filter.ID,
		FilterID: filter.FilterID,
		Value:    filter.Value,
		From:     filter.From,
		To:       filter.To,
	}
}

func MakeStorageBenefitWithFiltersToEntity(benefit *ent.Benefit) *BenefitWithFilters {
	result := &BenefitWithFilters{
		ID:         benefit.ID,
		Title:      benefit.Title,
		Content:    benefit.Content,
		Bonus:      benefit.Bonus,
		VideoURL:   benefit.VideoURL,
		SourceURL:  benefit.SourceURL,
		Filters:    make([]*BenefitFilter, 0),
		Categories: make([]*BenefitCategory, 0),
	}

	if benefit.Edges.BenefitFilters != nil {
		for _, filter := range benefit.Edges.BenefitFilters {
			result.Filters = append(result.Filters, MakeStorageBenefitFilterToEntity(filter))
		}
	}

	if benefit.Edges.BenefitCategories != nil {
		for _, bc := range benefit.Edges.BenefitCategories {
			if bc.Edges.Category != nil {
				result.Categories = append(result.Categories, &BenefitCategory{
					ID:          bc.Edges.Category.ID,
					Name:        bc.Edges.Category.Name,
					Description: bc.Edges.Category.Description,
				})
			}
		}
	}

	return result
}
