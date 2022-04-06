package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler/fiberhandler"
	"github.com/artnoi43/todong/lib/handler/ginhandler"
	"github.com/artnoi43/todong/lib/handler/gorillahandler"
)

// MapHandlers map ginHandler/fiberHandler methods to some strings from enums.
func MapHandlers(
	g *ginhandler.GinHandler,
	f *fiberhandler.FiberHandler,
	gr *gorillahandler.GorillaHandler,
) (
	map[string]func(*gin.Context),
	map[string]func(*fiber.Ctx) error,
	map[string]http.HandlerFunc,
) {
	MapGinHandlers := map[string]func(*gin.Context){
		enums.HandlerRegister:    g.Register,
		enums.HandlerLogin:       g.Login,
		enums.HandlerCreateTodo:  g.CreateTodo,
		enums.HandlerGetTodo:     g.GetTodo,
		enums.HandlerUpdateTodo:  g.UpdateTodo,
		enums.HandlerDeleteTodo:  g.DeleteTodo,
		enums.HandlerNewPassword: g.NewPassword,
		enums.HandlerDeleteUser:  g.DeleteUser,
		enums.HandlerTestAuth:    g.TestAuth,
	}
	MapFiberHandlers := map[string]func(*fiber.Ctx) error{
		enums.HandlerRegister:    f.Register,
		enums.HandlerLogin:       f.Login,
		enums.HandlerCreateTodo:  f.CreateTodo,
		enums.HandlerGetTodo:     f.GetTodo,
		enums.HandlerUpdateTodo:  f.UpdateTodo,
		enums.HandlerDeleteTodo:  f.DeleteTodo,
		enums.HandlerNewPassword: f.NewPassword,
		enums.HandlerDeleteUser:  f.DeleteUser,
	}
	MapGorillaHandlers := map[string]http.HandlerFunc{
		enums.HandlerRegister:    gr.Register,
		enums.HandlerLogin:       gr.Login,
		enums.HandlerCreateTodo:  gr.CreateTodo,
		enums.HandlerGetTodo:     gr.GetTodo,
		enums.HandlerUpdateTodo:  gr.UpdateTodo,
		enums.HandlerDeleteTodo:  gr.DeleteTodo,
		enums.HandlerNewPassword: gr.NewPassword,
		enums.HandlerDeleteUser:  gr.DeleteUser,
	}
	return MapGinHandlers, MapFiberHandlers, MapGorillaHandlers
}
