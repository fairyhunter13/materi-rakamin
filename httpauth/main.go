package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

// Docs: [![Run in Postman](https://run.pstmn.io/button.svg)]
// (https://app.getpostman.com/run-collection/5f5201d4f6c0699ca078?action=collection%2Fimport)

type User struct {
	Name   string
	Email  string
	Gender string
}

var (
	users = []User{
		{
			Name:   "John",
			Email:  "john@example.com",
			Gender: "male",
		},
		{
			Name:   "Sarah",
			Email:  "sarah@example.com",
			Gender: "female",
		},
		{
			Name:   "Alvin",
			Email:  "alvin@example.com",
			Gender: "male",
		},
	}
)

func main() {
	app := fiber.New(fiber.Config{
		Immutable:     true,
		CaseSensitive: true,
		AppName:       "Materi Rakamin v1.0.0",
	})
	addBasicRoute(app)
	addJWTRoute(app)
	log.SetFlags(log.Llongfile)

	chanServer := make(chan os.Signal, 1)
	signal.Notify(chanServer, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	host := ":8080"
	go func() {
		<-chanServer

		log.Printf("Server is shutting down in the %s.", host)
		err := app.Shutdown()
		if err != nil {
			log.Printf("Error in shutting down the server: %v.", err)
		}
	}()

	log.Printf("Server is running in the %s.", host)
	err := app.Listen(host)
	if err != nil {
		log.Printf("Error in running the server: %v.", err)
	}
}
