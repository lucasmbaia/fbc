package main

import (
  "fmt"
  "testing"
)

func Test_New_Merkle_Tree(t *testing.T) {
  transactions := [][]byte{{'A'}, {'B'}, {'C'}}

  merkle := NewMerkleTree(transactions)

  fmt.Printf("%x", merkle.Root.Data)
}
