package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

/**
 * Get RPC URL from environment variable
 */

func getRpcUrl() string {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	rpcUrl := os.Getenv("RPC_URL")
	fmt.Println("RPC_URL: v%", rpcUrl)
	return rpcUrl
}

func getBlockchainClient() *ethclient.Client {
	client, err := ethclient.Dial(getRpcUrl())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func getLastBlock() *big.Int {
	client := getBlockchainClient()
	blockNumber, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return blockNumber.Number
}

func main() {

	fmt.Println(getLastBlock())
}
