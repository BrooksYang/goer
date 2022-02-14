package auth

import (
	v1 "goer/app/http/controllers/v1"
	authRequest "goer/app/http/requests/auth"
	"goer/app/models/user"
	"goer/global"
	"goer/global/errno"
	"goer/pkg/auth"
	"goer/pkg/http"
	"goer/pkg/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	v1.BaseController
}

// Register
// @Summary      Register
// @Tags         Auth
// @Accept       multipart/form-data
// @Produce      json
// @Param        type         formData  integer        true   "Register type, 1-email, 2-phone"
// @Param        name         formData  string         false  "Name"
// @Param        email        formData  string         false  "Email"
// @Param        phone        formData  string         false  "Phone"
// @Param        password     formData  string         true   "Password"
// @Param        referral_id  formData  int            false  "Referral id"
// @Success      200          {object}  auth.JwtToken  "Register"
// @Router       /v1/auth/register [post]
func (a *AuthController) Register(c *gin.Context) {
	var request authRequest.RegisterRequest
	if ok := http.Validate(c, &request); !ok {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	// user
	newUser := user.User{
		Uid:          user.GenerateUid(),
		Name:         request.Name,
		Email:        request.Email,
		Phone:        request.Phone,
		Password:     string(password),
		GoogleStatus: string(user.GoogleStatusUnbind),
		IsValid:      true,
		Gender:       string(user.GenderSecret),
	}

	// Check if account exists
	if newUser.AccountExists() {
		response.Fail(c, errno.AccountExists)
		return
	}

	// Get parent
	var parent user.User
	global.DB.Where("uid", request.ReferralId).Limit(1).Find(&parent)

	err := global.DB.Transaction(func(tx *gorm.DB) (err error) {
		// parent invite count
		if parent.ID > 0 {
			parent.InviteCount++
			err = tx.Select("invite_count").Save(&parent).Error
			if err != nil {
				return
			}
		}

		// pid & depth
		newUser.Pid = parent.ID
		newUser.Depth = parent.Depth + 1

		// Create
		err = tx.Create(&newUser).Error
		if err != nil {
			return
		}

		return nil
	})

	if err != nil {
		response.Fail(c, errno.InternalServerError)
		return
	}

	// generate token
	token, _ := auth.NewJWT().CreateToken(newUser.ID)

	// Single sign on
	newUser.Sso = token.AccessToken
	global.DB.Select("sso").Save(&newUser)

	response.Data(c, token)
}

// Login
// @Summary      Login
// @Tags         Auth
// @Accept       multipart/form-data
// @Produce      json
// @Param        account   formData  string         true  "Account"
// @Param        password  formData  string         true  "Password"
// @Success      200       {object}  auth.JwtToken  "login"
// @Router       /v1/auth/login [post]
func (a AuthController) Login(c *gin.Context) {
	var request authRequest.LoginRequest
	if ok := http.Validate(c, &request); !ok {
		return
	}

	// Find user
	var authUser user.User
	global.DB.Where("email = ?", request.Account).
		Or("phone = ?", request.Account).
		First(&authUser)
	if authUser.ID == 0 {
		response.Fail(c, errno.InvalidAccount)
		return
	}

	// account locked
	if !authUser.IsValid {
		response.Fail(c, errno.AccountLocked)
		return
	}

	// check password
	res := authUser.CheckPassword(request.Password)
	if !res {
		response.Fail(c, errno.InvalidAccount)
		return
	}

	// generate token
	token, _ := auth.NewJWT().CreateToken(authUser.ID)

	// Single sign on
	authUser.Sso = token.AccessToken
	global.DB.Select("sso").Save(&authUser)

	response.Data(c, token)
}

// Profile
// @Summary      Profile
// @Description  Get user profile
// @Security     Bearer
// @Tags         Auth
// @Accept       multipart/form-data
// @Produce      json
// @Success      200  {object}  user.User  "Get user profile"
// @Router       /v1/auth/profile [get]
func (a AuthController) Profile(c *gin.Context) {
	// Find user
	authUser := auth.User(c)

	response.Data(c, authUser)
}

// UpdateProfile
// @Summary      Update profile
// @Description  Update user profile
// @Security     Bearer
// @Tags         Auth
// @Accept       multipart/form-data
// @Produce      json
// @Param        username  formData  string     false  "Username"
// @Param        gender    formData  string     false  "secret,male,female"
// @Param        age       formData  int        false  "Age"
// @Param        avatar    formData  string     false  "Avatar"
// @Success      200       {object}  user.User  "Update user profile"
// @Router       /v1/auth/profile [PATCH]
func (a AuthController) UpdateProfile(c *gin.Context) {
	var request authRequest.ProfileRequest
	if ok := http.Validate(c, &request); !ok {
		return
	}

	// Find user
	authUser := auth.User(c)

	// Update profile
	if request.Username != "" {
		authUser.Username = request.Username
	}
	if request.Gender != "" {
		authUser.Gender = request.Gender
	}
	if request.Age != 0 {
		authUser.Age = request.Age
	}
	if request.Avatar != "" {
		authUser.Avatar = request.Avatar
	}

	if request.Username != "" || request.Gender != "" || request.Age != 0 || request.Avatar != "" {
		global.DB.Select("Username", "Gender", "Age", "Avatar").Save(&authUser)
	}

	response.Data(c, authUser)
}
