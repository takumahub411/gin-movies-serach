package routes

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/controllers"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("signup", controllers.SignupGetHandler)
	r.POST("signup", controllers.SignupPostHandler)
	r.GET("success_signup", controllers.SignupCompGetHandler)
	r.GET("/emailVerify/:username/:id", controllers.EmailVerifyGetHandler)
	r.GET("/emailCheck", controllers.EmailCheckGetHandler)
	r.POST("/emailCheck", controllers.EmailCheckPostHandler)
	r.GET("/passwordReset/:username/:id", controllers.PasswordResetGetHandler)
	r.POST("/passwordReset/:username/:id", controllers.PasswordResetPostHandler)
	r.GET("/success_passwordReset", controllers.SuccessPasswordResetPostHandler)
	r.GET("login", controllers.LoginGetHandler)
	r.POST("login", controllers.LoginPostHandler)

}
