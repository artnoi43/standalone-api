package user

import (
	"github.com/artworkk/standalone-api/lib/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	pg     *gorm.DB
	config auth.Config
}

func NewHandler(db *gorm.DB, conf auth.Config) Handler {
	return &handler{
		pg:     db,
		config: conf,
	}
}

type Handler interface {
	Register(*fiber.Ctx) error
	Login(*fiber.Ctx) error
	GetUsers(*fiber.Ctx) error
}
