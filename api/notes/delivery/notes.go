package delivery

import (
	"examplego/middleware"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.Engine, handler NotesHandler) {
	posts := r.Group("/api/notes")
	{

		posts.POST("/",middleware.Authorized, handler.create)
		posts.GET("/",middleware.Authorized, list)
		posts.GET("/:id",middleware.Authorized,read)
		posts.DELETE("/:id",middleware.Authorized, remove)
		posts.PATCH("/:id",middleware.Authorized, update)
	}
}
