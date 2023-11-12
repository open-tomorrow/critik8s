package main

import (
	"monitor-backend/amqp"
	"monitor-backend/api"

	"github.com/gin-gonic/gin"
)

func main() {

	conn := amqp.Init()
	defer conn.Close()

	engine := gin.Default()
	engine.Use(api.CORSMiddleware())

	api.Routes(engine)

	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
