package entity

type (
	GetRequest struct{}

	GetResponse struct {
		Profile User `json:"profile"`
	}
)
