package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/wmnsk/milenage"
)

func main() {
	var (
		ks  = flag.String("k", "00112233445566778899aabbccddeeff", "K in hex string")
		ops = flag.String("op", "00112233445566778899aabbccddeeff", "OP in hex string")
	)
	flag.Parse()

	k, err := hex.DecodeString(*ks)
	if err != nil {
		log.Fatalf("Invalid K \"%s\": %+v", *ks, err)
	}
	op, err := hex.DecodeString(*ops)
	if err != nil {
		log.Fatalf("Invalid OP \"%s\": %+v", *ops, err)
	}

	opc, err := milenage.ComputeOPc(k, op)
	if err != nil {
		log.Fatalf("Failed to compute: %+v", err)
	}

	fmt.Printf("%x\n", opc)
}
