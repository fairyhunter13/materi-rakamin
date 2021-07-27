package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var cacheConn *redis.Client

const (
	keyPublisher = `publisher`
)

func initCache() {
	cacheConn = redis.NewClient(
		&redis.Options{
			Addr:     cfg.Cache.Address,
			Password: cfg.Cache.Password,
			DB:       cfg.Cache.DB,
		},
	)

	ctx := context.Background()
	cacheConn.Del(ctx, keyPublisher)
	err := cacheConn.Get(ctx, keyPublisher).Err()
	if err == redis.Nil {
		pub := Publisher{
			Name:              "Generic Publishing Company",
			FoundedIn:         time.Now().Add(-50 * 365 * 24 * time.Hour),
			NumberOfEmployees: 1000,
		}
		byteJson, _ := json.Marshal(&pub)
		err = cacheConn.Set(ctx, keyPublisher, string(byteJson), 0).Err()
		if err != nil {
			log.Printf("Error in setting the data for seeding: %v.", err)
			os.Exit(1)
		}
		return
	}
	if err != nil {
		log.Printf("Error in getting the data from the redis: %v.", err)
		os.Exit(1)
		return
	}
}

func addCacheHandler(app *fiber.App) {
	app.Get("/cache/publisher", func(c *fiber.Ctx) (err error) {
		var (
			errList []Error
			byteRes []byte
		)
		defer printJSONErr(c, &errList, &err)

		byteRes, err = cacheConn.Get(c.UserContext(), keyPublisher).Bytes()
		if err == redis.Nil {
			err = fmt.Errorf("data with key %s is not found", keyPublisher)
			return
		}

		if err != nil {
			log.Printf("Error in getting the publisher data: %v.", err)
			return
		}

		c.Type("json", "utf-8")
		c.Write(byteRes)
		return
	})
}
