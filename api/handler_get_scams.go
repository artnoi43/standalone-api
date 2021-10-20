package api

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetScams(ctx *fiber.Ctx) error {
	var tokenInfos []datamodel.TokenInfo
	h.pg.WithContext(ctx.Context()).Model(&datamodel.TokenInfo{}).Where(&datamodel.TokenInfo{
		IsScam: true,
	}).Find(&tokenInfos)

	return ctx.JSON(tokenInfos)
}
