package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type Config struct {
	StorageEndpoint  string `mapstructure:"STORAGE_ENDPOINT"`
	StorageAccessKey string `mapstructure:"STORAGE_ACCESS_KEY"`
	StorageSecretKey string `mapstructure:"STORAGE_SECRET_KEY"`
	StorageRegion    string `mapstructure:"STORAGE_REGION"`
	StorageBucket    string `mapstructure:"STORAGE_BUCKET"`
}

func loadConfig(paths ...string) (c *Config, err error) {
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	c = new(Config)
	err = viper.Unmarshal(c)
	return
}

const (
	testJSON = `
		{
			"hello" : "hi!",
			"Rakamin": "Ini untuk materi rakamin"
		}
	`
	testFilename = `test.json`
)

func main() {
	cfg, err := loadConfig("", "externalsvc/", "./externalsvc")
	if err != nil {
		log.Printf("Error in loading the config: %v.", err)
		return
	}

	client, err := minio.New(cfg.StorageEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV2(cfg.StorageAccessKey, cfg.StorageSecretKey, ""),
		Secure: true,
		Region: cfg.StorageRegion,
	})
	if err != nil {
		log.Printf("Error in initializing the new client: %v.", err)
		return
	}

	ctx := context.Background()
	isExist, err := client.BucketExists(ctx, cfg.StorageBucket)
	if err != nil {
		log.Printf("Error in checking the bucket: %v.", err)
		return
	}

	if !isExist {
		log.Printf("Bucket %s is not exist!", cfg.StorageBucket)
		return
	}

	isObjectExist := true
	objectInfo, err := client.StatObject(ctx, cfg.StorageBucket, testFilename, minio.GetObjectOptions{})
	if err != nil {
		errResp := minio.ToErrorResponse(err)
		if errResp.Code != "NoSuchKey" {
			log.Printf("Error in getting the object info: %v.", err)
			return
		}
		err = nil
		isObjectExist = false
	}

	if isObjectExist {
		fmt.Println("ObjectInfo:")
		fmt.Printf("%+v\n", objectInfo)
	}

	if isObjectExist {
		obj, err := client.GetObject(ctx, cfg.StorageBucket, testFilename, minio.GetObjectOptions{})
		if err != nil {
			log.Printf("Error in getting the object: %v.", err)
			return
		}

		sb := new(strings.Builder)
		_, err = io.Copy(sb, obj)
		if err != nil {
			log.Printf("Error in copying from the object reader: %v.", err)
			return
		}

		fmt.Printf("File \"%s\":\n", testFilename)
		fmt.Println(sb.String())
		return
	}

	strReader := strings.NewReader(testJSON)
	uploadInfo, err := client.PutObject(ctx, cfg.StorageBucket, testFilename, strReader, int64(strReader.Len()), minio.PutObjectOptions{})
	if err != nil {
		log.Printf("Error in uploading the file #%s: %v.", testFilename, err)
		return
	}

	log.Println("Uploading succeeded!")
	fmt.Println("UploadInfo:")
	fmt.Printf("%+v\n", uploadInfo)
}
