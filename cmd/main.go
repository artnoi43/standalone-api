package main

import (
	"log"

	"github.com/artworkk/standalone-api/api/auth"
	"github.com/artworkk/standalone-api/api/tokeninfo"
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

	app := fiber.New()
	app.Use(logger.New())
	authHandler := auth.NewHandler(db)
	tokenInfoHandler := tokeninfo.NewHandler(db)

	api := app.Group("/api")
	authAPI := api.Group("/auth")
	authAPI.Post("/register", authHandler.Register)
	authAPI.Post("/login", authHandler.Login)
	tokenInfoAPI := api.Group("/tokens")
	tokenInfoAPI.Get("/info/:tokenAddress", tokenInfoHandler.GetInfo)
	tokenInfoAPI.Get("/pending", tokenInfoHandler.GetPending)
	tokenInfoAPI.Get("/scams", tokenInfoHandler.GetScams)
	tokenInfoAPI.Post("/report-scam", tokenInfoHandler.ReportScam)
	tokenInfoAPI.Post("/approve-scam", tokenInfoHandler.ApproveScam)
	tokenInfoAPI.Delete("/delete-pending/:tokenAddress", tokenInfoHandler.DeletePending)

	log.Fatal(app.Listen(conf.Port))
}
