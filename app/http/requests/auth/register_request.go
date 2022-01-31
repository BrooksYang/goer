package authRequest

import "goer/pkg/http"

type RegisterRequest struct {
	Type       int    `form:"type" binding:"required,eq=1|eq=2"` // 1-邮箱注册，2-手机注册
	Name       string `form:"name"`
	Email      string `form:"email" binding:"required_if=Type 1,email"`
	Phone      string `form:"phone" binding:"required_if=Type 2"`
	Password   string `form:"password" binding:"required,min=6"`
	ReferralId uint   `form:"referral_id"`
}

func (req RegisterRequest) Messages() http.ValidatorMessages {
	return http.ValidatorMessages{
		"Type.required":       "Type is required",
		"Type.eq=1|eq=2":      "Type is invalid",
		"Email.required_if":   "Email is required",
		"Email.email":         "Email is invalid",
		"Phone.required_if":   "Phone is required",
		"Password.required":   "Password is required",
		"Password.min":        "Password must be at least 6 characters",
		"ReferralId.required": "Referral id is required",
	}
}
