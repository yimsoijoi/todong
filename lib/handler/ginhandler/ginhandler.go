package ginhandler

import (
	"github.com/yimsoijoi/todong/lib/middleware"
	"github.com/yimsoijoi/todong/lib/store"
)

type GinHandler struct {
	DataGateway store.DataGateway
	Config      *middleware.Config
}
