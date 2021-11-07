package user

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetUsers(ctx *fiber.Ctx) error {
	var users []datamodel.User
	h.pg.WithContext(ctx.Context()).Preload("Todos").Find(&users)
	return ctx.Status(200).JSON(users)
}
