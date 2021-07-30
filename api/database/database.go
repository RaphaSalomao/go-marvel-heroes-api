package database

import (
	"log"

	"br.com.github/raphasalomao/go-marvel-heroes-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dbUri := "host=localhost user=postgres dbname=marvel-api sslmode=disable password=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dbUri))
	if err != nil {
		log.Fatalf("Failed to open database connection %v", err)
	}
	log.Println("Database initialized")
	DB = db
	MigrateModel()
}

func MigrateModel() {
	db := DB.Exec(model.HeroesDDL)
	if db.Error != nil {
		log.Fatal("Database error", db.Error)
	}
}
