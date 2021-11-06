package tokeninfo

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
)

type ReportReq struct {
	Address string `json:"address"`
	Chain string   `json:"chain"`
}

func (h *handler) ReportScam(ctx *fiber.Ctx) error {
	var req ReportReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(map[string]string{
			"error": "bad request",
		})
	}

	// Range over request body addresses
	h.pg.WithContext(ctx.Context()).Create(&datamodel.TokenInfo{
		Address:     req.Address,
		Chain:       req.Chain,
		IsScam:      false,
		PendingScam: true,
	})

	return ctx.Status(201).JSON(map[string]bool{
		"success": true,
	})
}
