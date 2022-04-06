package httpserver

import (
	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/lib/handler"
	"github.com/yimsoijoi/todong/lib/httpserver/fiberserver"
	"github.com/yimsoijoi/todong/lib/httpserver/ginserver"
	"github.com/yimsoijoi/todong/lib/httpserver/gorillaserver"
	"github.com/yimsoijoi/todong/lib/middleware"
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
		case enums.Gorilla:
			return gorillaserver.New()
		}
	}
	panic("invalid server type")
}
