package main

import (
	"fmt"
	"math/rand"
	"crypto/sha256"
	"time"
)



func main() {
		rand.Seed(time.Now().UnixNano())
        block := createRandomBlock()
        fmt.Printf("%x\n", block)

        fmt.Printf("%x", sha256.Sum256(block))
}
// func generatePrivateKey() string {
// 	a := rand.Uint64()

//     sha := sha256.Sum256([]byte
// }

func createRandomBlock()  []byte {
 	block := make([]byte, 32)
	rand.Read(block)
	return block

}
