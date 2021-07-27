package main

import "github.com/gofiber/fiber/v2"

type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func printJSONErr(c *fiber.Ctx, errList *[]Error, err *error) {
	if c == nil || errList == nil || err == nil || *err == nil {
		return
	}

	*err = c.JSON(append(*errList, castError(*err)))
}

func castError(err error) Error {
	return Error{err.Error()}
}
