package models

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Title           string `gorm:"unique"`
	Year            int64
	Casts           []string `json:"cast" gorm:"-:all"`
	Genres          []string `json:"genres" gorm:"-:all"`
	Href            string
	Extract         string
	Thumbnail       string
	ThumbnailWidth  int16 `json:"thumbnail_width"`
	ThumbnailHeight int16 `json:"thumbnail_height"`
	Cast            string
	Genre           string
}

type MovieUpdate struct {
	Id      string
	Title   string
	Year    string
	Extract string
	Cast    string
	Genre   string
}

//
//func init() {
//	// jsonファイルを読み込む
//	jsonData, _ := os.ReadFile("./movieData.json")
//
//	// 構造体のスライスを作成
//	var movie []Movie
//	// *ここで何らかのエラーが起きる可能性が高い気がしますが、適宜調べて対応してみてください
//	err := json.Unmarshal([]byte(jsonData), &movie)
//	if err != nil {
//		fmt.Println("error:", err)
//	}
//
//}
