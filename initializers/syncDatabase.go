package initializers

import "movies-search-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Movie{})
	DB.AutoMigrate(&models.Review{})
}
