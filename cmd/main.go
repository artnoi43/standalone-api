package main

import (
	"log"

	"github.com/artworkk/standalone-api/api"
	"github.com/artworkk/standalone-api/config"
	"github.com/artworkk/standalone-api/lib/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Println("App started")
	conf, err := config.LoadConfig("config")
	if err != nil {
		panic(err.Error())
	}
	db, err := postgres.New(conf.Postgres)
	if err != nil {
		panic(err.Error())
	}

	handler := api.NewHandler(db)
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/info/:tokenAddress", handler.GetInfo)
	app.Get("/pending", handler.GetPending)
	app.Get("/scams", handler.GetScams)
	app.Post("/report-scam", handler.ReportScam)
	app.Post("/approve-scam", handler.ApproveScam)
	app.Delete("/delete-pending/:tokenAddress", handler.DeletePending)

	log.Fatal(app.Listen(conf.Port))
}
