package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

func downloadHandler(app *fiber.App) {
	app.Get("/files/:filename", func(c *fiber.Ctx) (err error) {
		file := c.Params("filename")

		var obj *minio.Object
		obj, err = client.GetObject(c.UserContext(), cfg.Storage.Bucket, PathGCS+file, minio.GetObjectOptions{})
		if err != nil {
			return
		}

		c.SendStream(obj)
		return
	})
}
