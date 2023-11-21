package utils

import (
	"context"
	"fmt"
	"github.com/anypay/scanner/service"
	"math/big"
	"strconv"
	"time"
)

type ChainType string

const (
	ChainTypeBsc ChainType = "bsc"
	ChainTypeEth ChainType = "eth"
	EthChainId             = 60
	BSCChainId             = 100
	TronChainId            = 200
)

func BlockTime(timestamp uint64) string {
	// 你的Unix时间戳，以秒为单位
	timestampInterval := int64(timestamp)

	// 创建一个Time类型的值，使用Unix函数将时间戳转换为Time
	t := time.Unix(timestampInterval, 0)

	// 设置时区为东八区（中国标准时间）
	cst := time.FixedZone("CST", 8*60*60)

	// 使用In函数将时间转换为指定时区的时间
	cstTime := t.In(cst)
	// 格式化为日期时间字符串
	formattedTime := cstTime.Format("2006-01-02 15:04:05")
	return formattedTime
}

func FormatAmount(data []uint8) string {
	amountInt := new(big.Int)
	amountInt.SetBytes(data)
	amountValue := fmt.Sprintf("%d", amountInt)
	return amountValue
}

func GetUsdtAmount(data []uint8) string {
	if len(data) < 18 {
		return ""
	}
	b := new(big.Int)
	value := b.SetBytes(data[:])
	amount := fmt.Sprintf("%d", value)
	length := len(amount)
	return amount[:length-18]
}

func FormatAddress(address string) string {
	return "0x" + address[26:]
}

func SaveEthBlock(blockNumber int64) {
	service.GetRedis().Set(context.Background(), "eth_block", blockNumber, 0)
}

func SaveBscBlock(blockNumber int64) {
	service.GetRedis().Set(context.Background(), "bsc_block", blockNumber, 0)
}

func SaveTronBlock(blockNumber int64) {
	service.GetRedis().Set(context.Background(), "tron_block", blockNumber, 0)
}

func GetEthBlock() int64 {
	str, err := service.GetRedis().Get(context.Background(), "eth_block").Result()
	if err != nil {
		return 0
	}
	tmp, err := strconv.Atoi(str)
	blockNumber := int64(tmp)
	return blockNumber
}

func GetBscBlock() int64 {
	str, err := service.GetRedis().Get(context.Background(), "bsc_block").Result()
	if err != nil {
		return 0
	}
	tmp, err := strconv.Atoi(str)
	blockNumber := int64(tmp)
	return blockNumber
}

func GetTronBlock() int64 {
	str, err := service.GetRedis().Get(context.Background(), "tron_block").Result()
	if err != nil {
		return 0
	}
	tmp, err := strconv.Atoi(str)
	blockNumber := int64(tmp)
	return blockNumber
}
