package entity

type (
	FilterCriteria struct {
		FilterID int     `json:"filter_id"`
		Value    *string `json:"value,omitempty"`
		From     *string `json:"from,omitempty"`
		To       *string `json:"to,omitempty"`
	}

	ListRequest struct {
		Limit   int              `json:"limit"`
		Offset  int              `json:"offset"`
		Search  string           `json:"search,omitempty"`
		Filters []FilterCriteria `json:"filters,omitempty"`
	}

	ListResponse struct {
		Benefits []*BenefitWithFilters `json:"benefits"`
		Total    int                   `json:"total"`
	}
)
