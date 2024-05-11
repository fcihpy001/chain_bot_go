package chain

import (
	"encoding/json"
	"fcihpy.com/chainBot/model"
	"fcihpy.com/chainBot/service"
	"fcihpy.com/chainBot/utils"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	isTronRunning = false
)

func ScanTronBlock() {
	for {
		fmt.Println("开始执行任务..")
		startTronBlock := utils.GetTronBlock()
		url := fmt.Sprintf("https://apilist.tronscan.org/api/token_trc20/transfers?contract_address=TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t&block=%d", startTronBlock)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("TRON-PRO-API-KEY", os.Getenv("TRON_KEY"))

		client := http.Client{}
		resp, _ := client.Do(req)
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		fmt.Println(string(body))

		var respone model.ApiResponse
		err := json.Unmarshal(body, &respone)
		if err != nil || respone.Total == 0 || len(respone.TokenTransfers) == 0 {
			timer := time.AfterFunc(10*time.Second, ScanTronBlock)
			select {
			case <-timer.C:
			}
			fmt.Println("异常退出")
			return
		}

		for _, tron := range respone.TokenTransfers {
			if tron.EventType == "Transfer" && tron.ContractAddress == os.Getenv("TRON_USDT_CONTRACT") {
				transaction := model.Transaction{}
				transaction.BlockTime = utils.BlockTime(uint64(tron.BlockTs))
				transaction.BlockNumber = tron.Block
				transaction.ContractAddress = os.Getenv("TRON_USDT_CONTRACT")
				transaction.ChainId = utils.TronChainId
				transaction.CoinName = "USDT"
				transaction.From = tron.FromAddress
				transaction.To = tron.ToAddress
				transaction.AmountHex = ""
				transaction.Amount = tron.Quant
				transaction.TxHash = tron.TransactionId
				transaction.TxIndex = 1
				service.GetDB().Save(&transaction)
			}
		}
		utils.SaveTronBlock(startTronBlock + 1)
	}
}

// 获取账户余额排名在前1万名的地址
func GetTronAccountRank() {
	page := 1
	offset := (page - 1) * 100
	url := fmt.Sprintf("https://apilist.tronscanapi.com/api/account/list?sort=-balance&limit=100&start=%d", offset)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("TRON-PRO-API-KEY", os.Getenv("TRON_KEY"))

	client := http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	var response model.ApiResponse
	err := json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("account", err)
	}
	for _, acc := range response.Data {
		fmt.Printf("account:%s--balance:%s", acc.Address, acc.Balance)
	}

}
