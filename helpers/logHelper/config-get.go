package logHelper

import (
	"demo/protos"
	"demo/services/zlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetConfig(c *gin.Context) {
	data := protos.LogConfig{
		Level: zlog.GetLevel(),
	}
	c.JSON(http.StatusOK, protos.Success(data))
}
