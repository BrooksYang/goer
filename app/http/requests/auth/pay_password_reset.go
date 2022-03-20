package authRequest

import "github.com/goer-project/goer/form"

type PayPasswordReset struct {
	Password                string `form:"password" json:"password" binding:"required"`
	PayPassword             string `form:"pay_password" json:"pay_password" binding:"required,min=6"`
	PayPasswordConfirmation string `form:"pay_password_confirmation" json:"pay_password_confirmation" binding:"eqfield=PayPassword"`
}

func (req PayPasswordReset) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Password.required":               "Password is required",
		"PayPassword.required":            "The pay password is required",
		"PayPassword.min":                 "The pay password must be at least 6 characters",
		"PayPasswordConfirmation.eqfield": "The pay password confirmation does not match",
	}
}
