package main

import (
  "fmt"
  "crypto/sha256"
  "encoding/hex"
  //"bytes"
  "encoding/binary"
  "strings"
  "unsafe"
  "strconv"
  "math/big"
)

func main() {
  header := "01000000" + "81cd02ab7e569e8bcd9317e2fe99f2de44d49ab2b8851ba4a308000000000000" + "e320b6c2fffc8d750423db8b1eb942ae710e951ed797f7affc8892b0f1fc122b" + "c7f5d74d" + "f2b9441a" + "42a14695"

  str, _ := hex.DecodeString(header)
  //fmt.Println(str)

  firstRound := sha256.Sum256(str)
  secondRound := sha256.Sum256(firstRound[:])
  hash := secondRound[:]

  //teste := HashToBig(&secondRound)

  if HashToBig(&secondRound).Cmp(UnsignegNumber(440711666)) <= 0 {
    fmt.Println("DEU CERTO")
  }

  fmt.Printf("%x\n", HashToBig(&secondRound))
  fmt.Println(hash)
  fmt.Printf("%x\n", hash)

  //fmt.Printf("<%s>\n", fmtBits(hash))
  //fmt.Printf("%b\n", 255)

  //fmt.Println(hash[:1])
  for i := len(hash) - 1; i > 1; i-- {
    fmt.Printf("%x\n", hash[i-1:i])
  }

  fmt.Printf("%x\n", hash[:1])
  var hashString []string
 // buf := make([]byte, len(hash))

  for i := 0; i < len(hash); i++ {
    //str := fmt.Sprintf("%x", binary.LittleEndian.Uint16(hash[i:]))

    //if str == "0" {
      //str = "0000"
    //}

    //fmt.Printf("%x\n", *(*uint16)(unsafe.Pointer(&hash[i])))
    //hashString = append(hashString, str)
    //buf = append(buf, []byte(fmt.Sprintf("%x", binary.LittleEndian.Uint16(hash[2:]))))
    //binary.LittleEndian.Uint16(hash[i:i+1])
    //fmt.Println(hash[i:i+1])
  }

  //t := make([]byte, 2)
  //l := *(*uint16)(unsafe.Pointer(&hash[0]))
  //l := binary.LittleEndian.Uint16(hash[2:])

  //fmt.Println(len(hash))
  //binary.BigEndian.PutUint16(t, l)
  //fmt.Printf("%x\n", l)
  //fmt.Printf("%x\n", t)
  //fmt.Printf("%x\n", binary.LittleEndian.Uint16(hash[2:]))
  //fmt.Println(reflect.TypeOf(fmt.Sprintf("%x", binary.LittleEndian.Uint16(hash[2:]))))
  fmt.Println(hashString)
  fmt.Println(string(littleAnd(hash)))
  //p, _ := hex.DecodeString(strconv.FormatInt(2504433986, 16))
  p, _ := hex.DecodeString(strconv.FormatInt(1, 16))
  fmt.Println(string(littleAnd(p)))

  target := big.NewInt(2504433986)
  //target.Lsh(target, uint(32))
  fmt.Printf("%x\n", target)
  fmt.Println("01000000")

  var bits uint64 = 486604799
  var exp uint
  var mant uint64
  var pepe uint64

  exp = uint(bits >> 24)
  mant = bits & 0xffffff
  pepe = mant * (1 << (8 * (exp - 3)))
  //exp := 486604799 >> 24
  //mant := 486604799 & 0xffffff
  //d := 8 * (exp - 3)
  //pepe = float64(mant * (1 << (8 * (29 -3))))
  //target_hex := fmt.Sprintf("%064X", strconv.Itoa(pepe))

  fmt.Println(exp, mant, pepe)
  bigOne := big.NewInt(1)
  oneLsh256 := new(big.Int).Lsh(bigOne, 256)

  un := UnsignegNumber(486604799)

  /*if un.Sign() <= 0 {
    return big.NewInt(0)
  }*/

  d := new(big.Int).Add(un, bigOne)
  pp := new(big.Int).Div(oneLsh256, d)

  fmt.Printf("%x\n", pp.Int64())
  fmt.Println(fmt.Printf("%064X", pp))
  //fmt.Println(hex.DecodeString(target_hex))
  //fmt.Println(binary.LittleEndian.Uint64(hash))
  //end := fmt.Sprintf("%x", hash)
}

type Hash [32]byte

func HashToBig(hash *[32]byte) *big.Int {
  buf := *hash
  blen := len(buf)
  for i := 0; i < blen/2; i++ {
    buf[i], buf[blen-1-i] = buf[blen-1-i], buf[i]
  }

  return new(big.Int).SetBytes(buf[:])
}

func littleAnd(b []byte) []byte {
  var lt  []string
  var ilt []string
  //buf := make([]byte, len(b)/10)

  if len(b)%2 == 0 {
    for i := 0; i < len(b); i+=2 {
      str := fmt.Sprintf("%x", binary.LittleEndian.Uint16(b[i:]))

      if str == "0" {
	str = "0000"
      }

      lt = append(lt, str)
      /*tBuf := make([]byte, 2)
      le := binary.LittleEndian.Uint16(b[])
      le := *(*uint16)(unsafe.Pointer(&b[i]))
      fmt.Printf("%x\n", le)
      binary.BigEndian.PutUint16(tBuf, le)

      buf = append(buf, tBuf...)*/
    }
  } else {
    for i := 0; i < len(b); i+=2 {
      str := fmt.Sprintf("%x", *(*uint16)(unsafe.Pointer(&b[i])))
      fmt.Println(str)
    }
  }

  for i := len(lt) - 1; i >= 0; i-- {
    ilt = append(ilt, lt[i])
  }

  return []byte(strings.Join(ilt, ""))
}

func UnsignegNumber(n uint32) *big.Int {
  mant := n & 0x007fffff
  exp := uint(n >> 24)
  negative := n&0x00800000 != 0

  var bn *big.Int

  if exp <= 3 {
    mant >>= 8 * (3 - exp)
    bn = big.NewInt(int64(mant))
  } else {
    bn = big.NewInt(int64(mant))
    bn.Lsh(bn, 8*(exp - 3))
  }

  if negative {
    bn = bn.Neg(bn)
  }

  fmt.Println(fmt.Printf("%064X", bn))
  return bn
}
