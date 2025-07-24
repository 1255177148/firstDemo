package main

import (
	"fmt"
	"sync"
)

// ch1 只声明channel,没有初始化，为nil，没有初始化的话不能用ch1收发消息，
var ch1 chan int

// ch2 初始化一个channel，可以用来收发消息
var ch2 chan int = make(chan int)

// ch3 初始化一个带有3个缓冲区的channel
var ch3 chan string = make(chan string, 3)

func say() {
	for i := 0; i < 3; i++ {
		ch3 <- fmt.Sprintf("message from ch3: %d", i)
	}
	close(ch3) // 关闭通道，不然的话，主程序会一直等待
}

// sendOnly 只能发送消息到通道
func sendOnly(ch chan<- string, msg string) {
	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("message from ch: %d，message is ：%s", i, msg)
	}
	close(ch)
}

// readOnly 只能接受通道发送的消息
func readOnly(ch <-chan string) {
	for val := range ch {
		fmt.Printf("message from ch: %s", val)
		fmt.Println()
	}
}

var done = make(chan struct{})
var addChan = make(chan int)

// add 执行累加计算的函数
func add() {
	count := 0
	for n := range addChan {
		count += n
	}
	fmt.Println("累加结果为：", count)
	close(done)
}

// handleWork 使用channel完成并发累加任务
func handleWork(workCount int, wg *sync.WaitGroup) {

	work := func(id int) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			addChan <- 1
		}
	}
	for i := 0; i < workCount; i++ {
		go work(i)
	}
}

func main() {
	go say()
	for val := range ch3 {
		fmt.Println(val)
	}

	ch := make(chan string, 3)
	go sendOnly(ch, "hello")
	readOnly(ch)

	// 使用select来监听通道

	// timeout设置一个5秒之后发送的通道
	//timeout := time.After(1 * time.Second)
	//
	//for {
	//	select {
	//	case msg, ok := <-ch:
	//		if !ok {
	//			fmt.Println("Channel已关闭")
	//			return
	//		}
	//		fmt.Println("主goroutine收到：", msg)
	//	case <-timeout:
	//		fmt.Println("操作超时")
	//		return
	//	default:
	//		fmt.Println("主goroutine还未收到消息")
	//	}
	//}

	wg := sync.WaitGroup{}
	workCount := 1000
	wg.Add(workCount)
	go add()                   // 启动接受者，执行累加计算
	handleWork(workCount, &wg) // 启动发送者，发送执行累加的命令
	wg.Wait()
	close(addChan)
	// 等待done通道关闭，然后再结束程序
	<-done
}
