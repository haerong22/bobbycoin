package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/haerong22/bobbycoin/utils"
)

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Height       int    `json:"height"`
	Diffilculty  int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Diffilculty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func persistBlock(b *Block) {
	dbStorage.SaveBlock(b.Hash, utils.ToBytes(b))
}

func FindBlock(hash string) (*Block, error) {
	blockbytes := dbStorage.FindBlock(hash)
	if blockbytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockbytes)
	return block, nil
}

func createBlock(prevHash string, height, diff int) *Block {
	block := &Block{
		Hash:        "",
		PrevHash:    prevHash,
		Height:      height,
		Diffilculty: diff,
		Nonce:       0,
	}
	block.Transactions = Mempool().txToConfirm()
	block.mine()
	persistBlock(block)
	return block
}
