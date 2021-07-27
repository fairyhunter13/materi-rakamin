package main

import (
	"log"
	"os"

	"github.com/fairyhunter13/materi-rakamin/pkg/config"
)

var cfg = new(Config)

type Config struct {
	Host     string   `mapstructure:"HOST"`
	Database Database `mapstructure:"DATABASE"`
	Cache    Cache    `mapstructure:"CACHE"`
}

type Database struct {
	Host     string `mapstructure:"HOST"`
	Port     uint64 `mapstructure:"PORT"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
	DBName   string `mapstructure:"DB_NAME"`
}

type Cache struct {
	Address  string `mapstructure:"ADDRESS"`
	Password string `mapstructure:"PASSWORD"`
	DB       int    `mapstructure:"DB"`
}

func init() {
	log.SetFlags(log.Llongfile)
	err := config.LoadConfig(cfg, "dbcache/.env", "./dbcache/.env")
	if err != nil {
		log.Printf("Error in loading the config: %v.", err)
		os.Exit(1)
		return
	}

	initDB()
	initCache()
}
