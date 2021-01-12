package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func startGinServer() {
	//router := gin.New() // 干净的没有中间件
	router := gin.Default() // 会包含 logger 和 recover 中间件
	router.GET("/:name", func(c *gin.Context) {
		c.String(http.StatusOK, "%s 你好!", c.Params.ByName("name"))
	})
	router.Run(":3000")
}
