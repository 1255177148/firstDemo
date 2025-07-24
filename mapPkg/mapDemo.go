package main

import "fmt"

func main() {
	maps := make(map[int]string)
	maps[0] = "hello"
	maps[1] = "world"
	fmt.Println(maps)
	value, ok := maps[0]
	fmt.Println(value)
	fmt.Println(ok)
	val, ok := maps[2]
	fmt.Println(val)
	fmt.Println(ok)
	delete(maps, 0)
	fmt.Println(maps)
	modifyMap(maps)
	fmt.Println(maps)
}

func modifyMap(m map[int]string) {
	m[1] = "hello"
}
