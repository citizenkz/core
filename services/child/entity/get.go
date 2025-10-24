package entity

type GetRequest struct {
	ID    int
	Token string
}

type GetResponse struct {
	Child *Child `json:"child"`
}
