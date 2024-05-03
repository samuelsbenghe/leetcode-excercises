package main

import "testing"

func BenchmarkScoreOfString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scoreOfString(RandomString(10))
	}
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[i%len(letters)]
	}
	return string(b)
}
