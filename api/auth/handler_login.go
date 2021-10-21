package auth

import (
	"time"

	"github.com/artworkk/standalone-api/datamodel"
	"github.com/dgrijalva/jwt-go"
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

	// Find user in database
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

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.UUID,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte(h.config.JWTSecret))
	if err != nil {
		return ctx.Status(500).JSON(map[string]interface{}{
			"error": "could not login",
		})
	}

	return ctx.JSON(map[string]interface{}{
		"message": "logged in successfully",
		"token":   token,
	})
}
