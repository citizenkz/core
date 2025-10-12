package entity

type (
	SaveFilersRequest struct {
		UserID       int             `json:"user_id"`
		FilterValues []*FilterValues `json:"filter_values"`
	}

	SaveFilterResponse struct {
		UserFilters UserFilters `json:"user_filters"`
	}
)
