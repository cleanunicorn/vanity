package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"strings"

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

func main() {

	var key *ecdsa.PrivateKey
	var prefix string = "0xdeadc0de"

	for {
		key = check_prefix(prefix)
		if key != nil {
			break
		}
	}

	print(hex.EncodeToString(key.D.Bytes()), "\n")
}
