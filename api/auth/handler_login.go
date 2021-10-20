package auth

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) Login(ctx *fiber.Ctx) error {
	var req LoginReq
	if err := ctx.BodyParser(&req); err != nil {
		ctx.Status(400).JSON(map[string]interface{}{
			"error": "bad request",
		})
	}
	
	var user datamodel.User
	h.pg.Where("username = ?", req.Username).First(&user)
	if len(user.UUID) == 0 {
		return ctx.Status(400).JSON(map[string]interface{}{
			"error": "incorrect username or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"error": "incorrect username or password",
		})
	}

	return ctx.JSON(map[string]interface{}{
		"status": "logged in successfully",
	})
}