package authRequest

import "github.com/goer-project/goer/form"

type RegisterRequest struct {
	Type       int    `form:"type" binding:"required,oneof=1 2"` // 1-邮箱注册，2-手机注册
	Name       string `form:"name"`
	Email      string `form:"email" binding:"required_if=Type 1,omitempty,email"`
	Phone      string `form:"phone" binding:"required_if=Type 2"`
	Password   string `form:"password" binding:"required,min=6"`
	ReferralId uint   `form:"referral_id"`
}

func (req RegisterRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Type.required":       "Type is required",
		"Type.oneof":          "Type is invalid",
		"Email.required_if":   "Email is required",
		"Email.email":         "Email is invalid",
		"Phone.required_if":   "Phone is required",
		"Password.required":   "Password is required",
		"Password.min":        "Password must be at least 6 characters",
		"ReferralId.required": "Referral id is required",
	}
}
