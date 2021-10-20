package tokeninfo

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
	GetInfo(*fiber.Ctx) error
	GetPending(*fiber.Ctx) error
	GetScams(*fiber.Ctx) error
	ReportScam(*fiber.Ctx) error
	ApproveScam(*fiber.Ctx) error
	DeletePending(*fiber.Ctx) error
}
