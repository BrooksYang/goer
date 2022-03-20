package authRequest

import "github.com/goer-project/goer/form"

type LoginRequest struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required,min=6"`
}

func (req LoginRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Account.required":  "Account is required",
		"Password.required": "Password is required",
		"Password.min":      "Password must be at least 6 characters",
	}
}
