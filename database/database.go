package database

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"restAPI/models"

	"github.com/joho/godotenv"
)

var MyDB *gorm.DB

func SetupPostgres() error {
	_ = godotenv.Load("./database/.env")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dsn := "host=localhost user=postgres password=" + password + " dbname=" + dbname + " port=5432 sslmode=disable"
	var err error
	MyDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = MyDB.AutoMigrate(&models.Team{}, &models.Player{})
	if err != nil {
		return errors.New("Failed to automigrate models: " + err.Error())
	}

	return nil
}

func ClosePostgres() {
	db, err := MyDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
