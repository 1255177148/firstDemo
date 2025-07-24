// 并发的线程安全问题
// 使用锁完成计数累加并发
package main

import (
	"fmt"
	"sync"
)

// SafeCounter 一个线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc 增加计数
// 线程安全的操作
func (c *SafeCounter) Inc() {
	c.mu.Lock()         // 上互斥锁
	defer c.mu.Unlock() // 最后执行解锁
	c.count++
}

// Value 读取计数
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// UnsafeCounter 一个线程不安全的计数器
type UnsafeCounter struct {
	count int
}

// Inc 增加计数
func (c *UnsafeCounter) Inc() {
	c.count++
}

// Value 读取计数
func (c *UnsafeCounter) Value() int {
	return c.count
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2000)
	// 线程不安全的计数器
	counter := UnsafeCounter{count: 0}
	for i := 0; i < 1000; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Inc()
			}
		}(&wg)
	}
	// 线程安全的计数器
	safeCounter := SafeCounter{count: 0}
	for i := 0; i < 1000; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				safeCounter.Inc()
			}
		}(&wg)
	}
	wg.Wait()
	fmt.Println("线程不安全的计数器最终的计数为：", counter.Value())
	fmt.Println("线程安全的计数器最终的计数为：", safeCounter.Value())
}
