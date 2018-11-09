package core

import (
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
