package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var (
	POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB string
	MainDB                                                                      *sql.DB
)

func init() {
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB)

	var err error
	MainDB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("main: Could not connect to db %s %s", err, connStr)
	}

	MainDB.SetMaxOpenConns(10)
	MainDB.SetMaxIdleConns(2)
	// not sure about this one, depends on use case i guess
	MainDB.SetConnMaxLifetime(30 * time.Minute)

	err = MainDB.Ping()
	if err != nil {
		log.Fatalf("main: Error pinging database: %v %s", err, connStr)
	}

	log.Println("connected to db")
}

func main() {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		RequestMethods:    []string{"GET", "POST", "HEAD"},
	})

	apiV1 := app.Group("/api/v1")
	apiV1.Use(ApiKeyAuth)
	apiV1.Route("/news", func(news fiber.Router) {
		news.Get("/list", ListNews)
		news.Post("/edit/:id", ValidateEditNews, EditNews)
	})

	app.Listen(":3333")

}
