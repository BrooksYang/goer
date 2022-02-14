package authRequest

import "goer/pkg/http"

type SendEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
}

func (req SendEmailRequest) Messages() http.ValidatorMessages {
	return http.ValidatorMessages{
		"Email.required": "Email is required",
		"Email.email":    "Email is invalid",
	}
}
