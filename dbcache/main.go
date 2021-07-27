package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable:     true,
		CaseSensitive: true,
		AppName:       "Materi Rakamin v1.0.0",
	})
	addDatabaseHandler(app)
	addCacheHandler(app)
	addDBCacheHandler(app)

	chanServer := make(chan os.Signal, 1)
	signal.Notify(chanServer, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	go func() {
		<-chanServer

		log.Printf("Server is shutting down in the %s.", cfg.Host)
		err := app.Shutdown()
		if err != nil {
			log.Printf("Error in shutting down the server: %v.", err)
		}
	}()

	log.Printf("Server is running in the %s.", cfg.Host)
	log.Println("Press Ctrl + C to exit the server!")
	err := app.Listen(cfg.Host)
	if err != nil {
		log.Printf("Error in running the server: %v.", err)
	}
}
