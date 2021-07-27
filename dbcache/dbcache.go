package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

const keyAllBooks = "books"

func addDBCacheHandler(app *fiber.App) {
	app.Get("/dbcache/book", func(c *fiber.Ctx) (err error) {
		var (
			errList   []Error
			byteBooks []byte
		)
		defer printJSONErr(c, &errList, &err)

		ctx := c.UserContext()
		byteBooks, err = cacheConn.Get(ctx, keyAllBooks).Bytes()
		if err == redis.Nil {
			log.Printf("All books are not found on the redis!")
			var books []Book
			err = dbConn.Find(&books).Error
			if err != nil {
				log.Printf("Error in getting all books: %v.", err)
				return
			}

			byteBooks, err = json.Marshal(books)
			if err != nil {
				log.Printf("Error in marshaling all books: %v.", err)
				return
			}

			err = cacheConn.Set(ctx, keyAllBooks, byteBooks, 5*time.Minute).Err()
			if err != nil {
				log.Printf("Error in set the redis data of all books: %v.", err)
				return
			}

			err = c.JSON(books)
			return
		}

		if err != nil {
			log.Printf("Error in getting all books from the redis: %v.", err)
			return
		}

		log.Printf("All books are found in the redis!")
		c.Type("json", "utf-8")
		c.Write(byteBooks)
		return
	})
}
