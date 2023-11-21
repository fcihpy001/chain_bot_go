package model

import "gorm.io/gorm"

type Config struct {
	Mysql    Mysql
	Redis    Redis
	Version  string
	Dsn      string
	RedisURL string
	Server   Server
}

type Server struct {
	Host string
	Port string
}

type Mysql struct {
	Host     string
	Port     string
	Database string
	UserName string
	Password string
	Charset  string
	Loc      string
}
type Redis struct {
	Host string
	Port string
}

type Transaction struct {
	gorm.Model
	BlockTime       string
	BlockNumber     int64
	ContractAddress string
	CoinName        string
	ChainId         uint
	TxHash          string `gorm:"tx_hash;uniqueIndex:tx_index; type:char(66);"`
	TxIndex         uint   `gorm:"tx_index;uniqueIndex:tx_index"`
	From            string
	To              string
	Amount          string
	AmountHex       string
	Status          bool `gorm:"default:1"`
	Remark          string
}

type Tron struct {
	Block           int64  `json:"block"`
	BlockTs         int64  `json:"block_ts"`
	ContractAddress string `json:"contract_address"`
	TransactionId   string `json:"transaction_id"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	Quant           string `json:"quant"`
	EventType       string `json:"event_type"`
	Confirmed       bool   `json:"confirmed"`
}

type ApiResponse struct {
	Total          int           `json:"total"`
	TokenTransfers []Tron        `json:"token_transfers"`
	Data           []AccountData `json:"data"`
	// 可以添加其他如状态码、消息等字段
}

type AccountData struct {
	Address string
	Balance int64
}
