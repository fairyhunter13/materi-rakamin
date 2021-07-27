package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func initDB() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	var err error
	dbConn, err = gorm.Open(mysql.Open(connStr), &gorm.Config{
		SkipDefaultTransaction:                   true,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Printf("Error in creating the new connection: %v.", err)
		os.Exit(1)
		return
	}

	initDBData()
}

func initDBData() {
	err := dbConn.AutoMigrate(new(Book))
	if err != nil {
		log.Printf("Error in migrating the models: %v.", err)
		os.Exit(1)
		return
	}

	var book Book
	err = dbConn.Take(&book).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var books = []Book{
			{
				Title:         "The Life",
				Topic:         "Slice of Life",
				Author:        "Kyouya",
				DatePublished: time.Now().Add(-2 * 365 * 24 * time.Hour),
			},
			{
				Title:         "Logic Analysis",
				Topic:         "Mathematics",
				Author:        "Steven",
				DatePublished: time.Now().Add(-3 * 365 * 24 * time.Hour),
			},
			{
				Title:         "Art of Love",
				Topic:         "Romance",
				Author:        "Ai",
				DatePublished: time.Now().Add(-10 * 365 * 24 * time.Hour),
			},
		}
		err = dbConn.Create(&books).Error
		if err != nil {
			log.Printf("Error in seeding the data: %v.", err)
			os.Exit(1)
		}
		return
	}
	if err != nil {
		log.Printf("Error in checking if the data has been seeded: %v.", err)
		os.Exit(1)
		return
	}
}

func addDatabaseHandler(app *fiber.App) {
	app.Get("/db/book", func(c *fiber.Ctx) (err error) {
		var (
			books   []Book
			errList []Error
		)
		defer printJSONErr(c, &errList, &err)

		err = dbConn.Find(&books).Error
		if err != nil {
			log.Printf("Error in getting the books data: %v.", err)
			return
		}

		err = c.JSON(books)
		return
	})
}
