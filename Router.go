package main

import (
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	engine := gin.Default()

	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, Error(404, "无效的路由"))
	})

	engine.POST("/monitor/create", func(c *gin.Context) {
		c.JSON(http.StatusOK, CreateMonitor(c))
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "系统异常，请联系客服",
		"code": 1001,
	})
}
