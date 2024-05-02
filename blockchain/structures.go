package blockchain

import "fmt"

type Data struct {
	Amount   float32
	Sender   string
	Receiver string
}

// Create the block data structure
// A block contains this info:
type Block struct {
	Timestamp         float64 // the time when the block was created
	Amount            float32
	Nounce            int // Proof-of-Work
	Sender            string
	Receiver          string
	PreviousBlockHash []byte // the hash of the previous block
	CurrentBlockHash  []byte // the hash of the current block
}

// Prepare the Blockchain data structure:
type Blockchain struct {
	Blocks []*Block // a blockchain is a series of blocks
}

func (bb *Blockchain) LengthOf() {
	fmt.Println("length of block: ", len(bb.Blocks))
}
