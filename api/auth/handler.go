package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	pg *gorm.DB
}

func NewHandler(db *gorm.DB) Handler {
	return &handler{
		pg: db,
	}
}

type Handler interface {
	Register(*fiber.Ctx) error
	Login(*fiber.Ctx) error
}