// 模型定义，定义表结构，创建数据库连接
// 官方文档 https://gorm.io/zh_CN/docs/

package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Name     string  // 非null的字符串，如果为空，写到数据库里会是一个空字符而不是null
	Email    *string `gorm:"size:50"` // 定义该字段的长度为50，允许为null的字符串
	Age      uint8
	Birthday *time.Time `gorm:"type:datetime"` // 可以为null的日期
	gorm.Model
}

func create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())

	}
	//if err := db.AutoMigrate(&User{}); err != nil {
	//	panic("failed to auto migrate database: " + err.Error())
	//}
	user := &User{}
	err = create(db, user)
	if err != nil {
		return
	}
}
