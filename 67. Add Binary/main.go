// Given two binary strings a and b, return their sum as a binary string.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Explain the rules
	fmt.Println("Given two binary strings a and b, return their sum as a binary string.")
	// Test cases taken from leetcode problem
	fmt.Printf("Adding \"11\" and \"1\" (expects \"100\"): %v\n", addBinary("11", "1"))             // "100"
	fmt.Printf("Adding \"1010\" and \"1011\" (expects \"10101\"): %v\n", addBinary("1010", "1011")) // "10101"
	// Ask the user for input until they enter an empty string. This is to allow the user to test their own strings.
	for {
		var a, b string
		fmt.Print("Enter a binary string: ")
		fmt.Scanln(&a)
		if a == "" {
			break
		}
		fmt.Print("Enter another binary string: ")
		fmt.Scanln(&b)
		fmt.Printf("Adding \"%s\" and \"%s\": %v\n", a, b, addBinary(a, b))
	}
}

func addBinary(a string, b string) string {
	binA, err := strconv.ParseInt(a, 2, 32)
	if err != nil {
		panic(err)
	}
	binB, err := strconv.ParseInt(b, 2, 32)
	if err != nil {
		panic(err)
	}
	result := binA + binB
	return strconv.FormatInt(result, 2)
}
