package users

import (
	"bongdaphui/bongdaphui/database/models"
	"bongdaphui/bongdaphui/lib/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Team = models.Team
type User = models.User
type JSON = common.JSON

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	user := c.MustGet("user").(User)

	type RequestBody struct {
		Email  string `json:"email" binding:"required"`
		Mobile string `json:"mobile" binding:"required"`
		Avatar string `json:"avatar"`
	}

	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	if user.ID < 1 {
		c.AbortWithStatus(401)
		return
	}

	user.Email = requestBody.Email
	user.Mobile = requestBody.Mobile
	user.Avatar = requestBody.Avatar
	db.Save(&user)
	c.JSON(200, common.GenerateResponse(0, "", user.Serialize()))
}

func detail(c *gin.Context) {
	user := c.MustGet("user").(User)

	if user.ID < 1 {
		c.AbortWithStatus(401)
		return
	}

	c.JSON(200, common.GenerateResponse(0, "", user.Serialize()))
}
