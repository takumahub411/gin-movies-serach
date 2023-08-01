package routes

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/controllers"
)

func MovieRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/movies", controllers.MoviesGetHandler)
	incomingRoutes.GET("/movies/:id", controllers.MovieGetHandler)
	//incomingRoutes.POST("/movies/:id", controllers.PostMovie())
	//incomingRoutes.PUT("/movies/:id", controllers.UpdateMovie())
	//incomingRoutes.DELETE("/movies/:id", controllers.DeleteMovie())
}
