Error - 1
```
missing method mustEmbedUnimplementedBlockchainServer
```
solution:
in blockchain_grpc.pb.go

```
type BlockchainServer interface {
	AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error)
	GetBlocks(context.Context, *GetBlocksRequest) (*GetBlocksResponse, error)
	mustEmbedUnimplementedBlockchainServer()
}
```
missing 'mustEmbedUnimplementedBlockchainServer()'
and this is function of type UnimplementedBlockchainServer struct { }

cause: gRPC new version, UnimplementedXXXXXServer, [issue](https://github.com/grpc/grpc-go/issues/3669)

so correct type struct of BServer is
```
type BServer struct{ 
	proto.UnimplementedBlockchainServer
 }
```

Error - 2 in the server/main.go, AddBlock function

Error - 3

server seems work, but when client use cmd, the error will come out. <br>

error on client:
```
can not get blockchain: rpc error: code = Unavailable desc = error reading from server: EOF exit status 1
```
error on server:
```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x13d8326]

goroutine 18 [running]:
main.(*BlockchainServer).GetBlocks(0xc000021940, 0x1542960, 0xc00009e1e0, 0xc00009e210, 0xc000021940, 0xc00009e1e0, 0xc00005bba0)
```

cause:

in server/main.go, didn't create an instance of &BlockchainServer{}

```
proto.RegisterBlockchainServer(grpcServer, &BlockchainServer{})
```

solution:

```
proto.RegisterBlockchainServer(grpcServer, &BlockchainServer{
		Blockchain: bl.NewBlockchain(),
	})
```




