package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

var signingKey = []byte("secret")

type UserRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func addJWTRoute(app *fiber.App) {
	apiGroup := app.Group("/api")
	apiGroup.Post("/login", func(c *fiber.Ctx) (err error) {
		var req UserRequest
		err = c.BodyParser(&req)
		if err != nil {
			log.Printf("Error in parsing the JSON request: %v.", err)
			return
		}

		if req.User != "admin" || req.Password != "4dm1n" {
			err = c.SendStatus(fiber.StatusUnauthorized)
			return
		}

		signJwt := jwt.New(jwt.SigningMethodHS256)

		claims := signJwt.Claims.(jwt.MapClaims)
		claims["name"] = "Admin"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		token, err := signJwt.SignedString(signingKey)
		if err != nil {
			err = c.SendStatus(fiber.StatusInternalServerError)
			return
		}

		err = c.JSON(fiber.Map{"token": token})
		return
	})

	apiGroup.Use("/users", jwtware.New(jwtware.Config{
		SigningKey: signingKey,
	}))
	apiGroup.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})
}
