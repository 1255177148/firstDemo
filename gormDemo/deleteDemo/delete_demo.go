package main

import (
	"github.com/1255177148/firstDemo/gormDemo"
)

func main() {
	db := gormDemo.LinkDB()
	ids := []int{1, 2, 3}
	db.Model(&gormDemo.Blog{}).Where("id in (?)", ids).Delete(&gormDemo.Blog{})
}
