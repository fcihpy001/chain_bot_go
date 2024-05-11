package chain

import (
	"context"
	"fcihpy.com/chainBot/utils"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

func query_contract_event() {
	//合约地址
	contractAddress := common.HexToAddress("0x6deFd7e108708019E99fF84e6B2731D27d39be68")
	//websocket监听
	client, err := ethclient.Dial(os.Getenv("SEPOLIA_CHAIN_WSS"))
	if err != nil {
		fmt.Errorf("could not connect to local node: %v", err)
		return
	}
	fmt.Println("eth success")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(5879852), //生产环境中，从0开始，查询后修改区块记录，下一次就从后一个有记录的区块数开始
		ToBlock:   big.NewInt(5879853),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	//erc20, _ := json.NewErc20(contractAddress, client)
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Errorf("err:%v\n", err)
		return
	}
	for _, vLog := range logs {
		if len(vLog.Topics) == 0 {
			continue
		}
		event := vLog.Topics[0].Hex()
		if event == TransferEvent() { //对对应的事件进行对应的处理
			//fmt.Println(vLog.Data)
			fmt.Println("block_no", vLog.TxHash, utils.FormatAddress(vLog.Topics[1].Hex()), utils.FormatEth(vLog.Data))

			//data, err := erc20.ParseTransfer(vLog)
			//if err != nil {
			//	fmt.Errorf("err:%v\n", err)
			//	continue
			//}
			//fmt.Println(data.From.Hex(), data.To.Hex(), data.Value.Int64(), data.Raw.Data)
		}
	}
}

func TransferEvent() string {
	event := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")).Hex()
	return event
}

func monitor_contract_event() {
	//合约地址
	contractAddress := common.HexToAddress("0x6deFd7e108708019E99fF84e6B2731D27d39be68")
	//websocket监听
	client, err := ethclient.Dial(os.Getenv("SEPOLIA_CHAIN_WSS"))
	if err != nil {
		fmt.Errorf("could not connect to local node: %v", err)
		return
	}
	defer client.Close()

	fmt.Println("eth success")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {

			case TransferEvent():
				fmt.Println("transfer_info", vLog.TxHash, utils.FormatAddress(vLog.Topics[1].Hex()), utils.FormatEth(vLog.Data))

			}
			//buyEvent事件 这里的hash为事件名buyEvent(address,uint64)进行keccack256计算得出的
			//case "0xbf5ed60d1b60f93841065247988234acab61d25f0cdb23dae3507eb01809e42f":
			//	//这步是对监听到的DATA数据进行解析
			//	intr, err := abi.Events["buyEvent"].Inputs.UnpackValues(vLog.Data)
			//	if err != nil {
			//		panic(err)
			//	}
			//	//打印监听到的参数
			//	fmt.Println(intr[0].(common.Address).String(),intr[1].(*big.Int).Uint64())
			//	fmt.Println("交易hash",vLog.TxHash.String())
			//	//exerciseEvent事件 exerciseEvent(uint64,uint8,uint40,uint32,uint256)
			//case "0xc32df39ae9f5cc595be756fd57aa0a8976bf1c8a5aad5199dc2001efa4a384b5":
			//	//这步是对监听到的DATA数据进行解析
			//	intr, err := abi.Events["exerciseEvent"].Inputs.UnpackValues(vLog.Data)
			//	if err != nil {
			//		panic(err)
			//	}
			//	list:=[]interface{}{}
			//	for _,v:=range intr{
			//		list=append(list,v)
			//	}
			//	//打印监听到的参数
			//	fmt.Println(list)
			//	fmt.Println("交易hash",vLog.TxHash.String())
			//}
		}
	}
}
