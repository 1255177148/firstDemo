package mathPkg

import "fmt"

/*
*
简单的四则运算
*/
func math() {
	a, b := 1, 2
	sum := a + b
	fmt.Println(sum)
	sub := a - b
	fmt.Println(sub)
	mul := a * b
	fmt.Println(mul)
	div := a / b
	fmt.Println(div)
	mod := a % b // 取余
	fmt.Println(mod)
	a += 1
	fmt.Println(a)
	b -= 1
	fmt.Println(b)
}

/*
*
关系运算符
*/
func relationshipMath() {
	a := 1
	b := 5
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a < b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}

/*
逻辑运算符
*/
func logicMath() {
	a := true
	b := false
	fmt.Println(a && b)
	fmt.Println(!(a && b))
	fmt.Println(a || b)
}
