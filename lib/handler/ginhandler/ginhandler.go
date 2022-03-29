package ginhandler

import (
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

type GinHandler struct {
	DataGateway store.DataGateway
	Config      *middleware.Config
}
