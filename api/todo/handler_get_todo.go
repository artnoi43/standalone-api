package todo

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *handler) GetTodo(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// Issuer is iss
	userUuid := claims["iss"].(string)
	var todos []datamodel.Todo
	h.pg.WithContext(ctx.Context()).Model(&datamodel.Todo{}).Where("user_uuid = ?", userUuid).Find(&todos)
	return ctx.Status(201).JSON(todos)
}
