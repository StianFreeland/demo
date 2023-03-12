package routers

import (
	"demo/handlers/log"
	"github.com/gin-gonic/gin"
)

func setupLogRouter(engine *gin.Engine) {
	engine.GET("/log/config", log.GetConfig)
	engine.PUT("/log/config", log.UpdateConfig)
}
