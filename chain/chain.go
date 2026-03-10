package chain

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/poisonlocket/stutz-coin/models"
	"time"
)

func calculateHash(block models.Block) string {
	record := string(rune(block.Index)) + block.Timestamp + block.Content + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock models.Block, Content string) (models.Block, error) {

	var newBlock models.Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Content = Content
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock, nil
}

func ValidateBlock(newBlock, oldBlock models.Block) bool {
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

func ReplaceChain(newChain, currChain []models.Block) []models.Block {
	if len(currChain) < len(newChain) {
		return newChain
	}
	return currChain
}
