package main

import (
	"bytes"
	"strconv"
)

func IntToHex(x int64) []byte {
	return []byte(strconv.FormatInt(x, 16))
}

func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualCheckSum := pubKeyHash[(len(pubKeyHash) - addressChecksumLen):]
	version := pubKeyHash[0]
	pubKeyHash = pubKeyHash[1:(len(pubKeyHash) - addressChecksumLen)]
	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Equal(targetChecksum, actualCheckSum)
}
