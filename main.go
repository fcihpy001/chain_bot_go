package main

import (
	"fcihpy.com/chainBot/service"
)

func main() {

	//wei := utils.EthToWei(2.8)
	//println("afa", wei)
	//
	//utils.Wei2Eth("133330241232")
	//return
	//加载配置文件
	service.GetConfig()

	//query_contract_event()
	//_, cancel := context.WithCancel(context.Background())
	//
	//chain.ScanEthBlock()
	////chain.ScanBscBlock()
	////chain.ScanTronBlock()
	////chain.GetTronAccountRank()
	//
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//<-c
	//cancel()
}
