package db

import (
	"fmt"
	"os"

	"github.com/haerong22/bobbycoin/utils"
	bolt "go.etcd.io/bbolt"
)

const (
	dbName      = "blockchain"
	dataBucket  = "data"
	blockBucket = "blocks"
	checkpoint  = "checkpoint"
)

var db *bolt.DB

type DB struct{}

func (DB) FindBlock(hash string) []byte {
	return findBlock(hash)
}
func (DB) SaveBlock(hash string, data []byte) {
	saveBlock(hash, data)
}
func (DB) SaveChain(data []byte) {
	saveChain(data)
}
func (DB) LoadChain() []byte {
	return loadChain()
}
func (DB) DeleteAllBlocks() {
	emptyBlocks()
}

func getDbname() string {
	port := os.Args[2][6:]
	return fmt.Sprintf("%s_%s.db", dbName, port)
}

func InitDB() {
	if db == nil {
		// init db
		dbPointer, err := bolt.Open(getDbname(), 0600, nil)
		db = dbPointer
		utils.HandleErr(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = t.CreateBucketIfNotExists([]byte(blockBucket))
			return err
		})
		utils.HandleErr(err)
	}
}

func Close() {
	db.Close()
}

func saveBlock(hash string, data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func saveChain(data []byte) {
	err := db.Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func loadChain() []byte {
	var data []byte
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func findBlock(hash string) []byte {
	var data []byte
	db.View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blockBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}

func emptyBlocks() {
	db.Update(func(t *bolt.Tx) error {
		utils.HandleErr(t.DeleteBucket([]byte(blockBucket)))
		_, err := t.CreateBucket([]byte(blockBucket))
		utils.HandleErr(err)
		return nil
	})
}
