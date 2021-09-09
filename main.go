package main

import (
	"github.com/haerong22/bobbycoin/cli"
	"github.com/haerong22/bobbycoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
