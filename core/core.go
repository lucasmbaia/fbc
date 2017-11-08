package core

import (
  "errors"
  "time"
  "fmt"
  "strconv"
  "crypto/sha256"
)

type Block struct {
  Index		int
  PreviousHash	string
  Timestamp	int
  Data		string
  Hash		string
}

func CalcHash(index int, prevHash, data string) string {
  var hash = strconv.Itoa(index) + prevHash + data + strconv.FormatInt(time.Now().UTC().Unix(), 10)

  return fmt.Sprintf("%x", sha256.Sum256([]byte(hash)))
}

func CreateGenesis() Block {
  return Block {
    Index:	  0,
    PreviousHash: "0",
    Timestamp:	  int(time.Now().UTC().Unix()),
    Data:	  "Genesis Block",
    Hash:	  "444e29f318744b990d53e4ce66ce3d61ae8da1ab8f6d4f17776e9b02f9c0e6c3",
  }
}

func NextBlock(lastBlock Block) Block {
  var (
    index     = lastBlock.Index + 1
    timestamp = int(time.Now().UTC().Unix())
    data      = "Block " + strconv.Itoa(index)
    block     Block
  )

  block = Block{
    Index:	  index,
    PreviousHash: lastBlock.Hash,
    Timestamp:	  timestamp,
    Data:	  "Block " + strconv.Itoa(index),
    Hash:	  CalcHash(index, lastBlock.PreviousHash, data),
  }

  return block
}

func ValidateNewBlock(newBlock, lastBlock Block) error {
  if lastBlock.Index + 1 != newBlock.Index {
    return errors.New("Invalid Index!")
  } else if lastBlock.Hash != newBlock.PreviousHash {
    return errors.New("Invalid PreviousHash!")
  }

  return nil
}
