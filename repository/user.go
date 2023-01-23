package repository

import (
	"fmt"
)

type User struct {
	ID             uint64 `gorm:"primarykey"`
	Username       string
	Name           string
	Avatar         string
	Password       string
	Follower_count uint64
	Follow_count   uint64
}

type UserInfo struct {
	ID       uint   `json:"id"`
	Avatar   string `json:"avatar"`
	UserName string `json:"userName"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func GetUserByUsername(username string) User {
	var user User
	db.Where("username = ?", username).First(&user)
	return user
}

func GetUserById(Id uint64) User {
	var user User
	db.Where("id = ?", Id).First(&user)
	return user
}

func CreateUser(u *User) uint64 {
	db.Create(u)
	return u.ID
}

func GenerateA() {
	user := User{Name: "yly", Username: "1234", Password: "1234"}
	result := db.Create(&user)
	fmt.Printf("userId=%d,err=%s,affected=%d", user.ID, result.Error, result.RowsAffected)
}
