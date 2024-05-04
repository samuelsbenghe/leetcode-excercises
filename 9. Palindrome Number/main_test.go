package main

import (
	"math"
	"math/rand"
	"testing"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(RandomNumber(10))
	}
}

func RandomNumber(n int) int {
	min := int(math.Pow10(n - 1))
	max := int(math.Pow10(n)) - 1
	return rand.Intn(max-min+1) + min
}
