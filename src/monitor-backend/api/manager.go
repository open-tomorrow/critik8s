package api

import (
	"monitor-backend/amqp"
	"monitor-backend/config"
	"monitor-backend/models"
	"monitor-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var monitorConfig = config.GetMonitorConfig()

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "paolo",
	})
}

func Auth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, monitorConfig)
}

func Rules(c *gin.Context) {
	rule := utils.ParseFile[models.Rule]("rules/RWO.yaml")

	c.JSON(200, rule)
}

func Nodes(c *gin.Context) {
	nodes := amqp.GetNodes()
	c.JSON(200, nodes)
}

func Pods(c *gin.Context) {
	pods := amqp.GetPods()
	c.JSON(200, pods)
}
