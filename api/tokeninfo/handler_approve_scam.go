package tokeninfo

import (
	"github.com/artworkk/standalone-api/datamodel"
	"github.com/gofiber/fiber/v2"
)

type ApproveSpam struct {
	Address string `json:"address"`
	Chain   string `json:"chain"`
}

type ApproveReq []ApproveSpam

func (h *handler) ApproveScam(ctx *fiber.Ctx) error {
	var req ApproveReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(map[string]string{
			"error": "bad request",
		})
	}

	for _, token := range req {
		h.pg.WithContext(ctx.Context()).Model(&datamodel.TokenInfo{}).Where("address = ?", token.Address).Where("chain = ?", token.Chain).Updates(map[string]interface{}{
			"is_scam":      true,
			"pending_scam": false,
		})
	}

	return ctx.Status(201).JSON(map[string]bool{
		"success": true,
	})
}
