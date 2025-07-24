package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// 事件结构
type Event struct {
	BlockNumber uint64
	TxIndex     uint
	LogIndex    uint
	Data        string
}

// 按 TxIndex, LogIndex 排序事件
func sortEvents(events []Event) {
	sort.Slice(events, func(i, j int) bool {
		if events[i].TxIndex != events[j].TxIndex {
			return events[i].TxIndex < events[j].TxIndex
		}
		return events[i].LogIndex < events[j].LogIndex
	})
}

// 模拟写数据库
func writeEvents(blockNumber uint64, events []Event) {
	fmt.Printf("写入区块 %d 的 %d 个事件\n", blockNumber, len(events))
	for _, e := range events {
		fmt.Printf("  Tx %d Log %d Data: %s\n", e.TxIndex, e.LogIndex, e.Data)
		time.Sleep(10 * time.Millisecond)
	}
}

// 管理分块事件缓存和写入的控制器
type BlockEventManager struct {
	mu        sync.Mutex
	eventMap  map[uint64][]Event
	writeWg   sync.WaitGroup
	writeCh   chan uint64
	closeOnce sync.Once
	closed    chan struct{}
}

func NewBlockEventManager() *BlockEventManager {
	return &BlockEventManager{
		eventMap: make(map[uint64][]Event),
		writeCh:  make(chan uint64, 100),
		closed:   make(chan struct{}),
	}
}

// 添加事件
func (m *BlockEventManager) AddEvent(e Event) {
	m.mu.Lock()
	m.eventMap[e.BlockNumber] = append(m.eventMap[e.BlockNumber], e)
	m.mu.Unlock()

	// 通知写入事件，允许多次通知无害
	select {
	case m.writeCh <- e.BlockNumber:
	default:
	}
}

// 启动写入协程，监听写入请求
func (m *BlockEventManager) Run() {
	m.writeWg.Add(1)
	go func() {
		defer m.writeWg.Done()
		for {
			select {
			case blockNumber := <-m.writeCh:
				m.processBlock(blockNumber)
			case <-m.closed:
				return
			}
		}
	}()
}

// 处理单个区块事件写入
func (m *BlockEventManager) processBlock(blockNumber uint64) {
	m.mu.Lock()
	events := m.eventMap[blockNumber]
	if len(events) == 0 {
		m.mu.Unlock()
		return
	}
	// 拷贝事件后清空，防止重复写入
	eventsCopy := make([]Event, len(events))
	copy(eventsCopy, events)
	m.eventMap[blockNumber] = nil
	m.mu.Unlock()

	// 排序并写入
	sortEvents(eventsCopy)
	writeEvents(blockNumber, eventsCopy)
}

// 关闭写入器，等待完成
func (m *BlockEventManager) Close() {
	m.closeOnce.Do(func() {
		close(m.closed)
		m.writeWg.Wait()
	})
}

// -------- 模拟主流程 --------

func main() {
	manager := NewBlockEventManager()
	manager.Run()

	// 模拟并发接收乱序事件
	var wg sync.WaitGroup
	events := []Event{
		{BlockNumber: 1, TxIndex: 0, LogIndex: 0, Data: "A"},
		{BlockNumber: 1, TxIndex: 2, LogIndex: 0, Data: "C"},
		{BlockNumber: 1, TxIndex: 1, LogIndex: 0, Data: "B"},
		{BlockNumber: 2, TxIndex: 0, LogIndex: 1, Data: "D"},
		{BlockNumber: 2, TxIndex: 0, LogIndex: 0, Data: "E"},
		{BlockNumber: 3, TxIndex: 0, LogIndex: 0, Data: "F"},
	}

	wg.Add(len(events))
	for _, e := range events {
		ev := e
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // 模拟乱序
			fmt.Printf("接收事件: Block %d Tx %d Log %d Data %s\n", ev.BlockNumber, ev.TxIndex, ev.LogIndex, ev.Data)
			manager.AddEvent(ev)
		}()
	}

	wg.Wait()
	time.Sleep(500 * time.Millisecond) // 等待写入完成
	manager.Close()
}
