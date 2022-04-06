package gorillahandler

import (
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

type GorillaHandler struct {
	DataGateway store.DataGateway
	Config      *middleware.Config
}
