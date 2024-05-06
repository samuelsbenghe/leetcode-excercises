package main

import (
	"fmt"
)

func main() {
	// Explain the rules
	fmt.Println("Given a roman numeral, convert it to an integer.")
	// Test cases taken from leetcode problem
	fmt.Printf("Roman numeral \"III\" (expects 3): %d\n", romanToInt("III"))            // 3
	fmt.Printf("Roman numeral \"LVIII\" (expects 58): %d\n", romanToInt("LVIII"))       // 58
	fmt.Printf("Roman numeral \"MCMXCIV\" (expects 1994): %d\n", romanToInt("MCMXCIV")) // 1994
	// Ask the user for input until they enter an empty string. This is to allow the user to test their own roman numerals.
	for {
		var s string
		fmt.Print("Enter a roman numeral: ")
		fmt.Scanln(&s)
		if s == "" {
			break
		}
		fmt.Printf("Integer value of \"%s\": %d\n", s, romanToInt(s))
	}
}

func romanToInt(s string) int {
	// Create a map of roman numerals
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanNumerals[rune(s[i])] < romanNumerals[rune(s[i+1])] {
			result -= romanNumerals[rune(s[i])]
		} else {
			result += romanNumerals[rune(s[i])]
		}
	}
	return result
}
