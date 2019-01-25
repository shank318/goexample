package pingdom

import (
	"examplego/middleware"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.Engine) {
	posts := r.Group("/api/urls")
	{
		posts.POST("/",middleware.Authorized, create)
		posts.GET("/",middleware.Authorized, list)
		posts.GET("/:id",middleware.Authorized,read)
		posts.DELETE("/:id/:status",middleware.Authorized, updateStatus)
	}
}
