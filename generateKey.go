package main

import (
	"fmt"
	"math/rand"
	"crypto/sha256"
	"time"
)



func main() {
		rand.Seed(time.Now().UnixNano())
        // block := createRandomBlock()
        // fmt.Printf("%x\n", block)

        // fmt.Printf("%x\n", sha256.Sum256(block))

        privateKey := generatePrivateKey()
        fmt.Printf("%x\n", privateKey[0])

        publicKey := generatePublicKey(privateKey)

        // fmt.Printf("%x\n", sha256.Sum256(key[0][0]))

        fmt.Printf("%x\n", publicKey[0])
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

func generatePrivateKey() [256][2] []byte {
	var key [256][2] []byte
	for j := 0; j < 2; j++ {
		for i := 0; i < 256; i++ {
			key[i][j] = createRandomBlock()
		}
	}
	return 	key
}

func generatePublicKey(privateKey [256][2][]byte) [256][2] []byte{
	var publicKey [256][2] []byte
	for j := 0; j < 2; j++ {
		for i := 0; i < 256; i++ {
			block := sha256.Sum256(privateKey[i][j])
			// fmt.Printf("%x\n", block)
			publicKey[i][j] = block[:]
		}
	}
	return publicKey
}
