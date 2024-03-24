package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	postModels "instaclone/src/applicatation/posts/models"
	userModels "instaclone/src/applicatation/users/models"
)

func Initialize() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable timezone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	if err := db.AutoMigrate(&userModels.UserModel{}, &postModels.PostModel{}); err != nil {
		panic("Could not migrate the database")
	}

	return db
}
