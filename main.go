package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func check_prefix(prefix string) *ecdsa.PrivateKey {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	if strings.HasPrefix(strings.ToLower(addr.Hex()), prefix) {
		return key
	}

	return nil
}

func check_prefix_routine(c chan *ecdsa.PrivateKey, prefix string) {
	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)

	if strings.HasPrefix(strings.ToLower(addr.Hex()), prefix) {
		c <- key
		return
	}

	c <- nil
}

func main() {

	var prefix string = "0xdeadcode"

	c := make(chan *ecdsa.PrivateKey, 10)

	for i := 0; i < 10; i++ {
		go check_prefix_routine(c, prefix)
	}

	start := time.Now()
	count := 0

	for {
		count++
		key := <-c
		if key != nil {
			print(hex.EncodeToString(key.D.Bytes()), "\n")
			break
		} else {
			go check_prefix_routine(c, prefix)
		}

		elapsed := time.Now().Sub(start)
		if (count%10000 == 0) && (int(elapsed.Seconds()) > 0) {
			print("Speed: ", count/int(elapsed.Seconds()), " keys/sec \n")
		}
	}

}
