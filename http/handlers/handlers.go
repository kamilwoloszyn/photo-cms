package handlers

import "github.com/gin-gonic/gin"

func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func AddPaperTypeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "added",
	})
}

func DeletePaperTypeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "removed",
	})
}
