package entity

type (
	BenefitFilterRequest struct {
		FilterID int     `json:"filter_id"`
		Value    *string `json:"value,omitempty"`
		From     *string `json:"from,omitempty"`
		To       *string `json:"to,omitempty"`
	}

	CreateRequest struct {
		Title      string                 `json:"title"`
		Content    string                 `json:"content"`
		Bonus      string                 `json:"bonus"`
		VideoURL   *string                `json:"video_url,omitempty"`
		SourceURL  *string                `json:"source_url,omitempty"`
		Filters    []BenefitFilterRequest `json:"filters,omitempty"`
		Categories []int                  `json:"categories,omitempty"`
	}

	CreateResponse struct {
		Benefit *BenefitWithFilters `json:"benefit"`
	}
)
