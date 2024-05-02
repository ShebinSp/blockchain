package main

import (
	"fmt"

	"github.com/ShebinSp/blockchain_project/blockchain"
)

func main() {

	data1 := blockchain.Data{
		Amount:   11.11,
		Sender:   "2N3oefVeg6stiTb5Kh3ozCSkaqmx91FDbsm",
		Receiver: "0x32Be343B94f860124dC4fEe278FDCBD38C102D88",
	}
	data2 := blockchain.Data{
		Amount:   2,
		Sender:   "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		Receiver: "0xF79D6aFBb6dA890132F9D7c355e3015f15F3406F",
	}
	data3 := blockchain.Data{
		Amount:   90.3,
		Sender:   "0x617F2E2fD72FD9D5503197092aC168c91465E7f2",
		Receiver: "0x17F6AD8Ef982297579C203069C1DbfFE4348c372",
	}
	newblockchain := blockchain.NewBlockchain() // Initialize the blockchain
	// create 2 blocks and add 2 transactions
	newblockchain.AddBlock(data1) // first block containing one tx
	newblockchain.AddBlock(data2) // second block containing one tx
	newblockchain.AddBlock(data3)
	// Now print all the blocks and their contents
	for i, block := range newblockchain.Blocks { // iterate on each block
		fmt.Printf("Block ID : %d \n", i)                                        // print the block ID
		fmt.Println("Sender address: ", block.Sender)                            // wallet address of the sender
		fmt.Println("Receiver address: ", block.Receiver)                        // wallet address ot the receiver
		fmt.Println("Transfered amount: ", block.Amount)                         // transferd amount
		fmt.Printf("Timestamp : %f \n", block.Timestamp+float64(i))              // print the timestamp of the block, to make them different, we just add a value i
		fmt.Printf("Hash of the block : %s\n", block.CurrentBlockHash)           // print the hash of the block
		fmt.Printf("Hash of the previous Block : %s\n", block.PreviousBlockHash) // print the hash of the previous block
		fmt.Println("Nounce of the block: ", block.Nounce)
	} // our blockchain will be printed

	//fmt.Println("Length: ",blockchain.LengthOf())
}
