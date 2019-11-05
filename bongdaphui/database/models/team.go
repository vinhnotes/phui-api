package models

import (
	"bongdaphui/bongdaphui/lib/common"

	"github.com/jinzhu/gorm"
)

// Stadium data model
type Team struct {
	gorm.Model
	Name        string
	Description string `sql:"type:text;"`
	Level       uint
	Code        string
	CityID      uint
	WardID      uint
	DistrictID  uint
	Since       string
	Address     string
	Web         string
	FacebookID  string
	Cover       string
	Lng         float64
	Lat         float64
	User        User `gorm:"foreignkey:OwnerID"`
	OwnerID     uint `gorm:"column:owner_id" json:"owner_id"`
}

// Serialize serializes user data
func (team *Team) Serialize() common.JSON {
	return common.JSON{
		"id":          team.ID,
		"name":        team.Name,
		"description": team.Description,
		"level":       team.Level,
		"code":        team.Code,
		"city_id":     team.CityID,
		"ward_id":     team.WardID,
		"district_id": team.DistrictID,
		"since":       team.Since,
		"address":     team.Address,
		"web":         team.Web,
		"facebook_id": team.FacebookID,
		"cover":       team.Cover,
		"lng":         team.Lng,
		"lat":         team.Lat,
		"user":        team.User.Serialize(),
	}
}

func (team *Team) Read(m common.JSON) {
	team.ID = uint(m["id"].(float64))
	team.Name = m["name"].(string)
	team.Description = m["description"].(string)
	team.Level = m["level"].(uint)
	team.Code = m["code"].(string)
	team.CityID = m["city_id"].(uint)
	team.WardID = m["ward_id"].(uint)
	team.DistrictID = m["district_id"].(uint)
	team.Since = m["since"].(string)
	team.Address = m["address"].(string)
	team.Web = m["web"].(string)
	team.FacebookID = m["facebook_id"].(string)
	team.Cover = m["cover"].(string)
	team.Lng = m["lng"].(float64)
	team.Lat = m["lat"].(float64)
	team.User = m["user"].(User)
}
