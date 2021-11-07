package todo

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type NewTodoReq struct {
	Todo string `json:"todo"`
}

func (h *handler) NewTodo(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	// Issuer is iss
	userUuid := claims["iss"].(string)

	var dbUser datamodel.User
	h.pg.WithContext(ctx.Context()).Where("uuid = ?", userUuid).First(&dbUser)
	if dbUser.Username == "" {
		return ctx.Status(404).JSON(map[string]string{
			"error": "user not found",
		})
	}

	var req NewTodoReq
	if err := ctx.BodyParser(&req); err != nil {
		ctx.SendStatus(400)
	}
	newTodo := &datamodel.Todo{
		UUID:     uuid.NewString(),
		UserUUID: userUuid,
		Text:     req.Todo,
		Done:     false,
	}
	h.pg.WithContext(ctx.Context()).Create(newTodo)

	var todos []datamodel.Todo
	h.pg.WithContext(ctx.Context()).Model(&datamodel.Todo{}).Where("user_uuid = ?", userUuid).Find(&todos)
	return ctx.Status(201).JSON(todos)
}
