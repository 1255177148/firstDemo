package casewhen

import (
	"fmt"
	"github.com/1255177148/firstDemo/gormDemo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

// BuildCaseWhenExpr 生成单字段 CASE WHEN 表达式
func BuildCaseWhenExpr(field string, changeField string, values map[interface{}]interface{}) (clause.Expr, []interface{}) {
	sql := fmt.Sprintf("CASE %s ", field)
	var vars []any
	var ids []any

	for id, v := range values {
		sql += "WHEN ? THEN ? "
		vars = append(vars, id, v)
		ids = append(ids, id)
	}

	sql += fmt.Sprintf("ELSE %s END", changeField)
	return gorm.Expr(sql, vars...), ids
}

// BuildMultiCaseExpr 生成多字段 CASE WHEN 表达式
func BuildMultiCaseExpr(idField string, updateFields map[string]map[any]any) (map[string]interface{}, []any) {
	result := map[string]interface{}{}
	var ids []any
	first := true

	for field, values := range updateFields {
		sql := fmt.Sprintf("CASE %s ", idField)
		var vars []interface{}
		for id, v := range values {
			sql += "WHEN ? THEN ? "
			vars = append(vars, id, v)
			if first {
				ids = append(ids, id)
			}
		}
		sql += fmt.Sprintf("ELSE %s END", field)
		result[field] = gorm.Expr(sql, vars...)
		first = false
	}
	return result, ids
}

// BuildFromStructSlice 支持从结构体切片自动生成 CASE WHEN
// struct 里必须有 `ID` 字段，其他字段将自动生成表达式
func BuildFromStructSlice[T any](slice []T, idField string) (map[string]any, []any) {
	if len(slice) == 0 {
		return nil, nil
	}

	fields := map[string]map[any]any{}
	var ids []any

	tType := reflect.TypeOf(slice[0])
	for i := 0; i < tType.NumField(); i++ {
		field := tType.Field(i)
		if field.Name == "ID" {
			// 跳过id，因为默认id是主键，一般是根据主键来更新别的字段
			continue
		}
		fields[field.Name] = map[any]any{}
	}

	for _, item := range slice {
		v := reflect.ValueOf(item)
		id := v.FieldByName("ID").Interface()
		ids = append(ids, id)

		for i := 0; i < tType.NumField(); i++ {
			field := tType.Field(i)
			if field.Name == "ID" {
				continue
			}
			fields[field.Name][id] = v.Field(i).Interface()
		}
	}

	return BuildMultiCaseExpr(idField, fields)
}

// 根据case when修改单个字段的demo
func demoBuildCaseWhenExpr() {
	db := gormDemo.LinkDB()
	// 下面是根据case when修改单个字段
	values := map[any]any{
		1: "张大炮",
		2: "李四",
	}
	expr, ids := BuildCaseWhenExpr("id", "name", values)
	db.Model(&gormDemo.Blog{}).Where("id in ?", ids).Update("name", expr)
}

// case when 修改多个字段
func demoBuildMultiCaseExpr() {
	db := gormDemo.LinkDB()
	// 下面是case when 修改多个字段
	updateFields := map[string]map[any]any{
		"name": {1: "张三", 2: "老李"},
		"age":  {1: 20, 2: 21},
	}
	expr, ids := BuildMultiCaseExpr("id", updateFields)
	db.Model(&gormDemo.Blog{}).Where("id in ?", ids).Updates(expr)
}

// case when 修改多个字段，传入的是结构体参数
func demoBuildFromStructSlice() {
	db := gormDemo.LinkDB()
	updateFields := []struct {
		ID   int
		Name string
		Age  int
	}{
		{ID: 1, Name: "张大炮", Age: 21},
		{ID: 2, Name: "李四", Age: 22},
	}
	expr, ids := BuildFromStructSlice(updateFields, "id")
	db.Model(&gormDemo.Blog{}).Where("id in ?", ids).Updates(expr)
}
