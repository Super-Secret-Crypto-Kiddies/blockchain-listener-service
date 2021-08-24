package eth

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Listen(address string) {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/8aee8b73fc1e4c3a9bc776e3b4fe6804")

	if err != nil {
		panic(err)
	}

	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)

	if err != nil {
		panic(err)
	}

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())

			if err != nil {
				panic(err)
			}

			for _, transaction := range block.Transactions() {
				if transaction.To() != nil {
					if strings.ToLower(transaction.To().String()) == "0xd6f35edeb6b60a73dfb3eecfa0f7eca30d4c0957" {
						fmt.Println(transaction)
					}
				}
			}
		}
	}

}
