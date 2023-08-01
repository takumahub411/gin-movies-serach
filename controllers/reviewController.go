package controllers

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/initializers"
	"movies-search-go/models"
	"net/http"
	"strconv"
	"time"
)

func ReviewPostHandler(c *gin.Context) {
	movieId, _ := strconv.Atoi(c.Param("id"))

	session, _ := initializers.Store.Get(c.Request, initializers.SessionName)
	userId := session.Values["userId"].(string)

	var errMg []string

	if title := c.PostForm("title"); title == "" {
		errMg = append(errMg, "タイトルが入力されていません")
	}
	if content := c.PostForm("content"); content == "" {
		errMg = append(errMg, "レヴュー内容が入力されていません")
	}
	if star := c.PostForm("star"); star == "" {
		errMg = append(errMg, "星の数を選択してください")
	}
	star, _ := strconv.Atoi(c.PostForm("star"))

	if len(errMg) == 0 {
		//post review
		data := models.Review{
			ReviewId:  userId,
			MovieId:   movieId,
			Title:     c.PostForm("title"),
			Content:   c.PostForm("content"),
			Stars:     star,
			CreatedAt: time.Date(2020, 1, 2, 3, 4, 5, 123456789, time.Local),
		}
		initializers.DB.Create(&data)
	} else {
		c.HTML(http.StatusOK, "movies.html", gin.H{
			"error": errMg,
		})
		return
	}

}
