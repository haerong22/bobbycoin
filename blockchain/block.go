package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/haerong22/bobbycoin/db"
	"github.com/haerong22/bobbycoin/utils"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockbytes := db.Block(hash)
	if blockbytes == nil {
		fmt.Printf("NewestHash: %s\nHeight:%d\n", b.NewestHash, b.Height)
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockbytes)
	return block, nil
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevHash,
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}
