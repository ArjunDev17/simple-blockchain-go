package main

import (
	"fmt" // Importing the "fmt" package for formatted I/O, like printing to the console.

	"github.com/ArjunDev17/simple-blockchain-go/internal/blockchain" // Importing the "blockchain" package from the internal directory of the project.
)

func main() {
	// Create a new blockchain instance using the NewBlockchain function from the blockchain package.
	// This initializes a blockchain with a genesis block.
	bc := blockchain.NewBlockchain()

	// Add a new block to the blockchain with the data "Block 1 Data".
	// The blockchain automatically handles linking this block to the previous block.
	bc.AddBlock("Block 1 Data")

	// Add another block to the blockchain with the data "Block 2 Data".
	// This block is also linked to the previous block in the chain.
	bc.AddBlock("Block 2 Data")
	bc.AddBlock("Block 3 Data")

	// Loop through each block in the blockchain to display its details.
	for _, block := range bc.Blocks() {
		// Print the index of the block, indicating its position in the blockchain.
		fmt.Printf("Index: %d\n", block.Index)

		// Print the timestamp of when the block was created.
		fmt.Printf("Timestamp: %s\n", block.Timestamp)

		// Print the data stored in the block, which could represent transactions or other information.
		fmt.Printf("Data: %s\n", block.Data)

		// Print the hash of the previous block, which links the current block to the previous one.
		// This ensures the chain's integrity by making it difficult to alter previous blocks.
		fmt.Printf("PrevHash: %s\n", block.PrevHash)

		// Print the hash of the current block. This hash is unique to the block's contents (Index, Timestamp, Data, PrevHash, Nonce).
		// Any change in the block's data will result in a completely different hash, highlighting the integrity of the block.
		fmt.Printf("Hash: %s\n", block.Hash)

		// Print the nonce used to generate the block's hash during the mining process (proof of work).
		// The nonce is the number that miners adjust to find a hash that meets the blockchain's difficulty requirement.
		fmt.Printf("Nonce: %d\n", block.Nonce)

		// Print a newline for better readability between block details.
		fmt.Println()
	}
}
