package auth

import (
	v1 "goapp/app/http/controllers/v1"
	authRequest "goapp/app/http/requests/auth"
	"goapp/global"
	"goapp/global/errno"
	"goapp/pkg/auth"
	"goapp/pkg/form"
	"goapp/pkg/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PasswordController struct {
	v1.BaseController
}

// UpdatePassword
// @Summary   Update password
// @Security  Bearer
// @Tags      Auth
// @Accept    multipart/form-data
// @Produce   json
// @Param     old_password           formData  string  true  "old password"
// @Param     password               formData  string  true  "new password"
// @Param     password_confirmation  formData  string  true  "new password confirmation"
// @Success   200                        {string}  string  "OK"
// @Router    /v1/auth/password [PATCH]
func (a PasswordController) UpdatePassword(c *gin.Context) {
	var request authRequest.PasswordRequest
	if ok := form.Validate(c, &request); !ok {
		return
	}

	// Find user
	authUser := auth.User(c)

	// Check password
	res := authUser.CheckPassword(request.OldPassword)
	if !res {
		response.Fail(c, errno.InvalidPassword)
		return
	}

	// Update password
	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	authUser.Password = string(password)
	global.DB.Select("Password").Save(&authUser)

	response.Success(c)
}

// SetPayPassword
// @Summary   Set pay password
// @Security  Bearer
// @Tags      Auth
// @Accept    multipart/form-data
// @Produce   json
// @Param     pay_password               formData  string  true  "Pay password"
// @Param     pay_password_confirmation  formData  string  true  "Pay password confirmation"
// @Success   200                        {string}  string  "OK"
// @Router    /v1/auth/payPassword [POST]
func (a PasswordController) SetPayPassword(c *gin.Context) {
	var request authRequest.PayPasswordRequest
	if ok := form.Validate(c, &request); !ok {
		return
	}

	// Find user
	authUser := auth.User(c)

	// Has pay password
	if authUser.HasPayPassword() {
		response.Fail(c, errno.PayPasswordExists)
		return
	}

	// Update password
	payPassword, _ := bcrypt.GenerateFromPassword([]byte(request.PayPassword), bcrypt.DefaultCost)
	authUser.PayPassword = string(payPassword)
	global.DB.Select("PayPassword").Save(&authUser)

	response.Success(c)
}

// ResetPayPassword
// @Summary   Reset pay password
// @Security  Bearer
// @Tags      Auth
// @Accept    multipart/form-data
// @Produce   json
// @Param     password                   formData  string  true  "Password"
// @Param     pay_password               formData  string  true  "Pay password"
// @Param     pay_password_confirmation  formData  string  true  "Pay password confirmation"
// @Success   200                    {string}  string  "OK"
// @Router    /v1/auth/payPassword/reset [POST]
func (a PasswordController) ResetPayPassword(c *gin.Context) {
	var request authRequest.PayPasswordReset
	if ok := form.Validate(c, &request); !ok {
		return
	}

	// Find user
	authUser := auth.User(c)

	// Check password
	if !authUser.CheckPassword(request.Password) {
		response.Fail(c, errno.InvalidPassword)
		return
	}

	// Update pay password
	payPassword, _ := bcrypt.GenerateFromPassword([]byte(request.PayPassword), bcrypt.DefaultCost)
	authUser.PayPassword = string(payPassword)
	global.DB.Select("PayPassword").Save(&authUser)

	response.Success(c)
}
