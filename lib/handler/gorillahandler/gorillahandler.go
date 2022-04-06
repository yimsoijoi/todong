package gorillahandler

import (
	"github.com/yimsoijoi/todong/lib/middleware"
	"github.com/yimsoijoi/todong/lib/store"
)

type GorillaHandler struct {
	DataGateway store.DataGateway
	Config      *middleware.Config
}
