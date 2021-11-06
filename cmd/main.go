package main

import (
	"log"

	"github.com/artworkk/standalone-api/api/tokeninfo"
	"github.com/artworkk/standalone-api/api/user"
	"github.com/artworkk/standalone-api/config"
	"github.com/artworkk/standalone-api/lib/auth"
	"github.com/artworkk/standalone-api/lib/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Println("App started")
	conf, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalln("Failed to load config:", err)
	}
	db, err := postgres.New(conf.Postgres)
	if err != nil {
		log.Fatalln("Failed to open database:", err)
	}
	log.Println("JWT secret:", conf.Auth.JWTSecret)

	app := fiber.New()
	app.Use(logger.New())
	api := app.Group("/api")

	userApi := api.Group("/user")
	userHandler := user.NewHandler(db, conf.Auth)
	userApi.Post("/register", userHandler.Register)
	userApi.Post("/login", userHandler.Login)

	tokenInfoAPI := api.Group("/tokens")
	tokenInfoHandler := tokeninfo.NewHandler(db)
	tokenInfoAPI.Get("/info/:chain/:tokenAddress", tokenInfoHandler.GetInfo)
	tokenInfoAPI.Get("/pending", tokenInfoHandler.GetPending)
	tokenInfoAPI.Get("/scams", tokenInfoHandler.GetScams)
	tokenInfoAPI.Post("/report-scam", tokenInfoHandler.ReportScam)
	// These paths need authentication header
	tokenInfoAPI.Post("/approve-scam", auth.Authenticate(conf.Auth), tokenInfoHandler.ApproveScam)
	tokenInfoAPI.Delete("/delete-pending/:chain/:tokenAddress", auth.Authenticate(conf.Auth), tokenInfoHandler.DeletePending)

	log.Fatal(app.Listen(conf.Port))
}
