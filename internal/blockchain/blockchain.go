package blockchain

// Blockchain represents a series of linked blocks, forming the blockchain.
// Each block in the chain is connected to the previous one via its hash.
type Blockchain struct {
	blocks []*Block // A slice (dynamic array) to hold pointers to the blocks in the blockchain.
}

// NewBlockchain creates a new Blockchain instance with a genesis block.
// The genesis block is the first block in any blockchain, and it does not have a predecessor.
func NewBlockchain() *Blockchain {
	return &Blockchain{
		[]*Block{ // Initialize the blockchain with a slice containing just the genesis block.
			NewGenesisBlock(), // Call the function to create the genesis block.
		},
	}
}

// AddBlock adds a new block to the blockchain with the provided data.
// This function creates a new block, links it to the previous block, mines it, and then appends it to the chain.
func (bc *Blockchain) AddBlock(data string) {
	// Get the previous (latest) block in the blockchain.
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// Create a new block, passing the provided data, the hash of the previous block,
	// and the index which is the previous block's index incremented by 1.
	newBlock := NewBlock(data, prevBlock.Hash, prevBlock.Index+1)

	// Mine the new block to find a valid hash according to the proof-of-work algorithm.
	newBlock.MineBlock()

	// Append the newly mined block to the blockchain.
	bc.blocks = append(bc.blocks, newBlock)
}

// Blocks returns the slice of blocks that make up the blockchain.
// This method is useful for accessing the entire chain of blocks.
func (bc *Blockchain) Blocks() []*Block {
	return bc.blocks
}
