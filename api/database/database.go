package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	Host     string
	Port     string
	Dbname   string
	Password string
	User     string
	Sslmode  string
)

func InitDatabase() {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s port=%s", Host, User, Dbname, Sslmode, Password, Port)
	db, err := gorm.Open(postgres.Open(dbUri))
	if err != nil {
		log.Fatalf("Failed to open database connection %v", err)
	}
	log.Println("Database initialized")
	DB = db
	MigrateModel()
}

func MigrateModel() {
	db := DB.Exec(HeroesDDL)
	db = DB.Exec(HeroesDescriptionDDL)
	if db.Error != nil {
		log.Fatal("Database error", db.Error)
	}
}
