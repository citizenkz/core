package entity

type (
	UpdateRequest struct {
		ID         int                    `json:"id"`
		Title      string                 `json:"title"`
		Content    string                 `json:"content"`
		Bonus      string                 `json:"bonus"`
		VideoURL   *string                `json:"video_url,omitempty"`
		SourceURL  *string                `json:"source_url,omitempty"`
		Filters    []BenefitFilterRequest `json:"filters,omitempty"`
		Categories []int                  `json:"categories,omitempty"`
	}

	UpdateResponse struct {
		Benefit *BenefitWithFilters `json:"benefit"`
	}
)
