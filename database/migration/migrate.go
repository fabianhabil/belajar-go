package migration

import (
	"belajar-go/database"
	"belajar-go/model/entity"
	"log"
)

func DBMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if err != nil {
		log.Println(err)
	}

	log.Println("Database Migrated")
}
