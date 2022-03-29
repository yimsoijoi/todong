package httpserver

import (
	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler"
	"github.com/artnoi43/todong/lib/httpserver/fiberserver"
	"github.com/artnoi43/todong/lib/httpserver/ginserver"
	"github.com/artnoi43/todong/lib/middleware"
)

// Server abstracts different web frameworks (e.g. Fiber, Gorilla, and Gin)
type Server interface {
	SetUpRoutes(conf *middleware.Config, handler handler.Adaptor)
	Serve(addr string) error
}

func New(t enums.ServerType) Server {
	if t.IsValid() {
		switch t.ToUpper() {
		case enums.Gin:
			return ginserver.New()
		case enums.Fiber:
			return fiberserver.New()
		}
	}
	panic("invalid server type")
}
