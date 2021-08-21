package main

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

const (
	KeyForm = `uploads`
	PathGCS = "C/Hafiz_Putra_Ludyanto/"
)

func uploadHandler(app *fiber.App) {
	app.Post("/file", func(c *fiber.Ctx) (err error) {
		var form *multipart.Form
		form, err = c.MultipartForm()
		if err != nil {
			return
		}

		files := form.File[KeyForm]
		if len(files) <= 0 {
			err = errors.New("err no files")
			return
		}

		var (
			arrResp []map[string]string
			id      string
			ext     string
		)
		for _, file := range files {
			func() {
				var realFile multipart.File
				realFile, err = file.Open()
				if err != nil {
					return
				}
				defer realFile.Close()

				var path string
				id, ext, path = generateObjectPath(file)
				_, err = client.PutObject(c.UserContext(), cfg.Storage.Bucket, path, realFile, file.Size, minio.PutObjectOptions{})
			}()
			if err != nil {
				return
			}

			var mapResp = make(map[string]string)
			mapResp["uuid"] = id
			mapResp["ext"] = ext
			arrResp = append(arrResp, mapResp)
		}

		c.JSON(arrResp)
		return
	})
}

func generateObjectPath(f *multipart.FileHeader) (id string, ext string, path string) {
	ext = filepath.Ext(f.Filename)
	extArr := strings.Split(ext, ".")
	id = uuid.New().String()
	path = PathGCS + id + ext
	if len(extArr) > 1 {
		ext = extArr[1]
	}
	return
}
