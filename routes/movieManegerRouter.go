package routes

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/controllers"
)

func MovieManegeRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/manege", controllers.MoviesManegeGetHandler)
	incomingRoutes.GET("/manege/:id", controllers.MoviesEditGetHandler)

	incomingRoutes.POST("/manege/edit_confirm/:id", controllers.MoviesEditConfPostHandler)
	incomingRoutes.GET("/manege/delete/:id", controllers.MoviesDeleteGetHandler)
	incomingRoutes.GET("/edit_confirm", controllers.MoviesEditConfGetHandler)
}
