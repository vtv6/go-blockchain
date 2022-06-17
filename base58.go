package main

import (
	"bytes"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}

	return result
}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	base := big.NewInt(int64(len(b58Alphabet)))

	zeroBytes := 0
	for _, b := range input {
		if b == b58Alphabet[0] {
			zeroBytes++
		} else {
			break
		}
	}

	for _, b := range input {
		index := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, base)
		result.Add(result, big.NewInt(int64(index)))
	}
	return append(bytes.Repeat([]byte{0x00}, zeroBytes), result.Bytes()...)
}
