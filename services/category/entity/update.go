package entity

type (
	UpdateRequest struct {
		ID          int     `json:"id"`
		Name        string  `json:"name"`
		Description *string `json:"description,omitempty"`
	}

	UpdateResponse struct {
		Category *Category `json:"category"`
	}
)
