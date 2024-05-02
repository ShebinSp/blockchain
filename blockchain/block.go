package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

func CreateBlock(header, body string) {
	fmt.Println(header, "\n", body) // Show the block content
}

func (b *Block) SetHash() {

	// time.Time formats to string and then converted to bytes
	timestamp := []byte(fmt.Sprintf("%f", b.Timestamp)) // []byte(strconv.FormatInt(block.Timestamp, 10)) 

	// get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{timestamp, b.PreviousBlockHash, b.AllData}, []byte{}) // concatenate all the block data

	// hash the headers
	hash := sha256.Sum256(headers)

	// set the hash of the block
	b.CurrentBlockHash = hash[:]
}

// Create a function for new block generation and return that block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		float64(time.Now().Unix()),
		prevBlockHash,
		[]byte{},
		[]byte(data),
	} // the block is received

	// the block is hashed
	block.SetHash()

	return block
}

//create genesis block funtion that will return the first block. This is the first block on the chain
func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
