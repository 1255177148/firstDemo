package processControl

import "fmt"

func IfDemo1() {
	a := 10
	// 这里可以在表达式前声明一个变量，这个变量作用域只在if作用域内
	if b := 1; a > 10 {
		b = 2
		fmt.Println("a > 10")
	} else if c := 3; b > 1 {
		b = 3
		fmt.Println("b > 1")
	} else {
		fmt.Println("other")
		if c == 3 {
			fmt.Println("c == 3")
		}
		fmt.Println(a)
		fmt.Println(b)
	}
}

/*
普通的switch用法
*/
func SwitchDemo1() {
	a := "test switch"
	switch a {
	case "test":
		fmt.Println("a == test")
	case "t", "test switch":
		fmt.Println("a == t test")
	case "switch":
		fmt.Println("a == switch")
	default:
		fmt.Println("a != test")
	}
}

/*
switch用法二，switch里声明一个变量，该变量只在switch作用域内
*/
func SwitchDemo2() {
	switch b := 5; b {
	case 1:
		fmt.Println("b == 1")
	case 2:
		fmt.Println("b == 2")
	case 3:
		fmt.Println("b == 3")
	case 4, 5:
		fmt.Println("b == 4 || 5")
	default:
		fmt.Println("other")

	}
}

/*
switch的第三种写法，在case里写表达式
*/
func SwitchDemo3() {
	a := "test switch"
	b := 5
	switch {
	case a == "test":
		fmt.Println("a == test")
	case b == 1:
		fmt.Println("b == 1")
	case a == "test switch", b == 4:
		fmt.Println("a == test switch")
	case b == 5:
		fmt.Println("b == 5")
	default:
		fmt.Println("other")

	}
}

/*
switch的第四种写法，适用于接口和泛型
*/
func SwitchDemo4() {
	var d any
	d = 1
	switch t := d.(type) {
	case byte:
		fmt.Println("byte", t)
	case string:
		fmt.Println("string", t)
	case int:
		fmt.Println("int", t)
	default:
		fmt.Println("other")
	}
}
