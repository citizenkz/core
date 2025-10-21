package entity

type (
	DeleteRequest struct {
		Password string `json:"password"`
		Token    string `json:"-"`
	}

	DeleteResponse struct {
		Success bool `json:"success"`
	}
)
