package user

import (
	"github.com/lty/shortlink/admin/common/convertion/result"
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

func (u *User) AddUser() {
	err := db.DB.Create(&u).Error
	if err != nil {
		panic(err)
	}
}

func (u *User) GetUserByUsername() result.Result {
	return
}
