package entity

type (
	ListRequest struct {
		Limit  int    `json:"limit"`
		Offset int    `json:"offset"`
		Search string `json:"search,omitempty"`
	}

	ListResponse struct {
		Categories []*Category `json:"categories"`
		Total      int         `json:"total"`
	}
)
