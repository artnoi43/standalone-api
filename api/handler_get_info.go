package api

import (
	"errors"

	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *handler) GetInfo(ctx *fiber.Ctx) error {
	tokenAddr := ctx.Params("tokenAddress")
	var token *datamodel.TokenInfo
	result := h.pg.WithContext(ctx.Context()).Model(&datamodel.TokenInfo{}).Where(&datamodel.TokenInfo{
		Address: tokenAddr,
	}).First(&token)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(404).JSON(map[string]string{
			"warning": "tokeninfo not found",
		})
	}
	return ctx.Status(200).JSON(token)
}
