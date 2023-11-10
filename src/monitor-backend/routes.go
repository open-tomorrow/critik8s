package main

import (
	"github.com/gin-gonic/gin"
)

func routes(engine *gin.Engine) {

	engine.GET("/ping", ping)
	engine.GET("/auth", auth)
	engine.GET("/rules", rules)
}
