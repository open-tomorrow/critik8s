package api

import (
	"monitor-backend/config"
	"monitor-backend/models"
	"monitor-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var monitorConfig = config.Get()

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
