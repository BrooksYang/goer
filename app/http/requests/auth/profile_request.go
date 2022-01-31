package authRequest

import "goer/pkg/http"

type ProfileRequest struct {
	Username string `form:"username" binding:"alphanum"`
	Gender   string `form:"gender" binding:"oneof=male female secret"`
	Age      int64  `form:"age" binding:"number,gt=0"`
	Avatar   string `form:"avatar"`
}

func (req ProfileRequest) Messages() http.ValidatorMessages {
	return http.ValidatorMessages{
		"Username.alphanum": "Username must only contain letters and numbers",
		"Gender.oneof":      "The Gender is invalid",
		"Age.number":        "The age must be a number",
		"Age.gt":            "The age must be greater than 0",
	}
}
