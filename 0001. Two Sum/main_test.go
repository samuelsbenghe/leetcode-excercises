package main

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{2, 7, 11, 15, 42, 23, 22, 145, 53, 11, 44, 12, 44, 1, 65, 23, 77, 23, 15, 88, 32, 11, 86, 46, 89, 53, 73}, 126)
	}
}
