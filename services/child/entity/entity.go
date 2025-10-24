package entity

import "time"

type Child struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ChildFilter struct {
	ID       int    `json:"id"`
	ChildID  int    `json:"child_id"`
	FilterID int    `json:"filter_id"`
	Value    string `json:"value"`
}
