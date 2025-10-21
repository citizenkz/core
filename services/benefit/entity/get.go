package entity

type (
	GetRequest struct {
		ID int `json:"id"`
	}

	GetResponse struct {
		Benefit *BenefitWithFilters `json:"benefit"`
	}
)
