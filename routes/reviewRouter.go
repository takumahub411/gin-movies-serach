package routes

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/controllers"
)

func ReviewRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/movies/:id", controllers.ReviewPostHandler)
}
