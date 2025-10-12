package entity

type (
	DeleteRequest struct {
		ID int `json:"id"`
	}

	DeleteResponse struct {
		IsDeleted bool `json:"is_deleted"`
	}
)
