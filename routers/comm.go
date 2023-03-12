package routers

import (
	"demo/protos"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, protos.Success())
	})
	setupLogRouter(engine)
	return engine
}
