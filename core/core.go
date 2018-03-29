package core

import (
  "errors"
  "time"
  "fmt"
  "strconv"
  "crypto/sha256"
  "sync"
  "log"
)

type Block struct {
  Index		int	`json:""`
  PreviousHash	string	`json:""`
  Timestamp	int	`json:""`
  Data		string	`json:""`
  Hash		string	`json:""`
}

var (
  Blockchain  []Block
  Mutex	      = &sync.Mutex{}
)

func calcHashForBlock(block Block) string {
  return CalcHash(block.Index, block.Timestamp, block.PreviousHash, block.Data)
}

func CalcHash(index, timestamp int, prevHash, data string) string {
  var hash = strconv.Itoa(index) + strconv.Itoa(timestamp) + prevHash + data

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

func NextBlock() Block {
  var (
    lastBlock = getLastestBlock()
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
    Hash:	  CalcHash(index, timestamp, lastBlock.Hash, data),
  }

  return block
}

func AddBlock(block Block) {
  if err := validateNewBlock(block, getLastestBlock()); err != nil {
    log.Println("Error to add block: ", err)
    return
  }

  Mutex.Lock()
  Blockchain = append(Blockchain, block)
  Mutex.Unlock()

  return
}

func validateNewBlock(newBlock, lastBlock Block) error {
  if lastBlock.Index + 1 != newBlock.Index {
    return errors.New("Invalid Index!")
  } else if lastBlock.Hash != newBlock.PreviousHash {
    return errors.New("Invalid PreviousHash!")
  } else if calcHashForBlock(newBlock) != newBlock.Hash {
    return errors.New("Invalid Hash: " + calcHashForBlock(newBlock) + " # " + newBlock.Hash)
  }

  return nil
}

func getLastestBlock() Block {
  return Blockchain[len(Blockchain) - 1]
}
