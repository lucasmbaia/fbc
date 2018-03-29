package main

import (
  "log"
  "crypto/rsa"
  "crypto/rand"
  "crypto/x509"
  "io/ioutil"
  "fmt"
  "strings"
  "crypto/sha256"
)

func main() {
  var (
    key	    *rsa.PrivateKey
    err	    error
    rawkey  []string
    sha	    string
  )

  if key, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
    log.Fatalf("Erro ao gerar chave: ", err)
  }

  ioutil.WriteFile("priv.key", x509.MarshalPKCS1PrivateKey(key), 0777)

  rawkey = []string{"0x80", string(x509.MarshalPKCS1PrivateKey(key)), "0x01"}

  sha = fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(rawkey, ""))))
  sha = fmt.Sprintf("%x", sha256.Sum256([]byte(sha)))
  log.Println(sha)
}
