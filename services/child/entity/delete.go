package entity

type DeleteRequest struct {
	ID    int
	Token string
}

type DeleteResponse struct {
	Success bool `json:"success"`
}
