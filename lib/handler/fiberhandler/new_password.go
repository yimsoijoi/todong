package fiberhandler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/yimsoijoi/todong/datamodel"
	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/internal"
	"github.com/yimsoijoi/todong/lib/store"
	"github.com/yimsoijoi/todong/lib/utils"
)

// DeleteUser deletes a datamodel.User in database
// datamodel.User.UUID is used to target deletion
func (h *FiberHandler) NewPassword(c *fiber.Ctx) error {
	userInfo, err := utils.ExtractAndDecodeJwtFiber(c)
	if err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.JwtError, err)
		return c.Status(http.StatusInternalServerError).JSON(status)
	}
	uuid := userInfo.UserUuid

	// Parse new password JSON request
	var newPassReq internal.NewPasswordJson
	if err := c.BodyParser(&newPassReq); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}
	if len(newPassReq.NewPassword) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "blank password received",
		})
	}

	// Get user from DB
	ctx := c.Context()
	var targetUser datamodel.User
	if err := h.DataGateway.GetUserByUuid(ctx, uuid, &targetUser); err != nil {
		// Record not found
		if errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status": "user not found",
				"uuid":   uuid,
			})
		}
		// Other errors
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed to find target user",
			"uuid":   uuid,
			"error":  err.Error(),
		})
	}
	pw, err := utils.EncodeBcrypt([]byte(newPassReq.NewPassword))
	if err != nil {
		if errors.Is(enums.ErrPwTooShort, err) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to change password",
		})
	}
	// Update data in DB
	if err := h.DataGateway.ChangePassword(ctx, &targetUser, pw); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"errors": "failed to change password",
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":   "password change successful",
		"username": targetUser.Username,
	})
}
