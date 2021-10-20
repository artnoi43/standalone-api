package tokeninfo

import (
	"errors"

	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *handler) DeletePending(ctx *fiber.Ctx) error {
	tokenAddr := ctx.Params("tokenAddress")
	// Only delete those with pending_scam = true
	result := h.pg.WithContext(ctx.Context()).Where("pending_scam = ?", true).Delete(&datamodel.TokenInfo{}, tokenAddr)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return ctx.Status(404).JSON(map[string]string{
			"warning": "tokeninfo not found",
		})
	}
	return ctx.SendStatus(200)
}