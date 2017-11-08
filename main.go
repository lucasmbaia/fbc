package main

import (
  "log"
  "github.com/lucasmbaia/fbc/core"
)

func main() {
  var (
    genesis	= core.CreateGenesis()
    blockchain	[]core.Block
    block	core.Block
    previous	= genesis
  )

  blockchain = append(blockchain, genesis)

  for i := 0; i < 20; i++ {
    block = core.NextBlock(previous)
    blockchain = append(blockchain, block)
    previous = block
  }

  log.Println(blockchain)
}
