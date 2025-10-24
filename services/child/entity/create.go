package entity

import "time"

type CreateRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Token     string
}

type CreateResponse struct {
	Child *Child `json:"child"`
}
