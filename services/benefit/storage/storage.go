package storage

import (
	"context"
	"log/slog"

	"github.com/citizenkz/core/ent"
	"github.com/citizenkz/core/ent/benefit"
	"github.com/citizenkz/core/ent/benefitcategory"
	"github.com/citizenkz/core/ent/benefitfilter"
	"github.com/citizenkz/core/services/benefit/entity"
)

type storage struct {
	client *ent.Client
	log    *slog.Logger
}

type Storage interface {
	CreateBenefit(ctx context.Context, req *entity.CreateRequest) (*entity.BenefitWithFilters, error)
	GetBenefit(ctx context.Context, id int) (*entity.BenefitWithFilters, error)
	ListBenefits(ctx context.Context, req *entity.ListRequest) ([]*entity.BenefitWithFilters, int, error)
	UpdateBenefit(ctx context.Context, req *entity.UpdateRequest) (*entity.BenefitWithFilters, error)
	DeleteBenefit(ctx context.Context, id int) error
}

func New(client *ent.Client, log *slog.Logger) Storage {
	return &storage{
		client: client,
		log:    log,
	}
}

func (s *storage) CreateBenefit(ctx context.Context, req *entity.CreateRequest) (*entity.BenefitWithFilters, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		s.log.Error("failed to start transaction", slog.String("error", err.Error()))
		return nil, err
	}

	// Create the benefit
	benefitCreate := tx.Benefit.Create().
		SetTitle(req.Title).
		SetContent(req.Content).
		SetBonus(req.Bonus).
		SetNillableVideoURL(req.VideoURL).
		SetNillableSourceURL(req.SourceURL)

	benefit, err := benefitCreate.Save(ctx)
	if err != nil {
		s.log.Error("failed to save benefit", slog.String("error", err.Error()))
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
		}
		return nil, err
	}

	// Create benefit filters
	if len(req.Filters) > 0 {
		for _, filter := range req.Filters {
			_, err := tx.BenefitFilter.Create().
				SetBenefitID(benefit.ID).
				SetFilterID(filter.FilterID).
				SetNillableValue(filter.Value).
				SetNillableFrom(filter.From).
				SetNillableTo(filter.To).
				Save(ctx)
			if err != nil {
				s.log.Error("failed to save benefit filter", slog.String("error", err.Error()))
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
				}
				return nil, err
			}
		}
	}

	// Create benefit categories
	if len(req.Categories) > 0 {
		for _, categoryID := range req.Categories {
			_, err := tx.BenefitCategory.Create().
				SetBenefitID(benefit.ID).
				SetCategoryID(categoryID).
				Save(ctx)
			if err != nil {
				s.log.Error("failed to save benefit category", slog.String("error", err.Error()))
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
				}
				return nil, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		s.log.Error("failed to commit transaction", slog.String("error", err.Error()))
		return nil, err
	}

	// Fetch the created benefit with filters
	return s.GetBenefit(ctx, benefit.ID)
}

func (s *storage) GetBenefit(ctx context.Context, id int) (*entity.BenefitWithFilters, error) {
	benefit, err := s.client.Benefit.Query().
		Where(benefit.ID(id)).
		WithBenefitFilters().
		WithBenefitCategories(func(bcq *ent.BenefitCategoryQuery) {
			bcq.WithCategory()
		}).
		First(ctx)
	if err != nil {
		s.log.Error("failed to get benefit", slog.String("error", err.Error()))
		return nil, err
	}

	return entity.MakeStorageBenefitWithFiltersToEntity(benefit), nil
}

func (s *storage) ListBenefits(ctx context.Context, req *entity.ListRequest) ([]*entity.BenefitWithFilters, int, error) {
	query := s.client.Benefit.Query().
		WithBenefitFilters().
		WithBenefitCategories(func(bcq *ent.BenefitCategoryQuery) {
			bcq.WithCategory()
		})

	// Apply search filter
	if req.Search != "" {
		query = query.Where(
			benefit.Or(
				benefit.TitleContains(req.Search),
				benefit.ContentContains(req.Search),
				benefit.BonusContains(req.Search),
			),
		)
	}

	// Fetch all benefits first (we'll filter by criteria in application logic)
	allBenefits, err := query.All(ctx)
	if err != nil {
		s.log.Error("failed to list benefits", slog.String("error", err.Error()))
		return nil, 0, err
	}

	// Filter benefits based on filter criteria
	var filteredBenefits []*ent.Benefit
	if len(req.Filters) > 0 {
		for _, b := range allBenefits {
			if s.benefitMatchesFilters(b, req.Filters) {
				filteredBenefits = append(filteredBenefits, b)
			}
		}
	} else {
		filteredBenefits = allBenefits
	}

	total := len(filteredBenefits)

	// Apply pagination
	start := req.Offset
	if start > len(filteredBenefits) {
		start = len(filteredBenefits)
	}

	end := start + req.Limit
	if req.Limit == 0 || end > len(filteredBenefits) {
		end = len(filteredBenefits)
	}

	paginatedBenefits := filteredBenefits[start:end]

	result := make([]*entity.BenefitWithFilters, 0, len(paginatedBenefits))
	for _, b := range paginatedBenefits {
		result = append(result, entity.MakeStorageBenefitWithFiltersToEntity(b))
	}

	return result, total, nil
}

// benefitMatchesFilters checks if a benefit matches the filter criteria
// A benefit matches if:
// - It doesn't have a filter record for a given filter ID (passes through), OR
// - It has a filter record and the values match
func (s *storage) benefitMatchesFilters(benefit *ent.Benefit, filters []entity.FilterCriteria) bool {
	for _, filterCriteria := range filters {
		// Find if benefit has this filter
		hasFilter := false
		matchesValue := false

		if benefit.Edges.BenefitFilters != nil {
			for _, bf := range benefit.Edges.BenefitFilters {
				if bf.FilterID == filterCriteria.FilterID {
					hasFilter = true
					// Check if values match
					if s.filterValuesMatch(bf, &filterCriteria) {
						matchesValue = true
						break
					}
				}
			}
		}

		// If benefit has this filter but doesn't match the value, exclude it
		if hasFilter && !matchesValue {
			return false
		}
		// If benefit doesn't have this filter, it passes through (continue to next filter)
	}

	return true
}

// filterValuesMatch compares benefit filter values with criteria
func (s *storage) filterValuesMatch(bf *ent.BenefitFilter, criteria *entity.FilterCriteria) bool {
	// For single value filters
	if criteria.Value != nil {
		return bf.Value != nil && *bf.Value == *criteria.Value
	}

	// For range filters (from/to)
	if criteria.From != nil || criteria.To != nil {
		// Check if the benefit filter range overlaps with criteria range
		if bf.From != nil && bf.To != nil {
			// Both benefit and criteria use ranges
			if criteria.From != nil && criteria.To != nil {
				// Check for range overlap
				return !(*bf.To < *criteria.From || *bf.From > *criteria.To)
			}
			// Criteria has only from or to
			if criteria.From != nil {
				return *bf.To >= *criteria.From
			}
			if criteria.To != nil {
				return *bf.From <= *criteria.To
			}
		}
	}

	return false
}

func (s *storage) UpdateBenefit(ctx context.Context, req *entity.UpdateRequest) (*entity.BenefitWithFilters, error) {
	tx, err := s.client.Tx(ctx)
	if err != nil {
		s.log.Error("failed to start transaction", slog.String("error", err.Error()))
		return nil, err
	}

	// Update the benefit
	benefit, err := tx.Benefit.UpdateOneID(req.ID).
		SetTitle(req.Title).
		SetContent(req.Content).
		SetBonus(req.Bonus).
		SetNillableVideoURL(req.VideoURL).
		SetNillableSourceURL(req.SourceURL).
		Save(ctx)
	if err != nil {
		s.log.Error("failed to update benefit", slog.String("error", err.Error()))
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
		}
		return nil, err
	}

	// Delete existing benefit filters
	_, err = tx.BenefitFilter.Delete().
		Where(benefitfilter.BenefitID(benefit.ID)).
		Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete old benefit filters", slog.String("error", err.Error()))
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
		}
		return nil, err
	}

	// Create new benefit filters
	if len(req.Filters) > 0 {
		for _, filter := range req.Filters {
			_, err := tx.BenefitFilter.Create().
				SetBenefitID(benefit.ID).
				SetFilterID(filter.FilterID).
				SetNillableValue(filter.Value).
				SetNillableFrom(filter.From).
				SetNillableTo(filter.To).
				Save(ctx)
			if err != nil {
				s.log.Error("failed to save benefit filter", slog.String("error", err.Error()))
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
				}
				return nil, err
			}
		}
	}

	// Delete existing benefit categories
	if len(req.Categories) > 0 {
		_, err = tx.BenefitCategory.Delete().
			Where(benefitcategory.BenefitID(benefit.ID)).
			Exec(ctx)
		if err != nil {
			s.log.Error("failed to delete old benefit categories", slog.String("error", err.Error()))
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
			}
			return nil, err
		}
	}

	// Create new benefit categories
	if len(req.Categories) > 0 {
		for _, categoryID := range req.Categories {
			_, err := tx.BenefitCategory.Create().
				SetBenefitID(benefit.ID).
				SetCategoryID(categoryID).
				Save(ctx)
			if err != nil {
				s.log.Error("failed to save benefit category", slog.String("error", err.Error()))
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					s.log.Error("failed to rollback transaction", slog.String("error", rollbackErr.Error()))
				}
				return nil, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		s.log.Error("failed to commit transaction", slog.String("error", err.Error()))
		return nil, err
	}

	// Fetch the updated benefit with filters
	return s.GetBenefit(ctx, benefit.ID)
}

func (s *storage) DeleteBenefit(ctx context.Context, id int) error {
	err := s.client.Benefit.DeleteOneID(id).Exec(ctx)
	if err != nil {
		s.log.Error("failed to delete benefit", slog.String("error", err.Error()))
		return err
	}

	return nil
}
