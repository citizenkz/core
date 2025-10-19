package entity

type (
	SaveFilersRequest struct {
		Token        string          `json:"user_id"`
		FilterValues []*FilterValues `json:"filter_values"`
	}

	SaveFilterResponse struct {
		UserFilters UserFilters `json:"user_filters"`
	}
)
