package blockchain

import (
	"crypto/sha256" // Importing the sha256 package to generate cryptographic hash functions
	"encoding/hex"  // Importing the hex package to encode the hash output to a string
	"strconv"       // Importing the strconv package to convert integers to strings
	"time"          // Importing the time package to work with timestamps
)

// Block represents each 'item' in the blockchain
// It holds the data and the necessary information to ensure the integrity and order of the blocks.
type Block struct {
	Index     int    // Index: The position of the block in the blockchain
	Timestamp string // Timestamp: The time when the block was created, stored as a string
	Data      string // Data: The information or transactions stored in the block
	PrevHash  string // PrevHash: The hash of the previous block in the chain, linking the blocks together
	Hash      string // Hash: The hash of this block, ensuring data integrity and proof of the block's validity
	Nonce     int    // Nonce: A counter used during the proof-of-work algorithm to find a valid hash
}

// CalculateHash generates the hash of the block using its contents
// The hash serves as a unique identifier for the block and ensures the data's integrity.
func (b *Block) CalculateHash() string {
	// Concatenate all the block's fields into a single string
	// This string is used as the input for generating the hash
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevHash + strconv.Itoa(b.Nonce)

	// Create a new SHA-256 hash object
	h := sha256.New()

	// Write the concatenated string into the hash object
	h.Write([]byte(record))

	// Compute the final hash value by summing (finalizing) the hash
	// The result is a byte slice that we convert to a hexadecimal string
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// NewBlock creates a new block using the given data, the hash of the previous block, and the index of the block
// This function is called when a new block needs to be added to the blockchain.
func NewBlock(data, prevHash string, index int) *Block {
	// Create a new block with the provided data, previous block's hash, and index
	// The current timestamp is captured and the nonce is initialized to 0
	block := &Block{
		Index:     index,
		Timestamp: time.Now().String(), // Get the current time and convert it to a string
		Data:      data,
		PrevHash:  prevHash,
		Hash:      "", // The hash will be calculated and set below
		Nonce:     0,  // Nonce starts at 0 and will be used later in proof-of-work
	}

	// Calculate the block's hash based on its initial contents
	block.Hash = block.CalculateHash()
	return block
}

// NewGenesisBlock creates the first block in the blockchain, known as the genesis block
// The genesis block doesn't have a previous hash, so it's typically set to an empty string.
func NewGenesisBlock() *Block {
	// Call NewBlock with a special message "Genesis Block" and no previous hash
	// The index is set to 0 since this is the first block in the chain
	return NewBlock("Genesis Block", "Ram", 0)
}
