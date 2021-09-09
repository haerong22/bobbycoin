package main

import (
	"github.com/haerong22/bobbycoin/blockchain"
	"github.com/haerong22/bobbycoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
