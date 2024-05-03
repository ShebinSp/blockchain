package blockchain

// create the method that adds a new block to a blockchain
func (blockchain *Blockchain) AddBlock(data Data) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // the previous block is needed,

	// create a new block containing the data and the hash of the previous block
	newBlock := NewBlock(data, PreviousBlock.CurrentBlockHash)
	// add the newBlock to the chain to create a chain of blocks
	blockchain.Blocks = append(blockchain.Blocks, newBlock)

}

// The function that returns the whole blockchain and add the genesis to it first. The genesis block is the first ever minde
// block

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}

}

// Checking if the blocks are valid except the last index
func (b Blockchain) IntegrityCheck() bool {

	blocks := b.Blocks
	blockLen := len(blocks) - 1

	for i, block := range b.Blocks[:blockLen] {
		previousHash := string(block.CurrentBlockHash)
		currentHash := string(blocks[i+1].PreviousBlockHash)

		// comparing the Previous block hash and Current block hash
		if previousHash != currentHash && i+1 < blockLen {
			return false
		}

	}

	return true
}
