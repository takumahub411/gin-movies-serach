package controllers

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/initializers"
	"movies-search-go/models"
	"net/http"
)

func MoviesGetHandler(c *gin.Context) {
	session, _ := initializers.Store.Get(c.Request, initializers.SessionName)
	username := session.Values["username"]

	//get movies
	var movies []*models.Movie

	result := initializers.DB.Find(&movies)

	if result.RowsAffected == 0 {

	}

	c.HTML(http.StatusOK, "movies.html", gin.H{
		"name": username,
		"data": &movies,
	})
}

func MovieGetHandler(c *gin.Context) {

	//get movie
	var movies []*models.Movie

	id := c.Param("id")

	result := initializers.DB.Find(&movies, id)

	if result.RowsAffected == 0 {

	}

	c.HTML(http.StatusOK, "movie.html", gin.H{
		"data": &movies,
	})
}

func MoviesPostHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "manege.html", nil)
}
