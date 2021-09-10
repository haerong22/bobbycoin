package blockchain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/haerong22/bobbycoin/db"
	"github.com/haerong22/bobbycoin/utils"
)

type Block struct {
	Data        string `json:"data"`
	Hash        string `json:"hash"`
	PrevHash    string `json:"prevHash,omitempty"`
	Height      int    `json:"height"`
	Diffilculty int    `json:"difficulty"`
	Nonce       int    `json:"nonce"`
	Timestamp   int    `json:"timestamp"`
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

func (b *Block) mine() {
	target := strings.Repeat("0", b.Diffilculty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("Target:%s\nHash:%s\nNonce:%d\n\n", target, hash, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:        data,
		Hash:        "",
		PrevHash:    prevHash,
		Height:      height,
		Diffilculty: Blockchain().difficulty(),
		Nonce:       0,
	}
	block.mine()
	block.persist()
	return block
}
