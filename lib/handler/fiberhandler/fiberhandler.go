package fiberhandler

import (
	"github.com/yimsoijoi/todong/lib/middleware"
	"github.com/yimsoijoi/todong/lib/store"
)

type FiberHandler struct {
	DataGateway store.DataGateway
	Config      *middleware.Config
}
