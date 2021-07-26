package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/fairyhunter13/materi-rakamin/pkg/config"
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
	cfg := new(Config)
	err := config.LoadConfig(cfg, "externalsvc/.env", "./externalsvc/.env")
	if err != nil {
		log.Printf("Error in loading the config: %v.", err)
		return
	}

	client, err := minio.New(cfg.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV2(cfg.Storage.AccessKey, cfg.Storage.SecretKey, ""),
		Secure: true,
		Region: cfg.Storage.Region,
	})
	if err != nil {
		log.Printf("Error in initializing the new client: %v.", err)
		return
	}

	ctx := context.Background()
	isExist, err := client.BucketExists(ctx, cfg.Storage.Bucket)
	if err != nil {
		log.Printf("Error in checking the bucket: %v.", err)
		return
	}

	if !isExist {
		log.Printf("Bucket %s is not exist!", cfg.Storage.Bucket)
		return
	}

	isObjectExist := true
	objectInfo, err := client.StatObject(ctx, cfg.Storage.Bucket, testFilename, minio.GetObjectOptions{})
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
		obj, err := client.GetObject(ctx, cfg.Storage.Bucket, testFilename, minio.GetObjectOptions{})
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
	uploadInfo, err := client.PutObject(ctx, cfg.Storage.Bucket, testFilename, strReader, int64(strReader.Len()), minio.PutObjectOptions{})
	if err != nil {
		log.Printf("Error in uploading the file #%s: %v.", testFilename, err)
		return
	}

	log.Printf("Uploading the file #%s succeeded!", testFilename)
	fmt.Println("UploadInfo:")
	fmt.Printf("%+v\n", uploadInfo)
}
