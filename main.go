package main

import (
	"context"
	"log"
	"os"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethGotypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type BlockData struct {
	Number       *big.Int
	Hash         string
	ParentHash   string
	Nonce        uint64
	Sha3Uncles   string
	Miner        string
	Difficulty   *big.Int
	ExtraData    string
	Size         uint64
	GasLimit     uint64
	GasUsed      uint64
	Timestamp    uint64
	Transactions []string
}

/**
 * Get RPC URL from environment variable
 */

func getRpcUrl() string {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err, `for getRpcUrl`)
	}

	rpcUrl := os.Getenv("RPC_URL")
	return rpcUrl
}

/**
 * Get blockchain client
 */

func getBlockchainClient() *ethclient.Client {
	client, err := ethclient.Dial(getRpcUrl())
	if err != nil {
		log.Fatal(err, `for getBlockchainClient`)
	}
	return client
}

/**
 * Get last block number
 */

func getLastBlock() *big.Int {
	client := getBlockchainClient()
	blockNumber, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err, `for getLastBlock`)
	}
	return blockNumber.Number
}

/**
 * Get block data
 */

// func getBlockData() map[string]interface{} {
// 	lastBlock := getLastBlock()
// 	client := getBlockchainClient()
// 	block, err := client.BlockByNumber(context.Background(), lastBlock)
// 	if err != nil {
// 		log.Fatal(err, `for getBlockData`)
// 	}

// 	return map[string]interface{}{
// 		"Number":       block.Number(),
// 		"Hash":         block.Hash().Hex(),
// 		"ParentHash":   block.ParentHash().Hex(),
// 		"Nonce":        block.Nonce(),
// 		"Sha3Uncles":   block.UncleHash().Hex(),
// 		"Miner":        block.Coinbase().Hex(),
// 		"Difficulty":   block.Difficulty(),
// 		"ExtraData":    string(block.Extra()),
// 		"Size":         block.Size(),
// 		"GasLimit":     block.GasLimit(),
// 		"GasUsed":      block.GasUsed(),
// 		"Timestamp":    block.Time(),
// 		"Transactions": getTransactions(block),
// 	}
// }

func getTransactionData(tx string) map[string]interface{} {
	client := getBlockchainClient()
	txHash := common.HexToHash(tx)
	transaction, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err, ` for finding the transaction`)
	}

	if !isPending {
		log.Fatal(`tx pending`)
	}

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err, `for chainId`)
	}

	sender, err := ethGotypes.Sender(ethGotypes.NewLondonSigner(chainId), transaction)
	if err != nil {
		log.Fatal(err, `for gettig the sender`)
	}

	println(string(transaction.Data()), sender.Hex())

	return map[string]interface{}{
		"Hash":     transaction.Hash().Hex(),
		"Value":    transaction.Value().String(),
		"Gas":      transaction.Gas(),
		"GasPrice": transaction.GasPrice().String(),
		"Nonce":    transaction.Nonce(),
		"Data":     string(transaction.Data()),
		"From":     sender.Hex(),
		"To":       transaction.To().Hex(),
	}
}

// func getTransactions(block *ethGotypes.Block) []string {
// 	var transactions []string
// 	for _, tx := range block.Transactions() {
// 		transactions = append(transactions, tx.Hash().Hex())
// 	}
// 	return transactions
// }

func main() {
	println(getTransactionData("0xead0b1c9b70fab3656bbce96c0f052134790859e1bccdf4b5f27556baff1ab37"))
}
