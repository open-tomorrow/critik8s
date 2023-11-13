package api

import (
	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {

	engine.GET("/ping", Ping)
	engine.GET("/auth", Auth)
	engine.GET("/rules", Rules)
}
