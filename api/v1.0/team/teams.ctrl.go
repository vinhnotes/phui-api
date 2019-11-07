package teams

import (
	"bongdaphui/database/models"
	"bongdaphui/lib/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Team = models.Team
type User = models.User
type JSON = common.JSON

func create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Level       uint    `json:"level" binding:"required"`
		CityID      uint    `json:"city_id"`
		WardID      uint    `json:"ward_id"`
		DistrictID  uint    `json:"district_id"`
		Since       string  `json:"since"`
		Address     string  `json:"address" binding:"required"`
		Web         string  `json:"web"`
		FacebookID  string  `json:"facebook_id"`
		Cover       string  `json:"cover"`
		Lng         float64 `json:"lng"`
		Lat         float64 `json:"lat"`
	}
	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	user := c.MustGet("user").(User)
	team := Team{
		Name:        requestBody.Name,
		Description: requestBody.Description,
		Level:       requestBody.Level,
		CityID:      requestBody.CityID,
		WardID:      requestBody.WardID,
		DistrictID:  requestBody.DistrictID,
		Since:       requestBody.Since,
		Address:     requestBody.Address,
		Web:         requestBody.Web,
		FacebookID:  requestBody.FacebookID,
		Cover:       requestBody.Cover,
		Lng:         requestBody.Lng,
		Lat:         requestBody.Lat,
		OwnerID:     user.ID,
		User:        user,
	}
	db.NewRecord(team)
	db.Create(&team)
	c.JSON(200, common.GenerateResponse(0, "", team.Serialize()))
}

func list(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	offset := common.GetPage(c)
	cursor := c.Query("cursor")
	recent := c.Query("recent")

	var teams []Team

	if cursor == "" {
		if err := db.Preload("User").Limit(common.PageSize).Offset(offset).Order("id desc").Find(&teams).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		condition := "id < ?"
		if recent == "1" {
			condition = "id > ?"
		}
		if err := db.Preload("User").Limit(common.PageSize).Offset(offset).Order("id desc").Where(condition, cursor).Find(&teams).Error; err != nil {
			c.AbortWithStatus(500)
			return
		}
	}

	length := len(teams)
	serialized := make([]JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = teams[i].Serialize()
	}

	c.JSON(200, common.GenerateResponse(0, "", common.JSON{
		"teams": serialized,
		"total": len(serialized),
	}))
}

func read(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var team Team

	// auto preloads the related model
	// http://gorm.io/docs/preload.html#Auto-Preloading
	if err := db.Set("gorm:auto_preload", true).Where("id = ?", id).First(&team).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, team.Serialize())
}

func remove(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	var team Team
	if err := db.Where("id = ?", id).First(&team).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if team.OwnerID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	db.Delete(&team)
	c.Status(204)
}

func update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	user := c.MustGet("user").(User)

	type RequestBody struct {
		Name string `json:"name" binding:"required"`
	}

	var requestBody RequestBody

	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatus(400)
		return
	}

	var team Team
	if err := db.Preload("User").Where("id = ?", id).First(&team).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if team.OwnerID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	team.Name = requestBody.Name
	db.Save(&team)
	c.JSON(200, team.Serialize())
}
