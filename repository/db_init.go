package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	user := "root"
	password := "qwerty666"
	host := "127.0.0.1"
	port := "3306"
	database := "douyin"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		database,
		charset)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 写操作将不会进入事务, 提高性能
		PrepareStmt:            true, // 启用预编译缓存，提高性能
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	return err
}
