package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func CreateBlock(header, body string) {
	fmt.Println(header, "\n", body) // Show the block content
}

const targetDifficulty = 4 // Adjust the difficulty target as needed


func (b *Block) SetHash() {
	// Convert the target difficulty to a string with leading zeroes
	targetString := fmt.Sprintf("%0*d", targetDifficulty, 0)

	nonce := 0
	var hashString string

	for {
		timestamp := []byte(strconv.FormatFloat(b.Timestamp, 'f', -1, 64)) // Convert timestamp to bytes
		nonceBytes := []byte(strconv.Itoa(nonce))                          // Convert nonce to bytes

		headers := bytes.Join([][]byte{timestamp, b.PreviousBlockHash, nonceBytes}, []byte{}) // Concatenate block data
		hash := sha256.Sum256(headers)                                                        // Hash the block data

		hashString = fmt.Sprintf("%x", hash) // Convert hash to hexadecimal string

		// Check if the hash meets the target difficulty
		if hashString[:targetDifficulty] == targetString {
			break
		}
		nonce++
	}

	b.CurrentBlockHash = []byte(hashString) // Set the hash of the block
	b.Nounce = nonce                           // Set the proof-of-work value
}

// Create a function for new block generation and return that block
func NewBlock(data Data, prevBlockHash []byte) *Block {
	block := &Block{
		float64(time.Now().Unix()),
		data.Amount,
		0,
		data.Sender,
		data.Receiver,
		prevBlockHash,
		[]byte{},
	} // the block is received

	// the block is hashed
	block.SetHash()

	return block
}

// create genesis block funtion that will return the first block. This is the first block on the chain
func GenesisBlock() *Block {
	return NewBlock(Data{Receiver: "Genesis Block"}, []byte{})
}
