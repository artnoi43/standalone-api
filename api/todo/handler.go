package todo

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type handler struct {
	pg *gorm.DB
}

type Handler interface {
	GetTodo(*fiber.Ctx) error
	NewTodo(*fiber.Ctx) error
}

func NewHandler(db *gorm.DB) Handler {
	return &handler{
		pg: db,
	}
}
