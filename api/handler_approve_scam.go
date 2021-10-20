package api

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
)

type ApproveReq struct {
	Addresses []string `json:"addresses"`
}

func (h *handler) ApproveScam(ctx *fiber.Ctx) error {
	var req ApproveReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(map[string]string{
			"error": "bad request",
		})
	}

	// Range over request body addresses
	for _, address := range req.Addresses {
		h.pg.WithContext(ctx.Context()).Model(&datamodel.TokenInfo{}).Where("address = ?", address).Updates(map[string]interface{}{
			"is_scam":      true,
			"pending_scam": false,
		})
	}

	return ctx.Status(201).JSON(map[string]bool{
		"success": true,
	})
}
