package blockchain

import (
	"strings"
)

const Difficulty = 2

func (b *Block) MineBlock() {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", Difficulty)) {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
}
