package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

func check_prefix_routine(c chan *ecdsa.PrivateKey, prefix string, suffix string) {
	for {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)

		if strings.HasPrefix(strings.ToLower(addr.Hex()), prefix) && strings.HasSuffix(strings.ToLower(addr.Hex()), suffix) {
			c <- key
		} else {
			c <- nil
		}
	}
}

func main() {
	var prefix string
	fmt.Print("Prefix: ")
	fmt.Scanf("%s", &prefix)
	if !strings.HasPrefix(prefix, "0x") {
		prefix = "0x" + prefix
	}

	var suffix string
	fmt.Print("Suffix: ")
	fmt.Scanf("%s", &suffix)

	var threads int = runtime.NumCPU()
	fmt.Printf("Number of threads: [%d]", threads)
	fmt.Scanf("%d", &threads)

	c := make(chan *ecdsa.PrivateKey, threads)

	for i := 0; i < threads; i++ {
		go check_prefix_routine(c, prefix, suffix)
	}

	start := time.Now()
	count := 0

	for {
		count++
		key := <-c
		if key != nil {
			print("Private Key: ", hex.EncodeToString(key.D.Bytes()), "\n")
			print("Public Key: ", crypto.PubkeyToAddress(key.PublicKey).Hex(), "\n")
			break
		}
		elapsed := time.Now().Sub(start)
		if (count%100000 == 0) && (int(elapsed.Seconds()) > 0) {
			print("Speed: ", count/int(elapsed.Seconds()), " keys/sec, Total: ", count, "\n")
		}
	}
}
