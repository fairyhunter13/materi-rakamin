package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func addBasicRoute(app *fiber.App) {
	basicauthHandler := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "4dm1n",
		},
	})

	app.Use("/users", basicauthHandler).Get("/users", func(c *fiber.Ctx) (err error) {
		err = c.JSON(users)
		return
	})

}
