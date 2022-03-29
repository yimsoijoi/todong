package fiberhandler

import (
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

type FiberHandler struct {
	DataGateway store.DataGateway
	Config      *middleware.Config
}
