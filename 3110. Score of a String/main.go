// You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters. Return the score of s.
package main

import (
	"fmt"
)

func main() {
	// Explain the rules
	fmt.Println("The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.")
	// Test cases taken from leetcode problem
	fmt.Printf("Score of \"hello\" (expects 13): %d\n", scoreOfString("hello")) // 13
	fmt.Printf("Score of \"zaz\" (expects 13): %d\n", scoreOfString("zaz"))     // 50
	// Ask the user for input until they enter an empty string. This is to allow the user to test their own strings.
	for {
		var s string
		fmt.Print("Enter a string: ")
		fmt.Scanln(&s)
		if s == "" {
			break
		}
		fmt.Printf("Score of \"%s\": %d\n", s, scoreOfString(s))
	}
}

func scoreOfString(s string) int {
	score := 0

	// Iterate through the string
	for i := 0; i < len(s)-1; i++ {
		score += turnNegativeToPositive(int(rune(s[i]) - rune(s[i+1])))
	}

	return score
}

func turnNegativeToPositive(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
