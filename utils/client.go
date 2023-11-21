package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

var (
	ethClient *ethclient.Client
	bscClient *ethclient.Client
)

func EthClient() *ethclient.Client {
	if ethClient == nil {

		log.Println("env 文件读取功")
		dial, err := ethclient.Dial(os.Getenv("ETH_CHAIN_RPC"))
		if err != nil {
			log.Fatalf("连接以太坊节点失败：%v", err)
		}
		ethClient = dial
		log.Println("eth_client节点初始化成功...")
	}
	return ethClient
}

func BscClient() *ethclient.Client {
	if bscClient == nil {

		log.Println("env 文件读取功")
		dial, err := ethclient.Dial(os.Getenv("BSC_CHAIN_RPC"))
		if err != nil {
			log.Fatalf("连接币安节点失败：%v", err)
		}
		bscClient = dial
		fmt.Println("币安链节点初始化成功。。。")
	}
	return bscClient
}

func TronClient() {

}
