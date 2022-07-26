package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block ...
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// SetHash calc hash and set
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock ...
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()

	return block
}

// Blockchain ...
type Blockchain struct {
	blocks []*Block
}

// AddBlock add new block to blockchain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// To add a new block we need an existing block, but there's not blocks in our blockchain!
// So, in any blockchain, there must be at least one block

// NewGenesisBlock ...
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockChain ...
func NewBlockChain() *Blockchain {
	return &Blockchain{
		blocks: []*Block{NewGenesisBlock()},
	}
}
