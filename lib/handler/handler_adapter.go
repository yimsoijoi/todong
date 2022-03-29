package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler/fiberhandler"
	"github.com/artnoi43/todong/lib/handler/ginhandler"
	"github.com/artnoi43/todong/lib/middleware"
	"github.com/artnoi43/todong/lib/store"
)

// Adapter abstracts handlers of different web frameworks
type Adaptor interface {
	Gin(string) func(*gin.Context)
	Fiber(string) func(*fiber.Ctx) error
}

// adapter implements Adapter
type adapter struct {
	gin      *ginhandler.GinHandler
	fiber    *fiberhandler.FiberHandler
	ginMap   map[string]func(*gin.Context)
	fiberMap map[string]func(*fiber.Ctx) error
}

func (h *adapter) Gin(s string) func(*gin.Context) {
	if h.ginMap[s] == nil {
		panic(fmt.Sprintf("missing gin handlers for: %s", s))
	}
	return h.ginMap[s]
}

func (h *adapter) Fiber(s string) func(*fiber.Ctx) error {
	if h.fiberMap[s] == nil {
		panic(fmt.Sprintf("missing fiber handlers for: %s", s))
	}
	return h.fiberMap[s]
}

func NewAdaptor(t enums.ServerType, dataGateway store.DataGateway, conf *middleware.Config) Adaptor {
	var g *ginhandler.GinHandler
	var f *fiberhandler.FiberHandler
	switch t.ToUpper() {
	case enums.Gin:
		g = &ginhandler.GinHandler{
			DataGateway: dataGateway,
			Config:      conf,
		}
		mapGin, _ := MapHandlers(g, f)
		return &adapter{
			gin:      g,
			fiber:    nil,
			ginMap:   mapGin,
			fiberMap: nil,
		}
	case enums.Fiber:
		f = &fiberhandler.FiberHandler{
			DataGateway: dataGateway,
			Config:      conf,
		}
		_, mapFiber := MapHandlers(g, f)
		return &adapter{
			gin:      nil,
			fiber:    f,
			ginMap:   nil,
			fiberMap: mapFiber,
		}
	}
	panic("invalid server type")
}
