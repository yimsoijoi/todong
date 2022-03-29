package ginserver

import (
	"log"

	"github.com/artnoi43/todong/enums"
	"github.com/artnoi43/todong/lib/handler"
	"github.com/artnoi43/todong/lib/middleware"
)

func (g *ginServer) SetUpRoutes(conf *middleware.Config, handler handler.Adaptor) {
	auth, err := middleware.AuthenticateGin(conf)
	if err != nil {
		log.Fatal("error: failed to create JWT authentication middleware")
	}
	authMiddlewareFunc := auth.MiddlewareFunc()
	// Setup routes
	todoAPI := g.engine.Group("/todos")
	userAPI := g.engine.Group("/users")
	todoAPI.GET("/", authMiddlewareFunc, handler.Gin(enums.HandlerGetTodo))
	todoAPI.GET("/:uuid", authMiddlewareFunc, handler.Gin(enums.HandlerGetTodo))
	todoAPI.POST("/create", authMiddlewareFunc, handler.Gin(enums.HandlerCreateTodo))
	todoAPI.DELETE("/:uuid", authMiddlewareFunc, handler.Gin(enums.HandlerDeleteTodo))
	todoAPI.POST("/update/:uuid", authMiddlewareFunc, handler.Gin(enums.HandlerUpdateTodo))
	userAPI.POST("/register", handler.Gin(enums.HandlerRegister))
	userAPI.POST("/login", handler.Gin(enums.HandlerLogin))
	userAPI.DELETE("/", authMiddlewareFunc, handler.Gin(enums.HandlerDeleteUser))
	userAPI.POST("/new-password", authMiddlewareFunc, handler.Gin(enums.HandlerNewPassword))
	// For testing JWT
	userAPI.GET("/testauth", authMiddlewareFunc, handler.Gin(enums.HandlerTestAuth))
}
