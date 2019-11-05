package users

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.PUT("/update", update)
		users.GET("/", detail)
	}
}
