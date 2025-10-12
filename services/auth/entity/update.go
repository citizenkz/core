package entity

import "time"

type (
	UpdateRequest struct {
		ID        int        `json:"user_id"`
		FirstName string     `json:"first_name"`
		LastName  string     `json:"last_name"`
		BirthDate *time.Time `json:"birth_date"`
	}

	UpdateResponse struct {
		Profile User `json:"profile"`
	}
)
