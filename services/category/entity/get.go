package entity

type (
	GetRequest struct {
		ID int `json:"id"`
	}

	GetResponse struct {
		Category *Category `json:"category"`
	}
)
