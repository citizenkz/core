package entity

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Profile User   `json:"profile"`
		Token   string `json:"token"`
	}
)
