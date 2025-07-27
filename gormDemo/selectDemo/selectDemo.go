package main

import (
	"database/sql"
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

func linkDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	return db
}

func NewNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

type BlogResult struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	db := linkDB()
	blog := Blog{}
	db.First(&blog)
	fmt.Println(blog.Author.Birthday)

	lastBlog := Blog{}
	db.Last(&lastBlog)
	fmt.Println(lastBlog.Author.Name)

	var blogs []*Blog
	db.Where("name in ?", []string{"张飞", "刘备", "关羽"}).Find(&blogs)
	for _, blog := range blogs {
		fmt.Println(blog.Author.Name)
	}

	timeBlog := Blog{}
	db.Where("birthday", time.Date(1990, 1, 1, 8, 0, 0, 0, time.Local)).Take(&timeBlog)
	fmt.Println("根据时间来查询", timeBlog.Author.Name)

	var timeBlogs []*Blog
	start := time.Date(1990, 1, 1, 8, 0, 0, 0, time.Local)
	end := time.Now()
	db.Where("birthday between ? and ?", start, end).Find(&timeBlogs)
	fmt.Println("根据生日范围来查询---")
	for _, blog := range timeBlogs {
		fmt.Println(blog.Author.Name)
	}

	// 复杂查询

	// from子查询
	var blogArr []*Blog
	db.Table("(?) as d ", db.Model(&Blog{}).Select("name", "age", "deleted_at")).Where("age = ?", 18).Find(&blogArr)
	fmt.Println("from子查询---")
	for _, blog := range blogArr {
		fmt.Println(blog.Author.Name)
	}

	// where 子查询
	var blogList []*Blog
	db.Where("age > (?)", db.Table("blogs").Select("AVG(age)")).Limit(5).Find(&blogList)
	fmt.Println("where子查询---")
	for _, blog := range blogList {
		fmt.Println(blog.Author.Name)
	}

	// 使用map来接收查询的字段，这样可以适用于不清楚表字段的查询，也就是没有表结构体，或者结构体是动态的
	var resultMap map[string]interface{}
	db.Model(&Blog{}).Where("name = ?", "张三").Find(&resultMap)
	fmt.Println(resultMap)

	// 分批处理
	fmt.Println("分批处理查询---")
	var blogResult []BlogResult
	// FindInBatches反射结构体，不能反射map切片，所以只能用结构体来接收
	result := db.Table("blogs").Where("age > (?)", db.Table("blogs").Select("AVG(age)")).FindInBatches(&blogResult, 3, func(tx *gorm.DB, batch int) error {
		for _, blog := range blogResult {
			fmt.Println(blog.Name)
		}
		return nil
	})
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	// 分组统计
	var countMap []map[string]interface{}
	db.Table("blogs").Select("GROUP_CONCAT(name separator ',') as names", "count(1) as count", "age").Group("age").Find(&countMap)
	fmt.Println("分组统计------")
	for _, value := range countMap {
		fmt.Println(value)
	}
}
