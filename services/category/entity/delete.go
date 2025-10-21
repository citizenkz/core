package entity

type (
	DeleteRequest struct {
		ID int `json:"id"`
	}

	DeleteResponse struct {
		Success bool `json:"success"`
	}
)
