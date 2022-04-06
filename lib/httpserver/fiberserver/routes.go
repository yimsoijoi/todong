package fiberserver

import (
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/yimsoijoi/todong/enums"
	"github.com/yimsoijoi/todong/lib/handler"
	"github.com/yimsoijoi/todong/lib/middleware"
)

func (f *fiberServer) SetUpRoutes(conf *middleware.Config, handler handler.Adaptor) {
	f.app.Use(logger.New())
	authMiddlewareFunc := middleware.AuthenticateFiber(conf)
	todoApi := f.app.Group("/todos")
	userApi := f.app.Group("/users")
	userApi.Post("/register", handler.Fiber(enums.HandlerRegister))
	userApi.Post("/login", handler.Fiber(enums.HandlerLogin))
	userApi.Post("/new-password", authMiddlewareFunc, handler.Fiber(enums.HandlerNewPassword))
	userApi.Delete("/", authMiddlewareFunc, handler.Fiber(enums.HandlerDeleteUser))
	todoApi.Get("/", authMiddlewareFunc, handler.Fiber(enums.HandlerGetTodo))
	todoApi.Get("/:uuid", authMiddlewareFunc, handler.Fiber(enums.HandlerGetTodo))
	todoApi.Post("/create", authMiddlewareFunc, handler.Fiber(enums.HandlerCreateTodo))
	todoApi.Post("/update/:uuid", authMiddlewareFunc, handler.Fiber(enums.HandlerUpdateTodo))
	todoApi.Delete("/:uuid", authMiddlewareFunc, handler.Fiber(enums.HandlerDeleteTodo))
}
