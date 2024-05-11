package chain

import (
	"context"
	"encoding/hex"
	"fcihpy.com/chainBot/model"
	"fcihpy.com/chainBot/service"
	"fcihpy.com/chainBot/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"os"
	"strings"
)

var (
	isBscRunning = false
)

func ScanBscBlock() {
	if isBscRunning {
		return
	}
	isBscRunning = true

	//获取当前区块高度
	header, err := utils.BscClient().HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("获取区块高度有误:%v", err)
		return
	}
	log.Println("BSC链最新高度为:", header.Number.Int64())

	startBlock := utils.GetBscBlock()
	for i := startBlock; i <= header.Number.Int64(); i++ {
		getBSCBlockInfo(i)
	}
	isBscRunning = false
}

func getBSCBlockInfo(blockNumber int64) {
	log.Println("正在扫描区块：", blockNumber)
	block, err := utils.BscClient().BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Println("获取区块信息出错:", err)
		return
	}

	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue
		}
		if strings.ToLower(tx.To().Hex()) == strings.ToLower(os.Getenv("BSC_USDT_CONTRACT")) { //USDT

			bscEventHandler(tx, blockNumber, utils.BlockTime(block.Time()))
		}
	}
	utils.SaveBscBlock(blockNumber)
}

func bscEventHandler(tx *types.Transaction, blockNumber int64, blockTime string) {
	receipt, err := utils.BscClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Println("获取交易凭据出错-", err)
		return
	}
	fmt.Println("tx", receipt.TxHash)
	fmt.Println("addr", receipt.ContractAddress)
	//遍历事件列表，找到感觉兴趣的
	for _, event := range receipt.Logs {
		if event.Topics[0] == transferEventHash { //合约转账
			transaction := model.Transaction{}
			transaction.BlockTime = blockTime
			transaction.BlockNumber = blockNumber
			transaction.ContractAddress = os.Getenv("BSC_USDT_CONTRACT")
			transaction.ChainId = utils.BSCChainId
			transaction.CoinName = "USDT"
			transaction.From = utils.FormatAddress(event.Topics[1].Hex())
			transaction.To = utils.FormatAddress(event.Topics[2].Hex())
			transaction.AmountHex = strings.TrimLeft(hex.EncodeToString(event.Data), "0")
			transaction.Amount = utils.FormatAmount(event.Data[:])
			transaction.TxHash = event.TxHash.Hex()
			transaction.TxIndex = event.Index
			service.GetDB().Save(&transaction)
		}
	}
}
