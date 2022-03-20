package authRequest

import "github.com/goer-project/goer/form"

type PayPasswordRequest struct {
	PayPassword             string `form:"pay_password" json:"pay_password" binding:"required,min=6"`
	PayPasswordConfirmation string `form:"pay_password_confirmation" json:"pay_password_confirmation" binding:"eqfield=PayPassword"`
}

func (req PayPasswordRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"PayPassword.required":            "The pay password is required",
		"PayPassword.min":                 "The pay password must be at least 6 characters",
		"PayPasswordConfirmation.eqfield": "The pay password confirmation does not match",
	}
}
