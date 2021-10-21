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
	authHandler := auth.NewHandler(db, conf.Auth)
	tokenInfoHandler := tokeninfo.NewHandler(db)

	log.Println("JWT secret:", conf.Auth.JWTSecret)

	api := app.Group("/api")
	authAPI := api.Group("/auth")
	authAPI.Post("/register", authHandler.Register)
	authAPI.Post("/login", authHandler.Login)
	tokenInfoAPI := api.Group("/tokens")
	tokenInfoAPI.Get("/info/:tokenAddress", tokenInfoHandler.GetInfo)
	tokenInfoAPI.Get("/pending", tokenInfoHandler.GetPending)
	tokenInfoAPI.Get("/scams", tokenInfoHandler.GetScams)
	tokenInfoAPI.Post("/report-scam", tokenInfoHandler.ReportScam)
	// These paths need authentication header
	tokenInfoAPI.Post("/approve-scam", auth.Authenticate(conf.Auth), tokenInfoHandler.ApproveScam)
	tokenInfoAPI.Delete("/delete-pending/:tokenAddress", auth.Authenticate(conf.Auth), tokenInfoHandler.DeletePending)

	log.Fatal(app.Listen(conf.Port))
}
