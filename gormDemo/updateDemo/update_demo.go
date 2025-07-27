package main

import (
	"github.com/1255177148/firstDemo/gormDemo"
	"github.com/1255177148/firstDemo/gormDemo/utils/casewhen"
)

func main() {
	db := gormDemo.LinkDB()
	//db.Model(&gormDemo.Blog{}).Where("name = ?", "张三").Updates(map[string]interface{}{"name": "张大炮", "age": 25})

	// 使用row，也就是写原始的sql代码，这样可以用来处理复杂sql
	//	db.Exec(`
	//update blogs set name =
	//case
	//when id = ? then ?
	//when id = ? then ?
	//end
	//where id in (?,?)
	//`, 1, "张三", 2, "老李", 1, 2)

	// 下面是根据case when修改单个字段
	//values := map[any]any{
	//	1: "张大炮",
	//	2: "李四",
	//}
	//expr, ids := casewhen.BuildCaseWhenExpr("id", "name", values)
	//db.Model(&gormDemo.Blog{}).Where("id in ?", ids).Update("name", expr)

	// 下面是case when 修改多个字段
	//updateFields := map[string]map[any]any{
	//	"name": {1: "张三", 2: "老李"},
	//	"age":  {1: 20, 2: 21},
	//}
	//expr, ids := casewhen.BuildMultiCaseExpr("id", updateFields)
	//fmt.Println(expr)
	//db.Model(&gormDemo.Blog{}).Where("id in ?", ids).Updates(expr)

	// 下面是case when 修改多个字段，传入的是结构体切片
	updateFields := []struct {
		ID   int
		Name string
		Age  int
	}{
		{ID: 1, Name: "张大炮", Age: 21},
		{ID: 2, Name: "李四", Age: 22},
	}
	expr, ids := casewhen.BuildFromStructSlice(updateFields, "id")
	db.Model(&gormDemo.Blog{}).Where("id in ?", ids).Updates(expr)
}
