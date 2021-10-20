package api

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
)

type ReportReq struct {
	Addresses []string `json:"addresses"`
}

func (h *handler) ReportScam(ctx *fiber.Ctx) error {
	var req ReportReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(map[string]string{
			"error": "bad request",
		})
	}

	// Range over request body addresses
	for _, address := range req.Addresses {
		h.pg.WithContext(ctx.Context()).Create(&datamodel.TokenInfo{
			Address:     address,
			IsScam:      false,
			PendingScam: true,
		})
	}

	return ctx.Status(201).JSON(map[string]bool{
		"success": true,
	})
}
