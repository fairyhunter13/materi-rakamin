package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

type Product struct {
	Name         string          `json:"name"`
	Kind         string          `json:"kind"`
	Price        decimal.Decimal `json:"price"`
	Place        string          `json:"place"`
	Barcode      string          `json:"barcode"`
	PurchaseDate time.Time       `json:"purchase_date"`
}

var (
	users = []Product{
		{
			Name:         "Mie Samyang",
			Kind:         "Noodle",
			Price:        decimal.NewFromInt(20000),
			Place:        "Alfamart",
			Barcode:      "AA112233",
			PurchaseDate: time.Now().Add(-3 * 24 * time.Hour),
		},
		{
			Name:         "Mie Gaga 100",
			Kind:         "Noodle",
			Price:        decimal.NewFromInt(3000),
			Place:        "Alfamart",
			Barcode:      "BB112233",
			PurchaseDate: time.Now().Add(-1 * 24 * time.Hour),
		},
		{
			Name:         "Susu UHT Coklat 1 Liter",
			Kind:         "Milk",
			Price:        decimal.NewFromInt(20000),
			Place:        "Alfamart",
			Barcode:      "CC321123",
			PurchaseDate: time.Now().Add(-5 * time.Hour),
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
	log.Println("Press Ctrl + C to exit the server!")
	err := app.Listen(host)
	if err != nil {
		log.Printf("Error in running the server: %v.", err)
	}
}
