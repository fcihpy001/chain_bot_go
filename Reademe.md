# scanner 程序使用说明

- 参数1 chain_type
  - bsc
  - tron
  - eth


# 热编译步骤
. 安装air 插件,使用黑
```shell
go get -u github.com/cosmtrek/air
```

. 在项目当前目录执行air命令，会自动监控文件的变化
```shell
air
```

# 对于合约监听有三种方法
## 1.直接扫描链上区块

## 2.监听合约事件
   - 配置WSS rpc
```shell
client, err := ethclient.Dial(os.Getenv("SEPOLIA_CHAIN_WSS"))
```
   - 配置过滤条件，设置合约地址即可
```shell
query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}
```
   - 订阅事件，用channerl 接收
```shell
sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
```
   - select 轮询接收事件
```shell
for {
    select {
    case err := <-sub.Err():
        log.Fatal(err)
    case vLog := <-logs:
        switch vLog.Topics[0].Hex() {

        case TransferEvent():
            fmt.Println("transfer_info", vLog.TxHash, utils.FormatAddress(vLog.Topics[1].Hex()), utils.FormatEth(vLog.Data))

        }
    }
}
```
### 3.在历史记录中查询合约事件
  - 配置WSS rpc
```shell
client, err := ethclient.Dial(os.Getenv("SEPOLIA_CHAIN_WSS"))
```
- 配置过滤条件，设置合约地址，起始、终止的区块号
```shell
query := ethereum.FilterQuery{
    Addresses: []common.Address{
        contractAddress,
    },
}
```
  - 开始筛选
```shell
logs, err := client.FilterLogs(context.Background(), query)
```
  - 数据解析
```shell
for _, vLog := range logs {
    if len(vLog.Topics) == 0 {
        continue
    }
    event := vLog.Topics[0].Hex()
    if event == TransferEvent() { //对对应的事件进行对应的处理
        fmt.Println("block_no", vLog.TxHash, utils.FormatAddress(vLog.Topics[1].Hex()), utils.FormatEth(vLog.Data))
    }
}
```



# 对于数据的解析，可以手动处理，也可以通过ABI进行解析


### 使用代码发送一笔交易

### 安装ABIgen
```shell
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```


