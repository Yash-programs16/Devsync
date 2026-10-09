package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {

	key := make([]byte, 32)

	_, err := rand.Read(key)

	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(key))
}
