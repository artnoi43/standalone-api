package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	pg     *gorm.DB
	config Config
}

func NewHandler(db *gorm.DB, conf Config) Handler {
	return &handler{
		pg:     db,
		config: conf,
	}
}

type Handler interface {
	Register(*fiber.Ctx) error
	Login(*fiber.Ctx) error
}
