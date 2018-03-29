package main

import (
  "encoding/json"
  "log"
  "github.com/lucasmbaia/fbc/core"
)

func main() {
  var (
    genesis	= core.CreateGenesis()
    block	core.Block
  )

  core.Blockchain = append(core.Blockchain, genesis)

  //for i := 0; i < 10; i++ {
  for {
    block = core.NextBlock()
    core.AddBlock(block)

    j, _ := json.Marshal(block)
    log.Println("New Block add: ", string(j))
  }
}
