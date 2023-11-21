package main

import (
	"context"
	"github.com/anypay/scanner/service"
	"os"
	"os/signal"
)

func main() {

	//加载配置文件
	service.LoadEnvFile()

	_, cancel := context.WithCancel(context.Background())

	//chain.ScanEthBlock()
	//chain.ScanBscBlock()
	//chain.ScanTronBlock()
	//chain.GetTronAccountRank()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	cancel()
}
