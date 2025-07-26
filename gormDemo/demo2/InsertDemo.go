package main

import (
	"database/sql"
	"errors"
	"fmt"
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

func (b *Blog) BeforeCreate(tx *gorm.DB) (err error) {
	if b.Author.Birthday.IsZero() {
		b.Author.Birthday = time.Now()
	}
	return
}

func linkDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}

func parseTime(timeStr string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, errors.New("时间格式化出错")
	}
	return t, nil
}

func NewNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func main() {
	db := linkDB()
	//birthday, err := parseTime("1990-01-01")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	blogs := []*Blog{
		{Author: Author{
			Name: "赵云",
			Age:  19,
		},
			Email: NewNullString("elvis.zhan@foxmail.com")},
		{
			Author: Author{
				Name: "荀彧",
				Age:  21,
			},
			Email: NewNullString("elvis.zhan@foxmail.com"),
		},
	}
	//db.AutoMigrate(Blog{})

	// 新的泛型写法，类型在编译的时候就声明了
	//ctx := context.Background()
	//err = gorm.G[[]*Blog](db).Create(ctx, &blogs)// 批量插入或者单体插入

	//err := gorm.G[[]*Blog](db.Select("name", "age")).Create(ctx, &blogs) // 根据select的字段来插入，这里只插入Name和Age字段

	// 之前的传统写法，类型是运行的时候推导的
	db.Select("name", "age").Create(&blogs) // 根据select的字段来插入，这里只插入Name和Age字段

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println(blogs[0].ID)
}
