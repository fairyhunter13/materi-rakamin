package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fairyhunter13/materi-rakamin/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Storage Storage `mapstructure:"STORAGE"`
}

type Storage struct {
	Endpoint  string `mapstructure:"ENDPOINT"`
	AccessKey string `mapstructure:"ACCESS_KEY"`
	SecretKey string `mapstructure:"SECRET_KEY"`
	Region    string `mapstructure:"REGION"`
	Bucket    string `mapstructure:"BUCKET"`
}

var (
	client *minio.Client
	cfg    *Config
)

func init() {
	cfg = new(Config)
	err := config.LoadConfig(cfg, "externalsvc/.env", "./externalsvc/.env")
	if err != nil {
		log.Fatalf("Error in loading the config: %v.", err)
		return
	}

	client, err = minio.New(cfg.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV2(cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
		Secure: true,
		Region: cfg.Storage.Region,
	})
	if err != nil {
		log.Fatalf("Error in initializing the new client: %v.", err)
		return
	}
}

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
		AppName:   "Server Handson",
	})
	uploadHandler(app)
	downloadHandler(app)

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
