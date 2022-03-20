package authRequest

import "github.com/goer-project/goer/form"

type PasswordRequest struct {
	OldPassword          string `form:"old_password" binding:"required,min=6"`
	Password             string `form:"password" binding:"required,min=6"`
	PasswordConfirmation string `form:"password_confirmation" binding:"eqfield=Password"`
}

func (req PasswordRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"OldPassword.required":         "Old password is required",
		"OldPassword.min":              "Old password must be at least 6 characters",
		"Password.required":            "Password is required",
		"Password.min":                 "Password must be at least 6 characters",
		"PasswordConfirmation.eqfield": "Password confirmation does not match",
	}
}
