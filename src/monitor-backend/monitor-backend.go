package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	c := config()

	engine := gin.Default()
	engine.Use(CORSMiddleware())

	routes(engine, c)

	engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
