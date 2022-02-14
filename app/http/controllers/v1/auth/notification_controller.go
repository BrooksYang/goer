package auth

import (
	"fmt"
	"time"

	v1 "goer/app/http/controllers/v1"
	authRequest "goer/app/http/requests/auth"
	"goer/global"
	"goer/global/errno"
	"goer/pkg/helpers"
	"goer/pkg/http"
	"goer/pkg/mail"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	v1.BaseController
}

// SendEmailCode
// @Summary  Send Email Code
// @Tags     Auth
// @Accept   multipart/form-data
// @Produce  json
// @Param    email  formData  string  true  "Email"
// @Success  200    {string}  string  "OK"
// @Router   /v1/auth/code/email [POST]
func (*NotificationController) SendEmailCode(c *gin.Context) {
	var request authRequest.SendEmailRequest
	if ok := http.Validate(c, &request); !ok {
		return
	}

	// Generate code
	code := helpers.RandomNumber(global.Config.Code.Length)

	// Cache code
	global.Cache.Set("verify_code:"+request.Email, code, time.Minute*time.Duration(global.Config.Code.TTL))

	// Send email
	msg := fmt.Sprintf("Your verification code is: %s", code)
	res := mail.NewMailer().Send(request.Email, "Verification", msg)
	if !res {
		response.Fail(c, errno.ServiceMaintenance)
	}

	response.Success(c)
}
