package gormDemo

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Author struct {
	Name     string `gorm:"size:255;unique;not null"`
	Age      int
	Birthday time.Time `gorm:"type:datetime"`
}

//type Blog struct {
//	Author        //使用匿名字段
//	Email  string `gorm:"size:255;unique;not null"`
//}

type Blog struct {
	gorm.Model
	Author Author         `gorm:"embedded"` // Author里的所有字段都加上author前缀，embedded表示作为嵌入结构体，和上面的写法是等效的
	Email  sql.NullString // 表示是一个允许为null的字符串
}

func LinkDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}
