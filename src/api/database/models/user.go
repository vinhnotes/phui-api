package models

import (
	"bongdaphui/lib/common"

	"github.com/jinzhu/gorm"
)

// User data model
type User struct {
	gorm.Model
	Email        string
	Mobile       string
	Avatar       string
	Enable       bool
	DisplayName  string
	PasswordHash string
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"email":        u.Email,
		"display_name": u.DisplayName,
		"mobile":       u.Mobile,
		"avatar":       u.Avatar,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	if m["email"] != nil {
		u.Email = m["email"].(string)
	}
	if m["display_name"] != nil {
		u.Email = m["display_name"].(string)
	}
	if m["mobile"] != nil {
		u.Email = m["mobile"].(string)
	}
	if m["avatar"] != nil {
		u.Email = m["avatar"].(string)
	}
	if m["enable"] != nil {
		u.Enable = m["enable"].(bool)
	}
}
