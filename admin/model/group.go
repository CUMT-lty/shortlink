package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Gid       string
	Name      string
	Username  string
	SortOrder string
}

func (g *Group) TableName() string {
	return "groups"
}
