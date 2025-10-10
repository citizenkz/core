package entity

type (
	GetRequest struct {
		Token string
	}

	GetResponse struct {
		Profile User `json:"profile"`
	}
)
