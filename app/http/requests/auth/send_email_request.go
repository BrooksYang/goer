package authRequest

import "github.com/goer-project/goer/form"

type SendEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
}

func (req SendEmailRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Email.required": "Email is required",
		"Email.email":    "Email is invalid",
	}
}
