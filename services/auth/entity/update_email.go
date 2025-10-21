package entity

type (
	UpdateEmailRequest struct {
		Token    string `json:"-"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateEmailResponse struct {
		Profile User `json:"profile"`
	}
)
