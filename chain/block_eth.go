package chain

import (
	"context"
	"encoding/hex"
	"github.com/anypay/scanner/model"
	"github.com/anypay/scanner/service"
	"github.com/anypay/scanner/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"os"
	"strings"
)

var (
	isEthRunning = false
)

func ScanEthBlock() {
	if isRunning {
		return
	}
	isRunning = true

	//获取当前区块高度
	header, err := utils.EthClient().HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Printf("获取区块高度有误:%v", err)
		return
	}
	log.Println("ETH链最新高度为:", header.Number.Int64())

	startBlock := utils.GetEthBlock()
	for i := startBlock; i <= header.Number.Int64(); i++ {
		getEthBlockInfo(i)
	}
	isRunning = false
}

func getEthBlockInfo(blockNumber int64) {
	log.Println("正在扫描区块：", blockNumber)
	block, err := utils.EthClient().BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		log.Println("获取区块信息出错:", err)
		return
	}

	for _, tx := range block.Transactions() {
		if tx.To() == nil {
			continue
		}
		if strings.ToLower(tx.To().Hex()) == strings.ToLower(os.Getenv("ETH_USDT_CONTRACT")) { //USDT
			coinEventHandler(tx, blockNumber, utils.BlockTime(block.Time()))
		}
	}
	utils.SaveEthBlock(blockNumber)
}

func coinEventHandler(tx *types.Transaction, blockNumber int64, blockTime string) {
	receipt, err := utils.EthClient().TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Println("获取交易凭据出错-", err)
		return
	}

	//遍历事件列表，找到感觉兴趣的
	for _, event := range receipt.Logs {
		if event.Topics[0] == transferEventHash { //合约转账
			transaction := model.Transaction{}
			transaction.BlockTime = blockTime
			transaction.BlockNumber = blockNumber
			transaction.ContractAddress = os.Getenv("ETH_USDT_CONTRACT")
			transaction.ChainId = utils.EthChainId
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
