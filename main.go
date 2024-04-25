package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func getRpcUrl() string {
	rpcUrl := os.Getenv("RPC_URL")
	fmt.Println("RPC_URL: v%", rpcUrl)
	return rpcUrl
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(getRpcUrl())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client.BlockNumber())
}
