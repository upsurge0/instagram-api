package main

import (
	"crypto/sha512"
	"encoding"
	"encoding/hex"
	"log"
)

func hash(input string) string {
	first := sha512.New()
	first.Write([]byte(input))

	marshaler, ok := first.(encoding.BinaryMarshaler)
	if !ok {
		log.Fatal("first does not implement encoding.BinaryMarshaler")
	}
	_, err := marshaler.MarshalBinary()
	if err != nil {
		log.Fatal("unable to marshal hash:", err)
	}
	return hex.EncodeToString(first.Sum(nil))
}

func compare(input string, hashedInput string) bool {
	return hash(input) == hashedInput
}
