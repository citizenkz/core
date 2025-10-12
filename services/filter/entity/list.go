package entity

type (
	ListRequest struct {
		SearchQuery string `json:"search_query"`
		Limit       int    `json:"limit"`
		Offset      int    `json:"offset"`
	}

	ListResponse struct {
		Filters []*Filter `json:"filter"`
	}
)
