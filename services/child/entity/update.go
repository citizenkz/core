package entity

import "time"

type UpdateRequest struct {
	ID        int
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Token     string
}

type UpdateResponse struct {
	Child *Child `json:"child"`
}
