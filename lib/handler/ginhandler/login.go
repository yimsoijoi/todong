package ginhandler

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/yimsoijoi/todong/datamodel"
	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/internal"
	"github.com/yimsoijoi/todong/lib/store"
	"github.com/yimsoijoi/todong/lib/utils"
)

// Login authenticates username/password and return JWT token signed with configured secret
func (h *GinHandler) Login(c *gin.Context) {
	var req internal.AuthJson
	if err := c.BindJSON(&req); err != nil {
		status := utils.ErrStatus(enums.MapErrHandler.Unmarshal, err)
		c.AbortWithStatusJSON(http.StatusBadRequest, status)
		return
	}

	ctx := c.Request.Context()
	var user datamodel.User
	if err := h.DataGateway.GetUserByUsername(ctx, req.Username, &user); err != nil {
		if !errors.Is(err, store.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "login failed",
			})
			return
		}
	}
	// Null user from database, i.e. zero-valued user.UUID
	if user.UUID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "invalid username or password",
			"code":   1,
		})
		return
	}
	// Compare hashed password with bcrypt
	if err := utils.DecodeBcrypt(user.Password, []byte(req.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": "invalid username or password",
			"code":   2,
		})
		return
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed to generate token",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status":   "successful login",
		"username": user.Username,
		"userUuid": iss,
		"expire":   exp.String(),
		"token":    token,
	})
}
