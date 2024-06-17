package migration

import (
	"fmt"
	"learn-go-fiber/database"
	"learn-go-fiber/model/entity"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})

	if err != nil {
		log.Fatal("Failed to run migration: ", err)
	}

	fmt.Println("Successfully migrated the database.")
}
