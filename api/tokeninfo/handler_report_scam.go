package tokeninfo

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
)

type ReportScam struct {
	Address string `json:"address"`
	Chain   string `json:"chain"`
}

type ReportReq []ReportScam

func (h *handler) ReportScam(ctx *fiber.Ctx) error {
	var req ReportReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(map[string]string{
			"error": "bad request",
		})
	}

	for _, token := range req {
		h.pg.WithContext(ctx.Context()).Create(&datamodel.TokenInfo{
			Address:     token.Address,
			Chain:       token.Chain,
			IsScam:      false,
			PendingScam: true,
		})
	}

	return ctx.Status(201).JSON(map[string]bool{
		"success": true,
	})
}
