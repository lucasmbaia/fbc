package main

import (
  "time"
  "testing"
  "log"
  "github.com/lucasmbaia/fbc/core"
)

func TestCalcHash(t *testing.T) {
  log.Println(core.CalcHash(0, int(time.Now().UTC().Unix()), "0", "Test"))
}

func TestCreateGenesis(t *testing.T) {
  log.Println(core.CreateGenesis())
}

func TestNextBlock(t *testing.T) {
  var (
    block = core.CreateGenesis()
  )

  log.Println(core.NextBlock(block))
}
