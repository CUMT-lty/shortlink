package model

import (
	"github.com/lty/shortlink/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	RealName string
	Phone    string
	Email    string
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUserByUsername(username string) User { // TODO：这里看后期怎么改
	var user User
	db.DB.Where("username = ?", username).Take(&user)
	return user
}

func (u *User) InsertUser() error { // TODO：这里需不需要返回 id
	err := db.DB.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
