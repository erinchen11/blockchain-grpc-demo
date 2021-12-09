package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/erinchen11/blockchain-grpc-demo/proto"
	"google.golang.org/grpc"
)

var client proto.BlockchainClient

const address = "localhost:50051"

func main() {
	// use command line to interact with grpc server
	addFlag := flag.Bool("add", false, "Add a new block")
	listFlag := flag.Bool("list", false, "List all blocks")
	flag.Parse()

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot dial server: %v", err)
	}

	client = proto.NewBlockchainClient(conn)
	// different cmd to call the corresponding function
	if *addFlag {
		addBlock()
	}

	if *listFlag {
		getBlockchain()
	}
}

func addBlock() {
	
	// use current time as input data
	block, err := client.AddBlock(context.Background(), &proto.AddBlockRequest{
		Data: time.Now().String(),
	})
	if err != nil {
		log.Fatalf("can not add a block: %v", err)
	}
	log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockchain() {
	blockchain, err := client.GetBlocks(context.Background(), &proto.GetBlocksRequest{})
	if err != nil {
		log.Fatalf("can not get blockchain: %v", err)
	}

	log.Println("blocks on blockchain:")
	// get all blocks on blokchain
	for _, b := range blockchain.Blocks {
		log.Printf("hash %s, prev hash: %s, data: %s\n", b.Hash, b.PrevBlockHash, b.Data)
	}
}
