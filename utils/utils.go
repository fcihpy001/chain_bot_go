package utils

import (
	"context"
	"fcihpy.com/chainBot/service"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math"
	"math/big"
	"strconv"
	"strings"
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

// EthToWei eth单位安全转wei
// https://stackoverrun.com/cn/q/13021596
func EthToWei(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

func Wei2Eth(balance string) {
	//b := new(big.Int)
	//value := b.SetBytes(data[:])
	//amount := fmt.Sprintf("%d", value)

	fbalance := new(big.Float)
	fbalance.SetString(balance)
	println("ff:", fbalance)
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	println("balance:", ethValue.String())
}

func FormatEth(amount []byte) string {
	//byteArray := []byte{0x01, 0x00, 0x00} // 示例字节数组，你需要根据实际情况替换

	// 将字节数组转换为大整数
	bigInt := new(big.Int).SetBytes(amount)

	// 将大整数除以 10^18，得到以太币单位
	ethAmount := new(big.Float).SetInt(bigInt)
	ethAmount.Quo(ethAmount, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))

	eth := ethAmount.Text('f', 18)
	// 输出以太币单位
	fmt.Printf("ETH Amount1: %s\n", eth)
	return eth
}

// 获取ABI对象
func GetABI(abiJSON string) abi.ABI {
	wrapABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}
	return wrapABI
}
