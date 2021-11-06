package tokeninfo

import (
	"errors"

	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *handler) GetInfo(ctx *fiber.Ctx) error {
	tokenAddr := ctx.Params("tokenAddress")
	chain := ctx.Params("chain")
	var token *datamodel.TokenInfo
	result := h.pg.WithContext(ctx.Context()).Model(&datamodel.TokenInfo{}).Where(&datamodel.TokenInfo{
		Address: tokenAddr,
		Chain: chain,
	}).First(&token)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return ctx.Status(404).JSON(map[string]string{
			"warning": "tokeninfo not found",
		})
	}
	return ctx.Status(200).JSON(token)
}
