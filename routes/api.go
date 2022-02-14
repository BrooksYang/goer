package routes

import (
	"goer/app/http/controllers/v1/auth"
	"goer/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func MapApiRoutes(r *gin.Engine) {

	// v1 group
	v1 := r.Group("/v1")

	/**
	|--------------------------------------------------------------------------
	| Auth
	|--------------------------------------------------------------------------
	|
	| Here is where you can register API routes for your application. These
	| routes are loaded by the RouteServiceProvider within a group which
	| is assigned the "api" middleware group. Enjoy building your API!
	|
	*/

	// Create controller instance
	authController := auth.AuthController{}
	passwordController := auth.PasswordController{}
	notificationController := auth.NotificationController{}

	// Send email code
	v1.POST("auth/code/email", middleware.LimitPath("1-M"), notificationController.SendEmailCode)

	// Register
	v1.POST("auth/register", authController.Register)

	// Login
	v1.POST("auth/login", authController.Login)

	v1.Use(middleware.Auth())
	{
		// Profile
		v1.GET("auth/profile", authController.Profile)

		// Update Profile
		v1.PATCH("auth/profile", authController.UpdateProfile)

		// Update Password
		v1.PATCH("auth/password", passwordController.UpdatePassword)

		// Set pay password
		v1.POST("auth/payPassword", passwordController.SetPayPassword)

		// Reset pay password
		v1.POST("auth/payPassword/reset", passwordController.ResetPayPassword)
	}
}
