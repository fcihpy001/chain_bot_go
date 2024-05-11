package chain

import (
	"github.com/ethereum/go-ethereum/crypto"
	"time"
)

var (
	//startBlock int64 = 5873817
	isRunning = false

	transferEventHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)

// Task 定时任务
func Task() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				ScanEthBlock()
			}
		}
	}()
	<-make(chan interface{})
}

//func ScanEthBlock() {
//	if isRunning {
//		return
//	}
//	isRunning = true
//
//	//获取当前区块高度
//	header, err := utils.EthClient().HeaderByNumber(context.Background(), nil)
//	if err != nil {
//		log.Printf("获取区块高度有误:%v", err)
//		return
//	}
//	log.Println("ETH链最新高度为:", header.Number.Int64())
//
//	startBlock := utils.GetEthBlock()
//	for i := startBlock; i <= header.Number.Int64(); i++ {
//		getEthBlockInfo(i)
//	}
//	isRunning = false
//}

//
//func ScanBscBlock() {
//	if isRunning {
//		return
//	}
//	isRunning = true
//
//	//获取当前区块高度
//	header, err := utils.BscClient().HeaderByNumber(context.Background(), nil)
//	if err != nil {
//		log.Printf("获取区块高度有误:%v", err)
//		return
//	}
//	log.Println("ETH链最新高度为:", header.Number.Int64())
//
//	startBlock := utils.GetEthBlock()
//	for i := startBlock; i <= header.Number.Int64(); i++ {
//		getEthBlockInfo(i)
//	}
//	isRunning = false
//}
