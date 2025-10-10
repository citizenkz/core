package entity

type (
	UpdatePasswordRequest struct {
		OldPassword     string `json:"old_password"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	UpdatePasswordResponse struct {
		Profile User `json:"profile"`
	}
)
