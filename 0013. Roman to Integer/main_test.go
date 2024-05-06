package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt(RandomRomanNumeral(50))
	}
}

func RandomRomanNumeral(size int) string {
	letterOptions := []string{"I", "V", "X", "L", "C", "D", "M"}
	result := ""
	for i := 0; i < size; i++ {
		result += letterOptions[rand.Intn(len(letterOptions))]
	}
	return result
}
