package main

import (
  "crypto/sha256"
)

type MerkleRoot struct {
  Root *MerkleNodes
}

type MerkleNodes struct {
  Left	*MerkleNodes
  Right	*MerkleNodes
  Data	[]byte
}

func NewMerkleTree(transactions	[][]byte) *MerkleRoot {
  var (
    merkle []MerkleNodes
  )

  if len(transactions)%2 != 0 {
    transactions = append(transactions, transactions[len(transactions) - 1])
  }

  for _, transaction := range transactions {
    hash := sha256.Sum256(transaction)

    merkle = append(merkle, MerkleNodes{
      Data: hash[:],
    })
  }

  for i := 0; i < len(transactions)/2; i++ {
    var newMerkle []MerkleNodes

    for j := 0; j < len(merkle); j = j + 2 {
      left, right := merkle[j].Data, merkle[j + 1].Data
      hash := sha256.Sum256(append(left, right...))

      newMerkle = append(newMerkle, MerkleNodes{
	Left:	&merkle[j],
	Right:	&merkle[j + 1],
	Data:	hash[:],
      })
    }

    merkle = newMerkle
  }

  return &MerkleRoot{&merkle[0]}
}
