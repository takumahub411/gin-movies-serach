package main

import (
	"github.com/gin-gonic/gin"
	"movies-search-go/initializers"
	"movies-search-go/routes"
	"os"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectDb()
	initializers.SyncDatabase()
	initializers.SessionInit()
	// jsonファイルを読み込む
	//jsonData, _ := os.ReadFile("./movieData.json")
	//
	//// 構造体のスライスを作成
	//var movies []*models.Movie
	//
	//err := json.Unmarshal(jsonData, &movies)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//
	//for _, movie := range movies {
	//	movie.Cast = strings.Join(movie.Casts, ",")
	//	movie.Genre = strings.Join(movie.Genres, ",")
	//	initializers.DB.Create(&movie)
	//}

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.Static("images", "./images")

	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.MovieRoutes(r)
	routes.ReviewRoutes(r)
	routes.MovieManegeRoutes(r)
	r.Run(os.Getenv("port"))
}
