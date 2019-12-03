package apiv1

import (
	"bongdaphui/api/v1.0/auth"
	teams "bongdaphui/api/v1.0/team"
	users "bongdaphui/api/v1.0/user"
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// ApplyRoutes applies router to the gin Engine
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		v1.GET("/ping", ping)
		auth.ApplyRoutes(v1)
		teams.ApplyRoutes(v1)
		users.ApplyRoutes(v1)
	}
}
