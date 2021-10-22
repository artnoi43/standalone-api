package auth

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Authenticate(conf Config) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(conf.JWTSecret),
	})
}
