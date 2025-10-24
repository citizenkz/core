package entity

type SaveFiltersRequest struct {
	ChildID int                  `json:"child_id"`
	Filters []FilterValueRequest `json:"filters"`
	Token   string
}

type FilterValueRequest struct {
	FilterID int    `json:"filter_id"`
	Value    string `json:"value"`
}

type SaveFiltersResponse struct {
	Success bool `json:"success"`
}
