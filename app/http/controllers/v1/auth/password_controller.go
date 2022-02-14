package auth

import (
	v1 "goer/app/http/controllers/v1"
	authRequest "goer/app/http/requests/auth"
	"goer/global"
	"goer/global/errno"
	"goer/pkg/auth"
	"goer/pkg/http"
	"goer/pkg/response"

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
// @Success   200                    {string}  string  "OK"
// @Router    /v1/auth/password [PATCH]
func (a PasswordController) UpdatePassword(c *gin.Context) {
	var request authRequest.PasswordRequest
	if ok := http.Validate(c, &request); !ok {
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
