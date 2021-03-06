package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
)

func main() {
    client, err := ethclient.Dial("https://mainnet.infura.io/v3/47068b6932a94a3da3e2e5ffd427641a")
    if err != nil {
        log.Fatal(err)
    }

    account := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil 	{
  	    log.Fatal(err)
    }

   fmt.Println(balance)

   balance, err = client.BalanceAt(context.Background(), account, blockNumber)
   if err != nil {
       log.Fatal(err)
   }

   fmt.Println(balance)
}
