package middleware

import (
	"github.com/gin-gonic/gin"
)

func Authorized(c *gin.Context) {
	user :=c.GetHeader("username")
	if user != "1345" {
		c.AbortWithStatus(401)
		return
	}
}
