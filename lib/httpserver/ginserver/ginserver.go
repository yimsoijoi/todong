package ginserver

import "github.com/gin-gonic/gin"

type ginServer struct {
	engine *gin.Engine
}

func New() *ginServer {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return &ginServer{
		engine: r,
	}
}
