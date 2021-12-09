package blockchain

import (
	"crypto/sha256"
	"encoding/hex"

)

// define block

type Block struct {
	Hash string
	PrevHash string
	Data string
}

type Blockchain struct {
	Blocks []*Block
}

// method of Block
// Can use Block to set the Hash of its own.
// genrate a new block must set up its own Hash
func (b *Block) setHash()  {
	// new hash is generated from PrevHash + Data of current block
	newHash := sha256.Sum256([]byte(b.PrevHash + b.Data))
	//set up its own Hash, it's string, use hex to conver []byte to string
	b.Hash = hex.EncodeToString(newHash[:])

}

// pass data and prevhash and return a new block
func NewBlock(data string, prevHash string) *Block{
	// implement instance of Block
	block := &Block{
		Data: data,
		PrevHash: prevHash,
	}
	// set up hash of new block
	block.setHash()

	return block

} 

// Blockchain must have genesis block
// use NewBlock() to generate
func GenesisBlock() *Block{
	// No prevhash
	return NewBlock("Genesis Block", "")
}

// use genesis block to create a blockchain
// return *Blockchain
func NewBlockchain() *Blockchain{
	// think of the type of Blockchain is []*Block
	// and GenesisBlock() will return *Block
	return &Blockchain{[]*Block{GenesisBlock()}}
}

// Add new Block to Blockchain
// use new data as a Block to add to Blockchain and return the last Block
func (bc *Blockchain) AddBlock(data string) *Block {
	// because Blockchain's type is  slice
	// think: add new element to the end of slice
	lastBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := NewBlock(data, lastBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock

	// new idea: if blockchain's type is linklist
	// use property of linklist- pointer to add new block to blockchain
	// update the pointer of last block
	// update the pointer of new block (it become the last one)
}


