package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/poisonlocket/stutz-coin/models"
	"time"
)

func calculateHash(block models.Block) string {
	record := string(rune(block.Index)) + block.Timestamp + string(rune(block.BPM)) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock models.Block, BPM int) models.Block {
	var newBlock models.Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

func validateBlock(newBlock, oldBlock models.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func replaceChain(newChain, currChain []models.Block) []models.Block {
	if len(currChain) < len(newChain) {
		return newChain
	}
	return currChain
}
