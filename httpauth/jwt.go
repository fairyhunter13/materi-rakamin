package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

var signingKey = []byte("rakamin")

type UserRequest struct {
	Username string `json:"username"`
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

		if req.Username != "admin" || req.Password != "4dm1n" {
			err = c.SendStatus(fiber.StatusUnauthorized)
			return
		}

		signJwt := jwt.New(jwt.SigningMethodHS256)

		claims := signJwt.Claims.(jwt.MapClaims)
		claims["name"] = "Admin"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		signJwt.Claims = claims

		token, err := signJwt.SignedString(signingKey)
		if err != nil {
			err = c.SendStatus(fiber.StatusInternalServerError)
			return
		}

		err = c.JSON(fiber.Map{"token": token})
		return
	})

	apiGroup.Use("/products", jwtware.New(jwtware.Config{
		SigningKey: signingKey,
	}))
	apiGroup.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})
}
