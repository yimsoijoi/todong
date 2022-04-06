package fiberhandler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/yimsoijoi/todong/datamodel"
	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/internal"
	"github.com/yimsoijoi/todong/lib/utils"
)

func (h *FiberHandler) Register(c *fiber.Ctx) error {
	var req internal.AuthJson
	if err := c.BodyParser(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}

	ctx := c.Context()
	var targetUser datamodel.User
	_ = h.DataGateway.GetUserByUsername(ctx, req.Username, &targetUser)
	if targetUser.Username != "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"duplicate username": req.Username,
		})
	}

	pw, err := utils.EncodeBcrypt([]byte(req.Password))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"failed to generate password": err.Error(),
		})
	}
	user := datamodel.User{
		UUID:     uuid.NewString(),
		Username: req.Username,
		Password: pw,
	}
	if err := h.DataGateway.CreateUser(ctx, &user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("failed to create user: %s", req.Username),
		})
	}
	// user.Password will not leak because the field has "-" in its JSON tag
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":   "user registration successful",
		"username": user.Username,
		"userUuid": user.UUID,
	})
}
