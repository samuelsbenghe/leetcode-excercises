// Source: https://leetcode.com/problems/palindrome-number/
// Given an integer x, return true if x is a palindrome, and false otherwise.
package main

import (
	"fmt"
)

func main() {
	// Explain the rules
	fmt.Println("Given an integer x, return true if x is a palindrome, and false otherwise.")
	// Test cases taken from leetcode problem
	fmt.Printf("Is 121 a palindrome? %v\n", isPalindrome(121))   // true
	fmt.Printf("Is -121 a palindrome? %v\n", isPalindrome(-121)) // false
	fmt.Printf("Is 10 a palindrome? %v\n", isPalindrome(10))     // false
	// Ask the user for input until they enter an empty string. This is to allow the user to test their own numbers.
	for {
		fmt.Print("Enter an integer: ")
		var input int
		fmt.Scanln(&input)
		fmt.Printf("Is %d a palindrome? %v\n", input, isPalindrome(input))
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original, reversed := x, 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}
