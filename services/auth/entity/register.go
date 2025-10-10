package entity

import "time"

type (
	RegisterRequest struct {
		FirstName       string    `json:"first_name"`
		LastName        string    `json:"last_name"`
		Email           string    `json:"email"`
		Password        string    `json:"password"`
		ConfirmPassword string    `json:"confirm_password"`
		BirthDate       time.Time `json:"birth_date"`
	}

	RegisterResponse struct {
		Profile User   `json:"profile"`
		Token   string `json:"token"`
	}
)
