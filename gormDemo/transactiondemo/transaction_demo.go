// 事务demo

package main

import (
	"github.com/1255177148/firstDemo/gormDemo"
	"gorm.io/gorm"
)

func main() {
	db := gormDemo.LinkDB()
	err := saveData(db, 6, "赵四")
	if err != nil {
		return
	}
}

func saveData(db *gorm.DB, id uint, name string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var blog gormDemo.Blog
		tx.First(&blog, id)
		if err := tx.Model(&gormDemo.Blog{}).Where("id = ?", id).Update("name", name).Error; err != nil {
			return err
		}
		if err := tx.Model(&gormDemo.Blog{}).Where("id = ?", id).Update("name", blog.Author.Name).Error; err != nil {
			return err
		}
		return nil
	})
}
