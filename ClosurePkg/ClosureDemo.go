package ClosurePkg

import "fmt"

type A struct {
	a int
}

/*
这是结构体的方法
*/
func (a A) add(i int) int {
	a.a += i
	return a.a
}

// 声明函数变量
var function1 func(int) int

// 声明闭包,将一个匿名函数赋给function2，就变成了闭包
var function2 func(int) int = func(i int) int {
	i += i
	return i
}

/*
这是一个匿名函数作为返回值
*/
func squareMath(i int) func(int) int {
	return func(x int) int {
		return i * x
	}
}

/*
接收一个匿名函数作为参数，并执行了该匿名函数
*/
func applyFunc(x int, f func(i int) int) int {
	return f(x)
}

func Demo() {
	a := A{1}
	// 将结构体的add方法赋给function1
	function1 = a.add
	addRes := function1(2)
	fmt.Println(addRes)
	function2(2)              //调用闭包
	square := squareMath(2)   // 这里返回了内部匿名函数，并记住了i是2
	y := applyFunc(3, square) // 这里是执行了匿名函数，让i乘以x
	fmt.Println(y)
}
