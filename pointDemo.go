package main

import "fmt"

// 指针demo
var p1 *int
var p2 *string

func demoPoint() {
	i := 1
	p1 = &i
	pp := &p1
	s := "hello"
	p2 = &s
	fmt.Println(p1)
	fmt.Println(*p1)
	*p1 += 1
	fmt.Println(i)
	**pp += 1
	fmt.Println(i)
	fmt.Println(p2)
	fmt.Println(*p2)
}
