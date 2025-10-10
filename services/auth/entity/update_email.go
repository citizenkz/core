package entity

type (
	UpdateEmailRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateEmailResponse struct {
		Profile User `json:"profile"`
	}
)
