package authRequest

import "github.com/goer-project/goer/form"

type ProfileRequest struct {
	Username string `form:"username" binding:"alphanum"`
	Gender   string `form:"gender" binding:"oneof=male female secret"`
	Age      int64  `form:"age" binding:"number,gt=0"`
	Avatar   string `form:"avatar"`
}

func (req ProfileRequest) Messages() form.ValidatorMessages {
	return form.ValidatorMessages{
		"Username.alphanum": "Username must only contain letters and numbers",
		"Gender.oneof":      "The Gender is invalid",
		"Age.number":        "The age must be a number",
		"Age.gt":            "The age must be greater than 0",
	}
}
