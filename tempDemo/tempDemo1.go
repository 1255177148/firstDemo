package main

import (
	"database/sql"
	"log"
	"sync"
	"time"
)

// Transaction 表示区块链中的一笔交易
type Transaction struct {
	ID      string  `json:"id"`      // 交易ID
	From    string  `json:"from"`    // 发送方地址
	To      string  `json:"to"`      // 接收方地址
	Amount  float64 `json:"amount"`  // 转账金额
	Gas     uint64  `json:"gas"`     // 交易手续费
	Payload []byte  `json:"payload"` // 附加数据(智能合约调用等)
}

type Block struct {
	Height       uint64         // 区块高度/区块号
	Hash         string         // 当前区块的哈希值
	PrevHash     string         // 前一个区块的哈希值(形成链式结构)
	Timestamp    int64          // 区块创建时间戳(Unix时间)
	Transactions []*Transaction `json:"transactions"` // 交易列表
	Nonce        uint64         // 工作量证明(PoW)使用的随机数
	Difficulty   uint64         // 当前区块的挖矿难度
	Miner        string         // 矿工地址(打包该区块的节点)
	StateRoot    string         // 状态树根哈希(Merkle Patricia Tree)
	ReceiptsRoot string         // 交易回执树根哈希
	// 其他可能的字段...
}

type BlockProcessor struct {
	blockChan     chan *Block       // 接收新区块的通道
	pendingBlocks map[uint64]*Block // 等待队列(高度->区块)
	db            *sql.DB
	lastBlock     uint64 // 最后一个处理过的区块高度
	mu            sync.Mutex
	pendingCond   *sync.Cond    // 用于通知有新区块到达的条件变量
	shutdownChan  chan struct{} // 关闭信号
}

func NewBlockProcessor(db *sql.DB) *BlockProcessor {
	bp := &BlockProcessor{
		blockChan:     make(chan *Block, 100),
		pendingBlocks: make(map[uint64]*Block),
		db:            db,
		shutdownChan:  make(chan struct{}),
	}
	bp.pendingCond = sync.NewCond(&bp.mu)
	return bp
}

func (bp *BlockProcessor) Process() {
	go bp.processBlocks()
	go bp.monitorPendingBlocks()
}

// 主处理循环
func (bp *BlockProcessor) processBlocks() {
	for {
		select {
		case block := <-bp.blockChan:
			bp.mu.Lock()

			// 检查是否是我们期待的区块
			if block.Height == bp.lastBlock+1 {
				// 直接处理这个区块
				bp.processBlock(block)
				bp.checkPendingBlocks()
			} else if block.Height > bp.lastBlock+1 {
				// 未来区块，放入等待队列
				bp.pendingBlocks[block.Height] = block
				log.Printf("Added block %d to pending queue", block.Height)
			} else {
				// 旧区块，可以忽略或特殊处理
				log.Printf("Received old block %d, current height is %d",
					block.Height, bp.lastBlock)
			}

			bp.mu.Unlock()

		case <-bp.shutdownChan:
			return
		}
	}
}

// 监控并处理等待队列中的区块
func (bp *BlockProcessor) monitorPendingBlocks() {
	for {
		select {
		case <-time.After(5 * time.Second):
			bp.mu.Lock()
			bp.checkPendingBlocks()
			bp.mu.Unlock()
		case <-bp.shutdownChan:
			return
		}
	}
}

// 检查是否有等待中的区块可以处理
func (bp *BlockProcessor) checkPendingBlocks() {
	for {
		// 检查是否有下一个期待的区块
		nextBlock, exists := bp.pendingBlocks[bp.lastBlock+1]
		if !exists {
			break
		}

		// 处理这个区块
		bp.processBlock(nextBlock)

		// 从等待队列移除
		delete(bp.pendingBlocks, bp.lastBlock)

		// 通知其他可能等待的goroutine
		bp.pendingCond.Broadcast()
	}
}

// 处理单个区块
func (bp *BlockProcessor) processBlock(block *Block) error {
	// 使用事务保证原子性
	tx, err := bp.db.Begin()
	if err != nil {
		return err
	}

	// 写入区块头
	_, err = tx.Exec("INSERT INTO blocks(height, hash, prev_hash) VALUES (?, ?, ?)",
		block.Height, block.Hash, block.PrevHash)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 按顺序写入交易

	if err := tx.Commit(); err != nil {
		return err
	}

	// 更新最后处理的区块高度
	bp.lastBlock = block.Height
	log.Printf("Processed block %d", block.Height)

	return nil
}

// 添加新区块到处理器
func (bp *BlockProcessor) AddBlock(block *Block) {
	select {
	case bp.blockChan <- block:
	case <-time.After(1 * time.Second):
		log.Printf("Block channel full, dropping block %d", block.Height)
	}
}

// 停止处理器
func (bp *BlockProcessor) Stop() {
	close(bp.shutdownChan)
}

func main() {
	db, err := sql.Open("sqlite3", "./chain.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	processor := NewBlockProcessor(db)
	processor.Process()

	// 模拟接收区块(可能是乱序到达)
	for i := 0; i < 4; i++ {
		go func() {

		}()
	}
	processor.AddBlock(&Block{Height: 1, Hash: "hash1", PrevHash: "hash0"})
	processor.AddBlock(&Block{Height: 3, Hash: "hash3", PrevHash: "hash2"})
	processor.AddBlock(&Block{Height: 2, Hash: "hash2", PrevHash: "hash1"})
	processor.AddBlock(&Block{Height: 4, Hash: "hash4", PrevHash: "hash3"})

	time.Sleep(2 * time.Second)
	processor.Stop()
}
