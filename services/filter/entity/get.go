package entity

type (
	GetRequest struct {
		ID int `json:"id"`
	}

	GetResponse struct {
		Fitler Filter `json:"filter"`
	}
)
