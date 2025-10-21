package entity

type (
	CreateRequest struct {
		Name        string  `json:"name"`
		Description *string `json:"description,omitempty"`
	}

	CreateResponse struct {
		Category *Category `json:"category"`
	}
)
