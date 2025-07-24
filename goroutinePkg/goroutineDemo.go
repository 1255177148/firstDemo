package main

import (
	"fmt"
	"sync"
)

func sayHello(s string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done() // 表示此函数逻辑执行完毕之后再执行这行代码
	}
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2) // 表示要等待2个goroutine完成
	go sayHello("go run", &wg)
	go func(s string, wg *sync.WaitGroup) {
		defer wg.Done() // 表示这个goroutine完成了
		fmt.Println(s)
	}("go anonymity function", &wg)
	sayHello("go main", nil)
	wg.Wait() // 等待其他协程完成
}
