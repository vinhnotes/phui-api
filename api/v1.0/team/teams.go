package teams

import (
	"bongdaphui/lib/middlewares"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	posts := r.Group("/teams")
	{
		posts.POST("/", middlewares.Authorized, create)
		posts.GET("/", list)
		posts.GET("/:id", read)
		posts.DELETE("/:id", middlewares.Authorized, remove)
		posts.PATCH("/:id", middlewares.Authorized, update)
	}
}
