package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

func Hash(x []byte) []byte {
	h := md5.New()
	io.WriteString(h, "seed")
	io.WriteString(h, string(x))
	return h.Sum(nil)
}

func main() {
	var next []byte
	var old [][]byte
out:
	for {
		next = Hash(next)
		for _, item := range old {
			if bytes.Equal(next, item) {
				fmt.Printf("hash: %x\n", item)
				break out
			}
		}
		old = append(old, next)
	}
	fmt.Printf("next: %x\n", next)
}
