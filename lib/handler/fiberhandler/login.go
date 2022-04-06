package fiberhandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"github.com/yimsoijoi/todong/datamodel"
	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/internal"
	"github.com/yimsoijoi/todong/lib/store"
	"github.com/yimsoijoi/todong/lib/utils"
)

// Login authenticates username/password and return JWT token signed with configured secret
func (h *FiberHandler) Login(c *fiber.Ctx) error {
	var req internal.AuthJson
	if err := c.BodyParser(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		return c.Status(http.StatusBadRequest).JSON(status)
	}

	ctx := c.Context()
	var user datamodel.User
	if err := h.DataGateway.GetUserByUsername(ctx, req.Username, &user); err != nil {
		if !errors.Is(err, store.ErrRecordNotFound) {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "login failed",
			})
		}
	}
	// Null user from database, i.e. zero-valued user.UUID
	if user.UUID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status": "invalid username or password",
			"code":   1,
		})
	}
	// Compare hashed password with bcrypt
	if err := utils.DecodeBcrypt(user.Password, []byte(req.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status": "invalid username or password",
			"code":   2,
		})
	}
	// Generate claims (JWT info)
	iss := user.UUID
	// TODO: investigate if Local() is actually needed
	exp := time.Now().Add(24 * time.Hour).Local()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    iss,
		ExpiresAt: exp.Unix(),
	})
	// Generate JWT token from claims
	token, err := claims.SignedString([]byte(h.Config.SecretKey))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status": "failed to generate token",
			"error":  err.Error(),
		})
	}
	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"status":   "successful login",
		"username": user.Username,
		"userUuid": iss,
		"expire":   exp.String(),
		"token":    token,
	})
}
