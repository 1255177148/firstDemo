package vardefinition

// 全局变量的声明
var a string = "第一种全局变量"
var b string
var c = "第三种全局变量"
var (
	d uint64 = 1
	e bool
	f bool = true
)

/*
局部变量声明，和全局变量不一样的写法
*/
func method1() (var1 string, var2 uint8) {
	return "局部变量", 2
}

/*
局部变量声明，和全局变量不一样的第二种写法
*/
func method2() (var1 string, var2 uint8) {
	var1 = "局部变量"
	var2 = 3
	return
}

/*
局部变量声明，和全局变量不一样的第三种写法，声明后不赋值
*/
func method3() (var1 string, var2 uint8) {
	return
}

