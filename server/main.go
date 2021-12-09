package main

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/erinchen11/blockchain-grpc-demo/proto"
	bl "github.com/erinchen11/blockchain-grpc-demo/server/blockchain"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	// use tcp to listen port 8000

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Can's listen on port 50051: %v", err)
	}

	fmt.Println("Already listen on port 50051")

	// create an instance of grpc server, it will return a type grpc.Server
	// pass it to RegisterBlockchainServer from blockchain_grpc.pb.go
	grpcServer := grpc.NewServer()
	// need grpc.server and BlockchainServer as parameters
	// Error3 - remember to instanitial BlockchainServer
	
	proto.RegisterBlockchainServer(grpcServer, &BlockchainServer{
		Blockchain: bl.NewBlockchain(),
	})
	// check the BlockchainServer, it's an interface
	// have to implement it

	fmt.Println("Start grpc server...")
	// run start the grpcServer, have to pass net.Listener
	// grpcServer.Serve(listener)

	// fmt.Println("Server running success")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Server running success")

}

// implement the BlockchainServer
// AddBlock, GetBlock

// error-1, see debug.md - proto.UnimplementedBlockchainServer
type BlockchainServer struct {
	Blockchain *bl.Blockchain
	proto.UnimplementedBlockchainServer
}

// Server's methods

// Add a new block to blockchain by request
func (s *BlockchainServer) AddBlock(ctx context.Context, request *proto.AddBlockRequest) (*proto.AddBlockResponse, error) {
	//
	block := s.Blockchain.AddBlock(request.Data)

	response := &proto.AddBlockResponse{
		Hash: block.Hash,
	}

	return response, nil

	// error - 2, cannot use &resp (value of type *string) as *proto.AddBlockResponse value in return statement
	// because type of resp is string, and type of AddBlockResponse not only string
	// but can use struct to assign Hash
	// resp := block.Hash
	// return &resp, nil

}

// Get a blockchain ([]*Block)
func (s *BlockchainServer) GetBlocks(ctx context.Context, request *proto.GetBlocksRequest) (*proto.GetBlocksResponse, error) {

	// use Blockchain.Blocks
	// need an instance of GetBlockchainResponse
	// and append Blockchain.Blocks to the instance then return the instance
	response := new(proto.GetBlocksResponse)

	// response's type is Blocks []*Block
	// get the each block of blockchain
	// use the block's information to map proto.Block, and add the block to response
	for _, block := range s.Blockchain.Blocks {
		response.Blocks = append(response.Blocks, &proto.Block{
			PrevBlockHash: block.PrevHash,
			Hash:          block.Hash,
			Data:          block.Data,
		})
	}

	return response, nil
}
