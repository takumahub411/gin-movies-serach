package routes

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/mypage", controllers.MypageGetHandler)
	incomingRoutes.GET("/passwordChange", controllers.PasswordChangeGetHandler)
	incomingRoutes.POST("/passwordChange", controllers.PasswordChangePostHandler)
	incomingRoutes.POST("/passwordReset", controllers.PasswordResetPostHandler)
	incomingRoutes.GET("/passwordReset", controllers.PasswordResetGetHandler)
}
