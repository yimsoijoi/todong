package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/lib/handler/fiberhandler"
	"github.com/yimsoijoi/todong/lib/handler/ginhandler"
	"github.com/yimsoijoi/todong/lib/handler/gorillahandler"
	"github.com/yimsoijoi/todong/lib/middleware"
	"github.com/yimsoijoi/todong/lib/store"
)

// Adapter abstracts handlers of different web frameworks
type Adaptor interface {
	Gin(string) func(*gin.Context)
	Fiber(string) func(*fiber.Ctx) error
	Gorilla(string) http.HandlerFunc
}

// adapter implements Adapter
type adapter struct {
	gin        *ginhandler.GinHandler
	fiber      *fiberhandler.FiberHandler
	gorilla    *gorillahandler.GorillaHandler
	ginMap     map[string]func(*gin.Context)
	fiberMap   map[string]func(*fiber.Ctx) error
	gorillaMap map[string]http.HandlerFunc
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
func (h *adapter) Gorilla(s string) http.HandlerFunc {
	if h.gorillaMap[s] == nil {
		panic(fmt.Sprintf("missing gorilla handlers for: %s", s))

	}
	return h.gorillaMap[s]
}

func NewAdaptor(t enums.ServerType, dataGateway store.DataGateway, conf *middleware.Config) Adaptor {
	var g *ginhandler.GinHandler
	var f *fiberhandler.FiberHandler
	var gr *gorillahandler.GorillaHandler
	switch t.ToUpper() {
	case enums.Gin:
		g = &ginhandler.GinHandler{
			DataGateway: dataGateway,
			Config:      conf,
		}
		mapGin, _, _ := MapHandlers(g, f, gr)
		return &adapter{
			gin:        g,
			fiber:      nil,
			gorilla:    nil,
			ginMap:     mapGin,
			fiberMap:   nil,
			gorillaMap: nil,
		}
	case enums.Fiber:
		f = &fiberhandler.FiberHandler{
			DataGateway: dataGateway,
			Config:      conf,
		}
		_, mapFiber, _ := MapHandlers(g, f, gr)
		return &adapter{
			gin:        nil,
			fiber:      f,
			gorilla:    nil,
			ginMap:     nil,
			fiberMap:   mapFiber,
			gorillaMap: nil,
		}
	case enums.Gorilla:
		gr = &gorillahandler.GorillaHandler{
			DataGateway: dataGateway,
			Config:      conf,
		}
		_, _, mapGorilla := MapHandlers(g, f, gr)
		return &adapter{
			gin:        nil,
			fiber:      nil,
			gorilla:    gr,
			ginMap:     nil,
			fiberMap:   nil,
			gorillaMap: mapGorilla,
		}
	}
	panic("invalid server type")
}
