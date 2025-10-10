package entity

type (
	ForgetPasswordRequest struct {
		Email string `json:"email"`
	}

	ForgetPasswordResponse struct {
		AttemptID int `json:"attempt_id"`
		RetryTime int `json:"retry_time"`
	}

	ForgetPasswordConfirmRequest struct {
		AttemptID int    `json:"attempt_id"`
		OtpCode   string `json:"otp_code"`
	}

	ForgetPasswordConfirmResponse struct {
		Profile User `json:"profile"`
	}
)
