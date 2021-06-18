package main

import (
	"context"
	// "math/big"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// func main() {
// 	client, err := ethclient.Dial("0.0.0.0")
// 	ctx := context.Background()
// 	block, err := client.BlockByNumber(ctx, big.NewInt(123))
// 	if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Println("we have a connection") // we'll use this in the upcoming sections
// 	fmt.Println("blocknumber",block)
// }
func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/4d377fafbdaf48198ed4770145ff9fa1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
	account := common.HexToAddress("0x2FF312a35e4bBD2158997D90CC81bC298059F2D1")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 25893180161173005034
}
