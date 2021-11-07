package user

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) Register(ctx *fiber.Ctx) error {
	var req RegisterReq
	err := ctx.BodyParser(&req)
	if err != nil {
		ctx.Status(400).JSON(map[string]interface{}{
			"error": "bad request",
		})
	}

	// Find existing username
	var user datamodel.User
	h.pg.WithContext(ctx.Context()).Where("username = ?", req.Username).First(&user)
	if user.Username != "" {
		return ctx.Status(400).JSON(map[string]string{
			"error": "duplicate username",
		})
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	newUser := datamodel.User{
		UUID:     uuid.NewString(),
		Username: req.Username,
		Password: pw,
	}

	h.pg.WithContext(ctx.Context()).Create(&newUser)
	return ctx.JSON(map[string]interface{}{
		"message": "registered successfully",
	})
}
