package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkAddBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addBinary(GenerateBinaryString(10), GenerateBinaryString(10))
	}
}

func GenerateBinaryString(n int) string {
	binaryString := "1"
	for i := 0; i < n; i++ {
		// randomly generate a 0 or 1
		randomBit := rand.Intn(2)
		binaryString += strconv.Itoa(randomBit)
	}
	return binaryString
}
