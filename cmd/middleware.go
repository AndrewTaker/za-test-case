package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type keyType struct{}

var NewsKey keyType

func ValidateEditNews(c *fiber.Ctx) error {
	log.Printf("ValidateEditNews: got request %s", c.Request())

	var news NewsEditRequest

	if err := c.BodyParser(&news); err != nil {
		log.Printf("ValidateEditNews: err reading body %s", c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, "Could not read body")
	}

	if err := validate.Struct(&news); err != nil {
		log.Printf("ValidateEditNews: err validating body %s", c.Request())
		return ResponseGenericError(c, fiber.StatusBadRequest, err.Error())
	}

	c.Locals(NewsKey, news)
	log.Printf("ValidateEditNews: validated %+v %s", news, c.Request())

	return c.Next()
}

func ApiKeyAuth(c *fiber.Ctx) error {
	log.Printf("ApiKeyAuth: got request %s", c.Request())
	header := c.Get("Authorization")

	if !strings.HasPrefix(header, "Bearer ") {
		log.Printf("ApiKeyAuth: no Bearer prefix %s", c.Request())
		return ResponseGenericError(c, fiber.StatusUnauthorized, ErrUnauthorized.Error())
	}

	apiKey := strings.TrimPrefix(header, "Bearer ")

	if apiKey == "" {
		log.Printf("ApiKeyAuth: no api key %s", c.Request())
		return ResponseGenericError(c, fiber.StatusUnauthorized, ErrUnauthorized.Error())
	}

	var r int8
	MainDB.QueryRow("SELECT COUNT(API_KEY) FROM USERS WHERE API_KEY = $1", apiKey).Scan(&r)
	if r != 1 {
		log.Printf("ApiKeyAuth: query for api key != 1 %d %s", r, c.Request())
		return ResponseGenericError(c, fiber.StatusUnauthorized, ErrUnauthorized.Error())
	}

	log.Printf("ApiKeyAuth: passed authentication %s", c.Request())

	return c.Next()
}
