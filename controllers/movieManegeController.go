package controllers

import (
	"fmt"
	"movies-search-go/initializers"
	"movies-search-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MoviesManegeGetHandler(c *gin.Context) {

	var movies []*models.Movie

	result := initializers.DB.Find(&movies)

	if result.RowsAffected == 0 {

	}

	c.HTML(http.StatusOK, "manege.html", gin.H{"data": &movies})
}

func MoviesEditGetHandler(c *gin.Context) {

	id := c.Param("id")
	var movies []*models.Movie

	result := initializers.DB.First(&movies, id)

	if result.RowsAffected == 0 {
		//省略
	}
	fmt.Println(&movies)

	c.HTML(http.StatusOK, "edit.html", gin.H{"data": &movies})
}

func MoviesEditConfPostHandler(c *gin.Context) {

	var errMg []string

	if c.PostForm("title") == "" {
		errMg = append(errMg, "タイトルは必須項目です")
	}
	if c.PostForm("year") == "" {
		errMg = append(errMg, "公開年は必須項目です。")
	}
	if c.PostForm("cast") == "" {
		errMg = append(errMg, "キャストは必須項目です。")
	}
	if c.PostForm("extract") == "" {
		errMg = append(errMg, "説明文は必須項目です。")
	}
	if c.PostForm("genre") == "" {
		errMg = append(errMg, "ジャンルは必須項目です。")
	}

	Id := c.PostForm("id")
	Title := c.PostForm("title")
	Year := c.PostForm("year")
	Cast := c.PostForm("cast")
	Extract := c.PostForm("extract")
	Genre := c.PostForm("genre")

	c.HTML(http.StatusOK, "edit_confirm.html", gin.H{
		"error":   errMg,
		"Id":      Id,
		"Title":   Title,
		"Year":    Year,
		"Cast":    Cast,
		"Extract": Extract,
		"Genre":   Genre,
	})

	if len(errMg) == 0 {
		//update
		result := initializers.DB.
			Select([5]string{"title", "year", "cast", "extract", "genre"}).
			Where("id = ?", Id).
			Update(map[string]interface{}{"title": Title, "year": Year, "cast": Cast, "extract": Extract, "genre": Genre})
	}

}

func MoviesDeleteGetHandler(c *gin.Context) {
	id := c.Param("id")
	var movies []*models.Movie

	result := initializers.DB.Delete(&movies, id)

	fmt.Println(result)

	if result.Error != nil {
		return
	}

	c.HTML(http.StatusOK, "delete_Comp.html", nil)

}

func MoviesEditConfGetHandler(c *gin.Context) {

	var movies models.MovieUpdate

	movies.Id = c.PostForm("id")
	movies.Title = c.PostForm("title")
	movies.Year = c.PostForm("year")
	movies.Cast = c.PostForm("cast")
	movies.Extract = c.PostForm("extract")
	movies.Genre = c.PostForm("genre")

	c.HTML(http.StatusOK, "edit_confirm.html", gin.H{
		"data": movies,
	})
}
