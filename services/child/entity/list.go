package entity

type ListRequest struct {
	Token  string
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListResponse struct {
	Children []*Child `json:"children"`
	Total    int      `json:"total"`
}
