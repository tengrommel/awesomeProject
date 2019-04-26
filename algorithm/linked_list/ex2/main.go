package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func main() {
	//h := hmac.New(sha1.New, key)
	//h.Write(bytes)
	//return fmt.Sprintf(hex.EncodeToString(h.Sum(nil)))
	key := []byte("123456")
	h := hmac.New(sha1.New, key)
	h.Write([]byte("aaaaaa"))
	fmt.Printf("%x\n", h.Sum(nil))
}
