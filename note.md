## Methods for gRPC Server

- AddBlock
- GetBlock

## Thinking Flow

### Server and Client

Build a server with gRPC, implement some functionality of the server.
Afer Build an server, must Listen on it, to see what happen to the server.
like client send request to server, the server will get the request
and handle the request, then server will response the corresponding operation to client.

### Blockchain and Block

motivation: for implementing blockchain grpc server.

thinking : <br>
what kind of data struct of Block and Blockchain <br>
what kind of functionality of Block and Blockchain <br>

data struct: 
- block : hash of previous block, data which store in block, hash of current block
- blockchain : slice of block

functionality:
- block : create genesis block, create a new block
- blockchain : use genesis block to create a blockchain, add a new block to blockchain.

the functionality of blockchain grpc server will be implemented by above.

