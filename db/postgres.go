package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// using to call DatabseConnection func to create a db client for us
func init() {
	DatabaseConnection()
}

var DB *gorm.DB

type Movie struct {
	ID        string `gorm:"primerykey"`
	Title     string
	Genre     string
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpDatedAt time.Time `gorm:"autoupdatedTime:false"`
}

func DatabaseConnection() {
	host := "172.17.0.2"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "mypass"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(Movie{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")

}
