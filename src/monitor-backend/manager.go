package main

import (
	"monitor-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var config = getConfig()

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "paolo",
	})
}

func auth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config)
}

func rules(c *gin.Context) {
	rule := utils.ParseFile[Rule]("rules/RWO.yaml")

	c.JSON(200, rule)
}
