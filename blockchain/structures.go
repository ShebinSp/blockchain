package blockchain

// Create the block data structure
// A block contains this info:
type Block struct {
	Timestamp         float64 // the time when the block was created
	PreviousBlockHash []byte  // the hash of the previous block
	CurrentBlockHash  []byte  // the hash of the current block
	AllData           []byte  // the data or transactions (body info)
}

// Prepare the Blockchain data structure:
type Blockchain struct {
	Blocks []*Block // a blockchain is a series of blocks
}
